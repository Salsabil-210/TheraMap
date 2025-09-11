package types

import (

)

type UserModel struct {
	ID        string `bson:"_id,omitempty" json:"id"`
	Name      string `bson:"name" json:"name"`
	Surname   string `bson:"surname" json:"surname"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"-"`
	Role      string `bson:"role" json:"role"`
	CreatedAt int64  `bson:"created_at" json:"created_at"`
	UpdatedAt int64  `bson:"updated_at" json:"updated_at"`
}

func (e UserModel) TableName() string {
	return "users"
}