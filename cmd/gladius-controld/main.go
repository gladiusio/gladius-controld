package main

import (
	"flag"

	"github.com/gladiusio/gladius-controld/controld"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()
	controld.InitializeNodeManager()
}
