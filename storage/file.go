package storage

import (
    "encoding/json"
    "errors"
    "os"
    "langcards/models"
)

type FileStore struct {
    filename string
    cards    []models.Card
    nextID   int64
}

func NewFileStore(filename string) (*FileStore, error) {
	s := &FileStore{
		filename: filename,
		cards: []models.Card{},
		nextID: 1,
	}

	f, err := os.Open(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return s, nil
		}
		return nil, err
	}
	defer f.Close()

	var loaded []models.Card
	if err := json.NewDecoder(f).Decode(&loaded); err != nil {
		return nil, err
	}
	s.cards = loaded

	var maxID int64 
	for _, c := range s.cards {
		if c.ID > maxID {
			maxID = c.ID
		}
	}

	s.nextID = maxID + 1
	if s.nextID < 1 {
		s.nextID = 1
	}


	return s, nil
} 


func (s *FileStore) save() error {
    // Создаём/очищаем файл (truncate)
    f, err := os.Create(s.filename)
    if err != nil {
        return err
    }
    defer f.Close()

    enc := json.NewEncoder(f)
    enc.SetIndent("", "  ") // делаем файл читаемым человеком
    return enc.Encode(s.cards)
}

func (s *FileStore) Add(c models.Card) (models.Card, error) {
	c.ID = s.nextID
	s.nextID++

	s.cards = append(s.cards, c)

	if err := s.save() ; err != nil {
		s.cards = s.cards[:len(s.cards) - 1]
		s.nextID--
		return models.Card{}, err
	}

	return c, nil
} 

func (s *FileStore) All() []models.Card {
    return append([]models.Card(nil), s.cards...)
}