package iso20022

// Message used to report to a corporate on the actual set up up of the account, related services and mandates.
type AccountReport1 struct {

	// Characteristics of the account.
	Account *CustomerAccount1 `xml:"Acct"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Specifies target and actual dates.
	ContractDates *AccountContract3 `xml:"CtrctDts,omitempty"`

	// Information specifying the account mandate.
	Mandate []*OperationMandate1 `xml:"Mndt,omitempty"`

	// Unique and unambiguous identification of the account used as a reference for the opening of another account.
	ReferenceAccount *CashAccount16 `xml:"RefAcct,omitempty"`

	// Unique and unambiguous identification of the account where to transfer the balance.
	BalanceTransferAccount *AccountForAction1 `xml:"BalTrfAcct,omitempty"`

	// Identification of  the transfer account servicer.
	TransferAccountServicerIdentification *BranchAndFinancialInstitutionIdentification4 `xml:"TrfAcctSvcrId,omitempty"`
}

func (a *AccountReport1) AddAccount() *CustomerAccount1 {
	a.Account = new(CustomerAccount1)
	return a.Account
}

func (a *AccountReport1) AddUnderlyingMasterAgreement() *ContractDocument1 {
	a.UnderlyingMasterAgreement = new(ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountReport1) AddContractDates() *AccountContract3 {
	a.ContractDates = new(AccountContract3)
	return a.ContractDates
}

func (a *AccountReport1) AddMandate() *OperationMandate1 {
	newValue := new(OperationMandate1)
	a.Mandate = append(a.Mandate, newValue)
	return newValue
}

func (a *AccountReport1) AddReferenceAccount() *CashAccount16 {
	a.ReferenceAccount = new(CashAccount16)
	return a.ReferenceAccount
}

func (a *AccountReport1) AddBalanceTransferAccount() *AccountForAction1 {
	a.BalanceTransferAccount = new(AccountForAction1)
	return a.BalanceTransferAccount
}

func (a *AccountReport1) AddTransferAccountServicerIdentification() *BranchAndFinancialInstitutionIdentification4 {
	a.TransferAccountServicerIdentification = new(BranchAndFinancialInstitutionIdentification4)
	return a.TransferAccountServicerIdentification
}

// Set of elements used to provide details of the account report.
type AccountReport11 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	Identification *Max35Text `xml:"Id"`

	// Sequential number of the report, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each report sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the report, as assigned by the account servicer. It is increased incrementally for each report sent.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account report is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount20 `xml:"Acct"`

	// Identifies the parent account of the account for which the report has been issued.
	RelatedAccount *CashAccount16 `xml:"RltdAcct,omitempty"`

	// Set of elements used to provide general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest2 `xml:"Intrst,omitempty"`

	// Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balance []*CashBalance3 `xml:"Bal,omitempty"`

	// Set of elements used to provide summary information on entries.
	TransactionsSummary *TotalTransactions2 `xml:"TxsSummry,omitempty"`

	// Set of elements used to specify an entry in the report.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry2 `xml:"Ntry,omitempty"`

	// Further details of the account report.
	AdditionalReportInformation *Max500Text `xml:"AddtlRptInf,omitempty"`
}

func (a *AccountReport11) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountReport11) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountReport11) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountReport11) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountReport11) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountReport11) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountReport11) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountReport11) AddAccount() *CashAccount20 {
	a.Account = new(CashAccount20)
	return a.Account
}

func (a *AccountReport11) AddRelatedAccount() *CashAccount16 {
	a.RelatedAccount = new(CashAccount16)
	return a.RelatedAccount
}

func (a *AccountReport11) AddInterest() *AccountInterest2 {
	newValue := new(AccountInterest2)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountReport11) AddBalance() *CashBalance3 {
	newValue := new(CashBalance3)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountReport11) AddTransactionsSummary() *TotalTransactions2 {
	a.TransactionsSummary = new(TotalTransactions2)
	return a.TransactionsSummary
}

func (a *AccountReport11) AddEntry() *ReportEntry2 {
	newValue := new(ReportEntry2)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountReport11) SetAdditionalReportInformation(value string) {
	a.AdditionalReportInformation = (*Max500Text)(&value)
}

// Provides further details of the account report.
type AccountReport12 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the report.
	//
	// Usage: The pagination of the report is only allowed when agreed between the parties.
	ReportPagination *Pagination `xml:"RptPgntn,omitempty"`

	// Sequential number of the report, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each report sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the report, as assigned by the account servicer. It is increased incrementally for each report sent.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account report is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the report has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest3 `xml:"Intrst,omitempty"`

	// Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balance []*CashBalance3 `xml:"Bal,omitempty"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions3 `xml:"TxsSummry,omitempty"`

	// Set of elements used to specify an entry in the report.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry3 `xml:"Ntry,omitempty"`

	// Further details of the account report.
	AdditionalReportInformation *Max500Text `xml:"AddtlRptInf,omitempty"`
}

func (a *AccountReport12) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountReport12) AddReportPagination() *Pagination {
	a.ReportPagination = new(Pagination)
	return a.ReportPagination
}

func (a *AccountReport12) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountReport12) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountReport12) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountReport12) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountReport12) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountReport12) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountReport12) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountReport12) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountReport12) AddInterest() *AccountInterest3 {
	newValue := new(AccountInterest3)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountReport12) AddBalance() *CashBalance3 {
	newValue := new(CashBalance3)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountReport12) AddTransactionsSummary() *TotalTransactions3 {
	a.TransactionsSummary = new(TotalTransactions3)
	return a.TransactionsSummary
}

func (a *AccountReport12) AddEntry() *ReportEntry3 {
	newValue := new(ReportEntry3)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountReport12) SetAdditionalReportInformation(value string) {
	a.AdditionalReportInformation = (*Max500Text)(&value)
}

// Reports to a corporate on the actual set up up of the account, related services and mandates.
type AccountReport15 struct {

	// Characteristics of the account.
	Account *CustomerAccount5 `xml:"Acct"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Specifies target and actual dates.
	ContractDates *AccountContract3 `xml:"CtrctDts,omitempty"`

	// Information specifying the account mandate.
	Mandate []*OperationMandate2 `xml:"Mndt,omitempty"`

	// Definition of a group of parties.
	Group []*Group1 `xml:"Grp,omitempty"`

	// Unique and unambiguous identification of the account used as a reference for the opening of another account.
	ReferenceAccount *CashAccount24 `xml:"RefAcct,omitempty"`

	// Unique and unambiguous identification of the account where to transfer the balance.
	BalanceTransferAccount *AccountForAction1 `xml:"BalTrfAcct,omitempty"`

	// Identification of  the transfer account servicer.
	TransferAccountServicerIdentification *BranchAndFinancialInstitutionIdentification5 `xml:"TrfAcctSvcrId,omitempty"`
}

func (a *AccountReport15) AddAccount() *CustomerAccount5 {
	a.Account = new(CustomerAccount5)
	return a.Account
}

func (a *AccountReport15) AddUnderlyingMasterAgreement() *ContractDocument1 {
	a.UnderlyingMasterAgreement = new(ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountReport15) AddContractDates() *AccountContract3 {
	a.ContractDates = new(AccountContract3)
	return a.ContractDates
}

func (a *AccountReport15) AddMandate() *OperationMandate2 {
	newValue := new(OperationMandate2)
	a.Mandate = append(a.Mandate, newValue)
	return newValue
}

func (a *AccountReport15) AddGroup() *Group1 {
	newValue := new(Group1)
	a.Group = append(a.Group, newValue)
	return newValue
}

func (a *AccountReport15) AddReferenceAccount() *CashAccount24 {
	a.ReferenceAccount = new(CashAccount24)
	return a.ReferenceAccount
}

func (a *AccountReport15) AddBalanceTransferAccount() *AccountForAction1 {
	a.BalanceTransferAccount = new(AccountForAction1)
	return a.BalanceTransferAccount
}

func (a *AccountReport15) AddTransferAccountServicerIdentification() *BranchAndFinancialInstitutionIdentification5 {
	a.TransferAccountServicerIdentification = new(BranchAndFinancialInstitutionIdentification5)
	return a.TransferAccountServicerIdentification
}

// Provides further details of the account report.
type AccountReport16 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the report.
	//
	// Usage: The pagination of the report is only allowed when agreed between the parties.
	ReportPagination *Pagination `xml:"RptPgntn,omitempty"`

	// Sequential number of the report, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each report sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the report, as assigned by the account servicer. It is increased incrementally for each report sent.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account report is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the report has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest3 `xml:"Intrst,omitempty"`

	// Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balance []*CashBalance3 `xml:"Bal,omitempty"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions4 `xml:"TxsSummry,omitempty"`

	// Set of elements used to specify an entry in the report.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry4 `xml:"Ntry,omitempty"`

	// Further details of the account report.
	AdditionalReportInformation *Max500Text `xml:"AddtlRptInf,omitempty"`
}

func (a *AccountReport16) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountReport16) AddReportPagination() *Pagination {
	a.ReportPagination = new(Pagination)
	return a.ReportPagination
}

func (a *AccountReport16) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountReport16) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountReport16) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountReport16) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountReport16) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountReport16) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountReport16) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountReport16) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountReport16) AddInterest() *AccountInterest3 {
	newValue := new(AccountInterest3)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountReport16) AddBalance() *CashBalance3 {
	newValue := new(CashBalance3)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountReport16) AddTransactionsSummary() *TotalTransactions4 {
	a.TransactionsSummary = new(TotalTransactions4)
	return a.TransactionsSummary
}

func (a *AccountReport16) AddEntry() *ReportEntry4 {
	newValue := new(ReportEntry4)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountReport16) SetAdditionalReportInformation(value string) {
	a.AdditionalReportInformation = (*Max500Text)(&value)
}

// Provides further details of the account report.
type AccountReport18 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the report.
	//
	// Usage: The pagination of the report is only allowed when agreed between the parties.
	ReportPagination *Pagination `xml:"RptPgntn,omitempty"`

	// Sequential number of the report, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each report sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the report, as assigned by the account servicer. It is increased incrementally for each report sent.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account report is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the report has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest3 `xml:"Intrst,omitempty"`

	// Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balance []*CashBalance3 `xml:"Bal,omitempty"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions4 `xml:"TxsSummry,omitempty"`

	// Specifies an entry in the report.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	//
	// Usage Rule:  In case of a Payments R-transaction the creditor / debtor referenced of the original payment initiation messages is also used for reporting of the R-transaction. The original debtor/creditor in the reporting of R-Transactions is not inverted.
	// Following elements all defined in the TransactionDetails in RelatedParties or RelatedAgents are impacted by this usage rule:
	// Creditor, UltimateCreditor, CreditorAccount, CreditorAgent, Debtor, UltimateDebtor, DebtorAccount and DebtorAgent.
	//
	Entry []*ReportEntry7 `xml:"Ntry,omitempty"`

	// Further details of the account report.
	AdditionalReportInformation *Max500Text `xml:"AddtlRptInf,omitempty"`
}

func (a *AccountReport18) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountReport18) AddReportPagination() *Pagination {
	a.ReportPagination = new(Pagination)
	return a.ReportPagination
}

func (a *AccountReport18) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountReport18) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountReport18) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountReport18) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountReport18) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountReport18) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountReport18) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountReport18) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountReport18) AddInterest() *AccountInterest3 {
	newValue := new(AccountInterest3)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountReport18) AddBalance() *CashBalance3 {
	newValue := new(CashBalance3)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountReport18) AddTransactionsSummary() *TotalTransactions4 {
	a.TransactionsSummary = new(TotalTransactions4)
	return a.TransactionsSummary
}

func (a *AccountReport18) AddEntry() *ReportEntry7 {
	newValue := new(ReportEntry7)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountReport18) SetAdditionalReportInformation(value string) {
	a.AdditionalReportInformation = (*Max500Text)(&value)
}

// Provides further details of the account report.
type AccountReport19 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the report.
	//
	// Usage: The pagination of the report is only allowed when agreed between the parties.
	ReportPagination *Pagination `xml:"RptPgntn,omitempty"`

	// Sequential number of the report, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each report sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the report, as assigned by the account servicer. It is increased incrementally for each report sent.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account report is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the report has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest3 `xml:"Intrst,omitempty"`

	// Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balance []*CashBalance7 `xml:"Bal,omitempty"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions5 `xml:"TxsSummry,omitempty"`

	// Specifies an entry in the report.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	//
	// Usage Rule:  In case of a Payments R-transaction the creditor / debtor referenced of the original payment initiation messages is also used for reporting of the R-transaction. The original debtor/creditor in the reporting of R-Transactions is not inverted.
	// Following elements all defined in the TransactionDetails in RelatedParties or RelatedAgents are impacted by this usage rule:
	// Creditor, UltimateCreditor, CreditorAccount, CreditorAgent, Debtor, UltimateDebtor, DebtorAccount and DebtorAgent.
	//
	Entry []*ReportEntry8 `xml:"Ntry,omitempty"`

	// Further details of the account report.
	AdditionalReportInformation *Max500Text `xml:"AddtlRptInf,omitempty"`
}

func (a *AccountReport19) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountReport19) AddReportPagination() *Pagination {
	a.ReportPagination = new(Pagination)
	return a.ReportPagination
}

func (a *AccountReport19) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountReport19) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountReport19) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountReport19) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountReport19) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountReport19) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountReport19) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountReport19) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountReport19) AddInterest() *AccountInterest3 {
	newValue := new(AccountInterest3)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountReport19) AddBalance() *CashBalance7 {
	newValue := new(CashBalance7)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountReport19) AddTransactionsSummary() *TotalTransactions5 {
	a.TransactionsSummary = new(TotalTransactions5)
	return a.TransactionsSummary
}

func (a *AccountReport19) AddEntry() *ReportEntry8 {
	newValue := new(ReportEntry8)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountReport19) SetAdditionalReportInformation(value string) {
	a.AdditionalReportInformation = (*Max500Text)(&value)
}

// Set of elements providing further details on the account report.
type AccountReport9 struct {

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

	// Range of time between the start date and the end date for which the account report is issued.
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
	Balance []*CashBalance1 `xml:"Bal,omitempty"`

	// Set of element providing summary information on entries.
	TransactionsSummary *TotalTransactions1 `xml:"TxsSummry,omitempty"`

	// Specifies the elements of an entry in the report.
	//
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry1 `xml:"Ntry,omitempty"`

	// Further details on the account report.
	AdditionalReportInformation *Max500Text `xml:"AddtlRptInf,omitempty"`
}

func (a *AccountReport9) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountReport9) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountReport9) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountReport9) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountReport9) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountReport9) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountReport9) AddAccount() *CashAccount13 {
	a.Account = new(CashAccount13)
	return a.Account
}

func (a *AccountReport9) AddRelatedAccount() *CashAccount7 {
	a.RelatedAccount = new(CashAccount7)
	return a.RelatedAccount
}

func (a *AccountReport9) AddInterest() *AccountInterest1 {
	newValue := new(AccountInterest1)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountReport9) AddBalance() *CashBalance1 {
	newValue := new(CashBalance1)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountReport9) AddTransactionsSummary() *TotalTransactions1 {
	a.TransactionsSummary = new(TotalTransactions1)
	return a.TransactionsSummary
}

func (a *AccountReport9) AddEntry() *ReportEntry1 {
	newValue := new(ReportEntry1)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountReport9) SetAdditionalReportInformation(value string) {
	a.AdditionalReportInformation = (*Max500Text)(&value)
}
