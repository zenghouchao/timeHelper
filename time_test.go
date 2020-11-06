package timeHelper

import (
	"testing"
	"time"
)

func TestToday(t *testing.T) {
	startTime, endTime := Today()
	if endTime-startTime != 86400 {
		t.Error("timestamp error today ", endTime-startTime)
	}
}

func TestWhatDay(t *testing.T) {
	today := WhatDay()
	if today != int(time.Now().Weekday()) {
		t.Error("week error")
	}
}

func TestTimeIntervalDay(t *testing.T) {

	if TimeIntervalDay("2020-11-01", "2020-12-01") != 30 ||
		TimeIntervalDay("2020-10-01", "2020-11-01") != 31 {
		t.Error("The number of days interval was calculated incorrectly")
	}
}
