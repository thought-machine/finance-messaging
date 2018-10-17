package iso20022

// Currency control registered contract amendment details.
type RegisteredContract1 struct {

	// Unique and unambiguous identification of the  contract registration amendment.
	ContractRegistrationAmendmentIdentification *Max35Text `xml:"CtrctRegnAmdmntId"`

	// Party registering the currency control contract.
	ReportingParty *TradeParty2 `xml:"RptgPty"`

	// Agent which registers the currency control contract.
	RegistrationAgent *BranchAndFinancialInstitutionIdentification5 `xml:"RegnAgt"`

	// Amendment details applied on one or several registered contracts.
	RegisteredContractAmendment []*RegisteredContract3 `xml:"RegdCtrctAmdmnt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RegisteredContract1) SetContractRegistrationAmendmentIdentification(value string) {
	r.ContractRegistrationAmendmentIdentification = (*Max35Text)(&value)
}

func (r *RegisteredContract1) AddReportingParty() *TradeParty2 {
	r.ReportingParty = new(TradeParty2)
	return r.ReportingParty
}

func (r *RegisteredContract1) AddRegistrationAgent() *BranchAndFinancialInstitutionIdentification5 {
	r.RegistrationAgent = new(BranchAndFinancialInstitutionIdentification5)
	return r.RegistrationAgent
}

func (r *RegisteredContract1) AddRegisteredContractAmendment() *RegisteredContract3 {
	newValue := new(RegisteredContract3)
	r.RegisteredContractAmendment = append(r.RegisteredContractAmendment, newValue)
	return newValue
}

func (r *RegisteredContract1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}

// Document that a user must file with an authorized servicer for each contract that involves foreign currency transactions with non residents.
type RegisteredContract2 struct {

	// Unique and unambiguous identification of the registered contract closure.
	RegisteredContractClosureIdentification *Max35Text `xml:"RegdCtrctClsrId"`

	// Party who registered the currency control contract.
	ReportingParty *TradeParty2 `xml:"RptgPty"`

	// Agent who registered the currency control contract.
	RegistrationAgent *BranchAndFinancialInstitutionIdentification5 `xml:"RegnAgt"`

	// Original registered contract identification.
	OriginalRegisteredContract *DocumentIdentification29 `xml:"OrgnlRegdCtrct"`

	// Priority of the registered contract closure.
	Priority *Priority2Code `xml:"Prty"`

	// Reason of the closure.
	ClosureReason *ContractClosureReason1Choice `xml:"ClsrRsn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RegisteredContract2) SetRegisteredContractClosureIdentification(value string) {
	r.RegisteredContractClosureIdentification = (*Max35Text)(&value)
}

func (r *RegisteredContract2) AddReportingParty() *TradeParty2 {
	r.ReportingParty = new(TradeParty2)
	return r.ReportingParty
}

func (r *RegisteredContract2) AddRegistrationAgent() *BranchAndFinancialInstitutionIdentification5 {
	r.RegistrationAgent = new(BranchAndFinancialInstitutionIdentification5)
	return r.RegistrationAgent
}

func (r *RegisteredContract2) AddOriginalRegisteredContract() *DocumentIdentification29 {
	r.OriginalRegisteredContract = new(DocumentIdentification29)
	return r.OriginalRegisteredContract
}

func (r *RegisteredContract2) SetPriority(value string) {
	r.Priority = (*Priority2Code)(&value)
}

func (r *RegisteredContract2) AddClosureReason() *ContractClosureReason1Choice {
	r.ClosureReason = new(ContractClosureReason1Choice)
	return r.ClosureReason
}

func (r *RegisteredContract2) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}

// Document that a user must file with an authorized servicer for each contract that involves foreign currency transactions with non residents.
type RegisteredContract3 struct {

	// Unique and unambiguous identification of the registered contract amendment request.
	RegisteredContractAmendmentIdentification *Max35Text `xml:"RegdCtrctAmdmntId"`

	// Identification of the original contract registration, as registered with the registration agent.
	OriginalRegisteredContractIdentification *Max35Text `xml:"OrgnlRegdCtrctId"`

	// Priority requested for the amendment of the registered contract.
	Priority *Priority2Code `xml:"Prty"`

	// Amendment details of the registered contract for the registered contract.
	Contract *UnderlyingContract1Choice `xml:"Ctrct"`

	// Contract balance on date of contract registration.
	ContractBalance []*ContractBalance1 `xml:"CtrctBal,omitempty"`

	// Type of the payment schedule provided in the contract.
	PaymentScheduleType *PaymentScheduleType1Choice `xml:"PmtSchdlTp,omitempty"`

	// Further additional information on the amendment.
	AdditionalInformation *Max1025Text `xml:"AddtlInf,omitempty"`

	// Documents provided as attachments to the contract registration amendment request.
	Attachment []*DocumentGeneralInformation3 `xml:"Attchmnt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RegisteredContract3) SetRegisteredContractAmendmentIdentification(value string) {
	r.RegisteredContractAmendmentIdentification = (*Max35Text)(&value)
}

func (r *RegisteredContract3) SetOriginalRegisteredContractIdentification(value string) {
	r.OriginalRegisteredContractIdentification = (*Max35Text)(&value)
}

func (r *RegisteredContract3) SetPriority(value string) {
	r.Priority = (*Priority2Code)(&value)
}

func (r *RegisteredContract3) AddContract() *UnderlyingContract1Choice {
	r.Contract = new(UnderlyingContract1Choice)
	return r.Contract
}

func (r *RegisteredContract3) AddContractBalance() *ContractBalance1 {
	newValue := new(ContractBalance1)
	r.ContractBalance = append(r.ContractBalance, newValue)
	return newValue
}

func (r *RegisteredContract3) AddPaymentScheduleType() *PaymentScheduleType1Choice {
	r.PaymentScheduleType = new(PaymentScheduleType1Choice)
	return r.PaymentScheduleType
}

func (r *RegisteredContract3) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max1025Text)(&value)
}

func (r *RegisteredContract3) AddAttachment() *DocumentGeneralInformation3 {
	newValue := new(DocumentGeneralInformation3)
	r.Attachment = append(r.Attachment, newValue)
	return newValue
}

func (r *RegisteredContract3) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}

// Document that a user must file with an authorized servicer for each contract that involves foreign currency transactions with non residents.
type RegisteredContract4 struct {

	// Unique identification of the contract registration creation, amendment or closure request.
	OriginalContractRegistrationRequest *Max35Text `xml:"OrgnlCtrctRegnReq,omitempty"`

	// Party registering the currency control contract.
	ReportingParty *TradeParty2 `xml:"RptgPty"`

	// Agent which registers the currency control contract.
	RegistrationAgent *BranchAndFinancialInstitutionIdentification5 `xml:"RegnAgt"`

	// Financial institution of the issuer of the contract.
	IssuerFinancialInstitution *BranchAndFinancialInstitutionIdentification5 `xml:"IssrFI"`

	// Details of the contract being registered.
	Contract *UnderlyingContract1Choice `xml:"Ctrct"`

	// Contract balance on date of contract registration.
	ContractBalance []*ContractBalance1 `xml:"CtrctBal,omitempty"`

	// Type of the payment schedule provided in the contract.
	PaymentScheduleType *PaymentScheduleType1Choice `xml:"PmtSchdlTp,omitempty"`

	// Unique and unambiguous identification of the registered contract as assigned by the registration agent.
	RegisteredContractIdentification *DocumentIdentification29 `xml:"RegdCtrctId"`

	// Identification of a previously defined registered contract linked to the same underlying contract.
	//
	// Usage:
	// This is the identification of a previous contract registration for the same underlying contract, which was closed. In most cases, this is used  when a reporting party moves from one registration agent to another.
	PreviousRegisteredContractIdentification *DocumentIdentification22 `xml:"PrvsRegdCtrctId,omitempty"`

	// Journal of previously closed registered contracts for the same underlying contract, which were requested at the same registration agent.
	RegisteredContractJournal []*RegisteredContractJournal1 `xml:"RegdCtrctJrnl,omitempty"`

	// Details on amendments to the registered contract.
	Amendment []*RegisteredContractAmendment1 `xml:"Amdmnt,omitempty"`

	// Provides the communication method for the submission of the registered contract.
	Submission *RegisteredContractCommunication1 `xml:"Submissn"`

	// Provides the communication method for the delivery of the registered contract.
	Delivery *RegisteredContractCommunication1 `xml:"Dlvry"`

	// Amount of money the borrower pays back to the lender outside of interests and charges.
	//
	// Usage:
	// Only applicable for loan contracts.
	LoanPrincipalAmount *ActiveCurrencyAndAmount `xml:"LnPrncplAmt,omitempty"`

	// Indicates whether the dates provided are estimates or not.
	EstimatedDateIndicator *TrueFalseIndicator `xml:"EstmtdDtInd"`

	// Indicates whether loan in which both the lender and the borrower are divisions of the same corporation or not.
	//
	// Usage:
	// Only applicable for loan contracts.
	InterCompanyLoan *TrueFalseIndicator `xml:"IntrCpnyLn"`

	// Further information on the registered contract.
	AdditionalInformation *Max1025Text `xml:"AddtlInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RegisteredContract4) SetOriginalContractRegistrationRequest(value string) {
	r.OriginalContractRegistrationRequest = (*Max35Text)(&value)
}

func (r *RegisteredContract4) AddReportingParty() *TradeParty2 {
	r.ReportingParty = new(TradeParty2)
	return r.ReportingParty
}

func (r *RegisteredContract4) AddRegistrationAgent() *BranchAndFinancialInstitutionIdentification5 {
	r.RegistrationAgent = new(BranchAndFinancialInstitutionIdentification5)
	return r.RegistrationAgent
}

func (r *RegisteredContract4) AddIssuerFinancialInstitution() *BranchAndFinancialInstitutionIdentification5 {
	r.IssuerFinancialInstitution = new(BranchAndFinancialInstitutionIdentification5)
	return r.IssuerFinancialInstitution
}

func (r *RegisteredContract4) AddContract() *UnderlyingContract1Choice {
	r.Contract = new(UnderlyingContract1Choice)
	return r.Contract
}

func (r *RegisteredContract4) AddContractBalance() *ContractBalance1 {
	newValue := new(ContractBalance1)
	r.ContractBalance = append(r.ContractBalance, newValue)
	return newValue
}

func (r *RegisteredContract4) AddPaymentScheduleType() *PaymentScheduleType1Choice {
	r.PaymentScheduleType = new(PaymentScheduleType1Choice)
	return r.PaymentScheduleType
}

func (r *RegisteredContract4) AddRegisteredContractIdentification() *DocumentIdentification29 {
	r.RegisteredContractIdentification = new(DocumentIdentification29)
	return r.RegisteredContractIdentification
}

func (r *RegisteredContract4) AddPreviousRegisteredContractIdentification() *DocumentIdentification22 {
	r.PreviousRegisteredContractIdentification = new(DocumentIdentification22)
	return r.PreviousRegisteredContractIdentification
}

func (r *RegisteredContract4) AddRegisteredContractJournal() *RegisteredContractJournal1 {
	newValue := new(RegisteredContractJournal1)
	r.RegisteredContractJournal = append(r.RegisteredContractJournal, newValue)
	return newValue
}

func (r *RegisteredContract4) AddAmendment() *RegisteredContractAmendment1 {
	newValue := new(RegisteredContractAmendment1)
	r.Amendment = append(r.Amendment, newValue)
	return newValue
}

func (r *RegisteredContract4) AddSubmission() *RegisteredContractCommunication1 {
	r.Submission = new(RegisteredContractCommunication1)
	return r.Submission
}

func (r *RegisteredContract4) AddDelivery() *RegisteredContractCommunication1 {
	r.Delivery = new(RegisteredContractCommunication1)
	return r.Delivery
}

func (r *RegisteredContract4) SetLoanPrincipalAmount(value, currency string) {
	r.LoanPrincipalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RegisteredContract4) SetEstimatedDateIndicator(value string) {
	r.EstimatedDateIndicator = (*TrueFalseIndicator)(&value)
}

func (r *RegisteredContract4) SetInterCompanyLoan(value string) {
	r.InterCompanyLoan = (*TrueFalseIndicator)(&value)
}

func (r *RegisteredContract4) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max1025Text)(&value)
}

func (r *RegisteredContract4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}

// Document that a user must file with an authorized servicer for each contract that involves foreign currency transactions with non residents.
type RegisteredContract5 struct {

	// Unique identification of the contract registration creation, amendment or closure request.
	OriginalContractRegistrationRequest *Max35Text `xml:"OrgnlCtrctRegnReq,omitempty"`

	// Financial institution of the issuer of the contract.
	IssuerFinancialInstitution *BranchAndFinancialInstitutionIdentification5 `xml:"IssrFI"`

	// Details of the contract being registered.
	Contract *UnderlyingContract1Choice `xml:"Ctrct"`

	// Contract balance on date of contract registration.
	ContractBalance []*ContractBalance1 `xml:"CtrctBal,omitempty"`

	// Type of the payment schedule provided in the contract.
	PaymentScheduleType *PaymentScheduleType1Choice `xml:"PmtSchdlTp,omitempty"`

	// Unique and unambiguous identification of the registered contract as assigned by the registration agent.
	RegisteredContractIdentification *DocumentIdentification29 `xml:"RegdCtrctId"`

	// Identification of a previously defined registered contract linked to the same underlying contract.
	//
	// Usage:
	// This is the identification of a previous contract registration for the same underlying contract, which was closed. In most cases, this is used  when a reporting party moves from one registration agent to another.
	PreviousRegisteredContractIdentification *DocumentIdentification22 `xml:"PrvsRegdCtrctId,omitempty"`

	// Journal of previously closed registered contracts for the same underlying contract, which were requested at the same registration agent.
	RegisteredContractJournal []*RegisteredContractJournal1 `xml:"RegdCtrctJrnl,omitempty"`

	// Details on amendments to the registered contract.
	Amendment []*RegisteredContractAmendment1 `xml:"Amdmnt,omitempty"`

	// Provides the communication method for the submission of the registered contract.
	Submission *RegisteredContractCommunication1 `xml:"Submissn"`

	// Provides the communication method for the delivery of the registered contract.
	Delivery *RegisteredContractCommunication1 `xml:"Dlvry"`

	// Amount of money the borrower pays back to the lender outside of interests and charges.
	//
	// Usage:
	// Only applicable for loan contracts.
	LoanPrincipalAmount *ActiveCurrencyAndAmount `xml:"LnPrncplAmt,omitempty"`

	// Indicates whether the dates provided are estimates or not.
	EstimatedDateIndicator *TrueFalseIndicator `xml:"EstmtdDtInd"`

	// Indicates whether loan in which both the lender and the borrower are divisions of the same corporation or not.
	//
	// Usage:
	// Only applicable for loan contracts.
	InterCompanyLoan *TrueFalseIndicator `xml:"IntrCpnyLn"`

	// Further information on the registered contract.
	AdditionalInformation *Max1025Text `xml:"AddtlInf,omitempty"`
}

func (r *RegisteredContract5) SetOriginalContractRegistrationRequest(value string) {
	r.OriginalContractRegistrationRequest = (*Max35Text)(&value)
}

func (r *RegisteredContract5) AddIssuerFinancialInstitution() *BranchAndFinancialInstitutionIdentification5 {
	r.IssuerFinancialInstitution = new(BranchAndFinancialInstitutionIdentification5)
	return r.IssuerFinancialInstitution
}

func (r *RegisteredContract5) AddContract() *UnderlyingContract1Choice {
	r.Contract = new(UnderlyingContract1Choice)
	return r.Contract
}

func (r *RegisteredContract5) AddContractBalance() *ContractBalance1 {
	newValue := new(ContractBalance1)
	r.ContractBalance = append(r.ContractBalance, newValue)
	return newValue
}

func (r *RegisteredContract5) AddPaymentScheduleType() *PaymentScheduleType1Choice {
	r.PaymentScheduleType = new(PaymentScheduleType1Choice)
	return r.PaymentScheduleType
}

func (r *RegisteredContract5) AddRegisteredContractIdentification() *DocumentIdentification29 {
	r.RegisteredContractIdentification = new(DocumentIdentification29)
	return r.RegisteredContractIdentification
}

func (r *RegisteredContract5) AddPreviousRegisteredContractIdentification() *DocumentIdentification22 {
	r.PreviousRegisteredContractIdentification = new(DocumentIdentification22)
	return r.PreviousRegisteredContractIdentification
}

func (r *RegisteredContract5) AddRegisteredContractJournal() *RegisteredContractJournal1 {
	newValue := new(RegisteredContractJournal1)
	r.RegisteredContractJournal = append(r.RegisteredContractJournal, newValue)
	return newValue
}

func (r *RegisteredContract5) AddAmendment() *RegisteredContractAmendment1 {
	newValue := new(RegisteredContractAmendment1)
	r.Amendment = append(r.Amendment, newValue)
	return newValue
}

func (r *RegisteredContract5) AddSubmission() *RegisteredContractCommunication1 {
	r.Submission = new(RegisteredContractCommunication1)
	return r.Submission
}

func (r *RegisteredContract5) AddDelivery() *RegisteredContractCommunication1 {
	r.Delivery = new(RegisteredContractCommunication1)
	return r.Delivery
}

func (r *RegisteredContract5) SetLoanPrincipalAmount(value, currency string) {
	r.LoanPrincipalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RegisteredContract5) SetEstimatedDateIndicator(value string) {
	r.EstimatedDateIndicator = (*TrueFalseIndicator)(&value)
}

func (r *RegisteredContract5) SetInterCompanyLoan(value string) {
	r.InterCompanyLoan = (*TrueFalseIndicator)(&value)
}

func (r *RegisteredContract5) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max1025Text)(&value)
}
