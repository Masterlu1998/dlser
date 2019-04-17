package handler

import (
	"dlser/common"
	"dlser/execute"
	"dlser/mysql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RequestHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	// 赋值
	url, name := reqBody["addr"].(string), reqBody["name"].(string)

	// 开启协程执行下载任务
	go execute.AddTask(mysql.DlTask{Addr: url, Name: name})

	// 返回响应
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	res := common.HandleRes(0, "开始下载", nil, nil)
	fmt.Fprintln(w, res)
}
