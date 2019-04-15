package common

type ResObj struct {
	Code int `json:"code"`
	Prompt string `json:"prompt"`
	Obj interface{} `json:"obj"`
	Err string `json:"err"`
}

