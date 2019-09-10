package main

import (
	"fmt"
	"gitlab.com/sandykarunia/fudge/server"
	"gitlab.com/sandykarunia/fudge/utils"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Check whether the we are in sudo environment or not
	if !utils.IsSudo() {
		fmt.Println("Please run this program as root, we need root to run the isolate binary")
		os.Exit(1)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	server.Start()
}
