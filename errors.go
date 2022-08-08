package cloudloyalty_client

import "fmt"

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

var DescriptionEN = map[int]string{
	ErrGeneralError:             "Request not processed or completed with error",
	ErrMalformedRequest:         "Malformed request",
	ErrClientNotFound:           "Client not found",
	ErrClientSuspended:          "Client account suspended",
	ErrShopNotFound:             "Shop not found",
	ErrIncorrectBonusAmount:     "Too many bonuses requested",
	ErrTooManyPurchases:         "Exceeded maximum number of purchases",
	ErrIncorrectReturnItem:      "Returning item wasn't purchased",
	ErrIncorrectReturnAmount:    "Refund amount exceeds the amount of the purchase",
	ErrIncorrectReturnPurchase:  "No such purchase for the return",
	ErrAlreadyProcessed:         "Purchase with this TxID already processed",
	ErrEmptyRowsArray:           "Purchase must include at least one item",
	ErrIncorrectPhone:           "Incorrect phone number",
	ErrPurchaseNotFound:         "Purchase not found",
	ErrDuplicatingPhone:         "Client with this phone number already exists",
	ErrDuplicatingCard:          "Client with this card already exists",
	ErrTooManyCodeRequests:      "Confirmation code sent too often",
	ErrEmptyPhone:               "The client's phone number is not specified",
	ErrDuplicatingExternalID:    "Client with this externalId already exists",
	ErrOrderNotFound:            "Order not found",
	ErrOrderAlreadyProcessed:    "Order already processed",
	ErrPromocodeNotFound:        "Promocode does not exist",
	ErrPromocodeNotApplicable:   "Promocode cannot be applied to this purchase",
	ErrPromocodeAlreadyUsed:     "Promocode has already been used maximum number of times",
	ErrPromocodeExpired:         "Promocode validity time is over or not yet started",
	ErrPromocodeRequiresClient:  "Promocode requires client to be authorized",
	ErrGiftCardNotFound:         "Gift card does not exist",
	ErrGiftCardNotApplicable:    "Gift card cannot be applied to this purchase",
	ErrGiftCardNotActivated:     "Gift card is not yet activated",
	ErrGiftCardExpired:          "Gift card expired",
	ErrGiftCardBlocked:          "Gift card blocked",
	ErrGiftCardNoFunds:          "Gift card ran out of funds",
	ErrGiftCardAlreadySold:      "Gift card already sold",
	ErrGiftCardAlreadyActivated: "Gift card already activated",
	ErrDuplicatingEmail:         "Client with this e-mail already exists",
	ErrTemporarilyUnavailable:   "Application temporarily unavailable",
}

var DescriptionRU = map[int]string{
	ErrGeneralError:             "Чек не обработан процессингом или обработан с ошибкой",
	ErrMalformedRequest:         "Запрос к процессингу составлен некорректно",
	ErrClientNotFound:           "Клиент не найден",
	ErrClientSuspended:          "Аккаунт клиента заблокирован",
	ErrShopNotFound:             "Не найден магазин",
	ErrIncorrectBonusAmount:     "Списание бонусов превышает допустимое значение",
	ErrTooManyPurchases:         "Превышено количество покупок",
	ErrIncorrectReturnItem:      "Возвращаемый товар не найден в чеке продажи",
	ErrIncorrectReturnAmount:    "Сумма возврата больше суммы продажи",
	ErrIncorrectReturnPurchase:  "Не найдена продажа для возврата",
	ErrAlreadyProcessed:         "Чек с данным номером уже обработан",
	ErrEmptyRowsArray:           "Продажа должна включать хотя бы один товар",
	ErrIncorrectPhone:           "Номер телефона клиента не валиден",
	ErrPurchaseNotFound:         "Продажа не найдена",
	ErrDuplicatingPhone:         "Клиент с таким номером телефона уже существует",
	ErrDuplicatingCard:          "Карта уже привязана к другому клиенту",
	ErrTooManyCodeRequests:      "Слишком частая отправка кода подтверждения",
	ErrEmptyPhone:               "У клиента не задан номер телефона",
	ErrDuplicatingExternalID:    "Клиент с таким внешним идентификатором уже существует",
	ErrOrderNotFound:            "Заказ не найден",
	ErrOrderAlreadyProcessed:    "Заказ уже обработан",
	ErrPromocodeNotFound:        "Промокод не найден",
	ErrPromocodeNotApplicable:   "Условия промокода не выполнены",
	ErrPromocodeAlreadyUsed:     "Промокод уже использован максимальное число раз",
	ErrPromocodeExpired:         "Время действия промокода еще не наступило или уже прошло",
	ErrPromocodeRequiresClient:  "Промокод работает только совместно с бонусным счетом",
	ErrGiftCardNotFound:         "Подарочная карта не найдена",
	ErrGiftCardNotApplicable:    "Условия применения подарочной карты не выполнены",
	ErrGiftCardNotActivated:     "Подарочная карта еще не активирована",
	ErrGiftCardExpired:          "Срок действия подарочной карты истёк или еще не наступил",
	ErrGiftCardBlocked:          "Подарочная карта заблокирована",
	ErrGiftCardNoFunds:          "На подарочной карте нет средств",
	ErrGiftCardAlreadySold:      "Подарочная карта уже продана",
	ErrGiftCardAlreadyActivated: "Подарочная карта уже активирована",
	ErrDuplicatingEmail:         "Клиент с таким e-mail уже существует",
	ErrTemporarilyUnavailable:   "Применение временно недоступно",
}

func Description(code int) string {
	if desc, ok := DescriptionEN[code]; ok {
		return desc
	}
	return fmt.Sprintf("Error #%d", code)
}

func RussianDescription(code int) string {
	if desc, ok := DescriptionRU[code]; ok {
		return desc
	}
	return fmt.Sprintf("Ошибка #%d", code)
}
