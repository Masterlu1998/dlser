package common

type ResObj struct {
	Code int `json:"code"`
	Prompt string `json:"prompt"`
	Obj interface{} `json:"obj"`
	Err string `json:"err"`
}

// 错误结构体格式
type Info struct {
	Msg string
	Code int
}

var (
	//----------------------------------------------------------
	// 错误的信息
	// 请求读取错误
	ReadRequestErrInfo = Info{ "请求读取错误", -101 } 

	// JSON解析错误
	JSONParseErrInfo = Info{ "JSON解析错误", -102 }

	// 文件已经被删除
	FileHasBeenDeletedErrInfo = Info{ "文件已经被删除", -103 }

	//-----------------------------------------------------------
	// 正确的信息
	// 数据库查询正确
	FindSuccessInfo = Info{ "查询成功", 100 }
	
	// 下载文件成功
	DownloadFileSuccessInfo = Info { "文件开始下载", 101 }

	// 删除文件成功
	DeleteFileSuccessInfo = Info { "删除文件成功", 102 }
)

