package main

import (
	"bufio"
	"fmt"
	"langcards/models"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) 
	result := CreateCard(reader)
	fmt.Println(result)
}

func ReadLine(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
} 

func ReadRequire(reader *bufio.Reader, prompt string) string {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		trimmed := strings.TrimSpace(input)
		if trimmed == "" {
			fmt.Println("String cannot be empty!")
			continue
		}  
		return trimmed
	}
}

func NormalizeLang(reader *bufio.Reader, prompt string) string {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		trimmed := strings.TrimSpace(input)
		if len(trimmed) != 2 {
			fmt.Println("Language code must be 2 characters long")
			continue
		}
		return strings.ToLower(trimmed)
	}
}

func CreateCard(reader *bufio.Reader) string  {
	var card models.Card
	card.Word = ReadRequire(reader, "Enter a word: ")
	card.Lang =  NormalizeLang(reader, "Enter a language (use 2 letters ISO-code): ")
	card.Assoc = ReadLine(reader, "Enter an association: ")
	card.Trans = ReadLine(reader, "Enter a translation: ")
	readyCard := FormatCard(card)
	return readyCard
}

func FormatCard(c models.Card) string{
	resultFormat := fmt.Sprintf("[#%d] 🗣 %s [%s] — 💡 %s — 📘 %s", c.ID, c.Word, c.Lang, c.Assoc, c.Trans)
	return resultFormat
}