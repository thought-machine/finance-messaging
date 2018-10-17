package iso20022

// Characteristics of the statement.
type Statement11 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Period for the statement.
	StatementPeriod *Period2Choice `xml:"StmtPrd"`

	// Frequency of the statement.
	Frequency *Frequency4Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis2Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`
}

func (s *Statement11) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement11) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement11) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement11) AddStatementPeriod() *Period2Choice {
	s.StatementPeriod = new(Period2Choice)
	return s.StatementPeriod
}

func (s *Statement11) AddFrequency() *Frequency4Choice {
	s.Frequency = new(Frequency4Choice)
	return s.Frequency
}

func (s *Statement11) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement11) AddStatementBasis() *StatementBasis2Choice {
	s.StatementBasis = new(StatementBasis2Choice)
	return s.StatementBasis
}

func (s *Statement11) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement11) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}

// General characteristics related to a statement which reports information.
type Statement12 struct {

	// Indicates whether the statement contains missing instructions only or all instructions.
	StatementType *CorporateActionStatementType1Code `xml:"StmtTp"`

	// Indicates whether the statement report on account holdings for corporate action events is for single account/multiple events or multiple accounts/single event.
	ReportingType *CorporateActionStatementReportingType1Code `xml:"RptgTp"`

	// Reference of the statement.
	StatementIdentification *Max35Text `xml:"StmtId"`

	// Sequential number of the statement.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`

	// Date of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency4Choice `xml:"Frqcy"`

	// Indicates whether the report is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Period during which identification deadline has been set.
	NotificationDeadlinePeriod *DateOrDateTimePeriodChoice `xml:"NtfctnDdlnPrd,omitempty"`
}

func (s *Statement12) SetStatementType(value string) {
	s.StatementType = (*CorporateActionStatementType1Code)(&value)
}

func (s *Statement12) SetReportingType(value string) {
	s.ReportingType = (*CorporateActionStatementReportingType1Code)(&value)
}

func (s *Statement12) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement12) SetReportNumber(value string) {
	s.ReportNumber = (*Max5NumericText)(&value)
}

func (s *Statement12) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement12) AddFrequency() *Frequency4Choice {
	s.Frequency = new(Frequency4Choice)
	return s.Frequency
}

func (s *Statement12) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement12) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement12) AddNotificationDeadlinePeriod() *DateOrDateTimePeriodChoice {
	s.NotificationDeadlinePeriod = new(DateOrDateTimePeriodChoice)
	return s.NotificationDeadlinePeriod
}

// Characteristics of the statement.
type Statement14 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency4Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp,omitempty"`

	// Specifies whether the statement is sorted by status or transaction.
	StatementStructure *StatementStructure1Code `xml:"StmtStr"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement14) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement14) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement14) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement14) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement14) AddFrequency() *Frequency4Choice {
	s.Frequency = new(Frequency4Choice)
	return s.Frequency
}

func (s *Statement14) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement14) SetStatementStructure(value string) {
	s.StatementStructure = (*StatementStructure1Code)(&value)
}

func (s *Statement14) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement15 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Period for the statement.
	StatementPeriod *Period2Choice `xml:"StmtPrd"`

	// Frequency of the statement.
	Frequency *Frequency4Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement15) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement15) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement15) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement15) AddStatementPeriod() *Period2Choice {
	s.StatementPeriod = new(Period2Choice)
	return s.StatementPeriod
}

func (s *Statement15) AddFrequency() *Frequency4Choice {
	s.Frequency = new(Frequency4Choice)
	return s.Frequency
}

func (s *Statement15) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement15) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement16 struct {

	// Date or period of the statement.
	StatementDateOrPeriod *DateAndPeriod1Choice `xml:"StmtDtOrPrd,omitempty"`

	// Frequency of the statement.
	Frequency *Frequency4Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis3Choice `xml:"StmtBsis,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementType *StatementType2Choice `xml:"StmtTp,omitempty"`
}

func (s *Statement16) AddStatementDateOrPeriod() *DateAndPeriod1Choice {
	s.StatementDateOrPeriod = new(DateAndPeriod1Choice)
	return s.StatementDateOrPeriod
}

func (s *Statement16) AddFrequency() *Frequency4Choice {
	s.Frequency = new(Frequency4Choice)
	return s.Frequency
}

func (s *Statement16) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement16) AddStatementBasis() *StatementBasis3Choice {
	s.StatementBasis = new(StatementBasis3Choice)
	return s.StatementBasis
}

func (s *Statement16) AddStatementType() *StatementType2Choice {
	s.StatementType = new(StatementType2Choice)
	return s.StatementType
}

// Characteristics of the statement.
type Statement17 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency4Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement17) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement17) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement17) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement17) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement17) AddFrequency() *Frequency4Choice {
	s.Frequency = new(Frequency4Choice)
	return s.Frequency
}

func (s *Statement17) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement17) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement19 struct {

	// Identification assigned by the portfolio transfer counterpart to unambiguously identify a portfolio transfer notification.
	CounterpartyPortfolioTransferNotificationReference *Max35Text `xml:"CtrPtyPrtflTrfNtfctnRef,omitempty"`

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement19) SetCounterpartyPortfolioTransferNotificationReference(value string) {
	s.CounterpartyPortfolioTransferNotificationReference = (*Max35Text)(&value)
}

func (s *Statement19) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement19) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement19) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement19) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement19) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement20 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency4Choice `xml:"Frqcy"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis3Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement is audited or not.
	AuditedIndicator *YesNoIndicator `xml:"AudtdInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`

	// Indicates whether the statement contains tax lot details.
	TaxLotIndicator *YesNoIndicator `xml:"TaxLotInd,omitempty"`
}

func (s *Statement20) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement20) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement20) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement20) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement20) AddFrequency() *Frequency4Choice {
	s.Frequency = new(Frequency4Choice)
	return s.Frequency
}

func (s *Statement20) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement20) AddStatementBasis() *StatementBasis3Choice {
	s.StatementBasis = new(StatementBasis3Choice)
	return s.StatementBasis
}

func (s *Statement20) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement20) SetAuditedIndicator(value string) {
	s.AuditedIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement20) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement20) SetTaxLotIndicator(value string) {
	s.TaxLotIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement21 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency4Choice `xml:"Frqcy"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis3Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`
}

func (s *Statement21) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement21) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement21) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement21) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement21) AddFrequency() *Frequency4Choice {
	s.Frequency = new(Frequency4Choice)
	return s.Frequency
}

func (s *Statement21) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement21) AddStatementBasis() *StatementBasis3Choice {
	s.StatementBasis = new(StatementBasis3Choice)
	return s.StatementBasis
}

func (s *Statement21) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement21) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}

// General characteristics related to a statement which reports information for a precise date.
type Statement3 struct {

	// Reference of the statement.
	Reference *Max35Text `xml:"Ref"`

	// Date of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Preparation date of the statement
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Frequency of the statement.
	Frequency *FrequencyCodeAndDSSCodeChoice `xml:"Frqcy,omitempty"`

	// Indicates whether the report is complete or contains changes only.
	UpdateType *StatementUpdateTypeCodeAndDSSCodeChoice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasisCodeAndDSSCodeChoice `xml:"StmtBsis,omitempty"`

	// Sequential number of the statement.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`
}

func (s *Statement3) SetReference(value string) {
	s.Reference = (*Max35Text)(&value)
}

func (s *Statement3) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement3) AddCreationDateTime() *DateAndDateTimeChoice {
	s.CreationDateTime = new(DateAndDateTimeChoice)
	return s.CreationDateTime
}

func (s *Statement3) AddFrequency() *FrequencyCodeAndDSSCodeChoice {
	s.Frequency = new(FrequencyCodeAndDSSCodeChoice)
	return s.Frequency
}

func (s *Statement3) AddUpdateType() *StatementUpdateTypeCodeAndDSSCodeChoice {
	s.UpdateType = new(StatementUpdateTypeCodeAndDSSCodeChoice)
	return s.UpdateType
}

func (s *Statement3) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement3) AddStatementBasis() *StatementBasisCodeAndDSSCodeChoice {
	s.StatementBasis = new(StatementBasisCodeAndDSSCodeChoice)
	return s.StatementBasis
}

func (s *Statement3) SetReportNumber(value string) {
	s.ReportNumber = (*Max5NumericText)(&value)
}

// Provides statement details such as the account owner identification (ie, the clearing member identification) and optionaly the non clearing member identification, the clearing account or the list of trade legs.
type Statement31 struct {

	// Identification that is common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId"`

	// Date of the statement.
	StatementDateAndTime *DateAndDateTimeChoice `xml:"StmtDtAndTm"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *StatementUpdateType1Code `xml:"UpdTp"`

	// Frequency of the statement.
	Frequency *EventFrequency6Code `xml:"Frqcy"`

	// Sequential number of the statement.
	ReportNumber *Exact5NumericText `xml:"RptNb,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement31) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement31) AddStatementDateAndTime() *DateAndDateTimeChoice {
	s.StatementDateAndTime = new(DateAndDateTimeChoice)
	return s.StatementDateAndTime
}

func (s *Statement31) SetUpdateType(value string) {
	s.UpdateType = (*StatementUpdateType1Code)(&value)
}

func (s *Statement31) SetFrequency(value string) {
	s.Frequency = (*EventFrequency6Code)(&value)
}

func (s *Statement31) SetReportNumber(value string) {
	s.ReportNumber = (*Exact5NumericText)(&value)
}

func (s *Statement31) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement32 struct {

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Frequency of the statement.
	Frequency *Frequency1Code `xml:"Frqcy"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`
}

func (s *Statement32) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement32) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement32) SetFrequency(value string) {
	s.Frequency = (*Frequency1Code)(&value)
}

func (s *Statement32) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

// Characteristics of the statement.
type Statement33 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency9Choice `xml:"Frqcy"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis3Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`
}

func (s *Statement33) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement33) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement33) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement33) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement33) AddFrequency() *Frequency9Choice {
	s.Frequency = new(Frequency9Choice)
	return s.Frequency
}

func (s *Statement33) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement33) AddStatementBasis() *StatementBasis3Choice {
	s.StatementBasis = new(StatementBasis3Choice)
	return s.StatementBasis
}

func (s *Statement33) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement33) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement39 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency25Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType15Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement39) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement39) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement39) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement39) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement39) AddFrequency() *Frequency25Choice {
	s.Frequency = new(Frequency25Choice)
	return s.Frequency
}

func (s *Statement39) AddUpdateType() *UpdateType15Choice {
	s.UpdateType = new(UpdateType15Choice)
	return s.UpdateType
}

func (s *Statement39) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// General characteristics related to a statement which reports information for a precise date.
type Statement4 struct {

	// Reference of the statement.
	Reference *Max35Text `xml:"Ref"`

	// Date of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Preparation date of the statement
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Frequency of the statement.
	Frequency *FrequencyCodeAndDSSCodeChoice `xml:"Frqcy,omitempty"`

	// Indicates whether the report is complete or contains changes only.
	UpdateType *StatementUpdateTypeCodeAndDSSCodeChoice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasisCodeAndDSSCodeChoice `xml:"StmtBsis,omitempty"`

	// Sequential number of the statement.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`

	// Indicates whether the statement is audited.
	AuditedIndicator *YesNoIndicator `xml:"AudtdInd"`
}

func (s *Statement4) SetReference(value string) {
	s.Reference = (*Max35Text)(&value)
}

func (s *Statement4) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement4) AddCreationDateTime() *DateAndDateTimeChoice {
	s.CreationDateTime = new(DateAndDateTimeChoice)
	return s.CreationDateTime
}

func (s *Statement4) AddFrequency() *FrequencyCodeAndDSSCodeChoice {
	s.Frequency = new(FrequencyCodeAndDSSCodeChoice)
	return s.Frequency
}

func (s *Statement4) AddUpdateType() *StatementUpdateTypeCodeAndDSSCodeChoice {
	s.UpdateType = new(StatementUpdateTypeCodeAndDSSCodeChoice)
	return s.UpdateType
}

func (s *Statement4) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement4) AddStatementBasis() *StatementBasisCodeAndDSSCodeChoice {
	s.StatementBasis = new(StatementBasisCodeAndDSSCodeChoice)
	return s.StatementBasis
}

func (s *Statement4) SetReportNumber(value string) {
	s.ReportNumber = (*Max5NumericText)(&value)
}

func (s *Statement4) SetAuditedIndicator(value string) {
	s.AuditedIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement40 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency25Choice `xml:"Frqcy"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType15Choice `xml:"UpdTp"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis7Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement is audited or not.
	AuditedIndicator *YesNoIndicator `xml:"AudtdInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`

	// Indicates whether the statement contains tax lot details.
	TaxLotIndicator *YesNoIndicator `xml:"TaxLotInd,omitempty"`
}

func (s *Statement40) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement40) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement40) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement40) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement40) AddFrequency() *Frequency25Choice {
	s.Frequency = new(Frequency25Choice)
	return s.Frequency
}

func (s *Statement40) AddUpdateType() *UpdateType15Choice {
	s.UpdateType = new(UpdateType15Choice)
	return s.UpdateType
}

func (s *Statement40) AddStatementBasis() *StatementBasis7Choice {
	s.StatementBasis = new(StatementBasis7Choice)
	return s.StatementBasis
}

func (s *Statement40) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement40) SetAuditedIndicator(value string) {
	s.AuditedIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement40) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement40) SetTaxLotIndicator(value string) {
	s.TaxLotIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement41 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the statement query message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency25Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType15Choice `xml:"UpdTp,omitempty"`

	// Specifies whether the statement is sorted by status or transaction.
	StatementStructure *StatementStructure1Code `xml:"StmtStr"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement41) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement41) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement41) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement41) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement41) AddFrequency() *Frequency25Choice {
	s.Frequency = new(Frequency25Choice)
	return s.Frequency
}

func (s *Statement41) AddUpdateType() *UpdateType15Choice {
	s.UpdateType = new(UpdateType15Choice)
	return s.UpdateType
}

func (s *Statement41) SetStatementStructure(value string) {
	s.StatementStructure = (*StatementStructure1Code)(&value)
}

func (s *Statement41) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement42 struct {

	// Date or period of the statement.
	StatementDateOrPeriod *DateAndPeriod1Choice `xml:"StmtDtOrPrd,omitempty"`

	// Frequency of the statement.
	Frequency *Frequency25Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType15Choice `xml:"UpdTp,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis7Choice `xml:"StmtBsis,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementType *StatementType5Choice `xml:"StmtTp,omitempty"`
}

func (s *Statement42) AddStatementDateOrPeriod() *DateAndPeriod1Choice {
	s.StatementDateOrPeriod = new(DateAndPeriod1Choice)
	return s.StatementDateOrPeriod
}

func (s *Statement42) AddFrequency() *Frequency25Choice {
	s.Frequency = new(Frequency25Choice)
	return s.Frequency
}

func (s *Statement42) AddUpdateType() *UpdateType15Choice {
	s.UpdateType = new(UpdateType15Choice)
	return s.UpdateType
}

func (s *Statement42) AddStatementBasis() *StatementBasis7Choice {
	s.StatementBasis = new(StatementBasis7Choice)
	return s.StatementBasis
}

func (s *Statement42) AddStatementType() *StatementType5Choice {
	s.StatementType = new(StatementType5Choice)
	return s.StatementType
}

// Characteristics of the statement.
type Statement43 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Period for the statement.
	StatementPeriod *Period2Choice `xml:"StmtPrd"`

	// Frequency of the statement.
	Frequency *Frequency25Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType15Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement43) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement43) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement43) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement43) AddStatementPeriod() *Period2Choice {
	s.StatementPeriod = new(Period2Choice)
	return s.StatementPeriod
}

func (s *Statement43) AddFrequency() *Frequency25Choice {
	s.Frequency = new(Frequency25Choice)
	return s.Frequency
}

func (s *Statement43) AddUpdateType() *UpdateType15Choice {
	s.UpdateType = new(UpdateType15Choice)
	return s.UpdateType
}

func (s *Statement43) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement44 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Period for the statement.
	StatementPeriod *Period2Choice `xml:"StmtPrd"`

	// Frequency of the statement.
	Frequency *Frequency25Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType15Choice `xml:"UpdTp,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis8Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`
}

func (s *Statement44) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement44) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement44) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement44) AddStatementPeriod() *Period2Choice {
	s.StatementPeriod = new(Period2Choice)
	return s.StatementPeriod
}

func (s *Statement44) AddFrequency() *Frequency25Choice {
	s.Frequency = new(Frequency25Choice)
	return s.Frequency
}

func (s *Statement44) AddUpdateType() *UpdateType15Choice {
	s.UpdateType = new(UpdateType15Choice)
	return s.UpdateType
}

func (s *Statement44) AddStatementBasis() *StatementBasis8Choice {
	s.StatementBasis = new(StatementBasis8Choice)
	return s.StatementBasis
}

func (s *Statement44) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement44) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement45 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency22Choice `xml:"Frqcy"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType15Choice `xml:"UpdTp"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis7Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`
}

func (s *Statement45) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement45) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement45) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement45) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement45) AddFrequency() *Frequency22Choice {
	s.Frequency = new(Frequency22Choice)
	return s.Frequency
}

func (s *Statement45) AddUpdateType() *UpdateType15Choice {
	s.UpdateType = new(UpdateType15Choice)
	return s.UpdateType
}

func (s *Statement45) AddStatementBasis() *StatementBasis7Choice {
	s.StatementBasis = new(StatementBasis7Choice)
	return s.StatementBasis
}

func (s *Statement45) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement45) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement46 struct {

	// Identification assigned by the portfolio transfer counterpart to unambiguously identify a portfolio transfer notification.
	CounterpartyPortfolioTransferNotificationReference *Max35Text `xml:"CtrPtyPrtflTrfNtfctnRef,omitempty"`

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType15Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement46) SetCounterpartyPortfolioTransferNotificationReference(value string) {
	s.CounterpartyPortfolioTransferNotificationReference = (*Max35Text)(&value)
}

func (s *Statement46) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement46) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement46) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement46) AddUpdateType() *UpdateType15Choice {
	s.UpdateType = new(UpdateType15Choice)
	return s.UpdateType
}

func (s *Statement46) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// General characteristics related to a statement which reports information.
type Statement47 struct {

	// Indicates whether the statement contains missing instructions only or all instructions.
	StatementType *CorporateActionStatementType1Code `xml:"StmtTp"`

	// Indicates whether the statement report on account holdings for corporate action events is for single account/multiple events or multiple accounts/single event.
	ReportingType *CorporateActionStatementReportingType1Code `xml:"RptgTp"`

	// Reference of the statement.
	StatementIdentification *Max35Text `xml:"StmtId"`

	// Sequential number of the statement.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`

	// Date of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency25Choice `xml:"Frqcy"`

	// Indicates whether the report is complete or contains changes only.
	UpdateType *UpdateType15Choice `xml:"UpdTp"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Period during which identification deadline has been set.
	NotificationDeadlinePeriod *DateOrDateTimePeriodChoice `xml:"NtfctnDdlnPrd,omitempty"`
}

func (s *Statement47) SetStatementType(value string) {
	s.StatementType = (*CorporateActionStatementType1Code)(&value)
}

func (s *Statement47) SetReportingType(value string) {
	s.ReportingType = (*CorporateActionStatementReportingType1Code)(&value)
}

func (s *Statement47) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement47) SetReportNumber(value string) {
	s.ReportNumber = (*Max5NumericText)(&value)
}

func (s *Statement47) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement47) AddFrequency() *Frequency25Choice {
	s.Frequency = new(Frequency25Choice)
	return s.Frequency
}

func (s *Statement47) AddUpdateType() *UpdateType15Choice {
	s.UpdateType = new(UpdateType15Choice)
	return s.UpdateType
}

func (s *Statement47) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement47) AddNotificationDeadlinePeriod() *DateOrDateTimePeriodChoice {
	s.NotificationDeadlinePeriod = new(DateOrDateTimePeriodChoice)
	return s.NotificationDeadlinePeriod
}

// General characteristics related to a statement which reports information.
type Statement48 struct {

	// Indicates whether the statement contains missing instructions only or all instructions.
	StatementType *CorporateActionStatementType1Code `xml:"StmtTp"`

	// Indicates whether the statement report on account holdings for corporate action events is for single account/multiple events or multiple accounts/single event.
	ReportingType *CorporateActionStatementReportingType1Code `xml:"RptgTp"`

	// Reference of the statement.
	StatementIdentification *RestrictedFINXMax16Text `xml:"StmtId"`

	// Sequential number of the statement.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`

	// Date of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency26Choice `xml:"Frqcy"`

	// Indicates whether the report is complete or contains changes only.
	UpdateType *UpdateType16Choice `xml:"UpdTp"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Period during which identification deadline has been set.
	NotificationDeadlinePeriod *DateOrDateTimePeriodChoice `xml:"NtfctnDdlnPrd,omitempty"`
}

func (s *Statement48) SetStatementType(value string) {
	s.StatementType = (*CorporateActionStatementType1Code)(&value)
}

func (s *Statement48) SetReportingType(value string) {
	s.ReportingType = (*CorporateActionStatementReportingType1Code)(&value)
}

func (s *Statement48) SetStatementIdentification(value string) {
	s.StatementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement48) SetReportNumber(value string) {
	s.ReportNumber = (*Max5NumericText)(&value)
}

func (s *Statement48) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement48) AddFrequency() *Frequency26Choice {
	s.Frequency = new(Frequency26Choice)
	return s.Frequency
}

func (s *Statement48) AddUpdateType() *UpdateType16Choice {
	s.UpdateType = new(UpdateType16Choice)
	return s.UpdateType
}

func (s *Statement48) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement48) AddNotificationDeadlinePeriod() *DateOrDateTimePeriodChoice {
	s.NotificationDeadlinePeriod = new(DateOrDateTimePeriodChoice)
	return s.NotificationDeadlinePeriod
}

// Characteristics of the statement.
type Statement49 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *RestrictedFINXMax16Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *RestrictedFINXMax16Text `xml:"StmtId,omitempty"`

	// Period for the statement.
	StatementPeriod *Period2Choice `xml:"StmtPrd"`

	// Frequency of the statement.
	Frequency *Frequency26Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType16Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement49) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement49) SetQueryReference(value string) {
	s.QueryReference = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement49) SetStatementIdentification(value string) {
	s.StatementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement49) AddStatementPeriod() *Period2Choice {
	s.StatementPeriod = new(Period2Choice)
	return s.StatementPeriod
}

func (s *Statement49) AddFrequency() *Frequency26Choice {
	s.Frequency = new(Frequency26Choice)
	return s.Frequency
}

func (s *Statement49) AddUpdateType() *UpdateType16Choice {
	s.UpdateType = new(UpdateType16Choice)
	return s.UpdateType
}

func (s *Statement49) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// General characteristics related to a statement which reports information for a defined period.
type Statement5 struct {

	// Reference of the statement.
	Reference *Max35Text `xml:"Ref"`

	// Period on which the statement is reporting.
	StatementPeriod *DatePeriodDetails `xml:"StmtPrd"`

	// Creation date of the statement.
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Frequency of the statement.
	Frequency *Frequency1Code `xml:"Frqcy,omitempty"`

	// Specifies if the statement is complete or only contains changes.
	UpdateType *StatementUpdateTypeCode `xml:"UpdTp"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Sequential number of the statement.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`
}

func (s *Statement5) SetReference(value string) {
	s.Reference = (*Max35Text)(&value)
}

func (s *Statement5) AddStatementPeriod() *DatePeriodDetails {
	s.StatementPeriod = new(DatePeriodDetails)
	return s.StatementPeriod
}

func (s *Statement5) AddCreationDateTime() *DateAndDateTimeChoice {
	s.CreationDateTime = new(DateAndDateTimeChoice)
	return s.CreationDateTime
}

func (s *Statement5) SetFrequency(value string) {
	s.Frequency = (*Frequency1Code)(&value)
}

func (s *Statement5) SetUpdateType(value string) {
	s.UpdateType = (*StatementUpdateTypeCode)(&value)
}

func (s *Statement5) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement5) SetReportNumber(value string) {
	s.ReportNumber = (*Max5NumericText)(&value)
}

// Characteristics of the statement.
type Statement50 struct {

	// Identification assigned by the portfolio transfer counterpart to unambiguously identify a portfolio transfer notification.
	CounterpartyPortfolioTransferNotificationReference *RestrictedFINXMax16Text `xml:"CtrPtyPrtflTrfNtfctnRef,omitempty"`

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *RestrictedFINXMax16Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType16Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement50) SetCounterpartyPortfolioTransferNotificationReference(value string) {
	s.CounterpartyPortfolioTransferNotificationReference = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement50) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement50) SetStatementIdentification(value string) {
	s.StatementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement50) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement50) AddUpdateType() *UpdateType16Choice {
	s.UpdateType = new(UpdateType16Choice)
	return s.UpdateType
}

func (s *Statement50) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement51 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *RestrictedFINXMax16Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *RestrictedFINXMax16Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency26Choice `xml:"Frqcy"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType16Choice `xml:"UpdTp"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis9Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement is audited or not.
	AuditedIndicator *YesNoIndicator `xml:"AudtdInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`

	// Indicates whether the statement contains tax lot details.
	TaxLotIndicator *YesNoIndicator `xml:"TaxLotInd,omitempty"`
}

func (s *Statement51) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement51) SetQueryReference(value string) {
	s.QueryReference = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement51) SetStatementIdentification(value string) {
	s.StatementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement51) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement51) AddFrequency() *Frequency26Choice {
	s.Frequency = new(Frequency26Choice)
	return s.Frequency
}

func (s *Statement51) AddUpdateType() *UpdateType16Choice {
	s.UpdateType = new(UpdateType16Choice)
	return s.UpdateType
}

func (s *Statement51) AddStatementBasis() *StatementBasis9Choice {
	s.StatementBasis = new(StatementBasis9Choice)
	return s.StatementBasis
}

func (s *Statement51) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement51) SetAuditedIndicator(value string) {
	s.AuditedIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement51) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement51) SetTaxLotIndicator(value string) {
	s.TaxLotIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement52 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *RestrictedFINXMax16Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *RestrictedFINXMax16Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency34Choice `xml:"Frqcy"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType16Choice `xml:"UpdTp"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis9Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`
}

func (s *Statement52) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement52) SetQueryReference(value string) {
	s.QueryReference = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement52) SetStatementIdentification(value string) {
	s.StatementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement52) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement52) AddFrequency() *Frequency34Choice {
	s.Frequency = new(Frequency34Choice)
	return s.Frequency
}

func (s *Statement52) AddUpdateType() *UpdateType16Choice {
	s.UpdateType = new(UpdateType16Choice)
	return s.UpdateType
}

func (s *Statement52) AddStatementBasis() *StatementBasis9Choice {
	s.StatementBasis = new(StatementBasis9Choice)
	return s.StatementBasis
}

func (s *Statement52) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement52) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement53 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *RestrictedFINXMax16Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *RestrictedFINXMax16Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency26Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType16Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement53) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement53) SetQueryReference(value string) {
	s.QueryReference = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement53) SetStatementIdentification(value string) {
	s.StatementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement53) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement53) AddFrequency() *Frequency26Choice {
	s.Frequency = new(Frequency26Choice)
	return s.Frequency
}

func (s *Statement53) AddUpdateType() *UpdateType16Choice {
	s.UpdateType = new(UpdateType16Choice)
	return s.UpdateType
}

func (s *Statement53) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement54 struct {

	// Date or period of the statement.
	StatementDateOrPeriod *DateAndPeriod1Choice `xml:"StmtDtOrPrd,omitempty"`

	// Frequency of the statement.
	Frequency *Frequency26Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType16Choice `xml:"UpdTp,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis9Choice `xml:"StmtBsis,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementType *StatementType6Choice `xml:"StmtTp,omitempty"`
}

func (s *Statement54) AddStatementDateOrPeriod() *DateAndPeriod1Choice {
	s.StatementDateOrPeriod = new(DateAndPeriod1Choice)
	return s.StatementDateOrPeriod
}

func (s *Statement54) AddFrequency() *Frequency26Choice {
	s.Frequency = new(Frequency26Choice)
	return s.Frequency
}

func (s *Statement54) AddUpdateType() *UpdateType16Choice {
	s.UpdateType = new(UpdateType16Choice)
	return s.UpdateType
}

func (s *Statement54) AddStatementBasis() *StatementBasis9Choice {
	s.StatementBasis = new(StatementBasis9Choice)
	return s.StatementBasis
}

func (s *Statement54) AddStatementType() *StatementType6Choice {
	s.StatementType = new(StatementType6Choice)
	return s.StatementType
}

// Characteristics of the statement.
type Statement55 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the statement query message sent to request this statement.
	QueryReference *RestrictedFINXMax16Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *RestrictedFINXMax16Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency26Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType16Choice `xml:"UpdTp,omitempty"`

	// Specifies whether the statement is sorted by status or transaction.
	StatementStructure *StatementStructure1Code `xml:"StmtStr"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement55) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement55) SetQueryReference(value string) {
	s.QueryReference = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement55) SetStatementIdentification(value string) {
	s.StatementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement55) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement55) AddFrequency() *Frequency26Choice {
	s.Frequency = new(Frequency26Choice)
	return s.Frequency
}

func (s *Statement55) AddUpdateType() *UpdateType16Choice {
	s.UpdateType = new(UpdateType16Choice)
	return s.UpdateType
}

func (s *Statement55) SetStatementStructure(value string) {
	s.StatementStructure = (*StatementStructure1Code)(&value)
}

func (s *Statement55) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// Characteristics of the statement.
type Statement56 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *RestrictedFINXMax16Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *RestrictedFINXMax16Text `xml:"StmtId,omitempty"`

	// Period for the statement.
	StatementPeriod *Period2Choice `xml:"StmtPrd"`

	// Frequency of the statement.
	Frequency *Frequency26Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType16Choice `xml:"UpdTp,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis12Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`
}

func (s *Statement56) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement56) SetQueryReference(value string) {
	s.QueryReference = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement56) SetStatementIdentification(value string) {
	s.StatementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement56) AddStatementPeriod() *Period2Choice {
	s.StatementPeriod = new(Period2Choice)
	return s.StatementPeriod
}

func (s *Statement56) AddFrequency() *Frequency26Choice {
	s.Frequency = new(Frequency26Choice)
	return s.Frequency
}

func (s *Statement56) AddUpdateType() *UpdateType16Choice {
	s.UpdateType = new(UpdateType16Choice)
	return s.UpdateType
}

func (s *Statement56) AddStatementBasis() *StatementBasis12Choice {
	s.StatementBasis = new(StatementBasis12Choice)
	return s.StatementBasis
}

func (s *Statement56) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement56) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}

// General characteristics related to a statement which reports information for a precise date.
type Statement59 struct {

	// Specifies the business role of the message sender and, therefore, the business relationship between the sender and the receiver (or the interests represented by them, in those cases where another entity is acting on behalf of the sender or receiver). The message is exchanged between two entities, one being the account servicer and the other the account owner, and the message can be used with either one as the sender.
	SenderBusinessRole *SenderBusinessRole1Code `xml:"SndrBizRole"`

	// Sequential number of the report.
	StatementNumber *Number3Choice `xml:"StmtNb,omitempty"`

	// Identification of the query message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of the statement.
	StatementIdentification *Max35Text `xml:"StmtId"`

	// Date and time when the statement was created.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Date period for which the statement was created.
	StatementPeriod *DatePeriod1Choice `xml:"StmtPrd"`

	// Frequency of the statement.
	Frequency *Frequency22Choice `xml:"Frqcy,omitempty"`

	// Granularity of the frequency used for the reporting.
	FrequencyGranularity *FrequencyGranularityType1Code `xml:"FrqcyGrnlrty,omitempty"`

	// Specifies whether the statement is complete or contains changes only.
	UpdateType *UpdateType4Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or updated information reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement59) SetSenderBusinessRole(value string) {
	s.SenderBusinessRole = (*SenderBusinessRole1Code)(&value)
}

func (s *Statement59) AddStatementNumber() *Number3Choice {
	s.StatementNumber = new(Number3Choice)
	return s.StatementNumber
}

func (s *Statement59) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement59) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement59) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement59) AddStatementPeriod() *DatePeriod1Choice {
	s.StatementPeriod = new(DatePeriod1Choice)
	return s.StatementPeriod
}

func (s *Statement59) AddFrequency() *Frequency22Choice {
	s.Frequency = new(Frequency22Choice)
	return s.Frequency
}

func (s *Statement59) SetFrequencyGranularity(value string) {
	s.FrequencyGranularity = (*FrequencyGranularityType1Code)(&value)
}

func (s *Statement59) AddUpdateType() *UpdateType4Choice {
	s.UpdateType = new(UpdateType4Choice)
	return s.UpdateType
}

func (s *Statement59) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

// General characteristics related to a statement which reports information for a precise date.
type Statement6 struct {

	// Reference of the statement.
	Reference *Max35Text `xml:"Ref"`

	// Date of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Preparation date of the statement
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Frequency of the statement.
	Frequency *FrequencyCodeAndDSSCode1Choice `xml:"Frqcy"`

	// Indicates whether the report is complete or contains changes only.
	UpdateType *StatementUpdateTypeCodeAndDSSCodeChoice `xml:"UpdTp"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasisCodeAndDSSCodeChoice `xml:"StmtBsis"`

	// Sequential number of the statement.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`

	// Indicates whether the statement is audited.
	AuditedIndicator *YesNoIndicator `xml:"AudtdInd"`
}

func (s *Statement6) SetReference(value string) {
	s.Reference = (*Max35Text)(&value)
}

func (s *Statement6) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement6) AddCreationDateTime() *DateAndDateTimeChoice {
	s.CreationDateTime = new(DateAndDateTimeChoice)
	return s.CreationDateTime
}

func (s *Statement6) AddFrequency() *FrequencyCodeAndDSSCode1Choice {
	s.Frequency = new(FrequencyCodeAndDSSCode1Choice)
	return s.Frequency
}

func (s *Statement6) AddUpdateType() *StatementUpdateTypeCodeAndDSSCodeChoice {
	s.UpdateType = new(StatementUpdateTypeCodeAndDSSCodeChoice)
	return s.UpdateType
}

func (s *Statement6) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement6) AddStatementBasis() *StatementBasisCodeAndDSSCodeChoice {
	s.StatementBasis = new(StatementBasisCodeAndDSSCodeChoice)
	return s.StatementBasis
}

func (s *Statement6) SetReportNumber(value string) {
	s.ReportNumber = (*Max5NumericText)(&value)
}

func (s *Statement6) SetAuditedIndicator(value string) {
	s.AuditedIndicator = (*YesNoIndicator)(&value)
}

// General characteristics related to a statement which reports information for a precise date.
type Statement7 struct {

	// Reference of the statement.
	Reference *Max35Text `xml:"Ref"`

	// Date of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Preparation date of the statement
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Frequency of the statement.
	Frequency *FrequencyCodeAndDSSCode1Choice `xml:"Frqcy"`

	// Indicates whether the report is complete or contains changes only.
	UpdateType *StatementUpdateTypeCodeAndDSSCodeChoice `xml:"UpdTp"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasisCodeAndDSSCodeChoice `xml:"StmtBsis"`

	// Sequential number of the statement.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`
}

func (s *Statement7) SetReference(value string) {
	s.Reference = (*Max35Text)(&value)
}

func (s *Statement7) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement7) AddCreationDateTime() *DateAndDateTimeChoice {
	s.CreationDateTime = new(DateAndDateTimeChoice)
	return s.CreationDateTime
}

func (s *Statement7) AddFrequency() *FrequencyCodeAndDSSCode1Choice {
	s.Frequency = new(FrequencyCodeAndDSSCode1Choice)
	return s.Frequency
}

func (s *Statement7) AddUpdateType() *StatementUpdateTypeCodeAndDSSCodeChoice {
	s.UpdateType = new(StatementUpdateTypeCodeAndDSSCodeChoice)
	return s.UpdateType
}

func (s *Statement7) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement7) AddStatementBasis() *StatementBasisCodeAndDSSCodeChoice {
	s.StatementBasis = new(StatementBasisCodeAndDSSCodeChoice)
	return s.StatementBasis
}

func (s *Statement7) SetReportNumber(value string) {
	s.ReportNumber = (*Max5NumericText)(&value)
}

// General characteristics related to a statement which reports information for a defined period.
type Statement8 struct {

	// Reference of the statement.
	Reference *Max35Text `xml:"Ref"`

	// Period on which the statement is reporting.
	StatementPeriod *DatePeriodDetails `xml:"StmtPrd"`

	// Creation date of the statement.
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Frequency of the statement.
	Frequency *EventFrequency1Code `xml:"Frqcy,omitempty"`

	// Specifies if the statement is complete or only contains changes.
	UpdateType *StatementUpdateTypeCode `xml:"UpdTp"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Sequential number of the statement.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`
}

func (s *Statement8) SetReference(value string) {
	s.Reference = (*Max35Text)(&value)
}

func (s *Statement8) AddStatementPeriod() *DatePeriodDetails {
	s.StatementPeriod = new(DatePeriodDetails)
	return s.StatementPeriod
}

func (s *Statement8) AddCreationDateTime() *DateAndDateTimeChoice {
	s.CreationDateTime = new(DateAndDateTimeChoice)
	return s.CreationDateTime
}

func (s *Statement8) SetFrequency(value string) {
	s.Frequency = (*EventFrequency1Code)(&value)
}

func (s *Statement8) SetUpdateType(value string) {
	s.UpdateType = (*StatementUpdateTypeCode)(&value)
}

func (s *Statement8) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement8) SetReportNumber(value string) {
	s.ReportNumber = (*Max5NumericText)(&value)
}
