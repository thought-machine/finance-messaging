package iso20022

// Information about a party's account.
type AccountParties10 struct {

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties5Choice `xml:"PrncplAcctPty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation10 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation10 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation10 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation10 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation10 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation10 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty7 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation10 `xml:"Grntr,omitempty"`

	// Entity that creates a trust or contributes assets to the trust.
	Settlor []*InvestmentAccountOwnershipInformation10 `xml:"Sttlr,omitempty"`
}

func (a *AccountParties10) AddPrincipalAccountParty() *AccountParties5Choice {
	a.PrincipalAccountParty = new(AccountParties5Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties10) AddSecondaryOwner() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties10) AddBeneficiary() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties10) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties10) AddLegalGuardian() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties10) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties10) AddAdministrator() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties10) AddOtherParty() *ExtendedParty7 {
	newValue := new(ExtendedParty7)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties10) AddGranter() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties10) AddSettlor() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.Settlor = append(a.Settlor, newValue)
	return newValue
}

// Information about a party's account.
type AccountParties11 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties6Choice `xml:"PrncplAcctPty,omitempty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation11 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation11 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation11 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation11 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath *InvestmentAccountOwnershipInformation11 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation11 `xml:"Admstr,omitempty"`

	// An other type of party.
	OtherParty []*ExtendedParty8 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation11 `xml:"Grntr,omitempty"`

	// Entity that creates a trust or contributes assets to the trust.
	Settlor []*InvestmentAccountOwnershipInformation11 `xml:"Sttlr,omitempty"`
}

func (a *AccountParties11) SetModificationScopeIndication(value string) {
	a.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (a *AccountParties11) AddPrincipalAccountParty() *AccountParties6Choice {
	a.PrincipalAccountParty = new(AccountParties6Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties11) AddSecondaryOwner() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties11) AddBeneficiary() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties11) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties11) AddLegalGuardian() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties11) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation11 {
	a.SuccessorOnDeath = new(InvestmentAccountOwnershipInformation11)
	return a.SuccessorOnDeath
}

func (a *AccountParties11) AddAdministrator() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties11) AddOtherParty() *ExtendedParty8 {
	newValue := new(ExtendedParty8)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties11) AddGranter() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties11) AddSettlor() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.Settlor = append(a.Settlor, newValue)
	return newValue
}

// Information about a party's account.
type AccountParties12 struct {

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties7Choice `xml:"PrncplAcctPty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation10 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation10 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation10 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation10 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation10 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation10 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty7 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation10 `xml:"Grntr,omitempty"`

	// Entity that creates a trust or contributes assets to the trust.
	Settlor []*InvestmentAccountOwnershipInformation10 `xml:"Sttlr,omitempty"`
}

func (a *AccountParties12) AddPrincipalAccountParty() *AccountParties7Choice {
	a.PrincipalAccountParty = new(AccountParties7Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties12) AddSecondaryOwner() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties12) AddBeneficiary() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties12) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties12) AddLegalGuardian() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties12) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties12) AddAdministrator() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties12) AddOtherParty() *ExtendedParty7 {
	newValue := new(ExtendedParty7)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties12) AddGranter() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties12) AddSettlor() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.Settlor = append(a.Settlor, newValue)
	return newValue
}

// Information about a party's account.
type AccountParties13 struct {

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties8Choice `xml:"PrncplAcctPty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation12 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation12 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation12 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation12 `xml:"LglGuardn,omitempty"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation12 `xml:"CtdnForMnr,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation12 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation12 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty9 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation12 `xml:"Grntr,omitempty"`

	// Entity that creates a trust or contributes assets to the trust.
	Settlor []*InvestmentAccountOwnershipInformation12 `xml:"Sttlr,omitempty"`

	// Party that registers its name with the issuer and the name used for the registration.
	RegisteredShareholderName *RegisteredShareholderName1Choice `xml:"RegdShrhldrNm,omitempty"`
}

func (a *AccountParties13) AddPrincipalAccountParty() *AccountParties8Choice {
	a.PrincipalAccountParty = new(AccountParties8Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties13) AddSecondaryOwner() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties13) AddBeneficiary() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties13) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties13) AddLegalGuardian() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties13) AddCustodianForMinor() *InvestmentAccountOwnershipInformation12 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation12)
	return a.CustodianForMinor
}

func (a *AccountParties13) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties13) AddAdministrator() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties13) AddOtherParty() *ExtendedParty9 {
	newValue := new(ExtendedParty9)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties13) AddGranter() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties13) AddSettlor() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.Settlor = append(a.Settlor, newValue)
	return newValue
}

func (a *AccountParties13) AddRegisteredShareholderName() *RegisteredShareholderName1Choice {
	a.RegisteredShareholderName = new(RegisteredShareholderName1Choice)
	return a.RegisteredShareholderName
}

// Information about a party's account.
type AccountParties14 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties9Choice `xml:"PrncplAcctPty,omitempty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner *InvestmentAccountOwnershipInformation13 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary *InvestmentAccountOwnershipInformation13 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney *InvestmentAccountOwnershipInformation13 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian *InvestmentAccountOwnershipInformation13 `xml:"LglGuardn,omitempty"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation13 `xml:"CtdnForMnr,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath *InvestmentAccountOwnershipInformation13 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person or organisation that has gone bankrupt.
	Administrator *InvestmentAccountOwnershipInformation13 `xml:"Admstr,omitempty"`

	// An other type of party.
	OtherParty *ExtendedParty10 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter *InvestmentAccountOwnershipInformation13 `xml:"Grntr,omitempty"`

	// Entity that creates a trust or contributes assets to the trust.
	Settlor *InvestmentAccountOwnershipInformation13 `xml:"Sttlr,omitempty"`

	// Party for which shares are to be registered.
	RegisteredShareholderName *RegisteredShareholderName1Choice `xml:"RegdShrhldrNm,omitempty"`
}

func (a *AccountParties14) SetModificationScopeIndication(value string) {
	a.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (a *AccountParties14) AddPrincipalAccountParty() *AccountParties9Choice {
	a.PrincipalAccountParty = new(AccountParties9Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties14) AddSecondaryOwner() *InvestmentAccountOwnershipInformation13 {
	a.SecondaryOwner = new(InvestmentAccountOwnershipInformation13)
	return a.SecondaryOwner
}

func (a *AccountParties14) AddBeneficiary() *InvestmentAccountOwnershipInformation13 {
	a.Beneficiary = new(InvestmentAccountOwnershipInformation13)
	return a.Beneficiary
}

func (a *AccountParties14) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation13 {
	a.PowerOfAttorney = new(InvestmentAccountOwnershipInformation13)
	return a.PowerOfAttorney
}

func (a *AccountParties14) AddLegalGuardian() *InvestmentAccountOwnershipInformation13 {
	a.LegalGuardian = new(InvestmentAccountOwnershipInformation13)
	return a.LegalGuardian
}

func (a *AccountParties14) AddCustodianForMinor() *InvestmentAccountOwnershipInformation13 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation13)
	return a.CustodianForMinor
}

func (a *AccountParties14) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation13 {
	a.SuccessorOnDeath = new(InvestmentAccountOwnershipInformation13)
	return a.SuccessorOnDeath
}

func (a *AccountParties14) AddAdministrator() *InvestmentAccountOwnershipInformation13 {
	a.Administrator = new(InvestmentAccountOwnershipInformation13)
	return a.Administrator
}

func (a *AccountParties14) AddOtherParty() *ExtendedParty10 {
	a.OtherParty = new(ExtendedParty10)
	return a.OtherParty
}

func (a *AccountParties14) AddGranter() *InvestmentAccountOwnershipInformation13 {
	a.Granter = new(InvestmentAccountOwnershipInformation13)
	return a.Granter
}

func (a *AccountParties14) AddSettlor() *InvestmentAccountOwnershipInformation13 {
	a.Settlor = new(InvestmentAccountOwnershipInformation13)
	return a.Settlor
}

func (a *AccountParties14) AddRegisteredShareholderName() *RegisteredShareholderName1Choice {
	a.RegisteredShareholderName = new(RegisteredShareholderName1Choice)
	return a.RegisteredShareholderName
}

// Information about a party's account.
type AccountParties15 struct {

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties10Choice `xml:"PrncplAcctPty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation14 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation14 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation14 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation14 `xml:"LglGuardn,omitempty"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation14 `xml:"CtdnForMnr,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation14 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation14 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty11 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation14 `xml:"Grntr,omitempty"`

	// Entity that creates a trust or contributes assets to the trust.
	Settlor []*InvestmentAccountOwnershipInformation14 `xml:"Sttlr,omitempty"`

	// Party that makes, or participates in the making of, decisions that affect the whole, or a substantial part, of the business of a customer of a reporting entity or that has the capacity to affect significantly the financial standing of a customer of a reporting entity. Typically, this is a controlling person of a corporate (ownership type CORP).
	SeniorManagingOfficial []*InvestmentAccountOwnershipInformation14 `xml:"SnrMggOffcl,omitempty"`

	// Person appointed under the trust instrument to direct or restrain the trustees in relation to their administration of the trust. Typically, this is a controlling person of a trust (ownership type TRUS) or other non-individual organisation (ownership type ONIS).
	Protector []*InvestmentAccountOwnershipInformation14 `xml:"Prtctr,omitempty"`

	// Party that registers its name with the issuer and the name used for the registration.
	RegisteredShareholderName *RegisteredShareholderName1Choice `xml:"RegdShrhldrNm,omitempty"`
}

func (a *AccountParties15) AddPrincipalAccountParty() *AccountParties10Choice {
	a.PrincipalAccountParty = new(AccountParties10Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties15) AddSecondaryOwner() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties15) AddBeneficiary() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties15) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties15) AddLegalGuardian() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties15) AddCustodianForMinor() *InvestmentAccountOwnershipInformation14 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation14)
	return a.CustodianForMinor
}

func (a *AccountParties15) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties15) AddAdministrator() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties15) AddOtherParty() *ExtendedParty11 {
	newValue := new(ExtendedParty11)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties15) AddGranter() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties15) AddSettlor() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.Settlor = append(a.Settlor, newValue)
	return newValue
}

func (a *AccountParties15) AddSeniorManagingOfficial() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.SeniorManagingOfficial = append(a.SeniorManagingOfficial, newValue)
	return newValue
}

func (a *AccountParties15) AddProtector() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.Protector = append(a.Protector, newValue)
	return newValue
}

func (a *AccountParties15) AddRegisteredShareholderName() *RegisteredShareholderName1Choice {
	a.RegisteredShareholderName = new(RegisteredShareholderName1Choice)
	return a.RegisteredShareholderName
}

// Information about a party's account.
type AccountParties16 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties11Choice `xml:"PrncplAcctPty,omitempty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation15 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation15 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation15 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation15 `xml:"LglGuardn,omitempty"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation15 `xml:"CtdnForMnr,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation15 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation15 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty12 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation15 `xml:"Grntr,omitempty"`

	// Entity that creates a trust or contributes assets to the trust.
	Settlor []*InvestmentAccountOwnershipInformation15 `xml:"Sttlr,omitempty"`

	// Party that makes, or participates in the making of, decisions that affect the whole, or a substantial part, of the business of a customer of a reporting entity or that has the capacity to affect significantly the financial standing of a customer of a reporting entity. Typically, this is a controlling person of a corporate (ownership type CORP).
	SeniorManagingOfficial []*InvestmentAccountOwnershipInformation15 `xml:"SnrMggOffcl,omitempty"`

	// Person appointed under the trust instrument to direct or restrain the trustees in relation to their administration of the trust. Typically, this is a controlling person of a trust (ownership type TRUS) or other non-individual organisation (ownership type ONIS).
	Protector []*InvestmentAccountOwnershipInformation15 `xml:"Prtctr,omitempty"`

	// Party for which shares are to be registered.
	RegisteredShareholderName *RegisteredShareholderName1Choice `xml:"RegdShrhldrNm,omitempty"`
}

func (a *AccountParties16) SetModificationScopeIndication(value string) {
	a.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (a *AccountParties16) AddPrincipalAccountParty() *AccountParties11Choice {
	a.PrincipalAccountParty = new(AccountParties11Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties16) AddSecondaryOwner() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties16) AddBeneficiary() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties16) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties16) AddLegalGuardian() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties16) AddCustodianForMinor() *InvestmentAccountOwnershipInformation15 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation15)
	return a.CustodianForMinor
}

func (a *AccountParties16) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties16) AddAdministrator() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties16) AddOtherParty() *ExtendedParty12 {
	newValue := new(ExtendedParty12)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties16) AddGranter() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties16) AddSettlor() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.Settlor = append(a.Settlor, newValue)
	return newValue
}

func (a *AccountParties16) AddSeniorManagingOfficial() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.SeniorManagingOfficial = append(a.SeniorManagingOfficial, newValue)
	return newValue
}

func (a *AccountParties16) AddProtector() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.Protector = append(a.Protector, newValue)
	return newValue
}

func (a *AccountParties16) AddRegisteredShareholderName() *RegisteredShareholderName1Choice {
	a.RegisteredShareholderName = new(RegisteredShareholderName1Choice)
	return a.RegisteredShareholderName
}

// Any party who is related to an investment account.
type AccountParties4 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Single owner of the investment account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation4 `xml:"PmryOwnr,omitempty"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation4 `xml:"Trstee,omitempty"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation4 `xml:"CtdnForMnr,omitempty"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation4 `xml:"Nmnee,omitempty"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation4 `xml:"JntOwnr,omitempty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner *InvestmentAccountOwnershipInformation4 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary *InvestmentAccountOwnershipInformation4 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney *InvestmentAccountOwnershipInformation4 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian *InvestmentAccountOwnershipInformation4 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath *InvestmentAccountOwnershipInformation4 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authorithy to act on behalf of a person or organisation that has gone bankrupt.
	Administrator *InvestmentAccountOwnershipInformation4 `xml:"Admstr,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation4 `xml:"Grntr,omitempty"`

	// Settler role in the hedge funds industry.
	Settler []*InvestmentAccountOwnershipInformation4 `xml:"Sttlr,omitempty"`

	// An other type of party.
	OtherParty []*ExtendedParty1 `xml:"OthrPty,omitempty"`
}

func (a *AccountParties4) SetModificationScopeIndication(value string) {
	a.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (a *AccountParties4) AddPrimaryOwner() *InvestmentAccountOwnershipInformation4 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation4)
	return a.PrimaryOwner
}

func (a *AccountParties4) AddTrustee() *InvestmentAccountOwnershipInformation4 {
	newValue := new(InvestmentAccountOwnershipInformation4)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties4) AddCustodianForMinor() *InvestmentAccountOwnershipInformation4 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation4)
	return a.CustodianForMinor
}

func (a *AccountParties4) AddNominee() *InvestmentAccountOwnershipInformation4 {
	a.Nominee = new(InvestmentAccountOwnershipInformation4)
	return a.Nominee
}

func (a *AccountParties4) AddJointOwner() *InvestmentAccountOwnershipInformation4 {
	newValue := new(InvestmentAccountOwnershipInformation4)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}

func (a *AccountParties4) AddSecondaryOwner() *InvestmentAccountOwnershipInformation4 {
	a.SecondaryOwner = new(InvestmentAccountOwnershipInformation4)
	return a.SecondaryOwner
}

func (a *AccountParties4) AddBeneficiary() *InvestmentAccountOwnershipInformation4 {
	a.Beneficiary = new(InvestmentAccountOwnershipInformation4)
	return a.Beneficiary
}

func (a *AccountParties4) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation4 {
	a.PowerOfAttorney = new(InvestmentAccountOwnershipInformation4)
	return a.PowerOfAttorney
}

func (a *AccountParties4) AddLegalGuardian() *InvestmentAccountOwnershipInformation4 {
	a.LegalGuardian = new(InvestmentAccountOwnershipInformation4)
	return a.LegalGuardian
}

func (a *AccountParties4) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation4 {
	a.SuccessorOnDeath = new(InvestmentAccountOwnershipInformation4)
	return a.SuccessorOnDeath
}

func (a *AccountParties4) AddAdministrator() *InvestmentAccountOwnershipInformation4 {
	a.Administrator = new(InvestmentAccountOwnershipInformation4)
	return a.Administrator
}

func (a *AccountParties4) AddGranter() *InvestmentAccountOwnershipInformation4 {
	newValue := new(InvestmentAccountOwnershipInformation4)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties4) AddSettler() *InvestmentAccountOwnershipInformation4 {
	newValue := new(InvestmentAccountOwnershipInformation4)
	a.Settler = append(a.Settler, newValue)
	return newValue
}

func (a *AccountParties4) AddOtherParty() *ExtendedParty1 {
	newValue := new(ExtendedParty1)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

// Any party who is related to an investment account.
type AccountParties5 struct {

	// Single owner of the investment account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation5 `xml:"PmryOwnr"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation5 `xml:"Trstee"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation5 `xml:"CtdnForMnr"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation5 `xml:"Nmnee"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation5 `xml:"JntOwnr"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation5 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation5 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation5 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation5 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation5 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authorithy to act on behalf of a person or organisation that has gone bankrupt.
	Administrator *InvestmentAccountOwnershipInformation5 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty2 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation5 `xml:"Grntr,omitempty"`

	// Settler role in the hedge funds industry.
	Settler []*InvestmentAccountOwnershipInformation5 `xml:"Sttlr,omitempty"`
}

func (a *AccountParties5) AddPrimaryOwner() *InvestmentAccountOwnershipInformation5 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation5)
	return a.PrimaryOwner
}

func (a *AccountParties5) AddTrustee() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties5) AddCustodianForMinor() *InvestmentAccountOwnershipInformation5 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation5)
	return a.CustodianForMinor
}

func (a *AccountParties5) AddNominee() *InvestmentAccountOwnershipInformation5 {
	a.Nominee = new(InvestmentAccountOwnershipInformation5)
	return a.Nominee
}

func (a *AccountParties5) AddJointOwner() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}

func (a *AccountParties5) AddSecondaryOwner() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties5) AddBeneficiary() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties5) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties5) AddLegalGuardian() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties5) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties5) AddAdministrator() *InvestmentAccountOwnershipInformation5 {
	a.Administrator = new(InvestmentAccountOwnershipInformation5)
	return a.Administrator
}

func (a *AccountParties5) AddOtherParty() *ExtendedParty2 {
	newValue := new(ExtendedParty2)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties5) AddGranter() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties5) AddSettler() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.Settler = append(a.Settler, newValue)
	return newValue
}

// Party related to an investment account.
type AccountParties6 struct {

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties1Choice `xml:"PrncplAcctPty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation6 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation6 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation6 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation6 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation6 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authorithy to act on behalf of a person or organisation that has gone bankrupt.
	Administrator *InvestmentAccountOwnershipInformation6 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty3 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation6 `xml:"Grntr,omitempty"`

	// Settler role in the hedge funds industry.
	Settler []*InvestmentAccountOwnershipInformation6 `xml:"Sttlr,omitempty"`
}

func (a *AccountParties6) AddPrincipalAccountParty() *AccountParties1Choice {
	a.PrincipalAccountParty = new(AccountParties1Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties6) AddSecondaryOwner() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties6) AddBeneficiary() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties6) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties6) AddLegalGuardian() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties6) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties6) AddAdministrator() *InvestmentAccountOwnershipInformation6 {
	a.Administrator = new(InvestmentAccountOwnershipInformation6)
	return a.Administrator
}

func (a *AccountParties6) AddOtherParty() *ExtendedParty3 {
	newValue := new(ExtendedParty3)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties6) AddGranter() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties6) AddSettler() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.Settler = append(a.Settler, newValue)
	return newValue
}

// Party related to an investment account.
type AccountParties7 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties2Choice `xml:"PrncplAcctPty,omitempty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner *InvestmentAccountOwnershipInformation7 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary *InvestmentAccountOwnershipInformation7 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney *InvestmentAccountOwnershipInformation7 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian *InvestmentAccountOwnershipInformation7 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath *InvestmentAccountOwnershipInformation7 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authorithy to act on behalf of a person or organisation that has gone bankrupt.
	Administrator *InvestmentAccountOwnershipInformation7 `xml:"Admstr,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation7 `xml:"Grntr,omitempty"`

	// Settler role in the hedge funds industry.
	Settler []*InvestmentAccountOwnershipInformation7 `xml:"Sttlr,omitempty"`

	// An other type of party.
	OtherParty []*ExtendedParty4 `xml:"OthrPty,omitempty"`
}

func (a *AccountParties7) SetModificationScopeIndication(value string) {
	a.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (a *AccountParties7) AddPrincipalAccountParty() *AccountParties2Choice {
	a.PrincipalAccountParty = new(AccountParties2Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties7) AddSecondaryOwner() *InvestmentAccountOwnershipInformation7 {
	a.SecondaryOwner = new(InvestmentAccountOwnershipInformation7)
	return a.SecondaryOwner
}

func (a *AccountParties7) AddBeneficiary() *InvestmentAccountOwnershipInformation7 {
	a.Beneficiary = new(InvestmentAccountOwnershipInformation7)
	return a.Beneficiary
}

func (a *AccountParties7) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation7 {
	a.PowerOfAttorney = new(InvestmentAccountOwnershipInformation7)
	return a.PowerOfAttorney
}

func (a *AccountParties7) AddLegalGuardian() *InvestmentAccountOwnershipInformation7 {
	a.LegalGuardian = new(InvestmentAccountOwnershipInformation7)
	return a.LegalGuardian
}

func (a *AccountParties7) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation7 {
	a.SuccessorOnDeath = new(InvestmentAccountOwnershipInformation7)
	return a.SuccessorOnDeath
}

func (a *AccountParties7) AddAdministrator() *InvestmentAccountOwnershipInformation7 {
	a.Administrator = new(InvestmentAccountOwnershipInformation7)
	return a.Administrator
}

func (a *AccountParties7) AddGranter() *InvestmentAccountOwnershipInformation7 {
	newValue := new(InvestmentAccountOwnershipInformation7)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties7) AddSettler() *InvestmentAccountOwnershipInformation7 {
	newValue := new(InvestmentAccountOwnershipInformation7)
	a.Settler = append(a.Settler, newValue)
	return newValue
}

func (a *AccountParties7) AddOtherParty() *ExtendedParty4 {
	newValue := new(ExtendedParty4)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

// Party related to an investment account.
type AccountParties8 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties3Choice `xml:"PrncplAcctPty,omitempty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation8 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation8 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation8 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation8 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath *InvestmentAccountOwnershipInformation8 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authorithy to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation8 `xml:"Admstr,omitempty"`

	// An other type of party.
	OtherParty []*ExtendedParty5 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation8 `xml:"Grntr,omitempty"`

	// Settler role in the hedge funds industry.
	Settler []*InvestmentAccountOwnershipInformation8 `xml:"Sttlr,omitempty"`
}

func (a *AccountParties8) SetModificationScopeIndication(value string) {
	a.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (a *AccountParties8) AddPrincipalAccountParty() *AccountParties3Choice {
	a.PrincipalAccountParty = new(AccountParties3Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties8) AddSecondaryOwner() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties8) AddBeneficiary() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties8) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties8) AddLegalGuardian() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties8) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation8 {
	a.SuccessorOnDeath = new(InvestmentAccountOwnershipInformation8)
	return a.SuccessorOnDeath
}

func (a *AccountParties8) AddAdministrator() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties8) AddOtherParty() *ExtendedParty5 {
	newValue := new(ExtendedParty5)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties8) AddGranter() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties8) AddSettler() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.Settler = append(a.Settler, newValue)
	return newValue
}

// Party related to an investment account.
type AccountParties9 struct {

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties4Choice `xml:"PrncplAcctPty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation9 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation9 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation9 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation9 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation9 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authorithy to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation9 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty6 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation9 `xml:"Grntr,omitempty"`

	// Settler role in the hedge funds industry.
	Settler []*InvestmentAccountOwnershipInformation9 `xml:"Sttlr,omitempty"`
}

func (a *AccountParties9) AddPrincipalAccountParty() *AccountParties4Choice {
	a.PrincipalAccountParty = new(AccountParties4Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties9) AddSecondaryOwner() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties9) AddBeneficiary() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties9) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties9) AddLegalGuardian() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties9) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties9) AddAdministrator() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties9) AddOtherParty() *ExtendedParty6 {
	newValue := new(ExtendedParty6)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties9) AddGranter() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties9) AddSettler() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.Settler = append(a.Settler, newValue)
	return newValue
}
