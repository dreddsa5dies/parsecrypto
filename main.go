package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Cryptorank struct {
	Name      string
	Teg       string
	Timestrap time.Time
}

func main() {
	webPage := "https://cryptorank.io/"

	resp, err := http.Get(webPage)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	rows := make([]Cryptorank, 0)
	data := doc.Find(".data-table__table-content")
	data.First().Find("tbody").Find("tr").Each(func(iх int, sel *goquery.Selection) {
		row := new(Cryptorank)
		sel.Find("span").Each(func(i int, selection *goquery.Selection) {
			switch i {
			case 0:
				row.Name = selection.Text()
			case 1:
				row.Teg = selection.Text()
			case 3:
				row.Timestrap = time.Now()
			}
		})
		rows = append(rows, *row)
	})
	fmt.Println(rows)
}
