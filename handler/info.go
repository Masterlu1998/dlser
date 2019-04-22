package handler

import (
	"dlser/common"
	"dlser/mysql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func GetTaskListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 先设置响应头
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 读取请求参数
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		resJSON := common.HandleRes(-1, "读取失败", nil, nil)
		fmt.Fprintln(w, resJSON)
		return
	}
	defer r.Body.Close()
	type reqObj struct {
		Index     int       `json:"index"`
		PageSize  int       `json:"pageSize"`
		Keywords  string    `json:"keywords"`
		StartTime time.Time `json:"startTime"`
		EndTime   time.Time `json:"endTime"`
	}
	req := reqObj{}
	json.Unmarshal(data, &req)

	// 赋值参数
	index, pageSize, keywords, startTime, endTime := req.Index, req.PageSize, req.Keywords, req.StartTime, req.EndTime

	// 调用接口获取查询结果
	dlTask := new(mysql.DlTask)
	dlTasks := dlTask.FindTaskInfoList(index, pageSize, keywords, startTime, endTime)

	// 设置响应头，返回响应
	resObj := map[string]interface{}{"taskList": dlTasks}
	resJSON := common.HandleRes(0, "查询成功", resObj, nil)
	fmt.Fprintln(w, resJSON)
}
