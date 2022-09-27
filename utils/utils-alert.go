package utils

import (
	"image/color"
	"strconv"

	"github.com/ncruces/zenity"

	"github.com/faschulze/srv-stockwatcherservice-go/lib/constants"
	"github.com/faschulze/srv-stockwatcherservice-go/structs"
	"github.com/faschulze/srv-stockwatcherservice-go/structs/output"
)

// SendPositionLimitNotification func
func (srv StockwatcherService) SendPositionLimitNotification(checked []structs.Position) {
	windowContent := srv.PreparePositionNotification(checked)

	zenity.Info(windowContent.Content,
		zenity.Title(windowContent.Title),
		zenity.Color(windowContent.Color),
		zenity.Height(windowContent.Height),
		zenity.Width(windowContent.Width),
		zenity.InfoIcon)
}

// SendPositionErrorNotification func
func (srv StockwatcherService) SendPositionErrorNotification(errorPosition structs.Position, errorText string) {
	zenity.Error(errorText,
		zenity.Title("Error"),
		zenity.Color(color.RGBA{255, 0, 0, 1}), //red
		zenity.Height(constants.StandardHeightPopUpWindow),
		zenity.Width(constants.StandardWidthPopUpWindows),
		zenity.ErrorIcon)
}

// PreparePositionNotification func
func (srv StockwatcherService) PreparePositionNotification(checked []structs.Position) output.PopUpWindowContent {
	notificationText := ""
	for _, pos := range checked {
		switch pos.LimitType {
		case "S":
			diff := pos.CurrentPrice - pos.LimitPrice
			notificationText += "SELL: " + pos.Name + "(" + pos.Tickersymbol + ") for " +
				strconv.FormatFloat(float64(pos.CurrentPrice), 'f', 2, 64) + "€" + " (" +
				strconv.FormatFloat(float64(pos.LimitPrice), 'f', 2, 64) + "€" + " | +" +
				strconv.FormatFloat(float64(diff), 'f', 2, 64) +
				"€)" + "\n"
		case "B":
			diff := pos.CurrentPrice - pos.LimitPrice
			notificationText += "BUY: " + pos.Name + "(" + pos.Tickersymbol + ") for " +
				strconv.FormatFloat(float64(pos.CurrentPrice), 'f', 2, 64) + "€" + " (" +
				strconv.FormatFloat(float64(pos.LimitPrice), 'f', 2, 64) + "€" + " | " +
				strconv.FormatFloat(float64(diff), 'f', 2, 64) +
				"€)" + "\n"

		}
	}
	return output.PopUpWindowContent{
		Title:   "Limit reached!",
		Content: notificationText,
		Color:   color.Black,
		Height:  constants.StandardHeightPopUpWindow,
		Width:   constants.StandardWidthPopUpWindows,
	}
}
