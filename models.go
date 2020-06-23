package cloudloyalty_client

import (
	"time"

	"github.com/shopspring/decimal"
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
	PhoneNumber       string      `json:"phoneNumber,omitempty"`
	Card              int64       `json:"card,omitempty"`
	CardString        string      `json:"cardString,omitempty"`
	ExternalID        string      `json:"externalId,omitempty"`
	Surname           string      `json:"surname,omitempty"`
	Name              string      `json:"name,omitempty"`
	PatronymicName    string      `json:"patronymicName,omitempty"`
	FullName          string      `json:"fullName,omitempty"`
	Gender            int         `json:"gender"`
	Birthdate         string      `json:"birthdate,omitempty"`
	Email             string      `json:"email,omitempty"`
	Level             int         `json:"level"`
	IsEmailSubscribed bool        `json:"isEmailSubscribed"`
	IsPhoneSubscribed bool        `json:"isPhoneSubscribed"`
	ExtraFields       ExtraFields `json:"extraFields,omitempty"`
	Bonuses           int         `json:"bonuses"`
	PendingBonuses    int         `json:"pendingBonuses"`
	Children          []Child     `json:"children,omitempty"`
}

type GetBalanceReplyBonus struct {
	ExpireAt *time.Time `json:"expireAt,omitempty"`
	Amount   int        `json:"amount"`
}

type Child struct {
	Name      string `json:"name,omitempty"`
	Birthdate string `json:"birthdate,omitempty"`
	Gender    int    `json:"gender"`
}

// new-client

type NewClientQuery struct {
	Client  NewClientClient `json:"client"`
	Shop    *ShopQuery      `json:"shop,omitempty"`
	Cashier *CashierQuery   `json:"cashier,omitempty"`
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
	PhoneNumber       string      `json:"phoneNumber,omitempty"`
	Card              string      `json:"card,omitempty"`
	ExternalID        string      `json:"externalId,omitempty"`
	Surname           string      `json:"surname,omitempty"`
	Name              string      `json:"name,omitempty"`
	PatronymicName    string      `json:"patronymicName,omitempty"`
	FullName          string      `json:"fullName,omitempty"`
	Gender            int         `json:"gender,omitempty"`
	Birthdate         string      `json:"birthdate,omitempty"`
	Email             string      `json:"email,omitempty"`
	Level             int         `json:"level,omitempty"`
	IsEmailSubscribed *bool       `json:"isEmailSubscribed,omitempty"`
	IsPhoneSubscribed *bool       `json:"isPhoneSubscribed,omitempty"`
	ExtraFields       ExtraFields `json:"extraFields,omitempty"`
	Children          []Child     `json:"children,omitempty"`

	City   string `json:"city,omitempty"`   // deprecated
	Street string `json:"street,omitempty"` // deprecated
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
	PhoneNumber       string           `json:"phoneNumber,omitempty"`
	Card              string           `json:"card,omitempty"`
	ExternalID        string           `json:"externalId,omitempty"`
	IsAnonymousClient bool             `json:"isAnonymousClient,omitempty"`
	LoyaltyAction     string           `json:"loyaltyAction,omitempty"`
	TotalAmount       decimal.Decimal  `json:"totalAmount"`
	ApplyingAmount    *decimal.Decimal `json:"applyingAmount,omitempty"`
	CollectingAmount  *decimal.Decimal `json:"collectingAmount,omitempty"`
	ApplyBonuses      *int             `json:"applyBonuses,omitempty"`
	CollectBonuses    *int             `json:"collectBonuses,omitempty"`
	Promocode         string           `json:"promocode,omitempty"`
	Units             map[string]Unit  `json:"units,omitempty"`
	OrderID           string           `json:"orderId,omitempty"`
	ExecutedAt        *time.Time       `json:"executedAt,omitempty"`
	DoApplyBonuses    *bool            `json:"doApplyBonuses,omitempty"`   // obsolete
	DoCollectBonuses  *bool            `json:"doCollectBonuses,omitempty"` // obsolete
}

type Unit struct {
	ExternalID         string           `json:"externalId,omitempty"`
	Sku                string           `json:"sku"`
	ItemTitle          string           `json:"itemTitle"`
	ItemCount          *float64         `json:"itemCount,omitempty"`
	BuyingPrice        *decimal.Decimal `json:"buyingPrice,omitempty"`
	Price              decimal.Decimal  `json:"price"`
	Category           string           `json:"category,omitempty"`
	CategoryExternalID string           `json:"categoryExternalID,omitempty"`
	Amount             *decimal.Decimal `json:"amount,omitempty"`
	MinPrice           decimal.Decimal  `json:"minPrice,omitempty"`
}

type CalculatePurchaseReply struct {
	Calculation     CalculatePurchaseReplyCalculate `json:"calculation"`
	CalculatedUnits map[string]CalculatedUnit       `json:"calculatedUnits"`
	ClientBonuses   ClientBonusesReply              `json:"clientBonuses"`
}

type CalculatedUnit struct {
	IsPromocodeApplicable bool            `json:"isPromocodeApplicable"`
	PromocodeDiscount     decimal.Decimal `json:"promocodeDiscount"`
	BonusesDiscount       decimal.Decimal `json:"bonusesDiscount"`
	TotalDiscount         decimal.Decimal `json:"totalDiscount"`
	OriginalPrice         decimal.Decimal `json:"originalPrice"`
	OriginalAmount        decimal.Decimal `json:"originalAmount"`
	CalculatedAmount      decimal.Decimal `json:"calculatedAmount"`
}

type CalculatePurchaseReplyCalculate struct {
	AppliedBonuses     int             `json:"appliedBonuses"`
	CollectedBonuses   int             `json:"collectedBonuses"`
	AppliableBonuses   int             `json:"appliableBonuses"`   // obsolete
	CollectableBonuses int             `json:"collectableBonuses"` // obsolete
	MaxBonuses         int             `json:"maxBonuses"`
	PromocodeDiscount  decimal.Decimal `json:"promocodeDiscount"`
	BonusesDiscount    decimal.Decimal `json:"bonusesDiscount"`
	TotalDiscount      decimal.Decimal `json:"totalDiscount"`
	TotalAmount        decimal.Decimal `json:"totalAmount"` // obsolete
	RemainingAmount    decimal.Decimal `json:"remainingAmount"`
	PhoneNumber        string          `json:"phoneNumber,omitempty"` // obsolete`
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
	ExecutedAt                 time.Time                          `json:"executedAt"`
	ReceiptID                  IntAsIntOrString                   `json:"receiptId,omitempty"`
	SessionID                  IntAsIntOrString                   `json:"sessionId,omitempty"`
	ShopCode                   string                             `json:"shopCode"`
	ShopName                   string                             `json:"shopName"`
	Cashier                    string                             `json:"cashier,omitempty"`
	CashierID                  string                             `json:"cashierId,omitempty"`
	LoyaltyAction              string                             `json:"loyaltyAction"`
	TotalAmount                decimal.Decimal                    `json:"totalAmount"`
	ApplyingAmount             *decimal.Decimal                   `json:"applyingAmount,omitempty"`
	CollectingAmount           *decimal.Decimal                   `json:"collectingAmount,omitempty"`
	ApplyBonuses               int                                `json:"applyBonuses"`
	CollectBonuses             *int                               `json:"collectBonuses,omitempty"`
	Promocode                  *ApplyPurchaseTransactionPromocode `json:"promocode,omitempty"`
	Items                      []Item                             `json:"items,omitempty"`
	IsConfirmationCodeRequired bool                               `json:"isConfirmationCodeRequired,omitempty"`
	ConfirmationCode           string                             `json:"confirmationCode,omitempty"`
}

type ApplyPurchaseTransactionPromocode struct {
	Promocode         string          `json:"promocode"`
	PromocodeDiscount decimal.Decimal `json:"promocodeDiscount"`
}

type ApplyPurchaseReply struct {
	Confirmation ApplyPurchaseReplyConfirmation `json:"confirmation"`
}

type ApplyPurchaseReplyConfirmation struct {
	PhoneNumber string                    `json:"phoneNumber"`
	PurchaseID  string                    `json:"purchaseId"`
	TotalAmount decimal.Decimal           `json:"totalAmount"`
	PaidAmount  decimal.Decimal           `json:"paidAmount"`
	Bonuses     ApplyPurchaseReplyBonuses `json:"bonuses"`
}

type ApplyPurchaseReplyBonuses struct {
	Applied   int `json:"applied"`
	Collected int `json:"collected"`
	Pending   int `json:"pending"`
	Available int `json:"available"`
}

// calculate-return

// CalculateReturnQuery is a request model for /calculate-return API
type CalculateReturnQuery struct {
	Calculate CalculateReturnQueryCalculate `json:"calculate"`
}

// CalculateReturnQueryCalculate holds info for return calculation
type CalculateReturnQueryCalculate struct {
	PhoneNumber  string          `json:"phoneNumber,omitempty"`
	Card         string          `json:"card,omitempty"`
	ExternalID   string          `json:"externalId,omitempty"`
	PurchaseID   string          `json:"purchaseId"`
	RefundAmount decimal.Decimal `json:"refundAmount"`
	Items        []Item          `json:"items,omitempty"`
}

// CalculateReturnReply is a response model for /calculate-return API
type CalculateReturnReply struct {
	Calculation CalculateReturnReplyCalculate `json:"calculation"`
}

// CalculateReturnReplyCalculate holds results of return calculation
type CalculateReturnReplyCalculate struct {
	PurchaseID       string          `json:"purchaseId"`
	RefundAmount     decimal.Decimal `json:"refundAmount"`
	RecoveredBonuses int             `json:"recoveredBonuses"`
	CancelledBonuses int             `json:"cancelledBonuses"`
}

// apply-return

// ApplyReturnQuery is a request model for /apply-return API
type ApplyReturnQuery struct {
	Transaction ApplyReturnTransaction `json:"transaction"`
}

// ApplyReturnTransaction holds info for apply return
type ApplyReturnTransaction struct {
	PhoneNumber       string           `json:"phoneNumber,omitempty"`
	Card              string           `json:"card,omitempty"`
	ExternalID        string           `json:"externalId,omitempty"`
	IsAnonymousClient bool             `json:"isAnonymousClient,omitempty"`
	ID                string           `json:"id"`
	ExecutedAt        time.Time        `json:"executedAt"`
	PurchaseID        string           `json:"purchaseId"`
	ReceiptID         IntAsIntOrString `json:"receiptId,omitempty"`
	SessionID         IntAsIntOrString `json:"sessionId,omitempty"`
	ShopCode          string           `json:"shopCode"`
	ShopName          string           `json:"shopName"`
	Cashier           string           `json:"cashier,omitempty"`
	CashierID         string           `json:"cashierId,omitempty"`
	RefundAmount      decimal.Decimal  `json:"refundAmount"`
	Items             []Item           `json:"items,omitempty"`
}

// ApplyReturnReply is a response model for /apply-return API
type ApplyReturnReply struct {
	Confirmation ApplyReturnReplyConfirmation `json:"confirmation"`
}

// ApplyReturnReplyConfirmation holds results for apply return
type ApplyReturnReplyConfirmation struct {
	PhoneNumber      string          `json:"phoneNumber"` // deprecated
	RefundID         string          `json:"refundId"`
	RefundAmount     decimal.Decimal `json:"refundAmount"`
	RecoveredBonuses int             `json:"recoveredBonuses"`
	CancelledBonuses int             `json:"cancelledBonuses"`
}

type Item struct {
	ExternalID         string           `json:"externalId,omitempty"`
	SKU                string           `json:"sku"`
	ItemTitle          string           `json:"itemTitle"`
	ItemCount          float64          `json:"itemCount"`
	BuyingPrice        *decimal.Decimal `json:"buyingPrice,omitempty"`
	Price              decimal.Decimal  `json:"price"`
	Amount             *decimal.Decimal `json:"amount,omitempty"`
	Category           string           `json:"category,omitempty"`
	CategoryExternalID string           `json:"categoryExternalID,omitempty"`
	MinPrice           *decimal.Decimal `json:"minPrice,omitempty"`
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
	TotalAmount decimal.Decimal  `json:"totalAmount"`
	Loyalty     *SetOrderLoyalty `json:"loyalty,omitempty"`
	Promocode   string           `json:"promocode,omitempty"`
	Items       []Item           `json:"items,omitempty"`
}

type SetOrderLoyalty struct {
	Action                     string            `json:"action"`
	ApplyingAmount             *decimal.Decimal  `json:"applyingAmount,omitempty"`
	CollectingAmount           *decimal.Decimal  `json:"collectingAmount,omitempty"`
	ApplyBonuses               *IntAsIntOrString `json:"applyBonuses"`
	CollectBonuses             *int              `json:"collectBonuses,omitempty"`
	IsConfirmationCodeRequired bool              `json:"isConfirmationCodeRequired,omitempty"`
	ConfirmationCode           string            `json:"confirmationCode,omitempty"`
}

type SetOrderReply struct {
	OperationResult SetOrderOperationResult `json:"operationResult"`
	ClientBonuses   ClientBonusesReply      `json:"clientBonuses"`
}

type SetOrderOperationResult struct {
	AppliedBonuses   int             `json:"appliedBonuses"`
	CollectedBonuses int             `json:"collectedBonuses"`
	RemainingAmount  decimal.Decimal `json:"remainingAmount"`
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

// get-history

type GetHistoryQuery struct {
	Client     ClientQuery
	Pagination *PaginationQuery `json:"pagination,omitempty"`
}

type PaginationQuery struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type GetHistoryReply struct {
	History    []HistoryEntry  `json:"history"`
	Pagination PaginationReply `json:"pagination"`
}

type HistoryEntry struct {
	At                       time.Time             `json:"at"`
	Amount                   int                   `json:"amount"`
	Operation                string                `json:"operation"`
	OperationName            string                `json:"operationName"`
	OperationApplied         *HistoryPurchaseEntry `json:"OPERATION_APPLIED,omitempty"`
	OperationCollected       *HistoryPurchaseEntry `json:"OPERATION_COLLECTED,omitempty"`
	OperationExpired         *struct{}             `json:"OPERATION_EXPIRED,omitempty"`
	OperationRefunded        *HistoryReturnEntry   `json:"OPERATION_REFUNDED,omitempty"`
	OperationCancelled       *HistoryReturnEntry   `json:"OPERATION_CANCELLED,omitempty"`
	OperationReceived        *HistoryReceiveEntry  `json:"OPERATION_RECEIVED,omitempty"`
	OperationRecalled        *HistoryReceiveEntry  `json:"OPERATION_RECALLED,omitempty"`
	OperationApplyReverted   *HistoryPurchaseEntry `json:"OPERATION_APPLY_REVERTED,omitempty"`
	OperationCollectReverted *HistoryPurchaseEntry `json:"OPERATION_COLLECT_REVERTED,omitempty"`
	OperationCollectedFriend *HistoryPurchaseEntry `json:"OPERATION_COLLECTED_FRIEND,omitempty"`
	OperationOther           *struct{}             `json:"OPERATION_OTHER"`
}

type HistoryPurchaseEntry struct {
	PurchaseID  string          `json:"purchaseId"`
	ExecutedAt  time.Time       `json:"executedAt"`
	TotalAmount decimal.Decimal `json:"totalAmount"`
}

type HistoryReturnEntry struct {
	ReturnID     string          `json:"returnId"`
	ExecutedAt   time.Time       `json:"executedAt"`
	RefundAmount decimal.Decimal `json:"refundAmount"`
}

type HistoryReceiveEntry struct {
	ActionName string `json:"actionName"`
	Comment    string `json:"comment"`
}

type PaginationReply struct {
	Total int `json:"total"`
}

// issue-promocode

type IssuePromocodeQuery struct {
	Client *ClientQuery `json:"client,omitempty"`
	Code   string       `json:"code"`
}

type IssuePromocodeReply struct {
	Promocode string `json:"promocode"`
}

// revert-purchase

type RevertPurchaseQuery struct {
	Transaction RevertPurchaseTransaction `json:"transaction"`
}

type RevertPurchaseTransaction struct {
	ID         string `json:"id"`
	ExecutedAt string `json:"executedAt"`
	PurchaseID string `json:"purchaseId"`
	Cashier    string `json:"cashier,omitempty"`
	CashierID  string `json:"cashierId,omitempty"`
}

type RevertPurchaseReply struct {
	OperationResult RevertPurchaseOperationResult `json:"operationResult"`
}

type RevertPurchaseOperationResult struct {
	RefundedBonuses  int `json:"refundedBonuses"`
	CancelledBonuses int `json:"cancelledBonuses"`
}
