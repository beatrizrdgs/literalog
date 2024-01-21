package apis

import (
	"strconv"

	"github.com/literalog/library/pkg/models"
)

type GBookResponse struct {
	Kind       string `json:"kind"`
	TotalItems int64  `json:"totalItems"`
	Items      []Item `json:"items"`
}

type Item struct {
	Kind       string     `json:"kind"`
	ID         string     `json:"id"`
	Etag       string     `json:"etag"`
	SelfLink   string     `json:"selfLink"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
}

type VolumeInfo struct {
	Title               string               `json:"title"`
	Authors             []string             `json:"authors"`
	Publisher           string               `json:"publisher"`
	PublishedDate       string               `json:"publishedDate"`
	Description         string               `json:"description"`
	ReadingModes        string               `json:"readingModes"`
	PageCount           int64                `json:"pageCount"`
	PrintType           string               `json:"printType"`
	Categories          []string             `json:"categories"`
	MaturityRating      string               `json:"maturityRating"`
	AllowAnonLogging    bool                 `json:"allowAnonLogging"`
	ContentVersion      string               `json:"contentVersion"`
	ImageLinks          ImageLinks           `json:"imageLinks"`
	Language            string               `json:"language"`
	PreviewLink         string               `json:"previewLink"`
	InfoLink            string               `json:"infoLink"`
	CanonicalVolumeLink string               `json:"canonicalVolumeLink"`
	Subtitle            *string              `json:"subtitle,omitempty"`
	IndustryIdentifiers []IndustryIdentifier `json:"industryIdentifiers"`
}

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

type IndustryIdentifier struct {
	Identifier string `json:"identifier"`
}

func (g *GBookResponse) ToBooks() ([]models.Book, error) {
	books := make([]models.Book, g.TotalItems)
	for i, item := range g.Items {
		book, err := item.ToBook()
		if err != nil {
			return nil, err
		}
		books[i] = book
	}
	return books, nil
}

func (i *Item) ToBook() (models.Book, error) {
	isbns := make([]string, len(i.VolumeInfo.IndustryIdentifiers))
	for i, identifier := range i.VolumeInfo.IndustryIdentifiers {
		isbns[i] = identifier.Identifier
	}

	yearStr := i.VolumeInfo.PublishedDate[:4]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return models.Book{}, err
	}

	format := models.NewFormat(i.VolumeInfo.PrintType)

	return models.Book{
		Title:     i.VolumeInfo.Title,
		ISBN:      isbns,
		Year:      year,
		Publisher: i.VolumeInfo.Publisher,
		Language:  i.VolumeInfo.Language,
		Format:    format,
		PagesNo:   int(i.VolumeInfo.PageCount),
	}, nil
}
