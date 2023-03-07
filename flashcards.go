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
	var cardCount int
	fmt.Println("Input the number of cards:")
	fmt.Scan(&cardCount)
	terms := make([]string, cardCount)
	definitions := make([]string, cardCount)
	for i := 0; i < cardCount; i++ {
		fmt.Printf("Input the term for card #%d \n", i+1)
		termReader := bufio.NewReader(os.Stdin)
		defReader := bufio.NewReader(os.Stdin)
		term, _ := termReader.ReadString('\n')
		fmt.Printf("Input the definition for card #%d \n", i+1)
		def, _ := defReader.ReadString('\n')
		term = strings.TrimSpace(term)
		def = strings.TrimSpace(def)
		terms[i] = term
		definitions[i] = def
	}
	for i := 0; i < cardCount; i++ {
		fmt.Printf("Print the definition of \"%s\" \n", terms[i])
		ansReader := bufio.NewReader(os.Stdin)
		ans, _ := ansReader.ReadString('\n')
		ans = strings.TrimSpace(ans)
		if ans == definitions[i] {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Wrong. The right answer is \"%s\" \n", definitions[i])
		}
	}
	//second part of the project
	// 	if ans == def {
	// 		fmt.Println("Your answer is right!")
	// 	} else {
	// 		fmt.Println("Your answer is wrong...")
	// 	}
	// }
}
