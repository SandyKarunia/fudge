package main

import (
	"github.com/sandykarunia/fudge/server"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	server.Instance().Start()
}
