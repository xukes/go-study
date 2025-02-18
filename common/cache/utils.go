package cache

import (
	"fmt"
	"gorm.io/gorm"
)

func init() {
	fmt.Println("utils init")

	us := &User{
		DtUser.Age: 123,
		School:     "xxx",
	}
	fmt.Println(us)
}

type DtUser struct {
	gorm.Model
	Name   string `gorm:"column:user_name"`
	Age    int    `gorm:"column:age"`
	Gender string `gorm:"column:gender"`
}

func (DtUser) TableName() string {
	return "dt_user"
}

type User struct {
	DtUser
	School string
}
