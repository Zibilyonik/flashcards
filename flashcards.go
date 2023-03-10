package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var terms = new([]string)
var definitions = new([]string)
var cards = new([]string)

func addCard(terms *[]string, definitions *[]string) ([]string, []string) {
	fmt.Println("Input the term:")
	var term string
	fmt.Scanln(&term)
	for index, t := range *terms {
		if t == term {
			fmt.Printf("The term \"%s\" already exists. Try again:\n", t)
			fmt.Scanf("%s", term)
			index--
		}
	}
	*terms = append(*terms, term)
	fmt.Println("Input the definition:")
	var def string
	fmt.Scanln(&def)
	for index, d := range *definitions {
		if d == def {
			fmt.Printf("The definition \"%s\" already exists. Try again:\n", d)
			fmt.Scanf("%s", def)
			index--
		}
	}
	*definitions = append(*definitions, def)
	fmt.Println(*terms, "+", *definitions)
	return *terms, *definitions
}

func removeCard(terms *[]string, definitions *[]string) ([]string, []string) {
	fmt.Println("Which card?")
	cardReader := bufio.NewReader(os.Stdin)
	card, _ := cardReader.ReadString('\n')
	card = strings.TrimSpace(card)
	for index, t := range *terms {
		if t == card {
			*terms = append((*terms)[:index], (*terms)[index+1:]...)
			*definitions = append((*definitions)[:index], (*definitions)[index+1:]...)
			fmt.Println("The card has been removed.")
			break
		} else {
			fmt.Printf("Can't remove \"%s\": there is no such card.\n", card)
			break
		}
	}
	return *terms, *definitions
}

func importCards(terms *[]string, definitions *[]string) ([]string, []string) {
	fmt.Println("File name:")
	fileReader := bufio.NewReader(os.Stdin)
	file, _ := fileReader.ReadString('\n')
	file = strings.TrimSpace(file)
	data, err := os.Open(file)
	if err != nil {
		fmt.Println("File not found.")
	}
	decoder := json.NewDecoder(data)
	decoder.Decode(&terms)
	decoder.Decode(&definitions)
	return *terms, *definitions
}

func exportCards(terms []string, definitions []string) {
	var title string
	var cards = new([]string)
	fmt.Println("File name:")
	fmt.Scan(&title)
	file, err := os.Create(title)
	if err != nil {
		log.Fatal(err)
	}
	for i := range terms {
		(*cards)[i] = (terms)[i] + "\n" + (definitions)[i]
		_, err2 := fmt.Fprintln(file, (*cards)[i])
		if err2 != nil {
			log.Fatal(err2)
		}
	}

	defer file.Close()

	fmt.Printf("%d cards written successfully!", len(*cards))
}

func playGame(terms *[]string, definitions *[]string) {
	for i := 0; i < len(*terms); i++ {
		var wrongDefinition bool
		fmt.Printf("Print the definition of \"%s\" \n", (*terms)[i])
		ansReader := bufio.NewReader(os.Stdin)
		ans, _ := ansReader.ReadString('\n')
		ans = strings.TrimSpace(ans)
		if ans == (*definitions)[i] {
			fmt.Println("Correct!")
		} else {
			for j := 0; j < len(*terms); j++ {
				if ans == (*definitions)[j] {
					fmt.Printf("Wrong. The right answer is \"%s\". You've just written the definition of \"%s\" \n", (*definitions)[i], (*terms)[j])
					wrongDefinition = true
					break
				}
			}
			if !wrongDefinition {
				fmt.Printf("Wrong. The right answer is \"%s\" \n", (*definitions)[i])
			}
		}
	}
}

func main() {
	var action string
	for {
		fmt.Println("Input the action (add, remove, import, export, ask, exit):")
		fmt.Scanln(&action)
		switch action {
		case "add":
			addCard(terms, definitions)
		case "remove":
			removeCard(terms, definitions)
		case "import":
			importCards(terms, definitions)
		case "export":
			exportCards(*terms, *definitions)
		case "print":
			fmt.Println(cards)
		case "ask":
			playGame(terms, definitions)
		case "exit":
			fmt.Println("Bye bye!")
			return
		}
	}
}
