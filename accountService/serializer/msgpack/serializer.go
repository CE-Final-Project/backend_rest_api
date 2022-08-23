package msgpack

import (
	"github.com/vmihailenco/msgpack"
	"io"
)

// ToMSG serializes the given interface into a string based MessagePack format
func ToMSG(i interface{}, w io.Writer) error {
	e := msgpack.NewEncoder(w)

	return e.Encode(i)
}

// FromMSG deserializes the object from MessagePack string
// in an io.Reader to the given interface
func FromMSG(i interface{}, r io.Reader) error {
	d := msgpack.NewDecoder(r)
	return d.Decode(i)
}
