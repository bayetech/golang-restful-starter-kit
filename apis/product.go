package apis

import (
	"strconv"

	"github.com/bayetech/golang-restful-starter-kit/app"
	"github.com/bayetech/golang-restful-starter-kit/models"
	"github.com/go-ozzo/ozzo-routing"
)

type (
	// productService specifies the interface for the product service needed by productResource.
	productService interface {
		Get(rs app.RequestScope, id int) (*models.Product, error)
		GetAll(rs app.RequestScope) (*[]models.Product, error)
	}

	// productResource defines the handlers for the CRUD APIs.
	productResource struct {
		service productService
	}
)

func ServeProductResource(rg *routing.RouteGroup, service productService) {
	r := &productResource{service}
	rg.Get("/products/<id>", r.get)
	rg.Get("/products", r.getAll)
}

func (r *productResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *productResource) getAll(c *routing.Context) error {
	response, err := r.service.GetAll(app.GetRequestScope(c))
	if err != nil {
		return err
	}

	return c.Write(response)
}
