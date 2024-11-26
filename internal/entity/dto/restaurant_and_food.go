package dto

import (
	"banners_oto/internal/entity"

	"github.com/asaskevich/govalidator"
)

type Category struct {
	Id   uint64  `json:"id" valid:"-"`
	Name string  `json:"name" valid:"-"`
	Food []*Food `json:"food,omitempty" valid:"-"`
}

type CategoryArray struct {
	Payload []*Category `json:"payload" valid:"-"`
}

type RestaurantAndFood struct {
	Id               uint64      `json:"id" valid:"-"`
	Name             string      `json:"name" valid:"-"`
	ShortDescription string      `json:"short_description,omitempty" valid:"-"`
	LongDescription  string      `json:"long_description,omitempty" valid:"-"`
	ImgUrl           string      `json:"img_url" valid:"url"`
	Rating           float64     `json:"rating"`
	CommentCount     uint32      `json:"comment_count"`
	Categories       []*Category `json:"categories" valid:"-"`
}

type RestaurantAndFoodArray struct {
	Payload []*RestaurantAndFood `json:"payload" valid:"-"`
}

func (d *RestaurantAndFood) Validate() (bool, error) {
	return govalidator.ValidateStruct(d)
}

func NewCategory(c *entity.Category) *Category {
	return &Category{
		Id:   uint64(c.Id),
		Name: c.Name,
		Food: NewFoodInCategoryArr(c.Food),
	}
}

func NewCategoryArray(categories []*entity.Category) []*Category {
	if len(categories) == 0 {
		return make([]*Category, 0)
	}
	cDTO := make([]*Category, len(categories))
	for i, c := range categories {
		cDTO[i] = NewCategory(c)
	}
	return cDTO
}

func NewRestaurantAndFood(r *entity.Restaurant, categories []*entity.Category) *RestaurantAndFood {
	return &RestaurantAndFood{
		Id:               r.Id,
		Name:             r.Name,
		ShortDescription: r.ShortDescription,
		LongDescription:  r.LongDescription,
		ImgUrl:           r.ImgUrl,
		Rating:           r.Rating,
		CommentCount:     r.CommentCount,
		Categories:       NewCategoryArray(categories),
	}
}
