package user

// 请求格式
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 响应格式
type CreateResponse struct {
	Username string `json:"username"`
}