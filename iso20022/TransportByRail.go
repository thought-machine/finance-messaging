package iso20022

// Information related to the transportation of goods by rail.
type TransportByRail2 struct {

	// Identifies the location where the goods are received for transportation.
	PlaceOfReceipt *Max35Text `xml:"PlcOfRct"`

	// Identifies the location of delivery of the goods.
	PlaceOfDelivery *Max35Text `xml:"PlcOfDlvry"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	RailCarrierName *Max35Text `xml:"RailCrrierNm,omitempty"`
}

func (t *TransportByRail2) SetPlaceOfReceipt(value string) {
	t.PlaceOfReceipt = (*Max35Text)(&value)
}

func (t *TransportByRail2) SetPlaceOfDelivery(value string) {
	t.PlaceOfDelivery = (*Max35Text)(&value)
}

func (t *TransportByRail2) SetRailCarrierName(value string) {
	t.RailCarrierName = (*Max35Text)(&value)
}

// Information related to the transportation of goods by rail.
type TransportByRail3 struct {

	// Identifies the location where the goods are received for transportation.
	PlaceOfReceipt []*Max35Text `xml:"PlcOfRct,omitempty"`

	// Identifies the location of delivery of the goods.
	PlaceOfDelivery []*Max35Text `xml:"PlcOfDlvry"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	RailCarrierName *Max35Text `xml:"RailCrrierNm,omitempty"`
}

func (t *TransportByRail3) AddPlaceOfReceipt(value string) {
	t.PlaceOfReceipt = append(t.PlaceOfReceipt, (*Max35Text)(&value))
}

func (t *TransportByRail3) AddPlaceOfDelivery(value string) {
	t.PlaceOfDelivery = append(t.PlaceOfDelivery, (*Max35Text)(&value))
}

func (t *TransportByRail3) SetRailCarrierName(value string) {
	t.RailCarrierName = (*Max35Text)(&value)
}

// Information related to the transportation of goods by rail.
type TransportByRail4 struct {

	// Identifies the location where the goods are received for transportation.
	PlaceOfReceipt *Max35Text `xml:"PlcOfRct"`

	// Identifies the location of delivery of the goods.
	PlaceOfDelivery *Max35Text `xml:"PlcOfDlvry"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	RailCarrierName *Max70Text `xml:"RailCrrierNm,omitempty"`

	// Country in which the carrier of the goods, for example, shipping company, is located or registered.
	RailCarrierCountry *CountryCode `xml:"RailCrrierCtry,omitempty"`

	// Name of the carrier's (for example, shipping company's) agent that acts on behalf of the carrier and may be the issuer of transport documents relating to the underlying shipment.
	CarrierAgentName *Max70Text `xml:"CrrierAgtNm,omitempty"`

	// Country of registration of the carrier's (for example, shipping company's) agent that acts on behalf of the carrier and may be the issuer of transport documents relating to the underlying shipment.
	CarrierAgentCountry *CountryCode `xml:"CrrierAgtCtry,omitempty"`
}

func (t *TransportByRail4) SetPlaceOfReceipt(value string) {
	t.PlaceOfReceipt = (*Max35Text)(&value)
}

func (t *TransportByRail4) SetPlaceOfDelivery(value string) {
	t.PlaceOfDelivery = (*Max35Text)(&value)
}

func (t *TransportByRail4) SetRailCarrierName(value string) {
	t.RailCarrierName = (*Max70Text)(&value)
}

func (t *TransportByRail4) SetRailCarrierCountry(value string) {
	t.RailCarrierCountry = (*CountryCode)(&value)
}

func (t *TransportByRail4) SetCarrierAgentName(value string) {
	t.CarrierAgentName = (*Max70Text)(&value)
}

func (t *TransportByRail4) SetCarrierAgentCountry(value string) {
	t.CarrierAgentCountry = (*CountryCode)(&value)
}

// Information related to the transportation of goods by rail.
type TransportByRail5 struct {

	// Identifies the location where the goods are received for transportation.
	PlaceOfReceipt []*Max35Text `xml:"PlcOfRct,omitempty"`

	// Identifies the location of delivery of the goods.
	PlaceOfDelivery []*Max35Text `xml:"PlcOfDlvry"`

	// Identifies the party that is responsible for the conveyance of the goods from one place to another.
	RailCarrierName *Max70Text `xml:"RailCrrierNm,omitempty"`

	// Country in which the carrier of the goods, for example, shipping company, is located or registered.
	RailCarrierCountry *CountryCode `xml:"RailCrrierCtry,omitempty"`

	// Name of the carrier's (for example, shipping company's) agent that acts on behalf of the carrier and may be the issuer of transport documents relating to the underlying shipment.
	CarrierAgentName *Max70Text `xml:"CrrierAgtNm,omitempty"`

	// Country of registration of the carrier's (for example, shipping company's) agent that acts on behalf of the carrier and may be the issuer of transport documents relating to the underlying shipment.
	CarrierAgentCountry *CountryCode `xml:"CrrierAgtCtry,omitempty"`
}

func (t *TransportByRail5) AddPlaceOfReceipt(value string) {
	t.PlaceOfReceipt = append(t.PlaceOfReceipt, (*Max35Text)(&value))
}

func (t *TransportByRail5) AddPlaceOfDelivery(value string) {
	t.PlaceOfDelivery = append(t.PlaceOfDelivery, (*Max35Text)(&value))
}

func (t *TransportByRail5) SetRailCarrierName(value string) {
	t.RailCarrierName = (*Max70Text)(&value)
}

func (t *TransportByRail5) SetRailCarrierCountry(value string) {
	t.RailCarrierCountry = (*CountryCode)(&value)
}

func (t *TransportByRail5) SetCarrierAgentName(value string) {
	t.CarrierAgentName = (*Max70Text)(&value)
}

func (t *TransportByRail5) SetCarrierAgentCountry(value string) {
	t.CarrierAgentCountry = (*CountryCode)(&value)
}
