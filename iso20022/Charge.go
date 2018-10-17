package iso20022

// Amount of money associated with a service.
type Charge10 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType1 `xml:"Tp"`

	// Method used to calculate a charge.
	ChargeBasis *TaxationBasis1 `xml:"ChrgBsis,omitempty"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`
}

func (c *Charge10) AddType() *ChargeType1 {
	c.Type = new(ChargeType1)
	return c.Type
}

func (c *Charge10) AddChargeBasis() *TaxationBasis1 {
	c.ChargeBasis = new(TaxationBasis1)
	return c.ChargeBasis
}

func (c *Charge10) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge10) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Charge10) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

// Amount of money associated with a service.
type Charge11 struct {

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`

	// Type of service for which a charge is asked or paid.
	Type *ChargeType1 `xml:"Tp"`
}

func (c *Charge11) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge11) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Charge11) AddType() *ChargeType1 {
	c.Type = new(ChargeType1)
	return c.Type
}

// Identifies the different types of freight charges associated with goods.
type Charge12 struct {

	// Identifies whether the freight charges associated with the goods are "prepaid" or "collect".
	Type *FreightCharges1Code `xml:"Tp"`

	// Amount of money associated with a service.
	Charges []*ChargesDetails1 `xml:"Chrgs,omitempty"`
}

func (c *Charge12) SetType(value string) {
	c.Type = (*FreightCharges1Code)(&value)
}

func (c *Charge12) AddCharges() *ChargesDetails1 {
	newValue := new(ChargesDetails1)
	c.Charges = append(c.Charges, newValue)
	return newValue
}

// Identifies the different types of freight charges associated with goods.
type Charge13 struct {

	// Identifies whether the freight charges associated with the goods are "prepaid" or "collect".
	Type *FreightCharges1Code `xml:"Tp"`

	// Amount of money associated with a service.
	Charges []*ChargesDetails2 `xml:"Chrgs,omitempty"`
}

func (c *Charge13) SetType(value string) {
	c.Type = (*FreightCharges1Code)(&value)
}

func (c *Charge13) AddCharges() *ChargesDetails2 {
	newValue := new(ChargesDetails2)
	c.Charges = append(c.Charges, newValue)
	return newValue
}

// Amount of money associated with a service.
type Charge15 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType9Code `xml:"Tp"`

	// Type of service for which a charge is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`

	// Calculation basis for the charge or fee.
	CalculationBasis *CalculationBasis2Code `xml:"ClctnBsis,omitempty"`

	// Calculation basis for the charge or fee.
	ExtendedCalculationBasis *Extended350Code `xml:"XtndedClctnBsis,omitempty"`
}

func (c *Charge15) SetType(value string) {
	c.Type = (*ChargeType9Code)(&value)
}

func (c *Charge15) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Charge15) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge15) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Charge15) SetCalculationBasis(value string) {
	c.CalculationBasis = (*CalculationBasis2Code)(&value)
}

func (c *Charge15) SetExtendedCalculationBasis(value string) {
	c.ExtendedCalculationBasis = (*Extended350Code)(&value)
}

// Amount of money associated with a service.
type Charge16 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType10Code `xml:"Tp"`

	// Type of service for which a charge is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`
}

func (c *Charge16) SetType(value string) {
	c.Type = (*ChargeType10Code)(&value)
}

func (c *Charge16) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Charge16) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge16) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

// Amount of money associated with a service.
type Charge17 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType11Code `xml:"Tp"`

	// Type of service for which a charge is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Method used to calculate a charge.
	ChargeBasis *TaxationBasis2Code `xml:"ChrgBsis,omitempty"`

	// Method used to calculate a charge.
	ExtendedChargeBasis *Extended350Code `xml:"XtndedChrgBsis,omitempty"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`
}

func (c *Charge17) SetType(value string) {
	c.Type = (*ChargeType11Code)(&value)
}

func (c *Charge17) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Charge17) SetChargeBasis(value string) {
	c.ChargeBasis = (*TaxationBasis2Code)(&value)
}

func (c *Charge17) SetExtendedChargeBasis(value string) {
	c.ExtendedChargeBasis = (*Extended350Code)(&value)
}

func (c *Charge17) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge17) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Charge17) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

// Amount of money associated with a service.
type Charge18 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType11Code `xml:"Tp"`

	// Type of service for which a charge is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Method used to calculate a charge.
	ChargeBasis *TaxationBasis2Code `xml:"ChrgBsis,omitempty"`

	// Method used to calculate a charge.
	ExtendedChargeBasis *Extended350Code `xml:"XtndedChrgBsis,omitempty"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`
}

func (c *Charge18) SetType(value string) {
	c.Type = (*ChargeType11Code)(&value)
}

func (c *Charge18) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Charge18) SetChargeBasis(value string) {
	c.ChargeBasis = (*TaxationBasis2Code)(&value)
}

func (c *Charge18) SetExtendedChargeBasis(value string) {
	c.ExtendedChargeBasis = (*Extended350Code)(&value)
}

func (c *Charge18) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge18) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Charge18) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

// Amount of money associated with a service.
type Charge19 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType11Code `xml:"Tp"`

	// Type of service for which a charge is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`
}

func (c *Charge19) SetType(value string) {
	c.Type = (*ChargeType11Code)(&value)
}

func (c *Charge19) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Charge19) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge19) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

// Amount of money associated with a service.
type Charge20 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType12Code `xml:"Tp"`

	// Type of service for which a charge is asked or paid.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Method used to calculate a charge.
	ChargeBasis *TaxationBasis2Code `xml:"ChrgBsis,omitempty"`

	// Method used to calculate a charge.
	ExtendedChargeBasis *Extended350Code `xml:"XtndedChrgBsis,omitempty"`

	// Specifies the party that will bear the charges associated with a transfer.
	ChargeBearer *ChargeBearer1Code `xml:"ChrgBr,omitempty"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`
}

func (c *Charge20) SetType(value string) {
	c.Type = (*ChargeType12Code)(&value)
}

func (c *Charge20) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *Charge20) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charge20) SetChargeBasis(value string) {
	c.ChargeBasis = (*TaxationBasis2Code)(&value)
}

func (c *Charge20) SetExtendedChargeBasis(value string) {
	c.ExtendedChargeBasis = (*Extended350Code)(&value)
}

func (c *Charge20) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearer1Code)(&value)
}

func (c *Charge20) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

// Identifies the different types of freight charges associated with goods.
type Charge24 struct {

	// Identifies whether the freight charges associated with the goods are "prepaid" or "collect".
	Type *FreightCharges1Code `xml:"Tp"`

	// Amount of money associated with a service.
	Charges []*ChargesDetails3 `xml:"Chrgs,omitempty"`
}

func (c *Charge24) SetType(value string) {
	c.Type = (*FreightCharges1Code)(&value)
}

func (c *Charge24) AddCharges() *ChargesDetails3 {
	newValue := new(ChargesDetails3)
	c.Charges = append(c.Charges, newValue)
	return newValue
}

// Identifies the different types of freight charges associated with goods.
type Charge25 struct {

	// Identifies whether the freight charges associated with the goods are "prepaid" or "collect".
	Type *FreightCharges1Code `xml:"Tp"`

	// Amount of money associated with a service.
	Charges []*ChargesDetails4 `xml:"Chrgs,omitempty"`
}

func (c *Charge25) SetType(value string) {
	c.Type = (*FreightCharges1Code)(&value)
}

func (c *Charge25) AddCharges() *ChargesDetails4 {
	newValue := new(ChargesDetails4)
	c.Charges = append(c.Charges, newValue)
	return newValue
}

// Choice of formats for the type of charge.
type Charge26 struct {

	// Type of charge.
	Type *ChargeType4Choice `xml:"Tp"`

	// Charge amount or charge rate applied.
	ChargeApplied *AmountOrRate3Choice `xml:"ChrgApld"`
}

func (c *Charge26) AddType() *ChargeType4Choice {
	c.Type = new(ChargeType4Choice)
	return c.Type
}

func (c *Charge26) AddChargeApplied() *AmountOrRate3Choice {
	c.ChargeApplied = new(AmountOrRate3Choice)
	return c.ChargeApplied
}

// Amount of money associated with a service.
type Charge27 struct {

	// Type of charge.
	Type *ChargeType4Choice `xml:"Tp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Method used to calculate the charge.
	ChargeBasis *ChargeBasisType1Choice `xml:"ChrgBsis,omitempty"`

	// Specifies the party that will bear the charges associated with a transfer.
	ChargeBearer *ChargeBearer1Code `xml:"ChrgBr,omitempty"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`
}

func (c *Charge27) AddType() *ChargeType4Choice {
	c.Type = new(ChargeType4Choice)
	return c.Type
}

func (c *Charge27) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charge27) AddChargeBasis() *ChargeBasisType1Choice {
	c.ChargeBasis = new(ChargeBasisType1Choice)
	return c.ChargeBasis
}

func (c *Charge27) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearer1Code)(&value)
}

func (c *Charge27) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

// Amount of money associated with a service.
type Charge29 struct {

	// Type of charge.
	Type *ChargeType4Choice `xml:"Tp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Method used to calculate the charge.
	ChargeBasis *ChargeBasisType1Choice `xml:"ChrgBsis,omitempty"`

	// Specifies the party that will bear the charges associated with a transfer.
	ChargeBearer *ChargeBearer1Code `xml:"ChrgBr,omitempty"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification70Choice `xml:"RcptId,omitempty"`
}

func (c *Charge29) AddType() *ChargeType4Choice {
	c.Type = new(ChargeType4Choice)
	return c.Type
}

func (c *Charge29) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charge29) AddChargeBasis() *ChargeBasisType1Choice {
	c.ChargeBasis = new(ChargeBasisType1Choice)
	return c.ChargeBasis
}

func (c *Charge29) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearer1Code)(&value)
}

func (c *Charge29) AddRecipientIdentification() *PartyIdentification70Choice {
	c.RecipientIdentification = new(PartyIdentification70Choice)
	return c.RecipientIdentification
}

// Amount of money associated with a service.
type Charge4 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeTypeFormat2Choice `xml:"Tp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Method used to calculate a charge.
	ChargeBasis *TaxationBasis2Code `xml:"ChrgBsis,omitempty"`

	// Specifies the party that will bear the charges associated with a transfer.
	ChargeBearer *ChargeBearer1Code `xml:"ChrgBr,omitempty"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification1Choice `xml:"RcptId,omitempty"`
}

func (c *Charge4) AddType() *ChargeTypeFormat2Choice {
	c.Type = new(ChargeTypeFormat2Choice)
	return c.Type
}

func (c *Charge4) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charge4) SetChargeBasis(value string) {
	c.ChargeBasis = (*TaxationBasis2Code)(&value)
}

func (c *Charge4) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearer1Code)(&value)
}

func (c *Charge4) AddRecipientIdentification() *PartyIdentification1Choice {
	c.RecipientIdentification = new(PartyIdentification1Choice)
	return c.RecipientIdentification
}

// Amount of money associated with a service.
type Charge8 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType1 `xml:"Tp"`

	// Method used to calculate a charge.
	ChargeBasis *TaxationBasis1 `xml:"ChrgBsis,omitempty"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`

	// Party entitled to the amount of money resulting from a charge.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`
}

func (c *Charge8) AddType() *ChargeType1 {
	c.Type = new(ChargeType1)
	return c.Type
}

func (c *Charge8) AddChargeBasis() *TaxationBasis1 {
	c.ChargeBasis = new(TaxationBasis1)
	return c.ChargeBasis
}

func (c *Charge8) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge8) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *Charge8) AddRecipientIdentification() *PartyIdentification2Choice {
	c.RecipientIdentification = new(PartyIdentification2Choice)
	return c.RecipientIdentification
}

// Amount of money associated with a service.
type Charge9 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType2 `xml:"Tp"`

	// Amount of money asked or paid for the charge.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Calculation basis for the charge or fee.
	CalculationBasis *CalculationBasis1 `xml:"ClctnBsis,omitempty"`
}

func (c *Charge9) AddType() *ChargeType2 {
	c.Type = new(ChargeType2)
	return c.Type
}

func (c *Charge9) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (c *Charge9) AddCalculationBasis() *CalculationBasis1 {
	c.CalculationBasis = new(CalculationBasis1)
	return c.CalculationBasis
}
