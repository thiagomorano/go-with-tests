package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Zeny(10))

		assertBalance(t, wallet, Zeny(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}

		err := wallet.Withdraw(5)

		assertNoError(t, err)
		assertBalance(t, wallet, Zeny(15))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Zeny(10)
		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(100)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Zeny) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	if got != nil {
		t.Fatal("got an error but wanted none")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("got no error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q ", got, want)
	}
}
