package codewithbags

type FinOperations interface {
	Deposit(value float64)
	Withdraw(value float64)
	GetBalance() float64
}

type BankAccount struct {
	balance float64
}

func New(initBalance float64) *BankAccount {
	return &BankAccount{
		balance: initBalance,
	}
}

func (ba *BankAccount) Deposit(value float64) {
	ba.balance += value
}

func (ba *BankAccount) Withdraw(value float64) {
	ba.balance -= value
}

func (ba *BankAccount) GetBalance() float64 {
	return ba.balance
}
