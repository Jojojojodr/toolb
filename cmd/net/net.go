package net

import (
	"github.com/spf13/cobra"
)

var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "Net is a pallet of network tools.",
	Long: `Net is a pallet of network tools. It is a collection of tools that can be used to troubleshoot network issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
