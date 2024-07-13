package main

import "fmt"

func fizzBuzz(i int) string {
	switch {
	case i%15 == 0:
		return "FizzBuzz"
	case i%3 == 0:
		return "Fizz"
	case i%5 == 0:
		return "Buzz"
	}
	return nil
}

func main() {
	for i := 0; i < 101; i++ {
		message := fizzBuzz(i)
		fmt.Println(message)
	}
}
