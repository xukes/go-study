package main

import (
	"fmt"
	"strings"
)

func main() {
	sd := strings.IndexFunc("2439h234", ds)
	fmt.Printf("%d \n", sd)

	callback(3, Add)

	ch1 := make(chan int)

	go func() {
		close(ch1)
	}()
	select {
	case <-ch1:
		fmt.Println("close")
	}

}

func ds(ss int32) bool {
	if ss > 255 {
		return false
	}
	return true
}

func Add(a, b int) {
	fmt.Printf("%d, %d ,%d", a, b, a+b)
}
func callback(y int, f func(int, int)) {
	f(y, 2)
}
