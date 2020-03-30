package main

import (
	"os"
	_ "wordgame/config"
	_ "wordgame/library/database"
	"wordgame/routers"
)

func main() {
	r := routers.InitRouter()
	port := os.Getenv("HTTP_PORT")
	r.Run(":" + port)
}
