package partyHandler

type partyRequest struct {
	Title *string `json:"title"`
	Theme *string `json:"theme"`
}

type partyResponse struct {
	Parties []party `json:"parties"`
}

type party struct {
	Title       string      `json:"title"`
	Id          string      `json:"id"`
	TimeOfDay   []string    `json:"time_of_day"`
	Theme       string      `json:"theme"`
	Decorations *decoration `json:"decorations"`
	Drinks      *drink      `json:"drinks"`
	Food        *food       `json:"food"`
}

type decoration struct {
	Name        string
	Description string
	Link        string
}

type drink struct {
	Name        string
	Description string
	Recipe      string
}

type food struct {
	Name        string
	Description string
	Recipe      string
}
