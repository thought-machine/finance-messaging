package iso20022

// Redemption leg, or switch-out, of a switch transaction.
type SwitchRedemptionLegOrder2 struct {

	// Unique technical identifier for an instance of a leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Investment fund class related to an order.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Investment fund class related to an order.
	FinancialInstrumentQuantityChoice *FinancialInstrumentQuantity3Choice `xml:"FinInstrmQtyChc,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

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
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`
}

func (s *SwitchRedemptionLegOrder2) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchRedemptionLegOrder2) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	s.FinancialInstrumentDetails = new(FinancialInstrument6)
	return s.FinancialInstrumentDetails
}

func (s *SwitchRedemptionLegOrder2) AddFinancialInstrumentQuantityChoice() *FinancialInstrumentQuantity3Choice {
	s.FinancialInstrumentQuantityChoice = new(FinancialInstrumentQuantity3Choice)
	return s.FinancialInstrumentQuantityChoice
}

func (s *SwitchRedemptionLegOrder2) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchRedemptionLegOrder2) SetGroup1Or2Units(value string) {
	s.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (s *SwitchRedemptionLegOrder2) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (s *SwitchRedemptionLegOrder2) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (s *SwitchRedemptionLegOrder2) AddChargeDetails() *Charge8 {
	newValue := new(Charge8)
	s.ChargeDetails = append(s.ChargeDetails, newValue)
	return newValue
}

func (s *SwitchRedemptionLegOrder2) AddCommissionDetails() *Commission6 {
	newValue := new(Commission6)
	s.CommissionDetails = append(s.CommissionDetails, newValue)
	return newValue
}

func (s *SwitchRedemptionLegOrder2) AddTaxDetails() *Tax6 {
	newValue := new(Tax6)
	s.TaxDetails = append(s.TaxDetails, newValue)
	return newValue
}

func (s *SwitchRedemptionLegOrder2) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchRedemptionLegOrder2) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchRedemptionLegOrder2) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

// Redemption leg, or switch-out, of a switch transaction.
type SwitchRedemptionLegOrder3 struct {

	// Unique technical identifier for an instance of a leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Investment fund class related to an order.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls"`

	// Investment fund class related to an order.
	FinancialInstrumentQuantityChoice *FinancialInstrumentQuantity5Choice `xml:"FinInstrmQtyChc,omitempty"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Amount of money associated with a service.
	ChargeDetails []*Charge17 `xml:"ChrgDtls,omitempty"`

	// Commission linked to the execution of an investment fund order.
	CommissionDetails []*Commission10 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax16 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (s *SwitchRedemptionLegOrder3) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchRedemptionLegOrder3) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	s.FinancialInstrumentDetails = new(FinancialInstrument10)
	return s.FinancialInstrumentDetails
}

func (s *SwitchRedemptionLegOrder3) AddFinancialInstrumentQuantityChoice() *FinancialInstrumentQuantity5Choice {
	s.FinancialInstrumentQuantityChoice = new(FinancialInstrumentQuantity5Choice)
	return s.FinancialInstrumentQuantityChoice
}

func (s *SwitchRedemptionLegOrder3) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SwitchRedemptionLegOrder3) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchRedemptionLegOrder3) SetGroup1Or2Units(value string) {
	s.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (s *SwitchRedemptionLegOrder3) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SwitchRedemptionLegOrder3) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SwitchRedemptionLegOrder3) AddChargeDetails() *Charge17 {
	newValue := new(Charge17)
	s.ChargeDetails = append(s.ChargeDetails, newValue)
	return newValue
}

func (s *SwitchRedemptionLegOrder3) AddCommissionDetails() *Commission10 {
	newValue := new(Commission10)
	s.CommissionDetails = append(s.CommissionDetails, newValue)
	return newValue
}

func (s *SwitchRedemptionLegOrder3) AddTaxDetails() *Tax16 {
	newValue := new(Tax16)
	s.TaxDetails = append(s.TaxDetails, newValue)
	return newValue
}

func (s *SwitchRedemptionLegOrder3) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchRedemptionLegOrder3) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchRedemptionLegOrder3) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SwitchRedemptionLegOrder3) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SwitchRedemptionLegOrder3) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

// Redemption leg, or switch-out, of a switch transaction.
type SwitchRedemptionLegOrder6 struct {

	// Unique technical identifier for the instance of the leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Investment fund class related to the redemption leg of the order.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls"`

	// Amount of money, number of units or percentage for the subscription leg of the switch order.
	FinancialInstrumentQuantityChoice *FinancialInstrumentQuantity29Choice `xml:"FinInstrmQtyChc,omitempty"`

	// Account impacted by the investment fund order.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnit1Code `xml:"Grp1Or2Units,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Fees (charges/commission) and tax to be applied to the gross amount.
	TransactionOverhead *FeeAndTax1 `xml:"TxOvrhd,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters12 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Part of an investor's retained subscription amount that is returned by the fund in order to reimburse preliminary incentive/performance fees.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (s *SwitchRedemptionLegOrder6) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchRedemptionLegOrder6) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	s.FinancialInstrumentDetails = new(FinancialInstrument57)
	return s.FinancialInstrumentDetails
}

func (s *SwitchRedemptionLegOrder6) AddFinancialInstrumentQuantityChoice() *FinancialInstrumentQuantity29Choice {
	s.FinancialInstrumentQuantityChoice = new(FinancialInstrumentQuantity29Choice)
	return s.FinancialInstrumentQuantityChoice
}

func (s *SwitchRedemptionLegOrder6) AddInvestmentAccountDetails() *InvestmentAccount58 {
	s.InvestmentAccountDetails = new(InvestmentAccount58)
	return s.InvestmentAccountDetails
}

func (s *SwitchRedemptionLegOrder6) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchRedemptionLegOrder6) SetGroup1Or2Units(value string) {
	s.Group1Or2Units = (*UKTaxGroupUnit1Code)(&value)
}

func (s *SwitchRedemptionLegOrder6) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SwitchRedemptionLegOrder6) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SwitchRedemptionLegOrder6) AddTransactionOverhead() *FeeAndTax1 {
	s.TransactionOverhead = new(FeeAndTax1)
	return s.TransactionOverhead
}

func (s *SwitchRedemptionLegOrder6) AddSettlementAndCustodyDetails() *FundSettlementParameters12 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters12)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchRedemptionLegOrder6) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchRedemptionLegOrder6) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SwitchRedemptionLegOrder6) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SwitchRedemptionLegOrder6) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}
