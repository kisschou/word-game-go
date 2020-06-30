package routers

import (
	"wordgame/app/controllers"
	"wordgame/tdog/core"
)

func InitRouter() *core.HttpEngine {
	r := core.NewEngine()

	var FileController controllers.File
	fileRouter := r.Group("/file", &FileController.Base)
	{
		fileRouter.POST("/pub", FileController.Pub)
		fileRouter.POST("/pri", FileController.Pri)
		fileRouter.POST("/upload", FileController.Upload)
		fileRouter.POST("/modify", FileController.Modify)
	}

	return r
}
