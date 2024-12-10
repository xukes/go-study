package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	pb "go-study/proto"
	"net/http"
	"time"
)

type Person struct {
	//不能为空并且大于10
	Age      int       `form:"age" binding:"required,gt=10" json:"age"`
	Name     string    `form:"name" binding:"required" json:"name"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1" json:"birthday"`
	user     int       "how use it"
}
type Pattern struct {
	Age int
}

type NewUser struct {
	Person  //继承
	Pattern // 继承
	Per     Person
	School  string
}

func main() {
	b := NewUser{Person: Person{}, Per: Person{Age: 1, Name: "111", Birthday: time.Now(), user: 12}, School: "USDT"}
	fmt.Println(b)
	fmt.Println(b.Person.Age)
	bt, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bt))
	var newBa pb.Balance

	err = json.Unmarshal(bt, &newBa)
	if err != nil {
		return
	}

	userT := pb.UserType_Women
	getName(&userT)

	route := gin.Default()
	route.Use()
	v1 := route.Group("/v1")
	v2 := v1.Group("/gg")
	v2.POST("post/", handlePost)
	v2.POST("/upload", upload)

	err = route.Run(":8443")
	if err != nil {
		return
	}
}
func upload(c *gin.Context) {
	file, _ := c.FormFile("xxx")
	c.Param("sd")
	err := c.SaveUploadedFile(file, file.Filename)
	if err != nil {
		fmt.Println(err)
	}
	c.String(http.StatusOK, file.Filename)
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
