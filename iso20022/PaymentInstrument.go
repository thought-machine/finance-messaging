package iso20022

// Instrument that has or represents monetary value and is used to process a payment instruction.
type PaymentInstrument10 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Choice of payment instruments.
	PaymentInstrument *PaymentInstrument16Choice `xml:"PmtInstrm"`

	// Percentage of the dividend payment not to be reinvested.
	DividendPercentage *PercentageBoundedRate `xml:"DvddPctg,omitempty"`
}

func (p *PaymentInstrument10) SetSettlementCurrency(value string) {
	p.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (p *PaymentInstrument10) AddPaymentInstrument() *PaymentInstrument16Choice {
	p.PaymentInstrument = new(PaymentInstrument16Choice)
	return p.PaymentInstrument
}

func (p *PaymentInstrument10) SetDividendPercentage(value string) {
	p.DividendPercentage = (*PercentageBoundedRate)(&value)
}

// Instrument that has or represents monetary value and is used to process a payment instruction.
type PaymentInstrument11 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Choice of payment instruments.
	PaymentInstrument *PaymentInstrument17Choice `xml:"PmtInstrm"`
}

func (p *PaymentInstrument11) SetSettlementCurrency(value string) {
	p.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (p *PaymentInstrument11) AddPaymentInstrument() *PaymentInstrument17Choice {
	p.PaymentInstrument = new(PaymentInstrument17Choice)
	return p.PaymentInstrument
}

// Instrument that has or represents monetary value and is used to process a payment instruction.
type PaymentInstrument12 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Choice of payment instruments.
	PaymentInstrument *PaymentInstrument16Choice `xml:"PmtInstrm"`

	// Percentage of the dividend payment not to be reinvested, that is, to be paid in cash.
	DividendPercentage *PercentageBoundedRate `xml:"DvddPctg,omitempty"`
}

func (p *PaymentInstrument12) SetSettlementCurrency(value string) {
	p.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (p *PaymentInstrument12) AddPaymentInstrument() *PaymentInstrument16Choice {
	p.PaymentInstrument = new(PaymentInstrument16Choice)
	return p.PaymentInstrument
}

func (p *PaymentInstrument12) SetDividendPercentage(value string) {
	p.DividendPercentage = (*PercentageBoundedRate)(&value)
}

// Instrument used to process a payment instruction.
type PaymentInstrument13 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Percentage of the dividend payment not to be reinvested, that is, to be paid in cash.
	DividendPercentage *PercentageBoundedRate `xml:"DvddPctg,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a subscription payment.
	SubscriptionPaymentInstrument *PaymentInstrument18Choice `xml:"SbcptPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a redemption payment.
	RedemptionPaymentInstrument *PaymentInstrument19Choice `xml:"RedPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a dividend payment.
	DividendPaymentInstrument *PaymentInstrument19Choice `xml:"DvddPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a savings plan payment.
	SavingsPlanPaymentInstrument *PaymentInstrument18Choice `xml:"SvgsPlanPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for an interest payment.
	InterestPaymentInstrument *PaymentInstrument19Choice `xml:"IntrstPmtInstrm,omitempty"`
}

func (p *PaymentInstrument13) SetSettlementCurrency(value string) {
	p.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (p *PaymentInstrument13) SetDividendPercentage(value string) {
	p.DividendPercentage = (*PercentageBoundedRate)(&value)
}

func (p *PaymentInstrument13) AddSubscriptionPaymentInstrument() *PaymentInstrument18Choice {
	p.SubscriptionPaymentInstrument = new(PaymentInstrument18Choice)
	return p.SubscriptionPaymentInstrument
}

func (p *PaymentInstrument13) AddRedemptionPaymentInstrument() *PaymentInstrument19Choice {
	p.RedemptionPaymentInstrument = new(PaymentInstrument19Choice)
	return p.RedemptionPaymentInstrument
}

func (p *PaymentInstrument13) AddDividendPaymentInstrument() *PaymentInstrument19Choice {
	p.DividendPaymentInstrument = new(PaymentInstrument19Choice)
	return p.DividendPaymentInstrument
}

func (p *PaymentInstrument13) AddSavingsPlanPaymentInstrument() *PaymentInstrument18Choice {
	p.SavingsPlanPaymentInstrument = new(PaymentInstrument18Choice)
	return p.SavingsPlanPaymentInstrument
}

func (p *PaymentInstrument13) AddInterestPaymentInstrument() *PaymentInstrument19Choice {
	p.InterestPaymentInstrument = new(PaymentInstrument19Choice)
	return p.InterestPaymentInstrument
}

// Instrument that has or represents monetary value and is used to process a payment instruction.
type PaymentInstrument8 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Cash account to debit for the payment of a subscription or of a savings plan to an investment fund.
	CashAccountDetails []*CashAccount4 `xml:"CshAcctDtls"`

	// Settlement instructions for a payment by card.
	PaymentCardDetails *PaymentCard2 `xml:"PmtCardDtls"`

	// Settlement instructions for a payment by direct debit.
	DirectDebitDetails *DirectDebitMandate4 `xml:"DrctDbtDtls"`

	// Indicates whether the payment is done via cheque.
	Cheque *YesNoIndicator `xml:"Chq"`

	// Indicates whether the payment is done via draft.
	BankersDraft *YesNoIndicator `xml:"BkrsDrft"`
}

func (p *PaymentInstrument8) SetSettlementCurrency(value string) {
	p.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (p *PaymentInstrument8) AddCashAccountDetails() *CashAccount4 {
	newValue := new(CashAccount4)
	p.CashAccountDetails = append(p.CashAccountDetails, newValue)
	return newValue
}

func (p *PaymentInstrument8) AddPaymentCardDetails() *PaymentCard2 {
	p.PaymentCardDetails = new(PaymentCard2)
	return p.PaymentCardDetails
}

func (p *PaymentInstrument8) AddDirectDebitDetails() *DirectDebitMandate4 {
	p.DirectDebitDetails = new(DirectDebitMandate4)
	return p.DirectDebitDetails
}

func (p *PaymentInstrument8) SetCheque(value string) {
	p.Cheque = (*YesNoIndicator)(&value)
}

func (p *PaymentInstrument8) SetBankersDraft(value string) {
	p.BankersDraft = (*YesNoIndicator)(&value)
}

// Instrument that has or represents monetary value and is used to process a payment instruction.
type PaymentInstrument9 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Cash account to credit for the payment of the dividends or of the redeemed investments funds.
	CashAccountDetails []*CashAccount4 `xml:"CshAcctDtls"`

	// Settlement instructions for a payment by cheque.
	ChequeDetails *Cheque4 `xml:"ChqDtls"`

	// Settlement instructions for a payment by draft.
	BankersDraftDetails *Cheque4 `xml:"BkrsDrftDtls"`
}

func (p *PaymentInstrument9) SetSettlementCurrency(value string) {
	p.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (p *PaymentInstrument9) AddCashAccountDetails() *CashAccount4 {
	newValue := new(CashAccount4)
	p.CashAccountDetails = append(p.CashAccountDetails, newValue)
	return newValue
}

func (p *PaymentInstrument9) AddChequeDetails() *Cheque4 {
	p.ChequeDetails = new(Cheque4)
	return p.ChequeDetails
}

func (p *PaymentInstrument9) AddBankersDraftDetails() *Cheque4 {
	p.BankersDraftDetails = new(Cheque4)
	return p.BankersDraftDetails
}
