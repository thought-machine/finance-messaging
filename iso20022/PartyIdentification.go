package iso20022

// Entity involved in an activity.
type PartyIdentification1 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Postal address of a party.
	PostalAddress *PostalAddress1 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identification of a party.
	Identification *Party1Choice `xml:"Id,omitempty"`
}

func (p *PartyIdentification1) SetName(value string) {
	p.Name = (*Max70Text)(&value)
}

func (p *PartyIdentification1) AddPostalAddress() *PostalAddress1 {
	p.PostalAddress = new(PostalAddress1)
	return p.PostalAddress
}

func (p *PartyIdentification1) AddIdentification() *Party1Choice {
	p.Identification = new(Party1Choice)
	return p.Identification
}

// Identification of the party.
type PartyIdentification100 struct {

	// Unique identification of the party.
	Identification *PartyIdentification71Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification100) AddIdentification() *PartyIdentification71Choice {
	p.Identification = new(PartyIdentification71Choice)
	return p.Identification
}

func (p *PartyIdentification100) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

// Provides information about the beneficial owner of the securities.
type PartyIdentification101 struct {

	// Party that is the beneficial owner of the specified quantity of securities.
	OwnerIdentification *PartyIdentification104Choice `xml:"OwnrId"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Country in which a person is permanently domiciled (the place of a persons permanent home).
	DomicileCountry *CountryCode `xml:"DmclCtry,omitempty"`

	// Holder of the security certifies, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Quantity of securities belonging to the beneficial owner specified.
	OwnedSecuritiesQuantity *FinancialInstrumentQuantity15Choice `xml:"OwndSctiesQty"`

	// Type of certification which is required.
	CertificationType []*BeneficiaryCertificationType11Choice `xml:"CertfctnTp,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *RestrictedFINXMax350Text `xml:"CertfctnBrkdwn,omitempty"`
}

func (p *PartyIdentification101) AddOwnerIdentification() *PartyIdentification104Choice {
	p.OwnerIdentification = new(PartyIdentification104Choice)
	return p.OwnerIdentification
}

func (p *PartyIdentification101) AddAlternateIdentification() *AlternatePartyIdentification9 {
	newValue := new(AlternatePartyIdentification9)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

func (p *PartyIdentification101) SetDomicileCountry(value string) {
	p.DomicileCountry = (*CountryCode)(&value)
}

func (p *PartyIdentification101) AddNonDomicileCountry(value string) {
	p.NonDomicileCountry = append(p.NonDomicileCountry, (*CountryCode)(&value))
}

func (p *PartyIdentification101) AddOwnedSecuritiesQuantity() *FinancialInstrumentQuantity15Choice {
	p.OwnedSecuritiesQuantity = new(FinancialInstrumentQuantity15Choice)
	return p.OwnedSecuritiesQuantity
}

func (p *PartyIdentification101) AddCertificationType() *BeneficiaryCertificationType11Choice {
	newValue := new(BeneficiaryCertificationType11Choice)
	p.CertificationType = append(p.CertificationType, newValue)
	return newValue
}

func (p *PartyIdentification101) SetCertificationBreakdown(value string) {
	p.CertificationBreakdown = (*RestrictedFINXMax350Text)(&value)
}

// Provides information about identification of the party .
type PartyIdentification102 struct {

	// Identification of a party.
	Identification *PartyIdentification111Choice `xml:"Id"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Provides alternate identification for a party using an id type, a country code and a text field.
	AlternateIdentification []*AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentification102) AddIdentification() *PartyIdentification111Choice {
	p.Identification = new(PartyIdentification111Choice)
	return p.Identification
}

func (p *PartyIdentification102) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentification102) AddAlternateIdentification() *AlternatePartyIdentification9 {
	newValue := new(AlternatePartyIdentification9)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

// Identification of an entity involved in an activity.
type PartyIdentification103 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification58Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation3 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentification103) AddIdentification() *PartyIdentification58Choice {
	p.Identification = new(PartyIdentification58Choice)
	return p.Identification
}

func (p *PartyIdentification103) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentification103) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

func (p *PartyIdentification103) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentification103) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentification103) AddAdditionalInformation() *PartyTextInformation3 {
	p.AdditionalInformation = new(PartyTextInformation3)
	return p.AdditionalInformation
}

// Identification of an entity involved in an activity.
type PartyIdentification108 struct {

	// Identification of the party.
	Identification *PartyIdentification58Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentification108) AddIdentification() *PartyIdentification58Choice {
	p.Identification = new(PartyIdentification58Choice)
	return p.Identification
}

func (p *PartyIdentification108) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentification108) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Identification of the party.
type PartyIdentification109 struct {

	// Unique identification of the party.
	Identification *PartyIdentification114Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification109) AddIdentification() *PartyIdentification114Choice {
	p.Identification = new(PartyIdentification114Choice)
	return p.Identification
}

func (p *PartyIdentification109) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

// Identification of the party.
type PartyIdentification110 struct {

	// Unique identification of the party.
	Identification *PartyIdentification115Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification110) AddIdentification() *PartyIdentification115Choice {
	p.Identification = new(PartyIdentification115Choice)
	return p.Identification
}

func (p *PartyIdentification110) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

// Identification of the party.
type PartyIdentification111 struct {

	// Unique identification of the party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification111) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentification111) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

// Set of elements used to identify a person or an organisation.
type PartyIdentification112 struct {

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	Identification *Party10Choice `xml:"Id,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Contact defined for this party.
	ContactDetails []*Contacts3 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification112) AddIdentification() *Party10Choice {
	p.Identification = new(Party10Choice)
	return p.Identification
}

func (p *PartyIdentification112) SetName(value string) {
	p.Name = (*Max35Text)(&value)
}

func (p *PartyIdentification112) AddPostalAddress() *PostalAddress6 {
	p.PostalAddress = new(PostalAddress6)
	return p.PostalAddress
}

func (p *PartyIdentification112) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification112) AddContactDetails() *Contacts3 {
	newValue := new(Contacts3)
	p.ContactDetails = append(p.ContactDetails, newValue)
	return newValue
}

// Identification of a party.
type PartyIdentification113 struct {

	// Unique identification of the party.
	Party *PartyIdentification90Choice `xml:"Pty"`

	// Legal entity identification as an alternate identification for the party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification113) AddParty() *PartyIdentification90Choice {
	p.Party = new(PartyIdentification90Choice)
	return p.Party
}

func (p *PartyIdentification113) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

// Unique identification, as assigned by an organisation, to unambiguously identify a party.
type PartyIdentification116 struct {

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	PartyIdentification *OrganisationIdentification28 `xml:"PtyId"`

	// Identifier and name of an organisation that is allocated by an institution.
	LegalOrganisation *LegalOrganisation1 `xml:"LglOrg,omitempty"`

	// TaxParty
	TaxParty *TaxParty1 `xml:"TaxPty,omitempty"`
}

func (p *PartyIdentification116) AddPartyIdentification() *OrganisationIdentification28 {
	p.PartyIdentification = new(OrganisationIdentification28)
	return p.PartyIdentification
}

func (p *PartyIdentification116) AddLegalOrganisation() *LegalOrganisation1 {
	p.LegalOrganisation = new(LegalOrganisation1)
	return p.LegalOrganisation
}

func (p *PartyIdentification116) AddTaxParty() *TaxParty1 {
	p.TaxParty = new(TaxParty1)
	return p.TaxParty
}

// Identification of the party.
type PartyIdentification119 struct {

	// Unique identification of the party.
	Identification *PartyIdentification103Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification119) AddIdentification() *PartyIdentification103Choice {
	p.Identification = new(PartyIdentification103Choice)
	return p.Identification
}

func (p *PartyIdentification119) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

// Provides information about identification of the party .
type PartyIdentification120 struct {

	// Identification of a party.
	Identification *PartyIdentification58Choice `xml:"Id"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Provides alternate identification for a party using an id type, a country code and a text field.
	AlternateIdentification []*AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentification120) AddIdentification() *PartyIdentification58Choice {
	p.Identification = new(PartyIdentification58Choice)
	return p.Identification
}

func (p *PartyIdentification120) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentification120) AddAlternateIdentification() *AlternatePartyIdentification9 {
	newValue := new(AlternatePartyIdentification9)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
type PartyIdentification15 struct {

	// Country in which the organisation is registered.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Description of an organisation.
	Identification *Max35Text `xml:"Id"`
}

func (p *PartyIdentification15) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

func (p *PartyIdentification15) SetIdentification(value string) {
	p.Identification = (*Max35Text)(&value)
}

// Identification of an entity involved in an activity.
type PartyIdentification2 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification12Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentification2) AddIdentification() *PartyIdentification12Choice {
	p.Identification = new(PartyIdentification12Choice)
	return p.Identification
}

func (p *PartyIdentification2) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentification2) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentification2) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentification2) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party involved in the settlement chain.
type PartyIdentification21 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification2Choice `xml:"PtyId"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingReference *Max35Text `xml:"PrcgRef,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (p *PartyIdentification21) AddPartyIdentification() *PartyIdentification2Choice {
	p.PartyIdentification = new(PartyIdentification2Choice)
	return p.PartyIdentification
}

func (p *PartyIdentification21) SetProcessingReference(value string) {
	p.ProcessingReference = (*Max35Text)(&value)
}

func (p *PartyIdentification21) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentification21) AddContactPerson() *ContactIdentification2 {
	p.ContactPerson = new(ContactIdentification2)
	return p.ContactPerson
}

// Unique and unambiguous way to identify an organisation.
type PartyIdentification22 struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICOrBEI *AnyBICIdentifier `xml:"BICOrBEI"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	AlternativeIdentifier []*Max35Text `xml:"AltrntvIdr,omitempty"`
}

func (p *PartyIdentification22) SetBICOrBEI(value string) {
	p.BICOrBEI = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification22) AddAlternativeIdentifier(value string) {
	p.AlternativeIdentifier = append(p.AlternativeIdentifier, (*Max35Text)(&value))
}

// Choice of identification of a party. The party can be identified by providing a BIC or a proprietary code.
// Optionally, the client account number can also be provided.
type PartyIdentification23 struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICOrBEI *AnyBICIdentifier `xml:"BICOrBEI"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification1 `xml:"PrtryId"`

	// Name by which a party is known and which is usually used to identify that party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr,omitempty"`
}

func (p *PartyIdentification23) SetBICOrBEI(value string) {
	p.BICOrBEI = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification23) AddProprietaryIdentification() *GenericIdentification1 {
	p.ProprietaryIdentification = new(GenericIdentification1)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification23) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

// Identification of a person, or a non-financial institution.
type PartyIdentification25 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm"`

	// Unique and unambiguous identifier, as assigned to a party using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification4 `xml:"PrtryId,omitempty"`

	// Identification of a non-financial institution.
	BEI *BEIIdentifier `xml:"BEI,omitempty"`
}

func (p *PartyIdentification25) SetName(value string) {
	p.Name = (*Max70Text)(&value)
}

func (p *PartyIdentification25) AddProprietaryIdentification() *GenericIdentification4 {
	p.ProprietaryIdentification = new(GenericIdentification4)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification25) SetBEI(value string) {
	p.BEI = (*BEIIdentifier)(&value)
}

// Entity involved in an activity.
type PartyIdentification26 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm"`

	// Unique and unambiguous identifier assigned to a party using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification4 `xml:"PrtryId,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress5 `xml:"PstlAdr"`
}

func (p *PartyIdentification26) SetName(value string) {
	p.Name = (*Max70Text)(&value)
}

func (p *PartyIdentification26) AddProprietaryIdentification() *GenericIdentification4 {
	p.ProprietaryIdentification = new(GenericIdentification4)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification26) AddPostalAddress() *PostalAddress5 {
	p.PostalAddress = new(PostalAddress5)
	return p.PostalAddress
}

// Entity involved in an activity.
type PartyIdentification27 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm"`

	// Unique and unambiguous identifier assigned to a party using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification4 `xml:"PrtryId,omitempty"`

	// Specifies the country of the party.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PartyIdentification27) SetName(value string) {
	p.Name = (*Max70Text)(&value)
}

func (p *PartyIdentification27) AddProprietaryIdentification() *GenericIdentification4 {
	p.ProprietaryIdentification = new(GenericIdentification4)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification27) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

// Entity involved in an activity.
type PartyIdentification28 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm"`

	// Unique and unambiguous identifier assigned to a party using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification4 `xml:"PrtryId,omitempty"`
}

func (p *PartyIdentification28) SetName(value string) {
	p.Name = (*Max70Text)(&value)
}

func (p *PartyIdentification28) AddProprietaryIdentification() *GenericIdentification4 {
	p.ProprietaryIdentification = new(GenericIdentification4)
	return p.ProprietaryIdentification
}

// Unique and unambiguous way to identify an organisation.
type PartyIdentification3 struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICOrBEI *AnyBICIdentifier `xml:"BICOrBEI"`
}

func (p *PartyIdentification3) SetBICOrBEI(value string) {
	p.BICOrBEI = (*AnyBICIdentifier)(&value)
}

// Set of elements used to identify a person or an organisation.
type PartyIdentification32 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identification of a party.
	Identification *Party6Choice `xml:"Id,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Set of elements used to indicate how to contact the party.
	ContactDetails *ContactDetails2 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification32) SetName(value string) {
	p.Name = (*Max140Text)(&value)
}

func (p *PartyIdentification32) AddPostalAddress() *PostalAddress6 {
	p.PostalAddress = new(PostalAddress6)
	return p.PostalAddress
}

func (p *PartyIdentification32) AddIdentification() *Party6Choice {
	p.Identification = new(Party6Choice)
	return p.Identification
}

func (p *PartyIdentification32) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification32) AddContactDetails() *ContactDetails2 {
	p.ContactDetails = new(ContactDetails2)
	return p.ContactDetails
}

// Provides information about the beneficial owner of the securities.
type PartyIdentification33 struct {

	// Party that is the beneficial owner of the specified quantity of securities.
	OwnerIdentification *PartyIdentification10Choice `xml:"OwnrId"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Country in which a person is permanently domiciled (the place of a persons permanent home).
	DomicileCountry *CountryCode `xml:"DmclCtry,omitempty"`

	// Holder of the security certifies, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Quantity of securities belonging to the beneficial owner specified.
	OwnedSecuritiesQuantity *FinancialInstrumentQuantity1Choice `xml:"OwndSctiesQty"`

	// Type of certification which is required.
	CertificationType []*BeneficiaryCertificationType2Choice `xml:"CertfctnTp,omitempty"`

	// Provides details relative to the beneficial owner not included within structured fields of this message.
	DeclarationDetails *Max350Text `xml:"DclrtnDtls,omitempty"`
}

func (p *PartyIdentification33) AddOwnerIdentification() *PartyIdentification10Choice {
	p.OwnerIdentification = new(PartyIdentification10Choice)
	return p.OwnerIdentification
}

func (p *PartyIdentification33) AddAlternateIdentification() *AlternatePartyIdentification2 {
	newValue := new(AlternatePartyIdentification2)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

func (p *PartyIdentification33) SetDomicileCountry(value string) {
	p.DomicileCountry = (*CountryCode)(&value)
}

func (p *PartyIdentification33) AddNonDomicileCountry(value string) {
	p.NonDomicileCountry = append(p.NonDomicileCountry, (*CountryCode)(&value))
}

func (p *PartyIdentification33) AddOwnedSecuritiesQuantity() *FinancialInstrumentQuantity1Choice {
	p.OwnedSecuritiesQuantity = new(FinancialInstrumentQuantity1Choice)
	return p.OwnedSecuritiesQuantity
}

func (p *PartyIdentification33) AddCertificationType() *BeneficiaryCertificationType2Choice {
	newValue := new(BeneficiaryCertificationType2Choice)
	p.CertificationType = append(p.CertificationType, newValue)
	return newValue
}

func (p *PartyIdentification33) SetDeclarationDetails(value string) {
	p.DeclarationDetails = (*Max350Text)(&value)
}

// Provides information about identification of the party .
type PartyIdentification35 struct {

	// Identification of a party.
	Identification *PartyIdentification12Choice `xml:"Id"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Provides alternate identification for a party using an id type, a country code and a text field.
	AlternateIdentification []*AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentification35) AddIdentification() *PartyIdentification12Choice {
	p.Identification = new(PartyIdentification12Choice)
	return p.Identification
}

func (p *PartyIdentification35) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentification35) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

func (p *PartyIdentification35) AddAlternateIdentification() *AlternatePartyIdentification2 {
	newValue := new(AlternatePartyIdentification2)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

// Identification of an entity involved in an activity.
type PartyIdentification36 struct {

	// Identification of the party.
	Identification *PartyIdentification12Choice `xml:"Id"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentification36) AddIdentification() *PartyIdentification12Choice {
	p.Identification = new(PartyIdentification12Choice)
	return p.Identification
}

func (p *PartyIdentification36) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

// Identification and additional identification Information on the party.
type PartyIdentification39 struct {

	// Unique identification of the party.
	Identification *PartyIdentification9Choice `xml:"Id"`

	// Ancillary identification information about the party.
	AdditionalIdentificationInformation *PartyAdditionalIdentification2Choice `xml:"AddtlIdInf,omitempty"`
}

func (p *PartyIdentification39) AddIdentification() *PartyIdentification9Choice {
	p.Identification = new(PartyIdentification9Choice)
	return p.Identification
}

func (p *PartyIdentification39) AddAdditionalIdentificationInformation() *PartyAdditionalIdentification2Choice {
	p.AdditionalIdentificationInformation = new(PartyAdditionalIdentification2Choice)
	return p.AdditionalIdentificationInformation
}

// Set of elements used to identify a person or an organisation.
type PartyIdentification40 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identification of a party.
	Identification *PersonIdentification5 `xml:"Id,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Set of elements used to indicate how to contact the party.
	ContactDetails *ContactDetails2 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification40) SetName(value string) {
	p.Name = (*Max140Text)(&value)
}

func (p *PartyIdentification40) AddPostalAddress() *PostalAddress6 {
	p.PostalAddress = new(PostalAddress6)
	return p.PostalAddress
}

func (p *PartyIdentification40) AddIdentification() *PersonIdentification5 {
	p.Identification = new(PersonIdentification5)
	return p.Identification
}

func (p *PartyIdentification40) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification40) AddContactDetails() *ContactDetails2 {
	p.ContactDetails = new(ContactDetails2)
	return p.ContactDetails
}

// Set of elements used to identify a person or an organisation.
type PartyIdentification41 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identification of a party.
	Identification *Party8Choice `xml:"Id,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Set of elements used to indicate how to contact the party.
	ContactDetails *ContactDetails2 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification41) SetName(value string) {
	p.Name = (*Max140Text)(&value)
}

func (p *PartyIdentification41) AddPostalAddress() *PostalAddress6 {
	p.PostalAddress = new(PostalAddress6)
	return p.PostalAddress
}

func (p *PartyIdentification41) AddIdentification() *Party8Choice {
	p.Identification = new(Party8Choice)
	return p.Identification
}

func (p *PartyIdentification41) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification41) AddContactDetails() *ContactDetails2 {
	p.ContactDetails = new(ContactDetails2)
	return p.ContactDetails
}

// Set of elements used to identify a person or an organisation.
type PartyIdentification42 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identification of a party.
	Identification *Party10Choice `xml:"Id,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Set of elements used to indicate how to contact the party.
	ContactDetails *ContactDetails2 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification42) SetName(value string) {
	p.Name = (*Max140Text)(&value)
}

func (p *PartyIdentification42) AddPostalAddress() *PostalAddress6 {
	p.PostalAddress = new(PostalAddress6)
	return p.PostalAddress
}

func (p *PartyIdentification42) AddIdentification() *Party10Choice {
	p.Identification = new(Party10Choice)
	return p.Identification
}

func (p *PartyIdentification42) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification42) AddContactDetails() *ContactDetails2 {
	p.ContactDetails = new(ContactDetails2)
	return p.ContactDetails
}

// Set of elements used to identify a person or an organisation.
type PartyIdentification43 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identification of a party.
	Identification *Party11Choice `xml:"Id,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Set of elements used to indicate how to contact the party.
	ContactDetails *ContactDetails2 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification43) SetName(value string) {
	p.Name = (*Max140Text)(&value)
}

func (p *PartyIdentification43) AddPostalAddress() *PostalAddress6 {
	p.PostalAddress = new(PostalAddress6)
	return p.PostalAddress
}

func (p *PartyIdentification43) AddIdentification() *Party11Choice {
	p.Identification = new(Party11Choice)
	return p.Identification
}

func (p *PartyIdentification43) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification43) AddContactDetails() *ContactDetails2 {
	p.ContactDetails = new(ContactDetails2)
	return p.ContactDetails
}

// Unique and unambiguous way to identify an organisation.
type PartyIdentification44 struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	AlternativeIdentifier []*Max35Text `xml:"AltrntvIdr,omitempty"`
}

func (p *PartyIdentification44) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification44) AddAlternativeIdentifier(value string) {
	p.AlternativeIdentifier = append(p.AlternativeIdentifier, (*Max35Text)(&value))
}

// Set of elements used to identify a person or an organisation.
type PartyIdentification45 struct {

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	Identification *Party8Choice `xml:"Id,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Contact defined for this party.
	ContactDetails []*Contacts3 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification45) AddIdentification() *Party8Choice {
	p.Identification = new(Party8Choice)
	return p.Identification
}

func (p *PartyIdentification45) SetName(value string) {
	p.Name = (*Max35Text)(&value)
}

func (p *PartyIdentification45) AddPostalAddress() *PostalAddress6 {
	p.PostalAddress = new(PostalAddress6)
	return p.PostalAddress
}

func (p *PartyIdentification45) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification45) AddContactDetails() *Contacts3 {
	newValue := new(Contacts3)
	p.ContactDetails = append(p.ContactDetails, newValue)
	return newValue
}

// Identification of an entity involved in an activity.
type PartyIdentification46 struct {

	// Identification of the party.
	Identification *PartyIdentification44Choice `xml:"Id"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentification46) AddIdentification() *PartyIdentification44Choice {
	p.Identification = new(PartyIdentification44Choice)
	return p.Identification
}

func (p *PartyIdentification46) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

// Provides information about identification of the party .
type PartyIdentification47 struct {

	// Identification of a party.
	Identification *PartyIdentification39Choice `xml:"Id"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Provides alternate identification for a party using an id type, a country code and a text field.
	AlternateIdentification []*AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentification47) AddIdentification() *PartyIdentification39Choice {
	p.Identification = new(PartyIdentification39Choice)
	return p.Identification
}

func (p *PartyIdentification47) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentification47) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

func (p *PartyIdentification47) AddAlternateIdentification() *AlternatePartyIdentification2 {
	newValue := new(AlternatePartyIdentification2)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

// Identification of an entity involved in an activity.
type PartyIdentification48 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification44Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentification48) AddIdentification() *PartyIdentification44Choice {
	p.Identification = new(PartyIdentification44Choice)
	return p.Identification
}

func (p *PartyIdentification48) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentification48) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentification48) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentification48) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Provides information about the beneficial owner of the securities.
type PartyIdentification50 struct {

	// Party that is the beneficial owner of the specified quantity of securities.
	OwnerIdentification *PartyIdentification48Choice `xml:"OwnrId"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Country in which a person is permanently domiciled (the place of a persons permanent home).
	DomicileCountry *CountryCode `xml:"DmclCtry,omitempty"`

	// Holder of the security certifies, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Quantity of securities belonging to the beneficial owner specified.
	OwnedSecuritiesQuantity *FinancialInstrumentQuantity1Choice `xml:"OwndSctiesQty"`

	// Type of certification which is required.
	CertificationType []*BeneficiaryCertificationType2Choice `xml:"CertfctnTp,omitempty"`

	// Provides details relative to the beneficial owner not included within structured fields of this message.
	DeclarationDetails *Max350Text `xml:"DclrtnDtls,omitempty"`
}

func (p *PartyIdentification50) AddOwnerIdentification() *PartyIdentification48Choice {
	p.OwnerIdentification = new(PartyIdentification48Choice)
	return p.OwnerIdentification
}

func (p *PartyIdentification50) AddAlternateIdentification() *AlternatePartyIdentification2 {
	newValue := new(AlternatePartyIdentification2)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

func (p *PartyIdentification50) SetDomicileCountry(value string) {
	p.DomicileCountry = (*CountryCode)(&value)
}

func (p *PartyIdentification50) AddNonDomicileCountry(value string) {
	p.NonDomicileCountry = append(p.NonDomicileCountry, (*CountryCode)(&value))
}

func (p *PartyIdentification50) AddOwnedSecuritiesQuantity() *FinancialInstrumentQuantity1Choice {
	p.OwnedSecuritiesQuantity = new(FinancialInstrumentQuantity1Choice)
	return p.OwnedSecuritiesQuantity
}

func (p *PartyIdentification50) AddCertificationType() *BeneficiaryCertificationType2Choice {
	newValue := new(BeneficiaryCertificationType2Choice)
	p.CertificationType = append(p.CertificationType, newValue)
	return newValue
}

func (p *PartyIdentification50) SetDeclarationDetails(value string) {
	p.DeclarationDetails = (*Max350Text)(&value)
}

// Choice of identification of a party. The party can be identified by providing a BIC or a proprietary code.
// Optionally, the client account number can also be provided.
type PartyIdentification54 struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *AnyBICIdentifier `xml:"BIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification29 `xml:"PrtryId"`

	// Identification of a party with its name and address in free text.
	NameAndAddress *NameAndAddress13 `xml:"NmAndAdr,omitempty"`
}

func (p *PartyIdentification54) SetBIC(value string) {
	p.BIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification54) AddProprietaryIdentification() *GenericIdentification29 {
	p.ProprietaryIdentification = new(GenericIdentification29)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification54) AddNameAndAddress() *NameAndAddress13 {
	p.NameAndAddress = new(NameAndAddress13)
	return p.NameAndAddress
}

// Identification of an entity involved in an activity.
type PartyIdentification55 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification68Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentification55) AddIdentification() *PartyIdentification68Choice {
	p.Identification = new(PartyIdentification68Choice)
	return p.Identification
}

func (p *PartyIdentification55) AddAlternateIdentification() *AlternatePartyIdentification5 {
	p.AlternateIdentification = new(AlternatePartyIdentification5)
	return p.AlternateIdentification
}

func (p *PartyIdentification55) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Provides information about the beneficial owner of the securities.
type PartyIdentification56 struct {

	// Party that is the beneficial owner of the specified quantity of securities.
	OwnerIdentification *PartyIdentification48Choice `xml:"OwnrId"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Country in which a person is permanently domiciled (the place of a persons permanent home).
	DomicileCountry *CountryCode `xml:"DmclCtry,omitempty"`

	// Holder of the security certifies, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Quantity of securities belonging to the beneficial owner specified.
	OwnedSecuritiesQuantity *FinancialInstrumentQuantity1Choice `xml:"OwndSctiesQty"`

	// Type of certification which is required.
	CertificationType []*BeneficiaryCertificationType6Choice `xml:"CertfctnTp,omitempty"`

	// Provides details relative to the beneficial owner not included within structured fields of this message.
	DeclarationDetails *Max350Text `xml:"DclrtnDtls,omitempty"`
}

func (p *PartyIdentification56) AddOwnerIdentification() *PartyIdentification48Choice {
	p.OwnerIdentification = new(PartyIdentification48Choice)
	return p.OwnerIdentification
}

func (p *PartyIdentification56) AddAlternateIdentification() *AlternatePartyIdentification2 {
	newValue := new(AlternatePartyIdentification2)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

func (p *PartyIdentification56) SetDomicileCountry(value string) {
	p.DomicileCountry = (*CountryCode)(&value)
}

func (p *PartyIdentification56) AddNonDomicileCountry(value string) {
	p.NonDomicileCountry = append(p.NonDomicileCountry, (*CountryCode)(&value))
}

func (p *PartyIdentification56) AddOwnedSecuritiesQuantity() *FinancialInstrumentQuantity1Choice {
	p.OwnedSecuritiesQuantity = new(FinancialInstrumentQuantity1Choice)
	return p.OwnedSecuritiesQuantity
}

func (p *PartyIdentification56) AddCertificationType() *BeneficiaryCertificationType6Choice {
	newValue := new(BeneficiaryCertificationType6Choice)
	p.CertificationType = append(p.CertificationType, newValue)
	return newValue
}

func (p *PartyIdentification56) SetDeclarationDetails(value string) {
	p.DeclarationDetails = (*Max350Text)(&value)
}

// Describes the details of an organisation.
type PartyIdentification58 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Name by which a party is known and which is usually used to identify that party.
	LegalName *Max140Text `xml:"LglNm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress11 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous way to identify the party.
	Identification *Party13Choice `xml:"Id"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Indicate how to contact the party.
	ContactDetails *ContactDetails3 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification58) SetName(value string) {
	p.Name = (*Max140Text)(&value)
}

func (p *PartyIdentification58) SetLegalName(value string) {
	p.LegalName = (*Max140Text)(&value)
}

func (p *PartyIdentification58) AddPostalAddress() *PostalAddress11 {
	p.PostalAddress = new(PostalAddress11)
	return p.PostalAddress
}

func (p *PartyIdentification58) AddIdentification() *Party13Choice {
	p.Identification = new(Party13Choice)
	return p.Identification
}

func (p *PartyIdentification58) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification58) AddContactDetails() *ContactDetails3 {
	p.ContactDetails = new(ContactDetails3)
	return p.ContactDetails
}

// Identification of a party. The party can be identified by providing the party's name and optionally, the BIC, account number, address, clearing system identification or LEI can also be provided.
type PartyIdentification59 struct {

	// Identification of the party expressed as the party's name.
	PartyName *Max34Text `xml:"PtyNm,omitempty"`

	// Identification of the party expressed as a BIC and an optional alternative identifier.
	AnyBIC *PartyIdentification44 `xml:"AnyBIC,omitempty"`

	// Identification of the party's account number.
	AccountNumber *Max34Text `xml:"AcctNb,omitempty"`

	// Identification of the party's address.
	Address *Max105Text `xml:"Adr,omitempty"`

	// Choice of a clearing system identifier.
	ClearingSystemIdentification *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty"`

	// Identification of the Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`
}

func (p *PartyIdentification59) SetPartyName(value string) {
	p.PartyName = (*Max34Text)(&value)
}

func (p *PartyIdentification59) AddAnyBIC() *PartyIdentification44 {
	p.AnyBIC = new(PartyIdentification44)
	return p.AnyBIC
}

func (p *PartyIdentification59) SetAccountNumber(value string) {
	p.AccountNumber = (*Max34Text)(&value)
}

func (p *PartyIdentification59) SetAddress(value string) {
	p.Address = (*Max105Text)(&value)
}

func (p *PartyIdentification59) AddClearingSystemIdentification() *ClearingSystemIdentification2Choice {
	p.ClearingSystemIdentification = new(ClearingSystemIdentification2Choice)
	return p.ClearingSystemIdentification
}

func (p *PartyIdentification59) SetLegalEntityIdentifier(value string) {
	p.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

// Identification of a party by fund name, name and address or an LEI.
type PartyIdentification60 struct {

	// Identification of a fund.
	FundIdentification *Max35Text `xml:"FndId"`

	// Identification of the party expressed as name and an optional address and an optional alternative identifier.
	NameAndAddress *NameAndAddress8 `xml:"NmAndAdr,omitempty"`

	// Identification of the Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`
}

func (p *PartyIdentification60) SetFundIdentification(value string) {
	p.FundIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentification60) AddNameAndAddress() *NameAndAddress8 {
	p.NameAndAddress = new(NameAndAddress8)
	return p.NameAndAddress
}

func (p *PartyIdentification60) SetLegalEntityIdentifier(value string) {
	p.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

// Identification of a party.
type PartyIdentification62 struct {

	// Identification of the financial institution expressed as a BIC.
	BICFI *BICFIIdentifier `xml:"BICFI,omitempty"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification1 `xml:"PrtryId,omitempty"`

	// Name and address of the party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr,omitempty"`
}

func (p *PartyIdentification62) SetBICFI(value string) {
	p.BICFI = (*BICFIIdentifier)(&value)
}

func (p *PartyIdentification62) AddProprietaryIdentification() *GenericIdentification1 {
	p.ProprietaryIdentification = new(GenericIdentification1)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification62) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

// Identification of an entity involved in an activity.
type PartyIdentification63 struct {

	// Identification of the party.
	PartyIdentification *PartyIdentification75Choice `xml:"PtyId"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentification63) AddPartyIdentification() *PartyIdentification75Choice {
	p.PartyIdentification = new(PartyIdentification75Choice)
	return p.PartyIdentification
}

func (p *PartyIdentification63) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

// Identification of a party.
type PartyIdentification64 struct {

	// Identification of the party expressed as a BIC.
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC,omitempty"`

	// Unique and unambiguous identifier, as assigned to the party using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification1 `xml:"PrtryId,omitempty"`

	// Name and address of the party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr,omitempty"`
}

func (p *PartyIdentification64) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification64) AddProprietaryIdentification() *GenericIdentification1 {
	p.ProprietaryIdentification = new(GenericIdentification1)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification64) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

// Set of elements used to identify a person or an organisation.
type PartyIdentification71 struct {

	// Unique identification of the party.
	Identification *PartyIdentification40Choice `xml:"Id"`

	// Ancillary identification information about the party.
	AdditionalIdentificationInformation *PartyAdditionalIdentification2Choice `xml:"AddtlIdInf,omitempty"`
}

func (p *PartyIdentification71) AddIdentification() *PartyIdentification40Choice {
	p.Identification = new(PartyIdentification40Choice)
	return p.Identification
}

func (p *PartyIdentification71) AddAdditionalIdentificationInformation() *PartyAdditionalIdentification2Choice {
	p.AdditionalIdentificationInformation = new(PartyAdditionalIdentification2Choice)
	return p.AdditionalIdentificationInformation
}

// Set of elements used to identify an organization or a person.
type PartyIdentification72 struct {

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	PartyIdentification *PartyIdentification43 `xml:"PtyId"`

	// Identifier and name of an organisation that is allocated by an institution.
	LegalOrganisation *LegalOrganisation1 `xml:"LglOrg,omitempty"`

	// Tax registration details.
	TaxParty *TaxParty1 `xml:"TaxPty,omitempty"`
}

func (p *PartyIdentification72) AddPartyIdentification() *PartyIdentification43 {
	p.PartyIdentification = new(PartyIdentification43)
	return p.PartyIdentification
}

func (p *PartyIdentification72) AddLegalOrganisation() *LegalOrganisation1 {
	p.LegalOrganisation = new(LegalOrganisation1)
	return p.LegalOrganisation
}

func (p *PartyIdentification72) AddTaxParty() *TaxParty1 {
	p.TaxParty = new(TaxParty1)
	return p.TaxParty
}

// Identification of an entity involved in an activity.
type PartyIdentification75 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification44Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentification75) AddIdentification() *PartyIdentification44Choice {
	p.Identification = new(PartyIdentification44Choice)
	return p.Identification
}

func (p *PartyIdentification75) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentification75) AddAlternateIdentification() *AlternatePartyIdentification7 {
	p.AlternateIdentification = new(AlternatePartyIdentification7)
	return p.AlternateIdentification
}

func (p *PartyIdentification75) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentification75) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentification75) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Identification of a person or an organisation.
type PartyIdentification77 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress19 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identification of a party.
	Identification *Party11Choice `xml:"Id,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Set of elements used to indicate how to contact the party.
	ContactDetails *ContactDetails2 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification77) SetName(value string) {
	p.Name = (*Max140Text)(&value)
}

func (p *PartyIdentification77) AddPostalAddress() *PostalAddress19 {
	p.PostalAddress = new(PostalAddress19)
	return p.PostalAddress
}

func (p *PartyIdentification77) AddIdentification() *Party11Choice {
	p.Identification = new(Party11Choice)
	return p.Identification
}

func (p *PartyIdentification77) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification77) AddContactDetails() *ContactDetails2 {
	p.ContactDetails = new(ContactDetails2)
	return p.ContactDetails
}

// Identification of a person, a financial institution or a non-financial institution.
type PartyIdentification8 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress1 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous way of identifying an organisation or an individual person.
	Identification *Party2Choice `xml:"Id,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`
}

func (p *PartyIdentification8) SetName(value string) {
	p.Name = (*Max70Text)(&value)
}

func (p *PartyIdentification8) AddPostalAddress() *PostalAddress1 {
	p.PostalAddress = new(PostalAddress1)
	return p.PostalAddress
}

func (p *PartyIdentification8) AddIdentification() *Party2Choice {
	p.Identification = new(Party2Choice)
	return p.Identification
}

func (p *PartyIdentification8) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

// Identification of an entity involved in an activity.
type PartyIdentification91 struct {

	// Identification of the party.
	Identification *PartyIdentification44Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentification91) AddIdentification() *PartyIdentification44Choice {
	p.Identification = new(PartyIdentification44Choice)
	return p.Identification
}

func (p *PartyIdentification91) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentification91) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

// Provides information about identification of the party .
type PartyIdentification92 struct {

	// Identification of a party.
	Identification *PartyIdentification44Choice `xml:"Id"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides alternate identification for a party using an id type, a country code and a text field.
	AlternateIdentification []*AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentification92) AddIdentification() *PartyIdentification44Choice {
	p.Identification = new(PartyIdentification44Choice)
	return p.Identification
}

func (p *PartyIdentification92) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentification92) AddAlternateIdentification() *AlternatePartyIdentification7 {
	newValue := new(AlternatePartyIdentification7)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

// Provides information about the beneficial owner of the securities.
type PartyIdentification93 struct {

	// Party that is the beneficial owner of the specified quantity of securities.
	OwnerIdentification *PartyIdentification71Choice `xml:"OwnrId"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`

	// Country in which a person is permanently domiciled (the place of a persons permanent home).
	DomicileCountry *CountryCode `xml:"DmclCtry,omitempty"`

	// Holder of the security certifies, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Quantity of securities belonging to the beneficial owner specified.
	OwnedSecuritiesQuantity *FinancialInstrumentQuantity1Choice `xml:"OwndSctiesQty"`

	// Type of certification which is required.
	CertificationType []*BeneficiaryCertificationType10Choice `xml:"CertfctnTp,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *Max350Text `xml:"CertfctnBrkdwn,omitempty"`
}

func (p *PartyIdentification93) AddOwnerIdentification() *PartyIdentification71Choice {
	p.OwnerIdentification = new(PartyIdentification71Choice)
	return p.OwnerIdentification
}

func (p *PartyIdentification93) AddAlternateIdentification() *AlternatePartyIdentification7 {
	newValue := new(AlternatePartyIdentification7)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

func (p *PartyIdentification93) SetDomicileCountry(value string) {
	p.DomicileCountry = (*CountryCode)(&value)
}

func (p *PartyIdentification93) AddNonDomicileCountry(value string) {
	p.NonDomicileCountry = append(p.NonDomicileCountry, (*CountryCode)(&value))
}

func (p *PartyIdentification93) AddOwnedSecuritiesQuantity() *FinancialInstrumentQuantity1Choice {
	p.OwnedSecuritiesQuantity = new(FinancialInstrumentQuantity1Choice)
	return p.OwnedSecuritiesQuantity
}

func (p *PartyIdentification93) AddCertificationType() *BeneficiaryCertificationType10Choice {
	newValue := new(BeneficiaryCertificationType10Choice)
	p.CertificationType = append(p.CertificationType, newValue)
	return newValue
}

func (p *PartyIdentification93) SetCertificationBreakdown(value string) {
	p.CertificationBreakdown = (*Max350Text)(&value)
}

// Identification of a person or an organisation.
type PartyIdentification95 struct {

	// Identifier for an organisation that is allocated by an institution.
	Identification *PartyIdentification70Choice `xml:"Id,omitempty"`

	// Identification of the organisation with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`
}

func (p *PartyIdentification95) AddIdentification() *PartyIdentification70Choice {
	p.Identification = new(PartyIdentification70Choice)
	return p.Identification
}

func (p *PartyIdentification95) SetLegalEntityIdentifier(value string) {
	p.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

// Identification of a person or an organisation.
type PartyIdentification96 struct {

	// Identification of the organisation.
	Identification *PartyIdentification96Choice `xml:"Id,omitempty"`

	// Identification of the organisation with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`
}

func (p *PartyIdentification96) AddIdentification() *PartyIdentification96Choice {
	p.Identification = new(PartyIdentification96Choice)
	return p.Identification
}

func (p *PartyIdentification96) SetLegalEntityIdentifier(value string) {
	p.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

// Party involved in the settlement chain.
type PartyIdentification97 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification70Choice `xml:"PtyId"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingReference *Max35Text `xml:"PrcgRef,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (p *PartyIdentification97) AddPartyIdentification() *PartyIdentification70Choice {
	p.PartyIdentification = new(PartyIdentification70Choice)
	return p.PartyIdentification
}

func (p *PartyIdentification97) SetProcessingReference(value string) {
	p.ProcessingReference = (*Max35Text)(&value)
}

func (p *PartyIdentification97) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentification97) AddContactPerson() *ContactIdentification2 {
	p.ContactPerson = new(ContactIdentification2)
	return p.ContactPerson
}

// Identification of the party.
type PartyIdentification98 struct {

	// Unique identification of the party.
	Identification *PartyIdentification92Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification98) AddIdentification() *PartyIdentification92Choice {
	p.Identification = new(PartyIdentification92Choice)
	return p.Identification
}

func (p *PartyIdentification98) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

// Identification of the party.
type PartyIdentification99 struct {

	// Unique identification of the party.
	Identification *PartyIdentification93Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PartyIdentification99) AddIdentification() *PartyIdentification93Choice {
	p.Identification = new(PartyIdentification93Choice)
	return p.Identification
}

func (p *PartyIdentification99) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
