package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	module "github.com/leeaway/MyProtobuf"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/protobuf", func(c *gin.Context) {

		data := &module.User{
			Name: "张三",
			Age:  20,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	r.Run(":8080")

	//resp, err := http.Get("http://localhost:8080/protobuf")
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		defer resp.Body.Close()
	//		body, err := ioutil.ReadAll(resp.Body)
	//		if err != nil {
	//			fmt.Println(err)
	//		} else {
	//			user := &module.User{}
	//			proto.UnmarshalMerge(body, user)
	//			fmt.Println(*user)
	//		}
	//	}

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
