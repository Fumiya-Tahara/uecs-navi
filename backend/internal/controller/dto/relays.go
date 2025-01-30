package dto

type Relays struct {
	Relay1 bool `json:"relay_1"`
	Relay2 bool `json:"relay_2"`
	Relay3 bool `json:"relay_3"`
	Relay4 bool `json:"relay_4"`
	Relay5 bool `json:"relay_5"`
	Relay6 bool `json:"relay_6"`
	Relay7 bool `json:"relay_7"`
	Relay8 bool `json:"relay_8"`
}

func NewRelays(relay1, relay2, relay3, relay4, relay5, relay6, relay7, relay8 bool) *Relays {
	return &Relays{
		Relay1: relay1,
		Relay2: relay2,
		Relay3: relay3,
		Relay4: relay4,
		Relay5: relay5,
		Relay6: relay6,
		Relay7: relay7,
		Relay8: relay8,
	}
}
