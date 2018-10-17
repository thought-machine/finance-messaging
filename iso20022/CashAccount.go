package iso20022

// Information used for identifying an account.
type CashAccount11 struct {

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Identification *CashAccountIdentification1Choice `xml:"Id"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *FinancialInstitutionIdentification3Choice `xml:"AcctSvcr,omitempty"`
}

func (c *CashAccount11) AddIdentification() *CashAccountIdentification1Choice {
	c.Identification = new(CashAccountIdentification1Choice)
	return c.Identification
}

func (c *CashAccount11) AddAccountServicer() *FinancialInstitutionIdentification3Choice {
	c.AccountServicer = new(FinancialInstitutionIdentification3Choice)
	return c.AccountServicer
}

// Account to or from which a cash entry is made.
type CashAccount12 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *CashAccountIdentification1Choice `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm"`

	// Specifies the nature, or use, of the cash account.
	Type *CashAccountType1Code `xml:"Tp,omitempty"`

	// Specifies the nature, or use, of the cash account.
	ExtendedType *Extended350Code `xml:"XtndedTp,omitempty"`

	// Medium of exchange of value.
	Currency *ActiveCurrencyCode `xml:"Ccy"`

	// Specifies the current state of an account, eg, enabled or deleted.
	Status *AccountStatus1Code `xml:"Sts"`
}

func (c *CashAccount12) AddIdentification() *CashAccountIdentification1Choice {
	c.Identification = new(CashAccountIdentification1Choice)
	return c.Identification
}

func (c *CashAccount12) SetName(value string) {
	c.Name = (*Max35Text)(&value)
}

func (c *CashAccount12) SetType(value string) {
	c.Type = (*CashAccountType1Code)(&value)
}

func (c *CashAccount12) SetExtendedType(value string) {
	c.ExtendedType = (*Extended350Code)(&value)
}

func (c *CashAccount12) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CashAccount12) SetStatus(value string) {
	c.Status = (*AccountStatus1Code)(&value)
}

// Information used for identifying an account.
type CashAccount13 struct {

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Identification *AccountIdentification3Choice `xml:"Id"`

	// Nature, or use, of the account.
	Type *CashAccountType2 `xml:"Tp,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, assigned by the account servicing institution in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage : the account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Party that legally owns the account.
	Owner *PartyIdentification8 `xml:"Ownr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *BranchAndFinancialInstitutionIdentification3 `xml:"Svcr,omitempty"`
}

func (c *CashAccount13) AddIdentification() *AccountIdentification3Choice {
	c.Identification = new(AccountIdentification3Choice)
	return c.Identification
}

func (c *CashAccount13) AddType() *CashAccountType2 {
	c.Type = new(CashAccountType2)
	return c.Type
}

func (c *CashAccount13) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CashAccount13) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

func (c *CashAccount13) AddOwner() *PartyIdentification8 {
	c.Owner = new(PartyIdentification8)
	return c.Owner
}

func (c *CashAccount13) AddServicer() *BranchAndFinancialInstitutionIdentification3 {
	c.Servicer = new(BranchAndFinancialInstitutionIdentification3)
	return c.Servicer
}

// Set of elements used to identify an account.
type CashAccount16 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4Choice `xml:"Id"`

	// Specifies the nature, or use of the account.
	Type *CashAccountType2 `xml:"Tp,omitempty"`

	// Identification of the currency in which the account is held.
	//
	// Usage: Currency should only be used in case one and the same account number covers several currencies
	// and the initiating party needs to identify which currency needs to be used for settlement on the account.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (c *CashAccount16) AddIdentification() *AccountIdentification4Choice {
	c.Identification = new(AccountIdentification4Choice)
	return c.Identification
}

func (c *CashAccount16) AddType() *CashAccountType2 {
	c.Type = new(CashAccountType2)
	return c.Type
}

func (c *CashAccount16) SetCurrency(value string) {
	c.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccount16) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

// Provides information about the cash account.
type CashAccount17 struct {

	// Identification of the cash account.
	AccountIdentification *CashAccountIdentification1Choice `xml:"AcctId"`

	// Currency of the payment.
	PaymentCurrency *ActiveCurrencyCode `xml:"PmtCcy"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Identification of the cash correspondent back.
	CorrespondentBankIdentification *BICIdentifier `xml:"CrspdtBkId"`
}

func (c *CashAccount17) AddAccountIdentification() *CashAccountIdentification1Choice {
	c.AccountIdentification = new(CashAccountIdentification1Choice)
	return c.AccountIdentification
}

func (c *CashAccount17) SetPaymentCurrency(value string) {
	c.PaymentCurrency = (*ActiveCurrencyCode)(&value)
}

func (c *CashAccount17) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	c.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return c.AccountOwnerIdentification
}

func (c *CashAccount17) SetCorrespondentBankIdentification(value string) {
	c.CorrespondentBankIdentification = (*BICIdentifier)(&value)
}

// Provides information about the cash account.
type CashAccount18 struct {

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Identification of the cash account or the securities account from which the cash account is derived.
	AccountIdentification *AccountIdentification2Choice `xml:"AcctId"`

	// The cash balance type.
	BalanceType *CashBalanceType1FormatType `xml:"BalTp,omitempty"`
}

func (c *CashAccount18) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashAccount18) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	c.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return c.AccountOwnerIdentification
}

func (c *CashAccount18) AddAccountIdentification() *AccountIdentification2Choice {
	c.AccountIdentification = new(AccountIdentification2Choice)
	return c.AccountIdentification
}

func (c *CashAccount18) AddBalanceType() *CashBalanceType1FormatType {
	c.BalanceType = new(CashBalanceType1FormatType)
	return c.BalanceType
}

// Provides information about the cash account.
type CashAccount19 struct {

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Identification of the cash account or the securities account from which the cash account is derived.
	AccountIdentification *AccountIdentification2Choice `xml:"AcctId"`
}

func (c *CashAccount19) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashAccount19) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	c.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return c.AccountOwnerIdentification
}

func (c *CashAccount19) AddAccountIdentification() *AccountIdentification2Choice {
	c.AccountIdentification = new(AccountIdentification2Choice)
	return c.AccountIdentification
}

// Set of elements used to identify an account.
type CashAccount20 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4Choice `xml:"Id"`

	// Specifies the nature, or use of the account.
	Type *CashAccountType2 `xml:"Tp,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Party that legally owns the account.
	Owner *PartyIdentification32 `xml:"Ownr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *BranchAndFinancialInstitutionIdentification4 `xml:"Svcr,omitempty"`
}

func (c *CashAccount20) AddIdentification() *AccountIdentification4Choice {
	c.Identification = new(AccountIdentification4Choice)
	return c.Identification
}

func (c *CashAccount20) AddType() *CashAccountType2 {
	c.Type = new(CashAccountType2)
	return c.Type
}

func (c *CashAccount20) SetCurrency(value string) {
	c.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccount20) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

func (c *CashAccount20) AddOwner() *PartyIdentification32 {
	c.Owner = new(PartyIdentification32)
	return c.Owner
}

func (c *CashAccount20) AddServicer() *BranchAndFinancialInstitutionIdentification4 {
	c.Servicer = new(BranchAndFinancialInstitutionIdentification4)
	return c.Servicer
}

// Account to or from which a cash entry is made.
type CashAccount21 struct {

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *BICIdentifier `xml:"Svcr,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification5Choice `xml:"Id"`
}

func (c *CashAccount21) SetServicer(value string) {
	c.Servicer = (*BICIdentifier)(&value)
}

func (c *CashAccount21) AddIdentification() *AccountIdentification5Choice {
	c.Identification = new(AccountIdentification5Choice)
	return c.Identification
}

// Account to or from which a cash entry is made.
type CashAccount22 struct {

	// Medium of exchange of value.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *BICIdentifier `xml:"Svcr"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification5Choice `xml:"Id"`

	// Sub-division of a master or omnibus cash account.
	SecondaryAccount *CashAccount21 `xml:"ScndryAcct,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountTypeDescription *Max35Text `xml:"AcctTpDesc"`
}

func (c *CashAccount22) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CashAccount22) SetServicer(value string) {
	c.Servicer = (*BICIdentifier)(&value)
}

func (c *CashAccount22) AddIdentification() *AccountIdentification5Choice {
	c.Identification = new(AccountIdentification5Choice)
	return c.Identification
}

func (c *CashAccount22) AddSecondaryAccount() *CashAccount21 {
	c.SecondaryAccount = new(CashAccount21)
	return c.SecondaryAccount
}

func (c *CashAccount22) SetAccountTypeDescription(value string) {
	c.AccountTypeDescription = (*Max35Text)(&value)
}

// Provides the details to identify an account.
type CashAccount24 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4Choice `xml:"Id"`

	// Specifies the nature, or use of the account.
	Type *CashAccountType2Choice `xml:"Tp,omitempty"`

	// Identification of the currency in which the account is held.
	//
	// Usage: Currency should only be used in case one and the same account number covers several currencies
	// and the initiating party needs to identify which currency needs to be used for settlement on the account.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (c *CashAccount24) AddIdentification() *AccountIdentification4Choice {
	c.Identification = new(AccountIdentification4Choice)
	return c.Identification
}

func (c *CashAccount24) AddType() *CashAccountType2Choice {
	c.Type = new(CashAccountType2Choice)
	return c.Type
}

func (c *CashAccount24) SetCurrency(value string) {
	c.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccount24) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

// Provides the details to identify an account.
type CashAccount25 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4Choice `xml:"Id"`

	// Specifies the nature, or use of the account.
	Type *CashAccountType2Choice `xml:"Tp,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Party that legally owns the account.
	Owner *PartyIdentification43 `xml:"Ownr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *BranchAndFinancialInstitutionIdentification5 `xml:"Svcr,omitempty"`
}

func (c *CashAccount25) AddIdentification() *AccountIdentification4Choice {
	c.Identification = new(AccountIdentification4Choice)
	return c.Identification
}

func (c *CashAccount25) AddType() *CashAccountType2Choice {
	c.Type = new(CashAccountType2Choice)
	return c.Type
}

func (c *CashAccount25) SetCurrency(value string) {
	c.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccount25) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

func (c *CashAccount25) AddOwner() *PartyIdentification43 {
	c.Owner = new(PartyIdentification43)
	return c.Owner
}

func (c *CashAccount25) AddServicer() *BranchAndFinancialInstitutionIdentification5 {
	c.Servicer = new(BranchAndFinancialInstitutionIdentification5)
	return c.Servicer
}

// Account to or from which a cash entry is made.
type CashAccount26 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationAndName3 `xml:"Id"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification2Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	AccountServicerBranch *BranchData `xml:"AcctSvcrBrnch,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, ISA.
	InvestmentAccountType *InvestmentAccountType1Choice `xml:"InvstmtAcctTp,omitempty"`

	// Other identification such as national registration identification number, passport number.
	AccountOwnerOtherIdentification []*GenericIdentification46 `xml:"AcctOwnrOthrId,omitempty"`
}

func (c *CashAccount26) AddIdentification() *AccountIdentificationAndName3 {
	c.Identification = new(AccountIdentificationAndName3)
	return c.Identification
}

func (c *CashAccount26) AddAccountOwner() *PartyIdentification2Choice {
	c.AccountOwner = new(PartyIdentification2Choice)
	return c.AccountOwner
}

func (c *CashAccount26) AddAccountServicer() *PartyIdentification2Choice {
	c.AccountServicer = new(PartyIdentification2Choice)
	return c.AccountServicer
}

func (c *CashAccount26) AddAccountServicerBranch() *BranchData {
	c.AccountServicerBranch = new(BranchData)
	return c.AccountServicerBranch
}

func (c *CashAccount26) AddInvestmentAccountType() *InvestmentAccountType1Choice {
	c.InvestmentAccountType = new(InvestmentAccountType1Choice)
	return c.InvestmentAccountType
}

func (c *CashAccount26) AddAccountOwnerOtherIdentification() *GenericIdentification46 {
	newValue := new(GenericIdentification46)
	c.AccountOwnerOtherIdentification = append(c.AccountOwnerOtherIdentification, newValue)
	return newValue
}

// Set of elements used to identify an account.
type CashAccount27 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4Choice `xml:"Id"`

	// Specifies the nature, or use of the account.
	Type *CashAccountType2 `xml:"Tp,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Party that legally owns the account.
	Owner *PartyIdentification41 `xml:"Ownr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	//
	Servicer *BranchAndFinancialInstitutionIdentification5 `xml:"Svcr,omitempty"`
}

func (c *CashAccount27) AddIdentification() *AccountIdentification4Choice {
	c.Identification = new(AccountIdentification4Choice)
	return c.Identification
}

func (c *CashAccount27) AddType() *CashAccountType2 {
	c.Type = new(CashAccountType2)
	return c.Type
}

func (c *CashAccount27) SetCurrency(value string) {
	c.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccount27) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

func (c *CashAccount27) AddOwner() *PartyIdentification41 {
	c.Owner = new(PartyIdentification41)
	return c.Owner
}

func (c *CashAccount27) AddServicer() *BranchAndFinancialInstitutionIdentification5 {
	c.Servicer = new(BranchAndFinancialInstitutionIdentification5)
	return c.Servicer
}

// Set of elements used to identify an account.
type CashAccount28 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4Choice `xml:"Id"`

	// Identification of the currency in which the account is held.
	//
	// Usage: Currency should only be used in case one and the same account number covers several currencies
	// and the initiating party needs to identify which currency needs to be used for settlement on the account.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (c *CashAccount28) AddIdentification() *AccountIdentification4Choice {
	c.Identification = new(AccountIdentification4Choice)
	return c.Identification
}

func (c *CashAccount28) SetCurrency(value string) {
	c.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccount28) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

// Information used for identifying an account.
type CashAccount29 struct {

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Identification *CashAccountIdentification1Choice `xml:"Id"`

	// Institution that maintains the records where the account is held.
	AccountServicer *FinancialInstitutionIdentification3Choice `xml:"AcctSvcr,omitempty"`
}

func (c *CashAccount29) AddIdentification() *CashAccountIdentification1Choice {
	c.Identification = new(CashAccountIdentification1Choice)
	return c.Identification
}

func (c *CashAccount29) AddAccountServicer() *FinancialInstitutionIdentification3Choice {
	c.AccountServicer = new(FinancialInstitutionIdentification3Choice)
	return c.AccountServicer
}

// Information used for identifying an account.
type CashAccount3 struct {

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Identification *AccountIdentification1Choice `xml:"Id"`

	// Nature, or use, of the account.
	Type *CashAccountType3Code `xml:"Tp,omitempty"`

	// Identification of the currency in which the account is held.
	//
	// Usage: Currency should only be used in case one and the same account number covers several currencies and the initiating party needs to identify which currency needs to be used for settlement on the account.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, assigned by the account servicing institution in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage : the account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (c *CashAccount3) AddIdentification() *AccountIdentification1Choice {
	c.Identification = new(AccountIdentification1Choice)
	return c.Identification
}

func (c *CashAccount3) SetType(value string) {
	c.Type = (*CashAccountType3Code)(&value)
}

func (c *CashAccount3) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CashAccount3) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}

// Account to or from which a cash entry is made.
type CashAccount32 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationAndName5 `xml:"Id"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification70Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	AccountServicerBranch *BranchData `xml:"AcctSvcrBrnch,omitempty"`
}

func (c *CashAccount32) AddIdentification() *AccountIdentificationAndName5 {
	c.Identification = new(AccountIdentificationAndName5)
	return c.Identification
}

func (c *CashAccount32) AddAccountOwner() *PartyIdentification70Choice {
	c.AccountOwner = new(PartyIdentification70Choice)
	return c.AccountOwner
}

func (c *CashAccount32) AddAccountServicer() *PartyIdentification70Choice {
	c.AccountServicer = new(PartyIdentification70Choice)
	return c.AccountServicer
}

func (c *CashAccount32) AddAccountServicerBranch() *BranchData {
	c.AccountServicerBranch = new(BranchData)
	return c.AccountServicerBranch
}

// Account to or from which a cash entry is made.
type CashAccount33 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationAndName5 `xml:"Id"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification70Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *FinancialInstitutionIdentification7Choice `xml:"AcctSvcr,omitempty"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	AccountServicerBranch *BranchData `xml:"AcctSvcrBrnch,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, tax identification number. This may be an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	AccountOwnerOtherIdentification []*GenericIdentification82 `xml:"AcctOwnrOthrId,omitempty"`

	// Type of account.
	InvestmentAccountType *AccountType2Choice `xml:"InvstmtAcctTp,omitempty"`

	// Specifies if the account is for credits or debits.
	CreditDebit *CreditDebit3Code `xml:"CdtDbt,omitempty"`

	// Type of transaction for which the cash account is specified.
	SettlementInstructionReason *SettlementInstructionReason1Choice `xml:"SttlmInstrRsn,omitempty"`

	// Purpose of the cash account.
	CashAccountPurpose *CashAccountType3Choice `xml:"CshAcctPurp,omitempty"`

	// Specifies whether the account is the primary or secondary account if there are two accounts for the same purpose.
	CashAccountDesignation *AccountDesignation1Choice `xml:"CshAcctDsgnt,omitempty"`

	// Percentage of the dividend payment not to be reinvested, that is, to be paid in cash.
	DividendPercentage *PercentageBoundedRate `xml:"DvddPctg,omitempty"`
}

func (c *CashAccount33) SetSettlementCurrency(value string) {
	c.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (c *CashAccount33) AddIdentification() *AccountIdentificationAndName5 {
	c.Identification = new(AccountIdentificationAndName5)
	return c.Identification
}

func (c *CashAccount33) AddAccountOwner() *PartyIdentification70Choice {
	c.AccountOwner = new(PartyIdentification70Choice)
	return c.AccountOwner
}

func (c *CashAccount33) AddAccountServicer() *FinancialInstitutionIdentification7Choice {
	c.AccountServicer = new(FinancialInstitutionIdentification7Choice)
	return c.AccountServicer
}

func (c *CashAccount33) AddAccountServicerBranch() *BranchData {
	c.AccountServicerBranch = new(BranchData)
	return c.AccountServicerBranch
}

func (c *CashAccount33) AddAccountOwnerOtherIdentification() *GenericIdentification82 {
	newValue := new(GenericIdentification82)
	c.AccountOwnerOtherIdentification = append(c.AccountOwnerOtherIdentification, newValue)
	return newValue
}

func (c *CashAccount33) AddInvestmentAccountType() *AccountType2Choice {
	c.InvestmentAccountType = new(AccountType2Choice)
	return c.InvestmentAccountType
}

func (c *CashAccount33) SetCreditDebit(value string) {
	c.CreditDebit = (*CreditDebit3Code)(&value)
}

func (c *CashAccount33) AddSettlementInstructionReason() *SettlementInstructionReason1Choice {
	c.SettlementInstructionReason = new(SettlementInstructionReason1Choice)
	return c.SettlementInstructionReason
}

func (c *CashAccount33) AddCashAccountPurpose() *CashAccountType3Choice {
	c.CashAccountPurpose = new(CashAccountType3Choice)
	return c.CashAccountPurpose
}

func (c *CashAccount33) AddCashAccountDesignation() *AccountDesignation1Choice {
	c.CashAccountDesignation = new(AccountDesignation1Choice)
	return c.CashAccountDesignation
}

func (c *CashAccount33) SetDividendPercentage(value string) {
	c.DividendPercentage = (*PercentageBoundedRate)(&value)
}

// Information used for identifying an account.
type CashAccount34 struct {

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Identification *AccountIdentificationAndName5 `xml:"Id"`

	// Institution that maintains the records where the account is held.
	AccountServicer *FinancialInstitutionIdentification7Choice `xml:"AcctSvcr,omitempty"`
}

func (c *CashAccount34) AddIdentification() *AccountIdentificationAndName5 {
	c.Identification = new(AccountIdentificationAndName5)
	return c.Identification
}

func (c *CashAccount34) AddAccountServicer() *FinancialInstitutionIdentification7Choice {
	c.AccountServicer = new(FinancialInstitutionIdentification7Choice)
	return c.AccountServicer
}

// Account to or from which a cash entry is made.
type CashAccount4 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationAndName3 `xml:"Id"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification2Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	AccountServicerBranch *BranchData `xml:"AcctSvcrBrnch,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	InvestmentAccountType *FundCashAccount2Code `xml:"InvstmtAcctTp,omitempty"`

	// Purpose of the account/source fund type. This is typically linked to an investment product, eg, wrapper, PEP, ISA.
	ExtendedInvestmentAccountType *Extended350Code `xml:"XtndedInvstmtAcctTp,omitempty"`
}

func (c *CashAccount4) AddIdentification() *AccountIdentificationAndName3 {
	c.Identification = new(AccountIdentificationAndName3)
	return c.Identification
}

func (c *CashAccount4) AddAccountOwner() *PartyIdentification2Choice {
	c.AccountOwner = new(PartyIdentification2Choice)
	return c.AccountOwner
}

func (c *CashAccount4) AddAccountServicer() *PartyIdentification2Choice {
	c.AccountServicer = new(PartyIdentification2Choice)
	return c.AccountServicer
}

func (c *CashAccount4) AddAccountServicerBranch() *BranchData {
	c.AccountServicerBranch = new(BranchData)
	return c.AccountServicerBranch
}

func (c *CashAccount4) SetInvestmentAccountType(value string) {
	c.InvestmentAccountType = (*FundCashAccount2Code)(&value)
}

func (c *CashAccount4) SetExtendedInvestmentAccountType(value string) {
	c.ExtendedInvestmentAccountType = (*Extended350Code)(&value)
}

// Information used for identifying an account.
type CashAccount7 struct {

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Identification *AccountIdentification3Choice `xml:"Id"`

	// Nature, or use, of the account.
	Type *CashAccountType2 `xml:"Tp,omitempty"`

	// Identification of the currency in which the account is held.
	//
	// Usage: Currency should only be used in case one and the same account number covers several currencies and the initiating party needs to identify which currency needs to be used for settlement on the account.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Name of the account, assigned by the account servicing institution in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage : the account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (c *CashAccount7) AddIdentification() *AccountIdentification3Choice {
	c.Identification = new(AccountIdentification3Choice)
	return c.Identification
}

func (c *CashAccount7) AddType() *CashAccountType2 {
	c.Type = new(CashAccountType2)
	return c.Type
}

func (c *CashAccount7) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CashAccount7) SetName(value string) {
	c.Name = (*Max70Text)(&value)
}
