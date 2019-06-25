package gesture

import (
	"fmt"

	"github.com/nkjmsss/class_3s_project_enshu/middleware/models"
)

type history []*models.ReceiveData

const (
	maxLength   = 10000
	threshold   = 100000  // 10%
	maxDuration = 1000000 // 1s
)

var historyData = make(history, maxLength)

func (h *history) checkLength() int {
	for i, v := range *h {
		if v == nil {
			return i
		}
	}

	return len(*h)
}

func (h *history) push(d *models.ReceiveData) {
	l := h.checkLength()
	if l < maxLength {
		(*h)[l] = d
	} else {
		// remove oldest data and push new one
		*h = append((*h)[1:], d)
	}
}

func (h *history) now() *models.ReceiveData {
	l := h.checkLength()
	return (*h)[l-1]
}

func Log(d *models.ReceiveData) {
	historyData.push(d)
}

func DoTakeoff() bool {
	now := historyData.now()
	if now == nil {
		return false
	}
	if now.Right.Shape != models.OPEN || now.Left.Shape != models.OPEN {
		return false
	}

	var first *models.ReceiveData
	for _, v := range historyData {
		if v == nil {
			continue
		}
		if v.Right.Shape == models.OPEN && v.Left.Shape == models.OPEN { // 両手がパー
			if now.Time-v.Time < maxDuration {
				first = v
			}
		}
	}
	if first == nil {
		return false
	}

	// 両手が十分量上に動いているかの判定
	if now.Right.Y-first.Right.Y > threshold && now.Left.Y-first.Left.Y > threshold {
		return true
	}
	return false
}

func DoLand() bool {
	now := historyData.now()
	if now == nil {
		return false
	}
	fmt.Println("now: ", now.Time)
	if now.Right.Shape != models.LASSO || now.Left.Shape != models.LASSO {
		return false
	}

	var first *models.ReceiveData
	for _, v := range historyData {
		if v == nil {
			continue
		}
		if v.Right.Shape == models.LASSO && v.Left.Shape == models.LASSO { // 両手がチョキ
			if now.Time-v.Time < maxDuration {
				first = v
			}
		}
	}
	if first == nil {
		return false
	}

	fmt.Println("first: ", first.Time)

	// 両手が十分量下に動いているかの判定
	if first.Right.Y-now.Right.Y > threshold && first.Left.Y-now.Left.Y > threshold {
		return true
	}
	return false
}
