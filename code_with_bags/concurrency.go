package codewithbags

import (
	"fmt"
	"sync"
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
// 		ch1 <- struct{}{}
// 	}()

// 	go func() {
// 		ch2 <- struct{}{}
// 	}()
// }

// correct version
func DeadlockExample() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func() {
		ch1 <- struct{}{}
	}()

	go func() {
		ch2 <- struct{}{}
	}()

	<-ch1
	close(ch1)

	<-ch2
	close(ch2)
}
