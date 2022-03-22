package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
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

// GetAll - parsed data from cryptorank.io
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

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := ".secret/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		log.Fatalln(err)
	}
	return config.Client(context.Background(), tok)
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Write - saver data to google sheet
func Write(d []*Cryptorank) error {
	ctx := context.Background()
	b, err := ioutil.ReadFile(".secret/client_secrets.json")
	if err != nil {
		return fmt.Errorf("unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return fmt.Errorf("unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "1ngUptjK8GwupzyG-_5uZTP_oCMCAtJ-v8F85lO0D7lw"

	writeRange := "A1"

	var vr sheets.ValueRange

	for i := 0; i < len(d); i++ {
		tmp := []interface{}{d[i].Name, d[i].Tag, d[i].Timestrap.Format(time.RFC822)}
		vr.Values = append(vr.Values, tmp)

		_, err = srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Do()
		if err != nil {
			return fmt.Errorf("unable to retrieve data from sheet. %v", err)
		}
	}
	return nil
}
