package iso20022

// Details of the amendment.
type Amendment1 struct {

	// Sequence number assigned by the issuer to each proposed amendment of the undertaking.
	SequenceNumber *Max4AlphaNumericText `xml:"SeqNb"`

	// Date on which the proposed amendment is issued.
	DateOfIssuance *ISODate `xml:"DtOfIssnc"`

	// Identification of the undertaking.
	UndertakingIdentification *Undertaking7 `xml:"UdrtkgId"`

	// Party asked to advise the proposed amendment to the beneficiary or to another advising party at the request of the issuer.
	AdvisingParty *PartyIdentification43 `xml:"AdvsgPty,omitempty"`

	// Additional party asked to advise the proposed amendment.
	SecondAdvisingParty *PartyIdentification43 `xml:"ScndAdvsgPty,omitempty"`

	// Details concerning the requested termination of the undertaking.
	TerminationDetails *UndertakingTermination3 `xml:"TermntnDtls,omitempty"`

	// Requested increase or decrease to the amount of for the undertaking.
	UndertakingAmountAdjustment *UndertakingAmount2 `xml:"UdrtkgAmtAdjstmnt,omitempty"`

	// Requested new expiry terms for the undertaking.
	NewExpiryDetails *ExpiryDetails1 `xml:"NewXpryDtls,omitempty"`

	// Requested new beneficiary of the undertaking.
	NewBeneficiary *PartyIdentification43 `xml:"NewBnfcry,omitempty"`

	// Requested new terms and conditions of the undertaking.
	NewUndertakingTermsAndConditions []*Narrative1 `xml:"NewUdrtkgTermsAndConds,omitempty"`

	// Amendment details related to the local undertaking.
	LocalUndertaking *Undertaking11 `xml:"LclUdrtkg,omitempty"`

	// Indicates whether or not consent is requested from the beneficiary.
	BeneficiaryConsentRequestIndicator *YesNoIndicator `xml:"BnfcryCnsntReqInd,omitempty"`

	// Communication channel for delivery of the proposed amendment.
	DeliveryChannel *CommunicationChannel1 `xml:"DlvryChanl,omitempty"`

	// Document or template enclosed in the proposed amendment.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Additional information related to the proposed amendment.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (a *Amendment1) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max4AlphaNumericText)(&value)
}

func (a *Amendment1) SetDateOfIssuance(value string) {
	a.DateOfIssuance = (*ISODate)(&value)
}

func (a *Amendment1) AddUndertakingIdentification() *Undertaking7 {
	a.UndertakingIdentification = new(Undertaking7)
	return a.UndertakingIdentification
}

func (a *Amendment1) AddAdvisingParty() *PartyIdentification43 {
	a.AdvisingParty = new(PartyIdentification43)
	return a.AdvisingParty
}

func (a *Amendment1) AddSecondAdvisingParty() *PartyIdentification43 {
	a.SecondAdvisingParty = new(PartyIdentification43)
	return a.SecondAdvisingParty
}

func (a *Amendment1) AddTerminationDetails() *UndertakingTermination3 {
	a.TerminationDetails = new(UndertakingTermination3)
	return a.TerminationDetails
}

func (a *Amendment1) AddUndertakingAmountAdjustment() *UndertakingAmount2 {
	a.UndertakingAmountAdjustment = new(UndertakingAmount2)
	return a.UndertakingAmountAdjustment
}

func (a *Amendment1) AddNewExpiryDetails() *ExpiryDetails1 {
	a.NewExpiryDetails = new(ExpiryDetails1)
	return a.NewExpiryDetails
}

func (a *Amendment1) AddNewBeneficiary() *PartyIdentification43 {
	a.NewBeneficiary = new(PartyIdentification43)
	return a.NewBeneficiary
}

func (a *Amendment1) AddNewUndertakingTermsAndConditions() *Narrative1 {
	newValue := new(Narrative1)
	a.NewUndertakingTermsAndConditions = append(a.NewUndertakingTermsAndConditions, newValue)
	return newValue
}

func (a *Amendment1) AddLocalUndertaking() *Undertaking11 {
	a.LocalUndertaking = new(Undertaking11)
	return a.LocalUndertaking
}

func (a *Amendment1) SetBeneficiaryConsentRequestIndicator(value string) {
	a.BeneficiaryConsentRequestIndicator = (*YesNoIndicator)(&value)
}

func (a *Amendment1) AddDeliveryChannel() *CommunicationChannel1 {
	a.DeliveryChannel = new(CommunicationChannel1)
	return a.DeliveryChannel
}

func (a *Amendment1) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	a.EnclosedFile = append(a.EnclosedFile, newValue)
	return newValue
}

func (a *Amendment1) AddAdditionalInformation(value string) {
	a.AdditionalInformation = append(a.AdditionalInformation, (*Max2000Text)(&value))
}

// Details of the amendment.
type Amendment2 struct {

	// Contents of the related Undertaking Amendment message.
	UndertakingAmendmentMessage *UndertakingAmendmentMessage1 `xml:"UdrtkgAmdmntMsg"`

	// Additional information related to the first advising party.
	FirstAdvisingPartyAdditionalInformation *AdvisingPartyAdditionalInformation1 `xml:"FrstAdvsgPtyAddtlInf,omitempty"`

	// Additional information related to the second advising party.
	SecondAdvisingPartyAdditionalInformation *AdvisingPartyAdditionalInformation1 `xml:"ScndAdvsgPtyAddtlInf,omitempty"`

	// Details concerning the confirmation of the proposed amendment.
	ConfirmationDetails *UndertakingConfirmation1 `xml:"ConfDtls,omitempty"`

	// Digital signature of the party providing additional undertaking amendment advice details.
	DigitalSignature []*PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (a *Amendment2) AddUndertakingAmendmentMessage() *UndertakingAmendmentMessage1 {
	a.UndertakingAmendmentMessage = new(UndertakingAmendmentMessage1)
	return a.UndertakingAmendmentMessage
}

func (a *Amendment2) AddFirstAdvisingPartyAdditionalInformation() *AdvisingPartyAdditionalInformation1 {
	a.FirstAdvisingPartyAdditionalInformation = new(AdvisingPartyAdditionalInformation1)
	return a.FirstAdvisingPartyAdditionalInformation
}

func (a *Amendment2) AddSecondAdvisingPartyAdditionalInformation() *AdvisingPartyAdditionalInformation1 {
	a.SecondAdvisingPartyAdditionalInformation = new(AdvisingPartyAdditionalInformation1)
	return a.SecondAdvisingPartyAdditionalInformation
}

func (a *Amendment2) AddConfirmationDetails() *UndertakingConfirmation1 {
	a.ConfirmationDetails = new(UndertakingConfirmation1)
	return a.ConfirmationDetails
}

func (a *Amendment2) AddDigitalSignature() *PartyAndSignature2 {
	newValue := new(PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

// Details of the amendent request.
type Amendment3 struct {

	// Unique and unambiguous identifier assigned by the applicant to the undertaking amendment request.
	ApplicantRequestNumber *Max35Text `xml:"ApplcntReqNb"`

	// Identification of the undertaking.
	UndertakingIdentification *Undertaking9 `xml:"UdrtkgId"`

	// Party requesting the issuance of the amendment.
	Applicant *PartyIdentification43 `xml:"Applcnt"`

	// Details concerning the requested termination of the undertaking.
	TerminationDetails *UndertakingTermination3 `xml:"TermntnDtls,omitempty"`

	// Indication of the amount of increase or decrease to the undertaking amount.
	IncreaseDecreaseAmount *UndertakingAmount2 `xml:"IncrDcrAmt,omitempty"`

	// Requested new expiry terms for the undertaking.
	NewExpiryDetails *ExpiryDetails2 `xml:"NewXpryDtls,omitempty"`

	// Requested new beneficiary of the undertaking.
	NewBeneficiary *Beneficiary1 `xml:"NewBnfcry,omitempty"`

	// Requested new terms and conditions of the undertaking.
	NewUndertakingTermsAndConditions []*Narrative1 `xml:"NewUdrtkgTermsAndConds,omitempty"`

	// Amendment details related to the counter-undertaking.
	CounterUndertaking *Undertaking10 `xml:"CntrUdrtkg,omitempty"`

	// Communication channel for delivery of the amendment.
	DeliveryChannel *CommunicationChannel1 `xml:"DlvryChanl,omitempty"`

	// Document or template enclosed in the request.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Additional information related to the request.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (a *Amendment3) SetApplicantRequestNumber(value string) {
	a.ApplicantRequestNumber = (*Max35Text)(&value)
}

func (a *Amendment3) AddUndertakingIdentification() *Undertaking9 {
	a.UndertakingIdentification = new(Undertaking9)
	return a.UndertakingIdentification
}

func (a *Amendment3) AddApplicant() *PartyIdentification43 {
	a.Applicant = new(PartyIdentification43)
	return a.Applicant
}

func (a *Amendment3) AddTerminationDetails() *UndertakingTermination3 {
	a.TerminationDetails = new(UndertakingTermination3)
	return a.TerminationDetails
}

func (a *Amendment3) AddIncreaseDecreaseAmount() *UndertakingAmount2 {
	a.IncreaseDecreaseAmount = new(UndertakingAmount2)
	return a.IncreaseDecreaseAmount
}

func (a *Amendment3) AddNewExpiryDetails() *ExpiryDetails2 {
	a.NewExpiryDetails = new(ExpiryDetails2)
	return a.NewExpiryDetails
}

func (a *Amendment3) AddNewBeneficiary() *Beneficiary1 {
	a.NewBeneficiary = new(Beneficiary1)
	return a.NewBeneficiary
}

func (a *Amendment3) AddNewUndertakingTermsAndConditions() *Narrative1 {
	newValue := new(Narrative1)
	a.NewUndertakingTermsAndConditions = append(a.NewUndertakingTermsAndConditions, newValue)
	return newValue
}

func (a *Amendment3) AddCounterUndertaking() *Undertaking10 {
	a.CounterUndertaking = new(Undertaking10)
	return a.CounterUndertaking
}

func (a *Amendment3) AddDeliveryChannel() *CommunicationChannel1 {
	a.DeliveryChannel = new(CommunicationChannel1)
	return a.DeliveryChannel
}

func (a *Amendment3) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	a.EnclosedFile = append(a.EnclosedFile, newValue)
	return newValue
}

func (a *Amendment3) AddAdditionalInformation(value string) {
	a.AdditionalInformation = append(a.AdditionalInformation, (*Max2000Text)(&value))
}

// Details of the amendment.
type Amendment6 struct {

	// Contents of the related proposed Undertaking Amendment message.
	UndertakingAmendmentMessage *UndertakingAmendmentMessage1 `xml:"UdrtkgAmdmntMsg"`

	// Unique and unambiguous identifier assigned by the applicant to the undertaking.
	ApplicantReferenceNumber *Max35Text `xml:"ApplcntRefNb,omitempty"`

	// Additional information related to the notification.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (a *Amendment6) AddUndertakingAmendmentMessage() *UndertakingAmendmentMessage1 {
	a.UndertakingAmendmentMessage = new(UndertakingAmendmentMessage1)
	return a.UndertakingAmendmentMessage
}

func (a *Amendment6) SetApplicantReferenceNumber(value string) {
	a.ApplicantReferenceNumber = (*Max35Text)(&value)
}

func (a *Amendment6) AddAdditionalInformation(value string) {
	a.AdditionalInformation = append(a.AdditionalInformation, (*Max2000Text)(&value))
}

// Details of the amendment.
type Amendment7 struct {

	// Identification of the proposed amendment.
	AmendmentIdentification *Amendment8 `xml:"AmdmntId"`

	// Proposed undertaking amendment status.
	AmendmentStatus *UndertakingStatus2Code `xml:"AmdmntSts"`
}

func (a *Amendment7) AddAmendmentIdentification() *Amendment8 {
	a.AmendmentIdentification = new(Amendment8)
	return a.AmendmentIdentification
}

func (a *Amendment7) SetAmendmentStatus(value string) {
	a.AmendmentStatus = (*UndertakingStatus2Code)(&value)
}

// Amendment identification.
type Amendment8 struct {

	// Unique and unambiguous identifier assigned by the issuer to the undertaking, for example the guarantee or standby number.
	Identification *Max35Text `xml:"Id"`

	// Sequence number assigned by the issuer to each amendment of the undertaking.
	SequenceNumber *Max4AlphaNumericText `xml:"SeqNb"`

	// Unique and unambiguous identifier assigned by the beneficiary to the undertaking.
	BeneficiaryReferenceNumber *Max35Text `xml:"BnfcryRefNb,omitempty"`

	// Party that issues the undertaking.
	Issuer *PartyIdentification43 `xml:"Issr"`
}

func (a *Amendment8) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Amendment8) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max4AlphaNumericText)(&value)
}

func (a *Amendment8) SetBeneficiaryReferenceNumber(value string) {
	a.BeneficiaryReferenceNumber = (*Max35Text)(&value)
}

func (a *Amendment8) AddIssuer() *PartyIdentification43 {
	a.Issuer = new(PartyIdentification43)
	return a.Issuer
}

// Details of the amendment.
type Amendment9 struct {

	// Contents of the related UndertakingAmendmentResponse message.
	UndertakingAmendmentResponseMessage *UndertakingAmendmentResponseMessage1 `xml:"UdrtkgAmdmntRspnMsg"`
}

func (a *Amendment9) AddUndertakingAmendmentResponseMessage() *UndertakingAmendmentResponseMessage1 {
	a.UndertakingAmendmentResponseMessage = new(UndertakingAmendmentResponseMessage1)
	return a.UndertakingAmendmentResponseMessage
}
