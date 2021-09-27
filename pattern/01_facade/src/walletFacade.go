package main

//pattern facade
/*
Реализовать паттерн «фасад».
Объяснить применимость паттерна,
его плюсы и минусы,
а также реальные примеры использования данного примера на практике.
*/

import "fmt"

type walletFacade struct { //simulation of a complex system
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

func newWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &walletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &notification{},
		ledger:       &ledger{},
	}
	fmt.Println("Account created")
	return walletFacacde
}

func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID) // check acc name
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode) // check security code
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount) // top up your balance
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID) // check acc name
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode) // check security code
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount) // top down your balance
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}
