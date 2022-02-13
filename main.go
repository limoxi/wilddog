package main

import (
	"flag"
	"github.com/limoxi/ghost"

	_ "wilddog/api"
	_ "wilddog/db"
	_ "wilddog/middleware"
)

var mod = flag.String("mod", "vendor", "")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 && args[0] == "sync" {
		ghost.SyncDB()
		return
	}
	ghost.RunWebServer()
}
