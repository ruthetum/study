package main

import (
	"stress/user"
	"time"
)

func main() {
	u := user.New("1")

	_ = u.GetWaitingTicket("room1")

	time.Sleep(1 * time.Second)

	ok, _ := u.Polling()
	if !ok {
		// 반복
		return
	}

	time.Sleep(1 * time.Second)

	err := u.Login()
	if err != nil {
		return
	}

	time.Sleep(1 * time.Second)

	_ = u.Logout()
}
