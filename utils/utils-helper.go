package utils

import "time"

const (
	Saturday = "Saturday"
	Sunday   = "Sunday"
)

// isOutsideOfBusinessHours func
func (srv StockwatcherService) isOutsideOfBusinessHours(currentTime time.Time) bool {
	weekday := currentTime.Weekday().String()
	if weekday != Saturday && weekday != Sunday {
		openingHour := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 7, 30, 0, 0, currentTime.Location())
		closingHour := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 22, 30, 0, 0, currentTime.Location())
		if currentTime.After(openingHour) && currentTime.Before(closingHour) {
			return false
		}
	}
	return true
}
