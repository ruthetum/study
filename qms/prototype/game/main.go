package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	sig   = "entrance"
	sep   = "-"
	limit = 10
)

type Rooms struct {
	Connection map[string]int
	*sync.RWMutex
}

var (
	rooms Rooms
)

func main() {
	rooms = Rooms{
		Connection: make(map[string]int),
		RWMutex:    new(sync.RWMutex),
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/enter", enter)
	e.GET("/exit", exit)
	e.GET("/monitor", monitor)

	// Start server
	e.Logger.Fatal(e.Start(":1324"))
}

func enter(c echo.Context) error {
	token := c.QueryParam("token")

	if ok := connect(token); !ok {
		return c.String(http.StatusUnauthorized, "invalid token")
	}

	return c.String(http.StatusOK, "ok")
}

func Parse(ticket string) (roomID, userID string, sequence int64, ok bool) {
	split := strings.Split(ticket, sep)
	if len(split) != 4 {
		return "", "", 0, false
	}

	if split[3] != sig {
		return "", "", 0, false
	}

	sequence, err := strconv.ParseInt(split[2], 10, 64)
	if err != nil {
		return "", "", 0, false
	}

	return split[0], split[1], sequence, true
}

func connect(ticket string) bool {
	rooms.Lock()
	defer rooms.Unlock()

	roomID, userID, sequence, ok := Parse(ticket)
	if !ok {
		return false
	}
	log.Println(fmt.Sprintf("connect roomID=%s, userID=%s, sequence=%d", roomID, userID, sequence))

	remain, ok := rooms.Connection[roomID]
	if !ok {
		remain = 0
	}

	if remain > limit {
		return false
	}

	rooms.Connection[roomID] = remain + 1
	log.Println(fmt.Sprintf("room %s, remain %d", roomID, rooms.Connection[roomID]))

	// Amount 피드백
	// amount := sequence + limit - int64(rooms.Connection[roomID])
	// go feedback(roomID, amount, 0)
	return true
}

func exit(c echo.Context) error {
	token := c.QueryParam("token")
	disconnect(token)
	return c.String(http.StatusOK, "ok")
}

func disconnect(token string) {
	rooms.Lock()
	defer rooms.Unlock()

	roomID, userID, sequence, ok := Parse(token)
	if !ok {
		return
	}
	log.Println(fmt.Sprintf("disconnect roomID=%s, userID=%s, sequence=%d", roomID, userID, sequence))

	remain, ok := rooms.Connection[roomID]
	if !ok {
		return
	}
	if remain < 1 {
		return
	}
	rooms.Connection[roomID] = remain - 1
	log.Println(fmt.Sprintf("room %s, remain %d", roomID, rooms.Connection[roomID]))

	// Increment 피드백
	go feedback(roomID, 0, 1)
	return
}

func feedback(roomID string, amount, increment int64) {
	type Body struct {
		RoomID    string `json:"room_id"`
		Amount    int64  `json:"amount"`
		Increment int64  `json:"increment"`
	}
	body := Body{
		RoomID:    roomID,
		Amount:    amount,
		Increment: increment,
	}

	b, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}
	buf := bytes.NewBuffer(b)
	_, err = http.Post("http://localhost:1323/feedback", "application/json", buf)
	if err != nil {
		log.Fatalln(err)
	}
}

func monitor(c echo.Context) error {
	rooms.RLock()
	defer rooms.RUnlock()

	return c.JSON(http.StatusOK, rooms.Connection)
}
