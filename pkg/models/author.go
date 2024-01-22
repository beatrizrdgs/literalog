package models

import "github.com/google/uuid"

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AuthorRequest struct {
	Name string `json:"name"`
}

func (r *AuthorRequest) ToAuthor() *Author {
	return NewAuthor(r.Name)
}

func NewAuthor(name string) *Author {
	return &Author{
		ID:   uuid.NewString(),
		Name: name,
	}
}
