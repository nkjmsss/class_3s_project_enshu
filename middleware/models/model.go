package models

import (
	"fmt"
)

const (
	UNKNOWN = iota
	NOT_TRACKED
	OPEN
	CLOSED
	LASSO
)

const (
	_ = iota
	TAKEOFF
	LAND
)

type hand struct {
	X     int `json:"x"`
	Y     int `json:"y"`
	Z     int `json:"z"`
	Shape int `json:"shape"`
}

type SendData struct {
	Time      int  `json:"time"` // micro sec
	RightHand hand `json:",inline"`
	Command   int  `json:"command"` // 離陸:1, 着陸:2
}

type ReceiveData struct {
	Right hand `json:"right"`
	Left  hand `json:"left"`
	Time  int  `json:"-"` // milli second
}

func (h *hand) String() string {
	return fmt.Sprintf(
		"X: %d\nY: %d\nZ: %d\nShape: %d\n",
		h.X,
		h.Y,
		h.Z,
		h.Shape,
	)
}
