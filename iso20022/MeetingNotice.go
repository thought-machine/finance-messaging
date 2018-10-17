package iso20022

// Information about the shareholders meeting, specifying the participation requirements and the voting procedures. Alternatively, it may indicate where such information may be obtained.
type MeetingNotice2 struct {

	// Identification assigned to a general meeting by the party notifying the meeting. It must be unique for the party notifying the meeting.
	MeetingIdentification *Max35Text `xml:"MtgId,omitempty"`

	// Identification assigned to a meeting by the issuer. It must be unique for the issuer.
	IssuerMeetingIdentification *Max35Text `xml:"IssrMtgId,omitempty"`

	// Specifies the type of security holders meeting.
	Type *MeetingType2Code `xml:"Tp"`

	// Classifies the type of meeting.
	Classification *MeetingTypeClassification1Code `xml:"Clssfctn,omitempty"`

	// This code can be used in case another meeting classifications is required.
	ExtendedClassification *Extended350Code `xml:"XtndedClssfctn,omitempty"`

	// Official meeting announcement date.
	AnnouncementDate *ISODate `xml:"AnncmntDt,omitempty"`

	// Indicates whether physical participation to a meeting is required in order to be allowed to vote.
	AttendanceRequired *YesNoIndicator `xml:"AttndncReqrd"`

	// Indicates how to order the attendance card or to give notice of attendance.
	AttendanceConfirmationInformation *Max350Text `xml:"AttndncConfInf,omitempty"`

	// Date and time by which the beneficial owner or agent must notify of their intention to participate in a meeting. This deadline is set by an intermediary.
	AttendanceConfirmationDeadline *DateFormat2Choice `xml:"AttndncConfDdln,omitempty"`

	// Date and time by which the beneficial owner or agent must notify of their intention to participate in a meeting (STP mode). This deadline is set by an intermediary.
	AttendanceConfirmationSTPDeadline *DateFormat2Choice `xml:"AttndncConfSTPDdln,omitempty"`

	// Date and time by which the attendance to the meeting should be confirmed. This deadline is set by the issuer.
	AttendanceConfirmationMarketDeadline *DateFormat2Choice `xml:"AttndncConfMktDdln,omitempty"`

	// Address to use over the www (HTTP) service where addtional information on the meeting may be found.
	AdditionalDocumentationURLAddress *Max256Text `xml:"AddtlDcmnttnURLAdr,omitempty"`

	// Date and time by which security holders can propose amendments or new resolutions. This deadline is set by an intermediary.
	ResolutionProposalDeadline *DateFormat2Choice `xml:"RsltnPrpslDdln,omitempty"`

	// Date and time by which security holders can propose amendments or new resolutions. This deadline is set by the issuer.
	ResolutionProposalMarketDeadline *DateFormat2Choice `xml:"RsltnPrpslMktDdln,omitempty"`

	// Specifies the minimum stake in share capital or cash value or number of security holders required to table resolutions.
	ResolutionProposalThreshold *Max350Text `xml:"RsltnPrpslThrshld,omitempty"`

	// Specifies the minimum stake in share capital or cash value or number of security holders required to table resolutions. This minimum is expressed as a percentage.
	ResolutionProposalThresholdPercentage *PercentageRate `xml:"RsltnPrpslThrshldPctg,omitempty"`

	// Number of securities admitted to the vote, expressed as an amount and a currency.
	TotalNumberOfSecuritiesOutstanding *CurrencyAndAmount `xml:"TtlNbOfSctiesOutsdng,omitempty"`

	// Number of rights admitted to the vote.
	TotalNumberOfVotingRights *Number `xml:"TtlNbOfVtngRghts,omitempty"`

	// Address where the information on the proxy should be sent.
	ProxyAppointmentNotificationAddress *PostalAddress1 `xml:"PrxyAppntmntNtfctnAdr,omitempty"`

	// Indicates that no proxy is allowed for a meeting.
	ProxyNotAllowed *ProxyNotAllowedCode `xml:"PrxyNotAllwd,omitempty"`

	// Specifies the elements required to assign a proxy.
	Proxy *ProxyAppointmentInformation1 `xml:"Prxy,omitempty"`

	// Contact person at the party organising the meeting, at the issuer or at an intermediary.
	ContactPersonDetails []*MeetingContactPerson1 `xml:"CtctPrsnDtls,omitempty"`

	// Date on which a company publishes the results of its meeting.
	ResultPublicationDate *DateFormat3Choice `xml:"RsltPblctnDt,omitempty"`
}

func (m *MeetingNotice2) SetMeetingIdentification(value string) {
	m.MeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingNotice2) SetIssuerMeetingIdentification(value string) {
	m.IssuerMeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingNotice2) SetType(value string) {
	m.Type = (*MeetingType2Code)(&value)
}

func (m *MeetingNotice2) SetClassification(value string) {
	m.Classification = (*MeetingTypeClassification1Code)(&value)
}

func (m *MeetingNotice2) SetExtendedClassification(value string) {
	m.ExtendedClassification = (*Extended350Code)(&value)
}

func (m *MeetingNotice2) SetAnnouncementDate(value string) {
	m.AnnouncementDate = (*ISODate)(&value)
}

func (m *MeetingNotice2) SetAttendanceRequired(value string) {
	m.AttendanceRequired = (*YesNoIndicator)(&value)
}

func (m *MeetingNotice2) SetAttendanceConfirmationInformation(value string) {
	m.AttendanceConfirmationInformation = (*Max350Text)(&value)
}

func (m *MeetingNotice2) AddAttendanceConfirmationDeadline() *DateFormat2Choice {
	m.AttendanceConfirmationDeadline = new(DateFormat2Choice)
	return m.AttendanceConfirmationDeadline
}

func (m *MeetingNotice2) AddAttendanceConfirmationSTPDeadline() *DateFormat2Choice {
	m.AttendanceConfirmationSTPDeadline = new(DateFormat2Choice)
	return m.AttendanceConfirmationSTPDeadline
}

func (m *MeetingNotice2) AddAttendanceConfirmationMarketDeadline() *DateFormat2Choice {
	m.AttendanceConfirmationMarketDeadline = new(DateFormat2Choice)
	return m.AttendanceConfirmationMarketDeadline
}

func (m *MeetingNotice2) SetAdditionalDocumentationURLAddress(value string) {
	m.AdditionalDocumentationURLAddress = (*Max256Text)(&value)
}

func (m *MeetingNotice2) AddResolutionProposalDeadline() *DateFormat2Choice {
	m.ResolutionProposalDeadline = new(DateFormat2Choice)
	return m.ResolutionProposalDeadline
}

func (m *MeetingNotice2) AddResolutionProposalMarketDeadline() *DateFormat2Choice {
	m.ResolutionProposalMarketDeadline = new(DateFormat2Choice)
	return m.ResolutionProposalMarketDeadline
}

func (m *MeetingNotice2) SetResolutionProposalThreshold(value string) {
	m.ResolutionProposalThreshold = (*Max350Text)(&value)
}

func (m *MeetingNotice2) SetResolutionProposalThresholdPercentage(value string) {
	m.ResolutionProposalThresholdPercentage = (*PercentageRate)(&value)
}

func (m *MeetingNotice2) SetTotalNumberOfSecuritiesOutstanding(value, currency string) {
	m.TotalNumberOfSecuritiesOutstanding = NewCurrencyAndAmount(value, currency)
}

func (m *MeetingNotice2) SetTotalNumberOfVotingRights(value string) {
	m.TotalNumberOfVotingRights = (*Number)(&value)
}

func (m *MeetingNotice2) AddProxyAppointmentNotificationAddress() *PostalAddress1 {
	m.ProxyAppointmentNotificationAddress = new(PostalAddress1)
	return m.ProxyAppointmentNotificationAddress
}

func (m *MeetingNotice2) SetProxyNotAllowed(value string) {
	m.ProxyNotAllowed = (*ProxyNotAllowedCode)(&value)
}

func (m *MeetingNotice2) AddProxy() *ProxyAppointmentInformation1 {
	m.Proxy = new(ProxyAppointmentInformation1)
	return m.Proxy
}

func (m *MeetingNotice2) AddContactPersonDetails() *MeetingContactPerson1 {
	newValue := new(MeetingContactPerson1)
	m.ContactPersonDetails = append(m.ContactPersonDetails, newValue)
	return newValue
}

func (m *MeetingNotice2) AddResultPublicationDate() *DateFormat3Choice {
	m.ResultPublicationDate = new(DateFormat3Choice)
	return m.ResultPublicationDate
}

// Information about the shareholders meeting, specifying the participation requirements and the voting procedures. Alternatively, it may indicate where such information may be obtained.
type MeetingNotice3 struct {

	// Identification assigned to a general meeting by the party notifying the meeting. It must be unique for the party notifying the meeting.
	MeetingIdentification *Max35Text `xml:"MtgId,omitempty"`

	// Identification assigned to a meeting by the issuer. It must be unique for the issuer.
	IssuerMeetingIdentification *Max35Text `xml:"IssrMtgId,omitempty"`

	// Specifies the type of security holders meeting.
	Type *MeetingType2Code `xml:"Tp"`

	// Classifies the type of meeting.
	Classification *MeetingTypeClassification1Choice `xml:"Clssfctn,omitempty"`

	// Official meeting announcement date.
	AnnouncementDate *ISODate `xml:"AnncmntDt,omitempty"`

	// Indicates whether physical participation to a meeting is required in order to be allowed to vote.
	AttendanceRequired *YesNoIndicator `xml:"AttndncReqrd"`

	// Indicates how to order the attendance card or to give notice of attendance.
	AttendanceConfirmationInformation *Max350Text `xml:"AttndncConfInf,omitempty"`

	// Date and time by which the beneficial owner or agent must notify of their intention to participate in a meeting. This deadline is set by an intermediary.
	AttendanceConfirmationDeadline *DateFormat2Choice `xml:"AttndncConfDdln,omitempty"`

	// Date and time by which the beneficial owner or agent must notify of their intention to participate in a meeting (STP mode). This deadline is set by an intermediary.
	AttendanceConfirmationSTPDeadline *DateFormat2Choice `xml:"AttndncConfSTPDdln,omitempty"`

	// Date and time by which the attendance to the meeting should be confirmed. This deadline is set by the issuer.
	AttendanceConfirmationMarketDeadline *DateFormat2Choice `xml:"AttndncConfMktDdln,omitempty"`

	// Address to use over the www (HTTP) service where addtional information on the meeting may be found.
	AdditionalDocumentationURLAddress *Max256Text `xml:"AddtlDcmnttnURLAdr,omitempty"`

	// Additional procedural information about the general meeting, specifying the participation requirements and the voting procedures. Alternatively, it may indicate where such information may be obtained.
	AdditionalProcedureDetails []*AdditionalRights1 `xml:"AddtlPrcdrDtls,omitempty"`

	// Number of securities admitted to the vote, expressed as an amount and a currency.
	TotalNumberOfSecuritiesOutstanding *ActiveCurrencyAndAmount `xml:"TtlNbOfSctiesOutsdng,omitempty"`

	// Number of rights admitted to the vote.
	TotalNumberOfVotingRights *Number `xml:"TtlNbOfVtngRghts,omitempty"`

	// Address where the information on the proxy should be sent.
	ProxyAppointmentNotificationAddress *PostalAddress1 `xml:"PrxyAppntmntNtfctnAdr,omitempty"`

	// Choice to signalize whether proxy is allowed.
	ProxyChoice *Proxy1Choice `xml:"PrxyChc,omitempty"`

	// Contact person at the party organising the meeting, at the issuer or at an intermediary.
	ContactPersonDetails []*MeetingContactPerson1 `xml:"CtctPrsnDtls,omitempty"`

	// Date on which a company publishes the results of its meeting.
	ResultPublicationDate *DateFormat3Choice `xml:"RsltPblctnDt,omitempty"`
}

func (m *MeetingNotice3) SetMeetingIdentification(value string) {
	m.MeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingNotice3) SetIssuerMeetingIdentification(value string) {
	m.IssuerMeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingNotice3) SetType(value string) {
	m.Type = (*MeetingType2Code)(&value)
}

func (m *MeetingNotice3) AddClassification() *MeetingTypeClassification1Choice {
	m.Classification = new(MeetingTypeClassification1Choice)
	return m.Classification
}

func (m *MeetingNotice3) SetAnnouncementDate(value string) {
	m.AnnouncementDate = (*ISODate)(&value)
}

func (m *MeetingNotice3) SetAttendanceRequired(value string) {
	m.AttendanceRequired = (*YesNoIndicator)(&value)
}

func (m *MeetingNotice3) SetAttendanceConfirmationInformation(value string) {
	m.AttendanceConfirmationInformation = (*Max350Text)(&value)
}

func (m *MeetingNotice3) AddAttendanceConfirmationDeadline() *DateFormat2Choice {
	m.AttendanceConfirmationDeadline = new(DateFormat2Choice)
	return m.AttendanceConfirmationDeadline
}

func (m *MeetingNotice3) AddAttendanceConfirmationSTPDeadline() *DateFormat2Choice {
	m.AttendanceConfirmationSTPDeadline = new(DateFormat2Choice)
	return m.AttendanceConfirmationSTPDeadline
}

func (m *MeetingNotice3) AddAttendanceConfirmationMarketDeadline() *DateFormat2Choice {
	m.AttendanceConfirmationMarketDeadline = new(DateFormat2Choice)
	return m.AttendanceConfirmationMarketDeadline
}

func (m *MeetingNotice3) SetAdditionalDocumentationURLAddress(value string) {
	m.AdditionalDocumentationURLAddress = (*Max256Text)(&value)
}

func (m *MeetingNotice3) AddAdditionalProcedureDetails() *AdditionalRights1 {
	newValue := new(AdditionalRights1)
	m.AdditionalProcedureDetails = append(m.AdditionalProcedureDetails, newValue)
	return newValue
}

func (m *MeetingNotice3) SetTotalNumberOfSecuritiesOutstanding(value, currency string) {
	m.TotalNumberOfSecuritiesOutstanding = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MeetingNotice3) SetTotalNumberOfVotingRights(value string) {
	m.TotalNumberOfVotingRights = (*Number)(&value)
}

func (m *MeetingNotice3) AddProxyAppointmentNotificationAddress() *PostalAddress1 {
	m.ProxyAppointmentNotificationAddress = new(PostalAddress1)
	return m.ProxyAppointmentNotificationAddress
}

func (m *MeetingNotice3) AddProxyChoice() *Proxy1Choice {
	m.ProxyChoice = new(Proxy1Choice)
	return m.ProxyChoice
}

func (m *MeetingNotice3) AddContactPersonDetails() *MeetingContactPerson1 {
	newValue := new(MeetingContactPerson1)
	m.ContactPersonDetails = append(m.ContactPersonDetails, newValue)
	return newValue
}

func (m *MeetingNotice3) AddResultPublicationDate() *DateFormat3Choice {
	m.ResultPublicationDate = new(DateFormat3Choice)
	return m.ResultPublicationDate
}

// Information about the shareholders meeting, specifying the participation requirements and the voting procedures. Alternatively, it may indicate where such information may be obtained.
type MeetingNotice4 struct {

	// Identification assigned to the general meeting by the party notifying the meeting. It must be unique for the party notifying the meeting.
	MeetingIdentification *Max35Text `xml:"MtgId,omitempty"`

	// Identification assigned to the meeting by the issuer. It must be unique for the issuer.
	IssuerMeetingIdentification *Max35Text `xml:"IssrMtgId,omitempty"`

	// Specifies the type of security holders meeting.
	Type *MeetingType3Code `xml:"Tp"`

	// Classifies the type of meeting.
	Classification *MeetingTypeClassification1Choice `xml:"Clssfctn,omitempty"`

	// Official meeting announcement date.
	AnnouncementDate *ISODate `xml:"AnncmntDt,omitempty"`

	// Indicates whether physical participation to the meeting is required in order to be allowed to vote.
	AttendanceRequired *YesNoIndicator `xml:"AttndncReqrd,omitempty"`

	// Indicates how to order the attendance card or to give notice of attendance.
	AttendanceConfirmationInformation *Max350Text `xml:"AttndncConfInf,omitempty"`

	// Date and time by which the beneficial owner or agent must notify of its intention to participate in the meeting. This deadline is set by an intermediary.
	AttendanceConfirmationDeadline *DateFormat29Choice `xml:"AttndncConfDdln,omitempty"`

	// Date and time by which the beneficial owner or agent must notify of its intention to participate in the meeting (STP mode). This deadline is set by an intermediary.
	AttendanceConfirmationSTPDeadline *DateFormat29Choice `xml:"AttndncConfSTPDdln,omitempty"`

	// Date and time by which the attendance to the meeting should be confirmed. This deadline is set by the issuer.
	AttendanceConfirmationMarketDeadline *DateFormat29Choice `xml:"AttndncConfMktDdln,omitempty"`

	// Address to use over the www (HTTP) service where additional information on the meeting may be found.
	AdditionalDocumentationURLAddress *Max256Text `xml:"AddtlDcmnttnURLAdr,omitempty"`

	// Additional procedural information about the general meeting, specifying the participation requirements and the voting procedures. Alternatively, this may indicate where such information may be obtained.
	AdditionalProcedureDetails []*AdditionalRights2 `xml:"AddtlPrcdrDtls,omitempty"`

	// Number of securities admitted to the vote, expressed as an amount and a currency.
	TotalNumberOfSecuritiesOutstanding *UnitOrFaceAmount1Choice `xml:"TtlNbOfSctiesOutsdng,omitempty"`

	// Number of rights admitted to the vote.
	TotalNumberOfVotingRights *Number `xml:"TtlNbOfVtngRghts,omitempty"`

	// Address where the information on the proxy should be sent.
	ProxyAppointmentNotificationAddress *PostalAddress1 `xml:"PrxyAppntmntNtfctnAdr,omitempty"`

	// Indicates whether a proxy is allowed.
	ProxyChoice *Proxy2Choice `xml:"PrxyChc,omitempty"`

	// Contact person at the party organising the meeting, at the issuer or at an intermediary.
	ContactPersonDetails []*MeetingContactPerson2 `xml:"CtctPrsnDtls,omitempty"`

	// Date on which the company publishes the results of its meeting.
	ResultPublicationDate *DateFormat3Choice `xml:"RsltPblctnDt,omitempty"`
}

func (m *MeetingNotice4) SetMeetingIdentification(value string) {
	m.MeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingNotice4) SetIssuerMeetingIdentification(value string) {
	m.IssuerMeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingNotice4) SetType(value string) {
	m.Type = (*MeetingType3Code)(&value)
}

func (m *MeetingNotice4) AddClassification() *MeetingTypeClassification1Choice {
	m.Classification = new(MeetingTypeClassification1Choice)
	return m.Classification
}

func (m *MeetingNotice4) SetAnnouncementDate(value string) {
	m.AnnouncementDate = (*ISODate)(&value)
}

func (m *MeetingNotice4) SetAttendanceRequired(value string) {
	m.AttendanceRequired = (*YesNoIndicator)(&value)
}

func (m *MeetingNotice4) SetAttendanceConfirmationInformation(value string) {
	m.AttendanceConfirmationInformation = (*Max350Text)(&value)
}

func (m *MeetingNotice4) AddAttendanceConfirmationDeadline() *DateFormat29Choice {
	m.AttendanceConfirmationDeadline = new(DateFormat29Choice)
	return m.AttendanceConfirmationDeadline
}

func (m *MeetingNotice4) AddAttendanceConfirmationSTPDeadline() *DateFormat29Choice {
	m.AttendanceConfirmationSTPDeadline = new(DateFormat29Choice)
	return m.AttendanceConfirmationSTPDeadline
}

func (m *MeetingNotice4) AddAttendanceConfirmationMarketDeadline() *DateFormat29Choice {
	m.AttendanceConfirmationMarketDeadline = new(DateFormat29Choice)
	return m.AttendanceConfirmationMarketDeadline
}

func (m *MeetingNotice4) SetAdditionalDocumentationURLAddress(value string) {
	m.AdditionalDocumentationURLAddress = (*Max256Text)(&value)
}

func (m *MeetingNotice4) AddAdditionalProcedureDetails() *AdditionalRights2 {
	newValue := new(AdditionalRights2)
	m.AdditionalProcedureDetails = append(m.AdditionalProcedureDetails, newValue)
	return newValue
}

func (m *MeetingNotice4) AddTotalNumberOfSecuritiesOutstanding() *UnitOrFaceAmount1Choice {
	m.TotalNumberOfSecuritiesOutstanding = new(UnitOrFaceAmount1Choice)
	return m.TotalNumberOfSecuritiesOutstanding
}

func (m *MeetingNotice4) SetTotalNumberOfVotingRights(value string) {
	m.TotalNumberOfVotingRights = (*Number)(&value)
}

func (m *MeetingNotice4) AddProxyAppointmentNotificationAddress() *PostalAddress1 {
	m.ProxyAppointmentNotificationAddress = new(PostalAddress1)
	return m.ProxyAppointmentNotificationAddress
}

func (m *MeetingNotice4) AddProxyChoice() *Proxy2Choice {
	m.ProxyChoice = new(Proxy2Choice)
	return m.ProxyChoice
}

func (m *MeetingNotice4) AddContactPersonDetails() *MeetingContactPerson2 {
	newValue := new(MeetingContactPerson2)
	m.ContactPersonDetails = append(m.ContactPersonDetails, newValue)
	return newValue
}

func (m *MeetingNotice4) AddResultPublicationDate() *DateFormat3Choice {
	m.ResultPublicationDate = new(DateFormat3Choice)
	return m.ResultPublicationDate
}
