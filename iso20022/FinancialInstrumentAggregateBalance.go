package iso20022

// Aggregated position of holdings held in a securities account for a specified financial instrument.
type FinancialInstrumentAggregateBalance1 struct {

	// Date of the line of holding in the statement.
	ItemDate *ISODate `xml:"ItmDt"`

	// Balances and sub-balances attributed to the holding.
	Holdings *FinancialInstrumentAggregateBalance1Choice `xml:"Hldgs"`

	// Details on the price value, type and source.
	Price []*Price6 `xml:"Pric,omitempty"`
}

func (f *FinancialInstrumentAggregateBalance1) SetItemDate(value string) {
	f.ItemDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAggregateBalance1) AddHoldings() *FinancialInstrumentAggregateBalance1Choice {
	f.Holdings = new(FinancialInstrumentAggregateBalance1Choice)
	return f.Holdings
}

func (f *FinancialInstrumentAggregateBalance1) AddPrice() *Price6 {
	newValue := new(Price6)
	f.Price = append(f.Price, newValue)
	return newValue
}

// Aggregated position held in a securities account for a specified financial instrument.
type FinancialInstrumentAggregateBalance2 struct {

	// Balance of settled transactions.
	SettledBalance *FinancialInstrumentQuantity1Choice `xml:"SttldBal,omitempty"`

	// Balance of settled transactions and transactions pending settlement.
	TradedBalance *FinancialInstrumentQuantity1Choice `xml:"TraddBal,omitempty"`

	// Breakdown of the balances of holdings into sub-balances.
	BalanceBreakdown []*SubBalanceBreakdown1 `xml:"BalBrkdwn,omitempty"`
}

func (f *FinancialInstrumentAggregateBalance2) AddSettledBalance() *FinancialInstrumentQuantity1Choice {
	f.SettledBalance = new(FinancialInstrumentQuantity1Choice)
	return f.SettledBalance
}

func (f *FinancialInstrumentAggregateBalance2) AddTradedBalance() *FinancialInstrumentQuantity1Choice {
	f.TradedBalance = new(FinancialInstrumentQuantity1Choice)
	return f.TradedBalance
}

func (f *FinancialInstrumentAggregateBalance2) AddBalanceBreakdown() *SubBalanceBreakdown1 {
	newValue := new(SubBalanceBreakdown1)
	f.BalanceBreakdown = append(f.BalanceBreakdown, newValue)
	return newValue
}
