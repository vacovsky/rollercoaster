package main

import (
	"encoding/json"
	"regexp"

	"bytes"
)

func extractTokenNames(input []byte, regex string) [][]byte {
	namesRegex := regex
	rex := regexp.MustCompile(namesRegex)
	return rex.FindAll(input, -1)
}

func extractTokens(input []byte, buffer []string) [][]byte {
	bufferBuilder := buffer[0] + "(.*?)" + buffer[1]
	rex := regexp.MustCompile(bufferBuilder)
	return rex.FindAll(input, -1)
}

func convertToJSON(data [][]byte, buffer []string) []byte {
	dmap := map[string]string{}
	for _, d := range data {
		cleanD := bytes.Replace(d, []byte(buffer[0]), []byte{}, -1)
		cleanD = bytes.Replace(cleanD, []byte(buffer[1]), []byte{}, -1)
		dmap[string(cleanD)] = ""
	}
	jdata, err := json.MarshalIndent(dmap, "", "    ")
	checkError(err)
	return jdata
}

func mapKeyPairs(input [][]byte, buffer []string) []map[string][]byte {
	tokenMapS := []map[string][]byte{}
	for _, pv := range input {
		tokenMap := map[string]*json.RawMessage{}

		json.Unmarshal(pv, &tokenMap)

		tempMap := map[string][]byte{}

		for k, v := range tokenMap {
			j, err := json.Marshal(&v)
			checkError(err)
			tempMap[buffer[0]+k+buffer[1]] = j
		}
		tokenMapS = append(tokenMapS, tempMap)
	}
	return tokenMapS
}
