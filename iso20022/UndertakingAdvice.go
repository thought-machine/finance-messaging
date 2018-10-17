package iso20022

// Details of the advice for the issuance of an undertaking.
type UndertakingAdvice1 struct {

	// Contents of the related UndertakingIssuance message.
	UndertakingIssuanceMessage *UndertakingIssuanceMessage `xml:"UdrtkgIssncMsg"`

	// Additional information related to the first advising party.
	FirstAdvisingPartyAdditionalInformation *AdvisingPartyAdditionalInformation1 `xml:"FrstAdvsgPtyAddtlInf,omitempty"`

	// Additional information related to the second advising party.
	SecondAdvisingPartyAdditionalInformation *AdvisingPartyAdditionalInformation1 `xml:"ScndAdvsgPtyAddtlInf,omitempty"`

	// Details related to the confirmation of the undertaking.
	ConfirmationDetails *UndertakingConfirmation1 `xml:"ConfDtls,omitempty"`

	// Digital signature of the party providing additional undertaking advice details.
	DigitalSignature []*PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingAdvice1) AddUndertakingIssuanceMessage() *UndertakingIssuanceMessage {
	u.UndertakingIssuanceMessage = new(UndertakingIssuanceMessage)
	return u.UndertakingIssuanceMessage
}

func (u *UndertakingAdvice1) AddFirstAdvisingPartyAdditionalInformation() *AdvisingPartyAdditionalInformation1 {
	u.FirstAdvisingPartyAdditionalInformation = new(AdvisingPartyAdditionalInformation1)
	return u.FirstAdvisingPartyAdditionalInformation
}

func (u *UndertakingAdvice1) AddSecondAdvisingPartyAdditionalInformation() *AdvisingPartyAdditionalInformation1 {
	u.SecondAdvisingPartyAdditionalInformation = new(AdvisingPartyAdditionalInformation1)
	return u.SecondAdvisingPartyAdditionalInformation
}

func (u *UndertakingAdvice1) AddConfirmationDetails() *UndertakingConfirmation1 {
	u.ConfirmationDetails = new(UndertakingConfirmation1)
	return u.ConfirmationDetails
}

func (u *UndertakingAdvice1) AddDigitalSignature() *PartyAndSignature2 {
	newValue := new(PartyAndSignature2)
	u.DigitalSignature = append(u.DigitalSignature, newValue)
	return newValue
}

// Details of the advice for the issuance of an undertaking.
type UndertakingAdvice2 struct {

	// Unique and unambiguous identifier assigned by the applicant to the undertaking.
	ApplicantReferenceNumber *Max35Text `xml:"ApplcntRefNb"`

	// Party obligated to reimburse the issuer.
	Obligor *PartyIdentification43 `xml:"Oblgr,omitempty"`

	// Contents of the related UndertakingIssuance message.
	UndertakingIssuanceMessage *UndertakingIssuanceMessage `xml:"UdrtkgIssncMsg"`

	// Medium used to issue the original undertaking.
	OriginalIssuedMedium *PresentationMedium1Code `xml:"OrgnlIssdMdm,omitempty"`

	// Document or template enclosed in the notification.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Additional information related to the undertaking notification.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *UndertakingAdvice2) SetApplicantReferenceNumber(value string) {
	u.ApplicantReferenceNumber = (*Max35Text)(&value)
}

func (u *UndertakingAdvice2) AddObligor() *PartyIdentification43 {
	u.Obligor = new(PartyIdentification43)
	return u.Obligor
}

func (u *UndertakingAdvice2) AddUndertakingIssuanceMessage() *UndertakingIssuanceMessage {
	u.UndertakingIssuanceMessage = new(UndertakingIssuanceMessage)
	return u.UndertakingIssuanceMessage
}

func (u *UndertakingAdvice2) SetOriginalIssuedMedium(value string) {
	u.OriginalIssuedMedium = (*PresentationMedium1Code)(&value)
}

func (u *UndertakingAdvice2) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	u.EnclosedFile = append(u.EnclosedFile, newValue)
	return newValue
}

func (u *UndertakingAdvice2) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}
