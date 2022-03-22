package coingecko

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	cg "github.com/superoo7/go-gecko/v3"
)

func Get() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	cg := cg.NewClient(httpClient)

	list, err := cg.CoinsList()
	if err != nil {
		log.Fatal(err)
	}

	vc := []string{"usd"}
	for k := range *list {
		ids := []string{(*list)[k].ID}
		sp, err := cg.SimplePrice(ids, vc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("%s usd %f", strings.TrimLeft((*list)[k].Name, "0.5X Long"), (*sp)[(*list)[k].ID]["usd"]))
	}
}
