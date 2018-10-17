package iso20022

// Information about a redemption multiple order.
type RedemptionMultipleOrderInstruction1 struct {

	// General information related to the order.
	MultipleOrderDetails *RedemptionMultipleOrder2 `xml:"MltplOrdrDtls"`

	// Information related to an intermediary.
	IntermediaryDetails []*Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Message is a copy.
	CopyDetails *CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionMultipleOrderInstruction1) AddMultipleOrderDetails() *RedemptionMultipleOrder2 {
	r.MultipleOrderDetails = new(RedemptionMultipleOrder2)
	return r.MultipleOrderDetails
}

func (r *RedemptionMultipleOrderInstruction1) AddIntermediaryDetails() *Intermediary4 {
	newValue := new(Intermediary4)
	r.IntermediaryDetails = append(r.IntermediaryDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleOrderInstruction1) AddCopyDetails() *CopyInformation1 {
	r.CopyDetails = new(CopyInformation1)
	return r.CopyDetails
}

func (r *RedemptionMultipleOrderInstruction1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}

// Information about a redemption multiple order.
type RedemptionMultipleOrderInstruction2 struct {

	// General information related to the order.
	MultipleOrderDetails *RedemptionMultipleOrder3 `xml:"MltplOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionMultipleOrderInstruction2) AddMultipleOrderDetails() *RedemptionMultipleOrder3 {
	r.MultipleOrderDetails = new(RedemptionMultipleOrder3)
	return r.MultipleOrderDetails
}

func (r *RedemptionMultipleOrderInstruction2) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleOrderInstruction2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
