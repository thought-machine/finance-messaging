package iso20022

// Amount of money to be transferred between the debtor and creditor, before deduction of charges, expressed in the currency of the debtor's account, and to be transferred into a different currency.
type EquivalentAmount struct {

	// Amount of money to be transferred between debtor and creditor, before deduction of charges, expressed in the currency of the debtor's account, and to be transferred in a different currency.
	//
	// Usage : Currency of the amount is expressed in the currency of the debtor's account, but the amount to be transferred is in another currency. The first agent will convert the amount and currency to the to be transferred amount and currency, eg, 'pay equivalent of 100000 EUR in JPY'(and account is in EUR).
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Specifies the currency of the to be transferred amount, which is different from the currency of the debtor's account.
	CurrencyOfTransfer *CurrencyCode `xml:"CcyOfTrf"`
}

func (e *EquivalentAmount) SetAmount(value, currency string) {
	e.Amount = NewCurrencyAndAmount(value, currency)
}

func (e *EquivalentAmount) SetCurrencyOfTransfer(value string) {
	e.CurrencyOfTransfer = (*CurrencyCode)(&value)
}

// Amount of money to be moved between the debtor and creditor, expressed in the currency of the debtor's account, and the currency in which the amount is to be moved.
type EquivalentAmount2 struct {

	// Amount of money to be moved between debtor and creditor, before deduction of charges, expressed in the currency of the debtor's account, and to be moved in a different currency.
	// Usage: The first agent will convert the equivalent amount into the amount to be moved.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Specifies the currency of the to be transferred amount, which is different from the currency of the debtor's account.
	CurrencyOfTransfer *ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf"`
}

func (e *EquivalentAmount2) SetAmount(value, currency string) {
	e.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EquivalentAmount2) SetCurrencyOfTransfer(value string) {
	e.CurrencyOfTransfer = (*ActiveOrHistoricCurrencyCode)(&value)
}
