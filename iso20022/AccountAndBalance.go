package iso20022

// Provides account and balance information.
type AccountAndBalance1 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails1 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance1) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance1) AddAccountOwner() *PartyIdentification13Choice {
	a.AccountOwner = new(PartyIdentification13Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance1) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance1) AddBalance() *CorporateActionBalanceDetails1 {
	a.Balance = new(CorporateActionBalanceDetails1)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance10 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails3 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance10) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance10) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance10) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance10) AddBalance() *CorporateActionBalanceDetails3 {
	a.Balance = new(CorporateActionBalanceDetails3)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance11 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails2 `xml:"Bal"`
}

func (a *AccountAndBalance11) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance11) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance11) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance11) AddBalance() *CorporateActionBalanceDetails2 {
	a.Balance = new(CorporateActionBalanceDetails2)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance15 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails10 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance15) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance15) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance15) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance15) AddBalance() *CorporateActionBalanceDetails10 {
	a.Balance = new(CorporateActionBalanceDetails10)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance16 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails11 `xml:"Bal"`
}

func (a *AccountAndBalance16) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance16) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance16) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance16) AddBalance() *CorporateActionBalanceDetails11 {
	a.Balance = new(CorporateActionBalanceDetails11)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance17 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails12 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance17) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance17) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance17) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance17) AddBalance() *CorporateActionBalanceDetails12 {
	a.Balance = new(CorporateActionBalanceDetails12)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance2 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails3 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance2) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance2) AddAccountOwner() *PartyIdentification13Choice {
	a.AccountOwner = new(PartyIdentification13Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance2) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance2) AddBalance() *CorporateActionBalanceDetails3 {
	a.Balance = new(CorporateActionBalanceDetails3)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance21 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails17 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance21) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance21) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance21) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance21) AddBalance() *CorporateActionBalanceDetails17 {
	a.Balance = new(CorporateActionBalanceDetails17)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance22 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails18 `xml:"Bal"`
}

func (a *AccountAndBalance22) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance22) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance22) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance22) AddBalance() *CorporateActionBalanceDetails18 {
	a.Balance = new(CorporateActionBalanceDetails18)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance25 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails21 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance25) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance25) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance25) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance25) AddBalance() *CorporateActionBalanceDetails21 {
	a.Balance = new(CorporateActionBalanceDetails21)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance26 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails22 `xml:"Bal"`
}

func (a *AccountAndBalance26) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance26) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance26) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance26) AddBalance() *CorporateActionBalanceDetails22 {
	a.Balance = new(CorporateActionBalanceDetails22)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance3 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails2 `xml:"Bal"`
}

func (a *AccountAndBalance3) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance3) AddAccountOwner() *PartyIdentification13Choice {
	a.AccountOwner = new(PartyIdentification13Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance3) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance3) AddBalance() *CorporateActionBalanceDetails2 {
	a.Balance = new(CorporateActionBalanceDetails2)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance33 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails29 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance33) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance33) AddAccountOwner() *PartyIdentification92Choice {
	a.AccountOwner = new(PartyIdentification92Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance33) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance33) AddBalance() *CorporateActionBalanceDetails29 {
	a.Balance = new(CorporateActionBalanceDetails29)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance34 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails31 `xml:"Bal"`
}

func (a *AccountAndBalance34) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance34) AddAccountOwner() *PartyIdentification92Choice {
	a.AccountOwner = new(PartyIdentification92Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance34) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance34) AddBalance() *CorporateActionBalanceDetails31 {
	a.Balance = new(CorporateActionBalanceDetails31)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance35 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails32 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance35) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance35) AddAccountOwner() *PartyIdentification92Choice {
	a.AccountOwner = new(PartyIdentification92Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance35) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance35) AddBalance() *CorporateActionBalanceDetails32 {
	a.Balance = new(CorporateActionBalanceDetails32)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance36 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat5Choice `xml:"ConfdBal"`
}

func (a *AccountAndBalance36) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance36) AddConfirmedBalance() *BalanceFormat5Choice {
	a.ConfirmedBalance = new(BalanceFormat5Choice)
	return a.ConfirmedBalance
}

// Provides account and balance information.
type AccountAndBalance37 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat11Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails34 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance37) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountAndBalance37) AddAccountOwner() *PartyIdentification103Choice {
	a.AccountOwner = new(PartyIdentification103Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance37) AddSafekeepingPlace() *SafekeepingPlaceFormat11Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat11Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance37) AddBalance() *CorporateActionBalanceDetails34 {
	a.Balance = new(CorporateActionBalanceDetails34)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance38 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat11Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails35 `xml:"Bal"`
}

func (a *AccountAndBalance38) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountAndBalance38) AddAccountOwner() *PartyIdentification103Choice {
	a.AccountOwner = new(PartyIdentification103Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance38) AddSafekeepingPlace() *SafekeepingPlaceFormat11Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat11Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance38) AddBalance() *CorporateActionBalanceDetails35 {
	a.Balance = new(CorporateActionBalanceDetails35)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance39 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat11Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails36 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance39) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountAndBalance39) AddAccountOwner() *PartyIdentification103Choice {
	a.AccountOwner = new(PartyIdentification103Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance39) AddSafekeepingPlace() *SafekeepingPlaceFormat11Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat11Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance39) AddBalance() *CorporateActionBalanceDetails36 {
	a.Balance = new(CorporateActionBalanceDetails36)
	return a.Balance
}

// Provides account and balance information.
type AccountAndBalance4 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat1Choice `xml:"ConfdBal"`
}

func (a *AccountAndBalance4) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance4) AddConfirmedBalance() *BalanceFormat1Choice {
	a.ConfirmedBalance = new(BalanceFormat1Choice)
	return a.ConfirmedBalance
}

// Provides account and balance information.
type AccountAndBalance40 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat7Choice `xml:"ConfdBal"`
}

func (a *AccountAndBalance40) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountAndBalance40) AddConfirmedBalance() *BalanceFormat7Choice {
	a.ConfirmedBalance = new(BalanceFormat7Choice)
	return a.ConfirmedBalance
}

// Provides account and balance information.
type AccountAndBalance9 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails1 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance9) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance9) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance9) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance9) AddBalance() *CorporateActionBalanceDetails1 {
	a.Balance = new(CorporateActionBalanceDetails1)
	return a.Balance
}
