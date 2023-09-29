package app

import (
	"ascii-art/internals/check"
	"fmt"
	"os"
	"strings"
)

func Run() {
	if len(os.Args) > 3 || len(os.Args) <= 1 {
		return
	}

	input := os.Args[1]

	banner := "./banners/standard.txt"
	if len(os.Args) > 2 {
		banner = os.Args[2]
	}
	if !check.Ascii(input) {
		return
	}

	elemsMap := make(map[rune][]string)
	if len(os.Args) > 2 && os.Args[2] == banner {
		switch banner {
		case "standard":
			banner = "./banners/standard.txt"
		case "shadow":
			banner = "./banners/shadow.txt"
		case "thinkertoy":
			banner = "./banners/thinkertoy.txt"
		default:
			fmt.Println("Type the correct banner name")
			return
		}
	}
	data, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("Banner has been deleted")
		return
	}
	changed := strings.ReplaceAll(string(data), "\r\n", "\n")
	sliceData := strings.Split(changed, "\n")              // splits standart.txt into multiple substrings by enters
	input = strings.ReplaceAll(input, "\\n", string('\n')) // replace occurrences of the "\\n" with the newline character '\n'
	splittedArr := strings.Split(input, string('\n'))
	termWidth := check.GetTerminalWidth()

	for i := ' '; i <= '~'; i++ {
		var strs []string
		for j := 0; j < 8; j++ {
			res := (int(i-' ') * 9) + j + 1
			strs = append(strs, sliceData[res])
		}
		elemsMap[i] = strs
	}

	res := ""
	if check.Valid(splittedArr) {
		for _, el := range splittedArr {
			if len(el) > 0 {
				line := getAsciiArtLine(el, elemsMap)
				firstline := len(line) / 8
				if firstline > termWidth {
					fmt.Println("The output of your text does not fit in the terminal")
					return
				}
				res += line
			} else {
				res += string('\n')
			}
		}
	} else {
		for i := 0; i < len(splittedArr)-1; i++ { // handling empty input
			res = "\n" + res
		}
	}
	fmt.Print(res)
}

func getAsciiArtLine(str string, mapOfEl map[rune][]string) string {
	res := ""
	for i := 0; i < 8; i++ {
		for _, elem := range str {
			res += mapOfEl[elem][i]
		}
		res += "\n"
	}
	return res
}
