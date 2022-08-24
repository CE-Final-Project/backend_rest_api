package json

import (
	"encoding/json"
	"github.com/ce-final-project/backend_rest_api/accountService/core"
	"github.com/pkg/errors"
)

type Account struct{}

func (a *Account) Decode(input []byte) (*core.Account, error) {
	account := &core.Account{}
	if err := json.Unmarshal(input, account); err != nil {
		return nil, errors.Wrap(err, "serializer.Account.Decode")
	}
	return account, nil
}

func (a *Account) Encode(input *core.Account) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Account.Encode")
	}
	return rawMsg, nil
}
