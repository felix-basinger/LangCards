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