package utils

import "log"

// StockwatcherService struct
type StockwatcherService struct {
	Logger                    *log.Logger
	IsReloadLimitedModeActive bool
	IsLogPrinted              bool
}
