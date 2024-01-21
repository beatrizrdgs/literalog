package models

import "github.com/google/uuid"

type Genre struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type GenreRequest struct {
	Name string `json:"name" bson:"name"`
}

func NewGenre(req GenreRequest) *Genre {
	return &Genre{
		Id:   uuid.NewString(),
		Name: req.Name,
	}
}
