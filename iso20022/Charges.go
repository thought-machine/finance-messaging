package iso20022

// Specifies the type of charge and the amont.
type Charges1 struct {

	// Type of charges.
	Type *ChargeType2FormatChoice `xml:"Tp,omitempty"`

	// Amount of charges.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (c *Charges1) AddType() *ChargeType2FormatChoice {
	c.Type = new(ChargeType2FormatChoice)
	return c.Type
}

func (c *Charges1) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAndAmount(value, currency)
}

// Set of elements used to provide information on the charges related to the payment transaction.
type Charges2 struct {

	// Transaction charges to be paid by the charge bearer.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Agent that takes the transaction charges or to which the transaction charges are due.
	Agent *BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
}

func (c *Charges2) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charges2) AddAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.Agent = new(BranchAndFinancialInstitutionIdentification5)
	return c.Agent
}

// Provides further details on the charges related to the payment transaction.
type Charges3 struct {

	// Total of all charges and taxes applied to the entry.
	TotalChargesAndTaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty"`

	// Provides details of the individual charges record.
	Record []*ChargesRecord1 `xml:"Rcrd,omitempty"`
}

func (c *Charges3) SetTotalChargesAndTaxAmount(value, currency string) {
	c.TotalChargesAndTaxAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charges3) AddRecord() *ChargesRecord1 {
	newValue := new(ChargesRecord1)
	c.Record = append(c.Record, newValue)
	return newValue
}

// Provides further details on the charges related to the payment transaction.
type Charges4 struct {

	// Total of all charges and taxes applied to the entry.
	TotalChargesAndTaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty"`

	// Provides details of the individual charges record.
	Record []*ChargesRecord2 `xml:"Rcrd,omitempty"`
}

func (c *Charges4) SetTotalChargesAndTaxAmount(value, currency string) {
	c.TotalChargesAndTaxAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charges4) AddRecord() *ChargesRecord2 {
	newValue := new(ChargesRecord2)
	c.Record = append(c.Record, newValue)
	return newValue
}

// Charges related to a payment obligation contracted between two financial institutions related to the financing of a commercial transaction.
type Charges5 struct {

	// Bank which will pay the charges.
	ChargesPayer *BankRole1Code `xml:"ChrgsPyer"`

	// Bank which will receive the charges.
	ChargesPayee *BankRole1Code `xml:"ChrgsPyee"`

	// Amount of the charges taken by the payer.
	Amount *CurrencyAndAmount `xml:"Amt,omitempty"`

	// Amount of the charges expressed as a percentage of the amount paid by the obligor bank.
	Percentage *PercentageRate `xml:"Pctg,omitempty"`

	// Type of charges. For example: transaction charges, financing charges, deferred payment, interests.
	Type *Max35Text `xml:"Tp,omitempty"`
}

func (c *Charges5) SetChargesPayer(value string) {
	c.ChargesPayer = (*BankRole1Code)(&value)
}

func (c *Charges5) SetChargesPayee(value string) {
	c.ChargesPayee = (*BankRole1Code)(&value)
}

func (c *Charges5) SetAmount(value, currency string) {
	c.Amount = NewCurrencyAndAmount(value, currency)
}

func (c *Charges5) SetPercentage(value string) {
	c.Percentage = (*PercentageRate)(&value)
}

func (c *Charges5) SetType(value string) {
	c.Type = (*Max35Text)(&value)
}
