package iso20022

// Set of elements used to provide the total sum of entries per bank transaction code.
type TotalsPerBankTransactionCode2 struct {

	// Number of individual entries for the bank transaction code.
	NumberOfEntries *Max15NumericText `xml:"NbOfNtries,omitempty"`

	// Total of all individual entries included in the report.
	Sum *DecimalNumber `xml:"Sum,omitempty"`

	// Total amount that is the result of the netted amounts for all debit and credit entries per bank transaction code.
	TotalNetEntryAmount *DecimalNumber `xml:"TtlNetNtryAmt,omitempty"`

	// Indicates whether the total net entry amount is a credit or a debit amount.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Indicates whether the bank transaction code is related to booked or forecast items.
	ForecastIndicator *TrueFalseIndicator `xml:"FcstInd,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd"`

	// Set of elements used to indicate when the booked amount of money will become available, that is can be accessed and starts generating interest.
	Availability []*CashBalanceAvailability2 `xml:"Avlbty,omitempty"`
}

func (t *TotalsPerBankTransactionCode2) SetNumberOfEntries(value string) {
	t.NumberOfEntries = (*Max15NumericText)(&value)
}

func (t *TotalsPerBankTransactionCode2) SetSum(value string) {
	t.Sum = (*DecimalNumber)(&value)
}

func (t *TotalsPerBankTransactionCode2) SetTotalNetEntryAmount(value string) {
	t.TotalNetEntryAmount = (*DecimalNumber)(&value)
}

func (t *TotalsPerBankTransactionCode2) SetCreditDebitIndicator(value string) {
	t.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (t *TotalsPerBankTransactionCode2) SetForecastIndicator(value string) {
	t.ForecastIndicator = (*TrueFalseIndicator)(&value)
}

func (t *TotalsPerBankTransactionCode2) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	t.BankTransactionCode = new(BankTransactionCodeStructure4)
	return t.BankTransactionCode
}

func (t *TotalsPerBankTransactionCode2) AddAvailability() *CashBalanceAvailability2 {
	newValue := new(CashBalanceAvailability2)
	t.Availability = append(t.Availability, newValue)
	return newValue
}

// Set of elements used to provide the total sum of entries per bank transaction code.
type TotalsPerBankTransactionCode3 struct {

	// Number of individual entries for the bank transaction code.
	NumberOfEntries *Max15NumericText `xml:"NbOfNtries,omitempty"`

	// Total of all individual entries included in the report.
	Sum *DecimalNumber `xml:"Sum,omitempty"`

	// Total debit or credit amount that is the result of the netted amounts for all debit and credit entries per bank transaction code.
	TotalNetEntry *AmountAndDirection35 `xml:"TtlNetNtry,omitempty"`

	// Indicates whether the bank transaction code is related to booked or forecast items.
	ForecastIndicator *TrueFalseIndicator `xml:"FcstInd,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd"`

	// Set of elements used to indicate when the booked amount of money will become available, that is can be accessed and starts generating interest.
	Availability []*CashBalanceAvailability2 `xml:"Avlbty,omitempty"`
}

func (t *TotalsPerBankTransactionCode3) SetNumberOfEntries(value string) {
	t.NumberOfEntries = (*Max15NumericText)(&value)
}

func (t *TotalsPerBankTransactionCode3) SetSum(value string) {
	t.Sum = (*DecimalNumber)(&value)
}

func (t *TotalsPerBankTransactionCode3) AddTotalNetEntry() *AmountAndDirection35 {
	t.TotalNetEntry = new(AmountAndDirection35)
	return t.TotalNetEntry
}

func (t *TotalsPerBankTransactionCode3) SetForecastIndicator(value string) {
	t.ForecastIndicator = (*TrueFalseIndicator)(&value)
}

func (t *TotalsPerBankTransactionCode3) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	t.BankTransactionCode = new(BankTransactionCodeStructure4)
	return t.BankTransactionCode
}

func (t *TotalsPerBankTransactionCode3) AddAvailability() *CashBalanceAvailability2 {
	newValue := new(CashBalanceAvailability2)
	t.Availability = append(t.Availability, newValue)
	return newValue
}

// Set of elements used to provide the total sum of entries per bank transaction code.
type TotalsPerBankTransactionCode4 struct {

	// Number of individual entries for the bank transaction code.
	NumberOfEntries *Max15NumericText `xml:"NbOfNtries,omitempty"`

	// Total of all individual entries included in the report.
	Sum *DecimalNumber `xml:"Sum,omitempty"`

	// Total debit or credit amount that is the result of the netted amounts for all debit and credit entries per bank transaction code.
	TotalNetEntry *AmountAndDirection35 `xml:"TtlNetNtry,omitempty"`

	// Indicates whether the bank transaction code is related to booked or forecast items.
	ForecastIndicator *TrueFalseIndicator `xml:"FcstInd,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd"`

	// Set of elements used to indicate when the booked amount of money will become available, that is can be accessed and starts generating interest.
	Availability []*CashAvailability1 `xml:"Avlbty,omitempty"`
}

func (t *TotalsPerBankTransactionCode4) SetNumberOfEntries(value string) {
	t.NumberOfEntries = (*Max15NumericText)(&value)
}

func (t *TotalsPerBankTransactionCode4) SetSum(value string) {
	t.Sum = (*DecimalNumber)(&value)
}

func (t *TotalsPerBankTransactionCode4) AddTotalNetEntry() *AmountAndDirection35 {
	t.TotalNetEntry = new(AmountAndDirection35)
	return t.TotalNetEntry
}

func (t *TotalsPerBankTransactionCode4) SetForecastIndicator(value string) {
	t.ForecastIndicator = (*TrueFalseIndicator)(&value)
}

func (t *TotalsPerBankTransactionCode4) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	t.BankTransactionCode = new(BankTransactionCodeStructure4)
	return t.BankTransactionCode
}

func (t *TotalsPerBankTransactionCode4) AddAvailability() *CashAvailability1 {
	newValue := new(CashAvailability1)
	t.Availability = append(t.Availability, newValue)
	return newValue
}
