package service

import (
	"github.com/Surpass-w/htpasswd/pkg/htpasswd"
)

func Run(user, passwd, filePath string) error {
	// 生成密码文件
	return htpasswd.SetPassword(filePath, user, passwd, htpasswd.HashBCrypt)
}
