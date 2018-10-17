package iso20022

// Specification of the price type.
type PriceType1 struct {

	// Structured format.
	Structured *TypeOfPrice7Code `xml:"Strd"`

	// Additional information about the type of charge.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (p *PriceType1) SetStructured(value string) {
	p.Structured = (*TypeOfPrice7Code)(&value)
}

func (p *PriceType1) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

// Specification of the price type.
type PriceType2 struct {

	// Structured format.
	Structured *TypeOfPrice6Code `xml:"Strd"`

	// Additional information about the type of charge.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (p *PriceType2) SetStructured(value string) {
	p.Structured = (*TypeOfPrice6Code)(&value)
}

func (p *PriceType2) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}
