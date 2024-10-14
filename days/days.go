package days

import (
	"encoding/json"
	"fmt"
	"log"
	"macodelife/jarvis/bark"
	"macodelife/jarvis/config"
	"strings"
	"time"

	"github.com/6tail/lunar-go/calendar"
	"github.com/supabase-community/supabase-go"
)

var (
	Loc *time.Location
	Now TimeNow
)

var Moments = []Moment{}

var Reminders = []Reminder{}

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
		log.Fatalln("failed to load time location", err)
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

func fetchMoments(client *supabase.Client) {
	body, _, err := client.From("moments").Select("*", "exact", false).Execute()
	if err != nil {
		log.Fatalln("failed to fetch moments", err)
		return
	}

	var data []Moment
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalln("failed to unmarshal moments", err)
		return
	}

	Moments = data
}

func fetchReminders(client *supabase.Client) {
	body, _, err := client.From("reminders").Select("*", "exact", false).Execute()
	if err != nil {
		log.Fatalln("failed to fetch reminders", err)
		return
	}

	var data []Reminder
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalln("failed to unmarshal reminders", err)
		return
	}

	Reminders = data
}

func initClient() *supabase.Client {
	client, err := supabase.NewClient(config.SupabaseUrl, config.SupabaseKey, &supabase.ClientOptions{})
	if err != nil {
		log.Fatalln("failed to initialize supabase client", err)
		return nil
	}

	return client
}

func Push() {
	initTime()
	client := initClient()
	fetchMoments(client)
	fetchReminders(client)

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
