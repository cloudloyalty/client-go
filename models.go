package cloudloyalty_client

import (
	"time"
)

const (
	LoyaltyActionNone         = "none"
	LoyaltyActionApply        = "apply"
	LoyaltyActionCollect      = "collect"
	LoyaltyActionApplyCollect = "apply-collect"
)

const (
	GenderUnknown = 0
	GenderMale    = 1
	GenderFemale  = 2
)

type ErrorReply struct {
	ErrorCode   int    `json:"errorCode"`
	Description string `json:"description"`
	Hint        string `json:"hint,omitempty"`
}

// get-balance

type GetBalanceQuery struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Card        string `json:"card,omitempty"`
	ExternalID  string `json:"externalId,omitempty"`
}

type GetBalanceReply struct {
	Client      GetBalanceReplyClient  `json:"client"`
	Bonuses     []GetBalanceReplyBonus `json:"bonuses"`
	WalletsLink string                 `json:"walletsLink,omitempty"`
}

type GetBalanceReplyClient struct {
	PhoneNumber       string            `json:"phoneNumber,omitempty"`
	Card              int64             `json:"card,omitempty"`
	CardString        string            `json:"cardString,omitempty"`
	ExternalID        string            `json:"externalId,omitempty"`
	Surname           string            `json:"surname,omitempty"`
	Name              string            `json:"name,omitempty"`
	PatronymicName    string            `json:"patronymicName,omitempty"`
	FullName          string            `json:"fullName,omitempty"`
	Gender            int               `json:"gender"`
	Birthdate         string            `json:"birthdate,omitempty"`
	Email             string            `json:"email,omitempty"`
	Level             int               `json:"level"`
	IsEmailSubscribed bool              `json:"isEmailSubscribed"`
	IsPhoneSubscribed bool              `json:"isPhoneSubscribed"`
	ExtraFields       map[string]string `json:"extraFields,omitempty"`
	Bonuses           int               `json:"bonuses"`
	PendingBonuses    int               `json:"pendingBonuses"`
	Children          []Child           `json:"children,omitempty"`
}

type GetBalanceReplyBonus struct {
	ExpireAt string `json:"expireAt,omitempty"`
	Amount   int    `json:"amount"`
}

type Child struct {
	Name      string `json:"name,omitempty"`
	Birthdate string `json:"birthdate,omitempty"`
	Gender    int    `json:"gender"`
}

// new-client

type NewClientQuery struct {
	Client  NewClientClient `json:"client"`
	Shop    *ShopQuery      `json:"shop"`
	Cashier *CashierQuery   `json:"cashier"`
}

type ShopQuery struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type CashierQuery struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type NewClientClient struct {
	PhoneNumber       string            `json:"phoneNumber,omitempty"`
	Card              string            `json:"card,omitempty"`
	ExternalID        string            `json:"externalId,omitempty"`
	Surname           string            `json:"surname,omitempty"`
	Name              string            `json:"name,omitempty"`
	PatronymicName    string            `json:"patronymicName,omitempty"`
	FullName          string            `json:"fullName,omitempty"`
	Gender            int               `json:"gender,omitempty"`
	Birthdate         string            `json:"birthdate,omitempty"`
	Email             string            `json:"email,omitempty"`
	Level             int               `json:"level,omitempty"`
	IsEmailSubscribed *bool             `json:"isEmailSubscribed,omitempty"`
	IsPhoneSubscribed *bool             `json:"isPhoneSubscribed,omitempty"`
	ExtraFields       map[string]string `json:"extraFields,omitempty"`
	Children          []Child           `json:"children,omitempty"`
}

type NewClientReply struct {
	Client      GetBalanceReplyClient  `json:"client"`
	Bonuses     []GetBalanceReplyBonus `json:"bonuses"`
	WalletsLink string                 `json:"walletsLink,omitempty"`
}

// update-client

type UpdateClientQuery struct {
	PhoneNumber string          `json:"phoneNumber,omitempty"`
	Card        string          `json:"card,omitempty"`
	ExternalID  string          `json:"externalId,omitempty"`
	Client      NewClientClient `json:"client"`
}

type UpdateClientReply struct {
	Client      GetBalanceReplyClient  `json:"client"`
	Bonuses     []GetBalanceReplyBonus `json:"bonuses"`
	WalletsLink string                 `json:"walletsLink,omitempty"`
}

// calculate-purchase

type CalculatePurchaseQuery struct {
	Calculate CalculatePurchaseQueryCalculate `json:"calculate"`
}

type CalculatePurchaseQueryCalculate struct {
	PhoneNumber       string          `json:"phoneNumber,omitempty"`
	Card              string          `json:"card,omitempty"`
	ExternalID        string          `json:"externalId,omitempty"`
	IsAnonymousClient bool            `json:"isAnonymousClient,omitempty"`
	LoyaltyAction     string          `json:"loyaltyAction,omitempty"`
	TotalAmount       float64         `json:"totalAmount"`
	ApplyingAmount    *float64        `json:"applyingAmount,omitempty"`
	CollectingAmount  *float64        `json:"collectingAmount,omitempty"`
	ApplyBonuses      *int            `json:"applyBonuses,omitempty"`
	CollectBonuses    *int            `json:"collectBonuses,omitempty"`
	Promocode         string          `json:"promocode,omitempty"`
	Units             map[string]Unit `json:"units,omitempty"`
	OrderID           string          `json:"orderId,omitempty"`
	ExecutedAt        *time.Time      `json:"executedAt,omitempty"`
	DoApplyBonuses    *bool           `json:"doApplyBonuses,omitempty"`   // obsolete
	DoCollectBonuses  *bool           `json:"doCollectBonuses,omitempty"` // obsolete
}

type Unit struct {
	ExternalID         string   `json:"externalId,omitempty"`
	Sku                string   `json:"sku"`
	ItemTitle          string   `json:"itemTitle"`
	ItemCount          *float64 `json:"itemCount,omitempty"`
	BuyingPrice        *float64 `json:"buyingPrice,omitempty"`
	Price              float64  `json:"price"`
	Category           string   `json:"category,omitempty"`
	CategoryExternalID string   `json:"categoryExternalID,omitempty"`
	Amount             *float64 `json:"amount,omitempty"`
	MinPrice           float64  `json:"minPrice,omitempty"`
}

type CalculatePurchaseReply struct {
	Calculation     CalculatePurchaseReplyCalculate `json:"calculation"`
	CalculatedUnits map[string]CalculatedUnit       `json:"calculatedUnits"`
	ClientBonuses   ClientBonusesReply              `json:"clientBonuses"`
}

type CalculatedUnit struct {
	IsPromocodeApplicable bool    `json:"isPromocodeApplicable"`
	PromocodeDiscount     float64 `json:"promocodeDiscount"`
	BonusesDiscount       float64 `json:"bonusesDiscount"`
	TotalDiscount         float64 `json:"totalDiscount"`
	OriginalPrice         float64 `json:"originalPrice"`
	OriginalAmount        float64 `json:"originalAmount"`
	CalculatedAmount      float64 `json:"calculatedAmount"`
}

type CalculatePurchaseReplyCalculate struct {
	AppliedBonuses     int     `json:"appliedBonuses"`
	CollectedBonuses   int     `json:"collectedBonuses"`
	AppliableBonuses   int     `json:"appliableBonuses"`   // obsolete
	CollectableBonuses int     `json:"collectableBonuses"` // obsolete
	MaxBonuses         int     `json:"maxBonuses"`
	PromocodeDiscount  float64 `json:"promocodeDiscount"`
	BonusesDiscount    float64 `json:"bonusesDiscount"`
	TotalDiscount      float64 `json:"totalDiscount"`
	TotalAmount        float64 `json:"totalAmount"` // obsolete
	RemainingAmount    float64 `json:"remainingAmount"`
	PhoneNumber        string  `json:"phoneNumber,omitempty"` // obsolete`
}

// apply-purchase

type ApplyPurchaseQuery struct {
	Transaction ApplyPurchaseTransaction `json:"transaction"`
}

type ApplyPurchaseTransaction struct {
	PhoneNumber                string                             `json:"phoneNumber,omitempty"`
	Card                       string                             `json:"card,omitempty"`
	ExternalID                 string                             `json:"externalId,omitempty"`
	IsAnonymousClient          bool                               `json:"isAnonymousClient,omitempty"`
	ID                         string                             `json:"id"`
	ExecutedAt                 string                             `json:"executedAt"`
	ReceiptID                  IntOrString                        `json:"receiptId,omitempty"`
	SessionID                  IntOrString                        `json:"sessionId,omitempty"`
	ShopCode                   string                             `json:"shopCode"`
	ShopName                   string                             `json:"shopName"`
	Cashier                    string                             `json:"cashier,omitempty"`
	CashierID                  string                             `json:"cashierId,omitempty"`
	LoyaltyAction              string                             `json:"loyaltyAction"`
	TotalAmount                float64                            `json:"totalAmount"`
	ApplyingAmount             *float64                           `json:"applyingAmount,omitempty"`
	CollectingAmount           *float64                           `json:"collectingAmount,omitempty"`
	ApplyBonuses               int                                `json:"applyBonuses"`
	CollectBonuses             *int                               `json:"collectBonuses,omitempty"`
	Promocode                  *ApplyPurchaseTransactionPromocode `json:"promocode,omitempty"`
	Items                      []Item                             `json:"items,omitempty"`
	IsConfirmationCodeRequired bool                               `json:"isConfirmationCodeRequired,omitempty"`
	ConfirmationCode           string                             `json:"confirmationCode,omitempty"`
}

type ApplyPurchaseTransactionPromocode struct {
	Promocode         string  `json:"promocode"`
	PromocodeDiscount float64 `json:"promocodeDiscount"`
}

type ApplyPurchaseReply struct {
	Confirmation ApplyPurchaseReplyConfirmation `json:"confirmation"`
}

type ApplyPurchaseReplyConfirmation struct {
	PhoneNumber string                    `json:"phoneNumber"`
	PurchaseID  string                    `json:"purchaseId"`
	TotalAmount float64                   `json:"totalAmount"`
	PaidAmount  float64                   `json:"paidAmount"`
	Bonuses     ApplyPurchaseReplyBonuses `json:"bonuses"`
}

type ApplyPurchaseReplyBonuses struct {
	Applied   int `json:"applied"`
	Collected int `json:"collected"`
	Pending   int `json:"pending"`
	Available int `json:"available"`
}

// calculate-return

type CalculateReturnQuery struct {
	Calculate CalculateReturnQueryCalculate `json:"calculate"`
}

type CalculateReturnQueryCalculate struct {
	PhoneNumber  string  `json:"phoneNumber,omitempty"`
	Card         string  `json:"card,omitempty"`
	ExternalID   string  `json:"externalId,omitempty"`
	PurchaseID   string  `json:"purchaseId"`
	RefundAmount float64 `json:"refundAmount"`
}

type CalculateReturnReply struct {
	Calculation CalculateReturnReplyCalculate `json:"calculation"`
}

type CalculateReturnReplyCalculate struct {
	RecoveredBonuses int `json:"recoveredBonuses"`
	CancelledBonuses int `json:"cancelledBonuses"`
}

// apply-return

type ApplyReturnQuery struct {
	Transaction ApplyReturnTransaction `json:"transaction"`
}

type ApplyReturnTransaction struct {
	PhoneNumber       string  `json:"phoneNumber,omitempty"`
	Card              string  `json:"card,omitempty"`
	ExternalID        string  `json:"externalId,omitempty"`
	IsAnonymousClient bool    `json:"isAnonymousClient,omitempty"`
	ID                string  `json:"id"`
	ExecutedAt        string  `json:"executedAt"`
	PurchaseID        string  `json:"purchaseId"`
	ReceiptID         int     `json:"receiptId,omitempty"`
	SessionID         int     `json:"sessionId,omitempty"`
	ShopCode          string  `json:"shopCode"`
	ShopName          string  `json:"shopName"`
	Cashier           string  `json:"cashier,omitempty"`
	CashierID         string  `json:"cashierId,omitempty"`
	RefundAmount      float64 `json:"refundAmount"`
	Items             []Item  `json:"items,omitempty"`
}

type ApplyReturnReply struct {
	Confirmation ApplyReturnReplyConfirmation `json:"confirmation"`
}

type ApplyReturnReplyConfirmation struct {
	PhoneNumber      string  `json:"phoneNumber"`
	RefundID         string  `json:"refundId"`
	RefundAmount     float64 `json:"refundAmount"`
	RecoveredBonuses int     `json:"recoveredBonuses"`
	CancelledBonuses int     `json:"cancelledBonuses"`
}

type Item struct {
	ExternalID         string   `json:"externalId,omitempty"`
	SKU                string   `json:"sku"`
	ItemTitle          string   `json:"itemTitle"`
	ItemCount          float64  `json:"itemCount"`
	BuyingPrice        *float64 `json:"buyingPrice,omitempty"`
	Price              float64  `json:"price"`
	Amount             *float64 `json:"amount,omitempty"`
	Category           string   `json:"category,omitempty"`
	CategoryExternalID string   `json:"categoryExternalID,omitempty"`
	MinPrice           float64  `json:"minPrice,omitempty"`
}

// adjust-balance

type AdjustBalanceQuery struct {
	Client            ClientQuery            `json:"client"`
	BalanceAdjustment BalanceAdjustmentQuery `json:"balanceAdjustment"`
}

type ClientQuery struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Card        string `json:"card,omitempty"`
	ExternalID  string `json:"externalId,omitempty"`
}

type BalanceAdjustmentQuery struct {
	AmountDelta          int    `json:"amountDelta"`
	ExpirationPeriodDays int    `json:"expirationPeriodDays,omitempty"`
	Comment              string `json:"comment,omitempty"`
	Notify               bool   `json:"notify,omitempty"`
}

type AdjustBalanceReply struct {
	ClientBonuses ClientBonusesReply `json:"clientBonuses"`
}

type ClientBonusesReply struct {
	Available int `json:"available"`
	Pending   int `json:"pending"`
	Reserved  int `json:"reserved"`
	Total     int `json:"total,omitempty"`
}

// set-order

type SetOrderQuery struct {
	Client ClientQuery `json:"client"`
	Order  OrderQuery  `json:"order"`
}

type OrderQuery struct {
	ID          string           `json:"id"`
	ExecutedAt  *time.Time       `json:"executedAt,omitempty"`
	ShopCode    string           `json:"shopCode"`
	ShopName    string           `json:"shopName"`
	TotalAmount float64          `json:"totalAmount"`
	Loyalty     *SetOrderLoyalty `json:"loyalty,omitempty"`
	Promocode   string           `json:"promocode,omitempty"`
	Items       []Item           `json:"items,omitempty"`
}

type SetOrderLoyalty struct {
	Action                     string       `json:"action"`
	ApplyingAmount             *float64     `json:"applyingAmount,omitempty"`
	CollectingAmount           *float64     `json:"collectingAmount,omitempty"`
	ApplyBonuses               *IntOrString `json:"applyBonuses"`
	CollectBonuses             *int         `json:"collectBonuses,omitempty"`
	IsConfirmationCodeRequired bool         `json:"isConfirmationCodeRequired,omitempty"`
	ConfirmationCode           string       `json:"confirmationCode,omitempty"`
}

type SetOrderReply struct {
	OperationResult SetOrderOperationResult `json:"operationResult"`
	ClientBonuses   ClientBonusesReply      `json:"clientBonuses"`
}

type SetOrderOperationResult struct {
	AppliedBonuses   int     `json:"appliedBonuses"`
	CollectedBonuses int     `json:"collectedBonuses"`
	RemainingAmount  float64 `json:"remainingAmount"`
}

// confirm-order

type ConfirmOrderQuery struct {
	OrderID    string     `json:"orderId"`
	ExecutedAt *time.Time `json:"executedAt,omitempty"`
}

type ConfirmOrderReply struct {
	OperationResult SetOrderOperationResult `json:"operationResult"`
	ClientBonuses   ClientBonusesReply      `json:"clientBonuses"`
}

// cancel-order

type CancelOrderQuery struct {
	OrderID    string     `json:"orderId"`
	ExecutedAt *time.Time `json:"executedAt,omitempty"`
}

type CancelOrderReply struct {
	OperationResult SetOrderOperationResult `json:"operationResult"`
	ClientBonuses   ClientBonusesReply      `json:"clientBonuses"`
}

// send-confirmation-code
type SendConfirmationCodeQuery struct {
	PhoneNumber       string `json:"phoneNumber,omitempty"`
	Card              string `json:"card,omitempty"`
	ExternalID        string `json:"externalId,omitempty"`
	IsAnonymousClient bool   `json:"isAnonymousClient"`
	To                string `json:"to,omitempty"`
}

type SendConfirmationCodeReply struct {
	Code      string    `json:"code"`
	MsgID     string    `json:"msgid"`
	ExpiresAt time.Time `json:"expiresAt"`
}
