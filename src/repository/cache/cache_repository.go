package cache

import (
	"encoding/json"
	"log"
	"orange-currency/src/domain/currency"
	"orange-currency/src/utils/errors"
	"strings"

	"golang.org/x/net/websocket"
)

var (
	CACHE = make(map[string]*currency.Currency)
)

func init() {
	go hitbtc()
}

// DbRepository interface
type DbRepository interface {
	GetBySymbol(string) (*currency.Currency, *errors.RestErr)
	All() ([2]*currency.Currency, *errors.RestErr)
}

type dbRepository struct{}

// NewRepository returns dbRepository interface
func NewRepository() DbRepository {
	return &dbRepository{}
}

func (repo *dbRepository) GetBySymbol(symbol string) (*currency.Currency, *errors.RestErr) {
	if symbol == "" {
		return nil, errors.NewBadRequestErr("invalid symbol name")
	}
	v, ok := CACHE[strings.ToUpper(symbol)]
	if !ok {
		return nil, errors.NewNotFoundErr("symbol no found")
	}
	return v, nil
}

func (repo *dbRepository) All() ([2]*currency.Currency, *errors.RestErr) {
	var result [2]*currency.Currency
	result[0] = CACHE["BTC"]
	result[1] = CACHE["ETH"]

	return result, nil
}

func hitbtc() {
	origin := "http://localhost:3000/"
	url := "wss://api.hitbtc.com/api/3/ws/public"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	subscribe := map[string]interface{}{
		"method": "subscribe",
		"ch":     "ticker/1s/batch",
		"params": map[string][]string{
			"symbols": {"ETHBTC", "BTCUSDT"},
		},
		// we can create "id" randomly, and save in the memory the ID
		// on the server stop, unsubscribe from ticker.
		"id": 19900211,
	}

	stringified, err := json.Marshal(subscribe)
	if err != nil {
		// panic on the ERROR, since the server rely on this data cache
		panic(err)
	}

	if _, err := ws.Write(stringified); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 768)
	var n int
	var data map[string]interface{}

	btc := currency.GetNewCurrency("BTC")
	eth := currency.GetNewCurrency("ETH")
	result := &currency.Currency{}

	for {
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		if err := json.Unmarshal(msg[:n], &data); err != nil {
			log.Fatal(err)
		}
		if data["data"] != nil {
			for k, v := range data["data"].(map[string]interface{}) {
				d := v.(map[string]interface{})
				if strings.Contains(k, "ETH") {
					eth.ID = "ETH"
					eth.FullName = "Ethereum"
					eth.FeeCurrency = "BTC"
					result = eth
				} else {
					btc.ID = "BTC"
					btc.FullName = "Bitcoin"
					btc.FeeCurrency = "USD"
					result = btc
				}
				result.Ask = d["a"].(string)
				result.Bid = d["b"].(string)
				result.Last = d["c"].(string)
				result.Open = d["o"].(string)
				result.Low = d["l"].(string)
				result.High = d["h"].(string)
			}
		}

		CACHE[result.ID] = result
	}
}
