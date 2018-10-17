package iso20022

// Set of elements to identify a proprietary price.
type ProprietaryPrice1 struct {

	// Identifies the type of price reported.
	Type *Max35Text `xml:"Tp"`

	// Proprietary price specification related to the underlying transaction.
	Price *CurrencyAndAmount `xml:"Pric"`
}

func (p *ProprietaryPrice1) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryPrice1) SetPrice(value, currency string) {
	p.Price = NewCurrencyAndAmount(value, currency)
}

// Set of elements used to identify a proprietary price.
type ProprietaryPrice2 struct {

	// Specifies the type of price.
	Type *Max35Text `xml:"Tp"`

	// Proprietary price specification related to the underlying transaction.
	Price *ActiveOrHistoricCurrencyAndAmount `xml:"Pric"`
}

func (p *ProprietaryPrice2) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryPrice2) SetPrice(value, currency string) {
	p.Price = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
