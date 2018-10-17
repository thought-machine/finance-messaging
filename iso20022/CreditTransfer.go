package iso20022

// Payment instrument between a debtor and a creditor, which flows through one or more financial institutions or systems.
type CreditTransfer3 struct {

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	Reference *Max35Text `xml:"Ref,omitempty"`

	// Party that receives an amount of money from the debtor. The creditor is also the credit account owner.
	CreditorDetails *Creditor2 `xml:"CdtrDtls"`

	// Party that owes an amount of money to the creditor. The debtor is also the debit account owner.
	DebtorDetails *Debtor2 `xml:"DbtrDtls,omitempty"`
}

func (c *CreditTransfer3) SetReference(value string) {
	c.Reference = (*Max35Text)(&value)
}

func (c *CreditTransfer3) AddCreditorDetails() *Creditor2 {
	c.CreditorDetails = new(Creditor2)
	return c.CreditorDetails
}

func (c *CreditTransfer3) AddDebtorDetails() *Debtor2 {
	c.DebtorDetails = new(Debtor2)
	return c.DebtorDetails
}

// Payment instrument between a debtor and a creditor, which flows through one or more financial institutions or systems.
type CreditTransfer4 struct {

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	Reference *Max35Text `xml:"Ref,omitempty"`

	// Party that receives an amount of money from the debtor. The creditor is also the credit account owner.
	CreditorDetails *Creditor2 `xml:"CdtrDtls,omitempty"`

	// Party that owes the cash to the creditor/final party. The debtor is also the debit account owner.
	DebtorDetails *Debtor2 `xml:"DbtrDtls"`
}

func (c *CreditTransfer4) SetReference(value string) {
	c.Reference = (*Max35Text)(&value)
}

func (c *CreditTransfer4) AddCreditorDetails() *Creditor2 {
	c.CreditorDetails = new(Creditor2)
	return c.CreditorDetails
}

func (c *CreditTransfer4) AddDebtorDetails() *Debtor2 {
	c.DebtorDetails = new(Debtor2)
	return c.DebtorDetails
}

// Payment instrument between a debtor and a creditor, which flows through one or more financial institutions or systems.
type CreditTransfer6 struct {

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	Reference *Max35Text `xml:"Ref,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor. In the context of the payment model, the debtor is also the debit account owner.
	Debtor *PartyIdentification2Choice `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *AccountIdentificationAndName3 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *FinancialInstitutionIdentification3Choice `xml:"DbtrAgt,omitempty"`

	// Identifies the account of the debtor's agent.
	DebtorAgentAccount *AccountIdentificationAndName3 `xml:"DbtrAgtAcct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent1 *FinancialInstitutionIdentification3Choice `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *AccountIdentificationAndName3 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent2 *FinancialInstitutionIdentification3Choice `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *AccountIdentificationAndName3 `xml:"IntrmyAgt2Acct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *FinancialInstitutionIdentification3Choice `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *AccountIdentificationAndName3 `xml:"CdtrAgtAcct,omitempty"`

	// Party that receives an amount of money from the debtor. In the context of the payment model, the creditor is also the credit account owner.
	Creditor *PartyIdentification2Choice `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *AccountIdentificationAndName3 `xml:"CdtrAcct"`
}

func (c *CreditTransfer6) SetReference(value string) {
	c.Reference = (*Max35Text)(&value)
}

func (c *CreditTransfer6) AddDebtor() *PartyIdentification2Choice {
	c.Debtor = new(PartyIdentification2Choice)
	return c.Debtor
}

func (c *CreditTransfer6) AddDebtorAccount() *AccountIdentificationAndName3 {
	c.DebtorAccount = new(AccountIdentificationAndName3)
	return c.DebtorAccount
}

func (c *CreditTransfer6) AddDebtorAgent() *FinancialInstitutionIdentification3Choice {
	c.DebtorAgent = new(FinancialInstitutionIdentification3Choice)
	return c.DebtorAgent
}

func (c *CreditTransfer6) AddDebtorAgentAccount() *AccountIdentificationAndName3 {
	c.DebtorAgentAccount = new(AccountIdentificationAndName3)
	return c.DebtorAgentAccount
}

func (c *CreditTransfer6) AddIntermediaryAgent1() *FinancialInstitutionIdentification3Choice {
	c.IntermediaryAgent1 = new(FinancialInstitutionIdentification3Choice)
	return c.IntermediaryAgent1
}

func (c *CreditTransfer6) AddIntermediaryAgent1Account() *AccountIdentificationAndName3 {
	c.IntermediaryAgent1Account = new(AccountIdentificationAndName3)
	return c.IntermediaryAgent1Account
}

func (c *CreditTransfer6) AddIntermediaryAgent2() *FinancialInstitutionIdentification3Choice {
	c.IntermediaryAgent2 = new(FinancialInstitutionIdentification3Choice)
	return c.IntermediaryAgent2
}

func (c *CreditTransfer6) AddIntermediaryAgent2Account() *AccountIdentificationAndName3 {
	c.IntermediaryAgent2Account = new(AccountIdentificationAndName3)
	return c.IntermediaryAgent2Account
}

func (c *CreditTransfer6) AddCreditorAgent() *FinancialInstitutionIdentification3Choice {
	c.CreditorAgent = new(FinancialInstitutionIdentification3Choice)
	return c.CreditorAgent
}

func (c *CreditTransfer6) AddCreditorAgentAccount() *AccountIdentificationAndName3 {
	c.CreditorAgentAccount = new(AccountIdentificationAndName3)
	return c.CreditorAgentAccount
}

func (c *CreditTransfer6) AddCreditor() *PartyIdentification2Choice {
	c.Creditor = new(PartyIdentification2Choice)
	return c.Creditor
}

func (c *CreditTransfer6) AddCreditorAccount() *AccountIdentificationAndName3 {
	c.CreditorAccount = new(AccountIdentificationAndName3)
	return c.CreditorAccount
}

// Payment instrument between a debtor and a creditor, which flows through one or more financial institutions or systems.
type CreditTransfer8 struct {

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	Reference *Max35Text `xml:"Ref,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor. In the context of the payment model, the debtor is also the debit account owner.
	Debtor *PartyIdentification113 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *AccountIdentificationAndName5 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *FinancialInstitutionIdentification10 `xml:"DbtrAgt,omitempty"`

	// Identifies the account of the debtor's agent.
	DebtorAgentAccount *AccountIdentificationAndName5 `xml:"DbtrAgtAcct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent1 *FinancialInstitutionIdentification10 `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *AccountIdentificationAndName5 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent2 *FinancialInstitutionIdentification10 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *AccountIdentificationAndName5 `xml:"IntrmyAgt2Acct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *FinancialInstitutionIdentification10 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *AccountIdentificationAndName5 `xml:"CdtrAgtAcct,omitempty"`

	// Party that receives an amount of money from the debtor. In the context of the payment model, the creditor is also the credit account owner.
	Creditor *PartyIdentification113 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *AccountIdentificationAndName5 `xml:"CdtrAcct"`
}

func (c *CreditTransfer8) SetReference(value string) {
	c.Reference = (*Max35Text)(&value)
}

func (c *CreditTransfer8) AddDebtor() *PartyIdentification113 {
	c.Debtor = new(PartyIdentification113)
	return c.Debtor
}

func (c *CreditTransfer8) AddDebtorAccount() *AccountIdentificationAndName5 {
	c.DebtorAccount = new(AccountIdentificationAndName5)
	return c.DebtorAccount
}

func (c *CreditTransfer8) AddDebtorAgent() *FinancialInstitutionIdentification10 {
	c.DebtorAgent = new(FinancialInstitutionIdentification10)
	return c.DebtorAgent
}

func (c *CreditTransfer8) AddDebtorAgentAccount() *AccountIdentificationAndName5 {
	c.DebtorAgentAccount = new(AccountIdentificationAndName5)
	return c.DebtorAgentAccount
}

func (c *CreditTransfer8) AddIntermediaryAgent1() *FinancialInstitutionIdentification10 {
	c.IntermediaryAgent1 = new(FinancialInstitutionIdentification10)
	return c.IntermediaryAgent1
}

func (c *CreditTransfer8) AddIntermediaryAgent1Account() *AccountIdentificationAndName5 {
	c.IntermediaryAgent1Account = new(AccountIdentificationAndName5)
	return c.IntermediaryAgent1Account
}

func (c *CreditTransfer8) AddIntermediaryAgent2() *FinancialInstitutionIdentification10 {
	c.IntermediaryAgent2 = new(FinancialInstitutionIdentification10)
	return c.IntermediaryAgent2
}

func (c *CreditTransfer8) AddIntermediaryAgent2Account() *AccountIdentificationAndName5 {
	c.IntermediaryAgent2Account = new(AccountIdentificationAndName5)
	return c.IntermediaryAgent2Account
}

func (c *CreditTransfer8) AddCreditorAgent() *FinancialInstitutionIdentification10 {
	c.CreditorAgent = new(FinancialInstitutionIdentification10)
	return c.CreditorAgent
}

func (c *CreditTransfer8) AddCreditorAgentAccount() *AccountIdentificationAndName5 {
	c.CreditorAgentAccount = new(AccountIdentificationAndName5)
	return c.CreditorAgentAccount
}

func (c *CreditTransfer8) AddCreditor() *PartyIdentification113 {
	c.Creditor = new(PartyIdentification113)
	return c.Creditor
}

func (c *CreditTransfer8) AddCreditorAccount() *AccountIdentificationAndName5 {
	c.CreditorAccount = new(AccountIdentificationAndName5)
	return c.CreditorAccount
}
