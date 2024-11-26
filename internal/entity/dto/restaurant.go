package dto

import (
	"banners_oto/internal/entity"

	"github.com/asaskevich/govalidator"
)

// нужно будет узнать минимальную длину описания и имени
type Restaurant struct {
	Id               uint64  `json:"id" valid:"-"`
	Name             string  `json:"name" valid:"-"`
	ShortDescription string  `json:"short_description,omitempty" valid:"-"`
	LongDescription  string  `json:"long_description,omitempty" valid:"-"`
	Address          string  `json:"address,omitempty" valid:"-"`
	ImgUrl           string  `json:"img_url" valid:"url"`
	Rating           float64 `json:"rating,omitempty"`
	CommentCount     uint32  `json:"comment_count,omitempty"`
}

type RestaurantArray struct {
	Payload []*Restaurant `json:"payload" valid:"-"`
}

func (d *Restaurant) Validate() (bool, error) {
	return govalidator.ValidateStruct(d)
}

func NewRestaurant(r *entity.Restaurant) *Restaurant {
	return &Restaurant{
		Id:               r.Id,
		Name:             r.Name,
		ShortDescription: r.ShortDescription,
		LongDescription:  r.LongDescription,
		Address:          r.Address,
		ImgUrl:           r.ImgUrl,
		Rating:           r.Rating,
		CommentCount:     r.CommentCount,
	}
}

func NewRestaurantArray(restArray []*entity.Restaurant) []*Restaurant {
	if restArray == nil {
		return []*Restaurant{}
	}
	restArrayDTO := make([]*Restaurant, len(restArray))
	for i, rest := range restArray {
		restArrayDTO[i] = NewRestaurant(rest)
	}
	return restArrayDTO
}
