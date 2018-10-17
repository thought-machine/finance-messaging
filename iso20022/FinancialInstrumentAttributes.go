package iso20022

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes15 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification5 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, that is, ownership, of the security.
	RegistrationForm *FormOfSecurity2Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency3Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus2Choice `xml:"PmtSts,omitempty"`

	// Indicates the direction of payment for asset or mortgage backed securities, that is, whether the repaid capital is distributed (payment direction is down) or capitalized (payment direction is up).
	PaymentDirection *PaymentDirection2Choice `xml:"PmtDrctn,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency3Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Indicates the  level of priority to claim on income and assets of the company in case of the payment of dividends and in the event of a bankruptcy, for example, ordinary/common stocks, preferred stocks, subordinated debt, etc.
	PreferenceToIncome *PreferenceToIncome2Choice `xml:"PrefToIncm,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType2Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number2Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *Number2Choice `xml:"PoolNb,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType1Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price2 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price2 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price2 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price2 `xml:"StrkPric,omitempty"`

	// Indicates the minimum tradable quantity of a security.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security by an ISIN.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification14 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes15) AddPlaceOfListing() *MarketIdentification5 {
	f.PlaceOfListing = new(MarketIdentification5)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes15) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes15) AddRegistrationForm() *FormOfSecurity2Choice {
	f.RegistrationForm = new(FormOfSecurity2Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes15) AddPaymentFrequency() *Frequency3Choice {
	f.PaymentFrequency = new(Frequency3Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes15) AddPaymentStatus() *SecuritiesPaymentStatus2Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus2Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes15) AddPaymentDirection() *PaymentDirection2Choice {
	f.PaymentDirection = new(PaymentDirection2Choice)
	return f.PaymentDirection
}

func (f *FinancialInstrumentAttributes15) AddVariableRateChangeFrequency() *Frequency3Choice {
	f.VariableRateChangeFrequency = new(Frequency3Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes15) AddPreferenceToIncome() *PreferenceToIncome2Choice {
	f.PreferenceToIncome = new(PreferenceToIncome2Choice)
	return f.PreferenceToIncome
}

func (f *FinancialInstrumentAttributes15) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes15) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes15) AddOptionType() *OptionType2Choice {
	f.OptionType = new(OptionType2Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes15) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes15) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes15) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes15) AddCouponAttachedNumber() *Number2Choice {
	f.CouponAttachedNumber = new(Number2Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes15) AddPoolNumber() *Number2Choice {
	f.PoolNumber = new(Number2Choice)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes15) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes15) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes15) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes15) AddMarketOrIndicativePrice() *PriceType1Choice {
	f.MarketOrIndicativePrice = new(PriceType1Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes15) AddExercisePrice() *Price2 {
	f.ExercisePrice = new(Price2)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes15) AddSubscriptionPrice() *Price2 {
	f.SubscriptionPrice = new(Price2)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes15) AddConversionPrice() *Price2 {
	f.ConversionPrice = new(Price2)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes15) AddStrikePrice() *Price2 {
	f.StrikePrice = new(Price2)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes15) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes15) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes15) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification14 {
	newValue := new(SecurityIdentification14)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes15) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}

// Description of the financial instrument.
type FinancialInstrumentAttributes16 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification2 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Initial issue/valuation price of a resulting security under a corporate action.
	IssuePrice *PriceFormat19Choice `xml:"IssePric,omitempty"`
}

func (f *FinancialInstrumentAttributes16) AddSecurityIdentification() *SecurityIdentification14 {
	f.SecurityIdentification = new(SecurityIdentification14)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes16) AddPlaceOfListing() *MarketIdentification2 {
	f.PlaceOfListing = new(MarketIdentification2)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes16) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes16) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes16) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes16) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes16) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) AddPreviousFactor() *RateFormat3Choice {
	f.PreviousFactor = new(RateFormat3Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes16) AddNextFactor() *RateFormat3Choice {
	f.NextFactor = new(RateFormat3Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes16) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes16) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes16) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes16) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes16) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes16) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes16) AddIssuePrice() *PriceFormat19Choice {
	f.IssuePrice = new(PriceFormat19Choice)
	return f.IssuePrice
}

// Description of the financial instrument.
type FinancialInstrumentAttributes17 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`

	// Quantity of entitled intermediate securities based on the balance of underlying securities.
	Quantity *DecimalNumber `xml:"Qty,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat1Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType9Choice `xml:"FrctnDspstn,omitempty"`

	// Quantity of intermediate securities awarded for a given quantity of underlying security.
	IntermediateSecuritiesToUnderlyingRatio *QuantityToQuantityRatio1 `xml:"IntrmdtSctiesToUndrlygRatio,omitempty"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *AmountPrice2 `xml:"MktPric,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *DateFormat16Choice `xml:"XpryDt"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateFormat16Choice `xml:"PstngDt"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period4 `xml:"TradgPrd,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat1Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal,omitempty"`
}

func (f *FinancialInstrumentAttributes17) AddSecurityIdentification() *SecurityIdentification14 {
	f.SecurityIdentification = new(SecurityIdentification14)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes17) SetQuantity(value string) {
	f.Quantity = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentAttributes17) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat1Choice {
	f.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat1Choice)
	return f.RenounceableEntitlementStatusType
}

func (f *FinancialInstrumentAttributes17) AddFractionDisposition() *FractionDispositionType9Choice {
	f.FractionDisposition = new(FractionDispositionType9Choice)
	return f.FractionDisposition
}

func (f *FinancialInstrumentAttributes17) AddIntermediateSecuritiesToUnderlyingRatio() *QuantityToQuantityRatio1 {
	f.IntermediateSecuritiesToUnderlyingRatio = new(QuantityToQuantityRatio1)
	return f.IntermediateSecuritiesToUnderlyingRatio
}

func (f *FinancialInstrumentAttributes17) AddMarketPrice() *AmountPrice2 {
	f.MarketPrice = new(AmountPrice2)
	return f.MarketPrice
}

func (f *FinancialInstrumentAttributes17) AddExpiryDate() *DateFormat16Choice {
	f.ExpiryDate = new(DateFormat16Choice)
	return f.ExpiryDate
}

func (f *FinancialInstrumentAttributes17) AddPostingDate() *DateFormat16Choice {
	f.PostingDate = new(DateFormat16Choice)
	return f.PostingDate
}

func (f *FinancialInstrumentAttributes17) AddTradingPeriod() *Period4 {
	f.TradingPeriod = new(Period4)
	return f.TradingPeriod
}

func (f *FinancialInstrumentAttributes17) AddUninstructedBalance() *BalanceFormat1Choice {
	f.UninstructedBalance = new(BalanceFormat1Choice)
	return f.UninstructedBalance
}

func (f *FinancialInstrumentAttributes17) AddInstructedBalance() *BalanceFormat1Choice {
	f.InstructedBalance = new(BalanceFormat1Choice)
	return f.InstructedBalance
}

// Description of the financial instrument.
type FinancialInstrumentAttributes18 struct {

	// Identification of a financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification2 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *PercentageRate `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *PercentageRate `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes18) AddSecurityIdentification() *SecurityIdentification14 {
	f.SecurityIdentification = new(SecurityIdentification14)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes18) AddPlaceOfListing() *MarketIdentification2 {
	f.PlaceOfListing = new(MarketIdentification2)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes18) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes18) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes18) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes18) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes18) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes18) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes18) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes18) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes18) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes18) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes18) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes18) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes18) SetPreviousFactor(value string) {
	f.PreviousFactor = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes18) SetNextFactor(value string) {
	f.NextFactor = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes18) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes18) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes18) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes18) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes18) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes18) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

// Description of the financial instrument.
type FinancialInstrumentAttributes19 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification2 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Percentage of the underlying assets of a fund that represents a debt and is in the scope of the EU Savings directive.
	PercentageOfDebtClaim *RateFormat3Choice `xml:"PctgOfDebtClm,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes19) AddSecurityIdentification() *SecurityIdentification14 {
	f.SecurityIdentification = new(SecurityIdentification14)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes19) AddPlaceOfListing() *MarketIdentification2 {
	f.PlaceOfListing = new(MarketIdentification2)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes19) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes19) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes19) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes19) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes19) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes19) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes19) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes19) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes19) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes19) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes19) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes19) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes19) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes19) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes19) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes19) AddPercentageOfDebtClaim() *RateFormat3Choice {
	f.PercentageOfDebtClaim = new(RateFormat3Choice)
	return f.PercentageOfDebtClaim
}

func (f *FinancialInstrumentAttributes19) AddPreviousFactor() *RateFormat3Choice {
	f.PreviousFactor = new(RateFormat3Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes19) AddNextFactor() *RateFormat3Choice {
	f.NextFactor = new(RateFormat3Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes19) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes19) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes19) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes19) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes20 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification5 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, this is, ownership, of the security.
	RegistrationForm *FormOfSecurity2Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency3Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus2Choice `xml:"PmtSts,omitempty"`

	// Indicates the direction of payment for asset or mortgage backed securities, this is, whether the repaid capital is distributed (payment direction is down) or capitalized (payment direction is up).
	PaymentDirection *PaymentDirection2Choice `xml:"PmtDrctn,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency3Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Indicates the  level of priority to claim on income and assets of the company in case of the payment of dividends and in the event of a bankruptcy, for example, ordinary/common stocks, preferred stocks, subordinated debt, etc.
	PreferenceToIncome *PreferenceToIncome2Choice `xml:"PrefToIncm,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType2Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number2Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *Number2Choice `xml:"PoolNb,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType1Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price2 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price2 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price2 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price2 `xml:"StrkPric,omitempty"`

	// Indicates the minimum tradable quantity of a security.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security by an ISIN.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification14 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes20) AddPlaceOfListing() *MarketIdentification5 {
	f.PlaceOfListing = new(MarketIdentification5)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes20) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes20) AddRegistrationForm() *FormOfSecurity2Choice {
	f.RegistrationForm = new(FormOfSecurity2Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes20) AddPaymentFrequency() *Frequency3Choice {
	f.PaymentFrequency = new(Frequency3Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes20) AddPaymentStatus() *SecuritiesPaymentStatus2Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus2Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes20) AddPaymentDirection() *PaymentDirection2Choice {
	f.PaymentDirection = new(PaymentDirection2Choice)
	return f.PaymentDirection
}

func (f *FinancialInstrumentAttributes20) AddVariableRateChangeFrequency() *Frequency3Choice {
	f.VariableRateChangeFrequency = new(Frequency3Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes20) AddPreferenceToIncome() *PreferenceToIncome2Choice {
	f.PreferenceToIncome = new(PreferenceToIncome2Choice)
	return f.PreferenceToIncome
}

func (f *FinancialInstrumentAttributes20) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes20) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes20) AddOptionType() *OptionType2Choice {
	f.OptionType = new(OptionType2Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes20) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes20) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes20) AddCouponAttachedNumber() *Number2Choice {
	f.CouponAttachedNumber = new(Number2Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes20) AddPoolNumber() *Number2Choice {
	f.PoolNumber = new(Number2Choice)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes20) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes20) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes20) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes20) AddMarketOrIndicativePrice() *PriceType1Choice {
	f.MarketOrIndicativePrice = new(PriceType1Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes20) AddExercisePrice() *Price2 {
	f.ExercisePrice = new(Price2)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes20) AddSubscriptionPrice() *Price2 {
	f.SubscriptionPrice = new(Price2)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes20) AddConversionPrice() *Price2 {
	f.ConversionPrice = new(Price2)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes20) AddStrikePrice() *Price2 {
	f.StrikePrice = new(Price2)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes20) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes20) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes20) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification14 {
	newValue := new(SecurityIdentification14)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes20) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes21 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification5 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, this is, ownership, of the security.
	RegistrationForm *FormOfSecurity2Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency3Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus2Choice `xml:"PmtSts,omitempty"`

	// Indicates the direction of payment for asset or mortgage backed securities, this is, whether the repaid capital is distributed (payment direction is down) or capitalized (payment direction is up).
	PaymentDirection *PaymentDirection2Choice `xml:"PmtDrctn,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency3Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Indicates the  level of priority to claim on income and assets of the company in case of the payment of dividends and in the event of a bankruptcy, for example, ordinary/common stocks, preferred stocks, subordinated debt, etc.
	PreferenceToIncome *PreferenceToIncome2Choice `xml:"PrefToIncm,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType2Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number2Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *Number2Choice `xml:"PoolNb,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown5 `xml:"QtyBrkdwn,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType1Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price2 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price2 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price2 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price2 `xml:"StrkPric,omitempty"`

	// Indicates the minimum tradable quantity of a security.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security by an ISIN.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification14 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes21) AddPlaceOfListing() *MarketIdentification5 {
	f.PlaceOfListing = new(MarketIdentification5)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes21) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes21) AddRegistrationForm() *FormOfSecurity2Choice {
	f.RegistrationForm = new(FormOfSecurity2Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes21) AddPaymentFrequency() *Frequency3Choice {
	f.PaymentFrequency = new(Frequency3Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes21) AddPaymentStatus() *SecuritiesPaymentStatus2Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus2Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes21) AddPaymentDirection() *PaymentDirection2Choice {
	f.PaymentDirection = new(PaymentDirection2Choice)
	return f.PaymentDirection
}

func (f *FinancialInstrumentAttributes21) AddVariableRateChangeFrequency() *Frequency3Choice {
	f.VariableRateChangeFrequency = new(Frequency3Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes21) AddPreferenceToIncome() *PreferenceToIncome2Choice {
	f.PreferenceToIncome = new(PreferenceToIncome2Choice)
	return f.PreferenceToIncome
}

func (f *FinancialInstrumentAttributes21) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes21) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes21) AddOptionType() *OptionType2Choice {
	f.OptionType = new(OptionType2Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes21) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes21) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes21) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes21) AddCouponAttachedNumber() *Number2Choice {
	f.CouponAttachedNumber = new(Number2Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes21) AddPoolNumber() *Number2Choice {
	f.PoolNumber = new(Number2Choice)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes21) AddQuantityBreakdown() *QuantityBreakdown5 {
	newValue := new(QuantityBreakdown5)
	f.QuantityBreakdown = append(f.QuantityBreakdown, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes21) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes21) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes21) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes21) AddMarketOrIndicativePrice() *PriceType1Choice {
	f.MarketOrIndicativePrice = new(PriceType1Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes21) AddExercisePrice() *Price2 {
	f.ExercisePrice = new(Price2)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes21) AddSubscriptionPrice() *Price2 {
	f.SubscriptionPrice = new(Price2)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes21) AddConversionPrice() *Price2 {
	f.ConversionPrice = new(Price2)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes21) AddStrikePrice() *Price2 {
	f.StrikePrice = new(Price2)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes21) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes21) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes21) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification14 {
	newValue := new(SecurityIdentification14)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes21) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}

// Description of the financial instrument.
type FinancialInstrumentAttributes3 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId"`

	// Quantity of entitled intermediate securities based on the balance of underlying securities.
	Quantity *DecimalNumber `xml:"Qty,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat1Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType3Choice `xml:"FrctnDspstn,omitempty"`

	// Quantity of intermediate securities awarded for a given quantity of underlying security.
	IntermediateSecuritiesToUnderlyingRatio *QuantityToQuantityRatio1 `xml:"IntrmdtSctiesToUndrlygRatio,omitempty"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *AmountPrice2 `xml:"MktPric,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *DateFormat5Choice `xml:"XpryDt"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateFormat5Choice `xml:"PstngDt"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3 `xml:"TradgPrd,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat1Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal,omitempty"`
}

func (f *FinancialInstrumentAttributes3) AddSecurityIdentification() *SecurityIdentification11 {
	f.SecurityIdentification = new(SecurityIdentification11)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes3) SetQuantity(value string) {
	f.Quantity = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentAttributes3) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat1Choice {
	f.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat1Choice)
	return f.RenounceableEntitlementStatusType
}

func (f *FinancialInstrumentAttributes3) AddFractionDisposition() *FractionDispositionType3Choice {
	f.FractionDisposition = new(FractionDispositionType3Choice)
	return f.FractionDisposition
}

func (f *FinancialInstrumentAttributes3) AddIntermediateSecuritiesToUnderlyingRatio() *QuantityToQuantityRatio1 {
	f.IntermediateSecuritiesToUnderlyingRatio = new(QuantityToQuantityRatio1)
	return f.IntermediateSecuritiesToUnderlyingRatio
}

func (f *FinancialInstrumentAttributes3) AddMarketPrice() *AmountPrice2 {
	f.MarketPrice = new(AmountPrice2)
	return f.MarketPrice
}

func (f *FinancialInstrumentAttributes3) AddExpiryDate() *DateFormat5Choice {
	f.ExpiryDate = new(DateFormat5Choice)
	return f.ExpiryDate
}

func (f *FinancialInstrumentAttributes3) AddPostingDate() *DateFormat5Choice {
	f.PostingDate = new(DateFormat5Choice)
	return f.PostingDate
}

func (f *FinancialInstrumentAttributes3) AddTradingPeriod() *Period3 {
	f.TradingPeriod = new(Period3)
	return f.TradingPeriod
}

func (f *FinancialInstrumentAttributes3) AddUninstructedBalance() *BalanceFormat1Choice {
	f.UninstructedBalance = new(BalanceFormat1Choice)
	return f.UninstructedBalance
}

func (f *FinancialInstrumentAttributes3) AddInstructedBalance() *BalanceFormat1Choice {
	f.InstructedBalance = new(BalanceFormat1Choice)
	return f.InstructedBalance
}

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes31 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Assessment of securities credit and investment risk.
	Rating *Rating1 `xml:"Ratg,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber *Max35Text `xml:"CertNb,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat3Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, ie, ownership, of the security.
	RegistrationForm *FormOfSecurity4Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency7Choice `xml:"PmtFrqcy,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency7Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType30Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle6Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType4Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Deadline by which a convertible security must be converted according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Date on which the current factor will be changed to the next factor.
	NextFactorDate *ISODate `xml:"NxtFctrDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// End ratio of principal outstanding to the original balance.
	EndFactor *BaseOneRate `xml:"EndFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Percentage of the underlying assets of a fund that represents a debt and is in the scope of the EU Savings directive.
	PercentageOfDebtClaims *PercentageRate `xml:"PctgOfDebtClms,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number1Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *GenericIdentification37 `xml:"PoolNb,omitempty"`

	// Number allocated by options exchanges to record that an option has undergone a change in its contract specifications (particularly adjustment of the strike price)
	VersionNumber *Number1Choice `xml:"VrsnNb,omitempty"`

	// Indicates whether the interest bearing security is convertible into another type of security.
	ConvertibleIndicator *YesNoIndicator `xml:"ConvtblInd,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the underlying security is owned by the writer of the entitlement.
	CoveredIndicator *YesNoIndicator `xml:"CvrdInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Indicates whether the warrants on a financial instrument (which has been traded cum warrants) will be attached on delivery.
	WarrantAttachedOnDelivery *YesNoIndicator `xml:"WarrtAttchdOnDlvry,omitempty"`

	// Indicates whether the payment of the coupon (interest) on an interest bearing instrument is off the normal schedule.
	OddCouponIndicator *YesNoIndicator `xml:"OddCpnInd,omitempty"`

	// Indicates whether, in the case of a debt security subject to redemption before maturity, such redemption could have an impact on the yield.
	RedemptionYieldImpact *YesNoIndicator `xml:"RedYldImpct,omitempty"`

	// Indicates whether the actual yield of an asset-backed security may vary according to the rate at which the underlying receivables or other financial assets are prepaid.
	YieldVariance *YesNoIndicator `xml:"YldVar,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price4 `xml:"ExrcPric,omitempty"`

	// Fixed price at which a new securities issue is being offered to the public.
	SubscriptionPrice *Price4 `xml:"SbcptPric,omitempty"`

	// Price at which a given convertible security can be converted to common stock .
	ConversionPrice *Price4 `xml:"ConvsPric,omitempty"`

	// Amount included in the Net Asset Value(NAV) that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *Price4 `xml:"TaxblIncmPerShr,omitempty"`

	// Indicates the minimum tradable quantity of a security.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Signed face amount and amortised value of security.
	FaceAmount *ImpliedCurrencyAndAmount `xml:"FaceAmt,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity18Choice `xml:"CtrctSz,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes31) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes31) AddRating() *Rating1 {
	f.Rating = new(Rating1)
	return f.Rating
}

func (f *FinancialInstrumentAttributes31) SetCertificateNumber(value string) {
	f.CertificateNumber = (*Max35Text)(&value)
}

func (f *FinancialInstrumentAttributes31) AddDayCountBasis() *InterestComputationMethodFormat3Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat3Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes31) AddRegistrationForm() *FormOfSecurity4Choice {
	f.RegistrationForm = new(FormOfSecurity4Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes31) AddPaymentFrequency() *Frequency7Choice {
	f.PaymentFrequency = new(Frequency7Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes31) AddVariableRateChangeFrequency() *Frequency7Choice {
	f.VariableRateChangeFrequency = new(Frequency7Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes31) AddClassificationType() *ClassificationType30Choice {
	f.ClassificationType = new(ClassificationType30Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes31) AddOptionStyle() *OptionStyle6Choice {
	f.OptionStyle = new(OptionStyle6Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes31) AddOptionType() *OptionType4Choice {
	f.OptionType = new(OptionType4Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes31) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes31) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetNextFactorDate(value string) {
	f.NextFactorDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetEndFactor(value string) {
	f.EndFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetPercentageOfDebtClaims(value string) {
	f.PercentageOfDebtClaims = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes31) AddCouponAttachedNumber() *Number1Choice {
	f.CouponAttachedNumber = new(Number1Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes31) AddPoolNumber() *GenericIdentification37 {
	f.PoolNumber = new(GenericIdentification37)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes31) AddVersionNumber() *Number1Choice {
	f.VersionNumber = new(Number1Choice)
	return f.VersionNumber
}

func (f *FinancialInstrumentAttributes31) SetConvertibleIndicator(value string) {
	f.ConvertibleIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetCoveredIndicator(value string) {
	f.CoveredIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetWarrantAttachedOnDelivery(value string) {
	f.WarrantAttachedOnDelivery = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetOddCouponIndicator(value string) {
	f.OddCouponIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetRedemptionYieldImpact(value string) {
	f.RedemptionYieldImpact = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetYieldVariance(value string) {
	f.YieldVariance = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) AddExercisePrice() *Price4 {
	f.ExercisePrice = new(Price4)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes31) AddSubscriptionPrice() *Price4 {
	f.SubscriptionPrice = new(Price4)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes31) AddConversionPrice() *Price4 {
	f.ConversionPrice = new(Price4)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes31) AddTaxableIncomePerShare() *Price4 {
	f.TaxableIncomePerShare = new(Price4)
	return f.TaxableIncomePerShare
}

func (f *FinancialInstrumentAttributes31) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes31) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes31) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes31) SetFaceAmount(value, currency string) {
	f.FaceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentAttributes31) AddContractSize() *FinancialInstrumentQuantity18Choice {
	f.ContractSize = new(FinancialInstrumentQuantity18Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes31) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}

// Description of the financial instrument.
type FinancialInstrumentAttributes32 struct {

	// Identification of a financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *PercentageRate `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *PercentageRate `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes32) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes32) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes32) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes32) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes32) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes32) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetPreviousFactor(value string) {
	f.PreviousFactor = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetNextFactor(value string) {
	f.NextFactor = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes32) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes32) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes32) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes32) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

// Description of the financial instrument.
type FinancialInstrumentAttributes33 struct {

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Percentage of the underlying assets of a fund that represents a debt and is in the scope of the EU Savings directive.
	PercentageOfDebtClaim *RateFormat3Choice `xml:"PctgOfDebtClm,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes33) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes33) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes33) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes33) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes33) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes33) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes33) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes33) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes33) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes33) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes33) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes33) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes33) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes33) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes33) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes33) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes33) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes33) AddPercentageOfDebtClaim() *RateFormat3Choice {
	f.PercentageOfDebtClaim = new(RateFormat3Choice)
	return f.PercentageOfDebtClaim
}

func (f *FinancialInstrumentAttributes33) AddPreviousFactor() *RateFormat3Choice {
	f.PreviousFactor = new(RateFormat3Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes33) AddNextFactor() *RateFormat3Choice {
	f.NextFactor = new(RateFormat3Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes33) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes33) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes33) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes33) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

// Description of the financial instrument.
type FinancialInstrumentAttributes34 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Initial issue/valuation price of a resulting security under a corporate action.
	IssuePrice *PriceFormat19Choice `xml:"IssePric,omitempty"`
}

func (f *FinancialInstrumentAttributes34) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes34) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes34) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes34) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes34) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes34) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes34) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes34) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes34) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes34) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes34) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes34) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes34) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes34) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes34) AddPreviousFactor() *RateFormat3Choice {
	f.PreviousFactor = new(RateFormat3Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes34) AddNextFactor() *RateFormat3Choice {
	f.NextFactor = new(RateFormat3Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes34) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes34) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes34) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes34) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes34) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes34) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes34) AddIssuePrice() *PriceFormat19Choice {
	f.IssuePrice = new(PriceFormat19Choice)
	return f.IssuePrice
}

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes35 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, this is, ownership, of the security.
	RegistrationForm *FormOfSecurity2Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency3Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus2Choice `xml:"PmtSts,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency3Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType2Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Rate of return anticipated on a bond when held until maturity date.
	YieldToMaturityRate *PercentageRate `xml:"YldToMtrtyRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number2Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *GenericIdentification37 `xml:"PoolNb,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType1Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price2 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price2 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price2 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price2 `xml:"StrkPric,omitempty"`

	// Indicates the minimum tradable quantity of a security.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security by an ISIN.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification14 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes35) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes35) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes35) AddRegistrationForm() *FormOfSecurity2Choice {
	f.RegistrationForm = new(FormOfSecurity2Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes35) AddPaymentFrequency() *Frequency3Choice {
	f.PaymentFrequency = new(Frequency3Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes35) AddPaymentStatus() *SecuritiesPaymentStatus2Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus2Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes35) AddVariableRateChangeFrequency() *Frequency3Choice {
	f.VariableRateChangeFrequency = new(Frequency3Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes35) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes35) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes35) AddOptionType() *OptionType2Choice {
	f.OptionType = new(OptionType2Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes35) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes35) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetYieldToMaturityRate(value string) {
	f.YieldToMaturityRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes35) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes35) AddCouponAttachedNumber() *Number2Choice {
	f.CouponAttachedNumber = new(Number2Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes35) AddPoolNumber() *GenericIdentification37 {
	f.PoolNumber = new(GenericIdentification37)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes35) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes35) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes35) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes35) AddMarketOrIndicativePrice() *PriceType1Choice {
	f.MarketOrIndicativePrice = new(PriceType1Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes35) AddExercisePrice() *Price2 {
	f.ExercisePrice = new(Price2)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes35) AddSubscriptionPrice() *Price2 {
	f.SubscriptionPrice = new(Price2)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes35) AddConversionPrice() *Price2 {
	f.ConversionPrice = new(Price2)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes35) AddStrikePrice() *Price2 {
	f.StrikePrice = new(Price2)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes35) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes35) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes35) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification14 {
	newValue := new(SecurityIdentification14)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes35) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes36 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, this is, ownership, of the security.
	RegistrationForm *FormOfSecurity2Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency3Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus2Choice `xml:"PmtSts,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency3Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType2Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Rate of return anticipated on a bond when held until maturity date.
	YieldToMaturityRate *PercentageRate `xml:"YldToMtrtyRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number2Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *GenericIdentification37 `xml:"PoolNb,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown15 `xml:"QtyBrkdwn,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType1Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price2 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price2 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price2 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price2 `xml:"StrkPric,omitempty"`

	// Indicates the minimum tradable quantity of a security.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security by an ISIN.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification14 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes36) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes36) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes36) AddRegistrationForm() *FormOfSecurity2Choice {
	f.RegistrationForm = new(FormOfSecurity2Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes36) AddPaymentFrequency() *Frequency3Choice {
	f.PaymentFrequency = new(Frequency3Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes36) AddPaymentStatus() *SecuritiesPaymentStatus2Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus2Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes36) AddVariableRateChangeFrequency() *Frequency3Choice {
	f.VariableRateChangeFrequency = new(Frequency3Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes36) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes36) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes36) AddOptionType() *OptionType2Choice {
	f.OptionType = new(OptionType2Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes36) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes36) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetYieldToMaturityRate(value string) {
	f.YieldToMaturityRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes36) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes36) AddCouponAttachedNumber() *Number2Choice {
	f.CouponAttachedNumber = new(Number2Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes36) AddPoolNumber() *GenericIdentification37 {
	f.PoolNumber = new(GenericIdentification37)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes36) AddQuantityBreakdown() *QuantityBreakdown15 {
	newValue := new(QuantityBreakdown15)
	f.QuantityBreakdown = append(f.QuantityBreakdown, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes36) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes36) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes36) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes36) AddMarketOrIndicativePrice() *PriceType1Choice {
	f.MarketOrIndicativePrice = new(PriceType1Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes36) AddExercisePrice() *Price2 {
	f.ExercisePrice = new(Price2)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes36) AddSubscriptionPrice() *Price2 {
	f.SubscriptionPrice = new(Price2)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes36) AddConversionPrice() *Price2 {
	f.ConversionPrice = new(Price2)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes36) AddStrikePrice() *Price2 {
	f.StrikePrice = new(Price2)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes36) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes36) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes36) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification14 {
	newValue := new(SecurityIdentification14)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes36) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes4 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification5 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, ie, ownership, of the security.
	RegistrationForm *FormOfSecurity2Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency3Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus2Choice `xml:"PmtSts,omitempty"`

	// Indicates the direction of payment for asset or mortgage backed securities, ie, whether the repaid capital is distributed (payment direction is down) or capitalized (payment direction is up).
	PaymentDirection *PaymentDirection2Choice `xml:"PmtDrctn,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency3Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Indicates the  level of priority to claim on income and assets of the company in case of the payment of dividends and in the event of a bankruptcy, for example, ordinary/common stocks, preferred stocks, subordinated debt, etc.
	PreferenceToIncome *PreferenceToIncome2Choice `xml:"PrefToIncm,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType2Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number2Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *Number2Choice `xml:"PoolNb,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown5 `xml:"QtyBrkdwn,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType1Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price2 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price2 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price2 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price2 `xml:"StrkPric,omitempty"`

	// Indicates the minimum tradable quantity of a security.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security by an ISIN.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification11 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes4) AddPlaceOfListing() *MarketIdentification5 {
	f.PlaceOfListing = new(MarketIdentification5)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes4) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes4) AddRegistrationForm() *FormOfSecurity2Choice {
	f.RegistrationForm = new(FormOfSecurity2Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes4) AddPaymentFrequency() *Frequency3Choice {
	f.PaymentFrequency = new(Frequency3Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes4) AddPaymentStatus() *SecuritiesPaymentStatus2Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus2Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes4) AddPaymentDirection() *PaymentDirection2Choice {
	f.PaymentDirection = new(PaymentDirection2Choice)
	return f.PaymentDirection
}

func (f *FinancialInstrumentAttributes4) AddVariableRateChangeFrequency() *Frequency3Choice {
	f.VariableRateChangeFrequency = new(Frequency3Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes4) AddPreferenceToIncome() *PreferenceToIncome2Choice {
	f.PreferenceToIncome = new(PreferenceToIncome2Choice)
	return f.PreferenceToIncome
}

func (f *FinancialInstrumentAttributes4) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes4) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes4) AddOptionType() *OptionType2Choice {
	f.OptionType = new(OptionType2Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes4) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes4) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes4) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes4) AddCouponAttachedNumber() *Number2Choice {
	f.CouponAttachedNumber = new(Number2Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes4) AddPoolNumber() *Number2Choice {
	f.PoolNumber = new(Number2Choice)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes4) AddQuantityBreakdown() *QuantityBreakdown5 {
	newValue := new(QuantityBreakdown5)
	f.QuantityBreakdown = append(f.QuantityBreakdown, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes4) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes4) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes4) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes4) AddMarketOrIndicativePrice() *PriceType1Choice {
	f.MarketOrIndicativePrice = new(PriceType1Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes4) AddExercisePrice() *Price2 {
	f.ExercisePrice = new(Price2)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes4) AddSubscriptionPrice() *Price2 {
	f.SubscriptionPrice = new(Price2)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes4) AddConversionPrice() *Price2 {
	f.ConversionPrice = new(Price2)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes4) AddStrikePrice() *Price2 {
	f.StrikePrice = new(Price2)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes4) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes4) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes4) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification11 {
	newValue := new(SecurityIdentification11)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes4) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}

// Description of the financial instrument.
type FinancialInstrumentAttributes43 struct {

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Percentage of the underlying assets of a fund that represents a debt, for example, in the context of the EU Savings directive.
	PercentageOfDebtClaim *RateFormat3Choice `xml:"PctgOfDebtClm,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Provides the ratio between the quantity of warrants and the quantity of underlying securities.
	WarrantParity *QuantityToQuantityRatio1 `xml:"WarrtParity,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes43) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes43) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes43) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes43) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes43) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes43) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes43) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes43) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes43) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes43) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes43) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes43) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes43) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes43) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes43) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes43) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes43) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes43) AddPercentageOfDebtClaim() *RateFormat3Choice {
	f.PercentageOfDebtClaim = new(RateFormat3Choice)
	return f.PercentageOfDebtClaim
}

func (f *FinancialInstrumentAttributes43) AddPreviousFactor() *RateFormat3Choice {
	f.PreviousFactor = new(RateFormat3Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes43) AddNextFactor() *RateFormat3Choice {
	f.NextFactor = new(RateFormat3Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes43) AddWarrantParity() *QuantityToQuantityRatio1 {
	f.WarrantParity = new(QuantityToQuantityRatio1)
	return f.WarrantParity
}

func (f *FinancialInstrumentAttributes43) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes43) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes43) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes43) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes44 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Assessment of securities credit and investment risk.
	Rating *Rating1 `xml:"Ratg,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber *Max35Text `xml:"CertNb,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat3Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, ie, ownership, of the security.
	RegistrationForm *FormOfSecurity4Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency7Choice `xml:"PmtFrqcy,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency7Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType30Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle6Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType4Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Deadline by which a convertible security must be converted according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Date on which the current factor will be changed to the next factor.
	NextFactorDate *ISODate `xml:"NxtFctrDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// End ratio of principal outstanding to the original balance.
	EndFactor *BaseOneRate `xml:"EndFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Percentage of the underlying assets of a fund that represents a debt, for example, in the context of the EU Savings directive.
	PercentageOfDebtClaims *PercentageRate `xml:"PctgOfDebtClms,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number1Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *GenericIdentification37 `xml:"PoolNb,omitempty"`

	// Number allocated by options exchanges to record that an option has undergone a change in its contract specifications (particularly adjustment of the strike price)
	VersionNumber *Number1Choice `xml:"VrsnNb,omitempty"`

	// Indicates whether the interest bearing security is convertible into another type of security.
	ConvertibleIndicator *YesNoIndicator `xml:"ConvtblInd,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the underlying security is owned by the writer of the entitlement.
	CoveredIndicator *YesNoIndicator `xml:"CvrdInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Indicates whether the warrants on a financial instrument (which has been traded cum warrants) will be attached on delivery.
	WarrantAttachedOnDelivery *YesNoIndicator `xml:"WarrtAttchdOnDlvry,omitempty"`

	// Indicates whether the payment of the coupon (interest) on an interest bearing instrument is off the normal schedule.
	OddCouponIndicator *YesNoIndicator `xml:"OddCpnInd,omitempty"`

	// Indicates whether, in the case of a debt security subject to redemption before maturity, such redemption could have an impact on the yield.
	RedemptionYieldImpact *YesNoIndicator `xml:"RedYldImpct,omitempty"`

	// Indicates whether the actual yield of an asset-backed security may vary according to the rate at which the underlying receivables or other financial assets are prepaid.
	YieldVariance *YesNoIndicator `xml:"YldVar,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price4 `xml:"ExrcPric,omitempty"`

	// Fixed price at which a new securities issue is being offered to the public.
	SubscriptionPrice *Price4 `xml:"SbcptPric,omitempty"`

	// Price at which a given convertible security can be converted to common stock .
	ConversionPrice *Price4 `xml:"ConvsPric,omitempty"`

	// Amount included in the Net Asset Value(NAV) that corresponds to gains directly or indirectly derived from interest payment, for example in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *Price4 `xml:"TaxblIncmPerShr,omitempty"`

	// Indicates the minimum tradable quantity of a security.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Signed face amount and amortised value of security.
	FaceAmount *ImpliedCurrencyAndAmount `xml:"FaceAmt,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity18Choice `xml:"CtrctSz,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes44) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes44) AddRating() *Rating1 {
	f.Rating = new(Rating1)
	return f.Rating
}

func (f *FinancialInstrumentAttributes44) SetCertificateNumber(value string) {
	f.CertificateNumber = (*Max35Text)(&value)
}

func (f *FinancialInstrumentAttributes44) AddDayCountBasis() *InterestComputationMethodFormat3Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat3Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes44) AddRegistrationForm() *FormOfSecurity4Choice {
	f.RegistrationForm = new(FormOfSecurity4Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes44) AddPaymentFrequency() *Frequency7Choice {
	f.PaymentFrequency = new(Frequency7Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes44) AddVariableRateChangeFrequency() *Frequency7Choice {
	f.VariableRateChangeFrequency = new(Frequency7Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes44) AddClassificationType() *ClassificationType30Choice {
	f.ClassificationType = new(ClassificationType30Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes44) AddOptionStyle() *OptionStyle6Choice {
	f.OptionStyle = new(OptionStyle6Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes44) AddOptionType() *OptionType4Choice {
	f.OptionType = new(OptionType4Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes44) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes44) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetNextFactorDate(value string) {
	f.NextFactorDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetEndFactor(value string) {
	f.EndFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes44) SetPercentageOfDebtClaims(value string) {
	f.PercentageOfDebtClaims = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes44) AddCouponAttachedNumber() *Number1Choice {
	f.CouponAttachedNumber = new(Number1Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes44) AddPoolNumber() *GenericIdentification37 {
	f.PoolNumber = new(GenericIdentification37)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes44) AddVersionNumber() *Number1Choice {
	f.VersionNumber = new(Number1Choice)
	return f.VersionNumber
}

func (f *FinancialInstrumentAttributes44) SetConvertibleIndicator(value string) {
	f.ConvertibleIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes44) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes44) SetCoveredIndicator(value string) {
	f.CoveredIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes44) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes44) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes44) SetWarrantAttachedOnDelivery(value string) {
	f.WarrantAttachedOnDelivery = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes44) SetOddCouponIndicator(value string) {
	f.OddCouponIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes44) SetRedemptionYieldImpact(value string) {
	f.RedemptionYieldImpact = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes44) SetYieldVariance(value string) {
	f.YieldVariance = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes44) AddExercisePrice() *Price4 {
	f.ExercisePrice = new(Price4)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes44) AddSubscriptionPrice() *Price4 {
	f.SubscriptionPrice = new(Price4)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes44) AddConversionPrice() *Price4 {
	f.ConversionPrice = new(Price4)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes44) AddTaxableIncomePerShare() *Price4 {
	f.TaxableIncomePerShare = new(Price4)
	return f.TaxableIncomePerShare
}

func (f *FinancialInstrumentAttributes44) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes44) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes44) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes44) SetFaceAmount(value, currency string) {
	f.FaceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentAttributes44) AddContractSize() *FinancialInstrumentQuantity18Choice {
	f.ContractSize = new(FinancialInstrumentQuantity18Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes44) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}

// Description of the financial instrument.
type FinancialInstrumentAttributes46 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`

	// Quantity of entitled intermediate securities based on the balance of underlying securities.
	Quantity *DecimalNumber `xml:"Qty,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat1Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType15Choice `xml:"FrctnDspstn,omitempty"`

	// Quantity of intermediate securities awarded for a given quantity of underlying security.
	IntermediateSecuritiesToUnderlyingRatio *QuantityToQuantityRatio1 `xml:"IntrmdtSctiesToUndrlygRatio,omitempty"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *AmountPrice2 `xml:"MktPric,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *DateFormat16Choice `xml:"XpryDt"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateFormat16Choice `xml:"PstngDt"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period4 `xml:"TradgPrd,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat1Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal,omitempty"`
}

func (f *FinancialInstrumentAttributes46) AddSecurityIdentification() *SecurityIdentification14 {
	f.SecurityIdentification = new(SecurityIdentification14)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes46) SetQuantity(value string) {
	f.Quantity = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentAttributes46) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat1Choice {
	f.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat1Choice)
	return f.RenounceableEntitlementStatusType
}

func (f *FinancialInstrumentAttributes46) AddFractionDisposition() *FractionDispositionType15Choice {
	f.FractionDisposition = new(FractionDispositionType15Choice)
	return f.FractionDisposition
}

func (f *FinancialInstrumentAttributes46) AddIntermediateSecuritiesToUnderlyingRatio() *QuantityToQuantityRatio1 {
	f.IntermediateSecuritiesToUnderlyingRatio = new(QuantityToQuantityRatio1)
	return f.IntermediateSecuritiesToUnderlyingRatio
}

func (f *FinancialInstrumentAttributes46) AddMarketPrice() *AmountPrice2 {
	f.MarketPrice = new(AmountPrice2)
	return f.MarketPrice
}

func (f *FinancialInstrumentAttributes46) AddExpiryDate() *DateFormat16Choice {
	f.ExpiryDate = new(DateFormat16Choice)
	return f.ExpiryDate
}

func (f *FinancialInstrumentAttributes46) AddPostingDate() *DateFormat16Choice {
	f.PostingDate = new(DateFormat16Choice)
	return f.PostingDate
}

func (f *FinancialInstrumentAttributes46) AddTradingPeriod() *Period4 {
	f.TradingPeriod = new(Period4)
	return f.TradingPeriod
}

func (f *FinancialInstrumentAttributes46) AddUninstructedBalance() *BalanceFormat1Choice {
	f.UninstructedBalance = new(BalanceFormat1Choice)
	return f.UninstructedBalance
}

func (f *FinancialInstrumentAttributes46) AddInstructedBalance() *BalanceFormat1Choice {
	f.InstructedBalance = new(BalanceFormat1Choice)
	return f.InstructedBalance
}

// Description of the financial instrument.
type FinancialInstrumentAttributes48 struct {

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Percentage of the underlying assets of a fund that represents a debt, for example, in the context of the EU Savings directive.
	PercentageOfDebtClaim *RateFormat3Choice `xml:"PctgOfDebtClm,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Provides the ratio between the quantity of warrants and the quantity of underlying securities.
	WarrantParity *QuantityToQuantityRatio1 `xml:"WarrtParity,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes48) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes48) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes48) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes48) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes48) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes48) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes48) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes48) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes48) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes48) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes48) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes48) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes48) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes48) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes48) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes48) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes48) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes48) AddPercentageOfDebtClaim() *RateFormat3Choice {
	f.PercentageOfDebtClaim = new(RateFormat3Choice)
	return f.PercentageOfDebtClaim
}

func (f *FinancialInstrumentAttributes48) AddPreviousFactor() *RateFormat12Choice {
	f.PreviousFactor = new(RateFormat12Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes48) AddNextFactor() *RateFormat12Choice {
	f.NextFactor = new(RateFormat12Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes48) AddWarrantParity() *QuantityToQuantityRatio1 {
	f.WarrantParity = new(QuantityToQuantityRatio1)
	return f.WarrantParity
}

func (f *FinancialInstrumentAttributes48) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes48) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes48) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes48) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

// Description of the financial instrument.
type FinancialInstrumentAttributes49 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Initial issue/valuation price of a resulting security under a corporate action.
	IssuePrice *PriceFormat19Choice `xml:"IssePric,omitempty"`
}

func (f *FinancialInstrumentAttributes49) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes49) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes49) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes49) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes49) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes49) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes49) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes49) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes49) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes49) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes49) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes49) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes49) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes49) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes49) AddPreviousFactor() *RateFormat12Choice {
	f.PreviousFactor = new(RateFormat12Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes49) AddNextFactor() *RateFormat12Choice {
	f.NextFactor = new(RateFormat12Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes49) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes49) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes49) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes49) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes49) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes49) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes49) AddIssuePrice() *PriceFormat19Choice {
	f.IssuePrice = new(PriceFormat19Choice)
	return f.IssuePrice
}

// Description of the financial instrument.
type FinancialInstrumentAttributes5 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification2 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Initial issue/valuation price of a resulting security under a corporate action.
	IssuePrice *PriceFormat11Choice `xml:"IssePric,omitempty"`
}

func (f *FinancialInstrumentAttributes5) AddSecurityIdentification() *SecurityIdentification11 {
	f.SecurityIdentification = new(SecurityIdentification11)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes5) AddPlaceOfListing() *MarketIdentification2 {
	f.PlaceOfListing = new(MarketIdentification2)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes5) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes5) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes5) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes5) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes5) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes5) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes5) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes5) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes5) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes5) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes5) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes5) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes5) AddPreviousFactor() *RateFormat3Choice {
	f.PreviousFactor = new(RateFormat3Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes5) AddNextFactor() *RateFormat3Choice {
	f.NextFactor = new(RateFormat3Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes5) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes5) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes5) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes5) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes5) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes5) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes5) AddIssuePrice() *PriceFormat11Choice {
	f.IssuePrice = new(PriceFormat11Choice)
	return f.IssuePrice
}

// Description of the financial instrument.
type FinancialInstrumentAttributes50 struct {

	// Identification of a financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *BaseOne14Rate `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *BaseOne14Rate `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes50) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes50) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes50) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes50) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes50) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes50) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes50) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes50) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes50) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes50) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes50) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes50) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes50) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes50) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes50) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOne14Rate)(&value)
}

func (f *FinancialInstrumentAttributes50) SetNextFactor(value string) {
	f.NextFactor = (*BaseOne14Rate)(&value)
}

func (f *FinancialInstrumentAttributes50) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes50) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes50) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes50) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes50) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes50) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

// Description of the financial instrument.
type FinancialInstrumentAttributes6 struct {

	// Identification of a financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId,omitempty"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification2 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *PercentageRate `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *PercentageRate `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes6) AddSecurityIdentification() *SecurityIdentification11 {
	f.SecurityIdentification = new(SecurityIdentification11)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes6) AddPlaceOfListing() *MarketIdentification2 {
	f.PlaceOfListing = new(MarketIdentification2)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes6) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes6) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes6) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes6) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes6) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes6) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes6) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes6) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes6) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes6) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes6) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes6) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes6) SetPreviousFactor(value string) {
	f.PreviousFactor = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes6) SetNextFactor(value string) {
	f.NextFactor = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes6) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes6) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes6) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes6) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes6) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes6) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes63 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat4Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, this is, ownership, of the security.
	RegistrationForm *FormOfSecurity6Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency23Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus5Choice `xml:"PmtSts,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency23Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType32Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle8Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType6Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Rate of return anticipated on a bond when held until maturity date.
	YieldToMaturityRate *PercentageRate `xml:"YldToMtrtyRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number22Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *GenericIdentification37 `xml:"PoolNb,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown31 `xml:"QtyBrkdwn,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType1Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price2 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price2 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price2 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price2 `xml:"StrkPric,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification19 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes63) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes63) AddDayCountBasis() *InterestComputationMethodFormat4Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat4Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes63) AddRegistrationForm() *FormOfSecurity6Choice {
	f.RegistrationForm = new(FormOfSecurity6Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes63) AddPaymentFrequency() *Frequency23Choice {
	f.PaymentFrequency = new(Frequency23Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes63) AddPaymentStatus() *SecuritiesPaymentStatus5Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus5Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes63) AddVariableRateChangeFrequency() *Frequency23Choice {
	f.VariableRateChangeFrequency = new(Frequency23Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes63) AddClassificationType() *ClassificationType32Choice {
	f.ClassificationType = new(ClassificationType32Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes63) AddOptionStyle() *OptionStyle8Choice {
	f.OptionStyle = new(OptionStyle8Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes63) AddOptionType() *OptionType6Choice {
	f.OptionType = new(OptionType6Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes63) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes63) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetYieldToMaturityRate(value string) {
	f.YieldToMaturityRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes63) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes63) AddCouponAttachedNumber() *Number22Choice {
	f.CouponAttachedNumber = new(Number22Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes63) AddPoolNumber() *GenericIdentification37 {
	f.PoolNumber = new(GenericIdentification37)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes63) AddQuantityBreakdown() *QuantityBreakdown31 {
	newValue := new(QuantityBreakdown31)
	f.QuantityBreakdown = append(f.QuantityBreakdown, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes63) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes63) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes63) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes63) AddMarketOrIndicativePrice() *PriceType1Choice {
	f.MarketOrIndicativePrice = new(PriceType1Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes63) AddExercisePrice() *Price2 {
	f.ExercisePrice = new(Price2)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes63) AddSubscriptionPrice() *Price2 {
	f.SubscriptionPrice = new(Price2)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes63) AddConversionPrice() *Price2 {
	f.ConversionPrice = new(Price2)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes63) AddStrikePrice() *Price2 {
	f.StrikePrice = new(Price2)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes63) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes63) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes63) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification19 {
	newValue := new(SecurityIdentification19)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes63) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes64 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat4Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, this is, ownership, of the security.
	RegistrationForm *FormOfSecurity6Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency23Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus5Choice `xml:"PmtSts,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency23Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType32Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle8Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType6Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Rate of return anticipated on a bond when held until maturity date.
	YieldToMaturityRate *PercentageRate `xml:"YldToMtrtyRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number22Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *GenericIdentification37 `xml:"PoolNb,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType1Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price2 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price2 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price2 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price2 `xml:"StrkPric,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification19 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes64) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes64) AddDayCountBasis() *InterestComputationMethodFormat4Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat4Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes64) AddRegistrationForm() *FormOfSecurity6Choice {
	f.RegistrationForm = new(FormOfSecurity6Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes64) AddPaymentFrequency() *Frequency23Choice {
	f.PaymentFrequency = new(Frequency23Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes64) AddPaymentStatus() *SecuritiesPaymentStatus5Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus5Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes64) AddVariableRateChangeFrequency() *Frequency23Choice {
	f.VariableRateChangeFrequency = new(Frequency23Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes64) AddClassificationType() *ClassificationType32Choice {
	f.ClassificationType = new(ClassificationType32Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes64) AddOptionStyle() *OptionStyle8Choice {
	f.OptionStyle = new(OptionStyle8Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes64) AddOptionType() *OptionType6Choice {
	f.OptionType = new(OptionType6Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes64) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes64) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetYieldToMaturityRate(value string) {
	f.YieldToMaturityRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes64) AddCouponAttachedNumber() *Number22Choice {
	f.CouponAttachedNumber = new(Number22Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes64) AddPoolNumber() *GenericIdentification37 {
	f.PoolNumber = new(GenericIdentification37)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes64) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes64) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes64) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes64) AddMarketOrIndicativePrice() *PriceType1Choice {
	f.MarketOrIndicativePrice = new(PriceType1Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes64) AddExercisePrice() *Price2 {
	f.ExercisePrice = new(Price2)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes64) AddSubscriptionPrice() *Price2 {
	f.SubscriptionPrice = new(Price2)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes64) AddConversionPrice() *Price2 {
	f.ConversionPrice = new(Price2)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes64) AddStrikePrice() *Price2 {
	f.StrikePrice = new(Price2)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes64) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes64) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes64) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification19 {
	newValue := new(SecurityIdentification19)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes64) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}

// Description of the financial instrument.
type FinancialInstrumentAttributes65 struct {

	// Identification of a financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat4Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType32Choice `xml:"ClssfctnTp,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *BaseOne14Rate `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *BaseOne14Rate `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes65) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes65) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes65) AddDayCountBasis() *InterestComputationMethodFormat4Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat4Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes65) AddClassificationType() *ClassificationType32Choice {
	f.ClassificationType = new(ClassificationType32Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes65) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes65) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes65) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes65) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes65) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes65) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes65) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes65) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes65) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes65) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes65) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOne14Rate)(&value)
}

func (f *FinancialInstrumentAttributes65) SetNextFactor(value string) {
	f.NextFactor = (*BaseOne14Rate)(&value)
}

func (f *FinancialInstrumentAttributes65) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes65) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes65) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes65) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

// Description of the financial instrument.
type FinancialInstrumentAttributes66 struct {

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat4Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType32Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle8Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Percentage of the underlying assets of a fund that represents a debt, for example, in the context of the EU Savings directive.
	PercentageOfDebtClaim *RateFormat3Choice `xml:"PctgOfDebtClm,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Provides the ratio between the quantity of warrants and the quantity of underlying securities.
	WarrantParity *QuantityToQuantityRatio1 `xml:"WarrtParity,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes66) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes66) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes66) AddDayCountBasis() *InterestComputationMethodFormat4Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat4Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes66) AddClassificationType() *ClassificationType32Choice {
	f.ClassificationType = new(ClassificationType32Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes66) AddOptionStyle() *OptionStyle8Choice {
	f.OptionStyle = new(OptionStyle8Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes66) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes66) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes66) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes66) AddPercentageOfDebtClaim() *RateFormat3Choice {
	f.PercentageOfDebtClaim = new(RateFormat3Choice)
	return f.PercentageOfDebtClaim
}

func (f *FinancialInstrumentAttributes66) AddPreviousFactor() *RateFormat12Choice {
	f.PreviousFactor = new(RateFormat12Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes66) AddNextFactor() *RateFormat12Choice {
	f.NextFactor = new(RateFormat12Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes66) AddWarrantParity() *QuantityToQuantityRatio1 {
	f.WarrantParity = new(QuantityToQuantityRatio1)
	return f.WarrantParity
}

func (f *FinancialInstrumentAttributes66) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes66) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

// Description of the financial instrument.
type FinancialInstrumentAttributes67 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat4Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType32Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle8Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument that may be instructed.
	MinimumQuantityToInstruct *FinancialInstrumentQuantity1Choice `xml:"MinQtyToInst,omitempty"`

	// Minimum multiple quantity of financial instrument that may be instructed.
	MinimumMultipleQuantityToInstruct *FinancialInstrumentQuantity1Choice `xml:"MinMltplQtyToInst,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Initial issue/valuation price of a resulting security under a corporate action.
	IssuePrice *PriceFormat45Choice `xml:"IssePric,omitempty"`
}

func (f *FinancialInstrumentAttributes67) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes67) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes67) AddDayCountBasis() *InterestComputationMethodFormat4Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat4Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes67) AddClassificationType() *ClassificationType32Choice {
	f.ClassificationType = new(ClassificationType32Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes67) AddOptionStyle() *OptionStyle8Choice {
	f.OptionStyle = new(OptionStyle8Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes67) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes67) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes67) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes67) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes67) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes67) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes67) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes67) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes67) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes67) AddPreviousFactor() *RateFormat12Choice {
	f.PreviousFactor = new(RateFormat12Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes67) AddNextFactor() *RateFormat12Choice {
	f.NextFactor = new(RateFormat12Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes67) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes67) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes67) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes67) AddMinimumQuantityToInstruct() *FinancialInstrumentQuantity1Choice {
	f.MinimumQuantityToInstruct = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumQuantityToInstruct
}

func (f *FinancialInstrumentAttributes67) AddMinimumMultipleQuantityToInstruct() *FinancialInstrumentQuantity1Choice {
	f.MinimumMultipleQuantityToInstruct = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumMultipleQuantityToInstruct
}

func (f *FinancialInstrumentAttributes67) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes67) AddIssuePrice() *PriceFormat45Choice {
	f.IssuePrice = new(PriceFormat45Choice)
	return f.IssuePrice
}

// Description of the financial instrument.
type FinancialInstrumentAttributes68 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId"`

	// Quantity of entitled intermediate securities based on the balance of underlying securities.
	Quantity *DecimalNumber `xml:"Qty,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat3Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType25Choice `xml:"FrctnDspstn,omitempty"`

	// Quantity of intermediate securities awarded for a given quantity of underlying security.
	IntermediateSecuritiesToUnderlyingRatio *QuantityToQuantityRatio1 `xml:"IntrmdtSctiesToUndrlygRatio,omitempty"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *AmountPrice2 `xml:"MktPric,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *DateFormat30Choice `xml:"XpryDt"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateFormat30Choice `xml:"PstngDt"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period4 `xml:"TradgPrd,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat5Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat5Choice `xml:"InstdBal,omitempty"`
}

func (f *FinancialInstrumentAttributes68) AddSecurityIdentification() *SecurityIdentification19 {
	f.SecurityIdentification = new(SecurityIdentification19)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes68) SetQuantity(value string) {
	f.Quantity = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentAttributes68) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat3Choice {
	f.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat3Choice)
	return f.RenounceableEntitlementStatusType
}

func (f *FinancialInstrumentAttributes68) AddFractionDisposition() *FractionDispositionType25Choice {
	f.FractionDisposition = new(FractionDispositionType25Choice)
	return f.FractionDisposition
}

func (f *FinancialInstrumentAttributes68) AddIntermediateSecuritiesToUnderlyingRatio() *QuantityToQuantityRatio1 {
	f.IntermediateSecuritiesToUnderlyingRatio = new(QuantityToQuantityRatio1)
	return f.IntermediateSecuritiesToUnderlyingRatio
}

func (f *FinancialInstrumentAttributes68) AddMarketPrice() *AmountPrice2 {
	f.MarketPrice = new(AmountPrice2)
	return f.MarketPrice
}

func (f *FinancialInstrumentAttributes68) AddExpiryDate() *DateFormat30Choice {
	f.ExpiryDate = new(DateFormat30Choice)
	return f.ExpiryDate
}

func (f *FinancialInstrumentAttributes68) AddPostingDate() *DateFormat30Choice {
	f.PostingDate = new(DateFormat30Choice)
	return f.PostingDate
}

func (f *FinancialInstrumentAttributes68) AddTradingPeriod() *Period4 {
	f.TradingPeriod = new(Period4)
	return f.TradingPeriod
}

func (f *FinancialInstrumentAttributes68) AddUninstructedBalance() *BalanceFormat5Choice {
	f.UninstructedBalance = new(BalanceFormat5Choice)
	return f.UninstructedBalance
}

func (f *FinancialInstrumentAttributes68) AddInstructedBalance() *BalanceFormat5Choice {
	f.InstructedBalance = new(BalanceFormat5Choice)
	return f.InstructedBalance
}

// Description of the financial instrument.
type FinancialInstrumentAttributes69 struct {

	// Identification of a financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId,omitempty"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat5Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *BaseOne14Rate `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *BaseOne14Rate `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity15Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity15Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes69) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes69) AddPlaceOfListing() *MarketIdentification4Choice {
	f.PlaceOfListing = new(MarketIdentification4Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes69) AddDayCountBasis() *InterestComputationMethodFormat5Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat5Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes69) AddClassificationType() *ClassificationType33Choice {
	f.ClassificationType = new(ClassificationType33Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes69) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes69) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOne14Rate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetNextFactor(value string) {
	f.NextFactor = (*BaseOne14Rate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes69) AddMinimumNominalQuantity() *FinancialInstrumentQuantity15Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes69) AddContractSize() *FinancialInstrumentQuantity15Choice {
	f.ContractSize = new(FinancialInstrumentQuantity15Choice)
	return f.ContractSize
}

// Description of the financial instrument.
type FinancialInstrumentAttributes7 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification2 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Percentage of the underlying assets of a fund that represents a debt and is in the scope of the EU Savings directive.
	PercentageOfDebtClaim *RateFormat3Choice `xml:"PctgOfDebtClm,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes7) AddSecurityIdentification() *SecurityIdentification11 {
	f.SecurityIdentification = new(SecurityIdentification11)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes7) AddPlaceOfListing() *MarketIdentification2 {
	f.PlaceOfListing = new(MarketIdentification2)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes7) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes7) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes7) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes7) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes7) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes7) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes7) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes7) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes7) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes7) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes7) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes7) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes7) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes7) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes7) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes7) AddPercentageOfDebtClaim() *RateFormat3Choice {
	f.PercentageOfDebtClaim = new(RateFormat3Choice)
	return f.PercentageOfDebtClaim
}

func (f *FinancialInstrumentAttributes7) AddPreviousFactor() *RateFormat3Choice {
	f.PreviousFactor = new(RateFormat3Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes7) AddNextFactor() *RateFormat3Choice {
	f.NextFactor = new(RateFormat3Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes7) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes7) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes7) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes7) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

// Description of the financial instrument.
type FinancialInstrumentAttributes70 struct {

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat5Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle9Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Percentage of the underlying assets of a fund that represents a debt, for example, in the context of the EU Savings directive.
	PercentageOfDebtClaim *RateFormat3Choice `xml:"PctgOfDebtClm,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Provides the ratio between the quantity of warrants and the quantity of underlying securities.
	WarrantParity *QuantityToQuantityRatio2 `xml:"WarrtParity,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity15Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity15Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes70) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes70) AddPlaceOfListing() *MarketIdentification4Choice {
	f.PlaceOfListing = new(MarketIdentification4Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes70) AddDayCountBasis() *InterestComputationMethodFormat5Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat5Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes70) AddClassificationType() *ClassificationType33Choice {
	f.ClassificationType = new(ClassificationType33Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes70) AddOptionStyle() *OptionStyle9Choice {
	f.OptionStyle = new(OptionStyle9Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes70) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes70) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes70) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes70) AddPercentageOfDebtClaim() *RateFormat3Choice {
	f.PercentageOfDebtClaim = new(RateFormat3Choice)
	return f.PercentageOfDebtClaim
}

func (f *FinancialInstrumentAttributes70) AddPreviousFactor() *RateFormat12Choice {
	f.PreviousFactor = new(RateFormat12Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes70) AddNextFactor() *RateFormat12Choice {
	f.NextFactor = new(RateFormat12Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes70) AddWarrantParity() *QuantityToQuantityRatio2 {
	f.WarrantParity = new(QuantityToQuantityRatio2)
	return f.WarrantParity
}

func (f *FinancialInstrumentAttributes70) AddMinimumNominalQuantity() *FinancialInstrumentQuantity15Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes70) AddContractSize() *FinancialInstrumentQuantity15Choice {
	f.ContractSize = new(FinancialInstrumentQuantity15Choice)
	return f.ContractSize
}

// Description of the financial instrument.
type FinancialInstrumentAttributes71 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat5Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle9Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity15Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument that may be instructed.
	MinimumQuantityToInstruct *FinancialInstrumentQuantity15Choice `xml:"MinQtyToInst,omitempty"`

	// Minimum multiple quantity of financial instrument that may be instructed.
	MinimumMultipleQuantityToInstruct *FinancialInstrumentQuantity15Choice `xml:"MinMltplQtyToInst,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity15Choice `xml:"CtrctSz,omitempty"`

	// Initial issue/valuation price of a resulting security under a corporate action.
	IssuePrice *PriceFormat57Choice `xml:"IssePric,omitempty"`
}

func (f *FinancialInstrumentAttributes71) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes71) AddPlaceOfListing() *MarketIdentification4Choice {
	f.PlaceOfListing = new(MarketIdentification4Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes71) AddDayCountBasis() *InterestComputationMethodFormat5Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat5Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes71) AddClassificationType() *ClassificationType33Choice {
	f.ClassificationType = new(ClassificationType33Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes71) AddOptionStyle() *OptionStyle9Choice {
	f.OptionStyle = new(OptionStyle9Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes71) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes71) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes71) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes71) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes71) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes71) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes71) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes71) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes71) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes71) AddPreviousFactor() *RateFormat12Choice {
	f.PreviousFactor = new(RateFormat12Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes71) AddNextFactor() *RateFormat12Choice {
	f.NextFactor = new(RateFormat12Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes71) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes71) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes71) AddMinimumNominalQuantity() *FinancialInstrumentQuantity15Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes71) AddMinimumQuantityToInstruct() *FinancialInstrumentQuantity15Choice {
	f.MinimumQuantityToInstruct = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumQuantityToInstruct
}

func (f *FinancialInstrumentAttributes71) AddMinimumMultipleQuantityToInstruct() *FinancialInstrumentQuantity15Choice {
	f.MinimumMultipleQuantityToInstruct = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumMultipleQuantityToInstruct
}

func (f *FinancialInstrumentAttributes71) AddContractSize() *FinancialInstrumentQuantity15Choice {
	f.ContractSize = new(FinancialInstrumentQuantity15Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes71) AddIssuePrice() *PriceFormat57Choice {
	f.IssuePrice = new(PriceFormat57Choice)
	return f.IssuePrice
}

// Description of the financial instrument.
type FinancialInstrumentAttributes73 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification20 `xml:"SctyId"`

	// Quantity of entitled intermediate securities based on the balance of underlying securities.
	Quantity *RestrictedFINDecimalNumber `xml:"Qty,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat4Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType32Choice `xml:"FrctnDspstn,omitempty"`

	// Quantity of intermediate securities awarded for a given quantity of underlying security.
	IntermediateSecuritiesToUnderlyingRatio *QuantityToQuantityRatio2 `xml:"IntrmdtSctiesToUndrlygRatio,omitempty"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *AmountPrice4 `xml:"MktPric,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *DateFormat41Choice `xml:"XpryDt"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateFormat41Choice `xml:"PstngDt"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period4 `xml:"TradgPrd,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat7Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat7Choice `xml:"InstdBal,omitempty"`
}

func (f *FinancialInstrumentAttributes73) AddSecurityIdentification() *SecurityIdentification20 {
	f.SecurityIdentification = new(SecurityIdentification20)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes73) SetQuantity(value string) {
	f.Quantity = (*RestrictedFINDecimalNumber)(&value)
}

func (f *FinancialInstrumentAttributes73) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat4Choice {
	f.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat4Choice)
	return f.RenounceableEntitlementStatusType
}

func (f *FinancialInstrumentAttributes73) AddFractionDisposition() *FractionDispositionType32Choice {
	f.FractionDisposition = new(FractionDispositionType32Choice)
	return f.FractionDisposition
}

func (f *FinancialInstrumentAttributes73) AddIntermediateSecuritiesToUnderlyingRatio() *QuantityToQuantityRatio2 {
	f.IntermediateSecuritiesToUnderlyingRatio = new(QuantityToQuantityRatio2)
	return f.IntermediateSecuritiesToUnderlyingRatio
}

func (f *FinancialInstrumentAttributes73) AddMarketPrice() *AmountPrice4 {
	f.MarketPrice = new(AmountPrice4)
	return f.MarketPrice
}

func (f *FinancialInstrumentAttributes73) AddExpiryDate() *DateFormat41Choice {
	f.ExpiryDate = new(DateFormat41Choice)
	return f.ExpiryDate
}

func (f *FinancialInstrumentAttributes73) AddPostingDate() *DateFormat41Choice {
	f.PostingDate = new(DateFormat41Choice)
	return f.PostingDate
}

func (f *FinancialInstrumentAttributes73) AddTradingPeriod() *Period4 {
	f.TradingPeriod = new(Period4)
	return f.TradingPeriod
}

func (f *FinancialInstrumentAttributes73) AddUninstructedBalance() *BalanceFormat7Choice {
	f.UninstructedBalance = new(BalanceFormat7Choice)
	return f.UninstructedBalance
}

func (f *FinancialInstrumentAttributes73) AddInstructedBalance() *BalanceFormat7Choice {
	f.InstructedBalance = new(BalanceFormat7Choice)
	return f.InstructedBalance
}

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes75 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat5Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, this is, ownership, of the security.
	RegistrationForm *FormOfSecurity7Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency27Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus6Choice `xml:"PmtSts,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency27Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle9Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType7Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Rate of return anticipated on a bond when held until maturity date.
	YieldToMaturityRate *PercentageRate `xml:"YldToMtrtyRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number23Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *GenericIdentification39 `xml:"PoolNb,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown33 `xml:"QtyBrkdwn,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType2Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price3 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price3 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price3 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price3 `xml:"StrkPric,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity15Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity15Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification32 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *RestrictedFINXMax350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes75) AddPlaceOfListing() *MarketIdentification4Choice {
	f.PlaceOfListing = new(MarketIdentification4Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes75) AddDayCountBasis() *InterestComputationMethodFormat5Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat5Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes75) AddRegistrationForm() *FormOfSecurity7Choice {
	f.RegistrationForm = new(FormOfSecurity7Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes75) AddPaymentFrequency() *Frequency27Choice {
	f.PaymentFrequency = new(Frequency27Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes75) AddPaymentStatus() *SecuritiesPaymentStatus6Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus6Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes75) AddVariableRateChangeFrequency() *Frequency27Choice {
	f.VariableRateChangeFrequency = new(Frequency27Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes75) AddClassificationType() *ClassificationType33Choice {
	f.ClassificationType = new(ClassificationType33Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes75) AddOptionStyle() *OptionStyle9Choice {
	f.OptionStyle = new(OptionStyle9Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes75) AddOptionType() *OptionType7Choice {
	f.OptionType = new(OptionType7Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes75) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes75) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetYieldToMaturityRate(value string) {
	f.YieldToMaturityRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes75) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes75) AddCouponAttachedNumber() *Number23Choice {
	f.CouponAttachedNumber = new(Number23Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes75) AddPoolNumber() *GenericIdentification39 {
	f.PoolNumber = new(GenericIdentification39)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes75) AddQuantityBreakdown() *QuantityBreakdown33 {
	newValue := new(QuantityBreakdown33)
	f.QuantityBreakdown = append(f.QuantityBreakdown, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes75) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes75) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes75) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes75) AddMarketOrIndicativePrice() *PriceType2Choice {
	f.MarketOrIndicativePrice = new(PriceType2Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes75) AddExercisePrice() *Price3 {
	f.ExercisePrice = new(Price3)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes75) AddSubscriptionPrice() *Price3 {
	f.SubscriptionPrice = new(Price3)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes75) AddConversionPrice() *Price3 {
	f.ConversionPrice = new(Price3)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes75) AddStrikePrice() *Price3 {
	f.StrikePrice = new(Price3)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes75) AddMinimumNominalQuantity() *FinancialInstrumentQuantity15Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes75) AddContractSize() *FinancialInstrumentQuantity15Choice {
	f.ContractSize = new(FinancialInstrumentQuantity15Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes75) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification32 {
	newValue := new(SecurityIdentification32)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes75) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes78 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat5Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, this is, ownership, of the security.
	RegistrationForm *FormOfSecurity7Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency27Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus6Choice `xml:"PmtSts,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency27Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle9Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType7Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Rate of return anticipated on a bond when held until maturity date.
	YieldToMaturityRate *PercentageRate `xml:"YldToMtrtyRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number23Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *GenericIdentification39 `xml:"PoolNb,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType2Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price3 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price3 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price3 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price3 `xml:"StrkPric,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity15Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity15Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification20 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *RestrictedFINXMax350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes78) AddPlaceOfListing() *MarketIdentification4Choice {
	f.PlaceOfListing = new(MarketIdentification4Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes78) AddDayCountBasis() *InterestComputationMethodFormat5Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat5Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes78) AddRegistrationForm() *FormOfSecurity7Choice {
	f.RegistrationForm = new(FormOfSecurity7Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes78) AddPaymentFrequency() *Frequency27Choice {
	f.PaymentFrequency = new(Frequency27Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes78) AddPaymentStatus() *SecuritiesPaymentStatus6Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus6Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes78) AddVariableRateChangeFrequency() *Frequency27Choice {
	f.VariableRateChangeFrequency = new(Frequency27Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes78) AddClassificationType() *ClassificationType33Choice {
	f.ClassificationType = new(ClassificationType33Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes78) AddOptionStyle() *OptionStyle9Choice {
	f.OptionStyle = new(OptionStyle9Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes78) AddOptionType() *OptionType7Choice {
	f.OptionType = new(OptionType7Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes78) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes78) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetYieldToMaturityRate(value string) {
	f.YieldToMaturityRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes78) AddCouponAttachedNumber() *Number23Choice {
	f.CouponAttachedNumber = new(Number23Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes78) AddPoolNumber() *GenericIdentification39 {
	f.PoolNumber = new(GenericIdentification39)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes78) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes78) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes78) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes78) AddMarketOrIndicativePrice() *PriceType2Choice {
	f.MarketOrIndicativePrice = new(PriceType2Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes78) AddExercisePrice() *Price3 {
	f.ExercisePrice = new(Price3)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes78) AddSubscriptionPrice() *Price3 {
	f.SubscriptionPrice = new(Price3)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes78) AddConversionPrice() *Price3 {
	f.ConversionPrice = new(Price3)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes78) AddStrikePrice() *Price3 {
	f.StrikePrice = new(Price3)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes78) AddMinimumNominalQuantity() *FinancialInstrumentQuantity15Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes78) AddContractSize() *FinancialInstrumentQuantity15Choice {
	f.ContractSize = new(FinancialInstrumentQuantity15Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes78) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification20 {
	newValue := new(SecurityIdentification20)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes78) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Description of the financial instrument.
type FinancialInstrumentAttributes79 struct {

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat4Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType32Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle8Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Annualised interest rate of a financial instrument used to calculate the actual interest rate of the coupon or the accrued interest.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Percentage of the underlying assets of a fund that represents a debt, for example, in the context of the EU Savings directive.
	PercentageOfDebtClaim *RateFormat3Choice `xml:"PctgOfDebtClm,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Provides the ratio between the quantity of warrants and the quantity of underlying securities.
	WarrantParity *QuantityToQuantityRatio1 `xml:"WarrtParity,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes79) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes79) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes79) AddDayCountBasis() *InterestComputationMethodFormat4Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat4Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes79) AddClassificationType() *ClassificationType32Choice {
	f.ClassificationType = new(ClassificationType32Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes79) AddOptionStyle() *OptionStyle8Choice {
	f.OptionStyle = new(OptionStyle8Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes79) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes79) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes79) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes79) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes79) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes79) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes79) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes79) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes79) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes79) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes79) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes79) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes79) AddPercentageOfDebtClaim() *RateFormat3Choice {
	f.PercentageOfDebtClaim = new(RateFormat3Choice)
	return f.PercentageOfDebtClaim
}

func (f *FinancialInstrumentAttributes79) AddPreviousFactor() *RateFormat12Choice {
	f.PreviousFactor = new(RateFormat12Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes79) AddNextFactor() *RateFormat12Choice {
	f.NextFactor = new(RateFormat12Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes79) AddWarrantParity() *QuantityToQuantityRatio1 {
	f.WarrantParity = new(QuantityToQuantityRatio1)
	return f.WarrantParity
}

func (f *FinancialInstrumentAttributes79) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes79) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes8 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification5 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, ie, ownership, of the security.
	RegistrationForm *FormOfSecurity2Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency3Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus2Choice `xml:"PmtSts,omitempty"`

	// Indicates the direction of payment for asset or mortgage backed securities, ie, whether the repaid capital is distributed (payment direction is down) or capitalized (payment direction is up).
	PaymentDirection *PaymentDirection2Choice `xml:"PmtDrctn,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency3Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Indicates the  level of priority to claim on income and assets of the company in case of the payment of dividends and in the event of a bankruptcy, for example, ordinary/common stocks, preferred stocks, subordinated debt, etc.
	PreferenceToIncome *PreferenceToIncome2Choice `xml:"PrefToIncm,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType2Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number2Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *Number2Choice `xml:"PoolNb,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType1Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price2 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price2 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price2 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price2 `xml:"StrkPric,omitempty"`

	// Indicates the minimum tradable quantity of a security.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security by an ISIN.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification11 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes8) AddPlaceOfListing() *MarketIdentification5 {
	f.PlaceOfListing = new(MarketIdentification5)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes8) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes8) AddRegistrationForm() *FormOfSecurity2Choice {
	f.RegistrationForm = new(FormOfSecurity2Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes8) AddPaymentFrequency() *Frequency3Choice {
	f.PaymentFrequency = new(Frequency3Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes8) AddPaymentStatus() *SecuritiesPaymentStatus2Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus2Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes8) AddPaymentDirection() *PaymentDirection2Choice {
	f.PaymentDirection = new(PaymentDirection2Choice)
	return f.PaymentDirection
}

func (f *FinancialInstrumentAttributes8) AddVariableRateChangeFrequency() *Frequency3Choice {
	f.VariableRateChangeFrequency = new(Frequency3Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes8) AddPreferenceToIncome() *PreferenceToIncome2Choice {
	f.PreferenceToIncome = new(PreferenceToIncome2Choice)
	return f.PreferenceToIncome
}

func (f *FinancialInstrumentAttributes8) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes8) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes8) AddOptionType() *OptionType2Choice {
	f.OptionType = new(OptionType2Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes8) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes8) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes8) AddCouponAttachedNumber() *Number2Choice {
	f.CouponAttachedNumber = new(Number2Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes8) AddPoolNumber() *Number2Choice {
	f.PoolNumber = new(Number2Choice)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes8) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes8) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes8) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes8) AddMarketOrIndicativePrice() *PriceType1Choice {
	f.MarketOrIndicativePrice = new(PriceType1Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes8) AddExercisePrice() *Price2 {
	f.ExercisePrice = new(Price2)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes8) AddSubscriptionPrice() *Price2 {
	f.SubscriptionPrice = new(Price2)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes8) AddConversionPrice() *Price2 {
	f.ConversionPrice = new(Price2)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes8) AddStrikePrice() *Price2 {
	f.StrikePrice = new(Price2)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes8) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes8) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes8) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification11 {
	newValue := new(SecurityIdentification11)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes8) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}

// Description of the financial instrument.
type FinancialInstrumentAttributes80 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat4Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType32Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle8Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Annualised interest rate of a financial instrument used to calculate the actual interest rate of the coupon or the accrued interest.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument that may be instructed.
	MinimumQuantityToInstruct *FinancialInstrumentQuantity1Choice `xml:"MinQtyToInst,omitempty"`

	// Minimum multiple quantity of financial instrument that may be instructed.
	MinimumMultipleQuantityToInstruct *FinancialInstrumentQuantity1Choice `xml:"MinMltplQtyToInst,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Initial issue/valuation price of a resulting security under a corporate action.
	IssuePrice *PriceFormat45Choice `xml:"IssePric,omitempty"`
}

func (f *FinancialInstrumentAttributes80) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes80) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes80) AddDayCountBasis() *InterestComputationMethodFormat4Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat4Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes80) AddClassificationType() *ClassificationType32Choice {
	f.ClassificationType = new(ClassificationType32Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes80) AddOptionStyle() *OptionStyle8Choice {
	f.OptionStyle = new(OptionStyle8Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes80) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes80) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) AddPreviousFactor() *RateFormat12Choice {
	f.PreviousFactor = new(RateFormat12Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes80) AddNextFactor() *RateFormat12Choice {
	f.NextFactor = new(RateFormat12Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes80) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes80) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes80) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes80) AddMinimumQuantityToInstruct() *FinancialInstrumentQuantity1Choice {
	f.MinimumQuantityToInstruct = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumQuantityToInstruct
}

func (f *FinancialInstrumentAttributes80) AddMinimumMultipleQuantityToInstruct() *FinancialInstrumentQuantity1Choice {
	f.MinimumMultipleQuantityToInstruct = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumMultipleQuantityToInstruct
}

func (f *FinancialInstrumentAttributes80) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes80) AddIssuePrice() *PriceFormat45Choice {
	f.IssuePrice = new(PriceFormat45Choice)
	return f.IssuePrice
}

// Description of the financial instrument.
type FinancialInstrumentAttributes81 struct {

	// Identification of a financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat4Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType32Choice `xml:"ClssfctnTp,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *BaseOne14Rate `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *BaseOne14Rate `xml:"NxtFctr,omitempty"`

	// Annualised interest rate of a financial instrument used to calculate the actual interest rate of the coupon or the accrued interest.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes81) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes81) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes81) AddDayCountBasis() *InterestComputationMethodFormat4Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat4Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes81) AddClassificationType() *ClassificationType32Choice {
	f.ClassificationType = new(ClassificationType32Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes81) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes81) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes81) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes81) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes81) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes81) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes81) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes81) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes81) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes81) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes81) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOne14Rate)(&value)
}

func (f *FinancialInstrumentAttributes81) SetNextFactor(value string) {
	f.NextFactor = (*BaseOne14Rate)(&value)
}

func (f *FinancialInstrumentAttributes81) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes81) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes81) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes81) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

// Description of the financial instrument.
type FinancialInstrumentAttributes83 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat5Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle9Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Annualised interest rate of a financial instrument used to calculate the actual interest rate of the coupon or the accrued interest.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity15Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument that may be instructed.
	MinimumQuantityToInstruct *FinancialInstrumentQuantity15Choice `xml:"MinQtyToInst,omitempty"`

	// Minimum multiple quantity of financial instrument that may be instructed.
	MinimumMultipleQuantityToInstruct *FinancialInstrumentQuantity15Choice `xml:"MinMltplQtyToInst,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity15Choice `xml:"CtrctSz,omitempty"`

	// Initial issue/valuation price of a resulting security under a corporate action.
	IssuePrice *PriceFormat57Choice `xml:"IssePric,omitempty"`
}

func (f *FinancialInstrumentAttributes83) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes83) AddPlaceOfListing() *MarketIdentification4Choice {
	f.PlaceOfListing = new(MarketIdentification4Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes83) AddDayCountBasis() *InterestComputationMethodFormat5Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat5Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes83) AddClassificationType() *ClassificationType33Choice {
	f.ClassificationType = new(ClassificationType33Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes83) AddOptionStyle() *OptionStyle9Choice {
	f.OptionStyle = new(OptionStyle9Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes83) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes83) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes83) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes83) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes83) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes83) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes83) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes83) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes83) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes83) AddPreviousFactor() *RateFormat12Choice {
	f.PreviousFactor = new(RateFormat12Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes83) AddNextFactor() *RateFormat12Choice {
	f.NextFactor = new(RateFormat12Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes83) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes83) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes83) AddMinimumNominalQuantity() *FinancialInstrumentQuantity15Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes83) AddMinimumQuantityToInstruct() *FinancialInstrumentQuantity15Choice {
	f.MinimumQuantityToInstruct = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumQuantityToInstruct
}

func (f *FinancialInstrumentAttributes83) AddMinimumMultipleQuantityToInstruct() *FinancialInstrumentQuantity15Choice {
	f.MinimumMultipleQuantityToInstruct = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumMultipleQuantityToInstruct
}

func (f *FinancialInstrumentAttributes83) AddContractSize() *FinancialInstrumentQuantity15Choice {
	f.ContractSize = new(FinancialInstrumentQuantity15Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes83) AddIssuePrice() *PriceFormat57Choice {
	f.IssuePrice = new(PriceFormat57Choice)
	return f.IssuePrice
}

// Description of the financial instrument.
type FinancialInstrumentAttributes84 struct {

	// Identification of a financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId,omitempty"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat5Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *BaseOne14Rate `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *BaseOne14Rate `xml:"NxtFctr,omitempty"`

	// Annualised interest rate of a financial instrument used to calculate the actual interest rate of the coupon or the accrued interest.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity15Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity15Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes84) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes84) AddPlaceOfListing() *MarketIdentification4Choice {
	f.PlaceOfListing = new(MarketIdentification4Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes84) AddDayCountBasis() *InterestComputationMethodFormat5Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat5Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes84) AddClassificationType() *ClassificationType33Choice {
	f.ClassificationType = new(ClassificationType33Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes84) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes84) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes84) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes84) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes84) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes84) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes84) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes84) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes84) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes84) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes84) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOne14Rate)(&value)
}

func (f *FinancialInstrumentAttributes84) SetNextFactor(value string) {
	f.NextFactor = (*BaseOne14Rate)(&value)
}

func (f *FinancialInstrumentAttributes84) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes84) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes84) AddMinimumNominalQuantity() *FinancialInstrumentQuantity15Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes84) AddContractSize() *FinancialInstrumentQuantity15Choice {
	f.ContractSize = new(FinancialInstrumentQuantity15Choice)
	return f.ContractSize
}

// Description of the financial instrument.
type FinancialInstrumentAttributes85 struct {

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat5Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle9Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Annualised interest rate of a financial instrument used to calculate the actual interest rate of the coupon or the accrued interest.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Percentage of the underlying assets of a fund that represents a debt, for example, in the context of the EU Savings directive.
	PercentageOfDebtClaim *RateFormat3Choice `xml:"PctgOfDebtClm,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Provides the ratio between the quantity of warrants and the quantity of underlying securities.
	WarrantParity *QuantityToQuantityRatio2 `xml:"WarrtParity,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity15Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity15Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes85) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes85) AddPlaceOfListing() *MarketIdentification4Choice {
	f.PlaceOfListing = new(MarketIdentification4Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes85) AddDayCountBasis() *InterestComputationMethodFormat5Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat5Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes85) AddClassificationType() *ClassificationType33Choice {
	f.ClassificationType = new(ClassificationType33Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes85) AddOptionStyle() *OptionStyle9Choice {
	f.OptionStyle = new(OptionStyle9Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes85) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes85) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes85) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes85) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes85) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes85) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes85) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes85) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes85) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes85) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes85) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes85) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes85) AddPercentageOfDebtClaim() *RateFormat3Choice {
	f.PercentageOfDebtClaim = new(RateFormat3Choice)
	return f.PercentageOfDebtClaim
}

func (f *FinancialInstrumentAttributes85) AddPreviousFactor() *RateFormat12Choice {
	f.PreviousFactor = new(RateFormat12Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes85) AddNextFactor() *RateFormat12Choice {
	f.NextFactor = new(RateFormat12Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes85) AddWarrantParity() *QuantityToQuantityRatio2 {
	f.WarrantParity = new(QuantityToQuantityRatio2)
	return f.WarrantParity
}

func (f *FinancialInstrumentAttributes85) AddMinimumNominalQuantity() *FinancialInstrumentQuantity15Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes85) AddContractSize() *FinancialInstrumentQuantity15Choice {
	f.ContractSize = new(FinancialInstrumentQuantity15Choice)
	return f.ContractSize
}
