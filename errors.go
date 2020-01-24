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

var russianText = map[int]string{
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

func RussianDescription(code int) string {
	if desc, ok := russianText[code]; ok {
		return desc
	}
	return fmt.Sprintf("Ошибка #%d", code)
}
