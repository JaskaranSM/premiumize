package main

import (
	"fmt"
	"log"
	"github.com/jaskaranSM/premiumize"
)

func main() {
	api_key := "<your_api_key>"
	client := premiumize.NewPremiumizeClient(api_key)
	resp, err := client.ListTransfers()
	if err != nil {
		log.Fatal(err)
	}

	for _, transfer := range resp.Transfers {
		fmt.Println(transfer.Name)
	}
}
