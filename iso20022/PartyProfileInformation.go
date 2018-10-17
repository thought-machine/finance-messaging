package iso20022

// Information to support the Know Your Customer processes.
type PartyProfileInformation1 struct {

	// Indicates whether the certificate type has been obtained and verified.
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd"`

	// Identification of the person who validated the document.
	ValidatingParty *Max140Text `xml:"VldtngPty,omitempty"`

	// Identification of the person who checked the document.
	CheckingParty *Max140Text `xml:"ChckngPty,omitempty"`

	// Identification of the person who is responsible for the document.
	ResponsibleParty *Max140Text `xml:"RspnsblPty,omitempty"`

	// Identifies the type of certificate.
	CertificateType *CertificateType1Code `xml:"CertTp"`

	// Identifies the type of certificate.
	ExtendedCertificateType *Extended350Code `xml:"XtndedCertTp"`

	// Date at which the certification check has been performed.
	CheckingDate *ISODate `xml:"ChckngDt,omitempty"`

	// Specifies how frequently the check is performed.
	CheckingFrequency *EventFrequency1Code `xml:"ChckngFrqcy,omitempty"`

	// Specifies the date at which the next certification check will be performed.
	NextRevisionDate *ISODate `xml:"NxtRvsnDt,omitempty"`

	// Limits between which a person's salary is estimated.
	SalaryRange *Max35Text `xml:"SlryRg,omitempty"`

	// Indicates the main source of revenue.
	SourceOfWealth *Max140Text `xml:"SrcOfWlth,omitempty"`
}

func (p *PartyProfileInformation1) SetCertificationIndicator(value string) {
	p.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (p *PartyProfileInformation1) SetValidatingParty(value string) {
	p.ValidatingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation1) SetCheckingParty(value string) {
	p.CheckingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation1) SetResponsibleParty(value string) {
	p.ResponsibleParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation1) SetCertificateType(value string) {
	p.CertificateType = (*CertificateType1Code)(&value)
}

func (p *PartyProfileInformation1) SetExtendedCertificateType(value string) {
	p.ExtendedCertificateType = (*Extended350Code)(&value)
}

func (p *PartyProfileInformation1) SetCheckingDate(value string) {
	p.CheckingDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation1) SetCheckingFrequency(value string) {
	p.CheckingFrequency = (*EventFrequency1Code)(&value)
}

func (p *PartyProfileInformation1) SetNextRevisionDate(value string) {
	p.NextRevisionDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation1) SetSalaryRange(value string) {
	p.SalaryRange = (*Max35Text)(&value)
}

func (p *PartyProfileInformation1) SetSourceOfWealth(value string) {
	p.SourceOfWealth = (*Max140Text)(&value)
}

// Information to support the Know Your Customer processes.
type PartyProfileInformation2 struct {

	// Indicates whether the certificate type has been obtained and verified.
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd"`

	// Identification of the person who validated the document.
	ValidatingParty *Max140Text `xml:"VldtngPty,omitempty"`

	// Identification of the person who checked the document.
	CheckingParty *Max140Text `xml:"ChckngPty,omitempty"`

	// Identification of the person who is responsible for the document.
	ResponsibleParty *Max140Text `xml:"RspnsblPty,omitempty"`

	// Type of certificate.
	CertificateType *CertificationType1Choice `xml:"CertTp"`

	// Date at which the certification check has been performed.
	CheckingDate *ISODate `xml:"ChckngDt,omitempty"`

	// Specifies how frequently the check is performed.
	CheckingFrequency *EventFrequency1Code `xml:"ChckngFrqcy,omitempty"`

	// Specifies the date at which the next certification check will be performed.
	NextRevisionDate *ISODate `xml:"NxtRvsnDt,omitempty"`

	// Limits between which a person's salary is estimated.
	SalaryRange *Max35Text `xml:"SlryRg,omitempty"`

	// Indicates the main source of revenue.
	SourceOfWealth *Max140Text `xml:"SrcOfWlth,omitempty"`

	// Specifies an assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Specifies the customer’s money laundering risk.
	RiskLevel *RiskLevel1Choice `xml:"RskLvl,omitempty"`
}

func (p *PartyProfileInformation2) SetCertificationIndicator(value string) {
	p.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (p *PartyProfileInformation2) SetValidatingParty(value string) {
	p.ValidatingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation2) SetCheckingParty(value string) {
	p.CheckingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation2) SetResponsibleParty(value string) {
	p.ResponsibleParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation2) AddCertificateType() *CertificationType1Choice {
	p.CertificateType = new(CertificationType1Choice)
	return p.CertificateType
}

func (p *PartyProfileInformation2) SetCheckingDate(value string) {
	p.CheckingDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation2) SetCheckingFrequency(value string) {
	p.CheckingFrequency = (*EventFrequency1Code)(&value)
}

func (p *PartyProfileInformation2) SetNextRevisionDate(value string) {
	p.NextRevisionDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation2) SetSalaryRange(value string) {
	p.SalaryRange = (*Max35Text)(&value)
}

func (p *PartyProfileInformation2) SetSourceOfWealth(value string) {
	p.SourceOfWealth = (*Max140Text)(&value)
}

func (p *PartyProfileInformation2) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	p.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return p.CustomerConductClassification
}

func (p *PartyProfileInformation2) AddRiskLevel() *RiskLevel1Choice {
	p.RiskLevel = new(RiskLevel1Choice)
	return p.RiskLevel
}

// Information to support the Know Your Customer (KYC) processes.
type PartyProfileInformation3 struct {

	// Indicates whether the certificate type has been obtained and verified.
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Identification of the person who validated the document.
	ValidatingParty *Max140Text `xml:"VldtngPty,omitempty"`

	// Identification of the person who checked the document.
	CheckingParty *Max140Text `xml:"ChckngPty,omitempty"`

	// Identification of the person who is responsible for the document.
	ResponsibleParty *Max140Text `xml:"RspnsblPty,omitempty"`

	// Type of certificate.
	CertificateType *CertificationType1Choice `xml:"CertTp,omitempty"`

	// Date at which the certification check has been performed.
	CheckingDate *ISODate `xml:"ChckngDt,omitempty"`

	// Specifies how frequently the check is performed.
	CheckingFrequency *EventFrequency1Code `xml:"ChckngFrqcy,omitempty"`

	// Specifies the date at which the next certification check will be performed.
	NextRevisionDate *ISODate `xml:"NxtRvsnDt,omitempty"`

	// Limits between which a person's salary is estimated.
	SalaryRange *Max35Text `xml:"SlryRg,omitempty"`

	// Indicates the main source of revenue.
	SourceOfWealth *Max140Text `xml:"SrcOfWlth,omitempty"`

	// Specifies an assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Specifies the customer’s money laundering risk.
	RiskLevel *RiskLevel1Choice `xml:"RskLvl,omitempty"`

	// Specifies the type of due diligence checks carried out on the investor. For definitions of ordinary, simple and enhanced know your customer checks, local market regulations should be consulted.
	KnowYourCustomerCheckType *KYCCheckType1Choice `xml:"KnowYourCstmrChckTp,omitempty"`
}

func (p *PartyProfileInformation3) SetCertificationIndicator(value string) {
	p.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (p *PartyProfileInformation3) SetValidatingParty(value string) {
	p.ValidatingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation3) SetCheckingParty(value string) {
	p.CheckingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation3) SetResponsibleParty(value string) {
	p.ResponsibleParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation3) AddCertificateType() *CertificationType1Choice {
	p.CertificateType = new(CertificationType1Choice)
	return p.CertificateType
}

func (p *PartyProfileInformation3) SetCheckingDate(value string) {
	p.CheckingDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation3) SetCheckingFrequency(value string) {
	p.CheckingFrequency = (*EventFrequency1Code)(&value)
}

func (p *PartyProfileInformation3) SetNextRevisionDate(value string) {
	p.NextRevisionDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation3) SetSalaryRange(value string) {
	p.SalaryRange = (*Max35Text)(&value)
}

func (p *PartyProfileInformation3) SetSourceOfWealth(value string) {
	p.SourceOfWealth = (*Max140Text)(&value)
}

func (p *PartyProfileInformation3) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	p.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return p.CustomerConductClassification
}

func (p *PartyProfileInformation3) AddRiskLevel() *RiskLevel1Choice {
	p.RiskLevel = new(RiskLevel1Choice)
	return p.RiskLevel
}

func (p *PartyProfileInformation3) AddKnowYourCustomerCheckType() *KYCCheckType1Choice {
	p.KnowYourCustomerCheckType = new(KYCCheckType1Choice)
	return p.KnowYourCustomerCheckType
}

// Information to support the Know Your Customer (KYC) processes.
type PartyProfileInformation4 struct {

	// Indicates whether the certificate type has been obtained and verified.
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Identification of the person who validated the document.
	ValidatingParty *Max140Text `xml:"VldtngPty,omitempty"`

	// Identification of the person who checked the document.
	CheckingParty *Max140Text `xml:"ChckngPty,omitempty"`

	// Identification of the person who is responsible for the document.
	ResponsibleParty *Max140Text `xml:"RspnsblPty,omitempty"`

	// Type of certificate.
	CertificateType *CertificationType1Choice `xml:"CertTp,omitempty"`

	// Date at which the certification check has been performed.
	CheckingDate *ISODate `xml:"ChckngDt,omitempty"`

	// Specifies how frequently the check is performed.
	CheckingFrequency *EventFrequency1Code `xml:"ChckngFrqcy,omitempty"`

	// Specifies the date at which the next certification check will be performed.
	NextRevisionDate *ISODate `xml:"NxtRvsnDt,omitempty"`

	// Limits between which a person's salary is estimated.
	SalaryRange *Max35Text `xml:"SlryRg,omitempty"`

	// Indicates the main source of revenue.
	SourceOfWealth *Max140Text `xml:"SrcOfWlth,omitempty"`

	// Specifies an assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Specifies the customer’s money laundering risk.
	RiskLevel *RiskLevel1Choice `xml:"RskLvl,omitempty"`

	// Specifies the type of due diligence checks carried out on the investor. For definitions of ordinary, simple and enhanced know your customer checks, local market regulations should be consulted.
	KnowYourCustomerCheckType *KYCCheckType1Choice `xml:"KnowYourCstmrChckTp,omitempty"`

	// Specifies whether a customer has been checked in a Know Your Customer (KYC) database.
	KnowYourCustomerDatabaseCheck *DataBaseCheck1 `xml:"KnowYourCstmrDBChck,omitempty"`
}

func (p *PartyProfileInformation4) SetCertificationIndicator(value string) {
	p.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (p *PartyProfileInformation4) SetValidatingParty(value string) {
	p.ValidatingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation4) SetCheckingParty(value string) {
	p.CheckingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation4) SetResponsibleParty(value string) {
	p.ResponsibleParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation4) AddCertificateType() *CertificationType1Choice {
	p.CertificateType = new(CertificationType1Choice)
	return p.CertificateType
}

func (p *PartyProfileInformation4) SetCheckingDate(value string) {
	p.CheckingDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation4) SetCheckingFrequency(value string) {
	p.CheckingFrequency = (*EventFrequency1Code)(&value)
}

func (p *PartyProfileInformation4) SetNextRevisionDate(value string) {
	p.NextRevisionDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation4) SetSalaryRange(value string) {
	p.SalaryRange = (*Max35Text)(&value)
}

func (p *PartyProfileInformation4) SetSourceOfWealth(value string) {
	p.SourceOfWealth = (*Max140Text)(&value)
}

func (p *PartyProfileInformation4) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	p.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return p.CustomerConductClassification
}

func (p *PartyProfileInformation4) AddRiskLevel() *RiskLevel1Choice {
	p.RiskLevel = new(RiskLevel1Choice)
	return p.RiskLevel
}

func (p *PartyProfileInformation4) AddKnowYourCustomerCheckType() *KYCCheckType1Choice {
	p.KnowYourCustomerCheckType = new(KYCCheckType1Choice)
	return p.KnowYourCustomerCheckType
}

func (p *PartyProfileInformation4) AddKnowYourCustomerDatabaseCheck() *DataBaseCheck1 {
	p.KnowYourCustomerDatabaseCheck = new(DataBaseCheck1)
	return p.KnowYourCustomerDatabaseCheck
}

// Information to support the Know Your Customer (KYC) processes.
type PartyProfileInformation5 struct {

	// Indicates whether the certificate type has been obtained and verified.
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Identification of the person who validated the document.
	ValidatingParty *Max140Text `xml:"VldtngPty,omitempty"`

	// Identification of the person who checked the document.
	CheckingParty *Max140Text `xml:"ChckngPty,omitempty"`

	// Identification of the person who is responsible for the document.
	ResponsibleParty *Max140Text `xml:"RspnsblPty,omitempty"`

	// Type of certificate.
	CertificateType *CertificationType1Choice `xml:"CertTp,omitempty"`

	// Date at which the certification check has been performed.
	CheckingDate *ISODate `xml:"ChckngDt,omitempty"`

	// Specifies how frequently the check is performed.
	CheckingFrequency *EventFrequency1Code `xml:"ChckngFrqcy,omitempty"`

	// Specifies the date at which the next certification check will be performed.
	NextRevisionDate *ISODate `xml:"NxtRvsnDt,omitempty"`

	// Limits between which a person's salary is estimated.
	SalaryRange *Max35Text `xml:"SlryRg,omitempty"`

	// Indicates the main source of revenue.
	SourceOfWealth *Max140Text `xml:"SrcOfWlth,omitempty"`

	// Specifies an assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Specifies the customer’s money laundering risk.
	RiskLevel *RiskLevel2Choice `xml:"RskLvl,omitempty"`

	// Specifies the type of due diligence checks carried out on the investor or account owner. For definitions of ordinary, simple and enhanced know your customer checks, local market regulations should be consulted.
	KnowYourCustomerCheckType *KYCCheckType1Choice `xml:"KnowYourCstmrChckTp,omitempty"`

	// Specifies whether a customer has been checked in a Know Your Customer (KYC) database.
	KnowYourCustomerDatabaseCheck *DataBaseCheck1 `xml:"KnowYourCstmrDBChck,omitempty"`
}

func (p *PartyProfileInformation5) SetCertificationIndicator(value string) {
	p.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (p *PartyProfileInformation5) SetValidatingParty(value string) {
	p.ValidatingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation5) SetCheckingParty(value string) {
	p.CheckingParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation5) SetResponsibleParty(value string) {
	p.ResponsibleParty = (*Max140Text)(&value)
}

func (p *PartyProfileInformation5) AddCertificateType() *CertificationType1Choice {
	p.CertificateType = new(CertificationType1Choice)
	return p.CertificateType
}

func (p *PartyProfileInformation5) SetCheckingDate(value string) {
	p.CheckingDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation5) SetCheckingFrequency(value string) {
	p.CheckingFrequency = (*EventFrequency1Code)(&value)
}

func (p *PartyProfileInformation5) SetNextRevisionDate(value string) {
	p.NextRevisionDate = (*ISODate)(&value)
}

func (p *PartyProfileInformation5) SetSalaryRange(value string) {
	p.SalaryRange = (*Max35Text)(&value)
}

func (p *PartyProfileInformation5) SetSourceOfWealth(value string) {
	p.SourceOfWealth = (*Max140Text)(&value)
}

func (p *PartyProfileInformation5) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	p.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return p.CustomerConductClassification
}

func (p *PartyProfileInformation5) AddRiskLevel() *RiskLevel2Choice {
	p.RiskLevel = new(RiskLevel2Choice)
	return p.RiskLevel
}

func (p *PartyProfileInformation5) AddKnowYourCustomerCheckType() *KYCCheckType1Choice {
	p.KnowYourCustomerCheckType = new(KYCCheckType1Choice)
	return p.KnowYourCustomerCheckType
}

func (p *PartyProfileInformation5) AddKnowYourCustomerDatabaseCheck() *DataBaseCheck1 {
	p.KnowYourCustomerDatabaseCheck = new(DataBaseCheck1)
	return p.KnowYourCustomerDatabaseCheck
}
