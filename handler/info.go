package handler

import (
	"dlser/common"
	"dlser/mysql"
	"encoding/json"
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
		common.HandleErr(w, common.ReadRequestErrInfo.Code, err, common.ReadRequestErrInfo.Msg)
		return
	}
	defer r.Body.Close()
	var reqObj struct {
		Index     int       `json:"index"`
		PageSize  int       `json:"pageSize"`
		Keywords  string    `json:"keywords"`
		StartTime time.Time `json:"startTime"`
		EndTime   time.Time `json:"endTime"`
	}
	// req := reqObj{}
	err = json.Unmarshal(data, &reqObj)
	if err != nil {
		common.HandleErr(w, common.JSONParseErrInfo.Code, err, common.JSONParseErrInfo.Msg)
		return
	}
	// 赋值参数
	index, pageSize, keywords, startTime, endTime := reqObj.Index, reqObj.PageSize, reqObj.Keywords, reqObj.StartTime, reqObj.EndTime

	// 调用接口获取查询结果
	dlTask := new(mysql.DlTask)
	dlTasks := dlTask.FindTaskInfoList(index, pageSize, keywords, startTime, endTime)

	// 返回响应
	resObj := map[string]interface{}{"taskList": dlTasks}
	common.HandleSuc(w, common.FindSuccessInfo.Code, resObj, common.FindSuccessInfo.Msg)
}
