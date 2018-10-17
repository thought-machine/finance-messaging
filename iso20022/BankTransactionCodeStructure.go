package iso20022

// Set of elements to fully identify the type of the bank transaction entry.
type BankTransactionCodeStructure1 struct {

	// Specifies the domain, the family and the sub-family of the bank transaction code, in a structured and hierarchical format.
	//
	// Usage: If a specific family or subfamily code cannot be provided, the generic family code defined for the domain or the generic subfamily code defined for the family should be provided.
	Domain *BankTransactionCodeStructure2 `xml:"Domn,omitempty"`

	// Proprietary identification of the bank transaction code, as defined by the issuer.
	Proprietary *ProprietaryBankTransactionCodeStructure1 `xml:"Prtry,omitempty"`
}

func (b *BankTransactionCodeStructure1) AddDomain() *BankTransactionCodeStructure2 {
	b.Domain = new(BankTransactionCodeStructure2)
	return b.Domain
}

func (b *BankTransactionCodeStructure1) AddProprietary() *ProprietaryBankTransactionCodeStructure1 {
	b.Proprietary = new(ProprietaryBankTransactionCodeStructure1)
	return b.Proprietary
}

// Code of the underlying bank transaction.
type BankTransactionCodeStructure2 struct {

	// Specifies the business area of the underlying transaction.
	Code *ExternalBankTransactionDomainCode `xml:"Cd"`

	// Specifies the family and the sub-family of the bank transaction code, within a specific domain, in a structured and hierarchical format.
	Family *BankTransactionCodeStructure3 `xml:"Fmly"`
}

func (b *BankTransactionCodeStructure2) SetCode(value string) {
	b.Code = (*ExternalBankTransactionDomainCode)(&value)
}

func (b *BankTransactionCodeStructure2) AddFamily() *BankTransactionCodeStructure3 {
	b.Family = new(BankTransactionCodeStructure3)
	return b.Family
}

// Code of the underlying bank transaction.
type BankTransactionCodeStructure3 struct {

	// Specifies the family within a domain.
	Code *ExternalBankTransactionFamilyCode `xml:"Cd"`

	// Specifies the sub-product family within a specific family.
	SubFamilyCode *ExternalBankTransactionSubFamilyCode `xml:"SubFmlyCd"`
}

func (b *BankTransactionCodeStructure3) SetCode(value string) {
	b.Code = (*ExternalBankTransactionFamilyCode)(&value)
}

func (b *BankTransactionCodeStructure3) SetSubFamilyCode(value string) {
	b.SubFamilyCode = (*ExternalBankTransactionSubFamilyCode)(&value)
}

// Set of elements used to identify the type or operations code of a transaction entry.
type BankTransactionCodeStructure4 struct {

	// Set of elements used to provide the domain, the family and the sub-family of the bank transaction code, in a structured and hierarchical format.
	//
	// Usage: If a specific family or sub-family code cannot be provided, the generic family code defined for the domain or the generic sub-family code defined for the family should be provided.
	Domain *BankTransactionCodeStructure5 `xml:"Domn,omitempty"`

	// Bank transaction code in a proprietary form, as defined by the issuer.
	Proprietary *ProprietaryBankTransactionCodeStructure1 `xml:"Prtry,omitempty"`
}

func (b *BankTransactionCodeStructure4) AddDomain() *BankTransactionCodeStructure5 {
	b.Domain = new(BankTransactionCodeStructure5)
	return b.Domain
}

func (b *BankTransactionCodeStructure4) AddProprietary() *ProprietaryBankTransactionCodeStructure1 {
	b.Proprietary = new(ProprietaryBankTransactionCodeStructure1)
	return b.Proprietary
}

// Set of elements used to identify the type or operations code of a transaction entry.
type BankTransactionCodeStructure5 struct {

	// Specifies the business area of the underlying transaction.
	Code *ExternalBankTransactionDomain1Code `xml:"Cd"`

	// Specifies the family and the sub-family of the bank transaction code, within a specific domain, in a structured and hierarchical format.
	Family *BankTransactionCodeStructure6 `xml:"Fmly"`
}

func (b *BankTransactionCodeStructure5) SetCode(value string) {
	b.Code = (*ExternalBankTransactionDomain1Code)(&value)
}

func (b *BankTransactionCodeStructure5) AddFamily() *BankTransactionCodeStructure6 {
	b.Family = new(BankTransactionCodeStructure6)
	return b.Family
}

// Set of elements used to identify the type or operations code of a transaction entry.
type BankTransactionCodeStructure6 struct {

	// Specifies the family within a domain.
	Code *ExternalBankTransactionFamily1Code `xml:"Cd"`

	// Specifies the sub-product family within a specific family.
	SubFamilyCode *ExternalBankTransactionSubFamily1Code `xml:"SubFmlyCd"`
}

func (b *BankTransactionCodeStructure6) SetCode(value string) {
	b.Code = (*ExternalBankTransactionFamily1Code)(&value)
}

func (b *BankTransactionCodeStructure6) SetSubFamilyCode(value string) {
	b.SubFamilyCode = (*ExternalBankTransactionSubFamily1Code)(&value)
}
