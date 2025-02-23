package main

func generate(count int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}
