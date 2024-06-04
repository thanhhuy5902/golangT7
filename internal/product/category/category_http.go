package category

import (
	"cakho.com/tudye/domain/product"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CategoryHttpDelivery struct {
	api          *echo.Group
	categoryRepo CategoryRepository
}

func (c *CategoryHttpDelivery) Get(e echo.Context) error {
	categories, err := c.categoryRepo.FindAll(e.Request().Context())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, categories)
}

func (c *CategoryHttpDelivery) Create(e echo.Context) error {
	category := new(product.Category)
	if err := e.Bind(category); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	err := c.categoryRepo.Create(e.Request().Context(), *category)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, category)
}

func (c *CategoryHttpDelivery) Update(e echo.Context) error {
	category := new(product.Category)
	if err := e.Bind(category); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	err := c.categoryRepo.Update(e.Request().Context(), *category)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, category)
}

func (c *CategoryHttpDelivery) Delete(e echo.Context) error {
	id := e.Param("id")
	if id == "" {
		return e.JSON(http.StatusBadRequest, "id is required")
	}

	err := c.categoryRepo.Delete(e.Request().Context(), id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, "success")
}

func (c *CategoryHttpDelivery) GetByCategoryId(e echo.Context) error {
	categoryId := e.Param("categoryId")
	if categoryId == "" {
		return e.JSON(http.StatusBadRequest, "categoryId is required")
	}
	categoryItems, err := c.categoryRepo.FindById(e.Request().Context(), categoryId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)

	}
	return e.JSON(http.StatusOK, categoryItems)
}

func NewCategoryHttpDelivery(api *echo.Group, categoryRepo CategoryRepository) *CategoryHttpDelivery {
	categoryHttpDelivery := &CategoryHttpDelivery{
		api:          api,
		categoryRepo: categoryRepo,
	}
	categoryHttpDelivery.api.GET("", categoryHttpDelivery.Get)
	categoryHttpDelivery.api.POST("", categoryHttpDelivery.Create)
	categoryHttpDelivery.api.PUT("", categoryHttpDelivery.Update)
	categoryHttpDelivery.api.DELETE("/:id", categoryHttpDelivery.Delete)
	return categoryHttpDelivery
}
