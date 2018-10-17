package iso20022

// Safekeeping or investment account. A safekeeping account is an account on which a securities entry is made. An investment account is an account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type SafekeepingAccount1 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether the securities in the account are fungible, ie, interchangeable.
	FungibleIndicator *YesNoIndicator `xml:"FngbInd"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary1 `xml:"IntrmyInf,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification1Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification1Choice `xml:"AcctSvcr,omitempty"`
}

func (s *SafekeepingAccount1) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SafekeepingAccount1) SetFungibleIndicator(value string) {
	s.FungibleIndicator = (*YesNoIndicator)(&value)
}

func (s *SafekeepingAccount1) SetName(value string) {
	s.Name = (*Max35Text)(&value)
}

func (s *SafekeepingAccount1) SetDesignation(value string) {
	s.Designation = (*Max35Text)(&value)
}

func (s *SafekeepingAccount1) AddIntermediaryInformation() *Intermediary1 {
	newValue := new(Intermediary1)
	s.IntermediaryInformation = append(s.IntermediaryInformation, newValue)
	return newValue
}

func (s *SafekeepingAccount1) AddAccountOwner() *PartyIdentification1Choice {
	s.AccountOwner = new(PartyIdentification1Choice)
	return s.AccountOwner
}

func (s *SafekeepingAccount1) AddAccountServicer() *PartyIdentification1Choice {
	s.AccountServicer = new(PartyIdentification1Choice)
	return s.AccountServicer
}

// Safekeeping or investment account. A safekeeping account is an account on which a securities entry is made. An investment account is an account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type SafekeepingAccount2 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether the securities in the account are fungible, ie, interchangeable.
	FungibleIndicator *YesNoIndicator `xml:"FngbInd"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary11 `xml:"IntrmyInf,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification2Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (s *SafekeepingAccount2) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SafekeepingAccount2) SetFungibleIndicator(value string) {
	s.FungibleIndicator = (*YesNoIndicator)(&value)
}

func (s *SafekeepingAccount2) SetName(value string) {
	s.Name = (*Max35Text)(&value)
}

func (s *SafekeepingAccount2) SetDesignation(value string) {
	s.Designation = (*Max35Text)(&value)
}

func (s *SafekeepingAccount2) AddIntermediaryInformation() *Intermediary11 {
	newValue := new(Intermediary11)
	s.IntermediaryInformation = append(s.IntermediaryInformation, newValue)
	return newValue
}

func (s *SafekeepingAccount2) AddAccountOwner() *PartyIdentification2Choice {
	s.AccountOwner = new(PartyIdentification2Choice)
	return s.AccountOwner
}

func (s *SafekeepingAccount2) AddAccountServicer() *PartyIdentification2Choice {
	s.AccountServicer = new(PartyIdentification2Choice)
	return s.AccountServicer
}

// A safekeeping account is an account on which a securities entry is made.
type SafekeepingAccount3 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification9Choice `xml:"AcctOwnr,omitempty"`

	// Identification of a subaccount within the safekeeping account
	SubAccountDetails *SubAccount2 `xml:"SubAcctDtls,omitempty"`

	// Quantity of securities in the sub-balance.
	InstructedBalance []*HoldingBalance4 `xml:"InstdBal"`

	// Owner of the voting rights.
	RightsHolder []*PartyIdentification9Choice `xml:"RghtsHldr,omitempty"`
}

func (s *SafekeepingAccount3) SetAccountIdentification(value string) {
	s.AccountIdentification = (*Max35Text)(&value)
}

func (s *SafekeepingAccount3) AddAccountOwner() *PartyIdentification9Choice {
	s.AccountOwner = new(PartyIdentification9Choice)
	return s.AccountOwner
}

func (s *SafekeepingAccount3) AddSubAccountDetails() *SubAccount2 {
	s.SubAccountDetails = new(SubAccount2)
	return s.SubAccountDetails
}

func (s *SafekeepingAccount3) AddInstructedBalance() *HoldingBalance4 {
	newValue := new(HoldingBalance4)
	s.InstructedBalance = append(s.InstructedBalance, newValue)
	return newValue
}

func (s *SafekeepingAccount3) AddRightsHolder() *PartyIdentification9Choice {
	newValue := new(PartyIdentification9Choice)
	s.RightsHolder = append(s.RightsHolder, newValue)
	return newValue
}

// A safekeeping account is an account on which a securities entry is made.
type SafekeepingAccount4 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification9Choice `xml:"AcctOwnr,omitempty"`

	// Identification of a subaccount within the safekeeping account
	SubAccountDetails *SubAccount2 `xml:"SubAcctDtls,omitempty"`

	// Quantity of securities in the sub-balance.
	InstructedBalance []*HoldingBalance5 `xml:"InstdBal"`

	// Owner of the voting rights.
	RightsHolder []*PartyIdentification9Choice `xml:"RghtsHldr,omitempty"`
}

func (s *SafekeepingAccount4) SetAccountIdentification(value string) {
	s.AccountIdentification = (*Max35Text)(&value)
}

func (s *SafekeepingAccount4) AddAccountOwner() *PartyIdentification9Choice {
	s.AccountOwner = new(PartyIdentification9Choice)
	return s.AccountOwner
}

func (s *SafekeepingAccount4) AddSubAccountDetails() *SubAccount2 {
	s.SubAccountDetails = new(SubAccount2)
	return s.SubAccountDetails
}

func (s *SafekeepingAccount4) AddInstructedBalance() *HoldingBalance5 {
	newValue := new(HoldingBalance5)
	s.InstructedBalance = append(s.InstructedBalance, newValue)
	return newValue
}

func (s *SafekeepingAccount4) AddRightsHolder() *PartyIdentification9Choice {
	newValue := new(PartyIdentification9Choice)
	s.RightsHolder = append(s.RightsHolder, newValue)
	return newValue
}

// A safekeeping account is an account on which a securities entry is made.
type SafekeepingAccount6 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification40Choice `xml:"AcctOwnr,omitempty"`

	// Identification of a subaccount within the safekeeping account
	SubAccountDetails *SubAccount2 `xml:"SubAcctDtls,omitempty"`

	// Quantity of securities in the sub-balance.
	InstructedBalance []*HoldingBalance8 `xml:"InstdBal"`

	// Owner of the voting rights.
	RightsHolder []*PartyIdentification40Choice `xml:"RghtsHldr,omitempty"`
}

func (s *SafekeepingAccount6) SetAccountIdentification(value string) {
	s.AccountIdentification = (*Max35Text)(&value)
}

func (s *SafekeepingAccount6) AddAccountOwner() *PartyIdentification40Choice {
	s.AccountOwner = new(PartyIdentification40Choice)
	return s.AccountOwner
}

func (s *SafekeepingAccount6) AddSubAccountDetails() *SubAccount2 {
	s.SubAccountDetails = new(SubAccount2)
	return s.SubAccountDetails
}

func (s *SafekeepingAccount6) AddInstructedBalance() *HoldingBalance8 {
	newValue := new(HoldingBalance8)
	s.InstructedBalance = append(s.InstructedBalance, newValue)
	return newValue
}

func (s *SafekeepingAccount6) AddRightsHolder() *PartyIdentification40Choice {
	newValue := new(PartyIdentification40Choice)
	s.RightsHolder = append(s.RightsHolder, newValue)
	return newValue
}

// Account on which a securities entry is made.
type SafekeepingAccount7 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the account,  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Holdings of level 1.
	AccountSubLevel1 []*AccountSubLevel11 `xml:"AcctSubLvl1,omitempty"`

	// Difference in holdings between the safekeeping account and the sub-accounts of level 1.
	AccountSubLevel1Difference []*AggregateHoldingBalance2 `xml:"AcctSubLvl1Diff,omitempty"`
}

func (s *SafekeepingAccount7) AddAccountIdentification() *SecuritiesAccount19 {
	s.AccountIdentification = new(SecuritiesAccount19)
	return s.AccountIdentification
}

func (s *SafekeepingAccount7) AddAccountOwner() *PartyIdentification100 {
	s.AccountOwner = new(PartyIdentification100)
	return s.AccountOwner
}

func (s *SafekeepingAccount7) AddAccountServicer() *PartyIdentification100 {
	s.AccountServicer = new(PartyIdentification100)
	return s.AccountServicer
}

func (s *SafekeepingAccount7) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	s.BeneficialOwner = append(s.BeneficialOwner, newValue)
	return newValue
}

func (s *SafekeepingAccount7) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	s.BalanceDetails = append(s.BalanceDetails, newValue)
	return newValue
}

func (s *SafekeepingAccount7) AddAccountSubLevel1() *AccountSubLevel11 {
	newValue := new(AccountSubLevel11)
	s.AccountSubLevel1 = append(s.AccountSubLevel1, newValue)
	return newValue
}

func (s *SafekeepingAccount7) AddAccountSubLevel1Difference() *AggregateHoldingBalance2 {
	newValue := new(AggregateHoldingBalance2)
	s.AccountSubLevel1Difference = append(s.AccountSubLevel1Difference, newValue)
	return newValue
}
