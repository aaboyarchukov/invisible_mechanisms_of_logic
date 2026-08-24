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

## Покрытие тестами

При неполном покрытии тестами, можно легко показать, что в коде могут содержаться баги.

Наша реализация:

```go
package codewithbags

type AverageCalculator struct {
}

func NewAverageCalculator() *AverageCalculator {
	return &AverageCalculator{}
}

func (c *AverageCalculator) calculateAverage(numbers []int) float64 {
	countNumbers := len(numbers)
	if countNumbers == 0 {
		return 0
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return float64(sum) / float64(countNumbers)
}
```

Тесты:

```go
package codewithbags

import "testing"

func TestAverageCalculator(t *testing.T) {
	calculator := NewAverageCalculator()

	cases := []struct {
		numbers  []int
		expected float64
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0},
		{[]int{}, 0.0},
		{[]int{10, 20, 30}, 20.0},
	}

	for _, c := range cases {
		result := calculator.calculateAverage(c.numbers)

		if result != c.expected {
			t.Errorf("Expected %f, got %f", c.expected, result)
		}
	}
}
```

В данном случае код покрыт на 100% и точно проверяет наш код:

```bash
invisible_mechanisms_of_logic/code_with_bags/average_calculator.go:6:	NewAverageCalculator	100.0%

invisible_mechanisms_of_logic/code_with_bags/average_calculator.go:10:	calculateAverage	100.0%
```

```bash
PASS
ok  	invisible_mechanisms_of_logic/code_with_bags	0.448s
```

Но если мы уберем один кейс, который проверяет, что при нулевой длина входного массива - у нас должен быть 0, тогда мы упускаем эту проверку, следовательно, если бы сама функция была написана некорректно, относительно определенного поведения (отсутствовала проверка на длину входящего массива), тогда были бы баги.

## Калькулатор оценок

В данном случае было много подводных камней и инвариантов, которые необходимо было проверить.

Я написал несколько видов тестов:

1. Глазами - когда ты в голове прогоняешь работу своего кода
2. Руками - когда ты просто выносишь пару кейсов с котрыми запускаешь код
3. Автотесты - прогоняешь автоматически все инварианты и краевые случаи, которые видишь
4. Фаззинг тестирование - кроме своих инвариантов - прогоняются и "мусорные данные"

Из всех способов самым полезным оказался фаззинг - ведь именно этот способ выявил ошибки переполнения типа и его проверки, что говорит о том, что даже в хорошо написанном коде и полностью протестированным - получилось, что все еще есть баги.

Реализация:

```go
package codewithbags

import (
	"errors"
	"math"
)

type GradeCalculator struct{}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{}
}

var (
	ErrNegativeGrade = errors.New("grade is negative")
	ErrEmptyGrades   = errors.New("grades are empty")
	ErrOverflow      = errors.New("int overflow")
	ErrGradeTooLarge = errors.New("grade is too large")
)

var (
	ZeroAvgGrade    = 0.0
	InvalidAvgGrade = -1.0
	MaxGrade        = 100
)

// conds:
// - zero len grades
// - negative grades
// - grades are numbers not runes (it's check by compiler)
func (gc *GradeCalculator) calculateAverage(grades []int) (float64, error) {
	countGrades := len(grades)

	if countGrades == 0 {
		return InvalidAvgGrade, ErrEmptyGrades
	}

	if countGrades > math.MaxInt/MaxGrade {
		return InvalidAvgGrade, ErrOverflow
	}

	sum := 0
	for _, grade := range grades {
		if grade < 0 {
			return InvalidAvgGrade, ErrNegativeGrade
		}

		if grade > MaxGrade {
			return InvalidAvgGrade, ErrGradeTooLarge
		}

		sum += grade
	}

	return float64(sum) / float64(countGrades), nil
}
```

Тесты:

```go
package codewithbags

import (
	"encoding/binary"
	"errors"
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
)

// 1. handle test with viewing through the eyes

// 2. simple auto test
func TestGradeCalculator_Simple(t *testing.T) {
	t.Helper()

	gradeCalculator := NewGradeCalculator()

	cases := []struct {
		grades   []int
		expected float64
		err      error
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0, nil},                                     // happy-path case
		{[]int{}, InvalidAvgGrade, ErrEmptyGrades},                           // no grades
		{[]int{10, 20, -30}, InvalidAvgGrade, ErrNegativeGrade},              // negative grades
		{[]int{-10, -20, -30}, InvalidAvgGrade, ErrNegativeGrade},            // negative sum
		{[]int{1000, 12222, 23123123123}, InvalidAvgGrade, ErrGradeTooLarge}, // too large grade
	}

	for _, c := range cases {
		result, err := gradeCalculator.calculateAverage(c.grades)

		require.Equalf(t, c.err, err, "Unexpected err: %w, expected = %w", err, c.err)
		require.Equalf(t, c.expected, result, "Err with calculate: expected = %f, result = %f", c.expected, result)

	}
}

func intsToBytes(nums []int) []byte {
	buf := make([]byte, 0, len(nums)*2)
	for _, n := range nums {
		buf = binary.AppendVarint(buf, int64(n))
	}
	return buf
}

func bytesToInts(data []byte) []int {
	nums := make([]int, 0, len(data))
	for len(data) > 0 {
		v, n := binary.Varint(data)
		if n <= 0 {
			break
		}

		if int64(int(v)) != v {
			break
		}
		nums = append(nums, int(v))
		data = data[n:]
	}
	return nums
}

var matchErrorsMap = map[error]bool{
	ErrEmptyGrades:   true,
	ErrGradeTooLarge: true,
	ErrInvalidValue:  true,
	ErrNegativeGrade: true,
}

func matchErrors(targetErr error) bool {
	return matchErrorsMap[targetErr]
}

// 3. fazzing test
func FuzzGradeCalculator(t *testing.F) {
	t.Helper()
	gradeCalculator := NewGradeCalculator()

	var (
		fakeDataBufferSize int = 100
		arraySize          int = 10

		minRange = 0
		maxRange = 100
	)

	fake := faker.New()
	for range fakeDataBufferSize {
		templeArray := make([]int, 0, arraySize)
		for range arraySize {
			templeArray = append(templeArray, int((float64(fake.IntBetween(minRange, maxRange)))))
		}

		numbers := intsToBytes(templeArray)
		t.Add(numbers)
	}

	t.Fuzz(func(t *testing.T, grades []byte) {
		numbers := bytesToInts(grades)
		result, err := gradeCalculator.calculateAverage(numbers)

		if err != nil && !matchErrors(err) {
			t.Fatalf("Unexpected err: %v", err)
		}

		if err != nil {
			require.Equal(t, result, -1.0)

			return
		}

		require.GreaterOrEqual(t, result, 0.0, "Result is negative is wrong")

	})
}

// 4. handle test with handing run programm
func TestRunGradeCalculator(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	result, err := gradeCalculator.calculateAverage([]int{1, 2, 3})

	if err != nil {
		t.Fatalf("didn't expect a mistake, but err is %v", err)
	}

	if result != 2 {
		t.Fatalf("result invalid = %f", result)
	}

	result, err = gradeCalculator.calculateAverage([]int{0})

	if err != nil {
		t.Fatalf("didn't expect a mistake, but err is %v", err)
	}

	if result != 0 {
		t.Fatalf("result invalid = %f", result)
	}

	result, err = gradeCalculator.calculateAverage([]int{})

	if !errors.Is(err, ErrEmptyGrades) {
		t.Fatalf("didn't expect a mistake, but err is %v", err)
	}

	if result != InvalidAvgGrade {
		t.Fatalf("result invalid = %f", result)
	}

	result, err = gradeCalculator.calculateAverage([]int{-1, -2, -3})

	if !errors.Is(err, ErrNegativeGrade) {
		t.Fatalf("didn't expect a mistake, but err is %v", err)
	}

	if result != InvalidAvgGrade {
		t.Fatalf("result invalid = %f", result)
	}
}

```
