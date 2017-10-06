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

			tokens = append(tokens, extractTokenNames(input, "("+strings.Join(args.keyNames, "|")+")=\".*\"\\s*?("+strings.Join(args.valueNames, "|")+")=\".*\"")...)

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
				for _, argggg := range tokens {
					re := fmt.Sprintf("("+strings.Join(args.keyNames, "|")+")=\"%s\"\\s*("+strings.Join(args.valueNames, "|")+")=", t)
					re3 := regexp.MustCompile(re)
					if re3.Match(argggg) {
						argggg = re3.ReplaceAll(argggg, []byte(""))
						argggg = bytes.Replace(argggg, []byte("\""), []byte(""), -1)
						tfound := false
						for _, exclude := range args.excluded {
							if strings.HasPrefix(string(argggg), exclude) {
								tfound = true
							}
						}
						if !tfound {
							tokenValueMap[string(t)] = string(argggg)
						}
					}
				}
			}

			deps := map[string]mermaidNode{}

			for k, v := range tokenValueMap {
				for _, ss := range Configuration.DependencyStrings {
					v = strings.ToLower(v)
					ss = strings.ToLower(ss)
					if strings.Contains(v, ss) {

						for _, mappedItem := range Configuration.Targets {
							if sliceElementInString(mappedItem.Identifiers, v) {
								k = mappedItem.FriendlyName
							}
						}

						parseProtocols := func(fName, fInput string) mermaidNode {
							r := mermaidNode{
								Name: fName,
							}
							for _, ci := range Configuration.TargetCategories {
								if sliceElementInString(ci.Identifiers, fInput) {
									r.Form = ci.Category
									r.Link = ""
								}
							}
							if r.Form == "" {
								r.Form = "???"
								r.Link = ""
							}
							return r
						}
						nk := parseProtocols(k, v)
						if !stringInSliceElement(v, Configuration.ExcludedIdentifiers) {
							deps[k] = nk
						}
					}
				}
				// if !stringInSlice(app, order) {
				// 	order = append(order, app)
				// }
				buildChartList(app, deps, category)
			}
		}
		for _, v := range allContent {
			writeContent("All Applications", v, []string{}, category)
		}
	}
	// for _, dwa := range order {
	// 	fmt.Println(dwa)
	// }
}

func tokenize(input []byte, buffers []string, tokenMap map[string]string) string {
	inputS := string(input)
	for k, v := range tokenMap {
		if v != "false" && v != "true" {
			inputS = strings.Replace(inputS, "\""+v+"\"", buffers[0]+k+buffers[1], 1)
		}
	}
	return inputS
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

func outputToStdout(data []byte) {
	fmt.Print(string(data))
}
