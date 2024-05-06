package cmd

import (
	"bytes"
	"fmt"
	"github.com/foomo/htpasswd"
	"os/exec"
)

func run(user, passwd, filePath string) error {
	// 生成密码文件
	if passwd == "" {
		// 通过 openssl 生成密码
		c := exec.Command("openssl", "rand", "-base64", "16")
		var out bytes.Buffer
		c.Stdout = &out
		err := c.Run()
		if err != nil {
			return err
		}
		passwd = out.String()
		fmt.Println("passwd:", passwd)
	}
	return htpasswd.SetPassword(filePath, user, passwd, htpasswd.HashBCrypt)
}
