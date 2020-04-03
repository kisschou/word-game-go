package unit

import (
	"encoding/json"
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
