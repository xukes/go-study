package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	pb "go-study/proto"
	"time"
)

type Person struct {
	//不能为空并且大于10
	Age      int       `form:"age" binding:"required,gt=10" json:"age"`
	Name     string    `form:"name" binding:"required" json:"name"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1" json:"birthday"`
	user     int       "how use it"
}

func main() {
	b := pb.Balance{}
	fmt.Println(b)

	userT := pb.UserType_Women
	getName(&userT)

	route := gin.Default()
	route.POST("post/", handlePost)

	err := route.Run(":8443")
	if err != nil {
		return
	}

}
func handlePost(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(500, fmt.Sprint(err))
		return
	}
	c.JSON(200, person)
}

func getName(userType *pb.UserType) {
	fmt.Println(userType)
}

type Orm struct {
	//alias *alias
	//db    dbQuerier
	isTx bool
}

func getArr(sad int16, len **int) {
	**len++
}
func (o *Orm) Delete(md interface{}, cols ...string) (int64, error) {
	return 1, nil
}
