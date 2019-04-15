package handler

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"dlser/common"
	"dlser/mysql"
)

func GetTaskList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
		resJSON := common.HandleRes(-1, "读取失败", nil, nil)
		fmt.Fprintln(w, resJSON)
		return
	}
	var reqObj map[string]interface{}
	json.Unmarshal(data, &reqObj)
	index, pageSize, keywords := int(reqObj["index"].(float64)), int(reqObj["pageSize"].(float64)), reqObj["keywords"].(string)
	dlTask := new(mysql.DlTask)
	dlTasks := dlTask.FindTaskList(index, pageSize, keywords)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	resObj := map[string]interface{}{ "taskList": dlTasks }
	resJSON := common.HandleRes(0, "查询成功", resObj, nil)
	fmt.Fprintln(w, resJSON)

}