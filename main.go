package main

import (
	_ "github.com/icattlecoder/godaemon"
	"wordgame/app/routers"
)

func main() {
	routers.InitRouter().Run()
}
