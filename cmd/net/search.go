package net

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	query string
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "This command searches google.",
	Long: `This command searches google. It is used to search google for a specific query.
	
	Example:
		net search -q "Kaas in nederland"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Searching for: ", query)
		out, err := exec.Command("python", "scripts/search.py", query).CombinedOutput()
		if err != nil {
			log.Fatal("Command did not run successfully\n", err)
			return
		}
		fmt.Printf("Output: \n%s\n", out)
	},
}

func init() {
	searchCmd.Flags().StringVarP(&query, "query", "q", "Kaas in nederland", "The search query")

	NetCmd.AddCommand(searchCmd)
}
