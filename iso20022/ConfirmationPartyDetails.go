package iso20022

// Parties used for acting parties that apply either to the whole message or to individual sides.
type ConfirmationPartyDetails1 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation5 `xml:"AddtlInf,omitempty"`
}

func (c *ConfirmationPartyDetails1) AddIdentification() *PartyIdentification32Choice {
	c.Identification = new(PartyIdentification32Choice)
	return c.Identification
}

func (c *ConfirmationPartyDetails1) AddAlternateIdentification() *AlternatePartyIdentification5 {
	c.AlternateIdentification = new(AlternatePartyIdentification5)
	return c.AlternateIdentification
}

func (c *ConfirmationPartyDetails1) SetProcessingIdentification(value string) {
	c.ProcessingIdentification = (*Max35Text)(&value)
}

func (c *ConfirmationPartyDetails1) AddAdditionalInformation() *PartyTextInformation5 {
	c.AdditionalInformation = new(PartyTextInformation5)
	return c.AdditionalInformation
}

// Parties used for acting parties that apply either to the whole message or to individual sides.
type ConfirmationPartyDetails2 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation5 `xml:"AddtlInf,omitempty"`

	// Capacity of customer placing the order. Primarily used by futures exchanges to indicate the CTI code (customer type indicator) as required by the US CFTC (Commodity Futures Trading Commission).
	InvestorCapacity *InvestorCapacity3Choice `xml:"InvstrCpcty,omitempty"`

	// Capacity of customer placing the order. Primarily used by futures exchanges to indicate the CTI code (customer type indicator) as required by the US CFTC (Commodity Futures Trading Commission).
	TradingPartyCapacity *TradingPartyCapacity1Choice `xml:"TradgPtyCpcty,omitempty"`
}

func (c *ConfirmationPartyDetails2) AddIdentification() *PartyIdentification32Choice {
	c.Identification = new(PartyIdentification32Choice)
	return c.Identification
}

func (c *ConfirmationPartyDetails2) AddAlternateIdentification() *AlternatePartyIdentification5 {
	c.AlternateIdentification = new(AlternatePartyIdentification5)
	return c.AlternateIdentification
}

func (c *ConfirmationPartyDetails2) SetProcessingIdentification(value string) {
	c.ProcessingIdentification = (*Max35Text)(&value)
}

func (c *ConfirmationPartyDetails2) AddAdditionalInformation() *PartyTextInformation5 {
	c.AdditionalInformation = new(PartyTextInformation5)
	return c.AdditionalInformation
}

func (c *ConfirmationPartyDetails2) AddInvestorCapacity() *InvestorCapacity3Choice {
	c.InvestorCapacity = new(InvestorCapacity3Choice)
	return c.InvestorCapacity
}

func (c *ConfirmationPartyDetails2) AddTradingPartyCapacity() *TradingPartyCapacity1Choice {
	c.TradingPartyCapacity = new(TradingPartyCapacity1Choice)
	return c.TradingPartyCapacity
}

// Parties used for acting parties that apply either to the whole message or to individual sides.
type ConfirmationPartyDetails3 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount3 `xml:"SfkpgAcct,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	CashDetails *AccountIdentification3Choice `xml:"CshDtls,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation5 `xml:"AddtlInf,omitempty"`

	// Capacity of customer placing the order. Primarily used by futures exchanges to indicate the CTI code (customer type indicator) as required by the US CFTC (Commodity Futures Trading Commission).
	PartyCapacity *TradingPartyCapacity2Choice `xml:"PtyCpcty,omitempty"`
}

func (c *ConfirmationPartyDetails3) AddIdentification() *PartyIdentification32Choice {
	c.Identification = new(PartyIdentification32Choice)
	return c.Identification
}

func (c *ConfirmationPartyDetails3) AddSafekeepingAccount() *SecuritiesAccount3 {
	c.SafekeepingAccount = new(SecuritiesAccount3)
	return c.SafekeepingAccount
}

func (c *ConfirmationPartyDetails3) AddCashDetails() *AccountIdentification3Choice {
	c.CashDetails = new(AccountIdentification3Choice)
	return c.CashDetails
}

func (c *ConfirmationPartyDetails3) AddAlternateIdentification() *AlternatePartyIdentification5 {
	c.AlternateIdentification = new(AlternatePartyIdentification5)
	return c.AlternateIdentification
}

func (c *ConfirmationPartyDetails3) SetProcessingIdentification(value string) {
	c.ProcessingIdentification = (*Max35Text)(&value)
}

func (c *ConfirmationPartyDetails3) AddAdditionalInformation() *PartyTextInformation5 {
	c.AdditionalInformation = new(PartyTextInformation5)
	return c.AdditionalInformation
}

func (c *ConfirmationPartyDetails3) AddPartyCapacity() *TradingPartyCapacity2Choice {
	c.PartyCapacity = new(TradingPartyCapacity2Choice)
	return c.PartyCapacity
}

// Parties used for acting parties that apply either to the whole message or to individual sides.
type ConfirmationPartyDetails4 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount3 `xml:"SfkpgAcct,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	CashDetails *AccountIdentification3Choice `xml:"CshDtls,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation5 `xml:"AddtlInf,omitempty"`
}

func (c *ConfirmationPartyDetails4) AddIdentification() *PartyIdentification32Choice {
	c.Identification = new(PartyIdentification32Choice)
	return c.Identification
}

func (c *ConfirmationPartyDetails4) AddSafekeepingAccount() *SecuritiesAccount3 {
	c.SafekeepingAccount = new(SecuritiesAccount3)
	return c.SafekeepingAccount
}

func (c *ConfirmationPartyDetails4) AddCashDetails() *AccountIdentification3Choice {
	c.CashDetails = new(AccountIdentification3Choice)
	return c.CashDetails
}

func (c *ConfirmationPartyDetails4) AddAlternateIdentification() *AlternatePartyIdentification5 {
	c.AlternateIdentification = new(AlternatePartyIdentification5)
	return c.AlternateIdentification
}

func (c *ConfirmationPartyDetails4) SetProcessingIdentification(value string) {
	c.ProcessingIdentification = (*Max35Text)(&value)
}

func (c *ConfirmationPartyDetails4) AddAdditionalInformation() *PartyTextInformation5 {
	c.AdditionalInformation = new(PartyTextInformation5)
	return c.AdditionalInformation
}

// Parties used for acting parties that apply either to the whole message or to individual sides.
type ConfirmationPartyDetails5 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation5 `xml:"AddtlInf,omitempty"`

	// Indicates whether the confirmation party is a member of the investor protection association required, eg, as per regulation.
	InvestorProtectionAssociationMembership *YesNoIndicator `xml:"InvstrPrtcnAssoctnMmbsh,omitempty"`
}

func (c *ConfirmationPartyDetails5) AddIdentification() *PartyIdentification32Choice {
	c.Identification = new(PartyIdentification32Choice)
	return c.Identification
}

func (c *ConfirmationPartyDetails5) AddAlternateIdentification() *AlternatePartyIdentification5 {
	c.AlternateIdentification = new(AlternatePartyIdentification5)
	return c.AlternateIdentification
}

func (c *ConfirmationPartyDetails5) SetProcessingIdentification(value string) {
	c.ProcessingIdentification = (*Max35Text)(&value)
}

func (c *ConfirmationPartyDetails5) AddAdditionalInformation() *PartyTextInformation5 {
	c.AdditionalInformation = new(PartyTextInformation5)
	return c.AdditionalInformation
}

func (c *ConfirmationPartyDetails5) SetInvestorProtectionAssociationMembership(value string) {
	c.InvestorProtectionAssociationMembership = (*YesNoIndicator)(&value)
}

// Parties used for acting parties that apply either to the whole message or to individual sides.
type ConfirmationPartyDetails6 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount3 `xml:"SfkpgAcct,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	CashDetails *AccountIdentification3Choice `xml:"CshDtls,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation5 `xml:"AddtlInf,omitempty"`

	// Capacity of customer placing the order. Primarily used by futures exchanges to indicate the CTI code (customer type indicator) as required by the US CFTC (Commodity Futures Trading Commission).
	PartyCapacity *TradingPartyCapacity2Choice `xml:"PtyCpcty,omitempty"`

	// Indicates whether the confirmation party is a member of the investor protection association required, eg, as per regulation.
	InvestorProtectionAssociationMembership *YesNoIndicator `xml:"InvstrPrtcnAssoctnMmbsh,omitempty"`
}

func (c *ConfirmationPartyDetails6) AddIdentification() *PartyIdentification32Choice {
	c.Identification = new(PartyIdentification32Choice)
	return c.Identification
}

func (c *ConfirmationPartyDetails6) AddSafekeepingAccount() *SecuritiesAccount3 {
	c.SafekeepingAccount = new(SecuritiesAccount3)
	return c.SafekeepingAccount
}

func (c *ConfirmationPartyDetails6) AddCashDetails() *AccountIdentification3Choice {
	c.CashDetails = new(AccountIdentification3Choice)
	return c.CashDetails
}

func (c *ConfirmationPartyDetails6) AddAlternateIdentification() *AlternatePartyIdentification5 {
	c.AlternateIdentification = new(AlternatePartyIdentification5)
	return c.AlternateIdentification
}

func (c *ConfirmationPartyDetails6) SetProcessingIdentification(value string) {
	c.ProcessingIdentification = (*Max35Text)(&value)
}

func (c *ConfirmationPartyDetails6) AddAdditionalInformation() *PartyTextInformation5 {
	c.AdditionalInformation = new(PartyTextInformation5)
	return c.AdditionalInformation
}

func (c *ConfirmationPartyDetails6) AddPartyCapacity() *TradingPartyCapacity2Choice {
	c.PartyCapacity = new(TradingPartyCapacity2Choice)
	return c.PartyCapacity
}

func (c *ConfirmationPartyDetails6) SetInvestorProtectionAssociationMembership(value string) {
	c.InvestorProtectionAssociationMembership = (*YesNoIndicator)(&value)
}
