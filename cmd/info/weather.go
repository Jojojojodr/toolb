package info

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	API_KEY string = "cc0582fc84e09298349ff4b21dc76c5e"
)

var (
	lat string
	lon string
	hour bool
	day bool
)

type IPInfo struct {
    IP       string  `json:"ip"`
    City     string  `json:"city"`
    Region   string  `json:"region"`
    Country  string  `json:"country"`
    Loc      string  `json:"loc"`
    Org      string  `json:"org"`
    Postal   string  `json:"postal"`
    Timezone string  `json:"timezone"`
    Readme   string  `json:"readme"`
}

type Weather struct {
	Timezone string `json:"timezone"`
	TimezoneOffset int `json:"timezone_offset"`
	Current struct {
		Dt int64 `json:"dt"`
		Temp float64 `json:"temp"`
		Weather []struct {
			Main string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"current"`
	Hourly []struct {
		Temp float64 `json:"temp"`
		Weather []struct {
			Main string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"hourly"`
	Daily []struct {
		Summary string `json:"summary"`
		Temp struct {
			Day float64 `json:"day"`
			Morning float64 `json:"morn"`
			Evening float64 `json:"eve"`
			Night float64 `json:"night"`
		} `json:"temp"`
		Weather []struct {
			Main string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"daily"`
}

func getLocation() (lat, lon string, err error) {
    resp, err := http.Get("https://ipinfo.io/json")
    if err != nil {
        return "", "", fmt.Errorf("failed to get location: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", "", fmt.Errorf("failed to read response body: %v", err)
    }

    var ipInfo IPInfo
    err = json.Unmarshal(body, &ipInfo)
    if err != nil {
        return "", "", fmt.Errorf("failed to unmarshal JSON: %v", err)
    }

    loc := ipInfo.Loc
    coords := strings.Split(loc, ",")
    if len(coords) != 2 {
        return "", "", fmt.Errorf("invalid location format: %v", loc)
    }

    lat = coords[0]
    lon = coords[1]
    return lat, lon, nil
}

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://api.openweathermap.org/data/3.0/onecall?lat=" + lat + "&lon=" + lon + "&units=metric" + "&appid=" + API_KEY


		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			panic("Weather API call failed")
		} 
		
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		
		var weather Weather
		err = json.Unmarshal(body, &weather)
		if err != nil {
			panic(err)
		}

		timezone := weather.Timezone
		current := weather.Current
		hourly := weather.Hourly
		daily := weather.Daily
		currtime := current.Dt

		unixTime := time.Unix(currtime, 0)

		fmt.Printf("Weather\n\n")
		fmt.Printf("Current: \n%s\n%s, %s, Temp: %.0fC, %s\n\n", unixTime, timezone, current.Weather[0].Main, current.Temp, current.Weather[0].Description)

		if hour {
			for _, h := range hourly {
				fmt.Printf("Hourly: \n%.0fC, %s\n", h.Temp, h.Weather[0].Description)
			}
			fmt.Printf("\n")
		}

		if day {
			for _, d := range daily {
				fmt.Printf("Dialy: %s,\nDay temp %.0fC, Morning temp %.0fC, Evening temp %.0fC, Night temp %.0fC,\n%s\n\n", d.Summary, d.Temp.Day, d.Temp.Morning, d.Temp.Evening, d.Temp.Night, d.Weather[0].Description)
			}
		}
	},
}

func init() {
	latitude, longitude, err := getLocation()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

	lat = latitude
	lon = longitude

	weatherCmd.Flags().BoolVarP(&hour, "hourly", "o", false, "Show hourly weather")
	weatherCmd.Flags().BoolVarP(&day, "daily", "d", false, "Show daily weather")

	InfoCmd.AddCommand(weatherCmd)
}
