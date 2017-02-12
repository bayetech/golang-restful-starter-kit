package daos

import (
  "github.com/bayetech/golang-restful-starter-kit/app"
  "github.com/bayetech/golang-restful-starter-kit/models"
)

type ProductDAO struct{}

func NewProductDAO() *ProductDAO {
  return &ProductDAO{}
}

func (dao *ProductDAO) Get(rs app.RequestScope, id int) (*models.Product, error) {
  var product models.Product
  err := rs.Tx().Select().From("products").Model(id, &product)
  return &product, err
}