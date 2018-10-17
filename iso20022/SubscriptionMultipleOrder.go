package iso20022

// Order to invest the investor's principal in an investment fund.
type SubscriptionMultipleOrder2 struct {

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *CountryCode `xml:"PlcOfTrad,omitempty"`

	// Date the investor places the order.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *ISODateTime `xml:"XpryDtTm,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1 `xml:"CxlRght,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Additional information about the beneficial owner.
	BeneficiaryDetails *IndividualPerson2 `xml:"BnfcryDtls,omitempty"`

	// Order to invest the investor's principal in an investment fund.
	IndividualOrderDetails []*SubscriptionOrder4 `xml:"IndvOrdrDtls"`

	// Payment transaction resulting from the investment fund order execution.
	BulkCashSettlementDetails *PaymentTransaction19 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (s *SubscriptionMultipleOrder2) SetPlaceOfTrade(value string) {
	s.PlaceOfTrade = (*CountryCode)(&value)
}

func (s *SubscriptionMultipleOrder2) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionMultipleOrder2) SetExpiryDateTime(value string) {
	s.ExpiryDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionMultipleOrder2) AddCancellationRight() *CancellationRight1 {
	s.CancellationRight = new(CancellationRight1)
	return s.CancellationRight
}

func (s *SubscriptionMultipleOrder2) AddInvestmentAccountDetails() *InvestmentAccount13 {
	s.InvestmentAccountDetails = new(InvestmentAccount13)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionMultipleOrder2) AddBeneficiaryDetails() *IndividualPerson2 {
	s.BeneficiaryDetails = new(IndividualPerson2)
	return s.BeneficiaryDetails
}

func (s *SubscriptionMultipleOrder2) AddIndividualOrderDetails() *SubscriptionOrder4 {
	newValue := new(SubscriptionOrder4)
	s.IndividualOrderDetails = append(s.IndividualOrderDetails, newValue)
	return newValue
}

func (s *SubscriptionMultipleOrder2) AddBulkCashSettlementDetails() *PaymentTransaction19 {
	s.BulkCashSettlementDetails = new(PaymentTransaction19)
	return s.BulkCashSettlementDetails
}

// Order to invest the investor's principal in an investment fund.
type SubscriptionMultipleOrder3 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *PlaceOfTradeIdentification1Choice `xml:"PlcOfTrad,omitempty"`

	// Date the investor places the order.
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

	// Additional information about the beneficial owner.
	BeneficiaryDetails *IndividualPerson9 `xml:"BnfcryDtls,omitempty"`

	// Order to invest the investor's principal in an investment fund.
	IndividualOrderDetails []*SubscriptionOrder6 `xml:"IndvOrdrDtls"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction23 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (s *SubscriptionMultipleOrder3) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SubscriptionMultipleOrder3) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return s.PlaceOfTrade
}

func (s *SubscriptionMultipleOrder3) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionMultipleOrder3) AddExpiryDateTime() *DateAndDateTimeChoice {
	s.ExpiryDateTime = new(DateAndDateTimeChoice)
	return s.ExpiryDateTime
}

func (s *SubscriptionMultipleOrder3) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SubscriptionMultipleOrder3) SetCancellationRight(value string) {
	s.CancellationRight = (*CancellationRight1Code)(&value)
}

func (s *SubscriptionMultipleOrder3) SetExtendedCancellationRight(value string) {
	s.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (s *SubscriptionMultipleOrder3) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionMultipleOrder3) AddBeneficiaryDetails() *IndividualPerson9 {
	s.BeneficiaryDetails = new(IndividualPerson9)
	return s.BeneficiaryDetails
}

func (s *SubscriptionMultipleOrder3) AddIndividualOrderDetails() *SubscriptionOrder6 {
	newValue := new(SubscriptionOrder6)
	s.IndividualOrderDetails = append(s.IndividualOrderDetails, newValue)
	return newValue
}

func (s *SubscriptionMultipleOrder3) SetTotalSettlementAmount(value, currency string) {
	s.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionMultipleOrder3) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionMultipleOrder3) AddBulkCashSettlementDetails() *PaymentTransaction23 {
	s.BulkCashSettlementDetails = new(PaymentTransaction23)
	return s.BulkCashSettlementDetails
}

// Order to invest the investor's principal in an investment fund.
type SubscriptionMultipleOrder4 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *PlaceOfTradeIdentification1Choice `xml:"PlcOfTrad,omitempty"`

	// Date the investor places the order.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *DateAndDateTimeChoice `xml:"XpryDtTm,omitempty"`

	// Future date at which the investor requests the order to be executed. The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1Code `xml:"CxlRght,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	ExtendedCancellationRight *Extended350Code `xml:"XtndedCxlRght,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls"`

	// Additional information about the beneficial owner.
	BeneficiaryDetails *IndividualPerson9 `xml:"BnfcryDtls,omitempty"`

	// Order to invest the investor's principal in an investment fund.
	IndividualOrderDetails []*SubscriptionOrder8 `xml:"IndvOrdrDtls"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction23 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (s *SubscriptionMultipleOrder4) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SubscriptionMultipleOrder4) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return s.PlaceOfTrade
}

func (s *SubscriptionMultipleOrder4) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionMultipleOrder4) AddExpiryDateTime() *DateAndDateTimeChoice {
	s.ExpiryDateTime = new(DateAndDateTimeChoice)
	return s.ExpiryDateTime
}

func (s *SubscriptionMultipleOrder4) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SubscriptionMultipleOrder4) SetCancellationRight(value string) {
	s.CancellationRight = (*CancellationRight1Code)(&value)
}

func (s *SubscriptionMultipleOrder4) SetExtendedCancellationRight(value string) {
	s.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (s *SubscriptionMultipleOrder4) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionMultipleOrder4) AddBeneficiaryDetails() *IndividualPerson9 {
	s.BeneficiaryDetails = new(IndividualPerson9)
	return s.BeneficiaryDetails
}

func (s *SubscriptionMultipleOrder4) AddIndividualOrderDetails() *SubscriptionOrder8 {
	newValue := new(SubscriptionOrder8)
	s.IndividualOrderDetails = append(s.IndividualOrderDetails, newValue)
	return newValue
}

func (s *SubscriptionMultipleOrder4) SetTotalSettlementAmount(value, currency string) {
	s.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionMultipleOrder4) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionMultipleOrder4) AddBulkCashSettlementDetails() *PaymentTransaction23 {
	s.BulkCashSettlementDetails = new(PaymentTransaction23)
	return s.BulkCashSettlementDetails
}

// Order to invest the investor's principal in an investment fund.
type SubscriptionMultipleOrder6 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *PlaceOfTradeIdentification1Choice `xml:"PlcOfTrad,omitempty"`

	// Date and time the order is placed by the investor or its agent.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *DateAndDateTimeChoice `xml:"XpryDtTm,omitempty"`

	// Future date at which the investor requests the order to be executed. The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Cancellation right of the investor with respect to the investment fund order.
	CancellationRight *CancellationRight1Choice `xml:"CxlRght,omitempty"`

	// Account impacted by the investment fund order.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails []*IndividualPerson31 `xml:"BnfcryDtls,omitempty"`

	// Order to invest the investor's principal in an investment fund.
	IndividualOrderDetails []*SubscriptionOrder14 `xml:"IndvOrdrDtls"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction70 `xml:"BlkCshSttlmDtls,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`
}

func (s *SubscriptionMultipleOrder6) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SubscriptionMultipleOrder6) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return s.PlaceOfTrade
}

func (s *SubscriptionMultipleOrder6) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionMultipleOrder6) AddExpiryDateTime() *DateAndDateTimeChoice {
	s.ExpiryDateTime = new(DateAndDateTimeChoice)
	return s.ExpiryDateTime
}

func (s *SubscriptionMultipleOrder6) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SubscriptionMultipleOrder6) AddCancellationRight() *CancellationRight1Choice {
	s.CancellationRight = new(CancellationRight1Choice)
	return s.CancellationRight
}

func (s *SubscriptionMultipleOrder6) AddInvestmentAccountDetails() *InvestmentAccount58 {
	s.InvestmentAccountDetails = new(InvestmentAccount58)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionMultipleOrder6) AddBeneficiaryDetails() *IndividualPerson31 {
	newValue := new(IndividualPerson31)
	s.BeneficiaryDetails = append(s.BeneficiaryDetails, newValue)
	return newValue
}

func (s *SubscriptionMultipleOrder6) AddIndividualOrderDetails() *SubscriptionOrder14 {
	newValue := new(SubscriptionOrder14)
	s.IndividualOrderDetails = append(s.IndividualOrderDetails, newValue)
	return newValue
}

func (s *SubscriptionMultipleOrder6) AddBulkCashSettlementDetails() *PaymentTransaction70 {
	s.BulkCashSettlementDetails = new(PaymentTransaction70)
	return s.BulkCashSettlementDetails
}

func (s *SubscriptionMultipleOrder6) SetTotalSettlementAmount(value, currency string) {
	s.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}
