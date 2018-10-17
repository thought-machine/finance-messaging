package iso20022

// Statistical data related to the price change of a security.
type StatisticsByUserDefinedTimePeriod1 struct {

	// Reference period for the valuation.
	Period *DateTimePeriodDetails `xml:"Prd"`

	// Highest price for the referenced period.
	HighestPriceValue *PriceValue1 `xml:"HghstPricVal,omitempty"`

	// Lowest price for the referenced period.
	LowestPriceValue *PriceValue1 `xml:"LwstPricVal,omitempty"`

	// Change in price since the last valuation.
	PriceChange *PriceValue2 `xml:"PricChng,omitempty"`

	// Rate of income from the financial instrument, usually calculated as total dividends or coupon interest available to investors in the last year,divided by the current price.
	Yield *PercentageRate `xml:"Yld,omitempty"`
}

func (s *StatisticsByUserDefinedTimePeriod1) AddPeriod() *DateTimePeriodDetails {
	s.Period = new(DateTimePeriodDetails)
	return s.Period
}

func (s *StatisticsByUserDefinedTimePeriod1) AddHighestPriceValue() *PriceValue1 {
	s.HighestPriceValue = new(PriceValue1)
	return s.HighestPriceValue
}

func (s *StatisticsByUserDefinedTimePeriod1) AddLowestPriceValue() *PriceValue1 {
	s.LowestPriceValue = new(PriceValue1)
	return s.LowestPriceValue
}

func (s *StatisticsByUserDefinedTimePeriod1) AddPriceChange() *PriceValue2 {
	s.PriceChange = new(PriceValue2)
	return s.PriceChange
}

func (s *StatisticsByUserDefinedTimePeriod1) SetYield(value string) {
	s.Yield = (*PercentageRate)(&value)
}

// Statistical data related to the price change of a security.
type StatisticsByUserDefinedTimePeriod2 struct {

	// Reference period for the valuation.
	Period *DateOrDateTimePeriodChoice `xml:"Prd"`

	// Highest price for the referenced period.
	HighestPriceValue *PriceValue5 `xml:"HghstPricVal,omitempty"`

	// Lowest price for the referenced period.
	LowestPriceValue *PriceValue5 `xml:"LwstPricVal,omitempty"`

	// Change in price since the previous valuation date.
	PriceChange *PriceValueChange1 `xml:"PricChng,omitempty"`

	// Rate of income from the financial instrument, usually calculated as total dividends or coupon interest available to investors in the last year,divided by the current price.
	Yield *PercentageRate `xml:"Yld,omitempty"`
}

func (s *StatisticsByUserDefinedTimePeriod2) AddPeriod() *DateOrDateTimePeriodChoice {
	s.Period = new(DateOrDateTimePeriodChoice)
	return s.Period
}

func (s *StatisticsByUserDefinedTimePeriod2) AddHighestPriceValue() *PriceValue5 {
	s.HighestPriceValue = new(PriceValue5)
	return s.HighestPriceValue
}

func (s *StatisticsByUserDefinedTimePeriod2) AddLowestPriceValue() *PriceValue5 {
	s.LowestPriceValue = new(PriceValue5)
	return s.LowestPriceValue
}

func (s *StatisticsByUserDefinedTimePeriod2) AddPriceChange() *PriceValueChange1 {
	s.PriceChange = new(PriceValueChange1)
	return s.PriceChange
}

func (s *StatisticsByUserDefinedTimePeriod2) SetYield(value string) {
	s.Yield = (*PercentageRate)(&value)
}
