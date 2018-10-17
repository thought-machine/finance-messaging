package iso20022

// Instruction, initiated by the creditor, to debit a debtor's account in favour of the creditor. A direct debit can be pre-authorised or not. In most countries, authorisation is in the form of a mandate between the debtor and creditor.
type DirectDebitMandate2 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	DebtorAccountIdentification *CashAccountIdentification1Choice `xml:"DbtrAcctId"`

	// Party that owes the cash to the creditor/final party. The debtor is also the debit account owner.
	DebtorIdentification *PartyIdentification2Choice `xml:"DbtrId,omitempty"`

	// Party that receives an amount of money from the debtor. In the context of the payment model, the creditor is also the credit account owner.
	CreditorIdentification *PartyIdentification2Choice `xml:"CdtrId,omitempty"`

	// Financial institution that receives the direct debit instruction from the creditor or other authorised party.
	FirstAgent *FinancialInstitutionIdentification3Choice `xml:"FrstAgt"`

	// Financial institution that receives the payment transaction on behalf of the creditor, or other nominated party, and credits the account.
	FinalAgent *FinancialInstitutionIdentification3Choice `xml:"FnlAgt,omitempty"`

	// Reference assigned to a creditor by its financial institution, or relevant authority, authorising the creditor to take part in a direct debit scheme.
	RegistrationIdentification *Max35Text `xml:"RegnId,omitempty"`

	// Reference of the direct debit mandate that has been agreed upon by the debtor and creditor.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`
}

func (d *DirectDebitMandate2) AddDebtorAccountIdentification() *CashAccountIdentification1Choice {
	d.DebtorAccountIdentification = new(CashAccountIdentification1Choice)
	return d.DebtorAccountIdentification
}

func (d *DirectDebitMandate2) AddDebtorIdentification() *PartyIdentification2Choice {
	d.DebtorIdentification = new(PartyIdentification2Choice)
	return d.DebtorIdentification
}

func (d *DirectDebitMandate2) AddCreditorIdentification() *PartyIdentification2Choice {
	d.CreditorIdentification = new(PartyIdentification2Choice)
	return d.CreditorIdentification
}

func (d *DirectDebitMandate2) AddFirstAgent() *FinancialInstitutionIdentification3Choice {
	d.FirstAgent = new(FinancialInstitutionIdentification3Choice)
	return d.FirstAgent
}

func (d *DirectDebitMandate2) AddFinalAgent() *FinancialInstitutionIdentification3Choice {
	d.FinalAgent = new(FinancialInstitutionIdentification3Choice)
	return d.FinalAgent
}

func (d *DirectDebitMandate2) SetRegistrationIdentification(value string) {
	d.RegistrationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitMandate2) SetMandateIdentification(value string) {
	d.MandateIdentification = (*Max35Text)(&value)
}

// Instruction, initiated by the creditor, to debit a debtor's account in favour of the creditor. A direct debit can be pre-authorised or not. In most countries, authorisation is in the form of a mandate between the debtor and creditor.
type DirectDebitMandate4 struct {

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *AccountIdentificationAndName3 `xml:"DbtrAcct"`

	// Party that owes the cash to the creditor/final party. The debtor is also the debit account owner.
	Debtor *PartyIdentification2Choice `xml:"Dbtr,omitempty"`

	// Number assigned by a tax authority to an entity.
	DebtorTaxIdentificationNumber *Max35Text `xml:"DbtrTaxIdNb,omitempty"`

	// Number assigned by a national registration authority to an entity.
	DebtorNationalRegistrationNumber *Max35Text `xml:"DbtrNtlRegnNb,omitempty"`

	// Party that receives an amount of money from the debtor. In the context of the payment model, the creditor is also the credit account owner.
	Creditor *PartyIdentification2Choice `xml:"Cdtr,omitempty"`

	// Financial institution that receives the direct debit instruction from the creditor or other authorised party.
	DebtorAgent *FinancialInstitutionIdentification3Choice `xml:"DbtrAgt"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	DebtorAgentBranch *BranchData `xml:"DbtrAgtBrnch,omitempty"`

	// Financial institution that receives the payment transaction on behalf of the creditor, or other nominated party, and credits the account.
	CreditorAgent *FinancialInstitutionIdentification3Choice `xml:"CdtrAgt,omitempty"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	CreditorAgentBranch *BranchData `xml:"CdtrAgtBrnch,omitempty"`

	// Reference assigned to a creditor by its financial institution, or relevant authority, authorising the creditor to take part in a direct debit scheme.
	RegistrationIdentification *Max35Text `xml:"RegnId,omitempty"`

	// Reference of the direct debit mandate that has been agreed upon by the debtor and creditor.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`
}

func (d *DirectDebitMandate4) AddDebtorAccount() *AccountIdentificationAndName3 {
	d.DebtorAccount = new(AccountIdentificationAndName3)
	return d.DebtorAccount
}

func (d *DirectDebitMandate4) AddDebtor() *PartyIdentification2Choice {
	d.Debtor = new(PartyIdentification2Choice)
	return d.Debtor
}

func (d *DirectDebitMandate4) SetDebtorTaxIdentificationNumber(value string) {
	d.DebtorTaxIdentificationNumber = (*Max35Text)(&value)
}

func (d *DirectDebitMandate4) SetDebtorNationalRegistrationNumber(value string) {
	d.DebtorNationalRegistrationNumber = (*Max35Text)(&value)
}

func (d *DirectDebitMandate4) AddCreditor() *PartyIdentification2Choice {
	d.Creditor = new(PartyIdentification2Choice)
	return d.Creditor
}

func (d *DirectDebitMandate4) AddDebtorAgent() *FinancialInstitutionIdentification3Choice {
	d.DebtorAgent = new(FinancialInstitutionIdentification3Choice)
	return d.DebtorAgent
}

func (d *DirectDebitMandate4) AddDebtorAgentBranch() *BranchData {
	d.DebtorAgentBranch = new(BranchData)
	return d.DebtorAgentBranch
}

func (d *DirectDebitMandate4) AddCreditorAgent() *FinancialInstitutionIdentification3Choice {
	d.CreditorAgent = new(FinancialInstitutionIdentification3Choice)
	return d.CreditorAgent
}

func (d *DirectDebitMandate4) AddCreditorAgentBranch() *BranchData {
	d.CreditorAgentBranch = new(BranchData)
	return d.CreditorAgentBranch
}

func (d *DirectDebitMandate4) SetRegistrationIdentification(value string) {
	d.RegistrationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitMandate4) SetMandateIdentification(value string) {
	d.MandateIdentification = (*Max35Text)(&value)
}

// Direct debit mandate parties and accounts.
type DirectDebitMandate5 struct {

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *AccountIdentificationAndName5 `xml:"DbtrAcct"`

	// Party that owes the cash to the creditor/final party. The debtor is also the debit account owner.
	Debtor *PartyIdentification70Choice `xml:"Dbtr,omitempty"`

	// Number assigned by a tax authority to an entity.
	DebtorTaxIdentificationNumber *Max35Text `xml:"DbtrTaxIdNb,omitempty"`

	// Number assigned by a national registration authority to an entity.
	DebtorNationalRegistrationNumber *Max35Text `xml:"DbtrNtlRegnNb,omitempty"`

	// Party that receives an amount of money from the debtor. In the context of the payment model, the creditor is also the credit account owner.
	Creditor *PartyIdentification70Choice `xml:"Cdtr,omitempty"`

	// Financial institution that receives the direct debit instruction from the creditor or other authorised party.
	DebtorAgent *FinancialInstitutionIdentification7Choice `xml:"DbtrAgt"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	DebtorAgentBranch *BranchData `xml:"DbtrAgtBrnch,omitempty"`

	// Financial institution that receives the payment transaction on behalf of the creditor, or other nominated party, and credits the account.
	CreditorAgent *FinancialInstitutionIdentification7Choice `xml:"CdtrAgt,omitempty"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	CreditorAgentBranch *BranchData `xml:"CdtrAgtBrnch,omitempty"`

	// Reference assigned to a creditor by its financial institution, or relevant authority, authorising the creditor to take part in a direct debit scheme.
	RegistrationIdentification *Max35Text `xml:"RegnId,omitempty"`

	// Reference of the direct debit mandate that has been agreed upon by the debtor and creditor.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`
}

func (d *DirectDebitMandate5) AddDebtorAccount() *AccountIdentificationAndName5 {
	d.DebtorAccount = new(AccountIdentificationAndName5)
	return d.DebtorAccount
}

func (d *DirectDebitMandate5) AddDebtor() *PartyIdentification70Choice {
	d.Debtor = new(PartyIdentification70Choice)
	return d.Debtor
}

func (d *DirectDebitMandate5) SetDebtorTaxIdentificationNumber(value string) {
	d.DebtorTaxIdentificationNumber = (*Max35Text)(&value)
}

func (d *DirectDebitMandate5) SetDebtorNationalRegistrationNumber(value string) {
	d.DebtorNationalRegistrationNumber = (*Max35Text)(&value)
}

func (d *DirectDebitMandate5) AddCreditor() *PartyIdentification70Choice {
	d.Creditor = new(PartyIdentification70Choice)
	return d.Creditor
}

func (d *DirectDebitMandate5) AddDebtorAgent() *FinancialInstitutionIdentification7Choice {
	d.DebtorAgent = new(FinancialInstitutionIdentification7Choice)
	return d.DebtorAgent
}

func (d *DirectDebitMandate5) AddDebtorAgentBranch() *BranchData {
	d.DebtorAgentBranch = new(BranchData)
	return d.DebtorAgentBranch
}

func (d *DirectDebitMandate5) AddCreditorAgent() *FinancialInstitutionIdentification7Choice {
	d.CreditorAgent = new(FinancialInstitutionIdentification7Choice)
	return d.CreditorAgent
}

func (d *DirectDebitMandate5) AddCreditorAgentBranch() *BranchData {
	d.CreditorAgentBranch = new(BranchData)
	return d.CreditorAgentBranch
}

func (d *DirectDebitMandate5) SetRegistrationIdentification(value string) {
	d.RegistrationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitMandate5) SetMandateIdentification(value string) {
	d.MandateIdentification = (*Max35Text)(&value)
}

// Instruction, initiated by the creditor, to debit a debtor's account in favour of the creditor. A direct debit can be pre-authorised or not. In most countries, authorisation is in the form of a mandate between the debtor and creditor.
type DirectDebitMandate6 struct {

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *AccountIdentificationAndName5 `xml:"DbtrAcct"`

	// Party that owes the cash to the creditor/final party. The debtor is also the debit account owner.
	Debtor *PartyIdentification113 `xml:"Dbtr,omitempty"`

	// Number assigned by a tax authority to an entity.
	DebtorTaxIdentificationNumber *Max35Text `xml:"DbtrTaxIdNb,omitempty"`

	// Number assigned by a national registration authority to an entity.
	DebtorNationalRegistrationNumber *Max35Text `xml:"DbtrNtlRegnNb,omitempty"`

	// Party that receives an amount of money from the debtor. In the context of the payment model, the creditor is also the credit account owner.
	Creditor *PartyIdentification113 `xml:"Cdtr,omitempty"`

	// Financial institution that receives the direct debit instruction from the creditor or other authorised party.
	DebtorAgent *FinancialInstitutionIdentification10 `xml:"DbtrAgt"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	DebtorAgentBranch *BranchData `xml:"DbtrAgtBrnch,omitempty"`

	// Financial institution that receives the payment transaction on behalf of the creditor, or other nominated party, and credits the account.
	CreditorAgent *FinancialInstitutionIdentification10 `xml:"CdtrAgt,omitempty"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	CreditorAgentBranch *BranchData `xml:"CdtrAgtBrnch,omitempty"`

	// Reference assigned to a creditor by its financial institution, or relevant authority, authorising the creditor to take part in a direct debit scheme.
	RegistrationIdentification *Max35Text `xml:"RegnId,omitempty"`

	// Reference of the direct debit mandate that has been agreed upon by the debtor and creditor.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`
}

func (d *DirectDebitMandate6) AddDebtorAccount() *AccountIdentificationAndName5 {
	d.DebtorAccount = new(AccountIdentificationAndName5)
	return d.DebtorAccount
}

func (d *DirectDebitMandate6) AddDebtor() *PartyIdentification113 {
	d.Debtor = new(PartyIdentification113)
	return d.Debtor
}

func (d *DirectDebitMandate6) SetDebtorTaxIdentificationNumber(value string) {
	d.DebtorTaxIdentificationNumber = (*Max35Text)(&value)
}

func (d *DirectDebitMandate6) SetDebtorNationalRegistrationNumber(value string) {
	d.DebtorNationalRegistrationNumber = (*Max35Text)(&value)
}

func (d *DirectDebitMandate6) AddCreditor() *PartyIdentification113 {
	d.Creditor = new(PartyIdentification113)
	return d.Creditor
}

func (d *DirectDebitMandate6) AddDebtorAgent() *FinancialInstitutionIdentification10 {
	d.DebtorAgent = new(FinancialInstitutionIdentification10)
	return d.DebtorAgent
}

func (d *DirectDebitMandate6) AddDebtorAgentBranch() *BranchData {
	d.DebtorAgentBranch = new(BranchData)
	return d.DebtorAgentBranch
}

func (d *DirectDebitMandate6) AddCreditorAgent() *FinancialInstitutionIdentification10 {
	d.CreditorAgent = new(FinancialInstitutionIdentification10)
	return d.CreditorAgent
}

func (d *DirectDebitMandate6) AddCreditorAgentBranch() *BranchData {
	d.CreditorAgentBranch = new(BranchData)
	return d.CreditorAgentBranch
}

func (d *DirectDebitMandate6) SetRegistrationIdentification(value string) {
	d.RegistrationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitMandate6) SetMandateIdentification(value string) {
	d.MandateIdentification = (*Max35Text)(&value)
}
