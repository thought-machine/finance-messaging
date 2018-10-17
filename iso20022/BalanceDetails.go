package iso20022

// Balance related details for a portfolio.
type BalanceDetails5 struct {

	// Balance type.
	Type *BalanceType6Choice `xml:"Tp"`

	// Unrealised gain or loss.
	Unrealised *Unrealised1Code `xml:"Urlsd,omitempty"`

	// Balance amount.
	Amount *AmountAndDirection31 `xml:"Amt"`

	// Detailed balance information.
	DetailedBalance []*BalanceDetails6 `xml:"DtldBal,omitempty"`
}

func (b *BalanceDetails5) AddType() *BalanceType6Choice {
	b.Type = new(BalanceType6Choice)
	return b.Type
}

func (b *BalanceDetails5) SetUnrealised(value string) {
	b.Unrealised = (*Unrealised1Code)(&value)
}

func (b *BalanceDetails5) AddAmount() *AmountAndDirection31 {
	b.Amount = new(AmountAndDirection31)
	return b.Amount
}

func (b *BalanceDetails5) AddDetailedBalance() *BalanceDetails6 {
	newValue := new(BalanceDetails6)
	b.DetailedBalance = append(b.DetailedBalance, newValue)
	return newValue
}

// Balance related details for a portfolio.
type BalanceDetails6 struct {

	// Category of the financial asset balance type.
	Category *FinancialAssetTypeCategory1Code `xml:"Ctgy,omitempty"`

	// Balance type.
	Type *BalanceType7Choice `xml:"Tp,omitempty"`

	// Unrealised gain or loss.
	Unrealised *Unrealised1Code `xml:"Urlsd,omitempty"`

	// Balance amount.
	Amount *AmountAndDirection31 `xml:"Amt"`
}

func (b *BalanceDetails6) SetCategory(value string) {
	b.Category = (*FinancialAssetTypeCategory1Code)(&value)
}

func (b *BalanceDetails6) AddType() *BalanceType7Choice {
	b.Type = new(BalanceType7Choice)
	return b.Type
}

func (b *BalanceDetails6) SetUnrealised(value string) {
	b.Unrealised = (*Unrealised1Code)(&value)
}

func (b *BalanceDetails6) AddAmount() *AmountAndDirection31 {
	b.Amount = new(AmountAndDirection31)
	return b.Amount
}
