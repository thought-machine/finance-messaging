package iso20022

// Specifies cash parties in the framework of a corporate action event.
type CashParties10 struct {

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount52 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount54 `xml:"CdtrAgt,omitempty"`

	// Party that has reimbursed the account owner with funds to which they were legally entitled.
	MarketClaimCounterparty *PartyIdentificationAndAccount52 `xml:"MktClmCtrPty,omitempty"`
}

func (c *CashParties10) AddCreditor() *PartyIdentificationAndAccount52 {
	c.Creditor = new(PartyIdentificationAndAccount52)
	return c.Creditor
}

func (c *CashParties10) AddCreditorAgent() *PartyIdentificationAndAccount54 {
	c.CreditorAgent = new(PartyIdentificationAndAccount54)
	return c.CreditorAgent
}

func (c *CashParties10) AddMarketClaimCounterparty() *PartyIdentificationAndAccount52 {
	c.MarketClaimCounterparty = new(PartyIdentificationAndAccount52)
	return c.MarketClaimCounterparty
}

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties11 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount53 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount55 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount53 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount55 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties11) AddDebtor() *PartyIdentificationAndAccount53 {
	c.Debtor = new(PartyIdentificationAndAccount53)
	return c.Debtor
}

func (c *CashParties11) AddDebtorAgent() *PartyIdentificationAndAccount55 {
	c.DebtorAgent = new(PartyIdentificationAndAccount55)
	return c.DebtorAgent
}

func (c *CashParties11) AddCreditor() *PartyIdentificationAndAccount53 {
	c.Creditor = new(PartyIdentificationAndAccount53)
	return c.Creditor
}

func (c *CashParties11) AddCreditorAgent() *PartyIdentificationAndAccount55 {
	c.CreditorAgent = new(PartyIdentificationAndAccount55)
	return c.CreditorAgent
}

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties17 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount39 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount50 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount39 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount50 `xml:"CdtrAgt,omitempty"`

	// Financial institution through which the transaction must pass to reach the account with institution.
	Intermediary *PartyIdentificationAndAccount50 `xml:"Intrmy,omitempty"`
}

func (c *CashParties17) AddDebtor() *PartyIdentificationAndAccount39 {
	c.Debtor = new(PartyIdentificationAndAccount39)
	return c.Debtor
}

func (c *CashParties17) AddDebtorAgent() *PartyIdentificationAndAccount50 {
	c.DebtorAgent = new(PartyIdentificationAndAccount50)
	return c.DebtorAgent
}

func (c *CashParties17) AddCreditor() *PartyIdentificationAndAccount39 {
	c.Creditor = new(PartyIdentificationAndAccount39)
	return c.Creditor
}

func (c *CashParties17) AddCreditorAgent() *PartyIdentificationAndAccount50 {
	c.CreditorAgent = new(PartyIdentificationAndAccount50)
	return c.CreditorAgent
}

func (c *CashParties17) AddIntermediary() *PartyIdentificationAndAccount50 {
	c.Intermediary = new(PartyIdentificationAndAccount50)
	return c.Intermediary
}

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties18 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount80 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount80 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount80 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount80 `xml:"CdtrAgt,omitempty"`

	// Financial institution through which the transaction must pass to reach the account with institution.
	Intermediary *PartyIdentificationAndAccount80 `xml:"Intrmy,omitempty"`
}

func (c *CashParties18) AddDebtor() *PartyIdentificationAndAccount80 {
	c.Debtor = new(PartyIdentificationAndAccount80)
	return c.Debtor
}

func (c *CashParties18) AddDebtorAgent() *PartyIdentificationAndAccount80 {
	c.DebtorAgent = new(PartyIdentificationAndAccount80)
	return c.DebtorAgent
}

func (c *CashParties18) AddCreditor() *PartyIdentificationAndAccount80 {
	c.Creditor = new(PartyIdentificationAndAccount80)
	return c.Creditor
}

func (c *CashParties18) AddCreditorAgent() *PartyIdentificationAndAccount80 {
	c.CreditorAgent = new(PartyIdentificationAndAccount80)
	return c.CreditorAgent
}

func (c *CashParties18) AddIntermediary() *PartyIdentificationAndAccount80 {
	c.Intermediary = new(PartyIdentificationAndAccount80)
	return c.Intermediary
}

// Specifies cash parties in the framework of a corporate action event.
type CashParties2 struct {

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount17 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount18 `xml:"CdtrAgt,omitempty"`

	// Party that has reimbursed the account owner with funds to which they were legally entitled.
	MarketClaimCounterparty *PartyIdentificationAndAccount17 `xml:"MktClmCtrPty,omitempty"`
}

func (c *CashParties2) AddCreditor() *PartyIdentificationAndAccount17 {
	c.Creditor = new(PartyIdentificationAndAccount17)
	return c.Creditor
}

func (c *CashParties2) AddCreditorAgent() *PartyIdentificationAndAccount18 {
	c.CreditorAgent = new(PartyIdentificationAndAccount18)
	return c.CreditorAgent
}

func (c *CashParties2) AddMarketClaimCounterparty() *PartyIdentificationAndAccount17 {
	c.MarketClaimCounterparty = new(PartyIdentificationAndAccount17)
	return c.MarketClaimCounterparty
}

// Specifies cash parties in the framework of a corporate action event.
type CashParties21 struct {

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount52 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount101 `xml:"CdtrAgt,omitempty"`

	// Party that has reimbursed the account owner with funds to which they were legally entitled.
	MarketClaimCounterparty *PartyIdentificationAndAccount52 `xml:"MktClmCtrPty,omitempty"`
}

func (c *CashParties21) AddCreditor() *PartyIdentificationAndAccount52 {
	c.Creditor = new(PartyIdentificationAndAccount52)
	return c.Creditor
}

func (c *CashParties21) AddCreditorAgent() *PartyIdentificationAndAccount101 {
	c.CreditorAgent = new(PartyIdentificationAndAccount101)
	return c.CreditorAgent
}

func (c *CashParties21) AddMarketClaimCounterparty() *PartyIdentificationAndAccount52 {
	c.MarketClaimCounterparty = new(PartyIdentificationAndAccount52)
	return c.MarketClaimCounterparty
}

// Cash settlement chain parties and accounts.
type CashParties24 struct {

	// Party to which the payment amount must be ultimately delivered. In some cases, this may be a fund.
	//
	Creditor *PartyIdentificationAndAccount96 `xml:"Cdtr"`

	// Financial institution that services the cash account of the beneficiary (creditor). In some markets, this is also known as receiving agent. The creditor agent is the party where the payment amount must be ultimately delivered on behalf of the beneficiary (creditor), that is, the party where the beneficiary has its account.
	CreditorAgent *PartyIdentificationAndAccount97 `xml:"CdtrAgt"`

	// Financial institution through which the transaction must pass to reach the account with institution (creditor agent).
	Intermediary *PartyIdentificationAndAccount97 `xml:"Intrmy,omitempty"`

	// Financial institution through which the transaction must pass to reach the account with institution (creditor agent).
	Intermediary2 *PartyIdentificationAndAccount97 `xml:"Intrmy2,omitempty"`
}

func (c *CashParties24) AddCreditor() *PartyIdentificationAndAccount96 {
	c.Creditor = new(PartyIdentificationAndAccount96)
	return c.Creditor
}

func (c *CashParties24) AddCreditorAgent() *PartyIdentificationAndAccount97 {
	c.CreditorAgent = new(PartyIdentificationAndAccount97)
	return c.CreditorAgent
}

func (c *CashParties24) AddIntermediary() *PartyIdentificationAndAccount97 {
	c.Intermediary = new(PartyIdentificationAndAccount97)
	return c.Intermediary
}

func (c *CashParties24) AddIntermediary2() *PartyIdentificationAndAccount97 {
	c.Intermediary2 = new(PartyIdentificationAndAccount97)
	return c.Intermediary2
}

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties25 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount111 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount112 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount111 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount112 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties25) AddDebtor() *PartyIdentificationAndAccount111 {
	c.Debtor = new(PartyIdentificationAndAccount111)
	return c.Debtor
}

func (c *CashParties25) AddDebtorAgent() *PartyIdentificationAndAccount112 {
	c.DebtorAgent = new(PartyIdentificationAndAccount112)
	return c.DebtorAgent
}

func (c *CashParties25) AddCreditor() *PartyIdentificationAndAccount111 {
	c.Creditor = new(PartyIdentificationAndAccount111)
	return c.Creditor
}

func (c *CashParties25) AddCreditorAgent() *PartyIdentificationAndAccount112 {
	c.CreditorAgent = new(PartyIdentificationAndAccount112)
	return c.CreditorAgent
}

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties26 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount111 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount112 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount111 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount112 `xml:"CdtrAgt,omitempty"`

	// Financial institution through which the transaction must pass to reach the account with institution.
	Intermediary *PartyIdentificationAndAccount112 `xml:"Intrmy,omitempty"`
}

func (c *CashParties26) AddDebtor() *PartyIdentificationAndAccount111 {
	c.Debtor = new(PartyIdentificationAndAccount111)
	return c.Debtor
}

func (c *CashParties26) AddDebtorAgent() *PartyIdentificationAndAccount112 {
	c.DebtorAgent = new(PartyIdentificationAndAccount112)
	return c.DebtorAgent
}

func (c *CashParties26) AddCreditor() *PartyIdentificationAndAccount111 {
	c.Creditor = new(PartyIdentificationAndAccount111)
	return c.Creditor
}

func (c *CashParties26) AddCreditorAgent() *PartyIdentificationAndAccount112 {
	c.CreditorAgent = new(PartyIdentificationAndAccount112)
	return c.CreditorAgent
}

func (c *CashParties26) AddIntermediary() *PartyIdentificationAndAccount112 {
	c.Intermediary = new(PartyIdentificationAndAccount112)
	return c.Intermediary
}

// Specifies cash parties in the framework of a corporate action event.
type CashParties28 struct {

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount120 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount121 `xml:"CdtrAgt,omitempty"`

	// Party that has reimbursed the account owner with funds to which they were legally entitled.
	MarketClaimCounterparty *PartyIdentificationAndAccount120 `xml:"MktClmCtrPty,omitempty"`
}

func (c *CashParties28) AddCreditor() *PartyIdentificationAndAccount120 {
	c.Creditor = new(PartyIdentificationAndAccount120)
	return c.Creditor
}

func (c *CashParties28) AddCreditorAgent() *PartyIdentificationAndAccount121 {
	c.CreditorAgent = new(PartyIdentificationAndAccount121)
	return c.CreditorAgent
}

func (c *CashParties28) AddMarketClaimCounterparty() *PartyIdentificationAndAccount120 {
	c.MarketClaimCounterparty = new(PartyIdentificationAndAccount120)
	return c.MarketClaimCounterparty
}

// Specifies cash parties in the framework of a corporate action event.
type CashParties29 struct {

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount129 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount130 `xml:"CdtrAgt,omitempty"`

	// Party that has reimbursed the account owner with funds to which they were legally entitled.
	MarketClaimCounterparty *PartyIdentificationAndAccount129 `xml:"MktClmCtrPty,omitempty"`
}

func (c *CashParties29) AddCreditor() *PartyIdentificationAndAccount129 {
	c.Creditor = new(PartyIdentificationAndAccount129)
	return c.Creditor
}

func (c *CashParties29) AddCreditorAgent() *PartyIdentificationAndAccount130 {
	c.CreditorAgent = new(PartyIdentificationAndAccount130)
	return c.CreditorAgent
}

func (c *CashParties29) AddMarketClaimCounterparty() *PartyIdentificationAndAccount129 {
	c.MarketClaimCounterparty = new(PartyIdentificationAndAccount129)
	return c.MarketClaimCounterparty
}

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties3 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount20 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount15 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount20 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount15 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties3) AddDebtor() *PartyIdentificationAndAccount20 {
	c.Debtor = new(PartyIdentificationAndAccount20)
	return c.Debtor
}

func (c *CashParties3) AddDebtorAgent() *PartyIdentificationAndAccount15 {
	c.DebtorAgent = new(PartyIdentificationAndAccount15)
	return c.DebtorAgent
}

func (c *CashParties3) AddCreditor() *PartyIdentificationAndAccount20 {
	c.Creditor = new(PartyIdentificationAndAccount20)
	return c.Creditor
}

func (c *CashParties3) AddCreditorAgent() *PartyIdentificationAndAccount15 {
	c.CreditorAgent = new(PartyIdentificationAndAccount15)
	return c.CreditorAgent
}

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties30 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount133 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount134 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount133 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount134 `xml:"CdtrAgt,omitempty"`

	// Financial institution through which the transaction must pass to reach the account with institution.
	Intermediary *PartyIdentificationAndAccount134 `xml:"Intrmy,omitempty"`
}

func (c *CashParties30) AddDebtor() *PartyIdentificationAndAccount133 {
	c.Debtor = new(PartyIdentificationAndAccount133)
	return c.Debtor
}

func (c *CashParties30) AddDebtorAgent() *PartyIdentificationAndAccount134 {
	c.DebtorAgent = new(PartyIdentificationAndAccount134)
	return c.DebtorAgent
}

func (c *CashParties30) AddCreditor() *PartyIdentificationAndAccount133 {
	c.Creditor = new(PartyIdentificationAndAccount133)
	return c.Creditor
}

func (c *CashParties30) AddCreditorAgent() *PartyIdentificationAndAccount134 {
	c.CreditorAgent = new(PartyIdentificationAndAccount134)
	return c.CreditorAgent
}

func (c *CashParties30) AddIntermediary() *PartyIdentificationAndAccount134 {
	c.Intermediary = new(PartyIdentificationAndAccount134)
	return c.Intermediary
}

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties32 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount133 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount134 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount133 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount134 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties32) AddDebtor() *PartyIdentificationAndAccount133 {
	c.Debtor = new(PartyIdentificationAndAccount133)
	return c.Debtor
}

func (c *CashParties32) AddDebtorAgent() *PartyIdentificationAndAccount134 {
	c.DebtorAgent = new(PartyIdentificationAndAccount134)
	return c.DebtorAgent
}

func (c *CashParties32) AddCreditor() *PartyIdentificationAndAccount133 {
	c.Creditor = new(PartyIdentificationAndAccount133)
	return c.Creditor
}

func (c *CashParties32) AddCreditorAgent() *PartyIdentificationAndAccount134 {
	c.CreditorAgent = new(PartyIdentificationAndAccount134)
	return c.CreditorAgent
}

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties6 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount80 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount80 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount80 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount80 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties6) AddDebtor() *PartyIdentificationAndAccount80 {
	c.Debtor = new(PartyIdentificationAndAccount80)
	return c.Debtor
}

func (c *CashParties6) AddDebtorAgent() *PartyIdentificationAndAccount80 {
	c.DebtorAgent = new(PartyIdentificationAndAccount80)
	return c.DebtorAgent
}

func (c *CashParties6) AddCreditor() *PartyIdentificationAndAccount80 {
	c.Creditor = new(PartyIdentificationAndAccount80)
	return c.Creditor
}

func (c *CashParties6) AddCreditorAgent() *PartyIdentificationAndAccount80 {
	c.CreditorAgent = new(PartyIdentificationAndAccount80)
	return c.CreditorAgent
}

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties7 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount38 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount50 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount38 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount50 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties7) AddDebtor() *PartyIdentificationAndAccount38 {
	c.Debtor = new(PartyIdentificationAndAccount38)
	return c.Debtor
}

func (c *CashParties7) AddDebtorAgent() *PartyIdentificationAndAccount50 {
	c.DebtorAgent = new(PartyIdentificationAndAccount50)
	return c.DebtorAgent
}

func (c *CashParties7) AddCreditor() *PartyIdentificationAndAccount38 {
	c.Creditor = new(PartyIdentificationAndAccount38)
	return c.Creditor
}

func (c *CashParties7) AddCreditorAgent() *PartyIdentificationAndAccount50 {
	c.CreditorAgent = new(PartyIdentificationAndAccount50)
	return c.CreditorAgent
}

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties8 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount39 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount50 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount39 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount50 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties8) AddDebtor() *PartyIdentificationAndAccount39 {
	c.Debtor = new(PartyIdentificationAndAccount39)
	return c.Debtor
}

func (c *CashParties8) AddDebtorAgent() *PartyIdentificationAndAccount50 {
	c.DebtorAgent = new(PartyIdentificationAndAccount50)
	return c.DebtorAgent
}

func (c *CashParties8) AddCreditor() *PartyIdentificationAndAccount39 {
	c.Creditor = new(PartyIdentificationAndAccount39)
	return c.Creditor
}

func (c *CashParties8) AddCreditorAgent() *PartyIdentificationAndAccount50 {
	c.CreditorAgent = new(PartyIdentificationAndAccount50)
	return c.CreditorAgent
}

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties9 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount48 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount50 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount48 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount50 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties9) AddDebtor() *PartyIdentificationAndAccount48 {
	c.Debtor = new(PartyIdentificationAndAccount48)
	return c.Debtor
}

func (c *CashParties9) AddDebtorAgent() *PartyIdentificationAndAccount50 {
	c.DebtorAgent = new(PartyIdentificationAndAccount50)
	return c.DebtorAgent
}

func (c *CashParties9) AddCreditor() *PartyIdentificationAndAccount48 {
	c.Creditor = new(PartyIdentificationAndAccount48)
	return c.Creditor
}

func (c *CashParties9) AddCreditorAgent() *PartyIdentificationAndAccount50 {
	c.CreditorAgent = new(PartyIdentificationAndAccount50)
	return c.CreditorAgent
}
