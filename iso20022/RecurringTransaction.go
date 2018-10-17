package iso20022

// Financial loan (instalment) or a recurring transaction.
type RecurringTransaction1 struct {

	// Indicates the recurring/instalment occurrence of the transaction (1 = 1st instalment, 2 = 2nd instalment, etc.).
	SequenceNumber *Max2NumericText `xml:"SeqNb"`

	// Period unit between consecutive payments (for example day, month, year).
	PeriodUnit *Frequency4Code `xml:"PrdUnit"`

	// Number of period units between consecutive payments.
	InstalmentPeriod *Number `xml:"InstlmtPrd"`

	// Total number of instalment payments.
	TotalNumberOfPayments *Number `xml:"TtlNbOfPmts"`

	// Interest charged in percentage for the total amount of payments.
	InterestCharges *ImpliedCurrencyAndAmount `xml:"IntrstChrgs,omitempty"`
}

func (r *RecurringTransaction1) SetSequenceNumber(value string) {
	r.SequenceNumber = (*Max2NumericText)(&value)
}

func (r *RecurringTransaction1) SetPeriodUnit(value string) {
	r.PeriodUnit = (*Frequency4Code)(&value)
}

func (r *RecurringTransaction1) SetInstalmentPeriod(value string) {
	r.InstalmentPeriod = (*Number)(&value)
}

func (r *RecurringTransaction1) SetTotalNumberOfPayments(value string) {
	r.TotalNumberOfPayments = (*Number)(&value)
}

func (r *RecurringTransaction1) SetInterestCharges(value, currency string) {
	r.InterestCharges = NewImpliedCurrencyAndAmount(value, currency)
}

// Financial loan (instalment) or a recurring transaction.
type RecurringTransaction2 struct {

	// Type of instalment plan.
	InstalmentPlan []*InstalmentPlan1Code `xml:"InstlmtPlan,omitempty"`

	// Identification of the instalment plan.
	PlanIdentification *Max35Text `xml:"PlanId,omitempty"`

	// Indicates the recurring/instalment occurrence of the transaction (1 = 1st instalment, 2 = 2nd instalment, etc.).
	SequenceNumber *Number `xml:"SeqNb,omitempty"`

	// Period unit between consecutive payments (for example day, month, year).
	PeriodUnit *Frequency3Code `xml:"PrdUnit,omitempty"`

	// Number of period units between consecutive payments.
	InstalmentPeriod *Number `xml:"InstlmtPrd,omitempty"`

	// Total number of instalment payments.
	TotalNumberOfPayments *Number `xml:"TtlNbOfPmts,omitempty"`

	// Date of the first payment.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Cumulative amount of all the instalments.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Amount of the first payment.
	FirstAmount *ImpliedCurrencyAndAmount `xml:"FrstAmt,omitempty"`

	// Charges related to the transaction.
	Charges *ImpliedCurrencyAndAmount `xml:"Chrgs,omitempty"`
}

func (r *RecurringTransaction2) AddInstalmentPlan(value string) {
	r.InstalmentPlan = append(r.InstalmentPlan, (*InstalmentPlan1Code)(&value))
}

func (r *RecurringTransaction2) SetPlanIdentification(value string) {
	r.PlanIdentification = (*Max35Text)(&value)
}

func (r *RecurringTransaction2) SetSequenceNumber(value string) {
	r.SequenceNumber = (*Number)(&value)
}

func (r *RecurringTransaction2) SetPeriodUnit(value string) {
	r.PeriodUnit = (*Frequency3Code)(&value)
}

func (r *RecurringTransaction2) SetInstalmentPeriod(value string) {
	r.InstalmentPeriod = (*Number)(&value)
}

func (r *RecurringTransaction2) SetTotalNumberOfPayments(value string) {
	r.TotalNumberOfPayments = (*Number)(&value)
}

func (r *RecurringTransaction2) SetFirstPaymentDate(value string) {
	r.FirstPaymentDate = (*ISODate)(&value)
}

func (r *RecurringTransaction2) SetTotalAmount(value, currency string) {
	r.TotalAmount = NewCurrencyAndAmount(value, currency)
}

func (r *RecurringTransaction2) SetFirstAmount(value, currency string) {
	r.FirstAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (r *RecurringTransaction2) SetCharges(value, currency string) {
	r.Charges = NewImpliedCurrencyAndAmount(value, currency)
}

// Financial loan (instalment) or a recurring transaction.
type RecurringTransaction3 struct {

	// Date of first transfer.
	StartDate *ISODate `xml:"StartDt"`

	// Number of transfers to perform.
	NumberOfOccurrences *Number `xml:"NbOfOcrncs,omitempty"`

	// Date of last transfer.
	EndDate *ISODate `xml:"EndDt"`

	// Period of the recurring transfer.
	PeriodUnit *Frequency3Code `xml:"PrdUnit,omitempty"`

	// Day of the period when the transfer will be performed.
	IntervalDay *Number `xml:"IntrvlDay,omitempty"`
}

func (r *RecurringTransaction3) SetStartDate(value string) {
	r.StartDate = (*ISODate)(&value)
}

func (r *RecurringTransaction3) SetNumberOfOccurrences(value string) {
	r.NumberOfOccurrences = (*Number)(&value)
}

func (r *RecurringTransaction3) SetEndDate(value string) {
	r.EndDate = (*ISODate)(&value)
}

func (r *RecurringTransaction3) SetPeriodUnit(value string) {
	r.PeriodUnit = (*Frequency3Code)(&value)
}

func (r *RecurringTransaction3) SetIntervalDay(value string) {
	r.IntervalDay = (*Number)(&value)
}
