package iso20022

// Order to invest the investor's principal in an investment fund.
type SubscriptionBulkOrder2 struct {

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *CountryCode `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *ISODateTime `xml:"XpryDtTm,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1 `xml:"CxlRght,omitempty"`

	// Investment fund class related to an order.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Order to invest the investor's principal in an investment fund.
	IndividualOrderDetails []*SubscriptionOrder3 `xml:"IndvOrdrDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	BulkCashSettlementDetails *PaymentTransaction17 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (s *SubscriptionBulkOrder2) SetPlaceOfTrade(value string) {
	s.PlaceOfTrade = (*CountryCode)(&value)
}

func (s *SubscriptionBulkOrder2) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionBulkOrder2) SetExpiryDateTime(value string) {
	s.ExpiryDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionBulkOrder2) AddCancellationRight() *CancellationRight1 {
	s.CancellationRight = new(CancellationRight1)
	return s.CancellationRight
}

func (s *SubscriptionBulkOrder2) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	s.FinancialInstrumentDetails = new(FinancialInstrument6)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionBulkOrder2) AddIndividualOrderDetails() *SubscriptionOrder3 {
	newValue := new(SubscriptionOrder3)
	s.IndividualOrderDetails = append(s.IndividualOrderDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrder2) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (s *SubscriptionBulkOrder2) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (s *SubscriptionBulkOrder2) AddBulkCashSettlementDetails() *PaymentTransaction17 {
	s.BulkCashSettlementDetails = new(PaymentTransaction17)
	return s.BulkCashSettlementDetails
}

// Order to invest the investor's principal in an investment fund.
type SubscriptionBulkOrder3 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
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

	// Investment fund class related to an order.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls"`

	// Order to invest the investor's principal in an investment fund.
	IndividualOrderDetails []*SubscriptionOrder5 `xml:"IndvOrdrDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction23 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (s *SubscriptionBulkOrder3) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SubscriptionBulkOrder3) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return s.PlaceOfTrade
}

func (s *SubscriptionBulkOrder3) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionBulkOrder3) AddExpiryDateTime() *DateAndDateTimeChoice {
	s.ExpiryDateTime = new(DateAndDateTimeChoice)
	return s.ExpiryDateTime
}

func (s *SubscriptionBulkOrder3) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SubscriptionBulkOrder3) SetCancellationRight(value string) {
	s.CancellationRight = (*CancellationRight1Code)(&value)
}

func (s *SubscriptionBulkOrder3) SetExtendedCancellationRight(value string) {
	s.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (s *SubscriptionBulkOrder3) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	s.FinancialInstrumentDetails = new(FinancialInstrument10)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionBulkOrder3) AddIndividualOrderDetails() *SubscriptionOrder5 {
	newValue := new(SubscriptionOrder5)
	s.IndividualOrderDetails = append(s.IndividualOrderDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrder3) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SubscriptionBulkOrder3) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SubscriptionBulkOrder3) SetTotalSettlementAmount(value, currency string) {
	s.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionBulkOrder3) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionBulkOrder3) AddBulkCashSettlementDetails() *PaymentTransaction23 {
	s.BulkCashSettlementDetails = new(PaymentTransaction23)
	return s.BulkCashSettlementDetails
}

// Order to invest the investor's principal in an investment fund.
type SubscriptionBulkOrder4 struct {

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

	// Investment fund class related to an order.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls"`

	// Order to invest the investor's principal in an investment fund.
	IndividualOrderDetails []*SubscriptionOrder7 `xml:"IndvOrdrDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction23 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (s *SubscriptionBulkOrder4) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SubscriptionBulkOrder4) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return s.PlaceOfTrade
}

func (s *SubscriptionBulkOrder4) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionBulkOrder4) AddExpiryDateTime() *DateAndDateTimeChoice {
	s.ExpiryDateTime = new(DateAndDateTimeChoice)
	return s.ExpiryDateTime
}

func (s *SubscriptionBulkOrder4) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SubscriptionBulkOrder4) SetCancellationRight(value string) {
	s.CancellationRight = (*CancellationRight1Code)(&value)
}

func (s *SubscriptionBulkOrder4) SetExtendedCancellationRight(value string) {
	s.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (s *SubscriptionBulkOrder4) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	s.FinancialInstrumentDetails = new(FinancialInstrument10)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionBulkOrder4) AddIndividualOrderDetails() *SubscriptionOrder7 {
	newValue := new(SubscriptionOrder7)
	s.IndividualOrderDetails = append(s.IndividualOrderDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrder4) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SubscriptionBulkOrder4) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SubscriptionBulkOrder4) SetTotalSettlementAmount(value, currency string) {
	s.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionBulkOrder4) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionBulkOrder4) AddBulkCashSettlementDetails() *PaymentTransaction23 {
	s.BulkCashSettlementDetails = new(PaymentTransaction23)
	return s.BulkCashSettlementDetails
}

// Order to invest the investor's principal in an investment fund.
type SubscriptionBulkOrder5 struct {

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

	// Investment fund class related to the order.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls"`

	// Order to invest the investor's principal in an investment fund.
	IndividualOrderDetails []*SubscriptionOrder15 `xml:"IndvOrdrDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction70 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (s *SubscriptionBulkOrder5) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SubscriptionBulkOrder5) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return s.PlaceOfTrade
}

func (s *SubscriptionBulkOrder5) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionBulkOrder5) AddExpiryDateTime() *DateAndDateTimeChoice {
	s.ExpiryDateTime = new(DateAndDateTimeChoice)
	return s.ExpiryDateTime
}

func (s *SubscriptionBulkOrder5) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SubscriptionBulkOrder5) AddCancellationRight() *CancellationRight1Choice {
	s.CancellationRight = new(CancellationRight1Choice)
	return s.CancellationRight
}

func (s *SubscriptionBulkOrder5) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	s.FinancialInstrumentDetails = new(FinancialInstrument57)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionBulkOrder5) AddIndividualOrderDetails() *SubscriptionOrder15 {
	newValue := new(SubscriptionOrder15)
	s.IndividualOrderDetails = append(s.IndividualOrderDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrder5) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SubscriptionBulkOrder5) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SubscriptionBulkOrder5) SetTotalSettlementAmount(value, currency string) {
	s.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionBulkOrder5) AddBulkCashSettlementDetails() *PaymentTransaction70 {
	s.BulkCashSettlementDetails = new(PaymentTransaction70)
	return s.BulkCashSettlementDetails
}
