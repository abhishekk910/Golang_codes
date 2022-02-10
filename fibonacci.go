package main

import "fmt"

func fibonacci(n int) {
	a, b, c := 0, 1, 0
	for i := 1; i <= n; i++ {
		fmt.Printf("%d :- %d\n", i, a)
		c = a + b
		a = b
		b = c
	}
}

func main() {
	var num int
	fmt.Print("Enter a number : ")
	fmt.Scan(&num)
	fibonacci(num)

}
