package iso20022

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson10 struct {

	// Specifies the terms used to formally address a person.
	NamePrefix *NamePrefix1Code `xml:"NmPrfx,omitempty"`

	// Specifies the terms used to formally address a person.
	ExtendedNamePrefix *Extended350Code `xml:"XtndedNmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm"`

	// Second name of a person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Additional information about a person that follows a person's name, eg, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Specifies the gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Language in which a person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt"`

	// Country where a person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where a person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where a person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of a person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Country of taxation of an individual person.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country and residential status of an individual, for example, non-pernament resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress3 `xml:"PstlAdr"`

	// Nationality and legal status (minor or major) or rights that an individual may possess.
	Citizenship []*CitizenshipInformation `xml:"Ctznsh"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification9 `xml:"OthrId,omitempty"`
}

func (i *IndividualPerson10) SetNamePrefix(value string) {
	i.NamePrefix = (*NamePrefix1Code)(&value)
}

func (i *IndividualPerson10) SetExtendedNamePrefix(value string) {
	i.ExtendedNamePrefix = (*Extended350Code)(&value)
}

func (i *IndividualPerson10) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson10) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson10) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson10) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson10) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson10) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *IndividualPerson10) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson10) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson10) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson10) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson10) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson10) SetTaxationCountry(value string) {
	i.TaxationCountry = (*CountryCode)(&value)
}

func (i *IndividualPerson10) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson10) AddPostalAddress() *PostalAddress3 {
	newValue := new(PostalAddress3)
	i.PostalAddress = append(i.PostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson10) AddCitizenship() *CitizenshipInformation {
	newValue := new(CitizenshipInformation)
	i.Citizenship = append(i.Citizenship, newValue)
	return newValue
}

func (i *IndividualPerson10) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson10) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson10) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	i.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return i.PrimaryCommunicationAddress
}

func (i *IndividualPerson10) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	i.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return i.SecondaryCommunicationAddress
}

func (i *IndividualPerson10) AddOtherIdentification() *GenericIdentification9 {
	newValue := new(GenericIdentification9)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson11 struct {

	// Specifies the terms used to formally address a person.
	NamePrefix *NamePrefix1Code `xml:"NmPrfx,omitempty"`

	// Specifies the terms used to formally address a person.
	ExtendedNamePrefix *Extended350Code `xml:"XtndedNmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm"`

	// Second name of a person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Additional information about a person that follows a person's name, eg, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Specifies the gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Language in which a person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt"`

	// Country where a person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where a person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where a person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of a person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Country of taxation of an individual person or an organisation.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country and residential status of an individual, for example, non-pernament resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Address information to be either inserted, updated or deleted.
	ModifiedPostalAddress []*ModificationScope1 `xml:"ModfdPstlAdr,omitempty"`

	// Citizenship information to be inserted or deleted.
	ModifiedCitizenship []*ModificationScope3 `xml:"ModfdCtznsh,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Identification information to be either inserted or deleted.
	ModifiedOtherIdentification []*ModificationScope2 `xml:"ModfdOthrId,omitempty"`
}

func (i *IndividualPerson11) SetNamePrefix(value string) {
	i.NamePrefix = (*NamePrefix1Code)(&value)
}

func (i *IndividualPerson11) SetExtendedNamePrefix(value string) {
	i.ExtendedNamePrefix = (*Extended350Code)(&value)
}

func (i *IndividualPerson11) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson11) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson11) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *IndividualPerson11) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson11) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson11) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetTaxationCountry(value string) {
	i.TaxationCountry = (*CountryCode)(&value)
}

func (i *IndividualPerson11) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson11) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson11) AddModifiedPostalAddress() *ModificationScope1 {
	newValue := new(ModificationScope1)
	i.ModifiedPostalAddress = append(i.ModifiedPostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson11) AddModifiedCitizenship() *ModificationScope3 {
	newValue := new(ModificationScope3)
	i.ModifiedCitizenship = append(i.ModifiedCitizenship, newValue)
	return newValue
}

func (i *IndividualPerson11) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	i.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return i.PrimaryCommunicationAddress
}

func (i *IndividualPerson11) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	i.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return i.SecondaryCommunicationAddress
}

func (i *IndividualPerson11) AddModifiedOtherIdentification() *ModificationScope2 {
	newValue := new(ModificationScope2)
	i.ModifiedOtherIdentification = append(i.ModifiedOtherIdentification, newValue)
	return newValue
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson12 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Country and residential status of an individual, for example, non-pernament resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification11 `xml:"OthrId,omitempty"`
}

func (i *IndividualPerson12) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson12) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson12) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson12) AddOtherIdentification() *GenericIdentification11 {
	newValue := new(GenericIdentification11)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson13 struct {

	// Name received at birth, eg, maiden name.
	BirthName *Max35Text `xml:"BirthNm"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Unique and unambiguous identification of a person, eg passport.
	Identification *PersonIdentification2 `xml:"Id,omitempty"`

	// Postal address of a party.
	Address *LongPostalAddress2Choice `xml:"Adr,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingParty *PartyIdentification9Choice `xml:"EmplngPty,omitempty"`

	// Specifies details related to the attendance card.
	AttendanceCardDetails *AttendanceCard1 `xml:"AttndncCardDtls"`
}

func (i *IndividualPerson13) SetBirthName(value string) {
	i.BirthName = (*Max35Text)(&value)
}

func (i *IndividualPerson13) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson13) AddIdentification() *PersonIdentification2 {
	i.Identification = new(PersonIdentification2)
	return i.Identification
}

func (i *IndividualPerson13) AddAddress() *LongPostalAddress2Choice {
	i.Address = new(LongPostalAddress2Choice)
	return i.Address
}

func (i *IndividualPerson13) AddEmployingParty() *PartyIdentification9Choice {
	i.EmployingParty = new(PartyIdentification9Choice)
	return i.EmployingParty
}

func (i *IndividualPerson13) AddAttendanceCardDetails() *AttendanceCard1 {
	i.AttendanceCardDetails = new(AttendanceCard1)
	return i.AttendanceCardDetails
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson14 struct {

	// Name received at birth, eg, maiden name.
	BirthName *Max35Text `xml:"BirthNm"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Unique and unambiguous identification of a person, eg passport.
	Identification *PersonIdentification2 `xml:"Id,omitempty"`

	// Postal address of a party.
	Address *PostalAddress1 `xml:"Adr,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingParty *PartyIdentification9Choice `xml:"EmplngPty,omitempty"`
}

func (i *IndividualPerson14) SetBirthName(value string) {
	i.BirthName = (*Max35Text)(&value)
}

func (i *IndividualPerson14) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson14) AddIdentification() *PersonIdentification2 {
	i.Identification = new(PersonIdentification2)
	return i.Identification
}

func (i *IndividualPerson14) AddAddress() *PostalAddress1 {
	i.Address = new(PostalAddress1)
	return i.Address
}

func (i *IndividualPerson14) AddEmployingParty() *PartyIdentification9Choice {
	i.EmployingParty = new(PartyIdentification9Choice)
	return i.EmployingParty
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson16 struct {

	// Name received at birth, eg, maiden name.
	BirthName *Max35Text `xml:"BirthNm"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Unique and unambiguous identification of a person, eg, passport.
	Identification *PersonIdentification6 `xml:"Id,omitempty"`

	// Postal address of a party.
	Address *LongPostalAddress2Choice `xml:"Adr,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingParty *PartyIdentification9Choice `xml:"EmplngPty,omitempty"`
}

func (i *IndividualPerson16) SetBirthName(value string) {
	i.BirthName = (*Max35Text)(&value)
}

func (i *IndividualPerson16) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson16) AddIdentification() *PersonIdentification6 {
	i.Identification = new(PersonIdentification6)
	return i.Identification
}

func (i *IndividualPerson16) AddAddress() *LongPostalAddress2Choice {
	i.Address = new(LongPostalAddress2Choice)
	return i.Address
}

func (i *IndividualPerson16) AddEmployingParty() *PartyIdentification9Choice {
	i.EmployingParty = new(PartyIdentification9Choice)
	return i.EmployingParty
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson17 struct {

	// Name received at birth, eg, maiden name.
	BirthName *Max35Text `xml:"BirthNm"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Unique and unambiguous identification of a person, eg, passport.
	Identification *PersonIdentification6 `xml:"Id,omitempty"`

	// Postal address of a party.
	Address *LongPostalAddress2Choice `xml:"Adr,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingParty *PartyIdentification9Choice `xml:"EmplngPty,omitempty"`

	// Specifies details related to the attendance card.
	AttendanceCardDetails *AttendanceCard2 `xml:"AttndncCardDtls"`
}

func (i *IndividualPerson17) SetBirthName(value string) {
	i.BirthName = (*Max35Text)(&value)
}

func (i *IndividualPerson17) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson17) AddIdentification() *PersonIdentification6 {
	i.Identification = new(PersonIdentification6)
	return i.Identification
}

func (i *IndividualPerson17) AddAddress() *LongPostalAddress2Choice {
	i.Address = new(LongPostalAddress2Choice)
	return i.Address
}

func (i *IndividualPerson17) AddEmployingParty() *PartyIdentification9Choice {
	i.EmployingParty = new(PartyIdentification9Choice)
	return i.EmployingParty
}

func (i *IndividualPerson17) AddAttendanceCardDetails() *AttendanceCard2 {
	i.AttendanceCardDetails = new(AttendanceCard2)
	return i.AttendanceCardDetails
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson2 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Country and residential status of an individual, for example, non-pernament resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification8 `xml:"OthrId,omitempty"`
}

func (i *IndividualPerson2) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson2) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson2) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson2) AddOtherIdentification() *GenericIdentification8 {
	newValue := new(GenericIdentification8)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson20 struct {

	// Term used to address a person.
	NamePrefix *NamePrefix1Choice `xml:"NmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm"`

	// Second name of a person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Additional information about a person that follows a person's name, eg, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Specifies the gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Language in which a person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt"`

	// Country where a person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where a person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where a person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of a person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Country of taxation of an individual person.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country and residential status of an individual, for example, non-pernament resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress3 `xml:"PstlAdr"`

	// Nationality and legal status (minor or major) or rights that an individual may possess.
	Citizenship []*CitizenshipInformation `xml:"Ctznsh"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification44 `xml:"OthrId,omitempty"`

	// Additional regulatory information about the investor that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`
}

func (i *IndividualPerson20) AddNamePrefix() *NamePrefix1Choice {
	i.NamePrefix = new(NamePrefix1Choice)
	return i.NamePrefix
}

func (i *IndividualPerson20) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson20) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson20) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson20) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson20) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson20) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *IndividualPerson20) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson20) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson20) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson20) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson20) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson20) SetTaxationCountry(value string) {
	i.TaxationCountry = (*CountryCode)(&value)
}

func (i *IndividualPerson20) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson20) AddPostalAddress() *PostalAddress3 {
	newValue := new(PostalAddress3)
	i.PostalAddress = append(i.PostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson20) AddCitizenship() *CitizenshipInformation {
	newValue := new(CitizenshipInformation)
	i.Citizenship = append(i.Citizenship, newValue)
	return newValue
}

func (i *IndividualPerson20) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson20) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson20) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	i.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return i.PrimaryCommunicationAddress
}

func (i *IndividualPerson20) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	i.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return i.SecondaryCommunicationAddress
}

func (i *IndividualPerson20) AddOtherIdentification() *GenericIdentification44 {
	newValue := new(GenericIdentification44)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

func (i *IndividualPerson20) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	i.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return i.AdditionalRegulatoryInformation
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson21 struct {

	// Term used to address a person.
	NamePrefix *NamePrefix1Choice `xml:"NmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm"`

	// Second name of a person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Additional information about a person that follows a person's name, eg, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Specifies the gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Language in which a person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt"`

	// Country where a person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where a person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where a person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of a person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Country of taxation of an individual person.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country and residential status of an individual, for example, non-pernament resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Address information to be either inserted, updated or deleted.
	ModifiedPostalAddress []*ModificationScope1 `xml:"ModfdPstlAdr,omitempty"`

	// Citizenship information to be inserted or deleted.
	ModifiedCitizenship []*ModificationScope3 `xml:"ModfdCtznsh,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Identification information to be either inserted or deleted.
	ModifiedOtherIdentification []*ModificationScope17 `xml:"ModfdOthrId,omitempty"`

	// Additional regulatory information about the investor that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`

	// Specifies if due diligence checks on the political exposure of the investor have been carried out and whether these checks are national or foreign. (A politically exposed person is someone who has been entrusted with a prominent public function, or an individual who is closely related to such a person.)
	PoliticallyExposedPersonType *PoliticalExposureType1Choice `xml:"PltclyXpsdPrsnTp,omitempty"`
}

func (i *IndividualPerson21) AddNamePrefix() *NamePrefix1Choice {
	i.NamePrefix = new(NamePrefix1Choice)
	return i.NamePrefix
}

func (i *IndividualPerson21) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson21) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson21) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson21) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson21) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson21) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *IndividualPerson21) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson21) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson21) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson21) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson21) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson21) SetTaxationCountry(value string) {
	i.TaxationCountry = (*CountryCode)(&value)
}

func (i *IndividualPerson21) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson21) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson21) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson21) AddModifiedPostalAddress() *ModificationScope1 {
	newValue := new(ModificationScope1)
	i.ModifiedPostalAddress = append(i.ModifiedPostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson21) AddModifiedCitizenship() *ModificationScope3 {
	newValue := new(ModificationScope3)
	i.ModifiedCitizenship = append(i.ModifiedCitizenship, newValue)
	return newValue
}

func (i *IndividualPerson21) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	i.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return i.PrimaryCommunicationAddress
}

func (i *IndividualPerson21) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	i.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return i.SecondaryCommunicationAddress
}

func (i *IndividualPerson21) AddModifiedOtherIdentification() *ModificationScope17 {
	newValue := new(ModificationScope17)
	i.ModifiedOtherIdentification = append(i.ModifiedOtherIdentification, newValue)
	return newValue
}

func (i *IndividualPerson21) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	i.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return i.AdditionalRegulatoryInformation
}

func (i *IndividualPerson21) AddPoliticallyExposedPersonType() *PoliticalExposureType1Choice {
	i.PoliticallyExposedPersonType = new(PoliticalExposureType1Choice)
	return i.PoliticallyExposedPersonType
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson22 struct {

	// Term used to address a person.
	NamePrefix *NamePrefix1Choice `xml:"NmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm"`

	// Second name of a person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Additional information about a person that follows a person's name, eg, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Specifies the gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Language in which a person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt"`

	// Country where a person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where a person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where a person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of a person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Country of taxation of an individual person.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country and residential status of an individual, for example, non-pernament resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress3 `xml:"PstlAdr"`

	// Nationality and legal status (minor or major) or rights that an individual may possess.
	Citizenship []*CitizenshipInformation `xml:"Ctznsh"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification44 `xml:"OthrId,omitempty"`

	// Additional regulatory information about the investor that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`

	// Specifies if due diligence checks on the political exposure of the investor have been carried out and whether these checks are national or foreign. (A politically exposed person is someone who has been entrusted with a prominent public function, or an individual who is closely related to such a person.)
	PoliticallyExposedPersonType *PoliticalExposureType1Choice `xml:"PltclyXpsdPrsnTp,omitempty"`
}

func (i *IndividualPerson22) AddNamePrefix() *NamePrefix1Choice {
	i.NamePrefix = new(NamePrefix1Choice)
	return i.NamePrefix
}

func (i *IndividualPerson22) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson22) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson22) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson22) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson22) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson22) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *IndividualPerson22) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson22) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson22) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson22) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson22) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson22) SetTaxationCountry(value string) {
	i.TaxationCountry = (*CountryCode)(&value)
}

func (i *IndividualPerson22) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson22) AddPostalAddress() *PostalAddress3 {
	newValue := new(PostalAddress3)
	i.PostalAddress = append(i.PostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson22) AddCitizenship() *CitizenshipInformation {
	newValue := new(CitizenshipInformation)
	i.Citizenship = append(i.Citizenship, newValue)
	return newValue
}

func (i *IndividualPerson22) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson22) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson22) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	i.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return i.PrimaryCommunicationAddress
}

func (i *IndividualPerson22) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	i.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return i.SecondaryCommunicationAddress
}

func (i *IndividualPerson22) AddOtherIdentification() *GenericIdentification44 {
	newValue := new(GenericIdentification44)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

func (i *IndividualPerson22) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	i.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return i.AdditionalRegulatoryInformation
}

func (i *IndividualPerson22) AddPoliticallyExposedPersonType() *PoliticalExposureType1Choice {
	i.PoliticallyExposedPersonType = new(PoliticalExposureType1Choice)
	return i.PoliticallyExposedPersonType
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson23 struct {

	// Term used to address a person.
	NamePrefix *NamePrefix1Choice `xml:"NmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm"`

	// Second name of a person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Additional information about a person that follows a person's name, for example, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Specifies the gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Language in which a person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt"`

	// Country where a person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where a person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where a person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of a person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Country of taxation of an individual person.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country and residential status of an individual, for example, non-permanent resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress []*PostalAddress3 `xml:"PstlAdr"`

	// Nationality and legal status (minor or major) or rights that an individual may possess.
	Citizenship []*CitizenshipInformation `xml:"Ctznsh"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification55 `xml:"OthrId,omitempty"`

	// Additional regulatory information about the investor that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`

	// Specifies if due diligence checks on the political exposure of the investor have been carried out and whether these checks are national or foreign. (A politically exposed person is someone who has been entrusted with a prominent public function, or an individual who is closely related to such a person.)
	PoliticallyExposedPersonType *PoliticalExposureType1Choice `xml:"PltclyXpsdPrsnTp,omitempty"`
}

func (i *IndividualPerson23) AddNamePrefix() *NamePrefix1Choice {
	i.NamePrefix = new(NamePrefix1Choice)
	return i.NamePrefix
}

func (i *IndividualPerson23) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson23) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson23) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson23) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson23) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson23) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *IndividualPerson23) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson23) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson23) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson23) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson23) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson23) SetTaxationCountry(value string) {
	i.TaxationCountry = (*CountryCode)(&value)
}

func (i *IndividualPerson23) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson23) AddPostalAddress() *PostalAddress3 {
	newValue := new(PostalAddress3)
	i.PostalAddress = append(i.PostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson23) AddCitizenship() *CitizenshipInformation {
	newValue := new(CitizenshipInformation)
	i.Citizenship = append(i.Citizenship, newValue)
	return newValue
}

func (i *IndividualPerson23) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson23) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson23) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	i.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return i.PrimaryCommunicationAddress
}

func (i *IndividualPerson23) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	i.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return i.SecondaryCommunicationAddress
}

func (i *IndividualPerson23) AddOtherIdentification() *GenericIdentification55 {
	newValue := new(GenericIdentification55)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

func (i *IndividualPerson23) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	i.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return i.AdditionalRegulatoryInformation
}

func (i *IndividualPerson23) AddPoliticallyExposedPersonType() *PoliticalExposureType1Choice {
	i.PoliticallyExposedPersonType = new(PoliticalExposureType1Choice)
	return i.PoliticallyExposedPersonType
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson24 struct {

	// Term used to address a person.
	NamePrefix *NamePrefix1Choice `xml:"NmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm"`

	// Second name of a person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Additional information about a person that follows a person's name, for example, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Specifies the gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Language in which a person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt"`

	// Country where a person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where a person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where a person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of a person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Country of taxation of an individual person.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country and residential status of an individual, for example, non-permanent resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Address information to be either inserted, updated or deleted.
	ModifiedPostalAddress []*ModificationScope1 `xml:"ModfdPstlAdr,omitempty"`

	// Citizenship information to be inserted or deleted.
	ModifiedCitizenship []*ModificationScope3 `xml:"ModfdCtznsh,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Identification information to be either inserted or deleted.
	ModifiedOtherIdentification []*ModificationScope23 `xml:"ModfdOthrId,omitempty"`

	// Additional regulatory information about the investor that is required in some markets to support anti-money laundering laws.
	AdditionalRegulatoryInformation *RegulatoryInformation1 `xml:"AddtlRgltryInf,omitempty"`

	// Specifies if due diligence checks on the political exposure of the investor have been carried out and whether these checks are national or foreign. (A politically exposed person is someone who has been entrusted with a prominent public function, or an individual who is closely related to such a person.)
	PoliticallyExposedPersonType *PoliticalExposureType1Choice `xml:"PltclyXpsdPrsnTp,omitempty"`
}

func (i *IndividualPerson24) AddNamePrefix() *NamePrefix1Choice {
	i.NamePrefix = new(NamePrefix1Choice)
	return i.NamePrefix
}

func (i *IndividualPerson24) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson24) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson24) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson24) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson24) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson24) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *IndividualPerson24) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson24) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson24) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson24) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson24) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson24) SetTaxationCountry(value string) {
	i.TaxationCountry = (*CountryCode)(&value)
}

func (i *IndividualPerson24) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson24) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson24) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson24) AddModifiedPostalAddress() *ModificationScope1 {
	newValue := new(ModificationScope1)
	i.ModifiedPostalAddress = append(i.ModifiedPostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson24) AddModifiedCitizenship() *ModificationScope3 {
	newValue := new(ModificationScope3)
	i.ModifiedCitizenship = append(i.ModifiedCitizenship, newValue)
	return newValue
}

func (i *IndividualPerson24) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	i.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return i.PrimaryCommunicationAddress
}

func (i *IndividualPerson24) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	i.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return i.SecondaryCommunicationAddress
}

func (i *IndividualPerson24) AddModifiedOtherIdentification() *ModificationScope23 {
	newValue := new(ModificationScope23)
	i.ModifiedOtherIdentification = append(i.ModifiedOtherIdentification, newValue)
	return newValue
}

func (i *IndividualPerson24) AddAdditionalRegulatoryInformation() *RegulatoryInformation1 {
	i.AdditionalRegulatoryInformation = new(RegulatoryInformation1)
	return i.AdditionalRegulatoryInformation
}

func (i *IndividualPerson24) AddPoliticallyExposedPersonType() *PoliticalExposureType1Choice {
	i.PoliticallyExposedPersonType = new(PoliticalExposureType1Choice)
	return i.PoliticallyExposedPersonType
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson25 struct {

	// Name received at birth, for example, maiden name.
	BirthName *Max35Text `xml:"BirthNm"`

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Unique and unambiguous identification of the person, for example, passport.
	Identification *PersonIdentification6 `xml:"Id,omitempty"`

	// Postal address of the party.
	Address *LongPostalAddress2Choice `xml:"Adr,omitempty"`

	// Organisation represented by the person, or for which the person works.
	EmployingParty *PartyIdentification40Choice `xml:"EmplngPty,omitempty"`
}

func (i *IndividualPerson25) SetBirthName(value string) {
	i.BirthName = (*Max35Text)(&value)
}

func (i *IndividualPerson25) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson25) AddIdentification() *PersonIdentification6 {
	i.Identification = new(PersonIdentification6)
	return i.Identification
}

func (i *IndividualPerson25) AddAddress() *LongPostalAddress2Choice {
	i.Address = new(LongPostalAddress2Choice)
	return i.Address
}

func (i *IndividualPerson25) AddEmployingParty() *PartyIdentification40Choice {
	i.EmployingParty = new(PartyIdentification40Choice)
	return i.EmployingParty
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson26 struct {

	// Name received at birth, for example, maiden name.
	BirthName *Max35Text `xml:"BirthNm"`

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Unique and unambiguous identification of the person, for example, passport.
	Identification *PersonIdentification6 `xml:"Id,omitempty"`

	// Postal address of the party.
	Address *LongPostalAddress2Choice `xml:"Adr,omitempty"`

	// Organisation represented by the person, or for which the person works.
	EmployingParty *PartyIdentification40Choice `xml:"EmplngPty,omitempty"`

	// Specifies details related to the attendance card.
	AttendanceCardDetails *AttendanceCard2 `xml:"AttndncCardDtls"`
}

func (i *IndividualPerson26) SetBirthName(value string) {
	i.BirthName = (*Max35Text)(&value)
}

func (i *IndividualPerson26) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson26) AddIdentification() *PersonIdentification6 {
	i.Identification = new(PersonIdentification6)
	return i.Identification
}

func (i *IndividualPerson26) AddAddress() *LongPostalAddress2Choice {
	i.Address = new(LongPostalAddress2Choice)
	return i.Address
}

func (i *IndividualPerson26) AddEmployingParty() *PartyIdentification40Choice {
	i.EmployingParty = new(PartyIdentification40Choice)
	return i.EmployingParty
}

func (i *IndividualPerson26) AddAttendanceCardDetails() *AttendanceCard2 {
	i.AttendanceCardDetails = new(AttendanceCard2)
	return i.AttendanceCardDetails
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson27 struct {

	// Term used to address the person.
	NamePrefix *NamePrefix1Choice `xml:"NmPrfx,omitempty"`

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Second name of the person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which the party is known and which is usually used to identify that person.
	Name *Max350Text `xml:"Nm"`

	// Additional information about the person that follows a person's name, for example, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Date on which the person was born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Country where the person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where the person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where the person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of the person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Address of the person.
	PostalAddress []*PostalAddress21 `xml:"PstlAdr"`

	// Nationality and legal status (minor or major) or rights that an individual may possess.
	Citizenship []*CitizenshipInformation `xml:"Ctznsh,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Specifies if due diligence checks on the political exposure of the investor or account servicer have been carried out and whether these checks are national or foreign. (A politically exposed person is someone who has been entrusted with a prominent public function, or an individual who is closely related to such a person.)
	PoliticallyExposedPersonType *PoliticalExposureType1Choice `xml:"PltclyXpsdPrsnTp,omitempty"`

	// Date of death.
	DeathDate *ISODate `xml:"DthDt,omitempty"`

	// Civil status of the individual person.
	CivilStatus *CivilStatus1Choice `xml:"CvlSts,omitempty"`

	// Highest level of education reached by the individual person.
	EducationLevel *Max35Text `xml:"EdctnLvl,omitempty"`

	// Information related to the person.
	FamilyInformation *PersonalInformation1 `xml:"FmlyInf,omitempty"`
}

func (i *IndividualPerson27) AddNamePrefix() *NamePrefix1Choice {
	i.NamePrefix = new(NamePrefix1Choice)
	return i.NamePrefix
}

func (i *IndividualPerson27) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson27) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson27) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson27) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson27) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson27) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson27) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson27) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson27) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson27) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson27) AddPostalAddress() *PostalAddress21 {
	newValue := new(PostalAddress21)
	i.PostalAddress = append(i.PostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson27) AddCitizenship() *CitizenshipInformation {
	newValue := new(CitizenshipInformation)
	i.Citizenship = append(i.Citizenship, newValue)
	return newValue
}

func (i *IndividualPerson27) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson27) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson27) AddPoliticallyExposedPersonType() *PoliticalExposureType1Choice {
	i.PoliticallyExposedPersonType = new(PoliticalExposureType1Choice)
	return i.PoliticallyExposedPersonType
}

func (i *IndividualPerson27) SetDeathDate(value string) {
	i.DeathDate = (*ISODate)(&value)
}

func (i *IndividualPerson27) AddCivilStatus() *CivilStatus1Choice {
	i.CivilStatus = new(CivilStatus1Choice)
	return i.CivilStatus
}

func (i *IndividualPerson27) SetEducationLevel(value string) {
	i.EducationLevel = (*Max35Text)(&value)
}

func (i *IndividualPerson27) AddFamilyInformation() *PersonalInformation1 {
	i.FamilyInformation = new(PersonalInformation1)
	return i.FamilyInformation
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson28 struct {

	// Term used to address the person.
	NamePrefix *NamePrefix1Choice `xml:"NmPrfx,omitempty"`

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Second name of the person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which the party is known and which is usually used to identify that person.
	Name *Max350Text `xml:"Nm"`

	// Additional information about the person that follows a person's name, for example, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Date on which the person was born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Country where the person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where the person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where the person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of the person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Information related to an address to be inserted, updated or deleted.
	ModifiedPostalAddress []*ModificationScope34 `xml:"ModfdPstlAdr,omitempty"`

	// Citizenship information to be inserted or deleted.
	ModifiedCitizenship []*ModificationScope3 `xml:"ModfdCtznsh,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Specifies if due diligence checks on the political exposure of the investor or account servicer have been carried out and whether these checks are national or foreign. (A politically exposed person is someone who has been entrusted with a prominent public function, or an individual who is closely related to such a person.)
	PoliticallyExposedPersonType *PoliticalExposureType1Choice `xml:"PltclyXpsdPrsnTp,omitempty"`

	// Date of death.
	DeathDate *ISODate `xml:"DthDt,omitempty"`

	// Civil status of the individual person.
	CivilStatus *CivilStatus1Choice `xml:"CvlSts,omitempty"`

	// Highest level of education reached by the individual person.
	EducationLevel *Max35Text `xml:"EdctnLvl,omitempty"`

	// Information related to the person.
	FamilyInformation *PersonalInformation1 `xml:"FmlyInf,omitempty"`
}

func (i *IndividualPerson28) AddNamePrefix() *NamePrefix1Choice {
	i.NamePrefix = new(NamePrefix1Choice)
	return i.NamePrefix
}

func (i *IndividualPerson28) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson28) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson28) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson28) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson28) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson28) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson28) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson28) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson28) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson28) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson28) AddModifiedPostalAddress() *ModificationScope34 {
	newValue := new(ModificationScope34)
	i.ModifiedPostalAddress = append(i.ModifiedPostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson28) AddModifiedCitizenship() *ModificationScope3 {
	newValue := new(ModificationScope3)
	i.ModifiedCitizenship = append(i.ModifiedCitizenship, newValue)
	return newValue
}

func (i *IndividualPerson28) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson28) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson28) AddPoliticallyExposedPersonType() *PoliticalExposureType1Choice {
	i.PoliticallyExposedPersonType = new(PoliticalExposureType1Choice)
	return i.PoliticallyExposedPersonType
}

func (i *IndividualPerson28) SetDeathDate(value string) {
	i.DeathDate = (*ISODate)(&value)
}

func (i *IndividualPerson28) AddCivilStatus() *CivilStatus1Choice {
	i.CivilStatus = new(CivilStatus1Choice)
	return i.CivilStatus
}

func (i *IndividualPerson28) SetEducationLevel(value string) {
	i.EducationLevel = (*Max35Text)(&value)
}

func (i *IndividualPerson28) AddFamilyInformation() *PersonalInformation1 {
	i.FamilyInformation = new(PersonalInformation1)
	return i.FamilyInformation
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson29 struct {

	// Term used to address the person.
	NamePrefix *NamePrefix1Choice `xml:"NmPrfx,omitempty"`

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Second name of the person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which the party is known and which is usually used to identify that person.
	Name *Max350Text `xml:"Nm"`

	// Address of the person.
	PostalAddress []*PostalAddress21 `xml:"PstlAdr"`
}

func (i *IndividualPerson29) AddNamePrefix() *NamePrefix1Choice {
	i.NamePrefix = new(NamePrefix1Choice)
	return i.NamePrefix
}

func (i *IndividualPerson29) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson29) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson29) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson29) AddPostalAddress() *PostalAddress21 {
	newValue := new(PostalAddress21)
	i.PostalAddress = append(i.PostalAddress, newValue)
	return newValue
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson30 struct {

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Second name of the person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which the party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Date on which the person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`
}

func (i *IndividualPerson30) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson30) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson30) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson30) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson30) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson31 struct {

	// Name by which the party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Date on which the person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Country and residential status of the individual, for example, non-permanent resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType2 `xml:"CtryAndResdtlSts,omitempty"`

	// Beneficial owner or its designated agent certifies that it complies with any holding or investment restrictions or requirements of the fund.
	BeneficiaryCertificationCompletion *BeneficiaryCertificationCompletion1Code `xml:"BnfcryCertfctnCmpltn,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification164 `xml:"OthrId,omitempty"`
}

func (i *IndividualPerson31) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson31) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson31) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType2 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType2)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson31) SetBeneficiaryCertificationCompletion(value string) {
	i.BeneficiaryCertificationCompletion = (*BeneficiaryCertificationCompletion1Code)(&value)
}

func (i *IndividualPerson31) AddOtherIdentification() *GenericIdentification164 {
	newValue := new(GenericIdentification164)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson32 struct {

	// Name by which the party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Date on which the person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Country and residential status of the individual, for example, non-permanent resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType2 `xml:"CtryAndResdtlSts,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification164 `xml:"OthrId,omitempty"`
}

func (i *IndividualPerson32) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson32) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson32) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType2 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType2)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson32) AddOtherIdentification() *GenericIdentification164 {
	newValue := new(GenericIdentification164)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson33 struct {

	// Term used to address the person.
	NamePrefix *NamePrefix1Choice `xml:"NmPrfx,omitempty"`

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Second name of the person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which the party is known and which is usually used to identify that person.
	Name *Max350Text `xml:"Nm"`

	// Additional information about the person that follows a person's name, for example, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Gender of the person.
	Gender *Gender1Code `xml:"Gndr,omitempty"`

	// Date on which the person was born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Country where the person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where the person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where the person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of the person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Information related to an address to be inserted, updated or deleted.
	ModifiedPostalAddress []*ModificationScope34 `xml:"ModfdPstlAdr,omitempty"`

	// Citizenship information to be inserted or deleted.
	ModifiedCitizenship []*ModificationScope39 `xml:"ModfdCtznsh,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Specifies if due diligence checks on the political exposure of the investor or account servicer have been carried out and whether these checks are national or foreign. (A politically exposed person is someone who has been entrusted with a prominent public function, or an individual who is closely related to such a person.)
	PoliticallyExposedPersonType *PoliticalExposureType1Choice `xml:"PltclyXpsdPrsnTp,omitempty"`

	// Date of death.
	DeathDate *ISODate `xml:"DthDt,omitempty"`

	// Civil status of the individual person.
	CivilStatus *CivilStatus1Choice `xml:"CvlSts,omitempty"`

	// Highest level of education reached by the individual person.
	EducationLevel *Max35Text `xml:"EdctnLvl,omitempty"`

	// Information related to the person.
	FamilyInformation *PersonalInformation1 `xml:"FmlyInf,omitempty"`
}

func (i *IndividualPerson33) AddNamePrefix() *NamePrefix1Choice {
	i.NamePrefix = new(NamePrefix1Choice)
	return i.NamePrefix
}

func (i *IndividualPerson33) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson33) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson33) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson33) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson33) SetGender(value string) {
	i.Gender = (*Gender1Code)(&value)
}

func (i *IndividualPerson33) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson33) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson33) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson33) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson33) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson33) AddModifiedPostalAddress() *ModificationScope34 {
	newValue := new(ModificationScope34)
	i.ModifiedPostalAddress = append(i.ModifiedPostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson33) AddModifiedCitizenship() *ModificationScope39 {
	newValue := new(ModificationScope39)
	i.ModifiedCitizenship = append(i.ModifiedCitizenship, newValue)
	return newValue
}

func (i *IndividualPerson33) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson33) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson33) AddPoliticallyExposedPersonType() *PoliticalExposureType1Choice {
	i.PoliticallyExposedPersonType = new(PoliticalExposureType1Choice)
	return i.PoliticallyExposedPersonType
}

func (i *IndividualPerson33) SetDeathDate(value string) {
	i.DeathDate = (*ISODate)(&value)
}

func (i *IndividualPerson33) AddCivilStatus() *CivilStatus1Choice {
	i.CivilStatus = new(CivilStatus1Choice)
	return i.CivilStatus
}

func (i *IndividualPerson33) SetEducationLevel(value string) {
	i.EducationLevel = (*Max35Text)(&value)
}

func (i *IndividualPerson33) AddFamilyInformation() *PersonalInformation1 {
	i.FamilyInformation = new(PersonalInformation1)
	return i.FamilyInformation
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson34 struct {

	// Term used to address the person.
	NamePrefix *NamePrefix1Choice `xml:"NmPrfx,omitempty"`

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Second name of the person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which the party is known and which is usually used to identify that person.
	Name *Max350Text `xml:"Nm"`

	// Additional information about the person that follows a person's name, for example, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Gender of the person.
	Gender *Gender1Code `xml:"Gndr,omitempty"`

	// Date on which the person was born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Country where the person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where the person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where the person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of the person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Address of the person.
	PostalAddress []*PostalAddress21 `xml:"PstlAdr"`

	// Nationality and legal status (minor or major).
	Citizenship []*CitizenshipInformation2 `xml:"Ctznsh,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Specifies if due diligence checks on the political exposure of the investor or account servicer have been carried out and whether these checks are national or foreign. (A politically exposed person is someone who has been entrusted with a prominent public function, or an individual who is closely related to such a person.)
	PoliticallyExposedPersonType *PoliticalExposureType1Choice `xml:"PltclyXpsdPrsnTp,omitempty"`

	// Date of death.
	DeathDate *ISODate `xml:"DthDt,omitempty"`

	// Civil status of the individual person.
	CivilStatus *CivilStatus1Choice `xml:"CvlSts,omitempty"`

	// Highest level of education reached by the individual person.
	EducationLevel *Max35Text `xml:"EdctnLvl,omitempty"`

	// Information related to the person.
	FamilyInformation *PersonalInformation1 `xml:"FmlyInf,omitempty"`
}

func (i *IndividualPerson34) AddNamePrefix() *NamePrefix1Choice {
	i.NamePrefix = new(NamePrefix1Choice)
	return i.NamePrefix
}

func (i *IndividualPerson34) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson34) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson34) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson34) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson34) SetGender(value string) {
	i.Gender = (*Gender1Code)(&value)
}

func (i *IndividualPerson34) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson34) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson34) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson34) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson34) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson34) AddPostalAddress() *PostalAddress21 {
	newValue := new(PostalAddress21)
	i.PostalAddress = append(i.PostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson34) AddCitizenship() *CitizenshipInformation2 {
	newValue := new(CitizenshipInformation2)
	i.Citizenship = append(i.Citizenship, newValue)
	return newValue
}

func (i *IndividualPerson34) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson34) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson34) AddPoliticallyExposedPersonType() *PoliticalExposureType1Choice {
	i.PoliticallyExposedPersonType = new(PoliticalExposureType1Choice)
	return i.PoliticallyExposedPersonType
}

func (i *IndividualPerson34) SetDeathDate(value string) {
	i.DeathDate = (*ISODate)(&value)
}

func (i *IndividualPerson34) AddCivilStatus() *CivilStatus1Choice {
	i.CivilStatus = new(CivilStatus1Choice)
	return i.CivilStatus
}

func (i *IndividualPerson34) SetEducationLevel(value string) {
	i.EducationLevel = (*Max35Text)(&value)
}

func (i *IndividualPerson34) AddFamilyInformation() *PersonalInformation1 {
	i.FamilyInformation = new(PersonalInformation1)
	return i.FamilyInformation
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson35 struct {

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Second name of the person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which the party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Gender of the person.
	Gender *Gender1Code `xml:"Gndr,omitempty"`

	// Date on which the person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`
}

func (i *IndividualPerson35) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson35) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson35) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson35) SetGender(value string) {
	i.Gender = (*Gender1Code)(&value)
}

func (i *IndividualPerson35) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson4 struct {

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm"`

	// Second name of a person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Specifies the gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`
}

func (i *IndividualPerson4) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson4) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson4) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson4) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson4) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson8 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm"`

	// Specifies the terms used to formally address a person.
	NamePrefix *NamePrefix1Code `xml:"NmPrfx,omitempty"`

	// Additional information about a person that follows a person's name, eg, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Specifies the gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Number assigned by a social security agency.
	SocialSecurityNumber *Max35Text `xml:"SclSctyNb,omitempty"`

	// Postal address of a party.
	IndividualInvestorAddress *PostalAddress1 `xml:"IndvInvstrAdr"`
}

func (i *IndividualPerson8) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *IndividualPerson8) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson8) SetNamePrefix(value string) {
	i.NamePrefix = (*NamePrefix1Code)(&value)
}

func (i *IndividualPerson8) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson8) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson8) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson8) SetSocialSecurityNumber(value string) {
	i.SocialSecurityNumber = (*Max35Text)(&value)
}

func (i *IndividualPerson8) AddIndividualInvestorAddress() *PostalAddress1 {
	i.IndividualInvestorAddress = new(PostalAddress1)
	return i.IndividualInvestorAddress
}

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson9 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Country and residential status of an individual, for example, non-pernament resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Beneficial owner or its designated agent certifies that it complies with any holding or investment restrictions or requirements of the fund.
	BeneficiaryCertificationCompletion *BeneficiaryCertificationCompletion1Code `xml:"BnfcryCertfctnCmpltn,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification11 `xml:"OthrId,omitempty"`
}

func (i *IndividualPerson9) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson9) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson9) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson9) SetBeneficiaryCertificationCompletion(value string) {
	i.BeneficiaryCertificationCompletion = (*BeneficiaryCertificationCompletion1Code)(&value)
}

func (i *IndividualPerson9) AddOtherIdentification() *GenericIdentification11 {
	newValue := new(GenericIdentification11)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}
