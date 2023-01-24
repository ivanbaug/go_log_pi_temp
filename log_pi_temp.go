package main

import (
	"bufio"
	"encoding/json"
        "flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type MConfigFile struct {
	ApiKey string
}

var configFile = flag.String("config", "config.json", "Location of the config file.")

func main() {
	flag.Parse()
	configData, err := os.ReadFile(*configFile)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		os.Exit(1)
	}

	var configs MConfigFile
	err = json.Unmarshal(configData, &configs)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		os.Exit(1)
	}

	//temperatureFileName := "temp"
	temperatureFileName := "/sys/class/thermal/thermal_zone0/temp"

	temperatureFile, err := os.Open(temperatureFileName)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		os.Exit(1)
	}
	rFile := bufio.NewReader(temperatureFile)
	tempStr, _, _ := rFile.ReadLine()
	cpuTemp, _ := strconv.ParseFloat(string(tempStr), 8)
	cpuTemp = cpuTemp / 1000

	// GET https://api.thingspeak.com/update?api_key=apikey&field1=0

	requestURL := fmt.Sprintf("https://api.thingspeak.com/update?api_key=%s&field1=%.2f", configs.ApiKey, cpuTemp)

	_, err = http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
}
