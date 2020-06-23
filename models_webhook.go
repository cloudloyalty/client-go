package cloudloyalty_client

const (
	EventSourceInternal = "internal"
	EventSourceAPI      = "api"
)

type Events []Event

type Event struct {
	EventName               string                               `json:"event"`
	EventID                 string                               `json:"eventId"`
	EventTime               string                               `json:"eventTime"`
	Source                  string                               `json:"source"`
	ClientNew               *ClientNewEvent                      `json:"EVENT_CLIENT_NEW"`
	ClientChanged           *ClientChangedEvent                  `json:"EVENT_CLIENT_CHANGED"`
	ClientBonusesChanged    *ClientBonusesChangedEvent           `json:"EVENT_CLIENT_BONUSES_CHANGED"`
	ClientEmailSubscribed   *ClientEmailSubscriptionChangedEvent `json:"EVENT_CLIENT_EMAIL_SUBSCRIBED"`
	ClientEmailUnsubscribed *ClientEmailSubscriptionChangedEvent `json:"EVENT_CLIENT_EMAIL_UNSUBSCRIBED"`
}

type ClientNewEvent struct {
	ClientData ClientChangedEventClientData `json:"clientData"`
}

type ClientChangedEvent struct {
	Client     ClientQuery                  `json:"client"`
	ClientData ClientChangedEventClientData `json:"clientData"`
}

type ClientChangedEventClientData struct {
	PhoneNumber       string      `json:"phoneNumber"`
	Card              string      `json:"card"`
	ExternalID        string      `json:"externalId"`
	Surname           string      `json:"surname"`
	Name              string      `json:"name"`
	PatronymicName    string      `json:"patronymicName"`
	FullName          string      `json:"fullName"`
	Gender            int         `json:"gender"`
	Birthdate         *Birthdate  `json:"birthdate"`
	Email             string      `json:"email"`
	Level             int         `json:"level"`
	IsEmailSubscribed bool        `json:"isEmailSubscribed"`
	IsPhoneSubscribed bool        `json:"isPhoneSubscribed"`
	ExtraFields       ExtraFields `json:"extraFields"`
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

type ClientEmailSubscriptionChangedEvent struct {
	Client ClientQuery `json:"client"`
	Email  string      `json:"email"`
}
