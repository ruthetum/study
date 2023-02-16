package main

import (
	"context"
	"fmt"
	"net/http"
	"queue/ticket"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

// TODO redis distributed lock (redsync pkg)
var (
	redisClient *redis.Client
	ctx         context.Context
)

const (
	waitingKey  = "queue:w"
	entranceKey = "queue:e"
)

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:16379",
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
	waitingNumber, err := redisClient.HIncrBy(ctx, waitingKey, req.RoomID, 1).Result()
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	// 대기번호, 진입번호가 없는 경우 초기화
	if waitingNumber == 1 {
		_, err := redisClient.HIncrBy(ctx, entranceKey, req.RoomID, 2).Result()
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
	}
	// Unlock

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
	entranceNumber, err := redisClient.HGet(ctx, entranceKey, token.RoomID).Int64()
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	// UnLock

	// 대기번호와 진입번호 비교 후 아직 진입 시기가 아닌 경우 대기자 수 출력
	if entranceNumber <= token.Sequence {
		return c.String(http.StatusNotFound, fmt.Sprintf("not yet, waiting for %d", token.Sequence-entranceNumber+1))
	}

	// 입장권 발급
	entranceTicket := ticket.NewEntrance(token.RoomID, token.UserID, token.Sequence).Token()
	return c.String(http.StatusOK, entranceTicket)
}

// 게임 서버 피드백

// 어드민 피드백

// 대기자 수 조회

// 트래픽 검사 후 알림
