package main

import (
	"bufio"
	"fmt"
	"langcards/models"
	"langcards/storage"
	"os"
	"regexp"
	"strings"
)

func main() {
	store := storage.NewMemoryStore()
	reader := bufio.NewReader(os.Stdin) 
	card, err := CreateCard(reader)
	if err != nil {
		fmt.Println("Creation card error")
		return
	}
	saved := store.Add(card)
	fmt.Println("Card successfully saved!")
	fmt.Println(FormatCard(saved))
	cards := store.All()
	if len(cards) == 0 {
		fmt.Println("Пока карточек нет")
		} else {
		for _, c := range cards {
			fmt.Println(FormatCard(c))
		}
	}
}

func ReadLine(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
} 

func ReadRequired(reader *bufio.Reader, prompt string) string {
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

func CreateCard(reader *bufio.Reader)  (models.Card, error)  {
	var card models.Card
	card.Word = ReadRequired(reader, "Enter a word: ")
	card.Lang =  NormalizeLang(reader, "Enter a language (use 2 letters ISO-code): ")
	card.Assoc = ReadRequired(reader, "Enter an association: ")
	card.Trans = ReadRequired(reader, "Enter a translation: ")
	
	for {
        if err := ValidateCard(card); err != nil {
            // Определяем, какое поле упало
            if fe, ok := err.(*models.FieldError); ok {
                fmt.Println("Error:", fe.Msg)
                switch fe.Field {
                case "word":
                    card.Word = ReadRequired(reader, "Enter a word: ")
                case "lang":
                    card.Lang = NormalizeLang(reader, "Enter a language (2-letter ISO code): ")
                case "assoc":
                    card.Assoc = ReadRequired(reader, "Enter an association: ")
                default:
                    // на всякий случай — повторим весь ввод
                    card.Word  = ReadRequired(reader, "Enter a word: ")
                    card.Lang  = NormalizeLang(reader, "Enter a language (2-letter ISO code): ")
                    card.Assoc = ReadRequired(reader, "Enter an association: ")
                    card.Trans = ReadLine(reader, "Enter a translation (optional): ")
                }
                // и крутим валидатор снова
                continue
            }
            // если пришла не FieldError (маловероятно) — вернём обычную ошибку
            return models.Card{}, err
        }
        // всё прошло — возвращаем валидную карточку
        return card, nil
    }
}

func FormatCard(c models.Card) string{
	resultFormat := fmt.Sprintf("[#%d] 🗣 %s [%s] — 💡 %s — 📘 %s", c.ID, c.Word, c.Lang, c.Assoc, c.Trans)
	return resultFormat
}

var reLang = regexp.MustCompile(`^[a-z]{2}$`)

func ValidateCard(c models.Card) error {
	if strings.TrimSpace(c.Word) == "" {
		return &models.FieldError{Field: "word", Msg: "word cannot be empty"}
	}

	if strings.TrimSpace(c.Lang) == "" {
		return &models.FieldError{Field: "lang", Msg: "language code is required"} 
	}

	if !reLang.MatchString(c.Lang) {
		return &models.FieldError{Field: "lang", Msg: "use 2-letter ISO code, e.g. it, en, de"}
	}

	if strings.TrimSpace(c.Assoc) == "" {
		return &models.FieldError{Field: "assoc", Msg: "association cannot be empty"} 
	}

	if strings.TrimSpace(c.Trans) == "" {
		return &models.FieldError{Field: "trans", Msg: "translation cannot be empty"} 
	}

	return nil
}


 
