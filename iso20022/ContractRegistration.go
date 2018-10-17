package iso20022

// Currency control contract registration details.
type ContractRegistration1 struct {

	// Unique and unambiguous identification of the contract registration.
	ContractRegistrationIdentification *Max35Text `xml:"CtrctRegnId"`

	// Party registering the currency control contract.
	ReportingParty *TradeParty2 `xml:"RptgPty"`

	// Agent which registers the currency control contract.
	RegistrationAgent *BranchAndFinancialInstitutionIdentification5 `xml:"RegnAgt"`

	// Details about the opening of the contract registration.
	ContractRegistrationOpening []*ContractRegistration2 `xml:"CtrctRegnOpng"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *ContractRegistration1) SetContractRegistrationIdentification(value string) {
	c.ContractRegistrationIdentification = (*Max35Text)(&value)
}

func (c *ContractRegistration1) AddReportingParty() *TradeParty2 {
	c.ReportingParty = new(TradeParty2)
	return c.ReportingParty
}

func (c *ContractRegistration1) AddRegistrationAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.RegistrationAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.RegistrationAgent
}

func (c *ContractRegistration1) AddContractRegistrationOpening() *ContractRegistration2 {
	newValue := new(ContractRegistration2)
	c.ContractRegistrationOpening = append(c.ContractRegistrationOpening, newValue)
	return newValue
}

func (c *ContractRegistration1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

// Document that a user must file with an authorized servicer for each contract that involves foreign currency transactions with non residents.
type ContractRegistration2 struct {

	// Unique and unambiguous identification of the registered contract opening.
	ContractRegistrationOpeningIdentification *Max35Text `xml:"CtrctRegnOpngId"`

	// Priority requested for the opening of the registered contract.
	Priority *Priority2Code `xml:"Prty"`

	// Details of the contract being registered.
	Contract *UnderlyingContract1Choice `xml:"Ctrct"`

	// contract balance on date of contract registration.
	ContractBalance []*ContractBalance1 `xml:"CtrctBal,omitempty"`

	// Type of the payment schedule provided in the contract.
	PaymentScheduleType *PaymentScheduleType1Choice `xml:"PmtSchdlTp,omitempty"`

	// Unique and unambiguous identification of a previous contract registration.
	PreviousRegistrationIdentification []*DocumentIdentification22 `xml:"PrvsRegnId,omitempty"`

	// Further details on the registered contract opening.
	AdditionalInformation *Max1025Text `xml:"AddtlInf,omitempty"`

	// Documents provided as attachments to the contract registration request.
	Attachment []*DocumentGeneralInformation3 `xml:"Attchmnt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *ContractRegistration2) SetContractRegistrationOpeningIdentification(value string) {
	c.ContractRegistrationOpeningIdentification = (*Max35Text)(&value)
}

func (c *ContractRegistration2) SetPriority(value string) {
	c.Priority = (*Priority2Code)(&value)
}

func (c *ContractRegistration2) AddContract() *UnderlyingContract1Choice {
	c.Contract = new(UnderlyingContract1Choice)
	return c.Contract
}

func (c *ContractRegistration2) AddContractBalance() *ContractBalance1 {
	newValue := new(ContractBalance1)
	c.ContractBalance = append(c.ContractBalance, newValue)
	return newValue
}

func (c *ContractRegistration2) AddPaymentScheduleType() *PaymentScheduleType1Choice {
	c.PaymentScheduleType = new(PaymentScheduleType1Choice)
	return c.PaymentScheduleType
}

func (c *ContractRegistration2) AddPreviousRegistrationIdentification() *DocumentIdentification22 {
	newValue := new(DocumentIdentification22)
	c.PreviousRegistrationIdentification = append(c.PreviousRegistrationIdentification, newValue)
	return newValue
}

func (c *ContractRegistration2) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max1025Text)(&value)
}

func (c *ContractRegistration2) AddAttachment() *DocumentGeneralInformation3 {
	newValue := new(DocumentGeneralInformation3)
	c.Attachment = append(c.Attachment, newValue)
	return newValue
}

func (c *ContractRegistration2) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
