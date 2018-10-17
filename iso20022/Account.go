package iso20022

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account1 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id,omitempty"`

	// Institution servicing an account and assigning the account identifier to the account owner.
	AccountServicer *PartyIdentification1Choice `xml:"AcctSvcr"`
}

func (a *Account1) AddIdentification() *AccountIdentification1 {
	a.Identification = new(AccountIdentification1)
	return a.Identification
}

func (a *Account1) AddAccountServicer() *PartyIdentification1Choice {
	a.AccountServicer = new(PartyIdentification1Choice)
	return a.AccountServicer
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account11 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification49Choice `xml:"AcctSvcr,omitempty"`
}

func (a *Account11) AddIdentification() *AccountIdentification1 {
	a.Identification = new(AccountIdentification1)
	return a.Identification
}

func (a *Account11) AddAccountServicer() *PartyIdentification49Choice {
	a.AccountServicer = new(PartyIdentification49Choice)
	return a.AccountServicer
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account13 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification26 `xml:"Id"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification49Choice `xml:"AcctSvcr,omitempty"`
}

func (a *Account13) AddIdentification() *AccountIdentification26 {
	a.Identification = new(AccountIdentification26)
	return a.Identification
}

func (a *Account13) AddAccountServicer() *PartyIdentification49Choice {
	a.AccountServicer = new(PartyIdentification49Choice)
	return a.AccountServicer
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account14 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Institution that maintains the records where the account is held.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (a *Account14) AddIdentification() *AccountIdentification1 {
	a.Identification = new(AccountIdentification1)
	return a.Identification
}

func (a *Account14) AddAccountServicer() *PartyIdentification2Choice {
	a.AccountServicer = new(PartyIdentification2Choice)
	return a.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type Account15 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Institution that maintains the records where the account is held.
	Servicer *PartyIdentification2Choice `xml:"Svcr"`
}

func (a *Account15) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Account15) SetDesignation(value string) {
	a.Designation = (*Max35Text)(&value)
}

func (a *Account15) AddServicer() *PartyIdentification2Choice {
	a.Servicer = new(PartyIdentification2Choice)
	return a.Servicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type Account16 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Institution that maintains the records where the account is held.
	Servicer *PartyIdentification2Choice `xml:"Svcr,omitempty"`
}

func (a *Account16) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Account16) SetDesignation(value string) {
	a.Designation = (*Max35Text)(&value)
}

func (a *Account16) AddServicer() *PartyIdentification2Choice {
	a.Servicer = new(PartyIdentification2Choice)
	return a.Servicer
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account18 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification26 `xml:"Id"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification71Choice `xml:"AcctSvcr,omitempty"`
}

func (a *Account18) AddIdentification() *AccountIdentification26 {
	a.Identification = new(AccountIdentification26)
	return a.Identification
}

func (a *Account18) AddAccountServicer() *PartyIdentification71Choice {
	a.AccountServicer = new(PartyIdentification71Choice)
	return a.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type Account19 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Institution that maintains the records where the account is held.
	Servicer *PartyIdentification70Choice `xml:"Svcr,omitempty"`
}

func (a *Account19) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Account19) SetDesignation(value string) {
	a.Designation = (*Max35Text)(&value)
}

func (a *Account19) AddServicer() *PartyIdentification70Choice {
	a.Servicer = new(PartyIdentification70Choice)
	return a.Servicer
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account2 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id,omitempty"`

	// Institution servicing an account and assigning the account identifier to the account owner.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr"`
}

func (a *Account2) AddIdentification() *AccountIdentification1 {
	a.Identification = new(AccountIdentification1)
	return a.Identification
}

func (a *Account2) AddAccountServicer() *PartyIdentification2Choice {
	a.AccountServicer = new(PartyIdentification2Choice)
	return a.AccountServicer
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account20 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Institution servicing an account and assigning the account identifier to the account owner.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr"`
}

func (a *Account20) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Account20) AddAccountServicer() *PartyIdentification70Choice {
	a.AccountServicer = new(PartyIdentification70Choice)
	return a.AccountServicer
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account21 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4 `xml:"Id"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification104Choice `xml:"AcctSvcr,omitempty"`
}

func (a *Account21) AddIdentification() *AccountIdentification4 {
	a.Identification = new(AccountIdentification4)
	return a.Identification
}

func (a *Account21) AddAccountServicer() *PartyIdentification104Choice {
	a.AccountServicer = new(PartyIdentification104Choice)
	return a.AccountServicer
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account22 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Institution servicing the account and assigning the account identifier to the account owner.
	AccountServicer *PartyIdentification113 `xml:"AcctSvcr,omitempty"`
}

func (a *Account22) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Account22) AddAccountServicer() *PartyIdentification113 {
	a.AccountServicer = new(PartyIdentification113)
	return a.AccountServicer
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account23 struct {

	// Identification of the account.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Information about the account to which the existing account is to be linked.
	RelatedAccountDetails *GenericIdentification1 `xml:"RltdAcctDtls,omitempty"`
}

func (a *Account23) SetAccountIdentification(value string) {
	a.AccountIdentification = (*Max35Text)(&value)
}

func (a *Account23) AddRelatedAccountDetails() *GenericIdentification1 {
	a.RelatedAccountDetails = new(GenericIdentification1)
	return a.RelatedAccountDetails
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type Account5 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification2Choice `xml:"Svcr"`
}

func (a *Account5) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Account5) SetDesignation(value string) {
	a.Designation = (*Max35Text)(&value)
}

func (a *Account5) AddServicer() *PartyIdentification2Choice {
	a.Servicer = new(PartyIdentification2Choice)
	return a.Servicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type Account6 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification2Choice `xml:"Svcr,omitempty"`
}

func (a *Account6) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Account6) SetDesignation(value string) {
	a.Designation = (*Max35Text)(&value)
}

func (a *Account6) AddServicer() *PartyIdentification2Choice {
	a.Servicer = new(PartyIdentification2Choice)
	return a.Servicer
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account7 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Institution servicing an account and assigning the account identifier to the account owner.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (a *Account7) AddIdentification() *AccountIdentification1 {
	a.Identification = new(AccountIdentification1)
	return a.Identification
}

func (a *Account7) AddAccountServicer() *PartyIdentification2Choice {
	a.AccountServicer = new(PartyIdentification2Choice)
	return a.AccountServicer
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account9 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification10Choice `xml:"AcctSvcr,omitempty"`
}

func (a *Account9) AddIdentification() *AccountIdentification1 {
	a.Identification = new(AccountIdentification1)
	return a.Identification
}

func (a *Account9) AddAccountServicer() *PartyIdentification10Choice {
	a.AccountServicer = new(PartyIdentification10Choice)
	return a.AccountServicer
}
