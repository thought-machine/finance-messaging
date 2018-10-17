package iso20022

// General characteristics related to a statement which reports information for a precise date.
type Report3 struct {

	// Sequential number of the report.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`

	// Gives the name and the reference of the query.
	QueryReference *QueryReference `xml:"QryRef,omitempty"`

	// Reference of the report.
	ReportIdentification *Max35Text `xml:"RptId,omitempty"`

	// Date of the statement.
	ReportDateTime *DateAndDateTime1Choice `xml:"RptDtTm"`

	// Specifies the regularity of an event.
	Frequency *Frequency4Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the report is complete or contains changes only.
	UpdateType *StatementUpdateTypeCodeAndDSSCodeChoice `xml:"UpdTp,omitempty"`

	// Notifies the type of report transmitted.
	NoticeType *GenericIdentification38 `xml:"NtceTp,omitempty"`
}

func (r *Report3) SetReportNumber(value string) {
	r.ReportNumber = (*Max5NumericText)(&value)
}

func (r *Report3) AddQueryReference() *QueryReference {
	r.QueryReference = new(QueryReference)
	return r.QueryReference
}

func (r *Report3) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *Report3) AddReportDateTime() *DateAndDateTime1Choice {
	r.ReportDateTime = new(DateAndDateTime1Choice)
	return r.ReportDateTime
}

func (r *Report3) AddFrequency() *Frequency4Choice {
	r.Frequency = new(Frequency4Choice)
	return r.Frequency
}

func (r *Report3) AddUpdateType() *StatementUpdateTypeCodeAndDSSCodeChoice {
	r.UpdateType = new(StatementUpdateTypeCodeAndDSSCodeChoice)
	return r.UpdateType
}

func (r *Report3) AddNoticeType() *GenericIdentification38 {
	r.NoticeType = new(GenericIdentification38)
	return r.NoticeType
}

// General characteristics of the report.
type Report4 struct {

	// Sequential number of the report.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Unique identification of the report.
	ReportIdentification *Max35Text `xml:"RptId,omitempty"`

	// Date and time of the report.
	ReportDateTime *DateAndDateTimeChoice `xml:"RptDtTm"`

	// Preparation date and time of the report.
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Previous report date and time.
	PreviousReportDateTime *DateAndDateTimeChoice `xml:"PrvsRptDtTm,omitempty"`

	// Specifies the frequency of the report.
	Frequency *Frequency8Choice `xml:"Frqcy"`

	// Specifies whether the report is complete or contains changes only.
	UpdateType *UpdateType4Choice `xml:"UpdTp"`

	// Specifies the type of balance on which the report is prepared.
	ReportBasis *StatementBasis6Choice `xml:"RptBsis"`

	// Period for which the report is given.
	ReportPeriod *DatePeriodDetails `xml:"RptPrd,omitempty"`

	// Specifies the source of the report.
	ReportSource *StatementSource1Choice `xml:"RptSrc,omitempty"`

	// Indicates whether the report is audited or not.
	AuditedIndicator *YesNoIndicator `xml:"AudtdInd,omitempty"`

	// Indicates whether there is activity or an information update reported in the report.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd,omitempty"`
}

func (r *Report4) SetReportNumber(value string) {
	r.ReportNumber = (*Max5NumericText)(&value)
}

func (r *Report4) SetQueryReference(value string) {
	r.QueryReference = (*Max35Text)(&value)
}

func (r *Report4) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *Report4) AddReportDateTime() *DateAndDateTimeChoice {
	r.ReportDateTime = new(DateAndDateTimeChoice)
	return r.ReportDateTime
}

func (r *Report4) AddCreationDateTime() *DateAndDateTimeChoice {
	r.CreationDateTime = new(DateAndDateTimeChoice)
	return r.CreationDateTime
}

func (r *Report4) AddPreviousReportDateTime() *DateAndDateTimeChoice {
	r.PreviousReportDateTime = new(DateAndDateTimeChoice)
	return r.PreviousReportDateTime
}

func (r *Report4) AddFrequency() *Frequency8Choice {
	r.Frequency = new(Frequency8Choice)
	return r.Frequency
}

func (r *Report4) AddUpdateType() *UpdateType4Choice {
	r.UpdateType = new(UpdateType4Choice)
	return r.UpdateType
}

func (r *Report4) AddReportBasis() *StatementBasis6Choice {
	r.ReportBasis = new(StatementBasis6Choice)
	return r.ReportBasis
}

func (r *Report4) AddReportPeriod() *DatePeriodDetails {
	r.ReportPeriod = new(DatePeriodDetails)
	return r.ReportPeriod
}

func (r *Report4) AddReportSource() *StatementSource1Choice {
	r.ReportSource = new(StatementSource1Choice)
	return r.ReportSource
}

func (r *Report4) SetAuditedIndicator(value string) {
	r.AuditedIndicator = (*YesNoIndicator)(&value)
}

func (r *Report4) SetActivityIndicator(value string) {
	r.ActivityIndicator = (*YesNoIndicator)(&value)
}

// Provides details on the settlement obligation report.
type Report5 struct {

	// Provides the identification for the non-clearing member. This is mandatory if the clearing member identification equals a general clearing member.
	NonClearingMember []*PartyIdentificationAndAccount31 `xml:"NonClrMmb,omitempty"`

	// Provides information about the settlement obligation details.
	SettlementObligationDetails []*SettlementObligation8 `xml:"SttlmOblgtnDtls"`
}

func (r *Report5) AddNonClearingMember() *PartyIdentificationAndAccount31 {
	newValue := new(PartyIdentificationAndAccount31)
	r.NonClearingMember = append(r.NonClearingMember, newValue)
	return newValue
}

func (r *Report5) AddSettlementObligationDetails() *SettlementObligation8 {
	newValue := new(SettlementObligation8)
	r.SettlementObligationDetails = append(r.SettlementObligationDetails, newValue)
	return newValue
}
