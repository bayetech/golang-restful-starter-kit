package models

// import "github.com/go-ozzo/ozzo-validation"
type Slide struct {
  Img string `json:"img" db:"image"`
  Url string `json:"url" db:"url"`
}

func (s Slide) AddImgHost() string {
  return "https://baye-media.oss-cn-shanghai.aliyuncs.com/uploads/slide_images/" + s.Img
}
