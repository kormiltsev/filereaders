package app

import (
	"github.com/kormiltsev/filereaders/internal/storage"
)

type DataTypeSwitcher interface {
	Do()
}

func ReturnTimezone(date int64) (string, error) {
	a, err := storage.ParseJsonTimeZones() //StartTime int64, Timezone string
	if err != nil {
		return "", err
	}
	// sort.Slice(a.HistoryTZ, func(i, j int) bool {
	// 	return a.HistoryTZ[i].StartTime > a.HistoryTZ[j].StartTime
	// })
	var d int64
	d = 1000000000
	answer := a.CurrentTZ
	for _, v := range a.HistoryTZ {
		if v.StartTime < date {
			if v.StartTime > d {
				d = v.StartTime
				answer = v.Timezone
			}
		}
	}
	return answer, nil
}
