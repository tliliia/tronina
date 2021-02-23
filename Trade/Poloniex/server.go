package poloniex

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	poloniexPublicAPIUrl string = "https://poloniex.com/public?command="
	poloniexTickerPath   string = "returnTicker"
)

//структура ответов на запreturnTciker и trade-streams
type TradeSymbol struct {
	Symbol string `json:"s"`
	Price  string `json:"p"`
}

type ticker struct {
	ID   int    `json:"id"`
	Last string `json:"last"`
}

func getTickers(respch chan TradeSymbol) {
	defer close(respch)

	rawurl := poloniexPublicAPIUrl + poloniexTickerPath
	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var rawTicks map[string]ticker
	err = json.Unmarshal(respBody, &rawTicks)
	if err != nil {
		return
	}
	for sKey, t := range rawTicks {
		st := TradeSymbol{
			Symbol: sKey,
			Price:  t.Last,
		}
		respch <- st
	}
}

func GetTickers() chan TradeSymbol {
	ch := make(chan TradeSymbol, 1)
	go getTickers(ch)
	return ch
}
