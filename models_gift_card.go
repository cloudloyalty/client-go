package cloudloyalty_client

import (
	"time"

	"github.com/shopspring/decimal"
)

type GiftCardQuery struct {
	Code string `json:"code"`
}

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
	Code string `json:"code"` // deprecated
	Kind string `json:"kind"`
}

type GenerateGiftCardReply struct {
	GiftCard GiftCardReply `json:"giftCard"`
}

type GiftCardStatus string

const (
	GiftCardStatusInactive GiftCardStatus = "INACTIVE"
	GiftCardStatusActive   GiftCardStatus = "ACTIVE"
	GiftCardStatusRedeemed GiftCardStatus = "REDEEMED"
	GiftCardStatusExpired  GiftCardStatus = "EXPIRED"
	GiftCardStatusBlocked  GiftCardStatus = "BLOCKED"
)

type GiftCardReply struct {
	Number      string          `json:"number"`
	Code        string          `json:"code"`
	SKU         string          `json:"sku,omitempty"`
	InitAmount  decimal.Decimal `json:"initAmount"`
	Balance     decimal.Decimal `json:"balance"`
	Status      GiftCardStatus  `json:"status"`
	ActivatedAt *time.Time      `json:"activatedAt,omitempty"`
	BlockedAt   *time.Time      `json:"blockedAt,omitempty"`
	ValidFrom   *time.Time      `json:"validFrom,omitempty"`
	ValidUntil  *time.Time      `json:"validUntil,omitempty"`
}
