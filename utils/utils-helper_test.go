package utils

import (
	"testing"
	"time"
)

func TestIsOutsideOfBusinessHours1(t *testing.T) {
	srv := StockwatcherService{}
	ti := time.Date(2022, 9, 29, 18, 30, 0, 0, time.Local)
	if srv.isOutsideOfBusinessHours(ti) {
		t.Fatal("(1) test variable is outside of business hours, but should be within business hours")
	}
}

func TestIsOutsideOfBusinessHours2(t *testing.T) {
	srv := StockwatcherService{}
	ti := time.Date(2022, 9, 29, 22, 35, 0, 0, time.Local)
	if !srv.isOutsideOfBusinessHours(ti) {
		t.Fatal("(2) test variable is inside of business hours, but should be outside of business hours")
	}
}

func TestIsOutsideOfBusinessHours3(t *testing.T) {
	srv := StockwatcherService{}
	ti := time.Date(2022, 10, 1, 18, 30, 0, 0, time.Local)
	if !srv.isOutsideOfBusinessHours(ti) {
		t.Fatal("(3) test variable is inside of business hours, but should be outside of business hours")
	}
}
