package Service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"os"
	"strings"
	"xorm.io/core"
)

const (
	userName = "sa"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "test_db"
)

var engine *xorm.Engine

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// CheckFileIsExist 判断文件存不存在
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func InitDB() *xorm.Engine {
	if engine != nil {
		return engine
	}
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	var err error = nil
	engine, err = xorm.NewEngine("mysql", path)
	if err != nil {
		fmt.Println(err)
	}
	engine.SetMapper(core.SameMapper{})
	return engine
}
