package Service

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	DB ConnectString `yaml:"db"`
}

type ConnectString struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Ip       string `yaml:"ip"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbname"`
}

var item *Config

func getConfig() {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(data, &item)
	if err != nil {
		fmt.Println(err)
	}
}

func GetConnectString() ConnectString {
	if item == nil {
		getConfig()
	}
	return item.DB
}
