package iso20022

// Identifies the mandate to be amended and gives details of the new mandate.
type MandateAmendment1 struct {

	// Set of elements used to provide information on the original messsage.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Set of elements used to provide detailed information on the amendment reason.
	AmendmentReason *AmendmentReasonInformation1 `xml:"AmdmntRsn"`

	// Set of elements used to provide the amended mandate data.
	Mandate *MandateInformation3 `xml:"Mndt"`

	// Set of elements used to provide the original mandate data.
	OriginalMandate *OriginalMandate1Choice `xml:"OrgnlMndt"`
}

func (m *MandateAmendment1) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAmendment1) AddAmendmentReason() *AmendmentReasonInformation1 {
	m.AmendmentReason = new(AmendmentReasonInformation1)
	return m.AmendmentReason
}

func (m *MandateAmendment1) AddMandate() *MandateInformation3 {
	m.Mandate = new(MandateInformation3)
	return m.Mandate
}

func (m *MandateAmendment1) AddOriginalMandate() *OriginalMandate1Choice {
	m.OriginalMandate = new(OriginalMandate1Choice)
	return m.OriginalMandate
}

// Identifies the mandate to be amended and gives details of the new mandate.
type MandateAmendment2 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the amendment reason.
	AmendmentReason *MandateAmendmentReason1 `xml:"AmdmntRsn"`

	// Provides the amended mandate data.
	Mandate *Mandate3 `xml:"Mndt"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate2Choice `xml:"OrgnlMndt"`
}

func (m *MandateAmendment2) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAmendment2) AddAmendmentReason() *MandateAmendmentReason1 {
	m.AmendmentReason = new(MandateAmendmentReason1)
	return m.AmendmentReason
}

func (m *MandateAmendment2) AddMandate() *Mandate3 {
	m.Mandate = new(Mandate3)
	return m.Mandate
}

func (m *MandateAmendment2) AddOriginalMandate() *OriginalMandate2Choice {
	m.OriginalMandate = new(OriginalMandate2Choice)
	return m.OriginalMandate
}

// Identifies the mandate to be amended and gives details of the new mandate.
type MandateAmendment3 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the amendment reason.
	AmendmentReason *MandateAmendmentReason1 `xml:"AmdmntRsn"`

	// Provides the amended mandate data.
	Mandate *Mandate3 `xml:"Mndt"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate2Choice `xml:"OrgnlMndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateAmendment3) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAmendment3) AddAmendmentReason() *MandateAmendmentReason1 {
	m.AmendmentReason = new(MandateAmendmentReason1)
	return m.AmendmentReason
}

func (m *MandateAmendment3) AddMandate() *Mandate3 {
	m.Mandate = new(Mandate3)
	return m.Mandate
}

func (m *MandateAmendment3) AddOriginalMandate() *OriginalMandate2Choice {
	m.OriginalMandate = new(OriginalMandate2Choice)
	return m.OriginalMandate
}

func (m *MandateAmendment3) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

// Identifies the mandate to be amended and gives details of the new mandate.
type MandateAmendment4 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the amendment reason.
	AmendmentReason *MandateAmendmentReason1 `xml:"AmdmntRsn"`

	// Provides the amended mandate data.
	Mandate *Mandate6 `xml:"Mndt"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate3Choice `xml:"OrgnlMndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateAmendment4) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAmendment4) AddAmendmentReason() *MandateAmendmentReason1 {
	m.AmendmentReason = new(MandateAmendmentReason1)
	return m.AmendmentReason
}

func (m *MandateAmendment4) AddMandate() *Mandate6 {
	m.Mandate = new(Mandate6)
	return m.Mandate
}

func (m *MandateAmendment4) AddOriginalMandate() *OriginalMandate3Choice {
	m.OriginalMandate = new(OriginalMandate3Choice)
	return m.OriginalMandate
}

func (m *MandateAmendment4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

// Identifies the mandate to be amended and gives details of the new mandate.
type MandateAmendment5 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the amendment reason.
	AmendmentReason *MandateAmendmentReason1 `xml:"AmdmntRsn"`

	// Provides the amended mandate data.
	Mandate *Mandate8 `xml:"Mndt"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate4Choice `xml:"OrgnlMndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateAmendment5) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAmendment5) AddAmendmentReason() *MandateAmendmentReason1 {
	m.AmendmentReason = new(MandateAmendmentReason1)
	return m.AmendmentReason
}

func (m *MandateAmendment5) AddMandate() *Mandate8 {
	m.Mandate = new(Mandate8)
	return m.Mandate
}

func (m *MandateAmendment5) AddOriginalMandate() *OriginalMandate4Choice {
	m.OriginalMandate = new(OriginalMandate4Choice)
	return m.OriginalMandate
}

func (m *MandateAmendment5) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
