package main

import (
	"Raven/src/Log"
	"Raven/src/Web/Router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	Log.Init()
	Router.InitRouter()
}
