package iso20022

// Set of elements qualifying the interest rate.
type Rate1 struct {

	// Percentage charged for the use of an amount of money, usually expressed at an annual rate. The interest rate is the ratio of the amount of interest paid during a certain period of time compared to the principal amount of the interest bearing financial instrument.
	// Example percentage rate : Rate expressed as a percentage, ie, in hundredths, eg, 0.7 is 7/10 of a percent, and 7.0 is 7%.
	// Example Textual rate : Rate is expressed as a text.
	Rate *RateTypeChoice `xml:"Rate"`

	// An amount range where the interest rate is applicable
	ValidityRange *CurrencyAndAmountRange `xml:"VldtyRg,omitempty"`
}

func (r *Rate1) AddRate() *RateTypeChoice {
	r.Rate = new(RateTypeChoice)
	return r.Rate
}

func (r *Rate1) AddValidityRange() *CurrencyAndAmountRange {
	r.ValidityRange = new(CurrencyAndAmountRange)
	return r.ValidityRange
}

// Set of elements qualifying the interest rate.
type Rate2 struct {

	// Indicates the sign of the rate.
	Sign *PlusOrMinusIndicator `xml:"Sgn,omitempty"`

	// Percentage charged for the use of an amount of money, usually expressed at an annual rate. The interest rate is the ratio of the amount of interest paid during a certain period of time compared to the principal amount of the interest bearing financial instrument.
	Rate *PercentageRate `xml:"Rate"`
}

func (r *Rate2) SetSign(value string) {
	r.Sign = (*PlusOrMinusIndicator)(&value)
}

func (r *Rate2) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

// Set of elements used to qualify the interest rate.
type Rate3 struct {

	// Specifies the type of interest rate.
	Type *RateType4Choice `xml:"Tp"`

	// An amount range where the interest rate is applicable.
	ValidityRange *CurrencyAndAmountRange2 `xml:"VldtyRg,omitempty"`
}

func (r *Rate3) AddType() *RateType4Choice {
	r.Type = new(RateType4Choice)
	return r.Type
}

func (r *Rate3) AddValidityRange() *CurrencyAndAmountRange2 {
	r.ValidityRange = new(CurrencyAndAmountRange2)
	return r.ValidityRange
}
