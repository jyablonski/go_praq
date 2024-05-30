package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// use pointers to pass the address of the wallet
// without it, any modifications to the receiver's fields would affect only the copy, not the original instance.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// use pointers to pass the address of the wallet
func (w *Wallet) Deposit(amount Bitcoin) Bitcoin {
	w.balance += amount
	return amount
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
