package iso20022

// Reference of an order, deal reference, client reference and master reference.
type InvestmentFundOrderExecution1 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Indicates whether a confirmation amendment message will follow the confirmation cancellation instruction or not.
	AmendmentIndicator *YesNoIndicator `xml:"AmdmntInd"`

	// Reference of an order, client or deal reference.
	OrderReferences []*InvestmentFundOrderExecution2 `xml:"OrdrRefs"`
}

func (i *InvestmentFundOrderExecution1) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrderExecution1) SetAmendmentIndicator(value string) {
	i.AmendmentIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentFundOrderExecution1) AddOrderReferences() *InvestmentFundOrderExecution2 {
	newValue := new(InvestmentFundOrderExecution2)
	i.OrderReferences = append(i.OrderReferences, newValue)
	return newValue
}

// Reference of an order, client or deal reference.
type InvestmentFundOrderExecution2 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef,omitempty"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`
}

func (i *InvestmentFundOrderExecution2) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrderExecution2) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrderExecution2) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}
