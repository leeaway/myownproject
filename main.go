package main

import (
	"example.com/m/v2/Huawei/problemlibrary"
	"example.com/m/v2/demo"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "net/http/pprof"
	"reflect"
)

func main() {
	//echart.EchartBar()
	//http.HandleFunc("/line", echart.Httpserver)
	//http.HandleFunc("/pie", echart.HttpserverPie)
	//http.HandleFunc("/bar3d", examples.Bar3dExamples{}.Examples)
	//http.HandleFunc("/geo", examples.GeoExamples{}.Examples)
	//http.HandleFunc("/gauge", examples.GaugeExamples{}.Examples)
	//http.HandleFunc("/tree", examples.TreeExamples{}.Examples)
	//http.HandleFunc("/graph", examples.GraphExamples{}.Examples)
	//http.HandleFunc("/map", examples.MapExamples{}.Examples)
	//http.HandleFunc("/center", examples.PageCenterLayoutExamples{}.Examples)
	//http.HandleFunc("/flex", examples.PageFlexLayoutExamples{}.Examples)
	//http.HandleFunc("/non", examples.PageNoneLayoutExamples{}.Examples)

	fmt.Println(problemlibrary.CalShoppingList())
}

func init() {
	//AOP
	demo.RegisterPoint(reflect.TypeOf((*HelloAop)(nil)))
	demo.RegisterAspect("main.HelloAop.TestAop", &Aspect{})
}

type Aspect struct{}

func (a *Aspect) Before(point *demo.JoinPoint) bool {
	fmt.Println("before")
	return true
}

func (a *Aspect) After(point *demo.JoinPoint) {
	fmt.Println("after")
}

func (a *Aspect) Finally(point *demo.JoinPoint) {
	fmt.Println("finally")
}

func (a *Aspect) GetAspectExpress() string {
	return ".*\\.HelloAop.TestAop"
}

type HelloAop struct {
}

func (h *HelloAop) TestAop() {
	fmt.Println("helloAop")
}

//r := gin.Default()
//r.GET("/protobuf", GETHandler)
//r.Run(":8080")

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

//var user []*User
//if errs := db.Table("user_tab").Where("id = ?", 1).Find(&user).GetErrors(); len(errs) > 0 {
//	fmt.Println(errs)
//}
//fmt.Println(user)

//func GETHandler(c *gin.Context) {
//	nums := c.Query("nums") //string
//	sum := int64(0)
//	for _, num := range nums {
//		sum += int64(num)
//	}
//	detail := c.DefaultQuery("name", "lee") + ":" + c.DefaultQuery("age", "100") + ":" + fmt.Sprintf("%v", sum)
//	c.JSON(http.StatusOK, &module.UserResponse{Detail: convert.String(detail)})
//}

//gorm:
//var (
//	//变量db 通过init直接初始化
//	db     *gorm.DB
//	err    error
//	dbinfo string
//)

//拼接数据库连接信息
//dbinfo = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "52991314wqq", "localhost", "3306", "goentrytask")
//
////初始化db
//db, err = gorm.Open("mysql", dbinfo)
//if err != nil {
//	fmt.Println("mysql打开失败", err)
//	return
//}

//type User struct {
//	ID       int64  `gorm:"column:id" json:"id"`
//	Unmae    string `gorm:"uname" json:"unmae"`
//	Nickname string `gorm:"nickname" json:"nickname"`
//}
