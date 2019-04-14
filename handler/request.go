package handler

import (
	"dlser/common"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"os/exec"
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
	fmt.Println(reqBody["addr"].(string))

	// 创建命令
	cmd := exec.Command("wget", reqBody["addr"].(string), "-P", "./download")

	// 执行命令
	err = cmd.Start()
	if err != nil {
		res := common.HandleRes(-1, "错误", nil, err)
		fmt.Fprintln(w, res)
		return
	}

	res := common.HandleRes(0, "开始下载", nil, nil)
	fmt.Fprintln(w, res)
}
