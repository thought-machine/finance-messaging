package iso20022

// Party and account details.
type PartyIdentificationAndAccount1 struct {

	// Identification of the party.
	Identification *PartyIdentification10Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount1) AddIdentification() *PartyIdentification10Choice {
	p.Identification = new(PartyIdentification10Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount1) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount1) AddSafekeepingAccount() *SecuritiesAccount13 {
	p.SafekeepingAccount = new(SecuritiesAccount13)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount1) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount1) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount1) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount100 struct {

	// Identification of the party.
	Identification *PartyIdentification83Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount100) AddIdentification() *PartyIdentification83Choice {
	p.Identification = new(PartyIdentification83Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount100) AddAlternateIdentification() *AlternatePartyIdentification5 {
	p.AlternateIdentification = new(AlternatePartyIdentification5)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount100) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount100) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount100) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount101 struct {

	// Identification of a party.
	Identification *PartyIdentification38Choice `xml:"Id"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount101) AddIdentification() *PartyIdentification38Choice {
	p.Identification = new(PartyIdentification38Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount101) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount101) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount101) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

func (p *PartyIdentificationAndAccount101) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

// Party involved in the settlement chain.
type PartyIdentificationAndAccount102 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification33Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	SubAccount *SubAccount4 `xml:"SubAcct,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (p *PartyIdentificationAndAccount102) AddPartyIdentification() *PartyIdentification33Choice {
	p.PartyIdentification = new(PartyIdentification33Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount102) SetAccountIdentification(value string) {
	p.AccountIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount102) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount102) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount102) AddSubAccount() *SubAccount4 {
	p.SubAccount = new(SubAccount4)
	return p.SubAccount
}

func (p *PartyIdentificationAndAccount102) AddContactPerson() *ContactIdentification2 {
	p.ContactPerson = new(ContactIdentification2)
	return p.ContactPerson
}

// Party and account details.
type PartyIdentificationAndAccount106 struct {

	// Identification of the party.
	Identification *PartyIdentification71Choice `xml:"Id"`

	// Legal Entity Identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount24 `xml:"SfkpgAcct,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount106) AddIdentification() *PartyIdentification71Choice {
	p.Identification = new(PartyIdentification71Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount106) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount106) AddAlternateIdentification() *AlternatePartyIdentification7 {
	p.AlternateIdentification = new(AlternatePartyIdentification7)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount106) AddSafekeepingAccount() *SecuritiesAccount24 {
	p.SafekeepingAccount = new(SecuritiesAccount24)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount106) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount106) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount106) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount107 struct {

	// Identification of the party.
	Identification *PartyIdentification71Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount107) AddIdentification() *PartyIdentification71Choice {
	p.Identification = new(PartyIdentification71Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount107) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount107) AddAlternateIdentification() *AlternatePartyIdentification7 {
	p.AlternateIdentification = new(AlternatePartyIdentification7)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount107) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount107) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount107) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount108 struct {

	// Identification of the party.
	Identification *PartyIdentification71Choice `xml:"Id,omitempty"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`

	// Nationality of the investor or country of incorporation (for a company).
	Nationality *CountryCode `xml:"Ntlty,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount108) AddIdentification() *PartyIdentification71Choice {
	p.Identification = new(PartyIdentification71Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount108) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount108) AddAlternateIdentification() *AlternatePartyIdentification7 {
	p.AlternateIdentification = new(AlternatePartyIdentification7)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount108) SetNationality(value string) {
	p.Nationality = (*CountryCode)(&value)
}

func (p *PartyIdentificationAndAccount108) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount108) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount108) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount109 struct {

	// Identification of the party.
	Identification *PartyIdentification71Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount109) AddIdentification() *PartyIdentification71Choice {
	p.Identification = new(PartyIdentification71Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount109) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount109) AddAlternateIdentification() *AlternatePartyIdentification7 {
	p.AlternateIdentification = new(AlternatePartyIdentification7)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount109) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount109) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount111 struct {

	// Identification of the party.
	Identification *PartyIdentification71Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification5Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification5Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification5Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation2 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount111) AddIdentification() *PartyIdentification71Choice {
	p.Identification = new(PartyIdentification71Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount111) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount111) AddAlternateIdentification() *AlternatePartyIdentification7 {
	p.AlternateIdentification = new(AlternatePartyIdentification7)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount111) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount111) AddChargesAccount() *CashAccountIdentification5Choice {
	p.ChargesAccount = new(CashAccountIdentification5Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount111) AddCommissionAccount() *CashAccountIdentification5Choice {
	p.CommissionAccount = new(CashAccountIdentification5Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount111) AddTaxAccount() *CashAccountIdentification5Choice {
	p.TaxAccount = new(CashAccountIdentification5Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount111) AddAdditionalInformation() *PartyTextInformation2 {
	p.AdditionalInformation = new(PartyTextInformation2)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount112 struct {

	// Identification of the party.
	Identification *PartyIdentification94Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification5Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification5Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification5Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation2 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount112) AddIdentification() *PartyIdentification94Choice {
	p.Identification = new(PartyIdentification94Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount112) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount112) AddAlternateIdentification() *AlternatePartyIdentification7 {
	p.AlternateIdentification = new(AlternatePartyIdentification7)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount112) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount112) AddChargesAccount() *CashAccountIdentification5Choice {
	p.ChargesAccount = new(CashAccountIdentification5Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount112) AddCommissionAccount() *CashAccountIdentification5Choice {
	p.CommissionAccount = new(CashAccountIdentification5Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount112) AddTaxAccount() *CashAccountIdentification5Choice {
	p.TaxAccount = new(CashAccountIdentification5Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount112) AddAdditionalInformation() *PartyTextInformation2 {
	p.AdditionalInformation = new(PartyTextInformation2)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount117 struct {

	// Identification of the party.
	Identification *PartyIdentification71Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentificationAndAccount117) AddIdentification() *PartyIdentification71Choice {
	p.Identification = new(PartyIdentification71Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount117) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount117) AddSafekeepingAccount() *SecuritiesAccount19 {
	p.SafekeepingAccount = new(SecuritiesAccount19)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount117) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount120 struct {

	// Identification of a party.
	Identification *PartyIdentification71Choice `xml:"Id"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount120) AddIdentification() *PartyIdentification71Choice {
	p.Identification = new(PartyIdentification71Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount120) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount120) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount120) AddAlternateIdentification() *AlternatePartyIdentification7 {
	p.AlternateIdentification = new(AlternatePartyIdentification7)
	return p.AlternateIdentification
}

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount121 struct {

	// Identification of a party.
	Identification *PartyIdentification94Choice `xml:"Id"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount121) AddIdentification() *PartyIdentification94Choice {
	p.Identification = new(PartyIdentification94Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount121) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount121) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount121) AddAlternateIdentification() *AlternatePartyIdentification7 {
	p.AlternateIdentification = new(AlternatePartyIdentification7)
	return p.AlternateIdentification
}

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount122 struct {

	// Identification of a party.
	Identification *PartyIdentification71Choice `xml:"Id"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount122) AddIdentification() *PartyIdentification71Choice {
	p.Identification = new(PartyIdentification71Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount122) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount122) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount122) AddAlternateIdentification() *AlternatePartyIdentification7 {
	newValue := new(AlternatePartyIdentification7)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

// Party involved in the settlement chain.
type PartyIdentificationAndAccount123 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification70Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingReference *Max35Text `xml:"PrcgRef,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	SubAccountDetails *SubAccount5 `xml:"SubAcctDtls,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (p *PartyIdentificationAndAccount123) AddPartyIdentification() *PartyIdentification70Choice {
	p.PartyIdentification = new(PartyIdentification70Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount123) SetAccountIdentification(value string) {
	p.AccountIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount123) SetProcessingReference(value string) {
	p.ProcessingReference = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount123) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount123) AddSubAccountDetails() *SubAccount5 {
	p.SubAccountDetails = new(SubAccount5)
	return p.SubAccountDetails
}

func (p *PartyIdentificationAndAccount123) AddContactPerson() *ContactIdentification2 {
	p.ContactPerson = new(ContactIdentification2)
	return p.ContactPerson
}

// Party involved in the settlement chain.
type PartyIdentificationAndAccount124 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification70Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingReference *Max35Text `xml:"PrcgRef,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`
}

func (p *PartyIdentificationAndAccount124) AddPartyIdentification() *PartyIdentification70Choice {
	p.PartyIdentification = new(PartyIdentification70Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount124) SetAccountIdentification(value string) {
	p.AccountIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount124) SetProcessingReference(value string) {
	p.ProcessingReference = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount124) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

// Party involved in the settlement chain.
type PartyIdentificationAndAccount125 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification70Choice `xml:"PtyId,omitempty"`

	// Identification of the account owned by the party.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlement *PartyIdentification70Choice `xml:"PlcOfSttlm"`
}

func (p *PartyIdentificationAndAccount125) AddPartyIdentification() *PartyIdentification70Choice {
	p.PartyIdentification = new(PartyIdentification70Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount125) SetAccountIdentification(value string) {
	p.AccountIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount125) AddPlaceOfSettlement() *PartyIdentification70Choice {
	p.PlaceOfSettlement = new(PartyIdentification70Choice)
	return p.PlaceOfSettlement
}

// Party involved in the settlement chain.
type PartyIdentificationAndAccount126 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification100Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	SubAccount *SubAccount5 `xml:"SubAcct,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (p *PartyIdentificationAndAccount126) AddPartyIdentification() *PartyIdentification100Choice {
	p.PartyIdentification = new(PartyIdentification100Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount126) SetAccountIdentification(value string) {
	p.AccountIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount126) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount126) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount126) AddSubAccount() *SubAccount5 {
	p.SubAccount = new(SubAccount5)
	return p.SubAccount
}

func (p *PartyIdentificationAndAccount126) AddContactPerson() *ContactIdentification2 {
	p.ContactPerson = new(ContactIdentification2)
	return p.ContactPerson
}

// Party and account details.
type PartyIdentificationAndAccount127 struct {

	// Identification of the party.
	Identification *PartyIdentification101Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification8 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount127) AddIdentification() *PartyIdentification101Choice {
	p.Identification = new(PartyIdentification101Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount127) AddAlternateIdentification() *AlternatePartyIdentification8 {
	p.AlternateIdentification = new(AlternatePartyIdentification8)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount127) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount127) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount127) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount128 struct {

	// Identification of a party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount128) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount128) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (p *PartyIdentificationAndAccount128) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentificationAndAccount128) AddAlternateIdentification() *AlternatePartyIdentification9 {
	newValue := new(AlternatePartyIdentification9)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount129 struct {

	// Identification of a party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount129) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount129) AddCashAccount() *CashAccountIdentification6Choice {
	p.CashAccount = new(CashAccountIdentification6Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount129) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentificationAndAccount129) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount130 struct {

	// Identification of a party.
	Identification *PartyIdentification113Choice `xml:"Id"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount130) AddIdentification() *PartyIdentification113Choice {
	p.Identification = new(PartyIdentification113Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount130) AddCashAccount() *CashAccountIdentification6Choice {
	p.CashAccount = new(CashAccountIdentification6Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount130) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentificationAndAccount130) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

// Party and account details.
type PartyIdentificationAndAccount131 struct {

	// Identification of the party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Legal Entity Identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount27 `xml:"SfkpgAcct,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation3 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount131) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount131) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount131) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount131) AddSafekeepingAccount() *SecuritiesAccount27 {
	p.SafekeepingAccount = new(SecuritiesAccount27)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount131) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount131) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentificationAndAccount131) AddAdditionalInformation() *PartyTextInformation3 {
	p.AdditionalInformation = new(PartyTextInformation3)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount133 struct {

	// Identification of the party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification6Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification6Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification6Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation4 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount133) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount133) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount133) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount133) AddCashAccount() *CashAccountIdentification6Choice {
	p.CashAccount = new(CashAccountIdentification6Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount133) AddChargesAccount() *CashAccountIdentification6Choice {
	p.ChargesAccount = new(CashAccountIdentification6Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount133) AddCommissionAccount() *CashAccountIdentification6Choice {
	p.CommissionAccount = new(CashAccountIdentification6Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount133) AddTaxAccount() *CashAccountIdentification6Choice {
	p.TaxAccount = new(CashAccountIdentification6Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount133) AddAdditionalInformation() *PartyTextInformation4 {
	p.AdditionalInformation = new(PartyTextInformation4)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount134 struct {

	// Identification of the party.
	Identification *PartyIdentification113Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification6Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification6Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification6Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification6Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation4 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount134) AddIdentification() *PartyIdentification113Choice {
	p.Identification = new(PartyIdentification113Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount134) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount134) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount134) AddCashAccount() *CashAccountIdentification6Choice {
	p.CashAccount = new(CashAccountIdentification6Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount134) AddChargesAccount() *CashAccountIdentification6Choice {
	p.ChargesAccount = new(CashAccountIdentification6Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount134) AddCommissionAccount() *CashAccountIdentification6Choice {
	p.CommissionAccount = new(CashAccountIdentification6Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount134) AddTaxAccount() *CashAccountIdentification6Choice {
	p.TaxAccount = new(CashAccountIdentification6Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount134) AddAdditionalInformation() *PartyTextInformation4 {
	p.AdditionalInformation = new(PartyTextInformation4)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount135 struct {

	// Identification of the party.
	Identification *PartyIdentification104Choice `xml:"Id,omitempty"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Nationality of the investor or country of incorporation (for a company).
	Nationality *CountryCode `xml:"Ntlty,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation3 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount135) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount135) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount135) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount135) SetNationality(value string) {
	p.Nationality = (*CountryCode)(&value)
}

func (p *PartyIdentificationAndAccount135) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (p *PartyIdentificationAndAccount135) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentificationAndAccount135) AddAdditionalInformation() *PartyTextInformation3 {
	p.AdditionalInformation = new(PartyTextInformation3)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount136 struct {

	// Identification of the party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation3 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount136) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount136) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount136) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount136) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (p *PartyIdentificationAndAccount136) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentificationAndAccount136) AddAdditionalInformation() *PartyTextInformation3 {
	p.AdditionalInformation = new(PartyTextInformation3)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount137 struct {

	// Identification of the party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation3 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount137) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount137) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount137) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount137) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentificationAndAccount137) AddAdditionalInformation() *PartyTextInformation3 {
	p.AdditionalInformation = new(PartyTextInformation3)
	return p.AdditionalInformation
}

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount14 struct {

	// Identification of a party.
	Identification *PartyIdentification10Choice `xml:"Id"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount14) AddIdentification() *PartyIdentification10Choice {
	p.Identification = new(PartyIdentification10Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount14) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount14) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount14) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

func (p *PartyIdentificationAndAccount14) AddAlternateIdentification() *AlternatePartyIdentification2 {
	newValue := new(AlternatePartyIdentification2)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

// Party and account details.
type PartyIdentificationAndAccount146 struct {

	// Identification of the party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentificationAndAccount146) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount146) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount146) AddSafekeepingAccount() *SecuritiesAccount30 {
	p.SafekeepingAccount = new(SecuritiesAccount30)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount146) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Party and account information.
type PartyIdentificationAndAccount147 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification113 `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`
}

func (p *PartyIdentificationAndAccount147) AddPartyIdentification() *PartyIdentification113 {
	p.PartyIdentification = new(PartyIdentification113)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount147) SetAccountIdentification(value string) {
	p.AccountIdentification = (*Max35Text)(&value)
}

// Party and account details.
type PartyIdentificationAndAccount15 struct {

	// Identification of the party.
	Identification *PartyIdentification30Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification5Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification5Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification5Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation2 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount15) AddIdentification() *PartyIdentification30Choice {
	p.Identification = new(PartyIdentification30Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount15) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount15) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount15) AddChargesAccount() *CashAccountIdentification5Choice {
	p.ChargesAccount = new(CashAccountIdentification5Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount15) AddCommissionAccount() *CashAccountIdentification5Choice {
	p.CommissionAccount = new(CashAccountIdentification5Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount15) AddTaxAccount() *CashAccountIdentification5Choice {
	p.TaxAccount = new(CashAccountIdentification5Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount15) AddAdditionalInformation() *PartyTextInformation2 {
	p.AdditionalInformation = new(PartyTextInformation2)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount16 struct {

	// Identification of the party.
	Identification *PartyIdentification10Choice `xml:"Id"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentificationAndAccount16) AddIdentification() *PartyIdentification10Choice {
	p.Identification = new(PartyIdentification10Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount16) AddSafekeepingAccount() *SecuritiesAccount13 {
	p.SafekeepingAccount = new(SecuritiesAccount13)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount16) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount17 struct {

	// Identification of a party.
	Identification *PartyIdentification10Choice `xml:"Id"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount17) AddIdentification() *PartyIdentification10Choice {
	p.Identification = new(PartyIdentification10Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount17) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount17) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount17) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

func (p *PartyIdentificationAndAccount17) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount18 struct {

	// Identification of a party.
	Identification *PartyIdentification30Choice `xml:"Id"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount18) AddIdentification() *PartyIdentification30Choice {
	p.Identification = new(PartyIdentification30Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount18) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount18) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount18) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

func (p *PartyIdentificationAndAccount18) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

// Party and account details.
type PartyIdentificationAndAccount19 struct {

	// Identification of the party.
	Identification *PartyIdentification10Choice `xml:"Id,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount19) AddIdentification() *PartyIdentification10Choice {
	p.Identification = new(PartyIdentification10Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount19) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount19) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentificationAndAccount19) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount19) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount19) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party involved in the settlement chain.
type PartyIdentificationAndAccount2 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification1Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingReference *Max35Text `xml:"PrcgRef,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`
}

func (p *PartyIdentificationAndAccount2) AddPartyIdentification() *PartyIdentification1Choice {
	p.PartyIdentification = new(PartyIdentification1Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount2) AddAccountIdentification() *AccountIdentification1 {
	p.AccountIdentification = new(AccountIdentification1)
	return p.AccountIdentification
}

func (p *PartyIdentificationAndAccount2) SetProcessingReference(value string) {
	p.ProcessingReference = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount2) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

// Party and account details.
type PartyIdentificationAndAccount20 struct {

	// Identification of the party.
	Identification *PartyIdentification10Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification5Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification5Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification5Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation2 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount20) AddIdentification() *PartyIdentification10Choice {
	p.Identification = new(PartyIdentification10Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount20) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount20) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount20) AddChargesAccount() *CashAccountIdentification5Choice {
	p.ChargesAccount = new(CashAccountIdentification5Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount20) AddCommissionAccount() *CashAccountIdentification5Choice {
	p.CommissionAccount = new(CashAccountIdentification5Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount20) AddTaxAccount() *CashAccountIdentification5Choice {
	p.TaxAccount = new(CashAccountIdentification5Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount20) AddAdditionalInformation() *PartyTextInformation2 {
	p.AdditionalInformation = new(PartyTextInformation2)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount21 struct {

	// Identification of the party.
	Identification *PartyIdentification10Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount21) AddIdentification() *PartyIdentification10Choice {
	p.Identification = new(PartyIdentification10Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount21) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount21) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount21) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount21) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account information.
type PartyIdentificationAndAccount3 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification2Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`
}

func (p *PartyIdentificationAndAccount3) AddPartyIdentification() *PartyIdentification2Choice {
	p.PartyIdentification = new(PartyIdentification2Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount3) AddAccountIdentification() *AccountIdentification1 {
	p.AccountIdentification = new(AccountIdentification1)
	return p.AccountIdentification
}

// Party and account details.
type PartyIdentificationAndAccount31 struct {

	// Identification of the party.
	Identification *PartyIdentification33Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification4 `xml:"AltrnId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`

	// Identifies the clearing member account at the Central counterparty through which the trade must be cleared (sometimes called position account).
	ClearingAccount *SecuritiesAccount18 `xml:"ClrAcct,omitempty"`
}

func (p *PartyIdentificationAndAccount31) AddIdentification() *PartyIdentification33Choice {
	p.Identification = new(PartyIdentification33Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount31) AddAlternateIdentification() *AlternatePartyIdentification4 {
	p.AlternateIdentification = new(AlternatePartyIdentification4)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount31) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

func (p *PartyIdentificationAndAccount31) AddClearingAccount() *SecuritiesAccount18 {
	p.ClearingAccount = new(SecuritiesAccount18)
	return p.ClearingAccount
}

// Party involved in the settlement chain.
type PartyIdentificationAndAccount32 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification33Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (p *PartyIdentificationAndAccount32) AddPartyIdentification() *PartyIdentification33Choice {
	p.PartyIdentification = new(PartyIdentification33Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount32) SetAccountIdentification(value string) {
	p.AccountIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount32) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount32) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount32) AddSubAccountDetails() *SubAccount1 {
	p.SubAccountDetails = new(SubAccount1)
	return p.SubAccountDetails
}

func (p *PartyIdentificationAndAccount32) AddContactPerson() *ContactIdentification2 {
	p.ContactPerson = new(ContactIdentification2)
	return p.ContactPerson
}

// Entity involved in an activity.
type PartyIdentificationAndAccount34 struct {

	// Identification of a party.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`
}

func (p *PartyIdentificationAndAccount34) AddIdentification() *PartyIdentification32Choice {
	p.Identification = new(PartyIdentification32Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount34) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

func (p *PartyIdentificationAndAccount34) AddAlternateIdentification() *AlternatePartyIdentification5 {
	p.AlternateIdentification = new(AlternatePartyIdentification5)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount34) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

// Party and account details.
type PartyIdentificationAndAccount35 struct {

	// Identification of the party.
	Identification *PartyIdentification49Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount35) AddIdentification() *PartyIdentification49Choice {
	p.Identification = new(PartyIdentification49Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount35) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount35) AddSafekeepingAccount() *SecuritiesAccount13 {
	p.SafekeepingAccount = new(SecuritiesAccount13)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount35) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount35) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount35) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount36 struct {

	// Identification of the party.
	Identification *PartyIdentification49Choice `xml:"Id,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount36) AddIdentification() *PartyIdentification49Choice {
	p.Identification = new(PartyIdentification49Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount36) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount36) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentificationAndAccount36) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount36) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount36) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount37 struct {

	// Identification of the party.
	Identification *PartyIdentification49Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount37) AddIdentification() *PartyIdentification49Choice {
	p.Identification = new(PartyIdentification49Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount37) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount37) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount37) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount37) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount38 struct {

	// Identification of the party.
	Identification *PartyIdentification49Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification5Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification5Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification5Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation2 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount38) AddIdentification() *PartyIdentification49Choice {
	p.Identification = new(PartyIdentification49Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount38) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount38) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount38) AddChargesAccount() *CashAccountIdentification5Choice {
	p.ChargesAccount = new(CashAccountIdentification5Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount38) AddCommissionAccount() *CashAccountIdentification5Choice {
	p.CommissionAccount = new(CashAccountIdentification5Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount38) AddTaxAccount() *CashAccountIdentification5Choice {
	p.TaxAccount = new(CashAccountIdentification5Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount38) AddAdditionalInformation() *PartyTextInformation2 {
	p.AdditionalInformation = new(PartyTextInformation2)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount39 struct {

	// Identification of the party.
	Identification *PartyIdentification43Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification5Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification5Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification5Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation2 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount39) AddIdentification() *PartyIdentification43Choice {
	p.Identification = new(PartyIdentification43Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount39) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount39) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount39) AddChargesAccount() *CashAccountIdentification5Choice {
	p.ChargesAccount = new(CashAccountIdentification5Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount39) AddCommissionAccount() *CashAccountIdentification5Choice {
	p.CommissionAccount = new(CashAccountIdentification5Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount39) AddTaxAccount() *CashAccountIdentification5Choice {
	p.TaxAccount = new(CashAccountIdentification5Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount39) AddAdditionalInformation() *PartyTextInformation2 {
	p.AdditionalInformation = new(PartyTextInformation2)
	return p.AdditionalInformation
}

// Party involved in the settlement chain.
type PartyIdentificationAndAccount4 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification2Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingReference *Max35Text `xml:"PrcgRef,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (p *PartyIdentificationAndAccount4) AddPartyIdentification() *PartyIdentification2Choice {
	p.PartyIdentification = new(PartyIdentification2Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount4) AddAccountIdentification() *AccountIdentification1 {
	p.AccountIdentification = new(AccountIdentification1)
	return p.AccountIdentification
}

func (p *PartyIdentificationAndAccount4) SetProcessingReference(value string) {
	p.ProcessingReference = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount4) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount4) AddSubAccountDetails() *SubAccount1 {
	p.SubAccountDetails = new(SubAccount1)
	return p.SubAccountDetails
}

func (p *PartyIdentificationAndAccount4) AddContactPerson() *ContactIdentification2 {
	p.ContactPerson = new(ContactIdentification2)
	return p.ContactPerson
}

// Party and account details.
type PartyIdentificationAndAccount40 struct {

	// Identification of the party.
	Identification *PartyIdentification43Choice `xml:"Id,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount40) AddIdentification() *PartyIdentification43Choice {
	p.Identification = new(PartyIdentification43Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount40) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount40) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentificationAndAccount40) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount40) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount40) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount41 struct {

	// Identification of the party.
	Identification *PartyIdentification43Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount41) AddIdentification() *PartyIdentification43Choice {
	p.Identification = new(PartyIdentification43Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount41) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount41) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount41) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount41) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount42 struct {

	// Identification of the party.
	Identification *PartyIdentification43Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount42) AddIdentification() *PartyIdentification43Choice {
	p.Identification = new(PartyIdentification43Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount42) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount42) AddSafekeepingAccount() *SecuritiesAccount13 {
	p.SafekeepingAccount = new(SecuritiesAccount13)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount42) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount42) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount42) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount43 struct {

	// Identification of the party.
	Identification *PartyIdentification43Choice `xml:"Id"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentificationAndAccount43) AddIdentification() *PartyIdentification43Choice {
	p.Identification = new(PartyIdentification43Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount43) AddSafekeepingAccount() *SecuritiesAccount13 {
	p.SafekeepingAccount = new(SecuritiesAccount13)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount43) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

// Party and account details.
type PartyIdentificationAndAccount44 struct {

	// Identification of the party.
	Identification *PartyIdentification45Choice `xml:"Id"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentificationAndAccount44) AddIdentification() *PartyIdentification45Choice {
	p.Identification = new(PartyIdentification45Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount44) AddSafekeepingAccount() *SecuritiesAccount13 {
	p.SafekeepingAccount = new(SecuritiesAccount13)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount44) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

// Party and account details.
type PartyIdentificationAndAccount45 struct {

	// Identification of the party.
	Identification *PartyIdentification45Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount45) AddIdentification() *PartyIdentification45Choice {
	p.Identification = new(PartyIdentification45Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount45) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount45) AddSafekeepingAccount() *SecuritiesAccount13 {
	p.SafekeepingAccount = new(SecuritiesAccount13)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount45) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount45) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount45) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount46 struct {

	// Identification of the party.
	Identification *PartyIdentification45Choice `xml:"Id,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount46) AddIdentification() *PartyIdentification45Choice {
	p.Identification = new(PartyIdentification45Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount46) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount46) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentificationAndAccount46) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount46) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount46) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount47 struct {

	// Identification of the party.
	Identification *PartyIdentification45Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount47) AddIdentification() *PartyIdentification45Choice {
	p.Identification = new(PartyIdentification45Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount47) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount47) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount47) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount47) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount48 struct {

	// Identification of the party.
	Identification *PartyIdentification45Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification5Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification5Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification5Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation2 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount48) AddIdentification() *PartyIdentification45Choice {
	p.Identification = new(PartyIdentification45Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount48) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount48) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount48) AddChargesAccount() *CashAccountIdentification5Choice {
	p.ChargesAccount = new(CashAccountIdentification5Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount48) AddCommissionAccount() *CashAccountIdentification5Choice {
	p.CommissionAccount = new(CashAccountIdentification5Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount48) AddTaxAccount() *CashAccountIdentification5Choice {
	p.TaxAccount = new(CashAccountIdentification5Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount48) AddAdditionalInformation() *PartyTextInformation2 {
	p.AdditionalInformation = new(PartyTextInformation2)
	return p.AdditionalInformation
}

// Party involved in the settlement chain.
type PartyIdentificationAndAccount5 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification2Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingReference *Max35Text `xml:"PrcgRef,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`
}

func (p *PartyIdentificationAndAccount5) AddPartyIdentification() *PartyIdentification2Choice {
	p.PartyIdentification = new(PartyIdentification2Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount5) AddAccountIdentification() *AccountIdentification1 {
	p.AccountIdentification = new(AccountIdentification1)
	return p.AccountIdentification
}

func (p *PartyIdentificationAndAccount5) SetProcessingReference(value string) {
	p.ProcessingReference = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount5) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

// Party and account details.
type PartyIdentificationAndAccount50 struct {

	// Identification of the party.
	Identification *PartyIdentification38Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification5Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification5Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification5Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation2 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount50) AddIdentification() *PartyIdentification38Choice {
	p.Identification = new(PartyIdentification38Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount50) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount50) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount50) AddChargesAccount() *CashAccountIdentification5Choice {
	p.ChargesAccount = new(CashAccountIdentification5Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount50) AddCommissionAccount() *CashAccountIdentification5Choice {
	p.CommissionAccount = new(CashAccountIdentification5Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount50) AddTaxAccount() *CashAccountIdentification5Choice {
	p.TaxAccount = new(CashAccountIdentification5Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount50) AddAdditionalInformation() *PartyTextInformation2 {
	p.AdditionalInformation = new(PartyTextInformation2)
	return p.AdditionalInformation
}

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount51 struct {

	// Identification of a party.
	Identification *PartyIdentification40Choice `xml:"Id"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount51) AddIdentification() *PartyIdentification40Choice {
	p.Identification = new(PartyIdentification40Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount51) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount51) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount51) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

func (p *PartyIdentificationAndAccount51) AddAlternateIdentification() *AlternatePartyIdentification2 {
	newValue := new(AlternatePartyIdentification2)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount52 struct {

	// Identification of a party.
	Identification *PartyIdentification40Choice `xml:"Id"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount52) AddIdentification() *PartyIdentification40Choice {
	p.Identification = new(PartyIdentification40Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount52) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount52) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount52) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

func (p *PartyIdentificationAndAccount52) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

// Party and account details.
type PartyIdentificationAndAccount53 struct {

	// Identification of the party.
	Identification *PartyIdentification40Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification5Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification5Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification5Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation2 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount53) AddIdentification() *PartyIdentification40Choice {
	p.Identification = new(PartyIdentification40Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount53) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount53) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount53) AddChargesAccount() *CashAccountIdentification5Choice {
	p.ChargesAccount = new(CashAccountIdentification5Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount53) AddCommissionAccount() *CashAccountIdentification5Choice {
	p.CommissionAccount = new(CashAccountIdentification5Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount53) AddTaxAccount() *CashAccountIdentification5Choice {
	p.TaxAccount = new(CashAccountIdentification5Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount53) AddAdditionalInformation() *PartyTextInformation2 {
	p.AdditionalInformation = new(PartyTextInformation2)
	return p.AdditionalInformation
}

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount54 struct {

	// Identification of a party.
	Identification *PartyIdentification42Choice `xml:"Id"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount54) AddIdentification() *PartyIdentification42Choice {
	p.Identification = new(PartyIdentification42Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount54) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount54) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount54) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

func (p *PartyIdentificationAndAccount54) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

// Party and account details.
type PartyIdentificationAndAccount55 struct {

	// Identification of the party.
	Identification *PartyIdentification42Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification5Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification5Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification5Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation2 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount55) AddIdentification() *PartyIdentification42Choice {
	p.Identification = new(PartyIdentification42Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount55) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount55) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount55) AddChargesAccount() *CashAccountIdentification5Choice {
	p.ChargesAccount = new(CashAccountIdentification5Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount55) AddCommissionAccount() *CashAccountIdentification5Choice {
	p.CommissionAccount = new(CashAccountIdentification5Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount55) AddTaxAccount() *CashAccountIdentification5Choice {
	p.TaxAccount = new(CashAccountIdentification5Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount55) AddAdditionalInformation() *PartyTextInformation2 {
	p.AdditionalInformation = new(PartyTextInformation2)
	return p.AdditionalInformation
}

// Entity involved in an activity.
type PartyIdentificationAndAccount6 struct {

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	PartyIdentification *PartyIdentification25 `xml:"PtyId"`

	// Unambiguous identification of an account held by Financing Requestor to First Agent. This account is requested to be used for crediting the amount financed, as a result of the financing process.
	CreditAccount *CashAccount7 `xml:"CdtAcct,omitempty"`

	// Unambiguous identification of an internal bank account used by First Agent to manage the line of credit granted to Financing Requestor. This account is requested to be used for managing the financing process.
	FinancingAccount *CashAccount7 `xml:"FincgAcct,omitempty"`
}

func (p *PartyIdentificationAndAccount6) AddPartyIdentification() *PartyIdentification25 {
	p.PartyIdentification = new(PartyIdentification25)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount6) AddCreditAccount() *CashAccount7 {
	p.CreditAccount = new(CashAccount7)
	return p.CreditAccount
}

func (p *PartyIdentificationAndAccount6) AddFinancingAccount() *CashAccount7 {
	p.FinancingAccount = new(CashAccount7)
	return p.FinancingAccount
}

// Party and account details.
type PartyIdentificationAndAccount77 struct {

	// Identification of the party.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount77) AddIdentification() *PartyIdentification32Choice {
	p.Identification = new(PartyIdentification32Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount77) AddAlternateIdentification() *AlternatePartyIdentification5 {
	p.AlternateIdentification = new(AlternatePartyIdentification5)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount77) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount77) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount77) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount78 struct {

	// Identification of the party.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Coded list to specify the side of the order.
	Side *ClearingSide1Code `xml:"Sd,omitempty"`

	// Identifies the clearing member account at the CCP through which the trade must be cleared (sometimes called position account).
	ClearingAccount *SecuritiesAccount20 `xml:"ClrAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount78) AddIdentification() *PartyIdentification32Choice {
	p.Identification = new(PartyIdentification32Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount78) AddAlternateIdentification() *AlternatePartyIdentification5 {
	p.AlternateIdentification = new(AlternatePartyIdentification5)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount78) SetSide(value string) {
	p.Side = (*ClearingSide1Code)(&value)
}

func (p *PartyIdentificationAndAccount78) AddClearingAccount() *SecuritiesAccount20 {
	p.ClearingAccount = new(SecuritiesAccount20)
	return p.ClearingAccount
}

func (p *PartyIdentificationAndAccount78) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount78) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Entity involved in an activity.
type PartyIdentificationAndAccount79 struct {

	// Information related to an identification, eg, party identification or account identification.
	Identification *PartyIdentification32Choice `xml:"Id,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification2Choice `xml:"CshAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount79) AddIdentification() *PartyIdentification32Choice {
	p.Identification = new(PartyIdentification32Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount79) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount79) AddCashAccount() *CashAccountIdentification2Choice {
	p.CashAccount = new(CashAccountIdentification2Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount79) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount79) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentificationAndAccount79) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

func (p *PartyIdentificationAndAccount79) AddAlternateIdentification() *AlternatePartyIdentification5 {
	p.AlternateIdentification = new(AlternatePartyIdentification5)
	return p.AlternateIdentification
}

// Party and account details.
type PartyIdentificationAndAccount80 struct {

	// Identification of the party.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Account to or from which a cash entry is made.
	CashAccount *CashAccountIdentification5Choice `xml:"CshAcct,omitempty"`

	// Account to be used for charges/fees if different from the account for payment.
	ChargesAccount *CashAccountIdentification5Choice `xml:"ChrgsAcct,omitempty"`

	// Account to be used for commission if different from the account for payment.
	CommissionAccount *CashAccountIdentification5Choice `xml:"ComssnAcct,omitempty"`

	// Account to be used for taxes if different from the account for payment.
	TaxAccount *CashAccountIdentification5Choice `xml:"TaxAcct,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation2 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount80) AddIdentification() *PartyIdentification32Choice {
	p.Identification = new(PartyIdentification32Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount80) AddAlternateIdentification() *AlternatePartyIdentification5 {
	p.AlternateIdentification = new(AlternatePartyIdentification5)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount80) AddCashAccount() *CashAccountIdentification5Choice {
	p.CashAccount = new(CashAccountIdentification5Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount80) AddChargesAccount() *CashAccountIdentification5Choice {
	p.ChargesAccount = new(CashAccountIdentification5Choice)
	return p.ChargesAccount
}

func (p *PartyIdentificationAndAccount80) AddCommissionAccount() *CashAccountIdentification5Choice {
	p.CommissionAccount = new(CashAccountIdentification5Choice)
	return p.CommissionAccount
}

func (p *PartyIdentificationAndAccount80) AddTaxAccount() *CashAccountIdentification5Choice {
	p.TaxAccount = new(CashAccountIdentification5Choice)
	return p.TaxAccount
}

func (p *PartyIdentificationAndAccount80) AddAdditionalInformation() *PartyTextInformation2 {
	p.AdditionalInformation = new(PartyTextInformation2)
	return p.AdditionalInformation
}

// Party and account details.
type PartyIdentificationAndAccount81 struct {

	// Identification of the party.
	Identification *PartyIdentification43Choice `xml:"Id,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Nationality of the investor or country of incorporation (for a company).
	Nationality *CountryCode `xml:"Ntlty,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount81) AddIdentification() *PartyIdentification43Choice {
	p.Identification = new(PartyIdentification43Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount81) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount81) SetNationality(value string) {
	p.Nationality = (*CountryCode)(&value)
}

func (p *PartyIdentificationAndAccount81) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount81) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount81) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Entity involved in an activity.
type PartyIdentificationAndAccount83 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification70Choice `xml:"Id"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Account in which cash is maintained.
	CashAccount *CashAccountIdentification2Choice `xml:"CshAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`

	// Entity involved in an activity.
	AlternateIdentification *AlternatePartyIdentification6 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount83) AddIdentification() *PartyIdentification70Choice {
	p.Identification = new(PartyIdentification70Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount83) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount83) AddCashAccount() *CashAccountIdentification2Choice {
	p.CashAccount = new(CashAccountIdentification2Choice)
	return p.CashAccount
}

func (p *PartyIdentificationAndAccount83) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount83) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

func (p *PartyIdentificationAndAccount83) AddAlternateIdentification() *AlternatePartyIdentification6 {
	p.AlternateIdentification = new(AlternatePartyIdentification6)
	return p.AlternateIdentification
}

// Party and account details.
type PartyIdentificationAndAccount86 struct {

	// Identification of the party.
	Identification *PartyIdentification43Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount86) AddIdentification() *PartyIdentification43Choice {
	p.Identification = new(PartyIdentification43Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount86) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount86) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount86) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

// Entity involved in an activity.
type PartyIdentificationAndAccount87 struct {

	// Unique and unambiguous way to identify an organisation.
	Identification *PartyIdentification70Choice `xml:"Id"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`

	// Entity involved in an activity.
	AlternateIdentification *AlternatePartyIdentification6 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount87) AddIdentification() *PartyIdentification70Choice {
	p.Identification = new(PartyIdentification70Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount87) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount87) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

func (p *PartyIdentificationAndAccount87) AddAlternateIdentification() *AlternatePartyIdentification6 {
	p.AlternateIdentification = new(AlternatePartyIdentification6)
	return p.AlternateIdentification
}

// Party involved in the settlement chain.
type PartyIdentificationAndAccount93 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification2Choice `xml:"PtyId,omitempty"`

	// Identification of the account owned by the party.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlement *PartyIdentification2Choice `xml:"PlcOfSttlm"`
}

func (p *PartyIdentificationAndAccount93) AddPartyIdentification() *PartyIdentification2Choice {
	p.PartyIdentification = new(PartyIdentification2Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount93) AddAccountIdentification() *AccountIdentification1 {
	p.AccountIdentification = new(AccountIdentification1)
	return p.AccountIdentification
}

func (p *PartyIdentificationAndAccount93) AddPlaceOfSettlement() *PartyIdentification2Choice {
	p.PlaceOfSettlement = new(PartyIdentification2Choice)
	return p.PlaceOfSettlement
}

// Party and account details.
type PartyIdentificationAndAccount95 struct {

	// Identification of the party.
	PartyIdentification *PartyIdentification71Choice `xml:"PtyId"`

	// Account to or from which a securities entry is made.
	AccountIdentification *SecuritiesAccount22 `xml:"AcctId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentificationAndAccount95) AddPartyIdentification() *PartyIdentification71Choice {
	p.PartyIdentification = new(PartyIdentification71Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount95) AddAccountIdentification() *SecuritiesAccount22 {
	p.AccountIdentification = new(SecuritiesAccount22)
	return p.AccountIdentification
}

func (p *PartyIdentificationAndAccount95) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

// Party and account information.
type PartyIdentificationAndAccount96 struct {

	// Identification of the party that legally owns the account.
	PartyIdentification *PartyIdentification64 `xml:"PtyId"`

	// Identification of the account.
	AccountIdentification *AccountIdentification26 `xml:"AcctId"`
}

func (p *PartyIdentificationAndAccount96) AddPartyIdentification() *PartyIdentification64 {
	p.PartyIdentification = new(PartyIdentification64)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount96) AddAccountIdentification() *AccountIdentification26 {
	p.AccountIdentification = new(AccountIdentification26)
	return p.AccountIdentification
}

// Party and account information.
type PartyIdentificationAndAccount97 struct {

	// Identification of the party that legally owns the account.
	PartyIdentification *PartyIdentification62 `xml:"PtyId"`

	// Identification of the account.
	AccountIdentification *AccountIdentification26 `xml:"AcctId,omitempty"`
}

func (p *PartyIdentificationAndAccount97) AddPartyIdentification() *PartyIdentification62 {
	p.PartyIdentification = new(PartyIdentification62)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount97) AddAccountIdentification() *AccountIdentification26 {
	p.AccountIdentification = new(AccountIdentification26)
	return p.AccountIdentification
}
