package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(traceCmd)
}

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP address",
	Long:  `Trace the IP address.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("Please provide an IP address")
		}
	},
}

type Ip struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Timezone string `json:"timezone"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
}

func showData(ip string) {
	if net.ParseIP(ip) == nil {
		c := color.New(color.FgRed)
		c.Print("Error: ")
		fmt.Println("Please provide a valid IP address")
		return
	}

	url := "https://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := Ip{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		fmt.Println("Unable to unmarshal the response.")
	}
	c := color.New(color.FgMagenta)
	c.Println("Data Fetched Successfully 🎉 🎉")
	fmt.Printf("IP: %s\nCity: %s\nRegion: %s\nCountry: %s\nLocation: %s\nTimezone: %s\nOrganization: %s\nPostal: %s\n", data.IP, data.City, data.Region, data.Country, data.Location, data.Timezone, data.Org, data.Postal)
}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Unable to get the response from the server")
	}

	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Unable to read the response.")
	}

	return responseByte
}
