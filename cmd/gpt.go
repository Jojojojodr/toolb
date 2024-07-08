package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var(
	KEY string
	msg string
	messages []ChatMessage
)

type ChatMessage struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

func updateMessages(userMsg, aiMsg string) {
    messages = append(messages, ChatMessage{Role: "user", Content: userMsg})
    messages = append(messages, ChatMessage{Role: "ai", Content: aiMsg})

    if len(messages) > 15 {
        messages = messages[len(messages)-15:]
    }
}

func saveMessagesToFile() error {
	var existingMessages []ChatMessage
    fileData, err := os.ReadFile("messages.json")
    if err == nil {
        _ = json.Unmarshal(fileData, &existingMessages)
    }

	combinedMessages := append(existingMessages, messages...)
	if len(combinedMessages) > 15 {
        combinedMessages = combinedMessages[len(combinedMessages)-15:]
    }

	combinedData, err := json.MarshalIndent(combinedMessages, "", "  ")
    if err != nil {
        return err
    }
	return os.WriteFile("messages.json", combinedData, 0644)
}

var gptCmd = &cobra.Command{
	Use:   "gpt",
	Short: "This command sends a message to the GPT-4o model.",
	Long: `This command sends a message to the GPT-4o model. It is used to send a message to the GPT-4o model and get a response.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := openai.NewClient(KEY)
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT4o,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: msg,
					},
				},
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}

		aiResponse := resp.Choices[0].Message.Content
		fmt.Printf("%s\n\n",aiResponse)

		updateMessages(msg, aiResponse)
		if err := saveMessagesToFile(); err != nil {
			fmt.Printf("Error saving messages to file: %v\n", err)
			return
		}

		fmt.Println("Messages saved to messages.json")
	},
}

func init() {
	env := godotenv.Load()
	if env != nil {
		fmt.Println("Error loading .env file")
	}
	
	KEY = os.Getenv("OPENAI_API_KEY")

	gptCmd.Flags().StringVarP(&msg, "msg", "m", "Hello", "The message to send to the GPT-4o model")
	
	rootCmd.AddCommand(gptCmd)
}
