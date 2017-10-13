package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func bubbleSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
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
