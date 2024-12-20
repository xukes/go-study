package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xukes/go-study/common"
	pb "github.com/xukes/go-study/proto"
	"net/http"
	"time"
)

type Person struct {
	//不能为空并且大于10
	Age      int       `form:"age" binding:"required,gt=10" json:"age"`
	Name     string    `form:"name" binding:"required" json:"name"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1" json:"birthday"`
	user     int
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
	route := gin.Default()
	route.Use(common.HandlerError, common.ShowHandleTime)
	//route.Use(common.HandlerError)
	v1 := route.Group("/v1")
	v2 := v1.Group("/gg")

	// gin 多个Use该处理逻辑是什么。

	v2.POST("post/", handlePost)
	v2.POST("/upload", upload)

	err := route.Run(":8443")
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
		c.Error(fmt.Errorf("this is an error")).SetMeta(&OrderMessage{Code: 29, Msg: err.Error(), Data: Orm{
			IsTx: false,
		}})
		return
	}
	if person.Age > 100 {
		err := c.Error(fmt.Errorf("this is an error")).SetMeta(&OrderMessage{Code: 29, Msg: "this is a error", Data: Orm{
			IsTx: false,
		}})

		if err != nil {
		}
		return
	}
	o := &OrderMessage{Data: person, Msg: "sds", Code: 23}
	c.JSON(200, o)
}

type OrderMessage struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func getName(userType *pb.UserType) {
	fmt.Println(userType)
}

type Orm struct {
	//alias *alias
	//db    dbQuerier
	IsTx bool `json:"isTx"`
}

func getArr(sad int16, len **int) {
	**len++
}
func (o *Orm) Delete(md interface{}, cols ...string) (int64, error) {
	return 1, nil
}
