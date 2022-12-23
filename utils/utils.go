package utils

import (
	"strings"
	"time"
)

type DateTime time.Time

const dateLayout = "01-02-2006"

func (j *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return err
	}
	*j = DateTime(t)
	return nil
}

func (j DateTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Format(dateLayout) + `"`), nil
}
