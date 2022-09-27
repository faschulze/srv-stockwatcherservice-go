package utils

import (
	"os"

	"github.com/gocarina/gocsv"

	"github.com/faschulze/srv-stockwatcherservice-go/lib/constants"
	"github.com/faschulze/srv-stockwatcherservice-go/structs"
	"github.com/faschulze/srv-stockwatcherservice-go/structs/input"
)

// InitListPositionsFromInput func
func (srv StockwatcherService) InitListPositionsFromInput(inputFileMap map[string]string) ([]structs.Position, error) {
	ret := []structs.Position{}
	var err error

	for key, inputFile := range inputFileMap {
		in, err := os.Open(inputFile)
		if err != nil {
			srv.Logger.Fatal("could not read inputfile for %v", err)
		}
		defer in.Close()

		switch key {
		case constants.CategoryFonds:
			positions := []input.InputFondsPosition{}
			if err := gocsv.UnmarshalFile(in, &positions); err != nil {
				srv.Logger.Fatal("could not unmarshal inputfile for crypto: %v", err)
			}
			for _, pos := range positions {
				ret = append(ret, structs.Position{
					Name:         pos.Name,
					Tickersymbol: "$" + pos.Tickersymbol,
					ReqURL:       constants.BASE_URL + "/etfs/detail/uebersicht.html?ID_NOTATION=" + pos.NotionID + "&ISIN=" + pos.ISIN,
					CurrentPrice: -1,
					PositionCategory: structs.PositionCategory{
						Category: constants.CategoryStock,
					},
					PositionLimit: structs.PositionLimit{
						LimitType:  pos.PositionLimitType,
						LimitPrice: pos.PositionLimit,
					},
				})

				if srv.IsLogPrinted {
					srv.Logger.Printf("add fonds position %s with values %+v", pos.Name, pos)
				}
			}
		case constants.CategoryIndex:
			positions := []input.InputIndexPosition{}
			if err := gocsv.UnmarshalFile(in, &positions); err != nil {
				srv.Logger.Fatal("could not unmarshal inputfile for crypto: %v", err)
			}
			for _, pos := range positions {
				ret = append(ret, structs.Position{
					Name:         pos.Name,
					Tickersymbol: "$" + pos.Tickersymbol,
					ReqURL:       constants.BASE_URL + "/indizes/" + pos.ISIN,
					CurrentPrice: -1,
					PositionCategory: structs.PositionCategory{
						Category: constants.CategoryStock,
					},
					PositionLimit: structs.PositionLimit{
						LimitType:  pos.PositionLimitType,
						LimitPrice: pos.PositionLimit,
					},
				})

				if srv.IsLogPrinted {
					srv.Logger.Printf("add index position %s with values %+v", pos.Name, pos)
				}
			}
		case constants.CategoryCrypto:
			positions := []input.InputCryptoPosition{}
			if err := gocsv.UnmarshalFile(in, &positions); err != nil {
				srv.Logger.Fatal("could not unmarshal inputfile for crypto: %v", err)
			}
			for _, pos := range positions {
				ret = append(ret, structs.Position{
					Name:         pos.Name,
					Tickersymbol: "$" + pos.Tickersymbol,
					ReqURL:       constants.BASE_URL + "/kryptowaehrungen/" + pos.NotionID,
					CurrentPrice: -1,
					PositionCategory: structs.PositionCategory{
						Category: constants.CategoryStock,
					},
					PositionLimit: structs.PositionLimit{
						LimitType:  pos.PositionLimitType,
						LimitPrice: pos.PositionLimit,
					},
				})

				if srv.IsLogPrinted {
					srv.Logger.Printf("add crypto position %s with values %+v", pos.Name, pos)
				}
			}
		case constants.CategoryStock:
			positions := []input.InputStockPosition{}
			if err := gocsv.UnmarshalFile(in, &positions); err != nil {
				srv.Logger.Fatal("could not unmarshal inputfile for stocks: %v", err)
			}
			for _, pos := range positions {
				ret = append(ret, structs.Position{
					Name:         pos.Name,
					Tickersymbol: "$" + pos.Tickersymbol,
					ReqURL:       constants.BASE_URL + "/aktien/detail/uebersicht.html?ID_NOTATION=" + pos.NotionID + "&ISIN=" + pos.ISIN,
					CurrentPrice: -1,
					PositionCategory: structs.PositionCategory{
						Category: constants.CategoryStock,
					},
					PositionLimit: structs.PositionLimit{
						LimitType:  pos.PositionLimitType,
						LimitPrice: pos.PositionLimit,
					},
				})

				if srv.IsLogPrinted {
					srv.Logger.Printf("add stock position %s with values %+v", pos.Name, pos)
				}
			}
		}
	}

	return ret, err
}
