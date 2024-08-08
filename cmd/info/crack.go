package info

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var (
	hashMethod string
	hash	   string
)

var crackCmd = &cobra.Command{
	Use:   "crack",
	Short: "This command cracks a hash.",
	Long: `This command cracks a hash. It is used to crack a hash using a bruteforce attack.
	
	Example:
        info crack -m sha1 -h 5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8
        info crack -m md5 -h 5f4dcc3b5aa765d61d8327deb882cf99
        info crack -m sha256 -h 5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Cracking hash: %s, %s\n\n", hash, hashMethod)

		passwords, err := ioutil.ReadFile("passwords.txt")
		if err != nil {
			log.Fatal("Could not read passwords file.")
			return
		}

		passwordList := strings.Split(string(passwords), "\n")

		for _, password := range passwordList {
			var hashedPassword string
			fmt.Printf("Trying password: %s\n", password)
			switch hashMethod {
			case "md5":
				hash := md5.Sum([]byte(password))
				hashedPassword = hex.EncodeToString(hash[:])
			case "sha1":
				hash := sha1.Sum([]byte(password))
				hashedPassword = hex.EncodeToString(hash[:])
			case "sha256":
				hash := sha256.Sum256([]byte(password))
				hashedPassword = hex.EncodeToString(hash[:])
			default:
				log.Fatal("Invalid hash method.")
				return
			}

			if hashedPassword == hash {
				fmt.Printf("Password found: %s\n", password)
				return
			}
		}

		fmt.Println("Password not found.")
	},
}

func init() {
	crackCmd.Flags().StringVarP(&hash, "hash", "x", "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8", "Args for the crack command")
	crackCmd.Flags().StringVarP(&hashMethod, "method", "m", "sha1", "Args for the crack command")

	InfoCmd.AddCommand(crackCmd)
}
