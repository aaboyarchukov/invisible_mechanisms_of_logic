# Асинхронное взаимодействие

Нам необходимо также рассмотреть проблемы возникающие при напимании асинхронного кода, так как они менее заметны при написании кода и отладки программы, для их поисков - необходимы определенные знания и навыки.

## Проблемы

1. **Race condition** - попытка изменить ресурс несколькими потоками программы
2. **Deadlock** - когда используемый поток блокируется навсегда, если два или более потоков ожидают друг друга для завершения какой-либо операции, это приводит к взаимной блокировке. Например, поток A ждёт, пока освободится ресурс, занятый потоком B, а поток B в свою очередь ждёт, пока освободится ресурс, занятый потоком A.
3. **Spurious Wakeups** - ложное использование потока, событие не произошло, но поток среагировал, что может привести к не предвиденным результатам
4. **Visibility Issues** - видимость состояний, при изменении состояния программы, ее премеенных, состояние может хранится в нескольких потоках, что затрудняет их синхронизацию, и что может привести к некорректной работе программы
5. **Forgotten Locks** - забытая блокировка ресурса, если не блокировать необходимы ресурс, это может привести к неопределенному поведению программы, а то и вовсе к ее поломке
6. **Mutual Exclusion** - взаимное исключение, сложность в работе с блокировками может приводить к снижению производительности программы, а также к излишней блокировке потоков
7. **Unusing Thread Pools** - не используя пул потоков - напрягаем аллокатор, что приводит к лишней нагрузке -> снижение производительности

Пример:

```java
public class RaceConditionExample {

    private static int counter = 0;

    public static void main(String[] args) {
        int numberOfThreads = 10;
        Thread[] threads = new Thread[numberOfThreads];

        for (int i = 0; i < numberOfThreads; i++) {
            threads[i] = new Thread(() -> {
                for (int j = 0; j < 100000; j++) {
                    counter++;
                }
            });
            threads[i].start();
        }

        for (int i = 0; i < numberOfThreads; i++) {
            try {
                threads[i].join();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        System.out.println("Final counter value: " + counter);
    }
}
```

Создаётся 10 потоков, каждый из которых увеличивает глобальный счетчик 100000 раз. Однако итоговое значение счетчика может быть неверным, почему? Объясните и напишите правильный вариант.

_Ответ:_

- здесь кроется проблема в том, что созданные потоки изменяют "одновременно" один и тот же ресурс, что приводит к неверным результатам, здесь необходимо блокировать ресурс перед изменением и разблокировать после, для того, чтобы каждая операция выполнилась атомарно; либо можно использовать атомики, но не знаю - есть ли они в java; также не совсем понимаю, как именно синхронизируются потоки между собой и главным потоком программы именно в Java, ведь основной поток может закончится раньше чем закончат работу остальные, подозреваю что точка `join` у потоков создается вот здесь: `threads[i].join()`

На Go

```go
package codewithbags

import (
	"fmt"
	"sync"
)

func RaceConditionExample() {
	counter := 0
	numberOfThreads := 10

	wg := sync.WaitGroup{}

	wg.Add(10)
	for range numberOfThreads {
		wg.Go(func() {
			counter++
		})
	}

	wg.Wait()

	fmt.Printf("Final counter value: %d", counter)
}
```

Исправленная версия:

```go
func RaceConditionExample() {
	counter := 0
	numberOfThreads := 10

	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	wg.Add(10)
	for range numberOfThreads {
		wg.Go(func() {
			mx.Lock()
			counter++
			mx.Unlock()
		})
	}

	wg.Wait()

	fmt.Printf("Final counter value: %d", counter)
}
```

Два потока конкурируют за общие ресурсы, но в результате возникает взаимная блокировка, почему? Объясните и напишите правильный вариант.

```java
public class DeadlockExample {

    private static final Object lock1 = new Object();
    private static final Object lock2 = new Object();

    public static void main(String[] args) {
        Thread thread1 = new Thread(() -> {
            synchronized (lock1) {
                System.out.println("Thread 1 acquired lock1");

                try { Thread.sleep(50); }
                catch (InterruptedException e) { e.printStackTrace(); }

                synchronized (lock2) {
                    System.out.println("Thread 1 acquired lock2");
                }
            }
        });

        Thread thread2 = new Thread(() -> {
            synchronized (lock2) {
                System.out.println("Thread 2 acquired lock2");

                try { Thread.sleep(50); }
                catch (InterruptedException e) { e.printStackTrace(); }

                synchronized (lock1) {
                    System.out.println("Thread 2 acquired lock1");
                }
            }
        });

        thread1.start();
        thread2.start();

        try {
            thread1.join();
            thread2.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println("Finished");
    }
}
```

Здесь проблема заключается в том, что два разных потока пытаются захватить одни и те же ресурсы, от этого один из процессов может заблокироваться навсегда и не продолжить работу.

_Ответ:_

- необходимо внимательно следить за ресурсами и не допускать случай, когда поток остается неиспользованным и заблокированным навсегда, в данном примере нам необходимо дождаться готовности и только потом захватить второй ресурс, а не захватывать его сразу

На Go:

```go
func DeadlockExample() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func() {
		<-ch2
		ch1 <- struct{}{}
	}()

	go func() {
		<-ch1
		ch2 <- struct{}{}
	}()
}
```

В Golang здесь можно показать это с каналами, запись и чтение из канала блокирующие операции, которые захватывают на себя горутины, чтобы разблокировать - необходимо либо прочитать при записи, либо записать при чтении И что важно - закрыть канал, чтобы не блокироваться в будущем.

Исправленная версия:

```go
func DeadlockExample() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	wg.Go(func() {
		<-ch2
		ch1 <- struct{}{}
	})

	wg.Go(func() {
		ch2 <- struct{}{}
		<-ch1
	})

	wg.Wait()
}
```

Или еще одна версия с захватом ресурса:

```go
func DeadlockExampleMutex() {
	mu1, mu2 := sync.Mutex{}, sync.Mutex{}

	var wg sync.WaitGroup
	wg.Add(2)

	wg.Go(func() {
		mu1.Lock()
		fmt.Println("goroutine 1 acquired mu1")

		time.Sleep(50 * time.Millisecond)

		mu2.Lock()
		fmt.Println("goroutine 1 acquired mu2")

		mu2.Unlock()
		mu1.Unlock()
	})

	wg.Go(func() {
		mu2.Lock()
		fmt.Println("goroutine 2 acquired mu2")

		time.Sleep(50 * time.Millisecond)

		mu1.Lock()
		fmt.Println("goroutine 2 acquired mu1")

		mu1.Unlock()
		mu2.Unlock()
	})

	wg.Wait()
}
```

Исправленная версия:

```go
func DeadlockExampleMutex() {
	mu1, mu2 := sync.Mutex{}, sync.Mutex{}

	var wg sync.WaitGroup
	wg.Add(2)

	wg.Go(func() {
		mu1.Lock()
		fmt.Println("goroutine 1 acquired mu1")

		time.Sleep(50 * time.Millisecond)

		mu1.Unlock()

	})

	wg.Go(func() {
		mu2.Lock()
		fmt.Println("goroutine 2 acquired mu2")

		time.Sleep(50 * time.Millisecond)

		mu2.Unlock()
	})

	wg.Wait()
}
```

Будет лучше разделить ресурсы по каждым горутинам.

## Примеры механизмов синхронизации

Приведем несколько примеров, где мы будем синхронизировать выполнение асинхронных вызовов:

1. Атомики
   Используется для облегчения взаимодействия с общими ресурсами, атомики позволяются выполнять операции без использования механизмов блокировки:

```go
package codewithbags

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCounter() {
	workers := 10
	var counter atomic.Int64
	wg := sync.WaitGroup{}

	for range workers {
		wg.Go(func() {
			counter.Add(1)
		})
	}

	wg.Wait()

	fmt.Printf("Result: %d", counter.Load())
}
```

Таким образом мы не блокируем с помощью `mutex` общий ресурс и наши воркеры спокойно изменяют ресурс.

2. Обработка событий
   В Go мы можем это сделать с помощью внутреннего механизма (каналы с буффером - циклический массив с блокировками)

```go
func SomeEvent() {
	time.Sleep(2 * time.Millisecond)
}

func Producer(ctx context.Context, events chan<- string) {
	producers := 10

	wg := sync.WaitGroup{}
	for range producers {
		wg.Go(func() {
			for n := 0; ; n++ {
				SomeEvent()
				select {
				case events <- fmt.Sprintf("Event%d", n):
				case <-ctx.Done():
					close(events)
					return
				}
			}
		})
	}
}

func Worker(ctx context.Context, id int, events <-chan string) {
	for {
		select {
		case val, ok := <-events:
			if !ok {
				return
			}
			fmt.Printf("worker %d: %s\n", id, val)
		case <-ctx.Done():
			return
		}
	}
}
```

Здесь мы бесконечно пишем и обрабатываем события.

3. FanIn
   Классический паттерн по объединению нескольких каналов в один

```go
func FanIn(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, in := range inputs {
		wg.Go(func() {
			defer wg.Done()
			for v := range in {
				out <- v
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
```

4. ErrGroup
   Синхронизация выполнения нескольких потоков - засчет отслеживания возникновения ошибок, при возникновении ошибки - все потоки отменяются

```go
func Fetch(ctx context.Context, url string) (string, error) {
	time.Sleep(2 * time.Millisecond)

	return "", nil
}

func FetchAll(ctx context.Context, urls []string) ([]string, error) {
	g, ctx := errgroup.WithContext(ctx)
	results := make([]string, len(urls))

	for i, url := range urls {
		g.Go(func() error {
			body, err := Fetch(ctx, url)
			if err != nil {
				return fmt.Errorf("fetch %s: %w", url, err)
			}

			results[i] = body
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}
```

5. InMemory AsyncSafty Cache
   Имплементация кеша, которая устойчив к асинхроному взаимодействию

```go
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]string)}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}
```
