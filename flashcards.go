package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
		var termExists bool
		for _, t := range terms {
			if t == term {
				fmt.Println("The term already exists. Try again:")
				termReader := bufio.NewReader(os.Stdin)
				term, _ = termReader.ReadString('\n')
				term = strings.TrimSpace(term)
				i--
				termExists = true
			}
		}
		for _, d := range definitions {
			if d == def {
				fmt.Println("The definition already exists. Try again:")
				defReader := bufio.NewReader(os.Stdin)
				def, _ = defReader.ReadString('\n')
				def = strings.TrimSpace(def)
				i--
				termExists = true
			}
		}
		if termExists {
			continue
		}
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
}
