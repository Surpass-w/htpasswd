package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
		if GlobalOptions.FilePath == "" {
			return fmt.Errorf("file path is empty")
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
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		user := GlobalOptions.User
		passwd := GlobalOptions.Password
		filePath := GlobalOptions.FilePath
		_ = os.Remove(filePath)
		return run(user, passwd, filePath)
	}
}
