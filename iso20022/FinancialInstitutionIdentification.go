package iso20022

// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
type FinancialInstitutionIdentification1 struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *BICIdentifier `xml:"BIC,omitempty"`

	// Unique and unambiguous identifier of a clearing system member, as assigned by the system or system administrator.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentificationChoice `xml:"ClrSysMmbId,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress1 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification3 `xml:"PrtryId,omitempty"`
}

func (f *FinancialInstitutionIdentification1) SetBIC(value string) {
	f.BIC = (*BICIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification1) AddClearingSystemMemberIdentification() *ClearingSystemMemberIdentificationChoice {
	f.ClearingSystemMemberIdentification = new(ClearingSystemMemberIdentificationChoice)
	return f.ClearingSystemMemberIdentification
}

func (f *FinancialInstitutionIdentification1) SetName(value string) {
	f.Name = (*Max70Text)(&value)
}

func (f *FinancialInstitutionIdentification1) AddPostalAddress() *PostalAddress1 {
	f.PostalAddress = new(PostalAddress1)
	return f.PostalAddress
}

func (f *FinancialInstitutionIdentification1) AddProprietaryIdentification() *GenericIdentification3 {
	f.ProprietaryIdentification = new(GenericIdentification3)
	return f.ProprietaryIdentification
}

// Identification of a financial institution.
type FinancialInstitutionIdentification10 struct {

	// Unique identification of the party.
	Party *FinancialInstitutionIdentification8Choice `xml:"Pty"`

	// Legal entity identification as an alternate identification for the party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (f *FinancialInstitutionIdentification10) AddParty() *FinancialInstitutionIdentification8Choice {
	f.Party = new(FinancialInstitutionIdentification8Choice)
	return f.Party
}

func (f *FinancialInstitutionIdentification10) SetLEI(value string) {
	f.LEI = (*LEIIdentifier)(&value)
}

// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
type FinancialInstitutionIdentification3 struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *BICIdentifier `xml:"BIC,omitempty"`

	// Unique and unambiguous identifier of a clearing system member, as assigned by the system or system administrator.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification3Choice `xml:"ClrSysMmbId,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress1 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification3 `xml:"PrtryId,omitempty"`
}

func (f *FinancialInstitutionIdentification3) SetBIC(value string) {
	f.BIC = (*BICIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification3) AddClearingSystemMemberIdentification() *ClearingSystemMemberIdentification3Choice {
	f.ClearingSystemMemberIdentification = new(ClearingSystemMemberIdentification3Choice)
	return f.ClearingSystemMemberIdentification
}

func (f *FinancialInstitutionIdentification3) SetName(value string) {
	f.Name = (*Max70Text)(&value)
}

func (f *FinancialInstitutionIdentification3) AddPostalAddress() *PostalAddress1 {
	f.PostalAddress = new(PostalAddress1)
	return f.PostalAddress
}

func (f *FinancialInstitutionIdentification3) AddProprietaryIdentification() *GenericIdentification3 {
	f.ProprietaryIdentification = new(GenericIdentification3)
	return f.ProprietaryIdentification
}

// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
type FinancialInstitutionIdentification6 struct {

	// Unique and unambiguous identifier of a clearing system member, as assigned by the system or system administrator.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification2Choice `xml:"ClrSysMmbId,omitempty"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification4 `xml:"PrtryId,omitempty"`

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *BICIdentifier `xml:"BIC,omitempty"`
}

func (f *FinancialInstitutionIdentification6) AddClearingSystemMemberIdentification() *ClearingSystemMemberIdentification2Choice {
	f.ClearingSystemMemberIdentification = new(ClearingSystemMemberIdentification2Choice)
	return f.ClearingSystemMemberIdentification
}

func (f *FinancialInstitutionIdentification6) AddProprietaryIdentification() *GenericIdentification4 {
	f.ProprietaryIdentification = new(GenericIdentification4)
	return f.ProprietaryIdentification
}

func (f *FinancialInstitutionIdentification6) SetBIC(value string) {
	f.BIC = (*BICIdentifier)(&value)
}

// Set of elements used to identify a financial institution.
type FinancialInstitutionIdentification7 struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *BICIdentifier `xml:"BIC,omitempty"`

	// Information used to identify a member within a clearing system.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`

	// Name by which an agent is known and which is usually used to identify that agent.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Unique identification of an agent, as assigned by an institution, using an identification scheme.
	Other *GenericFinancialIdentification1 `xml:"Othr,omitempty"`
}

func (f *FinancialInstitutionIdentification7) SetBIC(value string) {
	f.BIC = (*BICIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification7) AddClearingSystemMemberIdentification() *ClearingSystemMemberIdentification2 {
	f.ClearingSystemMemberIdentification = new(ClearingSystemMemberIdentification2)
	return f.ClearingSystemMemberIdentification
}

func (f *FinancialInstitutionIdentification7) SetName(value string) {
	f.Name = (*Max140Text)(&value)
}

func (f *FinancialInstitutionIdentification7) AddPostalAddress() *PostalAddress6 {
	f.PostalAddress = new(PostalAddress6)
	return f.PostalAddress
}

func (f *FinancialInstitutionIdentification7) AddOther() *GenericFinancialIdentification1 {
	f.Other = new(GenericFinancialIdentification1)
	return f.Other
}

// Set of elements used to identify a financial institution.
type FinancialInstitutionIdentification8 struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICFI *BICFIIdentifier `xml:"BICFI,omitempty"`

	// Information used to identify a member within a clearing system.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`

	// Name by which an agent is known and which is usually used to identify that agent.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Unique identification of an agent, as assigned by an institution, using an identification scheme.
	Other *GenericFinancialIdentification1 `xml:"Othr,omitempty"`
}

func (f *FinancialInstitutionIdentification8) SetBICFI(value string) {
	f.BICFI = (*BICFIIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification8) AddClearingSystemMemberIdentification() *ClearingSystemMemberIdentification2 {
	f.ClearingSystemMemberIdentification = new(ClearingSystemMemberIdentification2)
	return f.ClearingSystemMemberIdentification
}

func (f *FinancialInstitutionIdentification8) SetName(value string) {
	f.Name = (*Max140Text)(&value)
}

func (f *FinancialInstitutionIdentification8) AddPostalAddress() *PostalAddress6 {
	f.PostalAddress = new(PostalAddress6)
	return f.PostalAddress
}

func (f *FinancialInstitutionIdentification8) AddOther() *GenericFinancialIdentification1 {
	f.Other = new(GenericFinancialIdentification1)
	return f.Other
}

// Set of elements used to identify a financial institution.
type FinancialInstitutionIdentification9 struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICFI *BICFIIdentifier `xml:"BICFI,omitempty"`

	// Information used to identify a member within a clearing system.
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`

	// Unique identification of an agent, as assigned by an institution, using an identification scheme.
	Other *GenericFinancialIdentification1 `xml:"Othr,omitempty"`
}

func (f *FinancialInstitutionIdentification9) SetBICFI(value string) {
	f.BICFI = (*BICFIIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification9) AddClearingSystemMemberIdentification() *ClearingSystemMemberIdentification2 {
	f.ClearingSystemMemberIdentification = new(ClearingSystemMemberIdentification2)
	return f.ClearingSystemMemberIdentification
}

func (f *FinancialInstitutionIdentification9) AddOther() *GenericFinancialIdentification1 {
	f.Other = new(GenericFinancialIdentification1)
	return f.Other
}
