package unit

import (
	"encoding/json"
	"fmt"
	"os"
)

func Struct2Map(obj interface{}) (output map[string]interface{}) {
	j, _ := json.Marshal(obj)
	json.Unmarshal(j, &output)
	return
}

func Map2Struct(obj map[string]interface{}) (output interface{}) {
	j, _ := json.Marshal(obj)
	json.Unmarshal(j, &output)
	return
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func DirExistsAndCreate(path string) {
	if !FileExists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Mkdir path fail: ", err)
		}
	}
}
