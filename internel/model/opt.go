package model

// GenerateReqV1 生成密码请求
type GenerateReqV1 struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Path     string `json:"path"`
}
