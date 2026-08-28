package codewithbags

import (
	"fmt"
	"sync"
	"time"
)

// func RaceConditionExample() {
// 	counter := 0
// 	numberOfThreads := 10

// 	wg := sync.WaitGroup{}

// 	wg.Add(10)
// 	for range numberOfThreads {
// 		wg.Go(func() {
// 			counter++
// 		})
// 	}

// 	wg.Wait()

// 	fmt.Printf("Final counter value: %d", counter)
// }

// correct version
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

// func DeadlockExample() {
// 	ch1 := make(chan struct{})
// 	ch2 := make(chan struct{})

// 	go func() {
// 		<-ch2
// 		ch1 <- struct{}{}
// 	}()

// 	go func() {
// 		<-ch1
// 		ch2 <- struct{}{}
// 	}()
// }

// func DeadlockExampleMutex() {
// 	mu1, mu2 := sync.Mutex{}, sync.Mutex{}

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	wg.Go(func() {
// 		mu1.Lock()
// 		fmt.Println("goroutine 1 acquired mu1")

// 		time.Sleep(50 * time.Millisecond)

// 		mu2.Lock()
// 		fmt.Println("goroutine 1 acquired mu2")

// 		mu2.Unlock()
// 		mu1.Unlock()
// 	})

// 	wg.Go(func() {
// 		mu2.Lock()
// 		fmt.Println("goroutine 2 acquired mu2")

// 		time.Sleep(50 * time.Millisecond)

// 		mu1.Lock()
// 		fmt.Println("goroutine 2 acquired mu1")

// 		mu1.Unlock()
// 		mu2.Unlock()
// 	})

// 	wg.Wait()
// }

// correct version
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

// correct version
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
