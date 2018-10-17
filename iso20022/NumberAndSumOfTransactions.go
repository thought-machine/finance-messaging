package iso20022

// Set of elements providing the total sum of entries.
type NumberAndSumOfTransactions1 struct {

	// Number of individual entries included in the report.
	NumberOfEntries *Max15NumericText `xml:"NbOfNtries,omitempty"`

	// Total of all individual entries included in the report.
	Sum *DecimalNumber `xml:"Sum,omitempty"`
}

func (n *NumberAndSumOfTransactions1) SetNumberOfEntries(value string) {
	n.NumberOfEntries = (*Max15NumericText)(&value)
}

func (n *NumberAndSumOfTransactions1) SetSum(value string) {
	n.Sum = (*DecimalNumber)(&value)
}

// Set of elements providing the total sum of entries.
type NumberAndSumOfTransactions2 struct {

	// Number of individual entries included in the report.
	NumberOfEntries *Max15NumericText `xml:"NbOfNtries,omitempty"`

	// Total of all individual entries included in the report.
	Sum *DecimalNumber `xml:"Sum,omitempty"`

	// Resulting amount of the netted amounts for all debit and credit entries.
	TotalNetEntryAmount *DecimalNumber `xml:"TtlNetNtryAmt,omitempty"`

	// Indicates whether the total net entry amount is a credit or a debit amount.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`
}

func (n *NumberAndSumOfTransactions2) SetNumberOfEntries(value string) {
	n.NumberOfEntries = (*Max15NumericText)(&value)
}

func (n *NumberAndSumOfTransactions2) SetSum(value string) {
	n.Sum = (*DecimalNumber)(&value)
}

func (n *NumberAndSumOfTransactions2) SetTotalNetEntryAmount(value string) {
	n.TotalNetEntryAmount = (*DecimalNumber)(&value)
}

func (n *NumberAndSumOfTransactions2) SetCreditDebitIndicator(value string) {
	n.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

// Set of elements providing the total sum of entries.
type NumberAndSumOfTransactions3 struct {

	// Number of individual entries included in the report.
	NumberOfEntries *Max15NumericText `xml:"NbOfNtries,omitempty"`

	// Total of all individual entries included in the report.
	Sum *DecimalNumber `xml:"Sum,omitempty"`

	// Resulting amount of the netted amounts for all debit and credit entries.
	TotalNetEntryAmount *DecimalNumber `xml:"TtlNetNtryAmt,omitempty"`

	// Indicates whether the total net entry amount is a credit or a debit amount.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`
}

func (n *NumberAndSumOfTransactions3) SetNumberOfEntries(value string) {
	n.NumberOfEntries = (*Max15NumericText)(&value)
}

func (n *NumberAndSumOfTransactions3) SetSum(value string) {
	n.Sum = (*DecimalNumber)(&value)
}

func (n *NumberAndSumOfTransactions3) SetTotalNetEntryAmount(value string) {
	n.TotalNetEntryAmount = (*DecimalNumber)(&value)
}

func (n *NumberAndSumOfTransactions3) SetCreditDebitIndicator(value string) {
	n.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

// Set of elements providing the total sum of entries.
type NumberAndSumOfTransactions4 struct {

	// Number of individual entries included in the report.
	NumberOfEntries *Max15NumericText `xml:"NbOfNtries,omitempty"`

	// Total of all individual entries included in the report.
	Sum *DecimalNumber `xml:"Sum,omitempty"`

	// Resulting debit or credit amount of the netted amounts for all debit and credit entries.
	TotalNetEntry *AmountAndDirection35 `xml:"TtlNetNtry,omitempty"`
}

func (n *NumberAndSumOfTransactions4) SetNumberOfEntries(value string) {
	n.NumberOfEntries = (*Max15NumericText)(&value)
}

func (n *NumberAndSumOfTransactions4) SetSum(value string) {
	n.Sum = (*DecimalNumber)(&value)
}

func (n *NumberAndSumOfTransactions4) AddTotalNetEntry() *AmountAndDirection35 {
	n.TotalNetEntry = new(AmountAndDirection35)
	return n.TotalNetEntry
}
