package iso20022

// Specifies the value expressed as a rate type and a percentage rate.
type RateTypeAndPercentageRate1 struct {

	// Value expressed as a rate type.
	RateType *RateType28Choice `xml:"RateTp"`

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`
}

func (r *RateTypeAndPercentageRate1) AddRateType() *RateType28Choice {
	r.RateType = new(RateType28Choice)
	return r.RateType
}

func (r *RateTypeAndPercentageRate1) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

// Specifies the value expressed as a rate type and a percentage rate.
type RateTypeAndPercentageRate8 struct {

	// Value expressed as a rate type.
	RateType *RateType42Choice `xml:"RateTp"`

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`
}

func (r *RateTypeAndPercentageRate8) AddRateType() *RateType42Choice {
	r.RateType = new(RateType42Choice)
	return r.RateType
}

func (r *RateTypeAndPercentageRate8) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

// Specifies the value expressed as a rate type and a percentage rate.
type RateTypeAndPercentageRate9 struct {

	// Value expressed as a rate type.
	RateType *RateType46Choice `xml:"RateTp"`

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`
}

func (r *RateTypeAndPercentageRate9) AddRateType() *RateType46Choice {
	r.RateType = new(RateType46Choice)
	return r.RateType
}

func (r *RateTypeAndPercentageRate9) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}
