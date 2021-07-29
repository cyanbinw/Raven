package main

import (
	"bytes"
	"fmt"
	"github.com/WFallenDown/Raven/src/web/router"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
)

func main() {
	//runCommand()
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
