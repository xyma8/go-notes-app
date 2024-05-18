package models

import (
	"time"
)

type MySQLTimestamp time.Time

func (t *MySQLTimestamp) Scan(value interface{}) error {
	if value == nil {
		*t = MySQLTimestamp(time.Time{})
		return nil
	}
	parsedTime, err := time.Parse("2006-01-02 15:04:05", string(value.([]byte)))
	if err != nil {
		return err
	}
	*t = MySQLTimestamp(parsedTime)
	return nil
}

func (t MySQLTimestamp) MarshalJSON() ([]byte, error) {
	formattedTime := time.Time(t).Format("2006-01-02 15:04:05")
	return []byte(`"` + formattedTime + `"`), nil
}

type Note struct {
	Id           string         `json:"id"`
	Title        string         `json:"title"`
	NoteText     string         `json:"note_text"`
	AuthorName   string         `json:"author"`
	CreationDate MySQLTimestamp `json:"creation_time"`
	IsPublic     bool           `json:"is_public"`
}
