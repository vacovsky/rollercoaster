package main

import (
	"flag"
	"strings"
)

type flagsModel struct {
	keyNames   []string
	valueNames []string
	excluded   []string
	config     *string
}

func flagInit() flagsModel {
	model := flagsModel{}

	model.config = flag.String("config", "", "configuration file path")
	// keys and values prefixes.  In the input file, this will be something like <add name="blah" value="someblah">
	keyNamesString := flag.String("key-names", "", "Strings to search for when detecting token keys.  Separate strings with a comma (,) and no spaces.")
	valueNamesString := flag.String("value-names", "", "Strings to search for when detecting token values.  Separate strings with a comma (,) and no spaces.")
	excludedString := flag.String("exclude", "", "Values you wish to ignore if they start with anything in this list.  Separate strings with a comma (,) and no spaces.")

	flag.Parse()

	model.excluded = strings.Split(*excludedString, ",")
	model.keyNames = strings.Split(*keyNamesString, ",")
	model.valueNames = strings.Split(*valueNamesString, ",")
	return model
}
