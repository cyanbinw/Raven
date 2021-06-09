package main

import (
	"Raven/src/log"
	"Raven/src/web/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Init()
	router.InitRouter()
}
