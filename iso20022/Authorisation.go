package iso20022

// Autorisation of the mandate holder.
type Authorisation1 struct {

	// Minimum amount per transaction allowed by the mandate.
	MinimumAmountPerTransaction *ActiveCurrencyAndAmount `xml:"MinAmtPerTx"`

	// Maximum amount per transaction allowed by the mandate.
	MaximumAmountPerTransaction *ActiveCurrencyAndAmount `xml:"MaxAmtPerTx"`

	// Maximum amount allowed over a specific period of time.
	MaximumAmountByPeriod []*MaximumAmountByPeriod1 `xml:"MaxAmtByPrd,omitempty"`
}

func (a *Authorisation1) SetMinimumAmountPerTransaction(value, currency string) {
	a.MinimumAmountPerTransaction = NewActiveCurrencyAndAmount(value, currency)
}

func (a *Authorisation1) SetMaximumAmountPerTransaction(value, currency string) {
	a.MaximumAmountPerTransaction = NewActiveCurrencyAndAmount(value, currency)
}

func (a *Authorisation1) AddMaximumAmountByPeriod() *MaximumAmountByPeriod1 {
	newValue := new(MaximumAmountByPeriod1)
	a.MaximumAmountByPeriod = append(a.MaximumAmountByPeriod, newValue)
	return newValue
}

// Autorisation of the mandate holder.
type Authorisation2 struct {

	// Maximum amount allowed by the mandate for each transaction.
	MaximumAmountByTransaction *FixedAmountOrUnlimited1Choice `xml:"MaxAmtByTx,omitempty"`

	// Maximum amount allowed over a specific period of time.
	MaximumAmountByPeriod []*MaximumAmountByPeriod1 `xml:"MaxAmtByPrd,omitempty"`

	// Specifies the maximum amount for each  bulk submission.
	MaximumAmountByBulkSubmission *FixedAmountOrUnlimited1Choice `xml:"MaxAmtByBlkSubmissn,omitempty"`
}

func (a *Authorisation2) AddMaximumAmountByTransaction() *FixedAmountOrUnlimited1Choice {
	a.MaximumAmountByTransaction = new(FixedAmountOrUnlimited1Choice)
	return a.MaximumAmountByTransaction
}

func (a *Authorisation2) AddMaximumAmountByPeriod() *MaximumAmountByPeriod1 {
	newValue := new(MaximumAmountByPeriod1)
	a.MaximumAmountByPeriod = append(a.MaximumAmountByPeriod, newValue)
	return newValue
}

func (a *Authorisation2) AddMaximumAmountByBulkSubmission() *FixedAmountOrUnlimited1Choice {
	a.MaximumAmountByBulkSubmission = new(FixedAmountOrUnlimited1Choice)
	return a.MaximumAmountByBulkSubmission
}
