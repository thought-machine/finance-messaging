package iso20022

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type AggregateBalanceInformation1 struct {

	// Total quantity of financial instrument for the referenced holding.
	AggregateQuantity *BalanceQuantity1Choice `xml:"AggtQty"`

	// Total quantity of financial instrument for the referenced holding that is available.
	AvailableQuantity *BalanceQuantity1Choice `xml:"AvlblQty,omitempty"`

	// Total quantity of financial instrument for the referenced holding that is not available.
	NotAvailableQuantity *BalanceQuantity1Choice `xml:"NotAvlblQty,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	HoldingValue []*ActiveOrHistoricCurrencyAndAmount `xml:"HldgVal,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *ActiveOrHistoricCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	BookValue *ActiveOrHistoricCurrencyAndAmount `xml:"BookVal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc,omitempty"`

	// Security held on the account for which the balance is calculated.
	FinancialInstrumentDetails *FinancialInstrument4 `xml:"FinInstrmDtls"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation1 `xml:"PricDtls,omitempty"`

	// Currency exchange related to a securities order.
	ForeignExchangeDetails *ForeignExchangeTerms3 `xml:"FrgnXchgDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	BalanceBreakdownDetails []*SubBalanceInformation1 `xml:"BalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation `xml:"AddtlBalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace1 `xml:"BalAtSfkpgPlc,omitempty"`
}

func (a *AggregateBalanceInformation1) AddAggregateQuantity() *BalanceQuantity1Choice {
	a.AggregateQuantity = new(BalanceQuantity1Choice)
	return a.AggregateQuantity
}

func (a *AggregateBalanceInformation1) AddAvailableQuantity() *BalanceQuantity1Choice {
	a.AvailableQuantity = new(BalanceQuantity1Choice)
	return a.AvailableQuantity
}

func (a *AggregateBalanceInformation1) AddNotAvailableQuantity() *BalanceQuantity1Choice {
	a.NotAvailableQuantity = new(BalanceQuantity1Choice)
	return a.NotAvailableQuantity
}

func (a *AggregateBalanceInformation1) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation1) AddHoldingValue(value, currency string) {
	a.HoldingValue = append(a.HoldingValue, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (a *AggregateBalanceInformation1) SetAccruedInterestAmount(value, currency string) {
	a.AccruedInterestAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation1) SetBookValue(value, currency string) {
	a.BookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation1) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation1) AddFinancialInstrumentDetails() *FinancialInstrument4 {
	a.FinancialInstrumentDetails = new(FinancialInstrument4)
	return a.FinancialInstrumentDetails
}

func (a *AggregateBalanceInformation1) AddPriceDetails() *PriceInformation1 {
	newValue := new(PriceInformation1)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation1) AddForeignExchangeDetails() *ForeignExchangeTerms3 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms3)
	return a.ForeignExchangeDetails
}

func (a *AggregateBalanceInformation1) AddBalanceBreakdownDetails() *SubBalanceInformation1 {
	newValue := new(SubBalanceInformation1)
	a.BalanceBreakdownDetails = append(a.BalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation1) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation {
	newValue := new(AdditionalBalanceInformation)
	a.AdditionalBalanceBreakdownDetails = append(a.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation1) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace1 {
	newValue := new(AggregateBalancePerSafekeepingPlace1)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation12 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument other than a investment funds.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes20 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Elements used to calculate the valuation haircut.
	ValuationHaircutDetails *BasicCollateralValuation1Details `xml:"ValtnHrcutDtls,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *BalanceQuantity5Choice `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity5Choice `xml:"NotAvlblBal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation5 `xml:"PricDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms14 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts3 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts3 `xml:"InstrmCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation5 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation5 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace11 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation12) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation12) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes20 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes20)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation12) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation12) AddValuationHaircutDetails() *BasicCollateralValuation1Details {
	a.ValuationHaircutDetails = new(BasicCollateralValuation1Details)
	return a.ValuationHaircutDetails
}

func (a *AggregateBalanceInformation12) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation12) AddAvailableBalance() *BalanceQuantity5Choice {
	a.AvailableBalance = new(BalanceQuantity5Choice)
	return a.AvailableBalance
}

func (a *AggregateBalanceInformation12) AddNotAvailableBalance() *BalanceQuantity5Choice {
	a.NotAvailableBalance = new(BalanceQuantity5Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalanceInformation12) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation12) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation12) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation12) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation12) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation12) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation12) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation12) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation12) AddBalanceBreakdown() *SubBalanceInformation5 {
	newValue := new(SubBalanceInformation5)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation12) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation5 {
	newValue := new(AdditionalBalanceInformation5)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation12) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace11 {
	newValue := new(AggregateBalancePerSafekeepingPlace11)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation12) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation12) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation13 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes20 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Details of the swap contract.
	AdditionalDerivativeAttributes *DerivativeBasicAttributes1 `xml:"AddtlDerivAttrbts,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation5 `xml:"PricDtls"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms14 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts1 `xml:"AcctBaseCcyAmts"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts1 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts1 `xml:"AltrnRptgCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown4 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation6 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation6 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace12 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation13) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation13) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes20 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes20)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation13) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation13) AddAdditionalDerivativeAttributes() *DerivativeBasicAttributes1 {
	a.AdditionalDerivativeAttributes = new(DerivativeBasicAttributes1)
	return a.AdditionalDerivativeAttributes
}

func (a *AggregateBalanceInformation13) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation13) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation13) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation13) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation13) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation13) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation13) AddAccountBaseCurrencyAmounts() *BalanceAmounts1 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts1)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation13) AddInstrumentCurrencyAmounts() *BalanceAmounts1 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts1)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation13) AddAlternateReportingCurrencyAmounts() *BalanceAmounts1 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts1)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalanceInformation13) AddQuantityBreakdown() *QuantityBreakdown4 {
	newValue := new(QuantityBreakdown4)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation13) AddBalanceBreakdown() *SubBalanceInformation6 {
	newValue := new(SubBalanceInformation6)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation13) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation6 {
	newValue := new(AdditionalBalanceInformation6)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation13) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace12 {
	newValue := new(AggregateBalancePerSafekeepingPlace12)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation13) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation13) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation16 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument other than a investment funds.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Elements used to calculate the valuation haircut.
	ValuationHaircutDetails *BasicCollateralValuation1Details `xml:"ValtnHrcutDtls,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *BalanceQuantity5Choice `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity5Choice `xml:"NotAvlblBal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation5 `xml:"PricDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms14 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts3 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts3 `xml:"InstrmCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown13 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation9 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation9 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace15 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation16) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation16) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation16) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation16) AddValuationHaircutDetails() *BasicCollateralValuation1Details {
	a.ValuationHaircutDetails = new(BasicCollateralValuation1Details)
	return a.ValuationHaircutDetails
}

func (a *AggregateBalanceInformation16) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation16) AddAvailableBalance() *BalanceQuantity5Choice {
	a.AvailableBalance = new(BalanceQuantity5Choice)
	return a.AvailableBalance
}

func (a *AggregateBalanceInformation16) AddNotAvailableBalance() *BalanceQuantity5Choice {
	a.NotAvailableBalance = new(BalanceQuantity5Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalanceInformation16) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation16) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation16) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation16) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation16) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation16) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation16) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation16) AddQuantityBreakdown() *QuantityBreakdown13 {
	newValue := new(QuantityBreakdown13)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation16) AddBalanceBreakdown() *SubBalanceInformation9 {
	newValue := new(SubBalanceInformation9)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation16) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation9 {
	newValue := new(AdditionalBalanceInformation9)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation16) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace15 {
	newValue := new(AggregateBalancePerSafekeepingPlace15)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation16) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation16) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation17 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Details of the swap contract.
	AdditionalDerivativeAttributes *DerivativeBasicAttributes1 `xml:"AddtlDerivAttrbts,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation5 `xml:"PricDtls"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms14 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts1 `xml:"AcctBaseCcyAmts"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts1 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts1 `xml:"AltrnRptgCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown14 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation6 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation6 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace16 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation17) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation17) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation17) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation17) AddAdditionalDerivativeAttributes() *DerivativeBasicAttributes1 {
	a.AdditionalDerivativeAttributes = new(DerivativeBasicAttributes1)
	return a.AdditionalDerivativeAttributes
}

func (a *AggregateBalanceInformation17) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation17) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation17) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation17) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation17) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation17) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation17) AddAccountBaseCurrencyAmounts() *BalanceAmounts1 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts1)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation17) AddInstrumentCurrencyAmounts() *BalanceAmounts1 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts1)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation17) AddAlternateReportingCurrencyAmounts() *BalanceAmounts1 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts1)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalanceInformation17) AddQuantityBreakdown() *QuantityBreakdown14 {
	newValue := new(QuantityBreakdown14)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation17) AddBalanceBreakdown() *SubBalanceInformation6 {
	newValue := new(SubBalanceInformation6)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation17) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation6 {
	newValue := new(AdditionalBalanceInformation6)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation17) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace16 {
	newValue := new(AggregateBalancePerSafekeepingPlace16)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation17) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation17) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type AggregateBalanceInformation2 struct {

	// Total quantity of financial instrument for the referenced holding.
	AggregateQuantity *BalanceQuantity1Choice `xml:"AggtQty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	HoldingValue []*ActiveOrHistoricCurrencyAndAmount `xml:"HldgVal"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *ActiveOrHistoricCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	BookValue *ActiveOrHistoricCurrencyAndAmount `xml:"BookVal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc,omitempty"`

	// Security held on the account for which the balance is calculated.
	FinancialInstrumentDetails *FinancialInstrument4 `xml:"FinInstrmDtls"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation1 `xml:"PricDtls"`

	// Currency exchange related to a securities order.
	ForeignExchangeDetails *ForeignExchangeTerms3 `xml:"FrgnXchgDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	BalanceBreakdownDetails []*SubBalanceInformation1 `xml:"BalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation `xml:"AddtlBalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace2 `xml:"BalAtSfkpgPlc,omitempty"`
}

func (a *AggregateBalanceInformation2) AddAggregateQuantity() *BalanceQuantity1Choice {
	a.AggregateQuantity = new(BalanceQuantity1Choice)
	return a.AggregateQuantity
}

func (a *AggregateBalanceInformation2) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation2) AddHoldingValue(value, currency string) {
	a.HoldingValue = append(a.HoldingValue, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (a *AggregateBalanceInformation2) SetAccruedInterestAmount(value, currency string) {
	a.AccruedInterestAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation2) SetBookValue(value, currency string) {
	a.BookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation2) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation2) AddFinancialInstrumentDetails() *FinancialInstrument4 {
	a.FinancialInstrumentDetails = new(FinancialInstrument4)
	return a.FinancialInstrumentDetails
}

func (a *AggregateBalanceInformation2) AddPriceDetails() *PriceInformation1 {
	newValue := new(PriceInformation1)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation2) AddForeignExchangeDetails() *ForeignExchangeTerms3 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms3)
	return a.ForeignExchangeDetails
}

func (a *AggregateBalanceInformation2) AddBalanceBreakdownDetails() *SubBalanceInformation1 {
	newValue := new(SubBalanceInformation1)
	a.BalanceBreakdownDetails = append(a.BalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation2) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation {
	newValue := new(AdditionalBalanceInformation)
	a.AdditionalBalanceBreakdownDetails = append(a.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation2) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace2 {
	newValue := new(AggregateBalancePerSafekeepingPlace2)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation21 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument other than a investment funds.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Elements used to calculate the valuation haircut.
	ValuationHaircutDetails *BasicCollateralValuation1Details `xml:"ValtnHrcutDtls,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *BalanceQuantity5Choice `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity5Choice `xml:"NotAvlblBal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation5 `xml:"PricDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms14 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts3 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts3 `xml:"InstrmCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown23 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation11 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation11 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace20 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation21) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation21) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation21) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation21) AddValuationHaircutDetails() *BasicCollateralValuation1Details {
	a.ValuationHaircutDetails = new(BasicCollateralValuation1Details)
	return a.ValuationHaircutDetails
}

func (a *AggregateBalanceInformation21) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation21) AddAvailableBalance() *BalanceQuantity5Choice {
	a.AvailableBalance = new(BalanceQuantity5Choice)
	return a.AvailableBalance
}

func (a *AggregateBalanceInformation21) AddNotAvailableBalance() *BalanceQuantity5Choice {
	a.NotAvailableBalance = new(BalanceQuantity5Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalanceInformation21) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation21) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation21) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation21) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation21) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation21) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation21) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation21) AddQuantityBreakdown() *QuantityBreakdown23 {
	newValue := new(QuantityBreakdown23)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation21) AddBalanceBreakdown() *SubBalanceInformation11 {
	newValue := new(SubBalanceInformation11)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation21) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation11 {
	newValue := new(AdditionalBalanceInformation11)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation21) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace20 {
	newValue := new(AggregateBalancePerSafekeepingPlace20)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation21) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation21) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation22 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Details of the swap contract.
	AdditionalDerivativeAttributes *DerivativeBasicAttributes1 `xml:"AddtlDerivAttrbts,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation5 `xml:"PricDtls"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms14 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts1 `xml:"AcctBaseCcyAmts"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts1 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts1 `xml:"AltrnRptgCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown24 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation6 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation6 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace21 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation22) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation22) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation22) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation22) AddAdditionalDerivativeAttributes() *DerivativeBasicAttributes1 {
	a.AdditionalDerivativeAttributes = new(DerivativeBasicAttributes1)
	return a.AdditionalDerivativeAttributes
}

func (a *AggregateBalanceInformation22) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation22) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation22) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation22) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation22) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation22) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation22) AddAccountBaseCurrencyAmounts() *BalanceAmounts1 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts1)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation22) AddInstrumentCurrencyAmounts() *BalanceAmounts1 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts1)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation22) AddAlternateReportingCurrencyAmounts() *BalanceAmounts1 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts1)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalanceInformation22) AddQuantityBreakdown() *QuantityBreakdown24 {
	newValue := new(QuantityBreakdown24)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation22) AddBalanceBreakdown() *SubBalanceInformation6 {
	newValue := new(SubBalanceInformation6)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation22) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation6 {
	newValue := new(AdditionalBalanceInformation6)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation22) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace21 {
	newValue := new(AggregateBalancePerSafekeepingPlace21)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation22) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation22) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation25 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument other than a investment funds.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Elements used to calculate the valuation haircut.
	ValuationHaircutDetails *BasicCollateralValuation1Details `xml:"ValtnHrcutDtls,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *BalanceQuantity5Choice `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity5Choice `xml:"NotAvlblBal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation5 `xml:"PricDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms14 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts3 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts3 `xml:"InstrmCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown23 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation11 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation11 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace24 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation25) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation25) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation25) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation25) AddValuationHaircutDetails() *BasicCollateralValuation1Details {
	a.ValuationHaircutDetails = new(BasicCollateralValuation1Details)
	return a.ValuationHaircutDetails
}

func (a *AggregateBalanceInformation25) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation25) AddAvailableBalance() *BalanceQuantity5Choice {
	a.AvailableBalance = new(BalanceQuantity5Choice)
	return a.AvailableBalance
}

func (a *AggregateBalanceInformation25) AddNotAvailableBalance() *BalanceQuantity5Choice {
	a.NotAvailableBalance = new(BalanceQuantity5Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalanceInformation25) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation25) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation25) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation25) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation25) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation25) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation25) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation25) AddQuantityBreakdown() *QuantityBreakdown23 {
	newValue := new(QuantityBreakdown23)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation25) AddBalanceBreakdown() *SubBalanceInformation11 {
	newValue := new(SubBalanceInformation11)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation25) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation11 {
	newValue := new(AdditionalBalanceInformation11)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation25) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace24 {
	newValue := new(AggregateBalancePerSafekeepingPlace24)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation25) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation25) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation26 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Details of the swap contract.
	AdditionalDerivativeAttributes *DerivativeBasicAttributes1 `xml:"AddtlDerivAttrbts,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation5 `xml:"PricDtls"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms14 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts1 `xml:"AcctBaseCcyAmts"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts1 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts1 `xml:"AltrnRptgCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown24 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation6 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation6 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace25 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation26) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation26) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation26) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation26) AddAdditionalDerivativeAttributes() *DerivativeBasicAttributes1 {
	a.AdditionalDerivativeAttributes = new(DerivativeBasicAttributes1)
	return a.AdditionalDerivativeAttributes
}

func (a *AggregateBalanceInformation26) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation26) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation26) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation26) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation26) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation26) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation26) AddAccountBaseCurrencyAmounts() *BalanceAmounts1 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts1)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation26) AddInstrumentCurrencyAmounts() *BalanceAmounts1 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts1)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation26) AddAlternateReportingCurrencyAmounts() *BalanceAmounts1 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts1)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalanceInformation26) AddQuantityBreakdown() *QuantityBreakdown24 {
	newValue := new(QuantityBreakdown24)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation26) AddBalanceBreakdown() *SubBalanceInformation6 {
	newValue := new(SubBalanceInformation6)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation26) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation6 {
	newValue := new(AdditionalBalanceInformation6)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation26) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace25 {
	newValue := new(AggregateBalancePerSafekeepingPlace25)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation26) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation26) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type AggregateBalanceInformation3 struct {

	// Total quantity of financial instrument for the referenced holding.
	AggregateQuantity *BalanceQuantity1Choice `xml:"AggtQty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	HoldingValue []*ActiveOrHistoricCurrencyAndAmount `xml:"HldgVal"`

	// Previous total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	PreviousHoldingValue *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsHldgVal,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *ActiveOrHistoricCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Indicates whether the accrued interest is a positive or negative amount.
	AccruedInterestAmountSign *PlusOrMinusIndicator `xml:"AcrdIntrstAmtSgn,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	BookValue *ActiveOrHistoricCurrencyAndAmount `xml:"BookVal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc,omitempty"`

	// Security held on the account for which the balance is calculated.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation2 `xml:"PricDtls"`

	// Currency exchange related to a securities order.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	BalanceBreakdownDetails []*SubBalanceInformation2 `xml:"BalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation2 `xml:"AddtlBalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace4 `xml:"BalAtSfkpgPlc,omitempty"`
}

func (a *AggregateBalanceInformation3) AddAggregateQuantity() *BalanceQuantity1Choice {
	a.AggregateQuantity = new(BalanceQuantity1Choice)
	return a.AggregateQuantity
}

func (a *AggregateBalanceInformation3) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation3) AddHoldingValue(value, currency string) {
	a.HoldingValue = append(a.HoldingValue, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (a *AggregateBalanceInformation3) SetPreviousHoldingValue(value, currency string) {
	a.PreviousHoldingValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation3) SetAccruedInterestAmount(value, currency string) {
	a.AccruedInterestAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation3) SetAccruedInterestAmountSign(value string) {
	a.AccruedInterestAmountSign = (*PlusOrMinusIndicator)(&value)
}

func (a *AggregateBalanceInformation3) SetBookValue(value, currency string) {
	a.BookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation3) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation3) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	a.FinancialInstrumentDetails = new(FinancialInstrument13)
	return a.FinancialInstrumentDetails
}

func (a *AggregateBalanceInformation3) AddPriceDetails() *PriceInformation2 {
	newValue := new(PriceInformation2)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation3) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return a.ForeignExchangeDetails
}

func (a *AggregateBalanceInformation3) AddBalanceBreakdownDetails() *SubBalanceInformation2 {
	newValue := new(SubBalanceInformation2)
	a.BalanceBreakdownDetails = append(a.BalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation3) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation2 {
	newValue := new(AdditionalBalanceInformation2)
	a.AdditionalBalanceBreakdownDetails = append(a.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation3) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace4 {
	newValue := new(AggregateBalancePerSafekeepingPlace4)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation30 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument other than a investment funds.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Elements used to calculate the valuation haircut.
	ValuationHaircutDetails *BasicCollateralValuation1Details `xml:"ValtnHrcutDtls,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance6 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *Balance8 `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity8Choice `xml:"NotAvlblBal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation12 `xml:"PricDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms22 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts3 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts3 `xml:"InstrmCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown27 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation15 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation15 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace28 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation30) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation30) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation30) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation30) AddValuationHaircutDetails() *BasicCollateralValuation1Details {
	a.ValuationHaircutDetails = new(BasicCollateralValuation1Details)
	return a.ValuationHaircutDetails
}

func (a *AggregateBalanceInformation30) AddAggregateBalance() *Balance6 {
	a.AggregateBalance = new(Balance6)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation30) AddAvailableBalance() *Balance8 {
	a.AvailableBalance = new(Balance8)
	return a.AvailableBalance
}

func (a *AggregateBalanceInformation30) AddNotAvailableBalance() *BalanceQuantity8Choice {
	a.NotAvailableBalance = new(BalanceQuantity8Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalanceInformation30) AddSafekeepingPlace() *SafeKeepingPlace1 {
	a.SafekeepingPlace = new(SafeKeepingPlace1)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation30) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation30) AddPriceDetails() *PriceInformation12 {
	newValue := new(PriceInformation12)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation30) AddForeignExchangeDetails() *ForeignExchangeTerms22 {
	newValue := new(ForeignExchangeTerms22)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation30) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation30) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation30) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation30) AddQuantityBreakdown() *QuantityBreakdown27 {
	newValue := new(QuantityBreakdown27)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation30) AddBalanceBreakdown() *SubBalanceInformation15 {
	newValue := new(SubBalanceInformation15)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation30) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation15 {
	newValue := new(AdditionalBalanceInformation15)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation30) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace28 {
	newValue := new(AggregateBalancePerSafekeepingPlace28)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation30) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation30) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation31 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Details of the swap contract.
	AdditionalDerivativeAttributes *DerivativeBasicAttributes1 `xml:"AddtlDerivAttrbts,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance6 `xml:"AggtBal"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation12 `xml:"PricDtls"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms22 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts1 `xml:"AcctBaseCcyAmts"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts1 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts1 `xml:"AltrnRptgCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown28 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation14 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation14 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace29 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation31) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation31) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation31) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation31) AddAdditionalDerivativeAttributes() *DerivativeBasicAttributes1 {
	a.AdditionalDerivativeAttributes = new(DerivativeBasicAttributes1)
	return a.AdditionalDerivativeAttributes
}

func (a *AggregateBalanceInformation31) AddAggregateBalance() *Balance6 {
	a.AggregateBalance = new(Balance6)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation31) AddSafekeepingPlace() *SafeKeepingPlace1 {
	a.SafekeepingPlace = new(SafeKeepingPlace1)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation31) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation31) AddPriceDetails() *PriceInformation12 {
	newValue := new(PriceInformation12)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation31) AddForeignExchangeDetails() *ForeignExchangeTerms22 {
	newValue := new(ForeignExchangeTerms22)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation31) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation31) AddAccountBaseCurrencyAmounts() *BalanceAmounts1 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts1)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation31) AddInstrumentCurrencyAmounts() *BalanceAmounts1 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts1)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation31) AddAlternateReportingCurrencyAmounts() *BalanceAmounts1 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts1)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalanceInformation31) AddQuantityBreakdown() *QuantityBreakdown28 {
	newValue := new(QuantityBreakdown28)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation31) AddBalanceBreakdown() *SubBalanceInformation14 {
	newValue := new(SubBalanceInformation14)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation31) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation14 {
	newValue := new(AdditionalBalanceInformation14)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation31) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace29 {
	newValue := new(AggregateBalancePerSafekeepingPlace29)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation31) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation31) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation32 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument22 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Details of the swap contract.
	AdditionalDerivativeAttributes *DerivativeBasicAttributes2 `xml:"AddtlDerivAttrbts,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance10 `xml:"AggtBal"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation14 `xml:"PricDtls"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms31 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts5 `xml:"AcctBaseCcyAmts"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts5 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts5 `xml:"AltrnRptgCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown39 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation16 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation16 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace30 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *RestrictedFINXMax350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation32) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation32) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation32) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument22 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument22)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation32) AddAdditionalDerivativeAttributes() *DerivativeBasicAttributes2 {
	a.AdditionalDerivativeAttributes = new(DerivativeBasicAttributes2)
	return a.AdditionalDerivativeAttributes
}

func (a *AggregateBalanceInformation32) AddAggregateBalance() *Balance10 {
	a.AggregateBalance = new(Balance10)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation32) AddSafekeepingPlace() *SafeKeepingPlace2 {
	a.SafekeepingPlace = new(SafeKeepingPlace2)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation32) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation32) AddPriceDetails() *PriceInformation14 {
	newValue := new(PriceInformation14)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation32) AddForeignExchangeDetails() *ForeignExchangeTerms31 {
	newValue := new(ForeignExchangeTerms31)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation32) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation32) AddAccountBaseCurrencyAmounts() *BalanceAmounts5 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts5)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation32) AddInstrumentCurrencyAmounts() *BalanceAmounts5 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts5)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation32) AddAlternateReportingCurrencyAmounts() *BalanceAmounts5 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts5)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalanceInformation32) AddQuantityBreakdown() *QuantityBreakdown39 {
	newValue := new(QuantityBreakdown39)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation32) AddBalanceBreakdown() *SubBalanceInformation16 {
	newValue := new(SubBalanceInformation16)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation32) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation16 {
	newValue := new(AdditionalBalanceInformation16)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation32) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace30 {
	newValue := new(AggregateBalancePerSafekeepingPlace30)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation32) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (a *AggregateBalanceInformation32) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation33 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument other than a investment funds.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument22 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Elements used to calculate the valuation haircut.
	ValuationHaircutDetails *BasicCollateralValuation1Details `xml:"ValtnHrcutDtls,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance10 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *Balance12 `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity12Choice `xml:"NotAvlblBal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation14 `xml:"PricDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms31 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts4 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts4 `xml:"InstrmCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown40 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation17 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation17 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace31 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *RestrictedFINXMax350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateBalanceInformation33) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation33) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation33) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument22 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument22)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation33) AddValuationHaircutDetails() *BasicCollateralValuation1Details {
	a.ValuationHaircutDetails = new(BasicCollateralValuation1Details)
	return a.ValuationHaircutDetails
}

func (a *AggregateBalanceInformation33) AddAggregateBalance() *Balance10 {
	a.AggregateBalance = new(Balance10)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation33) AddAvailableBalance() *Balance12 {
	a.AvailableBalance = new(Balance12)
	return a.AvailableBalance
}

func (a *AggregateBalanceInformation33) AddNotAvailableBalance() *BalanceQuantity12Choice {
	a.NotAvailableBalance = new(BalanceQuantity12Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalanceInformation33) AddSafekeepingPlace() *SafeKeepingPlace2 {
	a.SafekeepingPlace = new(SafeKeepingPlace2)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation33) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation33) AddPriceDetails() *PriceInformation14 {
	newValue := new(PriceInformation14)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation33) AddForeignExchangeDetails() *ForeignExchangeTerms31 {
	newValue := new(ForeignExchangeTerms31)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation33) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation33) AddAccountBaseCurrencyAmounts() *BalanceAmounts4 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts4)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation33) AddInstrumentCurrencyAmounts() *BalanceAmounts4 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts4)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation33) AddQuantityBreakdown() *QuantityBreakdown40 {
	newValue := new(QuantityBreakdown40)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation33) AddBalanceBreakdown() *SubBalanceInformation17 {
	newValue := new(SubBalanceInformation17)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation33) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation17 {
	newValue := new(AdditionalBalanceInformation17)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation33) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace31 {
	newValue := new(AggregateBalancePerSafekeepingPlace31)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation33) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (a *AggregateBalanceInformation33) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type AggregateBalanceInformation4 struct {

	// Total quantity of financial instrument for the referenced holding.
	AggregateQuantity *BalanceQuantity1Choice `xml:"AggtQty"`

	// Total quantity of financial instrument for the referenced holding that is available.
	AvailableQuantity *BalanceQuantity1Choice `xml:"AvlblQty,omitempty"`

	// Total quantity of financial instrument for the referenced holding that is not available.
	NotAvailableQuantity *BalanceQuantity1Choice `xml:"NotAvlblQty,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	HoldingValue []*ActiveOrHistoricCurrencyAndAmount `xml:"HldgVal,omitempty"`

	// Previous total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	PreviousHoldingValue *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsHldgVal,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *ActiveOrHistoricCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Indicates whether the accrued interest is a positive or negative amount.
	AccruedInterestAmountSign *PlusOrMinusIndicator `xml:"AcrdIntrstAmtSgn,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	BookValue *ActiveOrHistoricCurrencyAndAmount `xml:"BookVal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc,omitempty"`

	// Security held on the account for which the balance is calculated.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation2 `xml:"PricDtls,omitempty"`

	// Currency exchange related to a securities order.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	BalanceBreakdownDetails []*SubBalanceInformation2 `xml:"BalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation2 `xml:"AddtlBalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace3 `xml:"BalAtSfkpgPlc,omitempty"`
}

func (a *AggregateBalanceInformation4) AddAggregateQuantity() *BalanceQuantity1Choice {
	a.AggregateQuantity = new(BalanceQuantity1Choice)
	return a.AggregateQuantity
}

func (a *AggregateBalanceInformation4) AddAvailableQuantity() *BalanceQuantity1Choice {
	a.AvailableQuantity = new(BalanceQuantity1Choice)
	return a.AvailableQuantity
}

func (a *AggregateBalanceInformation4) AddNotAvailableQuantity() *BalanceQuantity1Choice {
	a.NotAvailableQuantity = new(BalanceQuantity1Choice)
	return a.NotAvailableQuantity
}

func (a *AggregateBalanceInformation4) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation4) AddHoldingValue(value, currency string) {
	a.HoldingValue = append(a.HoldingValue, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (a *AggregateBalanceInformation4) SetPreviousHoldingValue(value, currency string) {
	a.PreviousHoldingValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation4) SetAccruedInterestAmount(value, currency string) {
	a.AccruedInterestAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation4) SetAccruedInterestAmountSign(value string) {
	a.AccruedInterestAmountSign = (*PlusOrMinusIndicator)(&value)
}

func (a *AggregateBalanceInformation4) SetBookValue(value, currency string) {
	a.BookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalanceInformation4) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation4) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	a.FinancialInstrumentDetails = new(FinancialInstrument13)
	return a.FinancialInstrumentDetails
}

func (a *AggregateBalanceInformation4) AddPriceDetails() *PriceInformation2 {
	newValue := new(PriceInformation2)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation4) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return a.ForeignExchangeDetails
}

func (a *AggregateBalanceInformation4) AddBalanceBreakdownDetails() *SubBalanceInformation2 {
	newValue := new(SubBalanceInformation2)
	a.BalanceBreakdownDetails = append(a.BalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation4) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation2 {
	newValue := new(AdditionalBalanceInformation2)
	a.AdditionalBalanceBreakdownDetails = append(a.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation4) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace3 {
	newValue := new(AggregateBalancePerSafekeepingPlace3)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation8 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes8 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Details of the swap contract.
	AdditionalDerivativeAttributes *DerivativeBasicAttributes1 `xml:"AddtlDerivAttrbts,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation5 `xml:"PricDtls"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms1 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts1 `xml:"AcctBaseCcyAmts"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts1 `xml:"InstrmCcyAmts,omitempty"`

	// Valuation amounts provided in another currency than the base currency of the account.
	AlternateReportingCurrencyAmounts *BalanceAmounts1 `xml:"AltrnRptgCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown4 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation6 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation6 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace7 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension2 `xml:"Xtnsn,omitempty"`
}

func (a *AggregateBalanceInformation8) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation8) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes8 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes8)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation8) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation8) AddAdditionalDerivativeAttributes() *DerivativeBasicAttributes1 {
	a.AdditionalDerivativeAttributes = new(DerivativeBasicAttributes1)
	return a.AdditionalDerivativeAttributes
}

func (a *AggregateBalanceInformation8) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation8) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation8) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation8) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation8) AddForeignExchangeDetails() *ForeignExchangeTerms1 {
	newValue := new(ForeignExchangeTerms1)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation8) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation8) AddAccountBaseCurrencyAmounts() *BalanceAmounts1 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts1)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation8) AddInstrumentCurrencyAmounts() *BalanceAmounts1 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts1)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation8) AddAlternateReportingCurrencyAmounts() *BalanceAmounts1 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts1)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalanceInformation8) AddQuantityBreakdown() *QuantityBreakdown4 {
	newValue := new(QuantityBreakdown4)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation8) AddBalanceBreakdown() *SubBalanceInformation6 {
	newValue := new(SubBalanceInformation6)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation8) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation6 {
	newValue := new(AdditionalBalanceInformation6)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation8) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace7 {
	newValue := new(AggregateBalancePerSafekeepingPlace7)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation8) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation8) AddExtension() *Extension2 {
	newValue := new(Extension2)
	a.Extension = append(a.Extension, newValue)
	return newValue
}

// Overall position, in a single security, held in a securities account at a specified place of safekeeping.
type AggregateBalanceInformation9 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument other than a investment funds.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes8 `xml:"FinInstrmAttrbts,omitempty"`

	// Elements characterising an investment funds financial instrument.
	InvestmentFundsFinancialInstrumentAttributes *FinancialInstrument21 `xml:"InvstmtFndsFinInstrmAttrbts,omitempty"`

	// Elements used to calculate the valuation haircut.
	ValuationHaircutDetails *BasicCollateralValuation1Details `xml:"ValtnHrcutDtls,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *BalanceQuantity5Choice `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity5Choice `xml:"NotAvlblBal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies the corporate action options available to the account owner.
	CorporateActionOptionType *CorporateActionOption5Code `xml:"CorpActnOptnTp,omitempty"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation5 `xml:"PricDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms1 `xml:"FXDtls,omitempty"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyAmounts *BalanceAmounts3 `xml:"AcctBaseCcyAmts,omitempty"`

	// Valuation amounts provided in the currency of the financial instrument.
	InstrumentCurrencyAmounts *BalanceAmounts3 `xml:"InstrmCcyAmts,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation5 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation5 `xml:"AddtlBalBrkdwn,omitempty"`

	// Breakdown of positions per place of safekeeping (and optionally per place of listing).
	BalanceAtSafekeepingPlace []*AggregateBalancePerSafekeepingPlace8 `xml:"BalAtSfkpgPlc,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension2 `xml:"Xtnsn,omitempty"`
}

func (a *AggregateBalanceInformation9) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateBalanceInformation9) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes8 {
	a.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes8)
	return a.FinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation9) AddInvestmentFundsFinancialInstrumentAttributes() *FinancialInstrument21 {
	a.InvestmentFundsFinancialInstrumentAttributes = new(FinancialInstrument21)
	return a.InvestmentFundsFinancialInstrumentAttributes
}

func (a *AggregateBalanceInformation9) AddValuationHaircutDetails() *BasicCollateralValuation1Details {
	a.ValuationHaircutDetails = new(BasicCollateralValuation1Details)
	return a.ValuationHaircutDetails
}

func (a *AggregateBalanceInformation9) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalanceInformation9) AddAvailableBalance() *BalanceQuantity5Choice {
	a.AvailableBalance = new(BalanceQuantity5Choice)
	return a.AvailableBalance
}

func (a *AggregateBalanceInformation9) AddNotAvailableBalance() *BalanceQuantity5Choice {
	a.NotAvailableBalance = new(BalanceQuantity5Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalanceInformation9) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalanceInformation9) SetCorporateActionOptionType(value string) {
	a.CorporateActionOptionType = (*CorporateActionOption5Code)(&value)
}

func (a *AggregateBalanceInformation9) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation9) AddForeignExchangeDetails() *ForeignExchangeTerms1 {
	newValue := new(ForeignExchangeTerms1)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalanceInformation9) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalanceInformation9) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalanceInformation9) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalanceInformation9) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation9) AddBalanceBreakdown() *SubBalanceInformation5 {
	newValue := new(SubBalanceInformation5)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation9) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation5 {
	newValue := new(AdditionalBalanceInformation5)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalanceInformation9) AddBalanceAtSafekeepingPlace() *AggregateBalancePerSafekeepingPlace8 {
	newValue := new(AggregateBalancePerSafekeepingPlace8)
	a.BalanceAtSafekeepingPlace = append(a.BalanceAtSafekeepingPlace, newValue)
	return newValue
}

func (a *AggregateBalanceInformation9) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

func (a *AggregateBalanceInformation9) AddExtension() *Extension2 {
	newValue := new(Extension2)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
