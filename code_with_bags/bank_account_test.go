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
