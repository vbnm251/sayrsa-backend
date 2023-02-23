package models

import "time"

type User struct {
	Id         int    `json:"-"`
	Username   string `json:"username"`
	Bio        string `json:"bio"`
	Password   string `json:"password"`
	PublicKey  string `json:"public key"`
	PrivateKey string
	LastOnline time.Time `json:"last online"`
}
