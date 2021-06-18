package main

import (
	"Raven/src/log"
	"Raven/src/web/router"
	"bytes"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
)

func main() {
	//runCommand()
	log.Init()
	router.InitRouter()
}

func runCommand() {
	cmd := exec.Command("swag", "init")
	fmt.Println("Cmd", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out.String())
}
