package iso20022

// Identification of market in which a trade transaction has been executed.
type PlaceOfTradeIdentification1 struct {

	// Identification and type of the place of trade.
	MarketTypeAndIdentification *MarketIdentification84 `xml:"MktTpAndId,omitempty"`

	// Legal entity identification as an alternate identification for a place of trade.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PlaceOfTradeIdentification1) AddMarketTypeAndIdentification() *MarketIdentification84 {
	p.MarketTypeAndIdentification = new(MarketIdentification84)
	return p.MarketTypeAndIdentification
}

func (p *PlaceOfTradeIdentification1) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

// Identification of market in which a trade transaction has been executed.
type PlaceOfTradeIdentification2 struct {

	// Identification and type of the place of trade.
	MarketTypeAndIdentification *MarketIdentification90 `xml:"MktTpAndId,omitempty"`

	// Legal entity identification as an alternate identification for a place of trade.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PlaceOfTradeIdentification2) AddMarketTypeAndIdentification() *MarketIdentification90 {
	p.MarketTypeAndIdentification = new(MarketIdentification90)
	return p.MarketTypeAndIdentification
}

func (p *PlaceOfTradeIdentification2) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
