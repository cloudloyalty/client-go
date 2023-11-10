package cloudloyalty_client

import "github.com/shopspring/decimal"

type CalculateProductsRequest struct {
	Client            *ClientQuery                   `json:"client,omitempty"`
	Shop              ShopQuery                      `json:"shop"`
	Products          []CalculateProductsRequestItem `json:"products"`
	DiscountRoundStep *float64                       `json:"discountRoundStep,omitempty"`
}

type CalculateProductsRequestItem struct {
	ID               *string                    `json:"id,omitempty"`
	Product          CalculationQueryRowProduct `json:"product"`
	Qty              *float64                   `json:"qty,omitempty"`
	ExternalDiscount decimal.Decimal            `json:"externalDiscount,omitempty"`
	NoCollectBonuses bool                       `json:"noCollectBonuses,omitempty"`
	NoOffer          bool                       `json:"noOffer,omitempty"`
	MaxDiscount      *decimal.Decimal           `json:"maxDiscount,omitempty"`
}

type CalculateProductsReply struct {
	Products []CalculateProductsResultItem `json:"products"`
}

type CalculateProductsResultItem struct {
	ID            *string                              `json:"id,omitempty"`
	TotalDiscount decimal.Decimal                      `json:"totalDiscount"`
	Discounts     CalculateProductsResultItemDiscounts `json:"discounts"`
	Bonuses       CalculateProductsResultItemBonuses   `json:"bonuses"`
	Offers        []CalculationResultRowOffer          `json:"offers,omitempty"`
}

type CalculateProductsResultItemDiscounts struct {
	External decimal.Decimal `json:"external"`
	Offer    decimal.Decimal `json:"offer"`
	Rounding decimal.Decimal `json:"rounding"`
}

type CalculateProductsResultItemBonuses struct {
	Collected int `json:"collected"`
}
