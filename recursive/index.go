package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	//fact(7) sẽ tính 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println(fact(7))

	var fib func(n int) int

	//Định nghĩa cho hàm 
	fib = func(n int) int {
        if n < 2 {
            return n
        }

        return fib(n-1) + fib(n-2)
    }
    fmt.Println(fib(7))
}