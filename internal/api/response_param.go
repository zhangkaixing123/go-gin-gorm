package api

// 定义请求参数结构体

type TestResp struct {
	Name string `json:"name,omitempty"` // 绑定json参数
	Page int    `json:"page"`           // 绑定url参数
}
