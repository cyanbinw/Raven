package common

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	BusinessDB  ConnectString `yaml:"business_db"`
	LogDB       ConnectString `yaml:"log_db"`
	DefaultUser DefaultUser   `yaml:"default_user"`
}

type ConnectString struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Ip       string `yaml:"ip"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbname"`
}

type DefaultUser struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
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

func GetBusinessConnectString() ConnectString {
	if item == nil {
		getConfig()
	}
	return item.BusinessDB
}

func GetLogConnectString() ConnectString {
	if item == nil {
		getConfig()
	}
	return item.LogDB
}

func GetDefaultUserString() DefaultUser {
	if item == nil {
		getConfig()
	}
	return item.DefaultUser
}
