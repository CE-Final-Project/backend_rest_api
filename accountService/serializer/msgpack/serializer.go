package msgpack

import (
	"github.com/vmihailenco/msgpack"
	"io"
)

type Account struct{}

func (a *Account) Decode(v interface{}, r io.Reader) error {

	d := msgpack.NewDecoder(r)

	return d.Decode(v)
}

func (a *Account) Encode(v interface{}, w io.Writer) error {
	e := msgpack.NewEncoder(w)

	return e.Encode(v)
}
