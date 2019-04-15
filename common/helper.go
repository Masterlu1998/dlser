package common

import (
	"encoding/json"
)

func HandleRes(code int, prompt string, obj map[string]interface{}, err error) string {
	var resObj = ResObj{ Code: code, Prompt: prompt }
	if code >= 0 {
		// 成功返回
		if obj != nil {
			resObj.Obj = obj
		}
	} else {
		// 错误返回
		if err != nil {
			resObj.Err = err.Error()
		}
	}
	resJSON, _ := json.Marshal(resObj)
	return string(resJSON)
}