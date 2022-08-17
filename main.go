package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var pl = fmt.Println

func main() {
	pl("what is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		pl("Error: ", err)
	}
	pl("Hello,", name)

	// variabel
	var vname string = "akari kitou"
	var v1, v2, v3 int = 1, 2, 3
	var v4 = "hello"
	pl(vname, v1, v2, v3, v4)

	// convert
	cv := "50000"
	cv4, err := strconv.Atoi(cv)
	pl(cv4, err)
}

func IsEmail(s string) (string, error) {
	r, _ := regexp.Compile(`[\w._%+-]{1,20}@[\w.-]{2,20}.[A-za-z]{2,3}`)
	if r.MatchString(s) {
		return "valid email", nil
	} else {
		return "invalid email", fmt.Errorf("invalid email")

	}
}
