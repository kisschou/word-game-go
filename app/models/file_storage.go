package models

import (
	"errors"
	"time"

	"wordgame/tdog/core"
	"wordgame/tdog/lib"
)

type (
	FileStorageModel struct {
		Base   core.Model
		IpAddr string
	}

	FileStorage struct {
		Id         int64     `xorm:"pk default 0 BIGINT(20)"`
		UserId     int64     `xorm:"not null default 0 comment('文件所有者用户id') index BIGINT(20)"`
		RealName   string    `xorm:"not null default '' comment('文件真实的名称') VARCHAR(255)"`
		Name       string    `xorm:"not null default '' comment('文件名') VARCHAR(255)"`
		Suffix     string    `xorm:"not null default '' comment('后缀') VARCHAR(16)"`
		Path       string    `xorm:"not null default '' comment('路径') VARCHAR(255)"`
		Md5        string    `xorm:"not null default '' comment('文件的MD5值') VARCHAR(32)"`
		Sha1       string    `xorm:"not null default '' comment('文件的SHA1值') VARCHAR(40)"`
		Mime       string    `xorm:"not null comment('文件的MIME值') VARCHAR(50)"`
		Type       int       `xorm:"not null default 0 comment('类型') TINYINT(3)"`
		Size       int64     `xorm:"not null default 0 comment('大小') BIGINT(20)"`
		Operate    string    `xorm:"not null default '' comment('操作人') VARCHAR(255)"`
		CreateTime time.Time `xorm:"not null default '1000-01-01 00:00:00' comment('创建日期') DATETIME"`
		IsOpen     int       `xorm:"not null default 0 comment('是否公有(0否1是)') TINYINT(3)"`
	}
)

func (fs *FileStorageModel) IsExists(userId int64, md5 string, sha1 string) bool {
	fileStorageInfo := new(FileStorage)
	fs.Base.Sql.NewEngine()
	existsCount, _ := fs.Base.Sql.Engine.Where("user_id=?", userId).Where("md5=?", md5).Where("sha1=?", sha1).Count(fileStorageInfo)
	if existsCount > 0 {
		return true
	}
	return false
}

func (fs *FileStorageModel) Modify(fileId int64, userId int64, isOpen int) (err error) {
	fileStorageInfo := new(FileStorage)
	fs.Base.Sql.NewEngine()
	existsCount, _ := fs.Base.Sql.Engine.Where("file_id=?", fileId).Where("user_id=?", userId).Count(fileStorageInfo)
	if existsCount < 1 {
		err = errors.New("ERROR_FILE_MISSING")
		return
	}
	fileStorageInfo = new(FileStorage)
	fileStorageInfo.IsOpen = isOpen
	var affected int64
	affected, err = fs.Base.Sql.Engine.Where("file_id=?", fileId).Where("user_id=?", userId).Update(fileStorageInfo)
	if affected < 1 && err != nil {
		LoggerLib := new(lib.Logger)
		LoggerLib.Level = 0
		LoggerLib.Key = "error"
		LoggerLib.New(err.Error())
		err = errors.New("ERROR_FILE_UPDATE_FAIL")
		return
	}

	err = nil
	return
}

func (fs *FileStorageModel) Add(fileInfo map[string]interface{}) (resInfo map[string]interface{}, err error) {
	if fs.IsExists(fileInfo["user_id"].(int64), fileInfo["md5"].(string), fileInfo["sha1"].(string)) {
		err = errors.New("ERROR_FILE_EXISTS")
		return
	}

	// 文件id - 雪花算法
	fileStorageInfo := new(FileStorage)
	fs.Base.Sql.NewEngine()
	SnowFlakeLib := new(lib.SnowFlake)
	UtilLib := new(lib.Util)
	SnowFlakeLib.MachineId = UtilLib.GetMachineId()
	SnowFlakeLib.SN = UtilLib.RandInt64(1000, 9999)
	SnowFlakeLib.LastTimeStamp = time.Now().UnixNano() / 1000000
	fileStorageInfo.Id = SnowFlakeLib.New()

	fileStorageInfo.UserId = fileInfo["user_id"].(int64)
	fileStorageInfo.RealName = fileInfo["file_name"].(string)
	fileStorageInfo.Name = fileInfo["save_name"].(string)
	fileStorageInfo.Suffix = fileInfo["ext"].(string)
	fileStorageInfo.Path = fileInfo["file_path"].(string) + fileStorageInfo.Name + "." + fileStorageInfo.Suffix
	fileStorageInfo.Md5 = fileInfo["md5"].(string)
	fileStorageInfo.Sha1 = fileInfo["sha1"].(string)
	fileStorageInfo.Mime = fileInfo["mime"].(string)
	fileStorageInfo.Type = fileInfo["type"].(int)
	fileStorageInfo.Size = fileInfo["size"].(int64)
	fileStorageInfo.Operate = ""
	fileStorageInfo.CreateTime = time.Now()
	fileStorageInfo.IsOpen = fileInfo["is_open"].(int)

	// 入库
	var affected int64
	affected, err = fs.Base.Sql.Engine.Insert(fileStorageInfo)
	if affected < 1 && err != nil {
		LoggerLib := new(lib.Logger)
		LoggerLib.Level = 0
		LoggerLib.Key = "error"
		LoggerLib.New(err.Error())
		err = errors.New("ERROR_FILE_ADD_FAIL")
		return
	}

	err = nil
	resInfo = fileStorageInfo.BuildRes()
	return
}

func (fs *FileStorageModel) GetPub(fileId int64) (fileInfo map[string]interface{}, err error) {
	fileStorageInfo := new(FileStorage)
	fs.Base.Sql.NewEngine()
	result, err := fs.Base.Sql.Engine.Where("id=?", fileId).Where("is_open=?", 1).Get(fileStorageInfo)
	if !result || err != nil {
		err = errors.New("ERROR_FILE_MISSING")
		return
	}

	fileInfo = fileStorageInfo.BuildRes()
	return
}

func (fs *FileStorageModel) GetPri(fileId int64, userId int64) (fileInfo map[string]interface{}, err error) {
	fileStorageInfo := new(FileStorage)
	fs.Base.Sql.NewEngine()
	result, err := fs.Base.Sql.Engine.Where("id=?", fileId).Where("user_id=?", userId).Get(fileStorageInfo)
	if !result || err != nil {
		err = errors.New("ERROR_FILE_MISSING")
		return
	}

	fileInfo = fileStorageInfo.BuildRes()
	return
}

func (fs *FileStorage) BuildRes() map[string]interface{} {
	ConfigLib := new(lib.Config)
	fileInfo := make(map[string]interface{})
	fileInfo["file_id"] = fs.Id
	fileInfo["real_name"] = fs.RealName
	fileInfo["path"] = ConfigLib.Get("file.file_request_url").String() + fs.Path
	fileInfo["md5"] = fs.Md5
	fileInfo["sha1"] = fs.Sha1
	fileInfo["size"] = fs.Size
	return fileInfo
}
