package iso20022

// Set of elements providing specific information on the direct debit transaction and the related mandate.
type DirectDebitTransaction1 struct {

	// Set of elements used to provide further details related to a direct debit mandate signed between the creditor and the debtor.
	//
	// Usage: Mandate related information is to be used only when the direct debit relates to a mandate signed between the debtor and the creditor.
	MandateRelatedInformation *MandateRelatedInformation1 `xml:"MndtRltdInf,omitempty"`

	// Credit party that signs the direct debit mandate.
	CreditorSchemeIdentification *PartyIdentification8 `xml:"CdtrSchmeId,omitempty"`

	// Unique and unambiguous identification of the pre-notification which is sent separately from the direct debit instruction.
	//
	// Usage: The direct debit pre-notification is used to reconcile separately sent collection information with the direct debit transaction information.
	PreNotificationIdentification *Max35Text `xml:"PreNtfctnId,omitempty"`

	// Date on which the creditor notifies the debtor about the amount and date on which the direct debit instruction will be presented to the debtor's agent.
	PreNotificationDate *ISODate `xml:"PreNtfctnDt,omitempty"`
}

func (d *DirectDebitTransaction1) AddMandateRelatedInformation() *MandateRelatedInformation1 {
	d.MandateRelatedInformation = new(MandateRelatedInformation1)
	return d.MandateRelatedInformation
}

func (d *DirectDebitTransaction1) AddCreditorSchemeIdentification() *PartyIdentification8 {
	d.CreditorSchemeIdentification = new(PartyIdentification8)
	return d.CreditorSchemeIdentification
}

func (d *DirectDebitTransaction1) SetPreNotificationIdentification(value string) {
	d.PreNotificationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitTransaction1) SetPreNotificationDate(value string) {
	d.PreNotificationDate = (*ISODate)(&value)
}

// Set of elements used to provide specific information on the direct debit transaction and the related mandate.
type DirectDebitTransaction6 struct {

	// Set of elements used to provide further details of the direct debit mandate signed between the creditor and the debtor.
	MandateRelatedInformation *MandateRelatedInformation6 `xml:"MndtRltdInf,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification32 `xml:"CdtrSchmeId,omitempty"`

	// Unique and unambiguous identification of the pre-notification which is sent separately from the direct debit instruction.
	//
	// Usage: The direct debit pre-notification is used to reconcile separately sent collection information with the direct debit transaction information.
	PreNotificationIdentification *Max35Text `xml:"PreNtfctnId,omitempty"`

	// Date on which the creditor notifies the debtor about the amount and date on which the direct debit instruction will be presented to the debtor's agent.
	PreNotificationDate *ISODate `xml:"PreNtfctnDt,omitempty"`
}

func (d *DirectDebitTransaction6) AddMandateRelatedInformation() *MandateRelatedInformation6 {
	d.MandateRelatedInformation = new(MandateRelatedInformation6)
	return d.MandateRelatedInformation
}

func (d *DirectDebitTransaction6) AddCreditorSchemeIdentification() *PartyIdentification32 {
	d.CreditorSchemeIdentification = new(PartyIdentification32)
	return d.CreditorSchemeIdentification
}

func (d *DirectDebitTransaction6) SetPreNotificationIdentification(value string) {
	d.PreNotificationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitTransaction6) SetPreNotificationDate(value string) {
	d.PreNotificationDate = (*ISODate)(&value)
}

// Provides further specific details on the direct debit transaction and the related mandate.
type DirectDebitTransaction7 struct {

	// Provides further details of the direct debit mandate signed between the creditor and the debtor.
	MandateRelatedInformation *MandateRelatedInformation8 `xml:"MndtRltdInf,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Unique and unambiguous identification of the pre-notification which is sent separately from the direct debit instruction.
	//
	// Usage: The direct debit pre-notification is used to reconcile separately sent collection information with the direct debit transaction information.
	PreNotificationIdentification *Max35Text `xml:"PreNtfctnId,omitempty"`

	// Date on which the creditor notifies the debtor about the amount and date on which the direct debit instruction will be presented to the debtor's agent.
	PreNotificationDate *ISODate `xml:"PreNtfctnDt,omitempty"`
}

func (d *DirectDebitTransaction7) AddMandateRelatedInformation() *MandateRelatedInformation8 {
	d.MandateRelatedInformation = new(MandateRelatedInformation8)
	return d.MandateRelatedInformation
}

func (d *DirectDebitTransaction7) AddCreditorSchemeIdentification() *PartyIdentification43 {
	d.CreditorSchemeIdentification = new(PartyIdentification43)
	return d.CreditorSchemeIdentification
}

func (d *DirectDebitTransaction7) SetPreNotificationIdentification(value string) {
	d.PreNotificationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitTransaction7) SetPreNotificationDate(value string) {
	d.PreNotificationDate = (*ISODate)(&value)
}

// Provides further specific details on the direct debit transaction and the related mandate.
type DirectDebitTransaction8 struct {

	// Provides further details of the direct debit mandate signed between the creditor and the debtor.
	MandateRelatedInformation *MandateRelatedInformation10 `xml:"MndtRltdInf,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Unique and unambiguous identification of the pre-notification which is sent separately from the direct debit instruction.
	//
	// Usage: The direct debit pre-notification is used to reconcile separately sent collection information with the direct debit transaction information.
	PreNotificationIdentification *Max35Text `xml:"PreNtfctnId,omitempty"`

	// Date on which the creditor notifies the debtor about the amount and date on which the direct debit instruction will be presented to the debtor's agent.
	PreNotificationDate *ISODate `xml:"PreNtfctnDt,omitempty"`
}

func (d *DirectDebitTransaction8) AddMandateRelatedInformation() *MandateRelatedInformation10 {
	d.MandateRelatedInformation = new(MandateRelatedInformation10)
	return d.MandateRelatedInformation
}

func (d *DirectDebitTransaction8) AddCreditorSchemeIdentification() *PartyIdentification43 {
	d.CreditorSchemeIdentification = new(PartyIdentification43)
	return d.CreditorSchemeIdentification
}

func (d *DirectDebitTransaction8) SetPreNotificationIdentification(value string) {
	d.PreNotificationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitTransaction8) SetPreNotificationDate(value string) {
	d.PreNotificationDate = (*ISODate)(&value)
}

// Provides further specific details on the direct debit transaction and the related mandate.
type DirectDebitTransaction9 struct {

	// Provides further details of the direct debit mandate signed between the creditor and the debtor.
	MandateRelatedInformation *MandateRelatedInformation11 `xml:"MndtRltdInf,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Unique and unambiguous identification of the pre-notification which is sent separately from the direct debit instruction.
	//
	// Usage: The direct debit pre-notification is used to reconcile separately sent collection information with the direct debit transaction information.
	PreNotificationIdentification *Max35Text `xml:"PreNtfctnId,omitempty"`

	// Date on which the creditor notifies the debtor about the amount and date on which the direct debit instruction will be presented to the debtor's agent.
	PreNotificationDate *ISODate `xml:"PreNtfctnDt,omitempty"`
}

func (d *DirectDebitTransaction9) AddMandateRelatedInformation() *MandateRelatedInformation11 {
	d.MandateRelatedInformation = new(MandateRelatedInformation11)
	return d.MandateRelatedInformation
}

func (d *DirectDebitTransaction9) AddCreditorSchemeIdentification() *PartyIdentification43 {
	d.CreditorSchemeIdentification = new(PartyIdentification43)
	return d.CreditorSchemeIdentification
}

func (d *DirectDebitTransaction9) SetPreNotificationIdentification(value string) {
	d.PreNotificationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitTransaction9) SetPreNotificationDate(value string) {
	d.PreNotificationDate = (*ISODate)(&value)
}
