package iso20022

// Specifies the payment terms of the underlying transaction.
type PaymentTerms1 struct {

	// Specifies payment terms not present in a code list.
	OtherPaymentTerms *Max140Text `xml:"OthrPmtTerms"`

	// Specifies the payment period in coded form and a number of days.
	PaymentCode *PaymentPeriod1 `xml:"PmtCd"`

	// Specifies that the payment conditions apply to a percentage of the amount due.
	Percentage *PercentageRate `xml:"Pctg"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (p *PaymentTerms1) SetOtherPaymentTerms(value string) {
	p.OtherPaymentTerms = (*Max140Text)(&value)
}

func (p *PaymentTerms1) AddPaymentCode() *PaymentPeriod1 {
	p.PaymentCode = new(PaymentPeriod1)
	return p.PaymentCode
}

func (p *PaymentTerms1) SetPercentage(value string) {
	p.Percentage = (*PercentageRate)(&value)
}

func (p *PaymentTerms1) SetAmount(value, currency string) {
	p.Amount = NewCurrencyAndAmount(value, currency)
}

// Specifies the payment terms of the underlying transaction.
type PaymentTerms2 struct {

	// Specifies payment terms not present in a code list.
	OtherPaymentTerms *Max140Text `xml:"OthrPmtTerms"`

	// Specifies the payment period in coded form and a number of days.
	PaymentCode *PaymentPeriod2 `xml:"PmtCd"`

	// Specifies that the payment conditions apply to a percentage of the amount due.
	Percentage *PercentageRate `xml:"Pctg"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (p *PaymentTerms2) SetOtherPaymentTerms(value string) {
	p.OtherPaymentTerms = (*Max140Text)(&value)
}

func (p *PaymentTerms2) AddPaymentCode() *PaymentPeriod2 {
	p.PaymentCode = new(PaymentPeriod2)
	return p.PaymentCode
}

func (p *PaymentTerms2) SetPercentage(value string) {
	p.Percentage = (*PercentageRate)(&value)
}

func (p *PaymentTerms2) SetAmount(value, currency string) {
	p.Amount = NewCurrencyAndAmount(value, currency)
}

// Specifies the payment terms of the underlying transaction.
type PaymentTerms3 struct {

	// Due date specified for the payment terms.
	DueDate *ISODate `xml:"DueDt,omitempty"`

	// Payment period specified for these payment terms.
	PaymentPeriod *PaymentPeriod1 `xml:"PmtPrd,omitempty"`

	// Textual description of these payment terms.
	Description []*Max140Text `xml:"Desc,omitempty"`

	// Partial payment, expressed as a percentage, for the payment terms.
	PartialPaymentPercent *PercentageRate `xml:"PrtlPmtPct,omitempty"`

	// Direct debit mandate identification specified for these payment terms.
	DirectDebitMandateIdentification []*Max35Text `xml:"DrctDbtMndtId,omitempty"`

	// Monetary value used as a basis to calculate the discount in these payment terms.
	DiscountAmount *CurrencyAndAmount `xml:"DscntAmt,omitempty"`

	// Percent rate used to calculate the discount for these payment terms.
	DiscountPercentRate *PercentageRate `xml:"DscntPctRate,omitempty"`

	// Monetary value used as a basis to calculate the discount in these payment terms.
	DiscountBasisAmount *CurrencyAndAmount `xml:"DscntBsisAmt,omitempty"`

	// Monetary value used as a basis to calculate the penalty in the payment terms.
	PenaltyAmount *CurrencyAndAmount `xml:"PnltyAmt,omitempty"`

	// Percent rate used to calculate the penalty for these payment terms.
	PenaltyPercentRate *PercentageRate `xml:"PnltyPctRate,omitempty"`

	// Amount used as a basis to calculate the penalty amount.
	PenaltyBasisAmount *CurrencyAndAmount `xml:"PnltyBsisAmt,omitempty"`
}

func (p *PaymentTerms3) SetDueDate(value string) {
	p.DueDate = (*ISODate)(&value)
}

func (p *PaymentTerms3) AddPaymentPeriod() *PaymentPeriod1 {
	p.PaymentPeriod = new(PaymentPeriod1)
	return p.PaymentPeriod
}

func (p *PaymentTerms3) AddDescription(value string) {
	p.Description = append(p.Description, (*Max140Text)(&value))
}

func (p *PaymentTerms3) SetPartialPaymentPercent(value string) {
	p.PartialPaymentPercent = (*PercentageRate)(&value)
}

func (p *PaymentTerms3) AddDirectDebitMandateIdentification(value string) {
	p.DirectDebitMandateIdentification = append(p.DirectDebitMandateIdentification, (*Max35Text)(&value))
}

func (p *PaymentTerms3) SetDiscountAmount(value, currency string) {
	p.DiscountAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTerms3) SetDiscountPercentRate(value string) {
	p.DiscountPercentRate = (*PercentageRate)(&value)
}

func (p *PaymentTerms3) SetDiscountBasisAmount(value, currency string) {
	p.DiscountBasisAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTerms3) SetPenaltyAmount(value, currency string) {
	p.PenaltyAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTerms3) SetPenaltyPercentRate(value string) {
	p.PenaltyPercentRate = (*PercentageRate)(&value)
}

func (p *PaymentTerms3) SetPenaltyBasisAmount(value, currency string) {
	p.PenaltyBasisAmount = NewCurrencyAndAmount(value, currency)
}

// Specifies the payment terms of the underlying transaction.
type PaymentTerms4 struct {

	// Specifies the payment terms using a code or other means.
	PaymentTerms *PaymentCodeOrOther1Choice `xml:"PmtTerms"`

	// Specifies if it is a fixed amount or a percentage.
	AmountOrPercentage *AmountOrPercentage2Choice `xml:"AmtOrPctg"`
}

func (p *PaymentTerms4) AddPaymentTerms() *PaymentCodeOrOther1Choice {
	p.PaymentTerms = new(PaymentCodeOrOther1Choice)
	return p.PaymentTerms
}

func (p *PaymentTerms4) AddAmountOrPercentage() *AmountOrPercentage2Choice {
	p.AmountOrPercentage = new(AmountOrPercentage2Choice)
	return p.AmountOrPercentage
}

// Specifies the payment terms of the underlying transaction.
type PaymentTerms5 struct {

	// Specifies the payment terms using a code or other means.
	PaymentTerms *PaymentCodeOrOther2Choice `xml:"PmtTerms"`

	// Specifies if it is a fixed amount or a percentage.
	AmountOrPercentage *AmountOrPercentage2Choice `xml:"AmtOrPctg"`
}

func (p *PaymentTerms5) AddPaymentTerms() *PaymentCodeOrOther2Choice {
	p.PaymentTerms = new(PaymentCodeOrOther2Choice)
	return p.PaymentTerms
}

func (p *PaymentTerms5) AddAmountOrPercentage() *AmountOrPercentage2Choice {
	p.AmountOrPercentage = new(AmountOrPercentage2Choice)
	return p.AmountOrPercentage
}

// Specifies the payment terms of the underlying transaction.
type PaymentTerms6 struct {

	// Due date specified for the payment terms.
	DueDate *ISODate `xml:"DueDt,omitempty"`

	// Payment period specified for these payment terms.
	PaymentPeriod *PaymentPeriod1 `xml:"PmtPrd,omitempty"`

	// Textual description of these payment terms.
	Description []*Max140Text `xml:"Desc,omitempty"`

	// Partial payment, expressed as a percentage, for the payment terms.
	PartialPaymentPercent *PercentageRate `xml:"PrtlPmtPct,omitempty"`

	// Direct debit mandate identification specified for these payment terms.
	DirectDebitMandateIdentification []*Max35Text `xml:"DrctDbtMndtId,omitempty"`

	// Amount used as a basis to calculate the discount amount for these payment terms.
	BasisAmount *CurrencyAndAmount `xml:"BsisAmt,omitempty"`

	// Amount of money that results from the application of an agreed discount percentage to the basis amount and payable to the creditor.
	DiscountAmount *CurrencyAndAmount `xml:"DscntAmt,omitempty"`

	// Percent rate used to calculate the discount for these payment terms.
	DiscountPercentRate *PercentageRate `xml:"DscntPctRate,omitempty"`

	// Amount of money that results from the application of an agreed penalty percentage to the basis amount and payable by the creditor.
	PenaltyAmount *CurrencyAndAmount `xml:"PnltyAmt,omitempty"`

	// Percent rate used to calculate the penalty for these payment terms.
	PenaltyPercentRate *PercentageRate `xml:"PnltyPctRate,omitempty"`
}

func (p *PaymentTerms6) SetDueDate(value string) {
	p.DueDate = (*ISODate)(&value)
}

func (p *PaymentTerms6) AddPaymentPeriod() *PaymentPeriod1 {
	p.PaymentPeriod = new(PaymentPeriod1)
	return p.PaymentPeriod
}

func (p *PaymentTerms6) AddDescription(value string) {
	p.Description = append(p.Description, (*Max140Text)(&value))
}

func (p *PaymentTerms6) SetPartialPaymentPercent(value string) {
	p.PartialPaymentPercent = (*PercentageRate)(&value)
}

func (p *PaymentTerms6) AddDirectDebitMandateIdentification(value string) {
	p.DirectDebitMandateIdentification = append(p.DirectDebitMandateIdentification, (*Max35Text)(&value))
}

func (p *PaymentTerms6) SetBasisAmount(value, currency string) {
	p.BasisAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTerms6) SetDiscountAmount(value, currency string) {
	p.DiscountAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTerms6) SetDiscountPercentRate(value string) {
	p.DiscountPercentRate = (*PercentageRate)(&value)
}

func (p *PaymentTerms6) SetPenaltyAmount(value, currency string) {
	p.PenaltyAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentTerms6) SetPenaltyPercentRate(value string) {
	p.PenaltyPercentRate = (*PercentageRate)(&value)
}
