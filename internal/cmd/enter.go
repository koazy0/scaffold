package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"moyu/internal/cmd/migrate"
	"moyu/internal/cmd/server"
	"moyu/internal/cmd/user"
	"moyu/internal/common"
	"os"
)

var logger = common.Logs().Cat("cmd")
var rootCmd = &cobra.Command{
	Use:   "honey",
	Short: "deal with user model",
	Long:  `this is a long description with user command`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		logger.Infoln("hello user")
	},
}

func init() {
	rootCmd.AddCommand(
		migrate.MigrateCommand(),
		server.ServerCommand(),
		user.UserCommand(),
	)
}
func RootCommand() *cobra.Command {
	return rootCmd
}

// Execute 对main函数暴露出来的方法
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type FlagOptions struct {
	File    string
	Version bool
	DB      bool
	Menu    string
	Type    string
	Value   string
	Help    bool
}

var options FlagOptions
