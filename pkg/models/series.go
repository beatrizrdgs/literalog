package models

import "github.com/google/uuid"

type Series struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type SeriesRequest struct {
	Name string `json:"name"`
}

func NewSeries(req SeriesRequest) *Series {
	return &Series{
		ID:   uuid.NewString(),
		Name: req.Name,
	}
}
