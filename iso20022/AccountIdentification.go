package iso20022

// Unique identifier of an account, as assigned by the account servicer.
type AccountIdentification1 struct {

	// Unique identifier for an account. It is assigned by the account servicer using a proprietary identification scheme.
	Proprietary *SimpleIdentificationInformation `xml:"Prtry"`
}

func (a *AccountIdentification1) AddProprietary() *SimpleIdentificationInformation {
	a.Proprietary = new(SimpleIdentificationInformation)
	return a.Proprietary
}

// Provides account identification information.
type AccountIdentification10 struct {

	// Standard code to specify that announcement applies to all safekeeping accounts that own underlying financial instrument.
	IdentificationCode *SafekeepingAccountIdentification1Code `xml:"IdCd"`
}

func (a *AccountIdentification10) SetIdentificationCode(value string) {
	a.IdentificationCode = (*SafekeepingAccountIdentification1Code)(&value)
}

// Provides account identification information.
type AccountIdentification15 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`
}

func (a *AccountIdentification15) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification15) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountIdentification15) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

// Account information and detailed account holdings information report for corporate action events.
type AccountIdentification16 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Detailed account holdings information report for a corporate action event.
	CorporateActionEventAndBalance []*CorporateActionEventAndBalance3 `xml:"CorpActnEvtAndBal,omitempty"`
}

func (a *AccountIdentification16) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification16) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountIdentification16) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification16) AddCorporateActionEventAndBalance() *CorporateActionEventAndBalance3 {
	newValue := new(CorporateActionEventAndBalance3)
	a.CorporateActionEventAndBalance = append(a.CorporateActionEventAndBalance, newValue)
	return newValue
}

// Provides account identification information.
type AccountIdentification17 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`
}

func (a *AccountIdentification17) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification17) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountIdentification17) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

// Provides account identification information.
type AccountIdentification18 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat1Choice `xml:"ConfdBal"`
}

func (a *AccountIdentification18) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification18) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountIdentification18) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification18) AddConfirmedBalance() *BalanceFormat1Choice {
	a.ConfirmedBalance = new(BalanceFormat1Choice)
	return a.ConfirmedBalance
}

// Account information and detailed account holdings information report for corporate action events.
type AccountIdentification23 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Detailed account holdings information report for a corporate action event.
	CorporateActionEventAndBalance []*CorporateActionEventAndBalance5 `xml:"CorpActnEvtAndBal,omitempty"`
}

func (a *AccountIdentification23) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification23) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountIdentification23) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification23) AddCorporateActionEventAndBalance() *CorporateActionEventAndBalance5 {
	newValue := new(CorporateActionEventAndBalance5)
	a.CorporateActionEventAndBalance = append(a.CorporateActionEventAndBalance, newValue)
	return newValue
}

// Account information and detailed account holdings information report for corporate action events.
type AccountIdentification25 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Detailed account holdings information report for a corporate action event.
	CorporateActionEventAndBalance []*CorporateActionEventAndBalance7 `xml:"CorpActnEvtAndBal,omitempty"`
}

func (a *AccountIdentification25) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification25) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountIdentification25) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification25) AddCorporateActionEventAndBalance() *CorporateActionEventAndBalance7 {
	newValue := new(CorporateActionEventAndBalance7)
	a.CorporateActionEventAndBalance = append(a.CorporateActionEventAndBalance, newValue)
	return newValue
}

// Unique identifier of an account, as assigned by the account servicer.
type AccountIdentification26 struct {

	// Unique identifier for an account. It is assigned by the account servicer using a proprietary identification scheme.
	Proprietary *SimpleIdentificationInformation4 `xml:"Prtry"`
}

func (a *AccountIdentification26) AddProprietary() *SimpleIdentificationInformation4 {
	a.Proprietary = new(SimpleIdentificationInformation4)
	return a.Proprietary
}

// Identification of the account expressed with a data source scheme.
type AccountIdentification3 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Entity that assigns the information.
	Issuer *Max8Text `xml:"Issr"`

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Information *Exact4AlphaNumericText `xml:"Inf"`
}

func (a *AccountIdentification3) AddIdentification() *AccountIdentification1 {
	a.Identification = new(AccountIdentification1)
	return a.Identification
}

func (a *AccountIdentification3) SetIssuer(value string) {
	a.Issuer = (*Max8Text)(&value)
}

func (a *AccountIdentification3) SetInformation(value string) {
	a.Information = (*Exact4AlphaNumericText)(&value)
}

// Provides account identification information.
type AccountIdentification31 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`
}

func (a *AccountIdentification31) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification31) AddAccountOwner() *PartyIdentification92Choice {
	a.AccountOwner = new(PartyIdentification92Choice)
	return a.AccountOwner
}

func (a *AccountIdentification31) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return a.SafekeepingPlace
}

// Provides account identification information.
type AccountIdentification32 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat5Choice `xml:"ConfdBal"`
}

func (a *AccountIdentification32) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification32) AddAccountOwner() *PartyIdentification92Choice {
	a.AccountOwner = new(PartyIdentification92Choice)
	return a.AccountOwner
}

func (a *AccountIdentification32) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification32) AddConfirmedBalance() *BalanceFormat5Choice {
	a.ConfirmedBalance = new(BalanceFormat5Choice)
	return a.ConfirmedBalance
}

// Account information and detailed account holdings information report for corporate action events.
type AccountIdentification33 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Detailed account holdings information report for a corporate action event.
	CorporateActionEventAndBalance []*CorporateActionEventAndBalance9 `xml:"CorpActnEvtAndBal,omitempty"`
}

func (a *AccountIdentification33) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification33) AddAccountOwner() *PartyIdentification92Choice {
	a.AccountOwner = new(PartyIdentification92Choice)
	return a.AccountOwner
}

func (a *AccountIdentification33) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification33) AddCorporateActionEventAndBalance() *CorporateActionEventAndBalance9 {
	newValue := new(CorporateActionEventAndBalance9)
	a.CorporateActionEventAndBalance = append(a.CorporateActionEventAndBalance, newValue)
	return newValue
}

// Provides account identification information.
type AccountIdentification34 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat11Choice `xml:"SfkpgPlc,omitempty"`
}

func (a *AccountIdentification34) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountIdentification34) AddAccountOwner() *PartyIdentification103Choice {
	a.AccountOwner = new(PartyIdentification103Choice)
	return a.AccountOwner
}

func (a *AccountIdentification34) AddSafekeepingPlace() *SafekeepingPlaceFormat11Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat11Choice)
	return a.SafekeepingPlace
}

// Account information and detailed account holdings information report for corporate action events.
type AccountIdentification35 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat11Choice `xml:"SfkpgPlc,omitempty"`

	// Detailed account holdings information report for a corporate action event.
	CorporateActionEventAndBalance []*CorporateActionEventAndBalance10 `xml:"CorpActnEvtAndBal,omitempty"`
}

func (a *AccountIdentification35) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountIdentification35) AddAccountOwner() *PartyIdentification103Choice {
	a.AccountOwner = new(PartyIdentification103Choice)
	return a.AccountOwner
}

func (a *AccountIdentification35) AddSafekeepingPlace() *SafekeepingPlaceFormat11Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat11Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification35) AddCorporateActionEventAndBalance() *CorporateActionEventAndBalance10 {
	newValue := new(CorporateActionEventAndBalance10)
	a.CorporateActionEventAndBalance = append(a.CorporateActionEventAndBalance, newValue)
	return newValue
}

// Provides account identification information.
type AccountIdentification37 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat11Choice `xml:"SfkpgPlc,omitempty"`

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat7Choice `xml:"ConfdBal"`
}

func (a *AccountIdentification37) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountIdentification37) AddAccountOwner() *PartyIdentification103Choice {
	a.AccountOwner = new(PartyIdentification103Choice)
	return a.AccountOwner
}

func (a *AccountIdentification37) AddSafekeepingPlace() *SafekeepingPlaceFormat11Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat11Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification37) AddConfirmedBalance() *BalanceFormat7Choice {
	a.ConfirmedBalance = new(BalanceFormat7Choice)
	return a.ConfirmedBalance
}

// Provides account identification information.
type AccountIdentification39 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification95Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat26Choice `xml:"SfkpgPlc,omitempty"`
}

func (a *AccountIdentification39) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountIdentification39) AddAccountOwner() *PartyIdentification95Choice {
	a.AccountOwner = new(PartyIdentification95Choice)
	return a.AccountOwner
}

func (a *AccountIdentification39) AddSafekeepingPlace() *SafekeepingPlaceFormat26Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat26Choice)
	return a.SafekeepingPlace
}

// Unique identifier of an account, as assigned by the account servicer.
type AccountIdentification4 struct {

	// Unique identifier for an account. It is assigned by the account servicer using a proprietary identification scheme.
	Proprietary *SimpleIdentificationInformation1 `xml:"Prtry"`
}

func (a *AccountIdentification4) AddProprietary() *SimpleIdentificationInformation1 {
	a.Proprietary = new(SimpleIdentificationInformation1)
	return a.Proprietary
}

// Account information and detailed account holdings information report for corporate action events.
type AccountIdentification40 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Detailed account holdings information report for a corporate action event.
	CorporateActionEventAndBalance []*CorporateActionEventAndBalance11 `xml:"CorpActnEvtAndBal,omitempty"`
}

func (a *AccountIdentification40) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification40) AddAccountOwner() *PartyIdentification92Choice {
	a.AccountOwner = new(PartyIdentification92Choice)
	return a.AccountOwner
}

func (a *AccountIdentification40) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification40) AddCorporateActionEventAndBalance() *CorporateActionEventAndBalance11 {
	newValue := new(CorporateActionEventAndBalance11)
	a.CorporateActionEventAndBalance = append(a.CorporateActionEventAndBalance, newValue)
	return newValue
}

// Account information and detailed account holdings information report for corporate action events.
type AccountIdentification41 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat11Choice `xml:"SfkpgPlc,omitempty"`

	// Detailed account holdings information report for a corporate action event.
	CorporateActionEventAndBalance []*CorporateActionEventAndBalance12 `xml:"CorpActnEvtAndBal,omitempty"`
}

func (a *AccountIdentification41) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountIdentification41) AddAccountOwner() *PartyIdentification103Choice {
	a.AccountOwner = new(PartyIdentification103Choice)
	return a.AccountOwner
}

func (a *AccountIdentification41) AddSafekeepingPlace() *SafekeepingPlaceFormat11Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat11Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification41) AddCorporateActionEventAndBalance() *CorporateActionEventAndBalance12 {
	newValue := new(CorporateActionEventAndBalance12)
	a.CorporateActionEventAndBalance = append(a.CorporateActionEventAndBalance, newValue)
	return newValue
}

// Provides account identification information.
type AccountIdentification42 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat27Choice `xml:"SfkpgPlc,omitempty"`
}

func (a *AccountIdentification42) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountIdentification42) AddAccountOwner() *PartyIdentification103Choice {
	a.AccountOwner = new(PartyIdentification103Choice)
	return a.AccountOwner
}

func (a *AccountIdentification42) AddSafekeepingPlace() *SafekeepingPlaceFormat27Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat27Choice)
	return a.SafekeepingPlace
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type AccountIdentification5 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Description of the account.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Specifies the type of account.
	Type *GenericIdentification30 `xml:"Tp,omitempty"`
}

func (a *AccountIdentification5) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountIdentification5) SetName(value string) {
	a.Name = (*Max35Text)(&value)
}

func (a *AccountIdentification5) AddType() *GenericIdentification30 {
	a.Type = new(GenericIdentification30)
	return a.Type
}

// Account information and detailed account holdings information report for corporate action events.
type AccountIdentification6 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Detailed account holdings information report for a corporate action event.
	CorporateActionEventAndBalance []*CorporateActionEventAndBalance1 `xml:"CorpActnEvtAndBal,omitempty"`
}

func (a *AccountIdentification6) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification6) AddAccountOwner() *PartyIdentification13Choice {
	a.AccountOwner = new(PartyIdentification13Choice)
	return a.AccountOwner
}

func (a *AccountIdentification6) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification6) AddCorporateActionEventAndBalance() *CorporateActionEventAndBalance1 {
	newValue := new(CorporateActionEventAndBalance1)
	a.CorporateActionEventAndBalance = append(a.CorporateActionEventAndBalance, newValue)
	return newValue
}

// Provides account identification information.
type AccountIdentification7 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat1Choice `xml:"ConfdBal"`
}

func (a *AccountIdentification7) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification7) AddAccountOwner() *PartyIdentification13Choice {
	a.AccountOwner = new(PartyIdentification13Choice)
	return a.AccountOwner
}

func (a *AccountIdentification7) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification7) AddConfirmedBalance() *BalanceFormat1Choice {
	a.ConfirmedBalance = new(BalanceFormat1Choice)
	return a.ConfirmedBalance
}

// Provides account identification information.
type AccountIdentification8 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`
}

func (a *AccountIdentification8) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification8) AddAccountOwner() *PartyIdentification13Choice {
	a.AccountOwner = new(PartyIdentification13Choice)
	return a.AccountOwner
}

func (a *AccountIdentification8) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

// Provides account identification information.
type AccountIdentification9 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`
}

func (a *AccountIdentification9) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification9) AddAccountOwner() *PartyIdentification13Choice {
	a.AccountOwner = new(PartyIdentification13Choice)
	return a.AccountOwner
}

func (a *AccountIdentification9) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}
