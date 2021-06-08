package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/barjoio/utils/log"
)

// Colour theme definition
var colours = map[string]string{
	"black":              "#1e2227",
	"red":                "#ff4766",
	"green":              "#17ff93",
	"yellow":             "#ffdc69",
	"blue":               "#4e6bff",
	"magenta":            "#66a2ff",
	"cyan":               "#abd5ff",
	"white":              "#f0f8ff",
	"hi_black":           "#3a434d",
	"hi_red":             "#ff8398",
	"hi_green":           "#71faba",
	"hi_yellow":          "#ffe89e",
	"hi_blue":            "#6e86ff",
	"hi_magenta":         "#80b2ff",
	"hi_cyan":            "#c9e4ff",
	"hi_white":           "#ffffff",
	"real_black":         "#000000",
	"background0":        "#0c0d10",
	"background1":        "#111217",
	"background1_border": "#191b21",
	"background2":        "#21232b",
	"background2_border": "#30333d",
	"transparent":        "#00000000",
	"red80":              "#ff476680",
	"red40":              "#ff476640",
	"green80":            "#17ff9380",
	"green40":            "#17ff9340",
	"yellow80":           "#ffdc6980",
	"blue90":             "#303eff90",
	"blue80":             "#303eff80",
	"blue60":             "#303eff60",
	"blue40":             "#303eff40",
	"blue20":             "#303eff20",
	"blue10":             "#303eff10",
	"magenta40":          "#599aff40",
	"cyan80":             "#9ccdff80",
	"cyan40":             "#9ccdff40",
	"cyan20":             "#9ccdff20",
	"whiteC0":            "#f0f8ffC0",
	"white80":            "#f0f8ff80",
	"white40":            "#f0f8ff40",
	"white30":            "#f0f8ff30",
	"white20":            "#f0f8ff20",
	"white10":            "#f0f8ff10",
	"bluewhite":          "#deebff",
	"bluewhiteD0":        "#deebffD0",
	"bluewhiteC0":        "#deebffC0",
	"bluewhiteB0":        "#deebffB0",
	"bluewhiteA0":        "#deebffA0",
	"bluewhite90":        "#deebff90",
	"bluewhite80":        "#deebff80",
	"bluewhite60":        "#deebff60",
	"bluewhite50":        "#deebff50",
	"bluewhite40":        "#deebff40",
	"bluewhite30":        "#deebff30",
	"bluewhite20":        "#deebff20",
	"bluewhite10":        "#deebff10",
	"bluewhite08":        "#deebff08",
}

func main() {
	argv := os.Args[1:]
	argc := len(argv)

	for _, arg := range argv {
		if arg[:2] == "--" {
			re := regexp.MustCompile("--([^=]+)=?([^:]+)?:?(.+)?")
			submatch := re.FindStringSubmatch(arg)
			conversion := submatch[1]
			fileIn := submatch[2]
			fileOut := submatch[3]

			switch conversion {
			case "help":
				if argc > 1 {
					log.Info("Run --help on its own: akparse --help")
					os.Exit(1)
				}
				fmt.Println("Usage:\n\takparse --type=filein ...\n\takparse --type=filein:fileout ...\n\nDescription:\n\tUsed to parse any number of files with the anirak theme. If fileout\n\tisn't specified, the output file name will be 'filein.anirak'.\n\nTypes:\n\t--hex\t outputs: #ff0000\n\t--hex0\t outputs: 0xff0000\n\t--rgb\t outputs: rgb(255, 0, 0)\n\nExample:\n\takparse --hex=template.json:theme.json")
			case "hex":
				parse(fileIn, fileOut, 0)
			case "hex0":
				parse(fileIn, fileOut, 1)
			case "rgb":
				parse(fileIn, fileOut, 2)
			default:
				log.Error("Unknown argument: %s", conversion)
			}
		} else {
			log.ErrorFatal("Arguments must be formatted: --type=filein:fileout. Run --help for more info")
		}
	}
}

func parse(fileIn, fileOut string, conversion int) {
	file, err := ioutil.ReadFile(fileIn)
	log.ReportFatal(err)

	re := regexp.MustCompile("--[^--]+--")
	fileParsed := re.ReplaceAllStringFunc(string(file), func(s string) string {
		colourName := s[2 : len(s)-2]
		var value string
		switch conversion {
		case 0: // hex
			value = colours[colourName]
		case 1: // hex0
			value = "0x" + colours[colourName][1:]
		case 2: // rgb
			rgb, err := strconv.ParseUint(colours[colourName][1:], 16, 32)
			log.ReportFatal(err)
			value = fmt.Sprintf("rgb(%d, %d, %d)", uint8(rgb>>16), uint8((rgb>>8)&0xFF), uint8(rgb&0xFF))
		}
		return value
	})

	if fileOut == "" {
		fileOut = fileIn + ".anirak"
	}

	err = ioutil.WriteFile(fileOut, []byte(fileParsed), 0644)
	log.ReportFatal(err)
}
