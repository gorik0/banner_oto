package dto

import "banners_oto/internal/utils/alias"

type ResponseDetail struct {
	Detail string `json:"detail" valid:"-"`
}

type ResponseUrlPay struct {
	Url string `json:"url" valid:"-"`
}

type PayedOrderInfo struct {
	Id     alias.OrderId `json:"id"`
	Status string        `json:"status"`
}
