package categoryItem

import (
	"cakho.com/tudye/domain/product"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CategoryItemHttpDelivery struct {
	api              *echo.Group
	categoryItemRepo CategoryItemRepository
}

func (c *CategoryItemHttpDelivery) GetByCategoryItemId(e echo.Context) error {
	categoryId := e.Param("categoryId")
	if categoryId == "" {
		return e.JSON(http.StatusBadRequest, "categoryId is required")
	}
	categoryItems, err := c.categoryItemRepo.FindByCategoryId(e.Request().Context(), categoryId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, categoryItems)

}

func (c *CategoryItemHttpDelivery) Create(e echo.Context) error {
	categoryItem := new(product.CategoryItem)
	if err := e.Bind(categoryItem); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	err := c.categoryItemRepo.Create(e.Request().Context(), *categoryItem)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, categoryItem)
}

func (c *CategoryItemHttpDelivery) Update(e echo.Context) error {
	categoryItem := new(product.CategoryItem)
	if err := e.Bind(categoryItem); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	err := c.categoryItemRepo.Update(e.Request().Context(), *categoryItem)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, categoryItem)
}

func (c *CategoryItemHttpDelivery) Delete(e echo.Context) error {
	id := e.Param("id")
	if id == "" {
		return e.JSON(http.StatusBadRequest, "id is required")
	}

	err := c.categoryItemRepo.Delete(e.Request().Context(), id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, id)
}

func (c *CategoryItemHttpDelivery) GetAll(e echo.Context) error {
	categoryItems, err := c.categoryItemRepo.FindAll(e.Request().Context())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, categoryItems)
}

func (c *CategoryItemHttpDelivery) GetById(e echo.Context) error {
	id := e.Param("id")
	if id == "" {
		return e.JSON(http.StatusBadRequest, "id is required")
	}

	categoryItem, err := c.categoryItemRepo.FindById(e.Request().Context(), id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, categoryItem)
}

func NewCategoryItemHttpDelivery(api *echo.Group, categoryItemRepo CategoryItemRepository) *CategoryItemHttpDelivery {
	http := &CategoryItemHttpDelivery{api: api, categoryItemRepo: categoryItemRepo}
	api.GET("/all", http.GetAll)
	api.POST("", http.Create)
	api.PUT("", http.Update)
	api.DELETE("/:id", http.Delete)
	api.GET("/:id", http.GetById)
	api.GET("/:categoryId", http.GetByCategoryItemId)

	return http
}
