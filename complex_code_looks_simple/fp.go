package complexcodelookssimple

import (
	"errors"
	"math/big"
)

var (
	ErrNotEnoughBalance = errors.New("not enough balance")
	ErrInvalidValue     = errors.New("invalid value")
)

var (
	InvalidBalance = big.NewFloat(-1.0)
	EmptyBalance   = big.NewFloat(0.0)
)

type BankAccount struct {
	balance *big.Float
}

func New(initBalance *big.Float) *BankAccount {
	return &BankAccount{
		balance: initBalance,
	}
}

func (ba *BankAccount) Deposit(value *big.Float) (*BankAccount, error) {
	if value.Cmp(EmptyBalance) == -1 {
		return New(InvalidBalance), ErrInvalidValue
	}

	return New(ba.balance.Add(ba.balance, value)), nil
}

func (ba *BankAccount) Withdraw(value *big.Float) (*BankAccount, error) {
	if value.Cmp(EmptyBalance) == -1 {
		return New(InvalidBalance), ErrInvalidValue
	}

	if ba.balance.Cmp(value) == -1 {
		return New(InvalidBalance), ErrNotEnoughBalance
	}

	return New(ba.balance.Sub(ba.balance, value)), nil
}

func (ba *BankAccount) GetBalance() *big.Float {
	return ba.balance
}
