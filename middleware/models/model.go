package models

import (
	"fmt"
)

type Data struct {
	Time  int `json:"time"`
	X     int `json:"x"`
	Y     int `json:"y"`
	Z     int `json:"z"`
	Shape int `json:"shape"` // 手の形
}

func (d *Data) String() string {
	return fmt.Sprintf(
		"Time: %dms\nX: %d\nY: %d\nZ: %d\nShape: %d\n",
		d.Time,
		d.X,
		d.Y,
		d.Z,
		d.Shape,
	)
}
