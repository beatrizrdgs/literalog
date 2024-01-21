package models

import "github.com/google/uuid"

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AuthorRequest struct {
	Name string `json:"name"`
}

func NewAuthor(req AuthorRequest) *Author {
	return &Author{
		ID:   uuid.NewString(),
		Name: req.Name,
	}
}
