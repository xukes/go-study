package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/xukes/go-study/common"
	"github.com/xukes/go-study/common/cache"
	pb "github.com/xukes/go-study/proto"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net"
	"time"
)

type MyPerson interface {
	Less()
}
type MyPerson1 struct {
}
type MyPerson2 struct {
	Age  int
	Name string
}

func (m MyPerson2) Less() bool {
	return m.Age < 10
}

func (m *MyPerson1) Less() {
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

type HandleServiceServer struct {
	pb.UnimplementedHandleServiceServer
	db *gorm.DB
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/tradingbot?parseTime=false"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		Logger:                 common.GetLogger(),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	// 确保连接成功
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
		}
	}(sqlDB)

	addr := ":8082"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterHandleServiceServer(s, &HandleServiceServer{db: db})
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}
func (h *HandleServiceServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (resp *pb.SendMessageResponse, err error) {
	resp = &pb.SendMessageResponse{
		Success:   true,
		MessageId: "msg success",
	}
	insertData(req.Text, int(req.ChatId), h.db)

	fmt.Println(req.Text)
	return resp, nil
}

func insertData(name string, age int, db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var futureArr []chan int64
		for i := 0; i < age; i++ {
			fu := make(chan int64)
			futureArr = append(futureArr, fu)
			insertData2(tx, i, name, fu)
		}
		reader(futureArr)
		return nil
	}, &sql.TxOptions{ReadOnly: false})
	if err != nil {
	}
}
func insertData2(tx *gorm.DB, age int, name string, future chan<- int64) {
	go func() {
		time.Sleep(time.Second * time.Duration(rand.Int()%2))
		rr := tx.Create(&cache.DtUser{Model: gorm.Model{CreatedAt: time.Now(), DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: false}}, Name: fmt.Sprintf("%s%d", name, age), Age: age + 0, Gender: "man"})
		future <- rr.RowsAffected
	}()
}
func reader(futureArr []chan int64) {
	total := int64(0)
	for _, f := range futureArr {
		a := <-f
		total += a
	}
	fmt.Printf("total=%d\n", total)
}
func (h *HandleServiceServer) GetMessage(ctx context.Context, req *pb.GetMessageRps) (resp *pb.GetMessageResp, err error) {
	fmt.Println(req.BaseMsg.Msg)
	return &pb.GetMessageResp{ChatId: 23, Text: "msg"}, nil
}
