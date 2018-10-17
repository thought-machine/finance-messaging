package iso20022

// Amount of money due to a party as compensation for a service.
type Commission10 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType6Code `xml:"Tp,omitempty"`

	// Service for which the commission is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp,omitempty"`

	// Basis upon which a commission is charged, eg, flat fee.
	Basis *TaxationBasis4Code `xml:"Bsis,omitempty"`

	// Basis upon which a commission is charged, eg, flat fee.
	ExtendedBasis *Extended350Code `xml:"XtndedBsis,omitempty"`

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Party entitled to the amount of money resulting from a commission.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Voluntary non-enforcement of the right to all or part of a commission.
	WaivingDetails *CommissionWaiver3 `xml:"WvgDtls,omitempty"`
}

func (c *Commission10) SetType(value string) {
	c.Type = (*CommissionType6Code)(&value)
}

func (c *Commission10) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Commission10) SetBasis(value string) {
	c.Basis = (*TaxationBasis4Code)(&value)
}

func (c *Commission10) SetExtendedBasis(value string) {
	c.ExtendedBasis = (*Extended350Code)(&value)
}

func (c *Commission10) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission10) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Commission10) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

func (c *Commission10) SetCommercialAgreementReference(value string) {
	c.CommercialAgreementReference = (*Max35Text)(&value)
}

func (c *Commission10) AddWaivingDetails() *CommissionWaiver3 {
	c.WaivingDetails = new(CommissionWaiver3)
	return c.WaivingDetails
}

// Amount of money due to a party as compensation for a service.
type Commission11 struct {

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Service for which the commission is asked or paid.
	Type *CommissionType6Code `xml:"Tp,omitempty"`

	// Service for which the commission is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Indicates if the CommercialAgreementReference is a new reference or not.
	NewCommercialAgreementReferenceIndicator *YesNoIndicator `xml:"NewComrclAgrmtRefInd,omitempty"`
}

func (c *Commission11) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission11) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Commission11) SetType(value string) {
	c.Type = (*CommissionType6Code)(&value)
}

func (c *Commission11) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Commission11) SetCommercialAgreementReference(value string) {
	c.CommercialAgreementReference = (*Max35Text)(&value)
}

func (c *Commission11) SetNewCommercialAgreementReferenceIndicator(value string) {
	c.NewCommercialAgreementReferenceIndicator = (*YesNoIndicator)(&value)
}

// Amount of money due to a party as compensation for a service.
type Commission12 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType7Code `xml:"Tp"`

	// Service for which the commission is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Basis upon which a commission is charged, eg, flat fee.
	Basis *TaxationBasis4Code `xml:"Bsis,omitempty"`

	// Basis upon which a commission is charged, eg, flat fee.
	ExtendedBasis *Extended350Code `xml:"XtndedBsis,omitempty"`

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Party entitled to the amount of money resulting from a commission.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`
}

func (c *Commission12) SetType(value string) {
	c.Type = (*CommissionType7Code)(&value)
}

func (c *Commission12) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Commission12) SetBasis(value string) {
	c.Basis = (*TaxationBasis4Code)(&value)
}

func (c *Commission12) SetExtendedBasis(value string) {
	c.ExtendedBasis = (*Extended350Code)(&value)
}

func (c *Commission12) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission12) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

func (c *Commission12) SetCommercialAgreementReference(value string) {
	c.CommercialAgreementReference = (*Max35Text)(&value)
}

// Amount of money due to a party as compensation for a service.
type Commission16 struct {

	// Specification of the commission type.
	Type *CommissionType2Choice `xml:"Tp"`

	// Amount of money due to a party as compensation for a service.
	Commission *AmountOrRate2Choice `xml:"Comssn"`

	// Information related to an identification, eg, party identification or account identification.
	RecipientIdentification *PartyIdentification54 `xml:"RcptId,omitempty"`

	// Date at which an operation is triggered to calculate, for instance, a commission, fee, asset values, etc.
	CalculationDate *ISODate `xml:"ClctnDt,omitempty"`

	// Total value of the commissions for a specific trade.
	TotalCommission *AmountAndDirection29 `xml:"TtlComssn,omitempty"`

	// Amount that results of the calculation of VAT on net fees, according to the transaction current tariffs.
	TotalVATAmount *ActiveCurrencyAndAmount `xml:"TtlVATAmt,omitempty"`

	// Specifies the VAT rate.
	VATRate *BaseOneRate `xml:"VATRate,omitempty"`
}

func (c *Commission16) AddType() *CommissionType2Choice {
	c.Type = new(CommissionType2Choice)
	return c.Type
}

func (c *Commission16) AddCommission() *AmountOrRate2Choice {
	c.Commission = new(AmountOrRate2Choice)
	return c.Commission
}

func (c *Commission16) AddRecipientIdentification() *PartyIdentification54 {
	c.RecipientIdentification = new(PartyIdentification54)
	return c.RecipientIdentification
}

func (c *Commission16) SetCalculationDate(value string) {
	c.CalculationDate = (*ISODate)(&value)
}

func (c *Commission16) AddTotalCommission() *AmountAndDirection29 {
	c.TotalCommission = new(AmountAndDirection29)
	return c.TotalCommission
}

func (c *Commission16) SetTotalVATAmount(value, currency string) {
	c.TotalVATAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *Commission16) SetVATRate(value string) {
	c.VATRate = (*BaseOneRate)(&value)
}

// Amount of money due to a party as compensation for a service.
type Commission17 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType3Choice `xml:"Tp"`

	// Basis upon which a commission is charged, for example, flat fee.
	Basis *CommissionBasis1Choice `xml:"Bsis"`

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Party entitled to the amount of money resulting from a commission.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Voluntary non-enforcement of the right to part of a commission.
	WaivingDetails *CommissionWaiver4 `xml:"WvgDtls,omitempty"`
}

func (c *Commission17) AddType() *CommissionType3Choice {
	c.Type = new(CommissionType3Choice)
	return c.Type
}

func (c *Commission17) AddBasis() *CommissionBasis1Choice {
	c.Basis = new(CommissionBasis1Choice)
	return c.Basis
}

func (c *Commission17) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission17) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

func (c *Commission17) SetCommercialAgreementReference(value string) {
	c.CommercialAgreementReference = (*Max35Text)(&value)
}

func (c *Commission17) AddWaivingDetails() *CommissionWaiver4 {
	c.WaivingDetails = new(CommissionWaiver4)
	return c.WaivingDetails
}

// Amount of money due to a party as compensation for a service.
type Commission18 struct {

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate"`

	// Additional information about the type of markup.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *Commission18) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Commission18) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}

// Amount of money due to a party as compensation for a service.
type Commission19 struct {

	// Commission expressed as an amount of money.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Additional information about the type of commission.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *Commission19) SetAmount(value, currency string) {
	c.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *Commission19) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}

// Amount of money due to a party as compensation for a service.
type Commission21 struct {

	// Type of commission.
	CommissionType *CommissionType5Choice `xml:"ComssnTp"`

	// Commission amount or commission rate applied.
	CommissionApplied *AmountOrRate3Choice `xml:"ComssnApld"`
}

func (c *Commission21) AddCommissionType() *CommissionType5Choice {
	c.CommissionType = new(CommissionType5Choice)
	return c.CommissionType
}

func (c *Commission21) AddCommissionApplied() *AmountOrRate3Choice {
	c.CommissionApplied = new(AmountOrRate3Choice)
	return c.CommissionApplied
}

// Amount of money due to a party as compensation for a service.
type Commission22 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType3Choice `xml:"Tp"`

	// Basis upon which a commission is charged, for example, flat fee.
	Basis *CommissionBasis1Choice `xml:"Bsis,omitempty"`

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Party entitled to the amount of money resulting from a commission.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Voluntary non-enforcement of the right to part of a commission.
	WaivingDetails *CommissionWaiver4 `xml:"WvgDtls,omitempty"`
}

func (c *Commission22) AddType() *CommissionType3Choice {
	c.Type = new(CommissionType3Choice)
	return c.Type
}

func (c *Commission22) AddBasis() *CommissionBasis1Choice {
	c.Basis = new(CommissionBasis1Choice)
	return c.Basis
}

func (c *Commission22) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission22) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

func (c *Commission22) SetCommercialAgreementReference(value string) {
	c.CommercialAgreementReference = (*Max35Text)(&value)
}

func (c *Commission22) AddWaivingDetails() *CommissionWaiver4 {
	c.WaivingDetails = new(CommissionWaiver4)
	return c.WaivingDetails
}

// Amount of money due to a party as compensation for a service.
type Commission23 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType3Choice `xml:"Tp"`

	// Basis upon which a commission is charged, for example, flat fee.
	Basis *CommissionBasis1Choice `xml:"Bsis,omitempty"`

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Party entitled to the amount of money resulting from a commission.
	RecipientIdentification *PartyIdentification70Choice `xml:"RcptId,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Voluntary non-enforcement of the right to part of a commission.
	WaivingDetails *CommissionWaiver4 `xml:"WvgDtls,omitempty"`
}

func (c *Commission23) AddType() *CommissionType3Choice {
	c.Type = new(CommissionType3Choice)
	return c.Type
}

func (c *Commission23) AddBasis() *CommissionBasis1Choice {
	c.Basis = new(CommissionBasis1Choice)
	return c.Basis
}

func (c *Commission23) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission23) AddRecipientIdentification() *PartyIdentification70Choice {
	c.RecipientIdentification = new(PartyIdentification70Choice)
	return c.RecipientIdentification
}

func (c *Commission23) SetCommercialAgreementReference(value string) {
	c.CommercialAgreementReference = (*Max35Text)(&value)
}

func (c *Commission23) AddWaivingDetails() *CommissionWaiver4 {
	c.WaivingDetails = new(CommissionWaiver4)
	return c.WaivingDetails
}

// Amount of money due to a party as compensation for a service.
type Commission4 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType1 `xml:"Tp"`

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate"`
}

func (c *Commission4) AddType() *CommissionType1 {
	c.Type = new(CommissionType1)
	return c.Type
}

func (c *Commission4) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission4) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

// Amount of money due to a party as compensation for a service.
type Commission6 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType1 `xml:"Tp,omitempty"`

	// Basis upon which a commission is charged, eg, flat fee.
	Basis *TaxationBasis1 `xml:"Bsis,omitempty"`

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Party entitled to the amount of money resulting from a commission.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Voluntary non-enforcement of the right to all or part of a commission.
	WaivingDetails *CommissionWaiver2 `xml:"WvgDtls,omitempty"`
}

func (c *Commission6) AddType() *CommissionType1 {
	c.Type = new(CommissionType1)
	return c.Type
}

func (c *Commission6) AddBasis() *TaxationBasis1 {
	c.Basis = new(TaxationBasis1)
	return c.Basis
}

func (c *Commission6) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission6) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Commission6) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

func (c *Commission6) SetCommercialAgreementReference(value string) {
	c.CommercialAgreementReference = (*Max35Text)(&value)
}

func (c *Commission6) AddWaivingDetails() *CommissionWaiver2 {
	c.WaivingDetails = new(CommissionWaiver2)
	return c.WaivingDetails
}

// Amount of money due to a party as compensation for a service.
type Commission7 struct {

	// Commission expressed as an amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate"`

	// Service for which the commission is asked or paid.
	Type *CommissionType1 `xml:"Tp"`
}

func (c *Commission7) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission7) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Commission7) AddType() *CommissionType1 {
	c.Type = new(CommissionType1)
	return c.Type
}

// Amount of money due to a party as compensation for a service.
type Commission9 struct {

	// Service for which the commission is asked or paid.
	Type *CommissionType6Code `xml:"Tp"`

	// Service for which the commission is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Commission expressed as an amount of money.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Commission expressed as a percentage.
	Rate *PercentageRate `xml:"Rate"`
}

func (c *Commission9) SetType(value string) {
	c.Type = (*CommissionType6Code)(&value)
}

func (c *Commission9) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Commission9) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Commission9) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}
