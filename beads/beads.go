package main

import (
	"fmt"
	"strconv"

	"bytes"

	"strings"

	"regexp"
)

// var wg sync.WaitGroup

func main() {

	args := flagInit()
	// args := flagsModel{}
	initConfig(*args.config)
	createContentFolders()
	buildMap(args, scrapeAppTransforms())
}

func buildMap(blargs flagsModel, transforms map[string]string) {
	done := []string{}
	for _, cm := range Configuration.CategoryMeta {

		category := cm.Name
		for app, content := range transforms {
			tokens := [][]byte{}
			input := []byte(content)
			tokens = append(tokens, extractTokenNames(input, "("+strings.Join(
				blargs.keyNames, "|")+")=\".*\"\\s*?("+strings.Join(
				blargs.valueNames,
				"|")+")=\".*\"")...)
			cleanTokens := [][]byte{}
			for _, t := range tokens {
				r1 := regexp.MustCompile("(" + strings.Join(blargs.keyNames, "|") + ")=\"")
				r2 := regexp.MustCompile("\"\\s*?(" + strings.Join(blargs.valueNames, "|") + ")=\".*\"")

				t = r1.ReplaceAll(t, []byte(""))
				t = r2.ReplaceAll(t, []byte(""))

				cleanTokens = append(cleanTokens, t)
			}
			tokenValueMap := map[string]string{}
			for _, t := range cleanTokens {
				for _, token := range tokens {
					re := fmt.Sprintf("("+strings.Join(blargs.keyNames, "|")+")=\"%s\"\\s*("+strings.Join(blargs.valueNames, "|")+")=", t)
					re3 := regexp.MustCompile(re)
					if re3.Match(token) {
						token = re3.ReplaceAll(token, []byte(""))
						token = bytes.Replace(token, []byte("\""), []byte(""), -1)
						tfound := false
						for _, exclude := range blargs.excluded {
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

			// fmt.Println(category + ": " + strconv.Itoa(len(tokenValueMap)))
			buildChartList(app, deps, category)
		}

		for ll, v := range allContent {
			if !stringInSlice(ll, done) {
				writeContent("All Applications", v, []string{}, category)
				done = append(done, ll)
				fmt.Println(ll + ": " + strconv.Itoa(len(v)))
			}
		}
	}
}
