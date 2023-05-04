package main

import "fmt"

func main() {
	sum := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(sum(1))
	}

	num := nums(4)

	fmt.Println(num(4))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func nums(num int) func(int) int {
	num += 10
	return func(x int) int {
		sum := num + x
		return sum
	}
}
