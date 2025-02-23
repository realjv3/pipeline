package main

import "sync"

// square squares input range of integers using a fan-out/fan-in pattern
func square(in <-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	worker := func(in <-chan int) {
		for n := range in {
			out <- n * n
		}
		wg.Done()
	}

	wg.Add(3)
	go worker(in)
	go worker(in)
	go worker(in)

	// By putting wg.Wait() in a separate goroutine:
	// The function returns out immediately, so the caller (e.g. printInts) can start consuming values.
	// As the caller reads values from out, the workers can keep sending without blocking.
	// Once all workers finish processing, the separate goroutine closes out, signaling completion.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
