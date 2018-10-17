package iso20022

// Provides information about the beneficial owner of the securities.
type BeneficialOwner1 struct {

	// Identification of the party that is the beneficial owner of the specified securities.
	BeneficialOwnerIdentification *PartyIdentification2Choice `xml:"BnfclOwnrId"`

	// Additional identification of the party that is the beneficial owner of the specified securities.
	AdditionalIdentification *GenericIdentification16 `xml:"AddtlId,omitempty"`

	// Nationality of the beneficial owner.
	Nationality *CountryCode `xml:"Ntlty,omitempty"`

	// Country in which a person is permanently domiciled (the place of a persons permanent home).
	DomicileCountry *CountryCode `xml:"DmclCtry,omitempty"`

	// The holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry *CountryCode `xml:"NonDmclCtry,omitempty"`

	// Whether or not certification is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Type of certification which is required.
	CertificationType *BeneficiaryCertificationType1FormatChoice `xml:"CertfctnTp,omitempty"`

	// Provides declaration details narrative relative to the financial instrument, eg, beneficial ownership.
	DeclarationDetails *Max350Text `xml:"DclrtnDtls,omitempty"`

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification7 `xml:"SctyId,omitempty"`

	// Quantity of securities elected by to the beneficial owner.
	ElectedSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"ElctdSctiesQty"`
}

func (b *BeneficialOwner1) AddBeneficialOwnerIdentification() *PartyIdentification2Choice {
	b.BeneficialOwnerIdentification = new(PartyIdentification2Choice)
	return b.BeneficialOwnerIdentification
}

func (b *BeneficialOwner1) AddAdditionalIdentification() *GenericIdentification16 {
	b.AdditionalIdentification = new(GenericIdentification16)
	return b.AdditionalIdentification
}

func (b *BeneficialOwner1) SetNationality(value string) {
	b.Nationality = (*CountryCode)(&value)
}

func (b *BeneficialOwner1) SetDomicileCountry(value string) {
	b.DomicileCountry = (*CountryCode)(&value)
}

func (b *BeneficialOwner1) SetNonDomicileCountry(value string) {
	b.NonDomicileCountry = (*CountryCode)(&value)
}

func (b *BeneficialOwner1) SetCertificationIndicator(value string) {
	b.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (b *BeneficialOwner1) AddCertificationType() *BeneficiaryCertificationType1FormatChoice {
	b.CertificationType = new(BeneficiaryCertificationType1FormatChoice)
	return b.CertificationType
}

func (b *BeneficialOwner1) SetDeclarationDetails(value string) {
	b.DeclarationDetails = (*Max350Text)(&value)
}

func (b *BeneficialOwner1) AddSecurityIdentification() *SecurityIdentification7 {
	b.SecurityIdentification = new(SecurityIdentification7)
	return b.SecurityIdentification
}

func (b *BeneficialOwner1) AddElectedSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	b.ElectedSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return b.ElectedSecuritiesQuantity
}

// Provides information about the beneficial owner of the securities.
type BeneficialOwner2 struct {

	// Identification of the party that is the beneficial owner of the specified financial instrument.
	BeneficialOwnerIdentification *PartyIdentification100 `xml:"BnfclOwnrId"`

	// Additional identification of the party that is the beneficial owner of the specified financial instrument.
	AdditionalIdentification *AlternateIdentification4 `xml:"AddtlId,omitempty"`

	// Nationality of the beneficial owner.
	Nationality *CountryCode `xml:"Ntlty,omitempty"`

	// Country in which the person is permanently domiciled (the place of a person's permanent home).
	DomicileCountry *CountryCode `xml:"DmclCtry,omitempty"`

	// Country for which the holder of the financial instrument must specify that it is not domiciled. (The holder must certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.)
	NonDomicileCountry *CountryCode `xml:"NonDmclCtry,omitempty"`

	// Indicates whether certification is required from the account owner. (Yes means that certification is required. No means certification is not required.)
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Type of beneficial owner certification that is required.
	CertificationType *BeneficiaryCertificationType9Choice `xml:"CertfctnTp,omitempty"`

	// Declaration details related to the financial instrument, for example, beneficial ownership.
	DeclarationDetails *Max350Text `xml:"DclrtnDtls,omitempty"`
}

func (b *BeneficialOwner2) AddBeneficialOwnerIdentification() *PartyIdentification100 {
	b.BeneficialOwnerIdentification = new(PartyIdentification100)
	return b.BeneficialOwnerIdentification
}

func (b *BeneficialOwner2) AddAdditionalIdentification() *AlternateIdentification4 {
	b.AdditionalIdentification = new(AlternateIdentification4)
	return b.AdditionalIdentification
}

func (b *BeneficialOwner2) SetNationality(value string) {
	b.Nationality = (*CountryCode)(&value)
}

func (b *BeneficialOwner2) SetDomicileCountry(value string) {
	b.DomicileCountry = (*CountryCode)(&value)
}

func (b *BeneficialOwner2) SetNonDomicileCountry(value string) {
	b.NonDomicileCountry = (*CountryCode)(&value)
}

func (b *BeneficialOwner2) SetCertificationIndicator(value string) {
	b.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (b *BeneficialOwner2) AddCertificationType() *BeneficiaryCertificationType9Choice {
	b.CertificationType = new(BeneficiaryCertificationType9Choice)
	return b.CertificationType
}

func (b *BeneficialOwner2) SetDeclarationDetails(value string) {
	b.DeclarationDetails = (*Max350Text)(&value)
}
