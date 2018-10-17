package iso20022

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails1 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate2Choice `xml:"TermntnDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType5Choice `xml:"RateTp,omitempty"`

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Revaluation *RevaluationIndicator1Choice `xml:"Rvaltn,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework1Choice `xml:"LglFrmwk,omitempty"`

	// Identifies the computation method of accrued interest of the related financial instrument.
	InterestComputationMethod *InterestComputationMethodFormat1Choice `xml:"IntrstCmptnMtd,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Total number of collateral instructions involved in the transaction.
	TotalNumberOfCollateralInstructions *Exact3NumericText `xml:"TtlNbOfCollInstrs,omitempty"`

	// Principal amount of a trade (for second leg).
	DealAmount *AmountAndDirection4 `xml:"DealAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection4 `xml:"AcrdIntrstAmt,omitempty"`

	// Fixed amount of money that has to be paid (instead of interest) in the case of a recall or at the closing date.
	ForfeitAmount *AmountAndDirection4 `xml:"FrftAmt,omitempty"`

	// Difference between the amount of money of the first leg and the amount of the second leg of the transaction.
	PremiumAmount *AmountAndDirection4 `xml:"PrmAmt,omitempty"`

	// Amount of money to be settled per piece of collateral to terminate the transaction.
	TerminationAmountPerPieceOfCollateral *AmountAndDirection4 `xml:"TermntnAmtPerPcOfColl,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection4 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *Max140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails1) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails1) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails1) AddTerminationDate() *TerminationDate2Choice {
	s.TerminationDate = new(TerminationDate2Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails1) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails1) AddRateType() *RateType5Choice {
	s.RateType = new(RateType5Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails1) AddRevaluation() *RevaluationIndicator1Choice {
	s.Revaluation = new(RevaluationIndicator1Choice)
	return s.Revaluation
}

func (s *SecuritiesFinancingTransactionDetails1) AddLegalFramework() *LegalFramework1Choice {
	s.LegalFramework = new(LegalFramework1Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails1) AddInterestComputationMethod() *InterestComputationMethodFormat1Choice {
	s.InterestComputationMethod = new(InterestComputationMethodFormat1Choice)
	return s.InterestComputationMethod
}

func (s *SecuritiesFinancingTransactionDetails1) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails1) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails1) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails1) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails1) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails1) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails1) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails1) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails1) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails1) SetTotalNumberOfCollateralInstructions(value string) {
	s.TotalNumberOfCollateralInstructions = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails1) AddDealAmount() *AmountAndDirection4 {
	s.DealAmount = new(AmountAndDirection4)
	return s.DealAmount
}

func (s *SecuritiesFinancingTransactionDetails1) AddAccruedInterestAmount() *AmountAndDirection4 {
	s.AccruedInterestAmount = new(AmountAndDirection4)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails1) AddForfeitAmount() *AmountAndDirection4 {
	s.ForfeitAmount = new(AmountAndDirection4)
	return s.ForfeitAmount
}

func (s *SecuritiesFinancingTransactionDetails1) AddPremiumAmount() *AmountAndDirection4 {
	s.PremiumAmount = new(AmountAndDirection4)
	return s.PremiumAmount
}

func (s *SecuritiesFinancingTransactionDetails1) AddTerminationAmountPerPieceOfCollateral() *AmountAndDirection4 {
	s.TerminationAmountPerPieceOfCollateral = new(AmountAndDirection4)
	return s.TerminationAmountPerPieceOfCollateral
}

func (s *SecuritiesFinancingTransactionDetails1) AddTerminationTransactionAmount() *AmountAndDirection4 {
	s.TerminationTransactionAmount = new(AmountAndDirection4)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails1) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*Max140Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails11 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate2Choice `xml:"TermntnDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Earliest date/time at which the call back can take place.
	EarliestCallBackDate *DateAndDateTimeChoice `xml:"EarlstCallBckDt,omitempty"`

	// Date/time at which the commission is calculated.
	CommissionCalculationDate *DateAndDateTimeChoice `xml:"ComssnClctnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType5Choice `xml:"RateTp,omitempty"`

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Revaluation *RevaluationIndicator1Choice `xml:"Rvaltn,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework1Choice `xml:"LglFrmwk,omitempty"`

	// Identifies the computation method of accrued interest of the related financial instrument.
	InterestComputationMethod *InterestComputationMethodFormat1Choice `xml:"IntrstCmptnMtd,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate paid in the context of a securities financing transaction.
	ChargesRate *Rate2 `xml:"ChrgsRate,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Total number of collateral instructions involved in the transaction.
	TotalNumberOfCollateralInstructions *Exact3NumericText `xml:"TtlNbOfCollInstrs,omitempty"`

	// Principal amount of a trade (for second leg).
	DealAmount *AmountAndDirection4 `xml:"DealAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection4 `xml:"AcrdIntrstAmt,omitempty"`

	// Fixed amount of money that has to be paid (instead of interest) in the case of a recall or at the closing date.
	ForfeitAmount *AmountAndDirection4 `xml:"FrftAmt,omitempty"`

	// Difference between the amount of money of the first leg and the amount of the second leg of the transaction.
	PremiumAmount *AmountAndDirection4 `xml:"PrmAmt,omitempty"`

	// Amount of money to be settled per piece of collateral to terminate the transaction.
	TerminationAmountPerPieceOfCollateral *AmountAndDirection4 `xml:"TermntnAmtPerPcOfColl,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection4 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *Max140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails11) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails11) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails11) AddTerminationDate() *TerminationDate2Choice {
	s.TerminationDate = new(TerminationDate2Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails11) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails11) AddEarliestCallBackDate() *DateAndDateTimeChoice {
	s.EarliestCallBackDate = new(DateAndDateTimeChoice)
	return s.EarliestCallBackDate
}

func (s *SecuritiesFinancingTransactionDetails11) AddCommissionCalculationDate() *DateAndDateTimeChoice {
	s.CommissionCalculationDate = new(DateAndDateTimeChoice)
	return s.CommissionCalculationDate
}

func (s *SecuritiesFinancingTransactionDetails11) AddRateType() *RateType5Choice {
	s.RateType = new(RateType5Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails11) AddRevaluation() *RevaluationIndicator1Choice {
	s.Revaluation = new(RevaluationIndicator1Choice)
	return s.Revaluation
}

func (s *SecuritiesFinancingTransactionDetails11) AddLegalFramework() *LegalFramework1Choice {
	s.LegalFramework = new(LegalFramework1Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails11) AddInterestComputationMethod() *InterestComputationMethodFormat1Choice {
	s.InterestComputationMethod = new(InterestComputationMethodFormat1Choice)
	return s.InterestComputationMethod
}

func (s *SecuritiesFinancingTransactionDetails11) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails11) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails11) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails11) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails11) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails11) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails11) AddChargesRate() *Rate2 {
	s.ChargesRate = new(Rate2)
	return s.ChargesRate
}

func (s *SecuritiesFinancingTransactionDetails11) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails11) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails11) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails11) SetTotalNumberOfCollateralInstructions(value string) {
	s.TotalNumberOfCollateralInstructions = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails11) AddDealAmount() *AmountAndDirection4 {
	s.DealAmount = new(AmountAndDirection4)
	return s.DealAmount
}

func (s *SecuritiesFinancingTransactionDetails11) AddAccruedInterestAmount() *AmountAndDirection4 {
	s.AccruedInterestAmount = new(AmountAndDirection4)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails11) AddForfeitAmount() *AmountAndDirection4 {
	s.ForfeitAmount = new(AmountAndDirection4)
	return s.ForfeitAmount
}

func (s *SecuritiesFinancingTransactionDetails11) AddPremiumAmount() *AmountAndDirection4 {
	s.PremiumAmount = new(AmountAndDirection4)
	return s.PremiumAmount
}

func (s *SecuritiesFinancingTransactionDetails11) AddTerminationAmountPerPieceOfCollateral() *AmountAndDirection4 {
	s.TerminationAmountPerPieceOfCollateral = new(AmountAndDirection4)
	return s.TerminationAmountPerPieceOfCollateral
}

func (s *SecuritiesFinancingTransactionDetails11) AddTerminationTransactionAmount() *AmountAndDirection4 {
	s.TerminationTransactionAmount = new(AmountAndDirection4)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails11) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*Max140Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails17 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *AmountAndDirection8 `xml:"OpngSttlmAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection4 `xml:"TermntnTxAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *SettlementDate2Choice `xml:"OpngSttlmDt"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate2Choice `xml:"TermntnDt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails61 `xml:"SttlmParams,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType5Choice `xml:"RateTp,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails17) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails17) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails17) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails17) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails17) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails17) SetClientTripartyCollateralTransactionIdentification(value string) {
	s.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails17) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesFinancingTransactionDetails17) AddSafekeepingAccount() *SecuritiesAccount13 {
	s.SafekeepingAccount = new(SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesFinancingTransactionDetails17) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesFinancingTransactionDetails17) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesFinancingTransactionDetails17) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingTransactionDetails17) AddSettlementQuantity() *Quantity6Choice {
	s.SettlementQuantity = new(Quantity6Choice)
	return s.SettlementQuantity
}

func (s *SecuritiesFinancingTransactionDetails17) AddOpeningSettlementAmount() *AmountAndDirection8 {
	s.OpeningSettlementAmount = new(AmountAndDirection8)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingTransactionDetails17) AddTerminationTransactionAmount() *AmountAndDirection4 {
	s.TerminationTransactionAmount = new(AmountAndDirection4)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails17) AddOpeningSettlementDate() *SettlementDate2Choice {
	s.OpeningSettlementDate = new(SettlementDate2Choice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails17) AddTerminationDate() *TerminationDate2Choice {
	s.TerminationDate = new(TerminationDate2Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails17) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesFinancingTransactionDetails17) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	s.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return s.ExpectedSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails17) AddExpectedValueDate() *DateAndDateTimeChoice {
	s.ExpectedValueDate = new(DateAndDateTimeChoice)
	return s.ExpectedValueDate
}

func (s *SecuritiesFinancingTransactionDetails17) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesFinancingTransactionDetails17) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails17) SetSecuritiesFinancingTransactionType(value string) {
	s.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails17) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails17) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails17) AddSettlementParameters() *SettlementDetails61 {
	s.SettlementParameters = new(SettlementDetails61)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingTransactionDetails17) AddRateType() *RateType5Choice {
	s.RateType = new(RateType5Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails17) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails17) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails17) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails17) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails17) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails17) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails17) AddDeliveringSettlementParties() *SettlementParties11 {
	s.DeliveringSettlementParties = new(SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails17) AddReceivingSettlementParties() *SettlementParties11 {
	s.ReceivingSettlementParties = new(SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails17) AddInvestor() *PartyIdentification37Choice {
	s.Investor = new(PartyIdentification37Choice)
	return s.Investor
}

func (s *SecuritiesFinancingTransactionDetails17) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails19 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate2Choice `xml:"TermntnDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Earliest date/time at which the call back can take place.
	EarliestCallBackDate *DateAndDateTimeChoice `xml:"EarlstCallBckDt,omitempty"`

	// Date/time at which the commission is calculated.
	CommissionCalculationDate *DateAndDateTimeChoice `xml:"ComssnClctnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType5Choice `xml:"RateTp,omitempty"`

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Revaluation *RevaluationIndicator1Choice `xml:"Rvaltn,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework1Choice `xml:"LglFrmwk,omitempty"`

	// Identifies the computation method of accrued interest of the related financial instrument.
	InterestComputationMethod *InterestComputationMethodFormat1Choice `xml:"IntrstCmptnMtd,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate paid in the context of a securities financing transaction.
	ChargesRate *Rate2 `xml:"ChrgsRate,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Total number of collateral instructions involved in the transaction.
	TotalNumberOfCollateralInstructions *Exact3NumericText `xml:"TtlNbOfCollInstrs,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection4 `xml:"LclBrkrComssn,omitempty"`

	// Principal amount of a trade (for second leg).
	DealAmount *AmountAndDirection4 `xml:"DealAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection4 `xml:"AcrdIntrstAmt,omitempty"`

	// Fixed amount of money that has to be paid (instead of interest) in the case of a recall or at the closing date.
	ForfeitAmount *AmountAndDirection4 `xml:"FrftAmt,omitempty"`

	// Difference between the amount of money of the first leg and the amount of the second leg of the transaction.
	PremiumAmount *AmountAndDirection4 `xml:"PrmAmt,omitempty"`

	// Amount of money to be settled per piece of collateral to terminate the transaction.
	TerminationAmountPerPieceOfCollateral *AmountAndDirection4 `xml:"TermntnAmtPerPcOfColl,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection4 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *Max140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails19) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails19) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails19) AddTerminationDate() *TerminationDate2Choice {
	s.TerminationDate = new(TerminationDate2Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails19) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails19) AddEarliestCallBackDate() *DateAndDateTimeChoice {
	s.EarliestCallBackDate = new(DateAndDateTimeChoice)
	return s.EarliestCallBackDate
}

func (s *SecuritiesFinancingTransactionDetails19) AddCommissionCalculationDate() *DateAndDateTimeChoice {
	s.CommissionCalculationDate = new(DateAndDateTimeChoice)
	return s.CommissionCalculationDate
}

func (s *SecuritiesFinancingTransactionDetails19) AddRateType() *RateType5Choice {
	s.RateType = new(RateType5Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails19) AddRevaluation() *RevaluationIndicator1Choice {
	s.Revaluation = new(RevaluationIndicator1Choice)
	return s.Revaluation
}

func (s *SecuritiesFinancingTransactionDetails19) AddLegalFramework() *LegalFramework1Choice {
	s.LegalFramework = new(LegalFramework1Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails19) AddInterestComputationMethod() *InterestComputationMethodFormat1Choice {
	s.InterestComputationMethod = new(InterestComputationMethodFormat1Choice)
	return s.InterestComputationMethod
}

func (s *SecuritiesFinancingTransactionDetails19) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails19) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails19) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails19) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails19) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails19) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails19) AddChargesRate() *Rate2 {
	s.ChargesRate = new(Rate2)
	return s.ChargesRate
}

func (s *SecuritiesFinancingTransactionDetails19) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails19) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails19) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails19) SetTotalNumberOfCollateralInstructions(value string) {
	s.TotalNumberOfCollateralInstructions = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails19) AddLocalBrokerCommission() *AmountAndDirection4 {
	s.LocalBrokerCommission = new(AmountAndDirection4)
	return s.LocalBrokerCommission
}

func (s *SecuritiesFinancingTransactionDetails19) AddDealAmount() *AmountAndDirection4 {
	s.DealAmount = new(AmountAndDirection4)
	return s.DealAmount
}

func (s *SecuritiesFinancingTransactionDetails19) AddAccruedInterestAmount() *AmountAndDirection4 {
	s.AccruedInterestAmount = new(AmountAndDirection4)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails19) AddForfeitAmount() *AmountAndDirection4 {
	s.ForfeitAmount = new(AmountAndDirection4)
	return s.ForfeitAmount
}

func (s *SecuritiesFinancingTransactionDetails19) AddPremiumAmount() *AmountAndDirection4 {
	s.PremiumAmount = new(AmountAndDirection4)
	return s.PremiumAmount
}

func (s *SecuritiesFinancingTransactionDetails19) AddTerminationAmountPerPieceOfCollateral() *AmountAndDirection4 {
	s.TerminationAmountPerPieceOfCollateral = new(AmountAndDirection4)
	return s.TerminationAmountPerPieceOfCollateral
}

func (s *SecuritiesFinancingTransactionDetails19) AddTerminationTransactionAmount() *AmountAndDirection4 {
	s.TerminationTransactionAmount = new(AmountAndDirection4)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails19) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*Max140Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails2 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *AmountAndDirection8 `xml:"OpngSttlmAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection4 `xml:"TermntnTxAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *SettlementDate2Choice `xml:"OpngSttlmDt"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate2Choice `xml:"TermntnDt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails10 `xml:"SttlmParams,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType5Choice `xml:"RateTp,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties5 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties5 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification14Choice `xml:"Invstr,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails2) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails2) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails2) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails2) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails2) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails2) SetClientTripartyCollateralTransactionIdentification(value string) {
	s.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails2) AddAccountOwner() *PartyIdentification13Choice {
	s.AccountOwner = new(PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SecuritiesFinancingTransactionDetails2) AddSafekeepingAccount() *SecuritiesAccount13 {
	s.SafekeepingAccount = new(SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesFinancingTransactionDetails2) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesFinancingTransactionDetails2) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesFinancingTransactionDetails2) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingTransactionDetails2) AddSettlementQuantity() *Quantity6Choice {
	s.SettlementQuantity = new(Quantity6Choice)
	return s.SettlementQuantity
}

func (s *SecuritiesFinancingTransactionDetails2) AddOpeningSettlementAmount() *AmountAndDirection8 {
	s.OpeningSettlementAmount = new(AmountAndDirection8)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingTransactionDetails2) AddTerminationTransactionAmount() *AmountAndDirection4 {
	s.TerminationTransactionAmount = new(AmountAndDirection4)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails2) AddOpeningSettlementDate() *SettlementDate2Choice {
	s.OpeningSettlementDate = new(SettlementDate2Choice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails2) AddTerminationDate() *TerminationDate2Choice {
	s.TerminationDate = new(TerminationDate2Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails2) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesFinancingTransactionDetails2) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	s.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return s.ExpectedSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails2) AddExpectedValueDate() *DateAndDateTimeChoice {
	s.ExpectedValueDate = new(DateAndDateTimeChoice)
	return s.ExpectedValueDate
}

func (s *SecuritiesFinancingTransactionDetails2) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesFinancingTransactionDetails2) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails2) SetSecuritiesFinancingTransactionType(value string) {
	s.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails2) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails2) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails2) AddSettlementParameters() *SettlementDetails10 {
	s.SettlementParameters = new(SettlementDetails10)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingTransactionDetails2) AddRateType() *RateType5Choice {
	s.RateType = new(RateType5Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails2) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails2) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails2) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails2) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails2) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails2) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails2) AddDeliveringSettlementParties() *SettlementParties5 {
	s.DeliveringSettlementParties = new(SettlementParties5)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails2) AddReceivingSettlementParties() *SettlementParties5 {
	s.ReceivingSettlementParties = new(SettlementParties5)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails2) AddInvestor() *PartyIdentification14Choice {
	s.Investor = new(PartyIdentification14Choice)
	return s.Investor
}

func (s *SecuritiesFinancingTransactionDetails2) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails21 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *AmountAndDirection8 `xml:"OpngSttlmAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection4 `xml:"TermntnTxAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *SettlementDate2Choice `xml:"OpngSttlmDt"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate2Choice `xml:"TermntnDt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails73 `xml:"SttlmParams,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType5Choice `xml:"RateTp,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails21) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails21) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails21) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails21) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails21) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails21) SetClientTripartyCollateralTransactionIdentification(value string) {
	s.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails21) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesFinancingTransactionDetails21) AddSafekeepingAccount() *SecuritiesAccount13 {
	s.SafekeepingAccount = new(SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesFinancingTransactionDetails21) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesFinancingTransactionDetails21) AddPlaceOfTrade() *MarketIdentification78 {
	s.PlaceOfTrade = new(MarketIdentification78)
	return s.PlaceOfTrade
}

func (s *SecuritiesFinancingTransactionDetails21) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingTransactionDetails21) AddSettlementQuantity() *Quantity6Choice {
	s.SettlementQuantity = new(Quantity6Choice)
	return s.SettlementQuantity
}

func (s *SecuritiesFinancingTransactionDetails21) AddOpeningSettlementAmount() *AmountAndDirection8 {
	s.OpeningSettlementAmount = new(AmountAndDirection8)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingTransactionDetails21) AddTerminationTransactionAmount() *AmountAndDirection4 {
	s.TerminationTransactionAmount = new(AmountAndDirection4)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails21) AddOpeningSettlementDate() *SettlementDate2Choice {
	s.OpeningSettlementDate = new(SettlementDate2Choice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails21) AddTerminationDate() *TerminationDate2Choice {
	s.TerminationDate = new(TerminationDate2Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails21) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesFinancingTransactionDetails21) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	s.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return s.ExpectedSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails21) AddExpectedValueDate() *DateAndDateTimeChoice {
	s.ExpectedValueDate = new(DateAndDateTimeChoice)
	return s.ExpectedValueDate
}

func (s *SecuritiesFinancingTransactionDetails21) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesFinancingTransactionDetails21) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails21) SetSecuritiesFinancingTransactionType(value string) {
	s.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails21) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails21) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails21) AddSettlementParameters() *SettlementDetails73 {
	s.SettlementParameters = new(SettlementDetails73)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingTransactionDetails21) AddRateType() *RateType5Choice {
	s.RateType = new(RateType5Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails21) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails21) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails21) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails21) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails21) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails21) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails21) AddDeliveringSettlementParties() *SettlementParties11 {
	s.DeliveringSettlementParties = new(SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails21) AddReceivingSettlementParties() *SettlementParties11 {
	s.ReceivingSettlementParties = new(SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails21) AddInvestor() *PartyIdentification37Choice {
	s.Investor = new(PartyIdentification37Choice)
	return s.Investor
}

func (s *SecuritiesFinancingTransactionDetails21) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails24 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *AmountAndDirection8 `xml:"OpngSttlmAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection4 `xml:"TermntnTxAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *SettlementDate2Choice `xml:"OpngSttlmDt"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate2Choice `xml:"TermntnDt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails73 `xml:"SttlmParams,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType5Choice `xml:"RateTp,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails24) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails24) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails24) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails24) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails24) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails24) SetClientTripartyCollateralTransactionIdentification(value string) {
	s.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails24) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesFinancingTransactionDetails24) AddSafekeepingAccount() *SecuritiesAccount13 {
	s.SafekeepingAccount = new(SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesFinancingTransactionDetails24) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesFinancingTransactionDetails24) AddPlaceOfTrade() *MarketIdentification78 {
	s.PlaceOfTrade = new(MarketIdentification78)
	return s.PlaceOfTrade
}

func (s *SecuritiesFinancingTransactionDetails24) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingTransactionDetails24) AddSettlementQuantity() *Quantity6Choice {
	s.SettlementQuantity = new(Quantity6Choice)
	return s.SettlementQuantity
}

func (s *SecuritiesFinancingTransactionDetails24) AddOpeningSettlementAmount() *AmountAndDirection8 {
	s.OpeningSettlementAmount = new(AmountAndDirection8)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingTransactionDetails24) AddTerminationTransactionAmount() *AmountAndDirection4 {
	s.TerminationTransactionAmount = new(AmountAndDirection4)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails24) AddOpeningSettlementDate() *SettlementDate2Choice {
	s.OpeningSettlementDate = new(SettlementDate2Choice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails24) AddTerminationDate() *TerminationDate2Choice {
	s.TerminationDate = new(TerminationDate2Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails24) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesFinancingTransactionDetails24) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	s.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return s.ExpectedSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails24) AddExpectedValueDate() *DateAndDateTimeChoice {
	s.ExpectedValueDate = new(DateAndDateTimeChoice)
	return s.ExpectedValueDate
}

func (s *SecuritiesFinancingTransactionDetails24) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesFinancingTransactionDetails24) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails24) SetSecuritiesFinancingTransactionType(value string) {
	s.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails24) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails24) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails24) AddSettlementParameters() *SettlementDetails73 {
	s.SettlementParameters = new(SettlementDetails73)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingTransactionDetails24) AddRateType() *RateType5Choice {
	s.RateType = new(RateType5Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails24) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails24) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails24) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails24) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails24) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails24) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails24) AddDeliveringSettlementParties() *SettlementParties11 {
	s.DeliveringSettlementParties = new(SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails24) AddReceivingSettlementParties() *SettlementParties11 {
	s.ReceivingSettlementParties = new(SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails24) AddInvestor() *PartyIdentification37Choice {
	s.Investor = new(PartyIdentification37Choice)
	return s.Investor
}

func (s *SecuritiesFinancingTransactionDetails24) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails26 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *AmountAndDirection51 `xml:"OpngSttlmAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection21 `xml:"TermntnTxAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *SettlementDate10Choice `xml:"OpngSttlmDt"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate4Choice `xml:"TermntnDt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails98 `xml:"SttlmParams,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType35Choice `xml:"RateTp,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails26) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails26) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails26) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails26) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails26) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails26) SetClientTripartyCollateralTransactionIdentification(value string) {
	s.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails26) AddAccountOwner() *PartyIdentification98 {
	s.AccountOwner = new(PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesFinancingTransactionDetails26) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesFinancingTransactionDetails26) AddSafekeepingPlace() *SafeKeepingPlace1 {
	s.SafekeepingPlace = new(SafeKeepingPlace1)
	return s.SafekeepingPlace
}

func (s *SecuritiesFinancingTransactionDetails26) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesFinancingTransactionDetails26) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingTransactionDetails26) AddSettlementQuantity() *Quantity6Choice {
	s.SettlementQuantity = new(Quantity6Choice)
	return s.SettlementQuantity
}

func (s *SecuritiesFinancingTransactionDetails26) AddOpeningSettlementAmount() *AmountAndDirection51 {
	s.OpeningSettlementAmount = new(AmountAndDirection51)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingTransactionDetails26) AddTerminationTransactionAmount() *AmountAndDirection21 {
	s.TerminationTransactionAmount = new(AmountAndDirection21)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails26) AddOpeningSettlementDate() *SettlementDate10Choice {
	s.OpeningSettlementDate = new(SettlementDate10Choice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails26) AddTerminationDate() *TerminationDate4Choice {
	s.TerminationDate = new(TerminationDate4Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails26) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesFinancingTransactionDetails26) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	s.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return s.ExpectedSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails26) AddExpectedValueDate() *DateAndDateTimeChoice {
	s.ExpectedValueDate = new(DateAndDateTimeChoice)
	return s.ExpectedValueDate
}

func (s *SecuritiesFinancingTransactionDetails26) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesFinancingTransactionDetails26) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails26) SetSecuritiesFinancingTransactionType(value string) {
	s.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails26) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails26) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails26) AddSettlementParameters() *SettlementDetails98 {
	s.SettlementParameters = new(SettlementDetails98)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingTransactionDetails26) AddRateType() *RateType35Choice {
	s.RateType = new(RateType35Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails26) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails26) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails26) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails26) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails26) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails26) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails26) AddDeliveringSettlementParties() *SettlementParties36 {
	s.DeliveringSettlementParties = new(SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails26) AddReceivingSettlementParties() *SettlementParties36 {
	s.ReceivingSettlementParties = new(SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails26) AddInvestor() *PartyIdentification99 {
	s.Investor = new(PartyIdentification99)
	return s.Investor
}

func (s *SecuritiesFinancingTransactionDetails26) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails27 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate4Choice `xml:"TermntnDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Earliest date/time at which the call back can take place.
	EarliestCallBackDate *DateAndDateTimeChoice `xml:"EarlstCallBckDt,omitempty"`

	// Date/time at which the commission is calculated.
	CommissionCalculationDate *DateAndDateTimeChoice `xml:"ComssnClctnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType35Choice `xml:"RateTp,omitempty"`

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Revaluation *RevaluationIndicator3Choice `xml:"Rvaltn,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework3Choice `xml:"LglFrmwk,omitempty"`

	// Identifies the computation method of accrued interest of the related financial instrument.
	InterestComputationMethod *InterestComputationMethodFormat4Choice `xml:"IntrstCmptnMtd,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate paid in the context of a securities financing transaction.
	ChargesRate *Rate2 `xml:"ChrgsRate,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Total number of collateral instructions involved in the transaction.
	TotalNumberOfCollateralInstructions *Exact3NumericText `xml:"TtlNbOfCollInstrs,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection21 `xml:"LclBrkrComssn,omitempty"`

	// Principal amount of a trade (for second leg).
	DealAmount *AmountAndDirection21 `xml:"DealAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection21 `xml:"AcrdIntrstAmt,omitempty"`

	// Fixed amount of money that has to be paid (instead of interest) in the case of a recall or at the closing date.
	ForfeitAmount *AmountAndDirection21 `xml:"FrftAmt,omitempty"`

	// Difference between the amount of money of the first leg and the amount of the second leg of the transaction.
	PremiumAmount *AmountAndDirection21 `xml:"PrmAmt,omitempty"`

	// Amount of money to be settled per piece of collateral to terminate the transaction.
	TerminationAmountPerPieceOfCollateral *AmountAndDirection21 `xml:"TermntnAmtPerPcOfColl,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection21 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *Max140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails27) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails27) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails27) AddTerminationDate() *TerminationDate4Choice {
	s.TerminationDate = new(TerminationDate4Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails27) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails27) AddEarliestCallBackDate() *DateAndDateTimeChoice {
	s.EarliestCallBackDate = new(DateAndDateTimeChoice)
	return s.EarliestCallBackDate
}

func (s *SecuritiesFinancingTransactionDetails27) AddCommissionCalculationDate() *DateAndDateTimeChoice {
	s.CommissionCalculationDate = new(DateAndDateTimeChoice)
	return s.CommissionCalculationDate
}

func (s *SecuritiesFinancingTransactionDetails27) AddRateType() *RateType35Choice {
	s.RateType = new(RateType35Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails27) AddRevaluation() *RevaluationIndicator3Choice {
	s.Revaluation = new(RevaluationIndicator3Choice)
	return s.Revaluation
}

func (s *SecuritiesFinancingTransactionDetails27) AddLegalFramework() *LegalFramework3Choice {
	s.LegalFramework = new(LegalFramework3Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails27) AddInterestComputationMethod() *InterestComputationMethodFormat4Choice {
	s.InterestComputationMethod = new(InterestComputationMethodFormat4Choice)
	return s.InterestComputationMethod
}

func (s *SecuritiesFinancingTransactionDetails27) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails27) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails27) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails27) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails27) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails27) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails27) AddChargesRate() *Rate2 {
	s.ChargesRate = new(Rate2)
	return s.ChargesRate
}

func (s *SecuritiesFinancingTransactionDetails27) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails27) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails27) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails27) SetTotalNumberOfCollateralInstructions(value string) {
	s.TotalNumberOfCollateralInstructions = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails27) AddLocalBrokerCommission() *AmountAndDirection21 {
	s.LocalBrokerCommission = new(AmountAndDirection21)
	return s.LocalBrokerCommission
}

func (s *SecuritiesFinancingTransactionDetails27) AddDealAmount() *AmountAndDirection21 {
	s.DealAmount = new(AmountAndDirection21)
	return s.DealAmount
}

func (s *SecuritiesFinancingTransactionDetails27) AddAccruedInterestAmount() *AmountAndDirection21 {
	s.AccruedInterestAmount = new(AmountAndDirection21)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails27) AddForfeitAmount() *AmountAndDirection21 {
	s.ForfeitAmount = new(AmountAndDirection21)
	return s.ForfeitAmount
}

func (s *SecuritiesFinancingTransactionDetails27) AddPremiumAmount() *AmountAndDirection21 {
	s.PremiumAmount = new(AmountAndDirection21)
	return s.PremiumAmount
}

func (s *SecuritiesFinancingTransactionDetails27) AddTerminationAmountPerPieceOfCollateral() *AmountAndDirection21 {
	s.TerminationAmountPerPieceOfCollateral = new(AmountAndDirection21)
	return s.TerminationAmountPerPieceOfCollateral
}

func (s *SecuritiesFinancingTransactionDetails27) AddTerminationTransactionAmount() *AmountAndDirection21 {
	s.TerminationTransactionAmount = new(AmountAndDirection21)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails27) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*Max140Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails28 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate4Choice `xml:"TermntnDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Earliest date/time at which the call back can take place.
	EarliestCallBackDate *DateAndDateTimeChoice `xml:"EarlstCallBckDt,omitempty"`

	// Date/time at which the commission is calculated.
	CommissionCalculationDate *DateAndDateTimeChoice `xml:"ComssnClctnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType35Choice `xml:"RateTp,omitempty"`

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Revaluation *RevaluationIndicator3Choice `xml:"Rvaltn,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework3Choice `xml:"LglFrmwk,omitempty"`

	// Identifies the computation method of accrued interest of the related financial instrument.
	InterestComputationMethod *InterestComputationMethodFormat4Choice `xml:"IntrstCmptnMtd,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate paid in the context of a securities financing transaction.
	ChargesRate *Rate2 `xml:"ChrgsRate,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Total number of collateral instructions involved in the transaction.
	TotalNumberOfCollateralInstructions *Exact3NumericText `xml:"TtlNbOfCollInstrs,omitempty"`

	// Principal amount of a trade (for second leg).
	DealAmount *AmountAndDirection21 `xml:"DealAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection21 `xml:"AcrdIntrstAmt,omitempty"`

	// Fixed amount of money that has to be paid (instead of interest) in the case of a recall or at the closing date.
	ForfeitAmount *AmountAndDirection21 `xml:"FrftAmt,omitempty"`

	// Difference between the amount of money of the first leg and the amount of the second leg of the transaction.
	PremiumAmount *AmountAndDirection21 `xml:"PrmAmt,omitempty"`

	// Amount of money to be settled per piece of collateral to terminate the transaction.
	TerminationAmountPerPieceOfCollateral *AmountAndDirection21 `xml:"TermntnAmtPerPcOfColl,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection21 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *Max140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails28) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails28) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails28) AddTerminationDate() *TerminationDate4Choice {
	s.TerminationDate = new(TerminationDate4Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails28) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails28) AddEarliestCallBackDate() *DateAndDateTimeChoice {
	s.EarliestCallBackDate = new(DateAndDateTimeChoice)
	return s.EarliestCallBackDate
}

func (s *SecuritiesFinancingTransactionDetails28) AddCommissionCalculationDate() *DateAndDateTimeChoice {
	s.CommissionCalculationDate = new(DateAndDateTimeChoice)
	return s.CommissionCalculationDate
}

func (s *SecuritiesFinancingTransactionDetails28) AddRateType() *RateType35Choice {
	s.RateType = new(RateType35Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails28) AddRevaluation() *RevaluationIndicator3Choice {
	s.Revaluation = new(RevaluationIndicator3Choice)
	return s.Revaluation
}

func (s *SecuritiesFinancingTransactionDetails28) AddLegalFramework() *LegalFramework3Choice {
	s.LegalFramework = new(LegalFramework3Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails28) AddInterestComputationMethod() *InterestComputationMethodFormat4Choice {
	s.InterestComputationMethod = new(InterestComputationMethodFormat4Choice)
	return s.InterestComputationMethod
}

func (s *SecuritiesFinancingTransactionDetails28) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails28) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails28) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails28) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails28) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails28) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails28) AddChargesRate() *Rate2 {
	s.ChargesRate = new(Rate2)
	return s.ChargesRate
}

func (s *SecuritiesFinancingTransactionDetails28) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails28) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails28) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails28) SetTotalNumberOfCollateralInstructions(value string) {
	s.TotalNumberOfCollateralInstructions = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails28) AddDealAmount() *AmountAndDirection21 {
	s.DealAmount = new(AmountAndDirection21)
	return s.DealAmount
}

func (s *SecuritiesFinancingTransactionDetails28) AddAccruedInterestAmount() *AmountAndDirection21 {
	s.AccruedInterestAmount = new(AmountAndDirection21)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails28) AddForfeitAmount() *AmountAndDirection21 {
	s.ForfeitAmount = new(AmountAndDirection21)
	return s.ForfeitAmount
}

func (s *SecuritiesFinancingTransactionDetails28) AddPremiumAmount() *AmountAndDirection21 {
	s.PremiumAmount = new(AmountAndDirection21)
	return s.PremiumAmount
}

func (s *SecuritiesFinancingTransactionDetails28) AddTerminationAmountPerPieceOfCollateral() *AmountAndDirection21 {
	s.TerminationAmountPerPieceOfCollateral = new(AmountAndDirection21)
	return s.TerminationAmountPerPieceOfCollateral
}

func (s *SecuritiesFinancingTransactionDetails28) AddTerminationTransactionAmount() *AmountAndDirection21 {
	s.TerminationTransactionAmount = new(AmountAndDirection21)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails28) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*Max140Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails29 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate4Choice `xml:"TermntnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType35Choice `xml:"RateTp,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework3Choice `xml:"LglFrmwk,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection21 `xml:"AcrdIntrstAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection21 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *Max140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails29) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails29) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails29) AddTerminationDate() *TerminationDate4Choice {
	s.TerminationDate = new(TerminationDate4Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails29) AddRateType() *RateType35Choice {
	s.RateType = new(RateType35Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails29) AddLegalFramework() *LegalFramework3Choice {
	s.LegalFramework = new(LegalFramework3Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails29) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails29) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails29) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails29) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails29) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails29) AddAccruedInterestAmount() *AmountAndDirection21 {
	s.AccruedInterestAmount = new(AmountAndDirection21)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails29) AddTerminationTransactionAmount() *AmountAndDirection21 {
	s.TerminationTransactionAmount = new(AmountAndDirection21)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails29) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*Max140Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails3 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate2Choice `xml:"TermntnDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType5Choice `xml:"RateTp,omitempty"`

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Revaluation *RevaluationIndicator1Choice `xml:"Rvaltn,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework1Choice `xml:"LglFrmwk,omitempty"`

	// Identifies the computation method of accrued interest of the related financial instrument.
	InterestComputationMethod *InterestComputationMethodFormat1Choice `xml:"IntrstCmptnMtd,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Total number of collateral instructions involved in the transaction.
	TotalNumberOfCollateralInstructions *Exact3NumericText `xml:"TtlNbOfCollInstrs,omitempty"`

	// Principal amount of a trade (for second leg).
	DealAmount *AmountAndDirection4 `xml:"DealAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection4 `xml:"AcrdIntrstAmt,omitempty"`

	// Fixed amount of money that has to be paid (instead of interest) in the case of a recall or at the closing date.
	ForfeitAmount *AmountAndDirection4 `xml:"FrftAmt,omitempty"`

	// Difference between the amount of money of the first leg and the amount of the second leg of the transaction.
	PremiumAmount *AmountAndDirection4 `xml:"PrmAmt,omitempty"`

	// Amount of money to be settled per piece of collateral to terminate the transaction.
	TerminationAmountPerPieceOfCollateral *AmountAndDirection4 `xml:"TermntnAmtPerPcOfColl,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection4 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *Max140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails3) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails3) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails3) AddTerminationDate() *TerminationDate2Choice {
	s.TerminationDate = new(TerminationDate2Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails3) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails3) AddRateType() *RateType5Choice {
	s.RateType = new(RateType5Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails3) AddRevaluation() *RevaluationIndicator1Choice {
	s.Revaluation = new(RevaluationIndicator1Choice)
	return s.Revaluation
}

func (s *SecuritiesFinancingTransactionDetails3) AddLegalFramework() *LegalFramework1Choice {
	s.LegalFramework = new(LegalFramework1Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails3) AddInterestComputationMethod() *InterestComputationMethodFormat1Choice {
	s.InterestComputationMethod = new(InterestComputationMethodFormat1Choice)
	return s.InterestComputationMethod
}

func (s *SecuritiesFinancingTransactionDetails3) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails3) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails3) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails3) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails3) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails3) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails3) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails3) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails3) SetTotalNumberOfCollateralInstructions(value string) {
	s.TotalNumberOfCollateralInstructions = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails3) AddDealAmount() *AmountAndDirection4 {
	s.DealAmount = new(AmountAndDirection4)
	return s.DealAmount
}

func (s *SecuritiesFinancingTransactionDetails3) AddAccruedInterestAmount() *AmountAndDirection4 {
	s.AccruedInterestAmount = new(AmountAndDirection4)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails3) AddForfeitAmount() *AmountAndDirection4 {
	s.ForfeitAmount = new(AmountAndDirection4)
	return s.ForfeitAmount
}

func (s *SecuritiesFinancingTransactionDetails3) AddPremiumAmount() *AmountAndDirection4 {
	s.PremiumAmount = new(AmountAndDirection4)
	return s.PremiumAmount
}

func (s *SecuritiesFinancingTransactionDetails3) AddTerminationAmountPerPieceOfCollateral() *AmountAndDirection4 {
	s.TerminationAmountPerPieceOfCollateral = new(AmountAndDirection4)
	return s.TerminationAmountPerPieceOfCollateral
}

func (s *SecuritiesFinancingTransactionDetails3) AddTerminationTransactionAmount() *AmountAndDirection4 {
	s.TerminationTransactionAmount = new(AmountAndDirection4)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails3) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*Max140Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails30 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *RestrictedFINXMax16Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *RestrictedFINXMax16Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate5Choice `xml:"TermntnDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Earliest date/time at which the call back can take place.
	EarliestCallBackDate *DateAndDateTimeChoice `xml:"EarlstCallBckDt,omitempty"`

	// Date/time at which the commission is calculated.
	CommissionCalculationDate *DateAndDateTimeChoice `xml:"ComssnClctnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType67Choice `xml:"RateTp,omitempty"`

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Revaluation *RevaluationIndicator4Choice `xml:"Rvaltn,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework4Choice `xml:"LglFrmwk,omitempty"`

	// Identifies the computation method of accrued interest of the related financial instrument.
	InterestComputationMethod *InterestComputationMethodFormat5Choice `xml:"IntrstCmptnMtd,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName2 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate paid in the context of a securities financing transaction.
	ChargesRate *Rate2 `xml:"ChrgsRate,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName2Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Total number of collateral instructions involved in the transaction.
	TotalNumberOfCollateralInstructions *Exact3NumericText `xml:"TtlNbOfCollInstrs,omitempty"`

	// Principal amount of a trade (for second leg).
	DealAmount *AmountAndDirection59 `xml:"DealAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection59 `xml:"AcrdIntrstAmt,omitempty"`

	// Fixed amount of money that has to be paid (instead of interest) in the case of a recall or at the closing date.
	ForfeitAmount *AmountAndDirection59 `xml:"FrftAmt,omitempty"`

	// Difference between the amount of money of the first leg and the amount of the second leg of the transaction.
	PremiumAmount *AmountAndDirection59 `xml:"PrmAmt,omitempty"`

	// Amount of money to be settled per piece of collateral to terminate the transaction.
	TerminationAmountPerPieceOfCollateral *AmountAndDirection59 `xml:"TermntnAmtPerPcOfColl,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection59 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *RestrictedFINXMax140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails30) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails30) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails30) AddTerminationDate() *TerminationDate5Choice {
	s.TerminationDate = new(TerminationDate5Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails30) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails30) AddEarliestCallBackDate() *DateAndDateTimeChoice {
	s.EarliestCallBackDate = new(DateAndDateTimeChoice)
	return s.EarliestCallBackDate
}

func (s *SecuritiesFinancingTransactionDetails30) AddCommissionCalculationDate() *DateAndDateTimeChoice {
	s.CommissionCalculationDate = new(DateAndDateTimeChoice)
	return s.CommissionCalculationDate
}

func (s *SecuritiesFinancingTransactionDetails30) AddRateType() *RateType67Choice {
	s.RateType = new(RateType67Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails30) AddRevaluation() *RevaluationIndicator4Choice {
	s.Revaluation = new(RevaluationIndicator4Choice)
	return s.Revaluation
}

func (s *SecuritiesFinancingTransactionDetails30) AddLegalFramework() *LegalFramework4Choice {
	s.LegalFramework = new(LegalFramework4Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails30) AddInterestComputationMethod() *InterestComputationMethodFormat5Choice {
	s.InterestComputationMethod = new(InterestComputationMethodFormat5Choice)
	return s.InterestComputationMethod
}

func (s *SecuritiesFinancingTransactionDetails30) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails30) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails30) AddVariableRateSupport() *RateName2 {
	s.VariableRateSupport = new(RateName2)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails30) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails30) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails30) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails30) AddChargesRate() *Rate2 {
	s.ChargesRate = new(Rate2)
	return s.ChargesRate
}

func (s *SecuritiesFinancingTransactionDetails30) AddPricingRate() *RateOrName2Choice {
	s.PricingRate = new(RateOrName2Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails30) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails30) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails30) SetTotalNumberOfCollateralInstructions(value string) {
	s.TotalNumberOfCollateralInstructions = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails30) AddDealAmount() *AmountAndDirection59 {
	s.DealAmount = new(AmountAndDirection59)
	return s.DealAmount
}

func (s *SecuritiesFinancingTransactionDetails30) AddAccruedInterestAmount() *AmountAndDirection59 {
	s.AccruedInterestAmount = new(AmountAndDirection59)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails30) AddForfeitAmount() *AmountAndDirection59 {
	s.ForfeitAmount = new(AmountAndDirection59)
	return s.ForfeitAmount
}

func (s *SecuritiesFinancingTransactionDetails30) AddPremiumAmount() *AmountAndDirection59 {
	s.PremiumAmount = new(AmountAndDirection59)
	return s.PremiumAmount
}

func (s *SecuritiesFinancingTransactionDetails30) AddTerminationAmountPerPieceOfCollateral() *AmountAndDirection59 {
	s.TerminationAmountPerPieceOfCollateral = new(AmountAndDirection59)
	return s.TerminationAmountPerPieceOfCollateral
}

func (s *SecuritiesFinancingTransactionDetails30) AddTerminationTransactionAmount() *AmountAndDirection59 {
	s.TerminationTransactionAmount = new(AmountAndDirection59)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails30) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*RestrictedFINXMax140Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails32 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *RestrictedFINXMax16Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *RestrictedFINXMax16Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate5Choice `xml:"TermntnDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Earliest date/time at which the call back can take place.
	EarliestCallBackDate *DateAndDateTimeChoice `xml:"EarlstCallBckDt,omitempty"`

	// Date/time at which the commission is calculated.
	CommissionCalculationDate *DateAndDateTimeChoice `xml:"ComssnClctnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType67Choice `xml:"RateTp,omitempty"`

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Revaluation *RevaluationIndicator4Choice `xml:"Rvaltn,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework4Choice `xml:"LglFrmwk,omitempty"`

	// Identifies the computation method of accrued interest of the related financial instrument.
	InterestComputationMethod *InterestComputationMethodFormat5Choice `xml:"IntrstCmptnMtd,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName2 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate paid in the context of a securities financing transaction.
	ChargesRate *Rate2 `xml:"ChrgsRate,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName2Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Total number of collateral instructions involved in the transaction.
	TotalNumberOfCollateralInstructions *Exact3NumericText `xml:"TtlNbOfCollInstrs,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection59 `xml:"LclBrkrComssn,omitempty"`

	// Principal amount of a trade (for second leg).
	DealAmount *AmountAndDirection59 `xml:"DealAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection59 `xml:"AcrdIntrstAmt,omitempty"`

	// Fixed amount of money that has to be paid (instead of interest) in the case of a recall or at the closing date.
	ForfeitAmount *AmountAndDirection59 `xml:"FrftAmt,omitempty"`

	// Difference between the amount of money of the first leg and the amount of the second leg of the transaction.
	PremiumAmount *AmountAndDirection59 `xml:"PrmAmt,omitempty"`

	// Amount of money to be settled per piece of collateral to terminate the transaction.
	TerminationAmountPerPieceOfCollateral *AmountAndDirection59 `xml:"TermntnAmtPerPcOfColl,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection59 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *RestrictedFINXMax140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails32) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails32) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails32) AddTerminationDate() *TerminationDate5Choice {
	s.TerminationDate = new(TerminationDate5Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails32) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails32) AddEarliestCallBackDate() *DateAndDateTimeChoice {
	s.EarliestCallBackDate = new(DateAndDateTimeChoice)
	return s.EarliestCallBackDate
}

func (s *SecuritiesFinancingTransactionDetails32) AddCommissionCalculationDate() *DateAndDateTimeChoice {
	s.CommissionCalculationDate = new(DateAndDateTimeChoice)
	return s.CommissionCalculationDate
}

func (s *SecuritiesFinancingTransactionDetails32) AddRateType() *RateType67Choice {
	s.RateType = new(RateType67Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails32) AddRevaluation() *RevaluationIndicator4Choice {
	s.Revaluation = new(RevaluationIndicator4Choice)
	return s.Revaluation
}

func (s *SecuritiesFinancingTransactionDetails32) AddLegalFramework() *LegalFramework4Choice {
	s.LegalFramework = new(LegalFramework4Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails32) AddInterestComputationMethod() *InterestComputationMethodFormat5Choice {
	s.InterestComputationMethod = new(InterestComputationMethodFormat5Choice)
	return s.InterestComputationMethod
}

func (s *SecuritiesFinancingTransactionDetails32) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails32) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails32) AddVariableRateSupport() *RateName2 {
	s.VariableRateSupport = new(RateName2)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails32) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails32) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails32) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails32) AddChargesRate() *Rate2 {
	s.ChargesRate = new(Rate2)
	return s.ChargesRate
}

func (s *SecuritiesFinancingTransactionDetails32) AddPricingRate() *RateOrName2Choice {
	s.PricingRate = new(RateOrName2Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails32) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails32) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails32) SetTotalNumberOfCollateralInstructions(value string) {
	s.TotalNumberOfCollateralInstructions = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails32) AddLocalBrokerCommission() *AmountAndDirection59 {
	s.LocalBrokerCommission = new(AmountAndDirection59)
	return s.LocalBrokerCommission
}

func (s *SecuritiesFinancingTransactionDetails32) AddDealAmount() *AmountAndDirection59 {
	s.DealAmount = new(AmountAndDirection59)
	return s.DealAmount
}

func (s *SecuritiesFinancingTransactionDetails32) AddAccruedInterestAmount() *AmountAndDirection59 {
	s.AccruedInterestAmount = new(AmountAndDirection59)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails32) AddForfeitAmount() *AmountAndDirection59 {
	s.ForfeitAmount = new(AmountAndDirection59)
	return s.ForfeitAmount
}

func (s *SecuritiesFinancingTransactionDetails32) AddPremiumAmount() *AmountAndDirection59 {
	s.PremiumAmount = new(AmountAndDirection59)
	return s.PremiumAmount
}

func (s *SecuritiesFinancingTransactionDetails32) AddTerminationAmountPerPieceOfCollateral() *AmountAndDirection59 {
	s.TerminationAmountPerPieceOfCollateral = new(AmountAndDirection59)
	return s.TerminationAmountPerPieceOfCollateral
}

func (s *SecuritiesFinancingTransactionDetails32) AddTerminationTransactionAmount() *AmountAndDirection59 {
	s.TerminationTransactionAmount = new(AmountAndDirection59)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails32) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*RestrictedFINXMax140Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails33 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *RestrictedFINXMax16Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *RestrictedFINXMax16Text `xml:"ClsgLegId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *AmountAndDirection67 `xml:"OpngSttlmAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection59 `xml:"TermntnTxAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *SettlementDate15Choice `xml:"OpngSttlmDt"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate5Choice `xml:"TermntnDt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails107 `xml:"SttlmParams,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType67Choice `xml:"RateTp,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName2 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName2Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails33) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetPoolIdentification(value string) {
	s.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetClientTripartyCollateralTransactionIdentification(value string) {
	s.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) AddAccountOwner() *PartyIdentification109 {
	s.AccountOwner = new(PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesFinancingTransactionDetails33) AddSafekeepingAccount() *SecuritiesAccount30 {
	s.SafekeepingAccount = new(SecuritiesAccount30)
	return s.SafekeepingAccount
}

func (s *SecuritiesFinancingTransactionDetails33) AddSafekeepingPlace() *SafeKeepingPlace2 {
	s.SafekeepingPlace = new(SafeKeepingPlace2)
	return s.SafekeepingPlace
}

func (s *SecuritiesFinancingTransactionDetails33) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesFinancingTransactionDetails33) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingTransactionDetails33) AddSettlementQuantity() *Quantity10Choice {
	s.SettlementQuantity = new(Quantity10Choice)
	return s.SettlementQuantity
}

func (s *SecuritiesFinancingTransactionDetails33) AddOpeningSettlementAmount() *AmountAndDirection67 {
	s.OpeningSettlementAmount = new(AmountAndDirection67)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingTransactionDetails33) AddTerminationTransactionAmount() *AmountAndDirection59 {
	s.TerminationTransactionAmount = new(AmountAndDirection59)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails33) AddOpeningSettlementDate() *SettlementDate15Choice {
	s.OpeningSettlementDate = new(SettlementDate15Choice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails33) AddTerminationDate() *TerminationDate5Choice {
	s.TerminationDate = new(TerminationDate5Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails33) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesFinancingTransactionDetails33) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	s.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return s.ExpectedSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails33) AddExpectedValueDate() *DateAndDateTimeChoice {
	s.ExpectedValueDate = new(DateAndDateTimeChoice)
	return s.ExpectedValueDate
}

func (s *SecuritiesFinancingTransactionDetails33) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesFinancingTransactionDetails33) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails33) SetSecuritiesFinancingTransactionType(value string) {
	s.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) AddSettlementParameters() *SettlementDetails107 {
	s.SettlementParameters = new(SettlementDetails107)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingTransactionDetails33) AddRateType() *RateType67Choice {
	s.RateType = new(RateType67Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails33) AddVariableRateSupport() *RateName2 {
	s.VariableRateSupport = new(RateName2)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails33) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails33) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails33) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails33) AddPricingRate() *RateOrName2Choice {
	s.PricingRate = new(RateOrName2Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails33) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails33) AddDeliveringSettlementParties() *SettlementParties44 {
	s.DeliveringSettlementParties = new(SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails33) AddReceivingSettlementParties() *SettlementParties44 {
	s.ReceivingSettlementParties = new(SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails33) AddInvestor() *PartyIdentification110 {
	s.Investor = new(PartyIdentification110)
	return s.Investor
}

func (s *SecuritiesFinancingTransactionDetails33) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails34 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *RestrictedFINXMax16Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *RestrictedFINXMax16Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate5Choice `xml:"TermntnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType67Choice `xml:"RateTp,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework4Choice `xml:"LglFrmwk,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName2 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection59 `xml:"AcrdIntrstAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection59 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *RestrictedFINXMax140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails34) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails34) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails34) AddTerminationDate() *TerminationDate5Choice {
	s.TerminationDate = new(TerminationDate5Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails34) AddRateType() *RateType67Choice {
	s.RateType = new(RateType67Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails34) AddLegalFramework() *LegalFramework4Choice {
	s.LegalFramework = new(LegalFramework4Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails34) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails34) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails34) AddVariableRateSupport() *RateName2 {
	s.VariableRateSupport = new(RateName2)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails34) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails34) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails34) AddAccruedInterestAmount() *AmountAndDirection59 {
	s.AccruedInterestAmount = new(AmountAndDirection59)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails34) AddTerminationTransactionAmount() *AmountAndDirection59 {
	s.TerminationTransactionAmount = new(AmountAndDirection59)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails34) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*RestrictedFINXMax140Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails35 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty-agent's/service-provider's point of view.
	TripartyAgentServiceProviderCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtSvcPrvdrCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *AmountAndDirection51 `xml:"OpngSttlmAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection21 `xml:"TermntnTxAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *SettlementDate10Choice `xml:"OpngSttlmDt"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate4Choice `xml:"TermntnDt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType2Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails98 `xml:"SttlmParams,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType35Choice `xml:"RateTp,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails35) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	s.TripartyAgentServiceProviderCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetClientTripartyCollateralTransactionIdentification(value string) {
	s.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) AddAccountOwner() *PartyIdentification98 {
	s.AccountOwner = new(PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesFinancingTransactionDetails35) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesFinancingTransactionDetails35) AddSafekeepingPlace() *SafeKeepingPlace1 {
	s.SafekeepingPlace = new(SafeKeepingPlace1)
	return s.SafekeepingPlace
}

func (s *SecuritiesFinancingTransactionDetails35) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesFinancingTransactionDetails35) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingTransactionDetails35) AddSettlementQuantity() *Quantity6Choice {
	s.SettlementQuantity = new(Quantity6Choice)
	return s.SettlementQuantity
}

func (s *SecuritiesFinancingTransactionDetails35) AddOpeningSettlementAmount() *AmountAndDirection51 {
	s.OpeningSettlementAmount = new(AmountAndDirection51)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingTransactionDetails35) AddTerminationTransactionAmount() *AmountAndDirection21 {
	s.TerminationTransactionAmount = new(AmountAndDirection21)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails35) AddOpeningSettlementDate() *SettlementDate10Choice {
	s.OpeningSettlementDate = new(SettlementDate10Choice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails35) AddTerminationDate() *TerminationDate4Choice {
	s.TerminationDate = new(TerminationDate4Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails35) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesFinancingTransactionDetails35) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	s.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return s.ExpectedSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails35) AddExpectedValueDate() *DateAndDateTimeChoice {
	s.ExpectedValueDate = new(DateAndDateTimeChoice)
	return s.ExpectedValueDate
}

func (s *SecuritiesFinancingTransactionDetails35) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesFinancingTransactionDetails35) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails35) SetSecuritiesFinancingTransactionType(value string) {
	s.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) AddSettlementParameters() *SettlementDetails98 {
	s.SettlementParameters = new(SettlementDetails98)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingTransactionDetails35) AddRateType() *RateType35Choice {
	s.RateType = new(RateType35Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails35) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails35) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails35) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails35) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails35) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails35) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails35) AddDeliveringSettlementParties() *SettlementParties36 {
	s.DeliveringSettlementParties = new(SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails35) AddReceivingSettlementParties() *SettlementParties36 {
	s.ReceivingSettlementParties = new(SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails35) AddInvestor() *PartyIdentification99 {
	s.Investor = new(PartyIdentification99)
	return s.Investor
}

func (s *SecuritiesFinancingTransactionDetails35) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails36 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *RestrictedFINXMax16Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *RestrictedFINXMax16Text `xml:"ClsgLegId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty-agent's/service-provider's point of view.
	TripartyAgentServiceProviderCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtSvcPrvdrCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *AmountAndDirection67 `xml:"OpngSttlmAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection59 `xml:"TermntnTxAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *SettlementDate15Choice `xml:"OpngSttlmDt"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate5Choice `xml:"TermntnDt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType2Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails107 `xml:"SttlmParams,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType67Choice `xml:"RateTp,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName2 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName2Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails36) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails36) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails36) SetPoolIdentification(value string) {
	s.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails36) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails36) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	s.TripartyAgentServiceProviderCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails36) SetClientTripartyCollateralTransactionIdentification(value string) {
	s.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails36) AddAccountOwner() *PartyIdentification119 {
	s.AccountOwner = new(PartyIdentification119)
	return s.AccountOwner
}

func (s *SecuritiesFinancingTransactionDetails36) AddSafekeepingAccount() *SecuritiesAccount30 {
	s.SafekeepingAccount = new(SecuritiesAccount30)
	return s.SafekeepingAccount
}

func (s *SecuritiesFinancingTransactionDetails36) AddSafekeepingPlace() *SafeKeepingPlace2 {
	s.SafekeepingPlace = new(SafeKeepingPlace2)
	return s.SafekeepingPlace
}

func (s *SecuritiesFinancingTransactionDetails36) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesFinancingTransactionDetails36) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingTransactionDetails36) AddSettlementQuantity() *Quantity10Choice {
	s.SettlementQuantity = new(Quantity10Choice)
	return s.SettlementQuantity
}

func (s *SecuritiesFinancingTransactionDetails36) AddOpeningSettlementAmount() *AmountAndDirection67 {
	s.OpeningSettlementAmount = new(AmountAndDirection67)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingTransactionDetails36) AddTerminationTransactionAmount() *AmountAndDirection59 {
	s.TerminationTransactionAmount = new(AmountAndDirection59)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails36) AddOpeningSettlementDate() *SettlementDate15Choice {
	s.OpeningSettlementDate = new(SettlementDate15Choice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails36) AddTerminationDate() *TerminationDate5Choice {
	s.TerminationDate = new(TerminationDate5Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails36) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesFinancingTransactionDetails36) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	s.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return s.ExpectedSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails36) AddExpectedValueDate() *DateAndDateTimeChoice {
	s.ExpectedValueDate = new(DateAndDateTimeChoice)
	return s.ExpectedValueDate
}

func (s *SecuritiesFinancingTransactionDetails36) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesFinancingTransactionDetails36) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails36) SetSecuritiesFinancingTransactionType(value string) {
	s.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails36) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails36) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails36) AddSettlementParameters() *SettlementDetails107 {
	s.SettlementParameters = new(SettlementDetails107)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingTransactionDetails36) AddRateType() *RateType67Choice {
	s.RateType = new(RateType67Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails36) AddVariableRateSupport() *RateName2 {
	s.VariableRateSupport = new(RateName2)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails36) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails36) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails36) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails36) AddPricingRate() *RateOrName2Choice {
	s.PricingRate = new(RateOrName2Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails36) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails36) AddDeliveringSettlementParties() *SettlementParties44 {
	s.DeliveringSettlementParties = new(SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails36) AddReceivingSettlementParties() *SettlementParties44 {
	s.ReceivingSettlementParties = new(SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails36) AddInvestor() *PartyIdentification110 {
	s.Investor = new(PartyIdentification110)
	return s.Investor
}

func (s *SecuritiesFinancingTransactionDetails36) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails7 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate2Choice `xml:"TermntnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType5Choice `xml:"RateTp,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework1Choice `xml:"LglFrmwk,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection4 `xml:"AcrdIntrstAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection4 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *Max140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails7) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails7) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails7) AddTerminationDate() *TerminationDate2Choice {
	s.TerminationDate = new(TerminationDate2Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails7) AddRateType() *RateType5Choice {
	s.RateType = new(RateType5Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails7) AddLegalFramework() *LegalFramework1Choice {
	s.LegalFramework = new(LegalFramework1Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails7) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails7) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails7) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails7) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails7) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails7) AddAccruedInterestAmount() *AmountAndDirection4 {
	s.AccruedInterestAmount = new(AmountAndDirection4)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails7) AddTerminationTransactionAmount() *AmountAndDirection4 {
	s.TerminationTransactionAmount = new(AmountAndDirection4)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails7) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*Max140Text)(&value)
}

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails8 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *AmountAndDirection8 `xml:"OpngSttlmAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection4 `xml:"TermntnTxAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *SettlementDate2Choice `xml:"OpngSttlmDt"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate2Choice `xml:"TermntnDt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails29 `xml:"SttlmParams,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType5Choice `xml:"RateTp,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails8) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetClientTripartyCollateralTransactionIdentification(value string) {
	s.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesFinancingTransactionDetails8) AddSafekeepingAccount() *SecuritiesAccount13 {
	s.SafekeepingAccount = new(SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesFinancingTransactionDetails8) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesFinancingTransactionDetails8) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesFinancingTransactionDetails8) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingTransactionDetails8) AddSettlementQuantity() *Quantity6Choice {
	s.SettlementQuantity = new(Quantity6Choice)
	return s.SettlementQuantity
}

func (s *SecuritiesFinancingTransactionDetails8) AddOpeningSettlementAmount() *AmountAndDirection8 {
	s.OpeningSettlementAmount = new(AmountAndDirection8)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingTransactionDetails8) AddTerminationTransactionAmount() *AmountAndDirection4 {
	s.TerminationTransactionAmount = new(AmountAndDirection4)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails8) AddOpeningSettlementDate() *SettlementDate2Choice {
	s.OpeningSettlementDate = new(SettlementDate2Choice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails8) AddTerminationDate() *TerminationDate2Choice {
	s.TerminationDate = new(TerminationDate2Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails8) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesFinancingTransactionDetails8) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	s.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return s.ExpectedSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails8) AddExpectedValueDate() *DateAndDateTimeChoice {
	s.ExpectedValueDate = new(DateAndDateTimeChoice)
	return s.ExpectedValueDate
}

func (s *SecuritiesFinancingTransactionDetails8) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesFinancingTransactionDetails8) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails8) SetSecuritiesFinancingTransactionType(value string) {
	s.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) AddSettlementParameters() *SettlementDetails29 {
	s.SettlementParameters = new(SettlementDetails29)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingTransactionDetails8) AddRateType() *RateType5Choice {
	s.RateType = new(RateType5Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails8) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails8) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails8) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails8) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails8) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails8) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails8) AddDeliveringSettlementParties() *SettlementParties11 {
	s.DeliveringSettlementParties = new(SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails8) AddReceivingSettlementParties() *SettlementParties11 {
	s.ReceivingSettlementParties = new(SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails8) AddInvestor() *PartyIdentification37Choice {
	s.Investor = new(PartyIdentification37Choice)
	return s.Investor
}

func (s *SecuritiesFinancingTransactionDetails8) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}
