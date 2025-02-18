package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xukes/go-study/common"
	pb "github.com/xukes/go-study/proto"
	"google.golang.org/grpc"
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

	accountAddr := fmt.Sprintf("%s:%d", "127.0.0.1", 8082)
	accountConn, err := grpc.Dial(accountAddr, grpc.WithInsecure())
	if err != nil {
	}
	handleServiceClient := pb.NewHandleServiceClient(accountConn)
	h := &Hand{handleServiceClient}

	v2.POST("post/", h.handlePost)
	v2.POST("/upload", h.upload)

	err = route.Run(":8443")
	if err != nil {
		return
	}
}

type Hand struct {
	handleServiceClient pb.HandleServiceClient
}

func (h *Hand) upload(c *gin.Context) {
	file, _ := c.FormFile("xxx")
	c.Param("sd")
	err := c.SaveUploadedFile(file, file.Filename)
	if err != nil {
		fmt.Println(err)
	}
	c.String(http.StatusOK, file.Filename)
}

//	func handelError() {
//		fmt.Println("handelError")
//	}
func (h *Hand) handlePost(c *gin.Context) {
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

	context.WithCancel(context.Background())
	context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	context.WithTimeout(context.Background(), time.Second)
	//context.WithValue(c.Request.Context(), "person", person)
	//context.WithValue(c.Request.Context(), "age", person.Age)
	context.TODO()

	ctx := context.TODO()
	ctx = context.WithValue(ctx, "name", &Person{Age: 13})
	//ctx = context.WithValue(ctx, "key", "user")
	//ctx = context.WithValue(ctx, "key1", "user")
	//ctx = context.WithValue(ctx, "key2", "user")
	//ctx = context.WithValue(ctx, "key3", "user")
	//defer han()
	rsp, err1 := h.handleServiceClient.SendMessage(ctx, &pb.SendMessageRequest{
		ChatId: int64(person.Age),
		Text:   person.Name,
	})

	p, ok := ctx.Value("name").(*Person)
	if ok {
		p.Age = 200
		fmt.Println(*p)
	}
	proseContext(ctx)
	fmt.Println(err1, rsp)

	fmt.Println(uuid.New().String())

	o := &OrderMessage{Data: person, Msg: "sds", Code: 23}
	c.JSON(200, o)
}

func proseContext(ctx context.Context) {
	sd, ok := ctx.Value("name").(*Person)
	if ok {
		fmt.Println(*sd)
	}
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
