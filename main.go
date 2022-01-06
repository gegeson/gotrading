package main

import (
	"fmt"
	"main/app/models"
	"main/config"
	"main/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.Dbconnection)

	// apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	// tickerChannel := make(chan bitflyer.Ticker)
	// go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	// for ticker := range tickerChannel {
	// 	fmt.Println(ticker)
	// 	fmt.Println(ticker.GetMidPrice())
	// 	fmt.Println(ticker.DateTime())
	// }
}