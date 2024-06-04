package user

import "github.com/ServiceWeaver/weaver"

type User struct {
	weaver.AutoMarshal
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	DisplayName string `json:"display_name"`
	PhotoURL string `json:"photo_url"`
}

type UserRepositoryConfig struct {
	DsnUser string `toml:"dsn"`
}
