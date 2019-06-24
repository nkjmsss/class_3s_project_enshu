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

type hand struct {
	X     int `json:"x"`
	Y     int `json:"y"`
	Z     int `json:"z"`
	Shape int `json:"shape"`
}

type SendData struct {
	Time      int  `json:"time"`
	RightHand hand `json:",inline"`
}

type ReceiveData struct {
	Right hand `json:"right"`
	Left  hand `json:"left"`
}

func (h *hand) String() string {
	return fmt.Sprintf(
		"\nX: %d\nY: %d\nZ: %d\nShape: %d\n",
		h.X,
		h.Y,
		h.Z,
		h.Shape,
	)
}
