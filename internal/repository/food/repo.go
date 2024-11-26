package food

import (
	"banners_oto/internal/delivery/metrics"
	"banners_oto/internal/entity"
	"banners_oto/internal/utils/alias"
	cnst "banners_oto/internal/utils/constants"
	merr "banners_oto/internal/utils/myerrors"
	"context"
	"database/sql"
	"errors"

	"time"
)

type Repo interface {
	GetByRestId(ctx context.Context, restId alias.RestId) ([]*entity.Food, error)
	GetById(ctx context.Context, foodId alias.FoodId) (*entity.Food, error)
}
type RepoLayer struct {
	db *sql.DB
	m  *metrics.Metrics
}

// GetById implements Repo.
func (repo *RepoLayer) GetByRestId(ctx context.Context, restId alias.RestId) ([]*entity.Food, error) {
	timeNow := time.Now()
	rows, err := repo.db.QueryContext(ctx,
		`SELECT c.name, f.id, f.name, restaurant_id, weight, price, img_url FROM food as f
    JOIN category as c ON f.category_id=c.id WHERE restaurant_id = $1 ORDER BY category_id`, uint64(restId))
	msRequestTimeout := time.Since(timeNow)
	repo.m.DatabaseDuration.WithLabelValues(cnst.SELECT).Observe(float64(msRequestTimeout.Milliseconds()))
	if err != nil {
		return nil, err
	}
	food := []*entity.Food{}
	for rows.Next() {
		item := entity.Food{}
		err = rows.Scan(&item.Category, &item.Id, &item.Name, &item.RestaurantId,
			&item.Weight, &item.Price, &item.ImgUrl)
		if err != nil {
			return nil, err
		}
		food = append(food, &item)
	}
	return food, nil
}

func (repo *RepoLayer) GetById(ctx context.Context, foodId alias.FoodId) (*entity.Food, error) {
	timeNow := time.Now()
	row := repo.db.QueryRowContext(ctx,
		`SELECT id, name, restaurant_id, category_id, weight, price, img_url
        FROM food WHERE id=$1`, uint64(foodId))
	msRequestTimeout := time.Since(timeNow)
	repo.m.DatabaseDuration.WithLabelValues(cnst.SELECT).Observe(float64(msRequestTimeout.Milliseconds()))
	var item entity.Food
	err := row.Scan(&item.Id, &item.Name, &item.RestaurantId,
		&item.Category, &item.Weight, &item.Price, &item.ImgUrl)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merr.SqlNoRowsFoodRelation
		}
		return nil, err
	}
	return &item, nil
}

func NewRepo(db *sql.DB, m *metrics.Metrics) Repo {
	return &RepoLayer{}
}
