package iso20022

// Information that serves as a basis to debit an account.
type Mandate1 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId,omitempty"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation1 `xml:"Tp,omitempty"`

	// Provides details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences2 `xml:"Ocrncs,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveOrHistoricCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Provides information to identify the underlying documents associated with the mandate.
	ReferredDocument *ReferredDocumentInformation3 `xml:"RfrdDoc,omitempty"`
}

func (m *Mandate1) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *Mandate1) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *Mandate1) AddType() *MandateTypeInformation1 {
	m.Type = new(MandateTypeInformation1)
	return m.Type
}

func (m *Mandate1) AddOccurrences() *MandateOccurrences2 {
	m.Occurrences = new(MandateOccurrences2)
	return m.Occurrences
}

func (m *Mandate1) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (m *Mandate1) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (m *Mandate1) AddCreditorSchemeIdentification() *PartyIdentification43 {
	m.CreditorSchemeIdentification = new(PartyIdentification43)
	return m.CreditorSchemeIdentification
}

func (m *Mandate1) AddCreditor() *PartyIdentification43 {
	m.Creditor = new(PartyIdentification43)
	return m.Creditor
}

func (m *Mandate1) AddCreditorAccount() *CashAccount24 {
	m.CreditorAccount = new(CashAccount24)
	return m.CreditorAccount
}

func (m *Mandate1) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.CreditorAgent
}

func (m *Mandate1) AddUltimateCreditor() *PartyIdentification43 {
	m.UltimateCreditor = new(PartyIdentification43)
	return m.UltimateCreditor
}

func (m *Mandate1) AddDebtor() *PartyIdentification43 {
	m.Debtor = new(PartyIdentification43)
	return m.Debtor
}

func (m *Mandate1) AddDebtorAccount() *CashAccount24 {
	m.DebtorAccount = new(CashAccount24)
	return m.DebtorAccount
}

func (m *Mandate1) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.DebtorAgent
}

func (m *Mandate1) AddUltimateDebtor() *PartyIdentification43 {
	m.UltimateDebtor = new(PartyIdentification43)
	return m.UltimateDebtor
}

func (m *Mandate1) AddReferredDocument() *ReferredDocumentInformation3 {
	m.ReferredDocument = new(ReferredDocumentInformation3)
	return m.ReferredDocument
}

// Information that serves as a basis to debit an account.
type Mandate10 struct {

	// Unique identification, as assigned by the responsible party (such as the creditor) or agent (such as the debtor agent), to unambiguously identify the mandate.
	MandateIdentification []*Max35Text `xml:"MndtId,omitempty"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId"`

	// Specifies the transport authentication details related to the mandate.
	Authentication *MandateAuthentication1 `xml:"Authntcn,omitempty"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation2 `xml:"Tp,omitempty"`

	// Provides details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences4 `xml:"Ocrncs,omitempty"`

	// Specifies whether the direct debit instructions should be automatically re-submitted periodically when bilaterally agreed.
	TrackingIndicator *TrueFalseIndicator `xml:"TrckgInd"`

	// Amount different from the collection amount, as it includes the costs associated with the first debited amount.
	FirstCollectionAmount *ActiveCurrencyAndAmount `xml:"FrstColltnAmt,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Specifies the characteristics of the adjustment applied to the collection amount of a direct debit instruction.
	Adjustment *MandateAdjustment1 `xml:"Adjstmnt,omitempty"`

	// Provides the reason for the setup of the mandate.
	Reason *MandateSetupReason1Choice `xml:"Rsn,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Reference assigned by a creditor or ultimate creditor for internal usage for the mandate.
	MandateReference *Max35Text `xml:"MndtRef,omitempty"`

	// Provides information to identify the underlying documents associated with the mandate.
	ReferredDocument []*ReferredMandateDocument1 `xml:"RfrdDoc,omitempty"`

	// Additional information that cannot be captured in the structured elements within the message component.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *Mandate10) AddMandateIdentification(value string) {
	m.MandateIdentification = append(m.MandateIdentification, (*Max35Text)(&value))
}

func (m *Mandate10) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *Mandate10) AddAuthentication() *MandateAuthentication1 {
	m.Authentication = new(MandateAuthentication1)
	return m.Authentication
}

func (m *Mandate10) AddType() *MandateTypeInformation2 {
	m.Type = new(MandateTypeInformation2)
	return m.Type
}

func (m *Mandate10) AddOccurrences() *MandateOccurrences4 {
	m.Occurrences = new(MandateOccurrences4)
	return m.Occurrences
}

func (m *Mandate10) SetTrackingIndicator(value string) {
	m.TrackingIndicator = (*TrueFalseIndicator)(&value)
}

func (m *Mandate10) SetFirstCollectionAmount(value, currency string) {
	m.FirstCollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate10) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate10) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate10) AddAdjustment() *MandateAdjustment1 {
	m.Adjustment = new(MandateAdjustment1)
	return m.Adjustment
}

func (m *Mandate10) AddReason() *MandateSetupReason1Choice {
	m.Reason = new(MandateSetupReason1Choice)
	return m.Reason
}

func (m *Mandate10) AddCreditorSchemeIdentification() *PartyIdentification43 {
	m.CreditorSchemeIdentification = new(PartyIdentification43)
	return m.CreditorSchemeIdentification
}

func (m *Mandate10) AddCreditor() *PartyIdentification43 {
	m.Creditor = new(PartyIdentification43)
	return m.Creditor
}

func (m *Mandate10) AddCreditorAccount() *CashAccount24 {
	m.CreditorAccount = new(CashAccount24)
	return m.CreditorAccount
}

func (m *Mandate10) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.CreditorAgent
}

func (m *Mandate10) AddUltimateCreditor() *PartyIdentification43 {
	m.UltimateCreditor = new(PartyIdentification43)
	return m.UltimateCreditor
}

func (m *Mandate10) AddDebtor() *PartyIdentification43 {
	m.Debtor = new(PartyIdentification43)
	return m.Debtor
}

func (m *Mandate10) AddDebtorAccount() *CashAccount24 {
	m.DebtorAccount = new(CashAccount24)
	return m.DebtorAccount
}

func (m *Mandate10) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.DebtorAgent
}

func (m *Mandate10) AddUltimateDebtor() *PartyIdentification43 {
	m.UltimateDebtor = new(PartyIdentification43)
	return m.UltimateDebtor
}

func (m *Mandate10) SetMandateReference(value string) {
	m.MandateReference = (*Max35Text)(&value)
}

func (m *Mandate10) AddReferredDocument() *ReferredMandateDocument1 {
	newValue := new(ReferredMandateDocument1)
	m.ReferredDocument = append(m.ReferredDocument, newValue)
	return newValue
}

func (m *Mandate10) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

// Information that serves as a basis to debit an account.
type Mandate11 struct {

	// Unique identification, as assigned by the responsible party (such as the creditor) or agent (such as the debtor agent), to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId,omitempty"`

	// Specifies the transport authentication details related to the mandate.
	Authentication *MandateAuthentication1 `xml:"Authntcn,omitempty"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation2 `xml:"Tp,omitempty"`

	// Provides details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences4 `xml:"Ocrncs,omitempty"`

	// Specifies whether the direct debit instructions should be automatically re-submitted periodically when bilaterally agreed.
	TrackingIndicator *TrueFalseIndicator `xml:"TrckgInd"`

	// Amount different from the collection amount, as it includes the costs associated with the first debited amount.
	FirstCollectionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"FrstColltnAmt,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveOrHistoricCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Specifies the characteristics of the adjustment applied to the collection amount of a direct debit instruction.
	Adjustment *MandateAdjustment1 `xml:"Adjstmnt,omitempty"`

	// Provides the reason for the setup of the mandate.
	Reason *MandateSetupReason1Choice `xml:"Rsn,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Reference assigned by a creditor or ultimate creditor for internal usage for the mandate.
	MandateReference *Max35Text `xml:"MndtRef,omitempty"`

	// Provides information to identify the underlying documents associated with the mandate.
	ReferredDocument []*ReferredMandateDocument1 `xml:"RfrdDoc,omitempty"`
}

func (m *Mandate11) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *Mandate11) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *Mandate11) AddAuthentication() *MandateAuthentication1 {
	m.Authentication = new(MandateAuthentication1)
	return m.Authentication
}

func (m *Mandate11) AddType() *MandateTypeInformation2 {
	m.Type = new(MandateTypeInformation2)
	return m.Type
}

func (m *Mandate11) AddOccurrences() *MandateOccurrences4 {
	m.Occurrences = new(MandateOccurrences4)
	return m.Occurrences
}

func (m *Mandate11) SetTrackingIndicator(value string) {
	m.TrackingIndicator = (*TrueFalseIndicator)(&value)
}

func (m *Mandate11) SetFirstCollectionAmount(value, currency string) {
	m.FirstCollectionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (m *Mandate11) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (m *Mandate11) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (m *Mandate11) AddAdjustment() *MandateAdjustment1 {
	m.Adjustment = new(MandateAdjustment1)
	return m.Adjustment
}

func (m *Mandate11) AddReason() *MandateSetupReason1Choice {
	m.Reason = new(MandateSetupReason1Choice)
	return m.Reason
}

func (m *Mandate11) AddCreditorSchemeIdentification() *PartyIdentification43 {
	m.CreditorSchemeIdentification = new(PartyIdentification43)
	return m.CreditorSchemeIdentification
}

func (m *Mandate11) AddCreditor() *PartyIdentification43 {
	m.Creditor = new(PartyIdentification43)
	return m.Creditor
}

func (m *Mandate11) AddCreditorAccount() *CashAccount24 {
	m.CreditorAccount = new(CashAccount24)
	return m.CreditorAccount
}

func (m *Mandate11) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.CreditorAgent
}

func (m *Mandate11) AddUltimateCreditor() *PartyIdentification43 {
	m.UltimateCreditor = new(PartyIdentification43)
	return m.UltimateCreditor
}

func (m *Mandate11) AddDebtor() *PartyIdentification43 {
	m.Debtor = new(PartyIdentification43)
	return m.Debtor
}

func (m *Mandate11) AddDebtorAccount() *CashAccount24 {
	m.DebtorAccount = new(CashAccount24)
	return m.DebtorAccount
}

func (m *Mandate11) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.DebtorAgent
}

func (m *Mandate11) AddUltimateDebtor() *PartyIdentification43 {
	m.UltimateDebtor = new(PartyIdentification43)
	return m.UltimateDebtor
}

func (m *Mandate11) SetMandateReference(value string) {
	m.MandateReference = (*Max35Text)(&value)
}

func (m *Mandate11) AddReferredDocument() *ReferredMandateDocument1 {
	newValue := new(ReferredMandateDocument1)
	m.ReferredDocument = append(m.ReferredDocument, newValue)
	return newValue
}

// Information that serves as a basis to debit an account.
type Mandate2 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation1 `xml:"Tp,omitempty"`

	// Provides details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences2 `xml:"Ocrncs,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Provides information to identify the underlying documents associated with the mandate.
	ReferredDocument *ReferredDocumentInformation3 `xml:"RfrdDoc,omitempty"`
}

func (m *Mandate2) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *Mandate2) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *Mandate2) AddType() *MandateTypeInformation1 {
	m.Type = new(MandateTypeInformation1)
	return m.Type
}

func (m *Mandate2) AddOccurrences() *MandateOccurrences2 {
	m.Occurrences = new(MandateOccurrences2)
	return m.Occurrences
}

func (m *Mandate2) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate2) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate2) AddCreditorSchemeIdentification() *PartyIdentification43 {
	m.CreditorSchemeIdentification = new(PartyIdentification43)
	return m.CreditorSchemeIdentification
}

func (m *Mandate2) AddCreditor() *PartyIdentification43 {
	m.Creditor = new(PartyIdentification43)
	return m.Creditor
}

func (m *Mandate2) AddCreditorAccount() *CashAccount24 {
	m.CreditorAccount = new(CashAccount24)
	return m.CreditorAccount
}

func (m *Mandate2) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.CreditorAgent
}

func (m *Mandate2) AddUltimateCreditor() *PartyIdentification43 {
	m.UltimateCreditor = new(PartyIdentification43)
	return m.UltimateCreditor
}

func (m *Mandate2) AddDebtor() *PartyIdentification43 {
	m.Debtor = new(PartyIdentification43)
	return m.Debtor
}

func (m *Mandate2) AddDebtorAccount() *CashAccount24 {
	m.DebtorAccount = new(CashAccount24)
	return m.DebtorAccount
}

func (m *Mandate2) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.DebtorAgent
}

func (m *Mandate2) AddUltimateDebtor() *PartyIdentification43 {
	m.UltimateDebtor = new(PartyIdentification43)
	return m.UltimateDebtor
}

func (m *Mandate2) AddReferredDocument() *ReferredDocumentInformation3 {
	m.ReferredDocument = new(ReferredDocumentInformation3)
	return m.ReferredDocument
}

// Information that serves as a basis to debit an account.
type Mandate3 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId,omitempty"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation1 `xml:"Tp,omitempty"`

	// Provides details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences2 `xml:"Ocrncs,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Provides information to identify the underlying documents associated with the mandate.
	ReferredDocument *ReferredDocumentInformation3 `xml:"RfrdDoc,omitempty"`
}

func (m *Mandate3) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *Mandate3) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *Mandate3) AddType() *MandateTypeInformation1 {
	m.Type = new(MandateTypeInformation1)
	return m.Type
}

func (m *Mandate3) AddOccurrences() *MandateOccurrences2 {
	m.Occurrences = new(MandateOccurrences2)
	return m.Occurrences
}

func (m *Mandate3) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate3) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate3) AddCreditorSchemeIdentification() *PartyIdentification43 {
	m.CreditorSchemeIdentification = new(PartyIdentification43)
	return m.CreditorSchemeIdentification
}

func (m *Mandate3) AddCreditor() *PartyIdentification43 {
	m.Creditor = new(PartyIdentification43)
	return m.Creditor
}

func (m *Mandate3) AddCreditorAccount() *CashAccount24 {
	m.CreditorAccount = new(CashAccount24)
	return m.CreditorAccount
}

func (m *Mandate3) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.CreditorAgent
}

func (m *Mandate3) AddUltimateCreditor() *PartyIdentification43 {
	m.UltimateCreditor = new(PartyIdentification43)
	return m.UltimateCreditor
}

func (m *Mandate3) AddDebtor() *PartyIdentification43 {
	m.Debtor = new(PartyIdentification43)
	return m.Debtor
}

func (m *Mandate3) AddDebtorAccount() *CashAccount24 {
	m.DebtorAccount = new(CashAccount24)
	return m.DebtorAccount
}

func (m *Mandate3) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.DebtorAgent
}

func (m *Mandate3) AddUltimateDebtor() *PartyIdentification43 {
	m.UltimateDebtor = new(PartyIdentification43)
	return m.UltimateDebtor
}

func (m *Mandate3) AddReferredDocument() *ReferredDocumentInformation3 {
	m.ReferredDocument = new(ReferredDocumentInformation3)
	return m.ReferredDocument
}

// Information that serves as a basis to debit an account.
type Mandate4 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation1 `xml:"Tp,omitempty"`

	// Provides details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences2 `xml:"Ocrncs,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Provides information to identify the underlying documents associated with the mandate.
	ReferredDocument *ReferredDocumentInformation3 `xml:"RfrdDoc,omitempty"`

	// Additional information that cannot be captured in the structured elements within the message component.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *Mandate4) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *Mandate4) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *Mandate4) AddType() *MandateTypeInformation1 {
	m.Type = new(MandateTypeInformation1)
	return m.Type
}

func (m *Mandate4) AddOccurrences() *MandateOccurrences2 {
	m.Occurrences = new(MandateOccurrences2)
	return m.Occurrences
}

func (m *Mandate4) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate4) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate4) AddCreditorSchemeIdentification() *PartyIdentification43 {
	m.CreditorSchemeIdentification = new(PartyIdentification43)
	return m.CreditorSchemeIdentification
}

func (m *Mandate4) AddCreditor() *PartyIdentification43 {
	m.Creditor = new(PartyIdentification43)
	return m.Creditor
}

func (m *Mandate4) AddCreditorAccount() *CashAccount24 {
	m.CreditorAccount = new(CashAccount24)
	return m.CreditorAccount
}

func (m *Mandate4) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.CreditorAgent
}

func (m *Mandate4) AddUltimateCreditor() *PartyIdentification43 {
	m.UltimateCreditor = new(PartyIdentification43)
	return m.UltimateCreditor
}

func (m *Mandate4) AddDebtor() *PartyIdentification43 {
	m.Debtor = new(PartyIdentification43)
	return m.Debtor
}

func (m *Mandate4) AddDebtorAccount() *CashAccount24 {
	m.DebtorAccount = new(CashAccount24)
	return m.DebtorAccount
}

func (m *Mandate4) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.DebtorAgent
}

func (m *Mandate4) AddUltimateDebtor() *PartyIdentification43 {
	m.UltimateDebtor = new(PartyIdentification43)
	return m.UltimateDebtor
}

func (m *Mandate4) AddReferredDocument() *ReferredDocumentInformation3 {
	m.ReferredDocument = new(ReferredDocumentInformation3)
	return m.ReferredDocument
}

func (m *Mandate4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

// Information that serves as a basis to debit an account.
type Mandate5 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId,omitempty"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation1 `xml:"Tp,omitempty"`

	// Provides details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences3 `xml:"Ocrncs,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveOrHistoricCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Provides the reason for the setup of the mandate.
	Reason *MandateSetupReason1Choice `xml:"Rsn,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Provides information to identify the underlying documents associated with the mandate.
	ReferredDocument []*ReferredDocumentInformation6 `xml:"RfrdDoc,omitempty"`
}

func (m *Mandate5) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *Mandate5) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *Mandate5) AddType() *MandateTypeInformation1 {
	m.Type = new(MandateTypeInformation1)
	return m.Type
}

func (m *Mandate5) AddOccurrences() *MandateOccurrences3 {
	m.Occurrences = new(MandateOccurrences3)
	return m.Occurrences
}

func (m *Mandate5) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (m *Mandate5) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (m *Mandate5) AddReason() *MandateSetupReason1Choice {
	m.Reason = new(MandateSetupReason1Choice)
	return m.Reason
}

func (m *Mandate5) AddCreditorSchemeIdentification() *PartyIdentification43 {
	m.CreditorSchemeIdentification = new(PartyIdentification43)
	return m.CreditorSchemeIdentification
}

func (m *Mandate5) AddCreditor() *PartyIdentification43 {
	m.Creditor = new(PartyIdentification43)
	return m.Creditor
}

func (m *Mandate5) AddCreditorAccount() *CashAccount24 {
	m.CreditorAccount = new(CashAccount24)
	return m.CreditorAccount
}

func (m *Mandate5) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.CreditorAgent
}

func (m *Mandate5) AddUltimateCreditor() *PartyIdentification43 {
	m.UltimateCreditor = new(PartyIdentification43)
	return m.UltimateCreditor
}

func (m *Mandate5) AddDebtor() *PartyIdentification43 {
	m.Debtor = new(PartyIdentification43)
	return m.Debtor
}

func (m *Mandate5) AddDebtorAccount() *CashAccount24 {
	m.DebtorAccount = new(CashAccount24)
	return m.DebtorAccount
}

func (m *Mandate5) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.DebtorAgent
}

func (m *Mandate5) AddUltimateDebtor() *PartyIdentification43 {
	m.UltimateDebtor = new(PartyIdentification43)
	return m.UltimateDebtor
}

func (m *Mandate5) AddReferredDocument() *ReferredDocumentInformation6 {
	newValue := new(ReferredDocumentInformation6)
	m.ReferredDocument = append(m.ReferredDocument, newValue)
	return newValue
}

// Information that serves as a basis to debit an account.
type Mandate6 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId,omitempty"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation1 `xml:"Tp,omitempty"`

	// Provides details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences3 `xml:"Ocrncs,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Provides the reason for the setup of the mandate.
	Reason *MandateSetupReason1Choice `xml:"Rsn,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Provides information to identify the underlying documents associated with the mandate.
	ReferredDocument []*ReferredDocumentInformation6 `xml:"RfrdDoc,omitempty"`
}

func (m *Mandate6) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *Mandate6) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *Mandate6) AddType() *MandateTypeInformation1 {
	m.Type = new(MandateTypeInformation1)
	return m.Type
}

func (m *Mandate6) AddOccurrences() *MandateOccurrences3 {
	m.Occurrences = new(MandateOccurrences3)
	return m.Occurrences
}

func (m *Mandate6) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate6) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate6) AddReason() *MandateSetupReason1Choice {
	m.Reason = new(MandateSetupReason1Choice)
	return m.Reason
}

func (m *Mandate6) AddCreditorSchemeIdentification() *PartyIdentification43 {
	m.CreditorSchemeIdentification = new(PartyIdentification43)
	return m.CreditorSchemeIdentification
}

func (m *Mandate6) AddCreditor() *PartyIdentification43 {
	m.Creditor = new(PartyIdentification43)
	return m.Creditor
}

func (m *Mandate6) AddCreditorAccount() *CashAccount24 {
	m.CreditorAccount = new(CashAccount24)
	return m.CreditorAccount
}

func (m *Mandate6) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.CreditorAgent
}

func (m *Mandate6) AddUltimateCreditor() *PartyIdentification43 {
	m.UltimateCreditor = new(PartyIdentification43)
	return m.UltimateCreditor
}

func (m *Mandate6) AddDebtor() *PartyIdentification43 {
	m.Debtor = new(PartyIdentification43)
	return m.Debtor
}

func (m *Mandate6) AddDebtorAccount() *CashAccount24 {
	m.DebtorAccount = new(CashAccount24)
	return m.DebtorAccount
}

func (m *Mandate6) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.DebtorAgent
}

func (m *Mandate6) AddUltimateDebtor() *PartyIdentification43 {
	m.UltimateDebtor = new(PartyIdentification43)
	return m.UltimateDebtor
}

func (m *Mandate6) AddReferredDocument() *ReferredDocumentInformation6 {
	newValue := new(ReferredDocumentInformation6)
	m.ReferredDocument = append(m.ReferredDocument, newValue)
	return newValue
}

// Information that serves as a basis to debit an account.
type Mandate7 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification []*Max35Text `xml:"MndtId,omitempty"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation1 `xml:"Tp,omitempty"`

	// Provides details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences3 `xml:"Ocrncs,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Provides the reason for the setup of the mandate.
	Reason *MandateSetupReason1Choice `xml:"Rsn,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Provides information to identify the underlying documents associated with the mandate.
	ReferredDocument []*ReferredDocumentInformation6 `xml:"RfrdDoc,omitempty"`

	// Additional information that cannot be captured in the structured elements within the message component.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *Mandate7) AddMandateIdentification(value string) {
	m.MandateIdentification = append(m.MandateIdentification, (*Max35Text)(&value))
}

func (m *Mandate7) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *Mandate7) AddType() *MandateTypeInformation1 {
	m.Type = new(MandateTypeInformation1)
	return m.Type
}

func (m *Mandate7) AddOccurrences() *MandateOccurrences3 {
	m.Occurrences = new(MandateOccurrences3)
	return m.Occurrences
}

func (m *Mandate7) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate7) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate7) AddReason() *MandateSetupReason1Choice {
	m.Reason = new(MandateSetupReason1Choice)
	return m.Reason
}

func (m *Mandate7) AddCreditorSchemeIdentification() *PartyIdentification43 {
	m.CreditorSchemeIdentification = new(PartyIdentification43)
	return m.CreditorSchemeIdentification
}

func (m *Mandate7) AddCreditor() *PartyIdentification43 {
	m.Creditor = new(PartyIdentification43)
	return m.Creditor
}

func (m *Mandate7) AddCreditorAccount() *CashAccount24 {
	m.CreditorAccount = new(CashAccount24)
	return m.CreditorAccount
}

func (m *Mandate7) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.CreditorAgent
}

func (m *Mandate7) AddUltimateCreditor() *PartyIdentification43 {
	m.UltimateCreditor = new(PartyIdentification43)
	return m.UltimateCreditor
}

func (m *Mandate7) AddDebtor() *PartyIdentification43 {
	m.Debtor = new(PartyIdentification43)
	return m.Debtor
}

func (m *Mandate7) AddDebtorAccount() *CashAccount24 {
	m.DebtorAccount = new(CashAccount24)
	return m.DebtorAccount
}

func (m *Mandate7) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.DebtorAgent
}

func (m *Mandate7) AddUltimateDebtor() *PartyIdentification43 {
	m.UltimateDebtor = new(PartyIdentification43)
	return m.UltimateDebtor
}

func (m *Mandate7) AddReferredDocument() *ReferredDocumentInformation6 {
	newValue := new(ReferredDocumentInformation6)
	m.ReferredDocument = append(m.ReferredDocument, newValue)
	return newValue
}

func (m *Mandate7) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

// Information that serves as a basis to debit an account.
type Mandate8 struct {

	// Unique identification, as assigned by the responsible party (such as the creditor) or agent (such as the debtor agent), to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId,omitempty"`

	// Specifies the transport authentication details related to the mandate.
	Authentication *MandateAuthentication1 `xml:"Authntcn,omitempty"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation2 `xml:"Tp,omitempty"`

	// Provides details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences4 `xml:"Ocrncs,omitempty"`

	// Specifies whether the direct debit instructions should be automatically re-submitted periodically when bilaterally agreed.
	TrackingIndicator *TrueFalseIndicator `xml:"TrckgInd"`

	// Amount different from the collection amount, as it includes the costs associated with the first debited amount.
	FirstCollectionAmount *ActiveCurrencyAndAmount `xml:"FrstColltnAmt,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Specifies the characteristics of the adjustment applied to the collection amount of a direct debit instruction.
	Adjustment *MandateAdjustment1 `xml:"Adjstmnt,omitempty"`

	// Provides the reason for the setup of the mandate.
	Reason *MandateSetupReason1Choice `xml:"Rsn,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Reference assigned by a creditor or ultimate creditor for internal usage for the mandate.
	MandateReference *Max35Text `xml:"MndtRef,omitempty"`

	// Provides information to identify the underlying documents associated with the mandate.
	ReferredDocument []*ReferredMandateDocument1 `xml:"RfrdDoc,omitempty"`
}

func (m *Mandate8) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *Mandate8) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *Mandate8) AddAuthentication() *MandateAuthentication1 {
	m.Authentication = new(MandateAuthentication1)
	return m.Authentication
}

func (m *Mandate8) AddType() *MandateTypeInformation2 {
	m.Type = new(MandateTypeInformation2)
	return m.Type
}

func (m *Mandate8) AddOccurrences() *MandateOccurrences4 {
	m.Occurrences = new(MandateOccurrences4)
	return m.Occurrences
}

func (m *Mandate8) SetTrackingIndicator(value string) {
	m.TrackingIndicator = (*TrueFalseIndicator)(&value)
}

func (m *Mandate8) SetFirstCollectionAmount(value, currency string) {
	m.FirstCollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate8) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate8) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate8) AddAdjustment() *MandateAdjustment1 {
	m.Adjustment = new(MandateAdjustment1)
	return m.Adjustment
}

func (m *Mandate8) AddReason() *MandateSetupReason1Choice {
	m.Reason = new(MandateSetupReason1Choice)
	return m.Reason
}

func (m *Mandate8) AddCreditorSchemeIdentification() *PartyIdentification43 {
	m.CreditorSchemeIdentification = new(PartyIdentification43)
	return m.CreditorSchemeIdentification
}

func (m *Mandate8) AddCreditor() *PartyIdentification43 {
	m.Creditor = new(PartyIdentification43)
	return m.Creditor
}

func (m *Mandate8) AddCreditorAccount() *CashAccount24 {
	m.CreditorAccount = new(CashAccount24)
	return m.CreditorAccount
}

func (m *Mandate8) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.CreditorAgent
}

func (m *Mandate8) AddUltimateCreditor() *PartyIdentification43 {
	m.UltimateCreditor = new(PartyIdentification43)
	return m.UltimateCreditor
}

func (m *Mandate8) AddDebtor() *PartyIdentification43 {
	m.Debtor = new(PartyIdentification43)
	return m.Debtor
}

func (m *Mandate8) AddDebtorAccount() *CashAccount24 {
	m.DebtorAccount = new(CashAccount24)
	return m.DebtorAccount
}

func (m *Mandate8) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.DebtorAgent
}

func (m *Mandate8) AddUltimateDebtor() *PartyIdentification43 {
	m.UltimateDebtor = new(PartyIdentification43)
	return m.UltimateDebtor
}

func (m *Mandate8) SetMandateReference(value string) {
	m.MandateReference = (*Max35Text)(&value)
}

func (m *Mandate8) AddReferredDocument() *ReferredMandateDocument1 {
	newValue := new(ReferredMandateDocument1)
	m.ReferredDocument = append(m.ReferredDocument, newValue)
	return newValue
}

// Information that serves as a basis to debit an account.
type Mandate9 struct {

	// Unique identification, as assigned by the responsible party (such as the creditor) or agent (such as the debtor agent), to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId,omitempty"`

	// Specifies the transport authentication details related to the mandate.
	Authentication *MandateAuthentication1 `xml:"Authntcn,omitempty"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation2 `xml:"Tp,omitempty"`

	// Provides details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences4 `xml:"Ocrncs,omitempty"`

	// Specifies whether the direct debit instructions should be automatically re-submitted periodically when bilaterally agreed.
	TrackingIndicator *TrueFalseIndicator `xml:"TrckgInd"`

	// Amount different from the collection amount, as it includes the costs associated with the first debited amount.
	FirstCollectionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"FrstColltnAmt,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveOrHistoricCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Specifies the characteristics of the adjustment applied to the collection amount of a direct debit instruction.
	Adjustment *MandateAdjustment1 `xml:"Adjstmnt,omitempty"`

	// Provides the reason for the setup of the mandate.
	Reason *MandateSetupReason1Choice `xml:"Rsn,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Reference assigned by a creditor or ultimate creditor for internal usage for the mandate.
	MandateReference *Max35Text `xml:"MndtRef,omitempty"`

	// Provides information to identify the underlying documents associated with the mandate.
	ReferredDocument []*ReferredMandateDocument1 `xml:"RfrdDoc,omitempty"`
}

func (m *Mandate9) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *Mandate9) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *Mandate9) AddAuthentication() *MandateAuthentication1 {
	m.Authentication = new(MandateAuthentication1)
	return m.Authentication
}

func (m *Mandate9) AddType() *MandateTypeInformation2 {
	m.Type = new(MandateTypeInformation2)
	return m.Type
}

func (m *Mandate9) AddOccurrences() *MandateOccurrences4 {
	m.Occurrences = new(MandateOccurrences4)
	return m.Occurrences
}

func (m *Mandate9) SetTrackingIndicator(value string) {
	m.TrackingIndicator = (*TrueFalseIndicator)(&value)
}

func (m *Mandate9) SetFirstCollectionAmount(value, currency string) {
	m.FirstCollectionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (m *Mandate9) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (m *Mandate9) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (m *Mandate9) AddAdjustment() *MandateAdjustment1 {
	m.Adjustment = new(MandateAdjustment1)
	return m.Adjustment
}

func (m *Mandate9) AddReason() *MandateSetupReason1Choice {
	m.Reason = new(MandateSetupReason1Choice)
	return m.Reason
}

func (m *Mandate9) AddCreditorSchemeIdentification() *PartyIdentification43 {
	m.CreditorSchemeIdentification = new(PartyIdentification43)
	return m.CreditorSchemeIdentification
}

func (m *Mandate9) AddCreditor() *PartyIdentification43 {
	m.Creditor = new(PartyIdentification43)
	return m.Creditor
}

func (m *Mandate9) AddCreditorAccount() *CashAccount24 {
	m.CreditorAccount = new(CashAccount24)
	return m.CreditorAccount
}

func (m *Mandate9) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.CreditorAgent
}

func (m *Mandate9) AddUltimateCreditor() *PartyIdentification43 {
	m.UltimateCreditor = new(PartyIdentification43)
	return m.UltimateCreditor
}

func (m *Mandate9) AddDebtor() *PartyIdentification43 {
	m.Debtor = new(PartyIdentification43)
	return m.Debtor
}

func (m *Mandate9) AddDebtorAccount() *CashAccount24 {
	m.DebtorAccount = new(CashAccount24)
	return m.DebtorAccount
}

func (m *Mandate9) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.DebtorAgent
}

func (m *Mandate9) AddUltimateDebtor() *PartyIdentification43 {
	m.UltimateDebtor = new(PartyIdentification43)
	return m.UltimateDebtor
}

func (m *Mandate9) SetMandateReference(value string) {
	m.MandateReference = (*Max35Text)(&value)
}

func (m *Mandate9) AddReferredDocument() *ReferredMandateDocument1 {
	newValue := new(ReferredMandateDocument1)
	m.ReferredDocument = append(m.ReferredDocument, newValue)
	return newValue
}
