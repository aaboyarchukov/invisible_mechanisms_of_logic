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
