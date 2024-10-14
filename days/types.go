package days

import "time"

type Moment struct {
	Name  string     `json:"name"`
	Month time.Month `json:"month"`
	Day   int        `json:"day"`
	Lunar bool       `json:"lunar"`
}

type TimeNow struct {
	Time  time.Time
	Year  int
	Month time.Month
	Day   int
}

type Reminder struct {
	Day     int    `json:"day"`
	Message string `json:"message"`
}
