package iso20022

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount10 struct {

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification1Choice `xml:"OwnrId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary1 `xml:"IntrmyInf,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Indicates whether a security exists only as an electronic record, ie, there is no physical document representing the security.
	DematerialisedIndicator *YesNoIndicator `xml:"DmtrlsdInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Indicates whether the beneficial ownership certification has been sent, certifying that the beneficial owner is eligible to own a specific investment fund or investment fund class.
	BeneficiaryCertificationIndicator *YesNoIndicator `xml:"BnfcryCertfctnInd,omitempty"`

	// Place requested as the place of safekeeping.
	SafekeepingPlace *PartyIdentification1Choice `xml:"SfkpgPlc,omitempty"`

	// Party related to an account that is not the legal account owner, eg, the power of attorney.
	AccountServicer *PartyIdentification1Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount10) AddOwnerIdentification() *PartyIdentification1Choice {
	newValue := new(PartyIdentification1Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount10) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount10) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount10) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount10) AddIntermediaryInformation() *Intermediary1 {
	newValue := new(Intermediary1)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount10) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentAccount10) SetDematerialisedIndicator(value string) {
	i.DematerialisedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount10) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount10) SetBeneficiaryCertificationIndicator(value string) {
	i.BeneficiaryCertificationIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount10) AddSafekeepingPlace() *PartyIdentification1Choice {
	i.SafekeepingPlace = new(PartyIdentification1Choice)
	return i.SafekeepingPlace
}

func (i *InvestmentAccount10) AddAccountServicer() *PartyIdentification1Choice {
	i.AccountServicer = new(PartyIdentification1Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount11 struct {

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification1Choice `xml:"OwnrId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary1 `xml:"IntrmyInf,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Indicates whether a security exists only as an electronic record, ie, there is no physical document representing the security.
	DematerialisedIndicator *YesNoIndicator `xml:"DmtrlsdInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Indicates whether the beneficial ownership certification has been sent, certifying that the beneficial owner is eligible to own a specific investment fund or investment fund class.
	BeneficiaryCertificationIndicator *YesNoIndicator `xml:"BnfcryCertfctnInd,omitempty"`

	// Party related to an account that is not the legal account owner, eg, the power of attorney.
	AccountServicer *PartyIdentification1Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount11) AddOwnerIdentification() *PartyIdentification1Choice {
	newValue := new(PartyIdentification1Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount11) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount11) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount11) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount11) AddIntermediaryInformation() *Intermediary1 {
	newValue := new(Intermediary1)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount11) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentAccount11) SetDematerialisedIndicator(value string) {
	i.DematerialisedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount11) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount11) SetBeneficiaryCertificationIndicator(value string) {
	i.BeneficiaryCertificationIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount11) AddAccountServicer() *PartyIdentification1Choice {
	i.AccountServicer = new(PartyIdentification1Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount12 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary1 `xml:"IntrmyInf,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification1Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount12) AddIdentification() *AccountIdentification1 {
	i.Identification = new(AccountIdentification1)
	return i.Identification
}

func (i *InvestmentAccount12) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount12) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount12) AddIntermediaryInformation() *Intermediary1 {
	newValue := new(Intermediary1)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount12) AddAccountServicer() *PartyIdentification1Choice {
	i.AccountServicer = new(PartyIdentification1Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount13 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification2Choice `xml:"OwnrId,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount13) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount13) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount13) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount13) AddOwnerIdentification() *PartyIdentification2Choice {
	newValue := new(PartyIdentification2Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount13) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount14 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Identification of an individual person whom legally owns the account.
	IndividualOwnerIdentification *IndividualPersonIdentificationChoice `xml:"IndvOwnrId,omitempty"`

	// Identification of an organisation that legally owns the account.
	OrganisationOwnerIdentification *PartyIdentification2Choice `xml:"OrgOwnrId,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount14) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount14) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount14) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount14) AddIndividualOwnerIdentification() *IndividualPersonIdentificationChoice {
	i.IndividualOwnerIdentification = new(IndividualPersonIdentificationChoice)
	return i.IndividualOwnerIdentification
}

func (i *InvestmentAccount14) AddOrganisationOwnerIdentification() *PartyIdentification2Choice {
	i.OrganisationOwnerIdentification = new(PartyIdentification2Choice)
	return i.OrganisationOwnerIdentification
}

func (i *InvestmentAccount14) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount15 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Account type.
	Type *CashAccountType1 `xml:"Tp,omitempty"`
}

func (i *InvestmentAccount15) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount15) AddType() *CashAccountType1 {
	i.Type = new(CashAccountType1)
	return i.Type
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount20 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Account type.
	Type *FundCashAccount2Code `xml:"Tp,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	ExtendedType *Extended350Code `xml:"XtndedTp,omitempty"`
}

func (i *InvestmentAccount20) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount20) SetType(value string) {
	i.Type = (*FundCashAccount2Code)(&value)
}

func (i *InvestmentAccount20) SetExtendedType(value string) {
	i.ExtendedType = (*Extended350Code)(&value)
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount21 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification2Choice `xml:"OwnrId,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Counterparties eligibility as defined by article 24 of the EU MiFID Directive applicable to transactions executed by investment firms for eligible counterparties.
	OrderOriginatorEligibility *OrderOriginatorEligibility1Code `xml:"OrdrOrgtrElgblty,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`
}

func (i *InvestmentAccount21) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount21) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount21) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount21) AddOwnerIdentification() *PartyIdentification2Choice {
	newValue := new(PartyIdentification2Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount21) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount21) SetOrderOriginatorEligibility(value string) {
	i.OrderOriginatorEligibility = (*OrderOriginatorEligibility1Code)(&value)
}

func (i *InvestmentAccount21) AddSubAccountDetails() *SubAccount1 {
	i.SubAccountDetails = new(SubAccount1)
	return i.SubAccountDetails
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount22 struct {

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification2Choice `xml:"OwnrId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary11 `xml:"IntrmyInf,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Indicates whether a security exists only as an electronic record, ie, there is no physical document representing the security.
	DematerialisedIndicator *YesNoIndicator `xml:"DmtrlsdInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Beneficial owner or its designated agent certifies that it complies with any holding or investment restrictions or requirements of the fund.
	BeneficiaryCertificationCompletion *BeneficiaryCertificationCompletion1Code `xml:"BnfcryCertfctnCmpltn,omitempty"`

	// Place requested as the place of safekeeping.
	SafekeepingPlace *PartyIdentification2Choice `xml:"SfkpgPlc,omitempty"`

	// Party related to an account that is not the legal account owner, eg, the power of attorney.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount22) AddOwnerIdentification() *PartyIdentification2Choice {
	newValue := new(PartyIdentification2Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount22) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount22) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount22) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount22) AddIntermediaryInformation() *Intermediary11 {
	newValue := new(Intermediary11)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount22) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentAccount22) SetDematerialisedIndicator(value string) {
	i.DematerialisedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount22) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount22) SetBeneficiaryCertificationCompletion(value string) {
	i.BeneficiaryCertificationCompletion = (*BeneficiaryCertificationCompletion1Code)(&value)
}

func (i *InvestmentAccount22) AddSafekeepingPlace() *PartyIdentification2Choice {
	i.SafekeepingPlace = new(PartyIdentification2Choice)
	return i.SafekeepingPlace
}

func (i *InvestmentAccount22) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount24 struct {

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification2Choice `xml:"OwnrId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary10 `xml:"IntrmyInf,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Indicates whether a security exists only as an electronic record, ie, there is no physical document representing the security.
	DematerialisedIndicator *YesNoIndicator `xml:"DmtrlsdInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Beneficial owner or its designated agent certifies that it complies with any holding or investment restrictions or requirements of the fund.
	BeneficiaryCertificationCompletion *BeneficiaryCertificationCompletion1Code `xml:"BnfcryCertfctnCmpltn,omitempty"`

	// Party related to an account that is not the legal account owner, eg, the power of attorney.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`
}

func (i *InvestmentAccount24) AddOwnerIdentification() *PartyIdentification2Choice {
	newValue := new(PartyIdentification2Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount24) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount24) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount24) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount24) AddIntermediaryInformation() *Intermediary10 {
	newValue := new(Intermediary10)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount24) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentAccount24) SetDematerialisedIndicator(value string) {
	i.DematerialisedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount24) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount24) SetBeneficiaryCertificationCompletion(value string) {
	i.BeneficiaryCertificationCompletion = (*BeneficiaryCertificationCompletion1Code)(&value)
}

func (i *InvestmentAccount24) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount24) AddSubAccountDetails() *SubAccount1 {
	i.SubAccountDetails = new(SubAccount1)
	return i.SubAccountDetails
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount25 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary11 `xml:"IntrmyInf,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount25) AddIdentification() *AccountIdentification1 {
	i.Identification = new(AccountIdentification1)
	return i.Identification
}

func (i *InvestmentAccount25) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount25) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount25) AddIntermediaryInformation() *Intermediary11 {
	newValue := new(Intermediary11)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount25) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount26 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	Type *FundCashAccount3Code `xml:"Tp,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	ExtendedType *Extended350Code `xml:"XtndedTp,omitempty"`

	// Ownership status of the account, eg, joint owners.
	OwnershipType *AccountOwnershipType3Code `xml:"OwnrshTp"`

	// Ownership status of the account, eg, joint owners.
	ExtendedOwnershipType *Extended350Code `xml:"XtndedOwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemptionReason *TaxExemptReason1Code `xml:"TaxXmptnRsn,omitempty"`

	// Tax advantage specific to the account.
	ExtendedTaxExemptionReason *Extended350Code `xml:"XtndedTaxXmptnRsn,omitempty"`

	// Regularity at which a statement is issued.
	StatementFrequency *EventFrequency1Code `xml:"StmtFrqcy,omitempty"`

	// Regularity at which a statement is issued.
	ExtendedStatementFrequency *Extended350Code `xml:"XtndedStmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Method by which the tax (withholding tax) is to be processed i.e. either withheld at source or tax information reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod1Code `xml:"TaxWhldgMtd,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	FundDetails []*FinancialInstrument10 `xml:"FndDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount26) AddIdentification() *AccountIdentification1 {
	i.Identification = new(AccountIdentification1)
	return i.Identification
}

func (i *InvestmentAccount26) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount26) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount26) SetType(value string) {
	i.Type = (*FundCashAccount3Code)(&value)
}

func (i *InvestmentAccount26) SetExtendedType(value string) {
	i.ExtendedType = (*Extended350Code)(&value)
}

func (i *InvestmentAccount26) SetOwnershipType(value string) {
	i.OwnershipType = (*AccountOwnershipType3Code)(&value)
}

func (i *InvestmentAccount26) SetExtendedOwnershipType(value string) {
	i.ExtendedOwnershipType = (*Extended350Code)(&value)
}

func (i *InvestmentAccount26) SetTaxExemptionReason(value string) {
	i.TaxExemptionReason = (*TaxExemptReason1Code)(&value)
}

func (i *InvestmentAccount26) SetExtendedTaxExemptionReason(value string) {
	i.ExtendedTaxExemptionReason = (*Extended350Code)(&value)
}

func (i *InvestmentAccount26) SetStatementFrequency(value string) {
	i.StatementFrequency = (*EventFrequency1Code)(&value)
}

func (i *InvestmentAccount26) SetExtendedStatementFrequency(value string) {
	i.ExtendedStatementFrequency = (*Extended350Code)(&value)
}

func (i *InvestmentAccount26) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount26) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount26) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount26) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod1Code)(&value)
}

func (i *InvestmentAccount26) SetLetterIntentReference(value string) {
	i.LetterIntentReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount26) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount26) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount26) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount26) AddFundDetails() *FinancialInstrument10 {
	newValue := new(FinancialInstrument10)
	i.FundDetails = append(i.FundDetails, newValue)
	return newValue
}

func (i *InvestmentAccount26) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount26) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount27 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Specifies the current state of an account, eg, enabled or deleted.
	Status *AccountStatus2Code `xml:"Sts"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	Type *FundCashAccount3Code `xml:"Tp,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	ExtendedType *Extended350Code `xml:"XtndedTp,omitempty"`

	// Ownership status of the account, eg, joint owners.
	OwnershipType *AccountOwnershipType3Code `xml:"OwnrshTp"`

	// Ownership status of the account, eg, joint owners.
	ExtendedOwnershipType *Extended350Code `xml:"XtndedOwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemptionReason *TaxExemptReason1Code `xml:"TaxXmptnRsn,omitempty"`

	// Tax advantage specific to the account.
	ExtendedTaxExemptionReason *Extended350Code `xml:"XtndedTaxXmptnRsn,omitempty"`

	// Regularity at which a statement is issued.
	StatementFrequency *EventFrequency1Code `xml:"StmtFrqcy,omitempty"`

	// Regularity at which a statement is issued.
	ExtendedStatementFrequency *Extended350Code `xml:"XtndedStmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Method by which the tax (withholding tax) is to be processed i.e. either withheld at source or tax information reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod1Code `xml:"TaxWhldgMtd,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Detailed information about the investment fund associated to the account.
	FundsDetails []*FinancialInstrument10 `xml:"FndsDtls,omitempty"`

	// Part of the investment account to or from which cash entries are made.
	CashAccount []*CashAccount12 `xml:"CshAcct,omitempty"`

	// Part of the investment account to or from which securities entries are made.
	SecuritiesAccount []*SecuritiesAccount4 `xml:"SctiesAcct,omitempty"`
}

func (i *InvestmentAccount27) AddIdentification() *AccountIdentification1 {
	i.Identification = new(AccountIdentification1)
	return i.Identification
}

func (i *InvestmentAccount27) SetStatus(value string) {
	i.Status = (*AccountStatus2Code)(&value)
}

func (i *InvestmentAccount27) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount27) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount27) SetType(value string) {
	i.Type = (*FundCashAccount3Code)(&value)
}

func (i *InvestmentAccount27) SetExtendedType(value string) {
	i.ExtendedType = (*Extended350Code)(&value)
}

func (i *InvestmentAccount27) SetOwnershipType(value string) {
	i.OwnershipType = (*AccountOwnershipType3Code)(&value)
}

func (i *InvestmentAccount27) SetExtendedOwnershipType(value string) {
	i.ExtendedOwnershipType = (*Extended350Code)(&value)
}

func (i *InvestmentAccount27) SetTaxExemptionReason(value string) {
	i.TaxExemptionReason = (*TaxExemptReason1Code)(&value)
}

func (i *InvestmentAccount27) SetExtendedTaxExemptionReason(value string) {
	i.ExtendedTaxExemptionReason = (*Extended350Code)(&value)
}

func (i *InvestmentAccount27) SetStatementFrequency(value string) {
	i.StatementFrequency = (*EventFrequency1Code)(&value)
}

func (i *InvestmentAccount27) SetExtendedStatementFrequency(value string) {
	i.ExtendedStatementFrequency = (*Extended350Code)(&value)
}

func (i *InvestmentAccount27) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount27) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount27) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount27) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod1Code)(&value)
}

func (i *InvestmentAccount27) SetLetterIntentReference(value string) {
	i.LetterIntentReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount27) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount27) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount27) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount27) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount27) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount27) AddFundsDetails() *FinancialInstrument10 {
	newValue := new(FinancialInstrument10)
	i.FundsDetails = append(i.FundsDetails, newValue)
	return newValue
}

func (i *InvestmentAccount27) AddCashAccount() *CashAccount12 {
	newValue := new(CashAccount12)
	i.CashAccount = append(i.CashAccount, newValue)
	return newValue
}

func (i *InvestmentAccount27) AddSecuritiesAccount() *SecuritiesAccount4 {
	newValue := new(SecuritiesAccount4)
	i.SecuritiesAccount = append(i.SecuritiesAccount, newValue)
	return newValue
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount28 struct {

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	Type *FundCashAccount3Code `xml:"Tp,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	ExtendedType *Extended350Code `xml:"XtndedTp,omitempty"`

	// Ownership status of the account, eg, joint owners.
	OwnershipType *AccountOwnershipType3Code `xml:"OwnrshTp,omitempty"`

	// Ownership status of the account, eg, joint owners.
	ExtendedOwnershipType *Extended350Code `xml:"XtndedOwnrshTp,omitempty"`

	// Tax advantage specific to the account.
	TaxExemptionReason *TaxExemptReason1Code `xml:"TaxXmptnRsn,omitempty"`

	// Tax advantage specific to the account.
	ExtendedTaxExemptionReason *Extended350Code `xml:"XtndedTaxXmptnRsn,omitempty"`

	// Regularity at which a statement is issued.
	StatementFrequency *EventFrequency1Code `xml:"StmtFrqcy,omitempty"`

	// Regularity at which a statement is issued.
	ExtendedStatementFrequency *Extended350Code `xml:"XtndedStmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Method by which the tax (withholding tax) is to be processed i.e. either withheld at source or tax information reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod1Code `xml:"TaxWhldgMtd,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	ModifiedFundDetails []*ModificationScope12 `xml:"ModfdFndDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount28) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount28) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount28) SetType(value string) {
	i.Type = (*FundCashAccount3Code)(&value)
}

func (i *InvestmentAccount28) SetExtendedType(value string) {
	i.ExtendedType = (*Extended350Code)(&value)
}

func (i *InvestmentAccount28) SetOwnershipType(value string) {
	i.OwnershipType = (*AccountOwnershipType3Code)(&value)
}

func (i *InvestmentAccount28) SetExtendedOwnershipType(value string) {
	i.ExtendedOwnershipType = (*Extended350Code)(&value)
}

func (i *InvestmentAccount28) SetTaxExemptionReason(value string) {
	i.TaxExemptionReason = (*TaxExemptReason1Code)(&value)
}

func (i *InvestmentAccount28) SetExtendedTaxExemptionReason(value string) {
	i.ExtendedTaxExemptionReason = (*Extended350Code)(&value)
}

func (i *InvestmentAccount28) SetStatementFrequency(value string) {
	i.StatementFrequency = (*EventFrequency1Code)(&value)
}

func (i *InvestmentAccount28) SetExtendedStatementFrequency(value string) {
	i.ExtendedStatementFrequency = (*Extended350Code)(&value)
}

func (i *InvestmentAccount28) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount28) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount28) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount28) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod1Code)(&value)
}

func (i *InvestmentAccount28) SetLetterIntentReference(value string) {
	i.LetterIntentReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount28) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount28) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount28) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount28) AddModifiedFundDetails() *ModificationScope12 {
	newValue := new(ModificationScope12)
	i.ModifiedFundDetails = append(i.ModifiedFundDetails, newValue)
	return newValue
}

func (i *InvestmentAccount28) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount28) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount29 struct {

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Legal form of the fund, eg, UCITS, SICAV, OEIC, Unit Trust, and FCP.
	FundType *Max35Text `xml:"FndTp,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	SecurityDetails *FinancialInstrument10 `xml:"SctyDtls,omitempty"`

	// Identification of an individual person whom legally owns the account.
	IndividualOwnerIdentification *IndividualPersonIdentificationChoice `xml:"IndvOwnrId,omitempty"`

	// Identification of an organisation that legally owns the account.
	OrganisationOwnerIdentification *PartyIdentification5Choice `xml:"OrgOwnrId,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	Intermediary []*Intermediary7 `xml:"Intrmy,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount29) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount29) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount29) SetFundType(value string) {
	i.FundType = (*Max35Text)(&value)
}

func (i *InvestmentAccount29) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount29) AddSecurityDetails() *FinancialInstrument10 {
	i.SecurityDetails = new(FinancialInstrument10)
	return i.SecurityDetails
}

func (i *InvestmentAccount29) AddIndividualOwnerIdentification() *IndividualPersonIdentificationChoice {
	i.IndividualOwnerIdentification = new(IndividualPersonIdentificationChoice)
	return i.IndividualOwnerIdentification
}

func (i *InvestmentAccount29) AddOrganisationOwnerIdentification() *PartyIdentification5Choice {
	i.OrganisationOwnerIdentification = new(PartyIdentification5Choice)
	return i.OrganisationOwnerIdentification
}

func (i *InvestmentAccount29) AddIntermediary() *Intermediary7 {
	newValue := new(Intermediary7)
	i.Intermediary = append(i.Intermediary, newValue)
	return newValue
}

func (i *InvestmentAccount29) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount34 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, for example, wrapper, ISA.
	Type *AccountType1Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType1Choice `xml:"OwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason1Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason1Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Method by which the tax (withholding tax) is to be processed i.e. either withheld at source or tax information reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod2Code `xml:"TaxWhldgMtd,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	FundDetails []*FinancialInstrument29 `xml:"FndDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Specifies information about blocked accounts.
	BlockedStatus []*Blocked1 `xml:"BlckdSts,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType1Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`
}

func (i *InvestmentAccount34) AddIdentification() *AccountIdentification1 {
	i.Identification = new(AccountIdentification1)
	return i.Identification
}

func (i *InvestmentAccount34) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount34) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount34) AddType() *AccountType1Choice {
	i.Type = new(AccountType1Choice)
	return i.Type
}

func (i *InvestmentAccount34) AddOwnershipType() *OwnershipType1Choice {
	i.OwnershipType = new(OwnershipType1Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount34) AddTaxExemption() *TaxExemptionReason1Choice {
	i.TaxExemption = new(TaxExemptionReason1Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount34) AddStatementFrequency() *StatementFrequencyReason1Choice {
	i.StatementFrequency = new(StatementFrequencyReason1Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount34) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount34) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount34) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount34) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod2Code)(&value)
}

func (i *InvestmentAccount34) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount34) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount34) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount34) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount34) AddFundDetails() *FinancialInstrument29 {
	newValue := new(FinancialInstrument29)
	i.FundDetails = append(i.FundDetails, newValue)
	return newValue
}

func (i *InvestmentAccount34) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount34) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount34) AddBlockedStatus() *Blocked1 {
	newValue := new(Blocked1)
	i.BlockedStatus = append(i.BlockedStatus, newValue)
	return newValue
}

func (i *InvestmentAccount34) AddAccountUsageType() *AccountUsageType1Choice {
	i.AccountUsageType = new(AccountUsageType1Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount34) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount34) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount35 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Specifies the current state of an account, eg, enabled or deleted.
	Status *AccountStatus2Code `xml:"Sts"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, for example, wrapper, ISA.
	Type *AccountType1Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType1Choice `xml:"OwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason1Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason1Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Method by which the tax (withholding tax) is to be processed i.e. either withheld at source or tax information reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod2Code `xml:"TaxWhldgMtd,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Detailed information about the investment fund associated to the account.
	FundsDetails []*FinancialInstrument29 `xml:"FndsDtls,omitempty"`

	// Part of the investment account to or from which cash entries are made.
	CashAccount []*CashAccount12 `xml:"CshAcct,omitempty"`

	// Part of the investment account to or from which securities entries are made.
	SecuritiesAccount []*SecuritiesAccount4 `xml:"SctiesAcct,omitempty"`

	// Specifies information about blocked accounts.
	BlockedStatus *Blocked1 `xml:"BlckdSts,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType1Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`
}

func (i *InvestmentAccount35) AddIdentification() *AccountIdentification1 {
	i.Identification = new(AccountIdentification1)
	return i.Identification
}

func (i *InvestmentAccount35) SetStatus(value string) {
	i.Status = (*AccountStatus2Code)(&value)
}

func (i *InvestmentAccount35) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount35) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount35) AddType() *AccountType1Choice {
	i.Type = new(AccountType1Choice)
	return i.Type
}

func (i *InvestmentAccount35) AddOwnershipType() *OwnershipType1Choice {
	i.OwnershipType = new(OwnershipType1Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount35) AddTaxExemption() *TaxExemptionReason1Choice {
	i.TaxExemption = new(TaxExemptionReason1Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount35) AddStatementFrequency() *StatementFrequencyReason1Choice {
	i.StatementFrequency = new(StatementFrequencyReason1Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount35) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount35) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount35) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount35) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod2Code)(&value)
}

func (i *InvestmentAccount35) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount35) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount35) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount35) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount35) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount35) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount35) AddFundsDetails() *FinancialInstrument29 {
	newValue := new(FinancialInstrument29)
	i.FundsDetails = append(i.FundsDetails, newValue)
	return newValue
}

func (i *InvestmentAccount35) AddCashAccount() *CashAccount12 {
	newValue := new(CashAccount12)
	i.CashAccount = append(i.CashAccount, newValue)
	return newValue
}

func (i *InvestmentAccount35) AddSecuritiesAccount() *SecuritiesAccount4 {
	newValue := new(SecuritiesAccount4)
	i.SecuritiesAccount = append(i.SecuritiesAccount, newValue)
	return newValue
}

func (i *InvestmentAccount35) AddBlockedStatus() *Blocked1 {
	i.BlockedStatus = new(Blocked1)
	return i.BlockedStatus
}

func (i *InvestmentAccount35) AddAccountUsageType() *AccountUsageType1Choice {
	i.AccountUsageType = new(AccountUsageType1Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount35) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount35) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount36 struct {

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, for example, wrapper, ISA.
	Type *AccountType1Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType1Choice `xml:"OwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason1Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason1Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Method by which the tax (withholding tax) is to be processed i.e. either withheld at source or tax information reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod2Code `xml:"TaxWhldgMtd,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	ModifiedFundDetails []*ModificationScope13 `xml:"ModfdFndDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Specifies information about blocked accounts.
	BlockedStatus *Blocked1 `xml:"BlckdSts,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType1Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`
}

func (i *InvestmentAccount36) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount36) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount36) AddType() *AccountType1Choice {
	i.Type = new(AccountType1Choice)
	return i.Type
}

func (i *InvestmentAccount36) AddOwnershipType() *OwnershipType1Choice {
	i.OwnershipType = new(OwnershipType1Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount36) AddTaxExemption() *TaxExemptionReason1Choice {
	i.TaxExemption = new(TaxExemptionReason1Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount36) AddStatementFrequency() *StatementFrequencyReason1Choice {
	i.StatementFrequency = new(StatementFrequencyReason1Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount36) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount36) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount36) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount36) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod2Code)(&value)
}

func (i *InvestmentAccount36) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount36) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount36) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount36) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount36) AddModifiedFundDetails() *ModificationScope13 {
	newValue := new(ModificationScope13)
	i.ModifiedFundDetails = append(i.ModifiedFundDetails, newValue)
	return newValue
}

func (i *InvestmentAccount36) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount36) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount36) AddBlockedStatus() *Blocked1 {
	i.BlockedStatus = new(Blocked1)
	return i.BlockedStatus
}

func (i *InvestmentAccount36) AddAccountUsageType() *AccountUsageType1Choice {
	i.AccountUsageType = new(AccountUsageType1Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount36) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount36) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount37 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, for example, wrapper, ISA.
	Type *AccountType1Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType1Choice `xml:"OwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason1Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason1Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Specifies, for income on the fund that is to be reinvested, parameters for the reinvestment. If the reinvestment percentage is less than one hundred percent, the remaining percentage will be invested according to the investor’s subsequent instructions.
	ReinvestmentDetails []*Reinvestment1 `xml:"RinvstmtDtls,omitempty"`

	// Method by which the tax (withholding tax) is to be processed i.e. either withheld at source or tax information reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod2Code `xml:"TaxWhldgMtd,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	FundDetails []*FinancialInstrument29 `xml:"FndDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Specifies information about blocked accounts.
	BlockedStatus []*Blocked1 `xml:"BlckdSts,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType1Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`

	// Specifies the means by which the investor submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Specifies the category of the investment account.
	InvestmentAccountCategory *InvestmentAccountCategory1Choice `xml:"InvstmtAcctCtgy,omitempty"`
}

func (i *InvestmentAccount37) AddIdentification() *AccountIdentification1 {
	i.Identification = new(AccountIdentification1)
	return i.Identification
}

func (i *InvestmentAccount37) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount37) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount37) AddType() *AccountType1Choice {
	i.Type = new(AccountType1Choice)
	return i.Type
}

func (i *InvestmentAccount37) AddOwnershipType() *OwnershipType1Choice {
	i.OwnershipType = new(OwnershipType1Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount37) AddTaxExemption() *TaxExemptionReason1Choice {
	i.TaxExemption = new(TaxExemptionReason1Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount37) AddStatementFrequency() *StatementFrequencyReason1Choice {
	i.StatementFrequency = new(StatementFrequencyReason1Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount37) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount37) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount37) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount37) AddReinvestmentDetails() *Reinvestment1 {
	newValue := new(Reinvestment1)
	i.ReinvestmentDetails = append(i.ReinvestmentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount37) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod2Code)(&value)
}

func (i *InvestmentAccount37) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount37) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount37) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount37) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount37) AddFundDetails() *FinancialInstrument29 {
	newValue := new(FinancialInstrument29)
	i.FundDetails = append(i.FundDetails, newValue)
	return newValue
}

func (i *InvestmentAccount37) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount37) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount37) AddBlockedStatus() *Blocked1 {
	newValue := new(Blocked1)
	i.BlockedStatus = append(i.BlockedStatus, newValue)
	return newValue
}

func (i *InvestmentAccount37) AddAccountUsageType() *AccountUsageType1Choice {
	i.AccountUsageType = new(AccountUsageType1Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount37) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount37) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

func (i *InvestmentAccount37) AddTransactionChannelType() *TransactionChannelType1Choice {
	i.TransactionChannelType = new(TransactionChannelType1Choice)
	return i.TransactionChannelType
}

func (i *InvestmentAccount37) AddInvestmentAccountCategory() *InvestmentAccountCategory1Choice {
	i.InvestmentAccountCategory = new(InvestmentAccountCategory1Choice)
	return i.InvestmentAccountCategory
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount38 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Specifies the current state of an account, eg, enabled or deleted.
	Status *AccountStatus2Code `xml:"Sts"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, for example, wrapper, ISA.
	Type *AccountType1Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType1Choice `xml:"OwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason1Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason1Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Specifies, for income on the fund that is to be reinvested, parameters for the reinvestment. If the reinvestment percentage is less than one hundred percent, the remaining percentage will be invested according to the investor’s subsequent instructions.
	ReinvestmentDetails []*Reinvestment1 `xml:"RinvstmtDtls,omitempty"`

	// Method by which the tax (withholding tax) is to be processed i.e. either withheld at source or tax information reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod2Code `xml:"TaxWhldgMtd,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Detailed information about the investment fund associated to the account.
	FundsDetails []*FinancialInstrument29 `xml:"FndsDtls,omitempty"`

	// Part of the investment account to or from which cash entries are made.
	CashAccount []*CashAccount12 `xml:"CshAcct,omitempty"`

	// Part of the investment account to or from which securities entries are made.
	SecuritiesAccount []*SecuritiesAccount4 `xml:"SctiesAcct,omitempty"`

	// Specifies information about blocked accounts.
	BlockedStatus *Blocked1 `xml:"BlckdSts,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType1Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`

	// Specifies the means by which the investor submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Specifies the category of the investment account.
	InvestmentAccountCategory *InvestmentAccountCategory1Choice `xml:"InvstmtAcctCtgy,omitempty"`
}

func (i *InvestmentAccount38) AddIdentification() *AccountIdentification1 {
	i.Identification = new(AccountIdentification1)
	return i.Identification
}

func (i *InvestmentAccount38) SetStatus(value string) {
	i.Status = (*AccountStatus2Code)(&value)
}

func (i *InvestmentAccount38) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount38) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount38) AddType() *AccountType1Choice {
	i.Type = new(AccountType1Choice)
	return i.Type
}

func (i *InvestmentAccount38) AddOwnershipType() *OwnershipType1Choice {
	i.OwnershipType = new(OwnershipType1Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount38) AddTaxExemption() *TaxExemptionReason1Choice {
	i.TaxExemption = new(TaxExemptionReason1Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount38) AddStatementFrequency() *StatementFrequencyReason1Choice {
	i.StatementFrequency = new(StatementFrequencyReason1Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount38) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount38) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount38) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount38) AddReinvestmentDetails() *Reinvestment1 {
	newValue := new(Reinvestment1)
	i.ReinvestmentDetails = append(i.ReinvestmentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount38) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod2Code)(&value)
}

func (i *InvestmentAccount38) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount38) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount38) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount38) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount38) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount38) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount38) AddFundsDetails() *FinancialInstrument29 {
	newValue := new(FinancialInstrument29)
	i.FundsDetails = append(i.FundsDetails, newValue)
	return newValue
}

func (i *InvestmentAccount38) AddCashAccount() *CashAccount12 {
	newValue := new(CashAccount12)
	i.CashAccount = append(i.CashAccount, newValue)
	return newValue
}

func (i *InvestmentAccount38) AddSecuritiesAccount() *SecuritiesAccount4 {
	newValue := new(SecuritiesAccount4)
	i.SecuritiesAccount = append(i.SecuritiesAccount, newValue)
	return newValue
}

func (i *InvestmentAccount38) AddBlockedStatus() *Blocked1 {
	i.BlockedStatus = new(Blocked1)
	return i.BlockedStatus
}

func (i *InvestmentAccount38) AddAccountUsageType() *AccountUsageType1Choice {
	i.AccountUsageType = new(AccountUsageType1Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount38) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount38) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

func (i *InvestmentAccount38) AddTransactionChannelType() *TransactionChannelType1Choice {
	i.TransactionChannelType = new(TransactionChannelType1Choice)
	return i.TransactionChannelType
}

func (i *InvestmentAccount38) AddInvestmentAccountCategory() *InvestmentAccountCategory1Choice {
	i.InvestmentAccountCategory = new(InvestmentAccountCategory1Choice)
	return i.InvestmentAccountCategory
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount39 struct {

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, for example, wrapper, ISA.
	Type *AccountType1Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType1Choice `xml:"OwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason1Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason1Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Specifies, for income on the fund that is to be reinvested, parameters for the reinvestment. If the reinvestment percentage is less than one hundred percent, the remaining percentage will be invested according to the investor’s subsequent instructions.
	ReinvestmentDetails []*Reinvestment1 `xml:"RinvstmtDtls,omitempty"`

	// Method by which the tax (withholding tax) is to be processed i.e. either withheld at source or tax information reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod2Code `xml:"TaxWhldgMtd,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	ModifiedFundDetails []*ModificationScope13 `xml:"ModfdFndDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Specifies information about blocked accounts.
	BlockedStatus *Blocked1 `xml:"BlckdSts,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType1Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`

	// Specifies the means by which the investor submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Specifies the category of the investment account.
	InvestmentAccountCategory *InvestmentAccountCategory1Choice `xml:"InvstmtAcctCtgy,omitempty"`
}

func (i *InvestmentAccount39) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount39) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount39) AddType() *AccountType1Choice {
	i.Type = new(AccountType1Choice)
	return i.Type
}

func (i *InvestmentAccount39) AddOwnershipType() *OwnershipType1Choice {
	i.OwnershipType = new(OwnershipType1Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount39) AddTaxExemption() *TaxExemptionReason1Choice {
	i.TaxExemption = new(TaxExemptionReason1Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount39) AddStatementFrequency() *StatementFrequencyReason1Choice {
	i.StatementFrequency = new(StatementFrequencyReason1Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount39) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount39) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount39) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount39) AddReinvestmentDetails() *Reinvestment1 {
	newValue := new(Reinvestment1)
	i.ReinvestmentDetails = append(i.ReinvestmentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount39) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod2Code)(&value)
}

func (i *InvestmentAccount39) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount39) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount39) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount39) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount39) AddModifiedFundDetails() *ModificationScope13 {
	newValue := new(ModificationScope13)
	i.ModifiedFundDetails = append(i.ModifiedFundDetails, newValue)
	return newValue
}

func (i *InvestmentAccount39) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount39) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount39) AddBlockedStatus() *Blocked1 {
	i.BlockedStatus = new(Blocked1)
	return i.BlockedStatus
}

func (i *InvestmentAccount39) AddAccountUsageType() *AccountUsageType1Choice {
	i.AccountUsageType = new(AccountUsageType1Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount39) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount39) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

func (i *InvestmentAccount39) AddTransactionChannelType() *TransactionChannelType1Choice {
	i.TransactionChannelType = new(TransactionChannelType1Choice)
	return i.TransactionChannelType
}

func (i *InvestmentAccount39) AddInvestmentAccountCategory() *InvestmentAccountCategory1Choice {
	i.InvestmentAccountCategory = new(InvestmentAccountCategory1Choice)
	return i.InvestmentAccountCategory
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount40 struct {

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification2Choice `xml:"OwnrId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary25 `xml:"IntrmyInf,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Indicates whether a security exists only as an electronic record, ie, there is no physical document representing the security.
	DematerialisedIndicator *YesNoIndicator `xml:"DmtrlsdInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Beneficial owner or its designated agent certifies that it complies with any holding or investment restrictions or requirements of the fund.
	BeneficiaryCertificationCompletion *BeneficiaryCertificationCompletion1Code `xml:"BnfcryCertfctnCmpltn,omitempty"`

	// Place requested as the place of safekeeping.
	SafekeepingPlace *PartyIdentification2Choice `xml:"SfkpgPlc,omitempty"`

	// Institution that maintains the records where the account is held.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`
}

func (i *InvestmentAccount40) AddOwnerIdentification() *PartyIdentification2Choice {
	newValue := new(PartyIdentification2Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount40) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount40) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount40) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount40) AddIntermediaryInformation() *Intermediary25 {
	newValue := new(Intermediary25)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount40) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentAccount40) SetDematerialisedIndicator(value string) {
	i.DematerialisedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount40) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount40) SetBeneficiaryCertificationCompletion(value string) {
	i.BeneficiaryCertificationCompletion = (*BeneficiaryCertificationCompletion1Code)(&value)
}

func (i *InvestmentAccount40) AddSafekeepingPlace() *PartyIdentification2Choice {
	i.SafekeepingPlace = new(PartyIdentification2Choice)
	return i.SafekeepingPlace
}

func (i *InvestmentAccount40) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount40) AddSubAccountDetails() *SubAccount1 {
	i.SubAccountDetails = new(SubAccount1)
	return i.SubAccountDetails
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount41 struct {

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification2Choice `xml:"OwnrId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary26 `xml:"IntrmyInf,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Indicates whether a security exists only as an electronic record, ie, there is no physical document representing the security.
	DematerialisedIndicator *YesNoIndicator `xml:"DmtrlsdInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Beneficial owner or its designated agent certifies that it complies with any holding or investment restrictions or requirements of the fund.
	BeneficiaryCertificationCompletion *BeneficiaryCertificationCompletion1Code `xml:"BnfcryCertfctnCmpltn,omitempty"`

	// Institution that maintains the records where the account is held.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`
}

func (i *InvestmentAccount41) AddOwnerIdentification() *PartyIdentification2Choice {
	newValue := new(PartyIdentification2Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount41) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount41) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount41) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount41) AddIntermediaryInformation() *Intermediary26 {
	newValue := new(Intermediary26)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount41) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentAccount41) SetDematerialisedIndicator(value string) {
	i.DematerialisedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount41) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount41) SetBeneficiaryCertificationCompletion(value string) {
	i.BeneficiaryCertificationCompletion = (*BeneficiaryCertificationCompletion1Code)(&value)
}

func (i *InvestmentAccount41) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount41) AddSubAccountDetails() *SubAccount1 {
	i.SubAccountDetails = new(SubAccount1)
	return i.SubAccountDetails
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount42 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Party that legally owns the account.
	OwnerIdentification *PartyIdentification2Choice `xml:"OwnrId,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount42) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount42) AddOwnerIdentification() *PartyIdentification2Choice {
	i.OwnerIdentification = new(PartyIdentification2Choice)
	return i.OwnerIdentification
}

func (i *InvestmentAccount42) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount43 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, for example, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary27 `xml:"IntrmyInf,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount43) AddIdentification() *AccountIdentification1 {
	i.Identification = new(AccountIdentification1)
	return i.Identification
}

func (i *InvestmentAccount43) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount43) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount43) AddIntermediaryInformation() *Intermediary27 {
	newValue := new(Intermediary27)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount43) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount45 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Identification of the owner of the account.
	OwnerIdentification *OwnerIdentification1Choice `xml:"OwnrId,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount45) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount45) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount45) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount45) AddOwnerIdentification() *OwnerIdentification1Choice {
	i.OwnerIdentification = new(OwnerIdentification1Choice)
	return i.OwnerIdentification
}

func (i *InvestmentAccount45) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount46 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Specifies the current state of an account, for example, enabled or deleted.
	Status *AccountStatus2Code `xml:"Sts"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, for example, wrapper, ISA.
	Type *AccountType1Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType1Choice `xml:"OwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason1Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason1Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Specifies, for income on the fund that is to be reinvested, parameters for the reinvestment. If the reinvestment percentage is less than one hundred percent, the remaining percentage will be invested according to the investor’s subsequent instructions.
	ReinvestmentDetails []*Reinvestment1 `xml:"RinvstmtDtls,omitempty"`

	// Method by which the tax (withholding tax) is to be processed i.e. either withheld at source or tax information reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod2Code `xml:"TaxWhldgMtd,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Detailed information about the investment fund associated to the account.
	FundsDetails []*FinancialInstrument29 `xml:"FndsDtls,omitempty"`

	// Part of the investment account to or from which cash entries are made.
	CashAccount []*CashAccount12 `xml:"CshAcct,omitempty"`

	// Part of the investment account to or from which securities entries are made.
	SecuritiesAccount []*SecuritiesAccount4 `xml:"SctiesAcct,omitempty"`

	// Specifies information about blocked accounts.
	BlockedStatus []*Blocked1 `xml:"BlckdSts,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType1Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`

	// Specifies the means by which the investor submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Specifies the category of the investment account.
	InvestmentAccountCategory *InvestmentAccountCategory1Choice `xml:"InvstmtAcctCtgy,omitempty"`
}

func (i *InvestmentAccount46) AddIdentification() *AccountIdentification1 {
	i.Identification = new(AccountIdentification1)
	return i.Identification
}

func (i *InvestmentAccount46) SetStatus(value string) {
	i.Status = (*AccountStatus2Code)(&value)
}

func (i *InvestmentAccount46) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount46) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount46) AddType() *AccountType1Choice {
	i.Type = new(AccountType1Choice)
	return i.Type
}

func (i *InvestmentAccount46) AddOwnershipType() *OwnershipType1Choice {
	i.OwnershipType = new(OwnershipType1Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount46) AddTaxExemption() *TaxExemptionReason1Choice {
	i.TaxExemption = new(TaxExemptionReason1Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount46) AddStatementFrequency() *StatementFrequencyReason1Choice {
	i.StatementFrequency = new(StatementFrequencyReason1Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount46) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount46) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount46) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount46) AddReinvestmentDetails() *Reinvestment1 {
	newValue := new(Reinvestment1)
	i.ReinvestmentDetails = append(i.ReinvestmentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount46) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod2Code)(&value)
}

func (i *InvestmentAccount46) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount46) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount46) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount46) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount46) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount46) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount46) AddFundsDetails() *FinancialInstrument29 {
	newValue := new(FinancialInstrument29)
	i.FundsDetails = append(i.FundsDetails, newValue)
	return newValue
}

func (i *InvestmentAccount46) AddCashAccount() *CashAccount12 {
	newValue := new(CashAccount12)
	i.CashAccount = append(i.CashAccount, newValue)
	return newValue
}

func (i *InvestmentAccount46) AddSecuritiesAccount() *SecuritiesAccount4 {
	newValue := new(SecuritiesAccount4)
	i.SecuritiesAccount = append(i.SecuritiesAccount, newValue)
	return newValue
}

func (i *InvestmentAccount46) AddBlockedStatus() *Blocked1 {
	newValue := new(Blocked1)
	i.BlockedStatus = append(i.BlockedStatus, newValue)
	return newValue
}

func (i *InvestmentAccount46) AddAccountUsageType() *AccountUsageType1Choice {
	i.AccountUsageType = new(AccountUsageType1Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount46) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount46) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

func (i *InvestmentAccount46) AddTransactionChannelType() *TransactionChannelType1Choice {
	i.TransactionChannelType = new(TransactionChannelType1Choice)
	return i.TransactionChannelType
}

func (i *InvestmentAccount46) AddInvestmentAccountCategory() *InvestmentAccountCategory1Choice {
	i.InvestmentAccountCategory = new(InvestmentAccountCategory1Choice)
	return i.InvestmentAccountCategory
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount47 struct {

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, for example, wrapper, ISA.
	Type *AccountType1Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType1Choice `xml:"OwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason1Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason1Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Specifies, for income on the fund that is to be reinvested, parameters for the reinvestment. If the reinvestment percentage is less than one hundred percent, the remaining percentage will be invested according to the investor’s subsequent instructions.
	ReinvestmentDetails []*Reinvestment1 `xml:"RinvstmtDtls,omitempty"`

	// Method by which the tax (withholding tax) is to be processed i.e. either withheld at source or tax information reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod2Code `xml:"TaxWhldgMtd,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	ModifiedFundDetails []*ModificationScope13 `xml:"ModfdFndDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Specifies information about blocked accounts.
	BlockedStatus []*Blocked1 `xml:"BlckdSts,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType1Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`

	// Specifies the means by which the investor submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Specifies the category of the investment account.
	InvestmentAccountCategory *InvestmentAccountCategory1Choice `xml:"InvstmtAcctCtgy,omitempty"`
}

func (i *InvestmentAccount47) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount47) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount47) AddType() *AccountType1Choice {
	i.Type = new(AccountType1Choice)
	return i.Type
}

func (i *InvestmentAccount47) AddOwnershipType() *OwnershipType1Choice {
	i.OwnershipType = new(OwnershipType1Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount47) AddTaxExemption() *TaxExemptionReason1Choice {
	i.TaxExemption = new(TaxExemptionReason1Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount47) AddStatementFrequency() *StatementFrequencyReason1Choice {
	i.StatementFrequency = new(StatementFrequencyReason1Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount47) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount47) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount47) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount47) AddReinvestmentDetails() *Reinvestment1 {
	newValue := new(Reinvestment1)
	i.ReinvestmentDetails = append(i.ReinvestmentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount47) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod2Code)(&value)
}

func (i *InvestmentAccount47) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount47) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount47) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount47) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount47) AddModifiedFundDetails() *ModificationScope13 {
	newValue := new(ModificationScope13)
	i.ModifiedFundDetails = append(i.ModifiedFundDetails, newValue)
	return newValue
}

func (i *InvestmentAccount47) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount47) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount47) AddBlockedStatus() *Blocked1 {
	newValue := new(Blocked1)
	i.BlockedStatus = append(i.BlockedStatus, newValue)
	return newValue
}

func (i *InvestmentAccount47) AddAccountUsageType() *AccountUsageType1Choice {
	i.AccountUsageType = new(AccountUsageType1Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount47) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount47) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

func (i *InvestmentAccount47) AddTransactionChannelType() *TransactionChannelType1Choice {
	i.TransactionChannelType = new(TransactionChannelType1Choice)
	return i.TransactionChannelType
}

func (i *InvestmentAccount47) AddInvestmentAccountCategory() *InvestmentAccountCategory1Choice {
	i.InvestmentAccountCategory = new(InvestmentAccountCategory1Choice)
	return i.InvestmentAccountCategory
}

// Information about a securities account and its characteristics.
type InvestmentAccount49 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Type of account.
	Type *AccountType2Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType2Choice `xml:"OwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason2Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason2Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Specifies, for income on the fund or security that is to be reinvested, parameters for the reinvestment. If the reinvestment percentage is less than one hundred percent, the remaining percentage will be invested according to the investor’s or account owner's subsequent instructions.
	ReinvestmentDetails []*Reinvestment2 `xml:"RinvstmtDtls,omitempty"`

	// Method by which the tax (withholding tax) is to be processed, that is,  either withheld at source or tax information is reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod3Code `xml:"TaxWhldgMtd,omitempty"`

	// Details for the reporting of tax, for example, the country of taxation.
	TaxReporting []*TaxReporting1 `xml:"TaxRptg,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund or security associated to the account.
	FinancialInstrumentDetails []*FinancialInstrument51 `xml:"FinInstrmDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Specifies the account is blocked and other factors for the blocked account.
	BlockedStatus []*Blocked2 `xml:"BlckdSts,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType2Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor or account owner signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`

	// Specifies the means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Specifies the category of the account.
	InvestmentAccountCategory *InvestmentAccountCategory1Choice `xml:"InvstmtAcctCtgy,omitempty"`

	// Specifies whether the holdings in the account are eligible for pledging.
	Pledging *Eligible1Code `xml:"Pldgg,omitempty"`

	// Specifies whether the holdings in the account are used as collateral.
	Collateral *Collateral1Code `xml:"Coll,omitempty"`

	// Details of third party rights.
	ThirdPartyRights *ThirdPartyRights1 `xml:"ThrdPtyRghts,omitempty"`

	// Functionality permitted to a third party in the actions that may be taken on an account on behalf of the investor.
	PowerOfAttorneyLevelOfControl *LevelOfControl1Choice `xml:"PwrOfAttnyLvlOfCtrl,omitempty"`

	// Specifies if the account is regarded as domestic or non-domestic for reporting purposes.
	AccountingStatus *AccountingStatus1Choice `xml:"AcctgSts,omitempty"`

	// Legal opening date for the account.
	OpeningDate *DateAndDateTimeChoice `xml:"OpngDt,omitempty"`

	// Legal closing date for the account.
	ClosingDate *DateAndDateTimeChoice `xml:"ClsgDt,omitempty"`

	// Indicates whether the account can hold a negative position in a security.
	NegativeIndicator *YesNoIndicator `xml:"NegInd,omitempty"`

	// Order in which securities are moved from the account.
	ProcessingOrder *PositionEffect3Code `xml:"PrcgOrdr,omitempty"`

	// Specifies whether the investor assumes responsibility for the liability.
	Liability *Liability1Choice `xml:"Lblty,omitempty"`

	// Information used to determine fees and types of operations that can be carried out on the account.
	InvestorProfile []*InvestorProfile1 `xml:"InvstrPrfl,omitempty"`

	// Fiscal year, when not the same as the calendar year.
	FiscalYear *FiscalYear1Choice `xml:"FsclYr,omitempty"`
}

func (i *InvestmentAccount49) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}

func (i *InvestmentAccount49) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount49) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount49) AddType() *AccountType2Choice {
	i.Type = new(AccountType2Choice)
	return i.Type
}

func (i *InvestmentAccount49) AddOwnershipType() *OwnershipType2Choice {
	i.OwnershipType = new(OwnershipType2Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount49) AddTaxExemption() *TaxExemptionReason2Choice {
	i.TaxExemption = new(TaxExemptionReason2Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount49) AddStatementFrequency() *StatementFrequencyReason2Choice {
	i.StatementFrequency = new(StatementFrequencyReason2Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount49) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount49) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount49) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentAccount49) AddReinvestmentDetails() *Reinvestment2 {
	newValue := new(Reinvestment2)
	i.ReinvestmentDetails = append(i.ReinvestmentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount49) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod3Code)(&value)
}

func (i *InvestmentAccount49) AddTaxReporting() *TaxReporting1 {
	newValue := new(TaxReporting1)
	i.TaxReporting = append(i.TaxReporting, newValue)
	return newValue
}

func (i *InvestmentAccount49) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount49) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount49) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount49) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount49) AddFinancialInstrumentDetails() *FinancialInstrument51 {
	newValue := new(FinancialInstrument51)
	i.FinancialInstrumentDetails = append(i.FinancialInstrumentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount49) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount49) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount49) AddBlockedStatus() *Blocked2 {
	newValue := new(Blocked2)
	i.BlockedStatus = append(i.BlockedStatus, newValue)
	return newValue
}

func (i *InvestmentAccount49) AddAccountUsageType() *AccountUsageType2Choice {
	i.AccountUsageType = new(AccountUsageType2Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount49) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount49) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

func (i *InvestmentAccount49) AddTransactionChannelType() *TransactionChannelType1Choice {
	i.TransactionChannelType = new(TransactionChannelType1Choice)
	return i.TransactionChannelType
}

func (i *InvestmentAccount49) AddInvestmentAccountCategory() *InvestmentAccountCategory1Choice {
	i.InvestmentAccountCategory = new(InvestmentAccountCategory1Choice)
	return i.InvestmentAccountCategory
}

func (i *InvestmentAccount49) SetPledging(value string) {
	i.Pledging = (*Eligible1Code)(&value)
}

func (i *InvestmentAccount49) SetCollateral(value string) {
	i.Collateral = (*Collateral1Code)(&value)
}

func (i *InvestmentAccount49) AddThirdPartyRights() *ThirdPartyRights1 {
	i.ThirdPartyRights = new(ThirdPartyRights1)
	return i.ThirdPartyRights
}

func (i *InvestmentAccount49) AddPowerOfAttorneyLevelOfControl() *LevelOfControl1Choice {
	i.PowerOfAttorneyLevelOfControl = new(LevelOfControl1Choice)
	return i.PowerOfAttorneyLevelOfControl
}

func (i *InvestmentAccount49) AddAccountingStatus() *AccountingStatus1Choice {
	i.AccountingStatus = new(AccountingStatus1Choice)
	return i.AccountingStatus
}

func (i *InvestmentAccount49) AddOpeningDate() *DateAndDateTimeChoice {
	i.OpeningDate = new(DateAndDateTimeChoice)
	return i.OpeningDate
}

func (i *InvestmentAccount49) AddClosingDate() *DateAndDateTimeChoice {
	i.ClosingDate = new(DateAndDateTimeChoice)
	return i.ClosingDate
}

func (i *InvestmentAccount49) SetNegativeIndicator(value string) {
	i.NegativeIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount49) SetProcessingOrder(value string) {
	i.ProcessingOrder = (*PositionEffect3Code)(&value)
}

func (i *InvestmentAccount49) AddLiability() *Liability1Choice {
	i.Liability = new(Liability1Choice)
	return i.Liability
}

func (i *InvestmentAccount49) AddInvestorProfile() *InvestorProfile1 {
	newValue := new(InvestorProfile1)
	i.InvestorProfile = append(i.InvestorProfile, newValue)
	return newValue
}

func (i *InvestmentAccount49) AddFiscalYear() *FiscalYear1Choice {
	i.FiscalYear = new(FiscalYear1Choice)
	return i.FiscalYear
}

// Information about a securities account and its characteristics.
type InvestmentAccount50 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Status of the account, for example, enabled or deleted.
	Status *AccountStatus1Choice `xml:"Sts"`

	// Date the status is assigned.
	StatusDate *DateAndDateTime1Choice `xml:"StsDt,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Type of account.
	Type *AccountType2Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType2Choice `xml:"OwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason2Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason2Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Specifies, for income on the fund or security that is to be reinvested, parameters for the reinvestment. If the reinvestment percentage is less than one hundred percent, the remaining percentage will be invested according to the investor’s or account owner's subsequent instructions.
	ReinvestmentDetails []*Reinvestment2 `xml:"RinvstmtDtls,omitempty"`

	// Method by which the tax (withholding tax) is to be processed, that is,  either withheld at source or tax information is reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod3Code `xml:"TaxWhldgMtd,omitempty"`

	// Details for the reporting of tax, for example, the country of taxation.
	TaxReporting []*TaxReporting1 `xml:"TaxRptg,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund or security associated to the account.
	FinancialInstrumentDetails []*FinancialInstrument51 `xml:"FinInstrmDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Specifies the account is blocked and other factors for the blocked account.
	BlockedStatus []*Blocked2 `xml:"BlckdSts,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType2Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor or account owner signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`

	// Specifies the means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Specifies the category of the account.
	InvestmentAccountCategory *InvestmentAccountCategory1Choice `xml:"InvstmtAcctCtgy,omitempty"`

	// Specifies whether the holdings in the account are eligible for pledging.
	Pledging *Eligible1Code `xml:"Pldgg,omitempty"`

	// Specifies whether the holdings in the account are used as collateral.
	Collateral *Collateral1Code `xml:"Coll,omitempty"`

	// Details of third party rights.
	ThirdPartyRights *ThirdPartyRights1 `xml:"ThrdPtyRghts,omitempty"`

	// Functionality permitted to a third party in the actions that may be taken on an account on behalf of the investor.
	PowerOfAttorneyLevelOfControl *LevelOfControl1Choice `xml:"PwrOfAttnyLvlOfCtrl,omitempty"`

	// Specifies if the account is regarded as domestic or non-domestic for reporting purposes.
	AccountingStatus *AccountingStatus1Choice `xml:"AcctgSts,omitempty"`

	// Legal opening date for the account.
	OpeningDate *DateAndDateTimeChoice `xml:"OpngDt,omitempty"`

	// Legal closing date for the account.
	ClosingDate *DateAndDateTimeChoice `xml:"ClsgDt,omitempty"`

	// Indicates whether the account can hold a negative position in a security.
	NegativeIndicator *YesNoIndicator `xml:"NegInd,omitempty"`

	// Order in which securities are moved from the account.
	ProcessingOrder *PositionEffect3Code `xml:"PrcgOrdr,omitempty"`

	// Specifies whether the investor assumes responsibility for the liability.
	Liability *Liability1Choice `xml:"Lblty,omitempty"`

	// Information used to determine fees and types of operations that can be carried out on the account.
	InvestorProfile []*InvestorProfile1 `xml:"InvstrPrfl,omitempty"`

	// Fiscal year, when not the same as the calendar year.
	FiscalYear *FiscalYear1Choice `xml:"FsclYr,omitempty"`
}

func (i *InvestmentAccount50) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}

func (i *InvestmentAccount50) AddStatus() *AccountStatus1Choice {
	i.Status = new(AccountStatus1Choice)
	return i.Status
}

func (i *InvestmentAccount50) AddStatusDate() *DateAndDateTime1Choice {
	i.StatusDate = new(DateAndDateTime1Choice)
	return i.StatusDate
}

func (i *InvestmentAccount50) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount50) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount50) AddType() *AccountType2Choice {
	i.Type = new(AccountType2Choice)
	return i.Type
}

func (i *InvestmentAccount50) AddOwnershipType() *OwnershipType2Choice {
	i.OwnershipType = new(OwnershipType2Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount50) AddTaxExemption() *TaxExemptionReason2Choice {
	i.TaxExemption = new(TaxExemptionReason2Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount50) AddStatementFrequency() *StatementFrequencyReason2Choice {
	i.StatementFrequency = new(StatementFrequencyReason2Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount50) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount50) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount50) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentAccount50) AddReinvestmentDetails() *Reinvestment2 {
	newValue := new(Reinvestment2)
	i.ReinvestmentDetails = append(i.ReinvestmentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount50) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod3Code)(&value)
}

func (i *InvestmentAccount50) AddTaxReporting() *TaxReporting1 {
	newValue := new(TaxReporting1)
	i.TaxReporting = append(i.TaxReporting, newValue)
	return newValue
}

func (i *InvestmentAccount50) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount50) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount50) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount50) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount50) AddFinancialInstrumentDetails() *FinancialInstrument51 {
	newValue := new(FinancialInstrument51)
	i.FinancialInstrumentDetails = append(i.FinancialInstrumentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount50) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount50) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount50) AddBlockedStatus() *Blocked2 {
	newValue := new(Blocked2)
	i.BlockedStatus = append(i.BlockedStatus, newValue)
	return newValue
}

func (i *InvestmentAccount50) AddAccountUsageType() *AccountUsageType2Choice {
	i.AccountUsageType = new(AccountUsageType2Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount50) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount50) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

func (i *InvestmentAccount50) AddTransactionChannelType() *TransactionChannelType1Choice {
	i.TransactionChannelType = new(TransactionChannelType1Choice)
	return i.TransactionChannelType
}

func (i *InvestmentAccount50) AddInvestmentAccountCategory() *InvestmentAccountCategory1Choice {
	i.InvestmentAccountCategory = new(InvestmentAccountCategory1Choice)
	return i.InvestmentAccountCategory
}

func (i *InvestmentAccount50) SetPledging(value string) {
	i.Pledging = (*Eligible1Code)(&value)
}

func (i *InvestmentAccount50) SetCollateral(value string) {
	i.Collateral = (*Collateral1Code)(&value)
}

func (i *InvestmentAccount50) AddThirdPartyRights() *ThirdPartyRights1 {
	i.ThirdPartyRights = new(ThirdPartyRights1)
	return i.ThirdPartyRights
}

func (i *InvestmentAccount50) AddPowerOfAttorneyLevelOfControl() *LevelOfControl1Choice {
	i.PowerOfAttorneyLevelOfControl = new(LevelOfControl1Choice)
	return i.PowerOfAttorneyLevelOfControl
}

func (i *InvestmentAccount50) AddAccountingStatus() *AccountingStatus1Choice {
	i.AccountingStatus = new(AccountingStatus1Choice)
	return i.AccountingStatus
}

func (i *InvestmentAccount50) AddOpeningDate() *DateAndDateTimeChoice {
	i.OpeningDate = new(DateAndDateTimeChoice)
	return i.OpeningDate
}

func (i *InvestmentAccount50) AddClosingDate() *DateAndDateTimeChoice {
	i.ClosingDate = new(DateAndDateTimeChoice)
	return i.ClosingDate
}

func (i *InvestmentAccount50) SetNegativeIndicator(value string) {
	i.NegativeIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount50) SetProcessingOrder(value string) {
	i.ProcessingOrder = (*PositionEffect3Code)(&value)
}

func (i *InvestmentAccount50) AddLiability() *Liability1Choice {
	i.Liability = new(Liability1Choice)
	return i.Liability
}

func (i *InvestmentAccount50) AddInvestorProfile() *InvestorProfile1 {
	newValue := new(InvestorProfile1)
	i.InvestorProfile = append(i.InvestorProfile, newValue)
	return newValue
}

func (i *InvestmentAccount50) AddFiscalYear() *FiscalYear1Choice {
	i.FiscalYear = new(FiscalYear1Choice)
	return i.FiscalYear
}

// Information about a securities account and its characteristics.
type InvestmentAccount51 struct {

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Type of account.
	Type *AccountType2Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType2Choice `xml:"OwnrshTp,omitempty"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason2Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason2Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Specifies, for income on the fund or security that is to be reinvested, parameters for the reinvestment. If the reinvestment percentage is less than one hundred percent, the remaining percentage will be invested according to the investor’s or account owner's subsequent instructions.
	ReinvestmentDetails []*Reinvestment2 `xml:"RinvstmtDtls,omitempty"`

	// Method by which the tax (withholding tax) is to be processed, that is,  either withheld at source or tax information is reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod3Code `xml:"TaxWhldgMtd,omitempty"`

	// Details for the reporting of tax, for example, the country of taxation.
	TaxReporting []*TaxReporting1 `xml:"TaxRptg,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the financial instrument associated to the account that is to be modified or deleted.
	ModifiedFinancialInstrumentDetails []*ModificationScope29 `xml:"ModfdFinInstrmDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Specifies the account is blocked and other factors for the blocked account.
	BlockedStatus []*Blocked2 `xml:"BlckdSts,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType2Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor or account owner signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`

	// Specifies the means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Specifies the category of the account.
	InvestmentAccountCategory *InvestmentAccountCategory1Choice `xml:"InvstmtAcctCtgy,omitempty"`

	// Specifies whether the holdings in the account are eligible for pledging.
	Pledging *Eligible1Code `xml:"Pldgg,omitempty"`

	// Specifies whether the holdings in the account are used as collateral.
	Collateral *Collateral1Code `xml:"Coll,omitempty"`

	// Details of third party rights.
	ThirdPartyRights *ThirdPartyRights1 `xml:"ThrdPtyRghts,omitempty"`

	// Functionality permitted to a third party in the actions that may be taken on an account on behalf of the investor.
	PowerOfAttorneyLevelOfControl *LevelOfControl1Choice `xml:"PwrOfAttnyLvlOfCtrl,omitempty"`

	// Specifies if the account is regarded as domestic or non-domestic for reporting purposes.
	AccountingStatus *AccountingStatus1Choice `xml:"AcctgSts,omitempty"`

	// Legal opening date for the account.
	OpeningDate *DateAndDateTimeChoice `xml:"OpngDt,omitempty"`

	// Legal closing date for the account.
	ClosingDate *DateAndDateTimeChoice `xml:"ClsgDt,omitempty"`

	// Indicates whether the account can hold a negative position in a security.
	NegativeIndicator *YesNoIndicator `xml:"NegInd,omitempty"`

	// Order in which securities are moved from the account.
	ProcessingOrder *PositionEffect3Code `xml:"PrcgOrdr,omitempty"`

	// Specifies whether the investor assumes responsibility for the liability.
	Liability *Liability1Choice `xml:"Lblty,omitempty"`

	// Information related to the investor profile to be inserted, updated or deleted.
	ModifiedInvestorProfile []*ModificationScope32 `xml:"ModfdInvstrPrfl,omitempty"`

	// Fiscal year, when not the same as the calendar year.
	FiscalYear *FiscalYear1Choice `xml:"FsclYr,omitempty"`
}

func (i *InvestmentAccount51) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount51) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount51) AddType() *AccountType2Choice {
	i.Type = new(AccountType2Choice)
	return i.Type
}

func (i *InvestmentAccount51) AddOwnershipType() *OwnershipType2Choice {
	i.OwnershipType = new(OwnershipType2Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount51) AddTaxExemption() *TaxExemptionReason2Choice {
	i.TaxExemption = new(TaxExemptionReason2Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount51) AddStatementFrequency() *StatementFrequencyReason2Choice {
	i.StatementFrequency = new(StatementFrequencyReason2Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount51) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount51) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount51) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentAccount51) AddReinvestmentDetails() *Reinvestment2 {
	newValue := new(Reinvestment2)
	i.ReinvestmentDetails = append(i.ReinvestmentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount51) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod3Code)(&value)
}

func (i *InvestmentAccount51) AddTaxReporting() *TaxReporting1 {
	newValue := new(TaxReporting1)
	i.TaxReporting = append(i.TaxReporting, newValue)
	return newValue
}

func (i *InvestmentAccount51) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount51) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount51) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount51) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount51) AddModifiedFinancialInstrumentDetails() *ModificationScope29 {
	newValue := new(ModificationScope29)
	i.ModifiedFinancialInstrumentDetails = append(i.ModifiedFinancialInstrumentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount51) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount51) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount51) AddBlockedStatus() *Blocked2 {
	newValue := new(Blocked2)
	i.BlockedStatus = append(i.BlockedStatus, newValue)
	return newValue
}

func (i *InvestmentAccount51) AddAccountUsageType() *AccountUsageType2Choice {
	i.AccountUsageType = new(AccountUsageType2Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount51) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount51) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

func (i *InvestmentAccount51) AddTransactionChannelType() *TransactionChannelType1Choice {
	i.TransactionChannelType = new(TransactionChannelType1Choice)
	return i.TransactionChannelType
}

func (i *InvestmentAccount51) AddInvestmentAccountCategory() *InvestmentAccountCategory1Choice {
	i.InvestmentAccountCategory = new(InvestmentAccountCategory1Choice)
	return i.InvestmentAccountCategory
}

func (i *InvestmentAccount51) SetPledging(value string) {
	i.Pledging = (*Eligible1Code)(&value)
}

func (i *InvestmentAccount51) SetCollateral(value string) {
	i.Collateral = (*Collateral1Code)(&value)
}

func (i *InvestmentAccount51) AddThirdPartyRights() *ThirdPartyRights1 {
	i.ThirdPartyRights = new(ThirdPartyRights1)
	return i.ThirdPartyRights
}

func (i *InvestmentAccount51) AddPowerOfAttorneyLevelOfControl() *LevelOfControl1Choice {
	i.PowerOfAttorneyLevelOfControl = new(LevelOfControl1Choice)
	return i.PowerOfAttorneyLevelOfControl
}

func (i *InvestmentAccount51) AddAccountingStatus() *AccountingStatus1Choice {
	i.AccountingStatus = new(AccountingStatus1Choice)
	return i.AccountingStatus
}

func (i *InvestmentAccount51) AddOpeningDate() *DateAndDateTimeChoice {
	i.OpeningDate = new(DateAndDateTimeChoice)
	return i.OpeningDate
}

func (i *InvestmentAccount51) AddClosingDate() *DateAndDateTimeChoice {
	i.ClosingDate = new(DateAndDateTimeChoice)
	return i.ClosingDate
}

func (i *InvestmentAccount51) SetNegativeIndicator(value string) {
	i.NegativeIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount51) SetProcessingOrder(value string) {
	i.ProcessingOrder = (*PositionEffect3Code)(&value)
}

func (i *InvestmentAccount51) AddLiability() *Liability1Choice {
	i.Liability = new(Liability1Choice)
	return i.Liability
}

func (i *InvestmentAccount51) AddModifiedInvestorProfile() *ModificationScope32 {
	newValue := new(ModificationScope32)
	i.ModifiedInvestorProfile = append(i.ModifiedInvestorProfile, newValue)
	return newValue
}

func (i *InvestmentAccount51) AddFiscalYear() *FiscalYear1Choice {
	i.FiscalYear = new(FiscalYear1Choice)
	return i.FiscalYear
}

// Information about a securities account and its characteristics.
type InvestmentAccount52 struct {

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Legal form of the fund, for example, UCITS, SICAV, OEIC, Unit Trust, and FCP.
	FundType *Max35Text `xml:"FndTp,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	SecurityDetails *FinancialInstrument45 `xml:"SctyDtls,omitempty"`

	// Owner of the account.
	AccountOwner *AccountOwner1Choice `xml:"AcctOwnr,omitempty"`

	// Intermediary or other party related to the management of the account. In some markets, when this intermediary is a party acting on behalf of the investor for which it has opened an account at, for example, a central securities depository or international central securities depository, this party is known by the investor as the 'account controller'.
	Intermediary []*Intermediary33 `xml:"Intrmy,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount52) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount52) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount52) SetFundType(value string) {
	i.FundType = (*Max35Text)(&value)
}

func (i *InvestmentAccount52) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount52) AddSecurityDetails() *FinancialInstrument45 {
	i.SecurityDetails = new(FinancialInstrument45)
	return i.SecurityDetails
}

func (i *InvestmentAccount52) AddAccountOwner() *AccountOwner1Choice {
	i.AccountOwner = new(AccountOwner1Choice)
	return i.AccountOwner
}

func (i *InvestmentAccount52) AddIntermediary() *Intermediary33 {
	newValue := new(Intermediary33)
	i.Intermediary = append(i.Intermediary, newValue)
	return newValue
}

func (i *InvestmentAccount52) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

// Information about a securities account and its characteristics.
type InvestmentAccount53 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Identification of the owner of the account.
	OwnerIdentification *OwnerIdentification2Choice `xml:"OwnrId,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount53) SetAccountIdentification(value string) {
	i.AccountIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccount53) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount53) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount53) AddOwnerIdentification() *OwnerIdentification2Choice {
	i.OwnerIdentification = new(OwnerIdentification2Choice)
	return i.OwnerIdentification
}

func (i *InvestmentAccount53) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount54 struct {

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification70Choice `xml:"OwnrId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Identification of a related party or intermediary.
	IntermediaryInformation []*Intermediary34 `xml:"IntrmyInf,omitempty"`

	// Form, that is, ownership, of the security, that is, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Indicates whether a security exists only as an electronic record, that is, there is no physical document representing the security.
	DematerialisedIndicator *YesNoIndicator `xml:"DmtrlsdInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Beneficial owner or its designated agent certifies that it complies with any holding or investment restrictions or requirements of the fund.
	BeneficiaryCertificationCompletion *BeneficiaryCertificationCompletion1Code `xml:"BnfcryCertfctnCmpltn,omitempty"`

	// Place requested as the place of safekeeping.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Institution that maintains the records where the account is held.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount5 `xml:"SubAcctDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount14 `xml:"SttlmPtiesDtls,omitempty"`
}

func (i *InvestmentAccount54) AddOwnerIdentification() *PartyIdentification70Choice {
	newValue := new(PartyIdentification70Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount54) SetAccountIdentification(value string) {
	i.AccountIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccount54) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount54) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount54) AddIntermediaryInformation() *Intermediary34 {
	newValue := new(Intermediary34)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount54) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentAccount54) SetDematerialisedIndicator(value string) {
	i.DematerialisedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount54) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentAccount54) SetBeneficiaryCertificationCompletion(value string) {
	i.BeneficiaryCertificationCompletion = (*BeneficiaryCertificationCompletion1Code)(&value)
}

func (i *InvestmentAccount54) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return i.SafekeepingPlace
}

func (i *InvestmentAccount54) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount54) AddSubAccountDetails() *SubAccount5 {
	i.SubAccountDetails = new(SubAccount5)
	return i.SubAccountDetails
}

func (i *InvestmentAccount54) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount14 {
	i.SettlementPartiesDetails = new(DeliveringPartiesAndAccount14)
	return i.SettlementPartiesDetails
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount55 struct {

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification70Choice `xml:"OwnrId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Form, that is, ownership, of the security, that is, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Indicates whether a security exists only as an electronic record, that is, there is no physical document representing the security.
	DematerialisedIndicator *YesNoIndicator `xml:"DmtrlsdInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Beneficial owner or its designated agent certifies that it complies with any holding or investment restrictions or requirements of the fund.
	BeneficiaryCertificationCompletion *BeneficiaryCertificationCompletion1Code `xml:"BnfcryCertfctnCmpltn,omitempty"`

	// Institution that maintains the records where the account is held.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount5 `xml:"SubAcctDtls,omitempty"`
}

func (i *InvestmentAccount55) AddOwnerIdentification() *PartyIdentification70Choice {
	newValue := new(PartyIdentification70Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount55) SetAccountIdentification(value string) {
	i.AccountIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccount55) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount55) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount55) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentAccount55) SetDematerialisedIndicator(value string) {
	i.DematerialisedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount55) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentAccount55) SetBeneficiaryCertificationCompletion(value string) {
	i.BeneficiaryCertificationCompletion = (*BeneficiaryCertificationCompletion1Code)(&value)
}

func (i *InvestmentAccount55) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount55) AddSubAccountDetails() *SubAccount5 {
	i.SubAccountDetails = new(SubAccount5)
	return i.SubAccountDetails
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount56 struct {

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification70Choice `xml:"OwnrId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, for example, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary34 `xml:"IntrmyInf,omitempty"`

	// Form, that is, ownership, of the security, that is, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Indicates whether a security exists only as an electronic record, that is, there is no physical document representing the security.
	DematerialisedIndicator *YesNoIndicator `xml:"DmtrlsdInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Beneficial owner or its designated agent certifies that it complies with any holding or investment restrictions or requirements of the fund.
	BeneficiaryCertificationCompletion *BeneficiaryCertificationCompletion1Code `xml:"BnfcryCertfctnCmpltn,omitempty"`

	// Place requested as the place of safekeeping.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Institution that maintains the records where the account is held.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount5 `xml:"SubAcctDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount14 `xml:"SttlmPtiesDtls,omitempty"`
}

func (i *InvestmentAccount56) AddOwnerIdentification() *PartyIdentification70Choice {
	newValue := new(PartyIdentification70Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount56) SetAccountIdentification(value string) {
	i.AccountIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccount56) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount56) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount56) AddIntermediaryInformation() *Intermediary34 {
	newValue := new(Intermediary34)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount56) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentAccount56) SetDematerialisedIndicator(value string) {
	i.DematerialisedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount56) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentAccount56) SetBeneficiaryCertificationCompletion(value string) {
	i.BeneficiaryCertificationCompletion = (*BeneficiaryCertificationCompletion1Code)(&value)
}

func (i *InvestmentAccount56) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return i.SafekeepingPlace
}

func (i *InvestmentAccount56) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount56) AddSubAccountDetails() *SubAccount5 {
	i.SubAccountDetails = new(SubAccount5)
	return i.SubAccountDetails
}

func (i *InvestmentAccount56) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount14 {
	i.SettlementPartiesDetails = new(ReceivingPartiesAndAccount14)
	return i.SettlementPartiesDetails
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount57 struct {

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification70Choice `xml:"OwnrId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, for example, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary35 `xml:"IntrmyInf,omitempty"`

	// Form, that is, ownership, of the security, that is, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Indicates whether a security exists only as an electronic record, that is, there is no physical document representing the security.
	DematerialisedIndicator *YesNoIndicator `xml:"DmtrlsdInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Beneficial owner or its designated agent certifies that it complies with any holding or investment restrictions or requirements of the fund.
	BeneficiaryCertificationCompletion *BeneficiaryCertificationCompletion1Code `xml:"BnfcryCertfctnCmpltn,omitempty"`

	// Place requested as the place of safekeeping.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Institution that maintains the records where the account is held.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount5 `xml:"SubAcctDtls,omitempty"`
}

func (i *InvestmentAccount57) AddOwnerIdentification() *PartyIdentification70Choice {
	newValue := new(PartyIdentification70Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount57) SetAccountIdentification(value string) {
	i.AccountIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccount57) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount57) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount57) AddIntermediaryInformation() *Intermediary35 {
	newValue := new(Intermediary35)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount57) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentAccount57) SetDematerialisedIndicator(value string) {
	i.DematerialisedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount57) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentAccount57) SetBeneficiaryCertificationCompletion(value string) {
	i.BeneficiaryCertificationCompletion = (*BeneficiaryCertificationCompletion1Code)(&value)
}

func (i *InvestmentAccount57) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return i.SafekeepingPlace
}

func (i *InvestmentAccount57) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount57) AddSubAccountDetails() *SubAccount5 {
	i.SubAccountDetails = new(SubAccount5)
	return i.SubAccountDetails
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount58 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification113 `xml:"OwnrId,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification113 `xml:"AcctSvcr,omitempty"`

	// Counterparties eligibility as defined by article 24 of the EU MiFID Directive applicable to transactions executed by investment firms for eligible counterparties.
	OrderOriginatorEligibility *OrderOriginatorEligibility1Code `xml:"OrdrOrgtrElgblty,omitempty"`

	// Sub-account of the master or omnibus account.
	SubAccountDetails *SubAccount6 `xml:"SubAcctDtls,omitempty"`
}

func (i *InvestmentAccount58) SetAccountIdentification(value string) {
	i.AccountIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccount58) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount58) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount58) AddOwnerIdentification() *PartyIdentification113 {
	newValue := new(PartyIdentification113)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount58) AddAccountServicer() *PartyIdentification113 {
	i.AccountServicer = new(PartyIdentification113)
	return i.AccountServicer
}

func (i *InvestmentAccount58) SetOrderOriginatorEligibility(value string) {
	i.OrderOriginatorEligibility = (*OrderOriginatorEligibility1Code)(&value)
}

func (i *InvestmentAccount58) AddSubAccountDetails() *SubAccount6 {
	i.SubAccountDetails = new(SubAccount6)
	return i.SubAccountDetails
}

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount60 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Account type.
	Type *InvestmentAccountType1Choice `xml:"Tp,omitempty"`
}

func (i *InvestmentAccount60) SetAccountIdentification(value string) {
	i.AccountIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccount60) AddType() *InvestmentAccountType1Choice {
	i.Type = new(InvestmentAccountType1Choice)
	return i.Type
}

// Information about a securities account and its characteristics.
type InvestmentAccount61 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Type of account.
	Type *AccountType2Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType2Choice `xml:"OwnrshTp"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason2Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason2Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Specifies, for income on the fund or security that is to be reinvested, parameters for the reinvestment. If the reinvestment percentage is less than one hundred percent, the remaining percentage will be invested according to the investor’s or account owner's subsequent instructions.
	ReinvestmentDetails []*Reinvestment3 `xml:"RinvstmtDtls,omitempty"`

	// Method by which the tax (withholding tax) is to be processed, that is,  either withheld at source or tax information is reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod3Code `xml:"TaxWhldgMtd,omitempty"`

	// Details for the reporting of tax, for example, the country of taxation.
	TaxReporting []*TaxReporting2 `xml:"TaxRptg,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund or security associated to the account.
	FinancialInstrumentDetails []*FinancialInstrument56 `xml:"FinInstrmDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Specifies the account is blocked and other factors for the blocked account.
	BlockedStatus *BlockedStatusReason2Choice `xml:"BlckdSts,omitempty"`

	// Type of usage of the account.
	AccountUsageType *AccountUsageType2Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor or account owner signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Category of the account.
	InvestmentAccountCategory *InvestmentAccountCategory1Choice `xml:"InvstmtAcctCtgy,omitempty"`

	// Specifies whether the holdings in the account are eligible for pledging.
	Pledging *Eligible1Code `xml:"Pldgg,omitempty"`

	// Specifies whether the holdings in the account are used as collateral.
	Collateral *Collateral1Code `xml:"Coll,omitempty"`

	// Details of third party rights.
	ThirdPartyRights *ThirdPartyRights1 `xml:"ThrdPtyRghts,omitempty"`

	// Functionality permitted to a third party in the actions that may be taken on an account on behalf of the investor.
	PowerOfAttorneyLevelOfControl *LevelOfControl1Choice `xml:"PwrOfAttnyLvlOfCtrl,omitempty"`

	// Specifies if the account is regarded as domestic or non-domestic for reporting purposes.
	AccountingStatus *AccountingStatus1Choice `xml:"AcctgSts,omitempty"`

	// Legal opening date for the account.
	OpeningDate *DateAndDateTimeChoice `xml:"OpngDt,omitempty"`

	// Legal closing date for the account.
	ClosingDate *DateAndDateTimeChoice `xml:"ClsgDt,omitempty"`

	// Indicates whether the account can hold a negative position in a security.
	NegativeIndicator *YesNoIndicator `xml:"NegInd,omitempty"`

	// Order in which securities are moved from the account.
	ProcessingOrder *PositionEffect3Code `xml:"PrcgOrdr,omitempty"`

	// Specifies whether the investor assumes responsibility for the liability.
	Liability *Liability1Choice `xml:"Lblty,omitempty"`

	// Information used to determine fees and types of operations that can be carried out on the account.
	InvestorProfile []*InvestorProfile1 `xml:"InvstrPrfl,omitempty"`

	// Fiscal year, when not the same as the calendar year.
	FiscalYear *FiscalYear1Choice `xml:"FsclYr,omitempty"`
}

func (i *InvestmentAccount61) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}

func (i *InvestmentAccount61) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount61) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount61) AddType() *AccountType2Choice {
	i.Type = new(AccountType2Choice)
	return i.Type
}

func (i *InvestmentAccount61) AddOwnershipType() *OwnershipType2Choice {
	i.OwnershipType = new(OwnershipType2Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount61) AddTaxExemption() *TaxExemptionReason2Choice {
	i.TaxExemption = new(TaxExemptionReason2Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount61) AddStatementFrequency() *StatementFrequencyReason2Choice {
	i.StatementFrequency = new(StatementFrequencyReason2Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount61) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount61) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount61) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentAccount61) AddReinvestmentDetails() *Reinvestment3 {
	newValue := new(Reinvestment3)
	i.ReinvestmentDetails = append(i.ReinvestmentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount61) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod3Code)(&value)
}

func (i *InvestmentAccount61) AddTaxReporting() *TaxReporting2 {
	newValue := new(TaxReporting2)
	i.TaxReporting = append(i.TaxReporting, newValue)
	return newValue
}

func (i *InvestmentAccount61) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount61) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount61) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount61) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount61) AddFinancialInstrumentDetails() *FinancialInstrument56 {
	newValue := new(FinancialInstrument56)
	i.FinancialInstrumentDetails = append(i.FinancialInstrumentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount61) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount61) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount61) AddBlockedStatus() *BlockedStatusReason2Choice {
	i.BlockedStatus = new(BlockedStatusReason2Choice)
	return i.BlockedStatus
}

func (i *InvestmentAccount61) AddAccountUsageType() *AccountUsageType2Choice {
	i.AccountUsageType = new(AccountUsageType2Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount61) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount61) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

func (i *InvestmentAccount61) AddTransactionChannelType() *TransactionChannelType1Choice {
	i.TransactionChannelType = new(TransactionChannelType1Choice)
	return i.TransactionChannelType
}

func (i *InvestmentAccount61) AddInvestmentAccountCategory() *InvestmentAccountCategory1Choice {
	i.InvestmentAccountCategory = new(InvestmentAccountCategory1Choice)
	return i.InvestmentAccountCategory
}

func (i *InvestmentAccount61) SetPledging(value string) {
	i.Pledging = (*Eligible1Code)(&value)
}

func (i *InvestmentAccount61) SetCollateral(value string) {
	i.Collateral = (*Collateral1Code)(&value)
}

func (i *InvestmentAccount61) AddThirdPartyRights() *ThirdPartyRights1 {
	i.ThirdPartyRights = new(ThirdPartyRights1)
	return i.ThirdPartyRights
}

func (i *InvestmentAccount61) AddPowerOfAttorneyLevelOfControl() *LevelOfControl1Choice {
	i.PowerOfAttorneyLevelOfControl = new(LevelOfControl1Choice)
	return i.PowerOfAttorneyLevelOfControl
}

func (i *InvestmentAccount61) AddAccountingStatus() *AccountingStatus1Choice {
	i.AccountingStatus = new(AccountingStatus1Choice)
	return i.AccountingStatus
}

func (i *InvestmentAccount61) AddOpeningDate() *DateAndDateTimeChoice {
	i.OpeningDate = new(DateAndDateTimeChoice)
	return i.OpeningDate
}

func (i *InvestmentAccount61) AddClosingDate() *DateAndDateTimeChoice {
	i.ClosingDate = new(DateAndDateTimeChoice)
	return i.ClosingDate
}

func (i *InvestmentAccount61) SetNegativeIndicator(value string) {
	i.NegativeIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount61) SetProcessingOrder(value string) {
	i.ProcessingOrder = (*PositionEffect3Code)(&value)
}

func (i *InvestmentAccount61) AddLiability() *Liability1Choice {
	i.Liability = new(Liability1Choice)
	return i.Liability
}

func (i *InvestmentAccount61) AddInvestorProfile() *InvestorProfile1 {
	newValue := new(InvestorProfile1)
	i.InvestorProfile = append(i.InvestorProfile, newValue)
	return newValue
}

func (i *InvestmentAccount61) AddFiscalYear() *FiscalYear1Choice {
	i.FiscalYear = new(FiscalYear1Choice)
	return i.FiscalYear
}

// Information about a securities account and its characteristics.
type InvestmentAccount62 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Status of the account.
	AccountStatus *AccountStatus2 `xml:"AcctSts,omitempty"`

	// Specifies the account is blocked and other factors for the blocked account.
	BlockedStatus *BlockedStatusReason2Choice `xml:"BlckdSts,omitempty"`

	// Date the status is assigned.
	StatusDate *DateAndDateTime1Choice `xml:"StsDt,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Type of account.
	Type *AccountType2Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType2Choice `xml:"OwnrshTp,omitempty"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason2Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason2Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Specifies, for income on the fund or security that is to be reinvested, parameters for the reinvestment. If the reinvestment percentage is less than one hundred percent, the remaining percentage will be invested according to the investor’s or account owner's subsequent instructions.
	ReinvestmentDetails []*Reinvestment3 `xml:"RinvstmtDtls,omitempty"`

	// Method by which the tax (withholding tax) is to be processed, that is,  either withheld at source or tax information is reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod3Code `xml:"TaxWhldgMtd,omitempty"`

	// Details for the reporting of tax, for example, the country of taxation.
	TaxReporting []*TaxReporting2 `xml:"TaxRptg,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund or security associated to the account.
	FinancialInstrumentDetails []*FinancialInstrument56 `xml:"FinInstrmDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType2Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor or account owner signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Specifies the category of the account.
	InvestmentAccountCategory *InvestmentAccountCategory1Choice `xml:"InvstmtAcctCtgy,omitempty"`

	// Specifies whether the holdings in the account are eligible for pledging.
	Pledging *Eligible1Code `xml:"Pldgg,omitempty"`

	// Specifies whether the holdings in the account are used as collateral.
	Collateral *Collateral1Code `xml:"Coll,omitempty"`

	// Details of third party rights.
	ThirdPartyRights *ThirdPartyRights1 `xml:"ThrdPtyRghts,omitempty"`

	// Functionality permitted to a third party in the actions that may be taken on an account on behalf of the investor.
	PowerOfAttorneyLevelOfControl *LevelOfControl1Choice `xml:"PwrOfAttnyLvlOfCtrl,omitempty"`

	// Specifies if the account is regarded as domestic or non-domestic for reporting purposes.
	AccountingStatus *AccountingStatus1Choice `xml:"AcctgSts,omitempty"`

	// Legal opening date for the account.
	OpeningDate *DateAndDateTimeChoice `xml:"OpngDt,omitempty"`

	// Legal closing date for the account.
	ClosingDate *DateAndDateTimeChoice `xml:"ClsgDt,omitempty"`

	// Indicates whether the account can hold a negative position in a security.
	NegativeIndicator *YesNoIndicator `xml:"NegInd,omitempty"`

	// Order in which securities are moved from the account.
	ProcessingOrder *PositionEffect3Code `xml:"PrcgOrdr,omitempty"`

	// Specifies whether the investor assumes responsibility for the liability.
	Liability *Liability1Choice `xml:"Lblty,omitempty"`

	// Information used to determine fees and types of operations that can be carried out on the account.
	InvestorProfile []*InvestorProfile1 `xml:"InvstrPrfl,omitempty"`

	// Fiscal year, when not the same as the calendar year.
	FiscalYear *FiscalYear1Choice `xml:"FsclYr,omitempty"`
}

func (i *InvestmentAccount62) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}

func (i *InvestmentAccount62) AddAccountStatus() *AccountStatus2 {
	i.AccountStatus = new(AccountStatus2)
	return i.AccountStatus
}

func (i *InvestmentAccount62) AddBlockedStatus() *BlockedStatusReason2Choice {
	i.BlockedStatus = new(BlockedStatusReason2Choice)
	return i.BlockedStatus
}

func (i *InvestmentAccount62) AddStatusDate() *DateAndDateTime1Choice {
	i.StatusDate = new(DateAndDateTime1Choice)
	return i.StatusDate
}

func (i *InvestmentAccount62) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount62) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount62) AddType() *AccountType2Choice {
	i.Type = new(AccountType2Choice)
	return i.Type
}

func (i *InvestmentAccount62) AddOwnershipType() *OwnershipType2Choice {
	i.OwnershipType = new(OwnershipType2Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount62) AddTaxExemption() *TaxExemptionReason2Choice {
	i.TaxExemption = new(TaxExemptionReason2Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount62) AddStatementFrequency() *StatementFrequencyReason2Choice {
	i.StatementFrequency = new(StatementFrequencyReason2Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount62) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount62) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount62) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentAccount62) AddReinvestmentDetails() *Reinvestment3 {
	newValue := new(Reinvestment3)
	i.ReinvestmentDetails = append(i.ReinvestmentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount62) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod3Code)(&value)
}

func (i *InvestmentAccount62) AddTaxReporting() *TaxReporting2 {
	newValue := new(TaxReporting2)
	i.TaxReporting = append(i.TaxReporting, newValue)
	return newValue
}

func (i *InvestmentAccount62) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount62) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount62) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount62) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount62) AddFinancialInstrumentDetails() *FinancialInstrument56 {
	newValue := new(FinancialInstrument56)
	i.FinancialInstrumentDetails = append(i.FinancialInstrumentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount62) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount62) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount62) AddAccountUsageType() *AccountUsageType2Choice {
	i.AccountUsageType = new(AccountUsageType2Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount62) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount62) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

func (i *InvestmentAccount62) AddTransactionChannelType() *TransactionChannelType1Choice {
	i.TransactionChannelType = new(TransactionChannelType1Choice)
	return i.TransactionChannelType
}

func (i *InvestmentAccount62) AddInvestmentAccountCategory() *InvestmentAccountCategory1Choice {
	i.InvestmentAccountCategory = new(InvestmentAccountCategory1Choice)
	return i.InvestmentAccountCategory
}

func (i *InvestmentAccount62) SetPledging(value string) {
	i.Pledging = (*Eligible1Code)(&value)
}

func (i *InvestmentAccount62) SetCollateral(value string) {
	i.Collateral = (*Collateral1Code)(&value)
}

func (i *InvestmentAccount62) AddThirdPartyRights() *ThirdPartyRights1 {
	i.ThirdPartyRights = new(ThirdPartyRights1)
	return i.ThirdPartyRights
}

func (i *InvestmentAccount62) AddPowerOfAttorneyLevelOfControl() *LevelOfControl1Choice {
	i.PowerOfAttorneyLevelOfControl = new(LevelOfControl1Choice)
	return i.PowerOfAttorneyLevelOfControl
}

func (i *InvestmentAccount62) AddAccountingStatus() *AccountingStatus1Choice {
	i.AccountingStatus = new(AccountingStatus1Choice)
	return i.AccountingStatus
}

func (i *InvestmentAccount62) AddOpeningDate() *DateAndDateTimeChoice {
	i.OpeningDate = new(DateAndDateTimeChoice)
	return i.OpeningDate
}

func (i *InvestmentAccount62) AddClosingDate() *DateAndDateTimeChoice {
	i.ClosingDate = new(DateAndDateTimeChoice)
	return i.ClosingDate
}

func (i *InvestmentAccount62) SetNegativeIndicator(value string) {
	i.NegativeIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount62) SetProcessingOrder(value string) {
	i.ProcessingOrder = (*PositionEffect3Code)(&value)
}

func (i *InvestmentAccount62) AddLiability() *Liability1Choice {
	i.Liability = new(Liability1Choice)
	return i.Liability
}

func (i *InvestmentAccount62) AddInvestorProfile() *InvestorProfile1 {
	newValue := new(InvestorProfile1)
	i.InvestorProfile = append(i.InvestorProfile, newValue)
	return newValue
}

func (i *InvestmentAccount62) AddFiscalYear() *FiscalYear1Choice {
	i.FiscalYear = new(FiscalYear1Choice)
	return i.FiscalYear
}

// Information about a securities account and its characteristics.
type InvestmentAccount63 struct {

	// Change of account status is instructed.
	AccountStatusUpdateInstruction *AccountStatusUpdateInstruction1 `xml:"AcctStsUpdInstr,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Type of account.
	Type *AccountType2Choice `xml:"Tp,omitempty"`

	// Ownership status of the account, for example, joint owners.
	OwnershipType *OwnershipType2Choice `xml:"OwnrshTp,omitempty"`

	// Tax advantage specific to the account.
	TaxExemption *TaxExemptionReason2Choice `xml:"TaxXmptn,omitempty"`

	// Frequency at which a statement is issued.
	StatementFrequency *StatementFrequencyReason2Choice `xml:"StmtFrqcy,omitempty"`

	// Currency chosen for reporting purposes by the account owner in agreement with the account servicer.
	ReferenceCurrency *ActiveCurrencyCode `xml:"RefCcy,omitempty"`

	// Language for all communication concerning the account.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Specifies, for income on the fund or security that is to be reinvested, parameters for the reinvestment. If the reinvestment percentage is less than one hundred percent, the remaining percentage will be invested according to the investor’s or account owner's subsequent instructions.
	ReinvestmentDetails []*Reinvestment3 `xml:"RinvstmtDtls,omitempty"`

	// Method by which the tax (withholding tax) is to be processed, that is,  either withheld at source or tax information is reported to tax authorities or tax information is reported due to the provision of a tax certificate.
	TaxWithholdingMethod *TaxWithholdingMethod3Code `xml:"TaxWhldgMtd,omitempty"`

	// Details for the reporting of tax, for example, the country of taxation.
	TaxReporting []*TaxReporting2 `xml:"TaxRptg,omitempty"`

	// Details of the letter of intent.
	LetterIntentDetails *LetterIntent1 `xml:"LttrInttDtls,omitempty"`

	// Reference of an accumulation rights program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Number of account owners or related parties required to authorise transactions on the account.
	RequiredSignatoriesNumber *Number `xml:"ReqrdSgntriesNb,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the financial instrument associated to the account that is to be modified or deleted.
	ModifiedFinancialInstrumentDetails []*ModificationScope36 `xml:"ModfdFinInstrmDtls,omitempty"`

	// Parameters to be applied on deal amount for orders when the amount is a fractional number.
	RoundingDetails *RoundingParameters1 `xml:"RndgDtls,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Specifies the account is blocked and other factors for the blocked account.
	BlockedStatus *BlockedStatusReason2Choice `xml:"BlckdSts,omitempty"`

	// Specifies the type of usage of the account.
	AccountUsageType *AccountUsageType2Choice `xml:"AcctUsgTp,omitempty"`

	// Specifies if documentary evidence has been provided for the foreign resident.
	ForeignStatusCertification *Provided1Code `xml:"FrgnStsCertfctn,omitempty"`

	// Date the investor or account owner signs the open account form.
	AccountSignatureDateTime *DateAndDateTimeChoice `xml:"AcctSgntrDtTm,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Specifies the category of the account.
	InvestmentAccountCategory *InvestmentAccountCategory1Choice `xml:"InvstmtAcctCtgy,omitempty"`

	// Specifies whether the holdings in the account are eligible for pledging.
	Pledging *Eligible1Code `xml:"Pldgg,omitempty"`

	// Specifies whether the holdings in the account are used as collateral.
	Collateral *Collateral1Code `xml:"Coll,omitempty"`

	// Details of third party rights.
	ThirdPartyRights *ThirdPartyRights1 `xml:"ThrdPtyRghts,omitempty"`

	// Functionality permitted to a third party in the actions that may be taken on an account on behalf of the investor.
	PowerOfAttorneyLevelOfControl *LevelOfControl1Choice `xml:"PwrOfAttnyLvlOfCtrl,omitempty"`

	// Specifies if the account is regarded as domestic or non-domestic for reporting purposes.
	AccountingStatus *AccountingStatus1Choice `xml:"AcctgSts,omitempty"`

	// Legal opening date for the account.
	OpeningDate *DateAndDateTimeChoice `xml:"OpngDt,omitempty"`

	// Legal closing date for the account.
	ClosingDate *DateAndDateTimeChoice `xml:"ClsgDt,omitempty"`

	// Indicates whether the account can hold a negative position in a security.
	NegativeIndicator *YesNoIndicator `xml:"NegInd,omitempty"`

	// Order in which securities are moved from the account.
	ProcessingOrder *PositionEffect3Code `xml:"PrcgOrdr,omitempty"`

	// Specifies whether the investor assumes responsibility for the liability.
	Liability *Liability1Choice `xml:"Lblty,omitempty"`

	// Information related to the investor profile to be inserted, updated or deleted.
	ModifiedInvestorProfile []*ModificationScope32 `xml:"ModfdInvstrPrfl,omitempty"`

	// Fiscal year, when not the same as the calendar year.
	FiscalYear *FiscalYear1Choice `xml:"FsclYr,omitempty"`
}

func (i *InvestmentAccount63) AddAccountStatusUpdateInstruction() *AccountStatusUpdateInstruction1 {
	i.AccountStatusUpdateInstruction = new(AccountStatusUpdateInstruction1)
	return i.AccountStatusUpdateInstruction
}

func (i *InvestmentAccount63) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount63) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount63) AddType() *AccountType2Choice {
	i.Type = new(AccountType2Choice)
	return i.Type
}

func (i *InvestmentAccount63) AddOwnershipType() *OwnershipType2Choice {
	i.OwnershipType = new(OwnershipType2Choice)
	return i.OwnershipType
}

func (i *InvestmentAccount63) AddTaxExemption() *TaxExemptionReason2Choice {
	i.TaxExemption = new(TaxExemptionReason2Choice)
	return i.TaxExemption
}

func (i *InvestmentAccount63) AddStatementFrequency() *StatementFrequencyReason2Choice {
	i.StatementFrequency = new(StatementFrequencyReason2Choice)
	return i.StatementFrequency
}

func (i *InvestmentAccount63) SetReferenceCurrency(value string) {
	i.ReferenceCurrency = (*ActiveCurrencyCode)(&value)
}

func (i *InvestmentAccount63) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *InvestmentAccount63) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentAccount63) AddReinvestmentDetails() *Reinvestment3 {
	newValue := new(Reinvestment3)
	i.ReinvestmentDetails = append(i.ReinvestmentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount63) SetTaxWithholdingMethod(value string) {
	i.TaxWithholdingMethod = (*TaxWithholdingMethod3Code)(&value)
}

func (i *InvestmentAccount63) AddTaxReporting() *TaxReporting2 {
	newValue := new(TaxReporting2)
	i.TaxReporting = append(i.TaxReporting, newValue)
	return newValue
}

func (i *InvestmentAccount63) AddLetterIntentDetails() *LetterIntent1 {
	i.LetterIntentDetails = new(LetterIntent1)
	return i.LetterIntentDetails
}

func (i *InvestmentAccount63) SetAccumulationRightReference(value string) {
	i.AccumulationRightReference = (*Max35Text)(&value)
}

func (i *InvestmentAccount63) SetRequiredSignatoriesNumber(value string) {
	i.RequiredSignatoriesNumber = (*Number)(&value)
}

func (i *InvestmentAccount63) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount63) AddModifiedFinancialInstrumentDetails() *ModificationScope36 {
	newValue := new(ModificationScope36)
	i.ModifiedFinancialInstrumentDetails = append(i.ModifiedFinancialInstrumentDetails, newValue)
	return newValue
}

func (i *InvestmentAccount63) AddRoundingDetails() *RoundingParameters1 {
	i.RoundingDetails = new(RoundingParameters1)
	return i.RoundingDetails
}

func (i *InvestmentAccount63) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount63) AddBlockedStatus() *BlockedStatusReason2Choice {
	i.BlockedStatus = new(BlockedStatusReason2Choice)
	return i.BlockedStatus
}

func (i *InvestmentAccount63) AddAccountUsageType() *AccountUsageType2Choice {
	i.AccountUsageType = new(AccountUsageType2Choice)
	return i.AccountUsageType
}

func (i *InvestmentAccount63) SetForeignStatusCertification(value string) {
	i.ForeignStatusCertification = (*Provided1Code)(&value)
}

func (i *InvestmentAccount63) AddAccountSignatureDateTime() *DateAndDateTimeChoice {
	i.AccountSignatureDateTime = new(DateAndDateTimeChoice)
	return i.AccountSignatureDateTime
}

func (i *InvestmentAccount63) AddTransactionChannelType() *TransactionChannelType1Choice {
	i.TransactionChannelType = new(TransactionChannelType1Choice)
	return i.TransactionChannelType
}

func (i *InvestmentAccount63) AddInvestmentAccountCategory() *InvestmentAccountCategory1Choice {
	i.InvestmentAccountCategory = new(InvestmentAccountCategory1Choice)
	return i.InvestmentAccountCategory
}

func (i *InvestmentAccount63) SetPledging(value string) {
	i.Pledging = (*Eligible1Code)(&value)
}

func (i *InvestmentAccount63) SetCollateral(value string) {
	i.Collateral = (*Collateral1Code)(&value)
}

func (i *InvestmentAccount63) AddThirdPartyRights() *ThirdPartyRights1 {
	i.ThirdPartyRights = new(ThirdPartyRights1)
	return i.ThirdPartyRights
}

func (i *InvestmentAccount63) AddPowerOfAttorneyLevelOfControl() *LevelOfControl1Choice {
	i.PowerOfAttorneyLevelOfControl = new(LevelOfControl1Choice)
	return i.PowerOfAttorneyLevelOfControl
}

func (i *InvestmentAccount63) AddAccountingStatus() *AccountingStatus1Choice {
	i.AccountingStatus = new(AccountingStatus1Choice)
	return i.AccountingStatus
}

func (i *InvestmentAccount63) AddOpeningDate() *DateAndDateTimeChoice {
	i.OpeningDate = new(DateAndDateTimeChoice)
	return i.OpeningDate
}

func (i *InvestmentAccount63) AddClosingDate() *DateAndDateTimeChoice {
	i.ClosingDate = new(DateAndDateTimeChoice)
	return i.ClosingDate
}

func (i *InvestmentAccount63) SetNegativeIndicator(value string) {
	i.NegativeIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount63) SetProcessingOrder(value string) {
	i.ProcessingOrder = (*PositionEffect3Code)(&value)
}

func (i *InvestmentAccount63) AddLiability() *Liability1Choice {
	i.Liability = new(Liability1Choice)
	return i.Liability
}

func (i *InvestmentAccount63) AddModifiedInvestorProfile() *ModificationScope32 {
	newValue := new(ModificationScope32)
	i.ModifiedInvestorProfile = append(i.ModifiedInvestorProfile, newValue)
	return newValue
}

func (i *InvestmentAccount63) AddFiscalYear() *FiscalYear1Choice {
	i.FiscalYear = new(FiscalYear1Choice)
	return i.FiscalYear
}

// Information about a securities account and its characteristics.
type InvestmentAccount64 struct {

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Legal form of the fund, for example, UCITS, SICAV, OEIC, Unit Trust, and FCP.
	FundType *Max35Text `xml:"FndTp,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	SecurityDetails *FinancialInstrument55 `xml:"SctyDtls,omitempty"`

	// Owner of the account.
	AccountOwner *AccountOwner2Choice `xml:"AcctOwnr,omitempty"`

	// Intermediary or other party related to the management of the account.
	Intermediary []*Intermediary33 `xml:"Intrmy,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount64) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount64) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount64) SetFundType(value string) {
	i.FundType = (*Max35Text)(&value)
}

func (i *InvestmentAccount64) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount64) AddSecurityDetails() *FinancialInstrument55 {
	i.SecurityDetails = new(FinancialInstrument55)
	return i.SecurityDetails
}

func (i *InvestmentAccount64) AddAccountOwner() *AccountOwner2Choice {
	i.AccountOwner = new(AccountOwner2Choice)
	return i.AccountOwner
}

func (i *InvestmentAccount64) AddIntermediary() *Intermediary33 {
	newValue := new(Intermediary33)
	i.Intermediary = append(i.Intermediary, newValue)
	return newValue
}

func (i *InvestmentAccount64) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}
