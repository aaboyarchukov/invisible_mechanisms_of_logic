package complexcodelookssimple

import (
	"fmt"
	"math/rand/v2"
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
