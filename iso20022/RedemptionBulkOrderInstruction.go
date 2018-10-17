package iso20022

// Information about a redemption bulk order.
type RedemptionBulkOrderInstruction1 struct {

	// Common information related to all the orders.
	BulkOrderDetails *RedemptionBulkOrder2 `xml:"BlkOrdrDtls"`

	// Information related to an intermediary.
	IntermediaryDetails []*Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Message is a copy.
	CopyDetails *CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionBulkOrderInstruction1) AddBulkOrderDetails() *RedemptionBulkOrder2 {
	r.BulkOrderDetails = new(RedemptionBulkOrder2)
	return r.BulkOrderDetails
}

func (r *RedemptionBulkOrderInstruction1) AddIntermediaryDetails() *Intermediary4 {
	newValue := new(Intermediary4)
	r.IntermediaryDetails = append(r.IntermediaryDetails, newValue)
	return newValue
}

func (r *RedemptionBulkOrderInstruction1) AddCopyDetails() *CopyInformation1 {
	r.CopyDetails = new(CopyInformation1)
	return r.CopyDetails
}

func (r *RedemptionBulkOrderInstruction1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}

// Information about a redemption bulk order.
type RedemptionBulkOrderInstruction2 struct {

	// Common information related to all the orders.
	BulkOrderDetails *RedemptionBulkOrder3 `xml:"BlkOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionBulkOrderInstruction2) AddBulkOrderDetails() *RedemptionBulkOrder3 {
	r.BulkOrderDetails = new(RedemptionBulkOrder3)
	return r.BulkOrderDetails
}

func (r *RedemptionBulkOrderInstruction2) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionBulkOrderInstruction2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
