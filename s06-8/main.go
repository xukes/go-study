package main

import "fmt"

func main() {
	func(fv string) {
		fmt.Println(fv)
	}("hello world")

	fmt.Println(f())

}

func f() (ret int) {

	return 1
}
