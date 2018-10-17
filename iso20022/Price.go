package iso20022

// Type and information about a price.
type Price2 struct {

	// Specification of the price type.
	Type *YieldedOrValueType1Choice `xml:"Tp"`

	// Value of the price, for example, as a currency and value.
	Value *PriceRateOrAmountChoice `xml:"Val"`
}

func (p *Price2) AddType() *YieldedOrValueType1Choice {
	p.Type = new(YieldedOrValueType1Choice)
	return p.Type
}

func (p *Price2) AddValue() *PriceRateOrAmountChoice {
	p.Value = new(PriceRateOrAmountChoice)
	return p.Value
}

// Type and information about a price.
type Price3 struct {

	// Specification of the price type.
	Type *YieldedOrValueType1Choice `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmount1Choice `xml:"Val"`
}

func (p *Price3) AddType() *YieldedOrValueType1Choice {
	p.Type = new(YieldedOrValueType1Choice)
	return p.Type
}

func (p *Price3) AddValue() *PriceRateOrAmount1Choice {
	p.Value = new(PriceRateOrAmount1Choice)
	return p.Value
}

// Provides the value and optionaly the type of price.
type Price4 struct {

	// Value of the price.
	Value *PriceRateOrAmountChoice `xml:"Val"`

	// Specification of the price type.
	Type *PriceValueType7Code `xml:"Tp,omitempty"`
}

func (p *Price4) AddValue() *PriceRateOrAmountChoice {
	p.Value = new(PriceRateOrAmountChoice)
	return p.Value
}

func (p *Price4) SetType(value string) {
	p.Type = (*PriceValueType7Code)(&value)
}

// Provides the value, type and source of price.
type Price6 struct {

	// Value of the price expressed as a currency and value or as a rate.
	RateOrAmount *PriceRateOrAmountChoice `xml:"RateOrAmt"`

	// Specification of the price type.
	Type *TypeOfPrice13Code `xml:"Tp"`

	// Source for the price valuation.
	Source *PriceSource2Code `xml:"Src"`
}

func (p *Price6) AddRateOrAmount() *PriceRateOrAmountChoice {
	p.RateOrAmount = new(PriceRateOrAmountChoice)
	return p.RateOrAmount
}

func (p *Price6) SetType(value string) {
	p.Type = (*TypeOfPrice13Code)(&value)
}

func (p *Price6) SetSource(value string) {
	p.Source = (*PriceSource2Code)(&value)
}
