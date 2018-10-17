package iso20022

// Provides details on the account notification.
type OriginalNotificationReference1 struct {

	// Identifies the account to be credited with the incoming amount of money.
	Account *CashAccount16 `xml:"Acct,omitempty"`

	// Party that legally owns the account.
	AccountOwner *Party12Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Identifies the parent account of the account to be credited with the incoming amount of money.
	RelatedAccount *CashAccount16 `xml:"RltdAcct,omitempty"`

	// Sum of the amounts in all the Item entries.
	TotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *Party12Choice `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt,omitempty"`

	// Provides details of the expected amount on the account serviced by the account servicer.
	OriginalItem []*OriginalItem2 `xml:"OrgnlItm"`
}

func (o *OriginalNotificationReference1) AddAccount() *CashAccount16 {
	o.Account = new(CashAccount16)
	return o.Account
}

func (o *OriginalNotificationReference1) AddAccountOwner() *Party12Choice {
	o.AccountOwner = new(Party12Choice)
	return o.AccountOwner
}

func (o *OriginalNotificationReference1) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	o.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return o.AccountServicer
}

func (o *OriginalNotificationReference1) AddRelatedAccount() *CashAccount16 {
	o.RelatedAccount = new(CashAccount16)
	return o.RelatedAccount
}

func (o *OriginalNotificationReference1) SetTotalAmount(value, currency string) {
	o.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalNotificationReference1) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalNotificationReference1) AddDebtor() *Party12Choice {
	o.Debtor = new(Party12Choice)
	return o.Debtor
}

func (o *OriginalNotificationReference1) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalNotificationReference1) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.IntermediaryAgent
}

func (o *OriginalNotificationReference1) AddOriginalItem() *OriginalItem2 {
	newValue := new(OriginalItem2)
	o.OriginalItem = append(o.OriginalItem, newValue)
	return newValue
}

// Provides details on the account notification.
type OriginalNotificationReference2 struct {

	// Identifies the account to be credited with the incoming amount of money.
	Account *CashAccount16 `xml:"Acct,omitempty"`

	// Party that legally owns the account.
	AccountOwner *Party12Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Identifies the parent account of the account to be credited with the incoming amount of money.
	RelatedAccount *CashAccount16 `xml:"RltdAcct,omitempty"`

	// Sum of the amounts in all the Item entries.
	TotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *Party12Choice `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt,omitempty"`

	// Provides details of the expected amount on the account serviced by the account servicer.
	OriginalItemAndStatus []*OriginalItemAndStatus2 `xml:"OrgnlItmAndSts"`
}

func (o *OriginalNotificationReference2) AddAccount() *CashAccount16 {
	o.Account = new(CashAccount16)
	return o.Account
}

func (o *OriginalNotificationReference2) AddAccountOwner() *Party12Choice {
	o.AccountOwner = new(Party12Choice)
	return o.AccountOwner
}

func (o *OriginalNotificationReference2) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	o.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return o.AccountServicer
}

func (o *OriginalNotificationReference2) AddRelatedAccount() *CashAccount16 {
	o.RelatedAccount = new(CashAccount16)
	return o.RelatedAccount
}

func (o *OriginalNotificationReference2) SetTotalAmount(value, currency string) {
	o.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalNotificationReference2) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalNotificationReference2) AddDebtor() *Party12Choice {
	o.Debtor = new(Party12Choice)
	return o.Debtor
}

func (o *OriginalNotificationReference2) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalNotificationReference2) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.IntermediaryAgent
}

func (o *OriginalNotificationReference2) AddOriginalItemAndStatus() *OriginalItemAndStatus2 {
	newValue := new(OriginalItemAndStatus2)
	o.OriginalItemAndStatus = append(o.OriginalItemAndStatus, newValue)
	return newValue
}

// Provides details on the account notification.
type OriginalNotificationReference3 struct {

	// Identifies the account to be credited with the incoming amount of money.
	Account *CashAccount24 `xml:"Acct,omitempty"`

	// Party that legally owns the account.
	AccountOwner *Party12Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Identifies the parent account of the account to be credited with the incoming amount of money.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Sum of the amounts in all the Item entries.
	TotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *Party12Choice `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt,omitempty"`

	// Provides details of the expected amount on the account serviced by the account servicer.
	OriginalItemAndStatus []*OriginalItemAndStatus3 `xml:"OrgnlItmAndSts"`
}

func (o *OriginalNotificationReference3) AddAccount() *CashAccount24 {
	o.Account = new(CashAccount24)
	return o.Account
}

func (o *OriginalNotificationReference3) AddAccountOwner() *Party12Choice {
	o.AccountOwner = new(Party12Choice)
	return o.AccountOwner
}

func (o *OriginalNotificationReference3) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	o.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return o.AccountServicer
}

func (o *OriginalNotificationReference3) AddRelatedAccount() *CashAccount24 {
	o.RelatedAccount = new(CashAccount24)
	return o.RelatedAccount
}

func (o *OriginalNotificationReference3) SetTotalAmount(value, currency string) {
	o.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalNotificationReference3) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalNotificationReference3) AddDebtor() *Party12Choice {
	o.Debtor = new(Party12Choice)
	return o.Debtor
}

func (o *OriginalNotificationReference3) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalNotificationReference3) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.IntermediaryAgent
}

func (o *OriginalNotificationReference3) AddOriginalItemAndStatus() *OriginalItemAndStatus3 {
	newValue := new(OriginalItemAndStatus3)
	o.OriginalItemAndStatus = append(o.OriginalItemAndStatus, newValue)
	return newValue
}

// Provides details on the account notification.
type OriginalNotificationReference4 struct {

	// Identifies the account to be credited with the incoming amount of money.
	Account *CashAccount24 `xml:"Acct,omitempty"`

	// Party that legally owns the account.
	AccountOwner *Party12Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Identifies the parent account of the account to be credited with the incoming amount of money.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Sum of the amounts in all the Item entries.
	TotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *Party12Choice `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt,omitempty"`

	// Provides details of the expected amount on the account serviced by the account servicer.
	OriginalItem []*OriginalItem3 `xml:"OrgnlItm"`
}

func (o *OriginalNotificationReference4) AddAccount() *CashAccount24 {
	o.Account = new(CashAccount24)
	return o.Account
}

func (o *OriginalNotificationReference4) AddAccountOwner() *Party12Choice {
	o.AccountOwner = new(Party12Choice)
	return o.AccountOwner
}

func (o *OriginalNotificationReference4) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	o.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return o.AccountServicer
}

func (o *OriginalNotificationReference4) AddRelatedAccount() *CashAccount24 {
	o.RelatedAccount = new(CashAccount24)
	return o.RelatedAccount
}

func (o *OriginalNotificationReference4) SetTotalAmount(value, currency string) {
	o.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalNotificationReference4) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalNotificationReference4) AddDebtor() *Party12Choice {
	o.Debtor = new(Party12Choice)
	return o.Debtor
}

func (o *OriginalNotificationReference4) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalNotificationReference4) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.IntermediaryAgent
}

func (o *OriginalNotificationReference4) AddOriginalItem() *OriginalItem3 {
	newValue := new(OriginalItem3)
	o.OriginalItem = append(o.OriginalItem, newValue)
	return newValue
}

// Provides details on the account notification.
type OriginalNotificationReference5 struct {

	// Identifies the account to be credited with the incoming amount of money.
	Account *CashAccount24 `xml:"Acct,omitempty"`

	// Party that legally owns the account.
	AccountOwner *Party12Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Identifies the parent account of the account to be credited with the incoming amount of money.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Sum of the amounts in all the Item entries.
	TotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *Party12Choice `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt,omitempty"`

	// Provides details of the expected amount on the account serviced by the account servicer.
	OriginalItemAndStatus []*OriginalItemAndStatus4 `xml:"OrgnlItmAndSts"`
}

func (o *OriginalNotificationReference5) AddAccount() *CashAccount24 {
	o.Account = new(CashAccount24)
	return o.Account
}

func (o *OriginalNotificationReference5) AddAccountOwner() *Party12Choice {
	o.AccountOwner = new(Party12Choice)
	return o.AccountOwner
}

func (o *OriginalNotificationReference5) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	o.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return o.AccountServicer
}

func (o *OriginalNotificationReference5) AddRelatedAccount() *CashAccount24 {
	o.RelatedAccount = new(CashAccount24)
	return o.RelatedAccount
}

func (o *OriginalNotificationReference5) SetTotalAmount(value, currency string) {
	o.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalNotificationReference5) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalNotificationReference5) AddDebtor() *Party12Choice {
	o.Debtor = new(Party12Choice)
	return o.Debtor
}

func (o *OriginalNotificationReference5) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalNotificationReference5) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.IntermediaryAgent
}

func (o *OriginalNotificationReference5) AddOriginalItemAndStatus() *OriginalItemAndStatus4 {
	newValue := new(OriginalItemAndStatus4)
	o.OriginalItemAndStatus = append(o.OriginalItemAndStatus, newValue)
	return newValue
}

// Provides details on the account notification.
type OriginalNotificationReference6 struct {

	// Identifies the account to be credited with the incoming amount of money.
	Account *CashAccount24 `xml:"Acct,omitempty"`

	// Party that legally owns the account.
	AccountOwner *Party12Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Identifies the parent account of the account to be credited with the incoming amount of money.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Sum of the amounts in all the Item entries.
	TotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *Party12Choice `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt,omitempty"`

	// Provides details of the expected amount on the account serviced by the account servicer.
	OriginalItem []*OriginalItem4 `xml:"OrgnlItm"`
}

func (o *OriginalNotificationReference6) AddAccount() *CashAccount24 {
	o.Account = new(CashAccount24)
	return o.Account
}

func (o *OriginalNotificationReference6) AddAccountOwner() *Party12Choice {
	o.AccountOwner = new(Party12Choice)
	return o.AccountOwner
}

func (o *OriginalNotificationReference6) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	o.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return o.AccountServicer
}

func (o *OriginalNotificationReference6) AddRelatedAccount() *CashAccount24 {
	o.RelatedAccount = new(CashAccount24)
	return o.RelatedAccount
}

func (o *OriginalNotificationReference6) SetTotalAmount(value, currency string) {
	o.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalNotificationReference6) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalNotificationReference6) AddDebtor() *Party12Choice {
	o.Debtor = new(Party12Choice)
	return o.Debtor
}

func (o *OriginalNotificationReference6) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalNotificationReference6) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.IntermediaryAgent
}

func (o *OriginalNotificationReference6) AddOriginalItem() *OriginalItem4 {
	newValue := new(OriginalItem4)
	o.OriginalItem = append(o.OriginalItem, newValue)
	return newValue
}

// Provides details on the account notification.
type OriginalNotificationReference7 struct {

	// Identifies the account to be credited with the incoming amount of money.
	Account *CashAccount24 `xml:"Acct,omitempty"`

	// Party that legally owns the account.
	AccountOwner *Party12Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Identifies the parent account of the account to be credited with the incoming amount of money.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Sum of the amounts in all the Item entries.
	TotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *Party12Choice `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt,omitempty"`

	// Provides details of the expected amount on the account serviced by the account servicer.
	OriginalItemAndStatus []*OriginalItemAndStatus5 `xml:"OrgnlItmAndSts"`
}

func (o *OriginalNotificationReference7) AddAccount() *CashAccount24 {
	o.Account = new(CashAccount24)
	return o.Account
}

func (o *OriginalNotificationReference7) AddAccountOwner() *Party12Choice {
	o.AccountOwner = new(Party12Choice)
	return o.AccountOwner
}

func (o *OriginalNotificationReference7) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	o.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return o.AccountServicer
}

func (o *OriginalNotificationReference7) AddRelatedAccount() *CashAccount24 {
	o.RelatedAccount = new(CashAccount24)
	return o.RelatedAccount
}

func (o *OriginalNotificationReference7) SetTotalAmount(value, currency string) {
	o.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalNotificationReference7) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalNotificationReference7) AddDebtor() *Party12Choice {
	o.Debtor = new(Party12Choice)
	return o.Debtor
}

func (o *OriginalNotificationReference7) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalNotificationReference7) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.IntermediaryAgent
}

func (o *OriginalNotificationReference7) AddOriginalItemAndStatus() *OriginalItemAndStatus5 {
	newValue := new(OriginalItemAndStatus5)
	o.OriginalItemAndStatus = append(o.OriginalItemAndStatus, newValue)
	return newValue
}

// Provides details on the account notification.
type OriginalNotificationReference8 struct {

	// Identifies the account to be credited with the incoming amount of money.
	Account *CashAccount24 `xml:"Acct,omitempty"`

	// Party that legally owns the account.
	AccountOwner *Party12Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Identifies the parent account of the account to be credited with the incoming amount of money.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Sum of the amounts in all the Item entries.
	TotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *Party12Choice `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt,omitempty"`

	// Provides details of the expected amount on the account serviced by the account servicer.
	OriginalItem []*OriginalItem5 `xml:"OrgnlItm"`
}

func (o *OriginalNotificationReference8) AddAccount() *CashAccount24 {
	o.Account = new(CashAccount24)
	return o.Account
}

func (o *OriginalNotificationReference8) AddAccountOwner() *Party12Choice {
	o.AccountOwner = new(Party12Choice)
	return o.AccountOwner
}

func (o *OriginalNotificationReference8) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	o.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return o.AccountServicer
}

func (o *OriginalNotificationReference8) AddRelatedAccount() *CashAccount24 {
	o.RelatedAccount = new(CashAccount24)
	return o.RelatedAccount
}

func (o *OriginalNotificationReference8) SetTotalAmount(value, currency string) {
	o.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalNotificationReference8) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalNotificationReference8) AddDebtor() *Party12Choice {
	o.Debtor = new(Party12Choice)
	return o.Debtor
}

func (o *OriginalNotificationReference8) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalNotificationReference8) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.IntermediaryAgent
}

func (o *OriginalNotificationReference8) AddOriginalItem() *OriginalItem5 {
	newValue := new(OriginalItem5)
	o.OriginalItem = append(o.OriginalItem, newValue)
	return newValue
}
