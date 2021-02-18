package wallet

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds an error that occurs when trying to withdraw more than balance
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin is a type created from an existing one (int)
type Bitcoin int

// Stringer interface is defined in `fmt` package and lets you define how the type is printed
// when used with `%s` format string in prints
type Stringer interface {
	String() string
}

// Overriding how `%s` on Bitcoin is printed
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct
type Wallet struct {
	// private variable
	balance Bitcoin
}

// Deposit public method of Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Balance public method of Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw public method of Wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
