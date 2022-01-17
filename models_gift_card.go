package cloudloyalty_client

import (
	"time"

	"github.com/shopspring/decimal"
)

type ActivateGiftCardRequest struct {
	TxID     string        `json:"txid"`
	GiftCard GiftCardQuery `json:"giftCard"`
}

type ActivateGiftCardReply struct {
	GiftCard GiftCardReply `json:"giftCard"`
}

// DiscardGiftCardRequest request model for discard-gift-card request
type DiscardGiftCardRequest struct {
	TxID     string        `json:"txid"`
	GiftCard GiftCardQuery `json:"giftCard"`
}

// DiscardGiftCardReply response model for discard-gift-card request
type DiscardGiftCardReply struct {
}

type GenerateGiftCardRequest struct {
	Code string `json:"code"`
}

type GenerateGiftCardReply struct {
	GiftCard GiftCardReply `json:"giftCard"`
}

type GiftCardQuery struct {
	SKU  string `json:"sku"`
	Code string `json:"code"`
}

type GiftCardReply struct {
	Number     string          `json:"number"`
	Code       string          `json:"code"`
	InitAmount decimal.Decimal `json:"initAmount"`
	ValidUntil *time.Time      `json:"validUntil,omitempty"`
}
