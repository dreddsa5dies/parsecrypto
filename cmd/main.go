package main

import (
	"log"

	"github.com/dreddsa5dies/parsecrypto/coingecko"
)

func main() {
	log.Println("Starting... ok!")

	// data, err := cryptorank.GetAll()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	coingecko.Get()

	log.Println("Getting data... ok!")

	// err = cryptorank.Write(data)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Println("Data saved!")
}
