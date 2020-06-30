package lib

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	File struct {
		Filename   string              // 文件名
		Header     map[string][]string // 文件信息
		Size       int64               // 文件大小
		Body       multipart.File      // 文件数据
		Savename   string              // 存储名
		Ext        string              // 后缀
		FilePathAP string              // 存储目录(绝对路径)
		FilePath   string              // 存储目录(相对路径)
		Md5        string              // 文件md5值
		Sha1       string              // 文件sha1值
		MIME       string              // 文件的MIME值
		Type       int                 // 文件类型
	}
)

const (
	FILE_TYPE_OTHER = iota
	FILE_TYPE_PICTURE
	FILE_TYPE_DOCUMENT
	FILE_TYPE_MUSIC
	FILE_TYPE_VIDEO
)

func (file *File) GetSaveName() *File {
	CryptLib := new(Crypt)
	filename := ""
	filenameSplit := strings.Split(file.Filename, ".")
	file.Ext = filenameSplit[len(filenameSplit)-1]
	filenameSplit = filenameSplit[:len(filenameSplit)-1]
	filename = strings.Join(filenameSplit, "_")
	CryptLib.Str = filename + strconv.FormatInt(time.Now().UnixNano(), 10)
	file.Savename = CryptLib.Sha256()
	return file
}

func (file *File) GetFilePath() *File {
	file.MIME = file.Header["Content-Type"][0]

	ConfigLib := new(Config)
	UtilLib := new(Util)
	// 创建存储目录
	filePath := ConfigLib.Get("file.file_save_path").String()
	file.FilePath = "/data/upload/" + file.MIME + "/" + time.Now().Format("2006-01-02") + "/"
	filePath += file.FilePath
	UtilLib.DirExistsAndCreate(filePath)

	file.FilePathAP = filePath
	return file
}

func (file *File) GetFileType() *File {
	fileType := 0
	mimeSplit := strings.Split(file.MIME, "/")
	switch mimeSplit[0] {
	case "image":
		fileType = FILE_TYPE_PICTURE
		break

	case "video":
		fileType = FILE_TYPE_VIDEO
		break

	case "audio":
		fileType = FILE_TYPE_MUSIC
		break

	default:
		fileType = FILE_TYPE_OTHER
		break
	}

	UtilLib := new(Util)
	docMime := []string{
		"application/msword",
		"application/vnd.ms-excel",
		"application/pdf",
		"application/vnd.ms-powerpoint",
		"text/plain",
	}
	if UtilLib.InStringSlice(file.MIME, docMime) {
		fileType = FILE_TYPE_DOCUMENT
	}

	file.Type = fileType
	return file
}

func (file *File) GetFileHash() *File {
	f, err := os.Open(file.FilePathAP + file.Savename + "." + file.Ext)
	if err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		return file
	}
	defer f.Close()
	hash := md5.New()
	io.Copy(hash, f)
	file.Md5 = hex.EncodeToString(hash.Sum(nil))
	hash = sha1.New()
	io.Copy(hash, f)
	file.Sha1 = hex.EncodeToString(hash.Sum(nil))
	return file
}

func (file *File) BuildRes() map[string]interface{} {
	fileSaveRes := make(map[string]interface{})
	fileSaveRes["file_name"] = file.Filename
	fileSaveRes["save_name"] = file.Savename
	fileSaveRes["ext"] = file.Ext
	fileSaveRes["file_path_ap"] = file.FilePathAP
	fileSaveRes["file_path"] = file.FilePath
	fileSaveRes["md5"] = file.Md5
	fileSaveRes["sha1"] = file.Sha1
	fileSaveRes["mime"] = file.MIME
	fileSaveRes["type"] = file.Type
	fileSaveRes["size"] = file.Size
	return fileSaveRes
}

func (file *File) Save() map[string]interface{} {
	// 获取文件基础保存数据
	file = file.GetSaveName().GetFilePath().GetFileType()

	// 创建存储文件
	f, err := os.OpenFile(file.FilePathAP+file.Savename+"."+file.Ext, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		return nil
	}
	defer f.Close()

	// Copy
	io.Copy(f, file.Body)

	// 获取文件的sha1和md5数据
	return file.GetFileHash().BuildRes()
}

func (file *File) Del(delFile string) bool {
	err := os.Remove(delFile)
	if err != nil {
		logger := Logger{Level: 0, Key: "error"}
		logger.New(err.Error())
		return false
	}
	return true
}
