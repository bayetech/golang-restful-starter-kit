package daos

import (
  "github.com/bayetech/golang-restful-starter-kit/app"
  "github.com/bayetech/golang-restful-starter-kit/models"
  "github.com/go-ozzo/ozzo-dbx"
)

type SlideDAO struct{}

func NewSlideDAO() *SlideDAO {
  return &SlideDAO{}
}

func (dao *SlideDAO) Get(rs app.RequestScope) (*[]models.Slide, error) {
  slides := []models.Slide{}
  err := rs.Tx().Select().From("slides").
    Where(dbx.NewExp("enabled_at IS NOT NULL")).
    AndWhere(dbx.HashExp{"position": "Store"}).
    OrderBy("sort asc").
    All(&slides)

  return &slides, err
}
