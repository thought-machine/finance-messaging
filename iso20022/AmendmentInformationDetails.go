package iso20022

// Amendment information details providing the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails1 struct {

	// Original mandate identification that has been modified.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId,omitempty"`

	// Original creditor scheme identification that has been modified.
	OriginalCreditorSchemeIdentification *PartyIdentification8 `xml:"OrgnlCdtrSchmeId,omitempty"`

	// Original creditor agent that has been modified.
	OriginalCreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"OrgnlCdtrAgt,omitempty"`

	// Original creditor agent acount that has been modified.
	OriginalCreditorAgentAccount *CashAccount7 `xml:"OrgnlCdtrAgtAcct,omitempty"`

	// Original debtor that has been modified.
	OriginalDebtor *PartyIdentification8 `xml:"OrgnlDbtr,omitempty"`

	// Original debtor account that has been modified.
	OriginalDebtorAccount *CashAccount7 `xml:"OrgnlDbtrAcct,omitempty"`

	// Original debtor's agent that has been modified.
	OriginalDebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"OrgnlDbtrAgt,omitempty"`

	// Original debtor agent account that has been modified.
	OriginalDebtorAgentAccount *CashAccount7 `xml:"OrgnlDbtrAgtAcct,omitempty"`

	// Original final collection date that has been modified.
	OriginalFinalCollectionDate *ISODate `xml:"OrgnlFnlColltnDt,omitempty"`

	// Original frequency that has been modified.
	OriginalFrequency *Frequency1Code `xml:"OrgnlFrqcy,omitempty"`
}

func (a *AmendmentInformationDetails1) SetOriginalMandateIdentification(value string) {
	a.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails1) AddOriginalCreditorSchemeIdentification() *PartyIdentification8 {
	a.OriginalCreditorSchemeIdentification = new(PartyIdentification8)
	return a.OriginalCreditorSchemeIdentification
}

func (a *AmendmentInformationDetails1) AddOriginalCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	a.OriginalCreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return a.OriginalCreditorAgent
}

func (a *AmendmentInformationDetails1) AddOriginalCreditorAgentAccount() *CashAccount7 {
	a.OriginalCreditorAgentAccount = new(CashAccount7)
	return a.OriginalCreditorAgentAccount
}

func (a *AmendmentInformationDetails1) AddOriginalDebtor() *PartyIdentification8 {
	a.OriginalDebtor = new(PartyIdentification8)
	return a.OriginalDebtor
}

func (a *AmendmentInformationDetails1) AddOriginalDebtorAccount() *CashAccount7 {
	a.OriginalDebtorAccount = new(CashAccount7)
	return a.OriginalDebtorAccount
}

func (a *AmendmentInformationDetails1) AddOriginalDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	a.OriginalDebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return a.OriginalDebtorAgent
}

func (a *AmendmentInformationDetails1) AddOriginalDebtorAgentAccount() *CashAccount7 {
	a.OriginalDebtorAgentAccount = new(CashAccount7)
	return a.OriginalDebtorAgentAccount
}

func (a *AmendmentInformationDetails1) SetOriginalFinalCollectionDate(value string) {
	a.OriginalFinalCollectionDate = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails1) SetOriginalFrequency(value string) {
	a.OriginalFrequency = (*Frequency1Code)(&value)
}

// Provides further details on the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails10 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId,omitempty"`

	// Original creditor scheme identification that has been modified.
	OriginalCreditorSchemeIdentification *PartyIdentification43 `xml:"OrgnlCdtrSchmeId,omitempty"`

	// Original creditor agent that has been modified.
	OriginalCreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlCdtrAgt,omitempty"`

	// Original creditor agent account that has been modified.
	OriginalCreditorAgentAccount *CashAccount24 `xml:"OrgnlCdtrAgtAcct,omitempty"`

	// Original debtor that has been modified.
	OriginalDebtor *PartyIdentification43 `xml:"OrgnlDbtr,omitempty"`

	// Original debtor account that has been modified.
	OriginalDebtorAccount *CashAccount24 `xml:"OrgnlDbtrAcct,omitempty"`

	// Original debtor agent that has been modified.
	OriginalDebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlDbtrAgt,omitempty"`

	// Original debtor agent account that has been modified.
	OriginalDebtorAgentAccount *CashAccount24 `xml:"OrgnlDbtrAgtAcct,omitempty"`

	// Original final collection date that has been modified.
	OriginalFinalCollectionDate *ISODate `xml:"OrgnlFnlColltnDt,omitempty"`

	// Original frequency that has been modified.
	OriginalFrequency *Frequency21Choice `xml:"OrgnlFrqcy,omitempty"`

	// Original reason for the mandate to allow the user to distinguish between different mandates for the same creditor.
	OriginalReason *MandateSetupReason1Choice `xml:"OrgnlRsn,omitempty"`
}

func (a *AmendmentInformationDetails10) SetOriginalMandateIdentification(value string) {
	a.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails10) AddOriginalCreditorSchemeIdentification() *PartyIdentification43 {
	a.OriginalCreditorSchemeIdentification = new(PartyIdentification43)
	return a.OriginalCreditorSchemeIdentification
}

func (a *AmendmentInformationDetails10) AddOriginalCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalCreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalCreditorAgent
}

func (a *AmendmentInformationDetails10) AddOriginalCreditorAgentAccount() *CashAccount24 {
	a.OriginalCreditorAgentAccount = new(CashAccount24)
	return a.OriginalCreditorAgentAccount
}

func (a *AmendmentInformationDetails10) AddOriginalDebtor() *PartyIdentification43 {
	a.OriginalDebtor = new(PartyIdentification43)
	return a.OriginalDebtor
}

func (a *AmendmentInformationDetails10) AddOriginalDebtorAccount() *CashAccount24 {
	a.OriginalDebtorAccount = new(CashAccount24)
	return a.OriginalDebtorAccount
}

func (a *AmendmentInformationDetails10) AddOriginalDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalDebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalDebtorAgent
}

func (a *AmendmentInformationDetails10) AddOriginalDebtorAgentAccount() *CashAccount24 {
	a.OriginalDebtorAgentAccount = new(CashAccount24)
	return a.OriginalDebtorAgentAccount
}

func (a *AmendmentInformationDetails10) SetOriginalFinalCollectionDate(value string) {
	a.OriginalFinalCollectionDate = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails10) AddOriginalFrequency() *Frequency21Choice {
	a.OriginalFrequency = new(Frequency21Choice)
	return a.OriginalFrequency
}

func (a *AmendmentInformationDetails10) AddOriginalReason() *MandateSetupReason1Choice {
	a.OriginalReason = new(MandateSetupReason1Choice)
	return a.OriginalReason
}

// Provides further details on the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails11 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId,omitempty"`

	// Original creditor scheme identification that has been modified.
	OriginalCreditorSchemeIdentification *PartyIdentification43 `xml:"OrgnlCdtrSchmeId,omitempty"`

	// Original creditor agent that has been modified.
	OriginalCreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlCdtrAgt,omitempty"`

	// Original creditor agent account that has been modified.
	OriginalCreditorAgentAccount *CashAccount24 `xml:"OrgnlCdtrAgtAcct,omitempty"`

	// Original debtor that has been modified.
	OriginalDebtor *PartyIdentification43 `xml:"OrgnlDbtr,omitempty"`

	// Original debtor account that has been modified.
	OriginalDebtorAccount *CashAccount24 `xml:"OrgnlDbtrAcct,omitempty"`

	// Original debtor agent that has been modified.
	OriginalDebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlDbtrAgt,omitempty"`

	// Original debtor agent account that has been modified.
	OriginalDebtorAgentAccount *CashAccount24 `xml:"OrgnlDbtrAgtAcct,omitempty"`

	// Original final collection date that has been modified.
	OriginalFinalCollectionDate *ISODate `xml:"OrgnlFnlColltnDt,omitempty"`

	// Original frequency that has been modified.
	OriginalFrequency *Frequency36Choice `xml:"OrgnlFrqcy,omitempty"`

	// Original reason for the mandate to allow the user to distinguish between different mandates for the same creditor.
	OriginalReason *MandateSetupReason1Choice `xml:"OrgnlRsn,omitempty"`

	// Original number of tracking days that has been modified.
	OriginalTrackingDays *Exact2NumericText `xml:"OrgnlTrckgDays,omitempty"`
}

func (a *AmendmentInformationDetails11) SetOriginalMandateIdentification(value string) {
	a.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails11) AddOriginalCreditorSchemeIdentification() *PartyIdentification43 {
	a.OriginalCreditorSchemeIdentification = new(PartyIdentification43)
	return a.OriginalCreditorSchemeIdentification
}

func (a *AmendmentInformationDetails11) AddOriginalCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalCreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalCreditorAgent
}

func (a *AmendmentInformationDetails11) AddOriginalCreditorAgentAccount() *CashAccount24 {
	a.OriginalCreditorAgentAccount = new(CashAccount24)
	return a.OriginalCreditorAgentAccount
}

func (a *AmendmentInformationDetails11) AddOriginalDebtor() *PartyIdentification43 {
	a.OriginalDebtor = new(PartyIdentification43)
	return a.OriginalDebtor
}

func (a *AmendmentInformationDetails11) AddOriginalDebtorAccount() *CashAccount24 {
	a.OriginalDebtorAccount = new(CashAccount24)
	return a.OriginalDebtorAccount
}

func (a *AmendmentInformationDetails11) AddOriginalDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalDebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalDebtorAgent
}

func (a *AmendmentInformationDetails11) AddOriginalDebtorAgentAccount() *CashAccount24 {
	a.OriginalDebtorAgentAccount = new(CashAccount24)
	return a.OriginalDebtorAgentAccount
}

func (a *AmendmentInformationDetails11) SetOriginalFinalCollectionDate(value string) {
	a.OriginalFinalCollectionDate = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails11) AddOriginalFrequency() *Frequency36Choice {
	a.OriginalFrequency = new(Frequency36Choice)
	return a.OriginalFrequency
}

func (a *AmendmentInformationDetails11) AddOriginalReason() *MandateSetupReason1Choice {
	a.OriginalReason = new(MandateSetupReason1Choice)
	return a.OriginalReason
}

func (a *AmendmentInformationDetails11) SetOriginalTrackingDays(value string) {
	a.OriginalTrackingDays = (*Exact2NumericText)(&value)
}

// Set of elements used to provide the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails6 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId,omitempty"`

	// Original creditor scheme identification that has been modified.
	OriginalCreditorSchemeIdentification *PartyIdentification32 `xml:"OrgnlCdtrSchmeId,omitempty"`

	// Original creditor agent that has been modified.
	OriginalCreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"OrgnlCdtrAgt,omitempty"`

	// Original creditor agent acount that has been modified.
	OriginalCreditorAgentAccount *CashAccount16 `xml:"OrgnlCdtrAgtAcct,omitempty"`

	// Original debtor that has been modified.
	OriginalDebtor *PartyIdentification32 `xml:"OrgnlDbtr,omitempty"`

	// Original debtor account that has been modified.
	OriginalDebtorAccount *CashAccount16 `xml:"OrgnlDbtrAcct,omitempty"`

	// Original debtor agent that has been modified.
	OriginalDebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"OrgnlDbtrAgt,omitempty"`

	// Original debtor agent account that has been modified.
	OriginalDebtorAgentAccount *CashAccount16 `xml:"OrgnlDbtrAgtAcct,omitempty"`

	// Original final collection date that has been modified.
	OriginalFinalCollectionDate *ISODate `xml:"OrgnlFnlColltnDt,omitempty"`

	// Original frequency that has been modified.
	OriginalFrequency *Frequency1Code `xml:"OrgnlFrqcy,omitempty"`
}

func (a *AmendmentInformationDetails6) SetOriginalMandateIdentification(value string) {
	a.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails6) AddOriginalCreditorSchemeIdentification() *PartyIdentification32 {
	a.OriginalCreditorSchemeIdentification = new(PartyIdentification32)
	return a.OriginalCreditorSchemeIdentification
}

func (a *AmendmentInformationDetails6) AddOriginalCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	a.OriginalCreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return a.OriginalCreditorAgent
}

func (a *AmendmentInformationDetails6) AddOriginalCreditorAgentAccount() *CashAccount16 {
	a.OriginalCreditorAgentAccount = new(CashAccount16)
	return a.OriginalCreditorAgentAccount
}

func (a *AmendmentInformationDetails6) AddOriginalDebtor() *PartyIdentification32 {
	a.OriginalDebtor = new(PartyIdentification32)
	return a.OriginalDebtor
}

func (a *AmendmentInformationDetails6) AddOriginalDebtorAccount() *CashAccount16 {
	a.OriginalDebtorAccount = new(CashAccount16)
	return a.OriginalDebtorAccount
}

func (a *AmendmentInformationDetails6) AddOriginalDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	a.OriginalDebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return a.OriginalDebtorAgent
}

func (a *AmendmentInformationDetails6) AddOriginalDebtorAgentAccount() *CashAccount16 {
	a.OriginalDebtorAgentAccount = new(CashAccount16)
	return a.OriginalDebtorAgentAccount
}

func (a *AmendmentInformationDetails6) SetOriginalFinalCollectionDate(value string) {
	a.OriginalFinalCollectionDate = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails6) SetOriginalFrequency(value string) {
	a.OriginalFrequency = (*Frequency1Code)(&value)
}

// Set of elements used to provide the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails7 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId,omitempty"`

	// Original creditor scheme identification that has been modified.
	OriginalCreditorSchemeIdentification *PartyIdentification43 `xml:"OrgnlCdtrSchmeId,omitempty"`

	// Original creditor agent that has been modified.
	OriginalCreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlCdtrAgt,omitempty"`

	// Original debtor that has been modified.
	OriginalDebtor *PartyIdentification43 `xml:"OrgnlDbtr,omitempty"`

	// Original debtor account that has been modified.
	OriginalDebtorAccount *CashAccount16 `xml:"OrgnlDbtrAcct,omitempty"`

	// Original debtor agent that has been modified.
	OriginalDebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlDbtrAgt,omitempty"`

	// Original final collection date that has been modified.
	OriginalFinalCollectionDate *ISODate `xml:"OrgnlFnlColltnDt,omitempty"`

	// Original frequency that has been modified.
	OriginalFrequency *Frequency1Code `xml:"OrgnlFrqcy,omitempty"`
}

func (a *AmendmentInformationDetails7) SetOriginalMandateIdentification(value string) {
	a.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails7) AddOriginalCreditorSchemeIdentification() *PartyIdentification43 {
	a.OriginalCreditorSchemeIdentification = new(PartyIdentification43)
	return a.OriginalCreditorSchemeIdentification
}

func (a *AmendmentInformationDetails7) AddOriginalCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalCreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalCreditorAgent
}

func (a *AmendmentInformationDetails7) AddOriginalDebtor() *PartyIdentification43 {
	a.OriginalDebtor = new(PartyIdentification43)
	return a.OriginalDebtor
}

func (a *AmendmentInformationDetails7) AddOriginalDebtorAccount() *CashAccount16 {
	a.OriginalDebtorAccount = new(CashAccount16)
	return a.OriginalDebtorAccount
}

func (a *AmendmentInformationDetails7) AddOriginalDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalDebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalDebtorAgent
}

func (a *AmendmentInformationDetails7) SetOriginalFinalCollectionDate(value string) {
	a.OriginalFinalCollectionDate = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails7) SetOriginalFrequency(value string) {
	a.OriginalFrequency = (*Frequency1Code)(&value)
}

// Provides further details on the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails8 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId,omitempty"`

	// Original creditor scheme identification that has been modified.
	OriginalCreditorSchemeIdentification *PartyIdentification43 `xml:"OrgnlCdtrSchmeId,omitempty"`

	// Original creditor agent that has been modified.
	OriginalCreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlCdtrAgt,omitempty"`

	// Original creditor agent account that has been modified.
	OriginalCreditorAgentAccount *CashAccount24 `xml:"OrgnlCdtrAgtAcct,omitempty"`

	// Original debtor that has been modified.
	OriginalDebtor *PartyIdentification43 `xml:"OrgnlDbtr,omitempty"`

	// Original debtor account that has been modified.
	OriginalDebtorAccount *CashAccount24 `xml:"OrgnlDbtrAcct,omitempty"`

	// Original debtor agent that has been modified.
	OriginalDebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlDbtrAgt,omitempty"`

	// Original debtor agent account that has been modified.
	OriginalDebtorAgentAccount *CashAccount24 `xml:"OrgnlDbtrAgtAcct,omitempty"`

	// Original final collection date that has been modified.
	OriginalFinalCollectionDate *ISODate `xml:"OrgnlFnlColltnDt,omitempty"`

	// Original frequency that has been modified.
	OriginalFrequency *Frequency6Code `xml:"OrgnlFrqcy,omitempty"`
}

func (a *AmendmentInformationDetails8) SetOriginalMandateIdentification(value string) {
	a.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails8) AddOriginalCreditorSchemeIdentification() *PartyIdentification43 {
	a.OriginalCreditorSchemeIdentification = new(PartyIdentification43)
	return a.OriginalCreditorSchemeIdentification
}

func (a *AmendmentInformationDetails8) AddOriginalCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalCreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalCreditorAgent
}

func (a *AmendmentInformationDetails8) AddOriginalCreditorAgentAccount() *CashAccount24 {
	a.OriginalCreditorAgentAccount = new(CashAccount24)
	return a.OriginalCreditorAgentAccount
}

func (a *AmendmentInformationDetails8) AddOriginalDebtor() *PartyIdentification43 {
	a.OriginalDebtor = new(PartyIdentification43)
	return a.OriginalDebtor
}

func (a *AmendmentInformationDetails8) AddOriginalDebtorAccount() *CashAccount24 {
	a.OriginalDebtorAccount = new(CashAccount24)
	return a.OriginalDebtorAccount
}

func (a *AmendmentInformationDetails8) AddOriginalDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalDebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalDebtorAgent
}

func (a *AmendmentInformationDetails8) AddOriginalDebtorAgentAccount() *CashAccount24 {
	a.OriginalDebtorAgentAccount = new(CashAccount24)
	return a.OriginalDebtorAgentAccount
}

func (a *AmendmentInformationDetails8) SetOriginalFinalCollectionDate(value string) {
	a.OriginalFinalCollectionDate = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails8) SetOriginalFrequency(value string) {
	a.OriginalFrequency = (*Frequency6Code)(&value)
}

// Provides further details on the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails9 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId,omitempty"`

	// Original creditor scheme identification that has been modified.
	OriginalCreditorSchemeIdentification *PartyIdentification43 `xml:"OrgnlCdtrSchmeId,omitempty"`

	// Original creditor agent that has been modified.
	OriginalCreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlCdtrAgt,omitempty"`

	// Original debtor that has been modified.
	OriginalDebtor *PartyIdentification43 `xml:"OrgnlDbtr,omitempty"`

	// Original debtor account that has been modified.
	OriginalDebtorAccount *CashAccount24 `xml:"OrgnlDbtrAcct,omitempty"`

	// Original debtor agent that has been modified.
	OriginalDebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlDbtrAgt,omitempty"`

	// Original final collection date that has been modified.
	OriginalFinalCollectionDate *ISODate `xml:"OrgnlFnlColltnDt,omitempty"`

	// Original frequency that has been modified.
	OriginalFrequency *Frequency6Code `xml:"OrgnlFrqcy,omitempty"`
}

func (a *AmendmentInformationDetails9) SetOriginalMandateIdentification(value string) {
	a.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails9) AddOriginalCreditorSchemeIdentification() *PartyIdentification43 {
	a.OriginalCreditorSchemeIdentification = new(PartyIdentification43)
	return a.OriginalCreditorSchemeIdentification
}

func (a *AmendmentInformationDetails9) AddOriginalCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalCreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalCreditorAgent
}

func (a *AmendmentInformationDetails9) AddOriginalDebtor() *PartyIdentification43 {
	a.OriginalDebtor = new(PartyIdentification43)
	return a.OriginalDebtor
}

func (a *AmendmentInformationDetails9) AddOriginalDebtorAccount() *CashAccount24 {
	a.OriginalDebtorAccount = new(CashAccount24)
	return a.OriginalDebtorAccount
}

func (a *AmendmentInformationDetails9) AddOriginalDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalDebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalDebtorAgent
}

func (a *AmendmentInformationDetails9) SetOriginalFinalCollectionDate(value string) {
	a.OriginalFinalCollectionDate = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails9) SetOriginalFrequency(value string) {
	a.OriginalFrequency = (*Frequency6Code)(&value)
}
