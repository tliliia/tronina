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

func GetTickers() (tickers []TradeSymbol, err error) {
	rawurl := poloniexPublicAPIUrl + poloniexTickerPath
	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rawTicks map[string]ticker
	err = json.Unmarshal(body, &rawTicks)
	if err != nil {
		return nil, err
	}

	tickers = make([]TradeSymbol, 0, len(rawTicks))
	for sKey, t := range rawTicks {
		st := TradeSymbol{
			Symbol: sKey,
			Price:  t.Last,
		}
		tickers = append(tickers, st)
	}
	return tickers, nil
}
