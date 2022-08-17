package main

import (
	"bufio"
	"fmt"
	"os"
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
}
