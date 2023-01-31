package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("wallet deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		assertCorrectMessage(t, wallet, Bitcoin(10))
	})

	t.Run("wallet widthdrawal", func(t *testing.T) {
		wallet := Wallet{}

		err := wallet.Withdraw(Bitcoin(10))
		assertError(t, err)
		assertCorrectMessage(t, wallet, Bitcoin(0))

	})
}

func assertCorrectMessage(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	if wallet.Balance() != expected {
		t.Errorf("got %s want %s", wallet.Balance(), expected)
	}
}
func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but got none")
	}
}
