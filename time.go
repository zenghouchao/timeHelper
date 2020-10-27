package timeHelper

import (
	"time"
)

// Gets today's day timestamp
func Today() (int64, int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	startTime := tTime.Unix()
	endTime := tTime.AddDate(0, 0, 1).Unix()
	return startTime, endTime
}

// Get the timestamp of yesterday
func Yesterday() (int64, int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	endTime := tTime.Unix()
	startTime := endTime - 86400
	return startTime, endTime
}

// Gets the tomorrow day timestamp
func Tomorrow() (int64, int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, t.Location())
	startTime := tTime.Unix()
	endTime := startTime + 86400
	return startTime, endTime
}

// Get this week's timestamp
func Week() (int64, int64) {
	startTime := getThisMondayTime()
	endTime := startTime + (86400 * 7)
	return startTime, endTime
}

// Gets last week's timestamp
func LastWeek() (int64, int64) {
	endTime := getThisMondayTime()
	startTime := endTime - (86400 * 7)
	return startTime, endTime
}

// Get the next week timestamp
func NextWeek() (int64, int64) {
	mondayTime := getThisMondayTime()
	startTime := mondayTime + (86400 * 7)
	endTime := mondayTime + (86400 * 7 * 2)
	return startTime, endTime
}

// Gets the timestamp of the beginning and end of the month
func Month() (int64, int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	startTime := tTime.Unix()
	endTime := tTime.AddDate(0, 1, 0).Unix()
	return startTime, endTime
}

// Gets the timestamp of the beginning and end of the last month
func LastMonth() (int64, int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month()-1, 1, 0, 0, 0, 0, t.Location())
	startTime := tTime.Unix()
	endTime := tTime.AddDate(0, 1, 0).Unix()
	return startTime, endTime
}

// Gets the timestamp of the beginning and end of the year
func Year() (int64, int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
	startTime := tTime.Unix()
	endTime := tTime.AddDate(1, 0, 0).Unix()
	return startTime, endTime
}

// Gets the timestamp of the beginning and end of last year
func LastYear() (int64, int64) {
	t := time.Now()
	tTime := time.Date(t.Year()-1, 1, 1, 0, 0, 0, 0, t.Location())
	startTime := tTime.Unix()
	endTime := tTime.AddDate(1, 0, 0).Unix()
	return startTime, endTime
}

// Get a Monday timestamp
func getThisMondayTime() (mondayTime int64) {
	t := time.Now()
	offset := int(time.Sunday - t.Weekday())
	if offset == 0 {
		offset = -6
	} else {
		offset += 1
	}
	monday := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	mondayTime = monday.Unix()
	return
}

// Gets the timestamp from zero a few days ago to the present
func DayToNow(day int) (int64, int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month(), t.Day()-day, 0, 0, 0, 0, time.Local)
	startTime := tTime.Unix()
	return startTime, t.Unix()
}

// What day is it today
func WhatDay() (weekDay int) {
	weekDay = int(time.Now().Weekday())
	if weekDay == 0 {
		weekDay = 7
	}
	return weekDay
}

// Get the number of days of the month
func ThisMonthDayCount() (count int) {
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	hours := tTime.AddDate(0, 1, 0).Sub(tTime).Hours()
	count = int(hours / 24)
	return
}
