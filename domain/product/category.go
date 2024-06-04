package product

import (
	"github.com/ServiceWeaver/weaver"
)

type Category struct {
	weaver.AutoMarshal
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CategoryRepositoryConfig struct {
	Dsn string `toml:"dsn"`
}


