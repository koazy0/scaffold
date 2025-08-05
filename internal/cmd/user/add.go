package user

import (
	"fmt"
	"scaffold/internal/model"
	"scaffold/internal/service"
)

func init() {
	//解析命令行参数
	addCmd.Flags().BoolVarP(&add.Admin, "admin", "a", false, "add admin user")
	addCmd.Flags().StringVarP(&add.Username, "name", "n", "", "set user name")
	addCmd.Flags().StringVarP(&add.Username, "password", "p", "", "set user password") //暂时先这么写，最好通过命令框输入
	userCmd.AddCommand(addCmd)
}

type addOption struct {
	Admin    bool //创建admin角色
	Username string
	Password string
}

var add addOption

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add users",
	Long:  `this is a long description with add command`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

		if add.Username == "" {
			fmt.Println("please input username")
			_, err := fmt.Scanln(&add.Username)
			if err != nil {
				logger.Fatalln(err.Error())
			}
		}
		if add.Password != "" {
			fmt.Println("Warning: maybe should not show password in shell")
		} else {
			_, err := fmt.Scanln(&add.Username)
			if err != nil {
				logger.Fatalln(err.Error())
			}
		}

		var role int

		if add.Admin {
			role = 1
		} else {
			role = 2
		}

		_, err := service.User().CreateUser(cmd.Context(), model.UserSignUp{
			UserID:   add.Username,
			Username: add.Username,
			Password: add.Password,
		}, role)
		if err != nil {
			logger.Fatalf("create user failed:%s", err.Error())
		}
		logger.Infof("create User:%s Success", add.Username)

	},
}
