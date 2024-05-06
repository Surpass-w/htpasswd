package cmd

import (
	"github.com/foomo/htpasswd"
)

func run(user, passwd, filePath string) error {
	// 生成密码文件
	return htpasswd.SetPassword(filePath, user, passwd, htpasswd.HashBCrypt)
}
