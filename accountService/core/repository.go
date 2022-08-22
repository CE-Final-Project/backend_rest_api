package core

type AccountRepository interface {
	Find(playerId string, accountId string) (*Account, error)
	Store(account *Account) error
	Remove(playerId string, accountId string) (*Account, error)
}
