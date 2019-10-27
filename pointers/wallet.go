package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance Bitcoin
}

// 型を定義したことでメソッドを追加できる
type Bitcoin int

// HACK: なくても動くっちゃうごくんだが。
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//  "a pointer to a wallet".
func (w *Wallet) Deposite(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)

	// https://golang.org/ref/spec#Method_values
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
