package iso20022

// Provides information about pending balance and pending transactions.
type PendingBalance1 struct {

	// Signed quantity of balance.
	Balance *SignedQuantityFormat2 `xml:"Bal"`

	// Overall process covering the trade and settlement transactions of financial instruments.
	PendingTransactions []*SettlementTypeAndIdentification2 `xml:"PdgTxs,omitempty"`
}

func (p *PendingBalance1) AddBalance() *SignedQuantityFormat2 {
	p.Balance = new(SignedQuantityFormat2)
	return p.Balance
}

func (p *PendingBalance1) AddPendingTransactions() *SettlementTypeAndIdentification2 {
	newValue := new(SettlementTypeAndIdentification2)
	p.PendingTransactions = append(p.PendingTransactions, newValue)
	return newValue
}

// Provides information about pending balance and pending transactions.
type PendingBalance3 struct {

	// Signed quantity of balance.
	Balance *SignedQuantityFormat6 `xml:"Bal"`

	// Overall process covering the trade and settlement transactions of financial instruments.
	PendingTransactions []*SettlementTypeAndIdentification20 `xml:"PdgTxs,omitempty"`
}

func (p *PendingBalance3) AddBalance() *SignedQuantityFormat6 {
	p.Balance = new(SignedQuantityFormat6)
	return p.Balance
}

func (p *PendingBalance3) AddPendingTransactions() *SettlementTypeAndIdentification20 {
	newValue := new(SettlementTypeAndIdentification20)
	p.PendingTransactions = append(p.PendingTransactions, newValue)
	return newValue
}

// Provides information about pending balance and pending transactions.
type PendingBalance4 struct {

	// Signed quantity of balance.
	Balance *SignedQuantityFormat9 `xml:"Bal"`

	// Overall process covering the trade and settlement transactions of financial instruments.
	PendingTransactions []*SettlementTypeAndIdentification21 `xml:"PdgTxs,omitempty"`
}

func (p *PendingBalance4) AddBalance() *SignedQuantityFormat9 {
	p.Balance = new(SignedQuantityFormat9)
	return p.Balance
}

func (p *PendingBalance4) AddPendingTransactions() *SettlementTypeAndIdentification21 {
	newValue := new(SettlementTypeAndIdentification21)
	p.PendingTransactions = append(p.PendingTransactions, newValue)
	return newValue
}
