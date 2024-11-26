package search

import (
	"banners_oto/internal/entity/dto"
	"banners_oto/internal/usecase/search"
	cnst "banners_oto/internal/utils/constants"
	"banners_oto/internal/utils/functions"
	"banners_oto/internal/utils/myerrors"
	"net/http"

	"go.uber.org/zap"
)

type Delivery struct {
	ucSearch search.Usecase
	logger   *zap.Logger
}

func NewDelivery(ucs search.Usecase, loggerProps *zap.Logger) *Delivery {
	return &Delivery{
		ucSearch: ucs,
		logger:   loggerProps,
	}
}

func (h *Delivery) Search(w http.ResponseWriter, r *http.Request) {
	requestId := functions.GetCtxRequestId(r)
	str := r.URL.Query().Get("search")
	var rests []*dto.RestaurantAndFood
	var err error
	if str != "" {
		rests, err = h.ucSearch.Search(r.Context(), str)
		if err != nil {
			h.logger.Error(err.Error(), zap.String(cnst.RequestId, requestId))
			functions.ErrorResponse(w, myerrors.InternalServerRu, http.StatusInternalServerError)
			return
		}
	} else {
		return
	}
	restaurantAndFoodArray := &dto.RestaurantAndFoodArray{Payload: rests}
	functions.JsonResponse(w, restaurantAndFoodArray)
}
