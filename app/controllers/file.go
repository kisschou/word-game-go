package controllers

import (
	"net/http"
	"strconv"

	"wordgame/app/models"
	"wordgame/tdog/core"
	"wordgame/tdog/lib"
)

type (
	File struct {
		Base core.Controller
	}
)

// swagger:operation POST /file/pub file pub
// ---
// summary: 获取公共文件
// description: 获取公共文件
// parameters:
// - name: file_id
//   in: body
//   description: 文件id
//   type: string
//   required: true
// responses:
//   200: repoResp
//   401: badReq
func (file *File) Pub() {
	var fileId int64
	if _, ok := file.Base.Req.Params["file_id"]; ok {
		if len(file.Base.Req.Params["file_id"]) > 0 {
			var err error
			fileId, err = strconv.ParseInt(file.Base.Req.Params["file_id"][0], 10, 64)
			if err != nil || fileId < 1 {
				file.Base.Res.JSON(http.StatusInternalServerError, core.H{
					"code": "ERROR_REQUEST_PARAMS",
				})
				return
			}
		}
	}

	FileStorageModel := new(models.FileStorageModel)
	fileInfo, err := FileStorageModel.GetPub(fileId)
	if err != nil {
		file.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": err.Error(),
		})
		return
	}

	file.Base.Res.JSON(http.StatusOK, core.H{
		"message": "success",
		"data":    fileInfo,
	})
}

// swagger:operation POST /file/pri file pri
// ---
// summary: 获取私有文件
// description: 获取私有文件
// parameters:
// - name: file_id
//   in: body
//   description: 文件id
//   type: string
//   required: true
// - name: user_id
//   in: body
//   description: 用户id
//   type: string
//   required: true
// responses:
//   200: repoResp
//   401: badReq
func (file *File) Pri() {
	var fileId, userId int64
	if _, ok := file.Base.Req.Params["file_id"]; ok {
		if len(file.Base.Req.Params["file_id"]) > 0 {
			var err error
			fileId, err = strconv.ParseInt(file.Base.Req.Params["file_id"][0], 10, 64)
			if err != nil || fileId < 1 {
				file.Base.Res.JSON(http.StatusInternalServerError, core.H{
					"code": "ERROR_REQUEST_PARAMS",
				})
				return
			}
		}
	}

	if _, ok := file.Base.Req.Params["user_id"]; ok {
		if len(file.Base.Req.Params["user_id"]) > 0 {
			var err error
			userId, err = strconv.ParseInt(file.Base.Req.Params["user_id"][0], 10, 64)
			if err != nil || userId < 1 {
				file.Base.Res.JSON(http.StatusInternalServerError, core.H{
					"code": "ERROR_REQUEST_PARAMS",
				})
				return
			}
		}
	}

	FileStorageModel := new(models.FileStorageModel)
	fileInfo, err := FileStorageModel.GetPri(fileId, userId)
	if err != nil {
		file.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": err.Error(),
		})
		return
	}

	file.Base.Res.JSON(http.StatusOK, core.H{
		"message": "success",
		"data":    fileInfo,
	})
}

// swagger:operation POST /file/upload file upload
// ---
// summary: 文件上传
// description: 文件上传
// parameters:
// - name: file
//   in: body
//   description: 上传的文件资源
//   type: string
//   required: true
// - name: user_id
//   in: body
//   description: 用户id
//   type: string
//   required: true
// - name: is_open
//   in: body
//   description: 是否开放(0否1是)
//   type: int
//   required: false
// responses:
//   200: repoResp
//   401: badReq
func (file *File) Upload() {
	var userId int64
	isOpen := 0

	if _, ok := file.Base.Req.Params["user_id"]; ok {
		if len(file.Base.Req.Params["user_id"]) > 0 {
			var err error
			userId, err = strconv.ParseInt(file.Base.Req.Params["user_id"][0], 10, 64)
			if err != nil || userId < 1 {
				file.Base.Res.JSON(http.StatusInternalServerError, core.H{
					"code": "ERROR_REQUEST_PARAMS",
				})
				return
			}
		}
	}

	if _, ok := file.Base.Req.Params["is_open"]; ok {
		if len(file.Base.Req.Params["is_open"]) > 0 {
			var err error
			isOpen, err = strconv.Atoi(file.Base.Req.Params["is_open"][0])
			if err != nil || isOpen != 1 {
				isOpen = 0
			}
		}
	}

	// 文件处理
	FileLib := new(lib.File)
	FileLib.Filename = file.Base.Req.File.Filename
	FileLib.Header = file.Base.Req.File.Header
	FileLib.Size = file.Base.Req.File.Size
	FileLib.Body = file.Base.Req.File.Body
	fileInfo := FileLib.Save()
	fileInfo["user_id"] = userId
	fileInfo["is_open"] = isOpen

	FileStorageModel := new(models.FileStorageModel)
	FileStorageModel.IpAddr = file.Base.Req.IpAddr
	resInfo, err := FileStorageModel.Add(fileInfo)
	if err != nil {
		// 删除文件
		FileLib.Del(fileInfo["file_path_ap"].(string) + fileInfo["save_name"].(string) + "." + fileInfo["ext"].(string))

		file.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": err.Error(),
		})
		return
	}

	file.Base.Res.JSON(http.StatusOK, core.H{
		"message": "success",
		"data":    resInfo,
	})
}

// swagger:operation POST /file/modify file modify
// ---
// summary: 修改文件信息
// description: 修改文件信息
// parameters:
// - name: file_id
//   in: body
//   description: 文件id
//   type: string
//   required: true
// - name: user_id
//   in: body
//   description: 用户id
//   type: string
//   required: true
// - name: is_open
//   in: body
//   description: 是否开放(0否1是)
//   type: int
//   required: true
// responses:
//   200: repoResp
//   401: badReq
func (file *File) Modify() {
	var fileId, userId int64
	isOpen := 0

	if _, ok := file.Base.Req.Params["file_id"]; ok {
		if len(file.Base.Req.Params["file_id"]) > 0 {
			var err error
			fileId, err = strconv.ParseInt(file.Base.Req.Params["file_id"][0], 10, 64)
			if err != nil || fileId < 1 {
				file.Base.Res.JSON(http.StatusInternalServerError, core.H{
					"code": "ERROR_REQUEST_PARAMS",
				})
				return
			}
		}
	}

	if _, ok := file.Base.Req.Params["user_id"]; ok {
		if len(file.Base.Req.Params["user_id"]) > 0 {
			var err error
			userId, err = strconv.ParseInt(file.Base.Req.Params["user_id"][0], 10, 64)
			if err != nil || userId < 1 {
				file.Base.Res.JSON(http.StatusInternalServerError, core.H{
					"code": "ERROR_REQUEST_PARAMS",
				})
				return
			}
		}
	}

	if _, ok := file.Base.Req.Params["is_open"]; ok {
		if len(file.Base.Req.Params["is_open"]) > 0 {
			var err error
			isOpen, err = strconv.Atoi(file.Base.Req.Params["is_open"][0])
			if err != nil || isOpen != 1 {
				isOpen = 0
			}
		}
	}

	FileStorageModel := new(models.FileStorageModel)
	FileStorageModel.IpAddr = file.Base.Req.IpAddr
	err := FileStorageModel.Modify(fileId, userId, isOpen)
	if err != nil {
		file.Base.Res.JSON(http.StatusInternalServerError, core.H{
			"code": err.Error(),
		})
		return
	}
	file.Base.Res.JSON(http.StatusOK, core.H{
		"message": "success",
	})
}

func (file *File) Ping() {
	file.Base.Res.String(http.StatusOK, "Pong")
}
