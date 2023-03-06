package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// first part of the project
	// fmt.Println("Card:")
	// fmt.Println("Potato.")
	// fmt.Println("Definition:")
	// fmt.Println("POTATO")
	termReader := bufio.NewReader(os.Stdin)
	defReader := bufio.NewReader(os.Stdin)
	ansReader := bufio.NewReader(os.Stdin)
	term, _ := termReader.ReadString('\n')
	def, _ := defReader.ReadString('\n')
	ans, _ := ansReader.ReadString('\n')
	term = strings.TrimSpace(term)
	def = strings.TrimSpace(def)
	ans = strings.TrimSpace(ans)
	if ans == def {
		fmt.Println("Your answer is right!")
	} else {
		fmt.Println("Your answer is wrong...")
	}
}
