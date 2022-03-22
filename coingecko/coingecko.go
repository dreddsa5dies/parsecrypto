package coingecko

import (
	"net/http"
	"strings"
	"time"

	cg "github.com/superoo7/go-gecko/v3"
)

// CoingeckoPrice - data model
type CoingeckoPrice struct {
	Name      string
	PriceUSD  float32
	Timestrap time.Time
}

// NewCoingeckoPrice - new data
func NewCoingeckoPrice() *CoingeckoPrice {
	return &CoingeckoPrice{}
}

// GetAll - parse data from coingecko.com
func GetAll() ([]*CoingeckoPrice, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	cg := cg.NewClient(httpClient)

	list, err := cg.CoinsList()
	if err != nil {
		return nil, err
	}

	rows := make([]*CoingeckoPrice, 0)
	vc := []string{"usd"}
	for k := 0; k <= 64; k++ {
		row := NewCoingeckoPrice()
		ids := []string{(*list)[k].ID}
		sp, err := cg.SimplePrice(ids, vc)
		if err != nil {
			return nil, err
		}
		row.Name = strings.TrimLeft((*list)[k].Name, "0.5X Long")
		row.PriceUSD = (*sp)[(*list)[k].ID]["usd"]
		row.Timestrap = time.Now()
		rows = append(rows, row)
		time.Sleep(time.Millisecond * 50)
	}
	return rows, nil
}
