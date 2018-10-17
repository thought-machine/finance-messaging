package iso20022

// Set of elements used to provide identification information.
type IdentificationInformation1 struct {

	// Account owner that owes an amount of money or to whom an amount of money is due.
	Party *PartyIdentification32 `xml:"Pty,omitempty"`

	// Unambiguous identification of the account of a party.
	Account *AccountIdentification4Choice `xml:"Acct,omitempty"`

	// Financial institution servicing an account for a party.
	Agent *BranchAndFinancialInstitutionIdentification4 `xml:"Agt,omitempty"`
}

func (i *IdentificationInformation1) AddParty() *PartyIdentification32 {
	i.Party = new(PartyIdentification32)
	return i.Party
}

func (i *IdentificationInformation1) AddAccount() *AccountIdentification4Choice {
	i.Account = new(AccountIdentification4Choice)
	return i.Account
}

func (i *IdentificationInformation1) AddAgent() *BranchAndFinancialInstitutionIdentification4 {
	i.Agent = new(BranchAndFinancialInstitutionIdentification4)
	return i.Agent
}

// Provides the details of the identification information.
type IdentificationInformation2 struct {

	// Account owner that owes an amount of money or to whom an amount of money is due.
	Party *PartyIdentification43 `xml:"Pty,omitempty"`

	// Unambiguous identification of the account of a party.
	Account *AccountIdentification4Choice `xml:"Acct,omitempty"`

	// Financial institution servicing an account for a party.
	Agent *BranchAndFinancialInstitutionIdentification5 `xml:"Agt,omitempty"`
}

func (i *IdentificationInformation2) AddParty() *PartyIdentification43 {
	i.Party = new(PartyIdentification43)
	return i.Party
}

func (i *IdentificationInformation2) AddAccount() *AccountIdentification4Choice {
	i.Account = new(AccountIdentification4Choice)
	return i.Account
}

func (i *IdentificationInformation2) AddAgent() *BranchAndFinancialInstitutionIdentification5 {
	i.Agent = new(BranchAndFinancialInstitutionIdentification5)
	return i.Agent
}
