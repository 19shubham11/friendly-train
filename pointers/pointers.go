package pointers

import (
	"fmt"
	"errors"
)

const ErrInsufficientFunds = "cannot withdraw, insufficient funds"

type Bitcoin int

type Stringer interface {
	String() string
}


func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if (amount > w.balance) {
		return errors.New(ErrInsufficientFunds)
	}

	w.balance -= amount
	return nil
}