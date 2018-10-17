package iso20022

// Account to or from which a securities entry is made.
type SubAccountIdentification1 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether the securities in the account are fungible, ie, interchangeable.
	FungibleIndicator *YesNoIndicator `xml:"FngbInd"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForSubAccount []*AggregateBalanceInformation1 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification1) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SubAccountIdentification1) SetFungibleIndicator(value string) {
	s.FungibleIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification1) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification1) AddBalanceForSubAccount() *AggregateBalanceInformation1 {
	newValue := new(AggregateBalanceInformation1)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification10 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation8 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification10) AddAccountOwner() *PartyIdentification13Choice {
	s.AccountOwner = new(PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification10) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification10) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification10) AddBalanceForSubAccount() *AggregateBalanceInformation8 {
	newValue := new(AggregateBalanceInformation8)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification11 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation9 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification11) AddAccountOwner() *PartyIdentification13Choice {
	s.AccountOwner = new(PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification11) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification11) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification11) AddBalanceForSubAccount() *AggregateBalanceInformation9 {
	newValue := new(AggregateBalanceInformation9)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Sub-account reporting.
type SubAccountIdentification15 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails5 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification15) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification15) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification15) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification15) AddFinancialInstrumentDetails() *FinancialInstrumentDetails5 {
	newValue := new(FinancialInstrumentDetails5)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification16 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation13 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification16) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification16) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification16) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification16) AddBalanceForSubAccount() *AggregateBalanceInformation13 {
	newValue := new(AggregateBalanceInformation13)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification17 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation12 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification17) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification17) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification17) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification17) AddBalanceForSubAccount() *AggregateBalanceInformation12 {
	newValue := new(AggregateBalanceInformation12)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification2 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether the securities in the account are fungible, ie, interchangeable.
	FungibleIndicator *YesNoIndicator `xml:"FngbInd"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForSubAccount []*AggregateBalanceInformation2 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification2) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SubAccountIdentification2) SetFungibleIndicator(value string) {
	s.FungibleIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification2) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification2) AddBalanceForSubAccount() *AggregateBalanceInformation2 {
	newValue := new(AggregateBalanceInformation2)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Sub-account reporting.
type SubAccountIdentification21 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails9 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification21) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification21) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification21) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification21) AddFinancialInstrumentDetails() *FinancialInstrumentDetails9 {
	newValue := new(FinancialInstrumentDetails9)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification22 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation16 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification22) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification22) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification22) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification22) AddBalanceForSubAccount() *AggregateBalanceInformation16 {
	newValue := new(AggregateBalanceInformation16)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification23 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation17 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification23) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification23) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification23) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification23) AddBalanceForSubAccount() *AggregateBalanceInformation17 {
	newValue := new(AggregateBalanceInformation17)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification28 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation21 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification28) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification28) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification28) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification28) AddBalanceForSubAccount() *AggregateBalanceInformation21 {
	newValue := new(AggregateBalanceInformation21)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification29 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation22 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification29) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification29) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification29) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification29) AddBalanceForSubAccount() *AggregateBalanceInformation22 {
	newValue := new(AggregateBalanceInformation22)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification3 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether the securities in the account are fungible, ie, interchangeable.
	FungibleIndicator *YesNoIndicator `xml:"FngbInd"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForSubAccount []*AggregateBalanceInformation3 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification3) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SubAccountIdentification3) SetFungibleIndicator(value string) {
	s.FungibleIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification3) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification3) AddBalanceForSubAccount() *AggregateBalanceInformation3 {
	newValue := new(AggregateBalanceInformation3)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Sub-account reporting.
type SubAccountIdentification30 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails13 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification30) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification30) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification30) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification30) AddFinancialInstrumentDetails() *FinancialInstrumentDetails13 {
	newValue := new(FinancialInstrumentDetails13)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}

// Sub-account reporting.
type SubAccountIdentification34 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails17 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification34) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification34) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification34) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification34) AddFinancialInstrumentDetails() *FinancialInstrumentDetails17 {
	newValue := new(FinancialInstrumentDetails17)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification36 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Creation/cancellation of investment units on the books of the fund or its designated agent, as a result of executing an investment fund order.
	TransactionOnSubAccount []*InvestmentFundTransactionsByFund3 `xml:"TxOnSubAcct,omitempty"`
}

func (s *SubAccountIdentification36) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SubAccountIdentification36) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification36) AddTransactionOnSubAccount() *InvestmentFundTransactionsByFund3 {
	newValue := new(InvestmentFundTransactionsByFund3)
	s.TransactionOnSubAccount = append(s.TransactionOnSubAccount, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification37 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation25 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification37) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification37) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification37) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification37) AddBalanceForSubAccount() *AggregateBalanceInformation25 {
	newValue := new(AggregateBalanceInformation25)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification38 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation26 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification38) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification38) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification38) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification38) AddBalanceForSubAccount() *AggregateBalanceInformation26 {
	newValue := new(AggregateBalanceInformation26)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type SubAccountIdentification4 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Creation/cancellation of investment units on the books of the fund or its designated agent, as a result of executing an investment fund order.
	TransactionOnSubAccount []*InvestmentFundTransactionsByFund1 `xml:"TxOnSubAcct,omitempty"`
}

func (s *SubAccountIdentification4) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SubAccountIdentification4) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification4) AddTransactionOnSubAccount() *InvestmentFundTransactionsByFund1 {
	newValue := new(InvestmentFundTransactionsByFund1)
	s.TransactionOnSubAccount = append(s.TransactionOnSubAccount, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification42 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount25 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation30 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification42) AddAccountOwner() *PartyIdentification98 {
	s.AccountOwner = new(PartyIdentification98)
	return s.AccountOwner
}

func (s *SubAccountIdentification42) AddSafekeepingAccount() *SecuritiesAccount25 {
	s.SafekeepingAccount = new(SecuritiesAccount25)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification42) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification42) AddBalanceForSubAccount() *AggregateBalanceInformation30 {
	newValue := new(AggregateBalanceInformation30)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification43 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount25 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation31 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification43) AddAccountOwner() *PartyIdentification98 {
	s.AccountOwner = new(PartyIdentification98)
	return s.AccountOwner
}

func (s *SubAccountIdentification43) AddSafekeepingAccount() *SecuritiesAccount25 {
	s.SafekeepingAccount = new(SecuritiesAccount25)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification43) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification43) AddBalanceForSubAccount() *AggregateBalanceInformation31 {
	newValue := new(AggregateBalanceInformation31)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Provides all sub-account details.
type SubAccountIdentification44 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount25 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails20 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification44) AddAccountOwner() *PartyIdentification98 {
	s.AccountOwner = new(PartyIdentification98)
	return s.AccountOwner
}

func (s *SubAccountIdentification44) AddSafekeepingAccount() *SecuritiesAccount25 {
	s.SafekeepingAccount = new(SecuritiesAccount25)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification44) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification44) AddFinancialInstrumentDetails() *FinancialInstrumentDetails20 {
	newValue := new(FinancialInstrumentDetails20)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification45 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount34 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation32 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification45) AddAccountOwner() *PartyIdentification109 {
	s.AccountOwner = new(PartyIdentification109)
	return s.AccountOwner
}

func (s *SubAccountIdentification45) AddSafekeepingAccount() *SecuritiesAccount34 {
	s.SafekeepingAccount = new(SecuritiesAccount34)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification45) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification45) AddBalanceForSubAccount() *AggregateBalanceInformation32 {
	newValue := new(AggregateBalanceInformation32)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification46 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount34 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities subaccount.
	BalanceForSubAccount []*AggregateBalanceInformation33 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification46) AddAccountOwner() *PartyIdentification109 {
	s.AccountOwner = new(PartyIdentification109)
	return s.AccountOwner
}

func (s *SubAccountIdentification46) AddSafekeepingAccount() *SecuritiesAccount34 {
	s.SafekeepingAccount = new(SecuritiesAccount34)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification46) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification46) AddBalanceForSubAccount() *AggregateBalanceInformation33 {
	newValue := new(AggregateBalanceInformation33)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Provides all sub-account details.
type SubAccountIdentification47 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount34 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails23 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification47) AddAccountOwner() *PartyIdentification109 {
	s.AccountOwner = new(PartyIdentification109)
	return s.AccountOwner
}

func (s *SubAccountIdentification47) AddSafekeepingAccount() *SecuritiesAccount34 {
	s.SafekeepingAccount = new(SecuritiesAccount34)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification47) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification47) AddFinancialInstrumentDetails() *FinancialInstrumentDetails23 {
	newValue := new(FinancialInstrumentDetails23)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}

// Provides all sub-account details.
type SubAccountIdentification49 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount25 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails25 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification49) AddAccountOwner() *PartyIdentification98 {
	s.AccountOwner = new(PartyIdentification98)
	return s.AccountOwner
}

func (s *SubAccountIdentification49) AddSafekeepingAccount() *SecuritiesAccount25 {
	s.SafekeepingAccount = new(SecuritiesAccount25)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification49) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification49) AddFinancialInstrumentDetails() *FinancialInstrumentDetails25 {
	newValue := new(FinancialInstrumentDetails25)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification5 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether the securities in the account are fungible, ie, interchangeable.
	FungibleIndicator *YesNoIndicator `xml:"FngbInd"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForSubAccount []*AggregateBalanceInformation4 `xml:"BalForSubAcct,omitempty"`
}

func (s *SubAccountIdentification5) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SubAccountIdentification5) SetFungibleIndicator(value string) {
	s.FungibleIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification5) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification5) AddBalanceForSubAccount() *AggregateBalanceInformation4 {
	newValue := new(AggregateBalanceInformation4)
	s.BalanceForSubAccount = append(s.BalanceForSubAccount, newValue)
	return newValue
}

// Provides all sub-account details.
type SubAccountIdentification50 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount34 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails27 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification50) AddAccountOwner() *PartyIdentification119 {
	s.AccountOwner = new(PartyIdentification119)
	return s.AccountOwner
}

func (s *SubAccountIdentification50) AddSafekeepingAccount() *SecuritiesAccount34 {
	s.SafekeepingAccount = new(SecuritiesAccount34)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification50) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification50) AddFinancialInstrumentDetails() *FinancialInstrumentDetails27 {
	newValue := new(FinancialInstrumentDetails27)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}

// Account to or from which a securities entry is made.
type SubAccountIdentification6 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Creation/cancellation of investment units on the books of the fund or its designated agent, as a result of executing an investment fund order.
	TransactionOnSubAccount []*InvestmentFundTransactionsByFund2 `xml:"TxOnSubAcct,omitempty"`
}

func (s *SubAccountIdentification6) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SubAccountIdentification6) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification6) AddTransactionOnSubAccount() *InvestmentFundTransactionsByFund2 {
	newValue := new(InvestmentFundTransactionsByFund2)
	s.TransactionOnSubAccount = append(s.TransactionOnSubAccount, newValue)
	return newValue
}

// Sub-account reporting.
type SubAccountIdentification9 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount14 `xml:"SfkpgAcct"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*FinancialInstrumentDetails2 `xml:"FinInstrmDtls,omitempty"`
}

func (s *SubAccountIdentification9) AddAccountOwner() *PartyIdentification13Choice {
	s.AccountOwner = new(PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SubAccountIdentification9) AddSafekeepingAccount() *SecuritiesAccount14 {
	s.SafekeepingAccount = new(SecuritiesAccount14)
	return s.SafekeepingAccount
}

func (s *SubAccountIdentification9) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *SubAccountIdentification9) AddFinancialInstrumentDetails() *FinancialInstrumentDetails2 {
	newValue := new(FinancialInstrumentDetails2)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}
