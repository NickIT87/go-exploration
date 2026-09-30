package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}

// Starting create account
// Account created

// Starting add money to wallet
// Account Verified
// SecurityCode Verified
// Wallet balance added successfully
// Sending wallet credit notification
// Make ledger entry for accountId abc with txnType credit for amount 10

// Starting debit money from wallet
// Account Verified
// SecurityCode Verified
// Wallet balance is Sufficient
// Sending wallet debit notification
// Make ledger entry for accountId abc with txnType debit for amount 5
