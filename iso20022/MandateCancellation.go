package iso20022

// Identifies the mandate to be cancelled.
type MandateCancellation1 struct {

	// Set of elements used to provide information on the original messsage.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Set of elements used to provide detailed information on the cancellation reason.
	CancellationReason *CancellationReasonInformation2 `xml:"CxlRsn"`

	// Set of elements used to provide the original mandate data.
	OriginalMandate *OriginalMandate1Choice `xml:"OrgnlMndt"`
}

func (m *MandateCancellation1) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateCancellation1) AddCancellationReason() *CancellationReasonInformation2 {
	m.CancellationReason = new(CancellationReasonInformation2)
	return m.CancellationReason
}

func (m *MandateCancellation1) AddOriginalMandate() *OriginalMandate1Choice {
	m.OriginalMandate = new(OriginalMandate1Choice)
	return m.OriginalMandate
}

// Identifies the mandate to be cancelled.
type MandateCancellation2 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReason *PaymentCancellationReason1 `xml:"CxlRsn"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate2Choice `xml:"OrgnlMndt"`
}

func (m *MandateCancellation2) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateCancellation2) AddCancellationReason() *PaymentCancellationReason1 {
	m.CancellationReason = new(PaymentCancellationReason1)
	return m.CancellationReason
}

func (m *MandateCancellation2) AddOriginalMandate() *OriginalMandate2Choice {
	m.OriginalMandate = new(OriginalMandate2Choice)
	return m.OriginalMandate
}

// Identifies the mandate to be cancelled.
type MandateCancellation3 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReason *PaymentCancellationReason1 `xml:"CxlRsn"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate2Choice `xml:"OrgnlMndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateCancellation3) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateCancellation3) AddCancellationReason() *PaymentCancellationReason1 {
	m.CancellationReason = new(PaymentCancellationReason1)
	return m.CancellationReason
}

func (m *MandateCancellation3) AddOriginalMandate() *OriginalMandate2Choice {
	m.OriginalMandate = new(OriginalMandate2Choice)
	return m.OriginalMandate
}

func (m *MandateCancellation3) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

// Identifies the mandate to be cancelled.
type MandateCancellation4 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReason *PaymentCancellationReason1 `xml:"CxlRsn"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate3Choice `xml:"OrgnlMndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateCancellation4) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateCancellation4) AddCancellationReason() *PaymentCancellationReason1 {
	m.CancellationReason = new(PaymentCancellationReason1)
	return m.CancellationReason
}

func (m *MandateCancellation4) AddOriginalMandate() *OriginalMandate3Choice {
	m.OriginalMandate = new(OriginalMandate3Choice)
	return m.OriginalMandate
}

func (m *MandateCancellation4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

// Identifies the mandate to be cancelled.
type MandateCancellation5 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReason *PaymentCancellationReason1 `xml:"CxlRsn"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate4Choice `xml:"OrgnlMndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateCancellation5) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateCancellation5) AddCancellationReason() *PaymentCancellationReason1 {
	m.CancellationReason = new(PaymentCancellationReason1)
	return m.CancellationReason
}

func (m *MandateCancellation5) AddOriginalMandate() *OriginalMandate4Choice {
	m.OriginalMandate = new(OriginalMandate4Choice)
	return m.OriginalMandate
}

func (m *MandateCancellation5) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
