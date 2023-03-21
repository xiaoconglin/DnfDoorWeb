package mysql_gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func init() {
	log.Println("初始化数据库...")
	dsn := "root:lxc123123@tcp(localhost:3306)/myweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if db.Error != nil {
		fmt.Printf("database error %v", Db.Error)
	}
	Db = db
}
