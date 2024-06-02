package models

type Logbook struct {
	Id     string `json:"id" bson:"_id"`
	UserId string `json:"user_id" bson:"user_id"`
	BookId string `json:"book_id" bson:"book_id"`
	Status Status `json:"status" bson:"status"`
}

type LogbookRequest struct {
	UserId string `json:"user_id"`
	BookId string `json:"book_id"`
	Status Status `json:"status"`
}

type Status string

const (
	READING  Status = "reading"
	READ     Status = "read"
	TOBEREAD Status = "to_be_read"
)

func (s Status) String() string {
	return string(s)
}

func NewStatus(status string) Status {
	switch status {
	case "reading":
		return READING
	case "read":
		return READ
	case "to_be_read":
		return TOBEREAD
	default:
		return ""
	}
}

func NewLogbook(req LogbookRequest) Logbook {
	return Logbook{
		UserId: req.UserId,
		BookId: req.BookId,
		Status: NewStatus(req.Status.String()),
	}
}
