package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration is global config storage
var Configuration = configuration{}

func initConfig(confPath string) {
	Configuration.load(confPath)
}

type configuration struct {
	AccessToken         string   `json:"AccessToken"`
	HugoSiteRootPath    string   `json:"HugoSiteRootPath"`
	DependencyStrings   []string `json:"DependencyStrings"`
	FlowIngressPoints   []string `json:"FlowIngressPoints"`
	ExcludedIdentifiers []string `json:"ExcludedIdentifiers"`
	TargetCategories    []struct {
		Category     string   `json:"Category"`
		Identifiers  []string `json:"Identifiers"`
		Icon         string   `json:"Icon"`
		DisplayShape []string `json:"DisplayShape"`
	} `json:"TargetCategories"`
	Targets []struct {
		FriendlyName string   `json:"FriendlyName"`
		Identifiers  []string `json:"Identifiers"`
		SourceURL    string   `json:"SourceURL,omitempty"`
	} `json:"Targets"`
}

func (conf *configuration) load(confPath string) {
	fileContent, err := os.Open(confPath)
	if err != nil {
		fmt.Println("Could not open config file\n", err)
	}

	jsonParser := json.NewDecoder(fileContent)
	if err = jsonParser.Decode(&conf); err != nil {
		fmt.Println("Could not load config file. Check JSON formatting.\n", err)
	}
}
