package iso20022

// Stock lending deadline assigned to a borrower of the stock.
type BorrowerLendingDeadline1 struct {

	// Date/time set as the deadline to respond, with instructions, to an outstanding event, for which the underlying security is out on loan.
	StockLendingDeadline *DateFormat31Choice `xml:"StockLndgDdln"`

	// Party who has borrowed stocks on loan.
	Borrower *PartyIdentification92Choice `xml:"Brrwr"`
}

func (b *BorrowerLendingDeadline1) AddStockLendingDeadline() *DateFormat31Choice {
	b.StockLendingDeadline = new(DateFormat31Choice)
	return b.StockLendingDeadline
}

func (b *BorrowerLendingDeadline1) AddBorrower() *PartyIdentification92Choice {
	b.Borrower = new(PartyIdentification92Choice)
	return b.Borrower
}

// Stock lending deadline assigned to a borrower of the stock.
type BorrowerLendingDeadline2 struct {

	// Date/time set as the deadline to respond, with instructions, to an outstanding event, for which the underlying security is out on loan.
	StockLendingDeadline *DateFormat34Choice `xml:"StockLndgDdln"`

	// Party who has borrowed stocks on loan.
	Borrower *PartyIdentification103Choice `xml:"Brrwr"`
}

func (b *BorrowerLendingDeadline2) AddStockLendingDeadline() *DateFormat34Choice {
	b.StockLendingDeadline = new(DateFormat34Choice)
	return b.StockLendingDeadline
}

func (b *BorrowerLendingDeadline2) AddBorrower() *PartyIdentification103Choice {
	b.Borrower = new(PartyIdentification103Choice)
	return b.Borrower
}
