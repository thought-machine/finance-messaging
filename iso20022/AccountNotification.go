package iso20022

// Set of elements providing further details on the account notification.
type AccountNotification1 struct {

	// Unique and unambiguous identification of the account report, assigned by the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Sequential number of the report, assigned by the account servicer. It is increased incrementally for each report sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the report, assigned by the account servicer. It is increased incrementally for each report sent.
	//
	// Usage : in those scenarios where eg a paper statement is a legal requirement, the paper statement may have a different numbering than the electronic sequential number. Paper statements can for instance only be sent if movement on the account has taken place, whereas electronic statements can be sent eg each day, regardless of whether movements have taken place or not.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the report was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between the start date and the end date for which the account notification is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Specifies if this document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *CashAccount13 `xml:"Acct"`

	// Identifies the parent account of the reported account.
	RelatedAccount *CashAccount7 `xml:"RltdAcct,omitempty"`

	// Set of element providing summary information on entries.
	TransactionsSummary *TotalTransactions1 `xml:"TxsSummry,omitempty"`

	// Specifies the elements of an entry in the report.
	//
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*NotificationEntry1 `xml:"Ntry,omitempty"`

	// Further details on the account notification.
	AdditionalNotificationInformation *Max500Text `xml:"AddtlNtfctnInf,omitempty"`
}

func (a *AccountNotification1) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountNotification1) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification1) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification1) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountNotification1) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountNotification1) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountNotification1) AddAccount() *CashAccount13 {
	a.Account = new(CashAccount13)
	return a.Account
}

func (a *AccountNotification1) AddRelatedAccount() *CashAccount7 {
	a.RelatedAccount = new(CashAccount7)
	return a.RelatedAccount
}

func (a *AccountNotification1) AddTransactionsSummary() *TotalTransactions1 {
	a.TransactionsSummary = new(TotalTransactions1)
	return a.TransactionsSummary
}

func (a *AccountNotification1) AddEntry() *NotificationEntry1 {
	newValue := new(NotificationEntry1)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountNotification1) SetAdditionalNotificationInformation(value string) {
	a.AdditionalNotificationInformation = (*Max500Text)(&value)
}

// Provides details on the account notification.
type AccountNotification10 struct {

	// Unique identification, as assigned by the account owner, to unambiguously identify the account notification.
	Identification *Max35Text `xml:"Id"`

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
	Item []*NotificationItem5 `xml:"Itm"`
}

func (a *AccountNotification10) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountNotification10) AddAccount() *CashAccount24 {
	a.Account = new(CashAccount24)
	return a.Account
}

func (a *AccountNotification10) AddAccountOwner() *Party12Choice {
	a.AccountOwner = new(Party12Choice)
	return a.AccountOwner
}

func (a *AccountNotification10) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicer
}

func (a *AccountNotification10) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountNotification10) SetTotalAmount(value, currency string) {
	a.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AccountNotification10) SetExpectedValueDate(value string) {
	a.ExpectedValueDate = (*ISODate)(&value)
}

func (a *AccountNotification10) AddDebtor() *Party12Choice {
	a.Debtor = new(Party12Choice)
	return a.Debtor
}

func (a *AccountNotification10) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.DebtorAgent
}

func (a *AccountNotification10) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.IntermediaryAgent
}

func (a *AccountNotification10) AddItem() *NotificationItem5 {
	newValue := new(NotificationItem5)
	a.Item = append(a.Item, newValue)
	return newValue
}

// Provides further details of the account notification.
type AccountNotification11 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account notification.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the notification.
	//
	// Usage: The pagination of the notification is only allowed when agreed between the parties.
	NotificationPagination *Pagination `xml:"NtfctnPgntn,omitempty"`

	// Sequential number of the notification, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each notification sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the notification, as assigned by the account servicer. It is increased incrementally for each notification sent.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account notification is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the notification has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest3 `xml:"Intrst,omitempty"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions4 `xml:"TxsSummry,omitempty"`

	// Specifies an entry in the debit credit notification.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	//
	// Usage Rule:  In case of a Payments R-transaction the creditor / debtor referenced of the original payment initiation messages is also used for reporting of the R-transaction. The original debtor/creditor in the reporting of R-Transactions is not inverted.
	// Following elements all defined in the TransactionDetails in RelatedParties or RelatedAgents are impacted by this usage rule:
	// Creditor, UltimateCreditor, CreditorAccount, CreditorAgent, Debtor, UltimateDebtor, DebtorAccount and DebtorAgent.
	//
	Entry []*ReportEntry7 `xml:"Ntry,omitempty"`

	// Further details of the account notification.
	AdditionalNotificationInformation *Max500Text `xml:"AddtlNtfctnInf,omitempty"`
}

func (a *AccountNotification11) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountNotification11) AddNotificationPagination() *Pagination {
	a.NotificationPagination = new(Pagination)
	return a.NotificationPagination
}

func (a *AccountNotification11) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification11) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification11) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountNotification11) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountNotification11) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountNotification11) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountNotification11) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountNotification11) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountNotification11) AddInterest() *AccountInterest3 {
	newValue := new(AccountInterest3)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountNotification11) AddTransactionsSummary() *TotalTransactions4 {
	a.TransactionsSummary = new(TotalTransactions4)
	return a.TransactionsSummary
}

func (a *AccountNotification11) AddEntry() *ReportEntry7 {
	newValue := new(ReportEntry7)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountNotification11) SetAdditionalNotificationInformation(value string) {
	a.AdditionalNotificationInformation = (*Max500Text)(&value)
}

// Provides further details of the account notification.
type AccountNotification12 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account notification.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the notification.
	//
	// Usage: The pagination of the notification is only allowed when agreed between the parties.
	NotificationPagination *Pagination `xml:"NtfctnPgntn,omitempty"`

	// Sequential number of the notification, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each notification sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the notification, as assigned by the account servicer. It is increased incrementally for each notification sent.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account notification is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the notification has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest3 `xml:"Intrst,omitempty"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions5 `xml:"TxsSummry,omitempty"`

	// Specifies an entry in the debit credit notification.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	//
	// Usage Rule:  In case of a Payments R-transaction the creditor / debtor referenced of the original payment initiation messages is also used for reporting of the R-transaction. The original debtor/creditor in the reporting of R-Transactions is not inverted.
	// Following elements all defined in the TransactionDetails in RelatedParties or RelatedAgents are impacted by this usage rule:
	// Creditor, UltimateCreditor, CreditorAccount, CreditorAgent, Debtor, UltimateDebtor, DebtorAccount and DebtorAgent.
	//
	Entry []*ReportEntry8 `xml:"Ntry,omitempty"`

	// Further details of the account notification.
	AdditionalNotificationInformation *Max500Text `xml:"AddtlNtfctnInf,omitempty"`
}

func (a *AccountNotification12) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountNotification12) AddNotificationPagination() *Pagination {
	a.NotificationPagination = new(Pagination)
	return a.NotificationPagination
}

func (a *AccountNotification12) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification12) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification12) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountNotification12) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountNotification12) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountNotification12) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountNotification12) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountNotification12) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountNotification12) AddInterest() *AccountInterest3 {
	newValue := new(AccountInterest3)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountNotification12) AddTransactionsSummary() *TotalTransactions5 {
	a.TransactionsSummary = new(TotalTransactions5)
	return a.TransactionsSummary
}

func (a *AccountNotification12) AddEntry() *ReportEntry8 {
	newValue := new(ReportEntry8)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountNotification12) SetAdditionalNotificationInformation(value string) {
	a.AdditionalNotificationInformation = (*Max500Text)(&value)
}

// Provides details on the account notification.
type AccountNotification13 struct {

	// Unique identification, as assigned by the account owner, to unambiguously identify the account notification.
	Identification *Max35Text `xml:"Id"`

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
	Item []*NotificationItem6 `xml:"Itm"`
}

func (a *AccountNotification13) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountNotification13) AddAccount() *CashAccount24 {
	a.Account = new(CashAccount24)
	return a.Account
}

func (a *AccountNotification13) AddAccountOwner() *Party12Choice {
	a.AccountOwner = new(Party12Choice)
	return a.AccountOwner
}

func (a *AccountNotification13) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicer
}

func (a *AccountNotification13) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountNotification13) SetTotalAmount(value, currency string) {
	a.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AccountNotification13) SetExpectedValueDate(value string) {
	a.ExpectedValueDate = (*ISODate)(&value)
}

func (a *AccountNotification13) AddDebtor() *Party12Choice {
	a.Debtor = new(Party12Choice)
	return a.Debtor
}

func (a *AccountNotification13) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.DebtorAgent
}

func (a *AccountNotification13) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.IntermediaryAgent
}

func (a *AccountNotification13) AddItem() *NotificationItem6 {
	newValue := new(NotificationItem6)
	a.Item = append(a.Item, newValue)
	return newValue
}

// Set of elements used to provide details of the account notification.
type AccountNotification2 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account notification.
	Identification *Max35Text `xml:"Id"`

	// Sequential number of the notification, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each notification sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the notification, as assigned by the account servicer. It is increased incrementally for each notification sent.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account notification is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount20 `xml:"Acct"`

	// Identifies the parent account of the account for which the notification has been issued.
	RelatedAccount *CashAccount16 `xml:"RltdAcct,omitempty"`

	// Set of elements used to provide general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest2 `xml:"Intrst,omitempty"`

	// Set of elements used to provide summary information on entries.
	TransactionsSummary *TotalTransactions2 `xml:"TxsSummry,omitempty"`

	// Set of elements used to specify an entry in the debit credit notification.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry2 `xml:"Ntry,omitempty"`

	// Further details of the account notification.
	AdditionalNotificationInformation *Max500Text `xml:"AddtlNtfctnInf,omitempty"`
}

func (a *AccountNotification2) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountNotification2) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification2) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification2) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountNotification2) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountNotification2) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountNotification2) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountNotification2) AddAccount() *CashAccount20 {
	a.Account = new(CashAccount20)
	return a.Account
}

func (a *AccountNotification2) AddRelatedAccount() *CashAccount16 {
	a.RelatedAccount = new(CashAccount16)
	return a.RelatedAccount
}

func (a *AccountNotification2) AddInterest() *AccountInterest2 {
	newValue := new(AccountInterest2)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountNotification2) AddTransactionsSummary() *TotalTransactions2 {
	a.TransactionsSummary = new(TotalTransactions2)
	return a.TransactionsSummary
}

func (a *AccountNotification2) AddEntry() *ReportEntry2 {
	newValue := new(ReportEntry2)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountNotification2) SetAdditionalNotificationInformation(value string) {
	a.AdditionalNotificationInformation = (*Max500Text)(&value)
}

// Provides details on the account notification.
type AccountNotification4 struct {

	// Unique identification, as assigned by the account owner, to unambiguously identify the account notification.
	Identification *Max35Text `xml:"Id"`

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
	Item []*NotificationItem3 `xml:"Itm"`
}

func (a *AccountNotification4) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountNotification4) AddAccount() *CashAccount16 {
	a.Account = new(CashAccount16)
	return a.Account
}

func (a *AccountNotification4) AddAccountOwner() *Party12Choice {
	a.AccountOwner = new(Party12Choice)
	return a.AccountOwner
}

func (a *AccountNotification4) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicer
}

func (a *AccountNotification4) AddRelatedAccount() *CashAccount16 {
	a.RelatedAccount = new(CashAccount16)
	return a.RelatedAccount
}

func (a *AccountNotification4) SetTotalAmount(value, currency string) {
	a.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AccountNotification4) SetExpectedValueDate(value string) {
	a.ExpectedValueDate = (*ISODate)(&value)
}

func (a *AccountNotification4) AddDebtor() *Party12Choice {
	a.Debtor = new(Party12Choice)
	return a.Debtor
}

func (a *AccountNotification4) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.DebtorAgent
}

func (a *AccountNotification4) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.IntermediaryAgent
}

func (a *AccountNotification4) AddItem() *NotificationItem3 {
	newValue := new(NotificationItem3)
	a.Item = append(a.Item, newValue)
	return newValue
}

// Provides further details of the account notification.
type AccountNotification5 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account notification.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the notification.
	//
	// Usage: The pagination of the notification is only allowed when agreed between the parties.
	NotificationPagination *Pagination `xml:"NtfctnPgntn,omitempty"`

	// Sequential number of the notification, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each notification sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the notification, as assigned by the account servicer. It is increased incrementally for each notification sent.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account notification is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the notification has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest2 `xml:"Intrst,omitempty"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions2 `xml:"TxsSummry,omitempty"`

	// Set of elements used to specify an entry in the debit credit notification.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry3 `xml:"Ntry,omitempty"`

	// Further details of the account notification.
	AdditionalNotificationInformation *Max500Text `xml:"AddtlNtfctnInf,omitempty"`
}

func (a *AccountNotification5) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountNotification5) AddNotificationPagination() *Pagination {
	a.NotificationPagination = new(Pagination)
	return a.NotificationPagination
}

func (a *AccountNotification5) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification5) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification5) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountNotification5) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountNotification5) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountNotification5) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountNotification5) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountNotification5) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountNotification5) AddInterest() *AccountInterest2 {
	newValue := new(AccountInterest2)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountNotification5) AddTransactionsSummary() *TotalTransactions2 {
	a.TransactionsSummary = new(TotalTransactions2)
	return a.TransactionsSummary
}

func (a *AccountNotification5) AddEntry() *ReportEntry3 {
	newValue := new(ReportEntry3)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountNotification5) SetAdditionalNotificationInformation(value string) {
	a.AdditionalNotificationInformation = (*Max500Text)(&value)
}

// Provides details on the account notification.
type AccountNotification6 struct {

	// Unique identification, as assigned by the account owner, to unambiguously identify the account notification.
	Identification *Max35Text `xml:"Id"`

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
	Item []*NotificationItem4 `xml:"Itm"`
}

func (a *AccountNotification6) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountNotification6) AddAccount() *CashAccount24 {
	a.Account = new(CashAccount24)
	return a.Account
}

func (a *AccountNotification6) AddAccountOwner() *Party12Choice {
	a.AccountOwner = new(Party12Choice)
	return a.AccountOwner
}

func (a *AccountNotification6) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicer
}

func (a *AccountNotification6) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountNotification6) SetTotalAmount(value, currency string) {
	a.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AccountNotification6) SetExpectedValueDate(value string) {
	a.ExpectedValueDate = (*ISODate)(&value)
}

func (a *AccountNotification6) AddDebtor() *Party12Choice {
	a.Debtor = new(Party12Choice)
	return a.Debtor
}

func (a *AccountNotification6) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.DebtorAgent
}

func (a *AccountNotification6) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.IntermediaryAgent
}

func (a *AccountNotification6) AddItem() *NotificationItem4 {
	newValue := new(NotificationItem4)
	a.Item = append(a.Item, newValue)
	return newValue
}

// Provides further details of the account notification.
type AccountNotification7 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account notification.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the notification.
	//
	// Usage: The pagination of the notification is only allowed when agreed between the parties.
	NotificationPagination *Pagination `xml:"NtfctnPgntn,omitempty"`

	// Sequential number of the notification, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each notification sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the notification, as assigned by the account servicer. It is increased incrementally for each notification sent.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account notification is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the notification has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest2 `xml:"Intrst,omitempty"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions4 `xml:"TxsSummry,omitempty"`

	// Set of elements used to specify an entry in the debit credit notification.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry4 `xml:"Ntry,omitempty"`

	// Further details of the account notification.
	AdditionalNotificationInformation *Max500Text `xml:"AddtlNtfctnInf,omitempty"`
}

func (a *AccountNotification7) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountNotification7) AddNotificationPagination() *Pagination {
	a.NotificationPagination = new(Pagination)
	return a.NotificationPagination
}

func (a *AccountNotification7) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification7) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification7) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountNotification7) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountNotification7) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountNotification7) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountNotification7) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountNotification7) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountNotification7) AddInterest() *AccountInterest2 {
	newValue := new(AccountInterest2)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountNotification7) AddTransactionsSummary() *TotalTransactions4 {
	a.TransactionsSummary = new(TotalTransactions4)
	return a.TransactionsSummary
}

func (a *AccountNotification7) AddEntry() *ReportEntry4 {
	newValue := new(ReportEntry4)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountNotification7) SetAdditionalNotificationInformation(value string) {
	a.AdditionalNotificationInformation = (*Max500Text)(&value)
}
