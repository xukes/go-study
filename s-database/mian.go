package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type DtUser struct {
	gorm.Model
	Name   string
	Age    int
	Gender string
	// hundreds of fields
}

func (DtUser) TableName() string {
	return "dt_user"
}
func main() {
	dsn := "db:5432@tcp(47.83.1.32:3306)/tradingbot?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	// 确保连接成功
	sqlDB, err := db.DB()
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
		}
	}(sqlDB)

	tx := db.Begin()
	//tx.Create(&DtUser{Model: gorm.Model{CreatedAt: time.Now()}, Name: "xxxx", Age: 12241, Gender: "man"})

	//db.Create(&DtUser{Name: "xxx", Age: 12, Gender: "man"})
	//tx.Create(&DtUser{Model: gorm.Model{CreatedAt: time.Now()}, Name: "xxxx", Age: 112, Gender: "man"})
	//db.Commit()

	var userLists []DtUser
	userLists = make([]DtUser, 200)
	for i := 0; i < 200; i++ {
		uu := uuid.New()
		userLists[i] = DtUser{Model: gorm.Model{CreatedAt: time.Now()}, Name: uu.String(), Age: i + 400, Gender: "man"}
	}
	tx.CreateInBatches(userLists, 20)
	var dtUser []DtUser
	tx.Find(&dtUser, &DtUser{Model: gorm.Model{ID: 5}}, &DtUser{Age: 12})
	fmt.Println(dtUser)

	var ddd []DtUser
	err = tx.Table("dt_user").Select("dt_user.*").Joins("left join users on users.id = dt_user.id").Scan(&ddd).Error
	//db.Commit()
	if err != nil {
		return
	}

	tx.Commit()
}
