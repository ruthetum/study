package ticket

import (
	"errors"
	"strconv"
	"strings"
)

var entranceErr = errors.New("invalid waiting token")

const (
	waitingSig = "waiting"
	wSep       = "-"
)

type WaitingToken string

type Waiting struct {
	RoomID    string `json:"room_id"`
	UserID    string `json:"user_id"`
	Sequence  int64  `json:"sequence"`
	Signature string `json:"signature"`
}

func NewWaiting(roomID, userID string, sequence int64) Waiting {
	return Waiting{
		RoomID:    roomID,
		UserID:    userID,
		Sequence:  sequence,
		Signature: waitingSig,
	}
}

func (wt WaitingToken) Parse() (Waiting, error) {
	split := strings.Split(string(wt), wSep)
	if len(split) != 4 {
		return Waiting{}, entranceErr
	}

	if split[3] != waitingSig {
		return Waiting{}, entranceErr
	}

	s, e := strconv.ParseInt(split[2], 10, 64)
	if e != nil {
		return Waiting{}, entranceErr
	}

	w := Waiting{
		RoomID:   split[0],
		UserID:   split[1],
		Sequence: s,
	}

	return w, nil
}

func (w Waiting) Token() string {
	return strings.Join([]string{w.RoomID, w.UserID, strconv.FormatInt(w.Sequence, 10), w.Signature}, wSep)
}
