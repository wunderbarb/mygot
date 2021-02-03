// v0.2.3
// (C) DIEHL E., Jan 2021

package main

import (

	"fmt"
	"os"

	"github.com/wunderbarb/mygot/internal/common"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("needs a file name")
		os.Exit(1)
	}

	err := common.CreateHeader(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

