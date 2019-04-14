package common

import (
	"encoding/json"
)

// func handleErr(err error, resJson ResObj) {
// 	if err != nil {
// 		fmt.Println("解析错误")
// 		resJson.Code = -1
// 		resJson.Prompt = "解析错误"
// 		resJson.Err = err.Error()
// 		res, _ := json.Marshal(resJson)
// 	} else {

// 	}
// }

func HandleRes(code int, prompt string, obj map[string]interface{}, err error) string {
	var resObj ResObj = ResObj{ Code: code, Prompt: prompt }
	if code > 0 {
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
	resJson, _ := json.Marshal(resObj)
	return string(resJson)
}