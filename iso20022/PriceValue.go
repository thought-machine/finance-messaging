package iso20022

// Value given to a price.
type PriceValue1 struct {

	// Price expressed as a currency and value.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (p *PriceValue1) SetAmount(value, currency string) {
	p.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Value given to a price.
type PriceValue2 struct {

	// Price expressed as a currency and value.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Price expressed as a rate, ie, percentage.
	Rate *PercentageRate `xml:"Rt,omitempty"`
}

func (p *PriceValue2) SetAmount(value, currency string) {
	p.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (p *PriceValue2) SetRate(value string) {
	p.Rate = (*PercentageRate)(&value)
}

// Value given to a price.
type PriceValue5 struct {

	// Price expressed as a currency and value.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (p *PriceValue5) SetAmount(value, currency string) {
	p.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}
