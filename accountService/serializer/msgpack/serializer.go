package msgpack

import (
	"fmt"
	"github.com/ce-final-project/backend_rest_api/accountService/core"
	"github.com/vmihailenco/msgpack"
)

type Account struct{}

func (a *Account) Decode(input []byte) (*core.Account, error) {
	account := &core.Account{}
	if err := msgpack.Unmarshal(input, account); err != nil {
		return nil, fmt.Errorf("%v serializer.Account.Decode", err)
	}
	return account, nil
}

func (a *Account) Encode(input *core.Account) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("%v serializer.Account.Encode", err)
	}
	return rawMsg, nil
}
