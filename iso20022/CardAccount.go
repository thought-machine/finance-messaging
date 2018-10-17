package iso20022

// Way to identify a customer account or a relationship to its account affected for the transaction.
type CardAccount1 struct {

	// Method used by the cardholder and the terminal for the choice of the account.
	SelectionMethod *AccountChoiceMethod1Code `xml:"SelctnMtd,omitempty"`

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType2Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Party that legally owns the account.
	AccountOwner *NameAndAddress3 `xml:"AcctOwnr,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification30Choice `xml:"AcctIdr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`
}

func (c *CardAccount1) SetSelectionMethod(value string) {
	c.SelectionMethod = (*AccountChoiceMethod1Code)(&value)
}

func (c *CardAccount1) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType2Code)(&value)
}

func (c *CardAccount1) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount1) AddAccountOwner() *NameAndAddress3 {
	c.AccountOwner = new(NameAndAddress3)
	return c.AccountOwner
}

func (c *CardAccount1) AddAccountIdentifier() *AccountIdentification30Choice {
	c.AccountIdentifier = new(AccountIdentification30Choice)
	return c.AccountIdentifier
}

func (c *CardAccount1) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

// Customer account information.
type CardAccount10 struct {

	// Sequence number of the account data for multi-account deposit.
	AccountSequenceNumber *Number `xml:"AcctSeqNb,omitempty"`

	// Method used by the cardholder and the terminal for the choice of the account.
	SelectionMethod *AccountChoiceMethod1Code `xml:"SelctnMtd,omitempty"`

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType3Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Party that legally owns the account.
	AccountOwner *NameAndAddress3 `xml:"AcctOwnr,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr,omitempty"`

	// Internal account reference in case of credit account.
	CreditReference *Max35Text `xml:"CdtRef,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`

	// Balance of the account.
	Balance *AmountAndDirection43 `xml:"Bal,omitempty"`

	// Indicates if the balance must be displayed to the customer on the ATM.
	BalanceDisplayFlag *TrueFalseIndicator `xml:"BalDispFlg,omitempty"`

	// Indicates if this is the default account.
	DefaultAccountIndicator *TrueFalseIndicator `xml:"DfltAcctInd,omitempty"`

	// Limit of amounts for the customer.
	Limits *ATMTransactionAmounts8 `xml:"Lmts,omitempty"`
}

func (c *CardAccount10) SetAccountSequenceNumber(value string) {
	c.AccountSequenceNumber = (*Number)(&value)
}

func (c *CardAccount10) SetSelectionMethod(value string) {
	c.SelectionMethod = (*AccountChoiceMethod1Code)(&value)
}

func (c *CardAccount10) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType3Code)(&value)
}

func (c *CardAccount10) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount10) AddAccountOwner() *NameAndAddress3 {
	c.AccountOwner = new(NameAndAddress3)
	return c.AccountOwner
}

func (c *CardAccount10) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount10) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount10) SetCreditReference(value string) {
	c.CreditReference = (*Max35Text)(&value)
}

func (c *CardAccount10) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

func (c *CardAccount10) AddBalance() *AmountAndDirection43 {
	c.Balance = new(AmountAndDirection43)
	return c.Balance
}

func (c *CardAccount10) SetBalanceDisplayFlag(value string) {
	c.BalanceDisplayFlag = (*TrueFalseIndicator)(&value)
}

func (c *CardAccount10) SetDefaultAccountIndicator(value string) {
	c.DefaultAccountIndicator = (*TrueFalseIndicator)(&value)
}

func (c *CardAccount10) AddLimits() *ATMTransactionAmounts8 {
	c.Limits = new(ATMTransactionAmounts8)
	return c.Limits
}

// Customer account information.
type CardAccount11 struct {

	// Method used by the cardholder and the terminal for the choice of the account.
	SelectionMethod *AccountChoiceMethod1Code `xml:"SelctnMtd,omitempty"`

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType3Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr,omitempty"`

	// Internal account reference in case of credit account.
	CreditReference *Max35Text `xml:"CdtRef,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`
}

func (c *CardAccount11) SetSelectionMethod(value string) {
	c.SelectionMethod = (*AccountChoiceMethod1Code)(&value)
}

func (c *CardAccount11) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType3Code)(&value)
}

func (c *CardAccount11) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount11) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount11) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount11) SetCreditReference(value string) {
	c.CreditReference = (*Max35Text)(&value)
}

func (c *CardAccount11) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

// Customer account information.
type CardAccount12 struct {

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType3Code `xml:"AcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr,omitempty"`

	// Internal account reference in case of credit account.
	CreditReference *Max35Text `xml:"CdtRef,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`

	// Balance of the account.
	Balance *AmountAndDirection43 `xml:"Bal,omitempty"`

	// Indicates if the balance must be displayed to the customer on the ATM.
	BalanceDisplayFlag *TrueFalseIndicator `xml:"BalDispFlg,omitempty"`

	// Indicates if this is the default account.
	DefaultAccountIndicator *TrueFalseIndicator `xml:"DfltAcctInd,omitempty"`

	// Service allowed on the account.
	AllowedService []*ATMService19 `xml:"AllwdSvc,omitempty"`
}

func (c *CardAccount12) SetAccountType(value string) {
	c.AccountType = (*CardAccountType3Code)(&value)
}

func (c *CardAccount12) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount12) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount12) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount12) SetCreditReference(value string) {
	c.CreditReference = (*Max35Text)(&value)
}

func (c *CardAccount12) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

func (c *CardAccount12) AddBalance() *AmountAndDirection43 {
	c.Balance = new(AmountAndDirection43)
	return c.Balance
}

func (c *CardAccount12) SetBalanceDisplayFlag(value string) {
	c.BalanceDisplayFlag = (*TrueFalseIndicator)(&value)
}

func (c *CardAccount12) SetDefaultAccountIndicator(value string) {
	c.DefaultAccountIndicator = (*TrueFalseIndicator)(&value)
}

func (c *CardAccount12) AddAllowedService() *ATMService19 {
	newValue := new(ATMService19)
	c.AllowedService = append(c.AllowedService, newValue)
	return newValue
}

// Customer account information.
type CardAccount13 struct {

	// Type of account used for the transaction.
	AccountType *CardAccountType3Code `xml:"AcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr,omitempty"`

	// Internal account reference in case of credit account.
	CreditReference *Max35Text `xml:"CdtRef,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`

	// Balance of the account before the transfer.
	BalanceBefore *AmountAndDirection43 `xml:"BalBfr,omitempty"`

	// Balance of the account after the transfer.
	BalanceAfter *AmountAndDirection43 `xml:"BalAftr,omitempty"`
}

func (c *CardAccount13) SetAccountType(value string) {
	c.AccountType = (*CardAccountType3Code)(&value)
}

func (c *CardAccount13) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount13) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount13) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount13) SetCreditReference(value string) {
	c.CreditReference = (*Max35Text)(&value)
}

func (c *CardAccount13) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

func (c *CardAccount13) AddBalanceBefore() *AmountAndDirection43 {
	c.BalanceBefore = new(AmountAndDirection43)
	return c.BalanceBefore
}

func (c *CardAccount13) AddBalanceAfter() *AmountAndDirection43 {
	c.BalanceAfter = new(AmountAndDirection43)
	return c.BalanceAfter
}

// Account involved in the card transaction.
type CardAccount2 struct {

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType2Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Party that legally owns the account.
	AccountOwner *NameAndAddress3 `xml:"AcctOwnr,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification30Choice `xml:"AcctIdr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`

	//  Balance of the account.
	Balance *AmountAndDirection41 `xml:"Bal,omitempty"`

	//  Date of the balance.
	BalanceDate *ISODate `xml:"BalDt,omitempty"`
}

func (c *CardAccount2) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType2Code)(&value)
}

func (c *CardAccount2) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount2) AddAccountOwner() *NameAndAddress3 {
	c.AccountOwner = new(NameAndAddress3)
	return c.AccountOwner
}

func (c *CardAccount2) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount2) AddAccountIdentifier() *AccountIdentification30Choice {
	c.AccountIdentifier = new(AccountIdentification30Choice)
	return c.AccountIdentifier
}

func (c *CardAccount2) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

func (c *CardAccount2) AddBalance() *AmountAndDirection41 {
	c.Balance = new(AmountAndDirection41)
	return c.Balance
}

func (c *CardAccount2) SetBalanceDate(value string) {
	c.BalanceDate = (*ISODate)(&value)
}

// Customer account information.
type CardAccount3 struct {

	// Method used by the cardholder and the terminal for the choice of the account.
	SelectionMethod *AccountChoiceMethod1Code `xml:"SelctnMtd,omitempty"`

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType2Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Party that legally owns the account.
	AccountOwner *NameAndAddress3 `xml:"AcctOwnr,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`
}

func (c *CardAccount3) SetSelectionMethod(value string) {
	c.SelectionMethod = (*AccountChoiceMethod1Code)(&value)
}

func (c *CardAccount3) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType2Code)(&value)
}

func (c *CardAccount3) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount3) AddAccountOwner() *NameAndAddress3 {
	c.AccountOwner = new(NameAndAddress3)
	return c.AccountOwner
}

func (c *CardAccount3) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount3) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount3) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

// Customer account information.
type CardAccount4 struct {

	// Method used by the cardholder and the terminal for the choice of the account.
	SelectionMethod *AccountChoiceMethod1Code `xml:"SelctnMtd,omitempty"`

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType2Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Party that legally owns the account.
	AccountOwner *NameAndAddress3 `xml:"AcctOwnr,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr,omitempty"`

	// Internal account reference in case of credit account.
	CreditReference *Max35Text `xml:"CdtRef,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`

	// Balance of the account.
	Balance *AmountAndDirection43 `xml:"Bal,omitempty"`

	// Indicates if the balance must be displayed to the customer on the ATM.
	BalanceDisplayFlag *TrueFalseIndicator `xml:"BalDispFlg,omitempty"`

	// Indicates if this is the default account.
	DefaultAccountIndicator *TrueFalseIndicator `xml:"DfltAcctInd,omitempty"`
}

func (c *CardAccount4) SetSelectionMethod(value string) {
	c.SelectionMethod = (*AccountChoiceMethod1Code)(&value)
}

func (c *CardAccount4) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType2Code)(&value)
}

func (c *CardAccount4) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount4) AddAccountOwner() *NameAndAddress3 {
	c.AccountOwner = new(NameAndAddress3)
	return c.AccountOwner
}

func (c *CardAccount4) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount4) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount4) SetCreditReference(value string) {
	c.CreditReference = (*Max35Text)(&value)
}

func (c *CardAccount4) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

func (c *CardAccount4) AddBalance() *AmountAndDirection43 {
	c.Balance = new(AmountAndDirection43)
	return c.Balance
}

func (c *CardAccount4) SetBalanceDisplayFlag(value string) {
	c.BalanceDisplayFlag = (*TrueFalseIndicator)(&value)
}

func (c *CardAccount4) SetDefaultAccountIndicator(value string) {
	c.DefaultAccountIndicator = (*TrueFalseIndicator)(&value)
}

// Customer account information.
type CardAccount5 struct {

	// Method used by the cardholder and the terminal for the choice of the account.
	SelectionMethod *AccountChoiceMethod1Code `xml:"SelctnMtd,omitempty"`

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType2Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr,omitempty"`

	// Internal account reference in case of credit account.
	CreditReference *Max35Text `xml:"CdtRef,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`
}

func (c *CardAccount5) SetSelectionMethod(value string) {
	c.SelectionMethod = (*AccountChoiceMethod1Code)(&value)
}

func (c *CardAccount5) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType2Code)(&value)
}

func (c *CardAccount5) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount5) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount5) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount5) SetCreditReference(value string) {
	c.CreditReference = (*Max35Text)(&value)
}

func (c *CardAccount5) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

// Customer account information.
type CardAccount6 struct {

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType2Code `xml:"AcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr,omitempty"`

	// Internal account reference in case of credit account.
	CreditReference *Max35Text `xml:"CdtRef,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`

	// Balance of the account.
	Balance *AmountAndDirection43 `xml:"Bal,omitempty"`

	// Indicates if the balance must be displayed to the customer on the ATM.
	BalanceDisplayFlag *TrueFalseIndicator `xml:"BalDispFlg,omitempty"`

	// Indicates if this is the default account.
	DefaultAccountIndicator *TrueFalseIndicator `xml:"DfltAcctInd,omitempty"`

	// Services allowed on the account.
	AllowedServices []*ATMServiceType4Code `xml:"AllwdSvcs,omitempty"`

	// Withdrawal limits for the account.
	WithdrawalLimits *ATMTransactionAmounts4 `xml:"WdrwlLmts,omitempty"`

	// Deposit limits for the account.
	DepositLimits *ATMTransactionAmounts5 `xml:"DpstLmts,omitempty"`
}

func (c *CardAccount6) SetAccountType(value string) {
	c.AccountType = (*CardAccountType2Code)(&value)
}

func (c *CardAccount6) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount6) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount6) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount6) SetCreditReference(value string) {
	c.CreditReference = (*Max35Text)(&value)
}

func (c *CardAccount6) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

func (c *CardAccount6) AddBalance() *AmountAndDirection43 {
	c.Balance = new(AmountAndDirection43)
	return c.Balance
}

func (c *CardAccount6) SetBalanceDisplayFlag(value string) {
	c.BalanceDisplayFlag = (*TrueFalseIndicator)(&value)
}

func (c *CardAccount6) SetDefaultAccountIndicator(value string) {
	c.DefaultAccountIndicator = (*TrueFalseIndicator)(&value)
}

func (c *CardAccount6) AddAllowedServices(value string) {
	c.AllowedServices = append(c.AllowedServices, (*ATMServiceType4Code)(&value))
}

func (c *CardAccount6) AddWithdrawalLimits() *ATMTransactionAmounts4 {
	c.WithdrawalLimits = new(ATMTransactionAmounts4)
	return c.WithdrawalLimits
}

func (c *CardAccount6) AddDepositLimits() *ATMTransactionAmounts5 {
	c.DepositLimits = new(ATMTransactionAmounts5)
	return c.DepositLimits
}

// Customer account information.
type CardAccount7 struct {

	// Method used by the cardholder and the terminal for the choice of the account.
	SelectionMethod *AccountChoiceMethod1Code `xml:"SelctnMtd,omitempty"`

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType3Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Party that legally owns the account.
	AccountOwner *NameAndAddress3 `xml:"AcctOwnr,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`
}

func (c *CardAccount7) SetSelectionMethod(value string) {
	c.SelectionMethod = (*AccountChoiceMethod1Code)(&value)
}

func (c *CardAccount7) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType3Code)(&value)
}

func (c *CardAccount7) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount7) AddAccountOwner() *NameAndAddress3 {
	c.AccountOwner = new(NameAndAddress3)
	return c.AccountOwner
}

func (c *CardAccount7) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount7) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount7) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

// Customer account information.
type CardAccount8 struct {

	// Method used by the cardholder and the terminal for the choice of the account.
	SelectionMethod *AccountChoiceMethod1Code `xml:"SelctnMtd,omitempty"`

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType3Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Party that legally owns the account.
	AccountOwner *NameAndAddress3 `xml:"AcctOwnr,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr,omitempty"`

	// Internal account reference in case of credit account.
	CreditReference *Max35Text `xml:"CdtRef,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`

	// Balance of the account.
	Balance *AmountAndDirection43 `xml:"Bal,omitempty"`

	// Indicates if the balance must be displayed to the customer on the ATM.
	BalanceDisplayFlag *TrueFalseIndicator `xml:"BalDispFlg,omitempty"`

	// Indicates if this is the default account.
	DefaultAccountIndicator *TrueFalseIndicator `xml:"DfltAcctInd,omitempty"`
}

func (c *CardAccount8) SetSelectionMethod(value string) {
	c.SelectionMethod = (*AccountChoiceMethod1Code)(&value)
}

func (c *CardAccount8) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType3Code)(&value)
}

func (c *CardAccount8) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount8) AddAccountOwner() *NameAndAddress3 {
	c.AccountOwner = new(NameAndAddress3)
	return c.AccountOwner
}

func (c *CardAccount8) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount8) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount8) SetCreditReference(value string) {
	c.CreditReference = (*Max35Text)(&value)
}

func (c *CardAccount8) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}

func (c *CardAccount8) AddBalance() *AmountAndDirection43 {
	c.Balance = new(AmountAndDirection43)
	return c.Balance
}

func (c *CardAccount8) SetBalanceDisplayFlag(value string) {
	c.BalanceDisplayFlag = (*TrueFalseIndicator)(&value)
}

func (c *CardAccount8) SetDefaultAccountIndicator(value string) {
	c.DefaultAccountIndicator = (*TrueFalseIndicator)(&value)
}

// Customer account information.
type CardAccount9 struct {

	// Sequence number of the account data for multi-account deposit.
	AccountSequenceNumber *Number `xml:"AcctSeqNb,omitempty"`

	// Method used by the cardholder and the terminal for the choice of the account.
	SelectionMethod *AccountChoiceMethod1Code `xml:"SelctnMtd,omitempty"`

	// Type of cardholder account used for the transaction.
	SelectedAccountType *CardAccountType3Code `xml:"SelctdAcctTp,omitempty"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Party that legally owns the account.
	AccountOwner *NameAndAddress3 `xml:"AcctOwnr,omitempty"`

	// Identification of the currency in which the account is held.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr,omitempty"`

	// Internal account reference in case of credit account.
	CreditReference *Max35Text `xml:"CdtRef,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	Servicer *PartyIdentification72Choice `xml:"Svcr,omitempty"`
}

func (c *CardAccount9) SetAccountSequenceNumber(value string) {
	c.AccountSequenceNumber = (*Number)(&value)
}

func (c *CardAccount9) SetSelectionMethod(value string) {
	c.SelectionMethod = (*AccountChoiceMethod1Code)(&value)
}

func (c *CardAccount9) SetSelectedAccountType(value string) {
	c.SelectedAccountType = (*CardAccountType3Code)(&value)
}

func (c *CardAccount9) SetAccountName(value string) {
	c.AccountName = (*Max70Text)(&value)
}

func (c *CardAccount9) AddAccountOwner() *NameAndAddress3 {
	c.AccountOwner = new(NameAndAddress3)
	return c.AccountOwner
}

func (c *CardAccount9) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CardAccount9) AddAccountIdentifier() *AccountIdentification31Choice {
	c.AccountIdentifier = new(AccountIdentification31Choice)
	return c.AccountIdentifier
}

func (c *CardAccount9) SetCreditReference(value string) {
	c.CreditReference = (*Max35Text)(&value)
}

func (c *CardAccount9) AddServicer() *PartyIdentification72Choice {
	c.Servicer = new(PartyIdentification72Choice)
	return c.Servicer
}
