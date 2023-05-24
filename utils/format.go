package utils

import "encoding/json"

func ObjToString(o interface{}) string {
	b, _ := json.MarshalIndent(o, "", "  ")
	return string(b)
}
