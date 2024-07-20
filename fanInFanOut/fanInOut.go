package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"

	"github.com/lakshyabatman/go-learning/dataPipeline"
)

func fanIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}
	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()
	return fannedInStream
}

func primeFinder(done <-chan int, randIntStream <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}
	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randomeInt := <-randIntStream:
				if isPrime(randomeInt) {
					primes <- randomeInt
				}
			}
		}
	}()
	return primes
}

func main() {
	numOfCpu := runtime.NumCPU()
	done := make(chan int)
	randNumFetcher := func() int { return rand.Intn(10) }
	stream := dataPipeline.RepeatFunc(done, randNumFetcher)
	primeFinderChannels := make([]<-chan int, numOfCpu)
	for i := 0; i < numOfCpu; i++ {
		primeFinderChannels[i] = primeFinder(done, stream)
	}
	primes := fanIn(done, primeFinderChannels...)
	for i := range primes {
		fmt.Println(i)
	}

}
