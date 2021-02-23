package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"trade_package/poloniex"
)

func Save(ch chan poloniex.TradeSymbol) {
	file, err := os.OpenFile("trade.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}

	var wg sync.WaitGroup
	for {
		val, ok := <-ch
		if ok {
			wg.Add(1)
			go localSave(val, &wg, file)
		}
	}
	wg.Wait()
	fmt.Println("wg wait")
	file.Close()
}

func localSave(val poloniex.TradeSymbol, wg *sync.WaitGroup, file *os.File) {
	defer wg.Done()
	writeToFile(val, file)
}

func writeToFile(val poloniex.TradeSymbol, file *os.File) {
	jsonString, err := json.Marshal(val)
	if err != nil {
		fmt.Println("marshal err:", err)
	}
	_, err = file.WriteString(string(jsonString))
	if err != nil {
		fmt.Println("write err:", err)
	}
}
