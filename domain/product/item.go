package product

import "github.com/ServiceWeaver/weaver"

type Item struct {
	weaver.AutoMarshal
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
}

type ItemRepositoryConfig struct {
	Dsn string `toml:"dsn"`
}
