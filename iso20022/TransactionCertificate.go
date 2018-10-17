package iso20022

// Certificate against which all currency control transactions are registered.
type TransactionCertificate1 struct {

	// Unique and unambiguous identification of the transaction.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Reference of the transaction certificate.
	Certificate *DocumentIdentification28 `xml:"Cert"`

	// Cash account, linked to the registered contract, on which the cash entry is made.
	Account *CashAccount24 `xml:"Acct,omitempty"`

	// Country in which the bank account is located, when the bank which services the account is located in another country.
	BankAccountDomiciliationCountry *CountryCode `xml:"BkAcctDmcltnCtry,omitempty"`

	// Amendment indicator details.
	Amendment *DocumentAmendment1 `xml:"Amdmnt,omitempty"`

	// Record of the transaction certificate.
	CertificateRecord []*TransactionCertificateRecord1 `xml:"CertRcrd"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionCertificate1) SetTransactionIdentification(value string) {
	t.TransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionCertificate1) AddCertificate() *DocumentIdentification28 {
	t.Certificate = new(DocumentIdentification28)
	return t.Certificate
}

func (t *TransactionCertificate1) AddAccount() *CashAccount24 {
	t.Account = new(CashAccount24)
	return t.Account
}

func (t *TransactionCertificate1) SetBankAccountDomiciliationCountry(value string) {
	t.BankAccountDomiciliationCountry = (*CountryCode)(&value)
}

func (t *TransactionCertificate1) AddAmendment() *DocumentAmendment1 {
	t.Amendment = new(DocumentAmendment1)
	return t.Amendment
}

func (t *TransactionCertificate1) AddCertificateRecord() *TransactionCertificateRecord1 {
	newValue := new(TransactionCertificateRecord1)
	t.CertificateRecord = append(t.CertificateRecord, newValue)
	return newValue
}

func (t *TransactionCertificate1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Certificate in which all currency control transactions are registered.
type TransactionCertificate2 struct {

	// Reference of the transaction, that is the underlying payment instruction or statement entry.
	ReferredDocument *CertificateReference1 `xml:"RfrdDoc"`

	// Date of the underlying transaction.
	TransactionDate *ISODate `xml:"TxDt"`

	// Type of the transaction.
	TransactionType *Exact1NumericText `xml:"TxTp"`

	// User community specific instrument.
	//
	// Usage: This element is used to specify a local transaction type to further qualify the transaction type.
	LocalInstrument *Exact5NumericText `xml:"LclInstrm"`

	// Amount as provided in the transaction to be recorded under the contract.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (t *TransactionCertificate2) AddReferredDocument() *CertificateReference1 {
	t.ReferredDocument = new(CertificateReference1)
	return t.ReferredDocument
}

func (t *TransactionCertificate2) SetTransactionDate(value string) {
	t.TransactionDate = (*ISODate)(&value)
}

func (t *TransactionCertificate2) SetTransactionType(value string) {
	t.TransactionType = (*Exact1NumericText)(&value)
}

func (t *TransactionCertificate2) SetLocalInstrument(value string) {
	t.LocalInstrument = (*Exact5NumericText)(&value)
}

func (t *TransactionCertificate2) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAndAmount(value, currency)
}
