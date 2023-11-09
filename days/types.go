package days

import "time"

type Moment struct {
	Name  string
	Month time.Month
	Day   int
	Lunar bool
}

type TimeNow struct {
	Time  time.Time
	Year  int
	Month time.Month
	Day   int
}
