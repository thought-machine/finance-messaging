package iso20022

// An event determined by a corporation's board of directors, that changes the existing corporate capital structure or financial condition.
type CorporateAction1 struct {

	// Specifies the code of corporate action event, in free-text format.
	Code *Max35Text `xml:"Cd,omitempty"`

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Proprietary corporate action event information.
	Proprietary *Max35Text `xml:"Prtry,omitempty"`
}

func (c *CorporateAction1) SetCode(value string) {
	c.Code = (*Max35Text)(&value)
}

func (c *CorporateAction1) SetNumber(value string) {
	c.Number = (*Max35Text)(&value)
}

func (c *CorporateAction1) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}

// Provides information about the corporate action event.
type CorporateAction10 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate27 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action event.
	PeriodDetails *CorporateActionPeriod8 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action event.
	RateAndAmountDetails *CorporateActionRate35 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action event.
	PriceDetails *CorporateActionPrice17 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action.
	SecuritiesQuantity *CorporateActionQuantity5 `xml:"SctiesQty,omitempty"`

	// Number of days used for calculating the accrued interest amount.
	InterestAccruedNumberOfDays *Max3Number `xml:"IntrstAcrdNbOfDays,omitempty"`

	// Number of the coupon attached/associated with a security.
	CouponNumber []*IdentificationFormat1Choice `xml:"CpnNb,omitempty"`

	// Indicates whether certification/breakdown is required.
	// Yes = certification required.
	// No = no certification required.
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether there is restrictions apply to the corporate action event or not.
	// Yes = There is restrictions.
	// No = There is no restrictions.
	RestrictionIndicator *YesNoIndicator `xml:"RstrctnInd,omitempty"`

	// Indicates whether the holder is entitled to accrued interest.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Indicates whether a letter of guaranteed delivery can be submitted in order to participate in the offer on full eligible position. It is not intended for use in situations arising from failed or late trades.
	LetterOfGuaranteedDeliveryIndicator *YesNoIndicator `xml:"LttrOfGrntedDlvryInd,omitempty"`

	// Specifies the conditions in which a dividend is paid.
	DividendType *DividendTypeFormat3Choice `xml:"DvddTp,omitempty"`

	// Specifies the conversion type of an instrument.
	ConversionType *ConversionTypeFormat1Choice `xml:"ConvsTp,omitempty"`

	// Specifies the conditions in which the payment of the proceeds occurs.
	PaymentOccurrenceType *DistributionTypeFormat3Choice `xml:"PmtOcrncTp,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat3Choice `xml:"OfferTp,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat1Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage []*CorporateActionEventStageFormat3Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of the additional business process linked to a corporate action event such as a claim compensation or tax refund.
	AdditionalBusinessProcessIndicator []*AdditionalBusinessProcessFormat1Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of change announced.
	ChangeType []*CorporateActionChangeTypeFormat1Choice `xml:"ChngTp,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat9Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies whether the capital gain is in the scope of the EU Savings directive for the income realised upon the sale, refund or redemption of shares and units (...) (Article 6(1d)).
	CapitalGainInOutIndicator *CapitalGainFormat1Choice `xml:"CptlGnInOutInd,omitempty"`

	// Specifies whether the financial instrument calculates the taxable income per dividend/taxable income per share.
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculatedFormat1Choice `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies the effect on the holdings of electing a corporate action option.
	ElectionType *ElectionTypeFormat1Choice `xml:"ElctnTp,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat1Choice `xml:"LtryTp,omitempty"`

	// Specifies the certification format required, this is, physical or electronic format.
	CertificationType *CertificationTypeFormat1Choice `xml:"CertfctnTp,omitempty"`

	// New company's place of incorporation.
	NewPlaceOfIncorporation *Max350Text `xml:"NewPlcOfIncorprtn,omitempty"`

	// Provides additional information. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalInformation *CorporateActionNarrative24 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateAction10) AddDateDetails() *CorporateActionDate27 {
	c.DateDetails = new(CorporateActionDate27)
	return c.DateDetails
}

func (c *CorporateAction10) AddPeriodDetails() *CorporateActionPeriod8 {
	c.PeriodDetails = new(CorporateActionPeriod8)
	return c.PeriodDetails
}

func (c *CorporateAction10) AddRateAndAmountDetails() *CorporateActionRate35 {
	c.RateAndAmountDetails = new(CorporateActionRate35)
	return c.RateAndAmountDetails
}

func (c *CorporateAction10) AddPriceDetails() *CorporateActionPrice17 {
	c.PriceDetails = new(CorporateActionPrice17)
	return c.PriceDetails
}

func (c *CorporateAction10) AddSecuritiesQuantity() *CorporateActionQuantity5 {
	c.SecuritiesQuantity = new(CorporateActionQuantity5)
	return c.SecuritiesQuantity
}

func (c *CorporateAction10) SetInterestAccruedNumberOfDays(value string) {
	c.InterestAccruedNumberOfDays = (*Max3Number)(&value)
}

func (c *CorporateAction10) AddCouponNumber() *IdentificationFormat1Choice {
	newValue := new(IdentificationFormat1Choice)
	c.CouponNumber = append(c.CouponNumber, newValue)
	return newValue
}

func (c *CorporateAction10) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction10) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction10) SetRestrictionIndicator(value string) {
	c.RestrictionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction10) SetAccruedInterestIndicator(value string) {
	c.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction10) SetLetterOfGuaranteedDeliveryIndicator(value string) {
	c.LetterOfGuaranteedDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction10) AddDividendType() *DividendTypeFormat3Choice {
	c.DividendType = new(DividendTypeFormat3Choice)
	return c.DividendType
}

func (c *CorporateAction10) AddConversionType() *ConversionTypeFormat1Choice {
	c.ConversionType = new(ConversionTypeFormat1Choice)
	return c.ConversionType
}

func (c *CorporateAction10) AddPaymentOccurrenceType() *DistributionTypeFormat3Choice {
	c.PaymentOccurrenceType = new(DistributionTypeFormat3Choice)
	return c.PaymentOccurrenceType
}

func (c *CorporateAction10) AddOfferType() *OfferTypeFormat3Choice {
	newValue := new(OfferTypeFormat3Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateAction10) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat1Choice {
	c.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat1Choice)
	return c.RenounceableEntitlementStatusType
}

func (c *CorporateAction10) AddEventStage() *CorporateActionEventStageFormat3Choice {
	newValue := new(CorporateActionEventStageFormat3Choice)
	c.EventStage = append(c.EventStage, newValue)
	return newValue
}

func (c *CorporateAction10) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat1Choice {
	newValue := new(AdditionalBusinessProcessFormat1Choice)
	c.AdditionalBusinessProcessIndicator = append(c.AdditionalBusinessProcessIndicator, newValue)
	return newValue
}

func (c *CorporateAction10) AddChangeType() *CorporateActionChangeTypeFormat1Choice {
	newValue := new(CorporateActionChangeTypeFormat1Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateAction10) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat9Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat9Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateAction10) AddCapitalGainInOutIndicator() *CapitalGainFormat1Choice {
	c.CapitalGainInOutIndicator = new(CapitalGainFormat1Choice)
	return c.CapitalGainInOutIndicator
}

func (c *CorporateAction10) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculatedFormat1Choice {
	c.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculatedFormat1Choice)
	return c.TaxableIncomePerShareCalculated
}

func (c *CorporateAction10) AddElectionType() *ElectionTypeFormat1Choice {
	c.ElectionType = new(ElectionTypeFormat1Choice)
	return c.ElectionType
}

func (c *CorporateAction10) AddLotteryType() *LotteryTypeFormat1Choice {
	c.LotteryType = new(LotteryTypeFormat1Choice)
	return c.LotteryType
}

func (c *CorporateAction10) AddCertificationType() *CertificationTypeFormat1Choice {
	c.CertificationType = new(CertificationTypeFormat1Choice)
	return c.CertificationType
}

func (c *CorporateAction10) SetNewPlaceOfIncorporation(value string) {
	c.NewPlaceOfIncorporation = (*Max350Text)(&value)
}

func (c *CorporateAction10) AddAdditionalInformation() *CorporateActionNarrative24 {
	c.AdditionalInformation = new(CorporateActionNarrative24)
	return c.AdditionalInformation
}

// Provides information about the corporate action event.
type CorporateAction12 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate27 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action event.
	PeriodDetails *CorporateActionPeriod10 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action event.
	RateAndAmountDetails *CorporateActionRate43 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action event.
	PriceDetails *CorporateActionPrice42 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action.
	SecuritiesQuantity *CorporateActionQuantity5 `xml:"SctiesQty,omitempty"`

	// Number of days used for calculating the accrued interest amount.
	InterestAccruedNumberOfDays *Max3Number `xml:"IntrstAcrdNbOfDays,omitempty"`

	// Number of the coupon attached/associated with a security.
	CouponNumber []*IdentificationFormat1Choice `xml:"CpnNb,omitempty"`

	// Indicates whether certification/breakdown is required.
	// Yes = certification required.
	// No = no certification required.
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether there is restrictions apply to the corporate action event or not.
	// Yes = There is restrictions.
	// No = There is no restrictions.
	RestrictionIndicator *YesNoIndicator `xml:"RstrctnInd,omitempty"`

	// Indicates whether the holder is entitled to accrued interest.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Indicates whether a letter of guaranteed delivery can be submitted in order to participate in the offer on full eligible position. It is not intended for use in situations arising from failed or late trades.
	LetterOfGuaranteedDeliveryIndicator *YesNoIndicator `xml:"LttrOfGrntedDlvryInd,omitempty"`

	// Specifies the conditions in which a dividend is paid.
	DividendType *DividendTypeFormat5Choice `xml:"DvddTp,omitempty"`

	// Specifies the conversion type of an instrument.
	ConversionType *ConversionTypeFormat1Choice `xml:"ConvsTp,omitempty"`

	// Specifies the conditions in which the instructions and/or payment of the proceeds occurs.
	OccurrenceType *DistributionTypeFormat4Choice `xml:"OcrncTp,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat5Choice `xml:"OfferTp,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat1Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage []*CorporateActionEventStageFormat5Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of the additional business process linked to a corporate action event such as a claim compensation or tax refund.
	AdditionalBusinessProcessIndicator []*AdditionalBusinessProcessFormat1Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of change announced.
	ChangeType []*CorporateActionChangeTypeFormat1Choice `xml:"ChngTp,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat9Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies whether the capital gain is in the scope of the EU Savings directive for the income realised upon the sale, refund or redemption of shares and units (...) (Article 6(1d)).
	CapitalGainInOutIndicator *CapitalGainFormat1Choice `xml:"CptlGnInOutInd,omitempty"`

	// Specifies whether the financial instrument calculates the taxable income per dividend/taxable income per share.
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculatedFormat1Choice `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies the effect on the holdings of electing a corporate action option.
	ElectionType *ElectionTypeFormat1Choice `xml:"ElctnTp,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat1Choice `xml:"LtryTp,omitempty"`

	// Specifies the certification format required, this is, physical or electronic format.
	CertificationType *CertificationTypeFormat1Choice `xml:"CertfctnTp,omitempty"`

	// Specifies the type of consent announced.
	ConsentType *ConsentTypeFormat1Choice `xml:"CnsntTp,omitempty"`

	// Specifies the type of information event.
	InformationType *InformationTypeFormat1Choice `xml:"InfTp,omitempty"`

	// New company's place of incorporation.
	NewPlaceOfIncorporation *Max350Text `xml:"NewPlcOfIncorprtn,omitempty"`

	// Provides additional information. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalInformation *CorporateActionNarrative24 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateAction12) AddDateDetails() *CorporateActionDate27 {
	c.DateDetails = new(CorporateActionDate27)
	return c.DateDetails
}

func (c *CorporateAction12) AddPeriodDetails() *CorporateActionPeriod10 {
	c.PeriodDetails = new(CorporateActionPeriod10)
	return c.PeriodDetails
}

func (c *CorporateAction12) AddRateAndAmountDetails() *CorporateActionRate43 {
	c.RateAndAmountDetails = new(CorporateActionRate43)
	return c.RateAndAmountDetails
}

func (c *CorporateAction12) AddPriceDetails() *CorporateActionPrice42 {
	c.PriceDetails = new(CorporateActionPrice42)
	return c.PriceDetails
}

func (c *CorporateAction12) AddSecuritiesQuantity() *CorporateActionQuantity5 {
	c.SecuritiesQuantity = new(CorporateActionQuantity5)
	return c.SecuritiesQuantity
}

func (c *CorporateAction12) SetInterestAccruedNumberOfDays(value string) {
	c.InterestAccruedNumberOfDays = (*Max3Number)(&value)
}

func (c *CorporateAction12) AddCouponNumber() *IdentificationFormat1Choice {
	newValue := new(IdentificationFormat1Choice)
	c.CouponNumber = append(c.CouponNumber, newValue)
	return newValue
}

func (c *CorporateAction12) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction12) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction12) SetRestrictionIndicator(value string) {
	c.RestrictionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction12) SetAccruedInterestIndicator(value string) {
	c.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction12) SetLetterOfGuaranteedDeliveryIndicator(value string) {
	c.LetterOfGuaranteedDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction12) AddDividendType() *DividendTypeFormat5Choice {
	c.DividendType = new(DividendTypeFormat5Choice)
	return c.DividendType
}

func (c *CorporateAction12) AddConversionType() *ConversionTypeFormat1Choice {
	c.ConversionType = new(ConversionTypeFormat1Choice)
	return c.ConversionType
}

func (c *CorporateAction12) AddOccurrenceType() *DistributionTypeFormat4Choice {
	c.OccurrenceType = new(DistributionTypeFormat4Choice)
	return c.OccurrenceType
}

func (c *CorporateAction12) AddOfferType() *OfferTypeFormat5Choice {
	newValue := new(OfferTypeFormat5Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateAction12) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat1Choice {
	c.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat1Choice)
	return c.RenounceableEntitlementStatusType
}

func (c *CorporateAction12) AddEventStage() *CorporateActionEventStageFormat5Choice {
	newValue := new(CorporateActionEventStageFormat5Choice)
	c.EventStage = append(c.EventStage, newValue)
	return newValue
}

func (c *CorporateAction12) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat1Choice {
	newValue := new(AdditionalBusinessProcessFormat1Choice)
	c.AdditionalBusinessProcessIndicator = append(c.AdditionalBusinessProcessIndicator, newValue)
	return newValue
}

func (c *CorporateAction12) AddChangeType() *CorporateActionChangeTypeFormat1Choice {
	newValue := new(CorporateActionChangeTypeFormat1Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateAction12) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat9Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat9Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateAction12) AddCapitalGainInOutIndicator() *CapitalGainFormat1Choice {
	c.CapitalGainInOutIndicator = new(CapitalGainFormat1Choice)
	return c.CapitalGainInOutIndicator
}

func (c *CorporateAction12) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculatedFormat1Choice {
	c.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculatedFormat1Choice)
	return c.TaxableIncomePerShareCalculated
}

func (c *CorporateAction12) AddElectionType() *ElectionTypeFormat1Choice {
	c.ElectionType = new(ElectionTypeFormat1Choice)
	return c.ElectionType
}

func (c *CorporateAction12) AddLotteryType() *LotteryTypeFormat1Choice {
	c.LotteryType = new(LotteryTypeFormat1Choice)
	return c.LotteryType
}

func (c *CorporateAction12) AddCertificationType() *CertificationTypeFormat1Choice {
	c.CertificationType = new(CertificationTypeFormat1Choice)
	return c.CertificationType
}

func (c *CorporateAction12) AddConsentType() *ConsentTypeFormat1Choice {
	c.ConsentType = new(ConsentTypeFormat1Choice)
	return c.ConsentType
}

func (c *CorporateAction12) AddInformationType() *InformationTypeFormat1Choice {
	c.InformationType = new(InformationTypeFormat1Choice)
	return c.InformationType
}

func (c *CorporateAction12) SetNewPlaceOfIncorporation(value string) {
	c.NewPlaceOfIncorporation = (*Max350Text)(&value)
}

func (c *CorporateAction12) AddAdditionalInformation() *CorporateActionNarrative24 {
	c.AdditionalInformation = new(CorporateActionNarrative24)
	return c.AdditionalInformation
}

// Provides information about the corporate action event.
type CorporateAction13 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate30 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat6Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat1Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction13) AddDateDetails() *CorporateActionDate30 {
	c.DateDetails = new(CorporateActionDate30)
	return c.DateDetails
}

func (c *CorporateAction13) AddEventStage() *CorporateActionEventStageFormat6Choice {
	c.EventStage = new(CorporateActionEventStageFormat6Choice)
	return c.EventStage
}

func (c *CorporateAction13) AddLotteryType() *LotteryTypeFormat1Choice {
	c.LotteryType = new(LotteryTypeFormat1Choice)
	return c.LotteryType
}

// Provides information about the corporate action event.
type CorporateAction14 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate30 `xml:"DtDtls,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat1Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction14) AddDateDetails() *CorporateActionDate30 {
	c.DateDetails = new(CorporateActionDate30)
	return c.DateDetails
}

func (c *CorporateAction14) AddLotteryType() *LotteryTypeFormat1Choice {
	c.LotteryType = new(LotteryTypeFormat1Choice)
	return c.LotteryType
}

// Provides information about the CA event.
type CorporateAction2 struct {

	// Stage in the corporate action event life cycle.
	EventStage []*CorporateActionEventStage1FormatChoice `xml:"EvtStag,omitempty"`

	// Identifies the option that will be selected by default if no instruction is provided by account owner.
	DefaultOptionType *CorporateActionOption1FormatChoice `xml:"DfltOptnTp,omitempty"`

	// Identifies the option number that will be selected by default if no instruction is provided by account owner.
	DefaultOptionNumber *Exact3NumericText `xml:"DfltOptnNb,omitempty"`

	// The method of calculation of drawings and partial redemptions.
	CalculationMethod *CorporateActionCalculationMethod1FormatChoice `xml:"ClctnMtd,omitempty"`

	// Represents the presence of a back end odd lot provision and the quantity of equity required after proration to be eligible for this privilege.
	BackEndOddLotSecuritiesQuantity *UnitOrFaceAmountOrCode1Choice `xml:"BckEndOddLotSctiesQty,omitempty"`

	// Specifies that if an order is prorated holders of odd lots who tender their full position will not have tendered position prorated but rather accepted in full.
	FrontEndOddLotSecuritiesQuantity *UnitOrFaceAmountOrCode1Choice `xml:"FrntEndOddLotSctiesQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"MinExrcblSctiesQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"MinExrcblMltplSctiesQty,omitempty"`

	// Amount used when the called amount is not met by running the lottery with the base denomination.
	IncrementalDenomination *UnitOrFaceAmount1Choice `xml:"IncrmtlDnmtn,omitempty"`

	// New Denomination of the equity following, eg, an increase or decrease in nominal value.
	NewDenominationSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"NewDnmtnSctiesQty,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"NewBrdLotSctiesQty,omitempty"`

	// Quantity of securities the offeror/issuer will purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	SecuritiesQuantitySought *UnitOrFaceAmountOrCode1Choice `xml:"SctiesQtySght,omitempty"`

	// The minimum integral amount of securities that each account owner must have remaining after the called amounts are applied.
	BaseDenomination *UnitOrFaceAmount1Choice `xml:"BaseDnmtn,omitempty"`

	// Specifies the type of change announced.
	ChangeType []*CorporateActionChangeType1FormatChoice `xml:"ChngTp,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferType1FormatChoice `xml:"OfferTp,omitempty"`

	// Indicates whether there is restrictions apply to the CA event or not.
	//
	// Yes = There is restrictions.
	// No = There is no restrictions.
	RestrictionIndicator *YesNoIndicator `xml:"RstrctnInd,omitempty"`

	// Specifies if the issuer will allow the agent to accept partial elections. It is to allow split voting over options. It allows the client to elect more than one option to be selected per designated holding.
	PartialElectionIndicator *YesNoIndicator `xml:"PrtlElctnInd,omitempty"`

	// Specifies the effect on the holdings of electing a Corporate Action option.
	ElectionType *ElectionMovementType1FormatChoice `xml:"ElctnTp,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryType1FormatChoice `xml:"LtryTp,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification13 `xml:"IncmTp,omitempty"`

	// Specifies the conditions in which a dividend is paid.
	DividendType *CorporateActionFrequencyType1FormatChoice `xml:"DvddTp,omitempty"`

	// Type of intermediates securities distribution, eg, stock dividend, reverse right.
	IntermediateSecuritiesDistributionType *IntermediateSecurityDistributionType1FormatChoice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Number of the coupon attached/associated with a security.
	CouponNumber []*Max3NumericText `xml:"CpnNb,omitempty"`

	// Number of days used for calculating the accrued interest amount.
	InterestAccruedNumberOfDays *Number `xml:"IntrstAcrdNbOfDays,omitempty"`

	// New denomination currency of the inancial instrument.
	NewDenominationCurrency *ActiveCurrencyCode `xml:"NewDnmtnCcy,omitempty"`

	// Provides information about the dates related to a CA event.
	DateDetails *CorporateActionDate2 `xml:"DtDtls,omitempty"`

	// Provides information about the prices related to a CA event.
	PriceDetails []*CorporateActionPrice2 `xml:"PricDtls,omitempty"`

	// Provides information about the periods related to a CA event.
	PeriodDetails *CorporateActionPeriod1 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a CA event.
	RateAndAmountDetails *CorporateActionRate1 `xml:"RateAndAmtDtls,omitempty"`

	// Provides additional information.
	CorporateActionAdditionalInformation *CorporateActionNarrative1 `xml:"CorpActnAddtlInf,omitempty"`

	// Indicates whether certification is required from the account owner.
	CertificationRequiredIndicator *YesNoIndicator `xml:"CertfctnReqrdInd,omitempty"`

	// Type of certification which is required.
	CertificationType *BeneficiaryCertificationType1FormatChoice `xml:"CertfctnTp,omitempty"`

	// Specifies whether the capital gain is in the scope of the EU Savings directive for the income realised upon the sale, refund or redemption of shares and units (...) (Article 6(1d)).
	CapitalGain *EUCapitalGain2Code `xml:"CptlGn,omitempty"`

	// Specifies whether the financial instrument calculates the taxable income per dividend/taxable income per share.
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculated2Code `xml:"TaxblIncmPerShrClctd,omitempty"`

	// New companys place of incorporation.
	NewPlaceOfIncorporation *Max70Text `xml:"NewPlcOfIncorprtn,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableStatus1FormatChoice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Specifies the conversion type of an instrument.
	ConversionType *ConversionType1FormatChoice `xml:"ConvsTp,omitempty"`

	// Indicates whether redemption charges apply.
	RedemptionChargesAppliedIndicator *YesNoIndicator `xml:"RedChrgsApldInd,omitempty"`

	// Specifies whether the proceeds of the event will be distributed on a rolling basis rather than on a specific date.
	DistributionType *DistributionType1FormatChoice `xml:"DstrbtnTp,omitempty"`
}

func (c *CorporateAction2) AddEventStage() *CorporateActionEventStage1FormatChoice {
	newValue := new(CorporateActionEventStage1FormatChoice)
	c.EventStage = append(c.EventStage, newValue)
	return newValue
}

func (c *CorporateAction2) AddDefaultOptionType() *CorporateActionOption1FormatChoice {
	c.DefaultOptionType = new(CorporateActionOption1FormatChoice)
	return c.DefaultOptionType
}

func (c *CorporateAction2) SetDefaultOptionNumber(value string) {
	c.DefaultOptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateAction2) AddCalculationMethod() *CorporateActionCalculationMethod1FormatChoice {
	c.CalculationMethod = new(CorporateActionCalculationMethod1FormatChoice)
	return c.CalculationMethod
}

func (c *CorporateAction2) AddBackEndOddLotSecuritiesQuantity() *UnitOrFaceAmountOrCode1Choice {
	c.BackEndOddLotSecuritiesQuantity = new(UnitOrFaceAmountOrCode1Choice)
	return c.BackEndOddLotSecuritiesQuantity
}

func (c *CorporateAction2) AddFrontEndOddLotSecuritiesQuantity() *UnitOrFaceAmountOrCode1Choice {
	c.FrontEndOddLotSecuritiesQuantity = new(UnitOrFaceAmountOrCode1Choice)
	return c.FrontEndOddLotSecuritiesQuantity
}

func (c *CorporateAction2) AddMinimumExercisableSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	c.MinimumExercisableSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return c.MinimumExercisableSecuritiesQuantity
}

func (c *CorporateAction2) AddMinimumExercisableMultipleSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	c.MinimumExercisableMultipleSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return c.MinimumExercisableMultipleSecuritiesQuantity
}

func (c *CorporateAction2) AddIncrementalDenomination() *UnitOrFaceAmount1Choice {
	c.IncrementalDenomination = new(UnitOrFaceAmount1Choice)
	return c.IncrementalDenomination
}

func (c *CorporateAction2) AddNewDenominationSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	c.NewDenominationSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return c.NewDenominationSecuritiesQuantity
}

func (c *CorporateAction2) AddNewBoardLotSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	c.NewBoardLotSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return c.NewBoardLotSecuritiesQuantity
}

func (c *CorporateAction2) AddSecuritiesQuantitySought() *UnitOrFaceAmountOrCode1Choice {
	c.SecuritiesQuantitySought = new(UnitOrFaceAmountOrCode1Choice)
	return c.SecuritiesQuantitySought
}

func (c *CorporateAction2) AddBaseDenomination() *UnitOrFaceAmount1Choice {
	c.BaseDenomination = new(UnitOrFaceAmount1Choice)
	return c.BaseDenomination
}

func (c *CorporateAction2) AddChangeType() *CorporateActionChangeType1FormatChoice {
	newValue := new(CorporateActionChangeType1FormatChoice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateAction2) AddOfferType() *OfferType1FormatChoice {
	newValue := new(OfferType1FormatChoice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateAction2) SetRestrictionIndicator(value string) {
	c.RestrictionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction2) SetPartialElectionIndicator(value string) {
	c.PartialElectionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction2) AddElectionType() *ElectionMovementType1FormatChoice {
	c.ElectionType = new(ElectionMovementType1FormatChoice)
	return c.ElectionType
}

func (c *CorporateAction2) AddLotteryType() *LotteryType1FormatChoice {
	c.LotteryType = new(LotteryType1FormatChoice)
	return c.LotteryType
}

func (c *CorporateAction2) AddIncomeType() *GenericIdentification13 {
	c.IncomeType = new(GenericIdentification13)
	return c.IncomeType
}

func (c *CorporateAction2) AddDividendType() *CorporateActionFrequencyType1FormatChoice {
	c.DividendType = new(CorporateActionFrequencyType1FormatChoice)
	return c.DividendType
}

func (c *CorporateAction2) AddIntermediateSecuritiesDistributionType() *IntermediateSecurityDistributionType1FormatChoice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecurityDistributionType1FormatChoice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateAction2) AddCouponNumber(value string) {
	c.CouponNumber = append(c.CouponNumber, (*Max3NumericText)(&value))
}

func (c *CorporateAction2) SetInterestAccruedNumberOfDays(value string) {
	c.InterestAccruedNumberOfDays = (*Number)(&value)
}

func (c *CorporateAction2) SetNewDenominationCurrency(value string) {
	c.NewDenominationCurrency = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateAction2) AddDateDetails() *CorporateActionDate2 {
	c.DateDetails = new(CorporateActionDate2)
	return c.DateDetails
}

func (c *CorporateAction2) AddPriceDetails() *CorporateActionPrice2 {
	newValue := new(CorporateActionPrice2)
	c.PriceDetails = append(c.PriceDetails, newValue)
	return newValue
}

func (c *CorporateAction2) AddPeriodDetails() *CorporateActionPeriod1 {
	c.PeriodDetails = new(CorporateActionPeriod1)
	return c.PeriodDetails
}

func (c *CorporateAction2) AddRateAndAmountDetails() *CorporateActionRate1 {
	c.RateAndAmountDetails = new(CorporateActionRate1)
	return c.RateAndAmountDetails
}

func (c *CorporateAction2) AddCorporateActionAdditionalInformation() *CorporateActionNarrative1 {
	c.CorporateActionAdditionalInformation = new(CorporateActionNarrative1)
	return c.CorporateActionAdditionalInformation
}

func (c *CorporateAction2) SetCertificationRequiredIndicator(value string) {
	c.CertificationRequiredIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction2) AddCertificationType() *BeneficiaryCertificationType1FormatChoice {
	c.CertificationType = new(BeneficiaryCertificationType1FormatChoice)
	return c.CertificationType
}

func (c *CorporateAction2) SetCapitalGain(value string) {
	c.CapitalGain = (*EUCapitalGain2Code)(&value)
}

func (c *CorporateAction2) SetTaxableIncomePerShareCalculated(value string) {
	c.TaxableIncomePerShareCalculated = (*TaxableIncomePerShareCalculated2Code)(&value)
}

func (c *CorporateAction2) SetNewPlaceOfIncorporation(value string) {
	c.NewPlaceOfIncorporation = (*Max70Text)(&value)
}

func (c *CorporateAction2) AddRenounceableEntitlementStatusType() *RenounceableStatus1FormatChoice {
	c.RenounceableEntitlementStatusType = new(RenounceableStatus1FormatChoice)
	return c.RenounceableEntitlementStatusType
}

func (c *CorporateAction2) AddConversionType() *ConversionType1FormatChoice {
	c.ConversionType = new(ConversionType1FormatChoice)
	return c.ConversionType
}

func (c *CorporateAction2) SetRedemptionChargesAppliedIndicator(value string) {
	c.RedemptionChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction2) AddDistributionType() *DistributionType1FormatChoice {
	c.DistributionType = new(DistributionType1FormatChoice)
	return c.DistributionType
}

// Provides information about the corporate action event.
type CorporateAction24 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate41 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat6Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat1Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction24) AddDateDetails() *CorporateActionDate41 {
	c.DateDetails = new(CorporateActionDate41)
	return c.DateDetails
}

func (c *CorporateAction24) AddEventStage() *CorporateActionEventStageFormat6Choice {
	c.EventStage = new(CorporateActionEventStageFormat6Choice)
	return c.EventStage
}

func (c *CorporateAction24) AddLotteryType() *LotteryTypeFormat1Choice {
	c.LotteryType = new(LotteryTypeFormat1Choice)
	return c.LotteryType
}

// Provides information about the corporate action event.
type CorporateAction3 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate1 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action event.
	PeriodDetails *CorporateActionPeriod3 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action event.
	RateAndAmountDetails *CorporateActionRate3 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action event.
	PriceDetails *CorporateActionPrice3 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action.
	SecuritiesQuantity *CorporateActionQuantity1 `xml:"SctiesQty,omitempty"`

	// Number of days used for calculating the accrued interest amount.
	InterestAccruedNumberOfDays *Max3Number `xml:"IntrstAcrdNbOfDays,omitempty"`

	// Number of the coupon attached/associated with a security.
	CouponNumber []*IdentificationFormat1Choice `xml:"CpnNb,omitempty"`

	// Indicates whether certification is required from the account owner.
	// Yes = certification required.
	// No = no certification required.
	CertificationRequiredIndicator *YesNoIndicator `xml:"CertfctnReqrdInd,omitempty"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether there is restrictions apply to the corporate action event or not.
	// Yes = There is restrictions.
	// No = There is no restrictions.
	RestrictionIndicator *YesNoIndicator `xml:"RstrctnInd,omitempty"`

	// Indicates whether the holder is entitled to accrued interest.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Specifies the conditions in which a dividend is paid.
	DividendType *DividendTypeFormat1Choice `xml:"DvddTp,omitempty"`

	// Specifies the conversion type of an instrument.
	ConversionType *ConversionTypeFormat1Choice `xml:"ConvsTp,omitempty"`

	// Specifies whether the proceeds of the event will be distributed on a rolling basis rather than on a specific date.
	DistributionType *DistributionTypeFormat1Choice `xml:"DstrbtnTp,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat1Choice `xml:"OfferTp,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat1Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage []*CorporateActionEventStageFormat1Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of the additional business process linked to a corporate action event such as a claim compensation or tax refund.
	AdditionalBusinessProcessIndicator []*AdditionalBusinessProcessFormat1Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of change announced.
	ChangeType []*CorporateActionChangeTypeFormat1Choice `xml:"ChngTp,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat1Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies whether the capital gain is in the scope of the EU Savings directive for the income realised upon the sale, refund or redemption of shares and units (...) (Article 6(1d)).
	CapitalGainInOutIndicator *CapitalGainFormat1Choice `xml:"CptlGnInOutInd,omitempty"`

	// Specifies whether the financial instrument calculates the taxable income per dividend/taxable income per share.
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculatedFormat1Choice `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies the effect on the holdings of electing a corporate action option.
	ElectionType *ElectionTypeFormat1Choice `xml:"ElctnTp,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat1Choice `xml:"LtryTp,omitempty"`

	// Specifies the certification format required, ie, physical or electronic format.
	CertificationType *CertificationTypeFormat1Choice `xml:"CertfctnTp,omitempty"`

	// New company's place of incorporation.
	NewPlaceOfIncorporation *Max70Text `xml:"NewPlcOfIncorprtn,omitempty"`

	// Provides additional information. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalInformation *CorporateActionNarrative3 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateAction3) AddDateDetails() *CorporateActionDate1 {
	c.DateDetails = new(CorporateActionDate1)
	return c.DateDetails
}

func (c *CorporateAction3) AddPeriodDetails() *CorporateActionPeriod3 {
	c.PeriodDetails = new(CorporateActionPeriod3)
	return c.PeriodDetails
}

func (c *CorporateAction3) AddRateAndAmountDetails() *CorporateActionRate3 {
	c.RateAndAmountDetails = new(CorporateActionRate3)
	return c.RateAndAmountDetails
}

func (c *CorporateAction3) AddPriceDetails() *CorporateActionPrice3 {
	c.PriceDetails = new(CorporateActionPrice3)
	return c.PriceDetails
}

func (c *CorporateAction3) AddSecuritiesQuantity() *CorporateActionQuantity1 {
	c.SecuritiesQuantity = new(CorporateActionQuantity1)
	return c.SecuritiesQuantity
}

func (c *CorporateAction3) SetInterestAccruedNumberOfDays(value string) {
	c.InterestAccruedNumberOfDays = (*Max3Number)(&value)
}

func (c *CorporateAction3) AddCouponNumber() *IdentificationFormat1Choice {
	newValue := new(IdentificationFormat1Choice)
	c.CouponNumber = append(c.CouponNumber, newValue)
	return newValue
}

func (c *CorporateAction3) SetCertificationRequiredIndicator(value string) {
	c.CertificationRequiredIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction3) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction3) SetRestrictionIndicator(value string) {
	c.RestrictionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction3) SetAccruedInterestIndicator(value string) {
	c.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction3) AddDividendType() *DividendTypeFormat1Choice {
	c.DividendType = new(DividendTypeFormat1Choice)
	return c.DividendType
}

func (c *CorporateAction3) AddConversionType() *ConversionTypeFormat1Choice {
	c.ConversionType = new(ConversionTypeFormat1Choice)
	return c.ConversionType
}

func (c *CorporateAction3) AddDistributionType() *DistributionTypeFormat1Choice {
	c.DistributionType = new(DistributionTypeFormat1Choice)
	return c.DistributionType
}

func (c *CorporateAction3) AddOfferType() *OfferTypeFormat1Choice {
	newValue := new(OfferTypeFormat1Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateAction3) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat1Choice {
	c.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat1Choice)
	return c.RenounceableEntitlementStatusType
}

func (c *CorporateAction3) AddEventStage() *CorporateActionEventStageFormat1Choice {
	newValue := new(CorporateActionEventStageFormat1Choice)
	c.EventStage = append(c.EventStage, newValue)
	return newValue
}

func (c *CorporateAction3) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat1Choice {
	newValue := new(AdditionalBusinessProcessFormat1Choice)
	c.AdditionalBusinessProcessIndicator = append(c.AdditionalBusinessProcessIndicator, newValue)
	return newValue
}

func (c *CorporateAction3) AddChangeType() *CorporateActionChangeTypeFormat1Choice {
	newValue := new(CorporateActionChangeTypeFormat1Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateAction3) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat1Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat1Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateAction3) AddCapitalGainInOutIndicator() *CapitalGainFormat1Choice {
	c.CapitalGainInOutIndicator = new(CapitalGainFormat1Choice)
	return c.CapitalGainInOutIndicator
}

func (c *CorporateAction3) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculatedFormat1Choice {
	c.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculatedFormat1Choice)
	return c.TaxableIncomePerShareCalculated
}

func (c *CorporateAction3) AddElectionType() *ElectionTypeFormat1Choice {
	c.ElectionType = new(ElectionTypeFormat1Choice)
	return c.ElectionType
}

func (c *CorporateAction3) AddLotteryType() *LotteryTypeFormat1Choice {
	c.LotteryType = new(LotteryTypeFormat1Choice)
	return c.LotteryType
}

func (c *CorporateAction3) AddCertificationType() *CertificationTypeFormat1Choice {
	c.CertificationType = new(CertificationTypeFormat1Choice)
	return c.CertificationType
}

func (c *CorporateAction3) SetNewPlaceOfIncorporation(value string) {
	c.NewPlaceOfIncorporation = (*Max70Text)(&value)
}

func (c *CorporateAction3) AddAdditionalInformation() *CorporateActionNarrative3 {
	c.AdditionalInformation = new(CorporateActionNarrative3)
	return c.AdditionalInformation
}

// Provides information about the corporate action event.
type CorporateAction31 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate44 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action event.
	PeriodDetails *CorporateActionPeriod10 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action event.
	RateAndAmountDetails *CorporateActionRate66 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action event.
	PriceDetails *CorporateActionPrice57 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action.
	SecuritiesQuantity *CorporateActionQuantity7 `xml:"SctiesQty,omitempty"`

	// Number of days used for calculating the accrued interest amount.
	InterestAccruedNumberOfDays *Max3Number `xml:"IntrstAcrdNbOfDays,omitempty"`

	// Number of the coupon attached/associated with a security.
	CouponNumber []*IdentificationFormat3Choice `xml:"CpnNb,omitempty"`

	// Indicates whether certification/breakdown is required.
	// Yes = certification required.
	// No = no certification required.
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether there is restrictions apply to the corporate action event or not.
	// Yes = There is restrictions.
	// No = There is no restrictions.
	RestrictionIndicator *YesNoIndicator `xml:"RstrctnInd,omitempty"`

	// Indicates whether the holder is entitled to accrued interest.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Indicates whether a letter of guaranteed delivery can be submitted in order to participate in the offer on full eligible position. It is not intended for use in situations arising from failed or late trades.
	LetterOfGuaranteedDeliveryIndicator *YesNoIndicator `xml:"LttrOfGrntedDlvryInd,omitempty"`

	// Specifies the conditions in which a dividend is paid.
	DividendType *DividendTypeFormat9Choice `xml:"DvddTp,omitempty"`

	// Specifies the conversion type of an instrument.
	ConversionType *ConversionTypeFormat3Choice `xml:"ConvsTp,omitempty"`

	// Specifies the conditions in which the instructions and/or payment of the proceeds occurs.
	OccurrenceType *DistributionTypeFormat7Choice `xml:"OcrncTp,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat10Choice `xml:"OfferTp,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat3Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage []*CorporateActionEventStageFormat13Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of the additional business process linked to a corporate action event such as a claim compensation or tax refund.
	AdditionalBusinessProcessIndicator []*AdditionalBusinessProcessFormat9Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of change announced.
	ChangeType []*CorporateActionChangeTypeFormat5Choice `xml:"ChngTp,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat15Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies whether the capital gain is in the scope of the EU Savings directive for the income realised upon the sale, refund or redemption of shares and units (...) (Article 6(1d)).
	CapitalGainInOutIndicator *CapitalGainFormat3Choice `xml:"CptlGnInOutInd,omitempty"`

	// Specifies whether the financial instrument calculates the taxable income per dividend/taxable income per share.
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculatedFormat3Choice `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies the effect on the holdings of electing a corporate action option.
	ElectionType *ElectionTypeFormat3Choice `xml:"ElctnTp,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat4Choice `xml:"LtryTp,omitempty"`

	// Specifies the certification format required, this is, physical or electronic format.
	CertificationType *CertificationTypeFormat3Choice `xml:"CertfctnTp,omitempty"`

	// Specifies the type of consent announced.
	ConsentType *ConsentTypeFormat4Choice `xml:"CnsntTp,omitempty"`

	// Specifies the type of information event.
	InformationType *InformationTypeFormat4Choice `xml:"InfTp,omitempty"`

	// New company's place of incorporation.
	NewPlaceOfIncorporation *Max350Text `xml:"NewPlcOfIncorprtn,omitempty"`

	// Provides additional information. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalInformation *CorporateActionNarrative26 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateAction31) AddDateDetails() *CorporateActionDate44 {
	c.DateDetails = new(CorporateActionDate44)
	return c.DateDetails
}

func (c *CorporateAction31) AddPeriodDetails() *CorporateActionPeriod10 {
	c.PeriodDetails = new(CorporateActionPeriod10)
	return c.PeriodDetails
}

func (c *CorporateAction31) AddRateAndAmountDetails() *CorporateActionRate66 {
	c.RateAndAmountDetails = new(CorporateActionRate66)
	return c.RateAndAmountDetails
}

func (c *CorporateAction31) AddPriceDetails() *CorporateActionPrice57 {
	c.PriceDetails = new(CorporateActionPrice57)
	return c.PriceDetails
}

func (c *CorporateAction31) AddSecuritiesQuantity() *CorporateActionQuantity7 {
	c.SecuritiesQuantity = new(CorporateActionQuantity7)
	return c.SecuritiesQuantity
}

func (c *CorporateAction31) SetInterestAccruedNumberOfDays(value string) {
	c.InterestAccruedNumberOfDays = (*Max3Number)(&value)
}

func (c *CorporateAction31) AddCouponNumber() *IdentificationFormat3Choice {
	newValue := new(IdentificationFormat3Choice)
	c.CouponNumber = append(c.CouponNumber, newValue)
	return newValue
}

func (c *CorporateAction31) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction31) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction31) SetRestrictionIndicator(value string) {
	c.RestrictionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction31) SetAccruedInterestIndicator(value string) {
	c.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction31) SetLetterOfGuaranteedDeliveryIndicator(value string) {
	c.LetterOfGuaranteedDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction31) AddDividendType() *DividendTypeFormat9Choice {
	c.DividendType = new(DividendTypeFormat9Choice)
	return c.DividendType
}

func (c *CorporateAction31) AddConversionType() *ConversionTypeFormat3Choice {
	c.ConversionType = new(ConversionTypeFormat3Choice)
	return c.ConversionType
}

func (c *CorporateAction31) AddOccurrenceType() *DistributionTypeFormat7Choice {
	c.OccurrenceType = new(DistributionTypeFormat7Choice)
	return c.OccurrenceType
}

func (c *CorporateAction31) AddOfferType() *OfferTypeFormat10Choice {
	newValue := new(OfferTypeFormat10Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateAction31) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat3Choice {
	c.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat3Choice)
	return c.RenounceableEntitlementStatusType
}

func (c *CorporateAction31) AddEventStage() *CorporateActionEventStageFormat13Choice {
	newValue := new(CorporateActionEventStageFormat13Choice)
	c.EventStage = append(c.EventStage, newValue)
	return newValue
}

func (c *CorporateAction31) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat9Choice {
	newValue := new(AdditionalBusinessProcessFormat9Choice)
	c.AdditionalBusinessProcessIndicator = append(c.AdditionalBusinessProcessIndicator, newValue)
	return newValue
}

func (c *CorporateAction31) AddChangeType() *CorporateActionChangeTypeFormat5Choice {
	newValue := new(CorporateActionChangeTypeFormat5Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateAction31) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat15Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat15Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateAction31) AddCapitalGainInOutIndicator() *CapitalGainFormat3Choice {
	c.CapitalGainInOutIndicator = new(CapitalGainFormat3Choice)
	return c.CapitalGainInOutIndicator
}

func (c *CorporateAction31) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculatedFormat3Choice {
	c.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculatedFormat3Choice)
	return c.TaxableIncomePerShareCalculated
}

func (c *CorporateAction31) AddElectionType() *ElectionTypeFormat3Choice {
	c.ElectionType = new(ElectionTypeFormat3Choice)
	return c.ElectionType
}

func (c *CorporateAction31) AddLotteryType() *LotteryTypeFormat4Choice {
	c.LotteryType = new(LotteryTypeFormat4Choice)
	return c.LotteryType
}

func (c *CorporateAction31) AddCertificationType() *CertificationTypeFormat3Choice {
	c.CertificationType = new(CertificationTypeFormat3Choice)
	return c.CertificationType
}

func (c *CorporateAction31) AddConsentType() *ConsentTypeFormat4Choice {
	c.ConsentType = new(ConsentTypeFormat4Choice)
	return c.ConsentType
}

func (c *CorporateAction31) AddInformationType() *InformationTypeFormat4Choice {
	c.InformationType = new(InformationTypeFormat4Choice)
	return c.InformationType
}

func (c *CorporateAction31) SetNewPlaceOfIncorporation(value string) {
	c.NewPlaceOfIncorporation = (*Max350Text)(&value)
}

func (c *CorporateAction31) AddAdditionalInformation() *CorporateActionNarrative26 {
	c.AdditionalInformation = new(CorporateActionNarrative26)
	return c.AdditionalInformation
}

// Provides information about the corporate action event.
type CorporateAction32 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate45 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat14Choice `xml:"EvtStag,omitempty"`

	// Indicates whether the message is related to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat10Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat4Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction32) AddDateDetails() *CorporateActionDate45 {
	c.DateDetails = new(CorporateActionDate45)
	return c.DateDetails
}

func (c *CorporateAction32) AddEventStage() *CorporateActionEventStageFormat14Choice {
	c.EventStage = new(CorporateActionEventStageFormat14Choice)
	return c.EventStage
}

func (c *CorporateAction32) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat10Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat10Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateAction32) AddLotteryType() *LotteryTypeFormat4Choice {
	c.LotteryType = new(LotteryTypeFormat4Choice)
	return c.LotteryType
}

// Provides information about the corporate action event.
type CorporateAction33 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate49 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat14Choice `xml:"EvtStag,omitempty"`

	// Indicates that the additional business process relates to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat11Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat4Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction33) AddDateDetails() *CorporateActionDate49 {
	c.DateDetails = new(CorporateActionDate49)
	return c.DateDetails
}

func (c *CorporateAction33) AddEventStage() *CorporateActionEventStageFormat14Choice {
	c.EventStage = new(CorporateActionEventStageFormat14Choice)
	return c.EventStage
}

func (c *CorporateAction33) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat11Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat11Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateAction33) AddLotteryType() *LotteryTypeFormat4Choice {
	c.LotteryType = new(LotteryTypeFormat4Choice)
	return c.LotteryType
}

// Provides information about the corporate action event.
type CorporateAction34 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate49 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat14Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat4Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction34) AddDateDetails() *CorporateActionDate49 {
	c.DateDetails = new(CorporateActionDate49)
	return c.DateDetails
}

func (c *CorporateAction34) AddEventStage() *CorporateActionEventStageFormat14Choice {
	c.EventStage = new(CorporateActionEventStageFormat14Choice)
	return c.EventStage
}

func (c *CorporateAction34) AddLotteryType() *LotteryTypeFormat4Choice {
	c.LotteryType = new(LotteryTypeFormat4Choice)
	return c.LotteryType
}

// Provides information about the corporate action event.
type CorporateAction35 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate50 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat15Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat5Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction35) AddDateDetails() *CorporateActionDate50 {
	c.DateDetails = new(CorporateActionDate50)
	return c.DateDetails
}

func (c *CorporateAction35) AddEventStage() *CorporateActionEventStageFormat15Choice {
	c.EventStage = new(CorporateActionEventStageFormat15Choice)
	return c.EventStage
}

func (c *CorporateAction35) AddLotteryType() *LotteryTypeFormat5Choice {
	c.LotteryType = new(LotteryTypeFormat5Choice)
	return c.LotteryType
}

// Provides information about the corporate action event.
type CorporateAction36 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate50 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat15Choice `xml:"EvtStag,omitempty"`

	// Indicates that the additional business process relates to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat14Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat5Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction36) AddDateDetails() *CorporateActionDate50 {
	c.DateDetails = new(CorporateActionDate50)
	return c.DateDetails
}

func (c *CorporateAction36) AddEventStage() *CorporateActionEventStageFormat15Choice {
	c.EventStage = new(CorporateActionEventStageFormat15Choice)
	return c.EventStage
}

func (c *CorporateAction36) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat14Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat14Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateAction36) AddLotteryType() *LotteryTypeFormat5Choice {
	c.LotteryType = new(LotteryTypeFormat5Choice)
	return c.LotteryType
}

// Provides information about the corporate action event.
type CorporateAction38 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate54 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat15Choice `xml:"EvtStag,omitempty"`

	// Indicates whether the message is related to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat13Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat5Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction38) AddDateDetails() *CorporateActionDate54 {
	c.DateDetails = new(CorporateActionDate54)
	return c.DateDetails
}

func (c *CorporateAction38) AddEventStage() *CorporateActionEventStageFormat15Choice {
	c.EventStage = new(CorporateActionEventStageFormat15Choice)
	return c.EventStage
}

func (c *CorporateAction38) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat13Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat13Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateAction38) AddLotteryType() *LotteryTypeFormat5Choice {
	c.LotteryType = new(LotteryTypeFormat5Choice)
	return c.LotteryType
}

// Provides information about the corporate action event.
type CorporateAction40 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate58 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action event.
	PeriodDetails *CorporateActionPeriod10 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action event.
	RateAndAmountDetails *CorporateActionRate78 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action event.
	PriceDetails *CorporateActionPrice67 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action.
	SecuritiesQuantity *CorporateActionQuantity8 `xml:"SctiesQty,omitempty"`

	// Number of days used for calculating the accrued interest amount.
	InterestAccruedNumberOfDays *Max3Number `xml:"IntrstAcrdNbOfDays,omitempty"`

	// Number of the coupon attached/associated with a security.
	CouponNumber []*IdentificationFormat4Choice `xml:"CpnNb,omitempty"`

	// Indicates whether certification/breakdown is required.
	// Yes = certification required.
	// No = no certification required.
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether there is restrictions apply to the corporate action event or not.
	// Yes = There is restrictions.
	// No = There is no restrictions.
	RestrictionIndicator *YesNoIndicator `xml:"RstrctnInd,omitempty"`

	// Indicates whether the holder is entitled to accrued interest.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Indicates whether a letter of guaranteed delivery can be submitted in order to participate in the offer on full eligible position. It is not intended for use in situations arising from failed or late trades.
	LetterOfGuaranteedDeliveryIndicator *YesNoIndicator `xml:"LttrOfGrntedDlvryInd,omitempty"`

	// Specifies the conditions in which a dividend is paid.
	DividendType *DividendTypeFormat10Choice `xml:"DvddTp,omitempty"`

	// Specifies the conversion type of an instrument.
	ConversionType *ConversionTypeFormat4Choice `xml:"ConvsTp,omitempty"`

	// Specifies the conditions in which the instructions and/or payment of the proceeds occurs.
	OccurrenceType *DistributionTypeFormat8Choice `xml:"OcrncTp,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat11Choice `xml:"OfferTp,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat4Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage []*CorporateActionEventStageFormat20Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of the additional business process linked to a corporate action event such as a claim compensation or tax refund.
	AdditionalBusinessProcessIndicator []*AdditionalBusinessProcessFormat12Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of change announced.
	ChangeType []*CorporateActionChangeTypeFormat8Choice `xml:"ChngTp,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat18Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies whether the capital gain is in the scope of the EU Savings directive for the income realised upon the sale, refund or redemption of shares and units (...) (Article 6(1d)).
	CapitalGainInOutIndicator *CapitalGainFormat4Choice `xml:"CptlGnInOutInd,omitempty"`

	// Specifies whether the financial instrument calculates the taxable income per dividend/taxable income per share.
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculatedFormat4Choice `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies the effect on the holdings of electing a corporate action option.
	ElectionType *ElectionTypeFormat4Choice `xml:"ElctnTp,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat5Choice `xml:"LtryTp,omitempty"`

	// Specifies the certification format required, this is, physical or electronic format.
	CertificationType *CertificationTypeFormat4Choice `xml:"CertfctnTp,omitempty"`

	// Specifies the type of consent announced.
	ConsentType *ConsentTypeFormat5Choice `xml:"CnsntTp,omitempty"`

	// Specifies the type of information event.
	InformationType *InformationTypeFormat5Choice `xml:"InfTp,omitempty"`

	// New company's place of incorporation.
	NewPlaceOfIncorporation *RestrictedFINXMax350Text `xml:"NewPlcOfIncorprtn,omitempty"`

	// Provides additional information. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalInformation *CorporateActionNarrative39 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateAction40) AddDateDetails() *CorporateActionDate58 {
	c.DateDetails = new(CorporateActionDate58)
	return c.DateDetails
}

func (c *CorporateAction40) AddPeriodDetails() *CorporateActionPeriod10 {
	c.PeriodDetails = new(CorporateActionPeriod10)
	return c.PeriodDetails
}

func (c *CorporateAction40) AddRateAndAmountDetails() *CorporateActionRate78 {
	c.RateAndAmountDetails = new(CorporateActionRate78)
	return c.RateAndAmountDetails
}

func (c *CorporateAction40) AddPriceDetails() *CorporateActionPrice67 {
	c.PriceDetails = new(CorporateActionPrice67)
	return c.PriceDetails
}

func (c *CorporateAction40) AddSecuritiesQuantity() *CorporateActionQuantity8 {
	c.SecuritiesQuantity = new(CorporateActionQuantity8)
	return c.SecuritiesQuantity
}

func (c *CorporateAction40) SetInterestAccruedNumberOfDays(value string) {
	c.InterestAccruedNumberOfDays = (*Max3Number)(&value)
}

func (c *CorporateAction40) AddCouponNumber() *IdentificationFormat4Choice {
	newValue := new(IdentificationFormat4Choice)
	c.CouponNumber = append(c.CouponNumber, newValue)
	return newValue
}

func (c *CorporateAction40) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction40) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction40) SetRestrictionIndicator(value string) {
	c.RestrictionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction40) SetAccruedInterestIndicator(value string) {
	c.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction40) SetLetterOfGuaranteedDeliveryIndicator(value string) {
	c.LetterOfGuaranteedDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction40) AddDividendType() *DividendTypeFormat10Choice {
	c.DividendType = new(DividendTypeFormat10Choice)
	return c.DividendType
}

func (c *CorporateAction40) AddConversionType() *ConversionTypeFormat4Choice {
	c.ConversionType = new(ConversionTypeFormat4Choice)
	return c.ConversionType
}

func (c *CorporateAction40) AddOccurrenceType() *DistributionTypeFormat8Choice {
	c.OccurrenceType = new(DistributionTypeFormat8Choice)
	return c.OccurrenceType
}

func (c *CorporateAction40) AddOfferType() *OfferTypeFormat11Choice {
	newValue := new(OfferTypeFormat11Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateAction40) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat4Choice {
	c.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat4Choice)
	return c.RenounceableEntitlementStatusType
}

func (c *CorporateAction40) AddEventStage() *CorporateActionEventStageFormat20Choice {
	newValue := new(CorporateActionEventStageFormat20Choice)
	c.EventStage = append(c.EventStage, newValue)
	return newValue
}

func (c *CorporateAction40) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat12Choice {
	newValue := new(AdditionalBusinessProcessFormat12Choice)
	c.AdditionalBusinessProcessIndicator = append(c.AdditionalBusinessProcessIndicator, newValue)
	return newValue
}

func (c *CorporateAction40) AddChangeType() *CorporateActionChangeTypeFormat8Choice {
	newValue := new(CorporateActionChangeTypeFormat8Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateAction40) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat18Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat18Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateAction40) AddCapitalGainInOutIndicator() *CapitalGainFormat4Choice {
	c.CapitalGainInOutIndicator = new(CapitalGainFormat4Choice)
	return c.CapitalGainInOutIndicator
}

func (c *CorporateAction40) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculatedFormat4Choice {
	c.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculatedFormat4Choice)
	return c.TaxableIncomePerShareCalculated
}

func (c *CorporateAction40) AddElectionType() *ElectionTypeFormat4Choice {
	c.ElectionType = new(ElectionTypeFormat4Choice)
	return c.ElectionType
}

func (c *CorporateAction40) AddLotteryType() *LotteryTypeFormat5Choice {
	c.LotteryType = new(LotteryTypeFormat5Choice)
	return c.LotteryType
}

func (c *CorporateAction40) AddCertificationType() *CertificationTypeFormat4Choice {
	c.CertificationType = new(CertificationTypeFormat4Choice)
	return c.CertificationType
}

func (c *CorporateAction40) AddConsentType() *ConsentTypeFormat5Choice {
	c.ConsentType = new(ConsentTypeFormat5Choice)
	return c.ConsentType
}

func (c *CorporateAction40) AddInformationType() *InformationTypeFormat5Choice {
	c.InformationType = new(InformationTypeFormat5Choice)
	return c.InformationType
}

func (c *CorporateAction40) SetNewPlaceOfIncorporation(value string) {
	c.NewPlaceOfIncorporation = (*RestrictedFINXMax350Text)(&value)
}

func (c *CorporateAction40) AddAdditionalInformation() *CorporateActionNarrative39 {
	c.AdditionalInformation = new(CorporateActionNarrative39)
	return c.AdditionalInformation
}

// Provides information about the corporate action event.
type CorporateAction5 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate14 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action event.
	PeriodDetails *CorporateActionPeriod6 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action event.
	RateAndAmountDetails *CorporateActionRate16 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action event.
	PriceDetails *CorporateActionPrice17 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action.
	SecuritiesQuantity *CorporateActionQuantity3 `xml:"SctiesQty,omitempty"`

	// Number of days used for calculating the accrued interest amount.
	InterestAccruedNumberOfDays *Max3Number `xml:"IntrstAcrdNbOfDays,omitempty"`

	// Number of the coupon attached/associated with a security.
	CouponNumber []*IdentificationFormat1Choice `xml:"CpnNb,omitempty"`

	// Indicates whether certification is required from the account owner.
	// Yes = certification required.
	// No = no certification required.
	CertificationRequiredIndicator *YesNoIndicator `xml:"CertfctnReqrdInd,omitempty"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether there is restrictions apply to the corporate action event or not.
	// Yes = There is restrictions.
	// No = There is no restrictions.
	RestrictionIndicator *YesNoIndicator `xml:"RstrctnInd,omitempty"`

	// Indicates whether the holder is entitled to accrued interest.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Specifies the conditions in which a dividend is paid.
	DividendType *DividendTypeFormat3Choice `xml:"DvddTp,omitempty"`

	// Specifies the conversion type of an instrument.
	ConversionType *ConversionTypeFormat1Choice `xml:"ConvsTp,omitempty"`

	// Specifies whether the proceeds of the event will be distributed on a rolling basis rather than on a specific date.
	DistributionType *DistributionTypeFormat1Choice `xml:"DstrbtnTp,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat1Choice `xml:"OfferTp,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat1Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage []*CorporateActionEventStageFormat3Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of the additional business process linked to a corporate action event such as a claim compensation or tax refund.
	AdditionalBusinessProcessIndicator []*AdditionalBusinessProcessFormat1Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of change announced.
	ChangeType []*CorporateActionChangeTypeFormat1Choice `xml:"ChngTp,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat5Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies whether the capital gain is in the scope of the EU Savings directive for the income realised upon the sale, refund or redemption of shares and units (...) (Article 6(1d)).
	CapitalGainInOutIndicator *CapitalGainFormat1Choice `xml:"CptlGnInOutInd,omitempty"`

	// Specifies whether the financial instrument calculates the taxable income per dividend/taxable income per share.
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculatedFormat1Choice `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies the effect on the holdings of electing a corporate action option.
	ElectionType *ElectionTypeFormat1Choice `xml:"ElctnTp,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat1Choice `xml:"LtryTp,omitempty"`

	// Specifies the certification format required, this is, physical or electronic format.
	CertificationType *CertificationTypeFormat1Choice `xml:"CertfctnTp,omitempty"`

	// New company's place of incorporation.
	NewPlaceOfIncorporation *Max70Text `xml:"NewPlcOfIncorprtn,omitempty"`

	// Provides additional information. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalInformation *CorporateActionNarrative3 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateAction5) AddDateDetails() *CorporateActionDate14 {
	c.DateDetails = new(CorporateActionDate14)
	return c.DateDetails
}

func (c *CorporateAction5) AddPeriodDetails() *CorporateActionPeriod6 {
	c.PeriodDetails = new(CorporateActionPeriod6)
	return c.PeriodDetails
}

func (c *CorporateAction5) AddRateAndAmountDetails() *CorporateActionRate16 {
	c.RateAndAmountDetails = new(CorporateActionRate16)
	return c.RateAndAmountDetails
}

func (c *CorporateAction5) AddPriceDetails() *CorporateActionPrice17 {
	c.PriceDetails = new(CorporateActionPrice17)
	return c.PriceDetails
}

func (c *CorporateAction5) AddSecuritiesQuantity() *CorporateActionQuantity3 {
	c.SecuritiesQuantity = new(CorporateActionQuantity3)
	return c.SecuritiesQuantity
}

func (c *CorporateAction5) SetInterestAccruedNumberOfDays(value string) {
	c.InterestAccruedNumberOfDays = (*Max3Number)(&value)
}

func (c *CorporateAction5) AddCouponNumber() *IdentificationFormat1Choice {
	newValue := new(IdentificationFormat1Choice)
	c.CouponNumber = append(c.CouponNumber, newValue)
	return newValue
}

func (c *CorporateAction5) SetCertificationRequiredIndicator(value string) {
	c.CertificationRequiredIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction5) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction5) SetRestrictionIndicator(value string) {
	c.RestrictionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction5) SetAccruedInterestIndicator(value string) {
	c.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction5) AddDividendType() *DividendTypeFormat3Choice {
	c.DividendType = new(DividendTypeFormat3Choice)
	return c.DividendType
}

func (c *CorporateAction5) AddConversionType() *ConversionTypeFormat1Choice {
	c.ConversionType = new(ConversionTypeFormat1Choice)
	return c.ConversionType
}

func (c *CorporateAction5) AddDistributionType() *DistributionTypeFormat1Choice {
	c.DistributionType = new(DistributionTypeFormat1Choice)
	return c.DistributionType
}

func (c *CorporateAction5) AddOfferType() *OfferTypeFormat1Choice {
	newValue := new(OfferTypeFormat1Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateAction5) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat1Choice {
	c.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat1Choice)
	return c.RenounceableEntitlementStatusType
}

func (c *CorporateAction5) AddEventStage() *CorporateActionEventStageFormat3Choice {
	newValue := new(CorporateActionEventStageFormat3Choice)
	c.EventStage = append(c.EventStage, newValue)
	return newValue
}

func (c *CorporateAction5) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat1Choice {
	newValue := new(AdditionalBusinessProcessFormat1Choice)
	c.AdditionalBusinessProcessIndicator = append(c.AdditionalBusinessProcessIndicator, newValue)
	return newValue
}

func (c *CorporateAction5) AddChangeType() *CorporateActionChangeTypeFormat1Choice {
	newValue := new(CorporateActionChangeTypeFormat1Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateAction5) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat5Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat5Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateAction5) AddCapitalGainInOutIndicator() *CapitalGainFormat1Choice {
	c.CapitalGainInOutIndicator = new(CapitalGainFormat1Choice)
	return c.CapitalGainInOutIndicator
}

func (c *CorporateAction5) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculatedFormat1Choice {
	c.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculatedFormat1Choice)
	return c.TaxableIncomePerShareCalculated
}

func (c *CorporateAction5) AddElectionType() *ElectionTypeFormat1Choice {
	c.ElectionType = new(ElectionTypeFormat1Choice)
	return c.ElectionType
}

func (c *CorporateAction5) AddLotteryType() *LotteryTypeFormat1Choice {
	c.LotteryType = new(LotteryTypeFormat1Choice)
	return c.LotteryType
}

func (c *CorporateAction5) AddCertificationType() *CertificationTypeFormat1Choice {
	c.CertificationType = new(CertificationTypeFormat1Choice)
	return c.CertificationType
}

func (c *CorporateAction5) SetNewPlaceOfIncorporation(value string) {
	c.NewPlaceOfIncorporation = (*Max70Text)(&value)
}

func (c *CorporateAction5) AddAdditionalInformation() *CorporateActionNarrative3 {
	c.AdditionalInformation = new(CorporateActionNarrative3)
	return c.AdditionalInformation
}

// Provides information about the corporate action event.
type CorporateAction7 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate22 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action event.
	PeriodDetails *CorporateActionPeriod8 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action event.
	RateAndAmountDetails *CorporateActionRate27 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action event.
	PriceDetails *CorporateActionPrice17 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action.
	SecuritiesQuantity *CorporateActionQuantity5 `xml:"SctiesQty,omitempty"`

	// Number of days used for calculating the accrued interest amount.
	InterestAccruedNumberOfDays *Max3Number `xml:"IntrstAcrdNbOfDays,omitempty"`

	// Number of the coupon attached/associated with a security.
	CouponNumber []*IdentificationFormat1Choice `xml:"CpnNb,omitempty"`

	// Indicates whether certification/breakdown is required.
	// Yes = certification required.
	// No = no certification required.
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether there is restrictions apply to the corporate action event or not.
	// Yes = There is restrictions.
	// No = There is no restrictions.
	RestrictionIndicator *YesNoIndicator `xml:"RstrctnInd,omitempty"`

	// Indicates whether the holder is entitled to accrued interest.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Specifies the conditions in which a dividend is paid.
	DividendType *DividendTypeFormat3Choice `xml:"DvddTp,omitempty"`

	// Specifies the conversion type of an instrument.
	ConversionType *ConversionTypeFormat1Choice `xml:"ConvsTp,omitempty"`

	// Specifies the conditions in which the payment of the proceeds occurs.
	PaymentOccurrenceType *DistributionTypeFormat3Choice `xml:"PmtOcrncTp,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat1Choice `xml:"OfferTp,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat1Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage []*CorporateActionEventStageFormat3Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of the additional business process linked to a corporate action event such as a claim compensation or tax refund.
	AdditionalBusinessProcessIndicator []*AdditionalBusinessProcessFormat1Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of change announced.
	ChangeType []*CorporateActionChangeTypeFormat1Choice `xml:"ChngTp,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat9Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies whether the capital gain is in the scope of the EU Savings directive for the income realised upon the sale, refund or redemption of shares and units (...) (Article 6(1d)).
	CapitalGainInOutIndicator *CapitalGainFormat1Choice `xml:"CptlGnInOutInd,omitempty"`

	// Specifies whether the financial instrument calculates the taxable income per dividend/taxable income per share.
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculatedFormat1Choice `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies the effect on the holdings of electing a corporate action option.
	ElectionType *ElectionTypeFormat1Choice `xml:"ElctnTp,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat1Choice `xml:"LtryTp,omitempty"`

	// Specifies the certification format required, this is, physical or electronic format.
	CertificationType *CertificationTypeFormat1Choice `xml:"CertfctnTp,omitempty"`

	// New company's place of incorporation.
	NewPlaceOfIncorporation *Max70Text `xml:"NewPlcOfIncorprtn,omitempty"`

	// Provides additional information. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalInformation *CorporateActionNarrative3 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateAction7) AddDateDetails() *CorporateActionDate22 {
	c.DateDetails = new(CorporateActionDate22)
	return c.DateDetails
}

func (c *CorporateAction7) AddPeriodDetails() *CorporateActionPeriod8 {
	c.PeriodDetails = new(CorporateActionPeriod8)
	return c.PeriodDetails
}

func (c *CorporateAction7) AddRateAndAmountDetails() *CorporateActionRate27 {
	c.RateAndAmountDetails = new(CorporateActionRate27)
	return c.RateAndAmountDetails
}

func (c *CorporateAction7) AddPriceDetails() *CorporateActionPrice17 {
	c.PriceDetails = new(CorporateActionPrice17)
	return c.PriceDetails
}

func (c *CorporateAction7) AddSecuritiesQuantity() *CorporateActionQuantity5 {
	c.SecuritiesQuantity = new(CorporateActionQuantity5)
	return c.SecuritiesQuantity
}

func (c *CorporateAction7) SetInterestAccruedNumberOfDays(value string) {
	c.InterestAccruedNumberOfDays = (*Max3Number)(&value)
}

func (c *CorporateAction7) AddCouponNumber() *IdentificationFormat1Choice {
	newValue := new(IdentificationFormat1Choice)
	c.CouponNumber = append(c.CouponNumber, newValue)
	return newValue
}

func (c *CorporateAction7) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction7) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction7) SetRestrictionIndicator(value string) {
	c.RestrictionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction7) SetAccruedInterestIndicator(value string) {
	c.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction7) AddDividendType() *DividendTypeFormat3Choice {
	c.DividendType = new(DividendTypeFormat3Choice)
	return c.DividendType
}

func (c *CorporateAction7) AddConversionType() *ConversionTypeFormat1Choice {
	c.ConversionType = new(ConversionTypeFormat1Choice)
	return c.ConversionType
}

func (c *CorporateAction7) AddPaymentOccurrenceType() *DistributionTypeFormat3Choice {
	c.PaymentOccurrenceType = new(DistributionTypeFormat3Choice)
	return c.PaymentOccurrenceType
}

func (c *CorporateAction7) AddOfferType() *OfferTypeFormat1Choice {
	newValue := new(OfferTypeFormat1Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateAction7) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat1Choice {
	c.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat1Choice)
	return c.RenounceableEntitlementStatusType
}

func (c *CorporateAction7) AddEventStage() *CorporateActionEventStageFormat3Choice {
	newValue := new(CorporateActionEventStageFormat3Choice)
	c.EventStage = append(c.EventStage, newValue)
	return newValue
}

func (c *CorporateAction7) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat1Choice {
	newValue := new(AdditionalBusinessProcessFormat1Choice)
	c.AdditionalBusinessProcessIndicator = append(c.AdditionalBusinessProcessIndicator, newValue)
	return newValue
}

func (c *CorporateAction7) AddChangeType() *CorporateActionChangeTypeFormat1Choice {
	newValue := new(CorporateActionChangeTypeFormat1Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateAction7) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat9Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat9Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateAction7) AddCapitalGainInOutIndicator() *CapitalGainFormat1Choice {
	c.CapitalGainInOutIndicator = new(CapitalGainFormat1Choice)
	return c.CapitalGainInOutIndicator
}

func (c *CorporateAction7) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculatedFormat1Choice {
	c.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculatedFormat1Choice)
	return c.TaxableIncomePerShareCalculated
}

func (c *CorporateAction7) AddElectionType() *ElectionTypeFormat1Choice {
	c.ElectionType = new(ElectionTypeFormat1Choice)
	return c.ElectionType
}

func (c *CorporateAction7) AddLotteryType() *LotteryTypeFormat1Choice {
	c.LotteryType = new(LotteryTypeFormat1Choice)
	return c.LotteryType
}

func (c *CorporateAction7) AddCertificationType() *CertificationTypeFormat1Choice {
	c.CertificationType = new(CertificationTypeFormat1Choice)
	return c.CertificationType
}

func (c *CorporateAction7) SetNewPlaceOfIncorporation(value string) {
	c.NewPlaceOfIncorporation = (*Max70Text)(&value)
}

func (c *CorporateAction7) AddAdditionalInformation() *CorporateActionNarrative3 {
	c.AdditionalInformation = new(CorporateActionNarrative3)
	return c.AdditionalInformation
}

// An event determined by a corporation's board of directors, that changes the existing corporate capital structure or financial condition.
type CorporateAction9 struct {

	// Type of corporate action event, in a free-text format.
	EventType *Max35Text `xml:"EvtTp"`

	// Identification of a corporate action assigned by an official central body/entity within a given market.
	EventIdentification *Max35Text `xml:"EvtId"`
}

func (c *CorporateAction9) SetEventType(value string) {
	c.EventType = (*Max35Text)(&value)
}

func (c *CorporateAction9) SetEventIdentification(value string) {
	c.EventIdentification = (*Max35Text)(&value)
}
