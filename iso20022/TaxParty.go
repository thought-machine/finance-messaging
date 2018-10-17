package iso20022

// Details about the entity involved in the tax paid or to be paid.
type TaxParty1 struct {

	// Tax identification number of the creditor.
	TaxIdentification *Max35Text `xml:"TaxId,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	RegistrationIdentification *Max35Text `xml:"RegnId,omitempty"`

	// Type of tax payer.
	TaxType *Max35Text `xml:"TaxTp,omitempty"`
}

func (t *TaxParty1) SetTaxIdentification(value string) {
	t.TaxIdentification = (*Max35Text)(&value)
}

func (t *TaxParty1) SetRegistrationIdentification(value string) {
	t.RegistrationIdentification = (*Max35Text)(&value)
}

func (t *TaxParty1) SetTaxType(value string) {
	t.TaxType = (*Max35Text)(&value)
}

// Details about the entity involved in the tax paid or to be paid.
type TaxParty2 struct {

	// Tax identification number of the debtor.
	TaxIdentification *Max35Text `xml:"TaxId,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	RegistrationIdentification *Max35Text `xml:"RegnId,omitempty"`

	// Type of tax payer.
	TaxType *Max35Text `xml:"TaxTp,omitempty"`

	// Details of the authorised tax paying party.
	Authorisation *TaxAuthorisation1 `xml:"Authstn,omitempty"`
}

func (t *TaxParty2) SetTaxIdentification(value string) {
	t.TaxIdentification = (*Max35Text)(&value)
}

func (t *TaxParty2) SetRegistrationIdentification(value string) {
	t.RegistrationIdentification = (*Max35Text)(&value)
}

func (t *TaxParty2) SetTaxType(value string) {
	t.TaxType = (*Max35Text)(&value)
}

func (t *TaxParty2) AddAuthorisation() *TaxAuthorisation1 {
	t.Authorisation = new(TaxAuthorisation1)
	return t.Authorisation
}

// Details about the entity involved in the tax paid or to be paid.
type TaxParty3 struct {

	// Number assigned by a tax authority to an entity.
	TaxIdentification *Max35Text `xml:"TaxId,omitempty"`

	// Type of tax payer.
	TaxType *Max35Text `xml:"TaxTp,omitempty"`

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	RegistrationIdentification *Max35Text `xml:"RegnId,omitempty"`

	// Specification of the tax exemption reason.
	TaxExemptionReason []*TaxExemptionReasonFormatChoice `xml:"TaxXmptnRsn,omitempty"`
}

func (t *TaxParty3) SetTaxIdentification(value string) {
	t.TaxIdentification = (*Max35Text)(&value)
}

func (t *TaxParty3) SetTaxType(value string) {
	t.TaxType = (*Max35Text)(&value)
}

func (t *TaxParty3) SetRegistrationIdentification(value string) {
	t.RegistrationIdentification = (*Max35Text)(&value)
}

func (t *TaxParty3) AddTaxExemptionReason() *TaxExemptionReasonFormatChoice {
	newValue := new(TaxExemptionReasonFormatChoice)
	t.TaxExemptionReason = append(t.TaxExemptionReason, newValue)
	return newValue
}
