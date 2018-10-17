package iso20022

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification2 struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *BICIdentifier `xml:"BIC,omitempty"`

	// International Business Entity Identifier to uniquely identify business entities playing a role in the lifecycle of and events related to a financial instrument. (tentative - to be confirmed)
	IBEI *IBEIIdentifier `xml:"IBEI,omitempty"`

	// Code allocated to a non-financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BEI *BEIIdentifier `xml:"BEI,omitempty"`

	// Global Location Number. A non-significant reference number used to identify legal entities, functional entities, or physical entities according to the European Association for Numbering (EAN) numbering scheme rules.The number is used to retrieve detailed information that is linked to it.
	EANGLN *EANGLNIdentifier `xml:"EANGLN,omitempty"`

	// (United States) Clearing House Interbank Payments System (CHIPS) Universal Identification (UID) - identifies entities that own accounts at CHIPS participating financial institutions, through which CHIPS payments are effected. The CHIPS UID is assigned by the New York Clearing House.
	CHIPSUniversalIdentification *CHIPSUniversalIdentifier `xml:"USCHU,omitempty"`

	// Data Universal Numbering System. A unique identification number provided by Dun & Bradstreet to identify an organization.
	DUNS *DunsIdentifier `xml:"DUNS,omitempty"`

	// Unique and unambiguous assignment made by a specific bank to identify a relationship as defined between the bank and its client.
	BankPartyIdentification *Max35Text `xml:"BkPtyId,omitempty"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb,omitempty"`

	// Unique and unambiguous identifier for an organisation that is allocated by an institution.
	ProprietaryIdentification *GenericIdentification3 `xml:"PrtryId,omitempty"`
}

func (o *OrganisationIdentification2) SetBIC(value string) {
	o.BIC = (*BICIdentifier)(&value)
}

func (o *OrganisationIdentification2) SetIBEI(value string) {
	o.IBEI = (*IBEIIdentifier)(&value)
}

func (o *OrganisationIdentification2) SetBEI(value string) {
	o.BEI = (*BEIIdentifier)(&value)
}

func (o *OrganisationIdentification2) SetEANGLN(value string) {
	o.EANGLN = (*EANGLNIdentifier)(&value)
}

func (o *OrganisationIdentification2) SetCHIPSUniversalIdentification(value string) {
	o.CHIPSUniversalIdentification = (*CHIPSUniversalIdentifier)(&value)
}

func (o *OrganisationIdentification2) SetDUNS(value string) {
	o.DUNS = (*DunsIdentifier)(&value)
}

func (o *OrganisationIdentification2) SetBankPartyIdentification(value string) {
	o.BankPartyIdentification = (*Max35Text)(&value)
}

func (o *OrganisationIdentification2) SetTaxIdentificationNumber(value string) {
	o.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (o *OrganisationIdentification2) AddProprietaryIdentification() *GenericIdentification3 {
	o.ProprietaryIdentification = new(GenericIdentification3)
	return o.ProprietaryIdentification
}

// Unique identification, as assigned by an organisation, to unambiguously identify a party.
type OrganisationIdentification28 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identification of a party.
	Identification *OrganisationIdentification8 `xml:"Id,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Set of elements used to indicate how to contact the party.
	ContactDetails *ContactDetails2 `xml:"CtctDtls"`
}

func (o *OrganisationIdentification28) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *OrganisationIdentification28) AddPostalAddress() *PostalAddress6 {
	o.PostalAddress = new(PostalAddress6)
	return o.PostalAddress
}

func (o *OrganisationIdentification28) AddIdentification() *OrganisationIdentification8 {
	o.Identification = new(OrganisationIdentification8)
	return o.Identification
}

func (o *OrganisationIdentification28) SetCountryOfResidence(value string) {
	o.CountryOfResidence = (*CountryCode)(&value)
}

func (o *OrganisationIdentification28) AddContactDetails() *ContactDetails2 {
	o.ContactDetails = new(ContactDetails2)
	return o.ContactDetails
}

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification4 struct {

	// Code allocated to a financial institution or non financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICOrBEI *AnyBICIdentifier `xml:"BICOrBEI,omitempty"`

	// Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	Other []*GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

func (o *OrganisationIdentification4) SetBICOrBEI(value string) {
	o.BICOrBEI = (*AnyBICIdentifier)(&value)
}

func (o *OrganisationIdentification4) AddOther() *GenericOrganisationIdentification1 {
	newValue := new(GenericOrganisationIdentification1)
	o.Other = append(o.Other, newValue)
	return newValue
}

// Official identification of an organisation (legal entity) in a specific register.
type OrganisationIdentification5 struct {

	// Name of the register of legal entities.
	RegistrationNumber *Max35Text `xml:"RegnNb"`

	// Name of the register managed by a registration authority.
	RegisterName *Max35Text `xml:"RegrNm,omitempty"`
}

func (o *OrganisationIdentification5) SetRegistrationNumber(value string) {
	o.RegistrationNumber = (*Max35Text)(&value)
}

func (o *OrganisationIdentification5) SetRegisterName(value string) {
	o.RegisterName = (*Max35Text)(&value)
}

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification6 struct {

	// Code allocated to organisations by the ISO 9362 Registration Authority, under an international identification scheme, as described in the latest version of the standard ISO 9362 Banking (Banking telecommunication messages, Business Identifier Codes).
	BIC *AnyBICIdentifier `xml:"BIC,omitempty"`

	// Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	Other []*GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

func (o *OrganisationIdentification6) SetBIC(value string) {
	o.BIC = (*AnyBICIdentifier)(&value)
}

func (o *OrganisationIdentification6) AddOther() *GenericOrganisationIdentification1 {
	newValue := new(GenericOrganisationIdentification1)
	o.Other = append(o.Other, newValue)
	return newValue
}

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification7 struct {

	// Code allocated to an institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC,omitempty"`

	// Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	Other []*GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

func (o *OrganisationIdentification7) SetAnyBIC(value string) {
	o.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (o *OrganisationIdentification7) AddOther() *GenericOrganisationIdentification1 {
	newValue := new(GenericOrganisationIdentification1)
	o.Other = append(o.Other, newValue)
	return newValue
}

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification8 struct {

	// Code allocated to a financial institution or non financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC,omitempty"`

	// Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	Other []*GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

func (o *OrganisationIdentification8) SetAnyBIC(value string) {
	o.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (o *OrganisationIdentification8) AddOther() *GenericOrganisationIdentification1 {
	newValue := new(GenericOrganisationIdentification1)
	o.Other = append(o.Other, newValue)
	return newValue
}

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification9 struct {

	// Code allocated to an institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC,omitempty"`

	// Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	Other []*GenericOrganisationIdentification2 `xml:"Othr,omitempty"`
}

func (o *OrganisationIdentification9) SetAnyBIC(value string) {
	o.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (o *OrganisationIdentification9) AddOther() *GenericOrganisationIdentification2 {
	newValue := new(GenericOrganisationIdentification2)
	o.Other = append(o.Other, newValue)
	return newValue
}
