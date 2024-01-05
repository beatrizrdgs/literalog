package models

import (
	"strings"

	"github.com/google/uuid"
)

type Book struct {
	Id             string   `json:"id" bson:"_id"`
	Title          string   `json:"title" bson:"title"`
	AuthorId       string   `json:"author_id" bson:"author_id,omitempty"`
	Isbn           []string `json:"isbn" bson:"isbn,omitempty"`
	SeriesId       string   `json:"series_id" bson:"series_id,omitempty"`
	SeriesPosition int      `json:"series_position" bson:"series_position,omitempty"`
	Year           int      `json:"year" bson:"year,omitempty"`
	Publisher      string   `json:"publisher" bson:"publisher,omitempty"`
	Language       string   `json:"language" bson:"language,omitempty"`
	Format         Format   `json:"format" bson:"format,omitempty"`
	PagesNo        int      `json:"pages_no" bson:"pages_no,omitempty"`
	HoursNo        int      `json:"hours_no" bson:"hours_no,omitempty"`
	Genre          []string `json:"genre" bson:"genre,omitempty"`
	Blurb          string   `json:"blurb" bson:"blurb,omitempty"`
	Cover          string   `json:"cover" bson:"cover,omitempty"`
	NotABook       bool     `json:"not_a_book" bson:"not_a_book"`
}

type BookRequest struct {
	Title          string   `json:"title" bson:"title"`
	AuthorId       string   `json:"author_id" bson:"author_id,omitempty"`
	Isbn           []string `json:"isbn" bson:"isbn,omitempty"`
	SeriesId       string   `json:"series_id" bson:"series_id,omitempty"`
	SeriesPosition int      `json:"series_position" bson:"series_position,omitempty"`
	Year           int      `json:"year" bson:"year,omitempty"`
	Publisher      string   `json:"publisher" bson:"publisher,omitempty"`
	Language       string   `json:"language" bson:"language,omitempty"`
	Format         Format   `json:"format" bson:"format,omitempty"`
	PagesNo        int      `json:"pages_no" bson:"pages_no,omitempty"`
	HoursNo        int      `json:"hours_no" bson:"hours_no,omitempty"`
	Genre          []string `json:"genre" bson:"genre,omitempty"`
	Blurb          string   `json:"blurb" bson:"blurb,omitempty"`
	Cover          string   `json:"cover" bson:"cover,omitempty"`
	NotABook       bool     `json:"not_a_book" bson:"not_a_book"`
}

type Format string

const (
	Hardcover Format = "Hardcover"
	Paperback Format = "Paperback"
	Digital   Format = "Digital"
	Audio     Format = "Audio"
)

func NewFormat(s string) Format {
	switch strings.ToLower(s) {
	case "hardcover":
		return Hardcover
	case "paperback":
		return Paperback
	case "digital":
		return Digital
	case "audio":
		return Audio
	default:
		return ""
	}
}

func NewBook(req BookRequest) *Book {
	return &Book{
		Id:             uuid.NewString(),
		Title:          req.Title,
		AuthorId:       req.AuthorId,
		Isbn:           req.Isbn,
		SeriesId:       req.SeriesId,
		SeriesPosition: req.SeriesPosition,
		Year:           req.Year,
		Publisher:      req.Publisher,
		Language:       req.Language,
		Format:         req.Format,
		PagesNo:        req.PagesNo,
		HoursNo:        req.HoursNo,
		Genre:          req.Genre,
		Blurb:          req.Blurb,
		Cover:          req.Cover,
		NotABook:       req.NotABook,
	}
}
