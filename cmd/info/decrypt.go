package info

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var (
	Key string
	encryptedFile string
	decryptedFile string
)

func checkKey() ([]byte, error) {
	key, err := hex.DecodeString(Key)
	if err != nil {
		fmt.Println("Error decoding key:", err)
		return nil, err
	}

	if len(key) != 32 {
		fmt.Println("Invalid key size: must be 32 bytes")
		return nil, errors.New("invalid key size")
	}

	return key, nil
}

func decryptFile(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Error creating GCM:", err)
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println("Ciphertext too short")
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println("Error decrypting file:", err)
		return nil, err
	}

	return plaintext, nil
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ciphertext, err := ioutil.ReadFile(encryptedFile)
		if err != nil {
			fmt.Println("Error reading encrypted file:", err)
			return
		}

		key, err := checkKey()
		if err != nil {
			fmt.Println("Error checking key:", err)
			return
		}

		plaintext, err := decryptFile(key, ciphertext)
		if err != nil {
			fmt.Println("Error decrypting file:", err)
			return
		}

		err = ioutil.WriteFile(decryptedFile, plaintext, 0644)
		if err != nil {
			fmt.Println("Error writing decrypted file:", err)
			return
		}

		fmt.Println("File decrypted successfully")
	},
}

func init() {
	decryptCmd.Flags().StringVarP(&Key, "key", "k", "", "Key to decrypt the file")
	decryptCmd.Flags().StringVarP(&encryptedFile, "input", "i", "example.enc", "Encrypted file to decrypt")
	decryptCmd.Flags().StringVarP(&encryptedFile, "output", "o", "example_decrypted.txt", "Output file to write the decrypted content")
	decryptCmd.MarkFlagRequired("key")
	InfoCmd.AddCommand(decryptCmd)
}
