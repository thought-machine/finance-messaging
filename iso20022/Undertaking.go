package iso20022

// Independent undertaking, such as a demand guarantee or standby letter of credit, that provides financial assurance, to be honoured on the presentation of documents that comply with its terms and conditions.
type Undertaking1 struct {

	// Unique and unambiguous identifier assigned by the applicant to the undertaking.
	ApplicantReferenceNumber *Max35Text `xml:"ApplcntRefNb"`

	// Brief description of the purpose of the undertaking. Provided as information for the issuer reference.
	Purpose *Max350Text `xml:"Purp"`

	// Undertaking name.
	Name *UndertakingName1Code `xml:"Nm"`

	// Type of undertaking, for example, performance, payment.
	Type *UndertakingType1Choice `xml:"Tp"`

	// Party obligated to reimburse the issuer.
	Obligor *PartyIdentification43 `xml:"Oblgr"`

	// Party to be named in the undertaking as the “applicant” when different from the obligor.
	Applicant []*PartyIdentification43 `xml:"Applcnt,omitempty"`

	// Party that issues the undertaking (or counter-undertaking).
	Issuer *PartyIdentification43 `xml:"Issr"`

	// Ultimate party in whose favour the undertaking is to be issued.
	Beneficiary []*PartyIdentification43 `xml:"Bnfcry"`

	// Party asked to advise the undertaking to the beneficiary or to another advising party.
	AdvisingParty *PartyIdentification43 `xml:"AdvsgPty,omitempty"`

	// Additional party asked to advise the undertaking.
	SecondAdvisingParty *PartyIdentification43 `xml:"ScndAdvsgPty,omitempty"`

	// Party that adds its confirmation to the undertaking. For further clarification, reference the applicable rules to which the undertaking is subject.
	Confirmer *PartyIdentification43 `xml:"Cnfrmr,omitempty"`

	// Indicates whether the advising bank (confirmer) is requested to add its confirmation to the undertaking. The absence of this element indicates that the advising bank (confirmer) is not requested to add its confirmation to the undertaking.
	ConfirmationIndicator *YesNoIndicator `xml:"ConfInd,omitempty"`

	// Indicates whether the undertaking is a local or ancillary undertaking to be issued under a counter-undertaking.
	CounterUndertakingIndicator *YesNoIndicator `xml:"CntrUdrtkgInd"`

	// Details related to the counter undertaking.
	CounterUndertaking *Undertaking2 `xml:"CntrUdrtkg,omitempty"`

	// Details related to the amount of the undertaking.
	UndertakingAmount *UndertakingAmount1 `xml:"UdrtkgAmt"`

	// Details related to the expiry terms of the undertaking.
	ExpiryDetails *ExpiryDetails2 `xml:"XpryDtls"`

	// Party, in addition to the other parties specified in the undertaking, that is also related to the undertaking.
	AdditionalParty []*PartyAndType1 `xml:"AddtlPty,omitempty"`

	// Rules and laws governing the undertaking.
	GovernanceRulesAndLaw *GovernanceRules1 `xml:"GovncRulesAndLaw"`

	// Details of the underlying transaction for which the undertaking is issued.
	UnderlyingTransaction []*UnderlyingTradeTransaction1 `xml:"UndrlygTx,omitempty"`

	// Presentation details related to the undertaking.
	PresentationDetails *Presentation4 `xml:"PresntnDtls,omitempty"`

	// Wording details and text for the undertaking.
	UndertakingWording *UndertakingWording1 `xml:"UdrtkgWrdg"`

	// Indicates whether multiple demands are permitted.
	MultipleDemandIndicator *YesNoIndicator `xml:"MltplDmndInd,omitempty"`

	// Indicates whether partial demands/drawings are permitted.
	PartialDemandIndicator *YesNoIndicator `xml:"PrtlDmndInd,omitempty"`

	// Indicates whether the undertaking is transferable.
	TransferIndicator *YesNoIndicator `xml:"TrfInd,omitempty"`

	// Indicates whether the applicant/obligor or beneficiary is responsible for payment of the transfer charges.
	TransferChargesPayableBy *ExternalTypeOfParty1Code `xml:"TrfChrgsPyblBy,omitempty"`

	// Indicates whether the applicant/obligor or beneficiary is responsible for payment of the confirmation charges.
	ConfirmationChargesPayableBy *ExternalTypeOfParty1Code `xml:"ConfChrgsPyblBy,omitempty"`

	// Details related to a variation in amount that is automatically applied.
	AutomaticAmountVariation []*AutomaticVariation1 `xml:"AutomtcAmtVartn,omitempty"`

	// Details of the communication channel.
	DeliveryChannel *CommunicationChannel1 `xml:"DlvryChanl"`

	// Account nominated by the obligor to record the liability amount of the undertaking.
	ObligorLiabilityAccount *CashAccount28 `xml:"OblgrLbltyAcct,omitempty"`

	// Account nominated by the obligor for the settlement of charges related to the undertaking.
	ObligorChargeAccount *CashAccount28 `xml:"OblgrChrgAcct,omitempty"`

	// Account nominated by the obligor for the settlement of the amount claimed under the undertaking.
	ObligorSettlementAccount *CashAccount28 `xml:"OblgrSttlmAcct,omitempty"`

	// Document or template enclosed in the undertaking directly related to the undertaking to be issued.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Additional information related to the application for an undertaking.
	AdditionalApplicationInformation []*Max2000Text `xml:"AddtlApplInf,omitempty"`
}

func (u *Undertaking1) SetApplicantReferenceNumber(value string) {
	u.ApplicantReferenceNumber = (*Max35Text)(&value)
}

func (u *Undertaking1) SetPurpose(value string) {
	u.Purpose = (*Max350Text)(&value)
}

func (u *Undertaking1) SetName(value string) {
	u.Name = (*UndertakingName1Code)(&value)
}

func (u *Undertaking1) AddType() *UndertakingType1Choice {
	u.Type = new(UndertakingType1Choice)
	return u.Type
}

func (u *Undertaking1) AddObligor() *PartyIdentification43 {
	u.Obligor = new(PartyIdentification43)
	return u.Obligor
}

func (u *Undertaking1) AddApplicant() *PartyIdentification43 {
	newValue := new(PartyIdentification43)
	u.Applicant = append(u.Applicant, newValue)
	return newValue
}

func (u *Undertaking1) AddIssuer() *PartyIdentification43 {
	u.Issuer = new(PartyIdentification43)
	return u.Issuer
}

func (u *Undertaking1) AddBeneficiary() *PartyIdentification43 {
	newValue := new(PartyIdentification43)
	u.Beneficiary = append(u.Beneficiary, newValue)
	return newValue
}

func (u *Undertaking1) AddAdvisingParty() *PartyIdentification43 {
	u.AdvisingParty = new(PartyIdentification43)
	return u.AdvisingParty
}

func (u *Undertaking1) AddSecondAdvisingParty() *PartyIdentification43 {
	u.SecondAdvisingParty = new(PartyIdentification43)
	return u.SecondAdvisingParty
}

func (u *Undertaking1) AddConfirmer() *PartyIdentification43 {
	u.Confirmer = new(PartyIdentification43)
	return u.Confirmer
}

func (u *Undertaking1) SetConfirmationIndicator(value string) {
	u.ConfirmationIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking1) SetCounterUndertakingIndicator(value string) {
	u.CounterUndertakingIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking1) AddCounterUndertaking() *Undertaking2 {
	u.CounterUndertaking = new(Undertaking2)
	return u.CounterUndertaking
}

func (u *Undertaking1) AddUndertakingAmount() *UndertakingAmount1 {
	u.UndertakingAmount = new(UndertakingAmount1)
	return u.UndertakingAmount
}

func (u *Undertaking1) AddExpiryDetails() *ExpiryDetails2 {
	u.ExpiryDetails = new(ExpiryDetails2)
	return u.ExpiryDetails
}

func (u *Undertaking1) AddAdditionalParty() *PartyAndType1 {
	newValue := new(PartyAndType1)
	u.AdditionalParty = append(u.AdditionalParty, newValue)
	return newValue
}

func (u *Undertaking1) AddGovernanceRulesAndLaw() *GovernanceRules1 {
	u.GovernanceRulesAndLaw = new(GovernanceRules1)
	return u.GovernanceRulesAndLaw
}

func (u *Undertaking1) AddUnderlyingTransaction() *UnderlyingTradeTransaction1 {
	newValue := new(UnderlyingTradeTransaction1)
	u.UnderlyingTransaction = append(u.UnderlyingTransaction, newValue)
	return newValue
}

func (u *Undertaking1) AddPresentationDetails() *Presentation4 {
	u.PresentationDetails = new(Presentation4)
	return u.PresentationDetails
}

func (u *Undertaking1) AddUndertakingWording() *UndertakingWording1 {
	u.UndertakingWording = new(UndertakingWording1)
	return u.UndertakingWording
}

func (u *Undertaking1) SetMultipleDemandIndicator(value string) {
	u.MultipleDemandIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking1) SetPartialDemandIndicator(value string) {
	u.PartialDemandIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking1) SetTransferIndicator(value string) {
	u.TransferIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking1) SetTransferChargesPayableBy(value string) {
	u.TransferChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking1) SetConfirmationChargesPayableBy(value string) {
	u.ConfirmationChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking1) AddAutomaticAmountVariation() *AutomaticVariation1 {
	newValue := new(AutomaticVariation1)
	u.AutomaticAmountVariation = append(u.AutomaticAmountVariation, newValue)
	return newValue
}

func (u *Undertaking1) AddDeliveryChannel() *CommunicationChannel1 {
	u.DeliveryChannel = new(CommunicationChannel1)
	return u.DeliveryChannel
}

func (u *Undertaking1) AddObligorLiabilityAccount() *CashAccount28 {
	u.ObligorLiabilityAccount = new(CashAccount28)
	return u.ObligorLiabilityAccount
}

func (u *Undertaking1) AddObligorChargeAccount() *CashAccount28 {
	u.ObligorChargeAccount = new(CashAccount28)
	return u.ObligorChargeAccount
}

func (u *Undertaking1) AddObligorSettlementAccount() *CashAccount28 {
	u.ObligorSettlementAccount = new(CashAccount28)
	return u.ObligorSettlementAccount
}

func (u *Undertaking1) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	u.EnclosedFile = append(u.EnclosedFile, newValue)
	return newValue
}

func (u *Undertaking1) AddAdditionalApplicationInformation(value string) {
	u.AdditionalApplicationInformation = append(u.AdditionalApplicationInformation, (*Max2000Text)(&value))
}

// Details related to the undertaking.
type Undertaking10 struct {

	// Details related to the requested new amount for the counter-undertaking.
	NewUndertakingAmount *UndertakingAmount2 `xml:"NewUdrtkgAmt,omitempty"`

	// Details related to the requested new expiry terms for the counter-undertaking.
	NewExpiryDetails *ExpiryDetails1 `xml:"NewXpryDtls,omitempty"`
}

func (u *Undertaking10) AddNewUndertakingAmount() *UndertakingAmount2 {
	u.NewUndertakingAmount = new(UndertakingAmount2)
	return u.NewUndertakingAmount
}

func (u *Undertaking10) AddNewExpiryDetails() *ExpiryDetails1 {
	u.NewExpiryDetails = new(ExpiryDetails1)
	return u.NewExpiryDetails
}

// Details related to the local undertaking.
type Undertaking11 struct {

	// Details related to the requested new amount for the local undertaking.
	NewUndertakingAmount *UndertakingAmount2 `xml:"NewUdrtkgAmt,omitempty"`

	// Details related to the requested new expiry terms for the local undertaking.
	NewExpiryDetails *ExpiryDetails1 `xml:"NewXpryDtls,omitempty"`

	// Details related to the requested new beneficiary for the local undertaking.
	NewBeneficiary *PartyIdentification43 `xml:"NewBnfcry,omitempty"`

	// Details related to the requested new terms and conditions for the local undertaking.
	NewUndertakingTermsAndConditions *Narrative1 `xml:"NewUdrtkgTermsAndConds,omitempty"`

	// Details related to the delivery channel for the amended local undertaking.
	DeliveryChannel *CommunicationChannel1 `xml:"DlvryChanl,omitempty"`
}

func (u *Undertaking11) AddNewUndertakingAmount() *UndertakingAmount2 {
	u.NewUndertakingAmount = new(UndertakingAmount2)
	return u.NewUndertakingAmount
}

func (u *Undertaking11) AddNewExpiryDetails() *ExpiryDetails1 {
	u.NewExpiryDetails = new(ExpiryDetails1)
	return u.NewExpiryDetails
}

func (u *Undertaking11) AddNewBeneficiary() *PartyIdentification43 {
	u.NewBeneficiary = new(PartyIdentification43)
	return u.NewBeneficiary
}

func (u *Undertaking11) AddNewUndertakingTermsAndConditions() *Narrative1 {
	u.NewUndertakingTermsAndConditions = new(Narrative1)
	return u.NewUndertakingTermsAndConditions
}

func (u *Undertaking11) AddDeliveryChannel() *CommunicationChannel1 {
	u.DeliveryChannel = new(CommunicationChannel1)
	return u.DeliveryChannel
}

// Independent undertaking, such as a demand guarantee or standby letter of credit, that provides financial assurance, to be honoured on the presentation of documents that comply with its terms and conditions.
type Undertaking2 struct {

	// Undertaking name.
	Name *UndertakingName1Code `xml:"Nm,omitempty"`

	// Party in whose favour the counter-undertaking is issued.
	Beneficiary *PartyIdentification43 `xml:"Bnfcry,omitempty"`

	// Details related to the expiry terms of the counter-undertaking.
	ExpiryDetails *ExpiryDetails2 `xml:"XpryDtls,omitempty"`

	// Details related to the amount of the counter-undertaking.
	CounterUndertakingAmount *UndertakingAmount1 `xml:"CntrUdrtkgAmt,omitempty"`

	// Indicates whether the applicant/obligor or beneficiary is responsible for payment of the confirmation charges.
	ConfirmationChargesPayableBy *ExternalTypeOfParty1Code `xml:"ConfChrgsPyblBy,omitempty"`

	// Rules and laws governing the counter-undertaking.
	GovernanceRulesAndLaw *GovernanceRules1 `xml:"GovncRulesAndLaw,omitempty"`

	// Indication as to whether a claim is to utilise a standard claim form of the issuing institution.
	StandardClaimDocumentIndicator *YesNoIndicator `xml:"StdClmDocInd,omitempty"`

	// Additional information related to the counter-undertaking.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *Undertaking2) SetName(value string) {
	u.Name = (*UndertakingName1Code)(&value)
}

func (u *Undertaking2) AddBeneficiary() *PartyIdentification43 {
	u.Beneficiary = new(PartyIdentification43)
	return u.Beneficiary
}

func (u *Undertaking2) AddExpiryDetails() *ExpiryDetails2 {
	u.ExpiryDetails = new(ExpiryDetails2)
	return u.ExpiryDetails
}

func (u *Undertaking2) AddCounterUndertakingAmount() *UndertakingAmount1 {
	u.CounterUndertakingAmount = new(UndertakingAmount1)
	return u.CounterUndertakingAmount
}

func (u *Undertaking2) SetConfirmationChargesPayableBy(value string) {
	u.ConfirmationChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking2) AddGovernanceRulesAndLaw() *GovernanceRules1 {
	u.GovernanceRulesAndLaw = new(GovernanceRules1)
	return u.GovernanceRulesAndLaw
}

func (u *Undertaking2) SetStandardClaimDocumentIndicator(value string) {
	u.StandardClaimDocumentIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking2) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}

// Independent undertaking, such as a demand guarantee or standby letter of credit, that provides financial assurance, to be honoured on the presentation of documents that comply with its terms and conditions.
type Undertaking3 struct {

	// Unique and unambiguous identifier assigned by the issuer to the undertaking, for example the guarantee or standby number.
	Identification *Max35Text `xml:"Id"`

	// Name of undertaking such as, demand guarantee, standby letter of credit.
	Name *UndertakingIssuanceName1Code `xml:"Nm"`

	// Type of undertaking, for example, performance, payment.
	Type *UndertakingType1Choice `xml:"Tp,omitempty"`

	// Type of the undertaking issuance.
	IssuanceType *IssuanceType1Code `xml:"IssncTp"`

	// Party named in the undertaking as the “applicant”.
	Applicant []*PartyIdentification43 `xml:"Applcnt,omitempty"`

	// Party that issues the undertaking (or counter-undertaking).
	Issuer *PartyIdentification43 `xml:"Issr"`

	// Party in whose favour the undertaking (or counter-undertaking) is issued.
	Beneficiary []*PartyIdentification43 `xml:"Bnfcry"`

	// Date on which the undertaking is issued.
	DateOfIssuance *ISODate `xml:"DtOfIssnc"`

	// Location which is to be regarded as the place from which the undertaking is issued.
	PlaceOfIssue *PostalAddress12 `xml:"PlcOfIsse,omitempty"`

	// Party asked to advise the undertaking to the beneficiary or to another advising party at the request of the issuer.
	AdvisingParty *PartyIdentification43 `xml:"AdvsgPty,omitempty"`

	// Additional party asked to advise the undertaking.
	SecondAdvisingParty *PartyIdentification43 `xml:"ScndAdvsgPty,omitempty"`

	// Details related to the amount of the undertaking.
	UndertakingAmount *UndertakingAmount1 `xml:"UdrtkgAmt"`

	// Details related to the expiry of the undertaking.
	ExpiryDetails *ExpiryDetails1 `xml:"XpryDtls"`

	// Indicates whether or not the advising bank (confirmer) is requested to add its confirmation to the undertaking.
	ConfirmationIndicator *YesNoIndicator `xml:"ConfInd,omitempty"`

	// Indicates the type of party requested to add its confirmation to the undertaking.
	ConfirmationPartyType *ExternalTypeOfParty1Code `xml:"ConfPtyTp,omitempty"`

	// Party, in addition to the other parties specified in the undertaking, that is also related to the undertaking .
	AdditionalParty []*PartyAndType1 `xml:"AddtlPty,omitempty"`

	// Rules and laws governing the undertaking.
	GovernanceRulesAndLaw *GovernanceRules1 `xml:"GovncRulesAndLaw"`

	// Details of the underlying transaction for which the undertaking is issued.
	UnderlyingTransaction []*UnderlyingTradeTransaction1 `xml:"UndrlygTx,omitempty"`

	// Presentation details related to the undertaking.
	PresentationDetails *Presentation1 `xml:"PresntnDtls,omitempty"`

	// Terms and conditions of the undertaking.
	UndertakingTermsAndConditions []*Narrative1 `xml:"UdrtkgTermsAndConds"`

	// Indicates that multiple demands are not permitted.
	MultipleDemandIndicator *YesNoIndicator `xml:"MltplDmndInd,omitempty"`

	// Indicates that partial demands/drawings are not permitted.
	PartialDemandIndicator *YesNoIndicator `xml:"PrtlDmndInd,omitempty"`

	// Indicates whether the applicant/obligor or beneficiary is responsible for payment of the confirmation charges.
	ConfirmationChargesPayableBy *ExternalTypeOfParty1Code `xml:"ConfChrgsPyblBy,omitempty"`

	// Indicates whether the applicant/obligor or beneficiary is responsible for payment of the transfer charges.
	TransferChargesPayableBy *ExternalTypeOfParty1Code `xml:"TrfChrgsPyblBy,omitempty"`

	// Details related to a variation in amount that is automatically applied.
	AutomaticAmountVariation []*AutomaticVariation1 `xml:"AutomtcAmtVartn,omitempty"`

	// Details of the communication channel.
	DeliveryChannel *CommunicationChannel1 `xml:"DlvryChanl,omitempty"`

	// Indicates whether the undertaking is transferable.
	TransferIndicator *YesNoIndicator `xml:"TrfInd,omitempty"`

	// Document or template enclosed in the undertaking directly related to the issued undertaking.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Additional information related to the undertaking.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`

	// Details of the local or ancillary undertaking requested to be issued by a local or other issuing institution.
	RequestedLocalUndertaking *Undertaking4 `xml:"ReqdLclUdrtkg,omitempty"`
}

func (u *Undertaking3) SetIdentification(value string) {
	u.Identification = (*Max35Text)(&value)
}

func (u *Undertaking3) SetName(value string) {
	u.Name = (*UndertakingIssuanceName1Code)(&value)
}

func (u *Undertaking3) AddType() *UndertakingType1Choice {
	u.Type = new(UndertakingType1Choice)
	return u.Type
}

func (u *Undertaking3) SetIssuanceType(value string) {
	u.IssuanceType = (*IssuanceType1Code)(&value)
}

func (u *Undertaking3) AddApplicant() *PartyIdentification43 {
	newValue := new(PartyIdentification43)
	u.Applicant = append(u.Applicant, newValue)
	return newValue
}

func (u *Undertaking3) AddIssuer() *PartyIdentification43 {
	u.Issuer = new(PartyIdentification43)
	return u.Issuer
}

func (u *Undertaking3) AddBeneficiary() *PartyIdentification43 {
	newValue := new(PartyIdentification43)
	u.Beneficiary = append(u.Beneficiary, newValue)
	return newValue
}

func (u *Undertaking3) SetDateOfIssuance(value string) {
	u.DateOfIssuance = (*ISODate)(&value)
}

func (u *Undertaking3) AddPlaceOfIssue() *PostalAddress12 {
	u.PlaceOfIssue = new(PostalAddress12)
	return u.PlaceOfIssue
}

func (u *Undertaking3) AddAdvisingParty() *PartyIdentification43 {
	u.AdvisingParty = new(PartyIdentification43)
	return u.AdvisingParty
}

func (u *Undertaking3) AddSecondAdvisingParty() *PartyIdentification43 {
	u.SecondAdvisingParty = new(PartyIdentification43)
	return u.SecondAdvisingParty
}

func (u *Undertaking3) AddUndertakingAmount() *UndertakingAmount1 {
	u.UndertakingAmount = new(UndertakingAmount1)
	return u.UndertakingAmount
}

func (u *Undertaking3) AddExpiryDetails() *ExpiryDetails1 {
	u.ExpiryDetails = new(ExpiryDetails1)
	return u.ExpiryDetails
}

func (u *Undertaking3) SetConfirmationIndicator(value string) {
	u.ConfirmationIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking3) SetConfirmationPartyType(value string) {
	u.ConfirmationPartyType = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking3) AddAdditionalParty() *PartyAndType1 {
	newValue := new(PartyAndType1)
	u.AdditionalParty = append(u.AdditionalParty, newValue)
	return newValue
}

func (u *Undertaking3) AddGovernanceRulesAndLaw() *GovernanceRules1 {
	u.GovernanceRulesAndLaw = new(GovernanceRules1)
	return u.GovernanceRulesAndLaw
}

func (u *Undertaking3) AddUnderlyingTransaction() *UnderlyingTradeTransaction1 {
	newValue := new(UnderlyingTradeTransaction1)
	u.UnderlyingTransaction = append(u.UnderlyingTransaction, newValue)
	return newValue
}

func (u *Undertaking3) AddPresentationDetails() *Presentation1 {
	u.PresentationDetails = new(Presentation1)
	return u.PresentationDetails
}

func (u *Undertaking3) AddUndertakingTermsAndConditions() *Narrative1 {
	newValue := new(Narrative1)
	u.UndertakingTermsAndConditions = append(u.UndertakingTermsAndConditions, newValue)
	return newValue
}

func (u *Undertaking3) SetMultipleDemandIndicator(value string) {
	u.MultipleDemandIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking3) SetPartialDemandIndicator(value string) {
	u.PartialDemandIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking3) SetConfirmationChargesPayableBy(value string) {
	u.ConfirmationChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking3) SetTransferChargesPayableBy(value string) {
	u.TransferChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking3) AddAutomaticAmountVariation() *AutomaticVariation1 {
	newValue := new(AutomaticVariation1)
	u.AutomaticAmountVariation = append(u.AutomaticAmountVariation, newValue)
	return newValue
}

func (u *Undertaking3) AddDeliveryChannel() *CommunicationChannel1 {
	u.DeliveryChannel = new(CommunicationChannel1)
	return u.DeliveryChannel
}

func (u *Undertaking3) SetTransferIndicator(value string) {
	u.TransferIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking3) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	u.EnclosedFile = append(u.EnclosedFile, newValue)
	return newValue
}

func (u *Undertaking3) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}

func (u *Undertaking3) AddRequestedLocalUndertaking() *Undertaking4 {
	u.RequestedLocalUndertaking = new(Undertaking4)
	return u.RequestedLocalUndertaking
}

// Information about an undertaking.
type Undertaking4 struct {

	// Name of the requested local undertaking such as, demand guarantee, standby letter of credit, surety.
	Name *UndertakingName1Code `xml:"Nm"`

	// Type of the requested local undertaking such as performance, payment.
	Type *ExternalUndertakingType1Code `xml:"Tp"`

	// Party requested to be named in the local undertaking as the party on whose behalf the undertaking is issued.
	Applicant []*PartyIdentification43 `xml:"Applcnt"`

	// Party in whose favour the requested local undertaking is to be issued.
	Beneficiary []*PartyIdentification43 `xml:"Bnfcry"`

	// Date on which the requested local undertaking is to be issued.
	DateOfIssuance *ISODate `xml:"DtOfIssnc,omitempty"`

	// Party asked to advise the requested local undertaking to the beneficiary or to another advising party at the request of the local undertaking issuer.
	AdvisingParty *PartyIdentification43 `xml:"AdvsgPty,omitempty"`

	// Additional party asked to advise the requested local undertaking.
	SecondAdvisingParty *PartyIdentification43 `xml:"ScndAdvsgPty,omitempty"`

	// Details related to the amount of the local undertaking.
	LocalUndertakingAmount *UndertakingAmount1 `xml:"LclUdrtkgAmt"`

	// Details related to the expiry of the requested local undertaking.
	ExpiryDetails *ExpiryDetails1 `xml:"XpryDtls"`

	// Indicates whether or not the advising bank (confirmer) is requested to add its confirmation to the undertaking.
	ConfirmationIndicator *YesNoIndicator `xml:"ConfInd,omitempty"`

	// Party, in addition to the other parties specified in the requested local undertaking, that is also related to the requested local undertaking.
	AdditionalParty []*PartyAndType1 `xml:"AddtlPty,omitempty"`

	// Rules and laws governing the requested local undertaking.
	GovernanceRulesAndLaw *GovernanceRules1 `xml:"GovncRulesAndLaw"`

	// Details of the underlying transaction for which the undertaking is issued.
	UnderlyingTransaction []*UnderlyingTradeTransaction1 `xml:"UndrlygTx,omitempty"`

	// Presentation details related to the undertaking.
	PresentationDetails *Presentation1 `xml:"PresntnDtls,omitempty"`

	// Wording details and text for the requested local undertaking.
	UndertakingWording *UndertakingWording1 `xml:"UdrtkgWrdg"`

	// Indicates that multiple demands are not permitted.
	MultipleDemandIndicator *YesNoIndicator `xml:"MltplDmndInd,omitempty"`

	// Indicates that partial demands/drawings are not permitted.
	PartialDemandIndicator *YesNoIndicator `xml:"PrtlDmndInd,omitempty"`

	// Indicates whether the applicant/obligor or beneficiary is responsible for payment of the confirmation charges.
	ConfirmationChargesPayableBy *ExternalTypeOfParty1Code `xml:"ConfChrgsPyblBy,omitempty"`

	// Indicates whether the applicant/obligor or beneficiary is responsible for payment of the transfer charges.
	TransferChargesPayableBy *ExternalTypeOfParty1Code `xml:"TrfChrgsPyblBy,omitempty"`

	// Details related to a variation in amount that is automatically applied.
	AutomaticAmountVariation []*AutomaticVariation1 `xml:"AutomtcAmtVartn,omitempty"`

	// Details of the communication channel.
	DeliveryChannel *CommunicationChannel1 `xml:"DlvryChanl,omitempty"`

	// Indicates whether the requested local undertaking is transferable.
	TransferIndicator *YesNoIndicator `xml:"TrfInd,omitempty"`

	// Additional information related to the requested local undertaking.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *Undertaking4) SetName(value string) {
	u.Name = (*UndertakingName1Code)(&value)
}

func (u *Undertaking4) SetType(value string) {
	u.Type = (*ExternalUndertakingType1Code)(&value)
}

func (u *Undertaking4) AddApplicant() *PartyIdentification43 {
	newValue := new(PartyIdentification43)
	u.Applicant = append(u.Applicant, newValue)
	return newValue
}

func (u *Undertaking4) AddBeneficiary() *PartyIdentification43 {
	newValue := new(PartyIdentification43)
	u.Beneficiary = append(u.Beneficiary, newValue)
	return newValue
}

func (u *Undertaking4) SetDateOfIssuance(value string) {
	u.DateOfIssuance = (*ISODate)(&value)
}

func (u *Undertaking4) AddAdvisingParty() *PartyIdentification43 {
	u.AdvisingParty = new(PartyIdentification43)
	return u.AdvisingParty
}

func (u *Undertaking4) AddSecondAdvisingParty() *PartyIdentification43 {
	u.SecondAdvisingParty = new(PartyIdentification43)
	return u.SecondAdvisingParty
}

func (u *Undertaking4) AddLocalUndertakingAmount() *UndertakingAmount1 {
	u.LocalUndertakingAmount = new(UndertakingAmount1)
	return u.LocalUndertakingAmount
}

func (u *Undertaking4) AddExpiryDetails() *ExpiryDetails1 {
	u.ExpiryDetails = new(ExpiryDetails1)
	return u.ExpiryDetails
}

func (u *Undertaking4) SetConfirmationIndicator(value string) {
	u.ConfirmationIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking4) AddAdditionalParty() *PartyAndType1 {
	newValue := new(PartyAndType1)
	u.AdditionalParty = append(u.AdditionalParty, newValue)
	return newValue
}

func (u *Undertaking4) AddGovernanceRulesAndLaw() *GovernanceRules1 {
	u.GovernanceRulesAndLaw = new(GovernanceRules1)
	return u.GovernanceRulesAndLaw
}

func (u *Undertaking4) AddUnderlyingTransaction() *UnderlyingTradeTransaction1 {
	newValue := new(UnderlyingTradeTransaction1)
	u.UnderlyingTransaction = append(u.UnderlyingTransaction, newValue)
	return newValue
}

func (u *Undertaking4) AddPresentationDetails() *Presentation1 {
	u.PresentationDetails = new(Presentation1)
	return u.PresentationDetails
}

func (u *Undertaking4) AddUndertakingWording() *UndertakingWording1 {
	u.UndertakingWording = new(UndertakingWording1)
	return u.UndertakingWording
}

func (u *Undertaking4) SetMultipleDemandIndicator(value string) {
	u.MultipleDemandIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking4) SetPartialDemandIndicator(value string) {
	u.PartialDemandIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking4) SetConfirmationChargesPayableBy(value string) {
	u.ConfirmationChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking4) SetTransferChargesPayableBy(value string) {
	u.TransferChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking4) AddAutomaticAmountVariation() *AutomaticVariation1 {
	newValue := new(AutomaticVariation1)
	u.AutomaticAmountVariation = append(u.AutomaticAmountVariation, newValue)
	return newValue
}

func (u *Undertaking4) AddDeliveryChannel() *CommunicationChannel1 {
	u.DeliveryChannel = new(CommunicationChannel1)
	return u.DeliveryChannel
}

func (u *Undertaking4) SetTransferIndicator(value string) {
	u.TransferIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking4) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}

// Information about an undertaking.
type Undertaking6 struct {

	// Unique and unambiguous identifier assigned by the issuer to the undertaking, for example the guarantee or standby number.
	Identification *Max35Text `xml:"Id"`

	// Party that issues the undertaking.
	Issuer *PartyIdentification43 `xml:"Issr"`

	// Unique and unambiguous identifier assigned by the beneficiary to the undertaking.
	BeneficiaryReferenceNumber *Max35Text `xml:"BnfcryRefNb,omitempty"`
}

func (u *Undertaking6) SetIdentification(value string) {
	u.Identification = (*Max35Text)(&value)
}

func (u *Undertaking6) AddIssuer() *PartyIdentification43 {
	u.Issuer = new(PartyIdentification43)
	return u.Issuer
}

func (u *Undertaking6) SetBeneficiaryReferenceNumber(value string) {
	u.BeneficiaryReferenceNumber = (*Max35Text)(&value)
}

// Information about an undertaking.
type Undertaking7 struct {

	// Unique and unambiguous identifier assigned by the issuer to the undertaking, for example the guarantee or standby number.
	Identification *Max35Text `xml:"Id"`

	// Party that issues the undertaking.
	Issuer *PartyIdentification43 `xml:"Issr"`
}

func (u *Undertaking7) SetIdentification(value string) {
	u.Identification = (*Max35Text)(&value)
}

func (u *Undertaking7) AddIssuer() *PartyIdentification43 {
	u.Issuer = new(PartyIdentification43)
	return u.Issuer
}

// Information about an undertaking.
type Undertaking8 struct {

	// Unique and unambiguous identifier assigned by the issuer to the undertaking, for example the guarantee or standby number.
	Identification *Max35Text `xml:"Id"`

	// Party that issues the undertaking.
	Issuer *PartyIdentification43 `xml:"Issr"`

	// Unique and unambiguous identifier assigned by the applicant to the undertaking.
	ApplicantReferenceNumber *Max35Text `xml:"ApplcntRefNb,omitempty"`

	// Unique and unambiguous identifier assigned by the beneficiary to the undertaking.
	BeneficiaryReferenceNumber *Max35Text `xml:"BnfcryRefNb,omitempty"`
}

func (u *Undertaking8) SetIdentification(value string) {
	u.Identification = (*Max35Text)(&value)
}

func (u *Undertaking8) AddIssuer() *PartyIdentification43 {
	u.Issuer = new(PartyIdentification43)
	return u.Issuer
}

func (u *Undertaking8) SetApplicantReferenceNumber(value string) {
	u.ApplicantReferenceNumber = (*Max35Text)(&value)
}

func (u *Undertaking8) SetBeneficiaryReferenceNumber(value string) {
	u.BeneficiaryReferenceNumber = (*Max35Text)(&value)
}

// Information about an undertaking.
type Undertaking9 struct {

	// Unique and unambiguous identifier assigned by the issuer to the undertaking, for example the guarantee or standby number.
	Identification *Max35Text `xml:"Id"`

	// Party that issues the undertaking.
	Issuer *PartyIdentification43 `xml:"Issr"`

	// Unique and unambiguous identifier assigned by the applicant to the undertaking.
	ApplicantReferenceNumber *Max35Text `xml:"ApplcntRefNb,omitempty"`
}

func (u *Undertaking9) SetIdentification(value string) {
	u.Identification = (*Max35Text)(&value)
}

func (u *Undertaking9) AddIssuer() *PartyIdentification43 {
	u.Issuer = new(PartyIdentification43)
	return u.Issuer
}

func (u *Undertaking9) SetApplicantReferenceNumber(value string) {
	u.ApplicantReferenceNumber = (*Max35Text)(&value)
}
