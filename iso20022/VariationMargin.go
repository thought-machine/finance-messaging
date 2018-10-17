package iso20022

// Elements used to calculate the collateral margin call for the variation margin.
type VariationMargin1 struct {

	// Amount of unsecured exposure a counterparty will accept before issuing a margin call in the base currency.
	ThresholdAmount *ActiveCurrencyAndAmount `xml:"ThrshldAmt"`

	// Specifies if the threshold amount is secured or unsecured.
	ThresholdType *ThresholdType1Code `xml:"ThrshldTp,omitempty"`

	// Minimum amount to pay/receive as specified in the agreement in the base currency (to avoid the need to transfer an inconveniently small amount of variation margin).
	MinimumTransferAmount *ActiveCurrencyAndAmount `xml:"MinTrfAmt"`

	// Amount specified to avoid the need to transfer uneven amounts of collateral.
	RoundingAmount *ActiveCurrencyAndAmount `xml:"RndgAmt"`

	// Defines how the rounding amount was applied in the calculation. For example, should the amount of collateral required be rounded up, down, to the closer integral multiple specified or not rounded.
	RoundingMethod *RoundingMethod1Code `xml:"RndgMtd"`
}

func (v *VariationMargin1) SetThresholdAmount(value, currency string) {
	v.ThresholdAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (v *VariationMargin1) SetThresholdType(value string) {
	v.ThresholdType = (*ThresholdType1Code)(&value)
}

func (v *VariationMargin1) SetMinimumTransferAmount(value, currency string) {
	v.MinimumTransferAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (v *VariationMargin1) SetRoundingAmount(value, currency string) {
	v.RoundingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (v *VariationMargin1) SetRoundingMethod(value string) {
	v.RoundingMethod = (*RoundingMethod1Code)(&value)
}

// Margin required to cover the risk because of the price fluctuations occurred on the unsettled exposures towards central counterparty.
type VariationMargin3 struct {

	// Provides details about the security identification.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Margin required to cover the risk because of the price fluctuations occurred on the unsettled exposures towards the central counterparty.
	TotalVariationMargin []*TotalVariationMargin1 `xml:"TtlVartnMrgn"`

	// Net unrealised profit or loss on the value of the netted, gross and failing positions.
	TotalMarkToMarket *Amount2 `xml:"TtlMrkToMkt"`

	// Unrealised net loss calculated at the participant portfolio level.
	MarkToMarketNetted []*Amount2 `xml:"MrkToMktNetd,omitempty"`

	// Unrealised net loss calculated in that market/boundary.
	MarkToMarketGross []*Amount2 `xml:"MrkToMktGrss,omitempty"`

	// Sum of the unrealised loss without taking profit into consideration.
	MarkToMarketFails []*Amount2 `xml:"MrkToMktFls,omitempty"`

	// Haircut applied to the absolute value of the participants net positions. Calculation depends on a participants credit rating.
	FailsHaircut *Amount2 `xml:"FlsHrcut,omitempty"`
}

func (v *VariationMargin3) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	v.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return v.FinancialInstrumentIdentification
}

func (v *VariationMargin3) AddTotalVariationMargin() *TotalVariationMargin1 {
	newValue := new(TotalVariationMargin1)
	v.TotalVariationMargin = append(v.TotalVariationMargin, newValue)
	return newValue
}

func (v *VariationMargin3) AddTotalMarkToMarket() *Amount2 {
	v.TotalMarkToMarket = new(Amount2)
	return v.TotalMarkToMarket
}

func (v *VariationMargin3) AddMarkToMarketNetted() *Amount2 {
	newValue := new(Amount2)
	v.MarkToMarketNetted = append(v.MarkToMarketNetted, newValue)
	return newValue
}

func (v *VariationMargin3) AddMarkToMarketGross() *Amount2 {
	newValue := new(Amount2)
	v.MarkToMarketGross = append(v.MarkToMarketGross, newValue)
	return newValue
}

func (v *VariationMargin3) AddMarkToMarketFails() *Amount2 {
	newValue := new(Amount2)
	v.MarkToMarketFails = append(v.MarkToMarketFails, newValue)
	return newValue
}

func (v *VariationMargin3) AddFailsHaircut() *Amount2 {
	v.FailsHaircut = new(Amount2)
	return v.FailsHaircut
}
