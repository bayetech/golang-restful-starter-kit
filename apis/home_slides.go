package apis

import (
  // "strconv"

  "github.com/bayetech/golang-restful-starter-kit/app"
  "github.com/bayetech/golang-restful-starter-kit/models"
  "github.com/go-ozzo/ozzo-routing"
)

type (
  slideService interface {
    Get(rs app.RequestScope) (*[]models.Slide, error)
  }

  slideResource struct {
    service slideService
  }
)

func ServeSlideResource(rg *routing.RouteGroup, service slideService) {
  r := &slideResource{service}
  rg.Get("/home_slides", r.get)
}

func (r *slideResource) get(c *routing.Context) error {
  response, err := r.service.Get(app.GetRequestScope(c))
  if err != nil {
    return err
  }

  return c.Write(response)
}
