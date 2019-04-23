package common

import (
	"encoding/json"
	"net/http"
	"fmt"
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

func WriteHeader(w http.ResponseWriter, key string, val string) {
	w.Header().Add(key, val)
}

func HandleErr(w http.ResponseWriter, code int, err error, prompt string) {
	fmt.Println(err)
	resJSON := HandleRes(code, prompt, nil, err)
	fmt.Fprintln(w, resJSON)
}

func HandleSuc(w http.ResponseWriter, code int, obj map[string]interface{}, prompt string) {
	fmt.Println(obj)
	resJSON := HandleRes(code, prompt, obj, nil)
	fmt.Fprintln(w, resJSON)
}