package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var rootCmd = &cobra.Command{
	Use:           "htpasswd",
	Short:         "htpasswd is a simple utility to create and update htpasswd files.",
	SilenceErrors: true,
	SilenceUsage:  true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		if GlobalOptions.Password == "" {
			// 通过 openssl 生成密码
			GlobalOptions.Password = "ksp123456"
		}
		return nil
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.PersistentFlags().StringVarP(&GlobalOptions.User, "user", "u", "admin", "user")
	rootCmd.PersistentFlags().StringVarP(&GlobalOptions.Password, "password", "p", "", "passwd")
	rootCmd.PersistentFlags().StringVarP(&GlobalOptions.FilePath, "file", "f", "", "file path")
	rootCmd.PersistentFlags().StringVarP(&GlobalOptions.Dir, "dir", "d", "", "dir path")
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		user := GlobalOptions.User
		passwd := GlobalOptions.Password
		filePath := GlobalOptions.FilePath
		dir := &GlobalOptions.Dir
		if *dir != "" {
			filePath = filepath.Join(*dir, ".htpasswd")
		} else {
			_, err := os.Stat(filePath)
			if err == nil {
				err = os.Remove(filePath)
				if err != nil {
					return errors.Errorf("failed to remove the binary file: %s, err: %s", filePath, err.Error())
				}
			}
			fileDir := filepath.Dir(filePath)
			err = os.MkdirAll(fileDir, os.ModePerm)
			if err != nil {
				return errors.Errorf("failed to create directory: %s, err: %s", fileDir, err.Error())
			}
		}

		f, err := os.Create(filePath)
		if err != nil {
			return errors.Errorf("failed to create file: %s, err: %s", filePath, err.Error())
		}
		defer f.Close()
		return run(user, passwd, filePath)
	}
}
