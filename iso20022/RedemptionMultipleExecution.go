package iso20022

// Execution of a redemption order.
type RedemptionMultipleExecution2 struct {

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson2 `xml:"BnfcryDtls,omitempty"`

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *CountryCode `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1 `xml:"CxlRght,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Execution of a redemption order.
	IndividualExecutionDetails []*RedemptionExecution4 `xml:"IndvExctnDtls"`

	// Payment transaction resulting from the investment fund order execution.
	BulkCashSettlementDetails *PaymentTransaction15 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionMultipleExecution2) AddBeneficiaryDetails() *IndividualPerson2 {
	r.BeneficiaryDetails = new(IndividualPerson2)
	return r.BeneficiaryDetails
}

func (r *RedemptionMultipleExecution2) SetPlaceOfTrade(value string) {
	r.PlaceOfTrade = (*CountryCode)(&value)
}

func (r *RedemptionMultipleExecution2) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionMultipleExecution2) AddCancellationRight() *CancellationRight1 {
	r.CancellationRight = new(CancellationRight1)
	return r.CancellationRight
}

func (r *RedemptionMultipleExecution2) AddInvestmentAccountDetails() *InvestmentAccount13 {
	r.InvestmentAccountDetails = new(InvestmentAccount13)
	return r.InvestmentAccountDetails
}

func (r *RedemptionMultipleExecution2) AddIndividualExecutionDetails() *RedemptionExecution4 {
	newValue := new(RedemptionExecution4)
	r.IndividualExecutionDetails = append(r.IndividualExecutionDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleExecution2) AddBulkCashSettlementDetails() *PaymentTransaction15 {
	r.BulkCashSettlementDetails = new(PaymentTransaction15)
	return r.BulkCashSettlementDetails
}

// Execution of a redemption order.
type RedemptionMultipleExecution3 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson12 `xml:"BnfcryDtls,omitempty"`

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

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls"`

	// Execution of a redemption order.
	IndividualExecutionDetails []*RedemptionExecution6 `xml:"IndvExctnDtls"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	BulkCashSettlementDetails *PaymentTransaction22 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionMultipleExecution3) SetMasterReference(value string) {
	r.MasterReference = (*Max35Text)(&value)
}

func (r *RedemptionMultipleExecution3) AddBeneficiaryDetails() *IndividualPerson12 {
	r.BeneficiaryDetails = new(IndividualPerson12)
	return r.BeneficiaryDetails
}

func (r *RedemptionMultipleExecution3) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	r.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return r.PlaceOfTrade
}

func (r *RedemptionMultipleExecution3) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionMultipleExecution3) SetRequestedFutureTradeDate(value string) {
	r.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (r *RedemptionMultipleExecution3) SetCancellationRight(value string) {
	r.CancellationRight = (*CancellationRight1Code)(&value)
}

func (r *RedemptionMultipleExecution3) SetExtendedCancellationRight(value string) {
	r.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (r *RedemptionMultipleExecution3) AddInvestmentAccountDetails() *InvestmentAccount21 {
	r.InvestmentAccountDetails = new(InvestmentAccount21)
	return r.InvestmentAccountDetails
}

func (r *RedemptionMultipleExecution3) AddIndividualExecutionDetails() *RedemptionExecution6 {
	newValue := new(RedemptionExecution6)
	r.IndividualExecutionDetails = append(r.IndividualExecutionDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleExecution3) SetTotalSettlementAmount(value, currency string) {
	r.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionMultipleExecution3) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionMultipleExecution3) AddBulkCashSettlementDetails() *PaymentTransaction22 {
	r.BulkCashSettlementDetails = new(PaymentTransaction22)
	return r.BulkCashSettlementDetails
}

// Execution of a redemption order.
type RedemptionMultipleExecution5 struct {

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

	// Account impacted by the investment fund order execution.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails []*IndividualPerson32 `xml:"BnfcryDtls,omitempty"`

	// Execution of a redemption order.
	IndividualExecutionDetails []*RedemptionExecution15 `xml:"IndvExctnDtls"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction72 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionMultipleExecution5) SetAmendmentIndicator(value string) {
	r.AmendmentIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionMultipleExecution5) SetMasterReference(value string) {
	r.MasterReference = (*Max35Text)(&value)
}

func (r *RedemptionMultipleExecution5) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	r.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return r.PlaceOfTrade
}

func (r *RedemptionMultipleExecution5) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionMultipleExecution5) SetReceivedDateTime(value string) {
	r.ReceivedDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionMultipleExecution5) SetRequestedFutureTradeDate(value string) {
	r.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (r *RedemptionMultipleExecution5) AddCancellationRight() *CancellationRight1Choice {
	r.CancellationRight = new(CancellationRight1Choice)
	return r.CancellationRight
}

func (r *RedemptionMultipleExecution5) AddInvestmentAccountDetails() *InvestmentAccount58 {
	r.InvestmentAccountDetails = new(InvestmentAccount58)
	return r.InvestmentAccountDetails
}

func (r *RedemptionMultipleExecution5) AddBeneficiaryDetails() *IndividualPerson32 {
	newValue := new(IndividualPerson32)
	r.BeneficiaryDetails = append(r.BeneficiaryDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleExecution5) AddIndividualExecutionDetails() *RedemptionExecution15 {
	newValue := new(RedemptionExecution15)
	r.IndividualExecutionDetails = append(r.IndividualExecutionDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleExecution5) SetTotalSettlementAmount(value, currency string) {
	r.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionMultipleExecution5) AddBulkCashSettlementDetails() *PaymentTransaction72 {
	r.BulkCashSettlementDetails = new(PaymentTransaction72)
	return r.BulkCashSettlementDetails
}
