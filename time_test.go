package timeHelper

import "testing"

func TestToday(t *testing.T) {
	startTime, endTime := Today()
	if endTime-startTime != 86400 {
		t.Error("timestamp error today ", endTime-startTime)
	}
}
