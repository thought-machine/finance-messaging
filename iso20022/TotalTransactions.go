package iso20022

// Set of element providing summary information on entries.
type TotalTransactions1 struct {

	// Indicates the total number and sum of debit and credit entries.
	TotalEntries *NumberAndSumOfTransactions2 `xml:"TtlNtries,omitempty"`

	// Indicates the total number and sum of credit entries.
	TotalCreditEntries *NumberAndSumOfTransactions1 `xml:"TtlCdtNtries,omitempty"`

	// Indicates the total number and sum of debit entries.
	TotalDebitEntries *NumberAndSumOfTransactions1 `xml:"TtlDbtNtries,omitempty"`

	// Indicates the total number and sum of entries per bank transaction code.
	TotalEntriesPerBankTransactionCode []*NumberAndSumOfTransactionsPerBankTransactionCode1 `xml:"TtlNtriesPerBkTxCd,omitempty"`
}

func (t *TotalTransactions1) AddTotalEntries() *NumberAndSumOfTransactions2 {
	t.TotalEntries = new(NumberAndSumOfTransactions2)
	return t.TotalEntries
}

func (t *TotalTransactions1) AddTotalCreditEntries() *NumberAndSumOfTransactions1 {
	t.TotalCreditEntries = new(NumberAndSumOfTransactions1)
	return t.TotalCreditEntries
}

func (t *TotalTransactions1) AddTotalDebitEntries() *NumberAndSumOfTransactions1 {
	t.TotalDebitEntries = new(NumberAndSumOfTransactions1)
	return t.TotalDebitEntries
}

func (t *TotalTransactions1) AddTotalEntriesPerBankTransactionCode() *NumberAndSumOfTransactionsPerBankTransactionCode1 {
	newValue := new(NumberAndSumOfTransactionsPerBankTransactionCode1)
	t.TotalEntriesPerBankTransactionCode = append(t.TotalEntriesPerBankTransactionCode, newValue)
	return newValue
}

// Set of elements used to provide summary information on entries.
type TotalTransactions2 struct {

	// Specifies the total number and sum of debit and credit entries.
	TotalEntries *NumberAndSumOfTransactions2 `xml:"TtlNtries,omitempty"`

	// Specifies the total number and sum of credit entries.
	TotalCreditEntries *NumberAndSumOfTransactions1 `xml:"TtlCdtNtries,omitempty"`

	// Specifies the total number and sum of debit entries.
	TotalDebitEntries *NumberAndSumOfTransactions1 `xml:"TtlDbtNtries,omitempty"`

	// Specifies the total number and sum of entries per bank transaction code.
	TotalEntriesPerBankTransactionCode []*TotalsPerBankTransactionCode2 `xml:"TtlNtriesPerBkTxCd,omitempty"`
}

func (t *TotalTransactions2) AddTotalEntries() *NumberAndSumOfTransactions2 {
	t.TotalEntries = new(NumberAndSumOfTransactions2)
	return t.TotalEntries
}

func (t *TotalTransactions2) AddTotalCreditEntries() *NumberAndSumOfTransactions1 {
	t.TotalCreditEntries = new(NumberAndSumOfTransactions1)
	return t.TotalCreditEntries
}

func (t *TotalTransactions2) AddTotalDebitEntries() *NumberAndSumOfTransactions1 {
	t.TotalDebitEntries = new(NumberAndSumOfTransactions1)
	return t.TotalDebitEntries
}

func (t *TotalTransactions2) AddTotalEntriesPerBankTransactionCode() *TotalsPerBankTransactionCode2 {
	newValue := new(TotalsPerBankTransactionCode2)
	t.TotalEntriesPerBankTransactionCode = append(t.TotalEntriesPerBankTransactionCode, newValue)
	return newValue
}

// Set of elements used to provide summary information on entries.
type TotalTransactions3 struct {

	// Specifies the total number and sum of debit and credit entries.
	TotalEntries *NumberAndSumOfTransactions3 `xml:"TtlNtries,omitempty"`

	// Specifies the total number and sum of credit entries.
	TotalCreditEntries *NumberAndSumOfTransactions1 `xml:"TtlCdtNtries,omitempty"`

	// Specifies the total number and sum of debit entries.
	TotalDebitEntries *NumberAndSumOfTransactions1 `xml:"TtlDbtNtries,omitempty"`

	// Specifies the total number and sum of entries per bank transaction code.
	TotalEntriesPerBankTransactionCode []*TotalsPerBankTransactionCode2 `xml:"TtlNtriesPerBkTxCd,omitempty"`
}

func (t *TotalTransactions3) AddTotalEntries() *NumberAndSumOfTransactions3 {
	t.TotalEntries = new(NumberAndSumOfTransactions3)
	return t.TotalEntries
}

func (t *TotalTransactions3) AddTotalCreditEntries() *NumberAndSumOfTransactions1 {
	t.TotalCreditEntries = new(NumberAndSumOfTransactions1)
	return t.TotalCreditEntries
}

func (t *TotalTransactions3) AddTotalDebitEntries() *NumberAndSumOfTransactions1 {
	t.TotalDebitEntries = new(NumberAndSumOfTransactions1)
	return t.TotalDebitEntries
}

func (t *TotalTransactions3) AddTotalEntriesPerBankTransactionCode() *TotalsPerBankTransactionCode2 {
	newValue := new(TotalsPerBankTransactionCode2)
	t.TotalEntriesPerBankTransactionCode = append(t.TotalEntriesPerBankTransactionCode, newValue)
	return newValue
}

// Set of elements used to provide summary information on entries.
type TotalTransactions4 struct {

	// Specifies the total number and sum of debit and credit entries.
	TotalEntries *NumberAndSumOfTransactions4 `xml:"TtlNtries,omitempty"`

	// Specifies the total number and sum of credit entries.
	TotalCreditEntries *NumberAndSumOfTransactions1 `xml:"TtlCdtNtries,omitempty"`

	// Specifies the total number and sum of debit entries.
	TotalDebitEntries *NumberAndSumOfTransactions1 `xml:"TtlDbtNtries,omitempty"`

	// Specifies the total number and sum of entries per bank transaction code.
	TotalEntriesPerBankTransactionCode []*TotalsPerBankTransactionCode3 `xml:"TtlNtriesPerBkTxCd,omitempty"`
}

func (t *TotalTransactions4) AddTotalEntries() *NumberAndSumOfTransactions4 {
	t.TotalEntries = new(NumberAndSumOfTransactions4)
	return t.TotalEntries
}

func (t *TotalTransactions4) AddTotalCreditEntries() *NumberAndSumOfTransactions1 {
	t.TotalCreditEntries = new(NumberAndSumOfTransactions1)
	return t.TotalCreditEntries
}

func (t *TotalTransactions4) AddTotalDebitEntries() *NumberAndSumOfTransactions1 {
	t.TotalDebitEntries = new(NumberAndSumOfTransactions1)
	return t.TotalDebitEntries
}

func (t *TotalTransactions4) AddTotalEntriesPerBankTransactionCode() *TotalsPerBankTransactionCode3 {
	newValue := new(TotalsPerBankTransactionCode3)
	t.TotalEntriesPerBankTransactionCode = append(t.TotalEntriesPerBankTransactionCode, newValue)
	return newValue
}

// Set of elements used to provide summary information on entries.
type TotalTransactions5 struct {

	// Specifies the total number and sum of debit and credit entries.
	TotalEntries *NumberAndSumOfTransactions4 `xml:"TtlNtries,omitempty"`

	// Specifies the total number and sum of credit entries.
	TotalCreditEntries *NumberAndSumOfTransactions1 `xml:"TtlCdtNtries,omitempty"`

	// Specifies the total number and sum of debit entries.
	TotalDebitEntries *NumberAndSumOfTransactions1 `xml:"TtlDbtNtries,omitempty"`

	// Specifies the total number and sum of entries per bank transaction code.
	TotalEntriesPerBankTransactionCode []*TotalsPerBankTransactionCode4 `xml:"TtlNtriesPerBkTxCd,omitempty"`
}

func (t *TotalTransactions5) AddTotalEntries() *NumberAndSumOfTransactions4 {
	t.TotalEntries = new(NumberAndSumOfTransactions4)
	return t.TotalEntries
}

func (t *TotalTransactions5) AddTotalCreditEntries() *NumberAndSumOfTransactions1 {
	t.TotalCreditEntries = new(NumberAndSumOfTransactions1)
	return t.TotalCreditEntries
}

func (t *TotalTransactions5) AddTotalDebitEntries() *NumberAndSumOfTransactions1 {
	t.TotalDebitEntries = new(NumberAndSumOfTransactions1)
	return t.TotalDebitEntries
}

func (t *TotalTransactions5) AddTotalEntriesPerBankTransactionCode() *TotalsPerBankTransactionCode4 {
	newValue := new(TotalsPerBankTransactionCode4)
	t.TotalEntriesPerBankTransactionCode = append(t.TotalEntriesPerBankTransactionCode, newValue)
	return newValue
}
