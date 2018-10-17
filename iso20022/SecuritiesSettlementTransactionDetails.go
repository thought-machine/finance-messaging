package iso20022

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails10 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters5 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages1 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails26 `xml:"TradDtls,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount30 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails50 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties17 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection32 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails10) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters5 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters5)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails10) AddLinkages() *Linkages1 {
	newValue := new(Linkages1)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails10) AddTradeDetails() *SecuritiesTradeDetails26 {
	s.TradeDetails = new(SecuritiesTradeDetails26)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails10) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionDetails10) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails10) AddQuantityAndAccountDetails() *QuantityAndAccount30 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount30)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails10) AddSettlementParameters() *SettlementDetails50 {
	s.SettlementParameters = new(SettlementDetails50)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails10) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails10) AddDeliveringSettlementParties() *SettlementParties11 {
	s.DeliveringSettlementParties = new(SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails10) AddReceivingSettlementParties() *SettlementParties11 {
	s.ReceivingSettlementParties = new(SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails10) AddCashParties() *CashParties17 {
	s.CashParties = new(CashParties17)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails10) AddSettlementAmount() *AmountAndDirection32 {
	s.SettlementAmount = new(AmountAndDirection32)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails10) AddOtherAmounts() *OtherAmounts14 {
	s.OtherAmounts = new(OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails10) AddOtherBusinessParties() *OtherParties19 {
	s.OtherBusinessParties = new(OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails10) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails10) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails14 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters5 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages1 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails34 `xml:"TradDtls,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount30 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails68 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties17 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection37 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails14) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters5 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters5)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails14) AddLinkages() *Linkages1 {
	newValue := new(Linkages1)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails14) AddTradeDetails() *SecuritiesTradeDetails34 {
	s.TradeDetails = new(SecuritiesTradeDetails34)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails14) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionDetails14) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails14) AddQuantityAndAccountDetails() *QuantityAndAccount30 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount30)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails14) AddSettlementParameters() *SettlementDetails68 {
	s.SettlementParameters = new(SettlementDetails68)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails14) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails14) AddDeliveringSettlementParties() *SettlementParties11 {
	s.DeliveringSettlementParties = new(SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails14) AddReceivingSettlementParties() *SettlementParties11 {
	s.ReceivingSettlementParties = new(SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails14) AddCashParties() *CashParties17 {
	s.CashParties = new(CashParties17)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails14) AddSettlementAmount() *AmountAndDirection37 {
	s.SettlementAmount = new(AmountAndDirection37)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails14) AddOtherAmounts() *OtherAmounts14 {
	s.OtherAmounts = new(OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails14) AddOtherBusinessParties() *OtherParties19 {
	s.OtherBusinessParties = new(OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails14) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails14) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails15 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters6 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages1 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails32 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount29 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails68 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties26 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties26 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties17 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection37 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails15) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters6 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters6)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails15) AddLinkages() *Linkages1 {
	newValue := new(Linkages1)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails15) AddTradeDetails() *SecuritiesTradeDetails32 {
	s.TradeDetails = new(SecuritiesTradeDetails32)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails15) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails15) AddQuantityAndAccountDetails() *QuantityAndAccount29 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount29)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails15) AddSettlementParameters() *SettlementDetails68 {
	s.SettlementParameters = new(SettlementDetails68)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails15) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails15) AddDeliveringSettlementParties() *SettlementParties26 {
	s.DeliveringSettlementParties = new(SettlementParties26)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails15) AddReceivingSettlementParties() *SettlementParties26 {
	s.ReceivingSettlementParties = new(SettlementParties26)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails15) AddCashParties() *CashParties17 {
	s.CashParties = new(CashParties17)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails15) AddSettlementAmount() *AmountAndDirection37 {
	s.SettlementAmount = new(AmountAndDirection37)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails15) AddOtherAmounts() *OtherAmounts14 {
	s.OtherAmounts = new(OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails15) AddOtherBusinessParties() *OtherParties19 {
	s.OtherBusinessParties = new(OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails15) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails15) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails16 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters6 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages21 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails32 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount29 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails68 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties26 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties26 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties17 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection37 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails16) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters6 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters6)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails16) AddLinkages() *Linkages21 {
	newValue := new(Linkages21)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails16) AddTradeDetails() *SecuritiesTradeDetails32 {
	s.TradeDetails = new(SecuritiesTradeDetails32)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails16) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails16) AddQuantityAndAccountDetails() *QuantityAndAccount29 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount29)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails16) AddSettlementParameters() *SettlementDetails68 {
	s.SettlementParameters = new(SettlementDetails68)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails16) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails16) AddDeliveringSettlementParties() *SettlementParties26 {
	s.DeliveringSettlementParties = new(SettlementParties26)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails16) AddReceivingSettlementParties() *SettlementParties26 {
	s.ReceivingSettlementParties = new(SettlementParties26)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails16) AddCashParties() *CashParties17 {
	s.CashParties = new(CashParties17)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails16) AddSettlementAmount() *AmountAndDirection37 {
	s.SettlementAmount = new(AmountAndDirection37)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails16) AddOtherAmounts() *OtherAmounts14 {
	s.OtherAmounts = new(OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails16) AddOtherBusinessParties() *OtherParties19 {
	s.OtherBusinessParties = new(OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails16) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails16) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails2 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters5 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages1 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails26 `xml:"TradDtls,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount30 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails50 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties8 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection32 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails2) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters5 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters5)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails2) AddLinkages() *Linkages1 {
	newValue := new(Linkages1)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails2) AddTradeDetails() *SecuritiesTradeDetails26 {
	s.TradeDetails = new(SecuritiesTradeDetails26)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails2) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionDetails2) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails2) AddQuantityAndAccountDetails() *QuantityAndAccount30 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount30)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails2) AddSettlementParameters() *SettlementDetails50 {
	s.SettlementParameters = new(SettlementDetails50)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails2) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails2) AddDeliveringSettlementParties() *SettlementParties11 {
	s.DeliveringSettlementParties = new(SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails2) AddReceivingSettlementParties() *SettlementParties11 {
	s.ReceivingSettlementParties = new(SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails2) AddCashParties() *CashParties8 {
	s.CashParties = new(CashParties8)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails2) AddSettlementAmount() *AmountAndDirection32 {
	s.SettlementAmount = new(AmountAndDirection32)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails2) AddOtherAmounts() *OtherAmounts14 {
	s.OtherAmounts = new(OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails2) AddOtherBusinessParties() *OtherParties19 {
	s.OtherBusinessParties = new(OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails2) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails2) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails20 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters13 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages40 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails50 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount44 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails94 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties39 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties39 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties26 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection45 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts28 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters4 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails20) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters13 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters13)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails20) AddLinkages() *Linkages40 {
	newValue := new(Linkages40)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails20) AddTradeDetails() *SecuritiesTradeDetails50 {
	s.TradeDetails = new(SecuritiesTradeDetails50)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails20) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails20) AddQuantityAndAccountDetails() *QuantityAndAccount44 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount44)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails20) AddSettlementParameters() *SettlementDetails94 {
	s.SettlementParameters = new(SettlementDetails94)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails20) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails20) AddDeliveringSettlementParties() *SettlementParties39 {
	s.DeliveringSettlementParties = new(SettlementParties39)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails20) AddReceivingSettlementParties() *SettlementParties39 {
	s.ReceivingSettlementParties = new(SettlementParties39)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails20) AddCashParties() *CashParties26 {
	s.CashParties = new(CashParties26)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails20) AddSettlementAmount() *AmountAndDirection45 {
	s.SettlementAmount = new(AmountAndDirection45)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails20) AddOtherAmounts() *OtherAmounts28 {
	s.OtherAmounts = new(OtherAmounts28)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails20) AddOtherBusinessParties() *OtherParties27 {
	s.OtherBusinessParties = new(OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails20) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails20) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails21 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters13 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages38 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails50 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount44 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails94 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties39 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties39 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties26 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection45 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts28 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters4 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails21) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters13 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters13)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails21) AddLinkages() *Linkages38 {
	newValue := new(Linkages38)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails21) AddTradeDetails() *SecuritiesTradeDetails50 {
	s.TradeDetails = new(SecuritiesTradeDetails50)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails21) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails21) AddQuantityAndAccountDetails() *QuantityAndAccount44 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount44)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails21) AddSettlementParameters() *SettlementDetails94 {
	s.SettlementParameters = new(SettlementDetails94)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails21) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails21) AddDeliveringSettlementParties() *SettlementParties39 {
	s.DeliveringSettlementParties = new(SettlementParties39)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails21) AddReceivingSettlementParties() *SettlementParties39 {
	s.ReceivingSettlementParties = new(SettlementParties39)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails21) AddCashParties() *CashParties26 {
	s.CashParties = new(CashParties26)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails21) AddSettlementAmount() *AmountAndDirection45 {
	s.SettlementAmount = new(AmountAndDirection45)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails21) AddOtherAmounts() *OtherAmounts28 {
	s.OtherAmounts = new(OtherAmounts28)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails21) AddOtherBusinessParties() *OtherParties27 {
	s.OtherBusinessParties = new(OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails21) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails21) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails22 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters14 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages38 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails52 `xml:"TradDtls,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount43 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails94 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties26 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection45 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts28 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters4 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails22) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters14 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters14)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails22) AddLinkages() *Linkages38 {
	newValue := new(Linkages38)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails22) AddTradeDetails() *SecuritiesTradeDetails52 {
	s.TradeDetails = new(SecuritiesTradeDetails52)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails22) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionDetails22) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails22) AddQuantityAndAccountDetails() *QuantityAndAccount43 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount43)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails22) AddSettlementParameters() *SettlementDetails94 {
	s.SettlementParameters = new(SettlementDetails94)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails22) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails22) AddDeliveringSettlementParties() *SettlementParties36 {
	s.DeliveringSettlementParties = new(SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails22) AddReceivingSettlementParties() *SettlementParties36 {
	s.ReceivingSettlementParties = new(SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails22) AddCashParties() *CashParties26 {
	s.CashParties = new(CashParties26)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails22) AddSettlementAmount() *AmountAndDirection45 {
	s.SettlementAmount = new(AmountAndDirection45)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails22) AddOtherAmounts() *OtherAmounts28 {
	s.OtherAmounts = new(OtherAmounts28)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails22) AddOtherBusinessParties() *OtherParties27 {
	s.OtherBusinessParties = new(OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails22) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails22) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails23 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters17 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages48 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails65 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount54 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails113 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties58 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties58 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties30 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection85 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts35 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters5 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails23) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters17 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters17)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails23) AddLinkages() *Linkages48 {
	newValue := new(Linkages48)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails23) AddTradeDetails() *SecuritiesTradeDetails65 {
	s.TradeDetails = new(SecuritiesTradeDetails65)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails23) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails23) AddQuantityAndAccountDetails() *QuantityAndAccount54 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount54)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails23) AddSettlementParameters() *SettlementDetails113 {
	s.SettlementParameters = new(SettlementDetails113)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails23) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails23) AddDeliveringSettlementParties() *SettlementParties58 {
	s.DeliveringSettlementParties = new(SettlementParties58)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails23) AddReceivingSettlementParties() *SettlementParties58 {
	s.ReceivingSettlementParties = new(SettlementParties58)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails23) AddCashParties() *CashParties30 {
	s.CashParties = new(CashParties30)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails23) AddSettlementAmount() *AmountAndDirection85 {
	s.SettlementAmount = new(AmountAndDirection85)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails23) AddOtherAmounts() *OtherAmounts35 {
	s.OtherAmounts = new(OtherAmounts35)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails23) AddOtherBusinessParties() *OtherParties29 {
	s.OtherBusinessParties = new(OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails23) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters5 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters5)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails23) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails24 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters17 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages49 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails65 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount54 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails113 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties58 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties58 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties30 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection85 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts35 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters5 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails24) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters17 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters17)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails24) AddLinkages() *Linkages49 {
	newValue := new(Linkages49)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails24) AddTradeDetails() *SecuritiesTradeDetails65 {
	s.TradeDetails = new(SecuritiesTradeDetails65)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails24) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails24) AddQuantityAndAccountDetails() *QuantityAndAccount54 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount54)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails24) AddSettlementParameters() *SettlementDetails113 {
	s.SettlementParameters = new(SettlementDetails113)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails24) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails24) AddDeliveringSettlementParties() *SettlementParties58 {
	s.DeliveringSettlementParties = new(SettlementParties58)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails24) AddReceivingSettlementParties() *SettlementParties58 {
	s.ReceivingSettlementParties = new(SettlementParties58)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails24) AddCashParties() *CashParties30 {
	s.CashParties = new(CashParties30)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails24) AddSettlementAmount() *AmountAndDirection85 {
	s.SettlementAmount = new(AmountAndDirection85)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails24) AddOtherAmounts() *OtherAmounts35 {
	s.OtherAmounts = new(OtherAmounts35)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails24) AddOtherBusinessParties() *OtherParties29 {
	s.OtherBusinessParties = new(OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails24) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters5 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters5)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails24) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails25 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters18 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages48 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails66 `xml:"TradDtls,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount55 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails113 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties30 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection85 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts35 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters5 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails25) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters18 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters18)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails25) AddLinkages() *Linkages48 {
	newValue := new(Linkages48)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails25) AddTradeDetails() *SecuritiesTradeDetails66 {
	s.TradeDetails = new(SecuritiesTradeDetails66)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails25) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionDetails25) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails25) AddQuantityAndAccountDetails() *QuantityAndAccount55 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount55)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails25) AddSettlementParameters() *SettlementDetails113 {
	s.SettlementParameters = new(SettlementDetails113)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails25) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails25) AddDeliveringSettlementParties() *SettlementParties44 {
	s.DeliveringSettlementParties = new(SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails25) AddReceivingSettlementParties() *SettlementParties44 {
	s.ReceivingSettlementParties = new(SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails25) AddCashParties() *CashParties30 {
	s.CashParties = new(CashParties30)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails25) AddSettlementAmount() *AmountAndDirection85 {
	s.SettlementAmount = new(AmountAndDirection85)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails25) AddOtherAmounts() *OtherAmounts35 {
	s.OtherAmounts = new(OtherAmounts35)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails25) AddOtherBusinessParties() *OtherParties29 {
	s.OtherBusinessParties = new(OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails25) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters5 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters5)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails25) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails26 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters13 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages38 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails50 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount44 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails122 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties39 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties39 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties26 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection45 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts28 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters4 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails26) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters13 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters13)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails26) AddLinkages() *Linkages38 {
	newValue := new(Linkages38)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails26) AddTradeDetails() *SecuritiesTradeDetails50 {
	s.TradeDetails = new(SecuritiesTradeDetails50)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails26) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails26) AddQuantityAndAccountDetails() *QuantityAndAccount44 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount44)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails26) AddSettlementParameters() *SettlementDetails122 {
	s.SettlementParameters = new(SettlementDetails122)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails26) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails26) AddDeliveringSettlementParties() *SettlementParties39 {
	s.DeliveringSettlementParties = new(SettlementParties39)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails26) AddReceivingSettlementParties() *SettlementParties39 {
	s.ReceivingSettlementParties = new(SettlementParties39)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails26) AddCashParties() *CashParties26 {
	s.CashParties = new(CashParties26)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails26) AddSettlementAmount() *AmountAndDirection45 {
	s.SettlementAmount = new(AmountAndDirection45)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails26) AddOtherAmounts() *OtherAmounts28 {
	s.OtherAmounts = new(OtherAmounts28)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails26) AddOtherBusinessParties() *OtherParties27 {
	s.OtherBusinessParties = new(OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails26) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails26) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails27 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters13 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages40 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails50 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount44 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails122 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties39 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties39 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties26 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection45 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts28 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters4 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails27) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters13 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters13)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails27) AddLinkages() *Linkages40 {
	newValue := new(Linkages40)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails27) AddTradeDetails() *SecuritiesTradeDetails50 {
	s.TradeDetails = new(SecuritiesTradeDetails50)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails27) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails27) AddQuantityAndAccountDetails() *QuantityAndAccount44 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount44)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails27) AddSettlementParameters() *SettlementDetails122 {
	s.SettlementParameters = new(SettlementDetails122)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails27) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails27) AddDeliveringSettlementParties() *SettlementParties39 {
	s.DeliveringSettlementParties = new(SettlementParties39)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails27) AddReceivingSettlementParties() *SettlementParties39 {
	s.ReceivingSettlementParties = new(SettlementParties39)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails27) AddCashParties() *CashParties26 {
	s.CashParties = new(CashParties26)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails27) AddSettlementAmount() *AmountAndDirection45 {
	s.SettlementAmount = new(AmountAndDirection45)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails27) AddOtherAmounts() *OtherAmounts28 {
	s.OtherAmounts = new(OtherAmounts28)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails27) AddOtherBusinessParties() *OtherParties27 {
	s.OtherBusinessParties = new(OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails27) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails27) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails28 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters14 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages38 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails52 `xml:"TradDtls,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount43 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails122 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties26 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection45 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts28 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters4 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails28) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters14 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters14)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails28) AddLinkages() *Linkages38 {
	newValue := new(Linkages38)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails28) AddTradeDetails() *SecuritiesTradeDetails52 {
	s.TradeDetails = new(SecuritiesTradeDetails52)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails28) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionDetails28) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails28) AddQuantityAndAccountDetails() *QuantityAndAccount43 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount43)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails28) AddSettlementParameters() *SettlementDetails122 {
	s.SettlementParameters = new(SettlementDetails122)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails28) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails28) AddDeliveringSettlementParties() *SettlementParties36 {
	s.DeliveringSettlementParties = new(SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails28) AddReceivingSettlementParties() *SettlementParties36 {
	s.ReceivingSettlementParties = new(SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails28) AddCashParties() *CashParties26 {
	s.CashParties = new(CashParties26)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails28) AddSettlementAmount() *AmountAndDirection45 {
	s.SettlementAmount = new(AmountAndDirection45)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails28) AddOtherAmounts() *OtherAmounts28 {
	s.OtherAmounts = new(OtherAmounts28)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails28) AddOtherBusinessParties() *OtherParties27 {
	s.OtherBusinessParties = new(OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails28) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails28) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails29 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters17 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages48 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails65 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount59 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails138 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties58 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties58 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties30 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection85 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts35 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters5 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails29) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters17 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters17)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails29) AddLinkages() *Linkages48 {
	newValue := new(Linkages48)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails29) AddTradeDetails() *SecuritiesTradeDetails65 {
	s.TradeDetails = new(SecuritiesTradeDetails65)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails29) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails29) AddQuantityAndAccountDetails() *QuantityAndAccount59 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount59)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails29) AddSettlementParameters() *SettlementDetails138 {
	s.SettlementParameters = new(SettlementDetails138)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails29) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails29) AddDeliveringSettlementParties() *SettlementParties58 {
	s.DeliveringSettlementParties = new(SettlementParties58)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails29) AddReceivingSettlementParties() *SettlementParties58 {
	s.ReceivingSettlementParties = new(SettlementParties58)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails29) AddCashParties() *CashParties30 {
	s.CashParties = new(CashParties30)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails29) AddSettlementAmount() *AmountAndDirection85 {
	s.SettlementAmount = new(AmountAndDirection85)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails29) AddOtherAmounts() *OtherAmounts35 {
	s.OtherAmounts = new(OtherAmounts35)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails29) AddOtherBusinessParties() *OtherParties29 {
	s.OtherBusinessParties = new(OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails29) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters5 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters5)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails29) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails3 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters6 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages1 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails25 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount29 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails50 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties26 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties26 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties8 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection32 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails3) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters6 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters6)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails3) AddLinkages() *Linkages1 {
	newValue := new(Linkages1)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails3) AddTradeDetails() *SecuritiesTradeDetails25 {
	s.TradeDetails = new(SecuritiesTradeDetails25)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails3) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails3) AddQuantityAndAccountDetails() *QuantityAndAccount29 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount29)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails3) AddSettlementParameters() *SettlementDetails50 {
	s.SettlementParameters = new(SettlementDetails50)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails3) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails3) AddDeliveringSettlementParties() *SettlementParties26 {
	s.DeliveringSettlementParties = new(SettlementParties26)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails3) AddReceivingSettlementParties() *SettlementParties26 {
	s.ReceivingSettlementParties = new(SettlementParties26)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails3) AddCashParties() *CashParties8 {
	s.CashParties = new(CashParties8)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails3) AddSettlementAmount() *AmountAndDirection32 {
	s.SettlementAmount = new(AmountAndDirection32)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails3) AddOtherAmounts() *OtherAmounts14 {
	s.OtherAmounts = new(OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails3) AddOtherBusinessParties() *OtherParties19 {
	s.OtherBusinessParties = new(OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails3) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails3) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails30 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters17 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages49 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails65 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount59 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails138 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties58 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties58 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties30 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection85 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts35 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters5 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails30) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters17 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters17)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails30) AddLinkages() *Linkages49 {
	newValue := new(Linkages49)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails30) AddTradeDetails() *SecuritiesTradeDetails65 {
	s.TradeDetails = new(SecuritiesTradeDetails65)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails30) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails30) AddQuantityAndAccountDetails() *QuantityAndAccount59 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount59)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails30) AddSettlementParameters() *SettlementDetails138 {
	s.SettlementParameters = new(SettlementDetails138)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails30) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails30) AddDeliveringSettlementParties() *SettlementParties58 {
	s.DeliveringSettlementParties = new(SettlementParties58)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails30) AddReceivingSettlementParties() *SettlementParties58 {
	s.ReceivingSettlementParties = new(SettlementParties58)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails30) AddCashParties() *CashParties30 {
	s.CashParties = new(CashParties30)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails30) AddSettlementAmount() *AmountAndDirection85 {
	s.SettlementAmount = new(AmountAndDirection85)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails30) AddOtherAmounts() *OtherAmounts35 {
	s.OtherAmounts = new(OtherAmounts35)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails30) AddOtherBusinessParties() *OtherParties29 {
	s.OtherBusinessParties = new(OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails30) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters5 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters5)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails30) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails31 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters18 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages48 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails66 `xml:"TradDtls,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount60 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails138 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties30 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection85 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts35 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters5 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails31) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters18 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters18)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails31) AddLinkages() *Linkages48 {
	newValue := new(Linkages48)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails31) AddTradeDetails() *SecuritiesTradeDetails66 {
	s.TradeDetails = new(SecuritiesTradeDetails66)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails31) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionDetails31) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails31) AddQuantityAndAccountDetails() *QuantityAndAccount60 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount60)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails31) AddSettlementParameters() *SettlementDetails138 {
	s.SettlementParameters = new(SettlementDetails138)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails31) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails31) AddDeliveringSettlementParties() *SettlementParties44 {
	s.DeliveringSettlementParties = new(SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails31) AddReceivingSettlementParties() *SettlementParties44 {
	s.ReceivingSettlementParties = new(SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails31) AddCashParties() *CashParties30 {
	s.CashParties = new(CashParties30)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails31) AddSettlementAmount() *AmountAndDirection85 {
	s.SettlementAmount = new(AmountAndDirection85)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails31) AddOtherAmounts() *OtherAmounts35 {
	s.OtherAmounts = new(OtherAmounts35)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails31) AddOtherBusinessParties() *OtherParties29 {
	s.OtherBusinessParties = new(OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails31) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters5 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters5)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails31) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails4 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters6 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages21 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails25 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount29 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails50 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties26 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties26 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties8 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection32 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails4) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters6 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters6)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails4) AddLinkages() *Linkages21 {
	newValue := new(Linkages21)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails4) AddTradeDetails() *SecuritiesTradeDetails25 {
	s.TradeDetails = new(SecuritiesTradeDetails25)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails4) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails4) AddQuantityAndAccountDetails() *QuantityAndAccount29 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount29)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails4) AddSettlementParameters() *SettlementDetails50 {
	s.SettlementParameters = new(SettlementDetails50)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails4) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails4) AddDeliveringSettlementParties() *SettlementParties26 {
	s.DeliveringSettlementParties = new(SettlementParties26)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails4) AddReceivingSettlementParties() *SettlementParties26 {
	s.ReceivingSettlementParties = new(SettlementParties26)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails4) AddCashParties() *CashParties8 {
	s.CashParties = new(CashParties8)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails4) AddSettlementAmount() *AmountAndDirection32 {
	s.SettlementAmount = new(AmountAndDirection32)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails4) AddOtherAmounts() *OtherAmounts14 {
	s.OtherAmounts = new(OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails4) AddOtherBusinessParties() *OtherParties19 {
	s.OtherBusinessParties = new(OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails4) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails8 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters6 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages1 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails25 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount29 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails50 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties26 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties26 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties17 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection32 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails8) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters6 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters6)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails8) AddLinkages() *Linkages1 {
	newValue := new(Linkages1)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails8) AddTradeDetails() *SecuritiesTradeDetails25 {
	s.TradeDetails = new(SecuritiesTradeDetails25)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails8) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails8) AddQuantityAndAccountDetails() *QuantityAndAccount29 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount29)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails8) AddSettlementParameters() *SettlementDetails50 {
	s.SettlementParameters = new(SettlementDetails50)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails8) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails8) AddDeliveringSettlementParties() *SettlementParties26 {
	s.DeliveringSettlementParties = new(SettlementParties26)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails8) AddReceivingSettlementParties() *SettlementParties26 {
	s.ReceivingSettlementParties = new(SettlementParties26)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails8) AddCashParties() *CashParties17 {
	s.CashParties = new(CashParties17)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails8) AddSettlementAmount() *AmountAndDirection32 {
	s.SettlementAmount = new(AmountAndDirection32)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails8) AddOtherAmounts() *OtherAmounts14 {
	s.OtherAmounts = new(OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails8) AddOtherBusinessParties() *OtherParties19 {
	s.OtherBusinessParties = new(OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails8) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails8) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails9 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters6 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages21 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails25 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount29 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails50 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties26 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties26 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties17 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection32 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails9) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters6 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters6)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails9) AddLinkages() *Linkages21 {
	newValue := new(Linkages21)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails9) AddTradeDetails() *SecuritiesTradeDetails25 {
	s.TradeDetails = new(SecuritiesTradeDetails25)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails9) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails9) AddQuantityAndAccountDetails() *QuantityAndAccount29 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount29)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails9) AddSettlementParameters() *SettlementDetails50 {
	s.SettlementParameters = new(SettlementDetails50)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails9) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails9) AddDeliveringSettlementParties() *SettlementParties26 {
	s.DeliveringSettlementParties = new(SettlementParties26)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails9) AddReceivingSettlementParties() *SettlementParties26 {
	s.ReceivingSettlementParties = new(SettlementParties26)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails9) AddCashParties() *CashParties17 {
	s.CashParties = new(CashParties17)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails9) AddSettlementAmount() *AmountAndDirection32 {
	s.SettlementAmount = new(AmountAndDirection32)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails9) AddOtherAmounts() *OtherAmounts14 {
	s.OtherAmounts = new(OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails9) AddOtherBusinessParties() *OtherParties19 {
	s.OtherBusinessParties = new(OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails9) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails9) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
