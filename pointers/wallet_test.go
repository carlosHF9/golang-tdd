package main


import (
	"testing"
	
)

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {

	t.Helper()

	got := wallet.Balance()

	if got != want {

		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {

		t.Fatal("Wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {

	t.Helper()

	if got == nil {
		t.Fatal("Got an error but didn't need one")
	}
}

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		
	})

	t.Run("withdraw", func(t *testing.T) {
		startingBalance := Bitcoin(5)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(5))
		
		assertNoError(t, err)
		assertBalance(t, wallet, 0)

		
	})

	t.Run("widthdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

}


