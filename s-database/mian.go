package main

import (
	"database/sql"
	"fmt"
	"github.com/xukes/go-study/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type DtUser struct {
	gorm.Model
	Name   string `gorm:"column:user_name"`
	Age    int
	Gender string
	// hundreds of fields
}

func (DtUser) TableName() string {
	return "dt_user"
}
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/tradingbot?parseTime=false"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		Logger:                 common.GetLogger(),
		SkipDefaultTransaction: true,
	})

	//db, err := gorm.Open(mysql.Open(dsn))
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

	insertData(db)

	updateDate(db)

	//db.Create(&DtUser{Name: "xxx", Age: 12, Gender: "man"})
	//tx.Create(&DtUser{Model: gorm.Model{CreatedAt: time.Now()}, Name: "xxxx", Age: 112, Gender: "man"})
	//db.Commit()
	//tx := db.Begin()
	//var dtUser []DtUser
	//tx.Find(&dtUser, &DtUser{Model: gorm.Model{ID: 3}}, &DtUser{Age: 12241})
	//fmt.Println(dtUser)

	//var ddd []DtUser
	//err = tx.Table("dt_user").Select("dt_user.*").Joins("left join users on users.id = dt_user.id").Scan(&ddd).Error
	//db.Commit()

	//tx.Commit()
}

func updateDate(db *gorm.DB) {
	//db.Find(&dtUser, &DtUser{Model: gorm.Model{ID: 3}})
	// 根据条件更新
	result := db.Model(&DtUser{Model: gorm.Model{ID: 2}}).Updates(DtUser{Age: 11121, Name: "xqxxw", Model: gorm.Model{DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: true}}})

	//result := db.Unscoped().Delete(&DtUser{Model: gorm.Model{ID: 1}})
	fmt.Println(result.RowsAffected)
}

//	func (u *DtUser) BeforeCreate(tx *gorm.DB) (err error) {
//		return
//	}
//
//	func (u *DtUser) AfterCreate(tx *gorm.DB) (err error) {
//		if u.ID == 1 {
//		}
//		return
//	}
func insertData(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&DtUser{Model: gorm.Model{CreatedAt: time.Now(), DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: false}}, Name: "hello", Age: 12241, Gender: "man"}).Error; err != nil {
			return err
		}
		return nil
	}, &sql.TxOptions{ReadOnly: false})
	if err != nil {
	}
}
