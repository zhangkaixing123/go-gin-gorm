package api

// 定义请求参数结构体

type TestReq struct {
	Name string `form:"name" json:"name"` // 绑定json参数
	Page int    `form:"page" json:"page"` // 绑定url参数
}
