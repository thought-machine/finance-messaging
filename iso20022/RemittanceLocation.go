package iso20022

// Remittance information that provides all remittance address elements, that enables the matching, i.e.  reconciliation, of a payment with the items that the transaction in intended to settle.
type RemittanceLocation1 struct {

	// Unique and unambiguous identification of the remittance information, e.g. a remittance advice, which is sent separately from the payment instruction.
	RemittanceIdentification *Max35Text `xml:"RmtId,omitempty"`

	// Specifies the method used to deliver the remittance advice information.
	RemittanceLocationMethod *RemittanceLocationMethod1Code `xml:"RmtLctnMtd,omitempty"`

	// Electronic address to which an agent is to send the remittance information.
	RemittanceLocationElectronicAddress *Max256Text `xml:"RmtLctnElctrncAdr,omitempty"`

	// Postal address to which an agent is to send the remittance information.
	RemittanceLocationPostalAddress *NameAndAddress3 `xml:"RmtLctnPstlAdr,omitempty"`
}

func (r *RemittanceLocation1) SetRemittanceIdentification(value string) {
	r.RemittanceIdentification = (*Max35Text)(&value)
}

func (r *RemittanceLocation1) SetRemittanceLocationMethod(value string) {
	r.RemittanceLocationMethod = (*RemittanceLocationMethod1Code)(&value)
}

func (r *RemittanceLocation1) SetRemittanceLocationElectronicAddress(value string) {
	r.RemittanceLocationElectronicAddress = (*Max256Text)(&value)
}

func (r *RemittanceLocation1) AddRemittanceLocationPostalAddress() *NameAndAddress3 {
	r.RemittanceLocationPostalAddress = new(NameAndAddress3)
	return r.RemittanceLocationPostalAddress
}

// Set of elements used to provide information on the remittance advice.
type RemittanceLocation2 struct {

	// Unique identification, as assigned by the initiating party, to unambiguously identify the remittance information sent separately from the payment instruction, such as a remittance advice.
	RemittanceIdentification *Max35Text `xml:"RmtId,omitempty"`

	// Method used to deliver the remittance advice information.
	RemittanceLocationMethod *RemittanceLocationMethod2Code `xml:"RmtLctnMtd,omitempty"`

	// Electronic address to which an agent is to send the remittance information.
	RemittanceLocationElectronicAddress *Max2048Text `xml:"RmtLctnElctrncAdr,omitempty"`

	// Postal address to which an agent is to send the remittance information.
	RemittanceLocationPostalAddress *NameAndAddress10 `xml:"RmtLctnPstlAdr,omitempty"`
}

func (r *RemittanceLocation2) SetRemittanceIdentification(value string) {
	r.RemittanceIdentification = (*Max35Text)(&value)
}

func (r *RemittanceLocation2) SetRemittanceLocationMethod(value string) {
	r.RemittanceLocationMethod = (*RemittanceLocationMethod2Code)(&value)
}

func (r *RemittanceLocation2) SetRemittanceLocationElectronicAddress(value string) {
	r.RemittanceLocationElectronicAddress = (*Max2048Text)(&value)
}

func (r *RemittanceLocation2) AddRemittanceLocationPostalAddress() *NameAndAddress10 {
	r.RemittanceLocationPostalAddress = new(NameAndAddress10)
	return r.RemittanceLocationPostalAddress
}

// Provides information on the remittance advice.
type RemittanceLocation3 struct {

	// Unique identification, as assigned by the initiating party, to unambiguously identify the remittance information sent separately from the payment instruction, such as a remittance advice.
	RemittanceIdentification *Max35Text `xml:"RmtId,omitempty"`

	// Set of elements used to provide information on the location and/or delivery of the remittance information.
	RemittanceLocationDetails []*RemittanceLocationDetails1 `xml:"RmtLctnDtls"`

	// Identifies the underlying transaction.
	References *TransactionReferences4 `xml:"Refs"`
}

func (r *RemittanceLocation3) SetRemittanceIdentification(value string) {
	r.RemittanceIdentification = (*Max35Text)(&value)
}

func (r *RemittanceLocation3) AddRemittanceLocationDetails() *RemittanceLocationDetails1 {
	newValue := new(RemittanceLocationDetails1)
	r.RemittanceLocationDetails = append(r.RemittanceLocationDetails, newValue)
	return newValue
}

func (r *RemittanceLocation3) AddReferences() *TransactionReferences4 {
	r.References = new(TransactionReferences4)
	return r.References
}

// Set of elements used to provide information on the remittance advice.
type RemittanceLocation4 struct {

	// Unique identification, as assigned by the initiating party, to unambiguously identify the remittance information sent separately from the payment instruction, such as a remittance advice.
	RemittanceIdentification *Max35Text `xml:"RmtId,omitempty"`

	// Set of elements used to provide information on the location and/or delivery of the remittance information.
	RemittanceLocationDetails []*RemittanceLocationDetails1 `xml:"RmtLctnDtls,omitempty"`
}

func (r *RemittanceLocation4) SetRemittanceIdentification(value string) {
	r.RemittanceIdentification = (*Max35Text)(&value)
}

func (r *RemittanceLocation4) AddRemittanceLocationDetails() *RemittanceLocationDetails1 {
	newValue := new(RemittanceLocationDetails1)
	r.RemittanceLocationDetails = append(r.RemittanceLocationDetails, newValue)
	return newValue
}
