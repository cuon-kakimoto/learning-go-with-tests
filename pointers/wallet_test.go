package main

import "testing"
import "fmt"


func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin){
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error){
		t.Helper()
		if got == nil{
			// We've introduced t.Fatal which will stop the test if it is called. This is because we don't want to make any more assertions on the error returned if there isn't one around. Without this the test would carry on to the next step and panic because of a nil pointer.
			t.Fatal("didn't get an error but wanted one")
		}

		if got != want{
			t.Errorf("got %q, want %q", got, want)
		}
	}


	assertNoError := func(t *testing.T, got error){
		t.Helper()
		if got != nil{
			// We've introduced t.Fatal which will stop the test if it is called. This is because we don't want to make any more assertions on the error returned if there isn't one around. Without this the test would carry on to the next step and panic because of a nil pointer.
			t.Fatal("didn't get an error but wanted one")
		}
	}

	t.Run("Deposite", func(t *testing.T){
		wallet := Wallet{}

		// in our very secure wallet we don't want to expose our inner state to the rest of the world. We want to control access via methods.
		// 金額は直接触らせずにmethod経由で。
		wallet.Deposite(Bitcoin(10))

		fmt.Printf("address of balance in test is %v \n", &wallet.balance)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T){
		wallet := Wallet{balance: Bitcoin(20)}

		// in our very secure wallet we don't want to expose our inner state to the rest of the world. We want to control access via methods.
		// 金額は直接触らせずにmethod経由で。
		err := wallet.Withdraw(Bitcoin(10))

		fmt.Printf("address of balance in test is %v \n", &wallet.balance)

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T){
    startingBalance := Bitcoin(20)
    wallet := Wallet{startingBalance}
    err := wallet.Withdraw(Bitcoin(100))

    assertBalance(t, wallet, startingBalance)
		assertError(t,err, ErrInsufficientFunds)
	})

}
