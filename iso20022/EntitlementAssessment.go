package iso20022

// Specifies the parameters, such as dates, used to calculate the entitlement to vote at a general meeting.
type EntitlementAssessment1 struct {

	// Date by which the securities should be blocked. This deadline is set by an intermediary.
	SecuritiesBlockingDeadline *DateFormat2Choice `xml:"SctiesBlckgDdln,omitempty"`

	// Date by which the securities should be blocked. This deadline is set by the issuer. (STP mode)
	SecuritiesBlockingSTPDeadline *DateFormat2Choice `xml:"SctiesBlckgSTPDdln,omitempty"`

	// Date by which the securities should be blocked. This deadline is set by the issuer.
	SecuritiesBlockingMarketDeadline *DateFormat2Choice `xml:"SctiesBlckgMktDdln,omitempty"`

	// Date by which the blocking period for the securities should end.
	SecuritiesBlockingPeriodEndDate *ISODateTime `xml:"SctiesBlckgPrdEndDt,omitempty"`

	// Date at which the positions are struck to note which parties will receive the entitlement, e.g. record date, book close date...
	EntitlementFixingDate *DateFormat3Choice `xml:"EntitlmntFxgDt,omitempty"`

	// Date by which the securities have to be registered. This deadline is specified by an intermediary.
	RegistrationSecuritiesDeadline *DateFormat2Choice `xml:"RegnSctiesDdln,omitempty"`

	// Date by which the securities have to be registered. This deadline is specified by an intermediary (STP mode).
	RegistrationSecuritiesSTPDeadline *DateFormat2Choice `xml:"RegnSctiesSTPDdln,omitempty"`

	// Date by which the securities have to be registered. This deadline is set by the issuer.
	RegistrationSecuritiesMarketDeadline *DateFormat2Choice `xml:"RegnSctiesMktDdln,omitempty"`

	// Date by which the holder needs to register its intention to participate in the meeting process in order to be allowed to participate in the meeting event. This deadline is specified by an intermediary.
	RegistrationParticipationDeadline *DateFormat2Choice `xml:"RegnPrtcptnDdln,omitempty"`

	// Date by which the holder needs to register its intention to participate in the meeting process in order to be allowed to participate in the meeting event. This deadline is specified by an intermediary (STP mode).
	RegistrationParticipationSTPDeadline *DateFormat2Choice `xml:"RegnPrtcptnSTPDdln,omitempty"`

	// Date by which the holder needs to register its intention to participate in the meeting process in order to be allowed to participate in the meeting event. This deadline is set by the issuer.
	RegistrationParticipationMarketDeadline *DateFormat2Choice `xml:"RegnPrtcptnMktDdln,omitempty"`

	// Specifies the calculation method of the number of votes assigned to one security. This element should be used when the entitlement calculation rule is complex.
	EntitlementDescription *Max350Text `xml:"EntitlmntDesc,omitempty"`

	// Number of votes assigned to one security.
	EntitlementRatio *DecimalNumber `xml:"EntitlmntRatio,omitempty"`
}

func (e *EntitlementAssessment1) AddSecuritiesBlockingDeadline() *DateFormat2Choice {
	e.SecuritiesBlockingDeadline = new(DateFormat2Choice)
	return e.SecuritiesBlockingDeadline
}

func (e *EntitlementAssessment1) AddSecuritiesBlockingSTPDeadline() *DateFormat2Choice {
	e.SecuritiesBlockingSTPDeadline = new(DateFormat2Choice)
	return e.SecuritiesBlockingSTPDeadline
}

func (e *EntitlementAssessment1) AddSecuritiesBlockingMarketDeadline() *DateFormat2Choice {
	e.SecuritiesBlockingMarketDeadline = new(DateFormat2Choice)
	return e.SecuritiesBlockingMarketDeadline
}

func (e *EntitlementAssessment1) SetSecuritiesBlockingPeriodEndDate(value string) {
	e.SecuritiesBlockingPeriodEndDate = (*ISODateTime)(&value)
}

func (e *EntitlementAssessment1) AddEntitlementFixingDate() *DateFormat3Choice {
	e.EntitlementFixingDate = new(DateFormat3Choice)
	return e.EntitlementFixingDate
}

func (e *EntitlementAssessment1) AddRegistrationSecuritiesDeadline() *DateFormat2Choice {
	e.RegistrationSecuritiesDeadline = new(DateFormat2Choice)
	return e.RegistrationSecuritiesDeadline
}

func (e *EntitlementAssessment1) AddRegistrationSecuritiesSTPDeadline() *DateFormat2Choice {
	e.RegistrationSecuritiesSTPDeadline = new(DateFormat2Choice)
	return e.RegistrationSecuritiesSTPDeadline
}

func (e *EntitlementAssessment1) AddRegistrationSecuritiesMarketDeadline() *DateFormat2Choice {
	e.RegistrationSecuritiesMarketDeadline = new(DateFormat2Choice)
	return e.RegistrationSecuritiesMarketDeadline
}

func (e *EntitlementAssessment1) AddRegistrationParticipationDeadline() *DateFormat2Choice {
	e.RegistrationParticipationDeadline = new(DateFormat2Choice)
	return e.RegistrationParticipationDeadline
}

func (e *EntitlementAssessment1) AddRegistrationParticipationSTPDeadline() *DateFormat2Choice {
	e.RegistrationParticipationSTPDeadline = new(DateFormat2Choice)
	return e.RegistrationParticipationSTPDeadline
}

func (e *EntitlementAssessment1) AddRegistrationParticipationMarketDeadline() *DateFormat2Choice {
	e.RegistrationParticipationMarketDeadline = new(DateFormat2Choice)
	return e.RegistrationParticipationMarketDeadline
}

func (e *EntitlementAssessment1) SetEntitlementDescription(value string) {
	e.EntitlementDescription = (*Max350Text)(&value)
}

func (e *EntitlementAssessment1) SetEntitlementRatio(value string) {
	e.EntitlementRatio = (*DecimalNumber)(&value)
}

// Specifies the parameters, such as dates, used to calculate the entitlement to vote at a general meeting.
type EntitlementAssessment2 struct {

	// Date by which the securities should be blocked. This deadline is set by an intermediary.
	SecuritiesBlockingDeadline *DateFormat2Choice `xml:"SctiesBlckgDdln,omitempty"`

	// Date by which the securities should be blocked. This deadline is set by the issuer. (STP mode)
	SecuritiesBlockingSTPDeadline *DateFormat2Choice `xml:"SctiesBlckgSTPDdln,omitempty"`

	// Date by which the securities should be blocked. This deadline is set by the issuer.
	SecuritiesBlockingMarketDeadline *DateFormat2Choice `xml:"SctiesBlckgMktDdln,omitempty"`

	// Date by which the blocking period for the securities should end.
	SecuritiesBlockingPeriodEndDate *ISODateTime `xml:"SctiesBlckgPrdEndDt,omitempty"`

	// Date at which the positions are struck to note which parties will receive the entitlement, e.g. record date, book close date...
	EntitlementFixingDate *DateFormat3Choice `xml:"EntitlmntFxgDt,omitempty"`

	// Date by which the securities have to be registered. This deadline is specified by an intermediary.
	RegistrationSecuritiesDeadline *DateFormat2Choice `xml:"RegnSctiesDdln,omitempty"`

	// Date by which the securities have to be registered. This deadline is specified by an intermediary (STP mode).
	RegistrationSecuritiesSTPDeadline *DateFormat2Choice `xml:"RegnSctiesSTPDdln,omitempty"`

	// Date by which the securities have to be registered. This deadline is set by the issuer.
	RegistrationSecuritiesMarketDeadline *DateFormat2Choice `xml:"RegnSctiesMktDdln,omitempty"`

	// Date by which the holder needs to register its intention to participate in the meeting process in order to be allowed to participate in the meeting event. This deadline is specified by an intermediary.
	RegistrationParticipationDeadline *DateFormat2Choice `xml:"RegnPrtcptnDdln,omitempty"`

	// Date by which the holder needs to register its intention to participate in the meeting process in order to be allowed to participate in the meeting event. This deadline is specified by an intermediary (STP mode).
	RegistrationParticipationSTPDeadline *DateFormat2Choice `xml:"RegnPrtcptnSTPDdln,omitempty"`

	// Date by which the holder needs to register its intention to participate in the meeting process in order to be allowed to participate in the meeting event. This deadline is set by the issuer.
	RegistrationParticipationMarketDeadline *DateFormat2Choice `xml:"RegnPrtcptnMktDdln,omitempty"`

	// Number of votes assigned to one security.
	Entitlement *Entitlement1Choice `xml:"Entitlmnt,omitempty"`
}

func (e *EntitlementAssessment2) AddSecuritiesBlockingDeadline() *DateFormat2Choice {
	e.SecuritiesBlockingDeadline = new(DateFormat2Choice)
	return e.SecuritiesBlockingDeadline
}

func (e *EntitlementAssessment2) AddSecuritiesBlockingSTPDeadline() *DateFormat2Choice {
	e.SecuritiesBlockingSTPDeadline = new(DateFormat2Choice)
	return e.SecuritiesBlockingSTPDeadline
}

func (e *EntitlementAssessment2) AddSecuritiesBlockingMarketDeadline() *DateFormat2Choice {
	e.SecuritiesBlockingMarketDeadline = new(DateFormat2Choice)
	return e.SecuritiesBlockingMarketDeadline
}

func (e *EntitlementAssessment2) SetSecuritiesBlockingPeriodEndDate(value string) {
	e.SecuritiesBlockingPeriodEndDate = (*ISODateTime)(&value)
}

func (e *EntitlementAssessment2) AddEntitlementFixingDate() *DateFormat3Choice {
	e.EntitlementFixingDate = new(DateFormat3Choice)
	return e.EntitlementFixingDate
}

func (e *EntitlementAssessment2) AddRegistrationSecuritiesDeadline() *DateFormat2Choice {
	e.RegistrationSecuritiesDeadline = new(DateFormat2Choice)
	return e.RegistrationSecuritiesDeadline
}

func (e *EntitlementAssessment2) AddRegistrationSecuritiesSTPDeadline() *DateFormat2Choice {
	e.RegistrationSecuritiesSTPDeadline = new(DateFormat2Choice)
	return e.RegistrationSecuritiesSTPDeadline
}

func (e *EntitlementAssessment2) AddRegistrationSecuritiesMarketDeadline() *DateFormat2Choice {
	e.RegistrationSecuritiesMarketDeadline = new(DateFormat2Choice)
	return e.RegistrationSecuritiesMarketDeadline
}

func (e *EntitlementAssessment2) AddRegistrationParticipationDeadline() *DateFormat2Choice {
	e.RegistrationParticipationDeadline = new(DateFormat2Choice)
	return e.RegistrationParticipationDeadline
}

func (e *EntitlementAssessment2) AddRegistrationParticipationSTPDeadline() *DateFormat2Choice {
	e.RegistrationParticipationSTPDeadline = new(DateFormat2Choice)
	return e.RegistrationParticipationSTPDeadline
}

func (e *EntitlementAssessment2) AddRegistrationParticipationMarketDeadline() *DateFormat2Choice {
	e.RegistrationParticipationMarketDeadline = new(DateFormat2Choice)
	return e.RegistrationParticipationMarketDeadline
}

func (e *EntitlementAssessment2) AddEntitlement() *Entitlement1Choice {
	e.Entitlement = new(Entitlement1Choice)
	return e.Entitlement
}

// Specifies the parameters, such as dates, used to calculate the entitlement to vote at a general meeting.
type EntitlementAssessment3 struct {

	// Date by which the securities should be blocked. This deadline is set by an intermediary.
	SecuritiesBlockingDeadline *DateFormat29Choice `xml:"SctiesBlckgDdln,omitempty"`

	// Date by which the securities should be blocked. This deadline is set by the issuer (STP mode).
	SecuritiesBlockingSTPDeadline *DateFormat29Choice `xml:"SctiesBlckgSTPDdln,omitempty"`

	// Date by which the securities should be blocked. This deadline is set by the issuer.
	SecuritiesBlockingMarketDeadline *DateFormat29Choice `xml:"SctiesBlckgMktDdln,omitempty"`

	// Date by which the blocking period for the securities should end.
	SecuritiesBlockingPeriodEndDate *ISODateTime `xml:"SctiesBlckgPrdEndDt,omitempty"`

	// Date at which the positions are struck to note which parties will receive the entitlement, for example. record date, book close date.
	EntitlementFixingDate *DateFormat1 `xml:"EntitlmntFxgDt,omitempty"`

	// Date by which the securities have to be registered. This deadline is specified by an intermediary.
	RegistrationSecuritiesDeadline *DateFormat29Choice `xml:"RegnSctiesDdln,omitempty"`

	// Date by which the securities have to be registered. This deadline is specified by an intermediary (STP mode).
	RegistrationSecuritiesSTPDeadline *DateFormat29Choice `xml:"RegnSctiesSTPDdln,omitempty"`

	// Date by which the securities have to be registered. This deadline is set by the issuer.
	RegistrationSecuritiesMarketDeadline *DateFormat29Choice `xml:"RegnSctiesMktDdln,omitempty"`

	// Date by which the holder needs to register its intention to participate in the meeting process in order to be allowed to participate in the meeting event. This deadline is specified by an intermediary.
	RegistrationParticipationDeadline *DateFormat29Choice `xml:"RegnPrtcptnDdln,omitempty"`

	// Date by which the holder needs to register its intention to participate in the meeting process in order to be allowed to participate in the meeting event. This deadline is specified by an intermediary (STP mode).
	RegistrationParticipationSTPDeadline *DateFormat29Choice `xml:"RegnPrtcptnSTPDdln,omitempty"`

	// Date by which the holder needs to register its intention to participate in the meeting process in order to be allowed to participate in the meeting event. This deadline is set by the issuer.
	RegistrationParticipationMarketDeadline *DateFormat29Choice `xml:"RegnPrtcptnMktDdln,omitempty"`

	// Number of votes assigned to one security.
	Entitlement *Entitlement1Choice `xml:"Entitlmnt,omitempty"`
}

func (e *EntitlementAssessment3) AddSecuritiesBlockingDeadline() *DateFormat29Choice {
	e.SecuritiesBlockingDeadline = new(DateFormat29Choice)
	return e.SecuritiesBlockingDeadline
}

func (e *EntitlementAssessment3) AddSecuritiesBlockingSTPDeadline() *DateFormat29Choice {
	e.SecuritiesBlockingSTPDeadline = new(DateFormat29Choice)
	return e.SecuritiesBlockingSTPDeadline
}

func (e *EntitlementAssessment3) AddSecuritiesBlockingMarketDeadline() *DateFormat29Choice {
	e.SecuritiesBlockingMarketDeadline = new(DateFormat29Choice)
	return e.SecuritiesBlockingMarketDeadline
}

func (e *EntitlementAssessment3) SetSecuritiesBlockingPeriodEndDate(value string) {
	e.SecuritiesBlockingPeriodEndDate = (*ISODateTime)(&value)
}

func (e *EntitlementAssessment3) AddEntitlementFixingDate() *DateFormat1 {
	e.EntitlementFixingDate = new(DateFormat1)
	return e.EntitlementFixingDate
}

func (e *EntitlementAssessment3) AddRegistrationSecuritiesDeadline() *DateFormat29Choice {
	e.RegistrationSecuritiesDeadline = new(DateFormat29Choice)
	return e.RegistrationSecuritiesDeadline
}

func (e *EntitlementAssessment3) AddRegistrationSecuritiesSTPDeadline() *DateFormat29Choice {
	e.RegistrationSecuritiesSTPDeadline = new(DateFormat29Choice)
	return e.RegistrationSecuritiesSTPDeadline
}

func (e *EntitlementAssessment3) AddRegistrationSecuritiesMarketDeadline() *DateFormat29Choice {
	e.RegistrationSecuritiesMarketDeadline = new(DateFormat29Choice)
	return e.RegistrationSecuritiesMarketDeadline
}

func (e *EntitlementAssessment3) AddRegistrationParticipationDeadline() *DateFormat29Choice {
	e.RegistrationParticipationDeadline = new(DateFormat29Choice)
	return e.RegistrationParticipationDeadline
}

func (e *EntitlementAssessment3) AddRegistrationParticipationSTPDeadline() *DateFormat29Choice {
	e.RegistrationParticipationSTPDeadline = new(DateFormat29Choice)
	return e.RegistrationParticipationSTPDeadline
}

func (e *EntitlementAssessment3) AddRegistrationParticipationMarketDeadline() *DateFormat29Choice {
	e.RegistrationParticipationMarketDeadline = new(DateFormat29Choice)
	return e.RegistrationParticipationMarketDeadline
}

func (e *EntitlementAssessment3) AddEntitlement() *Entitlement1Choice {
	e.Entitlement = new(Entitlement1Choice)
	return e.Entitlement
}
