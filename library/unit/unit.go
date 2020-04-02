package unit

import (
	"encoding/json"
)

func Struct2Map(newStruct *struct{}) (newMap map[string]interface{}) {
	b, _ := json.Marshal(newStruct)
	json.Unmarshal([]byte(string(b)), &newMap)
	return
}
