package models

import (
	"fmt"
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

func (d *SendData) String() string {
	return fmt.Sprintf(
		"Time: %dms\nX: %d\nY: %d\nZ: %d\nShape: %d\n",
		d.Time,
		d.RightHand.X,
		d.RightHand.Y,
		d.RightHand.Z,
		d.RightHand.Shape,
	)
}
