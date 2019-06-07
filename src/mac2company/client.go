package main

import (

	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"io/ioutil"
	"strings"
)

const version = "1.0"

var body struct {
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
}

var mainCmd = &cobra.Command{
	Use:   "client",
	Short: "MAC lookup api client",
	Long:  "Simple client to interact with macaddress.io.",
	Run: func(cmd *cobra.Command, args []string) {
		runCmd()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version.",
	Long:  "The version of the dispatch service.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func runCmd() {
    fmt.Println("MAC address company lookup client. Type mac2company -h for help.")
}


var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Lookup a mac uuid.",
	Long:  "Lookup a mac uuid and return the company name.",
	Run: func(cmd *cobra.Command, args []string) {

		url := viper.GetString("config.url")
		apikey := viper.GetString("config.apikey")
		mac := viper.GetString("mac")
		output := viper.GetString("config.output")

		apiTarget := fmt.Sprintf("%s?apiKey=%s&output=%s&search=%s", url, apikey,output,mac)
		res, err := http.Get(apiTarget)
		if err != nil{
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()

		if err != nil {
			log.Fatal(err)
		}


        if strings.Contains(string(body), "Invalid") {
			log.Fatal("Invalid MAC address (API result)")
		}

		if string(body) != "" {
		fmt.Printf("The company associated with MAC: %s: is %s\n", mac, body)
		} 
	},
}


func init() {

	viper.SetConfigName("config") // no need to include file extension
	viper.AddConfigPath(".") // look in current directory
	viper.AddConfigPath("/etc/mac2company/")
	viper.AddConfigPath("$HOME/.mac2company/") // look in home directory

	
	err := viper.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		log.Fatal(err)
	}

	// Adding commands into the client

	mainCmd.AddCommand(versionCmd)
	mainCmd.AddCommand(lookupCmd)
	
	lookupFlags := lookupCmd.Flags()

	lookupFlags.String("mac", "", "mac of device you want to lookup.")
	viper.BindPFlag("mac", lookupFlags.Lookup("mac"))

}

func main() {
	mainCmd.Execute()

}



