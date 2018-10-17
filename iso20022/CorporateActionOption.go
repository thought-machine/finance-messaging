package iso20022

// Provides information about the CA option.
type CorporateActionOption1 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *CorporateActionEventStatus2FormatChoice `xml:"OptnAvlbtySts"`

	// Whether or not certification is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Type of certification which is required.
	CertificationType *BeneficiaryCertificationType1FormatChoice `xml:"CertfctnTp,omitempty"`

	// Identification of a temporary security used for processing reasons, eg, contra security used in the US.
	AssentedLineSecurityIdentification *SecurityIdentification7 `xml:"AssntdLineSctyId,omitempty"`

	// Identification of the safekeeping account held by an agent at the CSD.
	AgentSecuritiesAccountIdentification *Max35Text `xml:"AgtSctiesAcctId,omitempty"`

	// Identification of the cash account held by an agent at the CSD.
	AgentCashAccountIdentification *AccountIdentification2Choice `xml:"AgtCshAcctId,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferType1FormatChoice `xml:"OfferTp,omitempty"`

	// Type of intermediates securities distribution, eg, stock dividend, reverse right.
	IntermediateSecuritiesDistributionType *IntermediateSecurityDistributionType1FormatChoice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd"`

	// Provides information about the dates related to a CA option.
	DateDetails *CorporateActionDate4 `xml:"DtDtls,omitempty"`

	// Provides information about rates and amounts related to a CA option.
	RateAndAmountDetails *CorporateActionRate2 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a CA option.
	PriceDetails *CorporateActionPrice1 `xml:"PricDtls,omitempty"`

	// Provides information about the periods related to a CA option.
	PeriodDetails *CorporateActionPeriod2 `xml:"PrdDtls,omitempty"`

	// Provides information about the securities movement linked to the CA option.
	SecuritiesMovementDetails []*SecurityOption1 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the CA option.
	CashMovementDetails []*CashOption1 `xml:"CshMvmntDtls,omitempty"`

	// Provides information about the agents linked to the CA option.
	CorporateActionOtherAgentDetails []*CorporateActionAgent1 `xml:"CorpActnOthrAgtDtls,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1FormatChoice `xml:"FrctnDspstn,omitempty"`

	// ndicates whether redemption charges apply.
	RedemptionChargesAppliedIndicator *YesNoIndicator `xml:"RedChrgsApldInd,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeatures1FormatChoice `xml:"OptnFeatrs,omitempty"`

	// Provides additional information.
	CorporateActionAdditionalInformation *CorporateActionNarrative1 `xml:"CorpActnAddtlInf,omitempty"`
}

func (c *CorporateActionOption1) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption1) AddOptionType() *CorporateActionOption1FormatChoice {
	c.OptionType = new(CorporateActionOption1FormatChoice)
	return c.OptionType
}

func (c *CorporateActionOption1) AddOptionAvailabilityStatus() *CorporateActionEventStatus2FormatChoice {
	c.OptionAvailabilityStatus = new(CorporateActionEventStatus2FormatChoice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption1) SetCertificationIndicator(value string) {
	c.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption1) AddCertificationType() *BeneficiaryCertificationType1FormatChoice {
	c.CertificationType = new(BeneficiaryCertificationType1FormatChoice)
	return c.CertificationType
}

func (c *CorporateActionOption1) AddAssentedLineSecurityIdentification() *SecurityIdentification7 {
	c.AssentedLineSecurityIdentification = new(SecurityIdentification7)
	return c.AssentedLineSecurityIdentification
}

func (c *CorporateActionOption1) SetAgentSecuritiesAccountIdentification(value string) {
	c.AgentSecuritiesAccountIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionOption1) AddAgentCashAccountIdentification() *AccountIdentification2Choice {
	c.AgentCashAccountIdentification = new(AccountIdentification2Choice)
	return c.AgentCashAccountIdentification
}

func (c *CorporateActionOption1) AddOfferType() *OfferType1FormatChoice {
	newValue := new(OfferType1FormatChoice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption1) AddIntermediateSecuritiesDistributionType() *IntermediateSecurityDistributionType1FormatChoice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecurityDistributionType1FormatChoice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption1) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption1) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption1) AddDateDetails() *CorporateActionDate4 {
	c.DateDetails = new(CorporateActionDate4)
	return c.DateDetails
}

func (c *CorporateActionOption1) AddRateAndAmountDetails() *CorporateActionRate2 {
	c.RateAndAmountDetails = new(CorporateActionRate2)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption1) AddPriceDetails() *CorporateActionPrice1 {
	c.PriceDetails = new(CorporateActionPrice1)
	return c.PriceDetails
}

func (c *CorporateActionOption1) AddPeriodDetails() *CorporateActionPeriod2 {
	c.PeriodDetails = new(CorporateActionPeriod2)
	return c.PeriodDetails
}

func (c *CorporateActionOption1) AddSecuritiesMovementDetails() *SecurityOption1 {
	newValue := new(SecurityOption1)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption1) AddCashMovementDetails() *CashOption1 {
	newValue := new(CashOption1)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption1) AddCorporateActionOtherAgentDetails() *CorporateActionAgent1 {
	newValue := new(CorporateActionAgent1)
	c.CorporateActionOtherAgentDetails = append(c.CorporateActionOtherAgentDetails, newValue)
	return newValue
}

func (c *CorporateActionOption1) AddFractionDisposition() *FractionDispositionType1FormatChoice {
	c.FractionDisposition = new(FractionDispositionType1FormatChoice)
	return c.FractionDisposition
}

func (c *CorporateActionOption1) SetRedemptionChargesAppliedIndicator(value string) {
	c.RedemptionChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption1) AddOptionFeatures() *OptionFeatures1FormatChoice {
	newValue := new(OptionFeatures1FormatChoice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption1) AddCorporateActionAdditionalInformation() *CorporateActionNarrative1 {
	c.CorporateActionAdditionalInformation = new(CorporateActionNarrative1)
	return c.CorporateActionAdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption10 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption2Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat1Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat2Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat2Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus1Choice `xml:"OptnAvlbtySts,omitempty"`

	// Type of certification which is required.
	CertificationType []*BeneficiaryCertificationType1Choice `xml:"CertfctnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether redemption charges apply.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether or not certification is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate8 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod5 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate5 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice6 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity related to a corporate action option.
	SecuritiesQuantity *SecuritiesOption1 `xml:"SctiesQty,omitempty"`

	// Provides information about securities movement related to a corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption4 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption3 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information about the corporate action movement.
	AdditionalInformation *CorporateActionNarrative5 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption10) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption10) AddOptionType() *CorporateActionOption2Choice {
	c.OptionType = new(CorporateActionOption2Choice)
	return c.OptionType
}

func (c *CorporateActionOption10) AddFractionDisposition() *FractionDispositionType1Choice {
	c.FractionDisposition = new(FractionDispositionType1Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption10) AddOfferType() *OfferTypeFormat1Choice {
	newValue := new(OfferTypeFormat1Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption10) AddOptionFeatures() *OptionFeaturesFormat2Choice {
	newValue := new(OptionFeaturesFormat2Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption10) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat2Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat2Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption10) AddOptionAvailabilityStatus() *OptionAvailabilityStatus1Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus1Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption10) AddCertificationType() *BeneficiaryCertificationType1Choice {
	newValue := new(BeneficiaryCertificationType1Choice)
	c.CertificationType = append(c.CertificationType, newValue)
	return newValue
}

func (c *CorporateActionOption10) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption10) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption10) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption10) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption10) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption10) SetCertificationIndicator(value string) {
	c.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption10) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption10) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption10) AddSecurityIdentification() *SecurityIdentification11 {
	c.SecurityIdentification = new(SecurityIdentification11)
	return c.SecurityIdentification
}

func (c *CorporateActionOption10) AddDateDetails() *CorporateActionDate8 {
	c.DateDetails = new(CorporateActionDate8)
	return c.DateDetails
}

func (c *CorporateActionOption10) AddPeriodDetails() *CorporateActionPeriod5 {
	c.PeriodDetails = new(CorporateActionPeriod5)
	return c.PeriodDetails
}

func (c *CorporateActionOption10) AddRateAndAmountDetails() *CorporateActionRate5 {
	c.RateAndAmountDetails = new(CorporateActionRate5)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption10) AddPriceDetails() *CorporateActionPrice6 {
	c.PriceDetails = new(CorporateActionPrice6)
	return c.PriceDetails
}

func (c *CorporateActionOption10) AddSecuritiesQuantity() *SecuritiesOption1 {
	c.SecuritiesQuantity = new(SecuritiesOption1)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption10) AddSecuritiesMovementDetails() *SecuritiesOption4 {
	newValue := new(SecuritiesOption4)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption10) AddCashMovementDetails() *CashOption3 {
	newValue := new(CashOption3)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption10) AddAdditionalInformation() *CorporateActionNarrative5 {
	c.AdditionalInformation = new(CorporateActionNarrative5)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption100 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption10Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType19Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat5Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat9Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat9Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus1Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType5Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether redemption charges apply.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate29 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate45 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice28 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity related to a corporate action option.
	SecuritiesQuantity *SecuritiesOption23 `xml:"SctiesQty,omitempty"`

	// Provides information about securities movement related to a corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption40 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption31 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information about the corporate action movement.
	AdditionalInformation *CorporateActionNarrative20 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption100) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption100) AddOptionType() *CorporateActionOption10Choice {
	c.OptionType = new(CorporateActionOption10Choice)
	return c.OptionType
}

func (c *CorporateActionOption100) AddFractionDisposition() *FractionDispositionType19Choice {
	c.FractionDisposition = new(FractionDispositionType19Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption100) AddOfferType() *OfferTypeFormat5Choice {
	newValue := new(OfferTypeFormat5Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption100) AddOptionFeatures() *OptionFeaturesFormat9Choice {
	newValue := new(OptionFeaturesFormat9Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption100) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat9Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat9Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption100) AddOptionAvailabilityStatus() *OptionAvailabilityStatus1Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus1Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption100) AddCertificationBreakdownType() *BeneficiaryCertificationType5Choice {
	newValue := new(BeneficiaryCertificationType5Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption100) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption100) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption100) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption100) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption100) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption100) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption100) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption100) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption100) AddSecurityIdentification() *SecurityIdentification14 {
	c.SecurityIdentification = new(SecurityIdentification14)
	return c.SecurityIdentification
}

func (c *CorporateActionOption100) AddDateDetails() *CorporateActionDate29 {
	c.DateDetails = new(CorporateActionDate29)
	return c.DateDetails
}

func (c *CorporateActionOption100) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption100) AddRateAndAmountDetails() *CorporateActionRate45 {
	c.RateAndAmountDetails = new(CorporateActionRate45)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption100) AddPriceDetails() *CorporateActionPrice28 {
	c.PriceDetails = new(CorporateActionPrice28)
	return c.PriceDetails
}

func (c *CorporateActionOption100) AddSecuritiesQuantity() *SecuritiesOption23 {
	c.SecuritiesQuantity = new(SecuritiesOption23)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption100) AddSecuritiesMovementDetails() *SecuritiesOption40 {
	newValue := new(SecuritiesOption40)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption100) AddCashMovementDetails() *CashOption31 {
	newValue := new(CashOption31)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption100) AddAdditionalInformation() *CorporateActionNarrative20 {
	c.AdditionalInformation = new(CorporateActionNarrative20)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption101 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption10Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType19Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat5Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat12Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus1Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType5Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate29 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate44 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice28 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption23 `xml:"SctiesQty,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption40 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption32 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative20 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption101) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption101) AddOptionType() *CorporateActionOption10Choice {
	c.OptionType = new(CorporateActionOption10Choice)
	return c.OptionType
}

func (c *CorporateActionOption101) AddFractionDisposition() *FractionDispositionType19Choice {
	c.FractionDisposition = new(FractionDispositionType19Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption101) AddOfferType() *OfferTypeFormat5Choice {
	newValue := new(OfferTypeFormat5Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption101) AddOptionFeatures() *OptionFeaturesFormat12Choice {
	newValue := new(OptionFeaturesFormat12Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption101) AddOptionAvailabilityStatus() *OptionAvailabilityStatus1Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus1Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption101) AddCertificationBreakdownType() *BeneficiaryCertificationType5Choice {
	newValue := new(BeneficiaryCertificationType5Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption101) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption101) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption101) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption101) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption101) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption101) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption101) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption101) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption101) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption101) AddDateDetails() *CorporateActionDate29 {
	c.DateDetails = new(CorporateActionDate29)
	return c.DateDetails
}

func (c *CorporateActionOption101) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption101) AddRateAndAmountDetails() *CorporateActionRate44 {
	c.RateAndAmountDetails = new(CorporateActionRate44)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption101) AddPriceDetails() *CorporateActionPrice28 {
	c.PriceDetails = new(CorporateActionPrice28)
	return c.PriceDetails
}

func (c *CorporateActionOption101) AddSecuritiesQuantity() *SecuritiesOption23 {
	c.SecuritiesQuantity = new(SecuritiesOption23)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption101) AddSecuritiesMovementDetails() *SecuritiesOption40 {
	newValue := new(SecuritiesOption40)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption101) AddCashMovementDetails() *CashOption32 {
	newValue := new(CashOption32)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption101) AddAdditionalInformation() *CorporateActionNarrative20 {
	c.AdditionalInformation = new(CorporateActionNarrative20)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption102 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption11Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat1Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType23Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate18 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod9 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate46 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice30 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption42 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption30 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption102) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption102) AddOptionType() *CorporateActionOption11Choice {
	c.OptionType = new(CorporateActionOption11Choice)
	return c.OptionType
}

func (c *CorporateActionOption102) AddOptionFeatures() *OptionFeaturesFormat1Choice {
	newValue := new(OptionFeaturesFormat1Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption102) AddFractionDisposition() *FractionDispositionType23Choice {
	c.FractionDisposition = new(FractionDispositionType23Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption102) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption102) AddDateDetails() *CorporateActionDate18 {
	c.DateDetails = new(CorporateActionDate18)
	return c.DateDetails
}

func (c *CorporateActionOption102) AddPeriodDetails() *CorporateActionPeriod9 {
	c.PeriodDetails = new(CorporateActionPeriod9)
	return c.PeriodDetails
}

func (c *CorporateActionOption102) AddRateAndAmountDetails() *CorporateActionRate46 {
	c.RateAndAmountDetails = new(CorporateActionRate46)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption102) AddPriceDetails() *CorporateActionPrice30 {
	c.PriceDetails = new(CorporateActionPrice30)
	return c.PriceDetails
}

func (c *CorporateActionOption102) AddPlaceOfTrade() *MarketIdentification78 {
	c.PlaceOfTrade = new(MarketIdentification78)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption102) AddSecuritiesMovementDetails() *SecuritiesOption42 {
	newValue := new(SecuritiesOption42)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption102) AddCashMovementDetails() *CashOption30 {
	newValue := new(CashOption30)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption103 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption12Choice `xml:"OptnTp"`

	// Specifies how fractional amount/quantities are treated.
	FractionDisposition *FractionDispositionType17Choice `xml:"FrctnDspstn,omitempty"`

	// Type of changes affecting the security form.
	ChangeType []*CorporateActionChangeTypeFormat2Choice `xml:"ChngTp,omitempty"`

	// Specifies that the corporate action instruction is to be processed using the Available-for-Collateral pool.
	EligibleForCollateralIndicator *YesNoIndicator `xml:"ElgblForCollInd,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds.
	CurrencyToBuy *ActiveCurrencyCode `xml:"CcyToBuy,omitempty"`

	// Account servicer is instructed to sell the indicated currency in order to obtain the necessary currency to fund the transaction within this instruction message.
	CurrencyToSell *ActiveCurrencyCode `xml:"CcyToSell,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption2 `xml:"SctiesQty"`

	// Date/time at which the instructing party requests the instruction to be executed.
	ExecutionRequestedDateTime *DateAndDateTimeChoice `xml:"ExctnReqdDtTm,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate47 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice44 `xml:"PricDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative8 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption103) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption103) AddOptionType() *CorporateActionOption12Choice {
	c.OptionType = new(CorporateActionOption12Choice)
	return c.OptionType
}

func (c *CorporateActionOption103) AddFractionDisposition() *FractionDispositionType17Choice {
	c.FractionDisposition = new(FractionDispositionType17Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption103) AddChangeType() *CorporateActionChangeTypeFormat2Choice {
	newValue := new(CorporateActionChangeTypeFormat2Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateActionOption103) SetEligibleForCollateralIndicator(value string) {
	c.EligibleForCollateralIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption103) SetCurrencyToBuy(value string) {
	c.CurrencyToBuy = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption103) SetCurrencyToSell(value string) {
	c.CurrencyToSell = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption103) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption103) AddSecurityIdentification() *SecurityIdentification14 {
	c.SecurityIdentification = new(SecurityIdentification14)
	return c.SecurityIdentification
}

func (c *CorporateActionOption103) AddSecuritiesQuantity() *SecuritiesOption2 {
	c.SecuritiesQuantity = new(SecuritiesOption2)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption103) AddExecutionRequestedDateTime() *DateAndDateTimeChoice {
	c.ExecutionRequestedDateTime = new(DateAndDateTimeChoice)
	return c.ExecutionRequestedDateTime
}

func (c *CorporateActionOption103) AddRateAndAmountDetails() *CorporateActionRate47 {
	c.RateAndAmountDetails = new(CorporateActionRate47)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption103) AddPriceDetails() *CorporateActionPrice44 {
	c.PriceDetails = new(CorporateActionPrice44)
	return c.PriceDetails
}

func (c *CorporateActionOption103) AddAdditionalInformation() *CorporateActionNarrative8 {
	c.AdditionalInformation = new(CorporateActionNarrative8)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption111 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption11Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat1Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType23Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate18 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod11 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate46 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice30 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption42 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption39 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption111) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption111) AddOptionType() *CorporateActionOption11Choice {
	c.OptionType = new(CorporateActionOption11Choice)
	return c.OptionType
}

func (c *CorporateActionOption111) AddOptionFeatures() *OptionFeaturesFormat1Choice {
	newValue := new(OptionFeaturesFormat1Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption111) AddFractionDisposition() *FractionDispositionType23Choice {
	c.FractionDisposition = new(FractionDispositionType23Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption111) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption111) AddDateDetails() *CorporateActionDate18 {
	c.DateDetails = new(CorporateActionDate18)
	return c.DateDetails
}

func (c *CorporateActionOption111) AddPeriodDetails() *CorporateActionPeriod11 {
	c.PeriodDetails = new(CorporateActionPeriod11)
	return c.PeriodDetails
}

func (c *CorporateActionOption111) AddRateAndAmountDetails() *CorporateActionRate46 {
	c.RateAndAmountDetails = new(CorporateActionRate46)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption111) AddPriceDetails() *CorporateActionPrice30 {
	c.PriceDetails = new(CorporateActionPrice30)
	return c.PriceDetails
}

func (c *CorporateActionOption111) AddPlaceOfTrade() *MarketIdentification78 {
	c.PlaceOfTrade = new(MarketIdentification78)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption111) AddSecuritiesMovementDetails() *SecuritiesOption42 {
	newValue := new(SecuritiesOption42)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption111) AddCashMovementDetails() *CashOption39 {
	newValue := new(CashOption39)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption114 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption18Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType26Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat10Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat16Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus3Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType9Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Indicates whether the option, different from the default one, shall be applied by the account owner.
	AppliedOptionIndicator *YesNoIndicator `xml:"ApldOptnInd,omitempty"`

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate48 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate68 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice58 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption51 `xml:"SctiesQty,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption49 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption43 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative29 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption114) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption114) AddOptionType() *CorporateActionOption18Choice {
	c.OptionType = new(CorporateActionOption18Choice)
	return c.OptionType
}

func (c *CorporateActionOption114) AddFractionDisposition() *FractionDispositionType26Choice {
	c.FractionDisposition = new(FractionDispositionType26Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption114) AddOfferType() *OfferTypeFormat10Choice {
	newValue := new(OfferTypeFormat10Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption114) AddOptionFeatures() *OptionFeaturesFormat16Choice {
	newValue := new(OptionFeaturesFormat16Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption114) AddOptionAvailabilityStatus() *OptionAvailabilityStatus3Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus3Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption114) AddCertificationBreakdownType() *BeneficiaryCertificationType9Choice {
	newValue := new(BeneficiaryCertificationType9Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption114) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption114) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption114) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption114) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption114) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption114) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption114) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption114) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption114) SetAppliedOptionIndicator(value string) {
	c.AppliedOptionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption114) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption114) AddDateDetails() *CorporateActionDate48 {
	c.DateDetails = new(CorporateActionDate48)
	return c.DateDetails
}

func (c *CorporateActionOption114) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption114) AddRateAndAmountDetails() *CorporateActionRate68 {
	c.RateAndAmountDetails = new(CorporateActionRate68)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption114) AddPriceDetails() *CorporateActionPrice58 {
	c.PriceDetails = new(CorporateActionPrice58)
	return c.PriceDetails
}

func (c *CorporateActionOption114) AddSecuritiesQuantity() *SecuritiesOption51 {
	c.SecuritiesQuantity = new(SecuritiesOption51)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption114) AddSecuritiesMovementDetails() *SecuritiesOption49 {
	newValue := new(SecuritiesOption49)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption114) AddCashMovementDetails() *CashOption43 {
	newValue := new(CashOption43)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption114) AddAdditionalInformation() *CorporateActionNarrative29 {
	c.AdditionalInformation = new(CorporateActionNarrative29)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption115 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption18Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType26Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat10Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat17Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat15Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus3Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType9Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether redemption charges apply.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Indicates whether the option, different from the default one, shall be applied by the account owner.
	AppliedOptionIndicator *YesNoIndicator `xml:"ApldOptnInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate48 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate67 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice58 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity related to a corporate action option.
	SecuritiesQuantity *SecuritiesOption51 `xml:"SctiesQty,omitempty"`

	// Provides information about securities movement related to a corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption49 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption42 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information about the corporate action movement.
	AdditionalInformation *CorporateActionNarrative29 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption115) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption115) AddOptionType() *CorporateActionOption18Choice {
	c.OptionType = new(CorporateActionOption18Choice)
	return c.OptionType
}

func (c *CorporateActionOption115) AddFractionDisposition() *FractionDispositionType26Choice {
	c.FractionDisposition = new(FractionDispositionType26Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption115) AddOfferType() *OfferTypeFormat10Choice {
	newValue := new(OfferTypeFormat10Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption115) AddOptionFeatures() *OptionFeaturesFormat17Choice {
	newValue := new(OptionFeaturesFormat17Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption115) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat15Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat15Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption115) AddOptionAvailabilityStatus() *OptionAvailabilityStatus3Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus3Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption115) AddCertificationBreakdownType() *BeneficiaryCertificationType9Choice {
	newValue := new(BeneficiaryCertificationType9Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption115) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption115) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption115) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption115) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption115) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption115) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption115) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption115) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption115) SetAppliedOptionIndicator(value string) {
	c.AppliedOptionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption115) AddSecurityIdentification() *SecurityIdentification19 {
	c.SecurityIdentification = new(SecurityIdentification19)
	return c.SecurityIdentification
}

func (c *CorporateActionOption115) AddDateDetails() *CorporateActionDate48 {
	c.DateDetails = new(CorporateActionDate48)
	return c.DateDetails
}

func (c *CorporateActionOption115) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption115) AddRateAndAmountDetails() *CorporateActionRate67 {
	c.RateAndAmountDetails = new(CorporateActionRate67)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption115) AddPriceDetails() *CorporateActionPrice58 {
	c.PriceDetails = new(CorporateActionPrice58)
	return c.PriceDetails
}

func (c *CorporateActionOption115) AddSecuritiesQuantity() *SecuritiesOption51 {
	c.SecuritiesQuantity = new(SecuritiesOption51)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption115) AddSecuritiesMovementDetails() *SecuritiesOption49 {
	newValue := new(SecuritiesOption49)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption115) AddCashMovementDetails() *CashOption42 {
	newValue := new(CashOption42)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption115) AddAdditionalInformation() *CorporateActionNarrative29 {
	c.AdditionalInformation = new(CorporateActionNarrative29)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption116 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption21Choice `xml:"OptnTp"`

	// Party that owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Account on which a securities entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *SignedQuantityFormat7 `xml:"TtlElgblBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *SignedQuantityFormat7 `xml:"InstdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *SignedQuantityFormat7 `xml:"UinstdBal,omitempty"`

	// Quantity of securities that has been assigned the status indicated.
	StatusQuantity *Quantity6Choice `xml:"StsQty"`
}

func (c *CorporateActionOption116) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption116) AddOptionType() *CorporateActionOption21Choice {
	c.OptionType = new(CorporateActionOption21Choice)
	return c.OptionType
}

func (c *CorporateActionOption116) AddAccountOwner() *PartyIdentification92Choice {
	c.AccountOwner = new(PartyIdentification92Choice)
	return c.AccountOwner
}

func (c *CorporateActionOption116) SetSafekeepingAccount(value string) {
	c.SafekeepingAccount = (*Max35Text)(&value)
}

func (c *CorporateActionOption116) AddCashAccount() *CashAccountIdentification5Choice {
	c.CashAccount = new(CashAccountIdentification5Choice)
	return c.CashAccount
}

func (c *CorporateActionOption116) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	c.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return c.SafekeepingPlace
}

func (c *CorporateActionOption116) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption116) AddTotalEligibleBalance() *SignedQuantityFormat7 {
	c.TotalEligibleBalance = new(SignedQuantityFormat7)
	return c.TotalEligibleBalance
}

func (c *CorporateActionOption116) AddInstructedBalance() *SignedQuantityFormat7 {
	c.InstructedBalance = new(SignedQuantityFormat7)
	return c.InstructedBalance
}

func (c *CorporateActionOption116) AddUninstructedBalance() *SignedQuantityFormat7 {
	c.UninstructedBalance = new(SignedQuantityFormat7)
	return c.UninstructedBalance
}

func (c *CorporateActionOption116) AddStatusQuantity() *Quantity6Choice {
	c.StatusQuantity = new(Quantity6Choice)
	return c.StatusQuantity
}

// Provides information about the corporate action option.
type CorporateActionOption117 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption19Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat18Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType27Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate46 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod11 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate70 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice61 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification84 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption50 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption44 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption117) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption117) AddOptionType() *CorporateActionOption19Choice {
	c.OptionType = new(CorporateActionOption19Choice)
	return c.OptionType
}

func (c *CorporateActionOption117) AddOptionFeatures() *OptionFeaturesFormat18Choice {
	newValue := new(OptionFeaturesFormat18Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption117) AddFractionDisposition() *FractionDispositionType27Choice {
	c.FractionDisposition = new(FractionDispositionType27Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption117) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption117) AddDateDetails() *CorporateActionDate46 {
	c.DateDetails = new(CorporateActionDate46)
	return c.DateDetails
}

func (c *CorporateActionOption117) AddPeriodDetails() *CorporateActionPeriod11 {
	c.PeriodDetails = new(CorporateActionPeriod11)
	return c.PeriodDetails
}

func (c *CorporateActionOption117) AddRateAndAmountDetails() *CorporateActionRate70 {
	c.RateAndAmountDetails = new(CorporateActionRate70)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption117) AddPriceDetails() *CorporateActionPrice61 {
	c.PriceDetails = new(CorporateActionPrice61)
	return c.PriceDetails
}

func (c *CorporateActionOption117) AddPlaceOfTrade() *MarketIdentification84 {
	c.PlaceOfTrade = new(MarketIdentification84)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption117) AddSecuritiesMovementDetails() *SecuritiesOption50 {
	newValue := new(SecuritiesOption50)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption117) AddCashMovementDetails() *CashOption44 {
	newValue := new(CashOption44)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption118 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption20Choice `xml:"OptnTp"`

	// Specifies how fractional amount/quantities are treated.
	FractionDisposition *FractionDispositionType28Choice `xml:"FrctnDspstn,omitempty"`

	// Type of changes affecting the security form.
	ChangeType []*CorporateActionChangeTypeFormat6Choice `xml:"ChngTp,omitempty"`

	// Specifies that the corporate action instruction is to be processed using the Available-for-Collateral pool.
	EligibleForCollateralIndicator *YesNoIndicator `xml:"ElgblForCollInd,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds.
	CurrencyToBuy *ActiveCurrencyCode `xml:"CcyToBuy,omitempty"`

	// Account servicer is instructed to sell the indicated currency in order to obtain the necessary currency to fund the transaction within this instruction message.
	CurrencyToSell *ActiveCurrencyCode `xml:"CcyToSell,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption52 `xml:"SctiesQty"`

	// Date/time at which the instructing party requests the instruction to be executed.
	ExecutionRequestedDateTime *DateAndDateTimeChoice `xml:"ExctnReqdDtTm,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate71 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice60 `xml:"PricDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative32 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption118) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption118) AddOptionType() *CorporateActionOption20Choice {
	c.OptionType = new(CorporateActionOption20Choice)
	return c.OptionType
}

func (c *CorporateActionOption118) AddFractionDisposition() *FractionDispositionType28Choice {
	c.FractionDisposition = new(FractionDispositionType28Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption118) AddChangeType() *CorporateActionChangeTypeFormat6Choice {
	newValue := new(CorporateActionChangeTypeFormat6Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateActionOption118) SetEligibleForCollateralIndicator(value string) {
	c.EligibleForCollateralIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption118) SetCurrencyToBuy(value string) {
	c.CurrencyToBuy = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption118) SetCurrencyToSell(value string) {
	c.CurrencyToSell = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption118) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption118) AddSecurityIdentification() *SecurityIdentification19 {
	c.SecurityIdentification = new(SecurityIdentification19)
	return c.SecurityIdentification
}

func (c *CorporateActionOption118) AddSecuritiesQuantity() *SecuritiesOption52 {
	c.SecuritiesQuantity = new(SecuritiesOption52)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption118) AddExecutionRequestedDateTime() *DateAndDateTimeChoice {
	c.ExecutionRequestedDateTime = new(DateAndDateTimeChoice)
	return c.ExecutionRequestedDateTime
}

func (c *CorporateActionOption118) AddRateAndAmountDetails() *CorporateActionRate71 {
	c.RateAndAmountDetails = new(CorporateActionRate71)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption118) AddPriceDetails() *CorporateActionPrice60 {
	c.PriceDetails = new(CorporateActionPrice60)
	return c.PriceDetails
}

func (c *CorporateActionOption118) AddAdditionalInformation() *CorporateActionNarrative32 {
	c.AdditionalInformation = new(CorporateActionNarrative32)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption119 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption19Choice `xml:"OptnTp"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption53 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption45 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption119) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption119) AddOptionType() *CorporateActionOption19Choice {
	c.OptionType = new(CorporateActionOption19Choice)
	return c.OptionType
}

func (c *CorporateActionOption119) AddSecuritiesMovementDetails() *SecuritiesOption53 {
	newValue := new(SecuritiesOption53)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption119) AddCashMovementDetails() *CashOption45 {
	newValue := new(CashOption45)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption12 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption3Choice `xml:"OptnTp"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption5 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption5 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption12) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption12) AddOptionType() *CorporateActionOption3Choice {
	c.OptionType = new(CorporateActionOption3Choice)
	return c.OptionType
}

func (c *CorporateActionOption12) AddSecuritiesMovementDetails() *SecuritiesOption5 {
	newValue := new(SecuritiesOption5)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption12) AddCashMovementDetails() *CashOption5 {
	newValue := new(CashOption5)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption120 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption20Choice `xml:"OptnTp"`

	// Quantity of securities to which this instruction applies.
	InstructedQuantity *Quantity20Choice `xml:"InstdQty"`
}

func (c *CorporateActionOption120) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption120) AddOptionType() *CorporateActionOption20Choice {
	c.OptionType = new(CorporateActionOption20Choice)
	return c.OptionType
}

func (c *CorporateActionOption120) AddInstructedQuantity() *Quantity20Choice {
	c.InstructedQuantity = new(Quantity20Choice)
	return c.InstructedQuantity
}

// Provides information about the corporate action option.
type CorporateActionOption121 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption22Choice `xml:"OptnTp"`

	// Party that owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct,omitempty"`

	// Account on which a securities entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat11Choice `xml:"SfkpgPlc,omitempty"`

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId,omitempty"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *SignedQuantityFormat8 `xml:"TtlElgblBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *SignedQuantityFormat8 `xml:"InstdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *SignedQuantityFormat8 `xml:"UinstdBal,omitempty"`

	// Quantity of securities that has been assigned the status indicated.
	StatusQuantity *Quantity10Choice `xml:"StsQty"`
}

func (c *CorporateActionOption121) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption121) AddOptionType() *CorporateActionOption22Choice {
	c.OptionType = new(CorporateActionOption22Choice)
	return c.OptionType
}

func (c *CorporateActionOption121) AddAccountOwner() *PartyIdentification103Choice {
	c.AccountOwner = new(PartyIdentification103Choice)
	return c.AccountOwner
}

func (c *CorporateActionOption121) SetSafekeepingAccount(value string) {
	c.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (c *CorporateActionOption121) AddCashAccount() *CashAccountIdentification6Choice {
	c.CashAccount = new(CashAccountIdentification6Choice)
	return c.CashAccount
}

func (c *CorporateActionOption121) AddSafekeepingPlace() *SafekeepingPlaceFormat11Choice {
	c.SafekeepingPlace = new(SafekeepingPlaceFormat11Choice)
	return c.SafekeepingPlace
}

func (c *CorporateActionOption121) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption121) AddTotalEligibleBalance() *SignedQuantityFormat8 {
	c.TotalEligibleBalance = new(SignedQuantityFormat8)
	return c.TotalEligibleBalance
}

func (c *CorporateActionOption121) AddInstructedBalance() *SignedQuantityFormat8 {
	c.InstructedBalance = new(SignedQuantityFormat8)
	return c.InstructedBalance
}

func (c *CorporateActionOption121) AddUninstructedBalance() *SignedQuantityFormat8 {
	c.UninstructedBalance = new(SignedQuantityFormat8)
	return c.UninstructedBalance
}

func (c *CorporateActionOption121) AddStatusQuantity() *Quantity10Choice {
	c.StatusQuantity = new(Quantity10Choice)
	return c.StatusQuantity
}

// Provides information about the corporate action option.
type CorporateActionOption123 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption25Choice `xml:"OptnTp"`

	// Specifies how fractional amount/quantities are treated.
	FractionDisposition *FractionDispositionType29Choice `xml:"FrctnDspstn,omitempty"`

	// Type of changes affecting the security form.
	ChangeType []*CorporateActionChangeTypeFormat7Choice `xml:"ChngTp,omitempty"`

	// Specifies that the corporate action instruction is to be processed using the Available-for-Collateral pool.
	EligibleForCollateralIndicator *YesNoIndicator `xml:"ElgblForCollInd,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds.
	CurrencyToBuy *ActiveCurrencyCode `xml:"CcyToBuy,omitempty"`

	// Account servicer is instructed to sell the indicated currency in order to obtain the necessary currency to fund the transaction within this instruction message.
	CurrencyToSell *ActiveCurrencyCode `xml:"CcyToSell,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification20 `xml:"SctyId,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption54 `xml:"SctiesQty"`

	// Date/time at which the instructing party requests the instruction to be executed.
	ExecutionRequestedDateTime *DateAndDateTimeChoice `xml:"ExctnReqdDtTm,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate73 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice62 `xml:"PricDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative33 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption123) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption123) AddOptionType() *CorporateActionOption25Choice {
	c.OptionType = new(CorporateActionOption25Choice)
	return c.OptionType
}

func (c *CorporateActionOption123) AddFractionDisposition() *FractionDispositionType29Choice {
	c.FractionDisposition = new(FractionDispositionType29Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption123) AddChangeType() *CorporateActionChangeTypeFormat7Choice {
	newValue := new(CorporateActionChangeTypeFormat7Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateActionOption123) SetEligibleForCollateralIndicator(value string) {
	c.EligibleForCollateralIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption123) SetCurrencyToBuy(value string) {
	c.CurrencyToBuy = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption123) SetCurrencyToSell(value string) {
	c.CurrencyToSell = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption123) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption123) AddSecurityIdentification() *SecurityIdentification20 {
	c.SecurityIdentification = new(SecurityIdentification20)
	return c.SecurityIdentification
}

func (c *CorporateActionOption123) AddSecuritiesQuantity() *SecuritiesOption54 {
	c.SecuritiesQuantity = new(SecuritiesOption54)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption123) AddExecutionRequestedDateTime() *DateAndDateTimeChoice {
	c.ExecutionRequestedDateTime = new(DateAndDateTimeChoice)
	return c.ExecutionRequestedDateTime
}

func (c *CorporateActionOption123) AddRateAndAmountDetails() *CorporateActionRate73 {
	c.RateAndAmountDetails = new(CorporateActionRate73)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption123) AddPriceDetails() *CorporateActionPrice62 {
	c.PriceDetails = new(CorporateActionPrice62)
	return c.PriceDetails
}

func (c *CorporateActionOption123) AddAdditionalInformation() *CorporateActionNarrative33 {
	c.AdditionalInformation = new(CorporateActionNarrative33)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption124 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption26Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat19Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType30Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate52 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod11 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate74 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice63 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification90 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption55 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption46 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption124) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption124) AddOptionType() *CorporateActionOption26Choice {
	c.OptionType = new(CorporateActionOption26Choice)
	return c.OptionType
}

func (c *CorporateActionOption124) AddOptionFeatures() *OptionFeaturesFormat19Choice {
	newValue := new(OptionFeaturesFormat19Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption124) AddFractionDisposition() *FractionDispositionType30Choice {
	c.FractionDisposition = new(FractionDispositionType30Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption124) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption124) AddDateDetails() *CorporateActionDate52 {
	c.DateDetails = new(CorporateActionDate52)
	return c.DateDetails
}

func (c *CorporateActionOption124) AddPeriodDetails() *CorporateActionPeriod11 {
	c.PeriodDetails = new(CorporateActionPeriod11)
	return c.PeriodDetails
}

func (c *CorporateActionOption124) AddRateAndAmountDetails() *CorporateActionRate74 {
	c.RateAndAmountDetails = new(CorporateActionRate74)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption124) AddPriceDetails() *CorporateActionPrice63 {
	c.PriceDetails = new(CorporateActionPrice63)
	return c.PriceDetails
}

func (c *CorporateActionOption124) AddPlaceOfTrade() *MarketIdentification90 {
	c.PlaceOfTrade = new(MarketIdentification90)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption124) AddSecuritiesMovementDetails() *SecuritiesOption55 {
	newValue := new(SecuritiesOption55)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption124) AddCashMovementDetails() *CashOption46 {
	newValue := new(CashOption46)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption125 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption23Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType31Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat11Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat20Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat18Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus4Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType12Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether redemption charges apply.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Indicates whether the option, different from the default one, shall be applied by the account owner.
	AppliedOptionIndicator *YesNoIndicator `xml:"ApldOptnInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification20 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate55 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate76 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice65 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity related to a corporate action option.
	SecuritiesQuantity *SecuritiesOption56 `xml:"SctiesQty,omitempty"`

	// Provides information about securities movement related to a corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption57 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption47 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information about the corporate action movement.
	AdditionalInformation *CorporateActionNarrative36 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption125) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption125) AddOptionType() *CorporateActionOption23Choice {
	c.OptionType = new(CorporateActionOption23Choice)
	return c.OptionType
}

func (c *CorporateActionOption125) AddFractionDisposition() *FractionDispositionType31Choice {
	c.FractionDisposition = new(FractionDispositionType31Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption125) AddOfferType() *OfferTypeFormat11Choice {
	newValue := new(OfferTypeFormat11Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption125) AddOptionFeatures() *OptionFeaturesFormat20Choice {
	newValue := new(OptionFeaturesFormat20Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption125) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat18Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat18Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption125) AddOptionAvailabilityStatus() *OptionAvailabilityStatus4Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus4Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption125) AddCertificationBreakdownType() *BeneficiaryCertificationType12Choice {
	newValue := new(BeneficiaryCertificationType12Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption125) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption125) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption125) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption125) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption125) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption125) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption125) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption125) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption125) SetAppliedOptionIndicator(value string) {
	c.AppliedOptionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption125) AddSecurityIdentification() *SecurityIdentification20 {
	c.SecurityIdentification = new(SecurityIdentification20)
	return c.SecurityIdentification
}

func (c *CorporateActionOption125) AddDateDetails() *CorporateActionDate55 {
	c.DateDetails = new(CorporateActionDate55)
	return c.DateDetails
}

func (c *CorporateActionOption125) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption125) AddRateAndAmountDetails() *CorporateActionRate76 {
	c.RateAndAmountDetails = new(CorporateActionRate76)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption125) AddPriceDetails() *CorporateActionPrice65 {
	c.PriceDetails = new(CorporateActionPrice65)
	return c.PriceDetails
}

func (c *CorporateActionOption125) AddSecuritiesQuantity() *SecuritiesOption56 {
	c.SecuritiesQuantity = new(SecuritiesOption56)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption125) AddSecuritiesMovementDetails() *SecuritiesOption57 {
	newValue := new(SecuritiesOption57)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption125) AddCashMovementDetails() *CashOption47 {
	newValue := new(CashOption47)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption125) AddAdditionalInformation() *CorporateActionNarrative36 {
	c.AdditionalInformation = new(CorporateActionNarrative36)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption126 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption26Choice `xml:"OptnTp"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption58 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption48 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption126) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption126) AddOptionType() *CorporateActionOption26Choice {
	c.OptionType = new(CorporateActionOption26Choice)
	return c.OptionType
}

func (c *CorporateActionOption126) AddSecuritiesMovementDetails() *SecuritiesOption58 {
	newValue := new(SecuritiesOption58)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption126) AddCashMovementDetails() *CashOption48 {
	newValue := new(CashOption48)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption127 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption23Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType31Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat11Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat21Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus4Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType12Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Indicates whether the option, different from the default one, shall be applied by the account owner.
	AppliedOptionIndicator *YesNoIndicator `xml:"ApldOptnInd,omitempty"`

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate55 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate79 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice65 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption56 `xml:"SctiesQty,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption57 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption49 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative36 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption127) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption127) AddOptionType() *CorporateActionOption23Choice {
	c.OptionType = new(CorporateActionOption23Choice)
	return c.OptionType
}

func (c *CorporateActionOption127) AddFractionDisposition() *FractionDispositionType31Choice {
	c.FractionDisposition = new(FractionDispositionType31Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption127) AddOfferType() *OfferTypeFormat11Choice {
	newValue := new(OfferTypeFormat11Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption127) AddOptionFeatures() *OptionFeaturesFormat21Choice {
	newValue := new(OptionFeaturesFormat21Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption127) AddOptionAvailabilityStatus() *OptionAvailabilityStatus4Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus4Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption127) AddCertificationBreakdownType() *BeneficiaryCertificationType12Choice {
	newValue := new(BeneficiaryCertificationType12Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption127) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption127) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption127) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption127) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption127) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption127) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption127) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption127) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption127) SetAppliedOptionIndicator(value string) {
	c.AppliedOptionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption127) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption127) AddDateDetails() *CorporateActionDate55 {
	c.DateDetails = new(CorporateActionDate55)
	return c.DateDetails
}

func (c *CorporateActionOption127) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption127) AddRateAndAmountDetails() *CorporateActionRate79 {
	c.RateAndAmountDetails = new(CorporateActionRate79)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption127) AddPriceDetails() *CorporateActionPrice65 {
	c.PriceDetails = new(CorporateActionPrice65)
	return c.PriceDetails
}

func (c *CorporateActionOption127) AddSecuritiesQuantity() *SecuritiesOption56 {
	c.SecuritiesQuantity = new(SecuritiesOption56)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption127) AddSecuritiesMovementDetails() *SecuritiesOption57 {
	newValue := new(SecuritiesOption57)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption127) AddCashMovementDetails() *CashOption49 {
	newValue := new(CashOption49)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption127) AddAdditionalInformation() *CorporateActionNarrative36 {
	c.AdditionalInformation = new(CorporateActionNarrative36)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption128 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption29Choice `xml:"OptnTp"`

	// Quantity of securities to which this instruction applies.
	InstructedQuantity *Quantity40Choice `xml:"InstdQty"`
}

func (c *CorporateActionOption128) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption128) AddOptionType() *CorporateActionOption29Choice {
	c.OptionType = new(CorporateActionOption29Choice)
	return c.OptionType
}

func (c *CorporateActionOption128) AddInstructedQuantity() *Quantity40Choice {
	c.InstructedQuantity = new(Quantity40Choice)
	return c.InstructedQuantity
}

// Provides information about the corporate action option.
type CorporateActionOption129 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption18Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType26Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat10Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat22Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat15Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus3Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType9Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether redemption charges apply.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Indicates whether the option, different from the default one, shall be applied by the account owner.
	AppliedOptionIndicator *YesNoIndicator `xml:"ApldOptnInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate48 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate81 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice58 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity related to a corporate action option.
	SecuritiesQuantity *SecuritiesOption51 `xml:"SctiesQty,omitempty"`

	// Provides information about securities movement related to a corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption59 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption51 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information about the corporate action movement.
	AdditionalInformation *CorporateActionNarrative29 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption129) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption129) AddOptionType() *CorporateActionOption18Choice {
	c.OptionType = new(CorporateActionOption18Choice)
	return c.OptionType
}

func (c *CorporateActionOption129) AddFractionDisposition() *FractionDispositionType26Choice {
	c.FractionDisposition = new(FractionDispositionType26Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption129) AddOfferType() *OfferTypeFormat10Choice {
	newValue := new(OfferTypeFormat10Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption129) AddOptionFeatures() *OptionFeaturesFormat22Choice {
	newValue := new(OptionFeaturesFormat22Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption129) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat15Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat15Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption129) AddOptionAvailabilityStatus() *OptionAvailabilityStatus3Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus3Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption129) AddCertificationBreakdownType() *BeneficiaryCertificationType9Choice {
	newValue := new(BeneficiaryCertificationType9Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption129) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption129) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption129) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption129) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption129) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption129) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption129) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption129) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption129) SetAppliedOptionIndicator(value string) {
	c.AppliedOptionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption129) AddSecurityIdentification() *SecurityIdentification19 {
	c.SecurityIdentification = new(SecurityIdentification19)
	return c.SecurityIdentification
}

func (c *CorporateActionOption129) AddDateDetails() *CorporateActionDate48 {
	c.DateDetails = new(CorporateActionDate48)
	return c.DateDetails
}

func (c *CorporateActionOption129) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption129) AddRateAndAmountDetails() *CorporateActionRate81 {
	c.RateAndAmountDetails = new(CorporateActionRate81)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption129) AddPriceDetails() *CorporateActionPrice58 {
	c.PriceDetails = new(CorporateActionPrice58)
	return c.PriceDetails
}

func (c *CorporateActionOption129) AddSecuritiesQuantity() *SecuritiesOption51 {
	c.SecuritiesQuantity = new(SecuritiesOption51)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption129) AddSecuritiesMovementDetails() *SecuritiesOption59 {
	newValue := new(SecuritiesOption59)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption129) AddCashMovementDetails() *CashOption51 {
	newValue := new(CashOption51)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption129) AddAdditionalInformation() *CorporateActionNarrative29 {
	c.AdditionalInformation = new(CorporateActionNarrative29)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption130 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption18Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType26Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat10Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat22Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus3Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType9Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Indicates whether the option, different from the default one, shall be applied by the account owner.
	AppliedOptionIndicator *YesNoIndicator `xml:"ApldOptnInd,omitempty"`

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate48 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate80 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice58 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption51 `xml:"SctiesQty,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption59 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption50 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative29 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption130) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption130) AddOptionType() *CorporateActionOption18Choice {
	c.OptionType = new(CorporateActionOption18Choice)
	return c.OptionType
}

func (c *CorporateActionOption130) AddFractionDisposition() *FractionDispositionType26Choice {
	c.FractionDisposition = new(FractionDispositionType26Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption130) AddOfferType() *OfferTypeFormat10Choice {
	newValue := new(OfferTypeFormat10Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption130) AddOptionFeatures() *OptionFeaturesFormat22Choice {
	newValue := new(OptionFeaturesFormat22Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption130) AddOptionAvailabilityStatus() *OptionAvailabilityStatus3Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus3Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption130) AddCertificationBreakdownType() *BeneficiaryCertificationType9Choice {
	newValue := new(BeneficiaryCertificationType9Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption130) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption130) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption130) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption130) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption130) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption130) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption130) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption130) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption130) SetAppliedOptionIndicator(value string) {
	c.AppliedOptionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption130) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption130) AddDateDetails() *CorporateActionDate48 {
	c.DateDetails = new(CorporateActionDate48)
	return c.DateDetails
}

func (c *CorporateActionOption130) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption130) AddRateAndAmountDetails() *CorporateActionRate80 {
	c.RateAndAmountDetails = new(CorporateActionRate80)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption130) AddPriceDetails() *CorporateActionPrice58 {
	c.PriceDetails = new(CorporateActionPrice58)
	return c.PriceDetails
}

func (c *CorporateActionOption130) AddSecuritiesQuantity() *SecuritiesOption51 {
	c.SecuritiesQuantity = new(SecuritiesOption51)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption130) AddSecuritiesMovementDetails() *SecuritiesOption59 {
	newValue := new(SecuritiesOption59)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption130) AddCashMovementDetails() *CashOption50 {
	newValue := new(CashOption50)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption130) AddAdditionalInformation() *CorporateActionNarrative29 {
	c.AdditionalInformation = new(CorporateActionNarrative29)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption131 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption20Choice `xml:"OptnTp"`

	// Specifies how fractional amount/quantities are treated.
	FractionDisposition *FractionDispositionType28Choice `xml:"FrctnDspstn,omitempty"`

	// Type of changes affecting the security form.
	ChangeType []*CorporateActionChangeTypeFormat6Choice `xml:"ChngTp,omitempty"`

	// Specifies that the corporate action instruction is to be processed using the Available-for-Collateral pool.
	EligibleForCollateralIndicator *YesNoIndicator `xml:"ElgblForCollInd,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds.
	CurrencyToBuy *ActiveCurrencyCode `xml:"CcyToBuy,omitempty"`

	// Account servicer is instructed to sell the indicated currency in order to obtain the necessary currency to fund the transaction within this instruction message.
	CurrencyToSell *ActiveCurrencyCode `xml:"CcyToSell,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantityOrInstructedAmount *SecuritiesQuantityOrAmount1Choice `xml:"SctiesQtyOrInstdAmt"`

	// Date/time at which the instructing party requests the instruction to be executed.
	ExecutionRequestedDateTime *DateAndDateTimeChoice `xml:"ExctnReqdDtTm,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate71 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice60 `xml:"PricDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative32 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption131) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption131) AddOptionType() *CorporateActionOption20Choice {
	c.OptionType = new(CorporateActionOption20Choice)
	return c.OptionType
}

func (c *CorporateActionOption131) AddFractionDisposition() *FractionDispositionType28Choice {
	c.FractionDisposition = new(FractionDispositionType28Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption131) AddChangeType() *CorporateActionChangeTypeFormat6Choice {
	newValue := new(CorporateActionChangeTypeFormat6Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateActionOption131) SetEligibleForCollateralIndicator(value string) {
	c.EligibleForCollateralIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption131) SetCurrencyToBuy(value string) {
	c.CurrencyToBuy = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption131) SetCurrencyToSell(value string) {
	c.CurrencyToSell = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption131) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption131) AddSecurityIdentification() *SecurityIdentification19 {
	c.SecurityIdentification = new(SecurityIdentification19)
	return c.SecurityIdentification
}

func (c *CorporateActionOption131) AddSecuritiesQuantityOrInstructedAmount() *SecuritiesQuantityOrAmount1Choice {
	c.SecuritiesQuantityOrInstructedAmount = new(SecuritiesQuantityOrAmount1Choice)
	return c.SecuritiesQuantityOrInstructedAmount
}

func (c *CorporateActionOption131) AddExecutionRequestedDateTime() *DateAndDateTimeChoice {
	c.ExecutionRequestedDateTime = new(DateAndDateTimeChoice)
	return c.ExecutionRequestedDateTime
}

func (c *CorporateActionOption131) AddRateAndAmountDetails() *CorporateActionRate71 {
	c.RateAndAmountDetails = new(CorporateActionRate71)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption131) AddPriceDetails() *CorporateActionPrice60 {
	c.PriceDetails = new(CorporateActionPrice60)
	return c.PriceDetails
}

func (c *CorporateActionOption131) AddAdditionalInformation() *CorporateActionNarrative32 {
	c.AdditionalInformation = new(CorporateActionNarrative32)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption132 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption19Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat18Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType27Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate46 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod11 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate82 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice61 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification84 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption60 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption52 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption132) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption132) AddOptionType() *CorporateActionOption19Choice {
	c.OptionType = new(CorporateActionOption19Choice)
	return c.OptionType
}

func (c *CorporateActionOption132) AddOptionFeatures() *OptionFeaturesFormat18Choice {
	newValue := new(OptionFeaturesFormat18Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption132) AddFractionDisposition() *FractionDispositionType27Choice {
	c.FractionDisposition = new(FractionDispositionType27Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption132) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption132) AddDateDetails() *CorporateActionDate46 {
	c.DateDetails = new(CorporateActionDate46)
	return c.DateDetails
}

func (c *CorporateActionOption132) AddPeriodDetails() *CorporateActionPeriod11 {
	c.PeriodDetails = new(CorporateActionPeriod11)
	return c.PeriodDetails
}

func (c *CorporateActionOption132) AddRateAndAmountDetails() *CorporateActionRate82 {
	c.RateAndAmountDetails = new(CorporateActionRate82)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption132) AddPriceDetails() *CorporateActionPrice61 {
	c.PriceDetails = new(CorporateActionPrice61)
	return c.PriceDetails
}

func (c *CorporateActionOption132) AddPlaceOfTrade() *MarketIdentification84 {
	c.PlaceOfTrade = new(MarketIdentification84)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption132) AddSecuritiesMovementDetails() *SecuritiesOption60 {
	newValue := new(SecuritiesOption60)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption132) AddCashMovementDetails() *CashOption52 {
	newValue := new(CashOption52)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption133 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption23Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType31Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat11Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat23Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus4Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType12Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Indicates whether the option, different from the default one, shall be applied by the account owner.
	AppliedOptionIndicator *YesNoIndicator `xml:"ApldOptnInd,omitempty"`

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate55 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate83 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice65 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption56 `xml:"SctiesQty,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption61 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption53 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative36 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption133) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption133) AddOptionType() *CorporateActionOption23Choice {
	c.OptionType = new(CorporateActionOption23Choice)
	return c.OptionType
}

func (c *CorporateActionOption133) AddFractionDisposition() *FractionDispositionType31Choice {
	c.FractionDisposition = new(FractionDispositionType31Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption133) AddOfferType() *OfferTypeFormat11Choice {
	newValue := new(OfferTypeFormat11Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption133) AddOptionFeatures() *OptionFeaturesFormat23Choice {
	newValue := new(OptionFeaturesFormat23Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption133) AddOptionAvailabilityStatus() *OptionAvailabilityStatus4Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus4Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption133) AddCertificationBreakdownType() *BeneficiaryCertificationType12Choice {
	newValue := new(BeneficiaryCertificationType12Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption133) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption133) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption133) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption133) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption133) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption133) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption133) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption133) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption133) SetAppliedOptionIndicator(value string) {
	c.AppliedOptionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption133) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption133) AddDateDetails() *CorporateActionDate55 {
	c.DateDetails = new(CorporateActionDate55)
	return c.DateDetails
}

func (c *CorporateActionOption133) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption133) AddRateAndAmountDetails() *CorporateActionRate83 {
	c.RateAndAmountDetails = new(CorporateActionRate83)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption133) AddPriceDetails() *CorporateActionPrice65 {
	c.PriceDetails = new(CorporateActionPrice65)
	return c.PriceDetails
}

func (c *CorporateActionOption133) AddSecuritiesQuantity() *SecuritiesOption56 {
	c.SecuritiesQuantity = new(SecuritiesOption56)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption133) AddSecuritiesMovementDetails() *SecuritiesOption61 {
	newValue := new(SecuritiesOption61)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption133) AddCashMovementDetails() *CashOption53 {
	newValue := new(CashOption53)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption133) AddAdditionalInformation() *CorporateActionNarrative36 {
	c.AdditionalInformation = new(CorporateActionNarrative36)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption134 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption29Choice `xml:"OptnTp"`

	// Specifies how fractional amount/quantities are treated.
	FractionDisposition *FractionDispositionType29Choice `xml:"FrctnDspstn,omitempty"`

	// Type of changes affecting the security form.
	ChangeType []*CorporateActionChangeTypeFormat7Choice `xml:"ChngTp,omitempty"`

	// Specifies that the corporate action instruction is to be processed using the Available-for-Collateral pool.
	EligibleForCollateralIndicator *YesNoIndicator `xml:"ElgblForCollInd,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds.
	CurrencyToBuy *ActiveCurrencyCode `xml:"CcyToBuy,omitempty"`

	// Account servicer is instructed to sell the indicated currency in order to obtain the necessary currency to fund the transaction within this instruction message.
	CurrencyToSell *ActiveCurrencyCode `xml:"CcyToSell,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification20 `xml:"SctyId,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantityOrInstructedAmount *SecuritiesQuantityOrAmount1Choice `xml:"SctiesQtyOrInstdAmt"`

	// Date/time at which the instructing party requests the instruction to be executed.
	ExecutionRequestedDateTime *DateAndDateTimeChoice `xml:"ExctnReqdDtTm,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate73 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice62 `xml:"PricDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative33 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption134) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption134) AddOptionType() *CorporateActionOption29Choice {
	c.OptionType = new(CorporateActionOption29Choice)
	return c.OptionType
}

func (c *CorporateActionOption134) AddFractionDisposition() *FractionDispositionType29Choice {
	c.FractionDisposition = new(FractionDispositionType29Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption134) AddChangeType() *CorporateActionChangeTypeFormat7Choice {
	newValue := new(CorporateActionChangeTypeFormat7Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateActionOption134) SetEligibleForCollateralIndicator(value string) {
	c.EligibleForCollateralIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption134) SetCurrencyToBuy(value string) {
	c.CurrencyToBuy = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption134) SetCurrencyToSell(value string) {
	c.CurrencyToSell = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption134) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption134) AddSecurityIdentification() *SecurityIdentification20 {
	c.SecurityIdentification = new(SecurityIdentification20)
	return c.SecurityIdentification
}

func (c *CorporateActionOption134) AddSecuritiesQuantityOrInstructedAmount() *SecuritiesQuantityOrAmount1Choice {
	c.SecuritiesQuantityOrInstructedAmount = new(SecuritiesQuantityOrAmount1Choice)
	return c.SecuritiesQuantityOrInstructedAmount
}

func (c *CorporateActionOption134) AddExecutionRequestedDateTime() *DateAndDateTimeChoice {
	c.ExecutionRequestedDateTime = new(DateAndDateTimeChoice)
	return c.ExecutionRequestedDateTime
}

func (c *CorporateActionOption134) AddRateAndAmountDetails() *CorporateActionRate73 {
	c.RateAndAmountDetails = new(CorporateActionRate73)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption134) AddPriceDetails() *CorporateActionPrice62 {
	c.PriceDetails = new(CorporateActionPrice62)
	return c.PriceDetails
}

func (c *CorporateActionOption134) AddAdditionalInformation() *CorporateActionNarrative33 {
	c.AdditionalInformation = new(CorporateActionNarrative33)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption135 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption23Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType31Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat11Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat23Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat18Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus4Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType12Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether redemption charges apply.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Indicates whether the option, different from the default one, shall be applied by the account owner.
	AppliedOptionIndicator *YesNoIndicator `xml:"ApldOptnInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification20 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate55 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate84 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice65 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity related to a corporate action option.
	SecuritiesQuantity *SecuritiesOption56 `xml:"SctiesQty,omitempty"`

	// Provides information about securities movement related to a corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption61 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption54 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information about the corporate action movement.
	AdditionalInformation *CorporateActionNarrative36 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption135) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption135) AddOptionType() *CorporateActionOption23Choice {
	c.OptionType = new(CorporateActionOption23Choice)
	return c.OptionType
}

func (c *CorporateActionOption135) AddFractionDisposition() *FractionDispositionType31Choice {
	c.FractionDisposition = new(FractionDispositionType31Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption135) AddOfferType() *OfferTypeFormat11Choice {
	newValue := new(OfferTypeFormat11Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption135) AddOptionFeatures() *OptionFeaturesFormat23Choice {
	newValue := new(OptionFeaturesFormat23Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption135) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat18Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat18Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption135) AddOptionAvailabilityStatus() *OptionAvailabilityStatus4Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus4Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption135) AddCertificationBreakdownType() *BeneficiaryCertificationType12Choice {
	newValue := new(BeneficiaryCertificationType12Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption135) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption135) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption135) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption135) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption135) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption135) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption135) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption135) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption135) SetAppliedOptionIndicator(value string) {
	c.AppliedOptionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption135) AddSecurityIdentification() *SecurityIdentification20 {
	c.SecurityIdentification = new(SecurityIdentification20)
	return c.SecurityIdentification
}

func (c *CorporateActionOption135) AddDateDetails() *CorporateActionDate55 {
	c.DateDetails = new(CorporateActionDate55)
	return c.DateDetails
}

func (c *CorporateActionOption135) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption135) AddRateAndAmountDetails() *CorporateActionRate84 {
	c.RateAndAmountDetails = new(CorporateActionRate84)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption135) AddPriceDetails() *CorporateActionPrice65 {
	c.PriceDetails = new(CorporateActionPrice65)
	return c.PriceDetails
}

func (c *CorporateActionOption135) AddSecuritiesQuantity() *SecuritiesOption56 {
	c.SecuritiesQuantity = new(SecuritiesOption56)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption135) AddSecuritiesMovementDetails() *SecuritiesOption61 {
	newValue := new(SecuritiesOption61)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption135) AddCashMovementDetails() *CashOption54 {
	newValue := new(CashOption54)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption135) AddAdditionalInformation() *CorporateActionNarrative36 {
	c.AdditionalInformation = new(CorporateActionNarrative36)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption136 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption26Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat19Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType30Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate52 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod11 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate85 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice63 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification90 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption63 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption55 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption136) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption136) AddOptionType() *CorporateActionOption26Choice {
	c.OptionType = new(CorporateActionOption26Choice)
	return c.OptionType
}

func (c *CorporateActionOption136) AddOptionFeatures() *OptionFeaturesFormat19Choice {
	newValue := new(OptionFeaturesFormat19Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption136) AddFractionDisposition() *FractionDispositionType30Choice {
	c.FractionDisposition = new(FractionDispositionType30Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption136) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption136) AddDateDetails() *CorporateActionDate52 {
	c.DateDetails = new(CorporateActionDate52)
	return c.DateDetails
}

func (c *CorporateActionOption136) AddPeriodDetails() *CorporateActionPeriod11 {
	c.PeriodDetails = new(CorporateActionPeriod11)
	return c.PeriodDetails
}

func (c *CorporateActionOption136) AddRateAndAmountDetails() *CorporateActionRate85 {
	c.RateAndAmountDetails = new(CorporateActionRate85)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption136) AddPriceDetails() *CorporateActionPrice63 {
	c.PriceDetails = new(CorporateActionPrice63)
	return c.PriceDetails
}

func (c *CorporateActionOption136) AddPlaceOfTrade() *MarketIdentification90 {
	c.PlaceOfTrade = new(MarketIdentification90)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption136) AddSecuritiesMovementDetails() *SecuritiesOption63 {
	newValue := new(SecuritiesOption63)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption136) AddCashMovementDetails() *CashOption55 {
	newValue := new(CashOption55)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption19 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption2Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat1Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat5Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus1Choice `xml:"OptnAvlbtySts,omitempty"`

	// Type of certification which is required.
	CertificationType []*BeneficiaryCertificationType1Choice `xml:"CertfctnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Whether or not certification is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate15 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate15 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice16 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption15 `xml:"SctiesQty,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption14 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption10 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative5 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption19) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption19) AddOptionType() *CorporateActionOption2Choice {
	c.OptionType = new(CorporateActionOption2Choice)
	return c.OptionType
}

func (c *CorporateActionOption19) AddFractionDisposition() *FractionDispositionType1Choice {
	c.FractionDisposition = new(FractionDispositionType1Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption19) AddOfferType() *OfferTypeFormat1Choice {
	newValue := new(OfferTypeFormat1Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption19) AddOptionFeatures() *OptionFeaturesFormat5Choice {
	newValue := new(OptionFeaturesFormat5Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption19) AddOptionAvailabilityStatus() *OptionAvailabilityStatus1Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus1Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption19) AddCertificationType() *BeneficiaryCertificationType1Choice {
	newValue := new(BeneficiaryCertificationType1Choice)
	c.CertificationType = append(c.CertificationType, newValue)
	return newValue
}

func (c *CorporateActionOption19) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption19) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption19) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption19) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption19) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption19) SetCertificationIndicator(value string) {
	c.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption19) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption19) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption19) AddSecurityIdentification() *SecurityIdentification14 {
	c.SecurityIdentification = new(SecurityIdentification14)
	return c.SecurityIdentification
}

func (c *CorporateActionOption19) AddDateDetails() *CorporateActionDate15 {
	c.DateDetails = new(CorporateActionDate15)
	return c.DateDetails
}

func (c *CorporateActionOption19) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption19) AddRateAndAmountDetails() *CorporateActionRate15 {
	c.RateAndAmountDetails = new(CorporateActionRate15)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption19) AddPriceDetails() *CorporateActionPrice16 {
	c.PriceDetails = new(CorporateActionPrice16)
	return c.PriceDetails
}

func (c *CorporateActionOption19) AddSecuritiesQuantity() *SecuritiesOption15 {
	c.SecuritiesQuantity = new(SecuritiesOption15)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption19) AddSecuritiesMovementDetails() *SecuritiesOption14 {
	newValue := new(SecuritiesOption14)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption19) AddCashMovementDetails() *CashOption10 {
	newValue := new(CashOption10)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption19) AddAdditionalInformation() *CorporateActionNarrative5 {
	c.AdditionalInformation = new(CorporateActionNarrative5)
	return c.AdditionalInformation
}

// Provides information on a CA option.
type CorporateActionOption2 struct {

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`
}

func (c *CorporateActionOption2) AddOptionType() *CorporateActionOption1FormatChoice {
	c.OptionType = new(CorporateActionOption1FormatChoice)
	return c.OptionType
}

func (c *CorporateActionOption2) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

// Provides information about the corporate action option.
type CorporateActionOption20 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption2Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat1Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat6Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat6Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus1Choice `xml:"OptnAvlbtySts,omitempty"`

	// Type of certification which is required.
	CertificationType []*BeneficiaryCertificationType1Choice `xml:"CertfctnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether redemption charges apply.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether or not certification is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate15 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate15 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice16 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity related to a corporate action option.
	SecuritiesQuantity *SecuritiesOption15 `xml:"SctiesQty,omitempty"`

	// Provides information about securities movement related to a corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption13 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption11 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information about the corporate action movement.
	AdditionalInformation *CorporateActionNarrative5 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption20) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption20) AddOptionType() *CorporateActionOption2Choice {
	c.OptionType = new(CorporateActionOption2Choice)
	return c.OptionType
}

func (c *CorporateActionOption20) AddFractionDisposition() *FractionDispositionType1Choice {
	c.FractionDisposition = new(FractionDispositionType1Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption20) AddOfferType() *OfferTypeFormat1Choice {
	newValue := new(OfferTypeFormat1Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption20) AddOptionFeatures() *OptionFeaturesFormat6Choice {
	newValue := new(OptionFeaturesFormat6Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption20) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat6Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat6Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption20) AddOptionAvailabilityStatus() *OptionAvailabilityStatus1Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus1Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption20) AddCertificationType() *BeneficiaryCertificationType1Choice {
	newValue := new(BeneficiaryCertificationType1Choice)
	c.CertificationType = append(c.CertificationType, newValue)
	return newValue
}

func (c *CorporateActionOption20) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption20) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption20) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption20) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption20) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption20) SetCertificationIndicator(value string) {
	c.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption20) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption20) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption20) AddSecurityIdentification() *SecurityIdentification14 {
	c.SecurityIdentification = new(SecurityIdentification14)
	return c.SecurityIdentification
}

func (c *CorporateActionOption20) AddDateDetails() *CorporateActionDate15 {
	c.DateDetails = new(CorporateActionDate15)
	return c.DateDetails
}

func (c *CorporateActionOption20) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption20) AddRateAndAmountDetails() *CorporateActionRate15 {
	c.RateAndAmountDetails = new(CorporateActionRate15)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption20) AddPriceDetails() *CorporateActionPrice16 {
	c.PriceDetails = new(CorporateActionPrice16)
	return c.PriceDetails
}

func (c *CorporateActionOption20) AddSecuritiesQuantity() *SecuritiesOption15 {
	c.SecuritiesQuantity = new(SecuritiesOption15)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption20) AddSecuritiesMovementDetails() *SecuritiesOption13 {
	newValue := new(SecuritiesOption13)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption20) AddCashMovementDetails() *CashOption11 {
	newValue := new(CashOption11)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption20) AddAdditionalInformation() *CorporateActionNarrative5 {
	c.AdditionalInformation = new(CorporateActionNarrative5)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption21 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption4Choice `xml:"OptnTp"`

	// Specifies whether the quantity of financial instrument is a quantity of securities instructed or a quantity to receive.
	InstructedOrQuantityToReceive *InstructedOrQuantityToReceive1Choice `xml:"InstdOrQtyToRcv"`
}

func (c *CorporateActionOption21) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption21) AddOptionType() *CorporateActionOption4Choice {
	c.OptionType = new(CorporateActionOption4Choice)
	return c.OptionType
}

func (c *CorporateActionOption21) AddInstructedOrQuantityToReceive() *InstructedOrQuantityToReceive1Choice {
	c.InstructedOrQuantityToReceive = new(InstructedOrQuantityToReceive1Choice)
	return c.InstructedOrQuantityToReceive
}

// Provides information about the corporate action option.
type CorporateActionOption22 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption5Choice `xml:"OptnTp"`

	// Party that owns the account.
	AccountOwner *PartyIdentification41Choice `xml:"AcctOwnr,omitempty"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Account on which a securities entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *SignedQuantityFormat1 `xml:"TtlElgblBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *SignedQuantityFormat1 `xml:"InstdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *SignedQuantityFormat1 `xml:"UinstdBal,omitempty"`

	// Specifies whether the quantity of financial instrument is a status quantity or a quantity to receive.
	StatusQuantityOrQuantityToReceive *StatusOrQuantityToReceive1Choice `xml:"StsQtyOrQtyToRcv,omitempty"`
}

func (c *CorporateActionOption22) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption22) AddOptionType() *CorporateActionOption5Choice {
	c.OptionType = new(CorporateActionOption5Choice)
	return c.OptionType
}

func (c *CorporateActionOption22) AddAccountOwner() *PartyIdentification41Choice {
	c.AccountOwner = new(PartyIdentification41Choice)
	return c.AccountOwner
}

func (c *CorporateActionOption22) SetSafekeepingAccount(value string) {
	c.SafekeepingAccount = (*Max35Text)(&value)
}

func (c *CorporateActionOption22) AddCashAccount() *CashAccountIdentification5Choice {
	c.CashAccount = new(CashAccountIdentification5Choice)
	return c.CashAccount
}

func (c *CorporateActionOption22) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	c.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return c.SafekeepingPlace
}

func (c *CorporateActionOption22) AddSecurityIdentification() *SecurityIdentification14 {
	c.SecurityIdentification = new(SecurityIdentification14)
	return c.SecurityIdentification
}

func (c *CorporateActionOption22) AddTotalEligibleBalance() *SignedQuantityFormat1 {
	c.TotalEligibleBalance = new(SignedQuantityFormat1)
	return c.TotalEligibleBalance
}

func (c *CorporateActionOption22) AddInstructedBalance() *SignedQuantityFormat1 {
	c.InstructedBalance = new(SignedQuantityFormat1)
	return c.InstructedBalance
}

func (c *CorporateActionOption22) AddUninstructedBalance() *SignedQuantityFormat1 {
	c.UninstructedBalance = new(SignedQuantityFormat1)
	return c.UninstructedBalance
}

func (c *CorporateActionOption22) AddStatusQuantityOrQuantityToReceive() *StatusOrQuantityToReceive1Choice {
	c.StatusQuantityOrQuantityToReceive = new(StatusOrQuantityToReceive1Choice)
	return c.StatusQuantityOrQuantityToReceive
}

// Provides information about the corporate action option.
type CorporateActionOption23 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption3Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat1Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType12Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate18 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod9 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate20 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice21 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption18 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption12 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption23) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption23) AddOptionType() *CorporateActionOption3Choice {
	c.OptionType = new(CorporateActionOption3Choice)
	return c.OptionType
}

func (c *CorporateActionOption23) AddOptionFeatures() *OptionFeaturesFormat1Choice {
	newValue := new(OptionFeaturesFormat1Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption23) AddFractionDisposition() *FractionDispositionType12Choice {
	c.FractionDisposition = new(FractionDispositionType12Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption23) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption23) AddDateDetails() *CorporateActionDate18 {
	c.DateDetails = new(CorporateActionDate18)
	return c.DateDetails
}

func (c *CorporateActionOption23) AddPeriodDetails() *CorporateActionPeriod9 {
	c.PeriodDetails = new(CorporateActionPeriod9)
	return c.PeriodDetails
}

func (c *CorporateActionOption23) AddRateAndAmountDetails() *CorporateActionRate20 {
	c.RateAndAmountDetails = new(CorporateActionRate20)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption23) AddPriceDetails() *CorporateActionPrice21 {
	c.PriceDetails = new(CorporateActionPrice21)
	return c.PriceDetails
}

func (c *CorporateActionOption23) AddPlaceOfTrade() *MarketIdentification4 {
	c.PlaceOfTrade = new(MarketIdentification4)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption23) AddSecuritiesMovementDetails() *SecuritiesOption18 {
	newValue := new(SecuritiesOption18)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption23) AddCashMovementDetails() *CashOption12 {
	newValue := new(CashOption12)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption24 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption3Choice `xml:"OptnTp"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption19 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption5 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption24) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption24) AddOptionType() *CorporateActionOption3Choice {
	c.OptionType = new(CorporateActionOption3Choice)
	return c.OptionType
}

func (c *CorporateActionOption24) AddSecuritiesMovementDetails() *SecuritiesOption19 {
	newValue := new(SecuritiesOption19)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption24) AddCashMovementDetails() *CashOption5 {
	newValue := new(CashOption5)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption25 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption4Choice `xml:"OptnTp"`

	// Specifies how fractional amount/quantities are treated.
	FractionDisposition *FractionDispositionType10Choice `xml:"FrctnDspstn,omitempty"`

	// Type of changes affecting the security form.
	ChangeType []*CorporateActionChangeTypeFormat2Choice `xml:"ChngTp,omitempty"`

	// Specifies that the corporate action instruction is to be processed using the Available-for-Collateral pool.
	EligibleForCollateralIndicator *YesNoIndicator `xml:"ElgblForCollInd,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds.
	CurrencyToBuy *ActiveCurrencyCode `xml:"CcyToBuy,omitempty"`

	// Account servicer is instructed to sell the indicated currency in order to obtain the necessary currency to fund the transaction within this instruction message.
	CurrencyToSell *ActiveCurrencyCode `xml:"CcyToSell,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption2 `xml:"SctiesQty"`

	// Date/time at which the instructing party requests the instruction to be executed.
	ExecutionRequestedDateTime *DateAndDateTimeChoice `xml:"ExctnReqdDtTm,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate8 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice19 `xml:"PricDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative8 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption25) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption25) AddOptionType() *CorporateActionOption4Choice {
	c.OptionType = new(CorporateActionOption4Choice)
	return c.OptionType
}

func (c *CorporateActionOption25) AddFractionDisposition() *FractionDispositionType10Choice {
	c.FractionDisposition = new(FractionDispositionType10Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption25) AddChangeType() *CorporateActionChangeTypeFormat2Choice {
	newValue := new(CorporateActionChangeTypeFormat2Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateActionOption25) SetEligibleForCollateralIndicator(value string) {
	c.EligibleForCollateralIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption25) SetCurrencyToBuy(value string) {
	c.CurrencyToBuy = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption25) SetCurrencyToSell(value string) {
	c.CurrencyToSell = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption25) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption25) AddSecurityIdentification() *SecurityIdentification14 {
	c.SecurityIdentification = new(SecurityIdentification14)
	return c.SecurityIdentification
}

func (c *CorporateActionOption25) AddSecuritiesQuantity() *SecuritiesOption2 {
	c.SecuritiesQuantity = new(SecuritiesOption2)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption25) AddExecutionRequestedDateTime() *DateAndDateTimeChoice {
	c.ExecutionRequestedDateTime = new(DateAndDateTimeChoice)
	return c.ExecutionRequestedDateTime
}

func (c *CorporateActionOption25) AddRateAndAmountDetails() *CorporateActionRate8 {
	c.RateAndAmountDetails = new(CorporateActionRate8)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption25) AddPriceDetails() *CorporateActionPrice19 {
	c.PriceDetails = new(CorporateActionPrice19)
	return c.PriceDetails
}

func (c *CorporateActionOption25) AddAdditionalInformation() *CorporateActionNarrative8 {
	c.AdditionalInformation = new(CorporateActionNarrative8)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption26 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption5Choice `xml:"OptnTp"`

	// Party that owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Account on which a securities entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *SignedQuantityFormat1 `xml:"TtlElgblBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *SignedQuantityFormat1 `xml:"InstdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *SignedQuantityFormat1 `xml:"UinstdBal,omitempty"`

	// Specifies whether the quantity of financial instrument is a status quantity or a quantity to receive.
	StatusQuantityOrQuantityToReceive *StatusOrQuantityToReceive1Choice `xml:"StsQtyOrQtyToRcv,omitempty"`
}

func (c *CorporateActionOption26) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption26) AddOptionType() *CorporateActionOption5Choice {
	c.OptionType = new(CorporateActionOption5Choice)
	return c.OptionType
}

func (c *CorporateActionOption26) AddAccountOwner() *PartyIdentification36Choice {
	c.AccountOwner = new(PartyIdentification36Choice)
	return c.AccountOwner
}

func (c *CorporateActionOption26) SetSafekeepingAccount(value string) {
	c.SafekeepingAccount = (*Max35Text)(&value)
}

func (c *CorporateActionOption26) AddCashAccount() *CashAccountIdentification5Choice {
	c.CashAccount = new(CashAccountIdentification5Choice)
	return c.CashAccount
}

func (c *CorporateActionOption26) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	c.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return c.SafekeepingPlace
}

func (c *CorporateActionOption26) AddSecurityIdentification() *SecurityIdentification14 {
	c.SecurityIdentification = new(SecurityIdentification14)
	return c.SecurityIdentification
}

func (c *CorporateActionOption26) AddTotalEligibleBalance() *SignedQuantityFormat1 {
	c.TotalEligibleBalance = new(SignedQuantityFormat1)
	return c.TotalEligibleBalance
}

func (c *CorporateActionOption26) AddInstructedBalance() *SignedQuantityFormat1 {
	c.InstructedBalance = new(SignedQuantityFormat1)
	return c.InstructedBalance
}

func (c *CorporateActionOption26) AddUninstructedBalance() *SignedQuantityFormat1 {
	c.UninstructedBalance = new(SignedQuantityFormat1)
	return c.UninstructedBalance
}

func (c *CorporateActionOption26) AddStatusQuantityOrQuantityToReceive() *StatusOrQuantityToReceive1Choice {
	c.StatusQuantityOrQuantityToReceive = new(StatusOrQuantityToReceive1Choice)
	return c.StatusQuantityOrQuantityToReceive
}

// Provides information about the corporate action option.
type CorporateActionOption3 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption2Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat1Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat2Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat2Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus1Choice `xml:"OptnAvlbtySts,omitempty"`

	// Type of certification which is required.
	CertificationType []*BeneficiaryCertificationType1Choice `xml:"CertfctnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Whether or not certification is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate8 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod5 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate5 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice6 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption1 `xml:"SctiesQty,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption6 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption4 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative5 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption3) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption3) AddOptionType() *CorporateActionOption2Choice {
	c.OptionType = new(CorporateActionOption2Choice)
	return c.OptionType
}

func (c *CorporateActionOption3) AddFractionDisposition() *FractionDispositionType1Choice {
	c.FractionDisposition = new(FractionDispositionType1Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption3) AddOfferType() *OfferTypeFormat1Choice {
	newValue := new(OfferTypeFormat1Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption3) AddOptionFeatures() *OptionFeaturesFormat2Choice {
	newValue := new(OptionFeaturesFormat2Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption3) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat2Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat2Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption3) AddOptionAvailabilityStatus() *OptionAvailabilityStatus1Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus1Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption3) AddCertificationType() *BeneficiaryCertificationType1Choice {
	newValue := new(BeneficiaryCertificationType1Choice)
	c.CertificationType = append(c.CertificationType, newValue)
	return newValue
}

func (c *CorporateActionOption3) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption3) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption3) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption3) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption3) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption3) SetCertificationIndicator(value string) {
	c.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption3) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption3) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption3) AddSecurityIdentification() *SecurityIdentification11 {
	c.SecurityIdentification = new(SecurityIdentification11)
	return c.SecurityIdentification
}

func (c *CorporateActionOption3) AddDateDetails() *CorporateActionDate8 {
	c.DateDetails = new(CorporateActionDate8)
	return c.DateDetails
}

func (c *CorporateActionOption3) AddPeriodDetails() *CorporateActionPeriod5 {
	c.PeriodDetails = new(CorporateActionPeriod5)
	return c.PeriodDetails
}

func (c *CorporateActionOption3) AddRateAndAmountDetails() *CorporateActionRate5 {
	c.RateAndAmountDetails = new(CorporateActionRate5)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption3) AddPriceDetails() *CorporateActionPrice6 {
	c.PriceDetails = new(CorporateActionPrice6)
	return c.PriceDetails
}

func (c *CorporateActionOption3) AddSecuritiesQuantity() *SecuritiesOption1 {
	c.SecuritiesQuantity = new(SecuritiesOption1)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption3) AddSecuritiesMovementDetails() *SecuritiesOption6 {
	newValue := new(SecuritiesOption6)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption3) AddCashMovementDetails() *CashOption4 {
	newValue := new(CashOption4)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption3) AddAdditionalInformation() *CorporateActionNarrative5 {
	c.AdditionalInformation = new(CorporateActionNarrative5)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption35 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption10Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat1Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat9Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat9Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus1Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType5Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether redemption charges apply.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate15 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate25 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice28 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity related to a corporate action option.
	SecuritiesQuantity *SecuritiesOption23 `xml:"SctiesQty,omitempty"`

	// Provides information about securities movement related to a corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption24 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption16 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information about the corporate action movement.
	AdditionalInformation *CorporateActionNarrative20 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption35) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption35) AddOptionType() *CorporateActionOption10Choice {
	c.OptionType = new(CorporateActionOption10Choice)
	return c.OptionType
}

func (c *CorporateActionOption35) AddFractionDisposition() *FractionDispositionType1Choice {
	c.FractionDisposition = new(FractionDispositionType1Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption35) AddOfferType() *OfferTypeFormat1Choice {
	newValue := new(OfferTypeFormat1Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption35) AddOptionFeatures() *OptionFeaturesFormat9Choice {
	newValue := new(OptionFeaturesFormat9Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption35) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat9Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat9Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption35) AddOptionAvailabilityStatus() *OptionAvailabilityStatus1Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus1Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption35) AddCertificationBreakdownType() *BeneficiaryCertificationType5Choice {
	newValue := new(BeneficiaryCertificationType5Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption35) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption35) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption35) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption35) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption35) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption35) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption35) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption35) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption35) AddSecurityIdentification() *SecurityIdentification14 {
	c.SecurityIdentification = new(SecurityIdentification14)
	return c.SecurityIdentification
}

func (c *CorporateActionOption35) AddDateDetails() *CorporateActionDate15 {
	c.DateDetails = new(CorporateActionDate15)
	return c.DateDetails
}

func (c *CorporateActionOption35) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption35) AddRateAndAmountDetails() *CorporateActionRate25 {
	c.RateAndAmountDetails = new(CorporateActionRate25)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption35) AddPriceDetails() *CorporateActionPrice28 {
	c.PriceDetails = new(CorporateActionPrice28)
	return c.PriceDetails
}

func (c *CorporateActionOption35) AddSecuritiesQuantity() *SecuritiesOption23 {
	c.SecuritiesQuantity = new(SecuritiesOption23)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption35) AddSecuritiesMovementDetails() *SecuritiesOption24 {
	newValue := new(SecuritiesOption24)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption35) AddCashMovementDetails() *CashOption16 {
	newValue := new(CashOption16)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption35) AddAdditionalInformation() *CorporateActionNarrative20 {
	c.AdditionalInformation = new(CorporateActionNarrative20)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption36 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption10Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat1Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat9Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus1Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType5Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate15 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate25 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice28 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption23 `xml:"SctiesQty,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption25 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption17 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative20 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption36) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption36) AddOptionType() *CorporateActionOption10Choice {
	c.OptionType = new(CorporateActionOption10Choice)
	return c.OptionType
}

func (c *CorporateActionOption36) AddFractionDisposition() *FractionDispositionType1Choice {
	c.FractionDisposition = new(FractionDispositionType1Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption36) AddOfferType() *OfferTypeFormat1Choice {
	newValue := new(OfferTypeFormat1Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption36) AddOptionFeatures() *OptionFeaturesFormat9Choice {
	newValue := new(OptionFeaturesFormat9Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption36) AddOptionAvailabilityStatus() *OptionAvailabilityStatus1Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus1Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption36) AddCertificationBreakdownType() *BeneficiaryCertificationType5Choice {
	newValue := new(BeneficiaryCertificationType5Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption36) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption36) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption36) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption36) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption36) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption36) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption36) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption36) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption36) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption36) AddDateDetails() *CorporateActionDate15 {
	c.DateDetails = new(CorporateActionDate15)
	return c.DateDetails
}

func (c *CorporateActionOption36) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption36) AddRateAndAmountDetails() *CorporateActionRate25 {
	c.RateAndAmountDetails = new(CorporateActionRate25)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption36) AddPriceDetails() *CorporateActionPrice28 {
	c.PriceDetails = new(CorporateActionPrice28)
	return c.PriceDetails
}

func (c *CorporateActionOption36) AddSecuritiesQuantity() *SecuritiesOption23 {
	c.SecuritiesQuantity = new(SecuritiesOption23)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption36) AddSecuritiesMovementDetails() *SecuritiesOption25 {
	newValue := new(SecuritiesOption25)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption36) AddCashMovementDetails() *CashOption17 {
	newValue := new(CashOption17)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption36) AddAdditionalInformation() *CorporateActionNarrative20 {
	c.AdditionalInformation = new(CorporateActionNarrative20)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption37 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption11Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat1Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType12Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate18 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod9 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate26 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice30 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption26 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption18 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption37) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption37) AddOptionType() *CorporateActionOption11Choice {
	c.OptionType = new(CorporateActionOption11Choice)
	return c.OptionType
}

func (c *CorporateActionOption37) AddOptionFeatures() *OptionFeaturesFormat1Choice {
	newValue := new(OptionFeaturesFormat1Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption37) AddFractionDisposition() *FractionDispositionType12Choice {
	c.FractionDisposition = new(FractionDispositionType12Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption37) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption37) AddDateDetails() *CorporateActionDate18 {
	c.DateDetails = new(CorporateActionDate18)
	return c.DateDetails
}

func (c *CorporateActionOption37) AddPeriodDetails() *CorporateActionPeriod9 {
	c.PeriodDetails = new(CorporateActionPeriod9)
	return c.PeriodDetails
}

func (c *CorporateActionOption37) AddRateAndAmountDetails() *CorporateActionRate26 {
	c.RateAndAmountDetails = new(CorporateActionRate26)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption37) AddPriceDetails() *CorporateActionPrice30 {
	c.PriceDetails = new(CorporateActionPrice30)
	return c.PriceDetails
}

func (c *CorporateActionOption37) AddPlaceOfTrade() *MarketIdentification4 {
	c.PlaceOfTrade = new(MarketIdentification4)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption37) AddSecuritiesMovementDetails() *SecuritiesOption26 {
	newValue := new(SecuritiesOption26)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption37) AddCashMovementDetails() *CashOption18 {
	newValue := new(CashOption18)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption38 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption12Choice `xml:"OptnTp"`

	// Specifies how fractional amount/quantities are treated.
	FractionDisposition *FractionDispositionType10Choice `xml:"FrctnDspstn,omitempty"`

	// Type of changes affecting the security form.
	ChangeType []*CorporateActionChangeTypeFormat2Choice `xml:"ChngTp,omitempty"`

	// Specifies that the corporate action instruction is to be processed using the Available-for-Collateral pool.
	EligibleForCollateralIndicator *YesNoIndicator `xml:"ElgblForCollInd,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds.
	CurrencyToBuy *ActiveCurrencyCode `xml:"CcyToBuy,omitempty"`

	// Account servicer is instructed to sell the indicated currency in order to obtain the necessary currency to fund the transaction within this instruction message.
	CurrencyToSell *ActiveCurrencyCode `xml:"CcyToSell,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption2 `xml:"SctiesQty"`

	// Date/time at which the instructing party requests the instruction to be executed.
	ExecutionRequestedDateTime *DateAndDateTimeChoice `xml:"ExctnReqdDtTm,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate8 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice29 `xml:"PricDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative8 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption38) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption38) AddOptionType() *CorporateActionOption12Choice {
	c.OptionType = new(CorporateActionOption12Choice)
	return c.OptionType
}

func (c *CorporateActionOption38) AddFractionDisposition() *FractionDispositionType10Choice {
	c.FractionDisposition = new(FractionDispositionType10Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption38) AddChangeType() *CorporateActionChangeTypeFormat2Choice {
	newValue := new(CorporateActionChangeTypeFormat2Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateActionOption38) SetEligibleForCollateralIndicator(value string) {
	c.EligibleForCollateralIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption38) SetCurrencyToBuy(value string) {
	c.CurrencyToBuy = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption38) SetCurrencyToSell(value string) {
	c.CurrencyToSell = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption38) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption38) AddSecurityIdentification() *SecurityIdentification14 {
	c.SecurityIdentification = new(SecurityIdentification14)
	return c.SecurityIdentification
}

func (c *CorporateActionOption38) AddSecuritiesQuantity() *SecuritiesOption2 {
	c.SecuritiesQuantity = new(SecuritiesOption2)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption38) AddExecutionRequestedDateTime() *DateAndDateTimeChoice {
	c.ExecutionRequestedDateTime = new(DateAndDateTimeChoice)
	return c.ExecutionRequestedDateTime
}

func (c *CorporateActionOption38) AddRateAndAmountDetails() *CorporateActionRate8 {
	c.RateAndAmountDetails = new(CorporateActionRate8)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption38) AddPriceDetails() *CorporateActionPrice29 {
	c.PriceDetails = new(CorporateActionPrice29)
	return c.PriceDetails
}

func (c *CorporateActionOption38) AddAdditionalInformation() *CorporateActionNarrative8 {
	c.AdditionalInformation = new(CorporateActionNarrative8)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption39 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption11Choice `xml:"OptnTp"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption27 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption19 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption39) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption39) AddOptionType() *CorporateActionOption11Choice {
	c.OptionType = new(CorporateActionOption11Choice)
	return c.OptionType
}

func (c *CorporateActionOption39) AddSecuritiesMovementDetails() *SecuritiesOption27 {
	newValue := new(SecuritiesOption27)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption39) AddCashMovementDetails() *CashOption19 {
	newValue := new(CashOption19)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption4 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption3Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat1Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate6 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod4 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate4 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice5 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption3 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption2 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption4) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption4) AddOptionType() *CorporateActionOption3Choice {
	c.OptionType = new(CorporateActionOption3Choice)
	return c.OptionType
}

func (c *CorporateActionOption4) AddOptionFeatures() *OptionFeaturesFormat1Choice {
	newValue := new(OptionFeaturesFormat1Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption4) AddFractionDisposition() *FractionDispositionType1Choice {
	c.FractionDisposition = new(FractionDispositionType1Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption4) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption4) AddDateDetails() *CorporateActionDate6 {
	c.DateDetails = new(CorporateActionDate6)
	return c.DateDetails
}

func (c *CorporateActionOption4) AddPeriodDetails() *CorporateActionPeriod4 {
	c.PeriodDetails = new(CorporateActionPeriod4)
	return c.PeriodDetails
}

func (c *CorporateActionOption4) AddRateAndAmountDetails() *CorporateActionRate4 {
	c.RateAndAmountDetails = new(CorporateActionRate4)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption4) AddPriceDetails() *CorporateActionPrice5 {
	c.PriceDetails = new(CorporateActionPrice5)
	return c.PriceDetails
}

func (c *CorporateActionOption4) AddPlaceOfTrade() *MarketIdentification4 {
	c.PlaceOfTrade = new(MarketIdentification4)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption4) AddSecuritiesMovementDetails() *SecuritiesOption3 {
	newValue := new(SecuritiesOption3)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption4) AddCashMovementDetails() *CashOption2 {
	newValue := new(CashOption2)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption41 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption13Choice `xml:"OptnTp"`

	// Party that owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Account on which a securities entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *SignedQuantityFormat1 `xml:"TtlElgblBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *SignedQuantityFormat1 `xml:"InstdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *SignedQuantityFormat1 `xml:"UinstdBal,omitempty"`

	// Specifies whether the quantity of financial instrument is a status quantity or a quantity to receive.
	StatusQuantityOrQuantityToReceive *StatusOrQuantityToReceive1Choice `xml:"StsQtyOrQtyToRcv,omitempty"`
}

func (c *CorporateActionOption41) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption41) AddOptionType() *CorporateActionOption13Choice {
	c.OptionType = new(CorporateActionOption13Choice)
	return c.OptionType
}

func (c *CorporateActionOption41) AddAccountOwner() *PartyIdentification36Choice {
	c.AccountOwner = new(PartyIdentification36Choice)
	return c.AccountOwner
}

func (c *CorporateActionOption41) SetSafekeepingAccount(value string) {
	c.SafekeepingAccount = (*Max35Text)(&value)
}

func (c *CorporateActionOption41) AddCashAccount() *CashAccountIdentification5Choice {
	c.CashAccount = new(CashAccountIdentification5Choice)
	return c.CashAccount
}

func (c *CorporateActionOption41) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	c.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return c.SafekeepingPlace
}

func (c *CorporateActionOption41) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption41) AddTotalEligibleBalance() *SignedQuantityFormat1 {
	c.TotalEligibleBalance = new(SignedQuantityFormat1)
	return c.TotalEligibleBalance
}

func (c *CorporateActionOption41) AddInstructedBalance() *SignedQuantityFormat1 {
	c.InstructedBalance = new(SignedQuantityFormat1)
	return c.InstructedBalance
}

func (c *CorporateActionOption41) AddUninstructedBalance() *SignedQuantityFormat1 {
	c.UninstructedBalance = new(SignedQuantityFormat1)
	return c.UninstructedBalance
}

func (c *CorporateActionOption41) AddStatusQuantityOrQuantityToReceive() *StatusOrQuantityToReceive1Choice {
	c.StatusQuantityOrQuantityToReceive = new(StatusOrQuantityToReceive1Choice)
	return c.StatusQuantityOrQuantityToReceive
}

// Provides information about the corporate action option.
type CorporateActionOption42 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption12Choice `xml:"OptnTp"`

	// Specifies whether the quantity of financial instrument is a quantity of securities instructed or a quantity to receive.
	InstructedOrQuantityToReceive *InstructedOrQuantityToReceive1Choice `xml:"InstdOrQtyToRcv"`
}

func (c *CorporateActionOption42) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption42) AddOptionType() *CorporateActionOption12Choice {
	c.OptionType = new(CorporateActionOption12Choice)
	return c.OptionType
}

func (c *CorporateActionOption42) AddInstructedOrQuantityToReceive() *InstructedOrQuantityToReceive1Choice {
	c.InstructedOrQuantityToReceive = new(InstructedOrQuantityToReceive1Choice)
	return c.InstructedOrQuantityToReceive
}

// Provides information about the corporate action option.
type CorporateActionOption5 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption4Choice `xml:"OptnTp"`

	// Specifies how fractional amount/quantities are treated.
	FractionDisposition *FractionDispositionType2Choice `xml:"FrctnDspstn,omitempty"`

	// Type of changes affecting the security form.
	ChangeType []*CorporateActionChangeTypeFormat2Choice `xml:"ChngTp,omitempty"`

	// Specifies that the corporate action instruction is to be processed using the Available-for-Collateral pool.
	EligibleForCollateralIndicator *YesNoIndicator `xml:"ElgblForCollInd,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds.
	CurrencyToBuy *ActiveCurrencyCode `xml:"CcyToBuy,omitempty"`

	// Account servicer is instructed to sell the indicated currency in order to obtain the necessary currency to fund the transaction within this instruction message.
	CurrencyToSell *ActiveCurrencyCode `xml:"CcyToSell,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption2 `xml:"SctiesQty"`

	// Date/time at which the instructing party requests the instruction to be executed.
	ExecutionRequestedDateTime *DateAndDateTimeChoice `xml:"ExctnReqdDtTm,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate8 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice8 `xml:"PricDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative8 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption5) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption5) AddOptionType() *CorporateActionOption4Choice {
	c.OptionType = new(CorporateActionOption4Choice)
	return c.OptionType
}

func (c *CorporateActionOption5) AddFractionDisposition() *FractionDispositionType2Choice {
	c.FractionDisposition = new(FractionDispositionType2Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption5) AddChangeType() *CorporateActionChangeTypeFormat2Choice {
	newValue := new(CorporateActionChangeTypeFormat2Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateActionOption5) SetEligibleForCollateralIndicator(value string) {
	c.EligibleForCollateralIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption5) SetCurrencyToBuy(value string) {
	c.CurrencyToBuy = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption5) SetCurrencyToSell(value string) {
	c.CurrencyToSell = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption5) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption5) AddSecurityIdentification() *SecurityIdentification11 {
	c.SecurityIdentification = new(SecurityIdentification11)
	return c.SecurityIdentification
}

func (c *CorporateActionOption5) AddSecuritiesQuantity() *SecuritiesOption2 {
	c.SecuritiesQuantity = new(SecuritiesOption2)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption5) AddExecutionRequestedDateTime() *DateAndDateTimeChoice {
	c.ExecutionRequestedDateTime = new(DateAndDateTimeChoice)
	return c.ExecutionRequestedDateTime
}

func (c *CorporateActionOption5) AddRateAndAmountDetails() *CorporateActionRate8 {
	c.RateAndAmountDetails = new(CorporateActionRate8)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption5) AddPriceDetails() *CorporateActionPrice8 {
	c.PriceDetails = new(CorporateActionPrice8)
	return c.PriceDetails
}

func (c *CorporateActionOption5) AddAdditionalInformation() *CorporateActionNarrative8 {
	c.AdditionalInformation = new(CorporateActionNarrative8)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption51 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption10Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType19Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat3Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat9Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus1Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType5Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate15 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate36 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice28 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption23 `xml:"SctiesQty,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption33 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption24 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative20 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption51) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption51) AddOptionType() *CorporateActionOption10Choice {
	c.OptionType = new(CorporateActionOption10Choice)
	return c.OptionType
}

func (c *CorporateActionOption51) AddFractionDisposition() *FractionDispositionType19Choice {
	c.FractionDisposition = new(FractionDispositionType19Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption51) AddOfferType() *OfferTypeFormat3Choice {
	newValue := new(OfferTypeFormat3Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption51) AddOptionFeatures() *OptionFeaturesFormat9Choice {
	newValue := new(OptionFeaturesFormat9Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption51) AddOptionAvailabilityStatus() *OptionAvailabilityStatus1Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus1Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption51) AddCertificationBreakdownType() *BeneficiaryCertificationType5Choice {
	newValue := new(BeneficiaryCertificationType5Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption51) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption51) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption51) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption51) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption51) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption51) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption51) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption51) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption51) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption51) AddDateDetails() *CorporateActionDate15 {
	c.DateDetails = new(CorporateActionDate15)
	return c.DateDetails
}

func (c *CorporateActionOption51) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption51) AddRateAndAmountDetails() *CorporateActionRate36 {
	c.RateAndAmountDetails = new(CorporateActionRate36)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption51) AddPriceDetails() *CorporateActionPrice28 {
	c.PriceDetails = new(CorporateActionPrice28)
	return c.PriceDetails
}

func (c *CorporateActionOption51) AddSecuritiesQuantity() *SecuritiesOption23 {
	c.SecuritiesQuantity = new(SecuritiesOption23)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption51) AddSecuritiesMovementDetails() *SecuritiesOption33 {
	newValue := new(SecuritiesOption33)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption51) AddCashMovementDetails() *CashOption24 {
	newValue := new(CashOption24)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption51) AddAdditionalInformation() *CorporateActionNarrative20 {
	c.AdditionalInformation = new(CorporateActionNarrative20)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption52 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption10Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType19Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat3Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat9Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat9Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus1Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType5Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether redemption charges apply.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate15 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate37 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice28 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity related to a corporate action option.
	SecuritiesQuantity *SecuritiesOption23 `xml:"SctiesQty,omitempty"`

	// Provides information about securities movement related to a corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption38 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption25 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information about the corporate action movement.
	AdditionalInformation *CorporateActionNarrative20 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption52) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption52) AddOptionType() *CorporateActionOption10Choice {
	c.OptionType = new(CorporateActionOption10Choice)
	return c.OptionType
}

func (c *CorporateActionOption52) AddFractionDisposition() *FractionDispositionType19Choice {
	c.FractionDisposition = new(FractionDispositionType19Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption52) AddOfferType() *OfferTypeFormat3Choice {
	newValue := new(OfferTypeFormat3Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption52) AddOptionFeatures() *OptionFeaturesFormat9Choice {
	newValue := new(OptionFeaturesFormat9Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption52) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat9Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat9Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption52) AddOptionAvailabilityStatus() *OptionAvailabilityStatus1Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus1Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption52) AddCertificationBreakdownType() *BeneficiaryCertificationType5Choice {
	newValue := new(BeneficiaryCertificationType5Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption52) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption52) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption52) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption52) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption52) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption52) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption52) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption52) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption52) AddSecurityIdentification() *SecurityIdentification14 {
	c.SecurityIdentification = new(SecurityIdentification14)
	return c.SecurityIdentification
}

func (c *CorporateActionOption52) AddDateDetails() *CorporateActionDate15 {
	c.DateDetails = new(CorporateActionDate15)
	return c.DateDetails
}

func (c *CorporateActionOption52) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption52) AddRateAndAmountDetails() *CorporateActionRate37 {
	c.RateAndAmountDetails = new(CorporateActionRate37)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption52) AddPriceDetails() *CorporateActionPrice28 {
	c.PriceDetails = new(CorporateActionPrice28)
	return c.PriceDetails
}

func (c *CorporateActionOption52) AddSecuritiesQuantity() *SecuritiesOption23 {
	c.SecuritiesQuantity = new(SecuritiesOption23)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption52) AddSecuritiesMovementDetails() *SecuritiesOption38 {
	newValue := new(SecuritiesOption38)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption52) AddCashMovementDetails() *CashOption25 {
	newValue := new(CashOption25)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption52) AddAdditionalInformation() *CorporateActionNarrative20 {
	c.AdditionalInformation = new(CorporateActionNarrative20)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption53 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption11Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat1Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType23Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate18 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod9 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate38 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice30 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption35 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption26 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption53) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption53) AddOptionType() *CorporateActionOption11Choice {
	c.OptionType = new(CorporateActionOption11Choice)
	return c.OptionType
}

func (c *CorporateActionOption53) AddOptionFeatures() *OptionFeaturesFormat1Choice {
	newValue := new(OptionFeaturesFormat1Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption53) AddFractionDisposition() *FractionDispositionType23Choice {
	c.FractionDisposition = new(FractionDispositionType23Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption53) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption53) AddDateDetails() *CorporateActionDate18 {
	c.DateDetails = new(CorporateActionDate18)
	return c.DateDetails
}

func (c *CorporateActionOption53) AddPeriodDetails() *CorporateActionPeriod9 {
	c.PeriodDetails = new(CorporateActionPeriod9)
	return c.PeriodDetails
}

func (c *CorporateActionOption53) AddRateAndAmountDetails() *CorporateActionRate38 {
	c.RateAndAmountDetails = new(CorporateActionRate38)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption53) AddPriceDetails() *CorporateActionPrice30 {
	c.PriceDetails = new(CorporateActionPrice30)
	return c.PriceDetails
}

func (c *CorporateActionOption53) AddPlaceOfTrade() *MarketIdentification4 {
	c.PlaceOfTrade = new(MarketIdentification4)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption53) AddSecuritiesMovementDetails() *SecuritiesOption35 {
	newValue := new(SecuritiesOption35)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption53) AddCashMovementDetails() *CashOption26 {
	newValue := new(CashOption26)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

// Provides information about the corporate action option.
type CorporateActionOption57 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption12Choice `xml:"OptnTp"`

	// Specifies how fractional amount/quantities are treated.
	FractionDisposition *FractionDispositionType17Choice `xml:"FrctnDspstn,omitempty"`

	// Type of changes affecting the security form.
	ChangeType []*CorporateActionChangeTypeFormat2Choice `xml:"ChngTp,omitempty"`

	// Specifies that the corporate action instruction is to be processed using the Available-for-Collateral pool.
	EligibleForCollateralIndicator *YesNoIndicator `xml:"ElgblForCollInd,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds.
	CurrencyToBuy *ActiveCurrencyCode `xml:"CcyToBuy,omitempty"`

	// Account servicer is instructed to sell the indicated currency in order to obtain the necessary currency to fund the transaction within this instruction message.
	CurrencyToSell *ActiveCurrencyCode `xml:"CcyToSell,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption2 `xml:"SctiesQty"`

	// Date/time at which the instructing party requests the instruction to be executed.
	ExecutionRequestedDateTime *DateAndDateTimeChoice `xml:"ExctnReqdDtTm,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate8 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice29 `xml:"PricDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative8 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption57) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption57) AddOptionType() *CorporateActionOption12Choice {
	c.OptionType = new(CorporateActionOption12Choice)
	return c.OptionType
}

func (c *CorporateActionOption57) AddFractionDisposition() *FractionDispositionType17Choice {
	c.FractionDisposition = new(FractionDispositionType17Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption57) AddChangeType() *CorporateActionChangeTypeFormat2Choice {
	newValue := new(CorporateActionChangeTypeFormat2Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateActionOption57) SetEligibleForCollateralIndicator(value string) {
	c.EligibleForCollateralIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption57) SetCurrencyToBuy(value string) {
	c.CurrencyToBuy = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption57) SetCurrencyToSell(value string) {
	c.CurrencyToSell = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption57) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption57) AddSecurityIdentification() *SecurityIdentification14 {
	c.SecurityIdentification = new(SecurityIdentification14)
	return c.SecurityIdentification
}

func (c *CorporateActionOption57) AddSecuritiesQuantity() *SecuritiesOption2 {
	c.SecuritiesQuantity = new(SecuritiesOption2)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption57) AddExecutionRequestedDateTime() *DateAndDateTimeChoice {
	c.ExecutionRequestedDateTime = new(DateAndDateTimeChoice)
	return c.ExecutionRequestedDateTime
}

func (c *CorporateActionOption57) AddRateAndAmountDetails() *CorporateActionRate8 {
	c.RateAndAmountDetails = new(CorporateActionRate8)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption57) AddPriceDetails() *CorporateActionPrice29 {
	c.PriceDetails = new(CorporateActionPrice29)
	return c.PriceDetails
}

func (c *CorporateActionOption57) AddAdditionalInformation() *CorporateActionNarrative8 {
	c.AdditionalInformation = new(CorporateActionNarrative8)
	return c.AdditionalInformation
}

// Provides information about the corporate action option.
type CorporateActionOption6 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption4Choice `xml:"OptnTp"`

	// Specifies whether the quantity of financial instrument is a quantity of securities instructed or a quantity to receive.
	InstructedOrQuantityToReceive *InstructedOrQuantityToReceive1Choice `xml:"InstdOrQtyToRcv"`
}

func (c *CorporateActionOption6) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption6) AddOptionType() *CorporateActionOption4Choice {
	c.OptionType = new(CorporateActionOption4Choice)
	return c.OptionType
}

func (c *CorporateActionOption6) AddInstructedOrQuantityToReceive() *InstructedOrQuantityToReceive1Choice {
	c.InstructedOrQuantityToReceive = new(InstructedOrQuantityToReceive1Choice)
	return c.InstructedOrQuantityToReceive
}

// Provides information about the corporate action option.
type CorporateActionOption9 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption5Choice `xml:"OptnTp"`

	// Party that owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Account on which a securities entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId,omitempty"`

	// Total balance of securities eligible for this corporate action event. The entitlement calculation is based on this balance.
	TotalEligibleBalance *SignedQuantityFormat1 `xml:"TtlElgblBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *SignedQuantityFormat1 `xml:"InstdBal,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *SignedQuantityFormat1 `xml:"UinstdBal,omitempty"`

	// Specifies whether the quantity of financial instrument is a status quantity or a quantity to receive.
	StatusQuantityOrQuantityToReceive *StatusOrQuantityToReceive1Choice `xml:"StsQtyOrQtyToRcv,omitempty"`
}

func (c *CorporateActionOption9) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption9) AddOptionType() *CorporateActionOption5Choice {
	c.OptionType = new(CorporateActionOption5Choice)
	return c.OptionType
}

func (c *CorporateActionOption9) AddAccountOwner() *PartyIdentification13Choice {
	c.AccountOwner = new(PartyIdentification13Choice)
	return c.AccountOwner
}

func (c *CorporateActionOption9) SetSafekeepingAccount(value string) {
	c.SafekeepingAccount = (*Max35Text)(&value)
}

func (c *CorporateActionOption9) AddCashAccount() *CashAccountIdentification5Choice {
	c.CashAccount = new(CashAccountIdentification5Choice)
	return c.CashAccount
}

func (c *CorporateActionOption9) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	c.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return c.SafekeepingPlace
}

func (c *CorporateActionOption9) AddSecurityIdentification() *SecurityIdentification11 {
	c.SecurityIdentification = new(SecurityIdentification11)
	return c.SecurityIdentification
}

func (c *CorporateActionOption9) AddTotalEligibleBalance() *SignedQuantityFormat1 {
	c.TotalEligibleBalance = new(SignedQuantityFormat1)
	return c.TotalEligibleBalance
}

func (c *CorporateActionOption9) AddInstructedBalance() *SignedQuantityFormat1 {
	c.InstructedBalance = new(SignedQuantityFormat1)
	return c.InstructedBalance
}

func (c *CorporateActionOption9) AddUninstructedBalance() *SignedQuantityFormat1 {
	c.UninstructedBalance = new(SignedQuantityFormat1)
	return c.UninstructedBalance
}

func (c *CorporateActionOption9) AddStatusQuantityOrQuantityToReceive() *StatusOrQuantityToReceive1Choice {
	c.StatusQuantityOrQuantityToReceive = new(StatusOrQuantityToReceive1Choice)
	return c.StatusQuantityOrQuantityToReceive
}
