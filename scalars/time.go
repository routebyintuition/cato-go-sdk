package scalars

import (
	"fmt"
	"io"
)

type Time string

func (t *Time) UnmarshalGQL(v interface{}) error {
	*t = Time(v.(string))

	return nil
}

func (t Time) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, string(t))
}

func (t Time) GetInt64() string {

	return string(t)
}
