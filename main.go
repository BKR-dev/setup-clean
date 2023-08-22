package main

import (
	"fmt"
	"github.com/bkr-dev/setup-clean/cmd"
)

func main() {
	fmt.Println("hell yeah brother")
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
