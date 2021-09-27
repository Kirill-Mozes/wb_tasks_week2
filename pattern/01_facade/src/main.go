package main

import (
	"fmt"
	"log"
)

//client code
func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234) //create new acciunt with name abc password 1234
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10) // id, code, amount
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5) // id, code, amount
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
