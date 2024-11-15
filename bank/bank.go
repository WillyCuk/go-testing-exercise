package bank

import "errors"

type Account struct {
	Balance float64
}

func NewAccount(balance float64) *Account {
	return &Account{
		Balance: balance,
	}
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount to deposit")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount to withdraw")
	}
	if amount > a.Balance {
		return errors.New("insufficient funds")

	}
	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}
