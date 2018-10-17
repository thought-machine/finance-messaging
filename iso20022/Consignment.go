package iso20022

// Physical packaging of goods for transport.
type Consignment1 struct {

	// Total quantity of packaging units, eg number of boxes, containers, pallets, etc
	TotalQuantity *Quantity3 `xml:"TtlQty,omitempty"`

	// Total volume of goods shipped, eg number of cubic meters.
	TotalVolume *Quantity3 `xml:"TtlVol,omitempty"`

	// Total weight of goods shipped, eg number of kg, tons.
	TotalWeight *Quantity3 `xml:"TtlWght,omitempty"`
}

func (c *Consignment1) AddTotalQuantity() *Quantity3 {
	c.TotalQuantity = new(Quantity3)
	return c.TotalQuantity
}

func (c *Consignment1) AddTotalVolume() *Quantity3 {
	c.TotalVolume = new(Quantity3)
	return c.TotalVolume
}

func (c *Consignment1) AddTotalWeight() *Quantity3 {
	c.TotalWeight = new(Quantity3)
	return c.TotalWeight
}

// Specifies the arrangement of the transport of goods and services and the parties involved in this process.
type Consignment2 struct {

	// Party consigning goods as stipulated in the transport contract by the party ordering transport.
	Consignor *TradeParty1 `xml:"Consgnr,omitempty"`

	// Party to which goods are consigned.
	Consignee *TradeParty1 `xml:"Consgn,omitempty"`

	// Particular aircraft, vehicle, vessel or other device used for the transport of a consignment.
	TransportMeans []*TransportMeans3 `xml:"TrnsprtMeans,omitempty"`
}

func (c *Consignment2) AddConsignor() *TradeParty1 {
	c.Consignor = new(TradeParty1)
	return c.Consignor
}

func (c *Consignment2) AddConsignee() *TradeParty1 {
	c.Consignee = new(TradeParty1)
	return c.Consignee
}

func (c *Consignment2) AddTransportMeans() *TransportMeans3 {
	newValue := new(TransportMeans3)
	c.TransportMeans = append(c.TransportMeans, newValue)
	return newValue
}

// Physical packaging of goods for transport.
type Consignment3 struct {

	// Total quantity of packaging units, eg number of boxes, containers, pallets, etc
	TotalQuantity *Quantity10 `xml:"TtlQty,omitempty"`

	// Total volume of goods shipped, eg number of cubic meters.
	TotalVolume *Quantity10 `xml:"TtlVol,omitempty"`

	// Total weight of goods shipped, eg number of kg, tons.
	TotalWeight *Quantity10 `xml:"TtlWght,omitempty"`
}

func (c *Consignment3) AddTotalQuantity() *Quantity10 {
	c.TotalQuantity = new(Quantity10)
	return c.TotalQuantity
}

func (c *Consignment3) AddTotalVolume() *Quantity10 {
	c.TotalVolume = new(Quantity10)
	return c.TotalVolume
}

func (c *Consignment3) AddTotalWeight() *Quantity10 {
	c.TotalWeight = new(Quantity10)
	return c.TotalWeight
}

// Specifies the arrangement of the transport of goods and services and the parties involved in this process.
type Consignment4 struct {

	// Party consigning goods as stipulated in the transport contract by the party ordering transport.
	Consignor *TradeParty3 `xml:"Consgnr,omitempty"`

	// Party to which goods are consigned.
	Consignee *TradeParty3 `xml:"Consgn,omitempty"`

	// Particular aircraft, vehicle, vessel or other device used for the transport of a consignment.
	TransportMeans []*TransportMeans3 `xml:"TrnsprtMeans,omitempty"`
}

func (c *Consignment4) AddConsignor() *TradeParty3 {
	c.Consignor = new(TradeParty3)
	return c.Consignor
}

func (c *Consignment4) AddConsignee() *TradeParty3 {
	c.Consignee = new(TradeParty3)
	return c.Consignee
}

func (c *Consignment4) AddTransportMeans() *TransportMeans3 {
	newValue := new(TransportMeans3)
	c.TransportMeans = append(c.TransportMeans, newValue)
	return newValue
}
