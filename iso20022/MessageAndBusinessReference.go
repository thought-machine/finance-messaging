package iso20022

// Information about the message reference of the message for which the status is requested and the business reference of the transfer instruction.
type MessageAndBusinessReference1 struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference2 `xml:"PrvsRef"`

	// Reference to a linked message sent in a proprietary way or the reference of a system.
	OtherReference *AdditionalReference2 `xml:"OthrRef"`

	// Investment account information of the transfer message for which the status is requested.
	InvestmentAccountDetails *InvestmentAccount10 `xml:"InvstmtAcctDtls,omitempty"`

	// Business reference of the transfer instruction message.
	TransferReference *Max35Text `xml:"TrfRef,omitempty"`
}

func (m *MessageAndBusinessReference1) AddPreviousReference() *AdditionalReference2 {
	m.PreviousReference = new(AdditionalReference2)
	return m.PreviousReference
}

func (m *MessageAndBusinessReference1) AddOtherReference() *AdditionalReference2 {
	m.OtherReference = new(AdditionalReference2)
	return m.OtherReference
}

func (m *MessageAndBusinessReference1) AddInvestmentAccountDetails() *InvestmentAccount10 {
	m.InvestmentAccountDetails = new(InvestmentAccount10)
	return m.InvestmentAccountDetails
}

func (m *MessageAndBusinessReference1) SetTransferReference(value string) {
	m.TransferReference = (*Max35Text)(&value)
}

// Information to identify funds order(s).
type MessageAndBusinessReference10 struct {

	// Reference to a linked message that was previously sent.
	Reference *References62Choice `xml:"Ref,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *AdditionalReference8 `xml:"RltdRef,omitempty"`

	// Unique and unambiguous identifiers of one or more individual order instructions or individual order cancellation requests.
	OrderReference []*InvestmentFundOrder8 `xml:"OrdrRef,omitempty"`
}

func (m *MessageAndBusinessReference10) AddReference() *References62Choice {
	m.Reference = new(References62Choice)
	return m.Reference
}

func (m *MessageAndBusinessReference10) AddRelatedReference() *AdditionalReference8 {
	m.RelatedReference = new(AdditionalReference8)
	return m.RelatedReference
}

func (m *MessageAndBusinessReference10) AddOrderReference() *InvestmentFundOrder8 {
	newValue := new(InvestmentFundOrder8)
	m.OrderReference = append(m.OrderReference, newValue)
	return newValue
}

// Information about the message reference of the message for which the status is requested and the business reference of the order.
type MessageAndBusinessReference2 struct {

	// Reference to a linked message sent in a proprietary way or reference of a system.
	OtherReference *AdditionalReference3 `xml:"OthrRef"`

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference3 `xml:"PrvsRef"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	IndividualOrderReference []*Max35Text `xml:"IndvOrdrRef,omitempty"`

	// Account information of the order message for which the status is requested.
	InvestmentAccount *InvestmentAccount13 `xml:"InvstmtAcct,omitempty"`
}

func (m *MessageAndBusinessReference2) AddOtherReference() *AdditionalReference3 {
	m.OtherReference = new(AdditionalReference3)
	return m.OtherReference
}

func (m *MessageAndBusinessReference2) AddPreviousReference() *AdditionalReference3 {
	m.PreviousReference = new(AdditionalReference3)
	return m.PreviousReference
}

func (m *MessageAndBusinessReference2) AddIndividualOrderReference(value string) {
	m.IndividualOrderReference = append(m.IndividualOrderReference, (*Max35Text)(&value))
}

func (m *MessageAndBusinessReference2) AddInvestmentAccount() *InvestmentAccount13 {
	m.InvestmentAccount = new(InvestmentAccount13)
	return m.InvestmentAccount
}

// Information to identify funds order(s).
type MessageAndBusinessReference4 struct {

	// Reference to a linked message sent in a proprietary way or reference of a system.
	OtherReference *AdditionalReference3 `xml:"OthrRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Unique and unambiguous identifiers of one or more indiviudal order instructions or individual order cancellation requests.
	OrderReference []*InvestmentFundOrder2 `xml:"OrdrRef,omitempty"`
}

func (m *MessageAndBusinessReference4) AddOtherReference() *AdditionalReference3 {
	m.OtherReference = new(AdditionalReference3)
	return m.OtherReference
}

func (m *MessageAndBusinessReference4) AddPreviousReference() *AdditionalReference3 {
	m.PreviousReference = new(AdditionalReference3)
	return m.PreviousReference
}

func (m *MessageAndBusinessReference4) AddRelatedReference() *AdditionalReference3 {
	m.RelatedReference = new(AdditionalReference3)
	return m.RelatedReference
}

func (m *MessageAndBusinessReference4) AddOrderReference() *InvestmentFundOrder2 {
	newValue := new(InvestmentFundOrder2)
	m.OrderReference = append(m.OrderReference, newValue)
	return newValue
}

// Information to identify the funds order confirmations.
type MessageAndBusinessReference5 struct {

	// Reference to a linked message sent in a proprietary way or reference of a system.
	OtherReference *AdditionalReference3 `xml:"OthrRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Unique and unambiguous identifiers of one or more individual order instructions.
	OrderReference []*InvestmentFundOrder3 `xml:"OrdrRef,omitempty"`
}

func (m *MessageAndBusinessReference5) AddOtherReference() *AdditionalReference3 {
	m.OtherReference = new(AdditionalReference3)
	return m.OtherReference
}

func (m *MessageAndBusinessReference5) AddPreviousReference() *AdditionalReference3 {
	m.PreviousReference = new(AdditionalReference3)
	return m.PreviousReference
}

func (m *MessageAndBusinessReference5) AddRelatedReference() *AdditionalReference3 {
	m.RelatedReference = new(AdditionalReference3)
	return m.RelatedReference
}

func (m *MessageAndBusinessReference5) AddOrderReference() *InvestmentFundOrder3 {
	newValue := new(InvestmentFundOrder3)
	m.OrderReference = append(m.OrderReference, newValue)
	return newValue
}

// Information about the message reference of the message for which the status is requested and the business reference of the transfer instruction.
type MessageAndBusinessReference6 struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message sent in a proprietary way or the reference of a system.
	OtherReference *AdditionalReference3 `xml:"OthrRef,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Investment account information of the transfer message for which the status is requested.
	InvestmentAccountDetails *InvestmentAccount22 `xml:"InvstmtAcctDtls,omitempty"`
}

func (m *MessageAndBusinessReference6) AddPreviousReference() *AdditionalReference3 {
	m.PreviousReference = new(AdditionalReference3)
	return m.PreviousReference
}

func (m *MessageAndBusinessReference6) AddOtherReference() *AdditionalReference3 {
	m.OtherReference = new(AdditionalReference3)
	return m.OtherReference
}

func (m *MessageAndBusinessReference6) SetMasterReference(value string) {
	m.MasterReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference6) SetTransferReference(value string) {
	m.TransferReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference6) SetClientReference(value string) {
	m.ClientReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference6) SetCancellationReference(value string) {
	m.CancellationReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference6) AddInvestmentAccountDetails() *InvestmentAccount22 {
	m.InvestmentAccountDetails = new(InvestmentAccount22)
	return m.InvestmentAccountDetails
}

// Information about the message reference of the message for which the status is requested and the business reference of the transfer instruction.
type MessageAndBusinessReference7 struct {

	// Reference to the message or communication that was previously sent.
	Reference *References39Choice `xml:"Ref,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Investment account information of the transfer message for which the status is requested.
	InvestmentAccountDetails *InvestmentAccount40 `xml:"InvstmtAcctDtls,omitempty"`
}

func (m *MessageAndBusinessReference7) AddReference() *References39Choice {
	m.Reference = new(References39Choice)
	return m.Reference
}

func (m *MessageAndBusinessReference7) SetMasterReference(value string) {
	m.MasterReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference7) SetTransferReference(value string) {
	m.TransferReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference7) SetClientReference(value string) {
	m.ClientReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference7) SetCancellationReference(value string) {
	m.CancellationReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference7) AddInvestmentAccountDetails() *InvestmentAccount40 {
	m.InvestmentAccountDetails = new(InvestmentAccount40)
	return m.InvestmentAccountDetails
}

// Information about the message reference of the message for which the status is requested and the business reference of the transfer instruction.
type MessageAndBusinessReference8 struct {

	// Reference to the message or communication that was previously sent.
	Reference *References48Choice `xml:"Ref,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Investment account information of the transfer message for which the status is requested.
	InvestmentAccountDetails *InvestmentAccount57 `xml:"InvstmtAcctDtls,omitempty"`
}

func (m *MessageAndBusinessReference8) AddReference() *References48Choice {
	m.Reference = new(References48Choice)
	return m.Reference
}

func (m *MessageAndBusinessReference8) SetMasterReference(value string) {
	m.MasterReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference8) SetTransferReference(value string) {
	m.TransferReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference8) AddClientReference() *AdditionalReference7 {
	m.ClientReference = new(AdditionalReference7)
	return m.ClientReference
}

func (m *MessageAndBusinessReference8) SetCancellationReference(value string) {
	m.CancellationReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference8) AddInvestmentAccountDetails() *InvestmentAccount57 {
	m.InvestmentAccountDetails = new(InvestmentAccount57)
	return m.InvestmentAccountDetails
}
