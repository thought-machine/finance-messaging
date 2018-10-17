package iso20022

// Provides addtional information such as the taxation conditions.
type CorporateActionNarrative1 struct {

	// Provides conditional information related to the event, eg, an offer is subject to 50% acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *Max350Text `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, eg, not open to US/Canadian residents, QIB or SIL to be provided.
	InformationToComplyWith *Max350Text `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message.
	TaxationConditions *Max350Text `xml:"TaxtnConds,omitempty"`

	// Provide  the new name of a company following a name change
	NewCompanyName *Max350Text `xml:"NewCpnyNm,omitempty"`

	// Provides the entity making the offer and is different from the issuing company.
	Offeror *PartyIdentification2Choice `xml:"Offerr,omitempty"`

	// Provides the web address published for the event, ie the address for the Universal Resource Locator (URL), eg, used over the www (HTTP) service.
	URLAddress *Max256Text `xml:"URLAdr,omitempty"`

	// Provides additional information or specifies in more detail the content of a
	// message.
	AdditionalText *Max350Text `xml:"AddtlTxt,omitempty"`
}

func (c *CorporateActionNarrative1) SetInformationConditions(value string) {
	c.InformationConditions = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative1) SetInformationToComplyWith(value string) {
	c.InformationToComplyWith = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative1) SetTaxationConditions(value string) {
	c.TaxationConditions = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative1) SetNewCompanyName(value string) {
	c.NewCompanyName = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative1) AddOfferor() *PartyIdentification2Choice {
	c.Offeror = new(PartyIdentification2Choice)
	return c.Offeror
}

func (c *CorporateActionNarrative1) SetURLAddress(value string) {
	c.URLAddress = (*Max256Text)(&value)
}

func (c *CorporateActionNarrative1) SetAdditionalText(value string) {
	c.AdditionalText = (*Max350Text)(&value)
}

// Provides additional information such as the information conditions.
type CorporateActionNarrative10 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText []*Max350Text `xml:"AddtlTxt,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*Max350Text `xml:"PtyCtctNrrtv,omitempty"`
}

func (c *CorporateActionNarrative10) AddAdditionalText(value string) {
	c.AdditionalText = append(c.AdditionalText, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative10) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*Max350Text)(&value))
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative11 struct {

	// Provides additional information or specifies in more detail the content of a message.  This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation2 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation2 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation2 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation2 `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation2 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation2 `xml:"Dsclmr,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative *UpdatedAdditionalInformation2 `xml:"PtyCtctNrrtv,omitempty"`

	// Provides declaration details narrative relative to the financial instrument, for example, beneficial ownership.
	DeclarationDetails *UpdatedAdditionalInformation2 `xml:"DclrtnDtls,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *UpdatedAdditionalInformation2 `xml:"RegnDtls,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation *UpdatedAdditionalInformation2 `xml:"BsktOrIndxInf,omitempty"`
}

func (c *CorporateActionNarrative11) AddAdditionalText() *UpdatedAdditionalInformation2 {
	c.AdditionalText = new(UpdatedAdditionalInformation2)
	return c.AdditionalText
}

func (c *CorporateActionNarrative11) AddNarrativeVersion() *UpdatedAdditionalInformation2 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation2)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative11) AddInformationConditions() *UpdatedAdditionalInformation2 {
	c.InformationConditions = new(UpdatedAdditionalInformation2)
	return c.InformationConditions
}

func (c *CorporateActionNarrative11) AddInformationToComplyWith() *UpdatedAdditionalInformation2 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation2)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative11) AddTaxationConditions() *UpdatedAdditionalInformation2 {
	c.TaxationConditions = new(UpdatedAdditionalInformation2)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative11) AddDisclaimer() *UpdatedAdditionalInformation2 {
	c.Disclaimer = new(UpdatedAdditionalInformation2)
	return c.Disclaimer
}

func (c *CorporateActionNarrative11) AddPartyContactNarrative() *UpdatedAdditionalInformation2 {
	c.PartyContactNarrative = new(UpdatedAdditionalInformation2)
	return c.PartyContactNarrative
}

func (c *CorporateActionNarrative11) AddDeclarationDetails() *UpdatedAdditionalInformation2 {
	c.DeclarationDetails = new(UpdatedAdditionalInformation2)
	return c.DeclarationDetails
}

func (c *CorporateActionNarrative11) AddRegistrationDetails() *UpdatedAdditionalInformation2 {
	c.RegistrationDetails = new(UpdatedAdditionalInformation2)
	return c.RegistrationDetails
}

func (c *CorporateActionNarrative11) AddBasketOrIndexInformation() *UpdatedAdditionalInformation2 {
	c.BasketOrIndexInformation = new(UpdatedAdditionalInformation2)
	return c.BasketOrIndexInformation
}

// Provides additional information such as the information conditions.
type CorporateActionNarrative19 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText []*RestrictedFINXMax350Text `xml:"AddtlTxt,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*RestrictedFINXMax350Text `xml:"PtyCtctNrrtv,omitempty"`
}

func (c *CorporateActionNarrative19) AddAdditionalText(value string) {
	c.AdditionalText = append(c.AdditionalText, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative19) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*RestrictedFINXMax350Text)(&value))
}

// Provides additional information about the CA event.
type CorporateActionNarrative2 struct {

	// Provides conditional information related to the event, eg, an offer is subject to 50% acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *Max350Text `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, eg, not open to US/Canadian residents, QIB or SIL to be provided.
	InformationToComplyWith *Max350Text `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the SLA.
	TaxationConditions *Max350Text `xml:"TaxtnConds,omitempty"`

	// Provides declaration details narrative relative to the financial instrument, eg, beneficial ownership.
	DeclarationDetails *Max350Text `xml:"DclrtnDtls,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *Max350Text `xml:"RegnDtls,omitempty"`

	// Provides additional information or specifies in more detail the content of a message.
	AdditionalText *Max350Text `xml:"AddtlTxt,omitempty"`
}

func (c *CorporateActionNarrative2) SetInformationConditions(value string) {
	c.InformationConditions = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative2) SetInformationToComplyWith(value string) {
	c.InformationToComplyWith = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative2) SetTaxationConditions(value string) {
	c.TaxationConditions = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative2) SetDeclarationDetails(value string) {
	c.DeclarationDetails = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative2) SetRegistrationDetails(value string) {
	c.RegistrationDetails = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative2) SetAdditionalText(value string) {
	c.AdditionalText = (*Max350Text)(&value)
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative20 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation3 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation3 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation1 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation1 `xml:"InfToCmplyWth,omitempty"`

	// Provides restriction(s) on securities.
	SecurityRestriction *UpdatedAdditionalInformation1 `xml:"SctyRstrctn,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message  and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation1 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation1 `xml:"Dsclmr,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *UpdatedAdditionalInformation1 `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative20) AddAdditionalText() *UpdatedAdditionalInformation3 {
	c.AdditionalText = new(UpdatedAdditionalInformation3)
	return c.AdditionalText
}

func (c *CorporateActionNarrative20) AddNarrativeVersion() *UpdatedAdditionalInformation3 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation3)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative20) AddInformationConditions() *UpdatedAdditionalInformation1 {
	c.InformationConditions = new(UpdatedAdditionalInformation1)
	return c.InformationConditions
}

func (c *CorporateActionNarrative20) AddInformationToComplyWith() *UpdatedAdditionalInformation1 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation1)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative20) AddSecurityRestriction() *UpdatedAdditionalInformation1 {
	c.SecurityRestriction = new(UpdatedAdditionalInformation1)
	return c.SecurityRestriction
}

func (c *CorporateActionNarrative20) AddTaxationConditions() *UpdatedAdditionalInformation1 {
	c.TaxationConditions = new(UpdatedAdditionalInformation1)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative20) AddDisclaimer() *UpdatedAdditionalInformation1 {
	c.Disclaimer = new(UpdatedAdditionalInformation1)
	return c.Disclaimer
}

func (c *CorporateActionNarrative20) AddCertificationBreakdown() *UpdatedAdditionalInformation1 {
	c.CertificationBreakdown = new(UpdatedAdditionalInformation1)
	return c.CertificationBreakdown
}

// Provides additional information such as the registration details.
type CorporateActionNarrative21 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText []*Max350Text `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion []*Max350Text `xml:"NrrtvVrsn,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails []*Max350Text `xml:"RegnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*Max350Text `xml:"PtyCtctNrrtv,omitempty"`

	// Disclaimer relative to the information provided in the message. It may be ignored for automated processing. No information about the instruction itself is allowed here.
	Disclaimer []*Max350Text `xml:"Dsclmr,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation []*Max350Text `xml:"BsktOrIndxInf,omitempty"`

	// Provides information required for the certification/breakdown.
	CertificationBreakdown []*Max350Text `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative21) AddAdditionalText(value string) {
	c.AdditionalText = append(c.AdditionalText, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative21) AddNarrativeVersion(value string) {
	c.NarrativeVersion = append(c.NarrativeVersion, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative21) AddRegistrationDetails(value string) {
	c.RegistrationDetails = append(c.RegistrationDetails, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative21) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative21) AddDisclaimer(value string) {
	c.Disclaimer = append(c.Disclaimer, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative21) AddBasketOrIndexInformation(value string) {
	c.BasketOrIndexInformation = append(c.BasketOrIndexInformation, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative21) AddCertificationBreakdown(value string) {
	c.CertificationBreakdown = append(c.CertificationBreakdown, (*Max350Text)(&value))
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative24 struct {

	// Provides the entity making the offer and is different from the issuing company.
	Offeror []*UpdatedAdditionalInformation3 `xml:"Offerr,omitempty"`

	// Provides the new name of a company following a name change.
	NewCompanyName *UpdatedAdditionalInformation3 `xml:"NewCpnyNm,omitempty"`

	// Provides the web address published for the event, that is, the address for the Universal Resource Locator (URL), for example, used over the www (HTTP) service.
	URLAddress *UpdatedURLlnformation `xml:"URLAdr,omitempty"`
}

func (c *CorporateActionNarrative24) AddOfferor() *UpdatedAdditionalInformation3 {
	newValue := new(UpdatedAdditionalInformation3)
	c.Offeror = append(c.Offeror, newValue)
	return newValue
}

func (c *CorporateActionNarrative24) AddNewCompanyName() *UpdatedAdditionalInformation3 {
	c.NewCompanyName = new(UpdatedAdditionalInformation3)
	return c.NewCompanyName
}

func (c *CorporateActionNarrative24) AddURLAddress() *UpdatedURLlnformation {
	c.URLAddress = new(UpdatedURLlnformation)
	return c.URLAddress
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative26 struct {

	// Provides the entity making the offer and is different from the issuing company.
	Offeror []*UpdatedAdditionalInformation3 `xml:"Offerr,omitempty"`

	// Provides the new name of a company following a name change.
	NewCompanyName *UpdatedAdditionalInformation3 `xml:"NewCpnyNm,omitempty"`

	// Provides the web address published for the event, that is, the address for the Universal Resource Locator (URL), for example, used over the www (HTTP) service.
	URLAddress *UpdatedURLlnformation2 `xml:"URLAdr,omitempty"`
}

func (c *CorporateActionNarrative26) AddOfferor() *UpdatedAdditionalInformation3 {
	newValue := new(UpdatedAdditionalInformation3)
	c.Offeror = append(c.Offeror, newValue)
	return newValue
}

func (c *CorporateActionNarrative26) AddNewCompanyName() *UpdatedAdditionalInformation3 {
	c.NewCompanyName = new(UpdatedAdditionalInformation3)
	return c.NewCompanyName
}

func (c *CorporateActionNarrative26) AddURLAddress() *UpdatedURLlnformation2 {
	c.URLAddress = new(UpdatedURLlnformation2)
	return c.URLAddress
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative27 struct {

	// Provides additional information or specifies in more detail the content of a message.  This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation8 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation8 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation8 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation8 `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation8 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation8 `xml:"Dsclmr,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative *UpdatedAdditionalInformation8 `xml:"PtyCtctNrrtv,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *UpdatedAdditionalInformation8 `xml:"RegnDtls,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation *UpdatedAdditionalInformation8 `xml:"BsktOrIndxInf,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *UpdatedAdditionalInformation8 `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative27) AddAdditionalText() *UpdatedAdditionalInformation8 {
	c.AdditionalText = new(UpdatedAdditionalInformation8)
	return c.AdditionalText
}

func (c *CorporateActionNarrative27) AddNarrativeVersion() *UpdatedAdditionalInformation8 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation8)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative27) AddInformationConditions() *UpdatedAdditionalInformation8 {
	c.InformationConditions = new(UpdatedAdditionalInformation8)
	return c.InformationConditions
}

func (c *CorporateActionNarrative27) AddInformationToComplyWith() *UpdatedAdditionalInformation8 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation8)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative27) AddTaxationConditions() *UpdatedAdditionalInformation8 {
	c.TaxationConditions = new(UpdatedAdditionalInformation8)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative27) AddDisclaimer() *UpdatedAdditionalInformation8 {
	c.Disclaimer = new(UpdatedAdditionalInformation8)
	return c.Disclaimer
}

func (c *CorporateActionNarrative27) AddPartyContactNarrative() *UpdatedAdditionalInformation8 {
	c.PartyContactNarrative = new(UpdatedAdditionalInformation8)
	return c.PartyContactNarrative
}

func (c *CorporateActionNarrative27) AddRegistrationDetails() *UpdatedAdditionalInformation8 {
	c.RegistrationDetails = new(UpdatedAdditionalInformation8)
	return c.RegistrationDetails
}

func (c *CorporateActionNarrative27) AddBasketOrIndexInformation() *UpdatedAdditionalInformation8 {
	c.BasketOrIndexInformation = new(UpdatedAdditionalInformation8)
	return c.BasketOrIndexInformation
}

func (c *CorporateActionNarrative27) AddCertificationBreakdown() *UpdatedAdditionalInformation8 {
	c.CertificationBreakdown = new(UpdatedAdditionalInformation8)
	return c.CertificationBreakdown
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative28 struct {

	// Provides additional information or specifies in more detail the content of a message.  This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation1 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation1 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation1 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation1 `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation1 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation1 `xml:"Dsclmr,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative *UpdatedAdditionalInformation1 `xml:"PtyCtctNrrtv,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *UpdatedAdditionalInformation1 `xml:"RegnDtls,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation *UpdatedAdditionalInformation1 `xml:"BsktOrIndxInf,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *UpdatedAdditionalInformation1 `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative28) AddAdditionalText() *UpdatedAdditionalInformation1 {
	c.AdditionalText = new(UpdatedAdditionalInformation1)
	return c.AdditionalText
}

func (c *CorporateActionNarrative28) AddNarrativeVersion() *UpdatedAdditionalInformation1 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation1)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative28) AddInformationConditions() *UpdatedAdditionalInformation1 {
	c.InformationConditions = new(UpdatedAdditionalInformation1)
	return c.InformationConditions
}

func (c *CorporateActionNarrative28) AddInformationToComplyWith() *UpdatedAdditionalInformation1 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation1)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative28) AddTaxationConditions() *UpdatedAdditionalInformation1 {
	c.TaxationConditions = new(UpdatedAdditionalInformation1)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative28) AddDisclaimer() *UpdatedAdditionalInformation1 {
	c.Disclaimer = new(UpdatedAdditionalInformation1)
	return c.Disclaimer
}

func (c *CorporateActionNarrative28) AddPartyContactNarrative() *UpdatedAdditionalInformation1 {
	c.PartyContactNarrative = new(UpdatedAdditionalInformation1)
	return c.PartyContactNarrative
}

func (c *CorporateActionNarrative28) AddRegistrationDetails() *UpdatedAdditionalInformation1 {
	c.RegistrationDetails = new(UpdatedAdditionalInformation1)
	return c.RegistrationDetails
}

func (c *CorporateActionNarrative28) AddBasketOrIndexInformation() *UpdatedAdditionalInformation1 {
	c.BasketOrIndexInformation = new(UpdatedAdditionalInformation1)
	return c.BasketOrIndexInformation
}

func (c *CorporateActionNarrative28) AddCertificationBreakdown() *UpdatedAdditionalInformation1 {
	c.CertificationBreakdown = new(UpdatedAdditionalInformation1)
	return c.CertificationBreakdown
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative29 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation3 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation3 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation1 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation1 `xml:"InfToCmplyWth,omitempty"`

	// Provides restriction(s) on securities.
	SecurityRestriction *UpdatedAdditionalInformation1 `xml:"SctyRstrctn,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message  and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation1 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation1 `xml:"Dsclmr,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *UpdatedAdditionalInformation1 `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative29) AddAdditionalText() *UpdatedAdditionalInformation3 {
	c.AdditionalText = new(UpdatedAdditionalInformation3)
	return c.AdditionalText
}

func (c *CorporateActionNarrative29) AddNarrativeVersion() *UpdatedAdditionalInformation3 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation3)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative29) AddInformationConditions() *UpdatedAdditionalInformation1 {
	c.InformationConditions = new(UpdatedAdditionalInformation1)
	return c.InformationConditions
}

func (c *CorporateActionNarrative29) AddInformationToComplyWith() *UpdatedAdditionalInformation1 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation1)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative29) AddSecurityRestriction() *UpdatedAdditionalInformation1 {
	c.SecurityRestriction = new(UpdatedAdditionalInformation1)
	return c.SecurityRestriction
}

func (c *CorporateActionNarrative29) AddTaxationConditions() *UpdatedAdditionalInformation1 {
	c.TaxationConditions = new(UpdatedAdditionalInformation1)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative29) AddDisclaimer() *UpdatedAdditionalInformation1 {
	c.Disclaimer = new(UpdatedAdditionalInformation1)
	return c.Disclaimer
}

func (c *CorporateActionNarrative29) AddCertificationBreakdown() *UpdatedAdditionalInformation1 {
	c.CertificationBreakdown = new(UpdatedAdditionalInformation1)
	return c.CertificationBreakdown
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative3 struct {

	// Provides the entity making the offer and is different from the issuing company.
	Offeror *UpdatedAdditionalInformation3 `xml:"Offerr,omitempty"`

	// Provides the new name of a company following a name change.
	NewCompanyName *UpdatedAdditionalInformation3 `xml:"NewCpnyNm,omitempty"`

	// Provides the web address published for the event, that is, the address for the Universal Resource Locator (URL), for example, used over the www (HTTP) service.
	URLAddress *UpdatedURLlnformation `xml:"URLAdr,omitempty"`
}

func (c *CorporateActionNarrative3) AddOfferor() *UpdatedAdditionalInformation3 {
	c.Offeror = new(UpdatedAdditionalInformation3)
	return c.Offeror
}

func (c *CorporateActionNarrative3) AddNewCompanyName() *UpdatedAdditionalInformation3 {
	c.NewCompanyName = new(UpdatedAdditionalInformation3)
	return c.NewCompanyName
}

func (c *CorporateActionNarrative3) AddURLAddress() *UpdatedURLlnformation {
	c.URLAddress = new(UpdatedURLlnformation)
	return c.URLAddress
}

// Provides additional information such as the registration details.
type CorporateActionNarrative30 struct {

	// Provides information required for the registration.
	RegistrationDetails []*Max350Text `xml:"RegnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*Max350Text `xml:"PtyCtctNrrtv,omitempty"`

	// Provides information required for the certification/breakdown.
	CertificationBreakdown []*Max350Text `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative30) AddRegistrationDetails(value string) {
	c.RegistrationDetails = append(c.RegistrationDetails, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative30) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative30) AddCertificationBreakdown(value string) {
	c.CertificationBreakdown = append(c.CertificationBreakdown, (*Max350Text)(&value))
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative31 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText []*Max350Text `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields, - or narrative information not needed for automatic processing.
	NarrativeVersion []*Max350Text `xml:"NrrtvVrsn,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*Max350Text `xml:"PtyCtctNrrtv,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA)
	TaxationConditions []*Max350Text `xml:"TaxtnConds,omitempty"`
}

func (c *CorporateActionNarrative31) AddAdditionalText(value string) {
	c.AdditionalText = append(c.AdditionalText, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative31) AddNarrativeVersion(value string) {
	c.NarrativeVersion = append(c.NarrativeVersion, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative31) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative31) AddTaxationConditions(value string) {
	c.TaxationConditions = append(c.TaxationConditions, (*Max350Text)(&value))
}

// Provides additional information such as the information to comply with.
type CorporateActionNarrative32 struct {

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or Sophisticated Investor Letter (SIL) to be provided.
	InformationToComplyWith []*Max350Text `xml:"InfToCmplyWth,omitempty"`

	// Provides additional information on the delivery details of the outturned (derived) securities. This narrative is only to be used in case the securities are not eligible at the agent/custodian, and may not be used for settlement instructions.
	DeliveryDetails []*Max350Text `xml:"DlvryDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	ForeignExchangeInstructionsAdditionalInformation []*Max350Text `xml:"FXInstrsAddtlInf,omitempty"`

	// Provides additional details pertaining to the corporate action instruction.
	InstructionAdditionalInformation []*Max350Text `xml:"InstrAddtlInf,omitempty"`
}

func (c *CorporateActionNarrative32) AddInformationToComplyWith(value string) {
	c.InformationToComplyWith = append(c.InformationToComplyWith, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative32) AddDeliveryDetails(value string) {
	c.DeliveryDetails = append(c.DeliveryDetails, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative32) AddForeignExchangeInstructionsAdditionalInformation(value string) {
	c.ForeignExchangeInstructionsAdditionalInformation = append(c.ForeignExchangeInstructionsAdditionalInformation, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative32) AddInstructionAdditionalInformation(value string) {
	c.InstructionAdditionalInformation = append(c.InstructionAdditionalInformation, (*Max350Text)(&value))
}

// Provides additional information such as the information to comply with.
type CorporateActionNarrative33 struct {

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or Sophisticated Investor Letter (SIL) to be provided.
	InformationToComplyWith []*RestrictedFINXMax350Text `xml:"InfToCmplyWth,omitempty"`

	// Provides additional information on the delivery details of the outturned (derived) securities. This narrative is only to be used in case the securities are not eligible at the agent/custodian, and may not be used for settlement instructions.
	DeliveryDetails []*RestrictedFINXMax350Text `xml:"DlvryDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	ForeignExchangeInstructionsAdditionalInformation []*RestrictedFINXMax350Text `xml:"FXInstrsAddtlInf,omitempty"`

	// Provides additional details pertaining to the corporate action instruction.
	InstructionAdditionalInformation []*RestrictedFINXMax350Text `xml:"InstrAddtlInf,omitempty"`
}

func (c *CorporateActionNarrative33) AddInformationToComplyWith(value string) {
	c.InformationToComplyWith = append(c.InformationToComplyWith, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative33) AddDeliveryDetails(value string) {
	c.DeliveryDetails = append(c.DeliveryDetails, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative33) AddForeignExchangeInstructionsAdditionalInformation(value string) {
	c.ForeignExchangeInstructionsAdditionalInformation = append(c.ForeignExchangeInstructionsAdditionalInformation, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative33) AddInstructionAdditionalInformation(value string) {
	c.InstructionAdditionalInformation = append(c.InstructionAdditionalInformation, (*RestrictedFINXMax350Text)(&value))
}

// Provides additional information such as the registration details.
type CorporateActionNarrative34 struct {

	// Provides information required for the registration.
	RegistrationDetails []*RestrictedFINXMax350Text `xml:"RegnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*RestrictedFINXMax350Text `xml:"PtyCtctNrrtv,omitempty"`

	// Provides information required for the certification/breakdown.
	CertificationBreakdown []*RestrictedFINXMax350Text `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative34) AddRegistrationDetails(value string) {
	c.RegistrationDetails = append(c.RegistrationDetails, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative34) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative34) AddCertificationBreakdown(value string) {
	c.CertificationBreakdown = append(c.CertificationBreakdown, (*RestrictedFINXMax350Text)(&value))
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative35 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText []*RestrictedFINXMax350Text `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields, - or narrative information not needed for automatic processing.
	NarrativeVersion []*RestrictedFINXMax350Text `xml:"NrrtvVrsn,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*RestrictedFINXMax350Text `xml:"PtyCtctNrrtv,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA)
	TaxationConditions []*RestrictedFINXMax350Text `xml:"TaxtnConds,omitempty"`
}

func (c *CorporateActionNarrative35) AddAdditionalText(value string) {
	c.AdditionalText = append(c.AdditionalText, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative35) AddNarrativeVersion(value string) {
	c.NarrativeVersion = append(c.NarrativeVersion, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative35) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative35) AddTaxationConditions(value string) {
	c.TaxationConditions = append(c.TaxationConditions, (*RestrictedFINXMax350Text)(&value))
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative36 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation6 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation6 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation5 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation5 `xml:"InfToCmplyWth,omitempty"`

	// Provides restriction(s) on securities.
	SecurityRestriction *UpdatedAdditionalInformation5 `xml:"SctyRstrctn,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message  and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation5 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation5 `xml:"Dsclmr,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *UpdatedAdditionalInformation5 `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative36) AddAdditionalText() *UpdatedAdditionalInformation6 {
	c.AdditionalText = new(UpdatedAdditionalInformation6)
	return c.AdditionalText
}

func (c *CorporateActionNarrative36) AddNarrativeVersion() *UpdatedAdditionalInformation6 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation6)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative36) AddInformationConditions() *UpdatedAdditionalInformation5 {
	c.InformationConditions = new(UpdatedAdditionalInformation5)
	return c.InformationConditions
}

func (c *CorporateActionNarrative36) AddInformationToComplyWith() *UpdatedAdditionalInformation5 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation5)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative36) AddSecurityRestriction() *UpdatedAdditionalInformation5 {
	c.SecurityRestriction = new(UpdatedAdditionalInformation5)
	return c.SecurityRestriction
}

func (c *CorporateActionNarrative36) AddTaxationConditions() *UpdatedAdditionalInformation5 {
	c.TaxationConditions = new(UpdatedAdditionalInformation5)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative36) AddDisclaimer() *UpdatedAdditionalInformation5 {
	c.Disclaimer = new(UpdatedAdditionalInformation5)
	return c.Disclaimer
}

func (c *CorporateActionNarrative36) AddCertificationBreakdown() *UpdatedAdditionalInformation5 {
	c.CertificationBreakdown = new(UpdatedAdditionalInformation5)
	return c.CertificationBreakdown
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative37 struct {

	// Provides additional information or specifies in more detail the content of a message.  This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation5 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation5 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation5 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation5 `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation5 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation5 `xml:"Dsclmr,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative *UpdatedAdditionalInformation5 `xml:"PtyCtctNrrtv,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *UpdatedAdditionalInformation5 `xml:"RegnDtls,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation *UpdatedAdditionalInformation5 `xml:"BsktOrIndxInf,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *UpdatedAdditionalInformation5 `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative37) AddAdditionalText() *UpdatedAdditionalInformation5 {
	c.AdditionalText = new(UpdatedAdditionalInformation5)
	return c.AdditionalText
}

func (c *CorporateActionNarrative37) AddNarrativeVersion() *UpdatedAdditionalInformation5 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation5)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative37) AddInformationConditions() *UpdatedAdditionalInformation5 {
	c.InformationConditions = new(UpdatedAdditionalInformation5)
	return c.InformationConditions
}

func (c *CorporateActionNarrative37) AddInformationToComplyWith() *UpdatedAdditionalInformation5 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation5)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative37) AddTaxationConditions() *UpdatedAdditionalInformation5 {
	c.TaxationConditions = new(UpdatedAdditionalInformation5)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative37) AddDisclaimer() *UpdatedAdditionalInformation5 {
	c.Disclaimer = new(UpdatedAdditionalInformation5)
	return c.Disclaimer
}

func (c *CorporateActionNarrative37) AddPartyContactNarrative() *UpdatedAdditionalInformation5 {
	c.PartyContactNarrative = new(UpdatedAdditionalInformation5)
	return c.PartyContactNarrative
}

func (c *CorporateActionNarrative37) AddRegistrationDetails() *UpdatedAdditionalInformation5 {
	c.RegistrationDetails = new(UpdatedAdditionalInformation5)
	return c.RegistrationDetails
}

func (c *CorporateActionNarrative37) AddBasketOrIndexInformation() *UpdatedAdditionalInformation5 {
	c.BasketOrIndexInformation = new(UpdatedAdditionalInformation5)
	return c.BasketOrIndexInformation
}

func (c *CorporateActionNarrative37) AddCertificationBreakdown() *UpdatedAdditionalInformation5 {
	c.CertificationBreakdown = new(UpdatedAdditionalInformation5)
	return c.CertificationBreakdown
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative39 struct {

	// Provides the entity making the offer and is different from the issuing company.
	Offeror []*UpdatedAdditionalInformation6 `xml:"Offerr,omitempty"`

	// Provides the new name of a company following a name change.
	NewCompanyName *UpdatedAdditionalInformation6 `xml:"NewCpnyNm,omitempty"`

	// Provides the web address published for the event, that is, the address for the Universal Resource Locator (URL), for example, used over the www (HTTP) service.
	URLAddress *UpdatedURLlnformation3 `xml:"URLAdr,omitempty"`
}

func (c *CorporateActionNarrative39) AddOfferor() *UpdatedAdditionalInformation6 {
	newValue := new(UpdatedAdditionalInformation6)
	c.Offeror = append(c.Offeror, newValue)
	return newValue
}

func (c *CorporateActionNarrative39) AddNewCompanyName() *UpdatedAdditionalInformation6 {
	c.NewCompanyName = new(UpdatedAdditionalInformation6)
	return c.NewCompanyName
}

func (c *CorporateActionNarrative39) AddURLAddress() *UpdatedURLlnformation3 {
	c.URLAddress = new(UpdatedURLlnformation3)
	return c.URLAddress
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative4 struct {

	// Provides declaration details narrative relative to the financial instrument, for example, beneficial ownership.
	DeclarationDetails []*Max350Text `xml:"DclrtnDtls,omitempty"`

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText []*Max350Text `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion []*Max350Text `xml:"NrrtvVrsn,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails []*Max350Text `xml:"RegnDtls,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions []*Max350Text `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter) to be provided.
	InformationToComplyWith []*Max350Text `xml:"InfToCmplyWth,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*Max350Text `xml:"PtyCtctNrrtv,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA)
	TaxationConditions []*Max350Text `xml:"TaxtnConds,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation []*Max350Text `xml:"BsktOrIndxInf,omitempty"`
}

func (c *CorporateActionNarrative4) AddDeclarationDetails(value string) {
	c.DeclarationDetails = append(c.DeclarationDetails, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddAdditionalText(value string) {
	c.AdditionalText = append(c.AdditionalText, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddNarrativeVersion(value string) {
	c.NarrativeVersion = append(c.NarrativeVersion, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddRegistrationDetails(value string) {
	c.RegistrationDetails = append(c.RegistrationDetails, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddInformationConditions(value string) {
	c.InformationConditions = append(c.InformationConditions, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddInformationToComplyWith(value string) {
	c.InformationToComplyWith = append(c.InformationToComplyWith, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddTaxationConditions(value string) {
	c.TaxationConditions = append(c.TaxationConditions, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative4) AddBasketOrIndexInformation(value string) {
	c.BasketOrIndexInformation = append(c.BasketOrIndexInformation, (*Max350Text)(&value))
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative41 struct {

	// Provides additional information or specifies in more detail the content of a message.  This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation10 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation10 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation10 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation10 `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation10 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation10 `xml:"Dsclmr,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative *UpdatedAdditionalInformation10 `xml:"PtyCtctNrrtv,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *UpdatedAdditionalInformation10 `xml:"RegnDtls,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation *UpdatedAdditionalInformation10 `xml:"BsktOrIndxInf,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *UpdatedAdditionalInformation10 `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative41) AddAdditionalText() *UpdatedAdditionalInformation10 {
	c.AdditionalText = new(UpdatedAdditionalInformation10)
	return c.AdditionalText
}

func (c *CorporateActionNarrative41) AddNarrativeVersion() *UpdatedAdditionalInformation10 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation10)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative41) AddInformationConditions() *UpdatedAdditionalInformation10 {
	c.InformationConditions = new(UpdatedAdditionalInformation10)
	return c.InformationConditions
}

func (c *CorporateActionNarrative41) AddInformationToComplyWith() *UpdatedAdditionalInformation10 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation10)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative41) AddTaxationConditions() *UpdatedAdditionalInformation10 {
	c.TaxationConditions = new(UpdatedAdditionalInformation10)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative41) AddDisclaimer() *UpdatedAdditionalInformation10 {
	c.Disclaimer = new(UpdatedAdditionalInformation10)
	return c.Disclaimer
}

func (c *CorporateActionNarrative41) AddPartyContactNarrative() *UpdatedAdditionalInformation10 {
	c.PartyContactNarrative = new(UpdatedAdditionalInformation10)
	return c.PartyContactNarrative
}

func (c *CorporateActionNarrative41) AddRegistrationDetails() *UpdatedAdditionalInformation10 {
	c.RegistrationDetails = new(UpdatedAdditionalInformation10)
	return c.RegistrationDetails
}

func (c *CorporateActionNarrative41) AddBasketOrIndexInformation() *UpdatedAdditionalInformation10 {
	c.BasketOrIndexInformation = new(UpdatedAdditionalInformation10)
	return c.BasketOrIndexInformation
}

func (c *CorporateActionNarrative41) AddCertificationBreakdown() *UpdatedAdditionalInformation10 {
	c.CertificationBreakdown = new(UpdatedAdditionalInformation10)
	return c.CertificationBreakdown
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative5 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation3 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation3 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation1 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation1 `xml:"InfToCmplyWth,omitempty"`

	// Provides restriction(s) on securities.
	SecurityRestriction *UpdatedAdditionalInformation1 `xml:"SctyRstrctn,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message  and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation1 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation1 `xml:"Dsclmr,omitempty"`
}

func (c *CorporateActionNarrative5) AddAdditionalText() *UpdatedAdditionalInformation3 {
	c.AdditionalText = new(UpdatedAdditionalInformation3)
	return c.AdditionalText
}

func (c *CorporateActionNarrative5) AddNarrativeVersion() *UpdatedAdditionalInformation3 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation3)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative5) AddInformationConditions() *UpdatedAdditionalInformation1 {
	c.InformationConditions = new(UpdatedAdditionalInformation1)
	return c.InformationConditions
}

func (c *CorporateActionNarrative5) AddInformationToComplyWith() *UpdatedAdditionalInformation1 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation1)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative5) AddSecurityRestriction() *UpdatedAdditionalInformation1 {
	c.SecurityRestriction = new(UpdatedAdditionalInformation1)
	return c.SecurityRestriction
}

func (c *CorporateActionNarrative5) AddTaxationConditions() *UpdatedAdditionalInformation1 {
	c.TaxationConditions = new(UpdatedAdditionalInformation1)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative5) AddDisclaimer() *UpdatedAdditionalInformation1 {
	c.Disclaimer = new(UpdatedAdditionalInformation1)
	return c.Disclaimer
}

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative6 struct {

	// Provides additional information or specifies in more detail the content of a message.  This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation1 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation1 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation1 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation1 `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation1 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation1 `xml:"Dsclmr,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative *UpdatedAdditionalInformation1 `xml:"PtyCtctNrrtv,omitempty"`

	// Provides declaration details narrative relative to the financial instrument, for example, beneficial ownership.
	DeclarationDetails *UpdatedAdditionalInformation1 `xml:"DclrtnDtls,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *UpdatedAdditionalInformation1 `xml:"RegnDtls,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation *UpdatedAdditionalInformation1 `xml:"BsktOrIndxInf,omitempty"`
}

func (c *CorporateActionNarrative6) AddAdditionalText() *UpdatedAdditionalInformation1 {
	c.AdditionalText = new(UpdatedAdditionalInformation1)
	return c.AdditionalText
}

func (c *CorporateActionNarrative6) AddNarrativeVersion() *UpdatedAdditionalInformation1 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation1)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative6) AddInformationConditions() *UpdatedAdditionalInformation1 {
	c.InformationConditions = new(UpdatedAdditionalInformation1)
	return c.InformationConditions
}

func (c *CorporateActionNarrative6) AddInformationToComplyWith() *UpdatedAdditionalInformation1 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation1)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative6) AddTaxationConditions() *UpdatedAdditionalInformation1 {
	c.TaxationConditions = new(UpdatedAdditionalInformation1)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative6) AddDisclaimer() *UpdatedAdditionalInformation1 {
	c.Disclaimer = new(UpdatedAdditionalInformation1)
	return c.Disclaimer
}

func (c *CorporateActionNarrative6) AddPartyContactNarrative() *UpdatedAdditionalInformation1 {
	c.PartyContactNarrative = new(UpdatedAdditionalInformation1)
	return c.PartyContactNarrative
}

func (c *CorporateActionNarrative6) AddDeclarationDetails() *UpdatedAdditionalInformation1 {
	c.DeclarationDetails = new(UpdatedAdditionalInformation1)
	return c.DeclarationDetails
}

func (c *CorporateActionNarrative6) AddRegistrationDetails() *UpdatedAdditionalInformation1 {
	c.RegistrationDetails = new(UpdatedAdditionalInformation1)
	return c.RegistrationDetails
}

func (c *CorporateActionNarrative6) AddBasketOrIndexInformation() *UpdatedAdditionalInformation1 {
	c.BasketOrIndexInformation = new(UpdatedAdditionalInformation1)
	return c.BasketOrIndexInformation
}

// Provides additional information such as the registration details.
type CorporateActionNarrative7 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText []*Max350Text `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion []*Max350Text `xml:"NrrtvVrsn,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails []*Max350Text `xml:"RegnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*Max350Text `xml:"PtyCtctNrrtv,omitempty"`

	// Disclaimer relative to the information provided in the message. It may be ignored for automated processing. No information about the instruction itself is allowed here.
	Disclaimer []*Max350Text `xml:"Dsclmr,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation []*Max350Text `xml:"BsktOrIndxInf,omitempty"`
}

func (c *CorporateActionNarrative7) AddAdditionalText(value string) {
	c.AdditionalText = append(c.AdditionalText, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative7) AddNarrativeVersion(value string) {
	c.NarrativeVersion = append(c.NarrativeVersion, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative7) AddRegistrationDetails(value string) {
	c.RegistrationDetails = append(c.RegistrationDetails, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative7) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative7) AddDisclaimer(value string) {
	c.Disclaimer = append(c.Disclaimer, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative7) AddBasketOrIndexInformation(value string) {
	c.BasketOrIndexInformation = append(c.BasketOrIndexInformation, (*Max350Text)(&value))
}

// Provides additional information such as the information to comply with.
type CorporateActionNarrative8 struct {

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or Sophisticated Investor Letter (SIL) to be provided.
	InformationToComplyWith []*Max350Text `xml:"InfToCmplyWth,omitempty"`

	// Provides additional information on the delivery details of the outturned (derived) securities. This narrative is only to be used in case the securities are not eligible at the agent/custodian, and may not be used for settlement instructions.
	DeliveryDetails []*Max350Text `xml:"DlvryDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	ForeignExchangeInstructionsAdditionalInformation []*Max350Text `xml:"FXInstrsAddtlInf,omitempty"`

	// Disclaimer relative to the information provided in the message. It may be ignored for automated processing. No information about the instruction itself is allowed here.
	Disclaimer []*Max350Text `xml:"Dsclmr,omitempty"`

	// Provides additional details pertaining to the corporate action instruction.
	InstructionAdditionalInformation []*Max350Text `xml:"InstrAddtlInf,omitempty"`
}

func (c *CorporateActionNarrative8) AddInformationToComplyWith(value string) {
	c.InformationToComplyWith = append(c.InformationToComplyWith, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative8) AddDeliveryDetails(value string) {
	c.DeliveryDetails = append(c.DeliveryDetails, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative8) AddForeignExchangeInstructionsAdditionalInformation(value string) {
	c.ForeignExchangeInstructionsAdditionalInformation = append(c.ForeignExchangeInstructionsAdditionalInformation, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative8) AddDisclaimer(value string) {
	c.Disclaimer = append(c.Disclaimer, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative8) AddInstructionAdditionalInformation(value string) {
	c.InstructionAdditionalInformation = append(c.InstructionAdditionalInformation, (*Max350Text)(&value))
}
