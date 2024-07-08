package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	link string
	client = &http.Client {
		Timeout: time.Second * 2,
	}
)

func pingURL(link string) (int, error) {
	link = "http://" + link
	req, err := http.NewRequest("HEAD", link, nil)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return 0, err
	}

	resp.Body.Close()
	return resp.StatusCode, nil
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This command pings a URL.",
	Long: `This command pings a URL. It is used to check if a URL is reachable.`,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := pingURL(link); err != nil {
			fmt.Println(err)
		} else {
			switch resp {
				case 200:
					fmt.Println("Ping to", link, "was successful.", resp)
				case 404:
					fmt.Println("Ping to", link, "was successful but the page was not found.", resp)
				case 401:
					fmt.Println("Ping to", link, "was successful but the page is unauthorized.", resp)
				case 403:
					fmt.Println("Ping to", link, "was successful but the page is forbidden.", resp)
				case 500:
					fmt.Println("Ping to", link, "was successful but the page has an internal server error.", resp)
				case 503:
					fmt.Println("Ping to", link, "was successful but the page is unavailable.", resp)
			}
		}
	},
}

func init() {
	pingCmd.Flags().StringVarP(&link, "url", "u", "", "URL to ping")

	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}

	NetCmd.AddCommand(pingCmd)
	
}
