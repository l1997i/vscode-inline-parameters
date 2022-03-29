package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func inc(x, y, z int) (a, b, c int) {
	a = x + 1
	b = y + 1
	c = z + 1
	return
}

func sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

func main() {
	// Simple functions
	res := plus(plus(1, 2), 1)
	fmt.Println("1+2 =", res)

	any := plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", any)

	// Multiple return values
	x, y, z := inc(10, 100, 1000)
	fmt.Println(x, y, z)

	// Variadic function
	res = sum(1, 2, 3, 4, 5)
	fmt.Println(res)

	// Anonymous function
	sum := func(a, b, c int) int {
		return a + b + c
	}(3, 5, 7)

	fmt.Println("5+3+7 =", sum)
}