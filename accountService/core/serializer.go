package core

import "io"

type AccountSerializer interface {
	Decode(v interface{}, r io.Reader) error
	Encode(v interface{}, w io.Writer) error
}
