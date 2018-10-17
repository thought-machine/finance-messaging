package iso20022

// Information about a subscription multiple order.
type SubscriptionMultipleOrderInstruction1 struct {

	// Common information related to all the orders to be cancelled.
	MultipleOrderDetails *SubscriptionMultipleOrder2 `xml:"MltplOrdrDtls"`

	// Information related to an intermediary.
	IntermediaryDetails []*Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Message is a copy.
	CopyDetails *CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionMultipleOrderInstruction1) AddMultipleOrderDetails() *SubscriptionMultipleOrder2 {
	s.MultipleOrderDetails = new(SubscriptionMultipleOrder2)
	return s.MultipleOrderDetails
}

func (s *SubscriptionMultipleOrderInstruction1) AddIntermediaryDetails() *Intermediary4 {
	newValue := new(Intermediary4)
	s.IntermediaryDetails = append(s.IntermediaryDetails, newValue)
	return newValue
}

func (s *SubscriptionMultipleOrderInstruction1) AddCopyDetails() *CopyInformation1 {
	s.CopyDetails = new(CopyInformation1)
	return s.CopyDetails
}

func (s *SubscriptionMultipleOrderInstruction1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

// Information about a subscription multiple order.
type SubscriptionMultipleOrderInstruction2 struct {

	// Common information related to all the orders to be cancelled.
	MultipleOrderDetails *SubscriptionMultipleOrder3 `xml:"MltplOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionMultipleOrderInstruction2) AddMultipleOrderDetails() *SubscriptionMultipleOrder3 {
	s.MultipleOrderDetails = new(SubscriptionMultipleOrder3)
	return s.MultipleOrderDetails
}

func (s *SubscriptionMultipleOrderInstruction2) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionMultipleOrderInstruction2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
