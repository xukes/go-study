package main

import (
	"fmt"
)

func main() {
	// 连接到 MySQL 数据库

	pos := 4
	result, pos := fibonacci(pos)
	fmt.Printf("the %d-th fibonacci number is: %d\n", pos, result)

	a := Address{s: 2}
	a.sayHello()

	Address.sayHell(a)

	f := Adder(1, 3)
	fmt.Printf("%d", f(4))
}

func Adder(a int, b int) func(c int) int {
	f := func(d int) int {
		return b + a + d
	}
	return f
}

func (a *Address) sayHello() {
	fmt.Println(a.s)
}

func (Address) sayHell() {
	fmt.Println()
}

type Address struct {
	s int
}

func fibonacci(n int) (val, pos int) {
	if n <= 1 {
		val = 1
	} else {
		v1, _ := fibonacci(n - 1)
		v2, _ := fibonacci(n - 2)
		val = v1 + v2
	}
	pos = n
	return
}
