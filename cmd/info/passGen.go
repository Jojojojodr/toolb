package info

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitBytes  = "0123456789"
	symbolBytes = "!@#$%^&*()-_=+,.?/:;{}[]~"
)

var (
	length int
	digits bool
	symbols bool
	charset string
)

var passGenCmd = &cobra.Command{
	Use:   "password-gen",
	Short: "This command generates a random password.",
	Long: `This command generates a random password. It is used to generate a random password with the specified length and optional digits and symbols.`,
	Run: func(cmd *cobra.Command, args []string) {
		charset = letterBytes
		if digits {
			charset += digitBytes
		}
		if symbols {
			charset += symbolBytes
		}

		password := make([]byte, length)
		for i := range password {
			password[i] = charset[rand.Intn(len(charset))]
		}

		fmt.Printf("Generated password:\n%s\n", password)
	},
}

func init() {
	passGenCmd.Flags().IntVarP(&length, "length", "l", 8, "The length of the password")
	passGenCmd.Flags().BoolVarP(&digits, "digits", "d", false, "Include digits in the password")
	passGenCmd.Flags().BoolVarP(&symbols, "symbols", "s", false, "Include symbols in the password")

	InfoCmd.AddCommand(passGenCmd)
}
