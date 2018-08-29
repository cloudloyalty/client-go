package cloudloyalty_client

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

type errorReply struct {
	ErrorCode   int    `json:"errorCode"`
	Description string `json:"description"`
}

// get-balance

type GetBalanceQuery struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Card        string `json:"card,omitempty"`
	ExternalId  string `json:"externalId,omitempty"`
}

type GetBalanceReply struct {
	Client  GetBalanceReplyClient  `json:"client"`
	Bonuses []GetBalanceReplyBonus `json:"bonuses"`
}

type GetBalanceReplyClient struct {
	PhoneNumber       string            `json:"phoneNumber"`
	CardString        string            `json:"cardString"`
	ExternalId        string            `json:"externalId"`
	Surname           string            `json:"surname"`
	Name              string            `json:"name"`
	PatronymicName    string            `json:"patronymicName"`
	FullName          string            `json:"fullName"`
	Gender            int               `json:"gender"`
	Birthdate         string            `json:"birthdate"`
	Email             string            `json:"email"`
	Level             int               `json:"level"`
	IsEmailSubscribed bool              `json:"isEmailSubscribed"`
	IsPhoneSubscribed bool              `json:"isPhoneSubscribed"`
	ExtraFields       map[string]string `json:"extraFields"`
	Bonuses           int               `json:"bonuses"`
	PendingBonuses    int               `json:"pendingBonuses"`
}

type GetBalanceReplyBonus struct {
	ExpireAt string `json:"expireAt"`
	Amount   int    `json:"amount"`
}

// calculate-purchase

type CalculatePurchaseQuery struct {
	Calculate CalculatePurchaseQueryCalculate `json:"calculate"`
}

type CalculatePurchaseQueryCalculate struct {
	PhoneNumber       string          `json:"phoneNumber,omitempty"`
	Card              string          `json:"card,omitempty"`
	ExternalId        string          `json:"externalId,omitempty"`
	IsAnonymousClient bool            `json:"isAnonymousClient,omitempty"`
	LoyaltyAction     string          `json:"loyaltyAction,omitempty"`
	TotalAmount       float64         `json:"totalAmount"`
	ApplyingAmount    float64         `json:"applyingAmount,omitempty"`
	CollectingAmount  float64         `json:"collectingAmount,omitempty"`
	ApplyBonuses      int             `json:"applyBonuses,omitempty"`
	CollectBonuses    int             `json:"collectBonuses,omitempty"`
	Promocode         string          `json:"promocode,omitempty"`
	Units             map[string]Unit `json:"units,omitempty"`
}

type Unit struct {
	Sku         string  `json:"sku"`
	ItemTitle   string  `json:"itemTitle"`
	ItemCount   float64 `json:"itemCount,omitempty"`
	BuyingPrice float64 `json:"buyingPrice,omitempty"`
	Price       float64 `json:"price"`
	Category    string  `json:"category,omitempty"`
}

type CalculatePurchaseReply struct {
	Calculation     CalculatePurchaseReplyCalculate `json:"calculation"`
	CalculatedUnits map[string]CalculatedUnit       `json:"calculatedUnits"`
}

type CalculatedUnit struct {
	OriginalPrice     float64 `json:"originalPrice"`
	PromocodeDiscount float64 `json:"promocodeDiscount"`
	BonusesDiscount   float64 `json:"bonusesDiscount"`
	TotalDiscount     float64 `json:"totalDiscount"`
	CalculatedPrice   float64 `json:"calculatedPrice"`
}

type CalculatePurchaseReplyCalculate struct {
	AppliedBonuses    int     `json:"appliedBonuses"`
	CollectedBonuses  int     `json:"collectedBonuses"`
	MaxBonuses        int     `json:"maxBonuses"`
	PromocodeDiscount float64 `json:"promocodeDiscount"`
	BonusesDiscount   float64 `json:"bonusesDiscount"`
	TotalDiscount     float64 `json:"totalDiscount"`
	RemainingAmount   float64 `json:"remainingDiscount"`
}

// apply-purchase

type ApplyPurchaseQuery struct {
	Transaction ApplyPurchaseTransaction `json:"transaction"`
}

type ApplyPurchaseTransaction struct {
	PhoneNumber       string  `json:"phoneNumber,omitempty"`
	Card              string  `json:"card,omitempty"`
	IsAnonymousClient bool    `json:"isAnonymousClient,omitempty"`
	ID                string  `json:"id"`
	ExecutedAt        string  `json:"executedAt"`
	ReceiptID         string  `json:"receiptId,omitempty"`
	SessionID         string  `json:"sessionId,omitempty"`
	ShopCode          string  `json:"shopCode"`
	ShopName          string  `json:"shopName"`
	Cashier           string  `json:"cashier,omitempty"`
	CashierID         string  `json:"cashierId,omitempty"`
	LoyaltyAction     string  `json:"loyaltyAction"`
	TotalAmount       float64 `json:"totalAmount"`
	ApplyBonuses      int     `json:"applyBonuses"`
	CollectBonuses    int     `json:"collectBonuses,omitempty"`
	Items             []Item  `json:"items,omitempty"`
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

// apply-return

type ApplyReturnQuery struct {
	Transaction ApplyReturnTransaction `json:"transaction"`
}

type ApplyReturnTransaction struct {
	PhoneNumber       string  `json:"phoneNumber,omitempty"`
	Card              string  `json:"card,omitempty"`
	IsAnonymousClient bool    `json:"isAnonymousClient,omitempty"`
	ID                string  `json:"id"`
	ExecutedAt        string  `json:"executedAt"`
	PurchaseID        string  `json:"purchaseId"`
	ReceiptID         string  `json:"receiptId,omitempty"`
	SessionID         string  `json:"sessionId,omitempty"`
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
	Sku         string  `json:"sku"`
	ItemTitle   string  `json:"itemTitle"`
	ItemCount   float64 `json:"itemCount"`
	BuyingPrice float64 `json:"buyingPrice,omitempty"`
	Price       float64 `json:"price"`
	Amount      float64 `json:"amount,omitempty"`
	Category    string  `json:"category,omitempty"`
}
