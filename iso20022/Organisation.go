package iso20022

// Information which describes the organisation.
type Organisation12 struct {

	// Name by which a party is known and which is usually used to identify that party.
	FullLegalName *Max350Text `xml:"FullLglNm"`

	// Name used by a business for commercial purposes, although its registered legal name, used for contracts and other formal situations, may be another.
	TradingName *Max350Text `xml:"TradgNm,omitempty"`

	// Country in which the organisation has its business activity.
	CountryOfOperation *CountryCode `xml:"CtryOfOpr"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Is an operational address, for example, of a shared services center.
	OperationalAddress *PostalAddress6 `xml:"OprlAdr,omitempty"`

	// Is the address where the business activity is taking place.
	BusinessAddress *PostalAddress6 `xml:"BizAdr,omitempty"`

	// Is the address where the entity resides and is registered. More generically, it is the home address (Residential address).
	LegalAddress *PostalAddress6 `xml:"LglAdr"`

	// Address where invoices must be sent.
	BillingAddress *PostalAddress6 `xml:"BllgAdr,omitempty"`

	// Unique and unambiguous way of identifying an organisation.
	OrganisationIdentification *OrganisationIdentification8 `xml:"OrgId"`

	// Person in the customer's organisation who can be contacted by the account servicer in relation to the account(s) identified in this instruction.
	RepresentativeOfficer []*PartyIdentification40 `xml:"RprtvOffcr,omitempty"`

	// Person responsible of the treasury department within the customer’s organisation.
	TreasuryManager *PartyIdentification40 `xml:"TrsrMgr,omitempty"`

	// Person that has the mandate to delegate authority, to assign mandates to other individuals (mandate holders) to perform specific bank operations on the account.
	MainMandateHolder []*PartyIdentification40 `xml:"MainMndtHldr,omitempty"`

	// Person that may be the potential sender of a message related to the life cycle of the account.
	Sender []*PartyIdentification40 `xml:"Sndr,omitempty"`

	// Person that is officially and legally mandated to represent the organisation. Depending on legislation, the legal representative(s) might for instance be assigned by the Board, identified in the by-laws of the organisation, be publicly announced in the official journal of a country, etc.
	LegalRepresentative []*PartyIdentification40 `xml:"LglRprtv,omitempty"`
}

func (o *Organisation12) SetFullLegalName(value string) {
	o.FullLegalName = (*Max350Text)(&value)
}

func (o *Organisation12) SetTradingName(value string) {
	o.TradingName = (*Max350Text)(&value)
}

func (o *Organisation12) SetCountryOfOperation(value string) {
	o.CountryOfOperation = (*CountryCode)(&value)
}

func (o *Organisation12) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation12) AddOperationalAddress() *PostalAddress6 {
	o.OperationalAddress = new(PostalAddress6)
	return o.OperationalAddress
}

func (o *Organisation12) AddBusinessAddress() *PostalAddress6 {
	o.BusinessAddress = new(PostalAddress6)
	return o.BusinessAddress
}

func (o *Organisation12) AddLegalAddress() *PostalAddress6 {
	o.LegalAddress = new(PostalAddress6)
	return o.LegalAddress
}

func (o *Organisation12) AddBillingAddress() *PostalAddress6 {
	o.BillingAddress = new(PostalAddress6)
	return o.BillingAddress
}

func (o *Organisation12) AddOrganisationIdentification() *OrganisationIdentification8 {
	o.OrganisationIdentification = new(OrganisationIdentification8)
	return o.OrganisationIdentification
}

func (o *Organisation12) AddRepresentativeOfficer() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.RepresentativeOfficer = append(o.RepresentativeOfficer, newValue)
	return newValue
}

func (o *Organisation12) AddTreasuryManager() *PartyIdentification40 {
	o.TreasuryManager = new(PartyIdentification40)
	return o.TreasuryManager
}

func (o *Organisation12) AddMainMandateHolder() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.MainMandateHolder = append(o.MainMandateHolder, newValue)
	return newValue
}

func (o *Organisation12) AddSender() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.Sender = append(o.Sender, newValue)
	return newValue
}

func (o *Organisation12) AddLegalRepresentative() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.LegalRepresentative = append(o.LegalRepresentative, newValue)
	return newValue
}

// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
type Organisation13 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification4Choice `xml:"Id,omitempty"`

	// Purpose of the organisation, eg, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country of taxation of an organisation.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb,omitempty"`

	// Number assigned by a national registration authority to an entity.
	NationalRegistrationNumber *Max35Text `xml:"NtlRegnNb,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress3 `xml:"PstlAdr"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Additional regulatory information about the investor that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`
}

func (o *Organisation13) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *Organisation13) AddIdentification() *PartyIdentification4Choice {
	o.Identification = new(PartyIdentification4Choice)
	return o.Identification
}

func (o *Organisation13) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation13) SetTaxationCountry(value string) {
	o.TaxationCountry = (*CountryCode)(&value)
}

func (o *Organisation13) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation13) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation13) SetTaxIdentificationNumber(value string) {
	o.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (o *Organisation13) SetNationalRegistrationNumber(value string) {
	o.NationalRegistrationNumber = (*Max35Text)(&value)
}

func (o *Organisation13) AddPostalAddress() *PostalAddress3 {
	newValue := new(PostalAddress3)
	o.PostalAddress = append(o.PostalAddress, newValue)
	return newValue
}

func (o *Organisation13) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	o.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return o.PrimaryCommunicationAddress
}

func (o *Organisation13) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	o.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return o.SecondaryCommunicationAddress
}

func (o *Organisation13) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	o.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return o.AdditionalRegulatoryInformation
}

// Information which describes the organisation.
type Organisation14 struct {

	// Name by which a party is known and which is usually used to identify that party.
	FullLegalName *Max350Text `xml:"FullLglNm,omitempty"`

	// Unique and unambiguous way of identifying an organisation.
	OrganisationIdentification *OrganisationIdentification8 `xml:"OrgId"`
}

func (o *Organisation14) SetFullLegalName(value string) {
	o.FullLegalName = (*Max350Text)(&value)
}

func (o *Organisation14) AddOrganisationIdentification() *OrganisationIdentification8 {
	o.OrganisationIdentification = new(OrganisationIdentification8)
	return o.OrganisationIdentification
}

// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
type Organisation15 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification4Choice `xml:"Id,omitempty"`

	// Purpose of the organisation, eg, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country of taxation of an organisation.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb,omitempty"`

	// Number assigned by a national registration authority to an entity.
	NationalRegistrationNumber *Max35Text `xml:"NtlRegnNb,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	ModifiedPostalAddress []*ModificationScope1 `xml:"ModfdPstlAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Additional regulatory information about the investor that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`
}

func (o *Organisation15) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *Organisation15) AddIdentification() *PartyIdentification4Choice {
	o.Identification = new(PartyIdentification4Choice)
	return o.Identification
}

func (o *Organisation15) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation15) SetTaxationCountry(value string) {
	o.TaxationCountry = (*CountryCode)(&value)
}

func (o *Organisation15) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation15) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation15) SetTaxIdentificationNumber(value string) {
	o.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (o *Organisation15) SetNationalRegistrationNumber(value string) {
	o.NationalRegistrationNumber = (*Max35Text)(&value)
}

func (o *Organisation15) AddModifiedPostalAddress() *ModificationScope1 {
	newValue := new(ModificationScope1)
	o.ModifiedPostalAddress = append(o.ModifiedPostalAddress, newValue)
	return newValue
}

func (o *Organisation15) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	o.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return o.PrimaryCommunicationAddress
}

func (o *Organisation15) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	o.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return o.SecondaryCommunicationAddress
}

func (o *Organisation15) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	o.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return o.AdditionalRegulatoryInformation
}

// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
type Organisation16 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, for example, Dun & Bradstreet Identification.
	Identification *PartyIdentification4Choice `xml:"Id,omitempty"`

	// Purpose of the organisation, for example, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country of taxation of an organisation.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Tax identification information.
	TaxIdentification []*TaxIdentification2 `xml:"TaxId,omitempty"`

	// Number assigned by a national registration authority to an entity.
	NationalRegistrationNumber *Max35Text `xml:"NtlRegnNb,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress3 `xml:"PstlAdr"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Additional regulatory information about the investor that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`
}

func (o *Organisation16) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *Organisation16) AddIdentification() *PartyIdentification4Choice {
	o.Identification = new(PartyIdentification4Choice)
	return o.Identification
}

func (o *Organisation16) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation16) SetTaxationCountry(value string) {
	o.TaxationCountry = (*CountryCode)(&value)
}

func (o *Organisation16) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation16) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation16) AddTaxIdentification() *TaxIdentification2 {
	newValue := new(TaxIdentification2)
	o.TaxIdentification = append(o.TaxIdentification, newValue)
	return newValue
}

func (o *Organisation16) SetNationalRegistrationNumber(value string) {
	o.NationalRegistrationNumber = (*Max35Text)(&value)
}

func (o *Organisation16) AddPostalAddress() *PostalAddress3 {
	newValue := new(PostalAddress3)
	o.PostalAddress = append(o.PostalAddress, newValue)
	return newValue
}

func (o *Organisation16) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	o.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return o.PrimaryCommunicationAddress
}

func (o *Organisation16) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	o.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return o.SecondaryCommunicationAddress
}

func (o *Organisation16) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	o.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return o.AdditionalRegulatoryInformation
}

// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
type Organisation17 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, for example, Dun & Bradstreet Identification.
	Identification *PartyIdentification4Choice `xml:"Id,omitempty"`

	// Purpose of the organisation, for example, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country of taxation of an organisation.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Tax identification information.
	TaxIdentification *TaxIdentification2 `xml:"TaxId,omitempty"`

	// Number assigned by a national registration authority to an entity.
	NationalRegistrationNumber *Max35Text `xml:"NtlRegnNb,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	ModifiedPostalAddress []*ModificationScope1 `xml:"ModfdPstlAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Additional regulatory information about the investor that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`
}

func (o *Organisation17) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *Organisation17) AddIdentification() *PartyIdentification4Choice {
	o.Identification = new(PartyIdentification4Choice)
	return o.Identification
}

func (o *Organisation17) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation17) SetTaxationCountry(value string) {
	o.TaxationCountry = (*CountryCode)(&value)
}

func (o *Organisation17) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation17) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation17) AddTaxIdentification() *TaxIdentification2 {
	o.TaxIdentification = new(TaxIdentification2)
	return o.TaxIdentification
}

func (o *Organisation17) SetNationalRegistrationNumber(value string) {
	o.NationalRegistrationNumber = (*Max35Text)(&value)
}

func (o *Organisation17) AddModifiedPostalAddress() *ModificationScope1 {
	newValue := new(ModificationScope1)
	o.ModifiedPostalAddress = append(o.ModifiedPostalAddress, newValue)
	return newValue
}

func (o *Organisation17) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	o.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return o.PrimaryCommunicationAddress
}

func (o *Organisation17) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	o.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return o.SecondaryCommunicationAddress
}

func (o *Organisation17) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	o.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return o.AdditionalRegulatoryInformation
}

// Card acceptor performing the transaction.
type Organisation18 struct {

	// Identification of the card acceptor.
	Identification *GenericIdentification32 `xml:"Id"`

	// Name of the card acceptor as appearing on the receipt or the statement of account of the cardholder.
	// It correspond to the ISO 8583 field number 43.
	CommonName *Max70Text `xml:"CmonNm"`

	// Location of the card acceptor.
	// It correspond to the ISO 8583 field number 43.
	Location *CommunicationAddress5 `xml:"Lctn"`

	// Selected language of the card acceptor. Reference ISO 639-1 (alpha-2) andISO 639-2 (alpha-3).
	SelectedLanguage *LanguageCode `xml:"SelctdLang,omitempty"`

	// Additional card acceptor data required by a card scheme.
	SchemeData *Max140Text `xml:"SchmeData,omitempty"`
}

func (o *Organisation18) AddIdentification() *GenericIdentification32 {
	o.Identification = new(GenericIdentification32)
	return o.Identification
}

func (o *Organisation18) SetCommonName(value string) {
	o.CommonName = (*Max70Text)(&value)
}

func (o *Organisation18) AddLocation() *CommunicationAddress5 {
	o.Location = new(CommunicationAddress5)
	return o.Location
}

func (o *Organisation18) SetSelectedLanguage(value string) {
	o.SelectedLanguage = (*LanguageCode)(&value)
}

func (o *Organisation18) SetSchemeData(value string) {
	o.SchemeData = (*Max140Text)(&value)
}

// Card acceptor performing the transaction.
type Organisation19 struct {

	// Identification of the card acceptor.
	Identification *GenericIdentification32 `xml:"Id"`

	// Name of the card acceptor as appearing on the receipt or the statement of account of the cardholder.
	// It correspond to the ISO 8583, field number 43.
	CommonName *Max70Text `xml:"CmonNm"`

	// Selected language of the card acceptor. Reference ISO 639-1 (alpha-2) and ISO 639-2 (alpha-3).
	SelectedLanguage *LanguageCode `xml:"SelctdLang,omitempty"`

	// Additional card acceptor data required by a card scheme.
	SchemeData *Max140Text `xml:"SchmeData,omitempty"`
}

func (o *Organisation19) AddIdentification() *GenericIdentification32 {
	o.Identification = new(GenericIdentification32)
	return o.Identification
}

func (o *Organisation19) SetCommonName(value string) {
	o.CommonName = (*Max70Text)(&value)
}

func (o *Organisation19) SetSelectedLanguage(value string) {
	o.SelectedLanguage = (*LanguageCode)(&value)
}

func (o *Organisation19) SetSchemeData(value string) {
	o.SchemeData = (*Max140Text)(&value)
}

// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
type Organisation2 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification4Choice `xml:"Id,omitempty"`

	// Purpose of the organisation, eg, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country of taxation of an organisation.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb,omitempty"`

	// Number assigned by a national registration authority to an entity.
	NationalRegistrationNumber *Max35Text `xml:"NtlRegnNb,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress3 `xml:"PstlAdr"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`
}

func (o *Organisation2) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *Organisation2) AddIdentification() *PartyIdentification4Choice {
	o.Identification = new(PartyIdentification4Choice)
	return o.Identification
}

func (o *Organisation2) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation2) SetTaxationCountry(value string) {
	o.TaxationCountry = (*CountryCode)(&value)
}

func (o *Organisation2) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation2) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation2) SetTaxIdentificationNumber(value string) {
	o.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (o *Organisation2) SetNationalRegistrationNumber(value string) {
	o.NationalRegistrationNumber = (*Max35Text)(&value)
}

func (o *Organisation2) AddPostalAddress() *PostalAddress3 {
	newValue := new(PostalAddress3)
	o.PostalAddress = append(o.PostalAddress, newValue)
	return newValue
}

func (o *Organisation2) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	o.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return o.PrimaryCommunicationAddress
}

func (o *Organisation2) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	o.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return o.SecondaryCommunicationAddress
}

// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
type Organisation21 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Unique and unambiguous identifier for the organisation.
	Identification *PartyIdentification72Choice `xml:"Id,omitempty"`

	// Purpose of the organisation, for example,, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country of taxation of an organisation.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb,omitempty"`

	// Number assigned by a national registration authority to an entity.
	NationalRegistrationNumber *Max35Text `xml:"NtlRegnNb,omitempty"`

	// Postal address of a party.
	CorporateInvestorAddress *PostalAddress1 `xml:"CorpInvstrAdr"`
}

func (o *Organisation21) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *Organisation21) AddIdentification() *PartyIdentification72Choice {
	o.Identification = new(PartyIdentification72Choice)
	return o.Identification
}

func (o *Organisation21) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation21) SetTaxationCountry(value string) {
	o.TaxationCountry = (*CountryCode)(&value)
}

func (o *Organisation21) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation21) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation21) SetTaxIdentificationNumber(value string) {
	o.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (o *Organisation21) SetNationalRegistrationNumber(value string) {
	o.NationalRegistrationNumber = (*Max35Text)(&value)
}

func (o *Organisation21) AddCorporateInvestorAddress() *PostalAddress1 {
	o.CorporateInvestorAddress = new(PostalAddress1)
	return o.CorporateInvestorAddress
}

// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
type Organisation22 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Name of the organisation in short form.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Unique and unambiguous identifier for the organisation.
	Identification *PartyIdentification72Choice `xml:"Id,omitempty"`

	// Identification of the organisation with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Purpose of the organisation, for example, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Information related to an address to be inserted, updated or deleted.
	ModifiedPostalAddress []*ModificationScope34 `xml:"ModfdPstlAdr,omitempty"`

	// Type of organisation.
	TypeOfOrganisation *OrganisationType1Choice `xml:"TpOfOrg,omitempty"`
}

func (o *Organisation22) SetName(value string) {
	o.Name = (*Max350Text)(&value)
}

func (o *Organisation22) SetShortName(value string) {
	o.ShortName = (*Max35Text)(&value)
}

func (o *Organisation22) AddIdentification() *PartyIdentification72Choice {
	o.Identification = new(PartyIdentification72Choice)
	return o.Identification
}

func (o *Organisation22) SetLegalEntityIdentifier(value string) {
	o.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (o *Organisation22) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation22) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation22) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation22) AddModifiedPostalAddress() *ModificationScope34 {
	newValue := new(ModificationScope34)
	o.ModifiedPostalAddress = append(o.ModifiedPostalAddress, newValue)
	return newValue
}

func (o *Organisation22) AddTypeOfOrganisation() *OrganisationType1Choice {
	o.TypeOfOrganisation = new(OrganisationType1Choice)
	return o.TypeOfOrganisation
}

// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
type Organisation23 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Name of the organisation in short form.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress21 `xml:"PstlAdr"`
}

func (o *Organisation23) SetName(value string) {
	o.Name = (*Max350Text)(&value)
}

func (o *Organisation23) SetShortName(value string) {
	o.ShortName = (*Max35Text)(&value)
}

func (o *Organisation23) AddPostalAddress() *PostalAddress21 {
	newValue := new(PostalAddress21)
	o.PostalAddress = append(o.PostalAddress, newValue)
	return newValue
}

// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
type Organisation24 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Name of the organisation in short form.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Unique and unambiguous identifier for the organisation.
	Identification *PartyIdentification72Choice `xml:"Id,omitempty"`

	// Identification of the organisation with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Purpose of the organisation, for example, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress21 `xml:"PstlAdr"`

	// Type of organisation.
	TypeOfOrganisation *OrganisationType1Choice `xml:"TpOfOrg,omitempty"`
}

func (o *Organisation24) SetName(value string) {
	o.Name = (*Max350Text)(&value)
}

func (o *Organisation24) SetShortName(value string) {
	o.ShortName = (*Max35Text)(&value)
}

func (o *Organisation24) AddIdentification() *PartyIdentification72Choice {
	o.Identification = new(PartyIdentification72Choice)
	return o.Identification
}

func (o *Organisation24) SetLegalEntityIdentifier(value string) {
	o.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (o *Organisation24) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation24) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation24) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation24) AddPostalAddress() *PostalAddress21 {
	newValue := new(PostalAddress21)
	o.PostalAddress = append(o.PostalAddress, newValue)
	return newValue
}

func (o *Organisation24) AddTypeOfOrganisation() *OrganisationType1Choice {
	o.TypeOfOrganisation = new(OrganisationType1Choice)
	return o.TypeOfOrganisation
}

// Merchant performing the transaction.
type Organisation25 struct {

	// Identification of the merchant.
	Identification *GenericIdentification32 `xml:"Id,omitempty"`

	// Name of the merchant as appearing on the receipt.
	CommonName *Max70Text `xml:"CmonNm,omitempty"`

	// Location category of the place where the merchant actually performed the transaction.
	LocationCategory *LocationCategory1Code `xml:"LctnCtgy,omitempty"`

	// Location and contact information of the merchant performing the transaction.
	LocationAndContact *CommunicationAddress5 `xml:"LctnAndCtct,omitempty"`

	// Additional merchant data required by a card scheme.
	SchemeData *Max140Text `xml:"SchmeData,omitempty"`
}

func (o *Organisation25) AddIdentification() *GenericIdentification32 {
	o.Identification = new(GenericIdentification32)
	return o.Identification
}

func (o *Organisation25) SetCommonName(value string) {
	o.CommonName = (*Max70Text)(&value)
}

func (o *Organisation25) SetLocationCategory(value string) {
	o.LocationCategory = (*LocationCategory1Code)(&value)
}

func (o *Organisation25) AddLocationAndContact() *CommunicationAddress5 {
	o.LocationAndContact = new(CommunicationAddress5)
	return o.LocationAndContact
}

func (o *Organisation25) SetSchemeData(value string) {
	o.SchemeData = (*Max140Text)(&value)
}

// Merchant performing the transaction.
type Organisation26 struct {

	// Name of the merchant.
	CommonName *Max70Text `xml:"CmonNm"`

	// Location of the merchant.
	Address *Max140Text `xml:"Adr,omitempty"`

	// Country of the merchant.
	CountryCode *ISO3NumericCountryCode `xml:"CtryCd"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Identifier of the sponsored merchant assigned by the payment facilitator of their acquirer.
	RegisteredIdentifier *Max35Text `xml:"RegdIdr"`
}

func (o *Organisation26) SetCommonName(value string) {
	o.CommonName = (*Max70Text)(&value)
}

func (o *Organisation26) SetAddress(value string) {
	o.Address = (*Max140Text)(&value)
}

func (o *Organisation26) SetCountryCode(value string) {
	o.CountryCode = (*ISO3NumericCountryCode)(&value)
}

func (o *Organisation26) SetMerchantCategoryCode(value string) {
	o.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (o *Organisation26) SetRegisteredIdentifier(value string) {
	o.RegisteredIdentifier = (*Max35Text)(&value)
}

// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
type Organisation29 struct {

	// Name by which the organisation is known and which is usually used to identify that  organisation.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Name of the organisation in short form.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Unique and unambiguous identifier for the organisation.
	Identification *PartyIdentification72Choice `xml:"Id,omitempty"`

	// Identification of the organisation with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Purpose of the organisation, for example, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Information related to an address to be inserted, updated or deleted.
	ModifiedPostalAddress []*ModificationScope34 `xml:"ModfdPstlAdr,omitempty"`

	// Type of organisation.
	TypeOfOrganisation *OrganisationType1Choice `xml:"TpOfOrg,omitempty"`

	// Place of listing for shares in the organisation.
	PlaceOfListing []*MICIdentifier `xml:"PlcOfListg,omitempty"`
}

func (o *Organisation29) SetName(value string) {
	o.Name = (*Max350Text)(&value)
}

func (o *Organisation29) SetShortName(value string) {
	o.ShortName = (*Max35Text)(&value)
}

func (o *Organisation29) AddIdentification() *PartyIdentification72Choice {
	o.Identification = new(PartyIdentification72Choice)
	return o.Identification
}

func (o *Organisation29) SetLegalEntityIdentifier(value string) {
	o.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (o *Organisation29) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation29) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation29) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation29) AddModifiedPostalAddress() *ModificationScope34 {
	newValue := new(ModificationScope34)
	o.ModifiedPostalAddress = append(o.ModifiedPostalAddress, newValue)
	return newValue
}

func (o *Organisation29) AddTypeOfOrganisation() *OrganisationType1Choice {
	o.TypeOfOrganisation = new(OrganisationType1Choice)
	return o.TypeOfOrganisation
}

func (o *Organisation29) AddPlaceOfListing(value string) {
	o.PlaceOfListing = append(o.PlaceOfListing, (*MICIdentifier)(&value))
}

// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
type Organisation3 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification4Choice `xml:"Id,omitempty"`

	// Purpose of the organisation, eg, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country of taxation of an individual person or an organisation.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb,omitempty"`

	// Number assigned by a national registration authority to an entity.
	NationalRegistrationNumber *Max35Text `xml:"NtlRegnNb,omitempty"`

	// Address information to be either inserted, updated or deleted.
	ModifiedPostalAddress []*ModificationScope1 `xml:"ModfdPstlAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`
}

func (o *Organisation3) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *Organisation3) AddIdentification() *PartyIdentification4Choice {
	o.Identification = new(PartyIdentification4Choice)
	return o.Identification
}

func (o *Organisation3) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation3) SetTaxationCountry(value string) {
	o.TaxationCountry = (*CountryCode)(&value)
}

func (o *Organisation3) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation3) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation3) SetTaxIdentificationNumber(value string) {
	o.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (o *Organisation3) SetNationalRegistrationNumber(value string) {
	o.NationalRegistrationNumber = (*Max35Text)(&value)
}

func (o *Organisation3) AddModifiedPostalAddress() *ModificationScope1 {
	newValue := new(ModificationScope1)
	o.ModifiedPostalAddress = append(o.ModifiedPostalAddress, newValue)
	return newValue
}

func (o *Organisation3) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	o.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return o.PrimaryCommunicationAddress
}

func (o *Organisation3) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	o.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return o.SecondaryCommunicationAddress
}

// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
type Organisation30 struct {

	// Name by which the organisation is known and which is usually used to identify that  organisation.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Name of the organisation in short form.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Unique and unambiguous identifier for the organisation.
	Identification *PartyIdentification72Choice `xml:"Id,omitempty"`

	// Identification of the organisation with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Purpose of the organisation, for example, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress21 `xml:"PstlAdr,omitempty"`

	// Type of organisation.
	TypeOfOrganisation *OrganisationType1Choice `xml:"TpOfOrg,omitempty"`

	// Place of listing for shares in the organisation.
	PlaceOfListing []*MICIdentifier `xml:"PlcOfListg,omitempty"`
}

func (o *Organisation30) SetName(value string) {
	o.Name = (*Max350Text)(&value)
}

func (o *Organisation30) SetShortName(value string) {
	o.ShortName = (*Max35Text)(&value)
}

func (o *Organisation30) AddIdentification() *PartyIdentification72Choice {
	o.Identification = new(PartyIdentification72Choice)
	return o.Identification
}

func (o *Organisation30) SetLegalEntityIdentifier(value string) {
	o.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (o *Organisation30) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation30) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation30) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation30) AddPostalAddress() *PostalAddress21 {
	newValue := new(PostalAddress21)
	o.PostalAddress = append(o.PostalAddress, newValue)
	return newValue
}

func (o *Organisation30) AddTypeOfOrganisation() *OrganisationType1Choice {
	o.TypeOfOrganisation = new(OrganisationType1Choice)
	return o.TypeOfOrganisation
}

func (o *Organisation30) AddPlaceOfListing(value string) {
	o.PlaceOfListing = append(o.PlaceOfListing, (*MICIdentifier)(&value))
}

// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
type Organisation4 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification4Choice `xml:"Id,omitempty"`

	// Purpose of the organisation, eg, charity.
	Purpose *Max35Text `xml:"Purp,omitempty"`

	// Country of taxation of an organisation.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country in which the organisation is registered.
	RegistrationCountry *CountryCode `xml:"RegnCtry,omitempty"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb,omitempty"`

	// Number assigned by a national registration authority to an entity.
	NationalRegistrationNumber *Max35Text `xml:"NtlRegnNb,omitempty"`

	// Postal address of a party.
	CorporateInvestorAddress *PostalAddress1 `xml:"CorpInvstrAdr"`
}

func (o *Organisation4) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *Organisation4) AddIdentification() *PartyIdentification4Choice {
	o.Identification = new(PartyIdentification4Choice)
	return o.Identification
}

func (o *Organisation4) SetPurpose(value string) {
	o.Purpose = (*Max35Text)(&value)
}

func (o *Organisation4) SetTaxationCountry(value string) {
	o.TaxationCountry = (*CountryCode)(&value)
}

func (o *Organisation4) SetRegistrationCountry(value string) {
	o.RegistrationCountry = (*CountryCode)(&value)
}

func (o *Organisation4) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation4) SetTaxIdentificationNumber(value string) {
	o.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (o *Organisation4) SetNationalRegistrationNumber(value string) {
	o.NationalRegistrationNumber = (*Max35Text)(&value)
}

func (o *Organisation4) AddCorporateInvestorAddress() *PostalAddress1 {
	o.CorporateInvestorAddress = new(PostalAddress1)
	return o.CorporateInvestorAddress
}

// Merchant performing the transaction.
type Organisation5 struct {

	// Identification of the merchant.
	Identification *GenericIdentification32 `xml:"Id,omitempty"`

	// Name of the merchant as appearing on the receipt.
	CommonName *Max35Text `xml:"CmonNm,omitempty"`

	// Location category of the place where the merchant actually performed the transaction.
	LocationCategory *LocationCategory1Code `xml:"LctnCtgy,omitempty"`

	// Location of the merchant where the transaction took place, as appearing on the receipt.
	Address *Max70Text `xml:"Adr,omitempty"`

	// Country of the merchant where the transaction took place.
	CountryCode *ISO3ACountryCode `xml:"CtryCd,omitempty"`

	// Additional merchant data required by a card scheme.
	SchemeData *Max140Text `xml:"SchmeData,omitempty"`
}

func (o *Organisation5) AddIdentification() *GenericIdentification32 {
	o.Identification = new(GenericIdentification32)
	return o.Identification
}

func (o *Organisation5) SetCommonName(value string) {
	o.CommonName = (*Max35Text)(&value)
}

func (o *Organisation5) SetLocationCategory(value string) {
	o.LocationCategory = (*LocationCategory1Code)(&value)
}

func (o *Organisation5) SetAddress(value string) {
	o.Address = (*Max70Text)(&value)
}

func (o *Organisation5) SetCountryCode(value string) {
	o.CountryCode = (*ISO3ACountryCode)(&value)
}

func (o *Organisation5) SetSchemeData(value string) {
	o.SchemeData = (*Max140Text)(&value)
}

// Information which describes the organisation.
type Organisation6 struct {

	// Name by which a party is known and which is usually used to identify that party.
	FullLegalName *Max350Text `xml:"FullLglNm"`

	// Name used by a business for commercial purposes, although its registered legal name, used for contracts and other formal situations, may be another.
	TradingName *Max350Text `xml:"TradgNm,omitempty"`

	// Country in which the organisation has its business activity.
	CountryOfOperation *CountryCode `xml:"CtryOfOpr"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`

	// Is an operational address, for example, of a shared services center.
	OperationalAddress *PostalAddress6 `xml:"OprlAdr,omitempty"`

	// Is the address where the business activity is taking place.
	BusinessAddress *PostalAddress6 `xml:"BizAdr,omitempty"`

	// Is the address where the entity resides and is registered. More generically, it is the home address (Residential address).
	LegalAddress *PostalAddress6 `xml:"LglAdr"`

	// Unique and unambiguous way of identifying an organisation.
	OrganisationIdentification *OrganisationIdentification6 `xml:"OrgId"`

	// Person in the customer's organisation who can be contacted by the account servicer.
	RepresentativeOfficer []*PartyIdentification40 `xml:"RprtvOffcr,omitempty"`

	// Identification of the person responsible of the treasury department within an organisation.
	TreasuryManager *PartyIdentification40 `xml:"TrsrMgr,omitempty"`

	// Is the main mandate holder that will delegate some authority to other individuals (mandate holders) to perform some specific bank operations on the account.
	MainMandateHolder []*PartyIdentification40 `xml:"MainMndtHldr,omitempty"`

	// Potential sender of a message related to the life cyle of an account.
	Sender []*PartyIdentification40 `xml:"Sndr,omitempty"`
}

func (o *Organisation6) SetFullLegalName(value string) {
	o.FullLegalName = (*Max350Text)(&value)
}

func (o *Organisation6) SetTradingName(value string) {
	o.TradingName = (*Max350Text)(&value)
}

func (o *Organisation6) SetCountryOfOperation(value string) {
	o.CountryOfOperation = (*CountryCode)(&value)
}

func (o *Organisation6) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation6) AddOperationalAddress() *PostalAddress6 {
	o.OperationalAddress = new(PostalAddress6)
	return o.OperationalAddress
}

func (o *Organisation6) AddBusinessAddress() *PostalAddress6 {
	o.BusinessAddress = new(PostalAddress6)
	return o.BusinessAddress
}

func (o *Organisation6) AddLegalAddress() *PostalAddress6 {
	o.LegalAddress = new(PostalAddress6)
	return o.LegalAddress
}

func (o *Organisation6) AddOrganisationIdentification() *OrganisationIdentification6 {
	o.OrganisationIdentification = new(OrganisationIdentification6)
	return o.OrganisationIdentification
}

func (o *Organisation6) AddRepresentativeOfficer() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.RepresentativeOfficer = append(o.RepresentativeOfficer, newValue)
	return newValue
}

func (o *Organisation6) AddTreasuryManager() *PartyIdentification40 {
	o.TreasuryManager = new(PartyIdentification40)
	return o.TreasuryManager
}

func (o *Organisation6) AddMainMandateHolder() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.MainMandateHolder = append(o.MainMandateHolder, newValue)
	return newValue
}

func (o *Organisation6) AddSender() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.Sender = append(o.Sender, newValue)
	return newValue
}

// Information which describes the organisation.
type Organisation7 struct {

	// Name by which a party is known and which is usually used to identify that party.
	FullLegalName *Max350Text `xml:"FullLglNm"`

	// Name used by a business for commercial purposes, although its registered legal name, used for contracts and other formal situations, may be another.
	TradingName *Max350Text `xml:"TradgNm,omitempty"`

	// Country in which the organisation has its business activity.
	CountryOfOperation *CountryCode `xml:"CtryOfOpr"`

	// Date and time at which a given organisation was officially registered.
	RegistrationDate *ISODate `xml:"RegnDt"`

	// Is an operational address, for example, of a shared services center.
	OperationalAddress *PostalAddress6 `xml:"OprlAdr,omitempty"`

	// Is the address where the business activity is taking place.
	BusinessAddress *PostalAddress6 `xml:"BizAdr,omitempty"`

	// Is the address where the entity resides and is registered. More generically, it is the home address (Residential address).
	LegalAddress *PostalAddress6 `xml:"LglAdr"`

	// Unique and unambiguous way of identifying an organisation.
	OrganisationIdentification *OrganisationIdentification6 `xml:"OrgId"`

	// Person in the customer's organisation who can be contacted by the account servicer.
	RepresentativeOfficer []*PartyIdentification40 `xml:"RprtvOffcr,omitempty"`

	// Identification of the person responsible of the treasury department within an organisation.
	TreasuryManager *PartyIdentification40 `xml:"TrsrMgr,omitempty"`

	// Is the main mandate holder that will delegate some authority to other individuals (mandate holders) to perform some specific bank operations on the account.
	MainMandateHolder []*PartyIdentification40 `xml:"MainMndtHldr,omitempty"`

	// Potential sender of a message related to the life cyle of an account.
	Sender []*PartyIdentification40 `xml:"Sndr,omitempty"`
}

func (o *Organisation7) SetFullLegalName(value string) {
	o.FullLegalName = (*Max350Text)(&value)
}

func (o *Organisation7) SetTradingName(value string) {
	o.TradingName = (*Max350Text)(&value)
}

func (o *Organisation7) SetCountryOfOperation(value string) {
	o.CountryOfOperation = (*CountryCode)(&value)
}

func (o *Organisation7) SetRegistrationDate(value string) {
	o.RegistrationDate = (*ISODate)(&value)
}

func (o *Organisation7) AddOperationalAddress() *PostalAddress6 {
	o.OperationalAddress = new(PostalAddress6)
	return o.OperationalAddress
}

func (o *Organisation7) AddBusinessAddress() *PostalAddress6 {
	o.BusinessAddress = new(PostalAddress6)
	return o.BusinessAddress
}

func (o *Organisation7) AddLegalAddress() *PostalAddress6 {
	o.LegalAddress = new(PostalAddress6)
	return o.LegalAddress
}

func (o *Organisation7) AddOrganisationIdentification() *OrganisationIdentification6 {
	o.OrganisationIdentification = new(OrganisationIdentification6)
	return o.OrganisationIdentification
}

func (o *Organisation7) AddRepresentativeOfficer() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.RepresentativeOfficer = append(o.RepresentativeOfficer, newValue)
	return newValue
}

func (o *Organisation7) AddTreasuryManager() *PartyIdentification40 {
	o.TreasuryManager = new(PartyIdentification40)
	return o.TreasuryManager
}

func (o *Organisation7) AddMainMandateHolder() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.MainMandateHolder = append(o.MainMandateHolder, newValue)
	return newValue
}

func (o *Organisation7) AddSender() *PartyIdentification40 {
	newValue := new(PartyIdentification40)
	o.Sender = append(o.Sender, newValue)
	return newValue
}

// Merchant performing the transaction.
type Organisation8 struct {

	// Identification of the merchant.
	Identification *GenericIdentification32 `xml:"Id,omitempty"`

	// Name of the merchant as appearing on the receipt.
	CommonName *Max70Text `xml:"CmonNm,omitempty"`

	// Location category of the place where the merchant actually performed the transaction.
	LocationCategory *LocationCategory1Code `xml:"LctnCtgy,omitempty"`

	// Location of the merchant where the transaction took place, as appearing on the receipt.
	Address *Max140Text `xml:"Adr,omitempty"`

	// Country of the merchant where the transaction took place.
	CountryCode *ISO3NumericCountryCode `xml:"CtryCd,omitempty"`

	// Additional merchant data required by a card scheme.
	SchemeData *Max140Text `xml:"SchmeData,omitempty"`
}

func (o *Organisation8) AddIdentification() *GenericIdentification32 {
	o.Identification = new(GenericIdentification32)
	return o.Identification
}

func (o *Organisation8) SetCommonName(value string) {
	o.CommonName = (*Max70Text)(&value)
}

func (o *Organisation8) SetLocationCategory(value string) {
	o.LocationCategory = (*LocationCategory1Code)(&value)
}

func (o *Organisation8) SetAddress(value string) {
	o.Address = (*Max140Text)(&value)
}

func (o *Organisation8) SetCountryCode(value string) {
	o.CountryCode = (*ISO3NumericCountryCode)(&value)
}

func (o *Organisation8) SetSchemeData(value string) {
	o.SchemeData = (*Max140Text)(&value)
}

// Merchant performing the transaction.
type Organisation9 struct {

	// Identification of the merchant.
	Identification *GenericIdentification32 `xml:"Id"`

	// Name of the merchant as appearing on the receipt.
	CommonName *Max70Text `xml:"CmonNm,omitempty"`

	// Location category of the place where the merchant actually performed the transaction.
	LocationCategory *LocationCategory1Code `xml:"LctnCtgy,omitempty"`

	// Location of the merchant where the transaction took place, as appearing on the receipt.
	Address *Max140Text `xml:"Adr,omitempty"`

	// Country of the merchant where the transaction took place.
	CountryCode *ISO3NumericCountryCode `xml:"CtryCd,omitempty"`

	// Additional merchant data required by a card scheme.
	SchemeData *Max140Text `xml:"SchmeData,omitempty"`
}

func (o *Organisation9) AddIdentification() *GenericIdentification32 {
	o.Identification = new(GenericIdentification32)
	return o.Identification
}

func (o *Organisation9) SetCommonName(value string) {
	o.CommonName = (*Max70Text)(&value)
}

func (o *Organisation9) SetLocationCategory(value string) {
	o.LocationCategory = (*LocationCategory1Code)(&value)
}

func (o *Organisation9) SetAddress(value string) {
	o.Address = (*Max140Text)(&value)
}

func (o *Organisation9) SetCountryCode(value string) {
	o.CountryCode = (*ISO3NumericCountryCode)(&value)
}

func (o *Organisation9) SetSchemeData(value string) {
	o.SchemeData = (*Max140Text)(&value)
}
