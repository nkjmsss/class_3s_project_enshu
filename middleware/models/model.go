package models

import (
	"fmt"
)

type Data struct {
	Time  int
	X     int
	Y     int
	Z     int
	Shape int
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
