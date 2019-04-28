package main

import (
	"testing"
)

func TestFormat(t *testing.T) {
	datetime := NewDateTime()
	fromFormat := "Y-m-d H:i:s"
	toFormat := "d/m/Y H:i:s"
	expected := "10/01/2018 12:03:14"

	formatted := datetime.Format(fromFormat, toFormat, "2018-01-10 12:03:14")
	if formatted != expected {
		t.Error("Format is wrong ", formatted)
	}

	toFormat = "d/M/Y H:i:s"
	expected = "10/Feb/2018 12:03:14"

	formatted = datetime.Format(fromFormat, toFormat, "2018-02-10 12:03:14")
	if formatted != expected {
		t.Error("Format is wrong ", formatted)
	}
}

func TestNew(t *testing.T) {
	datetime := NewDateTime()
	if datetime.Year != "2006" {
		t.Error("Year must be '2006' got: ", datetime.Year)
	}

	if datetime.Month != "01" {
		t.Error("Month must be '01' got: ", datetime.Month)
	}

	if datetime.Day != "02" {
		t.Error("Day must be '02' got: ", datetime.Day)
	}

	if datetime.WeekDay != "Mon" {
		t.Error("WeekDay must be 'Mon' got: ", datetime.WeekDay)
	}

	if datetime.Hours != "15" {
		t.Error("Hours must be '15' got: ", datetime.Hours)
	}

	if datetime.Minutes != "04" {
		t.Error("Month must be '04' got: ", datetime.Month)
	}

	if datetime.Seconds != "05" {
		t.Error("Seconds must be '05' got: ", datetime.Seconds)
	}

	if datetime.AmOrPm != "PM" {
		t.Error("AmOrPm must be 'PM' got: ", datetime.AmOrPm)
	}

	if datetime.Milliseconds != .000 {
		t.Error("Milliseconds must be '.000' got: ", datetime.Milliseconds)
	}

	if datetime.Microseconds != .000000 {
		t.Error("Microseconds must be '.000000' got: ", datetime.Microseconds)
	}

	if datetime.Nanoseconds != .000000000 {
		t.Error("Nanoseconds must be '.000000000' got: ", datetime.Nanoseconds)
	}

	if datetime.TimezoneOffset != "0700" {
		t.Error("TimezoneOffset must be '0700' got: ", datetime.TimezoneOffset)
	}

	if datetime.Timezone != "MST" {
		t.Error("Timezone must be 'MST' got: ", datetime.Timezone)
	}
}
