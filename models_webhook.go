package cloudloyalty_client

const (
	EventSourceInternal = "internal"
	EventSourceAPI      = "api"
)

type Events []Event

type Event struct {
	EventName            string                     `json:"event"`
	EventID              string                     `json:"eventId"`
	EventTime            string                     `json:"eventTime"`
	Source               string                     `json:"source"`
	ClientNew            *ClientNewEvent            `json:"EVENT_CLIENT_NEW"`
	ClientChanged        *ClientChangedEvent        `json:"EVENT_CLIENT_CHANGED"`
	ClientBonusesChanged *ClientBonusesChangedEvent `json:"EVENT_CLIENT_BONUSES_CHANGED"`
}

type ClientNewEvent struct {
	ClientData GetBalanceReplyClient `json:"clientData"`
}

type ClientChangedEvent struct {
	Client     ClientQuery           `json:"client"`
	ClientData GetBalanceReplyClient `json:"clientData"`
}

type ClientBonusesChangedEvent struct {
	Client        ClientQuery                      `json:"client"`
	ClientBonuses ClientBonusesChangedEventBonuses `json:"clientBonuses"`
}

type ClientBonusesChangedEventBonuses struct {
	TotalAmount     int `json:"totalAmount"`
	AvailableAmount int `json:"availableAmount"`
	PendingAmount   int `json:"pendingAmount"`
}
