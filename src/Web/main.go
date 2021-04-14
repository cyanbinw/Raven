package main

import (
	"Raven/src/Web/Router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	Router.InitRouter()
}
