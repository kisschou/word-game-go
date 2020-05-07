package main

import (
	"wordgame/app/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run(":8000")
}
