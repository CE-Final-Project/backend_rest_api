package core

type AccountSerializer interface {
	Decode(input []byte) (*Account, error)
	Encode(input *Account) ([]byte, error)
}
