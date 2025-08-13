package models

type Card struct {
	ID    int64  `json:"id"`
    Word  string `json:"word"`
    Lang  string `json:"lang"`
    Assoc string `json:"assoc"`
    Trans string `json:"trans"`
}

