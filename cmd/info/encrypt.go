package info

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
)

var inputFile string

func generateKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
}

func writeOutputFile(gcm cipher.AEAD, nonce []byte, content []byte) error {
	ciphertext := gcm.Seal(nonce, nonce, content, nil)

	outputFile := strings.Split(inputFile, ".")[0] + ".enc"
		err := ioutil.WriteFile(outputFile, ciphertext, 0644)
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Encrypted file written to", outputFile)
		return nil
}

func encryptFile(key []byte) (cipher.AEAD, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return nil, nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Error creating GCM:", err)
		return nil, nil, err
	}
	
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("Error creating nonce:", err)
		return nil, nil, err
	}

	return gcm, nonce, nil
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
    	content, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		key, err := generateKey()
		if err != nil {
			fmt.Println("Error generating key:", err)
			return
		}

		fmt.Printf("Save this key to decrypt in the future\n\nKey: %x\n", key)
		
		gcm, nonce, err := encryptFile(key)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = writeOutputFile(gcm, nonce, content)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	encryptCmd.Flags().StringVarP(&inputFile, "input", "i", "example.txt", "Input file to encrypt")
	InfoCmd.AddCommand(encryptCmd)
}
