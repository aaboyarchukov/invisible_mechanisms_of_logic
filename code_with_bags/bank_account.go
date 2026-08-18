package codewithbags

import "errors"

var ErrNotEnoughBalance = errors.New("not enough balance")
var ErrInvalidValue = errors.New("invalid value")

type BankAccount struct {
	balance float64
}

func New(initBalance float64) *BankAccount {
	return &BankAccount{
		balance: initBalance,
	}
}

func (ba *BankAccount) Deposit(value float64) error {
	if value <= 0 {
		return ErrInvalidValue
	}

	ba.balance += value

	return nil
}

func (ba *BankAccount) Withdraw(value float64) error {
	if value <= 0 {
		return ErrInvalidValue
	}

	if ba.balance < value {
		return ErrNotEnoughBalance
	}

	ba.balance -= value

	return nil
}

func (ba *BankAccount) GetBalance() float64 {
	return ba.balance
}
