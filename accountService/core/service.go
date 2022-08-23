package core

type AccountService interface {
	Find(username string) ([]*Account, error)
	FindOne(playerId string) (*Account, error)
	Store(account *Account) error
	Remove(playerId string, accountId string) (*Account, error)
}
