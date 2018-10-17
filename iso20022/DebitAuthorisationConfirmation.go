package iso20022

// Indicates if the debit authorisation is granted or not.
type DebitAuthorisationConfirmation struct {

	// Code expressing the decision taken by the account owner relative to the request for debit authorization.
	DebitAuthorisation *YesNoIndicator `xml:"DbtAuthstn"`

	// Amount authorised for debit. The party providing the debit authority may want to authorise the amount less charges and they may only be prepared to approve the debit for value today rather than the original value date.
	AmountToDebit *CurrencyAndAmount `xml:"AmtToDbt,omitempty"`

	// Value date for debiting the amount.
	ValueDateToDebit *ISODate `xml:"ValDtToDbt,omitempty"`

	// Justification of the (partial) debit authorisation.
	Reason *Max140Text `xml:"Rsn,omitempty"`
}

func (d *DebitAuthorisationConfirmation) SetDebitAuthorisation(value string) {
	d.DebitAuthorisation = (*YesNoIndicator)(&value)
}

func (d *DebitAuthorisationConfirmation) SetAmountToDebit(value, currency string) {
	d.AmountToDebit = NewCurrencyAndAmount(value, currency)
}

func (d *DebitAuthorisationConfirmation) SetValueDateToDebit(value string) {
	d.ValueDateToDebit = (*ISODate)(&value)
}

func (d *DebitAuthorisationConfirmation) SetReason(value string) {
	d.Reason = (*Max140Text)(&value)
}

// Specifies whether or not the debit authorisation is granted.
type DebitAuthorisationConfirmation2 struct {

	// Code expressing the decision taken by the account owner relative to the request for debit authorization.
	DebitAuthorisation *YesNoIndicator `xml:"DbtAuthstn"`

	// Amount of money authorised for debit.
	// Usage: The party approving the debit may want to authorise the amount less charges and may only be prepared to approve the debit for value today rather than the original value date.
	AmountToDebit *ActiveCurrencyAndAmount `xml:"AmtToDbt,omitempty"`

	// Value date for debiting the amount.
	ValueDateToDebit *ISODate `xml:"ValDtToDbt,omitempty"`

	// Specifies the reason for the debit authorisation request.
	Reason *Max140Text `xml:"Rsn,omitempty"`
}

func (d *DebitAuthorisationConfirmation2) SetDebitAuthorisation(value string) {
	d.DebitAuthorisation = (*YesNoIndicator)(&value)
}

func (d *DebitAuthorisationConfirmation2) SetAmountToDebit(value, currency string) {
	d.AmountToDebit = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DebitAuthorisationConfirmation2) SetValueDateToDebit(value string) {
	d.ValueDateToDebit = (*ISODate)(&value)
}

func (d *DebitAuthorisationConfirmation2) SetReason(value string) {
	d.Reason = (*Max140Text)(&value)
}
