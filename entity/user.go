package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	UserName     *string            `json:"userName" bson:"userName" validate:"required,min=2,max=100"`
	Password     *string            `json:"password" bson:"password" validate:"required,min=6,max=100"`
	Email        *string            `json:"email" bson:"email" validate:"email,required"`
	DeviceToken  *string            `json:"device_token" bson:"device_token"`
	Token        *string            `json:"token" bson:"token"`
	RefreshToken *string            `json:"refreshToken" bson:"refreshToken"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
	UserId       string             `json:"userId" bson:"userId"`
}
