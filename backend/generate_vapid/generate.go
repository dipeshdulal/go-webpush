package main

import (
	"fmt"
	"log"

	webpush "github.com/SherClockHolmes/webpush-go"
)

func main() {
	fmt.Println("Generating vapid.")
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		log.Panicf("Error occured: %v", err)
	}

	fmt.Println("Private Key: ", privateKey)
	fmt.Println("Public Key: ", publicKey)
}
