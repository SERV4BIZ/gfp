package datetimes

import (
	"fmt"
	"time"
)

// DateTime is data struct for DateTime object
type DateTime struct {
	Year  int
	Month int
	Day   int

	Hour   int
	Minute int
	Second int
}

// DateTimeFactory is global create a new DateTime Object
func DateTimeFactory() *DateTime {
	return new(DateTime).Factory()
}

// DateTimeNew is global create a new DateTime Object
func DateTimeNew() *DateTime {
	return DateTimeFactory()
}

// Factory is create a new DateTime object
func (me *DateTime) Factory() *DateTime {
	return me.Now()
}

// New is create a new DateTime object
func (me *DateTime) New() *DateTime {
	return me.Factory()
}

// Now is set current date time
func (me *DateTime) Now() *DateTime {
	tm := time.Now()
	me.SetYear(tm.Year())
	me.SetMonth(int(tm.Month()))
	me.SetDay(tm.Day())
	me.SetHour(tm.Hour())
	me.SetMinute(tm.Minute())
	me.SetSecond(tm.Second())
	return me
}

// UTC is set current date time with UTC timzone
func (me *DateTime) UTC() *DateTime {
	tm := time.Now().UTC()
	me.SetYear(tm.Year())
	me.SetMonth(int(tm.Month()))
	me.SetDay(tm.Day())
	me.SetHour(tm.Hour())
	me.SetMinute(tm.Minute())
	me.SetSecond(tm.Second())
	return me
}

// GMT is same UTC function
func (me *DateTime) GMT() *DateTime {
	return me.UTC()
}

// SetDateTime is set date time from DateTime object
func (me *DateTime) SetDateTime(tmv *DateTime) *DateTime {
	me.SetYear(tmv.GetYear())
	me.SetMonth(tmv.GetMonth())
	me.SetDay(tmv.GetDay())
	me.SetHour(tmv.GetHour())
	me.SetMinute(tmv.GetMinute())
	me.SetSecond(tmv.GetSecond())
	return me
}

// SetValue is set date time from year month day hour minute second
func (me *DateTime) SetValue(year int, month int, day int, hour int, minute int, second int) *DateTime {
	me.SetYear(year)
	me.SetMonth(month)
	me.SetDay(day)
	me.SetHour(hour)
	me.SetMinute(minute)
	me.SetSecond(second)
	return me
}

// FromString is convert string to date time object
func (me *DateTime) FromString(val string) *DateTime {
	dt := DateTimeFromString(val)
	me.SetYear(dt.GetYear())
	me.SetMonth(dt.GetMonth())
	me.SetDay(dt.GetDay())
	me.SetHour(dt.GetHour())
	me.SetMinute(dt.GetMinute())
	me.SetSecond(dt.GetSecond())
	return me
}

// ToString is convert date time object to string
func (me *DateTime) ToString() string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", me.GetYear(), me.GetMonth(), me.GetDay(), me.GetHour(), me.GetMinute(), me.GetSecond())
}

// SetYear is set year
func (me *DateTime) SetYear(year int) *DateTime {
	me.Year = year
	return me
}

// SetMonth is set month
func (me *DateTime) SetMonth(month int) *DateTime {
	me.Month = month
	return me
}

// SetDay is set day
func (me *DateTime) SetDay(day int) *DateTime {
	me.Day = day
	return me
}

// SetHour is set hour
func (me *DateTime) SetHour(hour int) *DateTime {
	me.Hour = hour
	return me
}

// SetMinute is set minute
func (me *DateTime) SetMinute(minute int) *DateTime {
	me.Minute = minute
	return me
}

// SetSecond is set second
func (me *DateTime) SetSecond(second int) *DateTime {
	me.Second = second
	return me
}

// GetYear is get year
func (me *DateTime) GetYear() int {
	return me.Year
}

// GetMonth is get month
func (me *DateTime) GetMonth() int {
	return me.Month
}

// GetDay is get day
func (me *DateTime) GetDay() int {
	return me.Day
}

// GetHour is get hour
func (me *DateTime) GetHour() int {
	return me.Hour
}

// GetMinute is get minute
func (me *DateTime) GetMinute() int {
	return me.Minute
}

// GetSecond is get second
func (me *DateTime) GetSecond() int {
	return me.Second
}

// LastDayOfMonth is get last day of month
func (me *DateTime) LastDayOfMonth(year int, month int) int {
	tm := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	return tm.Day()
}

// AddYear is add year
func (me *DateTime) AddYear(year int) *DateTime {
	for i := 0; i < year; i++ {
		me.Year++
	}
	return me
}

// AddMonth is add month
func (me *DateTime) AddMonth(month int) *DateTime {
	for i := 0; i < month; i++ {
		me.Month++
		if me.Month > 12 {
			me.Month = 1
			me.AddYear(1)
		}
	}
	return me
}

// AddDay is add day
func (me *DateTime) AddDay(day int) *DateTime {
	maxday := me.LastDayOfMonth(me.GetYear(), me.GetMonth())
	for i := 0; i < day; i++ {
		me.Day++
		if me.Day > maxday {
			me.Day = 1

			me.AddMonth(1)
			maxday = me.LastDayOfMonth(me.GetYear(), me.GetMonth())
		}
	}
	return me
}

// AddHour is add hour
func (me *DateTime) AddHour(hour int) *DateTime {
	for i := 0; i < hour; i++ {
		me.Hour++
		if me.Hour >= 24 {
			me.Hour = 0
			me.AddDay(1)
		}
	}
	return me
}

// AddMinute is add minute
func (me *DateTime) AddMinute(minute int) *DateTime {
	for i := 0; i < minute; i++ {
		me.Minute++
		if me.Minute >= 60 {
			me.Minute = 0
			me.AddHour(1)
		}
	}
	return me
}

// AddSecond is add second
func (me *DateTime) AddSecond(second int) *DateTime {
	for i := 0; i < second; i++ {
		me.Second++
		if me.Second >= 60 {
			me.Second = 0
			me.AddMinute(1)
		}
	}
	return me
}

// SubYear is minus year
func (me *DateTime) SubYear(year int) *DateTime {
	for i := 0; i < year; i++ {
		me.Year--
	}
	return me
}

// MinusYear is same SubYear function
func (me *DateTime) MinusYear(year int) *DateTime {
	return me.SubYear(year)
}

// SubMonth is minus month
func (me *DateTime) SubMonth(month int) *DateTime {
	for i := 0; i < month; i++ {
		me.Month--
		if me.Month < 1 {
			me.Month = 12
			me.SubYear(1)
		}
	}
	return me
}

// MinusMonth is same SubMonth function
func (me *DateTime) MinusMonth(month int) *DateTime {
	return me.SubMonth(month)
}

// SubDay is minus day
func (me *DateTime) SubDay(day int) *DateTime {
	for i := 0; i < day; i++ {
		me.Day--
		if me.Day < 1 {
			me.SubMonth(1)
			maxday := me.LastDayOfMonth(me.GetYear(), me.GetMonth())
			me.Day = maxday
		}
	}
	return me
}

// MinusDay is same SubDay function
func (me *DateTime) MinusDay(day int) *DateTime {
	return me.SubDay(day)
}

// SubHour is minus hour
func (me *DateTime) SubHour(hour int) *DateTime {
	for i := 0; i < hour; i++ {
		me.Hour--
		if me.Hour < 0 {
			me.Hour = 23
			me.SubDay(1)
		}
	}
	return me
}

// MinusHour is same SubHour function
func (me *DateTime) MinusHour(hour int) *DateTime {
	return me.SubHour(hour)
}

// SubMinute is minus Minute
func (me *DateTime) SubMinute(minute int) *DateTime {
	for i := 0; i < minute; i++ {
		me.Minute--
		if me.Minute < 0 {
			me.Minute = 59
			me.SubHour(1)
		}
	}
	return me
}

// MinusMinute is same SubMinute function
func (me *DateTime) MinusMinute(minute int) *DateTime {
	return me.SubMinute(minute)
}

// SubSecond is minus second
func (me *DateTime) SubSecond(second int) *DateTime {
	for i := 0; i < second; i++ {
		me.Second--
		if me.Second < 0 {
			me.Second = 59
			me.SubMinute(1)
		}
	}
	return me
}

// MinusSecond is same SubSecond
func (me *DateTime) MinusSecond(second int) *DateTime {
	return me.SubSecond(second)
}

// AddDateTime is add date time from DateTime object
func (me *DateTime) AddDateTime(tm *DateTime) *DateTime {
	me.AddSecond(tm.GetSecond())
	me.AddMinute(tm.GetMinute())
	me.AddHour(tm.GetHour())
	me.AddDay(tm.GetDay())
	me.AddMonth(tm.GetMonth())
	me.AddYear(tm.GetYear())
	return me
}

// AddValue is add date time from year month day hour minute second
func (me *DateTime) AddValue(year int, month int, day int, hour int, minute int, second int) *DateTime {
	me.AddSecond(second)
	me.AddMinute(minute)
	me.AddHour(hour)
	me.AddDay(day)
	me.AddMonth(month)
	me.AddYear(year)
	return me
}

// SubDateTime is minus datetime from DateTime object
func (me *DateTime) SubDateTime(tm *DateTime) *DateTime {
	me.SubSecond(tm.GetSecond())
	me.SubMinute(tm.GetMinute())
	me.SubHour(tm.GetHour())
	me.SubDay(tm.GetDay())
	me.SubMonth(tm.GetMonth())
	me.SubYear(tm.GetYear())
	return me
}

// MinusDateTime is same SubDateTime function
func (me *DateTime) MinusDateTime(tm *DateTime) *DateTime {
	return me.SubDateTime(tm)
}

// SubValue is minus value from year month day hour minute second
func (me *DateTime) SubValue(year int, month int, day int, hour int, minute int, second int) *DateTime {
	me.SubSecond(second)
	me.SubMinute(minute)
	me.SubHour(hour)
	me.SubDay(day)
	me.SubMonth(month)
	me.SubYear(year)
	return me
}
