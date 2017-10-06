package main

import (
	"os"
	"time"
)

func writeContent(application string, input []string, tags []string, category string) {
	tagsStr := "["

	for _, v := range tags {
		tagsStr += "\"" + v + "\",\n"

	}
	tagsStr += `]`
	template := `---
draft: false
date: "` + time.Now().Format("2006-01-02T15:04:05") + `"
tags: ` + tagsStr + `
title: ` + application + `
categories: ["` + category + `"]
depmap: [ "graph LR",
`
	for _, v := range input {
		template += "\"" + v + "\",\n"
	}
	template += `]
---
`

	outputToFile(Configuration.HugoSiteRootPath+"/content/"+category+"/"+application+".md", []byte(template))
}

func createContentFolders() {
	for _, d := range Configuration.CategoryMeta {
		_ = os.Mkdir(Configuration.HugoSiteRootPath+"/content/"+d.Name, os.ModePerm)
	}
}
