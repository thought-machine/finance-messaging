package iso20022

// Organisation established primarily to provide financial services.
type BranchAndFinancialInstitutionIdentification struct {

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	FinancialInstitutionIdentification *FinancialInstitutionIdentification1 `xml:"FinInstnId"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	BranchIdentification *BranchData `xml:"BrnchId,omitempty"`
}

func (b *BranchAndFinancialInstitutionIdentification) AddFinancialInstitutionIdentification() *FinancialInstitutionIdentification1 {
	b.FinancialInstitutionIdentification = new(FinancialInstitutionIdentification1)
	return b.FinancialInstitutionIdentification
}

func (b *BranchAndFinancialInstitutionIdentification) AddBranchIdentification() *BranchData {
	b.BranchIdentification = new(BranchData)
	return b.BranchIdentification
}

// Unique and unambiguous identifier of a financial institution or a branch of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
type BranchAndFinancialInstitutionIdentification3 struct {

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	FinancialInstitutionIdentification *FinancialInstitutionIdentification5Choice `xml:"FinInstnId"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	BranchIdentification *BranchData `xml:"BrnchId,omitempty"`
}

func (b *BranchAndFinancialInstitutionIdentification3) AddFinancialInstitutionIdentification() *FinancialInstitutionIdentification5Choice {
	b.FinancialInstitutionIdentification = new(FinancialInstitutionIdentification5Choice)
	return b.FinancialInstitutionIdentification
}

func (b *BranchAndFinancialInstitutionIdentification3) AddBranchIdentification() *BranchData {
	b.BranchIdentification = new(BranchData)
	return b.BranchIdentification
}

// Set of elements used to uniquely and unambiguously identify a financial institution or a branch of a financial institution.
type BranchAndFinancialInstitutionIdentification4 struct {

	// Unique and unambiguous identification of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	FinancialInstitutionIdentification *FinancialInstitutionIdentification7 `xml:"FinInstnId"`

	// Identifies a specific branch of a financial institution.
	//
	// Usage: This component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	BranchIdentification *BranchData2 `xml:"BrnchId,omitempty"`
}

func (b *BranchAndFinancialInstitutionIdentification4) AddFinancialInstitutionIdentification() *FinancialInstitutionIdentification7 {
	b.FinancialInstitutionIdentification = new(FinancialInstitutionIdentification7)
	return b.FinancialInstitutionIdentification
}

func (b *BranchAndFinancialInstitutionIdentification4) AddBranchIdentification() *BranchData2 {
	b.BranchIdentification = new(BranchData2)
	return b.BranchIdentification
}

// Set of elements used to uniquely and unambiguously identify a financial institution or a branch of a financial institution.
type BranchAndFinancialInstitutionIdentification5 struct {

	// Unique and unambiguous identification of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	FinancialInstitutionIdentification *FinancialInstitutionIdentification8 `xml:"FinInstnId"`

	// Identifies a specific branch of a financial institution.
	//
	// Usage: This component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	BranchIdentification *BranchData2 `xml:"BrnchId,omitempty"`
}

func (b *BranchAndFinancialInstitutionIdentification5) AddFinancialInstitutionIdentification() *FinancialInstitutionIdentification8 {
	b.FinancialInstitutionIdentification = new(FinancialInstitutionIdentification8)
	return b.FinancialInstitutionIdentification
}

func (b *BranchAndFinancialInstitutionIdentification5) AddBranchIdentification() *BranchData2 {
	b.BranchIdentification = new(BranchData2)
	return b.BranchIdentification
}
