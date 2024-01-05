package models

import "github.com/google/uuid"

type Series struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type SeriesRequest struct {
	Name string `json:"name"`
}

func NewSeries(req SeriesRequest) *Series {
	return &Series{
		Id:   uuid.NewString(),
		Name: req.Name,
	}
}
