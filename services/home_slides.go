package services

import (
  "github.com/bayetech/golang-restful-starter-kit/app"
  "github.com/bayetech/golang-restful-starter-kit/models"
)

type slideDAO interface {
  Get(rs app.RequestScope) (*[]models.Slide, error)
}

type SlideService struct {
  dao slideDAO
}

func NewSlideService(dao slideDAO) *SlideService {
  return &SlideService{dao}
}

func (s *SlideService) Get(rs app.RequestScope) (*[]models.Slide, error) {
  slides, err := s.dao.Get(rs)
  for i := range *slides {
    (*slides)[i].Img = (*slides)[i].AddImgHost()
  }

  return slides, err
}
