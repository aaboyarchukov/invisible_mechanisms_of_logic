package codewithbags

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
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
