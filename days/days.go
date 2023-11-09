package days

import (
	"fmt"
	"macodelife/weather-cli/bark"
	"strings"
	"time"

	"github.com/6tail/lunar-go/calendar"
)

var (
	Loc *time.Location
	Now TimeNow
)

func (m *Moment) remainingDays() int {
	var year int

	month := m.Month
	day := m.Day

	if m.Lunar {
		// 优先获取上一年的农历，如果非当前年份，获取当年的农历
		date := calendar.NewLunarFromYmd(Now.Year-1, int(m.Month), m.Day).GetSolar()

		if date.GetYear() < Now.Year {
			date = calendar.NewLunarFromYmd(Now.Year, int(m.Month), m.Day).GetSolar()
		}

		month = time.Month(date.GetMonth())
		day = date.GetDay()
	}

	if month < Now.Month || (month == Now.Month && day < Now.Day) {
		year = Now.Year + 1
	} else {
		year = Now.Year
	}

	tick := time.Date(year, month, day, 0, 0, 0, 0, Loc)

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

func countdown() []string {
	initTime()

	var Moments = []Moment{
		{Name: "陈双的生日🎂", Month: time.January, Day: 13},
		{Name: "王一旋的生日🎂", Month: time.January, Day: 16},
		{Name: "蒋姐的生日🎂", Month: time.June, Day: 6, Lunar: true},
		{Name: "七七的生日🎂", Month: time.July, Day: 17},
		{Name: "凯哥的生日🎂", Month: time.September, Day: 11, Lunar: true},
		{Name: "结婚纪念日💍", Month: time.September, Day: 30},
		{Name: "老戴的生日🎂", Month: time.October, Day: 12, Lunar: true},
	}

	var upcomingDays []string

	for _, m := range Moments {
		remainingDays := m.remainingDays()

		if remainingDays < 30 {
			var s string

			if remainingDays == 0 {
				s = fmt.Sprintf("今天是「%s」", m.Name)
			} else if remainingDays <= 1 {
				s = fmt.Sprintf("⚠️ 明天是「%s」，千万不要忘了哦", m.Name)
			} else {
				s = fmt.Sprintf("距离「%s」还有 %s 天", m.Name, fmt.Sprint(remainingDays))
			}

			upcomingDays = append(upcomingDays, s)
		}
	}

	return upcomingDays
}

func Push() {
	upcomingDays := countdown()

	if upcomingDays != nil {
		bark.Push(&bark.BarkParams{
			Title: "🗓️ Days Matter 🥳",
			Body:  strings.Join(upcomingDays, "\n"),
		})
	}
}
