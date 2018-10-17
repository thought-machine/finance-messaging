package iso20022

// Provides transaction specific interest information that applies to the underlying transaction.
type TransactionInterest1 struct {

	// Identifies the amount of interest included in the entry amount.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Identifies whether the interest amount included in the entry amount is positive (CRDT) or negative (DBIT).
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies the type of interest.
	Type *InterestType1Choice `xml:"Tp,omitempty"`

	// Set of elements qualifying the interest rate.
	Rate []*Rate1 `xml:"Rate,omitempty"`

	// Range of time between a start date and an end date.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Underlying reason for the interest, eg, yearly credit interest on a savings account.
	Reason *Max35Text `xml:"Rsn,omitempty"`
}

func (t *TransactionInterest1) SetAmount(value, currency string) {
	t.Amount = NewCurrencyAndAmount(value, currency)
}

func (t *TransactionInterest1) SetCreditDebitIndicator(value string) {
	t.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (t *TransactionInterest1) AddType() *InterestType1Choice {
	t.Type = new(InterestType1Choice)
	return t.Type
}

func (t *TransactionInterest1) AddRate() *Rate1 {
	newValue := new(Rate1)
	t.Rate = append(t.Rate, newValue)
	return newValue
}

func (t *TransactionInterest1) AddFromToDate() *DateTimePeriodDetails {
	t.FromToDate = new(DateTimePeriodDetails)
	return t.FromToDate
}

func (t *TransactionInterest1) SetReason(value string) {
	t.Reason = (*Max35Text)(&value)
}

// Set of elements used to provide transaction specific interest information that applies to the underlying transaction.
type TransactionInterest2 struct {

	// Amount of interest included in the entry amount.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the interest amount included in the entry is credit or debit amount.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies the type of interest.
	Type *InterestType1Choice `xml:"Tp,omitempty"`

	// Set of elements used to qualify the interest rate.
	Rate []*Rate3 `xml:"Rate,omitempty"`

	// Range of time between a start date and an end date for the calculation of the interest.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Specifies the reason for the interest.
	Reason *Max35Text `xml:"Rsn,omitempty"`
}

func (t *TransactionInterest2) SetAmount(value, currency string) {
	t.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TransactionInterest2) SetCreditDebitIndicator(value string) {
	t.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (t *TransactionInterest2) AddType() *InterestType1Choice {
	t.Type = new(InterestType1Choice)
	return t.Type
}

func (t *TransactionInterest2) AddRate() *Rate3 {
	newValue := new(Rate3)
	t.Rate = append(t.Rate, newValue)
	return newValue
}

func (t *TransactionInterest2) AddFromToDate() *DateTimePeriodDetails {
	t.FromToDate = new(DateTimePeriodDetails)
	return t.FromToDate
}

func (t *TransactionInterest2) SetReason(value string) {
	t.Reason = (*Max35Text)(&value)
}

// Provide further details on transaction specific interest information that applies to the underlying transaction.
type TransactionInterest3 struct {

	// Total amount of interests and taxes included in the entry amount.
	TotalInterestAndTaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlIntrstAndTaxAmt,omitempty"`

	// Individual interest record.
	Record []*InterestRecord1 `xml:"Rcrd,omitempty"`
}

func (t *TransactionInterest3) SetTotalInterestAndTaxAmount(value, currency string) {
	t.TotalInterestAndTaxAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TransactionInterest3) AddRecord() *InterestRecord1 {
	newValue := new(InterestRecord1)
	t.Record = append(t.Record, newValue)
	return newValue
}
