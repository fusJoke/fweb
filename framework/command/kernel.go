package command

import "github.com/fusjoke/fweb/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	// 挂载AppCommand命令
	root.AddCommand(initAppCommand())
}
