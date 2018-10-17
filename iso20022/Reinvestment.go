package iso20022

// Reinvestment information.
type Reinvestment1 struct {

	// Investment fund for the reinvestment.
	FundDetails *FinancialInstrument29 `xml:"FndDtls"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Percentage of the reinvestment.
	ReinvestmentPercentage *PercentageRate `xml:"RinvstmtPctg"`
}

func (r *Reinvestment1) AddFundDetails() *FinancialInstrument29 {
	r.FundDetails = new(FinancialInstrument29)
	return r.FundDetails
}

func (r *Reinvestment1) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (r *Reinvestment1) SetReinvestmentPercentage(value string) {
	r.ReinvestmentPercentage = (*PercentageRate)(&value)
}

// Reinvestment information.
type Reinvestment2 struct {

	// Investment fund for the reinvestment.
	FinancialInstrumentDetails *FinancialInstrument51 `xml:"FinInstrmDtls"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Percentage of the reinvestment.
	ReinvestmentPercentage *PercentageRate `xml:"RinvstmtPctg"`
}

func (r *Reinvestment2) AddFinancialInstrumentDetails() *FinancialInstrument51 {
	r.FinancialInstrumentDetails = new(FinancialInstrument51)
	return r.FinancialInstrumentDetails
}

func (r *Reinvestment2) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *Reinvestment2) SetReinvestmentPercentage(value string) {
	r.ReinvestmentPercentage = (*PercentageRate)(&value)
}

// Reinvestment information.
type Reinvestment3 struct {

	// Investment fund for the reinvestment.
	FinancialInstrumentDetails *FinancialInstrument56 `xml:"FinInstrmDtls"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Percentage of the reinvestment.
	ReinvestmentPercentage *PercentageRate `xml:"RinvstmtPctg"`
}

func (r *Reinvestment3) AddFinancialInstrumentDetails() *FinancialInstrument56 {
	r.FinancialInstrumentDetails = new(FinancialInstrument56)
	return r.FinancialInstrumentDetails
}

func (r *Reinvestment3) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *Reinvestment3) SetReinvestmentPercentage(value string) {
	r.ReinvestmentPercentage = (*PercentageRate)(&value)
}
