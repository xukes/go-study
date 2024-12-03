package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	//不能为空并且大于10
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	chI := make(chan int)
	chS := make(chan string)
	go func() {
		for {
			select {
			case i := <-chI:
				fmt.Println(i)
			case s := <-chS:
				fmt.Println(s)
			}
		}
	}()
	defer close(chI)
	r := gin.Default()
	r.GET("/5lmh", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, fmt.Sprint(err))
			return
		}
		//cookie, _ := c.Cookie("key_cookie")
		//fmt.Printf("%s\n", cookie)
		if person.Age > 1000 {
			chI <- person.Age
		} else {
			chS <- person.Name
		}
		c.String(200, fmt.Sprintf("%#v", person))
	})
	err := r.Run()
	if err != nil {
		return
	}

}

type orm struct {
	//alias *alias
	//db    dbQuerier
	isTx bool
}

func getArr(sad int16, len **int) {
	**len++
}
func (o *orm) Delete(md interface{}, cols ...string) (int64, error) {
	return 1, nil
}
