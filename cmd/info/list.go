package info

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list files in the current directory",
	Long: `This command lists all files in the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("ls", "-l", ".").Output()
		if err != nil {
			log.Fatal("Command did not run successfully", err)
			return
		}
		fmt.Printf("output: \n%s\n", out)
	},
}

func init() {
	InfoCmd.AddCommand(lsCmd)
}
