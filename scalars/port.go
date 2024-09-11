package scalars

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

type Port string

func (p *Port) UnmarshalGQL(v interface{}) error {
	var srcInt int64

	srcInt, err := v.(json.Number).Int64()
	if err != nil {
		return fmt.Errorf("UnmarshalGQL failure to convert to Int64: %s", v)
	}

	port := Port(strconv.FormatInt(srcInt, 10))
	p = &port

	return nil
}

func (p Port) MarshalGQL(w io.Writer) {
	strInt64, err := strconv.ParseInt(string(p), 10, 64)
	if err != nil {
		fmt.Println("error in Port/MarshalGQL: ", err)
		fmt.Fprint(w, 0)
		return
	}

	fmt.Fprint(w, strInt64)
}
