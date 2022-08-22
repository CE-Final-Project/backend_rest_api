package core

type Account struct {
	PlayerId     string `json:"player_id" bson:"player_id" msgpack:"player_id"`
	Username     string `json:"username" bson:"username" msgpack:"username"`
	Email        string `json:"email" bson:"email" msgpack:"email" validate:"empty=false format=email"`
	PasswordHash string `json:"password_hash" bson:"password_hash" msgpack:"password_hash"`
	CreatedAt    int64  `json:"created_at" bson:"created_at" msgpack:"created_at"`
}
