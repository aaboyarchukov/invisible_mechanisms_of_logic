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
