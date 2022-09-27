package main

import (
	"flag"
	"log"
	"time"

	"github.com/faschulze/srv-stockwatcherservice-go/lib/constants"
	"github.com/faschulze/srv-stockwatcherservice-go/utils"
)

var (
	// Input parameter
	pCsvStocksInputPath  = flag.String("csv-stocks-input-path", "input-files/stock-limits.csv", "csv file path for stock limit inputs")
	pCsvCryptosInputPath = flag.String("csv-cryptos-input-path", "input-files/crypto-limits.csv", "csv file path for crypto limit inputs")
	pCsvFondsInputPath   = flag.String("csv-fonds-input-path", "input-files/fonds-limits.csv", "csv file path for fonds limit inputs")
	pCsvIndexesInputPath = flag.String("csv-indexes-input-path", "input-files/indexes-limits.csv", "csv file path for index limit inputs")

	// Console log parameter
	printPositionStructs = flag.Bool("print-positions", true, "print position structs and info output to console")

	// Reload parameter
	pReloadIntervalSeconds = flag.Int("reload-interval-seconds", 5, "can also show a report without requesting from external websites")
	pReloadLimitedMode     = flag.Bool("reload-limited-mode", true, "if false always send requests, if true reload only positions that can change during the weekends or public holidays, e.g. crypto or metals")
)

var (
	logger *log.Logger
)

func main() {

	logger = log.Default()

	srv := utils.StockwatcherService{
		Logger:                    logger,
		IsReloadLimitedModeActive: *pReloadLimitedMode,
		IsLogPrinted:              *printPositionStructs,
	}

	posMap := initPositionMap()
	if len(posMap) == 0 {
		srv.Logger.Fatal("no input files detected")
	}

	positions, err := srv.InitListPositionsFromInput(posMap)
	if err != nil {
		srv.Logger.Fatal("position list could not be initialized")
	}

	for {
		srv.Logger.Println("retrieve current price for all positions....")

		for index, position := range positions {
			price, err := srv.RetrieveCurrentPrice(position, time.Now())
			if err != nil {
				errText := "could not retrieve current price for " + position.Name + "(" + position.Tickersymbol + ")\n" + err.Error()
				srv.Logger.Printf(errText)
				srv.SendPositionErrorNotification(position, errText)
			}
			positions[index].CurrentPrice = price
			if srv.IsLogPrinted {
				curr := ""
				val := "points"
				if position.Category != constants.CategoryIndex {
					curr = "€"
					val = "price"
				}
				srv.Logger.Printf("current "+val+" for %s: %v"+curr+"\n", positions[index].Name, positions[index].CurrentPrice)
			}
		}

		srv.CheckPositionLimitsAndSendNotification(positions)

		if srv.IsLogPrinted {
			srv.Logger.Printf("wait for %v seconds\n", *pReloadIntervalSeconds)
		}
		time.Sleep(time.Duration(*pReloadIntervalSeconds) * time.Second)
	}
}

// initPositionMap func
func initPositionMap() map[string]string {
	posMap := make(map[string]string)

	if *pCsvStocksInputPath != "" {
		posMap[constants.CategoryStock] = *pCsvStocksInputPath
	}
	if *pCsvCryptosInputPath != "" {
		posMap[constants.CategoryCrypto] = *pCsvCryptosInputPath
	}
	if *pCsvIndexesInputPath != "" {
		posMap[constants.CategoryIndex] = *pCsvIndexesInputPath
	}
	if *pCsvFondsInputPath != "" {
		posMap[constants.CategoryFonds] = *pCsvFondsInputPath
	}

	return posMap
}
