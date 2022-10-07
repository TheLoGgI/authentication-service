package models

type UserAccount struct {
	Uid   string `json:"uid" bson:"_id,omitempty"`
	Name  string
	Email string
}

type NewUserAccountRequest struct {
	Name     string
	Email    string
	Password []byte
}
