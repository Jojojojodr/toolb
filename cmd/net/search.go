package net

import (
	"fmt"
    "log"
    "net/http"
    "net/url"
    "strings"

    "github.com/PuerkitoBio/goquery"
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
		fmt.Printf("Searching for: %s\n\n", query)
		searchURL := "https://www.google.com/search?q=" + url.QueryEscape(query)
		req, err := http.NewRequest("GET", searchURL, nil)
		if err != nil {
			log.Fatal("Failed to perform search:", err)
			return
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Failed to perform search:", err)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal("Failed to parse search results:", err)
			return
		}

		doc.Find("h3").Each(func(i int, s *goquery.Selection) {
			title := s.Text()
			link, exists := s.Parent().Attr("href")
			if exists {
				link = strings.TrimPrefix(link, "/url?q=")
				fmt.Printf("Result %d: %s\nLink: %s\n\n", i+1, title, link)
			}
		})
	},
}

func init() {
	searchCmd.Flags().StringVarP(&query, "query", "q", "Kaas in nederland", "The search query")

	NetCmd.AddCommand(searchCmd)
}
