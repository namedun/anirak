package main

import (
	"io/ioutil"
	"os"
	"regexp"

	"github.com/barjoco/go-utils/log"
)

// Colour theme definition
var colours = map[string]string{
	"black":                 "#1e2227",
	"red":                   "#ff4766",
	"green":                 "#17ff93",
	"yellow":                "#ffdc69",
	"blue":                  "#3743ee",
	"magenta":               "#4a8ffd",
	"cyan":                  "#75abff",
	"white":                 "#f5faff",
	"hi_black":              "#3a434d",
	"hi_red":                "#ff8398",
	"hi_green":              "#71faba",
	"hi_yellow":             "#ffe388",
	"hi_blue":               "#6485f1",
	"hi_magenta":            "#7aadff",
	"hi_cyan":               "#bad3ff",
	"hi_white":              "#ffffff",
	"white0.75":             "#def0ffbf",
	"white0.5":              "#def0ff80",
	"white0.4":              "#def0ff66",
	"white0.25":             "#def0ff40",
	"white0.1":              "#def0ff1a",
	"blue0.75":              "#3743eebf",
	"blue0.5":               "#3743ee80",
	"blue0.25":              "#3743ee40",
	"blue0.1":               "#3743ee1a",
	"magenta0.75":           "#7aadffbf",
	"magenta0.5":            "#7aadff80",
	"magenta0.25":           "#7aadff40",
	"magenta0.1":            "#7aadff1a",
	"red0.75":               "#ff4766bf",
	"red0.5":                "#ff476680",
	"green0.5":              "#17ff9380",
	"active_line_num":       "#adbac9",
	"inactive_line_num":     "#373e45",
	"line_highlight":        "#111217",
	"line_highlight_border": "#17181f",
	"background":            "#0c0d10",
}

func main() {
	argv := os.Args
	argc := len(argv)

	//
	//	Validate args
	//

	if argc != 3 {
		log.Error("Must have 2 arguments (recieved %d)", argc-1)
	}

	//
	//	Parse colour scheme
	//

	fileIn := argv[1]
	fileOut := argv[2]

	// Read file
	b, err := ioutil.ReadFile(fileIn)
	log.ReportFatal(err)

	// Replace --COLOUR-- with colours[COLOUR]
	re := regexp.MustCompile(`--[^--]+--`)
	fileParsed := re.ReplaceAllStringFunc(string(b), func(s string) string {
		return colours[s[2:len(s)-2]]
	})

	// Write to fileOut
	err = ioutil.WriteFile(fileOut, []byte(fileParsed), 0644)
	log.ReportFatal(err)
	log.Success("Theme has been parsed!")
}
