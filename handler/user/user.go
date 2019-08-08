package user

import "apiServerDemo/model"

// 单个用户请求格式
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 单个用户响应格式
type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
