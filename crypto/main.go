package main

import (
	"fmt"
	"sync"

	"fem.com/go/crypto/api"
)

// the main go rutine
func main() {

	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup

	for _, currency := range currencies {

		wg.Add(1)

		go func(currencyCode string) {
			getCurrency(currencyCode)
			wg.Done()
		}(currency)
	}

	wg.Wait()
}

// new go routine
func getCurrency(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}

}
