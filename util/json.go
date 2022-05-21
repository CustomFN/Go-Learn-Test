package util

import (
	"encoding/json"
	"fmt"
)

func Marshal(input interface{}) string {
	marshal, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(marshal)
}
