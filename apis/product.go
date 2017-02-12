package apis

import (
  "strconv"

  "github.com/go-ozzo/ozzo-routing"
  "github.com/bayetech/golang-restful-starter-kit/app"
  "github.com/bayetech/golang-restful-starter-kit/models"
)

type (
  // productService specifies the interface for the product service needed by productResource.
  productService interface {
    Get(rs app.RequestScope, id int) (*models.Product, error)
  }

  // productResource defines the handlers for the CRUD APIs.
  productResource struct {
    service productService
  }
)

// ServeProduct sets up the routing of product endpoints and the corresponding handlers.
func ServeProductResource(rg *routing.RouteGroup, service productService) {
  r := &productResource{service}
  rg.Get("/product/<id>", r.get)
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
