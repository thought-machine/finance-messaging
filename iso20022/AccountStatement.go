package iso20022

// Set of elements providing further details on the account statement.
type AccountStatement1 struct {

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

	// Range of time between the start date and the end date for which the account statement is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Specifies if this document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *CashAccount13 `xml:"Acct"`

	// Identifies the parent account of the reported account.
	RelatedAccount *CashAccount7 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest1 `xml:"Intrst,omitempty"`

	// Set of elements defining the balance(s).
	Balance []*CashBalance2 `xml:"Bal"`

	// Set of element providing summary information on entries.
	TransactionsSummary *TotalTransactions1 `xml:"TxsSummry,omitempty"`

	// Specifies the elements of an entry in the statement.
	//
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*StatementEntry1 `xml:"Ntry,omitempty"`

	// Further details on the account statement.
	AdditionalStatementInformation *Max500Text `xml:"AddtlStmtInf,omitempty"`
}

func (a *AccountStatement1) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountStatement1) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement1) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement1) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountStatement1) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountStatement1) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountStatement1) AddAccount() *CashAccount13 {
	a.Account = new(CashAccount13)
	return a.Account
}

func (a *AccountStatement1) AddRelatedAccount() *CashAccount7 {
	a.RelatedAccount = new(CashAccount7)
	return a.RelatedAccount
}

func (a *AccountStatement1) AddInterest() *AccountInterest1 {
	newValue := new(AccountInterest1)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountStatement1) AddBalance() *CashBalance2 {
	newValue := new(CashBalance2)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountStatement1) AddTransactionsSummary() *TotalTransactions1 {
	a.TransactionsSummary = new(TotalTransactions1)
	return a.TransactionsSummary
}

func (a *AccountStatement1) AddEntry() *StatementEntry1 {
	newValue := new(StatementEntry1)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountStatement1) SetAdditionalStatementInformation(value string) {
	a.AdditionalStatementInformation = (*Max500Text)(&value)
}

// Set of elements used to provide details of the account statement.
type AccountStatement2 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account statement.
	Identification *Max35Text `xml:"Id"`

	// Sequential number of the statement, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each statement sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the statement, as assigned by the account servicer. It is increased incrementally for each statement sent.
	//
	// Usage: Where a paper statement is a legal requirement, it may have a number different from the electronic sequential number. Paper statements could for instance only be sent if movement on the account has taken place, whereas electronic statements could be sent at the end of each reporting period, regardless of whether movements have taken place or not.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account statement is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount20 `xml:"Acct"`

	// Identifies the parent account of the account for which the statement has been issued.
	RelatedAccount *CashAccount16 `xml:"RltdAcct,omitempty"`

	// Set of elements used to provide general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest2 `xml:"Intrst,omitempty"`

	// Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balance []*CashBalance3 `xml:"Bal"`

	// Set of elements used to provide summary information on entries.
	TransactionsSummary *TotalTransactions2 `xml:"TxsSummry,omitempty"`

	// Set of elements used to specify an entry in the statement.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry2 `xml:"Ntry,omitempty"`

	// Further details of the account statement.
	AdditionalStatementInformation *Max500Text `xml:"AddtlStmtInf,omitempty"`
}

func (a *AccountStatement2) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountStatement2) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement2) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement2) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountStatement2) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountStatement2) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountStatement2) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountStatement2) AddAccount() *CashAccount20 {
	a.Account = new(CashAccount20)
	return a.Account
}

func (a *AccountStatement2) AddRelatedAccount() *CashAccount16 {
	a.RelatedAccount = new(CashAccount16)
	return a.RelatedAccount
}

func (a *AccountStatement2) AddInterest() *AccountInterest2 {
	newValue := new(AccountInterest2)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountStatement2) AddBalance() *CashBalance3 {
	newValue := new(CashBalance3)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountStatement2) AddTransactionsSummary() *TotalTransactions2 {
	a.TransactionsSummary = new(TotalTransactions2)
	return a.TransactionsSummary
}

func (a *AccountStatement2) AddEntry() *ReportEntry2 {
	newValue := new(ReportEntry2)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountStatement2) SetAdditionalStatementInformation(value string) {
	a.AdditionalStatementInformation = (*Max500Text)(&value)
}

// Provides further details of the account statement.
type AccountStatement3 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account statement.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the statement.
	//
	// Usage: The pagination of the statement is only allowed when agreed between the parties.
	StatementPagination *Pagination `xml:"StmtPgntn,omitempty"`

	// Sequential number of the statement, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each statement sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the statement, as assigned by the account servicer. It is increased incrementally for each statement sent.
	//
	// Usage: Where a paper statement is a legal requirement, it may have a number different from the electronic sequential number. Paper statements could for instance only be sent if movement on the account has taken place, whereas electronic statements could be sent at the end of each reporting period, regardless of whether movements have taken place or not.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account statement is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the statement has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest2 `xml:"Intrst,omitempty"`

	// Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balance []*CashBalance3 `xml:"Bal"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions2 `xml:"TxsSummry,omitempty"`

	// Set of elements used to specify an entry in the statement.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry3 `xml:"Ntry,omitempty"`

	// Further details of the account statement.
	AdditionalStatementInformation *Max500Text `xml:"AddtlStmtInf,omitempty"`
}

func (a *AccountStatement3) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountStatement3) AddStatementPagination() *Pagination {
	a.StatementPagination = new(Pagination)
	return a.StatementPagination
}

func (a *AccountStatement3) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement3) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement3) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountStatement3) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountStatement3) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountStatement3) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountStatement3) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountStatement3) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountStatement3) AddInterest() *AccountInterest2 {
	newValue := new(AccountInterest2)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountStatement3) AddBalance() *CashBalance3 {
	newValue := new(CashBalance3)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountStatement3) AddTransactionsSummary() *TotalTransactions2 {
	a.TransactionsSummary = new(TotalTransactions2)
	return a.TransactionsSummary
}

func (a *AccountStatement3) AddEntry() *ReportEntry3 {
	newValue := new(ReportEntry3)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountStatement3) SetAdditionalStatementInformation(value string) {
	a.AdditionalStatementInformation = (*Max500Text)(&value)
}

// Provides further details of the account statement.
type AccountStatement4 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account statement.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the statement.
	//
	// Usage: The pagination of the statement is only allowed when agreed between the parties.
	StatementPagination *Pagination `xml:"StmtPgntn,omitempty"`

	// Sequential number of the statement, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each statement sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the statement, as assigned by the account servicer. It is increased incrementally for each statement sent.
	//
	// Usage: Where a paper statement is a legal requirement, it may have a number different from the electronic sequential number. Paper statements could for instance only be sent if movement on the account has taken place, whereas electronic statements could be sent at the end of each reporting period, regardless of whether movements have taken place or not.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account statement is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the statement has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest2 `xml:"Intrst,omitempty"`

	// Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balance []*CashBalance3 `xml:"Bal"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions4 `xml:"TxsSummry,omitempty"`

	// Set of elements used to specify an entry in the statement.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry4 `xml:"Ntry,omitempty"`

	// Further details of the account statement.
	AdditionalStatementInformation *Max500Text `xml:"AddtlStmtInf,omitempty"`
}

func (a *AccountStatement4) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountStatement4) AddStatementPagination() *Pagination {
	a.StatementPagination = new(Pagination)
	return a.StatementPagination
}

func (a *AccountStatement4) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement4) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement4) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountStatement4) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountStatement4) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountStatement4) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountStatement4) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountStatement4) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountStatement4) AddInterest() *AccountInterest2 {
	newValue := new(AccountInterest2)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountStatement4) AddBalance() *CashBalance3 {
	newValue := new(CashBalance3)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountStatement4) AddTransactionsSummary() *TotalTransactions4 {
	a.TransactionsSummary = new(TotalTransactions4)
	return a.TransactionsSummary
}

func (a *AccountStatement4) AddEntry() *ReportEntry4 {
	newValue := new(ReportEntry4)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountStatement4) SetAdditionalStatementInformation(value string) {
	a.AdditionalStatementInformation = (*Max500Text)(&value)
}

// Provides further details of the account statement.
type AccountStatement5 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account statement.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the statement.
	//
	// Usage: The pagination of the statement is only allowed when agreed between the parties.
	StatementPagination *Pagination `xml:"StmtPgntn,omitempty"`

	// Sequential number of the statement, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each statement sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the statement, as assigned by the account servicer. It is increased incrementally for each statement sent.
	//
	// Usage: Where a paper statement is a legal requirement, it may have a number different from the electronic sequential number. Paper statements could for instance only be sent if movement on the account has taken place, whereas electronic statements could be sent at the end of each reporting period, regardless of whether movements have taken place or not.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account statement is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the statement has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest3 `xml:"Intrst,omitempty"`

	// Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balance []*CashBalance3 `xml:"Bal"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions4 `xml:"TxsSummry,omitempty"`

	// Specify an entry in the statement.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	//
	//
	// Usage Rule:  In case of a Payments R-transaction the creditor / debtor referenced of the original payment initiation messages is also used for reporting of the R-transaction. The original debtor/creditor in the reporting of R-Transactions is not inverted.
	// Following elements all defined in the TransactionDetails in RelatedParties or RelatedAgents are impacted by this usage rule:
	// Creditor, UltimateCreditor, CreditorAccount, CreditorAgent, Debtor, UltimateDebtor, DebtorAccount and DebtorAgent.
	//
	Entry []*ReportEntry7 `xml:"Ntry,omitempty"`

	// Further details of the account statement.
	AdditionalStatementInformation *Max500Text `xml:"AddtlStmtInf,omitempty"`
}

func (a *AccountStatement5) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountStatement5) AddStatementPagination() *Pagination {
	a.StatementPagination = new(Pagination)
	return a.StatementPagination
}

func (a *AccountStatement5) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement5) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement5) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountStatement5) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountStatement5) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountStatement5) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountStatement5) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountStatement5) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountStatement5) AddInterest() *AccountInterest3 {
	newValue := new(AccountInterest3)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountStatement5) AddBalance() *CashBalance3 {
	newValue := new(CashBalance3)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountStatement5) AddTransactionsSummary() *TotalTransactions4 {
	a.TransactionsSummary = new(TotalTransactions4)
	return a.TransactionsSummary
}

func (a *AccountStatement5) AddEntry() *ReportEntry7 {
	newValue := new(ReportEntry7)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountStatement5) SetAdditionalStatementInformation(value string) {
	a.AdditionalStatementInformation = (*Max500Text)(&value)
}

// Provides further details of the account statement.
type AccountStatement6 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account statement.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the statement.
	//
	// Usage: The pagination of the statement is only allowed when agreed between the parties.
	StatementPagination *Pagination `xml:"StmtPgntn,omitempty"`

	// Sequential number of the statement, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each statement sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the statement, as assigned by the account servicer. It is increased incrementally for each statement sent.
	//
	// Usage: Where a paper statement is a legal requirement, it may have a number different from the electronic sequential number. Paper statements could for instance only be sent if movement on the account has taken place, whereas electronic statements could be sent at the end of each reporting period, regardless of whether movements have taken place or not.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account statement is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the statement has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest3 `xml:"Intrst,omitempty"`

	// Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balance []*CashBalance7 `xml:"Bal"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions5 `xml:"TxsSummry,omitempty"`

	// Specify an entry in the statement.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	//
	//
	// Usage Rule:  In case of a Payments R-transaction the creditor / debtor referenced of the original payment initiation messages is also used for reporting of the R-transaction. The original debtor/creditor in the reporting of R-Transactions is not inverted.
	// Following elements all defined in the TransactionDetails in RelatedParties or RelatedAgents are impacted by this usage rule:
	// Creditor, UltimateCreditor, CreditorAccount, CreditorAgent, Debtor, UltimateDebtor, DebtorAccount and DebtorAgent.
	//
	Entry []*ReportEntry8 `xml:"Ntry,omitempty"`

	// Further details of the account statement.
	AdditionalStatementInformation *Max500Text `xml:"AddtlStmtInf,omitempty"`
}

func (a *AccountStatement6) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountStatement6) AddStatementPagination() *Pagination {
	a.StatementPagination = new(Pagination)
	return a.StatementPagination
}

func (a *AccountStatement6) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement6) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement6) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountStatement6) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountStatement6) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountStatement6) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountStatement6) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountStatement6) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountStatement6) AddInterest() *AccountInterest3 {
	newValue := new(AccountInterest3)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountStatement6) AddBalance() *CashBalance7 {
	newValue := new(CashBalance7)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountStatement6) AddTransactionsSummary() *TotalTransactions5 {
	a.TransactionsSummary = new(TotalTransactions5)
	return a.TransactionsSummary
}

func (a *AccountStatement6) AddEntry() *ReportEntry8 {
	newValue := new(ReportEntry8)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountStatement6) SetAdditionalStatementInformation(value string) {
	a.AdditionalStatementInformation = (*Max500Text)(&value)
}
