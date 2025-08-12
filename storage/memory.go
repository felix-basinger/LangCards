package storage

import "langcards/models"

type MemoryStore struct {
	cards []models.Card
	nextID int64
}

func (s *MemoryStore) Add(c models.Card) models.Card {
	c.ID = s.nextID
	s.nextID++
	s.cards = append(s.cards, c)
	return c
}

func (s *MemoryStore) All() []models.Card {
	out := make([]models.Card, len(s.cards))
	copy(out, s.cards)
	return out
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		cards: []models.Card{},
		nextID: 1,
	}
}