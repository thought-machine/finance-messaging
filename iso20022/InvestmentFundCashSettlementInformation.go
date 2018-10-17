package iso20022

// Settlement instructions to be used to transfer cash from the Debtor to the Creditor.
type InvestmentFundCashSettlementInformation3 struct {

	// Instrument that has or represents monetary value and is used to process a payment instruction for a subscription payment.
	SubscriptionPaymentInstrument *PaymentInstrument8 `xml:"SbcptPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a redemption payment.
	RedemptionPaymentInstrument *PaymentInstrument9 `xml:"RedPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a dividend payment.
	DividendPaymentInstrument *PaymentInstrument9 `xml:"DvddPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a savings plan payment.
	SavingsPlanPaymentInstrument *PaymentInstrument8 `xml:"SvgsPlanPmtInstrm,omitempty"`
}

func (i *InvestmentFundCashSettlementInformation3) AddSubscriptionPaymentInstrument() *PaymentInstrument8 {
	i.SubscriptionPaymentInstrument = new(PaymentInstrument8)
	return i.SubscriptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation3) AddRedemptionPaymentInstrument() *PaymentInstrument9 {
	i.RedemptionPaymentInstrument = new(PaymentInstrument9)
	return i.RedemptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation3) AddDividendPaymentInstrument() *PaymentInstrument9 {
	i.DividendPaymentInstrument = new(PaymentInstrument9)
	return i.DividendPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation3) AddSavingsPlanPaymentInstrument() *PaymentInstrument8 {
	i.SavingsPlanPaymentInstrument = new(PaymentInstrument8)
	return i.SavingsPlanPaymentInstrument
}

// Settlement instructions to be used to transfer cash from the Debtor to the Creditor.
type InvestmentFundCashSettlementInformation4 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a subscription payment.
	SubscriptionPaymentInstrument *PaymentInstrument8 `xml:"SbcptPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a redemption payment.
	RedemptionPaymentInstrument *PaymentInstrument9 `xml:"RedPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a dividend payment.
	DividendPaymentInstrument *PaymentInstrument9 `xml:"DvddPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a savings plan payment.
	SavingsPlanPaymentInstrument *PaymentInstrument8 `xml:"SvgsPlanPmtInstrm,omitempty"`
}

func (i *InvestmentFundCashSettlementInformation4) SetModificationScopeIndication(value string) {
	i.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (i *InvestmentFundCashSettlementInformation4) AddSubscriptionPaymentInstrument() *PaymentInstrument8 {
	i.SubscriptionPaymentInstrument = new(PaymentInstrument8)
	return i.SubscriptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation4) AddRedemptionPaymentInstrument() *PaymentInstrument9 {
	i.RedemptionPaymentInstrument = new(PaymentInstrument9)
	return i.RedemptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation4) AddDividendPaymentInstrument() *PaymentInstrument9 {
	i.DividendPaymentInstrument = new(PaymentInstrument9)
	return i.DividendPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation4) AddSavingsPlanPaymentInstrument() *PaymentInstrument8 {
	i.SavingsPlanPaymentInstrument = new(PaymentInstrument8)
	return i.SavingsPlanPaymentInstrument
}

// Settlement instructions to be used to transfer cash from the Debtor to the Creditor.
type InvestmentFundCashSettlementInformation5 struct {

	// Instrument that has or represents monetary value and is used to process a payment instruction for a subscription payment.
	SubscriptionPaymentInstrument *PaymentInstrument11 `xml:"SbcptPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a redemption payment.
	RedemptionPaymentInstrument *PaymentInstrument10 `xml:"RedPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a dividend payment.
	DividendPaymentInstrument *PaymentInstrument10 `xml:"DvddPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a savings plan payment.
	SavingsPlanPaymentInstrument *PaymentInstrument11 `xml:"SvgsPlanPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for an interest payment.
	InterestPaymentInstrument *PaymentInstrument10 `xml:"IntrstPmtInstrm,omitempty"`
}

func (i *InvestmentFundCashSettlementInformation5) AddSubscriptionPaymentInstrument() *PaymentInstrument11 {
	i.SubscriptionPaymentInstrument = new(PaymentInstrument11)
	return i.SubscriptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation5) AddRedemptionPaymentInstrument() *PaymentInstrument10 {
	i.RedemptionPaymentInstrument = new(PaymentInstrument10)
	return i.RedemptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation5) AddDividendPaymentInstrument() *PaymentInstrument10 {
	i.DividendPaymentInstrument = new(PaymentInstrument10)
	return i.DividendPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation5) AddSavingsPlanPaymentInstrument() *PaymentInstrument11 {
	i.SavingsPlanPaymentInstrument = new(PaymentInstrument11)
	return i.SavingsPlanPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation5) AddInterestPaymentInstrument() *PaymentInstrument10 {
	i.InterestPaymentInstrument = new(PaymentInstrument10)
	return i.InterestPaymentInstrument
}

// Settlement instructions to be used to transfer cash from the Debtor to the Creditor.
type InvestmentFundCashSettlementInformation6 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a subscription payment.
	SubscriptionPaymentInstrument *PaymentInstrument11 `xml:"SbcptPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a redemption payment.
	RedemptionPaymentInstrument *PaymentInstrument10 `xml:"RedPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a dividend payment.
	DividendPaymentInstrument *PaymentInstrument10 `xml:"DvddPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a savings plan payment.
	SavingsPlanPaymentInstrument *PaymentInstrument11 `xml:"SvgsPlanPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for an interest payment.
	InterestPaymentInstrument *PaymentInstrument10 `xml:"IntrstPmtInstrm,omitempty"`
}

func (i *InvestmentFundCashSettlementInformation6) SetModificationScopeIndication(value string) {
	i.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (i *InvestmentFundCashSettlementInformation6) AddSubscriptionPaymentInstrument() *PaymentInstrument11 {
	i.SubscriptionPaymentInstrument = new(PaymentInstrument11)
	return i.SubscriptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation6) AddRedemptionPaymentInstrument() *PaymentInstrument10 {
	i.RedemptionPaymentInstrument = new(PaymentInstrument10)
	return i.RedemptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation6) AddDividendPaymentInstrument() *PaymentInstrument10 {
	i.DividendPaymentInstrument = new(PaymentInstrument10)
	return i.DividendPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation6) AddSavingsPlanPaymentInstrument() *PaymentInstrument11 {
	i.SavingsPlanPaymentInstrument = new(PaymentInstrument11)
	return i.SavingsPlanPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation6) AddInterestPaymentInstrument() *PaymentInstrument10 {
	i.InterestPaymentInstrument = new(PaymentInstrument10)
	return i.InterestPaymentInstrument
}

// Settlement instructions to be used to transfer cash from the debtor to the creditor.
type InvestmentFundCashSettlementInformation7 struct {

	// Instrument that has or represents monetary value and is used to process a payment instruction for a subscription payment.
	SubscriptionPaymentInstrument *PaymentInstrument11 `xml:"SbcptPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a redemption payment.
	RedemptionPaymentInstrument *PaymentInstrument12 `xml:"RedPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a dividend payment.
	DividendPaymentInstrument *PaymentInstrument12 `xml:"DvddPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a savings plan payment.
	SavingsPlanPaymentInstrument *PaymentInstrument11 `xml:"SvgsPlanPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for an interest payment.
	InterestPaymentInstrument *PaymentInstrument12 `xml:"IntrstPmtInstrm,omitempty"`
}

func (i *InvestmentFundCashSettlementInformation7) AddSubscriptionPaymentInstrument() *PaymentInstrument11 {
	i.SubscriptionPaymentInstrument = new(PaymentInstrument11)
	return i.SubscriptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation7) AddRedemptionPaymentInstrument() *PaymentInstrument12 {
	i.RedemptionPaymentInstrument = new(PaymentInstrument12)
	return i.RedemptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation7) AddDividendPaymentInstrument() *PaymentInstrument12 {
	i.DividendPaymentInstrument = new(PaymentInstrument12)
	return i.DividendPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation7) AddSavingsPlanPaymentInstrument() *PaymentInstrument11 {
	i.SavingsPlanPaymentInstrument = new(PaymentInstrument11)
	return i.SavingsPlanPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation7) AddInterestPaymentInstrument() *PaymentInstrument12 {
	i.InterestPaymentInstrument = new(PaymentInstrument12)
	return i.InterestPaymentInstrument
}

// Settlement instructions to be used to transfer cash from the debtor to the creditor.
type InvestmentFundCashSettlementInformation8 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification2Code `xml:"ModScpIndctn"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a subscription payment.
	SubscriptionPaymentInstrument *PaymentInstrument11 `xml:"SbcptPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a redemption payment.
	RedemptionPaymentInstrument *PaymentInstrument12 `xml:"RedPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a dividend payment.
	DividendPaymentInstrument *PaymentInstrument12 `xml:"DvddPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for a savings plan payment.
	SavingsPlanPaymentInstrument *PaymentInstrument11 `xml:"SvgsPlanPmtInstrm,omitempty"`

	// Instrument that has or represents monetary value and is used to process a payment instruction for an interest payment.
	InterestPaymentInstrument *PaymentInstrument12 `xml:"IntrstPmtInstrm,omitempty"`
}

func (i *InvestmentFundCashSettlementInformation8) SetModificationScopeIndication(value string) {
	i.ModificationScopeIndication = (*DataModification2Code)(&value)
}

func (i *InvestmentFundCashSettlementInformation8) AddSubscriptionPaymentInstrument() *PaymentInstrument11 {
	i.SubscriptionPaymentInstrument = new(PaymentInstrument11)
	return i.SubscriptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation8) AddRedemptionPaymentInstrument() *PaymentInstrument12 {
	i.RedemptionPaymentInstrument = new(PaymentInstrument12)
	return i.RedemptionPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation8) AddDividendPaymentInstrument() *PaymentInstrument12 {
	i.DividendPaymentInstrument = new(PaymentInstrument12)
	return i.DividendPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation8) AddSavingsPlanPaymentInstrument() *PaymentInstrument11 {
	i.SavingsPlanPaymentInstrument = new(PaymentInstrument11)
	return i.SavingsPlanPaymentInstrument
}

func (i *InvestmentFundCashSettlementInformation8) AddInterestPaymentInstrument() *PaymentInstrument12 {
	i.InterestPaymentInstrument = new(PaymentInstrument12)
	return i.InterestPaymentInstrument
}
