package models

import "github.com/google/uuid"

type Genre struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type GenreRequest struct {
	Name string `json:"name" bson:"name"`
}

func NewGenre(req GenreRequest) *Genre {
	return &Genre{
		ID:   uuid.NewString(),
		Name: req.Name,
	}
}
