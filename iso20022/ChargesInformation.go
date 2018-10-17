package iso20022

// Information on the charges related to the payment transaction.
type ChargesInformation1 struct {

	// Transaction charges to be paid by the charge bearer.
	ChargesAmount *CurrencyAndAmount `xml:"ChrgsAmt"`

	// Party that takes the transaction charges or to which the transaction charges are due.
	ChargesParty *BranchAndFinancialInstitutionIdentification3 `xml:"ChrgsPty"`
}

func (c *ChargesInformation1) SetChargesAmount(value, currency string) {
	c.ChargesAmount = NewCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation1) AddChargesParty() *BranchAndFinancialInstitutionIdentification3 {
	c.ChargesParty = new(BranchAndFinancialInstitutionIdentification3)
	return c.ChargesParty
}

// Information on the charges related to the payment transaction.
type ChargesInformation3 struct {

	// Total of all charges and taxes applied to the entry.
	TotalChargesAndTaxAmount *CurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty"`

	// Transaction charges to be paid by the charge bearer.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Identifies the type of charge.
	Type *ChargeTypeChoice `xml:"Tp,omitempty"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	Bearer *ChargeBearerType1Code `xml:"Br,omitempty"`

	// Party that takes the transaction charges or to which the transaction charges are due.
	Party *BranchAndFinancialInstitutionIdentification3 `xml:"Pty,omitempty"`

	// Specifies tax details applied to charges.
	Tax *TaxCharges1 `xml:"Tax,omitempty"`
}

func (c *ChargesInformation3) SetTotalChargesAndTaxAmount(value, currency string) {
	c.TotalChargesAndTaxAmount = NewCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation3) SetAmount(value, currency string) {
	c.Amount = NewCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation3) AddType() *ChargeTypeChoice {
	c.Type = new(ChargeTypeChoice)
	return c.Type
}

func (c *ChargesInformation3) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *ChargesInformation3) SetBearer(value string) {
	c.Bearer = (*ChargeBearerType1Code)(&value)
}

func (c *ChargesInformation3) AddParty() *BranchAndFinancialInstitutionIdentification3 {
	c.Party = new(BranchAndFinancialInstitutionIdentification3)
	return c.Party
}

func (c *ChargesInformation3) AddTax() *TaxCharges1 {
	c.Tax = new(TaxCharges1)
	return c.Tax
}

// Set of elements used to provide information on the charges related to the payment transaction.
type ChargesInformation5 struct {

	// Transaction charges to be paid by the charge bearer.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Party that takes the transaction charges or to which the transaction charges are due.
	Party *BranchAndFinancialInstitutionIdentification4 `xml:"Pty"`
}

func (c *ChargesInformation5) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation5) AddParty() *BranchAndFinancialInstitutionIdentification4 {
	c.Party = new(BranchAndFinancialInstitutionIdentification4)
	return c.Party
}

// Set of elements used to provide information on the charges related to the payment transaction.
type ChargesInformation6 struct {

	// Total of all charges and taxes applied to the entry.
	TotalChargesAndTaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty"`

	// Transaction charges to be paid by the charge bearer.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the charges amount is a credit or a debit amount.
	// Usage: A zero amount is considered to be a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Specifies the type of charge.
	Type *ChargeType2Choice `xml:"Tp,omitempty"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	Bearer *ChargeBearerType1Code `xml:"Br,omitempty"`

	// Party that takes the transaction charges or to which the transaction charges are due.
	Party *BranchAndFinancialInstitutionIdentification4 `xml:"Pty,omitempty"`

	// Set of elements used to provide details on the tax applied to charges.
	Tax *TaxCharges2 `xml:"Tax,omitempty"`
}

func (c *ChargesInformation6) SetTotalChargesAndTaxAmount(value, currency string) {
	c.TotalChargesAndTaxAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation6) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation6) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *ChargesInformation6) AddType() *ChargeType2Choice {
	c.Type = new(ChargeType2Choice)
	return c.Type
}

func (c *ChargesInformation6) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *ChargesInformation6) SetBearer(value string) {
	c.Bearer = (*ChargeBearerType1Code)(&value)
}

func (c *ChargesInformation6) AddParty() *BranchAndFinancialInstitutionIdentification4 {
	c.Party = new(BranchAndFinancialInstitutionIdentification4)
	return c.Party
}

func (c *ChargesInformation6) AddTax() *TaxCharges2 {
	c.Tax = new(TaxCharges2)
	return c.Tax
}

// Set of elements used to provide information on the charges related to the payment transaction.
type ChargesInformation7 struct {

	// Transaction charges to be paid by the charge bearer.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Party that takes the transaction charges or to which the transaction charges are due.
	Party *BranchAndFinancialInstitutionIdentification5 `xml:"Pty"`
}

func (c *ChargesInformation7) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *ChargesInformation7) AddParty() *BranchAndFinancialInstitutionIdentification5 {
	c.Party = new(BranchAndFinancialInstitutionIdentification5)
	return c.Party
}
