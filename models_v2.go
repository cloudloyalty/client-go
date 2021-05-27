package cloudloyalty_client

import (
	"time"

	"github.com/shopspring/decimal"
)

type V2CalculatePurchaseRequest struct {
	OrderID          string           `json:"orderId"`
	CalculationQuery CalculationQuery `json:"calculationQuery"`
}

type SetPurchaseRequest struct {
	TxID             string           `json:"txid"`
	CalculationQuery CalculationQuery `json:"calculationQuery"`
}

type ConfirmTicketRequest struct {
	TxID       string  `json:"txid"`
	SessionNum *string `json:"sessionNum,omitempty"`
	ReceiptNum *string `json:"receiptNum,omitempty"`
	Ticket     string  `json:"ticket"`
}

type ConfirmTicketReply struct {
	ClientBonuses *ClientBonusesReply `json:"-"`
}

type DiscardTicketRequest struct {
	TxID   string `json:"txid"`
	Ticket string `json:"ticket"`
}

type DiscardTicketReply struct {
}

type CalculationQuery struct {
	Client            *ClientQuery          `json:"client,omitempty"`
	Shop              ShopQuery             `json:"shop"`
	Cashier           *CashierQuery         `json:"cashier,omitempty"`
	ExecutedAt        *time.Time            `json:"executedAt,omitempty"`
	OrderID           string                `json:"orderId,omitempty"`
	Rows              []CalculationQueryRow `json:"rows"`
	ApplyBonuses      *IntOrAuto            `json:"applyBonuses,omitempty"`
	CollectBonuses    *IntOrAuto            `json:"collectBonuses,omitempty"`
	ApplyFactor       *decimal.Decimal      `json:"applyFactor,omitempty"`
	CollectFactor     *decimal.Decimal      `json:"collectFactor,omitempty"`
	Promocode         string                `json:"promocode,omitempty"`
	DiscountRoundStep *float64              `json:"discountRoundStep,omitempty"`
}

type CalculationQueryRow struct {
	ID                *string                    `json:"id,omitempty"`
	Product           CalculationQueryRowProduct `json:"product"`
	Qty               float64                    `json:"qty"`
	AutoDiscount      decimal.Decimal            `json:"autoDiscount,omitempty"`
	ManualDiscount    decimal.Decimal            `json:"manualDiscount,omitempty"`
	NoApplyBonuses    bool                       `json:"noApplyBonuses,omitempty"`
	NoCollectBonuses  bool                       `json:"noCollectBonuses,omitempty"`
	NoPromocode       bool                       `json:"noPromocode,omitempty"`
	NoOffer           bool                       `json:"noOffer,omitempty"`
	MaxDiscount       *decimal.Decimal           `json:"maxDiscount,omitempty"`
	DiscountRoundStep *float64                   `json:"discountRoundStep,omitempty"`
}

type CalculationQueryRowProduct struct {
	ExternalID         string           `json:"externalId,omitempty"`
	SKU                string           `json:"sku"`
	Title              string           `json:"title"`
	Category           string           `json:"category,omitempty"`
	CategoryExternalID string           `json:"categoryExternalId,omitempty"`
	BuyingPrice        *decimal.Decimal `json:"buyingPrice,omitempty"`
	BlackPrice         decimal.Decimal  `json:"blackPrice"`
	RedPrice           *decimal.Decimal `json:"redPrice,omitempty"`
	MinPrice           decimal.Decimal  `json:"minPrice,omitempty"`
}

type V2CalculatePurchaseReply struct {
	CalculationResult CalculationResult `json:"calculationResult"`
}

type SetPurchaseReply struct {
	ClientBonuses     *ClientBonusesReply `json:"clientBonuses,omitempty"`
	CalculationResult CalculationResult   `json:"calculationResult"`
	Ticket            string              `json:"ticket"`
	ReceiptInfo       []ReceiptInfoLine   `json:"receiptInfo,omitempty"`
}

type CalculationResult struct {
	Rows      []CalculationResultRow      `json:"rows"`
	Summary   CalculationResultSummary    `json:"summary"`
	Bonuses   *CalculationResultBonuses   `json:"bonuses,omitempty"`
	Promocode *CalculationResultPromocode `json:"promocode,omitempty"`
	GiftCard  *CalculationResultPromocode `json:"giftCard,omitempty"`
}

type CalculationResultRow struct {
	ID            *string                      `json:"id,omitempty"`
	TotalDiscount decimal.Decimal              `json:"totalDiscount"`
	Discounts     CalculationResultDiscounts   `json:"discounts"`
	Bonuses       *CalculationResultRowBonuses `json:"bonuses,omitempty"`
	Offers        []CalculationResultRowOffer  `json:"offers,omitempty"`
}

type CalculationResultDiscounts struct {
	Auto      decimal.Decimal `json:"auto"`
	Manual    decimal.Decimal `json:"manual"`
	Bonuses   decimal.Decimal `json:"bonuses"`
	Promocode decimal.Decimal `json:"promocode"`
	Offer     decimal.Decimal `json:"offer"`
	Rounding  decimal.Decimal `json:"rounding"`
}

type CalculationResultRowBonuses struct {
	Applied   int `json:"applied"`
	Collected int `json:"collected"`
}

type CalculationResultRowOffer struct {
	ID      int             `json:"id"`
	Code    string          `json:"code"`
	Name    string          `json:"name"`
	Bonuses int             `json:"bonuses,omitempty"`
	Amount  decimal.Decimal `json:"amount,omitempty"`
}

type CalculationResultSummary struct {
	TotalDiscount decimal.Decimal            `json:"totalDiscount"`
	Discounts     CalculationResultDiscounts `json:"discounts"`
}

type CalculationResultBonuses struct {
	Applied    int                     `json:"applied"`
	Collected  int                     `json:"collected"`
	MaxToApply int                     `json:"maxToApply"`
	Error      *CalculationResultError `json:"error,omitempty"`
}

type CalculationResultPromocode struct {
	Applied bool                    `json:"applied"`
	Error   *CalculationResultError `json:"error,omitempty"`
}

type CalculationResultError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Hint        string `json:"hint,omitempty"`
}

type ReceiptInfoLine struct {
	Type    string                  `json:"type"`
	Text    *ReceiptInfoLineText    `json:"text,omitempty"`
	Table   *ReceiptInfoLineTable   `json:"table,omitempty"`
	Barcode *ReceiptInfoLineBarcode `json:"barcode,omitempty"`
}

type ReceiptInfoLineText struct {
	Align string `json:"align,omitempty"`
	Value string `json:"value"`
}

type ReceiptInfoLineTable struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

type ReceiptInfoLineBarcode struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type V2SetOrderRequest struct {
	OrderID          string           `json:"orderId"`
	CalculationQuery CalculationQuery `json:"calculationQuery"`
}

type V2SetOrderReply struct {
	ClientBonuses     *ClientBonusesReply `json:"clientBonuses,omitempty"`
	CalculationResult CalculationResult   `json:"calculationResult"`
}
