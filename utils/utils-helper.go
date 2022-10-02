package utils

import "time"

// isOutsideOfBusinessHours func
func (srv StockwatcherService) isOutsideOfBusinessHours(currentTime time.Time) bool {
	weekday := currentTime.Weekday().String()
	if weekday != "Saturday" && weekday != "Sunday" {
		//if currentTime != sa && currentTime != so && currentTime >= 7:30 && cucurrentTime <= 22:30{
		return false
	}

	return true
}
