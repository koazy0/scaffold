package user

import (
	"github.com/spf13/cobra"
	"scaffold/internal/common"
)

var logger = common.Logs().Cat("cmd/user")

func UserCommand() *cobra.Command {
	return userCmd
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "deal with user model",
	Long:  `this is a long description with user command`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		logger.Infoln("hello user")
		logger.Infoln(cmd.Use)
		logger.Infoln(cmd.Short)
		logger.Infoln(cmd.Usage())
		cmd.SilenceUsage = true
		logger.Infoln(cmd.Usage())
		cmd.SilenceErrors = true
		logger.Infoln(cmd.Usage())

		name, _ := cmd.Flags().GetString("name")
		logger.Infoln(name)
		logger.Infoln(len(args))
		logger.Infoln(args)

	},
}
