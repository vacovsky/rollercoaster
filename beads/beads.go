package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"bytes"

	"strings"

	"regexp"
)

func main() {
	args := flagInit()
	initConfig(*args.config)

	createContentFolders()

	transforms := scrapeAppTransforms()
	for _, cm := range Configuration.CategoryMeta {
		category := cm.Name
		for app, content := range transforms {
			tokens := [][]byte{}
			input := []byte(content)
			tokens = append(tokens, extractTokenNames(input, "("+strings.Join(
				args.keyNames, "|")+")=\".*\"\\s*?("+strings.Join(
				args.valueNames,
				"|")+")=\".*\"")...)
			cleanTokens := [][]byte{}
			for _, t := range tokens {
				r1 := regexp.MustCompile("(" + strings.Join(args.keyNames, "|") + ")=\"")
				r2 := regexp.MustCompile("\"\\s*?(" + strings.Join(args.valueNames, "|") + ")=\".*\"")

				t = r1.ReplaceAll(t, []byte(""))
				t = r2.ReplaceAll(t, []byte(""))

				cleanTokens = append(cleanTokens, t)
			}
			tokenValueMap := map[string]string{}
			for _, t := range cleanTokens {
				for _, token := range tokens {
					re := fmt.Sprintf("("+strings.Join(args.keyNames, "|")+")=\"%s\"\\s*("+strings.Join(args.valueNames, "|")+")=", t)
					re3 := regexp.MustCompile(re)
					if re3.Match(token) {
						token = re3.ReplaceAll(token, []byte(""))
						token = bytes.Replace(token, []byte("\""), []byte(""), -1)
						tfound := false
						for _, exclude := range args.excluded {
							if strings.HasPrefix(string(token), exclude) {
								tfound = true
							}
						}
						if !tfound {
							tokenValueMap[string(t)] = string(token)
						}
					}
				}
			}
			deps := buildMermaidNodes(tokenValueMap)
			buildChartList(app, deps, category)
		}
		for _, v := range allContent {
			writeContent("All Applications", v, []string{}, category)
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func sliceElementInString(slc []string, inp string) bool {
	for _, elem := range slc {
		if strings.Contains(inp, elem) {
			return true
		}
	}
	return false
}
func loadFile(path string) []byte {
	fmt.Println("Loading: ", path)
	file, err := ioutil.ReadFile(path)
	checkError(err)
	return file
}

func outputToFile(path string, data []byte) {
	err := ioutil.WriteFile(path, data, 0644)
	checkError(err)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func stringInSliceElement(a string, list []string) bool {
	for _, b := range list {
		if strings.Contains(b, a) {
			return true
		}
	}
	return false
}
func ensureFileExists(file, use string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Printf("Error: File \"%s\" does not exist. Please provide a valid file path for %s.\n", file, use)
		os.Exit(1)
	}
}
