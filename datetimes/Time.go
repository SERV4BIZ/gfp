package datetimes

import (
	"fmt"
	"time"
)

// Time is data struct for Time object
type Time struct {
	Hour   int
	Minute int
	Second int
}

// TimeFactory is global create a new Time object
func TimeFactory() *Time {
	return new(Time).Factory()
}

// Factory is create a new Time object
func (me *Time) Factory() *Time {
	return me.Now()
}

// Now is set current time
func (me *Time) Now() *Time {
	tm := time.Now()
	me.SetHour(tm.Hour())
	me.SetMinute(tm.Minute())
	me.SetSecond(tm.Second())
	return me
}

// UTC is set current time with UTC time
func (me *Time) UTC() *Time {
	tm := time.Now().UTC()
	me.SetHour(tm.Hour())
	me.SetMinute(tm.Minute())
	me.SetSecond(tm.Second())
	return me
}

// GMT is set current time with GMT time
func (me *Time) GMT() *Time {
	return me.UTC()
}

// SetTime is set time from Time object
func (me *Time) SetTime(tmv *Time) *Time {
	me.SetHour(tmv.GetHour())
	me.SetMinute(tmv.GetMinute())
	me.SetSecond(tmv.GetSecond())
	return me
}

// SetValue is set time from hour minute second
func (me *Time) SetValue(hour int, minute int, second int) *Time {
	me.SetHour(hour)
	me.SetMinute(minute)
	me.SetSecond(second)
	return me
}

// FromString is convert to Time from string
func (me *Time) FromString(val string) *Time {
	tm := TimeFromString(val)
	me.SetHour(tm.GetHour())
	me.SetMinute(tm.GetMinute())
	me.SetSecond(tm.GetSecond())
	return me
}

// ToString is convert Time to string
func (me *Time) ToString() string {
	return fmt.Sprintf("%02d:%02d:%02d", me.GetHour(), me.GetMinute(), me.GetSecond())
}

// SetHour is set hour
func (me *Time) SetHour(hour int) *Time {
	me.Hour = hour
	return me
}

// SetMinute is set minute
func (me *Time) SetMinute(minute int) *Time {
	me.Minute = minute
	return me
}

// SetSecond is set second
func (me *Time) SetSecond(second int) *Time {
	me.Second = second
	return me
}

// GetHour is get hour
func (me *Time) GetHour() int {
	return me.Hour
}

// GetMinute is get minute
func (me *Time) GetMinute() int {
	return me.Minute
}

// GetSecond is get second
func (me *Time) GetSecond() int {
	return me.Second
}

// AddHour is add hour
func (me *Time) AddHour(hour int) *Time {
	for i := 0; i < hour; i++ {
		me.Hour++
	}
	return me
}

// AddMinute is add minute
func (me *Time) AddMinute(minute int) *Time {
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
func (me *Time) AddSecond(second int) *Time {
	for i := 0; i < second; i++ {
		me.Second++
		if me.Second >= 60 {
			me.Second = 0
			me.AddMinute(1)
		}
	}
	return me
}

// SubHour is minus hour
func (me *Time) SubHour(hour int) *Time {
	for i := 0; i < hour; i++ {
		me.Hour--
	}
	return me
}

// MinusHour is same SubHour function
func (me *Time) MinusHour(hour int) *Time {
	return me.SubHour(hour)
}

// SubMinute is minus minute
func (me *Time) SubMinute(minute int) *Time {
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
func (me *Time) MinusMinute(minute int) *Time {
	return me.SubMinute(minute)
}

// SubSecond is minus second
func (me *Time) SubSecond(second int) *Time {
	for i := 0; i < second; i++ {
		me.Second--
		if me.Second < 0 {
			me.Second = 59
			me.SubMinute(1)
		}
	}
	return me
}

// MinusSecond is same SubSecond function
func (me *Time) MinusSecond(second int) *Time {
	return me.SubSecond(second)
}

// AddTime is add time from time object
func (me *Time) AddTime(tm *Time) *Time {
	me.AddSecond(tm.GetSecond())
	me.AddMinute(tm.GetMinute())
	me.AddHour(tm.GetHour())
	return me
}

// AddValue is add time from hour minute second
func (me *Time) AddValue(hour int, minute int, second int) *Time {
	me.AddSecond(second)
	me.AddMinute(minute)
	me.AddHour(hour)
	return me
}

// SubTime is minus time from time object
func (me *Time) SubTime(tm *Time) *Time {
	me.SubSecond(tm.GetSecond())
	me.SubMinute(tm.GetMinute())
	me.SubHour(tm.GetHour())
	return me
}

// MinusTime is same SubTime function
func (me *Time) MinusTime(tm *Time) *Time {
	me.SubSecond(tm.GetSecond())
	me.SubMinute(tm.GetMinute())
	me.SubHour(tm.GetHour())
	return me
}

// SubValue is minus time from hour minute second
func (me *Time) SubValue(hour int, minute int, second int) *Time {
	me.SubSecond(second)
	me.SubMinute(minute)
	me.SubHour(hour)
	return me
}

// MinusValue is same SubValue function
func (me *Time) MinusValue(hour int, minute int, second int) *Time {
	return me.SubValue(hour, minute, second)
}
