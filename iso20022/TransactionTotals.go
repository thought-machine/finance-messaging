package iso20022

// Transaction totals during the reconciliation period, for a certain type of transaction.
type TransactionTotals1 struct {

	// Identifier assigned by the merchant identifying a set of POI terminals performing some categories of transactions.
	POIGroupIdentification *Max35Text `xml:"POIGrpId,omitempty"`

	// Category of cards related the acceptance processing rules defined by the acquirer.
	CardProductProfile *Exact4NumericText `xml:"CardPdctPrfl,omitempty"`

	// Currency associated with the transaction totals.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Identification of the type of transaction.
	Type *TypeTransactionTotals1Code `xml:"Tp"`

	// Total number of transactions during a reconciliation period.
	TotalNumber *Max35NumericText `xml:"TtlNb"`

	// Total amount of a collection of transactions.
	CumulativeAmount *ImpliedCurrencyAndAmount `xml:"CmltvAmt"`
}

func (t *TransactionTotals1) SetPOIGroupIdentification(value string) {
	t.POIGroupIdentification = (*Max35Text)(&value)
}

func (t *TransactionTotals1) SetCardProductProfile(value string) {
	t.CardProductProfile = (*Exact4NumericText)(&value)
}

func (t *TransactionTotals1) SetCurrency(value string) {
	t.Currency = (*CurrencyCode)(&value)
}

func (t *TransactionTotals1) SetType(value string) {
	t.Type = (*TypeTransactionTotals1Code)(&value)
}

func (t *TransactionTotals1) SetTotalNumber(value string) {
	t.TotalNumber = (*Max35NumericText)(&value)
}

func (t *TransactionTotals1) SetCumulativeAmount(value, currency string) {
	t.CumulativeAmount = NewImpliedCurrencyAndAmount(value, currency)
}

// Transaction totals during the reconciliation period, for a certain type of transaction.
type TransactionTotals2 struct {

	// Identifier assigned by the merchant identifying a set of POI terminals performing some categories of transactions.
	POIGroupIdentification *Max35Text `xml:"POIGrpId,omitempty"`

	// Category of cards related the acceptance processing rules defined by the acquirer.
	CardProductProfile *Exact4NumericText `xml:"CardPdctPrfl,omitempty"`

	// Currency associated with the transaction totals.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Identification of the type of transaction.
	Type *TypeTransactionTotals2Code `xml:"Tp"`

	// Total number of transactions during a reconciliation period.
	TotalNumber *Max35NumericText `xml:"TtlNb"`

	// Total amount of a collection of transactions.
	CumulativeAmount *ImpliedCurrencyAndAmount `xml:"CmltvAmt"`
}

func (t *TransactionTotals2) SetPOIGroupIdentification(value string) {
	t.POIGroupIdentification = (*Max35Text)(&value)
}

func (t *TransactionTotals2) SetCardProductProfile(value string) {
	t.CardProductProfile = (*Exact4NumericText)(&value)
}

func (t *TransactionTotals2) SetCurrency(value string) {
	t.Currency = (*CurrencyCode)(&value)
}

func (t *TransactionTotals2) SetType(value string) {
	t.Type = (*TypeTransactionTotals2Code)(&value)
}

func (t *TransactionTotals2) SetTotalNumber(value string) {
	t.TotalNumber = (*Max35NumericText)(&value)
}

func (t *TransactionTotals2) SetCumulativeAmount(value, currency string) {
	t.CumulativeAmount = NewImpliedCurrencyAndAmount(value, currency)
}

// Transaction totals during the reconciliation period, for a certain type of transaction.
type TransactionTotals3 struct {

	// Identifier assigned by the merchant identifying a set of POI terminals performing some categories of transactions.
	POIGroupIdentification *Max35Text `xml:"POIGrpId,omitempty"`

	// Category of cards related the acceptance processing rules defined by the acquirer.
	CardProductProfile *Max35Text `xml:"CardPdctPrfl,omitempty"`

	// Currency associated with the transaction totals.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Identification of the type of transaction.
	Type *TypeTransactionTotals2Code `xml:"Tp"`

	// Total number of transactions during a reconciliation period.
	TotalNumber *Number `xml:"TtlNb"`

	// Total amount of a collection of transactions.
	CumulativeAmount *ImpliedCurrencyAndAmount `xml:"CmltvAmt"`
}

func (t *TransactionTotals3) SetPOIGroupIdentification(value string) {
	t.POIGroupIdentification = (*Max35Text)(&value)
}

func (t *TransactionTotals3) SetCardProductProfile(value string) {
	t.CardProductProfile = (*Max35Text)(&value)
}

func (t *TransactionTotals3) SetCurrency(value string) {
	t.Currency = (*CurrencyCode)(&value)
}

func (t *TransactionTotals3) SetType(value string) {
	t.Type = (*TypeTransactionTotals2Code)(&value)
}

func (t *TransactionTotals3) SetTotalNumber(value string) {
	t.TotalNumber = (*Number)(&value)
}

func (t *TransactionTotals3) SetCumulativeAmount(value, currency string) {
	t.CumulativeAmount = NewImpliedCurrencyAndAmount(value, currency)
}

// Totals of the reconciliation.
type TransactionTotals4 struct {

	// Total of credit transactions.
	TotalCredit *TransactionTotals5 `xml:"TtlCdt"`

	// Total of debit transactions.
	TotalDebit *TransactionTotals5 `xml:"TtlDbt"`

	// Additional count which may be utilised for reconciliation.
	TotalNumber *TransactionTotals6 `xml:"TtlNb,omitempty"`
}

func (t *TransactionTotals4) AddTotalCredit() *TransactionTotals5 {
	t.TotalCredit = new(TransactionTotals5)
	return t.TotalCredit
}

func (t *TransactionTotals4) AddTotalDebit() *TransactionTotals5 {
	t.TotalDebit = new(TransactionTotals5)
	return t.TotalDebit
}

func (t *TransactionTotals4) AddTotalNumber() *TransactionTotals6 {
	t.TotalNumber = new(TransactionTotals6)
	return t.TotalNumber
}

// Total of credit or debit transactions
type TransactionTotals5 struct {

	// Cumulative amount of all financial transactions.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Number of all financial transactions.
	Number *Number `xml:"Nb"`

	// Cumulative amount of all chargeback transactions exclusive of any fees.
	ChargeBackAmount *ImpliedCurrencyAndAmount `xml:"ChrgBckAmt"`

	// Total number of chargeback transactions.
	ChargeBackNumber *Number `xml:"ChrgBckNb"`

	// Cumulative amount of all reversal transactions exclusive of any fees.
	ReversalAmount *ImpliedCurrencyAndAmount `xml:"RvslAmt"`

	// Total number of reversal transactions.
	ReversalNumber *Number `xml:"RvslNb"`

	// Sum amount of all fees.
	FeeAmounts *ImpliedCurrencyAndAmount `xml:"FeeAmts,omitempty"`
}

func (t *TransactionTotals5) SetAmount(value, currency string) {
	t.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (t *TransactionTotals5) SetNumber(value string) {
	t.Number = (*Number)(&value)
}

func (t *TransactionTotals5) SetChargeBackAmount(value, currency string) {
	t.ChargeBackAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (t *TransactionTotals5) SetChargeBackNumber(value string) {
	t.ChargeBackNumber = (*Number)(&value)
}

func (t *TransactionTotals5) SetReversalAmount(value, currency string) {
	t.ReversalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (t *TransactionTotals5) SetReversalNumber(value string) {
	t.ReversalNumber = (*Number)(&value)
}

func (t *TransactionTotals5) SetFeeAmounts(value, currency string) {
	t.FeeAmounts = NewImpliedCurrencyAndAmount(value, currency)
}

// Additional count which may be utilised for reconciliation.
type TransactionTotals6 struct {

	// Sum number of all authorisation transactions.
	Authorisation *Number `xml:"Authstn,omitempty"`

	// Sum number of all reversed authorisation transactions.
	AuthorisationReversal *Number `xml:"AuthstnRvsl,omitempty"`

	// Sum number of all inquiry transactions.
	Inquiry *Number `xml:"Nqry,omitempty"`

	// Sum number of all reversed inquiry transactions.
	InquiryReversal *Number `xml:"NqryRvsl,omitempty"`

	// Sum number of all financial presentment payment transactions processed.
	Payments *Number `xml:"Pmts,omitempty"`

	// Sum number of all financial presentment payment transactions which have been reversed.
	PaymentReversal *Number `xml:"PmtRvsl,omitempty"`

	// Sum number of all financial presentment transactions processed.
	Transfer *Number `xml:"Trf,omitempty"`

	// Sum number of all reversal transactions processed.
	TransferReversal *Number `xml:"TrfRvsl,omitempty"`

	// Sum number of all fee collection transactions.
	FeeCollection *Number `xml:"FeeColltn,omitempty"`
}

func (t *TransactionTotals6) SetAuthorisation(value string) {
	t.Authorisation = (*Number)(&value)
}

func (t *TransactionTotals6) SetAuthorisationReversal(value string) {
	t.AuthorisationReversal = (*Number)(&value)
}

func (t *TransactionTotals6) SetInquiry(value string) {
	t.Inquiry = (*Number)(&value)
}

func (t *TransactionTotals6) SetInquiryReversal(value string) {
	t.InquiryReversal = (*Number)(&value)
}

func (t *TransactionTotals6) SetPayments(value string) {
	t.Payments = (*Number)(&value)
}

func (t *TransactionTotals6) SetPaymentReversal(value string) {
	t.PaymentReversal = (*Number)(&value)
}

func (t *TransactionTotals6) SetTransfer(value string) {
	t.Transfer = (*Number)(&value)
}

func (t *TransactionTotals6) SetTransferReversal(value string) {
	t.TransferReversal = (*Number)(&value)
}

func (t *TransactionTotals6) SetFeeCollection(value string) {
	t.FeeCollection = (*Number)(&value)
}

// Transaction totals during the reconciliation period, for a certain type of transaction.
type TransactionTotals7 struct {

	// Identifier assigned by the merchant identifying a set of POI terminals performing some categories of transactions.
	POIGroupIdentification *Max35Text `xml:"POIGrpId,omitempty"`

	// Category of cards related the acceptance processing rules defined by the acquirer.
	CardProductProfile *Max35Text `xml:"CardPdctPrfl,omitempty"`

	// Currency associated with the transaction totals.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Identification of the type of transaction.
	Type *TypeTransactionTotals2Code `xml:"Tp"`

	// Total number of transactions during a reconciliation period.
	TotalNumber *Number `xml:"TtlNb"`

	// Total amount of a collection of transactions.
	CumulativeAmount *ImpliedCurrencyAndAmount `xml:"CmltvAmt"`
}

func (t *TransactionTotals7) SetPOIGroupIdentification(value string) {
	t.POIGroupIdentification = (*Max35Text)(&value)
}

func (t *TransactionTotals7) SetCardProductProfile(value string) {
	t.CardProductProfile = (*Max35Text)(&value)
}

func (t *TransactionTotals7) SetCurrency(value string) {
	t.Currency = (*ActiveCurrencyCode)(&value)
}

func (t *TransactionTotals7) SetType(value string) {
	t.Type = (*TypeTransactionTotals2Code)(&value)
}

func (t *TransactionTotals7) SetTotalNumber(value string) {
	t.TotalNumber = (*Number)(&value)
}

func (t *TransactionTotals7) SetCumulativeAmount(value, currency string) {
	t.CumulativeAmount = NewImpliedCurrencyAndAmount(value, currency)
}
