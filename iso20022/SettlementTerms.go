package iso20022

// Specifies the beneficiary's account information for the settlement of a purchase of goods or services.
type SettlementTerms2 struct {

	// Financial institution that receives the payment transaction on behalf of an account owner, and posts the transaction into the account.
	CreditorAgent *FinancialInstitutionIdentification4Choice `xml:"CdtrAgt,omitempty"`

	// Account to be credited as a result of an instruction.
	CreditorAccount *CashAccount7 `xml:"CdtrAcct"`
}

func (s *SettlementTerms2) AddCreditorAgent() *FinancialInstitutionIdentification4Choice {
	s.CreditorAgent = new(FinancialInstitutionIdentification4Choice)
	return s.CreditorAgent
}

func (s *SettlementTerms2) AddCreditorAccount() *CashAccount7 {
	s.CreditorAccount = new(CashAccount7)
	return s.CreditorAccount
}

// Specifies the beneficiary's account information for the settlement of a purchase of goods or services.
type SettlementTerms3 struct {

	// Financial institution that receives the payment transaction on behalf of an account owner, and posts the transaction into the account.
	CreditorAgent *FinancialInstitutionIdentification4Choice `xml:"CdtrAgt,omitempty"`

	// Account to be credited as a result of an instruction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct"`
}

func (s *SettlementTerms3) AddCreditorAgent() *FinancialInstitutionIdentification4Choice {
	s.CreditorAgent = new(FinancialInstitutionIdentification4Choice)
	return s.CreditorAgent
}

func (s *SettlementTerms3) AddCreditorAccount() *CashAccount24 {
	s.CreditorAccount = new(CashAccount24)
	return s.CreditorAccount
}
