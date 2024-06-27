package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Surpass-w/htpasswd/internel/model"
	"github.com/Surpass-w/htpasswd/internel/service"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

const (
	DefaultUser         = "admin"
	DefaultPasswdLength = 16
)

var rootCmd = &cobra.Command{
	Use:           "htpasswd",
	Short:         "htpasswd is a tool to help generate htpasswd files.",
	SilenceErrors: true,
	SilenceUsage:  true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func SetVersion(version string) {
	rootCmd.SetVersionTemplate(`{{printf "Version: %s" .Version}}`)
	rootCmd.Version = version
}

func init() {
	var (
		GlobalOptions model.GenerateReqV1
		debug         bool
	)
	rootCmd.Flags().SortFlags = false
	rootCmd.Flags().StringVarP(&GlobalOptions.User, "user", "u", DefaultUser, "user")
	rootCmd.Flags().StringVarP(&GlobalOptions.Password, "password", "p", "", "passwd")
	rootCmd.Flags().StringVarP(&GlobalOptions.Path, "file", "f", "", "store path")
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "print request params")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		// 判断密码是为空,则生成随机密码
		if GlobalOptions.Password == "" {
			passwd, err := service.GenerateRandomPassword(DefaultPasswdLength)
			if err != nil {
				logrus.Errorf("generate random password failed: %s\n", err.Error())
				return err
			}
			GlobalOptions.Password = passwd
			fmt.Println(GlobalOptions.Password)
		}
		if debug {
			data, _ := json.MarshalIndent(GlobalOptions, "", "    ")
			fmt.Println(string(data))
		}
		dir := filepath.Dir(GlobalOptions.Path)
		info, err := os.Stat(dir)
		if err != nil {
			if os.IsNotExist(err) {
				err = os.MkdirAll(dir, os.ModePerm)
				if err != nil {
					return err
				}
			}
			return err
		}
		if !info.IsDir() {
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				return err
			}
		}

		f, err := os.OpenFile(GlobalOptions.Path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "OpenFile")
		}
		defer f.Close()
		return service.Run(GlobalOptions.User, GlobalOptions.Password, GlobalOptions.Path)
	}
}
