package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/faschulze/srv-stockwatcherservice-go/lib/constants"
	"github.com/faschulze/srv-stockwatcherservice-go/structs"
)

// StockwatcherService struct
type StockwatcherService struct {
	Logger                    *log.Logger
	IsReloadLimitedModeActive bool
	IsLogPrinted              bool
}

// RetrieveCurrentPrice func return Position Price Each in €
func (srv StockwatcherService) RetrieveCurrentPrice(s structs.Position, currentTime time.Time) (float32, error) {

	switch s.Category {
	case constants.CategoryIndex:
		if s.ReqURL == "" {
			return 0, fmt.Errorf("fatal empty request url")
		}
		return srv.getGenericStocksFondsIndexOrCryptoRet(s)
	case constants.CategoryCrypto:
		if s.ReqURL == "" {
			return 0, fmt.Errorf("fatal empty request url")
		}
		return srv.getGenericStocksFondsIndexOrCryptoRet(s)
	case constants.CategoryFonds:
		if s.ReqURL == "" {
			return 0, fmt.Errorf("fatal empty request url")
		}
		return srv.getGenericStocksFondsIndexOrCryptoRet(s)
	case constants.CategoryStock:
		if srv.IsReloadLimitedModeActive {
			if srv.isOutsideOfBusinessHours(currentTime) {
				return -1, fmt.Errorf("do not retrieve position value out of business hours for position: %s", s.Name)
			}
		}

		if s.ReqURL == "" {
			return 0, fmt.Errorf("fatal empty request url")
		}

		return srv.getGenericStocksFondsIndexOrCryptoRet(s)
	default:
		return 0, fmt.Errorf("unknown position category")
	}
}

// func getGenericStocksFondsIndexOrCryptoRet gets current price in euro
func (srv StockwatcherService) getGenericStocksFondsIndexOrCryptoRet(s structs.Position) (float32, error) {
	var client http.Client
	resp, err := client.Get(s.ReqURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	content := string(body)
	containsKeyElement := strings.Contains(content, "id=\"keyelement_kurs_update\">")
	if !containsKeyElement {
		return 0, fmt.Errorf("err: %s\n", fmt.Errorf("key element kurs update could not be found for "+s.Name))
	}

	substrContent := strings.Split(content, "id=\"keyelement_kurs_update\">")[1]
	substrContent = strings.Split(substrContent, "</div>")[0]
	substrContentIndex := strings.Index(substrContent, ",")
	runes := []rune(substrContent)
	eur := string(runes[substrContentIndex-6 : substrContentIndex])
	eur = eur[strings.Index(eur, ">")+1:]
	eur = strings.Replace(eur, ".", "", 1) // if eur > 1k, replace seperator
	cent := string(runes[substrContentIndex+1 : substrContentIndex+6])
	cent = cent[0:strings.Index(cent, "<")]
	currentPrice, err := strconv.ParseFloat(eur+"."+cent, 32)
	if err != nil {
		return 0, err
	}

	sum := float32(currentPrice)
	return sum, err

}

// CheckPositionLimitsAndSendNotification func
func (srv StockwatcherService) CheckPositionLimitsAndSendNotification(positions []structs.Position) {
	ret := []structs.Position{}
	for _, pos := range positions {
		switch pos.LimitType {
		case "B":
			if pos.CurrentPrice <= pos.LimitPrice {
				if srv.IsLogPrinted {
					srv.Logger.Printf("BUY limit reached: Position %s has limit of %v€ and is currently purchaseable for %v€\n", pos.Name, pos.LimitPrice, pos.CurrentPrice)
				}
				ret = append(ret, pos)
			}
		case "S":
			if pos.CurrentPrice >= pos.LimitPrice {
				if srv.IsLogPrinted {
					srv.Logger.Printf("SELL limit reached: Position %s has limit of %v€ and is currently sellable for %v€\n", pos.Name, pos.LimitPrice, pos.CurrentPrice)
				}
				ret = append(ret, pos)
			}
		}
	}

	if len(ret) != 0 {
		srv.SendPositionLimitNotification(ret)
	}
}

// isOutsideOfBusinessHours func
func (srv StockwatcherService) isOutsideOfBusinessHours(currentTime time.Time) bool {
	dt := time.Now().Weekday()
	weekday := dt.String()
	if weekday != "Saturday" && weekday != "Sunday" {
		//if currentTime != sa && currentTime != so && currentTime >= 7:30 && cucurrentTime <= 22:30{
		return false
	}

	return true
}
