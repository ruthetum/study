package ticket

import (
	"errors"
	"strconv"
	"strings"
)

var entranceError = errors.New("invalid entrance token")

const (
	entranceSig = "entrance"
	eSep        = "-"
)

type EntranceToken string

type Entrance struct {
	RoomID    string `json:"room_id"`
	UserID    string `json:"user_id"`
	Sequence  int64  `json:"sequence"`
	Signature string `json:"signature"`
}

func NewEntrance(roomID, userID string, sequence int64) Entrance {
	return Entrance{
		RoomID:    roomID,
		UserID:    userID,
		Sequence:  sequence,
		Signature: entranceSig,
	}
}

func (rt EntranceToken) Parse() (Entrance, error) {
	split := strings.Split(string(rt), wSep)
	if len(split) != 4 {
		return Entrance{}, entranceErr
	}

	if split[3] != waitingSig {
		return Entrance{}, entranceErr
	}

	s, e := strconv.ParseInt(split[2], 10, 64)
	if e != nil {
		return Entrance{}, entranceErr
	}

	entrance := Entrance{
		RoomID:   split[0],
		UserID:   split[1],
		Sequence: s,
	}

	return entrance, nil
}

func (e Entrance) Token() string {
	return strings.Join([]string{e.RoomID, e.UserID, strconv.FormatInt(e.Sequence, 10), e.Signature}, eSep)
}
