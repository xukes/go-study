package main

import "fmt"

func main() {
	func(fv string) {
		fmt.Println(fv)
	}("hello world")
}
