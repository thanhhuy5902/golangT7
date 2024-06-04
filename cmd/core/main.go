package main

import (
	"cakho.com/tudye/internal/product/category"
	"cakho.com/tudye/internal/product/categoryItem"
	"cakho.com/tudye/internal/product/item"
	"cakho.com/tudye/internal/user"
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {

	if err := weaver.Run(context.Background(), serve); err != nil {
		log.Fatal(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
	itemRepo         weaver.Ref[item.ItemRepository]
	categoryItemRepo weaver.Ref[categoryItem.CategoryItemRepository]
	categoryRepo     weaver.Ref[category.CategoryRepository]
	userRepo         weaver.Ref[user.UserRepository]
	endpoint         weaver.Listener
}

func serve(ctx context.Context, app *app) error {

	e := echo.New()
	itemApi := e.Group("/item")
	categoryItemApi := e.Group("/categoryItem-item")
	categoryApi := e.Group("/category")
	userApi := e.Group("/user")
	_ = user.NewUserHttpDelivery(userApi, app.userRepo.Get())
	_ = item.NewItemHttpDelivery(itemApi, app.itemRepo.Get())
	_ = categoryItem.NewCategoryItemHttpDelivery(categoryItemApi, app.categoryItemRepo.Get())
	_ = category.NewCategoryHttpDelivery(categoryApi, app.categoryRepo.Get())
	fmt.Printf("hello listener available on %v\n", app.endpoint)
	return e.Start("127.0.0.1:8080")
}
