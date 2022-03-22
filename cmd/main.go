package main

import (
	"log"

	"github.com/dreddsa5dies/parsecrypto/coingecko"
	"github.com/dreddsa5dies/parsecrypto/cryptorank"
	"github.com/dreddsa5dies/parsecrypto/service"
)

func main() {
	log.Println("Starting... ok!")

	crypto, err := cryptorank.GetAll()
	if err != nil {
		log.Fatalln(err)
	}

	coin, err := coingecko.GetAll()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Getting data... ok!")

	err = service.Write(crypto, coin)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Data saved!")
}
