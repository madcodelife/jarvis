package days

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRemainingDays(t *testing.T) {
	Loc, _ = time.LoadLocation("Asia/Shanghai")

	now := time.Date(2023, time.January, 10, 0, 0, 0, 0, Loc)
	Now = TimeNow{
		Time:  now,
		Year:  now.Year(),
		Month: now.Month(),
		Day:   now.Day(),
	}

	m1 := Moment{Name: "", Month: time.January, Day: 9}
	assert.Equal(t, m1.remainingDays(), 364)

	m2 := Moment{Name: "", Month: time.January, Day: 10}
	assert.Equal(t, m2.remainingDays(), 0)

	m3 := Moment{Name: "", Month: time.January, Day: 11}
	assert.Equal(t, m3.remainingDays(), 1)

	m4 := Moment{Name: "", Month: time.December, Day: 1}
	assert.Equal(t, m4.remainingDays(), 325)
}
