package iso20022

// Payment obligation contracted between two financial institutions related to the financing of a commercial transaction.
type PaymentObligation1 struct {

	// Bank that has to pay under the obligation.
	ObligorBank *BICIdentification1 `xml:"OblgrBk"`

	// Bank that will be paid under the obligation.
	RecipientBank *BICIdentification1 `xml:"RcptBk"`

	// Maximum amount that will be paid under the obligation.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Maximum amount that will be paid under the obligation, expressed as a percentage of the purchase order net amount.
	Percentage *PercentageRate `xml:"Pctg"`

	// Amount of the charges taken by the obligor bank.
	ChargesAmount *CurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Amount of the charges expressed as a percentage of the amount paid by the obligor bank.
	ChargesPercentage *PercentageRate `xml:"ChrgsPctg,omitempty"`

	// Date at which the obligation will expire.
	ExpiryDate *ISODate `xml:"XpryDt"`

	// Country of which the law governs the bank payment obligation.
	ApplicableLaw *CountryCode `xml:"AplblLaw,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	PaymentTerms []*PaymentTerms2 `xml:"PmtTerms,omitempty"`

	// Instruction between two clearing agents stipulating the cash transfer characteristics between the two parties.
	SettlementTerms *SettlementTerms2 `xml:"SttlmTerms,omitempty"`
}

func (p *PaymentObligation1) AddObligorBank() *BICIdentification1 {
	p.ObligorBank = new(BICIdentification1)
	return p.ObligorBank
}

func (p *PaymentObligation1) AddRecipientBank() *BICIdentification1 {
	p.RecipientBank = new(BICIdentification1)
	return p.RecipientBank
}

func (p *PaymentObligation1) SetAmount(value, currency string) {
	p.Amount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentObligation1) SetPercentage(value string) {
	p.Percentage = (*PercentageRate)(&value)
}

func (p *PaymentObligation1) SetChargesAmount(value, currency string) {
	p.ChargesAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentObligation1) SetChargesPercentage(value string) {
	p.ChargesPercentage = (*PercentageRate)(&value)
}

func (p *PaymentObligation1) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISODate)(&value)
}

func (p *PaymentObligation1) SetApplicableLaw(value string) {
	p.ApplicableLaw = (*CountryCode)(&value)
}

func (p *PaymentObligation1) AddPaymentTerms() *PaymentTerms2 {
	newValue := new(PaymentTerms2)
	p.PaymentTerms = append(p.PaymentTerms, newValue)
	return newValue
}

func (p *PaymentObligation1) AddSettlementTerms() *SettlementTerms2 {
	p.SettlementTerms = new(SettlementTerms2)
	return p.SettlementTerms
}

// Payment obligation contracted between two financial institutions related to the financing of a commercial transaction.
type PaymentObligation2 struct {

	// Bank that has to pay under the obligation.
	ObligorBank *BICIdentification1 `xml:"OblgrBk"`

	// Bank that will be paid under the obligation.
	RecipientBank *BICIdentification1 `xml:"RcptBk"`

	// Payment obligation amount specified as an amount or percentage.
	PaymentObligationAmount *AmountOrPercentage2Choice `xml:"PmtOblgtnAmt"`

	// Charges related to the payment obligation.
	Charges []*Charges5 `xml:"Chrgs,omitempty"`

	// Date at which the obligation will expire.
	ExpiryDate *ISODate `xml:"XpryDt"`

	// Rules which apply to the BPO (Bank Payment Obligation).
	ApplicableRules *BPOApplicableRules1Choice `xml:"AplblRules,omitempty"`

	// Country of which the law governs the bank payment obligation.
	ApplicableLaw *CountryCode `xml:"AplblLaw,omitempty"`

	// Location and forum for dispute resolution.
	PlaceOfJurisdiction *Location2 `xml:"PlcOfJursdctn,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	PaymentTerms []*PaymentTerms4 `xml:"PmtTerms,omitempty"`

	// Instruction between two clearing agents stipulating the cash transfer characteristics between the two parties.
	SettlementTerms *SettlementTerms3 `xml:"SttlmTerms,omitempty"`
}

func (p *PaymentObligation2) AddObligorBank() *BICIdentification1 {
	p.ObligorBank = new(BICIdentification1)
	return p.ObligorBank
}

func (p *PaymentObligation2) AddRecipientBank() *BICIdentification1 {
	p.RecipientBank = new(BICIdentification1)
	return p.RecipientBank
}

func (p *PaymentObligation2) AddPaymentObligationAmount() *AmountOrPercentage2Choice {
	p.PaymentObligationAmount = new(AmountOrPercentage2Choice)
	return p.PaymentObligationAmount
}

func (p *PaymentObligation2) AddCharges() *Charges5 {
	newValue := new(Charges5)
	p.Charges = append(p.Charges, newValue)
	return newValue
}

func (p *PaymentObligation2) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISODate)(&value)
}

func (p *PaymentObligation2) AddApplicableRules() *BPOApplicableRules1Choice {
	p.ApplicableRules = new(BPOApplicableRules1Choice)
	return p.ApplicableRules
}

func (p *PaymentObligation2) SetApplicableLaw(value string) {
	p.ApplicableLaw = (*CountryCode)(&value)
}

func (p *PaymentObligation2) AddPlaceOfJurisdiction() *Location2 {
	p.PlaceOfJurisdiction = new(Location2)
	return p.PlaceOfJurisdiction
}

func (p *PaymentObligation2) AddPaymentTerms() *PaymentTerms4 {
	newValue := new(PaymentTerms4)
	p.PaymentTerms = append(p.PaymentTerms, newValue)
	return newValue
}

func (p *PaymentObligation2) AddSettlementTerms() *SettlementTerms3 {
	p.SettlementTerms = new(SettlementTerms3)
	return p.SettlementTerms
}
