package iso20022

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace1 struct {

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
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation1 `xml:"PricDtls,omitempty"`

	// Currency exchange related to a securities order.
	ForeignExchangeDetails *ForeignExchangeTerms3 `xml:"FrgnXchgDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	BalanceBreakdownDetails []*SubBalanceInformation1 `xml:"BalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace1) AddAggregateQuantity() *BalanceQuantity1Choice {
	a.AggregateQuantity = new(BalanceQuantity1Choice)
	return a.AggregateQuantity
}

func (a *AggregateBalancePerSafekeepingPlace1) AddAvailableQuantity() *BalanceQuantity1Choice {
	a.AvailableQuantity = new(BalanceQuantity1Choice)
	return a.AvailableQuantity
}

func (a *AggregateBalancePerSafekeepingPlace1) AddNotAvailableQuantity() *BalanceQuantity1Choice {
	a.NotAvailableQuantity = new(BalanceQuantity1Choice)
	return a.NotAvailableQuantity
}

func (a *AggregateBalancePerSafekeepingPlace1) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace1) AddHoldingValue(value, currency string) {
	a.HoldingValue = append(a.HoldingValue, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (a *AggregateBalancePerSafekeepingPlace1) SetAccruedInterestAmount(value, currency string) {
	a.AccruedInterestAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace1) SetBookValue(value, currency string) {
	a.BookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace1) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace1) AddPriceDetails() *PriceInformation1 {
	newValue := new(PriceInformation1)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace1) AddForeignExchangeDetails() *ForeignExchangeTerms3 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms3)
	return a.ForeignExchangeDetails
}

func (a *AggregateBalancePerSafekeepingPlace1) AddBalanceBreakdownDetails() *SubBalanceInformation1 {
	newValue := new(SubBalanceInformation1)
	a.BalanceBreakdownDetails = append(a.BalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace1) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation {
	newValue := new(AdditionalBalanceInformation)
	a.AdditionalBalanceBreakdownDetails = append(a.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace11 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification5 `xml:"PlcOfListg,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *BalanceQuantity5Choice `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity5Choice `xml:"NotAvlblBal,omitempty"`

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

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace11) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace11) AddPlaceOfListing() *MarketIdentification5 {
	a.PlaceOfListing = new(MarketIdentification5)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace11) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace11) AddAvailableBalance() *BalanceQuantity5Choice {
	a.AvailableBalance = new(BalanceQuantity5Choice)
	return a.AvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace11) AddNotAvailableBalance() *BalanceQuantity5Choice {
	a.NotAvailableBalance = new(BalanceQuantity5Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace11) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace11) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace11) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace11) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace11) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace11) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace11) AddBalanceBreakdown() *SubBalanceInformation5 {
	newValue := new(SubBalanceInformation5)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace11) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation5 {
	newValue := new(AdditionalBalanceInformation5)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace11) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace12 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification5 `xml:"PlcOfListg,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

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

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace12) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace12) AddPlaceOfListing() *MarketIdentification5 {
	a.PlaceOfListing = new(MarketIdentification5)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace12) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace12) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace12) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace12) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace12) AddAccountBaseCurrencyAmounts() *BalanceAmounts1 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts1)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace12) AddInstrumentCurrencyAmounts() *BalanceAmounts1 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts1)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace12) AddAlternateReportingCurrencyAmounts() *BalanceAmounts1 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts1)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace12) AddQuantityBreakdown() *QuantityBreakdown4 {
	newValue := new(QuantityBreakdown4)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace12) AddBalanceBreakdown() *SubBalanceInformation6 {
	newValue := new(SubBalanceInformation6)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace12) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation6 {
	newValue := new(AdditionalBalanceInformation6)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace12) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace15 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *BalanceQuantity5Choice `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity5Choice `xml:"NotAvlblBal,omitempty"`

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

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace15) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace15) AddPlaceOfListing() *MarketIdentification3Choice {
	a.PlaceOfListing = new(MarketIdentification3Choice)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace15) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace15) AddAvailableBalance() *BalanceQuantity5Choice {
	a.AvailableBalance = new(BalanceQuantity5Choice)
	return a.AvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace15) AddNotAvailableBalance() *BalanceQuantity5Choice {
	a.NotAvailableBalance = new(BalanceQuantity5Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace15) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace15) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace15) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace15) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace15) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace15) AddQuantityBreakdown() *QuantityBreakdown13 {
	newValue := new(QuantityBreakdown13)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace15) AddBalanceBreakdown() *SubBalanceInformation9 {
	newValue := new(SubBalanceInformation9)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace15) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation9 {
	newValue := new(AdditionalBalanceInformation9)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace15) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace16 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

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

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace16) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace16) AddPlaceOfListing() *MarketIdentification3Choice {
	a.PlaceOfListing = new(MarketIdentification3Choice)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace16) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace16) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace16) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace16) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace16) AddAccountBaseCurrencyAmounts() *BalanceAmounts1 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts1)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace16) AddInstrumentCurrencyAmounts() *BalanceAmounts1 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts1)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace16) AddAlternateReportingCurrencyAmounts() *BalanceAmounts1 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts1)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace16) AddQuantityBreakdown() *QuantityBreakdown14 {
	newValue := new(QuantityBreakdown14)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace16) AddBalanceBreakdown() *SubBalanceInformation6 {
	newValue := new(SubBalanceInformation6)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace16) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation6 {
	newValue := new(AdditionalBalanceInformation6)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace16) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace2 struct {

	// Total quantity of financial instrument for the referenced holding.
	AggregateBalance *BalanceQuantity1Choice `xml:"AggtBal"`

	// Specifies the number of days used for calculating the accrued interest amount.
	DaysAccrued *Number `xml:"DaysAcrd,omitempty"`

	// Total value of a balance of the securities account for a specific financial instrument, expressed in one or more currencies.
	HoldingValue []*ActiveOrHistoricCurrencyAndAmount `xml:"HldgVal"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *ActiveOrHistoricCurrencyAndAmount `xml:"AcrdIntrstAmt,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	BookValue *ActiveOrHistoricCurrencyAndAmount `xml:"BookVal,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation1 `xml:"PricDtls"`

	// Currency exchange related to a securities order.
	ForeignExchangeDetails *ForeignExchangeTerms3 `xml:"FrgnXchgDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	BalanceBreakdownDetails []*SubBalanceInformation1 `xml:"BalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace2) AddAggregateBalance() *BalanceQuantity1Choice {
	a.AggregateBalance = new(BalanceQuantity1Choice)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace2) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace2) AddHoldingValue(value, currency string) {
	a.HoldingValue = append(a.HoldingValue, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (a *AggregateBalancePerSafekeepingPlace2) SetAccruedInterestAmount(value, currency string) {
	a.AccruedInterestAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace2) SetBookValue(value, currency string) {
	a.BookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace2) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace2) AddPriceDetails() *PriceInformation1 {
	newValue := new(PriceInformation1)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace2) AddForeignExchangeDetails() *ForeignExchangeTerms3 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms3)
	return a.ForeignExchangeDetails
}

func (a *AggregateBalancePerSafekeepingPlace2) AddBalanceBreakdownDetails() *SubBalanceInformation1 {
	newValue := new(SubBalanceInformation1)
	a.BalanceBreakdownDetails = append(a.BalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace2) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation {
	newValue := new(AdditionalBalanceInformation)
	a.AdditionalBalanceBreakdownDetails = append(a.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace20 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *BalanceQuantity5Choice `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity5Choice `xml:"NotAvlblBal,omitempty"`

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

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace20) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace20) AddPlaceOfListing() *MarketIdentification3Choice {
	a.PlaceOfListing = new(MarketIdentification3Choice)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace20) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace20) AddAvailableBalance() *BalanceQuantity5Choice {
	a.AvailableBalance = new(BalanceQuantity5Choice)
	return a.AvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace20) AddNotAvailableBalance() *BalanceQuantity5Choice {
	a.NotAvailableBalance = new(BalanceQuantity5Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace20) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace20) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace20) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace20) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace20) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace20) AddQuantityBreakdown() *QuantityBreakdown23 {
	newValue := new(QuantityBreakdown23)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace20) AddBalanceBreakdown() *SubBalanceInformation11 {
	newValue := new(SubBalanceInformation11)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace20) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation11 {
	newValue := new(AdditionalBalanceInformation11)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace20) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace21 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

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

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace21) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace21) AddPlaceOfListing() *MarketIdentification3Choice {
	a.PlaceOfListing = new(MarketIdentification3Choice)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace21) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace21) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace21) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace21) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace21) AddAccountBaseCurrencyAmounts() *BalanceAmounts1 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts1)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace21) AddInstrumentCurrencyAmounts() *BalanceAmounts1 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts1)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace21) AddAlternateReportingCurrencyAmounts() *BalanceAmounts1 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts1)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace21) AddQuantityBreakdown() *QuantityBreakdown24 {
	newValue := new(QuantityBreakdown24)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace21) AddBalanceBreakdown() *SubBalanceInformation6 {
	newValue := new(SubBalanceInformation6)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace21) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation6 {
	newValue := new(AdditionalBalanceInformation6)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace21) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace24 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Choice between formats for the entity to which the financial instruments are pledged.
	Pledgee *PledgeeFormat1Choice `xml:"Pldgee,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *BalanceQuantity5Choice `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity5Choice `xml:"NotAvlblBal,omitempty"`

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

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType12Choice `xml:"XpsrTp,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation11 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation11 `xml:"AddtlBalBrkdwn,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace24) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace24) AddPlaceOfListing() *MarketIdentification3Choice {
	a.PlaceOfListing = new(MarketIdentification3Choice)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace24) AddPledgee() *PledgeeFormat1Choice {
	a.Pledgee = new(PledgeeFormat1Choice)
	return a.Pledgee
}

func (a *AggregateBalancePerSafekeepingPlace24) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace24) AddAvailableBalance() *BalanceQuantity5Choice {
	a.AvailableBalance = new(BalanceQuantity5Choice)
	return a.AvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace24) AddNotAvailableBalance() *BalanceQuantity5Choice {
	a.NotAvailableBalance = new(BalanceQuantity5Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace24) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace24) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace24) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace24) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace24) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace24) AddQuantityBreakdown() *QuantityBreakdown23 {
	newValue := new(QuantityBreakdown23)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace24) AddExposureType() *ExposureType12Choice {
	a.ExposureType = new(ExposureType12Choice)
	return a.ExposureType
}

func (a *AggregateBalancePerSafekeepingPlace24) AddBalanceBreakdown() *SubBalanceInformation11 {
	newValue := new(SubBalanceInformation11)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace24) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation11 {
	newValue := new(AdditionalBalanceInformation11)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace24) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace25 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Choice between formats for the entity to which the financial instruments are pledged.
	Pledgee *PledgeeFormat1Choice `xml:"Pldgee,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

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

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType12Choice `xml:"XpsrTp,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation6 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation6 `xml:"AddtlBalBrkdwn,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace25) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace25) AddPlaceOfListing() *MarketIdentification3Choice {
	a.PlaceOfListing = new(MarketIdentification3Choice)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace25) AddPledgee() *PledgeeFormat1Choice {
	a.Pledgee = new(PledgeeFormat1Choice)
	return a.Pledgee
}

func (a *AggregateBalancePerSafekeepingPlace25) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace25) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace25) AddForeignExchangeDetails() *ForeignExchangeTerms14 {
	newValue := new(ForeignExchangeTerms14)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace25) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace25) AddAccountBaseCurrencyAmounts() *BalanceAmounts1 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts1)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace25) AddInstrumentCurrencyAmounts() *BalanceAmounts1 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts1)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace25) AddAlternateReportingCurrencyAmounts() *BalanceAmounts1 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts1)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace25) AddQuantityBreakdown() *QuantityBreakdown24 {
	newValue := new(QuantityBreakdown24)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace25) AddExposureType() *ExposureType12Choice {
	a.ExposureType = new(ExposureType12Choice)
	return a.ExposureType
}

func (a *AggregateBalancePerSafekeepingPlace25) AddBalanceBreakdown() *SubBalanceInformation6 {
	newValue := new(SubBalanceInformation6)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace25) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation6 {
	newValue := new(AdditionalBalanceInformation6)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace25) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace28 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specify the entity to which the financial instruments are pledged.
	Pledgee *Pledgee1 `xml:"Pldgee,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance6 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *Balance8 `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity8Choice `xml:"NotAvlblBal,omitempty"`

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

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation15 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation15 `xml:"AddtlBalBrkdwn,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace28) AddSafekeepingPlace() *SafeKeepingPlace1 {
	a.SafekeepingPlace = new(SafeKeepingPlace1)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace28) AddPlaceOfListing() *MarketIdentification3Choice {
	a.PlaceOfListing = new(MarketIdentification3Choice)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace28) AddPledgee() *Pledgee1 {
	a.Pledgee = new(Pledgee1)
	return a.Pledgee
}

func (a *AggregateBalancePerSafekeepingPlace28) AddAggregateBalance() *Balance6 {
	a.AggregateBalance = new(Balance6)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace28) AddAvailableBalance() *Balance8 {
	a.AvailableBalance = new(Balance8)
	return a.AvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace28) AddNotAvailableBalance() *BalanceQuantity8Choice {
	a.NotAvailableBalance = new(BalanceQuantity8Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace28) AddPriceDetails() *PriceInformation12 {
	newValue := new(PriceInformation12)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace28) AddForeignExchangeDetails() *ForeignExchangeTerms22 {
	newValue := new(ForeignExchangeTerms22)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace28) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace28) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace28) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace28) AddQuantityBreakdown() *QuantityBreakdown27 {
	newValue := new(QuantityBreakdown27)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace28) AddExposureType() *ExposureType16Choice {
	a.ExposureType = new(ExposureType16Choice)
	return a.ExposureType
}

func (a *AggregateBalancePerSafekeepingPlace28) AddBalanceBreakdown() *SubBalanceInformation15 {
	newValue := new(SubBalanceInformation15)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace28) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation15 {
	newValue := new(AdditionalBalanceInformation15)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace28) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace29 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Choice between formats for the entity to which the financial instruments are pledged.
	Pledgee *Pledgee1 `xml:"Pldgee,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance6 `xml:"AggtBal"`

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

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation14 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation14 `xml:"AddtlBalBrkdwn,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace29) AddSafekeepingPlace() *SafeKeepingPlace1 {
	a.SafekeepingPlace = new(SafeKeepingPlace1)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace29) AddPlaceOfListing() *MarketIdentification3Choice {
	a.PlaceOfListing = new(MarketIdentification3Choice)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace29) AddPledgee() *Pledgee1 {
	a.Pledgee = new(Pledgee1)
	return a.Pledgee
}

func (a *AggregateBalancePerSafekeepingPlace29) AddAggregateBalance() *Balance6 {
	a.AggregateBalance = new(Balance6)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace29) AddPriceDetails() *PriceInformation12 {
	newValue := new(PriceInformation12)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace29) AddForeignExchangeDetails() *ForeignExchangeTerms22 {
	newValue := new(ForeignExchangeTerms22)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace29) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace29) AddAccountBaseCurrencyAmounts() *BalanceAmounts1 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts1)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace29) AddInstrumentCurrencyAmounts() *BalanceAmounts1 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts1)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace29) AddAlternateReportingCurrencyAmounts() *BalanceAmounts1 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts1)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace29) AddQuantityBreakdown() *QuantityBreakdown28 {
	newValue := new(QuantityBreakdown28)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace29) AddExposureType() *ExposureType16Choice {
	a.ExposureType = new(ExposureType16Choice)
	return a.ExposureType
}

func (a *AggregateBalancePerSafekeepingPlace29) AddBalanceBreakdown() *SubBalanceInformation14 {
	newValue := new(SubBalanceInformation14)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace29) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation14 {
	newValue := new(AdditionalBalanceInformation14)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace29) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace3 struct {

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
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation2 `xml:"PricDtls,omitempty"`

	// Currency exchange related to a securities order.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	BalanceBreakdownDetails []*SubBalanceInformation2 `xml:"BalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation2 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace3) AddAggregateQuantity() *BalanceQuantity1Choice {
	a.AggregateQuantity = new(BalanceQuantity1Choice)
	return a.AggregateQuantity
}

func (a *AggregateBalancePerSafekeepingPlace3) AddAvailableQuantity() *BalanceQuantity1Choice {
	a.AvailableQuantity = new(BalanceQuantity1Choice)
	return a.AvailableQuantity
}

func (a *AggregateBalancePerSafekeepingPlace3) AddNotAvailableQuantity() *BalanceQuantity1Choice {
	a.NotAvailableQuantity = new(BalanceQuantity1Choice)
	return a.NotAvailableQuantity
}

func (a *AggregateBalancePerSafekeepingPlace3) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace3) AddHoldingValue(value, currency string) {
	a.HoldingValue = append(a.HoldingValue, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (a *AggregateBalancePerSafekeepingPlace3) SetPreviousHoldingValue(value, currency string) {
	a.PreviousHoldingValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace3) SetAccruedInterestAmount(value, currency string) {
	a.AccruedInterestAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace3) SetAccruedInterestAmountSign(value string) {
	a.AccruedInterestAmountSign = (*PlusOrMinusIndicator)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace3) SetBookValue(value, currency string) {
	a.BookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace3) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace3) AddPriceDetails() *PriceInformation2 {
	newValue := new(PriceInformation2)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace3) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return a.ForeignExchangeDetails
}

func (a *AggregateBalancePerSafekeepingPlace3) AddBalanceBreakdownDetails() *SubBalanceInformation2 {
	newValue := new(SubBalanceInformation2)
	a.BalanceBreakdownDetails = append(a.BalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace3) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation2 {
	newValue := new(AdditionalBalanceInformation2)
	a.AdditionalBalanceBreakdownDetails = append(a.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace30 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Choice between formats for the entity to which the financial instruments are pledged.
	Pledgee *Pledgee2 `xml:"Pldgee,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance10 `xml:"AggtBal"`

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

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation16 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation16 `xml:"AddtlBalBrkdwn,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *RestrictedFINXMax350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace30) AddSafekeepingPlace() *SafeKeepingPlace2 {
	a.SafekeepingPlace = new(SafeKeepingPlace2)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace30) AddPlaceOfListing() *MarketIdentification4Choice {
	a.PlaceOfListing = new(MarketIdentification4Choice)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace30) AddPledgee() *Pledgee2 {
	a.Pledgee = new(Pledgee2)
	return a.Pledgee
}

func (a *AggregateBalancePerSafekeepingPlace30) AddAggregateBalance() *Balance10 {
	a.AggregateBalance = new(Balance10)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace30) AddPriceDetails() *PriceInformation14 {
	newValue := new(PriceInformation14)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace30) AddForeignExchangeDetails() *ForeignExchangeTerms31 {
	newValue := new(ForeignExchangeTerms31)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace30) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace30) AddAccountBaseCurrencyAmounts() *BalanceAmounts5 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts5)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace30) AddInstrumentCurrencyAmounts() *BalanceAmounts5 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts5)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace30) AddAlternateReportingCurrencyAmounts() *BalanceAmounts5 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts5)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace30) AddQuantityBreakdown() *QuantityBreakdown39 {
	newValue := new(QuantityBreakdown39)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace30) AddExposureType() *ExposureType17Choice {
	a.ExposureType = new(ExposureType17Choice)
	return a.ExposureType
}

func (a *AggregateBalancePerSafekeepingPlace30) AddBalanceBreakdown() *SubBalanceInformation16 {
	newValue := new(SubBalanceInformation16)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace30) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation16 {
	newValue := new(AdditionalBalanceInformation16)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace30) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace31 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Specify the entity to which the financial instruments are pledged.
	Pledgee *Pledgee2 `xml:"Pldgee,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance10 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *Balance12 `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity12Choice `xml:"NotAvlblBal,omitempty"`

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

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Breakdown of the aggregate balance per meaningful sub-balances and availability.
	BalanceBreakdown []*SubBalanceInformation17 `xml:"BalBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdown []*AdditionalBalanceInformation17 `xml:"AddtlBalBrkdwn,omitempty"`

	// Provides additional information on the holding.
	HoldingAdditionalDetails *RestrictedFINXMax350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace31) AddSafekeepingPlace() *SafeKeepingPlace2 {
	a.SafekeepingPlace = new(SafeKeepingPlace2)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace31) AddPlaceOfListing() *MarketIdentification4Choice {
	a.PlaceOfListing = new(MarketIdentification4Choice)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace31) AddPledgee() *Pledgee2 {
	a.Pledgee = new(Pledgee2)
	return a.Pledgee
}

func (a *AggregateBalancePerSafekeepingPlace31) AddAggregateBalance() *Balance10 {
	a.AggregateBalance = new(Balance10)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace31) AddAvailableBalance() *Balance12 {
	a.AvailableBalance = new(Balance12)
	return a.AvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace31) AddNotAvailableBalance() *BalanceQuantity12Choice {
	a.NotAvailableBalance = new(BalanceQuantity12Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace31) AddPriceDetails() *PriceInformation14 {
	newValue := new(PriceInformation14)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace31) AddForeignExchangeDetails() *ForeignExchangeTerms31 {
	newValue := new(ForeignExchangeTerms31)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace31) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace31) AddAccountBaseCurrencyAmounts() *BalanceAmounts4 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts4)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace31) AddInstrumentCurrencyAmounts() *BalanceAmounts4 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts4)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace31) AddQuantityBreakdown() *QuantityBreakdown40 {
	newValue := new(QuantityBreakdown40)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace31) AddExposureType() *ExposureType17Choice {
	a.ExposureType = new(ExposureType17Choice)
	return a.ExposureType
}

func (a *AggregateBalancePerSafekeepingPlace31) AddBalanceBreakdown() *SubBalanceInformation17 {
	newValue := new(SubBalanceInformation17)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace31) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation17 {
	newValue := new(AdditionalBalanceInformation17)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace31) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace4 struct {

	// Total quantity of financial instrument for the referenced holding.
	AggregateBalance *BalanceQuantity1Choice `xml:"AggtBal"`

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

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc"`

	// Price of the financial instrument in one or more currencies.
	PriceDetails []*PriceInformation2 `xml:"PricDtls"`

	// Currency exchange related to a securities order.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	BalanceBreakdownDetails []*SubBalanceInformation2 `xml:"BalBrkdwnDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation2 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace4) AddAggregateBalance() *BalanceQuantity1Choice {
	a.AggregateBalance = new(BalanceQuantity1Choice)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace4) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace4) AddHoldingValue(value, currency string) {
	a.HoldingValue = append(a.HoldingValue, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (a *AggregateBalancePerSafekeepingPlace4) SetPreviousHoldingValue(value, currency string) {
	a.PreviousHoldingValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace4) SetAccruedInterestAmount(value, currency string) {
	a.AccruedInterestAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace4) SetAccruedInterestAmountSign(value string) {
	a.AccruedInterestAmountSign = (*PlusOrMinusIndicator)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace4) SetBookValue(value, currency string) {
	a.BookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AggregateBalancePerSafekeepingPlace4) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace4) AddPriceDetails() *PriceInformation2 {
	newValue := new(PriceInformation2)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace4) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return a.ForeignExchangeDetails
}

func (a *AggregateBalancePerSafekeepingPlace4) AddBalanceBreakdownDetails() *SubBalanceInformation2 {
	newValue := new(SubBalanceInformation2)
	a.BalanceBreakdownDetails = append(a.BalanceBreakdownDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace4) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation2 {
	newValue := new(AdditionalBalanceInformation2)
	a.AdditionalBalanceBreakdownDetails = append(a.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace7 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification5 `xml:"PlcOfListg,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

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

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace7) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace7) AddPlaceOfListing() *MarketIdentification5 {
	a.PlaceOfListing = new(MarketIdentification5)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace7) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace7) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace7) AddForeignExchangeDetails() *ForeignExchangeTerms1 {
	newValue := new(ForeignExchangeTerms1)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace7) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace7) AddAccountBaseCurrencyAmounts() *BalanceAmounts1 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts1)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace7) AddInstrumentCurrencyAmounts() *BalanceAmounts1 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts1)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace7) AddAlternateReportingCurrencyAmounts() *BalanceAmounts1 {
	a.AlternateReportingCurrencyAmounts = new(BalanceAmounts1)
	return a.AlternateReportingCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace7) AddQuantityBreakdown() *QuantityBreakdown4 {
	newValue := new(QuantityBreakdown4)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace7) AddBalanceBreakdown() *SubBalanceInformation6 {
	newValue := new(SubBalanceInformation6)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace7) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation6 {
	newValue := new(AdditionalBalanceInformation6)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace7) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}

// Net position of a segregated holding, in a single security, within the overall position held in a securities account at a specified place of safekeeping.
type AggregateBalancePerSafekeepingPlace8 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc"`

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification5 `xml:"PlcOfListg,omitempty"`

	// Total quantity of financial instruments of the balance.
	AggregateBalance *Balance1 `xml:"AggtBal"`

	// Total quantity of financial instruments of the balance that is available.
	AvailableBalance *BalanceQuantity5Choice `xml:"AvlblBal,omitempty"`

	// Total quantity of financial instruments of the balance that is not available.
	NotAvailableBalance *BalanceQuantity5Choice `xml:"NotAvlblBal,omitempty"`

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

	// Provides additional information on the holding.
	HoldingAdditionalDetails *Max350Text `xml:"HldgAddtlDtls,omitempty"`
}

func (a *AggregateBalancePerSafekeepingPlace8) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return a.SafekeepingPlace
}

func (a *AggregateBalancePerSafekeepingPlace8) AddPlaceOfListing() *MarketIdentification5 {
	a.PlaceOfListing = new(MarketIdentification5)
	return a.PlaceOfListing
}

func (a *AggregateBalancePerSafekeepingPlace8) AddAggregateBalance() *Balance1 {
	a.AggregateBalance = new(Balance1)
	return a.AggregateBalance
}

func (a *AggregateBalancePerSafekeepingPlace8) AddAvailableBalance() *BalanceQuantity5Choice {
	a.AvailableBalance = new(BalanceQuantity5Choice)
	return a.AvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace8) AddNotAvailableBalance() *BalanceQuantity5Choice {
	a.NotAvailableBalance = new(BalanceQuantity5Choice)
	return a.NotAvailableBalance
}

func (a *AggregateBalancePerSafekeepingPlace8) AddPriceDetails() *PriceInformation5 {
	newValue := new(PriceInformation5)
	a.PriceDetails = append(a.PriceDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace8) AddForeignExchangeDetails() *ForeignExchangeTerms1 {
	newValue := new(ForeignExchangeTerms1)
	a.ForeignExchangeDetails = append(a.ForeignExchangeDetails, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace8) SetDaysAccrued(value string) {
	a.DaysAccrued = (*Number)(&value)
}

func (a *AggregateBalancePerSafekeepingPlace8) AddAccountBaseCurrencyAmounts() *BalanceAmounts3 {
	a.AccountBaseCurrencyAmounts = new(BalanceAmounts3)
	return a.AccountBaseCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace8) AddInstrumentCurrencyAmounts() *BalanceAmounts3 {
	a.InstrumentCurrencyAmounts = new(BalanceAmounts3)
	return a.InstrumentCurrencyAmounts
}

func (a *AggregateBalancePerSafekeepingPlace8) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace8) AddBalanceBreakdown() *SubBalanceInformation5 {
	newValue := new(SubBalanceInformation5)
	a.BalanceBreakdown = append(a.BalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace8) AddAdditionalBalanceBreakdown() *AdditionalBalanceInformation5 {
	newValue := new(AdditionalBalanceInformation5)
	a.AdditionalBalanceBreakdown = append(a.AdditionalBalanceBreakdown, newValue)
	return newValue
}

func (a *AggregateBalancePerSafekeepingPlace8) SetHoldingAdditionalDetails(value string) {
	a.HoldingAdditionalDetails = (*Max350Text)(&value)
}
