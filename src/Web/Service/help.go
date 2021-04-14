package Service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"xorm.io/core"
)

type ConnectString struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Ip       string `yaml:"ip"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbname"`
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

	v := SetConnectString()

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{v.UserName, ":", v.Password, "@tcp(", v.Ip, ":", v.Port, ")/", v.DbName, "?charset=utf8"}, "")

	engine, err := xorm.NewEngine("mysql", path)
	if err != nil {
		fmt.Println(err)
	}
	engine.SetMapper(core.SameMapper{})
	return engine
}

func SetConnectString() ConnectString {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	item := ConnectString{}
	data, err := ioutil.ReadFile(dir + "/config.yml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(data, &item)
	if err != nil {
		fmt.Println(err)
	}
	return item
}
