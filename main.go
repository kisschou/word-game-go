package main

import (
	"wordgame/app/routers"
)

func main() {
	routers.InitRouter().Run()
}
