package datetimes

import (
	"fmt"
	"time"
)

// Date is data struct of Date object
type Date struct {
	Year  int
	Month int
	Day   int
}

// DateFactory is global create a new Date object
func DateFactory() *Date {
	return new(Date).Factory()
}

// DateNew is global create a new Date object
func DateNew() *Date {
	return DateFactory()
}

// Factory is create a new Date object
func (me *Date) Factory() *Date {
	return me.Now()
}

// New is create a new Date object
func (me *Date) New() *Date {
	return me.Factory()
}

// Now is current date
func (me *Date) Now() *Date {
	tm := time.Now()
	me.SetYear(tm.Year())
	me.SetMonth(int(tm.Month()))
	me.SetDay(tm.Day())
	return me
}

// UTC is set current date with utc time
func (me *Date) UTC() *Date {
	tm := time.Now().UTC()
	me.SetYear(tm.Year())
	me.SetMonth(int(tm.Month()))
	me.SetDay(tm.Day())
	return me
}

// GMT is same UTC function
func (me *Date) GMT() *Date {
	return me.UTC()
}

// SetDate is set date with other date
func (me *Date) SetDate(tmv *Date) *Date {
	me.SetYear(tmv.GetYear())
	me.SetMonth(tmv.GetMonth())
	me.SetDay(tmv.GetDay())
	return me
}

// SetValue is set date with year month day
func (me *Date) SetValue(year int, month int, day int) *Date {
	me.SetYear(year)
	me.SetMonth(month)
	me.SetDay(day)
	return me
}

// FromString is convert string to date object in format year-month-day
func (me *Date) FromString(val string) *Date {
	dt := DateFromString(val)
	me.SetYear(dt.GetYear())
	me.SetMonth(dt.GetMonth())
	me.SetDay(dt.GetDay())
	return me
}

// ToString is convert date to string
func (me *Date) ToString() string {
	return fmt.Sprintf("%04d-%02d-%02d", me.GetYear(), me.GetMonth(), me.GetDay())
}

// SetYear is set year
func (me *Date) SetYear(year int) *Date {
	me.Year = year
	return me
}

// SetMonth is set month
func (me *Date) SetMonth(month int) *Date {
	me.Month = month
	return me
}

// SetDay is set day
func (me *Date) SetDay(day int) *Date {
	me.Day = day
	return me
}

// GetYear is get year
func (me *Date) GetYear() int {
	return me.Year
}

// GetMonth is get month
func (me *Date) GetMonth() int {
	return me.Month
}

// GetDay is get day
func (me *Date) GetDay() int {
	return me.Day
}

// LastDayOfMonth is get last day of month
func (me *Date) LastDayOfMonth(year int, month int) int {
	tm := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	return tm.Day()
}

// AddYear is add year
func (me *Date) AddYear(year int) *Date {
	for i := 0; i < year; i++ {
		me.Year++
	}
	return me
}

// AddMonth is add month
func (me *Date) AddMonth(month int) *Date {
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
func (me *Date) AddDay(day int) *Date {
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

// SubYear is minus year
func (me *Date) SubYear(year int) *Date {
	for i := 0; i < year; i++ {
		me.Year--
	}
	return me
}

// MinusYear is same SubYear
func (me *Date) MinusYear(year int) *Date {
	return me.SubYear(year)
}

// SubMonth is minus month
func (me *Date) SubMonth(month int) *Date {
	for i := 0; i < month; i++ {
		me.Month--
		if me.Month < 1 {
			me.Month = 12
			me.SubYear(1)
		}
	}
	return me
}

// MinusMonth is same SubMonth
func (me *Date) MinusMonth(month int) *Date {
	return me.SubMonth(month)
}

// SubDay is minus day
func (me *Date) SubDay(day int) *Date {
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
func (me *Date) MinusDay(day int) *Date {
	return me.SubDay(day)
}

// AddDate is add date from other date object
func (me *Date) AddDate(tm *Date) *Date {
	me.AddDay(tm.GetDay())
	me.AddMonth(tm.GetMonth())
	me.AddYear(tm.GetYear())
	return me
}

// AddValue is add date from year month day
func (me *Date) AddValue(year int, month int, day int) *Date {
	me.AddDay(day)
	me.AddMonth(month)
	me.AddYear(year)
	return me
}

// SubDate is minus date from date object
func (me *Date) SubDate(tm *Date) *Date {
	me.SubDay(tm.GetDay())
	me.SubMonth(tm.GetMonth())
	me.SubYear(tm.GetYear())
	return me
}

// MinusDate is same SubDate function
func (me *Date) MinusDate(tm *Date) *Date {
	return me.SubDate(tm)
}

// SubValue is minus value from year month day
func (me *Date) SubValue(year int, month int, day int) *Date {
	me.SubDay(day)
	me.SubMonth(month)
	me.SubYear(year)
	return me
}

// MinusValue is same SubValue function
func (me *Date) MinusValue(year int, month int, day int) *Date {
	me.SubDay(day)
	me.SubMonth(month)
	me.SubYear(year)
	return me
}
