package cloudloyalty_client

const (
	ErrGeneralError             = 1
	ErrMalformedRequest         = 2
	ErrClientNotFound           = 3
	ErrClientSuspended          = 4
	ErrShopNotFound             = 5
	ErrIncorrectBonusAmount     = 6
	ErrTooManyPurchases         = 7
	ErrIncorrectReturnItem      = 10
	ErrIncorrectReturnAmount    = 11
	ErrIncorrectReturnPurchase  = 13
	ErrAlreadyProcessed         = 17
	ErrEmptyRowsArray           = 18
	ErrIncorrectPhone           = 20
	ErrPurchaseNotFound         = 21
	ErrPurchaseDiscarded        = 22
	ErrDuplicatingPhone         = 24
	ErrDuplicatingCard          = 25
	ErrTooManyCodeRequests      = 28
	ErrEmptyPhone               = 29
	ErrDuplicatingExternalID    = 30
	ErrOrderNotFound            = 31
	ErrOrderAlreadyProcessed    = 32
	ErrPromocodeNotFound        = 33
	ErrPromocodeNotApplicable   = 34
	ErrPromocodeAlreadyUsed     = 35
	ErrPromocodeExpired         = 36
	ErrPromocodeRequiresClient  = 37
	ErrGiftCardNotFound         = 40
	ErrGiftCardNotApplicable    = 41
	ErrGiftCardNotActivated     = 42
	ErrGiftCardExpired          = 43
	ErrGiftCardBlocked          = 44
	ErrGiftCardNoFunds          = 45
	ErrGiftCardAlreadySold      = 46
	ErrGiftCardAlreadyActivated = 47
	ErrDuplicatingEmail         = 48
	ErrTemporarilyUnavailable   = 49
)
