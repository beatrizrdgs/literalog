package models

type Book struct {
	Id      string   `json:"id" bson:"_id"`
	Title   string   `json:"title" bson:"title"`
	Author  string   `json:"author" bson:"author"`
	Isbn    string   `json:"isbn" bson:"isbn"`
	Year    int      `json:"year" bson:"year"`
	Lang    string   `json:"lang" bson:"lang"`
	Format  Format   `json:"format" bson:"format"`
	PagesNo int      `json:"pages_no" bson:"pages_no"`
	Genre   []string `json:"genre" bson:"genre"`
	Blurb   string   `json:"blurb" bson:"blurb"`
	Cover   string   `json:"cover" bson:"cover"`
}

type BookRequest struct {
	Title   string   `json:"title"`
	Author  string   `json:"author"`
	Isbn    string   `json:"isbn"`
	Year    int      `json:"year"`
	Lang    string   `json:"lang"`
	Format  Format   `json:"format"`
	PagesNo int      `json:"pages_no"`
	Genre   []string `json:"genre"`
	Blurb   string   `json:"blurb"`
	Cover   string   `json:"cover"`
}

type Format string

const (
	HARDCOVER Format = "hardcover"
	PAPERBACK Format = "paperback"
	EBOOK     Format = "ebook"
	AUDIO     Format = "audio"
)

func NewFormat(format string) Format {
	switch format {
	case "hardcover":
		return HARDCOVER
	case "paperback":
		return PAPERBACK
	case "ebook":
		return EBOOK
	case "audio":
		return AUDIO
	default:
		return ""
	}
}

func NewBook(req BookRequest) Book {
	return Book{
		Title:   req.Title,
		Author:  req.Author,
		Isbn:    req.Isbn,
		Year:    req.Year,
		Lang:    req.Lang,
		Format:  req.Format,
		PagesNo: req.PagesNo,
		Genre:   req.Genre,
		Blurb:   req.Blurb,
		Cover:   req.Cover,
	}
}
