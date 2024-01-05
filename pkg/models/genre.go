package models

import "github.com/google/uuid"

type Genre struct {
	Id  string `json:"id" bson:"_id"`
	Tag string `json:"tag " bson:"tag"`
}

type GenreRequest struct {
	Tag string `json:"tag"`
}

func NewGenre(req GenreRequest) *Genre {
	return &Genre{
		Id:  uuid.NewString(),
		Tag: req.Tag,
	}
}
