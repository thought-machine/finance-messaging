package iso20022

// Investment fund transactions for a specific financial instrument.
type InvestmentFundTransactionsByFund1 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification1Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Specifies the form, ie, ownership, of the security.
	Form *FormOfSecurity1Code `xml:"Form,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Process of buying, selling, switching or transferring fund units.
	TransactionDetails []*InvestmentFundTransaction2 `xml:"TxDtls"`

	// Balance of the financial instrument for this specific statement page.
	BalanceByPage *PaginationBalance1 `xml:"BalByPg,omitempty"`
}

func (i *InvestmentFundTransactionsByFund1) AddIdentification() *SecurityIdentification1Choice {
	i.Identification = new(SecurityIdentification1Choice)
	return i.Identification
}

func (i *InvestmentFundTransactionsByFund1) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *InvestmentFundTransactionsByFund1) SetSupplementaryIdentification(value string) {
	i.SupplementaryIdentification = (*Max35Text)(&value)
}

func (i *InvestmentFundTransactionsByFund1) SetForm(value string) {
	i.Form = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentFundTransactionsByFund1) SetClassType(value string) {
	i.ClassType = (*Max35Text)(&value)
}

func (i *InvestmentFundTransactionsByFund1) SetDistributionPolicy(value string) {
	i.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (i *InvestmentFundTransactionsByFund1) AddTransactionDetails() *InvestmentFundTransaction2 {
	newValue := new(InvestmentFundTransaction2)
	i.TransactionDetails = append(i.TransactionDetails, newValue)
	return newValue
}

func (i *InvestmentFundTransactionsByFund1) AddBalanceByPage() *PaginationBalance1 {
	i.BalanceByPage = new(PaginationBalance1)
	return i.BalanceByPage
}

// Investment fund transactions for a specific financial instrument.
type InvestmentFundTransactionsByFund2 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Process of buying, selling, switching or transferring fund units.
	TransactionDetails []*InvestmentFundTransaction3 `xml:"TxDtls"`

	// Balance of the financial instrument for this specific statement page.
	BalanceByPage *PaginationBalance1 `xml:"BalByPg,omitempty"`
}

func (i *InvestmentFundTransactionsByFund2) AddIdentification() *SecurityIdentification3Choice {
	i.Identification = new(SecurityIdentification3Choice)
	return i.Identification
}

func (i *InvestmentFundTransactionsByFund2) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *InvestmentFundTransactionsByFund2) SetSupplementaryIdentification(value string) {
	i.SupplementaryIdentification = (*Max35Text)(&value)
}

func (i *InvestmentFundTransactionsByFund2) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentFundTransactionsByFund2) SetClassType(value string) {
	i.ClassType = (*Max35Text)(&value)
}

func (i *InvestmentFundTransactionsByFund2) SetDistributionPolicy(value string) {
	i.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (i *InvestmentFundTransactionsByFund2) AddTransactionDetails() *InvestmentFundTransaction3 {
	newValue := new(InvestmentFundTransaction3)
	i.TransactionDetails = append(i.TransactionDetails, newValue)
	return newValue
}

func (i *InvestmentFundTransactionsByFund2) AddBalanceByPage() *PaginationBalance1 {
	i.BalanceByPage = new(PaginationBalance1)
	return i.BalanceByPage
}

// Investment fund transactions for a specific financial instrument.
type InvestmentFundTransactionsByFund3 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Form, that is, ownership, of the security, for example, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, for example, front end or back end, an income policy, for example, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, for example, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Income policy relating to a class type, that is, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Process of buying, selling, switching or transferring fund units.
	TransactionDetails []*InvestmentFundTransaction4 `xml:"TxDtls"`

	// Balance of the financial instrument for this specific statement page.
	BalanceByPage *PaginationBalance2 `xml:"BalByPg,omitempty"`
}

func (i *InvestmentFundTransactionsByFund3) AddIdentification() *SecurityIdentification3Choice {
	i.Identification = new(SecurityIdentification3Choice)
	return i.Identification
}

func (i *InvestmentFundTransactionsByFund3) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *InvestmentFundTransactionsByFund3) SetSupplementaryIdentification(value string) {
	i.SupplementaryIdentification = (*Max35Text)(&value)
}

func (i *InvestmentFundTransactionsByFund3) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentFundTransactionsByFund3) SetClassType(value string) {
	i.ClassType = (*Max35Text)(&value)
}

func (i *InvestmentFundTransactionsByFund3) SetDistributionPolicy(value string) {
	i.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (i *InvestmentFundTransactionsByFund3) AddTransactionDetails() *InvestmentFundTransaction4 {
	newValue := new(InvestmentFundTransaction4)
	i.TransactionDetails = append(i.TransactionDetails, newValue)
	return newValue
}

func (i *InvestmentFundTransactionsByFund3) AddBalanceByPage() *PaginationBalance2 {
	i.BalanceByPage = new(PaginationBalance2)
	return i.BalanceByPage
}
