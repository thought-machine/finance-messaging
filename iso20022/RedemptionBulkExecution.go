package iso20022

// Execution of a redemption order.
type RedemptionBulkExecution2 struct {

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *CountryCode `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1 `xml:"CxlRght,omitempty"`

	// Investment fund class to which an investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Execution of a redemption order.
	IndividualExecutionDetails []*RedemptionExecution3 `xml:"IndvExctnDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment transaction related to the execution of an investment fund transaction.
	BulkCashSettlementDetails *PaymentTransaction18 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionBulkExecution2) SetPlaceOfTrade(value string) {
	r.PlaceOfTrade = (*CountryCode)(&value)
}

func (r *RedemptionBulkExecution2) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionBulkExecution2) AddCancellationRight() *CancellationRight1 {
	r.CancellationRight = new(CancellationRight1)
	return r.CancellationRight
}

func (r *RedemptionBulkExecution2) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	r.FinancialInstrumentDetails = new(FinancialInstrument6)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionBulkExecution2) AddIndividualExecutionDetails() *RedemptionExecution3 {
	newValue := new(RedemptionExecution3)
	r.IndividualExecutionDetails = append(r.IndividualExecutionDetails, newValue)
	return newValue
}

func (r *RedemptionBulkExecution2) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (r *RedemptionBulkExecution2) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (r *RedemptionBulkExecution2) AddBulkCashSettlementDetails() *PaymentTransaction18 {
	r.BulkCashSettlementDetails = new(PaymentTransaction18)
	return r.BulkCashSettlementDetails
}

// Execution of a redemption order.
type RedemptionBulkExecution3 struct {

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

	// Execution of a redemption order.
	IndividualExecutionDetails []*RedemptionExecution5 `xml:"IndvExctnDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction22 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionBulkExecution3) SetMasterReference(value string) {
	r.MasterReference = (*Max35Text)(&value)
}

func (r *RedemptionBulkExecution3) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	r.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return r.PlaceOfTrade
}

func (r *RedemptionBulkExecution3) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionBulkExecution3) SetRequestedFutureTradeDate(value string) {
	r.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (r *RedemptionBulkExecution3) SetCancellationRight(value string) {
	r.CancellationRight = (*CancellationRight1Code)(&value)
}

func (r *RedemptionBulkExecution3) SetExtendedCancellationRight(value string) {
	r.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (r *RedemptionBulkExecution3) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	r.FinancialInstrumentDetails = new(FinancialInstrument10)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionBulkExecution3) AddIndividualExecutionDetails() *RedemptionExecution5 {
	newValue := new(RedemptionExecution5)
	r.IndividualExecutionDetails = append(r.IndividualExecutionDetails, newValue)
	return newValue
}

func (r *RedemptionBulkExecution3) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *RedemptionBulkExecution3) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (r *RedemptionBulkExecution3) SetTotalSettlementAmount(value, currency string) {
	r.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionBulkExecution3) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionBulkExecution3) AddBulkCashSettlementDetails() *PaymentTransaction22 {
	r.BulkCashSettlementDetails = new(PaymentTransaction22)
	return r.BulkCashSettlementDetails
}

// Execution of a redemption order.
type RedemptionBulkExecution5 struct {

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

	// Execution of a redemption order.
	IndividualExecutionDetails []*RedemptionExecution16 `xml:"IndvExctnDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction72 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionBulkExecution5) SetAmendmentIndicator(value string) {
	r.AmendmentIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionBulkExecution5) SetMasterReference(value string) {
	r.MasterReference = (*Max35Text)(&value)
}

func (r *RedemptionBulkExecution5) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	r.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return r.PlaceOfTrade
}

func (r *RedemptionBulkExecution5) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionBulkExecution5) SetReceivedDateTime(value string) {
	r.ReceivedDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionBulkExecution5) SetRequestedFutureTradeDate(value string) {
	r.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (r *RedemptionBulkExecution5) AddCancellationRight() *CancellationRight1Choice {
	r.CancellationRight = new(CancellationRight1Choice)
	return r.CancellationRight
}

func (r *RedemptionBulkExecution5) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	r.FinancialInstrumentDetails = new(FinancialInstrument57)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionBulkExecution5) AddIndividualExecutionDetails() *RedemptionExecution16 {
	newValue := new(RedemptionExecution16)
	r.IndividualExecutionDetails = append(r.IndividualExecutionDetails, newValue)
	return newValue
}

func (r *RedemptionBulkExecution5) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *RedemptionBulkExecution5) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (r *RedemptionBulkExecution5) SetTotalSettlementAmount(value, currency string) {
	r.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionBulkExecution5) AddBulkCashSettlementDetails() *PaymentTransaction72 {
	r.BulkCashSettlementDetails = new(PaymentTransaction72)
	return r.BulkCashSettlementDetails
}
