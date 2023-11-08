package days

import (
	"fmt"
	"macodelife/weather-cli/bark"
	"strings"
	"time"
)

var Loc *time.Location

func initTimezone() {
	// "Asia/Shanghai"
	l, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("failed to load time location", err)
		return
	}
	Loc = l
}

func countdown() {
	initTimezone()

	var Moments = []Moment{
		{Name: "陈双的生日", Month: time.January, Day: 13},
		{Name: "王一旋的生日", Month: time.January, Day: 16},
		{Name: "七七的生日", Month: time.July, Day: 17},
	}

	now, yearNow, monthNow, _ := getNow()

	var upcomingDays []string

	for _, m := range Moments {
		year := yearNow

		if m.Month < monthNow {
			year = yearNow + 1
		}

		tick := time.Date(year, m.Month, m.Day, 0, 0, 0, 0, Loc)

		duration := tick.Sub(now)

		remainingDays := int(duration.Hours() / 24)

		if remainingDays < 90 {
			upcomingDays = append(upcomingDays, fmt.Sprintf("距离「%s」还有 %s 天", m.Name, fmt.Sprint(remainingDays)))
		}
	}

	if upcomingDays != nil {
		bark.Push(&bark.BarkParams{
			Title: "🗓️ Days Matter",
			Body:  strings.Join(upcomingDays, "\n"),
		})
	}
}

func getNow() (time.Time, int, time.Month, int) {
	now := time.Now().In(Loc)
	yearNow := now.Year()
	monthNow := now.Month()
	dayNow := now.Day()

	return now, yearNow, monthNow, dayNow
}

func Push() {
	countdown()
}
