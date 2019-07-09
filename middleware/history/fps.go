package history

import (
	"github.com/nkjmsss/class_3s_project_enshu/middleware/models"
)

const (
	minInterval   = 100
	maxGoBackTime = 200
)

var lastsend int = 0

func DoSend() (bool, int) {
	now := historyData.now()
	if now.Time-lastsend < minInterval {
		return false, 0
	}

	shape := now.Right.Shape
	if !isValidShape(shape) {
		for i := historyData.checkLength() - 1; i > 0; i-- {
			if now.Time-historyData[i].Time > maxGoBackTime {
				break
			} else if isValidShape(historyData[i].Right.Shape) {
				shape = historyData[i].Right.Shape
				break
			}
		}
	}
	lastsend = now.Time
	return true, shape
}

func isValidShape(shape int) bool {
	return shape != models.UNKNOWN && shape != models.NOT_TRACKED
}
