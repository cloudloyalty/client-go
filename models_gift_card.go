package cloudloyalty_client

import (
	"time"

	"github.com/shopspring/decimal"
)

type SellGiftCardRequest struct {
	TxID     string        `json:"txid"`
	GiftCard GiftCardQuery `json:"giftCard"`
}

type SellGiftCardReply struct {
	GiftCard GiftCardReply `json:"giftCard"`
}

type GiftCardQuery struct {
	SKU  string `json:"sku"`
	Code string `json:"code"`
}

type GiftCardReply struct {
	Number     string          `json:"number,omitempty"`
	InitAmount decimal.Decimal `json:"initAmount"`
	ExpiredAt  *time.Time      `json:"expiredAt,omitempty"`
}
