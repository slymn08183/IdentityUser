package model

type User struct {
	Password    *string `json:"password" bson:"password"`
	Email       *string `json:"email" bson:"email"`
	DeviceToken *string `json:"device_token" bson:"deviceToken"`
}
