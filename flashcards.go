package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
}

func addCard(cards []Card) []Card {
	fmt.Println("Input the term:")
	var appended []Card
	var card Card
	term := readLine()
	for index := range cards {
		if cards[index].Term == term {
			fmt.Printf("The term \"%s\" already exists. Try again:\n", cards[index].Term)
			term = readLine()
			index--
		}
	}
	card.Term = term
	fmt.Println("Input the definition:")
	def := readLine()
	for index := range cards {
		if cards[index].Definition == def {
			fmt.Printf("The definition \"%s\" already exists. Try again:\n", cards[index].Definition)
			def = readLine()
			index--
		}
	}
	card.Definition = def
	appended = append(cards, card)
	fmt.Printf("The pair (\"%s\": \"%s\") has been added.\n", term, def)
	return appended
}

func removeCard(cards []Card) []Card {
	fmt.Println("Which card?")
	var removed []Card
	card := readLine()
	if len(cards) == 0 {
		fmt.Printf("Can't remove \"%s\": there is no such card.\n", card)
		return cards
	}
	for index := range cards {
		if cards[index].Term == card {
			removed = append(cards[:index], cards[index+1:]...)
			fmt.Println("The card has been removed.")
			return removed
		}
	}
	fmt.Printf("Can't remove \"%s\": there is no such card.\n", card)
	return cards
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func importCards() []Card {
	var cards []Card
	fmt.Println("File name:")
	fileName := readLine()
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found.", err)
		return cards
	}
	defer file.Close()
	cardsJSON, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return cards
	}
	json.Unmarshal(cardsJSON, &cards)
	fmt.Printf("%d cards have been loaded.\n", len(cards))
	return cards
}

func exportCards(cards []Card) {
	fmt.Println("File name:")
	title := readLine()
	file, err := os.Create(title)
	if err != nil {
		log.Fatal(err)
	}
	cardsJSON, _ := json.MarshalIndent(cards, "", " ")
	file.Write(cardsJSON)
	fmt.Printf("%d cards have been saved", len(cards))
}

func playGame(cards []Card) {
	fmt.Println("How many times to ask?")
	ask := readLine()
	count, err := strconv.Atoi(ask)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}
	for i := 0; i < count; i++ {
		var wrongDefinition bool
		var question = rand.Intn(len(cards) - 1)
		fmt.Printf("Print the definition of \"%s\" \n", cards[question].Term)
		ans := readLine()
		if ans == cards[question].Definition {
			fmt.Println("Correct!")
		} else {
			for j := 0; j < len(cards); j++ {
				if ans == cards[j].Definition {
					fmt.Printf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\" \n", cards[question].Definition, cards[j].Term)
					wrongDefinition = true
					break
				}
			}
			if !wrongDefinition {
				fmt.Printf("Wrong. The right answer is \"%s\" \n", cards[question].Term)
			}
		}
	}
}

func main() {
	var cards []Card
	for {
		fmt.Println("Input the action (add, remove, import, export, ask, exit):")
		action := readLine()
		switch action {
		case "add":
			cards = addCard(cards)
		case "remove":
			cards = removeCard(cards)
		case "import":
			cards = importCards()
		case "export":
			exportCards(cards)
		case "print":
			fmt.Println(cards)
		case "ask":
			playGame(cards)
		case "exit":
			fmt.Println("Bye bye!")
			return
		}
	}
}
