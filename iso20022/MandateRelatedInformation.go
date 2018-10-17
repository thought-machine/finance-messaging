package iso20022

// Set of elements used to provide further details related to a direct debit mandate signed between the creditor and the debtor.
type MandateRelatedInformation1 struct {

	// Reference of the direct debit mandate that has been signed between by the debtor and the creditor.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Date on which the direct debit mandate has been signed by the debtor.
	DateOfSignature *ISODate `xml:"DtOfSgntr,omitempty"`

	// Indicator notifying whether the underlying mandate is amended or not.
	AmendmentIndicator *TrueFalseIndicator `xml:"AmdmntInd,omitempty"`

	// List of direct debit mandate elements that have been modified.
	AmendmentInformationDetails *AmendmentInformationDetails1 `xml:"AmdmntInfDtls,omitempty"`

	// Additional security provisions, such as a digital signature, as provided by the debtor.
	ElectronicSignature *Max1025Text `xml:"ElctrncSgntr,omitempty"`

	// Date of the first collection of a direct debit as per the mandate.
	FirstCollectionDate *ISODate `xml:"FrstColltnDt,omitempty"`

	// Date of the final collection of a direct debit as per the mandate.
	FinalCollectionDate *ISODate `xml:"FnlColltnDt,omitempty"`

	// Regularity with which direct debit instructions are to be created and processed.
	Frequency *Frequency1Code `xml:"Frqcy,omitempty"`
}

func (m *MandateRelatedInformation1) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *MandateRelatedInformation1) SetDateOfSignature(value string) {
	m.DateOfSignature = (*ISODate)(&value)
}

func (m *MandateRelatedInformation1) SetAmendmentIndicator(value string) {
	m.AmendmentIndicator = (*TrueFalseIndicator)(&value)
}

func (m *MandateRelatedInformation1) AddAmendmentInformationDetails() *AmendmentInformationDetails1 {
	m.AmendmentInformationDetails = new(AmendmentInformationDetails1)
	return m.AmendmentInformationDetails
}

func (m *MandateRelatedInformation1) SetElectronicSignature(value string) {
	m.ElectronicSignature = (*Max1025Text)(&value)
}

func (m *MandateRelatedInformation1) SetFirstCollectionDate(value string) {
	m.FirstCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation1) SetFinalCollectionDate(value string) {
	m.FinalCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation1) SetFrequency(value string) {
	m.Frequency = (*Frequency1Code)(&value)
}

// Provides further details related to a direct debit mandate signed between the creditor and the debtor.
type MandateRelatedInformation10 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Date on which the direct debit mandate has been signed by the debtor.
	DateOfSignature *ISODate `xml:"DtOfSgntr,omitempty"`

	// Indicator notifying whether the underlying mandate is amended or not.
	AmendmentIndicator *TrueFalseIndicator `xml:"AmdmntInd,omitempty"`

	// List of mandate elements that have been modified.
	AmendmentInformationDetails *AmendmentInformationDetails10 `xml:"AmdmntInfDtls,omitempty"`

	// Additional security provisions, such as a digital signature, as provided by the debtor.
	ElectronicSignature *Max1025Text `xml:"ElctrncSgntr,omitempty"`

	// Date of the first collection of a direct debit as per the mandate.
	FirstCollectionDate *ISODate `xml:"FrstColltnDt,omitempty"`

	// Date of the final collection of a direct debit as per the mandate.
	FinalCollectionDate *ISODate `xml:"FnlColltnDt,omitempty"`

	// Regularity with which direct debit instructions are to be created and processed.
	Frequency *Frequency21Choice `xml:"Frqcy,omitempty"`

	// Reason for the direct debit mandate to allow the user to distinguish between different mandates for the same creditor.
	Reason *MandateSetupReason1Choice `xml:"Rsn,omitempty"`
}

func (m *MandateRelatedInformation10) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *MandateRelatedInformation10) SetDateOfSignature(value string) {
	m.DateOfSignature = (*ISODate)(&value)
}

func (m *MandateRelatedInformation10) SetAmendmentIndicator(value string) {
	m.AmendmentIndicator = (*TrueFalseIndicator)(&value)
}

func (m *MandateRelatedInformation10) AddAmendmentInformationDetails() *AmendmentInformationDetails10 {
	m.AmendmentInformationDetails = new(AmendmentInformationDetails10)
	return m.AmendmentInformationDetails
}

func (m *MandateRelatedInformation10) SetElectronicSignature(value string) {
	m.ElectronicSignature = (*Max1025Text)(&value)
}

func (m *MandateRelatedInformation10) SetFirstCollectionDate(value string) {
	m.FirstCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation10) SetFinalCollectionDate(value string) {
	m.FinalCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation10) AddFrequency() *Frequency21Choice {
	m.Frequency = new(Frequency21Choice)
	return m.Frequency
}

func (m *MandateRelatedInformation10) AddReason() *MandateSetupReason1Choice {
	m.Reason = new(MandateSetupReason1Choice)
	return m.Reason
}

// Provides further details related to a direct debit mandate signed between the creditor and the debtor.
type MandateRelatedInformation11 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Date on which the direct debit mandate has been signed by the debtor.
	DateOfSignature *ISODate `xml:"DtOfSgntr,omitempty"`

	// Indicator notifying whether the underlying mandate is amended or not.
	AmendmentIndicator *TrueFalseIndicator `xml:"AmdmntInd,omitempty"`

	// List of mandate elements that have been modified.
	AmendmentInformationDetails *AmendmentInformationDetails11 `xml:"AmdmntInfDtls,omitempty"`

	// Additional security provisions, such as a digital signature, as provided by the debtor.
	ElectronicSignature *Max1025Text `xml:"ElctrncSgntr,omitempty"`

	// Date of the first collection of a direct debit as per the mandate.
	FirstCollectionDate *ISODate `xml:"FrstColltnDt,omitempty"`

	// Date of the final collection of a direct debit as per the mandate.
	FinalCollectionDate *ISODate `xml:"FnlColltnDt,omitempty"`

	// Regularity with which direct debit instructions are to be created and processed.
	Frequency *Frequency36Choice `xml:"Frqcy,omitempty"`

	// Reason for the direct debit mandate to allow the user to distinguish between different mandates for the same creditor.
	Reason *MandateSetupReason1Choice `xml:"Rsn,omitempty"`

	// Specifies the number of days the direct debit instruction must be tracked.
	TrackingDays *Exact2NumericText `xml:"TrckgDays,omitempty"`
}

func (m *MandateRelatedInformation11) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *MandateRelatedInformation11) SetDateOfSignature(value string) {
	m.DateOfSignature = (*ISODate)(&value)
}

func (m *MandateRelatedInformation11) SetAmendmentIndicator(value string) {
	m.AmendmentIndicator = (*TrueFalseIndicator)(&value)
}

func (m *MandateRelatedInformation11) AddAmendmentInformationDetails() *AmendmentInformationDetails11 {
	m.AmendmentInformationDetails = new(AmendmentInformationDetails11)
	return m.AmendmentInformationDetails
}

func (m *MandateRelatedInformation11) SetElectronicSignature(value string) {
	m.ElectronicSignature = (*Max1025Text)(&value)
}

func (m *MandateRelatedInformation11) SetFirstCollectionDate(value string) {
	m.FirstCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation11) SetFinalCollectionDate(value string) {
	m.FinalCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation11) AddFrequency() *Frequency36Choice {
	m.Frequency = new(Frequency36Choice)
	return m.Frequency
}

func (m *MandateRelatedInformation11) AddReason() *MandateSetupReason1Choice {
	m.Reason = new(MandateSetupReason1Choice)
	return m.Reason
}

func (m *MandateRelatedInformation11) SetTrackingDays(value string) {
	m.TrackingDays = (*Exact2NumericText)(&value)
}

// Set of elements used to provide further details related to a direct debit mandate signed between the creditor and the debtor.
type MandateRelatedInformation6 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Date on which the direct debit mandate has been signed by the debtor.
	DateOfSignature *ISODate `xml:"DtOfSgntr,omitempty"`

	// Indicator notifying whether the underlying mandate is amended or not.
	AmendmentIndicator *TrueFalseIndicator `xml:"AmdmntInd,omitempty"`

	// List of mandate elements that have been modified.
	AmendmentInformationDetails *AmendmentInformationDetails6 `xml:"AmdmntInfDtls,omitempty"`

	// Additional security provisions, such as a digital signature, as provided by the debtor.
	ElectronicSignature *Max1025Text `xml:"ElctrncSgntr,omitempty"`

	// Date of the first collection of a direct debit as per the mandate.
	FirstCollectionDate *ISODate `xml:"FrstColltnDt,omitempty"`

	// Date of the final collection of a direct debit as per the mandate.
	FinalCollectionDate *ISODate `xml:"FnlColltnDt,omitempty"`

	// Regularity with which direct debit instructions are to be created and processed.
	Frequency *Frequency1Code `xml:"Frqcy,omitempty"`
}

func (m *MandateRelatedInformation6) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *MandateRelatedInformation6) SetDateOfSignature(value string) {
	m.DateOfSignature = (*ISODate)(&value)
}

func (m *MandateRelatedInformation6) SetAmendmentIndicator(value string) {
	m.AmendmentIndicator = (*TrueFalseIndicator)(&value)
}

func (m *MandateRelatedInformation6) AddAmendmentInformationDetails() *AmendmentInformationDetails6 {
	m.AmendmentInformationDetails = new(AmendmentInformationDetails6)
	return m.AmendmentInformationDetails
}

func (m *MandateRelatedInformation6) SetElectronicSignature(value string) {
	m.ElectronicSignature = (*Max1025Text)(&value)
}

func (m *MandateRelatedInformation6) SetFirstCollectionDate(value string) {
	m.FirstCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation6) SetFinalCollectionDate(value string) {
	m.FinalCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation6) SetFrequency(value string) {
	m.Frequency = (*Frequency1Code)(&value)
}

// Set of elements used to provide further details related to a direct debit mandate signed between the creditor and the debtor.
type MandateRelatedInformation7 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Date on which the direct debit mandate has been signed by the debtor.
	DateOfSignature *ISODate `xml:"DtOfSgntr,omitempty"`

	// Indicator notifying whether the underlying mandate is amended or not.
	AmendmentIndicator *TrueFalseIndicator `xml:"AmdmntInd,omitempty"`

	// List of mandate elements that have been modified.
	AmendmentInformationDetails *AmendmentInformationDetails7 `xml:"AmdmntInfDtls,omitempty"`

	// Additional security provisions, such as a digital signature, as provided by the debtor.
	ElectronicSignature *Max1025Text `xml:"ElctrncSgntr,omitempty"`

	// Date of the first collection of a direct debit as per the mandate.
	FirstCollectionDate *ISODate `xml:"FrstColltnDt,omitempty"`

	// Date of the final collection of a direct debit as per the mandate.
	FinalCollectionDate *ISODate `xml:"FnlColltnDt,omitempty"`

	// Regularity with which direct debit instructions are to be created and processed.
	Frequency *Frequency1Code `xml:"Frqcy,omitempty"`
}

func (m *MandateRelatedInformation7) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *MandateRelatedInformation7) SetDateOfSignature(value string) {
	m.DateOfSignature = (*ISODate)(&value)
}

func (m *MandateRelatedInformation7) SetAmendmentIndicator(value string) {
	m.AmendmentIndicator = (*TrueFalseIndicator)(&value)
}

func (m *MandateRelatedInformation7) AddAmendmentInformationDetails() *AmendmentInformationDetails7 {
	m.AmendmentInformationDetails = new(AmendmentInformationDetails7)
	return m.AmendmentInformationDetails
}

func (m *MandateRelatedInformation7) SetElectronicSignature(value string) {
	m.ElectronicSignature = (*Max1025Text)(&value)
}

func (m *MandateRelatedInformation7) SetFirstCollectionDate(value string) {
	m.FirstCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation7) SetFinalCollectionDate(value string) {
	m.FinalCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation7) SetFrequency(value string) {
	m.Frequency = (*Frequency1Code)(&value)
}

// Provides further details related to a direct debit mandate signed between the creditor and the debtor.
type MandateRelatedInformation8 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Date on which the direct debit mandate has been signed by the debtor.
	DateOfSignature *ISODate `xml:"DtOfSgntr,omitempty"`

	// Indicator notifying whether the underlying mandate is amended or not.
	AmendmentIndicator *TrueFalseIndicator `xml:"AmdmntInd,omitempty"`

	// List of mandate elements that have been modified.
	AmendmentInformationDetails *AmendmentInformationDetails8 `xml:"AmdmntInfDtls,omitempty"`

	// Additional security provisions, such as a digital signature, as provided by the debtor.
	ElectronicSignature *Max1025Text `xml:"ElctrncSgntr,omitempty"`

	// Date of the first collection of a direct debit as per the mandate.
	FirstCollectionDate *ISODate `xml:"FrstColltnDt,omitempty"`

	// Date of the final collection of a direct debit as per the mandate.
	FinalCollectionDate *ISODate `xml:"FnlColltnDt,omitempty"`

	// Regularity with which direct debit instructions are to be created and processed.
	Frequency *Frequency6Code `xml:"Frqcy,omitempty"`
}

func (m *MandateRelatedInformation8) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *MandateRelatedInformation8) SetDateOfSignature(value string) {
	m.DateOfSignature = (*ISODate)(&value)
}

func (m *MandateRelatedInformation8) SetAmendmentIndicator(value string) {
	m.AmendmentIndicator = (*TrueFalseIndicator)(&value)
}

func (m *MandateRelatedInformation8) AddAmendmentInformationDetails() *AmendmentInformationDetails8 {
	m.AmendmentInformationDetails = new(AmendmentInformationDetails8)
	return m.AmendmentInformationDetails
}

func (m *MandateRelatedInformation8) SetElectronicSignature(value string) {
	m.ElectronicSignature = (*Max1025Text)(&value)
}

func (m *MandateRelatedInformation8) SetFirstCollectionDate(value string) {
	m.FirstCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation8) SetFinalCollectionDate(value string) {
	m.FinalCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation8) SetFrequency(value string) {
	m.Frequency = (*Frequency6Code)(&value)
}

// Provides further details related to a direct debit mandate signed between the creditor and the debtor.
type MandateRelatedInformation9 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Date on which the direct debit mandate has been signed by the debtor.
	DateOfSignature *ISODate `xml:"DtOfSgntr,omitempty"`

	// Indicator notifying whether the underlying mandate is amended or not.
	AmendmentIndicator *TrueFalseIndicator `xml:"AmdmntInd,omitempty"`

	// List of mandate elements that have been modified.
	AmendmentInformationDetails *AmendmentInformationDetails9 `xml:"AmdmntInfDtls,omitempty"`

	// Additional security provisions, such as a digital signature, as provided by the debtor.
	ElectronicSignature *Max1025Text `xml:"ElctrncSgntr,omitempty"`

	// Date of the first collection of a direct debit as per the mandate.
	FirstCollectionDate *ISODate `xml:"FrstColltnDt,omitempty"`

	// Date of the final collection of a direct debit as per the mandate.
	FinalCollectionDate *ISODate `xml:"FnlColltnDt,omitempty"`

	// Regularity with which direct debit instructions are to be created and processed.
	Frequency *Frequency6Code `xml:"Frqcy,omitempty"`
}

func (m *MandateRelatedInformation9) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *MandateRelatedInformation9) SetDateOfSignature(value string) {
	m.DateOfSignature = (*ISODate)(&value)
}

func (m *MandateRelatedInformation9) SetAmendmentIndicator(value string) {
	m.AmendmentIndicator = (*TrueFalseIndicator)(&value)
}

func (m *MandateRelatedInformation9) AddAmendmentInformationDetails() *AmendmentInformationDetails9 {
	m.AmendmentInformationDetails = new(AmendmentInformationDetails9)
	return m.AmendmentInformationDetails
}

func (m *MandateRelatedInformation9) SetElectronicSignature(value string) {
	m.ElectronicSignature = (*Max1025Text)(&value)
}

func (m *MandateRelatedInformation9) SetFirstCollectionDate(value string) {
	m.FirstCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation9) SetFinalCollectionDate(value string) {
	m.FinalCollectionDate = (*ISODate)(&value)
}

func (m *MandateRelatedInformation9) SetFrequency(value string) {
	m.Frequency = (*Frequency6Code)(&value)
}
