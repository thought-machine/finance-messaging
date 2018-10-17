package iso20022

// Amount of money associated with a service.
type Fee1 struct {

	// Type of fee (charge/commission).
	Type *ChargeType5Choice `xml:"Tp"`

	// Method used to calculate the fee (charge/commission).
	Basis *ChargeBasis2Choice `xml:"Bsis,omitempty"`

	// Standard fee (charge/commission) amount as specified in the fund prospectus or agreed for the account.
	StandardAmount *ActiveCurrencyAndAmount `xml:"StdAmt,omitempty"`

	// Standard fee (charge/commission) rate used to calculate the amount of the charge or fee, as specified in the fund prospectus or agreed for the account.
	StandardRate *PercentageRate `xml:"StdRate,omitempty"`

	// Discount or waiver applied to the fee (charge/commission).
	DiscountDetails *ChargeOrCommissionDiscount1 `xml:"DscntDtls,omitempty"`

	// Requested fee (charge/commission) amount as agreed for the account.
	RequestedAmount *ActiveCurrencyAndAmount `xml:"ReqdAmt,omitempty"`

	// Requested rate used to calculate the amount of the fee (charge/commission), as agreed for the account.
	RequestedRate *PercentageRate `xml:"ReqdRate,omitempty"`

	// Reference to a sales agreement that overrides normal processing or the Service Level Agreement (SLA), such as a fee (charge/commission).
	NonStandardSLAReference *Max35Text `xml:"NonStdSLARef,omitempty"`

	// Party entitled to the amount of money resulting from a fee (charge/commission).
	RecipientIdentification *PartyIdentification113 `xml:"RcptId,omitempty"`
}

func (f *Fee1) AddType() *ChargeType5Choice {
	f.Type = new(ChargeType5Choice)
	return f.Type
}

func (f *Fee1) AddBasis() *ChargeBasis2Choice {
	f.Basis = new(ChargeBasis2Choice)
	return f.Basis
}

func (f *Fee1) SetStandardAmount(value, currency string) {
	f.StandardAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee1) SetStandardRate(value string) {
	f.StandardRate = (*PercentageRate)(&value)
}

func (f *Fee1) AddDiscountDetails() *ChargeOrCommissionDiscount1 {
	f.DiscountDetails = new(ChargeOrCommissionDiscount1)
	return f.DiscountDetails
}

func (f *Fee1) SetRequestedAmount(value, currency string) {
	f.RequestedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee1) SetRequestedRate(value string) {
	f.RequestedRate = (*PercentageRate)(&value)
}

func (f *Fee1) SetNonStandardSLAReference(value string) {
	f.NonStandardSLAReference = (*Max35Text)(&value)
}

func (f *Fee1) AddRecipientIdentification() *PartyIdentification113 {
	f.RecipientIdentification = new(PartyIdentification113)
	return f.RecipientIdentification
}

// Amount of money associated with a service.
type Fee2 struct {

	// Type of fee (charge/commission).
	Type *ChargeType5Choice `xml:"Tp"`

	// Method used to calculate the fee (charge/commission).
	Basis *ChargeBasis2Choice `xml:"Bsis,omitempty"`

	// Standard fee (charge/commission) amount as specified in the fund prospectus or agreed for the account.
	StandardAmount *ActiveCurrencyAndAmount `xml:"StdAmt,omitempty"`

	// Standard fee (charge/commission) rate used to calculate the amount of the charge or fee, as specified in the fund prospectus or agreed for the account.
	StandardRate *PercentageRate `xml:"StdRate,omitempty"`

	// Discount or waiver applied to the fee (charge/commission).
	DiscountDetails *ChargeOrCommissionDiscount1 `xml:"DscntDtls,omitempty"`

	// Fee (charge/commission) amount applied to the transaction.
	AppliedAmount *ActiveCurrencyAndAmount `xml:"ApldAmt,omitempty"`

	// Final rate used to calculate the fee (charge/commission) amount.
	AppliedRate *PercentageRate `xml:"ApldRate,omitempty"`

	// Reference to a sales agreement that overrides normal processing or the Service Level Agreement (SLA), such as a fee (charge/commission).
	NonStandardSLAReference *Max35Text `xml:"NonStdSLARef,omitempty"`

	// Party entitled to the amount of money resulting from a fee (charge/commission).
	RecipientIdentification *PartyIdentification113 `xml:"RcptId,omitempty"`

	// Indicates the information is provided for information purposes only. When the value is ‘false’ or ‘0’ the amount provided is taken into consideration in the transaction overhead. When the value is ‘true’ or ‘1’ the amount provided is not taken into consideration in the transaction overhead.
	InformativeIndicator *YesNoIndicator `xml:"InftvInd"`
}

func (f *Fee2) AddType() *ChargeType5Choice {
	f.Type = new(ChargeType5Choice)
	return f.Type
}

func (f *Fee2) AddBasis() *ChargeBasis2Choice {
	f.Basis = new(ChargeBasis2Choice)
	return f.Basis
}

func (f *Fee2) SetStandardAmount(value, currency string) {
	f.StandardAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee2) SetStandardRate(value string) {
	f.StandardRate = (*PercentageRate)(&value)
}

func (f *Fee2) AddDiscountDetails() *ChargeOrCommissionDiscount1 {
	f.DiscountDetails = new(ChargeOrCommissionDiscount1)
	return f.DiscountDetails
}

func (f *Fee2) SetAppliedAmount(value, currency string) {
	f.AppliedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee2) SetAppliedRate(value string) {
	f.AppliedRate = (*PercentageRate)(&value)
}

func (f *Fee2) SetNonStandardSLAReference(value string) {
	f.NonStandardSLAReference = (*Max35Text)(&value)
}

func (f *Fee2) AddRecipientIdentification() *PartyIdentification113 {
	f.RecipientIdentification = new(PartyIdentification113)
	return f.RecipientIdentification
}

func (f *Fee2) SetInformativeIndicator(value string) {
	f.InformativeIndicator = (*YesNoIndicator)(&value)
}

// Amount of money associated with a service.
type Fee3 struct {

	// Type of fee (charge/commission).
	Type *ChargeType5Choice `xml:"Tp,omitempty"`

	// Modified value of the standard fee (charge/commission) amount applied on the order (the standard fee (charge/commission) amount in the original individual order that has been repaired so that the order can be accepted).
	RepairedStandardAmount *ActiveCurrencyAndAmount `xml:"RprdStdAmt,omitempty"`

	// Modified value of the standard fee (charge/commission) rate applied on the order (the standard fee (charge/commission) rate in the original individual order that has been repaired so that the order can be accepted).
	RepairedStandardRate *PercentageRate `xml:"RprdStdRate,omitempty"`

	// Modified value of the discount amount applied on the order (the discount amount in the original individual order that has been repaired so that the order can be accepted).
	RepairedDiscountAmount *ActiveCurrencyAndAmount `xml:"RprdDscntAmt,omitempty"`

	// Modified value of the discount rate applied on the order (the discount rate in the original individual order that has been repaired so that the order can be accepted).
	RepairedDiscountRate *PercentageRate `xml:"RprdDscntRate,omitempty"`

	// Modified value of the requested fee (charge/commission) amount applied on the order (the requested fee (charge/commission) amount in the original individual order that has been repaired so that the order can be accepted).
	RepairedRequestedAmount *ActiveCurrencyAndAmount `xml:"RprdReqdAmt,omitempty"`

	// Modified value of the requested  fee (charge/commission) rate applied on the order (the requested fee (charge/commission) rate in the original individual order that has been repaired so that the order can be accepted).
	RepairedRequestedRate *PercentageRate `xml:"RprdReqdRate,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Indicates if the CommercialAgreementReference is a new reference or not.
	NewCommercialAgreementReferenceIndicator *YesNoIndicator `xml:"NewComrclAgrmtRefInd,omitempty"`
}

func (f *Fee3) AddType() *ChargeType5Choice {
	f.Type = new(ChargeType5Choice)
	return f.Type
}

func (f *Fee3) SetRepairedStandardAmount(value, currency string) {
	f.RepairedStandardAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee3) SetRepairedStandardRate(value string) {
	f.RepairedStandardRate = (*PercentageRate)(&value)
}

func (f *Fee3) SetRepairedDiscountAmount(value, currency string) {
	f.RepairedDiscountAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee3) SetRepairedDiscountRate(value string) {
	f.RepairedDiscountRate = (*PercentageRate)(&value)
}

func (f *Fee3) SetRepairedRequestedAmount(value, currency string) {
	f.RepairedRequestedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee3) SetRepairedRequestedRate(value string) {
	f.RepairedRequestedRate = (*PercentageRate)(&value)
}

func (f *Fee3) SetCommercialAgreementReference(value string) {
	f.CommercialAgreementReference = (*Max35Text)(&value)
}

func (f *Fee3) SetNewCommercialAgreementReferenceIndicator(value string) {
	f.NewCommercialAgreementReferenceIndicator = (*YesNoIndicator)(&value)
}
