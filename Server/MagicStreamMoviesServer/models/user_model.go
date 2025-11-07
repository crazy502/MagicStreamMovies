package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID              bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID          string        `json:"user_id" bson:"user_id"`
	FirstName       string        `json:"first_name" bson:"first_name" validate:"required,min=2,max=100"`
	LastName        string        `json:"last_name" bson:"last_name" validate:"required,min=2,max=100"`
	Email           string        `json:"email" bson:"email" validate:"email,required"`
	Password        string        `json:"password" bson:"password" validate:"required,min=6"`
	Role            string        `json:"role" bson:"role" validate:"oneof=ADMIN USER"`
	CreatedAt       time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time     `json:"update_at" bson:"update_at"`
	Token           string        `json:"token" bson:"token"`
	RefreshToken    string        `json:"refresh_token" bson:"refresh_token"`
	FavouriteGenres []Genre       `json:"favourite_genres" bson:"favourite_genres" validate:"required,dive"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=6"`
}

// 创建用户响应结构体，它将成为一个子集
// DTO模型，它代表数据传输对象，是一种用于在软件应用层之间传输数据的设计模式，仅用于保存数据
// 通常包含字段，属性，获取器和设置器，不包含业务逻辑
// 可序列化，设计用于轻松序列化为JSO你，XML等
// 使用场景：不暴露内部领域模型的情况下，向客户端返回数据的API响应
// 通过使用DTO，仅暴露需要暴露给客户端的数据
type UserResponse struct {
	UserId          string  `json:"user_id"`
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Email           string  `json:"email"`
	Role            string  `json:"role"`
	FavouriteGenres []Genre `json:"favourite_genres"`
}
