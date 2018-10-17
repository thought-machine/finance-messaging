package iso20022

// Set of elements providing information on the original amount and currency information.
type AmountAndCurrencyExchange2 struct {

	// Identifies the amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party and provides currency exchange info in case the instructed amount and/or currency is/are different from the entry amount and/or currency.
	InstructedAmount *AmountAndCurrencyExchangeDetails1 `xml:"InstdAmt,omitempty"`

	// Amount of the underlying transaction.
	TransactionAmount *AmountAndCurrencyExchangeDetails1 `xml:"TxAmt,omitempty"`

	// Identifies the countervalue amount and provides currency exchange information. Either the counter amount quoted in an FX deal, or the result of the currency information applied to an instructed amount, before deduction of charges.
	CounterValueAmount *AmountAndCurrencyExchangeDetails1 `xml:"CntrValAmt,omitempty"`

	// Information on the amount of money, based on terms of corporate action event and balance of underlying securities, entitled to/from the account owner.
	//
	// Amount of money, based on terms of corporate action event and balance of underlying securities, entitled to/from the account owner.
	// In those situations, this amount may alternatively be called entitled amount.
	AnnouncedPostingAmount *AmountAndCurrencyExchangeDetails1 `xml:"AnncdPstngAmt,omitempty"`

	// Provides proprietary amount information.
	ProprietaryAmount []*AmountAndCurrencyExchangeDetails2 `xml:"PrtryAmt,omitempty"`
}

func (a *AmountAndCurrencyExchange2) AddInstructedAmount() *AmountAndCurrencyExchangeDetails1 {
	a.InstructedAmount = new(AmountAndCurrencyExchangeDetails1)
	return a.InstructedAmount
}

func (a *AmountAndCurrencyExchange2) AddTransactionAmount() *AmountAndCurrencyExchangeDetails1 {
	a.TransactionAmount = new(AmountAndCurrencyExchangeDetails1)
	return a.TransactionAmount
}

func (a *AmountAndCurrencyExchange2) AddCounterValueAmount() *AmountAndCurrencyExchangeDetails1 {
	a.CounterValueAmount = new(AmountAndCurrencyExchangeDetails1)
	return a.CounterValueAmount
}

func (a *AmountAndCurrencyExchange2) AddAnnouncedPostingAmount() *AmountAndCurrencyExchangeDetails1 {
	a.AnnouncedPostingAmount = new(AmountAndCurrencyExchangeDetails1)
	return a.AnnouncedPostingAmount
}

func (a *AmountAndCurrencyExchange2) AddProprietaryAmount() *AmountAndCurrencyExchangeDetails2 {
	newValue := new(AmountAndCurrencyExchangeDetails2)
	a.ProprietaryAmount = append(a.ProprietaryAmount, newValue)
	return newValue
}

// Set of elements used to provide information on the original amount.
type AmountAndCurrencyExchange3 struct {

	// Identifies the amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party and provides currency exchange information in case the instructed amount and/or currency is/are different from the entry amount and/or currency.
	InstructedAmount *AmountAndCurrencyExchangeDetails3 `xml:"InstdAmt,omitempty"`

	// Amount of the underlying transaction.
	TransactionAmount *AmountAndCurrencyExchangeDetails3 `xml:"TxAmt,omitempty"`

	// Set of elements used to provide the countervalue amount and  currency exchange information.
	// Usage: This can be either the counter amount quoted in an FX deal, or the result of the currency information applied to an instructed amount, before deduction of charges.
	CounterValueAmount *AmountAndCurrencyExchangeDetails3 `xml:"CntrValAmt,omitempty"`

	// Amount of money, based on terms of corporate action event and balance of underlying securities, entitled to/from the account owner.
	// In some situations, this amount may alternatively be called entitled amount.
	AnnouncedPostingAmount *AmountAndCurrencyExchangeDetails3 `xml:"AnncdPstngAmt,omitempty"`

	// Set of elements used to provide information on the original amount and currency exchange.
	ProprietaryAmount []*AmountAndCurrencyExchangeDetails4 `xml:"PrtryAmt,omitempty"`
}

func (a *AmountAndCurrencyExchange3) AddInstructedAmount() *AmountAndCurrencyExchangeDetails3 {
	a.InstructedAmount = new(AmountAndCurrencyExchangeDetails3)
	return a.InstructedAmount
}

func (a *AmountAndCurrencyExchange3) AddTransactionAmount() *AmountAndCurrencyExchangeDetails3 {
	a.TransactionAmount = new(AmountAndCurrencyExchangeDetails3)
	return a.TransactionAmount
}

func (a *AmountAndCurrencyExchange3) AddCounterValueAmount() *AmountAndCurrencyExchangeDetails3 {
	a.CounterValueAmount = new(AmountAndCurrencyExchangeDetails3)
	return a.CounterValueAmount
}

func (a *AmountAndCurrencyExchange3) AddAnnouncedPostingAmount() *AmountAndCurrencyExchangeDetails3 {
	a.AnnouncedPostingAmount = new(AmountAndCurrencyExchangeDetails3)
	return a.AnnouncedPostingAmount
}

func (a *AmountAndCurrencyExchange3) AddProprietaryAmount() *AmountAndCurrencyExchangeDetails4 {
	newValue := new(AmountAndCurrencyExchangeDetails4)
	a.ProprietaryAmount = append(a.ProprietaryAmount, newValue)
	return newValue
}
