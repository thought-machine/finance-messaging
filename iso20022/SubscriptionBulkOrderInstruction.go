package iso20022

// Information about a subscription bulk order.
type SubscriptionBulkOrderInstruction1 struct {

	// Common information related to all the orders to be cancelled.
	BulkOrderDetails *SubscriptionBulkOrder2 `xml:"BlkOrdrDtls"`

	// Information related to an intermediary.
	IntermediaryDetails []*Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Message is a copy.
	CopyDetails *CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionBulkOrderInstruction1) AddBulkOrderDetails() *SubscriptionBulkOrder2 {
	s.BulkOrderDetails = new(SubscriptionBulkOrder2)
	return s.BulkOrderDetails
}

func (s *SubscriptionBulkOrderInstruction1) AddIntermediaryDetails() *Intermediary4 {
	newValue := new(Intermediary4)
	s.IntermediaryDetails = append(s.IntermediaryDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderInstruction1) AddCopyDetails() *CopyInformation1 {
	s.CopyDetails = new(CopyInformation1)
	return s.CopyDetails
}

func (s *SubscriptionBulkOrderInstruction1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

// Information about a subscription bulk order.
type SubscriptionBulkOrderInstruction2 struct {

	// Common information related to all the orders to be cancelled.
	BulkOrderDetails *SubscriptionBulkOrder3 `xml:"BlkOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionBulkOrderInstruction2) AddBulkOrderDetails() *SubscriptionBulkOrder3 {
	s.BulkOrderDetails = new(SubscriptionBulkOrder3)
	return s.BulkOrderDetails
}

func (s *SubscriptionBulkOrderInstruction2) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderInstruction2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
