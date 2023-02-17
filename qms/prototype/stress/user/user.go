package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID             string
	RoomID         string
	WaitingTicket  string
	EntranceTicket string
	Success        bool
	Status         bool
}

func New(id string) *User {
	return &User{
		ID:             id,
		RoomID:         "",
		WaitingTicket:  "",
		EntranceTicket: "",
		Success:        false,
		Status:         false,
	}
}

func (u *User) String() string {
	return "ID:" + u.ID + ", RoomID:" + u.RoomID + ", WaitingTicket:" + u.WaitingTicket + ", EntranceTicket:" + u.EntranceTicket + ", Success:" + fmt.Sprint(u.Success) + ", Status:" + fmt.Sprint(u.Status)
}

func (u *User) GetWaitingTicket(roomID string) error {
	u.RoomID = roomID

	type request struct {
		RoomID string `json:"room_id"`
		UserID string `json:"user_id"`
	}
	req := request{
		RoomID: u.RoomID,
		UserID: u.ID,
	}

	b, err := json.Marshal(req)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	buf := bytes.NewBuffer(b)
	resp, err := http.Post("http://localhost:1323/waiting", "application/json", buf)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	u.WaitingTicket = string(data)
	fmt.Println("waiting ticket:", u.WaitingTicket)
	return nil
}

func (u *User) Polling() (bool, error) {
	type request struct {
		Token string `json:"token"`
	}
	req := request{
		Token: u.WaitingTicket,
	}

	b, err := json.Marshal(req)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	buf := bytes.NewBuffer(b)
	resp, err := http.Post("http://localhost:1323/entrance", "application/json", buf)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}

	re := string(data)
	if re == "Bad Request" {
		log.Println("Bad Request")
		return false, nil
	}

	if re == "not yet" {
		log.Println("not yet")
		return false, nil
	}

	log.Println("entrance ticket:", re)
	u.EntranceTicket = re
	return true, nil
}

func (u *User) Login() error {
	resp, err := http.Get("http://localhost:1324/enter?token=" + u.EntranceTicket)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	result := string(data)
	if result != "ok" {
		log.Println("login failed")
	}

	log.Println("login success")
	u.Success = true
	return nil
}

func (u *User) Logout() error {
	resp, err := http.Get("http://localhost:1324/exit?token=" + u.EntranceTicket)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	result := string(data)
	if result != "ok" {
		log.Println("logout failed")
	}

	log.Println("logout success")
	u.Success = true
	return nil
}
