package dto

import (
	"banners_oto/internal/entity"
	"banners_oto/internal/utils/constants"
)

type Promocode struct {
	Id   uint64 `json:"id"`
	Code string `json:"code"`
	Date string `json:"date"`
	Sale uint8  `json:"sale"`
	Type string `json:"type,omitempty"`
	Rest string `json:"rest,omitempty"`
	Sum  uint64 `json:"sum,omitempty"`
}

func NewPromocode(promo *entity.Promocode) *Promocode {
	return &Promocode{
		Id:   promo.Id,
		Code: promo.Code,
		Sale: promo.Sale,
		Date: promo.Date.Format(constants.Timestamptz),
		Type: promo.Type,
		Rest: promo.RestName,
		Sum:  promo.Sum,
	}
}

type Promo struct {
	Id     uint64 `json:"code_id"`
	Code   string `json:"code"`
	NewSum uint64 `json:"new_sum"`
}

type PromocodeArray struct {
	Payload []*Promocode `json:"payload" valid:"-"`
}

func NewPromocodeArray(arr []*entity.Promocode) []*Promocode {
	codeArray := make([]*Promocode, len(arr))
	for i, o := range arr {
		codeArray[i] = NewPromocode(o)
	}
	return codeArray
}
