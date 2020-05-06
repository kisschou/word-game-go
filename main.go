package main

import (
	// "wordgame/app/routers"
	"wordgame/tdog/lib"
)

func main() {
	// routers.InitRouter().Run()
	var mysql lib.MySql
	mysql.NewEngine()
}
