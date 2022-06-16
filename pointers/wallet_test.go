package main


import (
	"testing"
)




func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {

		t.Helper()

		got := wallet.Balance()

		if got != want {

			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()

		if err == nil {

			t.Error("Wanted an error but didn't get one")
		}
	}
	


	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		
	})

	t.Run("withdraw", func(t *testing.T) {
		startingBalance := Bitcoin(4)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(5))
		
		assertBalance(t, wallet, -1)

		assertError(t, err)
	})

	


	

}