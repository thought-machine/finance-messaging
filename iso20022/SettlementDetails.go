package iso20022

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails1 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType1Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition1Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn1Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType1Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction1Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType1Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking1Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed1Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails1) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails1) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails1) AddSecuritiesTransactionType() *SecuritiesTransactionType1Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType1Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails1) AddSettlementTransactionCondition() *SettlementTransactionCondition1Choice {
	newValue := new(SettlementTransactionCondition1Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails1) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails1) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails1) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails1) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails1) AddDeliveryReturnReason() *DeliveryReturn1Choice {
	s.DeliveryReturnReason = new(DeliveryReturn1Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails1) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails1) AddExposureType() *ExposureType1Choice {
	s.ExposureType = new(ExposureType1Choice)
	return s.ExposureType
}

func (s *SettlementDetails1) AddFXStandingInstruction() *FXStandingInstruction1Choice {
	s.FXStandingInstruction = new(FXStandingInstruction1Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails1) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails1) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails1) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails1) AddRepurchaseType() *RepurchaseType1Choice {
	s.RepurchaseType = new(RepurchaseType1Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails1) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails1) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails1) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails1) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails1) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails1) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails1) AddTracking() *Tracking1Choice {
	s.Tracking = new(Tracking1Choice)
	return s.Tracking
}

func (s *SettlementDetails1) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails1) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails1) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails1) AddModificationCancellationAllowed() *ModificationCancellationAllowed1Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed1Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails1) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails10 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition1Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails10) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails10) AddSettlementTransactionCondition() *SettlementTransactionCondition1Choice {
	newValue := new(SettlementTransactionCondition1Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails10) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails10) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails10) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails10) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails10) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails10) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails10) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails10) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails10) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails10) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails10) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails10) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails10) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails10) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails100 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition19Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`
}

func (s *SettlementDetails100) AddSettlementTransactionCondition() *SettlementTransactionCondition19Choice {
	newValue := new(SettlementTransactionCondition19Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails100) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails100) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails100) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails100) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails100) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails100) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

// Details of settlement of a transaction.
type SettlementDetails101 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator6 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType21Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition20Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn3Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction4Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType14Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking4Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing6Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed4Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification30 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification30 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails101) AddHoldIndicator() *HoldIndicator6 {
	s.HoldIndicator = new(HoldIndicator6)
	return s.HoldIndicator
}

func (s *SettlementDetails101) AddPriority() *PriorityNumeric4Choice {
	s.Priority = new(PriorityNumeric4Choice)
	return s.Priority
}

func (s *SettlementDetails101) AddSecuritiesTransactionType() *SecuritiesTransactionType21Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType21Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails101) AddSettlementTransactionCondition() *SettlementTransactionCondition20Choice {
	newValue := new(SettlementTransactionCondition20Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails101) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails101) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails101) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails101) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails101) AddDeliveryReturnReason() *DeliveryReturn3Choice {
	s.DeliveryReturnReason = new(DeliveryReturn3Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails101) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails101) AddExposureType() *ExposureType16Choice {
	s.ExposureType = new(ExposureType16Choice)
	return s.ExposureType
}

func (s *SettlementDetails101) AddFXStandingInstruction() *FXStandingInstruction4Choice {
	s.FXStandingInstruction = new(FXStandingInstruction4Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails101) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails101) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails101) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails101) AddRepurchaseType() *RepurchaseType14Choice {
	s.RepurchaseType = new(RepurchaseType14Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails101) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails101) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails101) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails101) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails101) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails101) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails101) AddTracking() *Tracking4Choice {
	s.Tracking = new(Tracking4Choice)
	return s.Tracking
}

func (s *SettlementDetails101) AddAutomaticBorrowing() *AutomaticBorrowing6Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing6Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails101) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails101) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails101) AddModificationCancellationAllowed() *ModificationCancellationAllowed4Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed4Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails101) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails101) AddSecuritiesSubBalanceType() *GenericIdentification30 {
	s.SecuritiesSubBalanceType = new(GenericIdentification30)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails101) AddCashSubBalanceType() *GenericIdentification30 {
	s.CashSubBalanceType = new(GenericIdentification30)
	return s.CashSubBalanceType
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails102 struct {

	// Indicates the date as known by the two parties to be used for matching purposes when settlement of securities occurs.
	TradeDate *ISODateTime `xml:"TradDt"`

	// Provides details on either the delivering or receiving settlement parties.
	SettlementParties *SettlementParties5Choice `xml:"SttlmPties,omitempty"`

	// Indicates the collateral ownership.
	CollateralOwnership *CollateralOwnership2 `xml:"CollOwnrsh"`
}

func (s *SettlementDetails102) SetTradeDate(value string) {
	s.TradeDate = (*ISODateTime)(&value)
}

func (s *SettlementDetails102) AddSettlementParties() *SettlementParties5Choice {
	s.SettlementParties = new(SettlementParties5Choice)
	return s.SettlementParties
}

func (s *SettlementDetails102) AddCollateralOwnership() *CollateralOwnership2 {
	s.CollateralOwnership = new(CollateralOwnership2)
	return s.CollateralOwnership
}

// Details of settlement of a transaction.
type SettlementDetails103 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition21Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`
}

func (s *SettlementDetails103) AddSettlementTransactionCondition() *SettlementTransactionCondition21Choice {
	newValue := new(SettlementTransactionCondition21Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails103) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails103) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails103) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails103) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails103) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails103) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

// Details of settlement of a transaction.
type SettlementDetails104 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition22Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing8Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails104) AddPriority() *PriorityNumeric5Choice {
	s.Priority = new(PriorityNumeric5Choice)
	return s.Priority
}

func (s *SettlementDetails104) AddSettlementTransactionCondition() *SettlementTransactionCondition22Choice {
	newValue := new(SettlementTransactionCondition22Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails104) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails104) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails104) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails104) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails104) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails104) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails104) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails104) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails104) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails104) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails104) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails104) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails104) AddAutomaticBorrowing() *AutomaticBorrowing8Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing8Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails104) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails104) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails105 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition22Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction5Choice `xml:"FxStgInstr,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking5Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing8Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails105) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails105) AddPriority() *PriorityNumeric5Choice {
	s.Priority = new(PriorityNumeric5Choice)
	return s.Priority
}

func (s *SettlementDetails105) AddSettlementTransactionCondition() *SettlementTransactionCondition22Choice {
	newValue := new(SettlementTransactionCondition22Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails105) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails105) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails105) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails105) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails105) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails105) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails105) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails105) AddFXStandingInstruction() *FXStandingInstruction5Choice {
	s.FXStandingInstruction = new(FXStandingInstruction5Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails105) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails105) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails105) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails105) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails105) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails105) AddTracking() *Tracking5Choice {
	s.Tracking = new(Tracking5Choice)
	return s.Tracking
}

func (s *SettlementDetails105) AddAutomaticBorrowing() *AutomaticBorrowing8Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing8Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails105) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails105) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails106 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition22Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction5Choice `xml:"FxStgInstr,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking5Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing8Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails106) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails106) AddPriority() *PriorityNumeric5Choice {
	s.Priority = new(PriorityNumeric5Choice)
	return s.Priority
}

func (s *SettlementDetails106) AddSettlementTransactionCondition() *SettlementTransactionCondition22Choice {
	newValue := new(SettlementTransactionCondition22Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails106) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails106) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails106) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails106) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails106) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails106) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails106) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails106) AddFXStandingInstruction() *FXStandingInstruction5Choice {
	s.FXStandingInstruction = new(FXStandingInstruction5Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails106) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails106) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails106) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails106) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails106) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails106) AddTracking() *Tracking5Choice {
	s.Tracking = new(Tracking5Choice)
	return s.Tracking
}

func (s *SettlementDetails106) AddAutomaticBorrowing() *AutomaticBorrowing8Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing8Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails106) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails106) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails107 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition22Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails107) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails107) AddSettlementTransactionCondition() *SettlementTransactionCondition22Choice {
	newValue := new(SettlementTransactionCondition22Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails107) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails107) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails107) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails107) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails107) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails107) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails107) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails107) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails107) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails107) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails107) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails107) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails107) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails107) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails108 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType22Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition26Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType17Choice `xml:"RpTp,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`
}

func (s *SettlementDetails108) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails108) AddSecuritiesTransactionType() *SecuritiesTransactionType22Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType22Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails108) AddSettlementTransactionCondition() *SettlementTransactionCondition26Choice {
	newValue := new(SettlementTransactionCondition26Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails108) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails108) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails108) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails108) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails108) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails108) AddRepurchaseType() *RepurchaseType17Choice {
	s.RepurchaseType = new(RepurchaseType17Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails108) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails108) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

// Details of settlement of a transaction.
type SettlementDetails110 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType24Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade having caused the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration was to occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType19Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing8Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement was executed using a letter of guarantee or if the physical certificates were used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification47 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification47 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails110) AddPriority() *PriorityNumeric5Choice {
	s.Priority = new(PriorityNumeric5Choice)
	return s.Priority
}

func (s *SettlementDetails110) AddSecuritiesTransactionType() *SecuritiesTransactionType24Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType24Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails110) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails110) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails110) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails110) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails110) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails110) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails110) AddExposureType() *ExposureType17Choice {
	s.ExposureType = new(ExposureType17Choice)
	return s.ExposureType
}

func (s *SettlementDetails110) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails110) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails110) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails110) AddRepurchaseType() *RepurchaseType19Choice {
	s.RepurchaseType = new(RepurchaseType19Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails110) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails110) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails110) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails110) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails110) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails110) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails110) AddAutomaticBorrowing() *AutomaticBorrowing8Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing8Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails110) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails110) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails110) AddSecuritiesSubBalanceType() *GenericIdentification47 {
	s.SecuritiesSubBalanceType = new(GenericIdentification47)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails110) AddCashSubBalanceType() *GenericIdentification47 {
	s.CashSubBalanceType = new(GenericIdentification47)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails111 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator7 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType25Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition29Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn4Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction5Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType20Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking5Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing8Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed5Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification47 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification47 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails111) AddHoldIndicator() *HoldIndicator7 {
	s.HoldIndicator = new(HoldIndicator7)
	return s.HoldIndicator
}

func (s *SettlementDetails111) AddPriority() *PriorityNumeric5Choice {
	s.Priority = new(PriorityNumeric5Choice)
	return s.Priority
}

func (s *SettlementDetails111) AddSecuritiesTransactionType() *SecuritiesTransactionType25Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType25Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails111) AddSettlementTransactionCondition() *SettlementTransactionCondition29Choice {
	newValue := new(SettlementTransactionCondition29Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails111) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails111) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails111) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails111) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails111) AddDeliveryReturnReason() *DeliveryReturn4Choice {
	s.DeliveryReturnReason = new(DeliveryReturn4Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails111) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails111) AddExposureType() *ExposureType17Choice {
	s.ExposureType = new(ExposureType17Choice)
	return s.ExposureType
}

func (s *SettlementDetails111) AddFXStandingInstruction() *FXStandingInstruction5Choice {
	s.FXStandingInstruction = new(FXStandingInstruction5Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails111) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails111) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails111) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails111) AddRepurchaseType() *RepurchaseType20Choice {
	s.RepurchaseType = new(RepurchaseType20Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails111) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails111) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails111) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails111) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails111) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails111) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails111) AddTracking() *Tracking5Choice {
	s.Tracking = new(Tracking5Choice)
	return s.Tracking
}

func (s *SettlementDetails111) AddAutomaticBorrowing() *AutomaticBorrowing8Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing8Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails111) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails111) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails111) AddModificationCancellationAllowed() *ModificationCancellationAllowed5Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed5Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails111) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails111) AddSecuritiesSubBalanceType() *GenericIdentification47 {
	s.SecuritiesSubBalanceType = new(GenericIdentification47)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails111) AddCashSubBalanceType() *GenericIdentification47 {
	s.CashSubBalanceType = new(GenericIdentification47)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails112 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator7 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType25Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn4Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction5Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType20Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking5Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing8Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed5Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification47 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification47 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails112) AddHoldIndicator() *HoldIndicator7 {
	s.HoldIndicator = new(HoldIndicator7)
	return s.HoldIndicator
}

func (s *SettlementDetails112) AddPriority() *PriorityNumeric5Choice {
	s.Priority = new(PriorityNumeric5Choice)
	return s.Priority
}

func (s *SettlementDetails112) AddSecuritiesTransactionType() *SecuritiesTransactionType25Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType25Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails112) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails112) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails112) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails112) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails112) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails112) AddDeliveryReturnReason() *DeliveryReturn4Choice {
	s.DeliveryReturnReason = new(DeliveryReturn4Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails112) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails112) AddExposureType() *ExposureType17Choice {
	s.ExposureType = new(ExposureType17Choice)
	return s.ExposureType
}

func (s *SettlementDetails112) AddFXStandingInstruction() *FXStandingInstruction5Choice {
	s.FXStandingInstruction = new(FXStandingInstruction5Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails112) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails112) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails112) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails112) AddRepurchaseType() *RepurchaseType20Choice {
	s.RepurchaseType = new(RepurchaseType20Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails112) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails112) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails112) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails112) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails112) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails112) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails112) AddTracking() *Tracking5Choice {
	s.Tracking = new(Tracking5Choice)
	return s.Tracking
}

func (s *SettlementDetails112) AddAutomaticBorrowing() *AutomaticBorrowing8Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing8Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails112) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails112) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails112) AddModificationCancellationAllowed() *ModificationCancellationAllowed5Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed5Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails112) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails112) AddSecuritiesSubBalanceType() *GenericIdentification47 {
	s.SecuritiesSubBalanceType = new(GenericIdentification47)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails112) AddCashSubBalanceType() *GenericIdentification47 {
	s.CashSubBalanceType = new(GenericIdentification47)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails113 struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType25Choice `xml:"SctiesTxTp,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn4Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction5Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking5Choice `xml:"Trckg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed5Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification47 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification47 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails113) AddSecuritiesTransactionType() *SecuritiesTransactionType25Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType25Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails113) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails113) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails113) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails113) AddDeliveryReturnReason() *DeliveryReturn4Choice {
	s.DeliveryReturnReason = new(DeliveryReturn4Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails113) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails113) AddExposureType() *ExposureType17Choice {
	s.ExposureType = new(ExposureType17Choice)
	return s.ExposureType
}

func (s *SettlementDetails113) AddFXStandingInstruction() *FXStandingInstruction5Choice {
	s.FXStandingInstruction = new(FXStandingInstruction5Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails113) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails113) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails113) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails113) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails113) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails113) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails113) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails113) AddTracking() *Tracking5Choice {
	s.Tracking = new(Tracking5Choice)
	return s.Tracking
}

func (s *SettlementDetails113) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails113) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails113) AddModificationCancellationAllowed() *ModificationCancellationAllowed5Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed5Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails113) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails113) AddSecuritiesSubBalanceType() *GenericIdentification47 {
	s.SecuritiesSubBalanceType = new(GenericIdentification47)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails113) AddCashSubBalanceType() *GenericIdentification47 {
	s.CashSubBalanceType = new(GenericIdentification47)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails115 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator7 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType25Choice `xml:"SctiesTxTp"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType19Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails115) AddHoldIndicator() *HoldIndicator7 {
	s.HoldIndicator = new(HoldIndicator7)
	return s.HoldIndicator
}

func (s *SettlementDetails115) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails115) AddSecuritiesTransactionType() *SecuritiesTransactionType25Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType25Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails115) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails115) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails115) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails115) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails115) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails115) AddExposureType() *ExposureType17Choice {
	s.ExposureType = new(ExposureType17Choice)
	return s.ExposureType
}

func (s *SettlementDetails115) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails115) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails115) AddRepurchaseType() *RepurchaseType19Choice {
	s.RepurchaseType = new(RepurchaseType19Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails115) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails115) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails115) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails115) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails115) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails115) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails115) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails115) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails115) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails116 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator7 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType19Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails116) AddHoldIndicator() *HoldIndicator7 {
	s.HoldIndicator = new(HoldIndicator7)
	return s.HoldIndicator
}

func (s *SettlementDetails116) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails116) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails116) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails116) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails116) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails116) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails116) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails116) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails116) AddRepurchaseType() *RepurchaseType19Choice {
	s.RepurchaseType = new(RepurchaseType19Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails116) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails116) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails116) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails116) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails116) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails116) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails116) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails116) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails117 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType19Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails117) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails117) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails117) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails117) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails117) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails117) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails117) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails117) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails117) AddRepurchaseType() *RepurchaseType19Choice {
	s.RepurchaseType = new(RepurchaseType19Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails117) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails117) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails117) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails117) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails117) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails117) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails117) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails117) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails118 struct {

	// Indicates the date as known by the two parties to be used for matching purposes when settlement of securities occurs.
	TradeDate *ISODateTime `xml:"TradDt"`

	// Provides details on either the delivering or receiving settlement parties.
	SettlementParties *SettlementParties7Choice `xml:"SttlmPties,omitempty"`

	// Indicates the collateral ownership.
	CollateralOwnership *CollateralOwnership2 `xml:"CollOwnrsh"`
}

func (s *SettlementDetails118) SetTradeDate(value string) {
	s.TradeDate = (*ISODateTime)(&value)
}

func (s *SettlementDetails118) AddSettlementParties() *SettlementParties7Choice {
	s.SettlementParties = new(SettlementParties7Choice)
	return s.SettlementParties
}

func (s *SettlementDetails118) AddCollateralOwnership() *CollateralOwnership2 {
	s.CollateralOwnership = new(CollateralOwnership2)
	return s.CollateralOwnership
}

// Details of settlement of a transaction.
type SettlementDetails119 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator6 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType32Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn3Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction4Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType23Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking4Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing6Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed4Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification30 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification30 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails119) AddHoldIndicator() *HoldIndicator6 {
	s.HoldIndicator = new(HoldIndicator6)
	return s.HoldIndicator
}

func (s *SettlementDetails119) AddPriority() *PriorityNumeric4Choice {
	s.Priority = new(PriorityNumeric4Choice)
	return s.Priority
}

func (s *SettlementDetails119) AddSecuritiesTransactionType() *SecuritiesTransactionType32Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType32Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails119) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails119) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails119) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails119) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails119) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails119) AddDeliveryReturnReason() *DeliveryReturn3Choice {
	s.DeliveryReturnReason = new(DeliveryReturn3Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails119) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails119) AddExposureType() *ExposureType16Choice {
	s.ExposureType = new(ExposureType16Choice)
	return s.ExposureType
}

func (s *SettlementDetails119) AddFXStandingInstruction() *FXStandingInstruction4Choice {
	s.FXStandingInstruction = new(FXStandingInstruction4Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails119) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails119) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails119) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails119) AddRepurchaseType() *RepurchaseType23Choice {
	s.RepurchaseType = new(RepurchaseType23Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails119) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails119) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails119) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails119) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails119) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails119) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails119) AddTracking() *Tracking4Choice {
	s.Tracking = new(Tracking4Choice)
	return s.Tracking
}

func (s *SettlementDetails119) AddAutomaticBorrowing() *AutomaticBorrowing6Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing6Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails119) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails119) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails119) AddModificationCancellationAllowed() *ModificationCancellationAllowed4Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed4Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails119) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails119) AddSecuritiesSubBalanceType() *GenericIdentification30 {
	s.SecuritiesSubBalanceType = new(GenericIdentification30)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails119) AddCashSubBalanceType() *GenericIdentification30 {
	s.CashSubBalanceType = new(GenericIdentification30)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails120 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator6 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType32Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition20Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn3Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction4Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType23Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking4Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing6Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed4Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification30 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification30 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails120) AddHoldIndicator() *HoldIndicator6 {
	s.HoldIndicator = new(HoldIndicator6)
	return s.HoldIndicator
}

func (s *SettlementDetails120) AddPriority() *PriorityNumeric4Choice {
	s.Priority = new(PriorityNumeric4Choice)
	return s.Priority
}

func (s *SettlementDetails120) AddSecuritiesTransactionType() *SecuritiesTransactionType32Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType32Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails120) AddSettlementTransactionCondition() *SettlementTransactionCondition20Choice {
	newValue := new(SettlementTransactionCondition20Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails120) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails120) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails120) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails120) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails120) AddDeliveryReturnReason() *DeliveryReturn3Choice {
	s.DeliveryReturnReason = new(DeliveryReturn3Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails120) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails120) AddExposureType() *ExposureType16Choice {
	s.ExposureType = new(ExposureType16Choice)
	return s.ExposureType
}

func (s *SettlementDetails120) AddFXStandingInstruction() *FXStandingInstruction4Choice {
	s.FXStandingInstruction = new(FXStandingInstruction4Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails120) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails120) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails120) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails120) AddRepurchaseType() *RepurchaseType23Choice {
	s.RepurchaseType = new(RepurchaseType23Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails120) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails120) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails120) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails120) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails120) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails120) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails120) AddTracking() *Tracking4Choice {
	s.Tracking = new(Tracking4Choice)
	return s.Tracking
}

func (s *SettlementDetails120) AddAutomaticBorrowing() *AutomaticBorrowing6Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing6Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails120) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails120) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails120) AddModificationCancellationAllowed() *ModificationCancellationAllowed4Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed4Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails120) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails120) AddSecuritiesSubBalanceType() *GenericIdentification30 {
	s.SecuritiesSubBalanceType = new(GenericIdentification30)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails120) AddCashSubBalanceType() *GenericIdentification30 {
	s.CashSubBalanceType = new(GenericIdentification30)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails121 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator6 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType32Choice `xml:"SctiesTxTp"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType22Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails121) AddHoldIndicator() *HoldIndicator6 {
	s.HoldIndicator = new(HoldIndicator6)
	return s.HoldIndicator
}

func (s *SettlementDetails121) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails121) AddSecuritiesTransactionType() *SecuritiesTransactionType32Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType32Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails121) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails121) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails121) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails121) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails121) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails121) AddExposureType() *ExposureType16Choice {
	s.ExposureType = new(ExposureType16Choice)
	return s.ExposureType
}

func (s *SettlementDetails121) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails121) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails121) AddRepurchaseType() *RepurchaseType22Choice {
	s.RepurchaseType = new(RepurchaseType22Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails121) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails121) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails121) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails121) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails121) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails121) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails121) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails121) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails121) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails122 struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType32Choice `xml:"SctiesTxTp,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn3Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction4Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking4Choice `xml:"Trckg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed4Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification30 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification30 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails122) AddSecuritiesTransactionType() *SecuritiesTransactionType32Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType32Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails122) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails122) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails122) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails122) AddDeliveryReturnReason() *DeliveryReturn3Choice {
	s.DeliveryReturnReason = new(DeliveryReturn3Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails122) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails122) AddExposureType() *ExposureType16Choice {
	s.ExposureType = new(ExposureType16Choice)
	return s.ExposureType
}

func (s *SettlementDetails122) AddFXStandingInstruction() *FXStandingInstruction4Choice {
	s.FXStandingInstruction = new(FXStandingInstruction4Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails122) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails122) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails122) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails122) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails122) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails122) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails122) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails122) AddTracking() *Tracking4Choice {
	s.Tracking = new(Tracking4Choice)
	return s.Tracking
}

func (s *SettlementDetails122) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails122) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails122) AddModificationCancellationAllowed() *ModificationCancellationAllowed4Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed4Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails122) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails122) AddSecuritiesSubBalanceType() *GenericIdentification30 {
	s.SecuritiesSubBalanceType = new(GenericIdentification30)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails122) AddCashSubBalanceType() *GenericIdentification30 {
	s.CashSubBalanceType = new(GenericIdentification30)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails125 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType33Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition17Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType13Choice `xml:"RpTp,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`
}

func (s *SettlementDetails125) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails125) AddSecuritiesTransactionType() *SecuritiesTransactionType33Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType33Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails125) AddSettlementTransactionCondition() *SettlementTransactionCondition17Choice {
	newValue := new(SettlementTransactionCondition17Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails125) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails125) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails125) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails125) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails125) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails125) AddRepurchaseType() *RepurchaseType13Choice {
	s.RepurchaseType = new(RepurchaseType13Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails125) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails125) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

// Details of settlement of a transaction.
type SettlementDetails126 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType22Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails126) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails126) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails126) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails126) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails126) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails126) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails126) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails126) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails126) AddRepurchaseType() *RepurchaseType22Choice {
	s.RepurchaseType = new(RepurchaseType22Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails126) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails126) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails126) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails126) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails126) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails126) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails126) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails126) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails127 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator6 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType22Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails127) AddHoldIndicator() *HoldIndicator6 {
	s.HoldIndicator = new(HoldIndicator6)
	return s.HoldIndicator
}

func (s *SettlementDetails127) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails127) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails127) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails127) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails127) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails127) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails127) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails127) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails127) AddRepurchaseType() *RepurchaseType22Choice {
	s.RepurchaseType = new(RepurchaseType22Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails127) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails127) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails127) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails127) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails127) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails127) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails127) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails127) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails128 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType18Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade having caused the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration was to occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType22Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing6Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement was executed using a letter of guarantee or if the physical certificates were used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification30 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification30 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails128) AddPriority() *PriorityNumeric4Choice {
	s.Priority = new(PriorityNumeric4Choice)
	return s.Priority
}

func (s *SettlementDetails128) AddSecuritiesTransactionType() *SecuritiesTransactionType18Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType18Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails128) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails128) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails128) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails128) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails128) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails128) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails128) AddExposureType() *ExposureType16Choice {
	s.ExposureType = new(ExposureType16Choice)
	return s.ExposureType
}

func (s *SettlementDetails128) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails128) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails128) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails128) AddRepurchaseType() *RepurchaseType22Choice {
	s.RepurchaseType = new(RepurchaseType22Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails128) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails128) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails128) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails128) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails128) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails128) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails128) AddAutomaticBorrowing() *AutomaticBorrowing6Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing6Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails128) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails128) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails128) AddSecuritiesSubBalanceType() *GenericIdentification30 {
	s.SecuritiesSubBalanceType = new(GenericIdentification30)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails128) AddCashSubBalanceType() *GenericIdentification30 {
	s.CashSubBalanceType = new(GenericIdentification30)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails130 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator7 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType24Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails130) AddHoldIndicator() *HoldIndicator7 {
	s.HoldIndicator = new(HoldIndicator7)
	return s.HoldIndicator
}

func (s *SettlementDetails130) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails130) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails130) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails130) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails130) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails130) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails130) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails130) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails130) AddRepurchaseType() *RepurchaseType24Choice {
	s.RepurchaseType = new(RepurchaseType24Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails130) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails130) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails130) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails130) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails130) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails130) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails130) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails130) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails131 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType35Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition26Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType17Choice `xml:"RpTp,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`
}

func (s *SettlementDetails131) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails131) AddSecuritiesTransactionType() *SecuritiesTransactionType35Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType35Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails131) AddSettlementTransactionCondition() *SettlementTransactionCondition26Choice {
	newValue := new(SettlementTransactionCondition26Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails131) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails131) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails131) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails131) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails131) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails131) AddRepurchaseType() *RepurchaseType17Choice {
	s.RepurchaseType = new(RepurchaseType17Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails131) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails131) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

// Details of settlement of a transaction.
type SettlementDetails132 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator7 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType34Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn4Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction5Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType26Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking5Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing8Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed5Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification47 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification47 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails132) AddHoldIndicator() *HoldIndicator7 {
	s.HoldIndicator = new(HoldIndicator7)
	return s.HoldIndicator
}

func (s *SettlementDetails132) AddPriority() *PriorityNumeric5Choice {
	s.Priority = new(PriorityNumeric5Choice)
	return s.Priority
}

func (s *SettlementDetails132) AddSecuritiesTransactionType() *SecuritiesTransactionType34Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType34Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails132) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails132) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails132) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails132) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails132) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails132) AddDeliveryReturnReason() *DeliveryReturn4Choice {
	s.DeliveryReturnReason = new(DeliveryReturn4Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails132) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails132) AddExposureType() *ExposureType17Choice {
	s.ExposureType = new(ExposureType17Choice)
	return s.ExposureType
}

func (s *SettlementDetails132) AddFXStandingInstruction() *FXStandingInstruction5Choice {
	s.FXStandingInstruction = new(FXStandingInstruction5Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails132) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails132) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails132) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails132) AddRepurchaseType() *RepurchaseType26Choice {
	s.RepurchaseType = new(RepurchaseType26Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails132) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails132) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails132) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails132) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails132) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails132) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails132) AddTracking() *Tracking5Choice {
	s.Tracking = new(Tracking5Choice)
	return s.Tracking
}

func (s *SettlementDetails132) AddAutomaticBorrowing() *AutomaticBorrowing8Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing8Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails132) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails132) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails132) AddModificationCancellationAllowed() *ModificationCancellationAllowed5Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed5Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails132) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails132) AddSecuritiesSubBalanceType() *GenericIdentification47 {
	s.SecuritiesSubBalanceType = new(GenericIdentification47)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails132) AddCashSubBalanceType() *GenericIdentification47 {
	s.CashSubBalanceType = new(GenericIdentification47)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails133 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator7 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType34Choice `xml:"SctiesTxTp"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType24Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails133) AddHoldIndicator() *HoldIndicator7 {
	s.HoldIndicator = new(HoldIndicator7)
	return s.HoldIndicator
}

func (s *SettlementDetails133) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails133) AddSecuritiesTransactionType() *SecuritiesTransactionType34Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType34Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails133) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails133) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails133) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails133) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails133) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails133) AddExposureType() *ExposureType17Choice {
	s.ExposureType = new(ExposureType17Choice)
	return s.ExposureType
}

func (s *SettlementDetails133) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails133) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails133) AddRepurchaseType() *RepurchaseType24Choice {
	s.RepurchaseType = new(RepurchaseType24Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails133) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails133) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails133) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails133) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails133) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails133) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails133) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails133) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails133) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails134 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType24Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade having caused the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration was to occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType24Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing8Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement was executed using a letter of guarantee or if the physical certificates were used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification47 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification47 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails134) AddPriority() *PriorityNumeric5Choice {
	s.Priority = new(PriorityNumeric5Choice)
	return s.Priority
}

func (s *SettlementDetails134) AddSecuritiesTransactionType() *SecuritiesTransactionType24Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType24Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails134) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails134) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails134) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails134) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails134) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails134) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails134) AddExposureType() *ExposureType17Choice {
	s.ExposureType = new(ExposureType17Choice)
	return s.ExposureType
}

func (s *SettlementDetails134) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails134) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails134) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails134) AddRepurchaseType() *RepurchaseType24Choice {
	s.RepurchaseType = new(RepurchaseType24Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails134) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails134) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails134) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails134) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails134) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails134) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails134) AddAutomaticBorrowing() *AutomaticBorrowing8Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing8Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails134) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails134) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails134) AddSecuritiesSubBalanceType() *GenericIdentification47 {
	s.SecuritiesSubBalanceType = new(GenericIdentification47)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails134) AddCashSubBalanceType() *GenericIdentification47 {
	s.CashSubBalanceType = new(GenericIdentification47)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails137 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator7 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType34Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition29Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn4Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction5Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType26Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking5Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing8Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed5Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification47 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification47 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails137) AddHoldIndicator() *HoldIndicator7 {
	s.HoldIndicator = new(HoldIndicator7)
	return s.HoldIndicator
}

func (s *SettlementDetails137) AddPriority() *PriorityNumeric5Choice {
	s.Priority = new(PriorityNumeric5Choice)
	return s.Priority
}

func (s *SettlementDetails137) AddSecuritiesTransactionType() *SecuritiesTransactionType34Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType34Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails137) AddSettlementTransactionCondition() *SettlementTransactionCondition29Choice {
	newValue := new(SettlementTransactionCondition29Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails137) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails137) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails137) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails137) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails137) AddDeliveryReturnReason() *DeliveryReturn4Choice {
	s.DeliveryReturnReason = new(DeliveryReturn4Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails137) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails137) AddExposureType() *ExposureType17Choice {
	s.ExposureType = new(ExposureType17Choice)
	return s.ExposureType
}

func (s *SettlementDetails137) AddFXStandingInstruction() *FXStandingInstruction5Choice {
	s.FXStandingInstruction = new(FXStandingInstruction5Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails137) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails137) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails137) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails137) AddRepurchaseType() *RepurchaseType26Choice {
	s.RepurchaseType = new(RepurchaseType26Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails137) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails137) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails137) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails137) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails137) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails137) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails137) AddTracking() *Tracking5Choice {
	s.Tracking = new(Tracking5Choice)
	return s.Tracking
}

func (s *SettlementDetails137) AddAutomaticBorrowing() *AutomaticBorrowing8Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing8Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails137) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails137) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails137) AddModificationCancellationAllowed() *ModificationCancellationAllowed5Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed5Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails137) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails137) AddSecuritiesSubBalanceType() *GenericIdentification47 {
	s.SecuritiesSubBalanceType = new(GenericIdentification47)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails137) AddCashSubBalanceType() *GenericIdentification47 {
	s.CashSubBalanceType = new(GenericIdentification47)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails138 struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType34Choice `xml:"SctiesTxTp,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn4Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction5Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking5Choice `xml:"Trckg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed5Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification47 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification47 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails138) AddSecuritiesTransactionType() *SecuritiesTransactionType34Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType34Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails138) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails138) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails138) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails138) AddDeliveryReturnReason() *DeliveryReturn4Choice {
	s.DeliveryReturnReason = new(DeliveryReturn4Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails138) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails138) AddExposureType() *ExposureType17Choice {
	s.ExposureType = new(ExposureType17Choice)
	return s.ExposureType
}

func (s *SettlementDetails138) AddFXStandingInstruction() *FXStandingInstruction5Choice {
	s.FXStandingInstruction = new(FXStandingInstruction5Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails138) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails138) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails138) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails138) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails138) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails138) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails138) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails138) AddTracking() *Tracking5Choice {
	s.Tracking = new(Tracking5Choice)
	return s.Tracking
}

func (s *SettlementDetails138) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails138) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails138) AddModificationCancellationAllowed() *ModificationCancellationAllowed5Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed5Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails138) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails138) AddSecuritiesSubBalanceType() *GenericIdentification47 {
	s.SecuritiesSubBalanceType = new(GenericIdentification47)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails138) AddCashSubBalanceType() *GenericIdentification47 {
	s.CashSubBalanceType = new(GenericIdentification47)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails139 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType24Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails139) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails139) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails139) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails139) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails139) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails139) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails139) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails139) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails139) AddRepurchaseType() *RepurchaseType24Choice {
	s.RepurchaseType = new(RepurchaseType24Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails139) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails139) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails139) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails139) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails139) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails139) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails139) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails139) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails2 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition1Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails2) AddSettlementTransactionCondition() *SettlementTransactionCondition1Choice {
	newValue := new(SettlementTransactionCondition1Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails2) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails2) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails2) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails2) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails2) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails2) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails2) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails2) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails2) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails2) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails2) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails2) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails2) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails2) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails2) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails2) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails22 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator2 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType1Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition6Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn1Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType3Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction1Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType1Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking1Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed1Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification19 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification19 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails22) AddHoldIndicator() *HoldIndicator2 {
	s.HoldIndicator = new(HoldIndicator2)
	return s.HoldIndicator
}

func (s *SettlementDetails22) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails22) AddSecuritiesTransactionType() *SecuritiesTransactionType1Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType1Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails22) AddSettlementTransactionCondition() *SettlementTransactionCondition6Choice {
	newValue := new(SettlementTransactionCondition6Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails22) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails22) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails22) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails22) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails22) AddDeliveryReturnReason() *DeliveryReturn1Choice {
	s.DeliveryReturnReason = new(DeliveryReturn1Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails22) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails22) AddExposureType() *ExposureType3Choice {
	s.ExposureType = new(ExposureType3Choice)
	return s.ExposureType
}

func (s *SettlementDetails22) AddFXStandingInstruction() *FXStandingInstruction1Choice {
	s.FXStandingInstruction = new(FXStandingInstruction1Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails22) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails22) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails22) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails22) AddRepurchaseType() *RepurchaseType1Choice {
	s.RepurchaseType = new(RepurchaseType1Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails22) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails22) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails22) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails22) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails22) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails22) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails22) AddTracking() *Tracking1Choice {
	s.Tracking = new(Tracking1Choice)
	return s.Tracking
}

func (s *SettlementDetails22) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails22) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails22) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails22) AddModificationCancellationAllowed() *ModificationCancellationAllowed1Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed1Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails22) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails22) AddSecuritiesSubBalanceType() *GenericIdentification19 {
	s.SecuritiesSubBalanceType = new(GenericIdentification19)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails22) AddCashSubBalanceType() *GenericIdentification19 {
	s.CashSubBalanceType = new(GenericIdentification19)
	return s.CashSubBalanceType
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails23 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator2 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition6Choice `xml:"SttlmTxCond,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType1Choice `xml:"SctiesTxTp"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType4Choice `xml:"XpsrTp,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails23) AddHoldIndicator() *HoldIndicator2 {
	s.HoldIndicator = new(HoldIndicator2)
	return s.HoldIndicator
}

func (s *SettlementDetails23) AddSettlementTransactionCondition() *SettlementTransactionCondition6Choice {
	newValue := new(SettlementTransactionCondition6Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails23) AddSecuritiesTransactionType() *SecuritiesTransactionType1Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType1Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails23) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails23) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails23) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails23) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails23) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails23) AddExposureType() *ExposureType4Choice {
	s.ExposureType = new(ExposureType4Choice)
	return s.ExposureType
}

func (s *SettlementDetails23) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails23) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails23) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails23) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails23) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails23) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails23) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails23) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails23) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails23) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails23) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails23) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails24 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType3Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition6Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade having caused the collateral movement.
	ExposureType *ExposureType4Choice `xml:"XpsrTp,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration was to occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement was executed using a letter of guarantee or if the physical certificates were used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification19 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification19 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails24) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails24) AddSecuritiesTransactionType() *SecuritiesTransactionType3Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType3Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails24) AddSettlementTransactionCondition() *SettlementTransactionCondition6Choice {
	newValue := new(SettlementTransactionCondition6Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails24) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails24) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails24) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails24) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails24) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails24) AddExposureType() *ExposureType4Choice {
	s.ExposureType = new(ExposureType4Choice)
	return s.ExposureType
}

func (s *SettlementDetails24) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails24) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails24) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails24) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails24) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails24) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails24) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails24) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails24) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails24) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails24) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails24) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails24) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails24) AddSecuritiesSubBalanceType() *GenericIdentification19 {
	s.SecuritiesSubBalanceType = new(GenericIdentification19)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails24) AddCashSubBalanceType() *GenericIdentification19 {
	s.CashSubBalanceType = new(GenericIdentification19)
	return s.CashSubBalanceType
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails25 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType7Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition5Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType8Choice `xml:"RpTp,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`
}

func (s *SettlementDetails25) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails25) AddSecuritiesTransactionType() *SecuritiesTransactionType7Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType7Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails25) AddSettlementTransactionCondition() *SettlementTransactionCondition5Choice {
	newValue := new(SettlementTransactionCondition5Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails25) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails25) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails25) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails25) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails25) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails25) AddRepurchaseType() *RepurchaseType8Choice {
	s.RepurchaseType = new(RepurchaseType8Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails25) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails25) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails26 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType3Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition7Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade having caused the collateral movement.
	ExposureType *ExposureType5Choice `xml:"XpsrTp,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration was to occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement was executed using a letter of guarantee or if the physical certificates were used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification19 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification19 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails26) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails26) AddSecuritiesTransactionType() *SecuritiesTransactionType3Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType3Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails26) AddSettlementTransactionCondition() *SettlementTransactionCondition7Choice {
	newValue := new(SettlementTransactionCondition7Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails26) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails26) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails26) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails26) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails26) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails26) AddExposureType() *ExposureType5Choice {
	s.ExposureType = new(ExposureType5Choice)
	return s.ExposureType
}

func (s *SettlementDetails26) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails26) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails26) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails26) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails26) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails26) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails26) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails26) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails26) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails26) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails26) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails26) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails26) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails26) AddSecuritiesSubBalanceType() *GenericIdentification19 {
	s.SecuritiesSubBalanceType = new(GenericIdentification19)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails26) AddCashSubBalanceType() *GenericIdentification19 {
	s.CashSubBalanceType = new(GenericIdentification19)
	return s.CashSubBalanceType
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails27 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition6Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails27) AddSettlementTransactionCondition() *SettlementTransactionCondition6Choice {
	newValue := new(SettlementTransactionCondition6Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails27) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails27) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails27) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails27) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails27) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails27) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails27) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails27) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails27) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails27) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails27) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails27) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails27) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails27) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails27) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails27) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails28 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator2 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition6Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails28) AddHoldIndicator() *HoldIndicator2 {
	s.HoldIndicator = new(HoldIndicator2)
	return s.HoldIndicator
}

func (s *SettlementDetails28) AddSettlementTransactionCondition() *SettlementTransactionCondition6Choice {
	newValue := new(SettlementTransactionCondition6Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails28) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails28) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails28) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails28) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails28) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails28) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails28) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails28) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails28) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails28) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails28) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails28) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails28) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails28) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails28) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails28) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails29 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition7Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails29) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails29) AddSettlementTransactionCondition() *SettlementTransactionCondition7Choice {
	newValue := new(SettlementTransactionCondition7Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails29) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails29) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails29) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails29) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails29) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails29) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails29) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails29) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails29) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails29) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails29) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails29) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails29) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails29) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails3 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition1Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction1Choice `xml:"FxStgInstr,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking1Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails3) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails3) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails3) AddSettlementTransactionCondition() *SettlementTransactionCondition1Choice {
	newValue := new(SettlementTransactionCondition1Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails3) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails3) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails3) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails3) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails3) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails3) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails3) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails3) AddFXStandingInstruction() *FXStandingInstruction1Choice {
	s.FXStandingInstruction = new(FXStandingInstruction1Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails3) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails3) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails3) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails3) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails3) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails3) AddTracking() *Tracking1Choice {
	s.Tracking = new(Tracking1Choice)
	return s.Tracking
}

func (s *SettlementDetails3) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails3) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails3) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails30 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition7Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails30) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails30) AddSettlementTransactionCondition() *SettlementTransactionCondition7Choice {
	newValue := new(SettlementTransactionCondition7Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails30) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails30) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails30) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails30) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails30) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails30) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails30) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails30) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails30) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails30) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails30) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails30) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails30) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails30) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails30) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails31 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition7Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction1Choice `xml:"FxStgInstr,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking1Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails31) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails31) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails31) AddSettlementTransactionCondition() *SettlementTransactionCondition7Choice {
	newValue := new(SettlementTransactionCondition7Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails31) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails31) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails31) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails31) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails31) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails31) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails31) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails31) AddFXStandingInstruction() *FXStandingInstruction1Choice {
	s.FXStandingInstruction = new(FXStandingInstruction1Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails31) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails31) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails31) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails31) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails31) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails31) AddTracking() *Tracking1Choice {
	s.Tracking = new(Tracking1Choice)
	return s.Tracking
}

func (s *SettlementDetails31) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails31) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails31) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails4 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition2Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`
}

func (s *SettlementDetails4) AddSettlementTransactionCondition() *SettlementTransactionCondition2Choice {
	newValue := new(SettlementTransactionCondition2Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails4) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails4) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails4) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails4) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails4) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails4) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails42 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator4 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType9Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn1Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType3Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction1Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType1Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking1Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed1Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification20 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification20 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails42) AddHoldIndicator() *HoldIndicator4 {
	s.HoldIndicator = new(HoldIndicator4)
	return s.HoldIndicator
}

func (s *SettlementDetails42) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails42) AddSecuritiesTransactionType() *SecuritiesTransactionType9Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType9Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails42) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails42) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails42) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails42) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails42) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails42) AddDeliveryReturnReason() *DeliveryReturn1Choice {
	s.DeliveryReturnReason = new(DeliveryReturn1Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails42) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails42) AddExposureType() *ExposureType3Choice {
	s.ExposureType = new(ExposureType3Choice)
	return s.ExposureType
}

func (s *SettlementDetails42) AddFXStandingInstruction() *FXStandingInstruction1Choice {
	s.FXStandingInstruction = new(FXStandingInstruction1Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails42) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails42) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails42) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails42) AddRepurchaseType() *RepurchaseType1Choice {
	s.RepurchaseType = new(RepurchaseType1Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails42) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails42) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails42) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails42) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails42) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails42) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails42) AddTracking() *Tracking1Choice {
	s.Tracking = new(Tracking1Choice)
	return s.Tracking
}

func (s *SettlementDetails42) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails42) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails42) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails42) AddModificationCancellationAllowed() *ModificationCancellationAllowed1Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed1Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails42) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails42) AddSecuritiesSubBalanceType() *GenericIdentification20 {
	s.SecuritiesSubBalanceType = new(GenericIdentification20)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails42) AddCashSubBalanceType() *GenericIdentification20 {
	s.CashSubBalanceType = new(GenericIdentification20)
	return s.CashSubBalanceType
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails43 struct {

	// Indicates whether the trade is to be settled as principal or netted off against another trade.
	SettlementTransactionType *SettlementTransactionType1Choice `xml:"SttlmTxTp"`

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric3Choice `xml:"Prty,omitempty"`

	// Specifies if the Electronic Trade Confirmation (ETC) service provider is to generate or not a settlement instruction.
	SettlementInstructionGeneration *SettlementInstructionGeneration1Choice `xml:"SttlmInstrGnrtn,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition11Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership3Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade3Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility3Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem3Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType9Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction3Choice `xml:"FxStgInstr,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide3Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility3Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration6Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType11Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction3Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS3Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity3Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod3Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty3Choice `xml:"TaxCpcty,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification38 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking3Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing5Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee3Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed3Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails43) AddSettlementTransactionType() *SettlementTransactionType1Choice {
	s.SettlementTransactionType = new(SettlementTransactionType1Choice)
	return s.SettlementTransactionType
}

func (s *SettlementDetails43) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails43) AddPriority() *PriorityNumeric3Choice {
	s.Priority = new(PriorityNumeric3Choice)
	return s.Priority
}

func (s *SettlementDetails43) AddSettlementInstructionGeneration() *SettlementInstructionGeneration1Choice {
	s.SettlementInstructionGeneration = new(SettlementInstructionGeneration1Choice)
	return s.SettlementInstructionGeneration
}

func (s *SettlementDetails43) AddSettlementTransactionCondition() *SettlementTransactionCondition11Choice {
	newValue := new(SettlementTransactionCondition11Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails43) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails43) AddBeneficialOwnership() *BeneficialOwnership3Choice {
	s.BeneficialOwnership = new(BeneficialOwnership3Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails43) AddBlockTrade() *BlockTrade3Choice {
	s.BlockTrade = new(BlockTrade3Choice)
	return s.BlockTrade
}

func (s *SettlementDetails43) AddCCPEligibility() *CentralCounterPartyEligibility3Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility3Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails43) AddCashClearingSystem() *CashSettlementSystem3Choice {
	s.CashClearingSystem = new(CashSettlementSystem3Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails43) AddExposureType() *ExposureType9Choice {
	s.ExposureType = new(ExposureType9Choice)
	return s.ExposureType
}

func (s *SettlementDetails43) AddFXStandingInstruction() *FXStandingInstruction3Choice {
	s.FXStandingInstruction = new(FXStandingInstruction3Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails43) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SettlementDetails43) AddMarketClientSide() *MarketClientSide3Choice {
	s.MarketClientSide = new(MarketClientSide3Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails43) AddNettingEligibility() *NettingEligibility3Choice {
	s.NettingEligibility = new(NettingEligibility3Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails43) AddRegistration() *Registration6Choice {
	s.Registration = new(Registration6Choice)
	return s.Registration
}

func (s *SettlementDetails43) AddRepurchaseType() *RepurchaseType11Choice {
	s.RepurchaseType = new(RepurchaseType11Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails43) AddLegalRestrictions() *Restriction3Choice {
	s.LegalRestrictions = new(Restriction3Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails43) AddSecuritiesRTGS() *SecuritiesRTGS3Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS3Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails43) AddSettlingCapacity() *SettlingCapacity3Choice {
	s.SettlingCapacity = new(SettlingCapacity3Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails43) AddSettlementSystemMethod() *SettlementSystemMethod3Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod3Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails43) AddTaxCapacity() *TaxCapacityParty3Choice {
	s.TaxCapacity = new(TaxCapacityParty3Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails43) SetStampDutyIndicator(value string) {
	s.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails43) AddStampDutyTaxBasis() *GenericIdentification38 {
	s.StampDutyTaxBasis = new(GenericIdentification38)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails43) AddTracking() *Tracking3Choice {
	s.Tracking = new(Tracking3Choice)
	return s.Tracking
}

func (s *SettlementDetails43) AddAutomaticBorrowing() *AutomaticBorrowing5Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing5Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails43) AddLetterOfGuarantee() *LetterOfGuarantee3Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee3Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails43) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails43) AddModificationCancellationAllowed() *ModificationCancellationAllowed3Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed3Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails43) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails44 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator4 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType9Choice `xml:"SctiesTxTp"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType4Choice `xml:"XpsrTp,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails44) AddHoldIndicator() *HoldIndicator4 {
	s.HoldIndicator = new(HoldIndicator4)
	return s.HoldIndicator
}

func (s *SettlementDetails44) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails44) AddSecuritiesTransactionType() *SecuritiesTransactionType9Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType9Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails44) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails44) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails44) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails44) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails44) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails44) AddExposureType() *ExposureType4Choice {
	s.ExposureType = new(ExposureType4Choice)
	return s.ExposureType
}

func (s *SettlementDetails44) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails44) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails44) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails44) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails44) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails44) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails44) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails44) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails44) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails44) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails44) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails44) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails45 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType10Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade having caused the collateral movement.
	ExposureType *ExposureType4Choice `xml:"XpsrTp,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration was to occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement was executed using a letter of guarantee or if the physical certificates were used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification20 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification20 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails45) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails45) AddSecuritiesTransactionType() *SecuritiesTransactionType10Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType10Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails45) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails45) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails45) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails45) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails45) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails45) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails45) AddExposureType() *ExposureType4Choice {
	s.ExposureType = new(ExposureType4Choice)
	return s.ExposureType
}

func (s *SettlementDetails45) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails45) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails45) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails45) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails45) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails45) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails45) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails45) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails45) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails45) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails45) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails45) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails45) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails45) AddSecuritiesSubBalanceType() *GenericIdentification20 {
	s.SecuritiesSubBalanceType = new(GenericIdentification20)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails45) AddCashSubBalanceType() *GenericIdentification20 {
	s.CashSubBalanceType = new(GenericIdentification20)
	return s.CashSubBalanceType
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails47 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator4 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails47) AddHoldIndicator() *HoldIndicator4 {
	s.HoldIndicator = new(HoldIndicator4)
	return s.HoldIndicator
}

func (s *SettlementDetails47) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails47) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails47) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails47) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails47) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails47) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails47) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails47) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails47) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails47) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails47) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails47) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails47) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails47) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails47) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails47) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails47) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails48 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails48) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails48) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails48) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails48) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails48) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails48) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails48) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails48) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails48) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails48) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails48) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails48) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails48) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails48) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails48) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails48) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails48) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails49 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType12Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition5Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType8Choice `xml:"RpTp,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`
}

func (s *SettlementDetails49) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails49) AddSecuritiesTransactionType() *SecuritiesTransactionType12Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType12Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails49) AddSettlementTransactionCondition() *SettlementTransactionCondition5Choice {
	newValue := new(SettlementTransactionCondition5Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails49) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails49) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails49) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails49) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails49) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails49) AddRepurchaseType() *RepurchaseType8Choice {
	s.RepurchaseType = new(RepurchaseType8Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails49) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails49) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails5 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType2Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition2Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType1Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`
}

func (s *SettlementDetails5) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails5) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails5) AddSecuritiesTransactionType() *SecuritiesTransactionType2Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType2Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails5) AddSettlementTransactionCondition() *SettlementTransactionCondition2Choice {
	newValue := new(SettlementTransactionCondition2Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails5) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails5) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails5) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails5) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails5) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails5) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails5) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails5) AddRepurchaseType() *RepurchaseType1Choice {
	s.RepurchaseType = new(RepurchaseType1Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails5) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails5) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails5) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails5) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails5) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails5) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails5) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

// Details of settlement of a transaction.
type SettlementDetails50 struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType9Choice `xml:"SctiesTxTp,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn1Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType3Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction1Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking1Choice `xml:"Trckg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed1Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification20 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification20 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails50) AddSecuritiesTransactionType() *SecuritiesTransactionType9Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType9Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails50) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails50) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails50) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails50) AddDeliveryReturnReason() *DeliveryReturn1Choice {
	s.DeliveryReturnReason = new(DeliveryReturn1Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails50) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails50) AddExposureType() *ExposureType3Choice {
	s.ExposureType = new(ExposureType3Choice)
	return s.ExposureType
}

func (s *SettlementDetails50) AddFXStandingInstruction() *FXStandingInstruction1Choice {
	s.FXStandingInstruction = new(FXStandingInstruction1Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails50) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails50) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails50) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails50) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails50) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails50) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails50) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails50) AddTracking() *Tracking1Choice {
	s.Tracking = new(Tracking1Choice)
	return s.Tracking
}

func (s *SettlementDetails50) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails50) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails50) AddModificationCancellationAllowed() *ModificationCancellationAllowed1Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed1Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails50) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails50) AddSecuritiesSubBalanceType() *GenericIdentification20 {
	s.SecuritiesSubBalanceType = new(GenericIdentification20)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails50) AddCashSubBalanceType() *GenericIdentification20 {
	s.CashSubBalanceType = new(GenericIdentification20)
	return s.CashSubBalanceType
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails6 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType3Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition1Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade having caused the collateral movement.
	ExposureType *ExposureType1Choice `xml:"XpsrTp,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration was to occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement was executed using a letter of guarantee or if the physical certificates were used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails6) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails6) AddSecuritiesTransactionType() *SecuritiesTransactionType3Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType3Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails6) AddSettlementTransactionCondition() *SettlementTransactionCondition1Choice {
	newValue := new(SettlementTransactionCondition1Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails6) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails6) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails6) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails6) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails6) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails6) AddExposureType() *ExposureType1Choice {
	s.ExposureType = new(ExposureType1Choice)
	return s.ExposureType
}

func (s *SettlementDetails6) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails6) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails6) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails6) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails6) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails6) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails6) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails6) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails6) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails6) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails6) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails6) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails6) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails60 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition7Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails60) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails60) AddSettlementTransactionCondition() *SettlementTransactionCondition7Choice {
	newValue := new(SettlementTransactionCondition7Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails60) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails60) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails60) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails60) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails60) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails60) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails60) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails60) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails60) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails60) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails60) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails60) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails60) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails60) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails60) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails61 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition7Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails61) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails61) AddSettlementTransactionCondition() *SettlementTransactionCondition7Choice {
	newValue := new(SettlementTransactionCondition7Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails61) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails61) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails61) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails61) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails61) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails61) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails61) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails61) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails61) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails61) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails61) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails61) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails61) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails61) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails62 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition7Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction1Choice `xml:"FxStgInstr,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking1Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails62) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails62) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails62) AddSettlementTransactionCondition() *SettlementTransactionCondition7Choice {
	newValue := new(SettlementTransactionCondition7Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails62) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails62) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails62) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails62) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails62) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails62) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails62) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails62) AddFXStandingInstruction() *FXStandingInstruction1Choice {
	s.FXStandingInstruction = new(FXStandingInstruction1Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails62) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails62) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails62) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails62) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails62) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails62) AddTracking() *Tracking1Choice {
	s.Tracking = new(Tracking1Choice)
	return s.Tracking
}

func (s *SettlementDetails62) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails62) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails62) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails67 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator4 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType9Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn1Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType10Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction1Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType1Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity4Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking1Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed1Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification20 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification20 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails67) AddHoldIndicator() *HoldIndicator4 {
	s.HoldIndicator = new(HoldIndicator4)
	return s.HoldIndicator
}

func (s *SettlementDetails67) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails67) AddSecuritiesTransactionType() *SecuritiesTransactionType9Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType9Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails67) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails67) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails67) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails67) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails67) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails67) AddDeliveryReturnReason() *DeliveryReturn1Choice {
	s.DeliveryReturnReason = new(DeliveryReturn1Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails67) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails67) AddExposureType() *ExposureType10Choice {
	s.ExposureType = new(ExposureType10Choice)
	return s.ExposureType
}

func (s *SettlementDetails67) AddFXStandingInstruction() *FXStandingInstruction1Choice {
	s.FXStandingInstruction = new(FXStandingInstruction1Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails67) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails67) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails67) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails67) AddRepurchaseType() *RepurchaseType1Choice {
	s.RepurchaseType = new(RepurchaseType1Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails67) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails67) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails67) AddSettlingCapacity() *SettlingCapacity4Choice {
	s.SettlingCapacity = new(SettlingCapacity4Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails67) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails67) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails67) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails67) AddTracking() *Tracking1Choice {
	s.Tracking = new(Tracking1Choice)
	return s.Tracking
}

func (s *SettlementDetails67) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails67) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails67) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails67) AddModificationCancellationAllowed() *ModificationCancellationAllowed1Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed1Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails67) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails67) AddSecuritiesSubBalanceType() *GenericIdentification20 {
	s.SecuritiesSubBalanceType = new(GenericIdentification20)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails67) AddCashSubBalanceType() *GenericIdentification20 {
	s.CashSubBalanceType = new(GenericIdentification20)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails68 struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType9Choice `xml:"SctiesTxTp,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn1Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType10Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction1Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity4Choice `xml:"SttlgCpcty,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking1Choice `xml:"Trckg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed1Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification20 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification20 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails68) AddSecuritiesTransactionType() *SecuritiesTransactionType9Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType9Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails68) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails68) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails68) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails68) AddDeliveryReturnReason() *DeliveryReturn1Choice {
	s.DeliveryReturnReason = new(DeliveryReturn1Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails68) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails68) AddExposureType() *ExposureType10Choice {
	s.ExposureType = new(ExposureType10Choice)
	return s.ExposureType
}

func (s *SettlementDetails68) AddFXStandingInstruction() *FXStandingInstruction1Choice {
	s.FXStandingInstruction = new(FXStandingInstruction1Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails68) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails68) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails68) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails68) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails68) AddSettlingCapacity() *SettlingCapacity4Choice {
	s.SettlingCapacity = new(SettlingCapacity4Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails68) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails68) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails68) AddTracking() *Tracking1Choice {
	s.Tracking = new(Tracking1Choice)
	return s.Tracking
}

func (s *SettlementDetails68) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails68) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails68) AddModificationCancellationAllowed() *ModificationCancellationAllowed1Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed1Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails68) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails68) AddSecuritiesSubBalanceType() *GenericIdentification20 {
	s.SecuritiesSubBalanceType = new(GenericIdentification20)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails68) AddCashSubBalanceType() *GenericIdentification20 {
	s.CashSubBalanceType = new(GenericIdentification20)
	return s.CashSubBalanceType
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails69 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator4 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType9Choice `xml:"SctiesTxTp"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity4Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType10Choice `xml:"XpsrTp,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails69) AddHoldIndicator() *HoldIndicator4 {
	s.HoldIndicator = new(HoldIndicator4)
	return s.HoldIndicator
}

func (s *SettlementDetails69) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails69) AddSecuritiesTransactionType() *SecuritiesTransactionType9Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType9Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails69) AddSettlingCapacity() *SettlingCapacity4Choice {
	s.SettlingCapacity = new(SettlingCapacity4Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails69) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails69) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails69) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails69) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails69) AddExposureType() *ExposureType10Choice {
	s.ExposureType = new(ExposureType10Choice)
	return s.ExposureType
}

func (s *SettlementDetails69) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails69) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails69) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails69) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails69) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails69) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails69) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails69) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails69) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails69) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails69) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails69) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails7 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition1Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails7) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails7) AddSettlementTransactionCondition() *SettlementTransactionCondition1Choice {
	newValue := new(SettlementTransactionCondition1Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails7) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails7) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails7) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails7) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails7) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails7) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails7) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails7) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails7) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails7) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails7) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails7) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails7) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails7) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails7) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails7) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails70 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType10Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade having caused the collateral movement.
	ExposureType *ExposureType10Choice `xml:"XpsrTp,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration was to occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity4Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement was executed using a letter of guarantee or if the physical certificates were used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification20 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification20 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails70) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails70) AddSecuritiesTransactionType() *SecuritiesTransactionType10Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType10Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails70) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails70) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails70) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails70) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails70) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails70) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails70) AddExposureType() *ExposureType10Choice {
	s.ExposureType = new(ExposureType10Choice)
	return s.ExposureType
}

func (s *SettlementDetails70) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails70) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails70) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails70) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails70) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails70) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails70) AddSettlingCapacity() *SettlingCapacity4Choice {
	s.SettlingCapacity = new(SettlingCapacity4Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails70) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails70) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails70) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails70) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails70) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails70) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails70) AddSecuritiesSubBalanceType() *GenericIdentification20 {
	s.SecuritiesSubBalanceType = new(GenericIdentification20)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails70) AddCashSubBalanceType() *GenericIdentification20 {
	s.CashSubBalanceType = new(GenericIdentification20)
	return s.CashSubBalanceType
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails71 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition7Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity4Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails71) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails71) AddSettlementTransactionCondition() *SettlementTransactionCondition7Choice {
	newValue := new(SettlementTransactionCondition7Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails71) AddSettlingCapacity() *SettlingCapacity4Choice {
	s.SettlingCapacity = new(SettlingCapacity4Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails71) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails71) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails71) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails71) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails71) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails71) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails71) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails71) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails71) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails71) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails71) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails71) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails71) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails71) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails72 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition7Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity4Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction1Choice `xml:"FxStgInstr,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking1Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails72) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails72) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails72) AddSettlementTransactionCondition() *SettlementTransactionCondition7Choice {
	newValue := new(SettlementTransactionCondition7Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails72) AddSettlingCapacity() *SettlingCapacity4Choice {
	s.SettlingCapacity = new(SettlingCapacity4Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails72) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails72) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails72) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails72) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails72) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails72) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails72) AddFXStandingInstruction() *FXStandingInstruction1Choice {
	s.FXStandingInstruction = new(FXStandingInstruction1Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails72) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails72) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails72) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails72) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails72) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails72) AddTracking() *Tracking1Choice {
	s.Tracking = new(Tracking1Choice)
	return s.Tracking
}

func (s *SettlementDetails72) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails72) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails72) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails73 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition7Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity4Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails73) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails73) AddSettlementTransactionCondition() *SettlementTransactionCondition7Choice {
	newValue := new(SettlementTransactionCondition7Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails73) AddSettlingCapacity() *SettlingCapacity4Choice {
	s.SettlingCapacity = new(SettlingCapacity4Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails73) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails73) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails73) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails73) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails73) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails73) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails73) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails73) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails73) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails73) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails73) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails73) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails73) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails74 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator4 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity4Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails74) AddHoldIndicator() *HoldIndicator4 {
	s.HoldIndicator = new(HoldIndicator4)
	return s.HoldIndicator
}

func (s *SettlementDetails74) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails74) AddSettlingCapacity() *SettlingCapacity4Choice {
	s.SettlingCapacity = new(SettlingCapacity4Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails74) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails74) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails74) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails74) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails74) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails74) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails74) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails74) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails74) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails74) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails74) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails74) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails74) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails74) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails74) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails75 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity4Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails75) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails75) AddSettlingCapacity() *SettlingCapacity4Choice {
	s.SettlingCapacity = new(SettlingCapacity4Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails75) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails75) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails75) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails75) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails75) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails75) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails75) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails75) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails75) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails75) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails75) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails75) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails75) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails75) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails75) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails8 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition1Choice `xml:"SttlmTxCond,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType1Choice `xml:"SctiesTxTp"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType1Choice `xml:"XpsrTp,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails8) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails8) AddSettlementTransactionCondition() *SettlementTransactionCondition1Choice {
	newValue := new(SettlementTransactionCondition1Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails8) AddSecuritiesTransactionType() *SecuritiesTransactionType1Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType1Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails8) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails8) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails8) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails8) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails8) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails8) AddExposureType() *ExposureType1Choice {
	s.ExposureType = new(ExposureType1Choice)
	return s.ExposureType
}

func (s *SettlementDetails8) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails8) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails8) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails8) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails8) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails8) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails8) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails8) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails8) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails8) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails8) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails8) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails88 struct {

	// Indicates the date as known by the two parties to be used for matching purposes when settlement of securities occurs.
	TradeDate *ISODateTime `xml:"TradDt"`

	// Provides details on either the delivering or receiving settlement parties.
	SettlementParties *SettlementParties3Choice `xml:"SttlmPties,omitempty"`

	// Indicates the collateral ownership.
	CollateralOwnership *CollateralOwnership1 `xml:"CollOwnrsh"`
}

func (s *SettlementDetails88) SetTradeDate(value string) {
	s.TradeDate = (*ISODateTime)(&value)
}

func (s *SettlementDetails88) AddSettlementParties() *SettlementParties3Choice {
	s.SettlementParties = new(SettlementParties3Choice)
	return s.SettlementParties
}

func (s *SettlementDetails88) AddCollateralOwnership() *CollateralOwnership1 {
	s.CollateralOwnership = new(CollateralOwnership1)
	return s.CollateralOwnership
}

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails9 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition1Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails9) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails9) AddSettlementTransactionCondition() *SettlementTransactionCondition1Choice {
	newValue := new(SettlementTransactionCondition1Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails9) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails9) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails9) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails9) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails9) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails9) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails9) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails9) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails9) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails9) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails9) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails9) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails9) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails9) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails9) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails90 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType18Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade having caused the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration was to occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType12Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing6Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement was executed using a letter of guarantee or if the physical certificates were used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification30 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification30 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails90) AddPriority() *PriorityNumeric4Choice {
	s.Priority = new(PriorityNumeric4Choice)
	return s.Priority
}

func (s *SettlementDetails90) AddSecuritiesTransactionType() *SecuritiesTransactionType18Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType18Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails90) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails90) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails90) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails90) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails90) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails90) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails90) AddExposureType() *ExposureType16Choice {
	s.ExposureType = new(ExposureType16Choice)
	return s.ExposureType
}

func (s *SettlementDetails90) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails90) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails90) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails90) AddRepurchaseType() *RepurchaseType12Choice {
	s.RepurchaseType = new(RepurchaseType12Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails90) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails90) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails90) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails90) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails90) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails90) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails90) AddAutomaticBorrowing() *AutomaticBorrowing6Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing6Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails90) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails90) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails90) AddSecuritiesSubBalanceType() *GenericIdentification30 {
	s.SecuritiesSubBalanceType = new(GenericIdentification30)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails90) AddCashSubBalanceType() *GenericIdentification30 {
	s.CashSubBalanceType = new(GenericIdentification30)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails91 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType12Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails91) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails91) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails91) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails91) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails91) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails91) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails91) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails91) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails91) AddRepurchaseType() *RepurchaseType12Choice {
	s.RepurchaseType = new(RepurchaseType12Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails91) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails91) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails91) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails91) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails91) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails91) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails91) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails91) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails92 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator6 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType12Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails92) AddHoldIndicator() *HoldIndicator6 {
	s.HoldIndicator = new(HoldIndicator6)
	return s.HoldIndicator
}

func (s *SettlementDetails92) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails92) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails92) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails92) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails92) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails92) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails92) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails92) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails92) AddRepurchaseType() *RepurchaseType12Choice {
	s.RepurchaseType = new(RepurchaseType12Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails92) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails92) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails92) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails92) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails92) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails92) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails92) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails92) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails93 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator6 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType21Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn3Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction4Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType14Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking4Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing6Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed4Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification30 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification30 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails93) AddHoldIndicator() *HoldIndicator6 {
	s.HoldIndicator = new(HoldIndicator6)
	return s.HoldIndicator
}

func (s *SettlementDetails93) AddPriority() *PriorityNumeric4Choice {
	s.Priority = new(PriorityNumeric4Choice)
	return s.Priority
}

func (s *SettlementDetails93) AddSecuritiesTransactionType() *SecuritiesTransactionType21Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType21Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails93) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails93) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails93) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails93) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails93) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails93) AddDeliveryReturnReason() *DeliveryReturn3Choice {
	s.DeliveryReturnReason = new(DeliveryReturn3Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails93) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails93) AddExposureType() *ExposureType16Choice {
	s.ExposureType = new(ExposureType16Choice)
	return s.ExposureType
}

func (s *SettlementDetails93) AddFXStandingInstruction() *FXStandingInstruction4Choice {
	s.FXStandingInstruction = new(FXStandingInstruction4Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails93) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails93) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails93) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails93) AddRepurchaseType() *RepurchaseType14Choice {
	s.RepurchaseType = new(RepurchaseType14Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails93) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails93) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails93) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails93) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails93) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails93) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails93) AddTracking() *Tracking4Choice {
	s.Tracking = new(Tracking4Choice)
	return s.Tracking
}

func (s *SettlementDetails93) AddAutomaticBorrowing() *AutomaticBorrowing6Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing6Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails93) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails93) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails93) AddModificationCancellationAllowed() *ModificationCancellationAllowed4Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed4Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails93) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails93) AddSecuritiesSubBalanceType() *GenericIdentification30 {
	s.SecuritiesSubBalanceType = new(GenericIdentification30)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails93) AddCashSubBalanceType() *GenericIdentification30 {
	s.CashSubBalanceType = new(GenericIdentification30)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails94 struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType21Choice `xml:"SctiesTxTp,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn3Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction4Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking4Choice `xml:"Trckg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed4Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification30 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification30 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails94) AddSecuritiesTransactionType() *SecuritiesTransactionType21Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType21Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails94) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails94) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails94) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails94) AddDeliveryReturnReason() *DeliveryReturn3Choice {
	s.DeliveryReturnReason = new(DeliveryReturn3Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails94) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails94) AddExposureType() *ExposureType16Choice {
	s.ExposureType = new(ExposureType16Choice)
	return s.ExposureType
}

func (s *SettlementDetails94) AddFXStandingInstruction() *FXStandingInstruction4Choice {
	s.FXStandingInstruction = new(FXStandingInstruction4Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails94) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails94) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails94) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails94) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails94) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails94) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails94) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails94) AddTracking() *Tracking4Choice {
	s.Tracking = new(Tracking4Choice)
	return s.Tracking
}

func (s *SettlementDetails94) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails94) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails94) AddModificationCancellationAllowed() *ModificationCancellationAllowed4Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed4Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails94) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails94) AddSecuritiesSubBalanceType() *GenericIdentification30 {
	s.SecuritiesSubBalanceType = new(GenericIdentification30)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails94) AddCashSubBalanceType() *GenericIdentification30 {
	s.CashSubBalanceType = new(GenericIdentification30)
	return s.CashSubBalanceType
}

// Details of settlement of a transaction.
type SettlementDetails95 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator6 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType21Choice `xml:"SctiesTxTp"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType12Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails95) AddHoldIndicator() *HoldIndicator6 {
	s.HoldIndicator = new(HoldIndicator6)
	return s.HoldIndicator
}

func (s *SettlementDetails95) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails95) AddSecuritiesTransactionType() *SecuritiesTransactionType21Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType21Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails95) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails95) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails95) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails95) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails95) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails95) AddExposureType() *ExposureType16Choice {
	s.ExposureType = new(ExposureType16Choice)
	return s.ExposureType
}

func (s *SettlementDetails95) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails95) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails95) AddRepurchaseType() *RepurchaseType12Choice {
	s.RepurchaseType = new(RepurchaseType12Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails95) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails95) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails95) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails95) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails95) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails95) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails95) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails95) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails95) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails96 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition18Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing6Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails96) AddPriority() *PriorityNumeric4Choice {
	s.Priority = new(PriorityNumeric4Choice)
	return s.Priority
}

func (s *SettlementDetails96) AddSettlementTransactionCondition() *SettlementTransactionCondition18Choice {
	newValue := new(SettlementTransactionCondition18Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails96) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails96) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails96) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails96) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails96) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails96) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails96) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails96) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails96) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails96) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails96) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails96) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails96) AddAutomaticBorrowing() *AutomaticBorrowing6Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing6Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails96) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails96) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails97 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition18Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction4Choice `xml:"FxStgInstr,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking4Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing6Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails97) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails97) AddPriority() *PriorityNumeric4Choice {
	s.Priority = new(PriorityNumeric4Choice)
	return s.Priority
}

func (s *SettlementDetails97) AddSettlementTransactionCondition() *SettlementTransactionCondition18Choice {
	newValue := new(SettlementTransactionCondition18Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails97) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails97) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails97) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails97) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails97) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails97) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails97) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails97) AddFXStandingInstruction() *FXStandingInstruction4Choice {
	s.FXStandingInstruction = new(FXStandingInstruction4Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails97) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails97) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails97) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails97) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails97) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails97) AddTracking() *Tracking4Choice {
	s.Tracking = new(Tracking4Choice)
	return s.Tracking
}

func (s *SettlementDetails97) AddAutomaticBorrowing() *AutomaticBorrowing6Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing6Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails97) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails97) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails98 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition18Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails98) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails98) AddSettlementTransactionCondition() *SettlementTransactionCondition18Choice {
	newValue := new(SettlementTransactionCondition18Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails98) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails98) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails98) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails98) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails98) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails98) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails98) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails98) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails98) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails98) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails98) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails98) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails98) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails98) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

// Details of settlement of a transaction.
type SettlementDetails99 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType20Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition17Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType13Choice `xml:"RpTp,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`
}

func (s *SettlementDetails99) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails99) AddSecuritiesTransactionType() *SecuritiesTransactionType20Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType20Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails99) AddSettlementTransactionCondition() *SettlementTransactionCondition17Choice {
	newValue := new(SettlementTransactionCondition17Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails99) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails99) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails99) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails99) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails99) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails99) AddRepurchaseType() *RepurchaseType13Choice {
	s.RepurchaseType = new(RepurchaseType13Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails99) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails99) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}
