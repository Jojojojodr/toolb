package cmd

import (
	"fmt"
	"os"
	"github.com/Jojojojodr/toolb/cmd/info"
	"github.com/Jojojojodr/toolb/cmd/net"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "toolbox",
	Short: "Toolbox is a collection of tools.",
	Long: `Toolbox is a collection of tools. It is a collection of tools that can be used to troubleshoot issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the toolbox!")
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


func addSubCommandsPallettes() {
	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(info.InfoCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubCommandsPallettes()
}


