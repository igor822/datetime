package datetime

import (
	"strings"
	"time"
)

// DateTime struct containg date pieces
type DateTime struct {
	Year           string
	Month          string
	MonthShort     string
	Day            string
	WeekDay        string
	Hours          string
	Minutes        string
	Seconds        string
	AmOrPm         string
	Milliseconds   int
	Microseconds   int32
	Nanoseconds    int64
	TimezoneOffset string
	Timezone       string
}

// ShortMonthNames variable
var ShortMonthNames [12]string

const (
	YearStr       = "Y"
	MonthStr      = "m"
	MonthShortStr = "M"
	DayStr        = "d"
	HourStr       = "H"
	MinuteStr     = "i"
	SecondStr     = "s"
	TimezoneStr   = "O"
)

var DateTimeDefault DateTime

func (dateTime *DateTime) replaceFormat(format string) string {
	if strings.Contains(format, YearStr) {
		format = strings.Replace(format, YearStr, dateTime.Year, 1)
	}

	if strings.Contains(format, MonthStr) {
		format = strings.Replace(format, MonthStr, dateTime.Month, 1)
	}

	if strings.Contains(format, MonthShortStr) {
		format = strings.Replace(format, MonthShortStr, dateTime.MonthShort, 1)
	}

	if strings.Contains(format, DayStr) {
		format = strings.Replace(format, DayStr, dateTime.Day, 1)
	}

	if strings.Contains(format, HourStr) {
		format = strings.Replace(format, HourStr, dateTime.Hours, 1)
	}

	if strings.Contains(format, MinuteStr) {
		format = strings.Replace(format, MinuteStr, dateTime.Minutes, 1)
	}

	if strings.Contains(format, SecondStr) {
		format = strings.Replace(format, SecondStr, dateTime.Seconds, 1)
	}
	return format
}

// Format receiver to format strings
func (dateTime *DateTime) Format(fromFormat string, toFormat string, datetime string) string {
	fromFormatted := dateTime.replaceFormat(fromFormat)
	parseTime, _ := time.Parse(fromFormatted, datetime)
	toFormatted := dateTime.replaceFormat(toFormat)

	return parseTime.Format(toFormatted)
}

// NewDateTime function to create a new DateTime object
func NewDateTime() DateTime {
	return DateTime{
		Year:           "2006",
		Month:          "01",
		MonthShort:     "Jan",
		Day:            "02",
		WeekDay:        "Mon",
		Hours:          "15",
		Minutes:        "04",
		Seconds:        "05",
		AmOrPm:         "PM",
		Milliseconds:   .000,
		Microseconds:   .000000,
		Nanoseconds:    .000000000,
		TimezoneOffset: "0700",
		Timezone:       "MST",
	}
}
