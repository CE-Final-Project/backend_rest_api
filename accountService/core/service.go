package core

type AccountService interface {
	Find(playerId string, accountId string) (*Account, error)
	Store(account *Account) error
	Remove(playerId string, accountId string) (*Account, error)
}
