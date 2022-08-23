package json

import (
	"encoding/json"
	"io"
)

type Account struct{}

func (a *Account) Decode(v interface{}, r io.Reader) error {

	d := json.NewDecoder(r)

	return d.Decode(v)
}

func (a *Account) Encode(v interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(v)
}
