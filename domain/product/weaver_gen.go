// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package product

import (
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
)

// weaver.InstanceOf checks.

// weaver.Router checks.

// Local stub implementations.

// Client stub implementations.

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][24]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.24.2 (codegen
version v0.24.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

// Reflect stub implementations.

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*Category)(nil)

type __is_Category[T ~struct {
	weaver.AutoMarshal
	Id   string "json:\"id\""
	Name string "json:\"name\""
}] struct{}

var _ __is_Category[Category]

func (x *Category) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Category.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Id)
	enc.String(x.Name)
}

func (x *Category) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Category.WeaverUnmarshal: nil receiver"))
	}
	x.Id = dec.String()
	x.Name = dec.String()
}

var _ codegen.AutoMarshal = (*CategoryItem)(nil)

type __is_CategoryItem[T ~struct {
	weaver.AutoMarshal
	Id         string "json:\"id\""
	CategoryId string "json:\"category_id\""
	ItemId     string "json:\"item_id\""
}] struct{}

var _ __is_CategoryItem[CategoryItem]

func (x *CategoryItem) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("CategoryItem.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Id)
	enc.String(x.CategoryId)
	enc.String(x.ItemId)
}

func (x *CategoryItem) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("CategoryItem.WeaverUnmarshal: nil receiver"))
	}
	x.Id = dec.String()
	x.CategoryId = dec.String()
	x.ItemId = dec.String()
}

var _ codegen.AutoMarshal = (*Item)(nil)

type __is_Item[T ~struct {
	weaver.AutoMarshal
	Id          string "json:\"id\""
	Name        string "json:\"name\""
	Description string "json:\"description\""
	Photo       string "json:\"photo\""
}] struct{}

var _ __is_Item[Item]

func (x *Item) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Item.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Id)
	enc.String(x.Name)
	enc.String(x.Description)
	enc.String(x.Photo)
}

func (x *Item) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Item.WeaverUnmarshal: nil receiver"))
	}
	x.Id = dec.String()
	x.Name = dec.String()
	x.Description = dec.String()
	x.Photo = dec.String()
}
