package apis

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/literalog/library/pkg/models"
)

type GBooksAPI struct {
	Key     string
	BaseURL *url.URL
}

func NewGBooksAPI(key string, baseURL string) (*GBooksAPI, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &GBooksAPI{
		Key:     os.Getenv("GBOOKS_API_KEY"),
		BaseURL: u,
	}, nil
}

func (g *GBooksAPI) Get(ctx context.Context, isbn string) (*models.Book, error) {
	u := g.BaseURL
	u = u.JoinPath("volumes")

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("q", "isbn:"+isbn)
	q.Add("key", g.Key)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gBook := new(GBookResponse)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, gBook)
	if err != nil {
		return nil, err
	}

	books, err := gBook.ToBooks()
	if err != nil {
		return nil, err
	}

	if len(books) > 0 {
		return &books[0], nil
	}

	return nil, nil
}
