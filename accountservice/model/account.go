package model

type Account struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}
