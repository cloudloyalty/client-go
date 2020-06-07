package cloudloyalty_client

import (
	"time"

	"github.com/shopspring/decimal"
)

type V2CalculatePurchaseRequest struct {
	CalculationQuery CalculationQuery `json:"calculationQuery"`
}

type SetPurchaseRequest struct {
	TxID             string           `json:"txid"`
	CalculationQuery CalculationQuery `json:"calculationQuery"`
}

type ConfirmTicketRequest struct {
	TxID       string  `json:"txid"`
	SessionNum *string `json:"sessionNum"`
	ReceiptNum *string `json:"receiptNum"`
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
	Client            *ClientQuery          `json:"client"`
	Shop              ShopQuery             `json:"shop"`
	Cashier           *CashierQuery         `json:"cashier"`
	ExecutedAt        *time.Time            `json:"executedAt"`
	OrderID           string                `json:"orderId"`
	Rows              []CalculationQueryRow `json:"rows"`
	ApplyBonuses      *IntOrAuto            `json:"applyBonuses"`
	CollectBonuses    *IntOrAuto            `json:"collectBonuses"`
	ApplyFactor       *decimal.Decimal      `json:"applyFactor"`
	CollectFactor     *decimal.Decimal      `json:"collectFactor"`
	Promocode         string                `json:"promocode"`
	DiscountRoundStep *float64              `json:"discountRoundStep"`
}

type CalculationQueryRow struct {
	ID               *string                    `json:"id"`
	Product          CalculationQueryRowProduct `json:"product"`
	Qty              float64                    `json:"qty"`
	AutoDiscount     decimal.Decimal            `json:"autoDiscount"`
	ManualDiscount   decimal.Decimal            `json:"manualDiscount"`
	NoApplyBonuses   bool                       `json:"noApplyBonuses"`
	NoCollectBonuses bool                       `json:"noCollectBonuses"`
	NoPromocode      bool                       `json:"noPromocode"`
	NoOffer          bool                       `json:"noOffer"`
	MaxDiscount      *decimal.Decimal           `json:"maxDiscount"`
}

type CalculationQueryRowProduct struct {
	ExternalID         string           `json:"externalId"`
	SKU                string           `json:"sku"`
	Title              string           `json:"title"`
	Category           string           `json:"category"`
	CategoryExternalID string           `json:"categoryExternalId"`
	BuyingPrice        *decimal.Decimal `json:"buyingPrice"`
	BlackPrice         decimal.Decimal  `json:"blackPrice"`
	RedPrice           *decimal.Decimal `json:"redPrice"`
	MinPrice           decimal.Decimal  `json:"minPrice"`
}

type V2CalculatePurchaseReply struct {
	CalculationResult CalculationResult `json:"calculationResult"`
}

type SetPurchaseReply struct {
	ClientBonuses     *ClientBonusesReply `json:"clientBonuses,omitempty"`
	CalculationResult CalculationResult   `json:"calculationResult"`
	Ticket            string              `json:"ticket"`
}

type CalculationResult struct {
	Rows      []CalculationResultRow      `json:"rows"`
	Summary   CalculationResultSummary    `json:"summary"`
	Bonuses   *CalculationResultBonuses   `json:"bonuses,omitempty"`
	Promocode *CalculationResultPromocode `json:"promocode,omitempty"`
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
	ID      int                     `json:"-"`
	Error   *CalculationResultError `json:"error,omitempty"`
}

type CalculationResultError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Hint        string `json:"hint,omitempty"`
}
