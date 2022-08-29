package domain

type Account struct {
	AccountID    uint64 `json:"account_id" bson:"_id" db:"account_id"`
	PlayerID     string `json:"player_id" bson:"player_id" db:"player_id"`
	Username     string `json:"username" bson:"username" db:"username"`
	Email        string `json:"email" bson:"email" db:"email"`
	PasswordHash string `json:"password_hash" bson:"password_hash" db:"password_hash"`
	IsBan        bool   `json:"is_ban" bson:"is_ban" db:"is_ban"`
	CreatedAt    int64  `json:"created_at" bson:"created_at" db:"created_at"`
}

type AccountRequest struct {
	Username string `json:"username" bson:"username" db:"username"`
	Email    string `json:"email" bson:"email" db:"email"`
	Password string `json:"password" bson:"password" db:"password"`
}

type AccountResponse struct {
	AccountID uint64 `json:"account_id" bson:"_id" db:"account_id"`
	PlayerID  string `json:"player_id" bson:"player_id" db:"player_id"`
	Username  string `json:"username" bson:"username" db:"username"`
	Email     string `json:"email" bson:"email" db:"email"`
	IsBan     bool   `json:"is_ban" bson:"is_ban" db:"is_ban"`
	CreatedAt int64  `json:"created_at" bson:"created_at" db:"created_at"`
}
