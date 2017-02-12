package services

import (
  "github.com/bayetech/golang-restful-starter-kit/app"
  "github.com/bayetech/golang-restful-starter-kit/models"
)

// productDAO specifies the interface of the product DAO needed by ProductService.
type productDAO interface {
  // Get returns the product with the specified the product ID.
  Get(rs app.RequestScope, id int) (*models.Product, error)
}

// ProductService provides services related with products.
type ProductService struct {
  dao productDAO
}

func NewProductService(dao productDAO) *ProductService {
  return &ProductService{dao}
}

// Get returns the product with the specified the product ID.
func (s *ProductService) Get(rs app.RequestScope, id int) (*models.Product, error) {
  return s.dao.Get(rs, id)
}
