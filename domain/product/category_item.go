package product

import (
	"github.com/ServiceWeaver/weaver"
)

type CategoryItem struct {
	weaver.AutoMarshal
	Id         string `json:"id"`
	CategoryId string `json:"category_id"`
	ItemId     string `json:"item_id"`
}

type CategoryItemRepositoryConfig struct {
	Dsn string `toml:"dsn"`
}


