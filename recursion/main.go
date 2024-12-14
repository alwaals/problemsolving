package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(SumOfNatural(5))
	fmt.Println(fibonacci(7))
}
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
func SumOfNatural(n int) int {
	if n == 1 {
		return 1
	}
	return n + SumOfNatural(n-1)

}
func fibonacci(n int) int {
	if n == 0 || n == 1 || n ==2{
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
