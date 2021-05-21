package utils

import "encoding/json"

func JSON(a interface{}) string {
	if a == nil {
		return ""
	}
	b, _ := json.Marshal(a)
	return string(b)
}
