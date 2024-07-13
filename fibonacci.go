package main

import "fmt"

func generateFib(c chan int, r int) chan int {
	go func() {
		a, b := 0, 1
		for i := 0; i < r; i++ {
			a, b = b, a+b
			fmt.Println("Sending value", a)
			c <- a
		}
		close(c)
	}()
	return c
}

func main() {
	c := make(chan int)

	for x := range generateFib(c, 10) {
		fmt.Println(x)
	}

}
