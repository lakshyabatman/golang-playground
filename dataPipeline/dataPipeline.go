package dataPipeline

func RepeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():

			}
		}
	}()
	return stream
}

func Take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()
	return taken
}

// func main() {
// 	done := make(chan int)

// 	randNumFetcher := func() int { return rand.Intn(5000000) }
// 	defer close(done)

// 	for x := range take(done, repeatFunc(done, randNumFetcher), 10) {
// 		fmt.Println(x)
// 	}
// }
