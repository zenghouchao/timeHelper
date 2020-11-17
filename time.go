package timeHelper

import "time"

// Gets today's day timestamp
func Today() (startTime, endTime int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	startTime = tTime.Unix()
	endTime = tTime.AddDate(0, 0, 1).Unix()
	return
}

// Get the timestamp of yesterday
func Yesterday() (startTime, endTime int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	endTime = tTime.Unix()
	startTime = endTime - 86400
	return
}

// Gets the tomorrow day timestamp
func Tomorrow() (startTime, endTime int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, t.Location())
	startTime = tTime.Unix()
	endTime = startTime + 86400
	return
}

// Get this week's timestamp
func Week() (startTime, endTime int64) {
	startTime = getThisMondayTime()
	endTime = startTime + (86400 * 7)
	return
}

// Gets last week's timestamp
func LastWeek() (startTime, endTime int64) {
	endTime = getThisMondayTime()
	startTime = endTime - (86400 * 7)
	return
}

// Get the next week timestamp
func NextWeek() (startTime, endTime int64) {
	mondayTime := getThisMondayTime()
	startTime = mondayTime + (86400 * 7)
	endTime = mondayTime + (86400 * 7 * 2)
	return
}

// Gets the timestamp of the beginning and end of the month
func Month() (startTime, endTime int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	startTime = tTime.Unix()
	endTime = tTime.AddDate(0, 1, 0).Unix()
	return
}

// Gets the timestamp of the beginning and end of the last month
func LastMonth() (startTime, endTime int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month()-1, 1, 0, 0, 0, 0, t.Location())
	startTime = tTime.Unix()
	endTime = tTime.AddDate(0, 1, 0).Unix()
	return
}

// Gets the timestamp of the beginning and end of the year
func Year() (startTime, endTime int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
	startTime = tTime.Unix()
	endTime = tTime.AddDate(1, 0, 0).Unix()
	return
}

// Gets the timestamp of the beginning and end of last year
func LastYear() (startTime, endTime int64) {
	t := time.Now()
	tTime := time.Date(t.Year()-1, 1, 1, 0, 0, 0, 0, t.Location())
	startTime = tTime.Unix()
	endTime = tTime.AddDate(1, 0, 0).Unix()
	return
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

// Gets a timestamp from the wee hours of a few days ago to the wee hours of this morning
func DayToNowZeroPM(day int) (int64, int64) {
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month(), t.Day()-day, 0, 0, 0, 0, time.Local)
	startTime := tTime.Unix()
	zero := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
	return startTime, zero
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

// The number of days between two time intervals
func TimeIntervalDay(startTime, endTime string) int {
	t1, err := time.ParseInLocation("2006-01-02", startTime, time.Local)
	if err != nil {
		panic(err)
	}
	t2, terr := time.ParseInLocation("2006-01-02", endTime, time.Local)
	if terr != nil {
		panic(err)
	}
	day := (t2.Unix() - t1.Unix()) / 86400
	return int(day)
}

// Get a timestamp from the last few weeks
func RecentWeeks(week int, isThisWeek bool) (startTime, endTime int64) {
	if isThisWeek {
		// Contains this week
		_, endTime = Week()
	} else {
		endTime, _ = Week()
	}
	startTime = endTime - (86400*7)*int64(week)
	return
}

// Get a timestamp for the last few months
func RecentMonths(month int, isThisMonth bool) (startTime, endTime int64) {
	start, end := Month()

	if isThisMonth {
		// Contains this month
		month -= 1
		endTime = end
	} else {
		endTime = start
	}
	t := time.Now()
	tTime := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).AddDate(0, -month, 0)
	startTime = tTime.Unix()
	return
}
