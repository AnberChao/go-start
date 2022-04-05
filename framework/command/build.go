package command

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/JoeZhao1/go-start/framework/cobra"
)

// build相关的命令
func initBuildCommand() *cobra.Command {
	buildCommand.AddCommand(buildSelfCommand)
	buildCommand.AddCommand(buildBackendCommand)
	return buildCommand
}

var buildCommand = &cobra.Command{
	Use:   "build",
	Short: "编译相关命令",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			c.Help()
		}
		return nil
	},
}

var buildSelfCommand = &cobra.Command{
	Use:   "self",
	Short: "编译go-start命令",
	RunE: func(c *cobra.Command, args []string) error {
		path, err := exec.LookPath("go")
		if err != nil {
			log.Fatalln("start go: 请在Path路径中先安装go")
		}

		cmd := exec.Command(path, "build", "-o", "start", "./")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("go build error:")
			fmt.Println(string(out))
			fmt.Println("--------------")
			return err
		}
		fmt.Println("编译go-start成功")
		return nil
	},
}

var buildBackendCommand = &cobra.Command{
	Use:   "backend",
	Short: "编译程序",
	RunE: func(c *cobra.Command, args []string) error {
		return buildSelfCommand.RunE(c, args)
	},
}