package utils

import (
	"testing"
	"time"
)

func TestIsOutsideOfBusinessHours(t *testing.T) {
	srv := StockwatcherService{}
	ti := time.Date(2022, 9, 29, 18, 30, 0, 0, time.Local)
	if srv.isOutsideOfBusinessHours(ti) {
		t.Fatal("test variable is outside of business hours, but should be within business hours")
	}
}
