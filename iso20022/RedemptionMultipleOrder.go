package iso20022

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionMultipleOrder2 struct {

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *CountryCode `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *ISODateTime `xml:"XpryDtTm,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1 `xml:"CxlRght,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson2 `xml:"BnfcryDtls,omitempty"`

	// Instruction from an investor to sell investment fund units back to the fund.
	IndividualOrderDetails []*RedemptionOrder4 `xml:"IndvOrdrDtls"`

	// Payment transaction resulting from the investment fund order execution.
	BulkCashSettlementDetails *PaymentTransaction15 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionMultipleOrder2) SetPlaceOfTrade(value string) {
	r.PlaceOfTrade = (*CountryCode)(&value)
}

func (r *RedemptionMultipleOrder2) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionMultipleOrder2) SetExpiryDateTime(value string) {
	r.ExpiryDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionMultipleOrder2) AddCancellationRight() *CancellationRight1 {
	r.CancellationRight = new(CancellationRight1)
	return r.CancellationRight
}

func (r *RedemptionMultipleOrder2) AddInvestmentAccountDetails() *InvestmentAccount13 {
	r.InvestmentAccountDetails = new(InvestmentAccount13)
	return r.InvestmentAccountDetails
}

func (r *RedemptionMultipleOrder2) AddBeneficiaryDetails() *IndividualPerson2 {
	r.BeneficiaryDetails = new(IndividualPerson2)
	return r.BeneficiaryDetails
}

func (r *RedemptionMultipleOrder2) AddIndividualOrderDetails() *RedemptionOrder4 {
	newValue := new(RedemptionOrder4)
	r.IndividualOrderDetails = append(r.IndividualOrderDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleOrder2) AddBulkCashSettlementDetails() *PaymentTransaction15 {
	r.BulkCashSettlementDetails = new(PaymentTransaction15)
	return r.BulkCashSettlementDetails
}

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionMultipleOrder3 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *PlaceOfTradeIdentification1Choice `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *DateAndDateTimeChoice `xml:"XpryDtTm,omitempty"`

	// Future date at which the investor requests the order to be executed.
	// The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1Code `xml:"CxlRght,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	ExtendedCancellationRight *Extended350Code `xml:"XtndedCxlRght,omitempty"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson12 `xml:"BnfcryDtls,omitempty"`

	// Instruction from an investor to sell investment fund units back to the fund.
	IndividualOrderDetails []*RedemptionOrder6 `xml:"IndvOrdrDtls"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction21 `xml:"BlkCshSttlmDtls,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`
}

func (r *RedemptionMultipleOrder3) SetMasterReference(value string) {
	r.MasterReference = (*Max35Text)(&value)
}

func (r *RedemptionMultipleOrder3) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	r.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return r.PlaceOfTrade
}

func (r *RedemptionMultipleOrder3) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionMultipleOrder3) AddExpiryDateTime() *DateAndDateTimeChoice {
	r.ExpiryDateTime = new(DateAndDateTimeChoice)
	return r.ExpiryDateTime
}

func (r *RedemptionMultipleOrder3) SetRequestedFutureTradeDate(value string) {
	r.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (r *RedemptionMultipleOrder3) AddInvestmentAccountDetails() *InvestmentAccount21 {
	r.InvestmentAccountDetails = new(InvestmentAccount21)
	return r.InvestmentAccountDetails
}

func (r *RedemptionMultipleOrder3) SetCancellationRight(value string) {
	r.CancellationRight = (*CancellationRight1Code)(&value)
}

func (r *RedemptionMultipleOrder3) SetExtendedCancellationRight(value string) {
	r.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (r *RedemptionMultipleOrder3) AddBeneficiaryDetails() *IndividualPerson12 {
	r.BeneficiaryDetails = new(IndividualPerson12)
	return r.BeneficiaryDetails
}

func (r *RedemptionMultipleOrder3) AddIndividualOrderDetails() *RedemptionOrder6 {
	newValue := new(RedemptionOrder6)
	r.IndividualOrderDetails = append(r.IndividualOrderDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleOrder3) AddBulkCashSettlementDetails() *PaymentTransaction21 {
	r.BulkCashSettlementDetails = new(PaymentTransaction21)
	return r.BulkCashSettlementDetails
}

func (r *RedemptionMultipleOrder3) SetTotalSettlementAmount(value, currency string) {
	r.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionMultipleOrder3) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionMultipleOrder4 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *PlaceOfTradeIdentification1Choice `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *DateAndDateTimeChoice `xml:"XpryDtTm,omitempty"`

	// Future date at which the investor requests the order to be executed.
	// The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1Code `xml:"CxlRght,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	ExtendedCancellationRight *Extended350Code `xml:"XtndedCxlRght,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson12 `xml:"BnfcryDtls,omitempty"`

	// Instruction from an investor to sell investment fund units back to the fund.
	IndividualOrderDetails []*RedemptionOrder8 `xml:"IndvOrdrDtls"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction21 `xml:"BlkCshSttlmDtls,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`
}

func (r *RedemptionMultipleOrder4) SetMasterReference(value string) {
	r.MasterReference = (*Max35Text)(&value)
}

func (r *RedemptionMultipleOrder4) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	r.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return r.PlaceOfTrade
}

func (r *RedemptionMultipleOrder4) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionMultipleOrder4) AddExpiryDateTime() *DateAndDateTimeChoice {
	r.ExpiryDateTime = new(DateAndDateTimeChoice)
	return r.ExpiryDateTime
}

func (r *RedemptionMultipleOrder4) SetRequestedFutureTradeDate(value string) {
	r.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (r *RedemptionMultipleOrder4) SetCancellationRight(value string) {
	r.CancellationRight = (*CancellationRight1Code)(&value)
}

func (r *RedemptionMultipleOrder4) SetExtendedCancellationRight(value string) {
	r.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (r *RedemptionMultipleOrder4) AddInvestmentAccountDetails() *InvestmentAccount21 {
	r.InvestmentAccountDetails = new(InvestmentAccount21)
	return r.InvestmentAccountDetails
}

func (r *RedemptionMultipleOrder4) AddBeneficiaryDetails() *IndividualPerson12 {
	r.BeneficiaryDetails = new(IndividualPerson12)
	return r.BeneficiaryDetails
}

func (r *RedemptionMultipleOrder4) AddIndividualOrderDetails() *RedemptionOrder8 {
	newValue := new(RedemptionOrder8)
	r.IndividualOrderDetails = append(r.IndividualOrderDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleOrder4) AddBulkCashSettlementDetails() *PaymentTransaction21 {
	r.BulkCashSettlementDetails = new(PaymentTransaction21)
	return r.BulkCashSettlementDetails
}

func (r *RedemptionMultipleOrder4) SetTotalSettlementAmount(value, currency string) {
	r.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionMultipleOrder4) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionMultipleOrder6 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *PlaceOfTradeIdentification1Choice `xml:"PlcOfTrad,omitempty"`

	// Date and time the order is placed by the investor or its agent.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *DateAndDateTimeChoice `xml:"XpryDtTm,omitempty"`

	// Future date at which the investor requests the order to be executed.
	// The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Cancellation right of the investor with respect to the investment fund order.
	CancellationRight *CancellationRight1Choice `xml:"CxlRght,omitempty"`

	// Account impacted by the investment fund order.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails []*IndividualPerson32 `xml:"BnfcryDtls,omitempty"`

	// Instruction from an investor to sell investment fund units back to the fund.
	IndividualOrderDetails []*RedemptionOrder14 `xml:"IndvOrdrDtls"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction72 `xml:"BlkCshSttlmDtls,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`
}

func (r *RedemptionMultipleOrder6) SetMasterReference(value string) {
	r.MasterReference = (*Max35Text)(&value)
}

func (r *RedemptionMultipleOrder6) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	r.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return r.PlaceOfTrade
}

func (r *RedemptionMultipleOrder6) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionMultipleOrder6) AddExpiryDateTime() *DateAndDateTimeChoice {
	r.ExpiryDateTime = new(DateAndDateTimeChoice)
	return r.ExpiryDateTime
}

func (r *RedemptionMultipleOrder6) SetRequestedFutureTradeDate(value string) {
	r.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (r *RedemptionMultipleOrder6) AddCancellationRight() *CancellationRight1Choice {
	r.CancellationRight = new(CancellationRight1Choice)
	return r.CancellationRight
}

func (r *RedemptionMultipleOrder6) AddInvestmentAccountDetails() *InvestmentAccount58 {
	r.InvestmentAccountDetails = new(InvestmentAccount58)
	return r.InvestmentAccountDetails
}

func (r *RedemptionMultipleOrder6) AddBeneficiaryDetails() *IndividualPerson32 {
	newValue := new(IndividualPerson32)
	r.BeneficiaryDetails = append(r.BeneficiaryDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleOrder6) AddIndividualOrderDetails() *RedemptionOrder14 {
	newValue := new(RedemptionOrder14)
	r.IndividualOrderDetails = append(r.IndividualOrderDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleOrder6) AddBulkCashSettlementDetails() *PaymentTransaction72 {
	r.BulkCashSettlementDetails = new(PaymentTransaction72)
	return r.BulkCashSettlementDetails
}

func (r *RedemptionMultipleOrder6) SetTotalSettlementAmount(value, currency string) {
	r.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}
