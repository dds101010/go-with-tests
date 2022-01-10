package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

// Stringer interface in fmt
// lets you define how your type is printed when used with the %s format string in prints
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// In Go if a symbol (variables, types, functions et al) starts with a
// lowercase symbol then it is private outside the package it's defined in.
type Wallet struct {
	balance Bitcoin
}

// when you call a function or a method the arguments are copied
// func (w Wallet) Deposit(amount int) {
//  fmt.Printf("address of balance in test is %v \n", &w.balance)
// 	w.balance += amount
// }

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// The var keyword allows us to define values global to the package.
var ErrInsufficientFunds = errors.New("Insufficient balance")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
