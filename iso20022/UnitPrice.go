package iso20022

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice1 struct {

	// Type and information about a price.
	Type *TypeOfPrice2Code `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`
}

func (u *UnitPrice1) SetType(value string) {
	u.Type = (*TypeOfPrice2Code)(&value)
}

func (u *UnitPrice1) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice1) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice10 struct {

	// Type and information about a price.
	Type *TypeOfPrice10Code `xml:"Tp"`

	// Type and information about a price.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *ActiveCurrencyAnd13DecimalAmount `xml:"TaxblIncmPerShr,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculated2Code `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	ExtendedTaxableIncomePerShareCalculated *Extended350Code `xml:"XtndedTaxblIncmPerShrClctd,omitempty"`

	// Specifies the reason why the price is different from the current market price.
	PriceDifferenceReason *Max350Text `xml:"PricDiffRsn,omitempty"`
}

func (u *UnitPrice10) SetType(value string) {
	u.Type = (*TypeOfPrice10Code)(&value)
}

func (u *UnitPrice10) SetExtendedType(value string) {
	u.ExtendedType = (*Extended350Code)(&value)
}

func (u *UnitPrice10) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice10) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice10) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice10) SetTaxableIncomePerShare(value, currency string) {
	u.TaxableIncomePerShare = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (u *UnitPrice10) SetTaxableIncomePerShareCalculated(value string) {
	u.TaxableIncomePerShareCalculated = (*TaxableIncomePerShareCalculated2Code)(&value)
}

func (u *UnitPrice10) SetExtendedTaxableIncomePerShareCalculated(value string) {
	u.ExtendedTaxableIncomePerShareCalculated = (*Extended350Code)(&value)
}

func (u *UnitPrice10) SetPriceDifferenceReason(value string) {
	u.PriceDifferenceReason = (*Max350Text)(&value)
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice11 struct {

	// Type and information about a price.
	Type *TypeOfPrice10Code `xml:"Tp"`

	// Type and information about a price.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`
}

func (u *UnitPrice11) SetType(value string) {
	u.Type = (*TypeOfPrice10Code)(&value)
}

func (u *UnitPrice11) SetExtendedType(value string) {
	u.ExtendedType = (*Extended350Code)(&value)
}

func (u *UnitPrice11) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice11) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice12 struct {

	// Type and information about a price.
	Type *TypeOfPrice12Code `xml:"Tp"`

	// Type and information about a price.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Interest that has accumulated between the most recent payment of interest and the sale of the financial instrument.
	AccruedInterestNAV *ActiveOrHistoricCurrencyAndAmount `xml:"AcrdIntrstNAV,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *ActiveCurrencyAnd13DecimalAmount `xml:"TaxblIncmPerShr,omitempty"`
}

func (u *UnitPrice12) SetType(value string) {
	u.Type = (*TypeOfPrice12Code)(&value)
}

func (u *UnitPrice12) SetExtendedType(value string) {
	u.ExtendedType = (*Extended350Code)(&value)
}

func (u *UnitPrice12) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice12) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice12) SetAccruedInterestNAV(value, currency string) {
	u.AccruedInterestNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (u *UnitPrice12) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice12) SetTaxableIncomePerShare(value, currency string) {
	u.TaxableIncomePerShare = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice15 struct {

	// Type and information about a price.
	Type *TypeOfPrice9Code `xml:"Tp"`

	// Type and information about a price.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Value of the price, eg, as a currency and value.
	ValueInInvestmentCurrency []*PriceValue1 `xml:"ValInInvstmtCcy"`

	// Value of the price, eg, as a currency and value.
	ValueInAlternativeCurrency []*PriceValue1 `xml:"ValInAltrntvCcy,omitempty"`

	// Indicates whether the price information can be used for the execution of a transaction.
	ForExecutionIndicator *YesNoIndicator `xml:"ForExctnInd"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Ratio applied on the non-adjusted price.
	CalculationBasis *PercentageRate `xml:"ClctnBsis,omitempty"`

	// Indicates whether the price is an estimated price.
	EstimatedPriceIndicator *YesNoIndicator `xml:"EstmtdPricInd"`

	// Specifies the number of days from trade date that the counterparty on the other side of the trade should "given up" or divulged.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"TaxblIncmPerShr,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculated2Code `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	ExtendedTaxableIncomePerShareCalculated *Extended350Code `xml:"XtndedTaxblIncmPerShrClctd,omitempty"`

	// Amount included in the dividend that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerDividend *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"TaxblIncmPerDvdd,omitempty"`

	// Specifies whether dividend is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUDividendStatus *EUDividendStatus1Code `xml:"EUDvddSts,omitempty"`

	// Specifies whether dividend is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	ExtendedEUDividendStatus *Extended350Code `xml:"XtndedEUDvddSts,omitempty"`

	// Amount of money associated with a service.
	ChargeDetails []*Charge15 `xml:"ChrgDtls,omitempty"`

	// Information related to taxes that are due.
	TaxLiabilityDetails []*Tax17 `xml:"TaxLbltyDtls,omitempty"`

	// Information related to taxes that are paid back.
	TaxRefundDetails []*Tax17 `xml:"TaxRfndDtls,omitempty"`
}

func (u *UnitPrice15) SetType(value string) {
	u.Type = (*TypeOfPrice9Code)(&value)
}

func (u *UnitPrice15) SetExtendedType(value string) {
	u.ExtendedType = (*Extended350Code)(&value)
}

func (u *UnitPrice15) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice15) AddValueInInvestmentCurrency() *PriceValue1 {
	newValue := new(PriceValue1)
	u.ValueInInvestmentCurrency = append(u.ValueInInvestmentCurrency, newValue)
	return newValue
}

func (u *UnitPrice15) AddValueInAlternativeCurrency() *PriceValue1 {
	newValue := new(PriceValue1)
	u.ValueInAlternativeCurrency = append(u.ValueInAlternativeCurrency, newValue)
	return newValue
}

func (u *UnitPrice15) SetForExecutionIndicator(value string) {
	u.ForExecutionIndicator = (*YesNoIndicator)(&value)
}

func (u *UnitPrice15) SetCumDividendIndicator(value string) {
	u.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (u *UnitPrice15) SetCalculationBasis(value string) {
	u.CalculationBasis = (*PercentageRate)(&value)
}

func (u *UnitPrice15) SetEstimatedPriceIndicator(value string) {
	u.EstimatedPriceIndicator = (*YesNoIndicator)(&value)
}

func (u *UnitPrice15) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice15) SetTaxableIncomePerShare(value, currency string) {
	u.TaxableIncomePerShare = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (u *UnitPrice15) SetTaxableIncomePerShareCalculated(value string) {
	u.TaxableIncomePerShareCalculated = (*TaxableIncomePerShareCalculated2Code)(&value)
}

func (u *UnitPrice15) SetExtendedTaxableIncomePerShareCalculated(value string) {
	u.ExtendedTaxableIncomePerShareCalculated = (*Extended350Code)(&value)
}

func (u *UnitPrice15) SetTaxableIncomePerDividend(value, currency string) {
	u.TaxableIncomePerDividend = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (u *UnitPrice15) SetEUDividendStatus(value string) {
	u.EUDividendStatus = (*EUDividendStatus1Code)(&value)
}

func (u *UnitPrice15) SetExtendedEUDividendStatus(value string) {
	u.ExtendedEUDividendStatus = (*Extended350Code)(&value)
}

func (u *UnitPrice15) AddChargeDetails() *Charge15 {
	newValue := new(Charge15)
	u.ChargeDetails = append(u.ChargeDetails, newValue)
	return newValue
}

func (u *UnitPrice15) AddTaxLiabilityDetails() *Tax17 {
	newValue := new(Tax17)
	u.TaxLiabilityDetails = append(u.TaxLiabilityDetails, newValue)
	return newValue
}

func (u *UnitPrice15) AddTaxRefundDetails() *Tax17 {
	newValue := new(Tax17)
	u.TaxRefundDetails = append(u.TaxRefundDetails, newValue)
	return newValue
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice18 struct {

	// Specifies the unit price.
	UnitPrice *UnitOfMeasure3Choice `xml:"UnitPric"`

	// Price expressed as a currency and value.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Multiplication factor of measurement values. For example: goods that can be ordered by 36 pieces.
	Factor *Max15NumericText `xml:"Fctr,omitempty"`
}

func (u *UnitPrice18) AddUnitPrice() *UnitOfMeasure3Choice {
	u.UnitPrice = new(UnitOfMeasure3Choice)
	return u.UnitPrice
}

func (u *UnitPrice18) SetAmount(value, currency string) {
	u.Amount = NewCurrencyAndAmount(value, currency)
}

func (u *UnitPrice18) SetFactor(value string) {
	u.Factor = (*Max15NumericText)(&value)
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice19 struct {

	// Type of price.
	PriceType *UnitPriceType2Choice `xml:"PricTp"`

	// Value of the price, that is, as a currency and value.
	Value *PriceValue1 `xml:"Val"`
}

func (u *UnitPrice19) AddPriceType() *UnitPriceType2Choice {
	u.PriceType = new(UnitPriceType2Choice)
	return u.PriceType
}

func (u *UnitPrice19) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice20 struct {

	// Type of price.
	PriceType *UnitPriceType2Choice `xml:"PricTp"`

	// Value of the price, that is, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`
}

func (u *UnitPrice20) AddPriceType() *UnitPriceType2Choice {
	u.PriceType = new(UnitPriceType2Choice)
	return u.PriceType
}

func (u *UnitPrice20) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice20) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice21 struct {

	// Type and information about a price.
	Type *TypeOfPrice31Choice `xml:"Tp"`

	// Value of the price, for example, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Interest that has accumulated between the most recent payment of interest and the sale of the financial instrument.
	AccruedInterestNAV *ActiveOrHistoricCurrencyAndAmount `xml:"AcrdIntrstNAV,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *ActiveCurrencyAnd13DecimalAmount `xml:"TaxblIncmPerShr,omitempty"`
}

func (u *UnitPrice21) AddType() *TypeOfPrice31Choice {
	u.Type = new(TypeOfPrice31Choice)
	return u.Type
}

func (u *UnitPrice21) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice21) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice21) SetAccruedInterestNAV(value, currency string) {
	u.AccruedInterestNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (u *UnitPrice21) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice21) SetTaxableIncomePerShare(value, currency string) {
	u.TaxableIncomePerShare = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice22 struct {

	// Type and information about a price.
	Type *TypeOfPrice46Choice `xml:"Tp"`

	// Value of the price, for example, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *ActiveCurrencyAnd13DecimalAmount `xml:"TaxblIncmPerShr,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculated2Choice `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Reason why the price is different from the current market price.
	PriceDifferenceReason *Max350Text `xml:"PricDiffRsn,omitempty"`
}

func (u *UnitPrice22) AddType() *TypeOfPrice46Choice {
	u.Type = new(TypeOfPrice46Choice)
	return u.Type
}

func (u *UnitPrice22) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice22) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice22) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice22) SetTaxableIncomePerShare(value, currency string) {
	u.TaxableIncomePerShare = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (u *UnitPrice22) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculated2Choice {
	u.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculated2Choice)
	return u.TaxableIncomePerShareCalculated
}

func (u *UnitPrice22) SetPriceDifferenceReason(value string) {
	u.PriceDifferenceReason = (*Max350Text)(&value)
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice3 struct {

	// Type and information about a price.
	PriceType *TypeOfPrice2Code `xml:"PricTp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Interest that has accumulated between the most recent payment of interest and the sale of the financial instrument.
	AccruedInterestNAV *CurrencyAndAmount `xml:"AcrdIntrstNAV,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *CurrencyAndAmount `xml:"TaxblIncmPerShr,omitempty"`
}

func (u *UnitPrice3) SetPriceType(value string) {
	u.PriceType = (*TypeOfPrice2Code)(&value)
}

func (u *UnitPrice3) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice3) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice3) SetAccruedInterestNAV(value, currency string) {
	u.AccruedInterestNAV = NewCurrencyAndAmount(value, currency)
}

func (u *UnitPrice3) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice3) SetTaxableIncomePerShare(value, currency string) {
	u.TaxableIncomePerShare = NewCurrencyAndAmount(value, currency)
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice4 struct {

	// Type and information about a price.
	Type *TypeOfPrice8Code `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceValue1 `xml:"Val"`
}

func (u *UnitPrice4) SetType(value string) {
	u.Type = (*TypeOfPrice8Code)(&value)
}

func (u *UnitPrice4) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice5 struct {

	// Type and information about a price.
	Type *PriceType1 `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *AmountPrice1Choice `xml:"TaxblIncmPerShr,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculated1 `xml:"TaxblIncmPerShrClctd,omitempty"`
}

func (u *UnitPrice5) AddType() *PriceType1 {
	u.Type = new(PriceType1)
	return u.Type
}

func (u *UnitPrice5) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice5) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice5) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice5) AddTaxableIncomePerShare() *AmountPrice1Choice {
	u.TaxableIncomePerShare = new(AmountPrice1Choice)
	return u.TaxableIncomePerShare
}

func (u *UnitPrice5) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculated1 {
	u.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculated1)
	return u.TaxableIncomePerShareCalculated
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice6 struct {

	// Type and information about a price.
	Type *PriceType2 `xml:"Tp"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`

	// Value of the price, eg, as a currency and value.
	ValueInInvestmentCurrency []*PriceValue1 `xml:"ValInInvstmtCcy"`

	// Value of the price, eg, as a currency and value.
	ValueInAlternativeCurrency []*PriceValue1 `xml:"ValInAltrntvCcy,omitempty"`

	// Indicates whether the price information can be used for the execution of a transaction.
	ForExecutionIndicator *YesNoIndicator `xml:"ForExctnInd"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Ratio applied on the non-adjusted price.
	CalculationBasis *PercentageRate `xml:"ClctnBsis,omitempty"`

	// Specifies the number of days from trade date that the counterparty on the other side of the trade should "given up" or divulged.
	NumberOfDaysAccrued *Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount included in the NAV that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *AmountPrice1Choice `xml:"TaxblIncmPerShr,omitempty"`

	// Specifies whether the fund calculates a taxable interest per share (TIS).
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculated1 `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Amount of money associated with a service.
	ChargeDetails []*Charge9 `xml:"ChrgDtls,omitempty"`

	// Information related to taxes that are due.
	TaxLiabilityDetails []*Tax8 `xml:"TaxLbltyDtls,omitempty"`

	// Information related to taxes that are paid back.
	TaxRefundDetails []*Tax8 `xml:"TaxRfndDtls,omitempty"`
}

func (u *UnitPrice6) AddType() *PriceType2 {
	u.Type = new(PriceType2)
	return u.Type
}

func (u *UnitPrice6) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}

func (u *UnitPrice6) AddValueInInvestmentCurrency() *PriceValue1 {
	newValue := new(PriceValue1)
	u.ValueInInvestmentCurrency = append(u.ValueInInvestmentCurrency, newValue)
	return newValue
}

func (u *UnitPrice6) AddValueInAlternativeCurrency() *PriceValue1 {
	newValue := new(PriceValue1)
	u.ValueInAlternativeCurrency = append(u.ValueInAlternativeCurrency, newValue)
	return newValue
}

func (u *UnitPrice6) SetForExecutionIndicator(value string) {
	u.ForExecutionIndicator = (*YesNoIndicator)(&value)
}

func (u *UnitPrice6) SetCumDividendIndicator(value string) {
	u.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (u *UnitPrice6) SetCalculationBasis(value string) {
	u.CalculationBasis = (*PercentageRate)(&value)
}

func (u *UnitPrice6) SetNumberOfDaysAccrued(value string) {
	u.NumberOfDaysAccrued = (*Number)(&value)
}

func (u *UnitPrice6) AddTaxableIncomePerShare() *AmountPrice1Choice {
	u.TaxableIncomePerShare = new(AmountPrice1Choice)
	return u.TaxableIncomePerShare
}

func (u *UnitPrice6) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculated1 {
	u.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculated1)
	return u.TaxableIncomePerShareCalculated
}

func (u *UnitPrice6) AddChargeDetails() *Charge9 {
	newValue := new(Charge9)
	u.ChargeDetails = append(u.ChargeDetails, newValue)
	return newValue
}

func (u *UnitPrice6) AddTaxLiabilityDetails() *Tax8 {
	newValue := new(Tax8)
	u.TaxLiabilityDetails = append(u.TaxLiabilityDetails, newValue)
	return newValue
}

func (u *UnitPrice6) AddTaxRefundDetails() *Tax8 {
	newValue := new(Tax8)
	u.TaxRefundDetails = append(u.TaxRefundDetails, newValue)
	return newValue
}

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice9 struct {

	// Specifies the unit of measurement. For example, kilo, tons.
	UnitOfMeasureCode *UnitOfMeasure4Code `xml:"UnitOfMeasrCd"`

	// Identifies the unit of measure not present in the code list.
	OtherUnitOfMeasure *Max35Text `xml:"OthrUnitOfMeasr"`

	// Price expressed as a currency and value.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Multiplication factor of measurement values. For example: goods that can be ordered by 36 pieces.
	Factor *Max15NumericText `xml:"Fctr,omitempty"`
}

func (u *UnitPrice9) SetUnitOfMeasureCode(value string) {
	u.UnitOfMeasureCode = (*UnitOfMeasure4Code)(&value)
}

func (u *UnitPrice9) SetOtherUnitOfMeasure(value string) {
	u.OtherUnitOfMeasure = (*Max35Text)(&value)
}

func (u *UnitPrice9) SetAmount(value, currency string) {
	u.Amount = NewCurrencyAndAmount(value, currency)
}

func (u *UnitPrice9) SetFactor(value string) {
	u.Factor = (*Max15NumericText)(&value)
}
