package cloudloyalty_client

import "fmt"

const (
	ErrGeneralError            = 1
	ErrMalformedRequest        = 2
	ErrClientNotFound          = 3
	ErrClientSuspended         = 4
	ErrShopNotFound            = 5
	ErrIncorrectBonusAmount    = 6
	ErrIncorrectReturnItem     = 10
	ErrIncorrectReturnAmount   = 11
	ErrAlreadyProcessed        = 17
	ErrIncorrectPhone          = 20
	ErrPurchaseNotFound        = 21
	ErrDuplicatingPhone        = 24
	ErrDuplicatingCard         = 25
	ErrTooManyCodeRequests     = 28
	ErrEmptyPhone              = 29
	ErrDuplicatingExternalID   = 30
	ErrOrderNotFound           = 31
	ErrOrderAlreadyProcessed   = 32
	ErrPromocodeNotFound       = 33
	ErrPromocodeNotApplicable  = 34
	ErrPromocodeAlreadyUsed    = 35
	ErrPromocodeExpired        = 36
	ErrPromocodeRequiresClient = 37
)

var DescriptionEN = map[int]string{
	ErrGeneralError:            "Request has not been complete",
	ErrMalformedRequest:        "Malformed request or incorrect JSON",
	ErrClientNotFound:          "Client not found",
	ErrClientSuspended:         "Client suspended",
	ErrShopNotFound:            "Shop not found",
	ErrIncorrectBonusAmount:    "Too many bonuses requested",
	ErrIncorrectReturnItem:     "Returning item wasn't purchased",
	ErrIncorrectReturnAmount:   "Refund amount is greater than amount of purchase",
	ErrAlreadyProcessed:        "Transaction with this TxID already processed",
	ErrIncorrectPhone:          "Incorrect phone number",
	ErrPurchaseNotFound:        "Purchase transaction not found",
	ErrDuplicatingPhone:        "Client with this phone number already exists",
	ErrDuplicatingCard:         "Client with this card already exists",
	ErrTooManyCodeRequests:     "Too many confirmation codes requested",
	ErrEmptyPhone:              "Client has no phone number specified",
	ErrDuplicatingExternalID:   "Client with this externalId already exists",
	ErrOrderNotFound:           "Order not found",
	ErrOrderAlreadyProcessed:   "Order already processed",
	ErrPromocodeNotFound:       "Promocode does not exist",
	ErrPromocodeNotApplicable:  "Promocode cannot be applied to this purchase",
	ErrPromocodeAlreadyUsed:    "Promocode has already been used maximum times",
	ErrPromocodeExpired:        "Promocode expired or not yet started to operate",
	ErrPromocodeRequiresClient: "Promocode requires client to be authorized",
}

var DescriptionRU = map[int]string{
	ErrGeneralError:            "Чек не обработан процессингом или обработан с ошибкой",
	ErrMalformedRequest:        "В запросе к процессингу обнаружена ошибка или неверный вид JSON",
	ErrClientNotFound:          "Клиент не найден",
	ErrClientSuspended:         "Аккаунт клиента заблокирован",
	ErrShopNotFound:            "Не найден магазин",
	ErrIncorrectBonusAmount:    "Списание бонусов превышает допустимое значение",
	ErrIncorrectReturnItem:     "Возвращаемый товар не найден в чеке продажи",
	ErrIncorrectReturnAmount:   "Сумма возврата больше суммы продажи",
	ErrAlreadyProcessed:        "Чек с данным номером уже обработан",
	ErrIncorrectPhone:          "Номер телефона клиента не валиден",
	ErrPurchaseNotFound:        "Продажа не найдена",
	ErrDuplicatingPhone:        "Клиент с таким номером телефона уже существует",
	ErrDuplicatingCard:         "Карта уже привязана к другому клиенту",
	ErrTooManyCodeRequests:     "Слишком частая отправка кода подтверждения",
	ErrEmptyPhone:              "У клиента не задан номер телефона",
	ErrDuplicatingExternalID:   "Клиент с таким внешним идентификатором уже существует",
	ErrOrderNotFound:           "Заказ не найден",
	ErrOrderAlreadyProcessed:   "Заказ уже обработан",
	ErrPromocodeNotFound:       "Промокод не найден",
	ErrPromocodeNotApplicable:  "Условия промокода не выполнены",
	ErrPromocodeAlreadyUsed:    "Промокод уже использован максимальное число раз",
	ErrPromocodeExpired:        "Время действия промокода еще не наступило или уже прошло",
	ErrPromocodeRequiresClient: "Промокод работает только совместно с бонусным счетом",
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
