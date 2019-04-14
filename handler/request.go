package handler

import (
	"dlser/common"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"dlser/execute"
	"strings"
)

func RequestHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 定义返回结构
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	fmt.Println("收到下载请求！")
	// 解析请求
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res := common.HandleRes(-1, "解析错误", nil, err)
		fmt.Fprintln(w, res)
		return
	}

	var reqBody map[string]interface{}
	json.Unmarshal(data, &reqBody)
	url := reqBody["addr"].(string)
	urlSlice := strings.Split(url, "/")
	fileName := urlSlice[len(urlSlice) - 1]

	go execute.AddTask(execute.DlTask{ Addr: url, Name: fileName })

	res := common.HandleRes(0, "开始下载", nil, nil)
	fmt.Fprintln(w, res)
}
