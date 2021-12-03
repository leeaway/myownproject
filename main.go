package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	var user []*User
	if errs := db.Table("user_tab").Where("id = ?", 1).Find(&user).GetErrors(); len(errs) > 0 {
		fmt.Println(errs)
	}
	fmt.Println(user)
}

//gorm:
var (
	//变量db 通过init直接初始化
	db     *gorm.DB
	err    error
	dbinfo string
)

func init() {
	//拼接数据库连接信息
	dbinfo = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "52991314wqq", "localhost", "3306", "goentrytask")

	//初始化db
	db, err = gorm.Open("mysql", dbinfo)
	if err != nil {
		fmt.Println("mysql打开失败", err)
		return
	}
}

type User struct {
	ID       int64  `gorm:"column:id" json:"id"`
	Unmae    string `gorm:"uname" json:"unmae"`
	Nickname string `gorm:"nickname" json:"nickname"`
}
