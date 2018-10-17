package iso20022

// Subscription leg, or switch-in, of a switch order.
type SwitchSubscriptionLegOrder2 struct {

	// Unique technical identifier for an instance of a leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Investment fund class related to an order.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Investment fund class related to an order.
	FinancialInstrumentQuantityChoice *FinancialInstrumentQuantity4Choice `xml:"FinInstrmQtyChc,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Amount of money associated with a service.
	ChargeDetails []*Charge8 `xml:"ChrgDtls,omitempty"`

	// Commission linked to the execution of an investment fund order.
	CommissionDetails []*Commission6 `xml:"ComssnDtls,omitempty"`

	// Tax related to an investment fund order.
	TaxDetails []*Tax6 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *NameAndAddress4 `xml:"PhysDlvryDtls,omitempty"`
}

func (s *SwitchSubscriptionLegOrder2) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchSubscriptionLegOrder2) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	s.FinancialInstrumentDetails = new(FinancialInstrument6)
	return s.FinancialInstrumentDetails
}

func (s *SwitchSubscriptionLegOrder2) AddFinancialInstrumentQuantityChoice() *FinancialInstrumentQuantity4Choice {
	s.FinancialInstrumentQuantityChoice = new(FinancialInstrumentQuantity4Choice)
	return s.FinancialInstrumentQuantityChoice
}

func (s *SwitchSubscriptionLegOrder2) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchSubscriptionLegOrder2) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegOrder2) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegOrder2) AddChargeDetails() *Charge8 {
	newValue := new(Charge8)
	s.ChargeDetails = append(s.ChargeDetails, newValue)
	return newValue
}

func (s *SwitchSubscriptionLegOrder2) AddCommissionDetails() *Commission6 {
	newValue := new(Commission6)
	s.CommissionDetails = append(s.CommissionDetails, newValue)
	return newValue
}

func (s *SwitchSubscriptionLegOrder2) AddTaxDetails() *Tax6 {
	newValue := new(Tax6)
	s.TaxDetails = append(s.TaxDetails, newValue)
	return newValue
}

func (s *SwitchSubscriptionLegOrder2) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchSubscriptionLegOrder2) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchSubscriptionLegOrder2) AddPhysicalDeliveryDetails() *NameAndAddress4 {
	s.PhysicalDeliveryDetails = new(NameAndAddress4)
	return s.PhysicalDeliveryDetails
}

// Subscription leg, or switch-in, of a switch order.
type SwitchSubscriptionLegOrder3 struct {

	// Unique technical identifier for an instance of a leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Investment fund class related to an order.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls"`

	// Investment fund class related to an order.
	FinancialInstrumentQuantityChoice *FinancialInstrumentQuantity6Choice `xml:"FinInstrmQtyChc,omitempty"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Amount of money associated with a service.
	ChargeDetails []*Charge17 `xml:"ChrgDtls,omitempty"`

	// Commission linked to the execution of an investment fund order.
	CommissionDetails []*Commission10 `xml:"ComssnDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax16 `xml:"TaxDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *NameAndAddress4 `xml:"PhysDlvryDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (s *SwitchSubscriptionLegOrder3) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchSubscriptionLegOrder3) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	s.FinancialInstrumentDetails = new(FinancialInstrument10)
	return s.FinancialInstrumentDetails
}

func (s *SwitchSubscriptionLegOrder3) AddFinancialInstrumentQuantityChoice() *FinancialInstrumentQuantity6Choice {
	s.FinancialInstrumentQuantityChoice = new(FinancialInstrumentQuantity6Choice)
	return s.FinancialInstrumentQuantityChoice
}

func (s *SwitchSubscriptionLegOrder3) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SwitchSubscriptionLegOrder3) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchSubscriptionLegOrder3) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegOrder3) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegOrder3) AddChargeDetails() *Charge17 {
	newValue := new(Charge17)
	s.ChargeDetails = append(s.ChargeDetails, newValue)
	return newValue
}

func (s *SwitchSubscriptionLegOrder3) AddCommissionDetails() *Commission10 {
	newValue := new(Commission10)
	s.CommissionDetails = append(s.CommissionDetails, newValue)
	return newValue
}

func (s *SwitchSubscriptionLegOrder3) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchSubscriptionLegOrder3) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchSubscriptionLegOrder3) AddTaxDetails() *Tax16 {
	newValue := new(Tax16)
	s.TaxDetails = append(s.TaxDetails, newValue)
	return newValue
}

func (s *SwitchSubscriptionLegOrder3) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SwitchSubscriptionLegOrder3) AddPhysicalDeliveryDetails() *NameAndAddress4 {
	s.PhysicalDeliveryDetails = new(NameAndAddress4)
	return s.PhysicalDeliveryDetails
}

func (s *SwitchSubscriptionLegOrder3) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

// Subscription leg, or switch-in, of a switch order.
type SwitchSubscriptionLegOrder6 struct {

	// Unique technical identifier for the instance of the leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Investment fund class related to the subscription leg of the order.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls"`

	// Amount of money, number of units or percentage for the redemption leg of the switch order.
	FinancialInstrumentQuantityChoice *FinancialInstrumentQuantity26Choice `xml:"FinInstrmQtyChc,omitempty"`

	// Account impacted by the investment fund order.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Fees (charges/commission) and tax to be applied to the net amount.
	TransactionOverhead *FeeAndTax1 `xml:"TxOvrhd,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters11 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *NameAndAddress4 `xml:"PhysDlvryDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive/performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (s *SwitchSubscriptionLegOrder6) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchSubscriptionLegOrder6) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	s.FinancialInstrumentDetails = new(FinancialInstrument57)
	return s.FinancialInstrumentDetails
}

func (s *SwitchSubscriptionLegOrder6) AddFinancialInstrumentQuantityChoice() *FinancialInstrumentQuantity26Choice {
	s.FinancialInstrumentQuantityChoice = new(FinancialInstrumentQuantity26Choice)
	return s.FinancialInstrumentQuantityChoice
}

func (s *SwitchSubscriptionLegOrder6) AddInvestmentAccountDetails() *InvestmentAccount58 {
	s.InvestmentAccountDetails = new(InvestmentAccount58)
	return s.InvestmentAccountDetails
}

func (s *SwitchSubscriptionLegOrder6) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchSubscriptionLegOrder6) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegOrder6) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegOrder6) AddTransactionOverhead() *FeeAndTax1 {
	s.TransactionOverhead = new(FeeAndTax1)
	return s.TransactionOverhead
}

func (s *SwitchSubscriptionLegOrder6) AddSettlementAndCustodyDetails() *FundSettlementParameters11 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters11)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchSubscriptionLegOrder6) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchSubscriptionLegOrder6) AddPhysicalDeliveryDetails() *NameAndAddress4 {
	s.PhysicalDeliveryDetails = new(NameAndAddress4)
	return s.PhysicalDeliveryDetails
}

func (s *SwitchSubscriptionLegOrder6) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SwitchSubscriptionLegOrder6) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}
