package iso20022

// Unique and unambiguous way to identify an organisation.
type PartyIdentification44Choice struct {

	// Unique and unambiguous way to identify an organisation.
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous way to identify an organisation.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`

	// Unique and unambiguous way to identify an organisation.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PartyIdentification44Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification44Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

func (p *PartyIdentification44Choice) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}