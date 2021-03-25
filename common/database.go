package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)





var DB  *gorm.DB
func InitDB() *gorm.DB{
	DB, _ := gorm.Open("mysql","interview:acat@tcp(192.168.1.197:3306)/recruitment?charset=utf8mb4&parseTime=True&loc=Local")
	//DB,err := gorm.Open("mysql","root:jxc627@tcp(127.0.0.1:3306)/student")

	return DB
}

