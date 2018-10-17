package iso20022

// Execution of a subscription order.
type SubscriptionBulkExecution2 struct {

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *CountryCode `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1 `xml:"CxlRght,omitempty"`

	// Investment fund class to which an investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Execution of a subscription order.
	IndividualExecutionDetails []*SubscriptionExecution3 `xml:"IndvExctnDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	BulkCashSettlementDetails *PaymentTransaction16 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (s *SubscriptionBulkExecution2) SetPlaceOfTrade(value string) {
	s.PlaceOfTrade = (*CountryCode)(&value)
}

func (s *SubscriptionBulkExecution2) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionBulkExecution2) AddCancellationRight() *CancellationRight1 {
	s.CancellationRight = new(CancellationRight1)
	return s.CancellationRight
}

func (s *SubscriptionBulkExecution2) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	s.FinancialInstrumentDetails = new(FinancialInstrument6)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionBulkExecution2) AddIndividualExecutionDetails() *SubscriptionExecution3 {
	newValue := new(SubscriptionExecution3)
	s.IndividualExecutionDetails = append(s.IndividualExecutionDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkExecution2) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (s *SubscriptionBulkExecution2) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (s *SubscriptionBulkExecution2) AddBulkCashSettlementDetails() *PaymentTransaction16 {
	s.BulkCashSettlementDetails = new(PaymentTransaction16)
	return s.BulkCashSettlementDetails
}

// Execution of a subscription order.
type SubscriptionBulkExecution3 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *PlaceOfTradeIdentification1Choice `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Future date at which the investor requests the order to be executed.
	// The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1Code `xml:"CxlRght,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	ExtendedCancellationRight *Extended350Code `xml:"XtndedCxlRght,omitempty"`

	// Investment fund class to which an investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls"`

	// Execution of a subscription order.
	IndividualExecutionDetails []*SubscriptionExecution5 `xml:"IndvExctnDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	BulkCashSettlementDetails *PaymentTransaction24 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (s *SubscriptionBulkExecution3) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SubscriptionBulkExecution3) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return s.PlaceOfTrade
}

func (s *SubscriptionBulkExecution3) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionBulkExecution3) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SubscriptionBulkExecution3) SetCancellationRight(value string) {
	s.CancellationRight = (*CancellationRight1Code)(&value)
}

func (s *SubscriptionBulkExecution3) SetExtendedCancellationRight(value string) {
	s.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (s *SubscriptionBulkExecution3) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	s.FinancialInstrumentDetails = new(FinancialInstrument10)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionBulkExecution3) AddIndividualExecutionDetails() *SubscriptionExecution5 {
	newValue := new(SubscriptionExecution5)
	s.IndividualExecutionDetails = append(s.IndividualExecutionDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkExecution3) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SubscriptionBulkExecution3) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SubscriptionBulkExecution3) SetTotalSettlementAmount(value, currency string) {
	s.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionBulkExecution3) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionBulkExecution3) AddBulkCashSettlementDetails() *PaymentTransaction24 {
	s.BulkCashSettlementDetails = new(PaymentTransaction24)
	return s.BulkCashSettlementDetails
}

// Execution of a subscription order.
type SubscriptionBulkExecution4 struct {

	// Indicates whether the confirmation is an amendment of a previous confirmation.
	AmendmentIndicator *YesNoIndicator `xml:"AmdmntInd,omitempty"`

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *PlaceOfTradeIdentification1Choice `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor or its agent.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Date and time the order was received by the executing party, for example, the transfer agent.
	ReceivedDateTime *ISODateTime `xml:"RcvdDtTm,omitempty"`

	// Future date at which the investor requests the order to be executed.
	// The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Cancellation right of the investor with respect to the investment fund order.
	CancellationRight *CancellationRight1Choice `xml:"CxlRght,omitempty"`

	// Investment fund class to which the investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls"`

	// Execution of a subscription order.
	IndividualExecutionDetails []*SubscriptionExecution12 `xml:"IndvExctnDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction70 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (s *SubscriptionBulkExecution4) SetAmendmentIndicator(value string) {
	s.AmendmentIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionBulkExecution4) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SubscriptionBulkExecution4) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return s.PlaceOfTrade
}

func (s *SubscriptionBulkExecution4) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionBulkExecution4) SetReceivedDateTime(value string) {
	s.ReceivedDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionBulkExecution4) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SubscriptionBulkExecution4) AddCancellationRight() *CancellationRight1Choice {
	s.CancellationRight = new(CancellationRight1Choice)
	return s.CancellationRight
}

func (s *SubscriptionBulkExecution4) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	s.FinancialInstrumentDetails = new(FinancialInstrument57)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionBulkExecution4) AddIndividualExecutionDetails() *SubscriptionExecution12 {
	newValue := new(SubscriptionExecution12)
	s.IndividualExecutionDetails = append(s.IndividualExecutionDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkExecution4) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SubscriptionBulkExecution4) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SubscriptionBulkExecution4) SetTotalSettlementAmount(value, currency string) {
	s.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionBulkExecution4) AddBulkCashSettlementDetails() *PaymentTransaction70 {
	s.BulkCashSettlementDetails = new(PaymentTransaction70)
	return s.BulkCashSettlementDetails
}
