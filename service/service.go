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

	"github.com/dreddsa5dies/parsecrypto/coingecko"
	"github.com/dreddsa5dies/parsecrypto/cryptorank"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Write - saver data to google sheet
func Write(a []*cryptorank.Cryptorank, b []*coingecko.CoingeckoPrice) error {
	ctx := context.Background()
	c, err := ioutil.ReadFile(".secret/client_secrets.json")
	if err != nil {
		return fmt.Errorf("unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(c, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return fmt.Errorf("unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "1ngUptjK8GwupzyG-_5uZTP_oCMCAtJ-v8F85lO0D7lw"

	writeRange := "A:C"

	var vr sheets.ValueRange

	tmp := make([][]interface{}, 0)

	for i := 0; i <= len(a)-1; i++ {
		tmp = append(tmp, []interface{}{a[i].Name, a[i].Tag, a[i].Timestrap.Format(time.RFC822)})
	}

	tmp = append(tmp, []interface{}{"", "", ""})

	for i := 0; i <= len(b)-1; i++ {
		tmp = append(tmp, []interface{}{b[i].Name, b[i].PriceUSD, b[i].Timestrap.Format(time.RFC822)})
	}

	vr.Values = append(vr.Values, tmp...)

	_, err = srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve data from sheet. %v", err)
	}

	return nil
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
