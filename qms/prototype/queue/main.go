package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"queue/ticket"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mitchellh/mapstructure"
	"github.com/redis/go-redis/v9"
)

// TODO redis distributed lock (redsync pkg) - HIncBy는 분산락이 필요 없음
// TTL 설정 필요
var (
	redisClient *redis.Client
	ctx         context.Context
)

const (
	queueKey      = "queue:"
	waitingField  = "waiting"
	entranceField = "entrance"
	threshold     = 10
)

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx = context.Background()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/waiting", getWaitingTicket)
	e.POST("/entrance", getEntranceTicket)
	e.POST("/feedback", feedback)
	e.GET("/monitor", monitor)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// 대기표 발권
func getWaitingTicket(c echo.Context) error {
	// 요청
	type request struct {
		RoomID string `json:"room_id"`
		UserID string `json:"user_id"`
	}
	req := new(request)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// 대기번호 조회 및 값 증가
	// Lock
	waitingNumber, err := redisClient.HIncrBy(ctx, queueKey+req.RoomID, waitingField, 1).Result()
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	// 대기번호, 진입번호가 없는 경우 초기화
	if waitingNumber == 1 {
		_, err := redisClient.HIncrBy(ctx, queueKey+req.RoomID, entranceField, 2).Result()
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
	}
	// Unlock

	// 발급한 대기번호와 진입번호의 차이가 큰 경우 알림(로그)
	entranceNumber, err := redisClient.HGet(ctx, queueKey+req.RoomID, entranceField).Int64()
	if err != nil && waitingNumber-entranceNumber > threshold {
		log.Fatalln(fmt.Sprintf("room %s is congestion. waiting number: %d, entrance number: %d", req.RoomID, waitingNumber, entranceNumber))
	}

	// 대기표 발급
	waitingTicket := ticket.NewWaiting(req.RoomID, req.UserID, waitingNumber).Token()
	return c.String(http.StatusOK, waitingTicket)
}

// 대기표 폴링 및 입장권 발급
func getEntranceTicket(c echo.Context) error {
	// 요청
	type request struct {
		Token string `json:"token"`
	}
	req := new(request)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	token, err := ticket.WaitingToken(req.Token).Parse()
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// 진입 번호 조회
	// Lock
	entranceNumber, err := redisClient.HGet(ctx, queueKey+token.RoomID, entranceField).Int64()
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	// UnLock

	// 대기번호와 진입번호 비교 후 아직 진입 시기가 아닌 경우 대기자 수 출력
	if entranceNumber <= token.Sequence {
		log.Println(fmt.Sprintf("not yet, waiting for %d", token.Sequence-entranceNumber+1))
		return c.String(http.StatusBadRequest, "not yet")
	}

	// 입장권 발급
	entranceTicket := ticket.NewEntrance(token.RoomID, token.UserID, token.Sequence).Token()
	return c.String(http.StatusOK, entranceTicket)
}

// 게임 서버 피드백, 어드민 피드백
func feedback(c echo.Context) error {
	// 요청
	type request struct {
		RoomID    string `json:"room_id"`
		Amount    int64  `json:"amount"`
		Increment int64  `json:"increment"`
	}
	req := new(request)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Lock
	entranceNumber, err := redisClient.HGet(ctx, queueKey+req.RoomID, entranceField).Int64()
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	if req.Amount > 0 {
		// 직접 설정
		if entranceNumber >= req.Amount {
			return c.String(http.StatusBadRequest, "Bad Request")
		}

		if err := redisClient.HSet(ctx, queueKey+req.RoomID, entranceField, req.Amount).Err(); err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
	} else {
		// 증감값만 설정
		if err := redisClient.HIncrBy(ctx, queueKey+req.RoomID, entranceField, req.Increment).Err(); err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
	}

	// Unlock

	return c.String(http.StatusOK, "ok")
}

// 대기자 수 조회
func monitor(c echo.Context) error {
	roomID := c.QueryParam("room-id")

	type Info struct {
		Waiting  string `json:"waiting"`
		Entrance string `json:"entrance"`
	}

	type Response struct {
		RoomID string `json:"room_id"`
		Info   Info   `json:"info"`
		Amount int64  `json:"amount"`
	}

	// 전체 조회
	if roomID == "" {
		responses := make([]Response, 0)
		var keys []string
		var cursor uint64
		for {
			result, nextCursor, err := redisClient.Scan(ctx, cursor, queueKey+"*", 100).Result()
			cursor = nextCursor
			if err != nil {
				log.Println("Scan Error", err)
				return c.String(http.StatusBadRequest, "Scan Error")
			}
			keys = append(keys, result...)
			if nextCursor == 0 {
				break
			}
		}

		for _, key := range keys {
			result, err := redisClient.HGetAll(ctx, key).Result()
			if err != nil {
				return c.String(http.StatusBadRequest, "HGetAll Error")

			}

			info := new(Info)
			err = mapstructure.Decode(result, info)
			if err != nil {
				return c.String(http.StatusBadRequest, "mapstructure Decode Error")
			}

			rID := strings.Split(key, ":")
			waiting, _ := strconv.ParseInt(info.Waiting, 10, 64)
			entrance, _ := strconv.ParseInt(info.Entrance, 10, 64)

			response := Response{
				RoomID: rID[1],
				Info:   *info,
				Amount: entrance - waiting,
			}

			responses = append(responses, response)
		}

		return c.JSON(http.StatusOK, responses)
	}

	// 특정 방 조회
	result, err := redisClient.HGetAll(ctx, queueKey+roomID).Result()
	if err != nil {
		return c.String(http.StatusBadRequest, "HGetAll Error")

	}

	info := new(Info)
	err = mapstructure.Decode(result, info)
	if err != nil {
		return c.String(http.StatusBadRequest, "mapstructure Decode Error")
	}

	waiting, _ := strconv.ParseInt(info.Waiting, 10, 64)
	entrance, _ := strconv.ParseInt(info.Entrance, 10, 64)

	res := Response{
		RoomID: roomID,
		Info:   *info,
		Amount: entrance - waiting,
	}
	return c.JSON(http.StatusOK, res)
}
