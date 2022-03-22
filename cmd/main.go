package main

import (
	"log"

	"github.com/dreddsa5dies/parsecrypto/service"
)

func main() {
	log.Println("Start... ok!")

	data, err := service.GetAll()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Get data... ok!")

	err = service.Write(data)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Data save... ok!")
}
