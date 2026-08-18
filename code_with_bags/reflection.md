# Код с багами

В данном занятии необходимо было продемонстрировать, что кажущимся правильно работающий код, на самом деле содержит в себе ошибки:

Класс банковского аккаунта:

```go
package codewithbags

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
```

Тестовый класс:

```go
package codewithbags_test

import (
	codewithbags "invisible_mechanisms_of_logic/code_with_bags"
	"testing"
)

func TestBankAccount(t *testing.T) {

	cases := []struct {
		name            string
		initBalance     float64
		depositReplyes  int
		withDrawReplyes int
	}{
		{
			"Stress test to bank account",
			100,
			10,
			15,
		},
	}

	randomFloat := codewithbags.RangeFloat{
		Min: 10,
		Max: 100,
	}

	for _, c := range cases {
		testBankAccount := codewithbags.New(c.initBalance)
		var targetBalance float64

		t.Log("Testing deposit func to bank account: ")
		for iteration := range c.depositReplyes {
			randomValue := randomFloat.Random()

			testBankAccount.Deposit(randomValue)
			targetBalance = testBankAccount.GetBalance()

			t.Logf("balance after %d deposit iteration: %f", iteration, targetBalance)
		}

		targetBalance = testBankAccount.GetBalance()
		t.Logf("balance after all deposit iterations: %f", targetBalance)

		t.Log("Testing withdraw func to bank account: ")
		for iteration := range c.withDrawReplyes {
			randomValue := randomFloat.Random()

			testBankAccount.Withdraw(randomValue)
			targetBalance = testBankAccount.GetBalance()

			t.Logf("balance after %d withdraw iteration: %f", iteration, targetBalance)
		}

		targetBalance = testBankAccount.GetBalance()
		t.Logf("balance after all withdraw iterations: %f", targetBalance)

	}
}
```

Вывод:

```bash
=== RUN   TestBankAccount
    bank_account_test.go:33: Testing deposit func to bank account:
    bank_account_test.go:40: balance after 0 deposit iteration: 180.322144
    bank_account_test.go:40: balance after 1 deposit iteration: 241.764280
    bank_account_test.go:40: balance after 2 deposit iteration: 319.666683
    bank_account_test.go:40: balance after 3 deposit iteration: 398.090380
    bank_account_test.go:40: balance after 4 deposit iteration: 486.505390
    bank_account_test.go:40: balance after 5 deposit iteration: 566.058702
    bank_account_test.go:40: balance after 6 deposit iteration: 614.448318
    bank_account_test.go:40: balance after 7 deposit iteration: 656.957576
    bank_account_test.go:40: balance after 8 deposit iteration: 688.625208
    bank_account_test.go:40: balance after 9 deposit iteration: 741.681116
    bank_account_test.go:44: balance after all deposit iterations: 741.681116
    bank_account_test.go:46: Testing withdraw func to bank account:
    bank_account_test.go:53: balance after 0 withdraw iteration: 691.789037
    bank_account_test.go:53: balance after 1 withdraw iteration: 632.894923
    bank_account_test.go:53: balance after 2 withdraw iteration: 600.277389
    bank_account_test.go:53: balance after 3 withdraw iteration: 515.722152
    bank_account_test.go:53: balance after 4 withdraw iteration: 484.289645
    bank_account_test.go:53: balance after 5 withdraw iteration: 410.178892
    bank_account_test.go:53: balance after 6 withdraw iteration: 370.598813
    bank_account_test.go:53: balance after 7 withdraw iteration: 322.138275
    bank_account_test.go:53: balance after 8 withdraw iteration: 234.298755
    bank_account_test.go:53: balance after 9 withdraw iteration: 189.851855
    bank_account_test.go:53: balance after 10 withdraw iteration: 98.015953
    bank_account_test.go:53: balance after 11 withdraw iteration: 58.502125
    bank_account_test.go:53: balance after 12 withdraw iteration: 45.119635
    bank_account_test.go:53: balance after 13 withdraw iteration: 22.538782
    bank_account_test.go:53: balance after 14 withdraw iteration: -69.393762
    bank_account_test.go:57: balance after all withdraw iterations: -69.393762
```

Как видно из приведенного теста, в конечном итоге мы изменили состояние банковского счета до отрицательного значения, что непозволительно для банков, соответсвенно наш код содержит баги.

## Исправление

```go
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

```

Теперь результаты тестов:

```bash
=== RUN   TestBankAccount
    bank_account_test.go:33: Testing deposit func to bank account:
    bank_account_test.go:44: balance after 0 deposit iteration: 190.490930
    bank_account_test.go:44: balance after 1 deposit iteration: 204.813007
    bank_account_test.go:44: balance after 2 deposit iteration: 263.690831
    bank_account_test.go:44: balance after 3 deposit iteration: 333.431729
    bank_account_test.go:44: balance after 4 deposit iteration: 362.294753
    bank_account_test.go:44: balance after 5 deposit iteration: 411.479176
    bank_account_test.go:44: balance after 6 deposit iteration: 462.636959
    bank_account_test.go:44: balance after 7 deposit iteration: 562.006817
    bank_account_test.go:44: balance after 8 deposit iteration: 584.821861
    bank_account_test.go:44: balance after 9 deposit iteration: 598.347266
    bank_account_test.go:48: balance after all deposit iterations: 598.347266
    bank_account_test.go:50: Testing withdraw func to bank account:
    bank_account_test.go:61: balance after 0 withdraw iteration: 533.672484
    bank_account_test.go:61: balance after 1 withdraw iteration: 463.394529
    bank_account_test.go:61: balance after 2 withdraw iteration: 370.449643
    bank_account_test.go:61: balance after 3 withdraw iteration: 324.912777
    bank_account_test.go:61: balance after 4 withdraw iteration: 289.911897
    bank_account_test.go:61: balance after 5 withdraw iteration: 198.987811
    bank_account_test.go:61: balance after 6 withdraw iteration: 125.481043
    bank_account_test.go:61: balance after 7 withdraw iteration: 98.588489
    bank_account_test.go:61: balance after 8 withdraw iteration: 10.182729
    bank_account_test.go:56: withdraw error: not enough balance
    bank_account_test.go:61: balance after 9 withdraw iteration: 10.182729
    bank_account_test.go:56: withdraw error: not enough balance
    bank_account_test.go:61: balance after 10 withdraw iteration: 10.182729
    bank_account_test.go:56: withdraw error: not enough balance
    bank_account_test.go:61: balance after 11 withdraw iteration: 10.182729
    bank_account_test.go:56: withdraw error: not enough balance
    bank_account_test.go:61: balance after 12 withdraw iteration: 10.182729
    bank_account_test.go:56: withdraw error: not enough balance
    bank_account_test.go:61: balance after 13 withdraw iteration: 10.182729
    bank_account_test.go:56: withdraw error: not enough balance
    bank_account_test.go:61: balance after 14 withdraw iteration: 10.182729
    bank_account_test.go:65: balance after all withdraw iterations: 10.182729
```
