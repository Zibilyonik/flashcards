package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
	WrongCount uint32 `json:"wrongCount"`
}

func addCard(cards []Card, logs *[]string) ([]Card, *[]string) {
	fmt.Println("Input the term:")
	*logs = append(*logs, "Input the term:")
	var appended []Card
	var card Card
	var term = readLine(logs)
	for index := range cards {
		if cards[index].Term == term {
			fmt.Printf("The term \"%s\" already exists. Try again:\n", cards[index].Term)
			*logs = append(*logs, fmt.Sprintf("The term \"%s\" already exists. Try again:\n", cards[index].Term))
			term = readLine(logs)
			index--
		}
	}
	card.Term = term
	fmt.Println("Input the definition:")
	*logs = append(*logs, "Input the definition:")
	def := readLine(logs)
	for index := range cards {
		if cards[index].Definition == def {
			fmt.Printf("The definition \"%s\" already exists. Try again:\n", cards[index].Definition)
			*logs = append(*logs, fmt.Sprintf("The definition \"%s\" already exists. Try again:\n", cards[index].Definition))
			def = readLine(logs)
			index--
		}
	}
	card.Definition = def
	card.WrongCount = 0
	appended = append(cards, card)
	fmt.Printf("The pair (\"%s\": \"%s\") has been added.\n", term, def)
	*logs = append(*logs, fmt.Sprintf("The pair (\"%s\": \"%s\") has been added.\n", term, def))
	return appended, logs
}

func removeCard(cards []Card, logs *[]string) ([]Card, *[]string) {
	fmt.Println("Which card?")
	*logs = append(*logs, "Which card?")
	var removed []Card
	card := readLine(logs)
	if len(cards) == 0 {
		fmt.Printf("Can't remove \"%s\": there is no such card.\n", card)
		*logs = append(*logs, fmt.Sprintf("Can't remove \"%s\": there is no such card.\n", card))
		return cards, logs
	}
	for index := range cards {
		if cards[index].Term == card {
			removed = append(cards[:index], cards[index+1:]...)
			fmt.Println("The card has been removed.")
			*logs = append(*logs, "The card has been removed.")
			return removed, logs
		}
	}
	fmt.Printf("Can't remove \"%s\": there is no such card.\n", card)
	*logs = append(*logs, fmt.Sprintf("Can't remove \"%s\": there is no such card.\n", card))
	return cards, logs
}

func readLine(logs *[]string) string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	logger(logs, line)
	return strings.TrimSpace(line)
}

func logger(logs *[]string, input string) *[]string {
	*logs = append(*logs, input)
	return logs
}

func importCards(logs *[]string) ([]Card, *[]string) {
	var cards []Card
	fmt.Println("File name:")
	*logs = append(*logs, "File name:")
	fileName := readLine(logs)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found.", err)
		return cards, logs
	}
	defer file.Close()
	cardsJSON, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return cards, logs
	}
	json.Unmarshal(cardsJSON, &cards)
	fmt.Printf("%d cards have been loaded.\n", len(cards))
	return cards, logs
}

func exportCards(cards []Card, logs *[]string) *[]string {
	fmt.Println("File name:")
	title := readLine(logs)
	file, err := os.Create(title)
	if err != nil {
		log.Fatal(err)
	}
	cardsJSON, _ := json.MarshalIndent(cards, "", " ")
	file.Write(cardsJSON)
	fmt.Printf("%d cards have been saved", len(cards))
	*logs = append(*logs, fmt.Sprintf("%d cards have been saved", len(cards)))
	return logs
}

func playGame(cards []Card, logs *[]string) ([]Card, *[]string) {
	fmt.Println("How many times to ask?")
	*logs = append(*logs, "How many times to ask?")
	ask := readLine(logs)
	count, err := strconv.Atoi(ask)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		*logs = append(*logs, fmt.Sprintln("Error converting string to int:", err))
	}
	for i := 0; i < count; i++ {
		var wrongDefinition bool = false
		var question int = 0
		if len(cards) == 0 {
			fmt.Println("There are no cards added.")
			*logs = append(*logs, "There are no cards added.")
			break
		} else if len(cards) == 1 {
			question = 0
		} else {
			question = rand.Intn(len(cards))
			fmt.Println(question)
		}
		fmt.Printf("Print the definition of \"%s\" \n", cards[question].Term)
		*logs = append(*logs, fmt.Sprintf("Print the definition of \"%s\" \n", cards[question].Term))
		ans := readLine(logs)
		if ans == cards[question].Definition {
			*logs = append(*logs, "Correct!")
			fmt.Println("Correct!")
			continue
		} else {
			for j := 0; j < len(cards); j++ {
				if ans == cards[j].Definition {
					fmt.Printf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\" \n", cards[question].Definition, cards[j].Term)
					*logs = append(*logs, fmt.Sprintf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\" \n", cards[question].Definition, cards[j].Term))
					wrongDefinition = true
					break
				}
			}
			if !wrongDefinition {
				fmt.Printf("Wrong. The right answer is \"%s\" \n", cards[question].Definition)
				*logs = append(*logs, fmt.Sprintf("Wrong. The right answer is \"%s\" \n", cards[question].Definition))
			}
			cards[question].WrongCount++
		}
	}
	return cards, logs
}

func logCards(logs *[]string) *[]string {
	fmt.Println("File name:")
	*logs = append(*logs, "File name:")
	title := readLine(logs)
	file, err := os.Create(title)
	if err != nil {
		log.Fatal(err)
	}
	for i := range *logs {
		file.WriteString((*logs)[i])
	}
	fmt.Printf("The log has been saved.\n")
	*logs = append(*logs, fmt.Sprintf("The log has been saved.\n"))
	return logs
}

func hardestCard(cards []Card, logs *[]string) *[]string {
	var hardest []Card
	var max uint32
	for i := range cards {
		if cards[i].WrongCount > max {
			max = cards[i].WrongCount
		}
	}
	for i := range cards {
		if cards[i].WrongCount == max {
			hardest = append(hardest, cards[i])
		}
	}
	if max == 0 {
		fmt.Println("There are no cards with errors.")
		*logs = append(*logs, "There are no cards with errors.")
	} else if len(hardest) == 1 {
		fmt.Printf("The hardest card is \"%s\". You have %d errors answering it. \n", hardest[0].Term, hardest[0].WrongCount)
		*logs = append(*logs, fmt.Sprintf("The hardest card is \"%s\". You have %d errors answering it. \n", hardest[0].Term, hardest[0].WrongCount))
	} else {
		fmt.Printf("The hardest cards are \"%s\"", hardest[0].Term)
		*logs = append(*logs, fmt.Sprintf("The hardest cards are \"%s\"", hardest[0].Term))
		for i := 1; i < len(hardest); i++ {
			fmt.Printf(", \"%s\"", hardest[i].Term)
			*logs = append(*logs, fmt.Sprintf(", \"%s\"", hardest[i].Term))
		}
		fmt.Printf(". You have %d errors answering them. \n", hardest[0].WrongCount)
		*logs = append(*logs, fmt.Sprintf(". You have %d errors answering them. \n", hardest[0].WrongCount))
	}
	return logs
}

func resetStats(cards []Card, logs *[]string) ([]Card, *[]string) {
	for i := range cards {
		cards[i].WrongCount = 0
	}
	fmt.Println("Card statistics has been reset.")
	*logs = append(*logs, "Card statistics have been reset.")
	return cards, logs
}

func main() {
	var cards []Card
	var logs = new([]string)
	for {
		fmt.Println("Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats):")
		*logs = append(*logs, fmt.Sprintln("Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats):"))
		action := readLine(logs)
		switch action {
		case "add":
			cards, logs = addCard(cards, logs)
		case "remove":
			cards, logs = removeCard(cards, logs)
		case "import":
			cards, logs = importCards(logs)
		case "export":
			logs = exportCards(cards, logs)
		case "print":
			fmt.Println(cards)
		case "ask":
			cards, logs = playGame(cards, logs)
		case "exit":
			fmt.Println("Bye bye!")
			return
		case "log":
			logs = logCards(logs)
		case "hardest card":
			logs = hardestCard(cards, logs)
		case "reset stats":
			cards, logs = resetStats(cards, logs)
		}
	}
}
