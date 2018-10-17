package iso20022

// Set of unique and unambiguous ways to identify a person.
type PersonIdentification11 struct {

	// Number assigned by a license authority to a driver's license.
	DriverLicenseNumber *Max35Text `xml:"DrvrLicNb,omitempty"`

	// Country, state or province, issuer of the driver license.
	DriverLicenseLocation *Max35Text `xml:"DrvrLicLctn,omitempty"`

	// Name or title of the driver license.
	DriverLicenseName *Max35Text `xml:"DrvrLicNm,omitempty"`

	// Identification of the driver in the fleet of vehicle.
	DriverIdentification *Max35Text `xml:"DrvrId,omitempty"`

	// Number assigned by an agent to identify its customer.
	CustomerNumber *Max35Text `xml:"CstmrNb,omitempty"`

	// Number assigned by a social security agency.
	SocialSecurityNumber *Max35Text `xml:"SclSctyNb,omitempty"`

	// Number assigned by a government agency to identify foreign nationals.
	AlienRegistrationNumber *Max35Text `xml:"AlnRegnNb,omitempty"`

	// Number assigned by a passport authority to a passport.
	PassportNumber *Max35Text `xml:"PsptNb,omitempty"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb,omitempty"`

	// Number assigned by a national authority to an identity card.
	IdentityCardNumber *Max35Text `xml:"IdntyCardNb,omitempty"`

	// Number assigned to an employer by a registration authority.
	EmployerIdentificationNumber *Max35Text `xml:"MplyrIdNb,omitempty"`

	// Number assigned to an employee by a employer.
	EmployeeIdentificationNumber *Max35Text `xml:"MplyeeIdNb,omitempty"`

	// Identification of the job.
	JobNumber *Max35Text `xml:"JobNb,omitempty"`

	// Identification of the department.
	Department *Max35Text `xml:"Dept,omitempty"`

	// Address for electronic mail (e-mail).
	EmailAddress *Max256Text `xml:"EmailAdr,omitempty"`

	// Date and place of birth of a person.
	DateAndPlaceOfBirth *DateAndPlaceOfBirth `xml:"DtAndPlcOfBirth,omitempty"`

	// Unique identification of a person, as assigned by an institution, using an identification scheme.
	Other []*GenericIdentification4 `xml:"Othr,omitempty"`
}

func (p *PersonIdentification11) SetDriverLicenseNumber(value string) {
	p.DriverLicenseNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetDriverLicenseLocation(value string) {
	p.DriverLicenseLocation = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetDriverLicenseName(value string) {
	p.DriverLicenseName = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetDriverIdentification(value string) {
	p.DriverIdentification = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetCustomerNumber(value string) {
	p.CustomerNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetSocialSecurityNumber(value string) {
	p.SocialSecurityNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetAlienRegistrationNumber(value string) {
	p.AlienRegistrationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetPassportNumber(value string) {
	p.PassportNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetTaxIdentificationNumber(value string) {
	p.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetIdentityCardNumber(value string) {
	p.IdentityCardNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetEmployerIdentificationNumber(value string) {
	p.EmployerIdentificationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetEmployeeIdentificationNumber(value string) {
	p.EmployeeIdentificationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetJobNumber(value string) {
	p.JobNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetDepartment(value string) {
	p.Department = (*Max35Text)(&value)
}

func (p *PersonIdentification11) SetEmailAddress(value string) {
	p.EmailAddress = (*Max256Text)(&value)
}

func (p *PersonIdentification11) AddDateAndPlaceOfBirth() *DateAndPlaceOfBirth {
	p.DateAndPlaceOfBirth = new(DateAndPlaceOfBirth)
	return p.DateAndPlaceOfBirth
}

func (p *PersonIdentification11) AddOther() *GenericIdentification4 {
	newValue := new(GenericIdentification4)
	p.Other = append(p.Other, newValue)
	return newValue
}

// Unique and unambiguous way to identify a person.
type PersonIdentification2 struct {

	// Number assigned by a license authority to a driver's license.
	DriversLicenseNumber *Max35Text `xml:"DrvrsLicNb"`

	// Number assigned by a social security agency.
	SocialSecurityNumber *Max35Text `xml:"SclSctyNb"`

	// Number assigned by a government agency to identify foreign nationals.
	AlienRegistrationNumber *Max35Text `xml:"AlnRegnNb"`

	// Number assigned by a passport authority to a passport.
	PassportNumber *Max35Text `xml:"PsptNb"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb"`

	// Number assigned by a national authority to an identity card.
	IdentityCardNumber *Max35Text `xml:"IdntyCardNb"`

	// Number assigned to an employer by a registration authority.
	EmployerIdentificationNumber *Max35Text `xml:"MplyrIdNb"`

	// Identifier issued to a person for which no specific identifier has been defined.
	OtherIdentification *GenericIdentification4 `xml:"OthrId"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (p *PersonIdentification2) SetDriversLicenseNumber(value string) {
	p.DriversLicenseNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) SetSocialSecurityNumber(value string) {
	p.SocialSecurityNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) SetAlienRegistrationNumber(value string) {
	p.AlienRegistrationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) SetPassportNumber(value string) {
	p.PassportNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) SetTaxIdentificationNumber(value string) {
	p.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) SetIdentityCardNumber(value string) {
	p.IdentityCardNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) SetEmployerIdentificationNumber(value string) {
	p.EmployerIdentificationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification2) AddOtherIdentification() *GenericIdentification4 {
	p.OtherIdentification = new(GenericIdentification4)
	return p.OtherIdentification
}

func (p *PersonIdentification2) SetIssuer(value string) {
	p.Issuer = (*Max35Text)(&value)
}

// Unique and unambiguous way to identify a person.
type PersonIdentification3 struct {

	// Number assigned by a license authority to a driver's license.
	DriversLicenseNumber *Max35Text `xml:"DrvrsLicNb"`

	// Number assigned by an agent  to identify its customer.
	CustomerNumber *Max35Text `xml:"CstmrNb"`

	// Number assigned by a social security agency.
	SocialSecurityNumber *Max35Text `xml:"SclSctyNb"`

	// Number assigned by a government agency to identify foreign nationals.
	AlienRegistrationNumber *Max35Text `xml:"AlnRegnNb"`

	// Number assigned by a passport authority to a passport.
	PassportNumber *Max35Text `xml:"PsptNb"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb"`

	// Number assigned by a national authority to an identity card.
	IdentityCardNumber *Max35Text `xml:"IdntyCardNb"`

	// Number assigned to an employer by a registration authority.
	EmployerIdentificationNumber *Max35Text `xml:"MplyrIdNb"`

	// Date and place of birth of a person.
	DateAndPlaceOfBirth *DateAndPlaceOfBirth `xml:"DtAndPlcOfBirth"`

	// Identifier issued to a person for which no specific identifier has been defined.
	OtherIdentification *GenericIdentification4 `xml:"OthrId"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (p *PersonIdentification3) SetDriversLicenseNumber(value string) {
	p.DriversLicenseNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification3) SetCustomerNumber(value string) {
	p.CustomerNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification3) SetSocialSecurityNumber(value string) {
	p.SocialSecurityNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification3) SetAlienRegistrationNumber(value string) {
	p.AlienRegistrationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification3) SetPassportNumber(value string) {
	p.PassportNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification3) SetTaxIdentificationNumber(value string) {
	p.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification3) SetIdentityCardNumber(value string) {
	p.IdentityCardNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification3) SetEmployerIdentificationNumber(value string) {
	p.EmployerIdentificationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification3) AddDateAndPlaceOfBirth() *DateAndPlaceOfBirth {
	p.DateAndPlaceOfBirth = new(DateAndPlaceOfBirth)
	return p.DateAndPlaceOfBirth
}

func (p *PersonIdentification3) AddOtherIdentification() *GenericIdentification4 {
	p.OtherIdentification = new(GenericIdentification4)
	return p.OtherIdentification
}

func (p *PersonIdentification3) SetIssuer(value string) {
	p.Issuer = (*Max35Text)(&value)
}

// Unique and unambiguous way to identify a person.
type PersonIdentification5 struct {

	// Date and place of birth of a person.
	DateAndPlaceOfBirth *DateAndPlaceOfBirth `xml:"DtAndPlcOfBirth,omitempty"`

	// Unique identification of a person, as assigned by an institution, using an identification scheme.
	Other []*GenericPersonIdentification1 `xml:"Othr,omitempty"`
}

func (p *PersonIdentification5) AddDateAndPlaceOfBirth() *DateAndPlaceOfBirth {
	p.DateAndPlaceOfBirth = new(DateAndPlaceOfBirth)
	return p.DateAndPlaceOfBirth
}

func (p *PersonIdentification5) AddOther() *GenericPersonIdentification1 {
	newValue := new(GenericPersonIdentification1)
	p.Other = append(p.Other, newValue)
	return newValue
}

// Unique and unambiguous way to identify a person.
type PersonIdentification6 struct {

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Personal identification type.
	PersonIdentificationType *PersonIdentificationType1Choice `xml:"PrsnIdTp"`
}

func (p *PersonIdentification6) SetIssuer(value string) {
	p.Issuer = (*Max35Text)(&value)
}

func (p *PersonIdentification6) AddPersonIdentificationType() *PersonIdentificationType1Choice {
	p.PersonIdentificationType = new(PersonIdentificationType1Choice)
	return p.PersonIdentificationType
}

// Set of unique and unambiguous ways to identify a person.
type PersonIdentification7 struct {

	// Number assigned by a license authority to a driver's license.
	DriversLicenseNumber *Max35Text `xml:"DrvrsLicNb,omitempty"`

	// Number assigned by an agent to identify its customer.
	CustomerNumber *Max35Text `xml:"CstmrNb,omitempty"`

	// Number assigned by a social security agency.
	SocialSecurityNumber *Max35Text `xml:"SclSctyNb,omitempty"`

	// Number assigned by a government agency to identify foreign nationals.
	AlienRegistrationNumber *Max35Text `xml:"AlnRegnNb,omitempty"`

	// Number assigned by a passport authority to a passport.
	PassportNumber *Max35Text `xml:"PsptNb,omitempty"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb,omitempty"`

	// Number assigned by a national authority to an identity card.
	IdentityCardNumber *Max35Text `xml:"IdntyCardNb,omitempty"`

	// Number assigned to an employer by a registration authority.
	EmployerIdentificationNumber *Max35Text `xml:"MplyrIdNb,omitempty"`

	// Number assigned to an employee by a employer.
	EmployeeIdentificationNumber *Max35Text `xml:"MplyeeIdNb,omitempty"`

	// Address for electronic mail (e-mail).
	EmailAddress *Max256Text `xml:"EmailAdr,omitempty"`

	// Date and place of birth of a person.
	DateAndPlaceOfBirth *DateAndPlaceOfBirth `xml:"DtAndPlcOfBirth,omitempty"`

	// Unique identification of a person, as assigned by an institution, using an identification scheme.
	Other []*GenericIdentification4 `xml:"Othr,omitempty"`
}

func (p *PersonIdentification7) SetDriversLicenseNumber(value string) {
	p.DriversLicenseNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification7) SetCustomerNumber(value string) {
	p.CustomerNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification7) SetSocialSecurityNumber(value string) {
	p.SocialSecurityNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification7) SetAlienRegistrationNumber(value string) {
	p.AlienRegistrationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification7) SetPassportNumber(value string) {
	p.PassportNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification7) SetTaxIdentificationNumber(value string) {
	p.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification7) SetIdentityCardNumber(value string) {
	p.IdentityCardNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification7) SetEmployerIdentificationNumber(value string) {
	p.EmployerIdentificationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification7) SetEmployeeIdentificationNumber(value string) {
	p.EmployeeIdentificationNumber = (*Max35Text)(&value)
}

func (p *PersonIdentification7) SetEmailAddress(value string) {
	p.EmailAddress = (*Max256Text)(&value)
}

func (p *PersonIdentification7) AddDateAndPlaceOfBirth() *DateAndPlaceOfBirth {
	p.DateAndPlaceOfBirth = new(DateAndPlaceOfBirth)
	return p.DateAndPlaceOfBirth
}

func (p *PersonIdentification7) AddOther() *GenericIdentification4 {
	newValue := new(GenericIdentification4)
	p.Other = append(p.Other, newValue)
	return newValue
}
