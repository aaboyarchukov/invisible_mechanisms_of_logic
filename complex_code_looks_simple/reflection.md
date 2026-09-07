# Как код, выглядящий просто, может оказаться сложным

Нам дан пример:

```java
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

public class DateExample {
    public static void main(String[] args) {
        String dateString = "2024-05-13 14:30:00";
        SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        try {
            Date date = format.parse(dateString);
            System.out.println("Date: " + date);
        } catch (ParseException e) {
            e.printStackTrace();
        }
    }
}
```

Необходимо выяснить, какие есть недостатки у данного решения и сделать решение лучше.

Как выглядит на Go:

```go
package complexcodelookssimple

import (
	"fmt"
	"time"
)

func main() {
	dateString := "2024-05-13 14:30:00"

	date, err := time.Parse("2006-01-02 15:04:05", dateString)
	if err != nil {
		fmt.Printf("parse date: %v\n", err)
		panic("parse date")
	}

	fmt.Println("Date:", date)
}
```

Недостатки:

- формат даты вынес бы в отдельную переменную: "yyyy-MM-dd HH:mm:ss", либо использовал гоотвые константы (если такие есть в Java библиотеках)
- после нахождения исключения - мы сразу печатаем stackTrace и аварийно завершаем программу
- из-за аварийного завершения программы - мы не можем протестировать ошибочное поведение программы
- нет UTC - таймзоны, из-за этого будет подставлена UTC конкретной JVM - недетерменированное поведение

Улучшенный пример:

```go
package complexcodelookssimple

import (
	"fmt"
	"time"
)

const (
	inputDate = "2024-05-13 14:30:00"
	timeZone  = "Europe/Moscow"
)

func ParseDate(s string, loc *time.Location) (time.Time, error) {
	date, err := time.ParseInLocation(time.DateTime, s, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse date %q: %w", s, err)
	}

	return date, nil
}

func ExampleParseDate() {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		fmt.Printf("load location %q: %v\n", timeZone, err)
		return
	}

	date, err := ParseDate(inputDate, loc)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Date:", date)
}
```

Следующий пример:

```java
public class ThreadExample {
    private static int counter = 0;

    public static void main(String[] args) {
        Runnable task = () -> {
            for (int i = 0; i < 1000; i++) {
                counter++;
            }
        };

        Thread thread1 = new Thread(task);
        Thread thread2 = new Thread(task);

        thread1.start();
        thread2.start();

        try {
            thread1.join();
            thread2.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println("Counter: " + counter);
    }
}
```

Эквивалент на Go:

```go
func ThreadExample() {
	counter := 0

	wg := &sync.WaitGroup{}

	var countFunc func() = func() {
		for range 1000 {
			counter++
		}
	}

	wg.Go(func() {
		countFunc()
	})

	wg.Go(func() {
		countFunc()
	})

	wg.Wait()
}
```

Здесь простой пример гонки данных, решается блокировкой ресура при взаимном изменении разных потоков.

Улучшенный пример:

```go
package complexcodelookssimple

import (
	"fmt"
	"sync"
)

func ThreadExample(wg *sync.WaitGroup, mu *sync.Mutex) {
	counter := 0

	var countFunc func() = func() {
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	wg.Go(func() {
		countFunc()
	})

	wg.Go(func() {
		countFunc()
	})

	wg.Wait()

	fmt.Println(counter)
}
```

Следующий пример:

```java
import java.util.Random;

public class ComplexMultiThreadProcessing {
    private static final int SIZE = 1000000;
    private static final int THREADS = 4;
    private static final int[] data = new int[SIZE];
    private static volatile int sum = 0;

    public static void main(String[] args) {
        Random random = new Random();
        for (int i = 0; i < SIZE; i++) {
            data[i] = random.nextInt(100);
        }

        Thread[] threads = new Thread[THREADS];
        int chunkSize = SIZE / THREADS;

        for (int i = 0; i < THREADS; i++) {
            final int start = i * chunkSize;
            final int end = (i + 1) * chunkSize;
            threads[i] = new Thread(() -> {
                int localSum = 0;
                for (int j = start; j < end; j++) {
                    localSum += data[j];
                }
                synchronized (ComplexMultiThreadProcessing.class) {
                    sum += localSum;
                }
            });
            threads[i].start();
        }

        for (int i = 0; i < THREADS; i++) {
            try {
                threads[i].join();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        System.out.println("Sum of all elements: " + sum);
    }
}
```

На Go:

```go
func ComplexMultiThreadProcessing() {
	var (
		SIZE    = 1000000
		THREADS = 4
	)

	data := make([]int, SIZE)
	sum := 0

	randomEngine := rand.New(
		&rand.Rand{},
	)

	for i := range SIZE {
		data[i] = randomEngine.IntN(100)
	}

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	chunkSize := SIZE / THREADS
	for i := range THREADS {
		start := i * chunkSize
		end := start + chunkSize
		if i == THREADS-1 {
			end = SIZE
		}

		wg.Go(func() {
			localSum := 0
			for _, v := range data[start:end] {
				localSum += v
			}

			mu.Lock()
			sum += localSum
			mu.Unlock()
		})
	}

	wg.Wait()
	fmt.Println("Sum of all elements:", sum)
}
```

Данный пример построен на асинхронной работе нескольких потоков, которые суммируют числа.

Недостатки:

- магические числа
- непонятные переменные
- работа, которая выполняется внутри циклов

Улучшенный пример:

```go
func accumSum(data []int, chanAccumValue chan int, start, end int) {
	localSum := 0
	chunkSlice := data[start:end]

	for _, v := range chunkSlice {
		localSum += v
	}

	chanAccumValue <- localSum
}

func processRanges(indx, chunkSize, size, amountThreads int) (int, int) {
	start := indx * chunkSize
	end := start + chunkSize
	if indx == amountThreads-1 {
		end = size
	}

	return start, end
}

type bufferOfData struct {
	buffer []int
}

func NewBufferOfData(size int) *bufferOfData {

	return &bufferOfData{
		buffer: make([]int, size),
	}
}

func (b *bufferOfData) Map(mapFunc func(ranges int) int, ranges int) {
	size := len(b.buffer)
	for i := range size {
		b.buffer[i] = mapFunc(ranges)
	}
}

func (b *bufferOfData) AsyncAccumSum(amountThreads int, wg *sync.WaitGroup, sumChan chan int) {
	size := len(b.buffer)
	chunkSize := size / amountThreads

	for i := range amountThreads {
		start, end := processRanges(i, chunkSize, size, amountThreads)
		wg.Go(func() {
			accumSum(b.buffer, sumChan, start, end)
		})
	}

}

func ComplexMultiThreadProcessing() {
	var (
		SIZE      = 1000000
		THREADS   = 4
		intRanges = 100
		sum       = 0
	)

	data := NewBufferOfData(SIZE)

	var (
		wg            = &sync.WaitGroup{}
		resultSumChan = make(chan int, THREADS)
	)

	data.Map(rand.IntN, intRanges)
	data.AsyncAccumSum(THREADS, wg, resultSumChan)

	go func() {
		wg.Wait()
		close(resultSumChan)
	}()

	for localSum := range resultSumChan {
		sum += localSum
	}

	fmt.Println("Sum of all elements:", sum)
}
```

Мы изменили способ блокировки при общем доступе к ресурсу (вместо мьютекса - каналы), а также организовали более читаемый код, разделив более сложный код - на несколько чистых функций, уменьшив когнитивную и цикломатическую нагрузку основной функции.

Но для Java модели есть механизм ForkJoinPool - это специализированный пул потоков в Java для управления и выполнения задач, которые могут быть рекурсивно разделены на подзадачи.

```java
import java.util.Random;
import java.util.Arrays;
import java.util.concurrent.ForkJoinPool;

public class SimplifiedMultiThreadProcessing {
    private static final int SIZE = 1000000;
    private static final int[] data = new int[SIZE];

    public static void main(String[] args) {
        Random random = new Random();
        for (int i = 0; i < SIZE; i++) {
            data[i] = random.nextInt(100);
        }

        ForkJoinPool pool = new ForkJoinPool();

        int sum = pool.submit(() -> Arrays.stream(data).parallel().sum()).join();

        System.out.println("Sum of all elements: " + sum);
    }
}
```

Следующий пример необходимо было выразить в виде чистых функций применив техники разработки в подходе ФП:

```go
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
```
