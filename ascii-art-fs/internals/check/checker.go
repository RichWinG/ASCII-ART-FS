package check

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Valid(s []string) bool {
	for _, el := range s {
		if len(el) > 0 {
			return true
		}
	}
	return false
}

func Ascii(input string) bool {
	for _, el := range input {
		if (el < ' ' || el > '~') && el != '\n' {
			fmt.Println("Your input should be in ascii")
			return false
		}
	}
	return true
}

func GetTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, _ := cmd.Output()
	width, _ := strconv.Atoi(strings.Fields(string(output))[1])
	return width
}

// func Hash(data []byte) bool {
// 	hashSt := "ac85e83127e49ec42487f272d9b9db8b"
// 	if hashSt != fmt.Sprintf("%x", md5.Sum(data)) {
// 		fmt.Println("Standard.txt was changed")
// 		return false
// 	}
// 	return true
// }
