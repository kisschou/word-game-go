package main

import (
	"wordgame/library/config"
	_ "wordgame/library/database"
	"wordgame/routers"
)

func main() {
	r := routers.InitRouter()
	c := config.Config{}
	port := c.Get("http_port").(string)
	r.Run(":" + port)
}
