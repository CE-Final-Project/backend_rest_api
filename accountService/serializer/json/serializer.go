package json

import (
	"encoding/json"
	"fmt"
	"github.com/ce-final-project/backend_rest_api/accountService/core"
)

type Account struct{}

func (a *Account) Decode(input []byte) (*core.Account, error) {
	account := &core.Account{}
	if err := json.Unmarshal(input, account); err != nil {
		return nil, fmt.Errorf("%v serializer.Account.Decode", err)
	}
	return account, nil
}

func (a *Account) Encode(input *core.Account) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("%v serializaer.Account.Encode", err)
	}
	return rawMsg, nil
}
