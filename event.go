package simplemetrics

import "time"

type Event struct {
	Label       string
	Value       int // maybe int64
	OccuredDate time.Time
}
