package cmd

import (
	"fmt"
	"os"

	"github.com/Jojojojodr/toolb/cmd/info"
	"github.com/Jojojojodr/toolb/cmd/net"

	"github.com/spf13/cobra"
)

var (
	showVersion bool
)

var rootCmd = &cobra.Command{
	Use:   "toolb",
	Short: "Toolb is a collection of tools.",
	Long: `Toolb is a collection of tools. It is a collection of tools that can be used to troubleshoot issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Println("Version:", Version)
		} else {
			fmt.Printf("Welcome to the toolb!\n\n")
			cmd.Help()
		}
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
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Print the version of the toolb.")

	addSubCommandsPallettes()
}


