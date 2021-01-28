package datetimes

import (
	"strconv"
	"strings"
)

// TimeFromString is global convert to time object from string
func TimeFromString(val string) *Time {
	vals := strings.Split(val, ":")
	if len(vals) == 3 {
		hour, _ := strconv.ParseInt(vals[0], 10, 64)
		minute, _ := strconv.ParseInt(vals[1], 10, 64)
		second, _ := strconv.ParseInt(vals[2], 10, 64)

		tm := TimeFactory()
		tm.SetValue(int(hour), int(minute), int(second))
		return tm
	}

	tm := TimeFactory()
	tm.SetValue(0, 0, 0)
	return tm
}

// DateFromString is global convert to date object from string
func DateFromString(val string) *Date {
	vals := strings.Split(val, "-")
	if len(vals) == 3 {
		year, _ := strconv.ParseInt(vals[0], 10, 64)
		month, _ := strconv.ParseInt(vals[1], 10, 64)
		day, _ := strconv.ParseInt(vals[2], 10, 64)

		dt := DateFactory()
		dt.SetValue(int(year), int(month), int(day))
		return dt
	}
	dt := DateFactory()
	dt.SetValue(0, 0, 0)
	return dt
}

// DateTimeFromString is global convert to datetime object from string
func DateTimeFromString(val string) *DateTime {
	vals := strings.Split(strings.TrimSpace(val), " ")

	if len(vals) == 2 {
		date := DateFromString(vals[0])
		time := TimeFromString(vals[1])

		dtm := DateTimeFactory()
		dtm.SetValue(date.GetYear(), date.GetMonth(), date.GetDay(), time.GetHour(), time.GetMinute(), time.GetSecond())
		return dtm
	}

	dtm := DateTimeFactory()
	dtm.SetValue(0, 0, 0, 0, 0, 0)
	return dtm
}
