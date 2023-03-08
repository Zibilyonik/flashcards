package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cardInit(cardCount int) ([]string, []string) {
	terms := make([]string, cardCount)
	definitions := make([]string, cardCount)
	for i := 0; i < cardCount; i++ {
		fmt.Printf("Input the term for card #%d \n", i+1)
		termReader := bufio.NewReader(os.Stdin)
		term, _ := termReader.ReadString('\n')
		term = strings.TrimSpace(term)
		for index, t := range terms {
			if t == term {
				fmt.Printf("The term \"%s\" already exists. Try again:\n", t)
				termReader := bufio.NewReader(os.Stdin)
				term, _ = termReader.ReadString('\n')
				term = strings.TrimSpace(term)
				index--
				continue
			}
		}
		terms[i] = term
		fmt.Printf("Input the definition for card #%d \n", i+1)
		defReader := bufio.NewReader(os.Stdin)
		def, _ := defReader.ReadString('\n')
		def = strings.TrimSpace(def)
		for index, d := range definitions {
			if d == def {
				fmt.Printf("The definition \"%s\" already exists. Try again:\n", d)
				defReader := bufio.NewReader(os.Stdin)
				def, _ = defReader.ReadString('\n')
				def = strings.TrimSpace(def)
				index--
			}
		}
		definitions[i] = def
	}
	return terms, definitions
}

func playGame(terms []string, definitions []string) {
	for i := 0; i < len(terms); i++ {
		var wrongDefinition bool
		fmt.Printf("Print the definition of \"%s\" \n", terms[i])
		ansReader := bufio.NewReader(os.Stdin)
		ans, _ := ansReader.ReadString('\n')
		ans = strings.TrimSpace(ans)
		if ans == definitions[i] {
			fmt.Println("Correct!")
		} else {
			for j := 0; j < len(terms); j++ {
				if ans == definitions[j] {
					fmt.Printf("Wrong. The right answer is \"%s\". You've just written the definition of \"%s\" \n", definitions[i], terms[j])
					wrongDefinition = true
					break
				}
			}
			if !wrongDefinition {
				fmt.Printf("Wrong. The right answer is \"%s\" \n", definitions[i])
			}
		}
	}
}

func main() {
	var cardCount int
	fmt.Println("Input the number of cards:")
	fmt.Scan(&cardCount)
	terms, definitions := cardInit(cardCount)
	playGame(terms, definitions)
}
