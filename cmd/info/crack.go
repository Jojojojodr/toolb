package info

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	query string
)

var crackCmd = &cobra.Command{
	Use:   "crack",
	Short: "This command cracks a hash.",
	Long: `This command cracks a hash. It is used to crack a hash using a bruteforce attack.
	
	Example:
		info crack -q "sha1 5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8"
		info crack -q "md5 5f4dcc3b5aa765d61d8327deb882cf99"
		info crack -q "sha256 5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Cracking hash: %s\n\n", query)
		out, err := exec.Command("python", "scripts/bruteforce.py", query).CombinedOutput()
		if err != nil {
			log.Fatal("Command did not run successfully\n", err)
			return
		}
		fmt.Printf("Output: \n%s\n", out)
	},
}

func init() {
	crackCmd.Flags().StringVarP(&query, "query", "q", "sha1 5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8", "Args for the crack command")

	InfoCmd.AddCommand(crackCmd)
}
