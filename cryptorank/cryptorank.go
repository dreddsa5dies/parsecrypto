package cryptorank

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Cryptorank - data model
type Cryptorank struct {
	Name      string
	Tag       string
	Timestrap time.Time
}

// NewCryptorank - new data
func NewCryptorank() *Cryptorank {
	return &Cryptorank{}
}

// GetAll - parse data from cryptorank.io
func GetAll() ([]*Cryptorank, error) {
	webPage := "https://cryptorank.io/"

	resp, err := http.Get(webPage)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	rows := make([]*Cryptorank, 0)
	// первые 3 записи
	count := 3
	data := doc.Find(".data-table__table-content")
	data.First().Find("tbody").Find("tr").Each(func(iх int, sel *goquery.Selection) {
		if count != 0 {
			row := NewCryptorank()
			sel.Find("span").Each(func(i int, selection *goquery.Selection) {
				switch i {
				case 0:
					row.Name = selection.Text()
				case 1:
					row.Tag = selection.Text()
				}
			})
			row.Timestrap = time.Now()
			rows = append(rows, row)
			count--
		}
	})

	return rows, nil
}
