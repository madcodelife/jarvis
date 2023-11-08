package days

import (
	"fmt"
	"macodelife/weather-cli/bark"
	"strings"
	"time"
)

var (
	Loc *time.Location
	Now TimeNow
)

func (m *Moment) remainingDays() int {
	var year int

	if m.Month < Now.Month || (m.Month == Now.Month && m.Day < Now.Day) {
		year = Now.Year + 1
	} else {
		year = Now.Year
	}

	tick := time.Date(year, m.Month, m.Day, 0, 0, 0, 0, Loc)

	duration := tick.Sub(Now.Time)

	remainingDays := int(duration.Hours() / 24)

	return remainingDays
}

func initTime() {
	// "Asia/Shanghai"
	l, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("failed to load time location", err)
		return
	}
	Loc = l

	now := time.Now().In(Loc)
	startOfNow := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, Loc)

	Now = TimeNow{
		Time:  startOfNow,
		Year:  startOfNow.Year(),
		Month: startOfNow.Month(),
		Day:   startOfNow.Day(),
	}
}

func countdown() {
	initTime()

	var Moments = []Moment{
		{Name: "陈双的生日", Month: time.January, Day: 13},
		{Name: "王一旋的生日", Month: time.January, Day: 16},
		{Name: "七七的生日", Month: time.July, Day: 17},
	}

	var upcomingDays []string

	for _, m := range Moments {
		remainingDays := m.remainingDays()

		if remainingDays < 30 {
			var s string

			if remainingDays == 0 {
				s = fmt.Sprintf("今天是「%s」🥳 ", m.Name)
			} else if remainingDays <= 1 {
				s = fmt.Sprintf("⚠️ 明天是「%s」，千万不要忘了哦", m.Name)
			} else {
				s = fmt.Sprintf("距离「%s」还有 %s 天", m.Name, fmt.Sprint(remainingDays))
			}

			upcomingDays = append(upcomingDays, s)
		}
	}

	if upcomingDays != nil {
		bark.Push(&bark.BarkParams{
			Title: "🗓️ Days Matter",
			Body:  strings.Join(upcomingDays, "\n"),
		})
	}
}

func Push() {
	countdown()
}
