package models

type AuthModel struct {
	Username    string `json:"username" bson:"username"`
	Password    string `json:"password" bson:"username"`
	Token       string
	LoginStatus string
}
