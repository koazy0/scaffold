package user

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	//解析命令行参数
	lsCmd.Flags().BoolVarP(&ls.ListShowAll, "all", "l", false, "show all user")
	userCmd.AddCommand(lsCmd)
}

type lsOption struct {
	ListShowAll bool
}

var ls lsOption
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "show users",
	Long:  `this is a long description with ls command`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		logger.Infoln("hello ls")
		logger.Infoln("args:", args)
		logger.Infof("lsOption:%#v", ls)

		fmt.Println("user1")
		fmt.Println("user2")
		fmt.Println("user3")
		if ls.ListShowAll {
			fmt.Println("userAdmin")
		}
	},
}
