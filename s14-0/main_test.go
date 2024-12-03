package even

import (
	"fmt"
	"testing"
	"time"
)

func TestUcc(t *testing.T) {
	cdd := make(chan *int)

	go func() {
		for i := 1; i < 100; i++ {
			cdd <- &i
		}
	}()

	go func() {
		for {
			fmt.Println(<-cdd)
		}
	}()

	time.Sleep(time.Second * 2)
}
