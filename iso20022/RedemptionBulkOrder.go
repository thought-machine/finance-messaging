package iso20022

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionBulkOrder2 struct {

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

	// Instruction from an investor to sell investment fund units back to the fund.
	IndividualOrderDetails []*RedemptionOrder3 `xml:"IndvOrdrDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	BulkCashSettlementDetails *PaymentTransaction18 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionBulkOrder2) SetPlaceOfTrade(value string) {
	r.PlaceOfTrade = (*CountryCode)(&value)
}

func (r *RedemptionBulkOrder2) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionBulkOrder2) SetExpiryDateTime(value string) {
	r.ExpiryDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionBulkOrder2) AddCancellationRight() *CancellationRight1 {
	r.CancellationRight = new(CancellationRight1)
	return r.CancellationRight
}

func (r *RedemptionBulkOrder2) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	r.FinancialInstrumentDetails = new(FinancialInstrument6)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionBulkOrder2) AddIndividualOrderDetails() *RedemptionOrder3 {
	newValue := new(RedemptionOrder3)
	r.IndividualOrderDetails = append(r.IndividualOrderDetails, newValue)
	return newValue
}

func (r *RedemptionBulkOrder2) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (r *RedemptionBulkOrder2) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (r *RedemptionBulkOrder2) AddBulkCashSettlementDetails() *PaymentTransaction18 {
	r.BulkCashSettlementDetails = new(PaymentTransaction18)
	return r.BulkCashSettlementDetails
}

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionBulkOrder3 struct {

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

	// Instruction from an investor to sell investment fund units back to the fund.
	IndividualOrderDetails []*RedemptionOrder5 `xml:"IndvOrdrDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction21 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionBulkOrder3) SetMasterReference(value string) {
	r.MasterReference = (*Max35Text)(&value)
}

func (r *RedemptionBulkOrder3) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	r.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return r.PlaceOfTrade
}

func (r *RedemptionBulkOrder3) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionBulkOrder3) AddExpiryDateTime() *DateAndDateTimeChoice {
	r.ExpiryDateTime = new(DateAndDateTimeChoice)
	return r.ExpiryDateTime
}

func (r *RedemptionBulkOrder3) SetRequestedFutureTradeDate(value string) {
	r.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (r *RedemptionBulkOrder3) SetCancellationRight(value string) {
	r.CancellationRight = (*CancellationRight1Code)(&value)
}

func (r *RedemptionBulkOrder3) SetExtendedCancellationRight(value string) {
	r.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (r *RedemptionBulkOrder3) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	r.FinancialInstrumentDetails = new(FinancialInstrument10)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionBulkOrder3) AddIndividualOrderDetails() *RedemptionOrder5 {
	newValue := new(RedemptionOrder5)
	r.IndividualOrderDetails = append(r.IndividualOrderDetails, newValue)
	return newValue
}

func (r *RedemptionBulkOrder3) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *RedemptionBulkOrder3) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (r *RedemptionBulkOrder3) SetTotalSettlementAmount(value, currency string) {
	r.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionBulkOrder3) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionBulkOrder3) AddBulkCashSettlementDetails() *PaymentTransaction21 {
	r.BulkCashSettlementDetails = new(PaymentTransaction21)
	return r.BulkCashSettlementDetails
}

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionBulkOrder4 struct {

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

	// Instruction from an investor to sell investment fund units back to the fund.
	IndividualOrderDetails []*RedemptionOrder7 `xml:"IndvOrdrDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction21 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionBulkOrder4) SetMasterReference(value string) {
	r.MasterReference = (*Max35Text)(&value)
}

func (r *RedemptionBulkOrder4) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	r.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return r.PlaceOfTrade
}

func (r *RedemptionBulkOrder4) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionBulkOrder4) AddExpiryDateTime() *DateAndDateTimeChoice {
	r.ExpiryDateTime = new(DateAndDateTimeChoice)
	return r.ExpiryDateTime
}

func (r *RedemptionBulkOrder4) SetRequestedFutureTradeDate(value string) {
	r.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (r *RedemptionBulkOrder4) SetCancellationRight(value string) {
	r.CancellationRight = (*CancellationRight1Code)(&value)
}

func (r *RedemptionBulkOrder4) SetExtendedCancellationRight(value string) {
	r.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (r *RedemptionBulkOrder4) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	r.FinancialInstrumentDetails = new(FinancialInstrument10)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionBulkOrder4) AddIndividualOrderDetails() *RedemptionOrder7 {
	newValue := new(RedemptionOrder7)
	r.IndividualOrderDetails = append(r.IndividualOrderDetails, newValue)
	return newValue
}

func (r *RedemptionBulkOrder4) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *RedemptionBulkOrder4) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (r *RedemptionBulkOrder4) SetTotalSettlementAmount(value, currency string) {
	r.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionBulkOrder4) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionBulkOrder4) AddBulkCashSettlementDetails() *PaymentTransaction21 {
	r.BulkCashSettlementDetails = new(PaymentTransaction21)
	return r.BulkCashSettlementDetails
}

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionBulkOrder6 struct {

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

	// Instruction from an investor to sell investment fund units back to the fund.
	IndividualOrderDetails []*RedemptionOrder15 `xml:"IndvOrdrDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	BulkCashSettlementDetails *PaymentTransaction72 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionBulkOrder6) SetMasterReference(value string) {
	r.MasterReference = (*Max35Text)(&value)
}

func (r *RedemptionBulkOrder6) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	r.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return r.PlaceOfTrade
}

func (r *RedemptionBulkOrder6) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionBulkOrder6) AddExpiryDateTime() *DateAndDateTimeChoice {
	r.ExpiryDateTime = new(DateAndDateTimeChoice)
	return r.ExpiryDateTime
}

func (r *RedemptionBulkOrder6) SetRequestedFutureTradeDate(value string) {
	r.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (r *RedemptionBulkOrder6) AddCancellationRight() *CancellationRight1Choice {
	r.CancellationRight = new(CancellationRight1Choice)
	return r.CancellationRight
}

func (r *RedemptionBulkOrder6) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	r.FinancialInstrumentDetails = new(FinancialInstrument57)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionBulkOrder6) AddIndividualOrderDetails() *RedemptionOrder15 {
	newValue := new(RedemptionOrder15)
	r.IndividualOrderDetails = append(r.IndividualOrderDetails, newValue)
	return newValue
}

func (r *RedemptionBulkOrder6) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *RedemptionBulkOrder6) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (r *RedemptionBulkOrder6) SetTotalSettlementAmount(value, currency string) {
	r.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionBulkOrder6) AddBulkCashSettlementDetails() *PaymentTransaction72 {
	r.BulkCashSettlementDetails = new(PaymentTransaction72)
	return r.BulkCashSettlementDetails
}
