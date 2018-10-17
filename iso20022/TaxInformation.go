package iso20022

// Details about tax paid, or to be paid, to the government in accordance with the law, including pre-defined parameters such as thresholds and type of account.
type TaxInformation2 struct {

	// Tax Identification Number of the creditor.
	CreditorTaxIdentification *Max35Text `xml:"CdtrTaxId,omitempty"`

	// Type of tax payer (creditor).
	CreditorTaxType *Max35Text `xml:"CdtrTaxTp,omitempty"`

	// Tax Identification Number of the debtor.
	DebtorTaxIdentification *Max35Text `xml:"DbtrTaxId,omitempty"`

	// Tax reference information that is specific to a taxing agency.
	TaxReferenceNumber *Max140Text `xml:"TaxRefNb,omitempty"`

	// Total amount of money on which the tax is based.
	TotalTaxableBaseAmount *CurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`

	// Amount of money resulting from the calculation of the tax.
	TotalTaxAmount *CurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`

	// Date by which tax is due.
	TaxDate *ISODate `xml:"TaxDt,omitempty"`

	// Set of characteristics defining the type of tax.
	TaxTypeInformation []*TaxDetails `xml:"TaxTpInf,omitempty"`
}

func (t *TaxInformation2) SetCreditorTaxIdentification(value string) {
	t.CreditorTaxIdentification = (*Max35Text)(&value)
}

func (t *TaxInformation2) SetCreditorTaxType(value string) {
	t.CreditorTaxType = (*Max35Text)(&value)
}

func (t *TaxInformation2) SetDebtorTaxIdentification(value string) {
	t.DebtorTaxIdentification = (*Max35Text)(&value)
}

func (t *TaxInformation2) SetTaxReferenceNumber(value string) {
	t.TaxReferenceNumber = (*Max140Text)(&value)
}

func (t *TaxInformation2) SetTotalTaxableBaseAmount(value, currency string) {
	t.TotalTaxableBaseAmount = NewCurrencyAndAmount(value, currency)
}

func (t *TaxInformation2) SetTotalTaxAmount(value, currency string) {
	t.TotalTaxAmount = NewCurrencyAndAmount(value, currency)
}

func (t *TaxInformation2) SetTaxDate(value string) {
	t.TaxDate = (*ISODate)(&value)
}

func (t *TaxInformation2) AddTaxTypeInformation() *TaxDetails {
	newValue := new(TaxDetails)
	t.TaxTypeInformation = append(t.TaxTypeInformation, newValue)
	return newValue
}

// Details about tax paid, or to be paid, to the government in accordance with the law, including pre-defined parameters such as thresholds and type of account.
type TaxInformation3 struct {

	// Party on the credit side of the transaction to which the tax applies.
	Creditor *TaxParty1 `xml:"Cdtr,omitempty"`

	// Set of elements used to identify the party on the debit side of the transaction to which the tax applies.
	Debtor *TaxParty2 `xml:"Dbtr,omitempty"`

	// Territorial part of a country to which the tax payment is related.
	AdministrationZone *Max35Text `xml:"AdmstnZn,omitempty"`

	// Tax reference information that is specific to a taxing agency.
	ReferenceNumber *Max140Text `xml:"RefNb,omitempty"`

	// Method used to indicate the underlying business or how the tax is paid.
	Method *Max35Text `xml:"Mtd,omitempty"`

	// Total amount of money on which the tax is based.
	TotalTaxableBaseAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`

	// Total amount of money as result of the calculation of the tax.
	TotalTaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`

	// Date by which tax is due.
	Date *ISODate `xml:"Dt,omitempty"`

	// Sequential number of the tax report.
	SequenceNumber *Number `xml:"SeqNb,omitempty"`

	// Record of tax details.
	Record []*TaxRecord1 `xml:"Rcrd,omitempty"`
}

func (t *TaxInformation3) AddCreditor() *TaxParty1 {
	t.Creditor = new(TaxParty1)
	return t.Creditor
}

func (t *TaxInformation3) AddDebtor() *TaxParty2 {
	t.Debtor = new(TaxParty2)
	return t.Debtor
}

func (t *TaxInformation3) SetAdministrationZone(value string) {
	t.AdministrationZone = (*Max35Text)(&value)
}

func (t *TaxInformation3) SetReferenceNumber(value string) {
	t.ReferenceNumber = (*Max140Text)(&value)
}

func (t *TaxInformation3) SetMethod(value string) {
	t.Method = (*Max35Text)(&value)
}

func (t *TaxInformation3) SetTotalTaxableBaseAmount(value, currency string) {
	t.TotalTaxableBaseAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation3) SetTotalTaxAmount(value, currency string) {
	t.TotalTaxAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation3) SetDate(value string) {
	t.Date = (*ISODate)(&value)
}

func (t *TaxInformation3) SetSequenceNumber(value string) {
	t.SequenceNumber = (*Number)(&value)
}

func (t *TaxInformation3) AddRecord() *TaxRecord1 {
	newValue := new(TaxRecord1)
	t.Record = append(t.Record, newValue)
	return newValue
}

// Details about tax paid, or to be paid, to the government in accordance with the law, including pre-defined parameters such as thresholds and type of account.
type TaxInformation4 struct {

	// Party on the credit side of the transaction to which the tax applies.
	Creditor *TaxParty1 `xml:"Cdtr,omitempty"`

	// Identifies the party on the debit side of the transaction to which the tax applies.
	Debtor *TaxParty2 `xml:"Dbtr,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor, in this case, to the taxing authority.
	UltimateDebtor *TaxParty2 `xml:"UltmtDbtr,omitempty"`

	// Territorial part of a country to which the tax payment is related.
	AdministrationZone *Max35Text `xml:"AdmstnZone,omitempty"`

	// Tax reference information that is specific to a taxing agency.
	ReferenceNumber *Max140Text `xml:"RefNb,omitempty"`

	// Method used to indicate the underlying business or how the tax is paid.
	Method *Max35Text `xml:"Mtd,omitempty"`

	// Total amount of money on which the tax is based.
	TotalTaxableBaseAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`

	// Total amount of money as result of the calculation of the tax.
	TotalTaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`

	// Date by which tax is due.
	Date *ISODate `xml:"Dt,omitempty"`

	// Sequential number of the tax report.
	SequenceNumber *Number `xml:"SeqNb,omitempty"`

	// Record of tax details.
	Record []*TaxRecord1 `xml:"Rcrd,omitempty"`
}

func (t *TaxInformation4) AddCreditor() *TaxParty1 {
	t.Creditor = new(TaxParty1)
	return t.Creditor
}

func (t *TaxInformation4) AddDebtor() *TaxParty2 {
	t.Debtor = new(TaxParty2)
	return t.Debtor
}

func (t *TaxInformation4) AddUltimateDebtor() *TaxParty2 {
	t.UltimateDebtor = new(TaxParty2)
	return t.UltimateDebtor
}

func (t *TaxInformation4) SetAdministrationZone(value string) {
	t.AdministrationZone = (*Max35Text)(&value)
}

func (t *TaxInformation4) SetReferenceNumber(value string) {
	t.ReferenceNumber = (*Max140Text)(&value)
}

func (t *TaxInformation4) SetMethod(value string) {
	t.Method = (*Max35Text)(&value)
}

func (t *TaxInformation4) SetTotalTaxableBaseAmount(value, currency string) {
	t.TotalTaxableBaseAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation4) SetTotalTaxAmount(value, currency string) {
	t.TotalTaxAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation4) SetDate(value string) {
	t.Date = (*ISODate)(&value)
}

func (t *TaxInformation4) SetSequenceNumber(value string) {
	t.SequenceNumber = (*Number)(&value)
}

func (t *TaxInformation4) AddRecord() *TaxRecord1 {
	newValue := new(TaxRecord1)
	t.Record = append(t.Record, newValue)
	return newValue
}
