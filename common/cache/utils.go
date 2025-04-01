package cache

import (
	"fmt"
	"gorm.io/gorm"
)

func init() {
	fmt.Println("utils init")

	us := &User{
		DtUser{Name: "xxx", Age: 12, Gender: "women"},
		"xxx",
	}
	us.rel()
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

func (dt *DtUser) rel() string {
	return ""
}

type User struct {
	DtUser
	School string
}

func (u *User) rel() string {
	return ""
}
