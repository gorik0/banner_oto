package food

import (
	"banners_oto/internal/entity"
	"banners_oto/internal/utils/alias"
	"errors"

	foodRepo "banners_oto/internal/repository/food"
	"context"
)

type Usecase interface {
	GetByRestId(ctx context.Context, restId alias.RestId) ([]*entity.Category, error)
	GetById(ctx context.Context, foodId alias.FoodId) (*entity.Food, error)
}

type UsecaseLayer struct {
	repoFood foodRepo.Repo
}

// GetById implements Usecase.
func (u UsecaseLayer) GetById(ctx context.Context, foodId alias.FoodId) (*entity.Food, error) {

	dishes, err := u.repoFood.GetById(ctx, foodId)
	return dishes, err
}

// GetByRestId implements Usecase.
func (u UsecaseLayer) GetByRestId(ctx context.Context, restId alias.RestId) ([]*entity.Category, error) {
	dishes, err := u.repoFood.GetByRestId(ctx, restId)
	if err != nil {

		return nil, err
	}
	if dishes == nil {
		return nil, errors.New("Empty dishes ... ")
	}

	categories := []*entity.Category{}

	id := 0
	var category entity.Category
	for i, dish := range dishes {
		id++
		if i != 0 && dishes[i-1].Category == dish.Category {
			category.Food = append(category.Food, dish)
		} else {

			category = entity.Category{
				Id:   alias.CategoryId(id),
				Name: dish.Category,
				Food: []*entity.Food{dish},
			}
			categories = append(categories, &category)
		}

	}
	return categories, nil

}

func NewUsecaseLayer(repo foodRepo.Repo) Usecase {
	return UsecaseLayer{
		repoFood: repo,
	}
}
