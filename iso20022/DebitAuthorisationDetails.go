package iso20022

// Provides the reason for requesting a debit authorisation as well as the amount of the requested debit.
type DebitAuthorisationDetails struct {

	// Indicates the reason for cancellation.
	CancellationReason *CancellationReason1Code `xml:"CxlRsn"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	AmountToDebit *CurrencyAndAmount `xml:"AmtToDbt,omitempty"`

	// Value date for debiting the amount.
	ValueDateToDebit *ISODate `xml:"ValDtToDbt,omitempty"`
}

func (d *DebitAuthorisationDetails) SetCancellationReason(value string) {
	d.CancellationReason = (*CancellationReason1Code)(&value)
}

func (d *DebitAuthorisationDetails) SetAmountToDebit(value, currency string) {
	d.AmountToDebit = NewCurrencyAndAmount(value, currency)
}

func (d *DebitAuthorisationDetails) SetValueDateToDebit(value string) {
	d.ValueDateToDebit = (*ISODate)(&value)
}

// Provides the reason for requesting a debit authorisation as well as the amount of the requested debit.
type DebitAuthorisationDetails3 struct {

	// Specifies the reason for the cancellation request.
	CancellationReason *CancellationReason2Choice `xml:"CxlRsn"`

	// Amount of money requested for debit authorisation.
	AmountToDebit *ActiveOrHistoricCurrencyAndAmount `xml:"AmtToDbt,omitempty"`

	// Value date for debiting the amount.
	ValueDateToDebit *ISODate `xml:"ValDtToDbt,omitempty"`

	// Further details on the cancellation request reason.
	AdditionalCancellationReasonInformation []*Max105Text `xml:"AddtlCxlRsnInf,omitempty"`
}

func (d *DebitAuthorisationDetails3) AddCancellationReason() *CancellationReason2Choice {
	d.CancellationReason = new(CancellationReason2Choice)
	return d.CancellationReason
}

func (d *DebitAuthorisationDetails3) SetAmountToDebit(value, currency string) {
	d.AmountToDebit = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DebitAuthorisationDetails3) SetValueDateToDebit(value string) {
	d.ValueDateToDebit = (*ISODate)(&value)
}

func (d *DebitAuthorisationDetails3) AddAdditionalCancellationReasonInformation(value string) {
	d.AdditionalCancellationReasonInformation = append(d.AdditionalCancellationReasonInformation, (*Max105Text)(&value))
}
