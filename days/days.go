package days

import (
	"fmt"
	"macodelife/jarvis/bark"
	"strings"
	"time"

	"github.com/6tail/lunar-go/calendar"
)

var (
	Loc *time.Location
	Now TimeNow
)

var Moments = []Moment{
	{Name: "陈双的生日🎂", Month: time.January, Day: 13},
	{Name: "王一旋的生日🎂", Month: time.January, Day: 16},
	{Name: "蒋姐的生日🎂", Month: time.June, Day: 6, Lunar: true},
	{Name: "七七的生日🎂", Month: time.July, Day: 17},
	{Name: "凯哥的生日🎂", Month: time.September, Day: 11, Lunar: true},
	{Name: "结婚纪念日💍", Month: time.September, Day: 30},
	{Name: "老戴的生日🎂", Month: time.October, Day: 12, Lunar: true},
	{Name: "七七的生日🎂", Month: time.July, Day: 17},
}

var Reminders = []Reminder{
	{Day: 1, Message: "月底了，记得还信用卡💳"},
}

func (m *Moment) remainingDays() int {
	month := m.Month
	day := m.Day

	if m.Lunar {
		date := calendar.NewLunarFromYmd(Now.Year-1, int(m.Month), m.Day).GetSolar()
		if date.GetYear() < Now.Year {
			date = calendar.NewLunarFromYmd(Now.Year, int(m.Month), m.Day).GetSolar()
		}
		month = time.Month(date.GetMonth())
		day = date.GetDay()
	}

	var year int
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

func checkReminders() []string {
	var reminders []string
	for _, r := range Reminders {
		if r.Day == Now.Day {
			reminders = append(reminders, r.Message)
		}
	}

	return reminders
}

func Push() {
	initTime()

	upcomingDays := countdown()
	reminders := checkReminders()

	events := append(upcomingDays, reminders...)

	if events != nil {
		bark.Push(&bark.BarkParams{
			Title: "🗓️ Days Matter 🥳",
			Body:  strings.Join(events, "\n"),
		})
	}
}
