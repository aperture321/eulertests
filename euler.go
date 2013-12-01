package main

import "fmt"
import "time"

//implement for under 1000 fizzbuzz, add them together.
func nums(v int) int {
	res := 0
	for i := 0; i < v; i++ {
		if i%3 == 0 || i%5 == 0 {
			res += i
		}
	}
	return res
}

func fib() int {
	topval := 4000000 //4 million
	res := 0
	a := 1
	b := 0
	temp := 0
	for a < topval {
		temp = a
		a = a + b
		b = temp
		if a%2 == 0 {
			res += a
		}
	}
	return res
}

func main() {
	start := time.Now()
	result := nums(1000)
	fmt.Println("Multiples of 3 and 5: ", result)
	result = fib()
	fmt.Println("Sum of even fibs: ", result)
	elapsed := time.Since(start)
	fmt.Println("seconds ", elapsed)
}
