package main

import (
	"trade_package/poloniex"
	"trade_package/storage"
)

func main() {
	var ch chan poloniex.TradeSymbol = poloniex.GetTickers()
	storage.Save(ch)
	/*
			ticker := time.NewTicker(time.Second)
			done := make(chan bool)

			go func() {
				for {
					select {
					case <-done:
						return
					case t := <-ticker.C:
						fmt.Println("Tick at", t)

					}
				}
			}()

		time.Sleep(time.Minute)
		ticker.Stop()
		done <- true
		fmt.Println("Ticker stopped")
	*/
}
