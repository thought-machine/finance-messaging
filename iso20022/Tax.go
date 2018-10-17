package iso20022

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type Tax12 struct {

	// Type of tax applied.
	Type *TaxType9Code `xml:"Tp"`

	// Specifies types of tax not present in a code list.
	OtherTaxType *Max35Text `xml:"OthrTaxTp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (t *Tax12) SetType(value string) {
	t.Type = (*TaxType9Code)(&value)
}

func (t *Tax12) SetOtherTaxType(value string) {
	t.OtherTaxType = (*Max35Text)(&value)
}

func (t *Tax12) SetAmount(value, currency string) {
	t.Amount = NewCurrencyAndAmount(value, currency)
}

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type Tax13 struct {

	// Type of tax applied.
	Type *TaxType9Code `xml:"Tp"`

	// Specifies types of tax not present in a code list.
	OtherTaxType *Max35Text `xml:"OthrTaxTp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate"`
}

func (t *Tax13) SetType(value string) {
	t.Type = (*TaxType9Code)(&value)
}

func (t *Tax13) SetOtherTaxType(value string) {
	t.OtherTaxType = (*Max35Text)(&value)
}

func (t *Tax13) SetAmount(value, currency string) {
	t.Amount = NewCurrencyAndAmount(value, currency)
}

func (t *Tax13) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

// Tax related to an investment fund order.
type Tax14 struct {

	// Type of tax applied.
	Type *TaxType11Code `xml:"Tp"`

	// Type of tax applied.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Country where the tax is due.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *TaxExemptReason1Code `xml:"XmptnRsn,omitempty"`

	// Reason for a tax exemption.
	ExtendedExemptionReason *Extended350Code `xml:"XtndedXmptnRsn,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation6 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax14) SetType(value string) {
	t.Type = (*TaxType11Code)(&value)
}

func (t *Tax14) SetExtendedType(value string) {
	t.ExtendedType = (*Extended350Code)(&value)
}

func (t *Tax14) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax14) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *Tax14) SetCountry(value string) {
	t.Country = (*CountryCode)(&value)
}

func (t *Tax14) AddRecipientIdentification() *PartyIdentification2Choice {
	t.RecipientIdentification = new(PartyIdentification2Choice)
	return t.RecipientIdentification
}

func (t *Tax14) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax14) SetExemptionReason(value string) {
	t.ExemptionReason = (*TaxExemptReason1Code)(&value)
}

func (t *Tax14) SetExtendedExemptionReason(value string) {
	t.ExtendedExemptionReason = (*Extended350Code)(&value)
}

func (t *Tax14) AddTaxCalculationDetails() *TaxCalculationInformation6 {
	t.TaxCalculationDetails = new(TaxCalculationInformation6)
	return t.TaxCalculationDetails
}

// Tax related to an investment fund order.
type Tax15 struct {

	// Type of tax applied.
	Type *TaxType13Code `xml:"Tp"`

	// Type of tax applied.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	Basis *TaxationBasis2Code `xml:"Bsis,omitempty"`

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	ExtendedBasis *Extended350Code `xml:"XtndedBsis,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *TaxExemptReason1Code `xml:"XmptnRsn,omitempty"`

	// Reason for a tax exemption.
	ExtendedExemptionReason *Extended350Code `xml:"XtndedXmptnRsn,omitempty"`
}

func (t *Tax15) SetType(value string) {
	t.Type = (*TaxType13Code)(&value)
}

func (t *Tax15) SetExtendedType(value string) {
	t.ExtendedType = (*Extended350Code)(&value)
}

func (t *Tax15) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax15) SetBasis(value string) {
	t.Basis = (*TaxationBasis2Code)(&value)
}

func (t *Tax15) SetExtendedBasis(value string) {
	t.ExtendedBasis = (*Extended350Code)(&value)
}

func (t *Tax15) AddRecipientIdentification() *PartyIdentification2Choice {
	t.RecipientIdentification = new(PartyIdentification2Choice)
	return t.RecipientIdentification
}

func (t *Tax15) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax15) SetExemptionReason(value string) {
	t.ExemptionReason = (*TaxExemptReason1Code)(&value)
}

func (t *Tax15) SetExtendedExemptionReason(value string) {
	t.ExtendedExemptionReason = (*Extended350Code)(&value)
}

// Tax related to an investment fund order.
type Tax16 struct {

	// Type of tax applied.
	Type *TaxType10Code `xml:"Tp"`

	// Type of tax applied.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Country where the tax is due.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *TaxExemptReason1Code `xml:"XmptnRsn,omitempty"`

	// Reason for a tax exemption.
	ExtendedExemptionReason *Extended350Code `xml:"XtndedXmptnRsn,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation5 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax16) SetType(value string) {
	t.Type = (*TaxType10Code)(&value)
}

func (t *Tax16) SetExtendedType(value string) {
	t.ExtendedType = (*Extended350Code)(&value)
}

func (t *Tax16) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax16) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *Tax16) SetCountry(value string) {
	t.Country = (*CountryCode)(&value)
}

func (t *Tax16) AddRecipientIdentification() *PartyIdentification2Choice {
	t.RecipientIdentification = new(PartyIdentification2Choice)
	return t.RecipientIdentification
}

func (t *Tax16) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax16) SetExemptionReason(value string) {
	t.ExemptionReason = (*TaxExemptReason1Code)(&value)
}

func (t *Tax16) SetExtendedExemptionReason(value string) {
	t.ExtendedExemptionReason = (*Extended350Code)(&value)
}

func (t *Tax16) AddTaxCalculationDetails() *TaxCalculationInformation5 {
	t.TaxCalculationDetails = new(TaxCalculationInformation5)
	return t.TaxCalculationDetails
}

// Tax related to an investment fund order.
type Tax17 struct {

	// Type of tax applied.
	Type *TaxType12Code `xml:"Tp"`

	// Type of tax applied.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Amount of money resulting from the calculation of the tax.
	Amount []*ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Country where the tax is due.
	Country *CountryCode `xml:"Ctry"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation4 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax17) SetType(value string) {
	t.Type = (*TaxType12Code)(&value)
}

func (t *Tax17) SetExtendedType(value string) {
	t.ExtendedType = (*Extended350Code)(&value)
}

func (t *Tax17) AddAmount(value, currency string) {
	t.Amount = append(t.Amount, NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency))
}

func (t *Tax17) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *Tax17) SetCountry(value string) {
	t.Country = (*CountryCode)(&value)
}

func (t *Tax17) AddTaxCalculationDetails() *TaxCalculationInformation4 {
	t.TaxCalculationDetails = new(TaxCalculationInformation4)
	return t.TaxCalculationDetails
}

// Tax related to an investment fund order.
type Tax21 struct {

	// Type of tax.
	Type *TaxType1Choice `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Basis used to determine the capital gain or loss, for example, the purchase price.
	Basis *TaxBasis1Choice `xml:"Bsis"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *ExemptionReason1Choice `xml:"XmptnRsn,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation8 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax21) AddType() *TaxType1Choice {
	t.Type = new(TaxType1Choice)
	return t.Type
}

func (t *Tax21) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax21) AddBasis() *TaxBasis1Choice {
	t.Basis = new(TaxBasis1Choice)
	return t.Basis
}

func (t *Tax21) AddRecipientIdentification() *PartyIdentification2Choice {
	t.RecipientIdentification = new(PartyIdentification2Choice)
	return t.RecipientIdentification
}

func (t *Tax21) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax21) AddExemptionReason() *ExemptionReason1Choice {
	t.ExemptionReason = new(ExemptionReason1Choice)
	return t.ExemptionReason
}

func (t *Tax21) AddTaxCalculationDetails() *TaxCalculationInformation8 {
	t.TaxCalculationDetails = new(TaxCalculationInformation8)
	return t.TaxCalculationDetails
}

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type Tax22 struct {

	// Specifies the type of tax.
	Type *TaxType2Choice `xml:"Tp"`

	// Specifies the tax as an amount.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (t *Tax22) AddType() *TaxType2Choice {
	t.Type = new(TaxType2Choice)
	return t.Type
}

func (t *Tax22) SetAmount(value, currency string) {
	t.Amount = NewCurrencyAndAmount(value, currency)
}

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type Tax23 struct {

	// Specifies the type of tax.
	Type *TaxType2Choice `xml:"Tp"`

	// Specifies the tax as an amount or percentage.
	AmountOrPercentage *AmountOrPercentage2Choice `xml:"AmtOrPctg"`
}

func (t *Tax23) AddType() *TaxType2Choice {
	t.Type = new(TaxType2Choice)
	return t.Type
}

func (t *Tax23) AddAmountOrPercentage() *AmountOrPercentage2Choice {
	t.AmountOrPercentage = new(AmountOrPercentage2Choice)
	return t.AmountOrPercentage
}

// Tax related to an investment fund order.
type Tax25 struct {

	// Type of tax.
	Type *TaxType1Choice `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Basis used to determine the capital gain or loss, for example, the purchase price.
	Basis *TaxBasis1Choice `xml:"Bsis,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *ExemptionReason1Choice `xml:"XmptnRsn,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation8 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax25) AddType() *TaxType1Choice {
	t.Type = new(TaxType1Choice)
	return t.Type
}

func (t *Tax25) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax25) AddBasis() *TaxBasis1Choice {
	t.Basis = new(TaxBasis1Choice)
	return t.Basis
}

func (t *Tax25) AddRecipientIdentification() *PartyIdentification2Choice {
	t.RecipientIdentification = new(PartyIdentification2Choice)
	return t.RecipientIdentification
}

func (t *Tax25) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax25) AddExemptionReason() *ExemptionReason1Choice {
	t.ExemptionReason = new(ExemptionReason1Choice)
	return t.ExemptionReason
}

func (t *Tax25) AddTaxCalculationDetails() *TaxCalculationInformation8 {
	t.TaxCalculationDetails = new(TaxCalculationInformation8)
	return t.TaxCalculationDetails
}

// Tax related to an investment fund order.
type Tax28 struct {

	// Type of tax.
	Type *TaxType1Choice `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Basis used to determine the capital gain or loss, for example, the purchase price.
	Basis *TaxBasis1Choice `xml:"Bsis,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification70Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *ExemptionReason1Choice `xml:"XmptnRsn,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation8 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax28) AddType() *TaxType1Choice {
	t.Type = new(TaxType1Choice)
	return t.Type
}

func (t *Tax28) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax28) AddBasis() *TaxBasis1Choice {
	t.Basis = new(TaxBasis1Choice)
	return t.Basis
}

func (t *Tax28) AddRecipientIdentification() *PartyIdentification70Choice {
	t.RecipientIdentification = new(PartyIdentification70Choice)
	return t.RecipientIdentification
}

func (t *Tax28) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax28) AddExemptionReason() *ExemptionReason1Choice {
	t.ExemptionReason = new(ExemptionReason1Choice)
	return t.ExemptionReason
}

func (t *Tax28) AddTaxCalculationDetails() *TaxCalculationInformation8 {
	t.TaxCalculationDetails = new(TaxCalculationInformation8)
	return t.TaxCalculationDetails
}

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type Tax3 struct {

	// Type of tax applied.
	Type *TaxTypeFormat2Choice `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Basis used to determine the capital gain or loss, eg, the purchase price.
	Basis *TaxationBasis2Code `xml:"Bsis,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification1Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *TaxExemptionReasonFormatChoice `xml:"XmptnRsn,omitempty"`
}

func (t *Tax3) AddType() *TaxTypeFormat2Choice {
	t.Type = new(TaxTypeFormat2Choice)
	return t.Type
}

func (t *Tax3) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax3) SetBasis(value string) {
	t.Basis = (*TaxationBasis2Code)(&value)
}

func (t *Tax3) AddRecipientIdentification() *PartyIdentification1Choice {
	t.RecipientIdentification = new(PartyIdentification1Choice)
	return t.RecipientIdentification
}

func (t *Tax3) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax3) AddExemptionReason() *TaxExemptionReasonFormatChoice {
	t.ExemptionReason = new(TaxExemptionReasonFormatChoice)
	return t.ExemptionReason
}

// Tax related to an investment fund order.
type Tax30 struct {

	// Type of tax.
	Type *TaxType3Choice `xml:"Tp"`

	// Tax to be applied.
	Tax *TaxAmountOrRate4Choice `xml:"Tax,omitempty"`

	// Country where the tax is due.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for the tax exemption.
	ExemptionReason *ExemptionReason1Choice `xml:"XmptnRsn,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification113 `xml:"RcptId,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation9 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax30) AddType() *TaxType3Choice {
	t.Type = new(TaxType3Choice)
	return t.Type
}

func (t *Tax30) AddTax() *TaxAmountOrRate4Choice {
	t.Tax = new(TaxAmountOrRate4Choice)
	return t.Tax
}

func (t *Tax30) SetCountry(value string) {
	t.Country = (*CountryCode)(&value)
}

func (t *Tax30) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax30) AddExemptionReason() *ExemptionReason1Choice {
	t.ExemptionReason = new(ExemptionReason1Choice)
	return t.ExemptionReason
}

func (t *Tax30) AddRecipientIdentification() *PartyIdentification113 {
	t.RecipientIdentification = new(PartyIdentification113)
	return t.RecipientIdentification
}

func (t *Tax30) AddTaxCalculationDetails() *TaxCalculationInformation9 {
	t.TaxCalculationDetails = new(TaxCalculationInformation9)
	return t.TaxCalculationDetails
}

// Tax related to an investment fund order.
type Tax31 struct {

	// Type of tax.
	Type *TaxType3Choice `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax.
	AppliedAmount *ActiveCurrencyAndAmount `xml:"ApldAmt"`

	// Rate used to calculate the tax.
	AppliedRate *PercentageRate `xml:"ApldRate,omitempty"`

	// Country where the tax is due.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification113 `xml:"RcptId,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation10 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax31) AddType() *TaxType3Choice {
	t.Type = new(TaxType3Choice)
	return t.Type
}

func (t *Tax31) SetAppliedAmount(value, currency string) {
	t.AppliedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *Tax31) SetAppliedRate(value string) {
	t.AppliedRate = (*PercentageRate)(&value)
}

func (t *Tax31) SetCountry(value string) {
	t.Country = (*CountryCode)(&value)
}

func (t *Tax31) AddRecipientIdentification() *PartyIdentification113 {
	t.RecipientIdentification = new(PartyIdentification113)
	return t.RecipientIdentification
}

func (t *Tax31) AddTaxCalculationDetails() *TaxCalculationInformation10 {
	t.TaxCalculationDetails = new(TaxCalculationInformation10)
	return t.TaxCalculationDetails
}

// Tax related to an investment fund order.
type Tax32 struct {

	// Type of tax applied.
	Type *TaxType3Choice `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax. This amount is provided for information only.
	InformativeAmount *ActiveCurrencyAndAmount `xml:"InftvAmt,omitempty"`

	// Rate used to calculate the tax. This rate is provided for information only.
	InformativeRate *PercentageRate `xml:"InftvRate,omitempty"`

	// Country where the tax is due.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for the tax exemption.
	ExemptionReason *ExemptionReason1Choice `xml:"XmptnRsn,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification113 `xml:"RcptId,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation10 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax32) AddType() *TaxType3Choice {
	t.Type = new(TaxType3Choice)
	return t.Type
}

func (t *Tax32) SetInformativeAmount(value, currency string) {
	t.InformativeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *Tax32) SetInformativeRate(value string) {
	t.InformativeRate = (*PercentageRate)(&value)
}

func (t *Tax32) SetCountry(value string) {
	t.Country = (*CountryCode)(&value)
}

func (t *Tax32) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax32) AddExemptionReason() *ExemptionReason1Choice {
	t.ExemptionReason = new(ExemptionReason1Choice)
	return t.ExemptionReason
}

func (t *Tax32) AddRecipientIdentification() *PartyIdentification113 {
	t.RecipientIdentification = new(PartyIdentification113)
	return t.RecipientIdentification
}

func (t *Tax32) AddTaxCalculationDetails() *TaxCalculationInformation10 {
	t.TaxCalculationDetails = new(TaxCalculationInformation10)
	return t.TaxCalculationDetails
}

// Tax related to an investment fund order.
type Tax6 struct {

	// Type of tax applied.
	Type *TaxType1 `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *TaxExemptionReason1 `xml:"XmptnRsn,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation1 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax6) AddType() *TaxType1 {
	t.Type = new(TaxType1)
	return t.Type
}

func (t *Tax6) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax6) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *Tax6) AddRecipientIdentification() *PartyIdentification2Choice {
	t.RecipientIdentification = new(PartyIdentification2Choice)
	return t.RecipientIdentification
}

func (t *Tax6) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax6) AddExemptionReason() *TaxExemptionReason1 {
	t.ExemptionReason = new(TaxExemptionReason1)
	return t.ExemptionReason
}

func (t *Tax6) AddTaxCalculationDetails() *TaxCalculationInformation1 {
	t.TaxCalculationDetails = new(TaxCalculationInformation1)
	return t.TaxCalculationDetails
}

// Tax related to an investment fund order.
type Tax7 struct {

	// Type of tax applied.
	Type *TaxType2 `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Party that receives the tax. The recipient of, and the party entitled to, the tax may be two different parties.
	RecipientIdentification *PartyIdentification2Choice `xml:"RcptId,omitempty"`

	// Indicates whether a tax exemption applies.
	ExemptionIndicator *YesNoIndicator `xml:"XmptnInd"`

	// Reason for a tax exemption.
	ExemptionReason *TaxExemptionReason1 `xml:"XmptnRsn,omitempty"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation3 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax7) AddType() *TaxType2 {
	t.Type = new(TaxType2)
	return t.Type
}

func (t *Tax7) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax7) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *Tax7) AddRecipientIdentification() *PartyIdentification2Choice {
	t.RecipientIdentification = new(PartyIdentification2Choice)
	return t.RecipientIdentification
}

func (t *Tax7) SetExemptionIndicator(value string) {
	t.ExemptionIndicator = (*YesNoIndicator)(&value)
}

func (t *Tax7) AddExemptionReason() *TaxExemptionReason1 {
	t.ExemptionReason = new(TaxExemptionReason1)
	return t.ExemptionReason
}

func (t *Tax7) AddTaxCalculationDetails() *TaxCalculationInformation3 {
	t.TaxCalculationDetails = new(TaxCalculationInformation3)
	return t.TaxCalculationDetails
}

// Tax related to an investment fund order.
type Tax8 struct {

	// Type of tax applied.
	Type *TaxType3 `xml:"Tp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Country where the tax is due.
	Country *CountryCode `xml:"Ctry"`

	// Information used to calculate the tax.
	TaxCalculationDetails *TaxCalculationInformation2 `xml:"TaxClctnDtls,omitempty"`
}

func (t *Tax8) AddType() *TaxType3 {
	t.Type = new(TaxType3)
	return t.Type
}

func (t *Tax8) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Tax8) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *Tax8) SetCountry(value string) {
	t.Country = (*CountryCode)(&value)
}

func (t *Tax8) AddTaxCalculationDetails() *TaxCalculationInformation2 {
	t.TaxCalculationDetails = new(TaxCalculationInformation2)
	return t.TaxCalculationDetails
}
