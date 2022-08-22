package types

import (
	"encoding/json"
	"strings"
	"time"
)

type ReleaseDate time.Time

func (j *ReleaseDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = ReleaseDate(t)
	return nil
}

func (j ReleaseDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

func (j ReleaseDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
