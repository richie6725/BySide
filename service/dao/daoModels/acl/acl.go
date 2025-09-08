package aclDaoModel

import (
	"time"
)

type FieldName string

func (f FieldName) String() string {
	return string(f)
}

const (
	Username  FieldName = "username"
	Password  FieldName = "password"
	Roles     FieldName = "roles"
	CreatedAt FieldName = "created_at"
	UpdatedAt FieldName = "updated_at"
)

type User struct {
	Username  string    `bson:"username" json:"username"`
	Password  string    `bson:"password,omitempty" json:"-"`
	Roles     []string  `bson:"roles" json:"roles"` // ex: ["admin", "editor"]
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type Role struct {
	Name        string   `bson:"name" json:"name"` // ex: "admin"
	Permissions []string `bson:"permissions" json:"permissions"`
}

type Query struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
