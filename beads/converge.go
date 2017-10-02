package main

import (
	"encoding/json"
	"io/ioutil"
)

func checkForDucplicates(key, val string, input map[string]map[string]string) int {
	count := 0
	for name := range input {
		if ev, ok := input[name][key]; ok {
			if ev == val {
				count++
			}
		}
	}
	return count
}

func convergeFromPath(path, output string) {
	files, _ := ioutil.ReadDir(path)

	results := map[string]map[string]string{}
	fileMaps := map[string]map[string]string{}
	results["environment"] = map[string]string{}

	for _, f := range files {
		vals := loadFile(path + "/" + f.Name())
		thisFileMap := map[string]string{}
		json.Unmarshal(vals, &thisFileMap)
		fileMaps[f.Name()] = thisFileMap
	}

	for name, fm := range fileMaps {
		thisParsedMap := map[string]string{}
		for k, v := range fm {
			foundCount := checkForDucplicates(k, v, fileMaps)

			if foundCount >= 2 {
				results["environment"][k] = v
			} else {
				thisParsedMap[k] = v
			}
			results[name] = thisParsedMap
		}

		for name, blob := range results {
			smap, _ := json.MarshalIndent(blob, "", "    ")
			outputFilename := output + "/" + name
			if name == "environment" {
				outputFilename += ".json"
			}
			outputToFile(outputFilename, []byte(smap))
		}
	}
}
