package payment

import (
	"banners_oto/config"
	"banners_oto/internal/delivery/metrics"
	"banners_oto/internal/entity/dto"
	"banners_oto/internal/usecase/order"
	"banners_oto/internal/usecase/session"
	cnst "banners_oto/internal/utils/constants"
	"banners_oto/internal/utils/functions"
	"banners_oto/internal/utils/myerrors"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/satori/uuid"
	"go.uber.org/zap"
)

type Payment struct {
	logger    *zap.Logger
	ucOrder   order.Usecase
	ucSession session.Usecase
	cfg       *config.ProjectConfiguration
	metrics   *metrics.Metrics
}

func NewPaymentDelivery(ucOrderProps order.Usecase, ucSessionProps session.Usecase, loggerProps *zap.Logger, metrics *metrics.Metrics) *Payment {
	return &Payment{
		logger:    loggerProps,
		ucOrder:   ucOrderProps,
		ucSession: ucSessionProps,
		cfg:       &config.Config,
		metrics:   metrics,
	}
}

func (d *Payment) OrderGetPayUrl(w http.ResponseWriter, r *http.Request) {
	requestId := functions.GetCtxRequestId(r)
	email := functions.GetCtxEmail(r)
	if email == "" {
		d.logger.Error(myerrors.AuthorizedEn.Error(), zap.String(cnst.RequestId, requestId))
		functions.ErrorResponse(w, myerrors.UnauthorizedRu, http.StatusUnauthorized)
		return
	}

	basket, err := d.ucOrder.GetBasket(r.Context(), email)
	if err != nil {
		d.logger.Error(err.Error(), zap.String(cnst.RequestId, requestId))
		if errors.Is(err, myerrors.SqlNoRowsUserRelation) {
			w, err = functions.FlashCookie(r, w, d.ucSession, &d.cfg.Redis, d.metrics)
			if err != nil {
				d.logger.Error(err.Error(), zap.String(cnst.RequestId, requestId))
				functions.ErrorResponse(w, myerrors.InternalServerRu, http.StatusInternalServerError)
				return
			}
			functions.ErrorResponse(w, myerrors.UnauthorizedRu, http.StatusUnauthorized)
			return
		}
		if errors.Is(err, myerrors.SqlNoRowsOrderRelation) {
			functions.ErrorResponse(w, myerrors.NoBasketRu, http.StatusNotFound)
			return
		}
		functions.ErrorResponse(w, myerrors.InternalServerRu, http.StatusInternalServerError)
		return
	}
	// we need to retrieve basket sum
	bodyRequestYooMoney := fmt.Sprintf(`{
		"amount": {
			"value": %d,
			"currency": "RUB"
		},
		"payment_method_data": {
		"type": "bank_card",
			"card": {
			"cardholder": "MR CARDHOLDER",
				"csc": "213",
				"expiry_month": "12",
				"expiry_year": "2024",
				"number": "5555555555554477"
		}
		},
		"capture": true,
		"confirmation": {
		"type": "redirect",
			"return_url": "https://resto-go.online
		},
		"description": "Заказ №%d"
	}`, basket.Sum, basket.Id)
	requestBody := bytes.NewBuffer([]byte(bodyRequestYooMoney))

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.yookassa.ru/v3/payments", requestBody)
	if err != nil {
		d.logger.Error(err.Error(), zap.String(cnst.RequestId, requestId))
		functions.ErrorResponse(w, myerrors.InternalServerRu, http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotence-Key", uuid.NewV4().String())
	req.SetBasicAuth(d.cfg.StoreId, d.cfg.Payment.SecretKey)

	resp, err := client.Do(req)
	if err != nil {
		d.logger.Error(err.Error(), zap.String(cnst.RequestId, requestId))
		functions.ErrorResponse(w, myerrors.InternalServerRu, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		d.logger.Error(err.Error(), zap.String(cnst.RequestId, requestId))
		functions.ErrorResponse(w, myerrors.InternalServerRu, http.StatusInternalServerError)
		return
	}
	// намеренно не используем easyjson, поскольку проще немного во времени проиграть
	var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		d.logger.Error(err.Error(), zap.String(cnst.RequestId, requestId))
		functions.ErrorResponse(w, myerrors.InternalServerRu, http.StatusInternalServerError)
		return
	}

	// receiving 'confirmation_url'
	confirmationURL := data["confirmation"].(map[string]interface{})["confirmation_url"].(string)
	functions.JsonResponse(w, &dto.ResponseUrlPay{Url: confirmationURL})
}
