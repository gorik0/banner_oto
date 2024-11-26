package dto

import (
	"banners_oto/internal/entity"

	"github.com/asaskevich/govalidator"
)

type QuestionInput struct {
	Id     uint64 `json:"id"`
	Rating uint8  `json:"rating"`
}

func (d *QuestionInput) Validate() (bool, error) {
	return govalidator.ValidateStruct(d)
}

type QuestionInputArray struct {
	Payload []*QuestionInput `json:"payload" valid:"-"`
}

type Question struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	FocusId   string `json:"focus_id,omitempty"`
	ParamType string `json:"param_type"`
}

type QuestionArray struct {
	Payload []*Question `json:"payload" valid:"-"`
}

func (d *Question) Validate() (bool, error) {
	return govalidator.ValidateStruct(d)
}

func QuestionReturn(qArray []*entity.Question) []*Question {
	arr := []*Question{}
	for _, q := range qArray {
		qDTO := &Question{
			Id:        q.Id,
			Name:      q.Name,
			ParamType: q.ParamType,
			FocusId:   q.FocusId,
			Url:       q.Url,
		}
		arr = append(arr, qDTO)
	}
	return arr
}
