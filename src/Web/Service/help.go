package Service

import (
	"encoding/json"
	"fmt"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"os"
	"strings"
	"xorm.io/core"
)

type ConnectString struct {
	userName string
	password string
	ip       string
	port     string
	dbName   string
}

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
	JsonParse := NewJsonStruct()
	v := ConnectString{}
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	JsonParse.Load("./config.json", &v)
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{v.userName, ":", v.password, "@tcp(", v.ip, ":", v.port, ")/", v.dbName, "?charset=utf8"}, "")
	var err error = nil
	engine, err = xorm.NewEngine("mysql", path)
	if err != nil {
		fmt.Println(err)
	}
	engine.SetMapper(core.SameMapper{})
	return engine
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
