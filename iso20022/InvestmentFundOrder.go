package iso20022

// Reference of an order, order cancellation and master reference.
type InvestmentFundOrder1 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Reference of an order and order cancellation.
	OrderReferences []*InvestmentFundOrder5 `xml:"OrdrRefs"`
}

func (i *InvestmentFundOrder1) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder1) AddOrderReferences() *InvestmentFundOrder5 {
	newValue := new(InvestmentFundOrder5)
	i.OrderReferences = append(i.OrderReferences, newValue)
	return newValue
}

// References of an order confirmation and an order confirmation cancellation.
type InvestmentFundOrder11 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order execution, as assigned by the confirming party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`

	// Unique and unambiguous identifier for the order confirmation cancellation, as assigned by the confirming party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Reason for the cancellation of the confirmation.
	CancellationReason *CancellationReason31Choice `xml:"CxlRsn,omitempty"`
}

func (i *InvestmentFundOrder11) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder11) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder11) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder11) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder11) AddCancellationReason() *CancellationReason31Choice {
	i.CancellationReason = new(CancellationReason31Choice)
	return i.CancellationReason
}

// Reference of an order and of an order cancellation.
type InvestmentFundOrder2 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`

	// Account information of the individual order instruction or individual order cancellation request for which the status is requested.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument information of the individual order or individual order cancellation request for which the status is requested.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls,omitempty"`
}

func (i *InvestmentFundOrder2) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder2) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder2) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder2) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder2) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder2) AddInvestmentAccountDetails() *InvestmentAccount13 {
	i.InvestmentAccountDetails = new(InvestmentAccount13)
	return i.InvestmentAccountDetails
}

func (i *InvestmentFundOrder2) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	i.FinancialInstrumentDetails = new(FinancialInstrument10)
	return i.FinancialInstrumentDetails
}

// Reference of an order.
type InvestmentFundOrder3 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`

	// Account information of the individual order confirmation which the status is requested.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument information of the individual order confirmation for which the status is requested.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls,omitempty"`
}

func (i *InvestmentFundOrder3) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder3) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder3) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder3) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder3) AddInvestmentAccountDetails() *InvestmentAccount13 {
	i.InvestmentAccountDetails = new(InvestmentAccount13)
	return i.InvestmentAccountDetails
}

func (i *InvestmentFundOrder3) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	i.FinancialInstrumentDetails = new(FinancialInstrument10)
	return i.FinancialInstrumentDetails
}

// Identifies an order linked to an account opening.
type InvestmentFundOrder4 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef,omitempty"`

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`
}

func (i *InvestmentFundOrder4) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder4) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

// Reference of an order and order cancellation.
type InvestmentFundOrder5 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`
}

func (i *InvestmentFundOrder5) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder5) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder5) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

// Reference of an order and of an order cancellation.
type InvestmentFundOrder8 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Unique and unambiguous identifier for the order execution, as assigned by the confirming party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`

	// Account information of the individual order instruction or individual order cancellation request for which the status is requested.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument information of the individual order or individual order cancellation request for which the status is requested.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls,omitempty"`
}

func (i *InvestmentFundOrder8) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder8) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder8) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder8) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder8) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder8) AddInvestmentAccountDetails() *InvestmentAccount58 {
	i.InvestmentAccountDetails = new(InvestmentAccount58)
	return i.InvestmentAccountDetails
}

func (i *InvestmentFundOrder8) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	i.FinancialInstrumentDetails = new(FinancialInstrument57)
	return i.FinancialInstrumentDetails
}

// References of an order and order cancellation.
type InvestmentFundOrder9 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Reason for the cancellation.
	CancellationReason *CancellationReason32Choice `xml:"CxlRsn,omitempty"`
}

func (i *InvestmentFundOrder9) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder9) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder9) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder9) AddCancellationReason() *CancellationReason32Choice {
	i.CancellationReason = new(CancellationReason32Choice)
	return i.CancellationReason
}
