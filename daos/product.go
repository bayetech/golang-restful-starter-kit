package daos

import (
	"github.com/bayetech/golang-restful-starter-kit/app"
	"github.com/bayetech/golang-restful-starter-kit/models"
	"github.com/go-ozzo/ozzo-dbx"
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

func (dao *ProductDAO) GetAll(rs app.RequestScope) (*[]models.Product, error) {
	products := []models.Product{}
	err := rs.Tx().Select().From("products").Where(dbx.In("category_id", 4, 6, 7, 15)).All(&products)
	return &products, err
}
