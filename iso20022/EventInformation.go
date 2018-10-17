package iso20022

// Provides information about event of a corporate action.
type EventInformation1 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Provides the reference of the linked official corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Provides information about the identification of the last notification.
	LastNotificationIdentification *NotificationIdentification1 `xml:"LastNtfctnId,omitempty"`
}

func (e *EventInformation1) SetCorporateActionEventIdentification(value string) {
	e.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation1) SetOfficialCorporateActionEventIdentification(value string) {
	e.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation1) AddEventType() *CorporateActionEventType3Choice {
	e.EventType = new(CorporateActionEventType3Choice)
	return e.EventType
}

func (e *EventInformation1) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	e.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return e.MandatoryVoluntaryEventType
}

func (e *EventInformation1) AddLastNotificationIdentification() *NotificationIdentification1 {
	e.LastNotificationIdentification = new(NotificationIdentification1)
	return e.LastNotificationIdentification
}

// Provides information about event of a corporate action.
type EventInformation10 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Provides the reference of the linked official corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType58Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Provides information about the identification of the last notification.
	LastNotificationIdentification *NotificationIdentification4 `xml:"LastNtfctnId,omitempty"`
}

func (e *EventInformation10) SetCorporateActionEventIdentification(value string) {
	e.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (e *EventInformation10) SetOfficialCorporateActionEventIdentification(value string) {
	e.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (e *EventInformation10) AddEventType() *CorporateActionEventType58Choice {
	e.EventType = new(CorporateActionEventType58Choice)
	return e.EventType
}

func (e *EventInformation10) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	e.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return e.MandatoryVoluntaryEventType
}

func (e *EventInformation10) AddLastNotificationIdentification() *NotificationIdentification4 {
	e.LastNotificationIdentification = new(NotificationIdentification4)
	return e.LastNotificationIdentification
}

// Provides information about event of a corporate action.
type EventInformation3 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Provides the reference of the linked official corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType7Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Provides information about the identification of the last notification.
	LastNotificationIdentification *NotificationIdentification1 `xml:"LastNtfctnId,omitempty"`
}

func (e *EventInformation3) SetCorporateActionEventIdentification(value string) {
	e.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation3) SetOfficialCorporateActionEventIdentification(value string) {
	e.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation3) AddEventType() *CorporateActionEventType7Choice {
	e.EventType = new(CorporateActionEventType7Choice)
	return e.EventType
}

func (e *EventInformation3) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	e.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return e.MandatoryVoluntaryEventType
}

func (e *EventInformation3) AddLastNotificationIdentification() *NotificationIdentification1 {
	e.LastNotificationIdentification = new(NotificationIdentification1)
	return e.LastNotificationIdentification
}

// Provides information about event of a corporate action.
type EventInformation5 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Provides the reference of the linked official corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType11Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Provides information about the identification of the last notification.
	LastNotificationIdentification *NotificationIdentification1 `xml:"LastNtfctnId,omitempty"`
}

func (e *EventInformation5) SetCorporateActionEventIdentification(value string) {
	e.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation5) SetOfficialCorporateActionEventIdentification(value string) {
	e.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation5) AddEventType() *CorporateActionEventType11Choice {
	e.EventType = new(CorporateActionEventType11Choice)
	return e.EventType
}

func (e *EventInformation5) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	e.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return e.MandatoryVoluntaryEventType
}

func (e *EventInformation5) AddLastNotificationIdentification() *NotificationIdentification1 {
	e.LastNotificationIdentification = new(NotificationIdentification1)
	return e.LastNotificationIdentification
}

// Provides information about event of a corporate action.
type EventInformation7 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Provides the reference of the linked official corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType32Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Provides information about the identification of the last notification.
	LastNotificationIdentification *NotificationIdentification3 `xml:"LastNtfctnId,omitempty"`
}

func (e *EventInformation7) SetCorporateActionEventIdentification(value string) {
	e.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation7) SetOfficialCorporateActionEventIdentification(value string) {
	e.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation7) AddEventType() *CorporateActionEventType32Choice {
	e.EventType = new(CorporateActionEventType32Choice)
	return e.EventType
}

func (e *EventInformation7) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	e.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return e.MandatoryVoluntaryEventType
}

func (e *EventInformation7) AddLastNotificationIdentification() *NotificationIdentification3 {
	e.LastNotificationIdentification = new(NotificationIdentification3)
	return e.LastNotificationIdentification
}

// Provides information about event of a corporate action.
type EventInformation8 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Provides the reference of the linked official corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType36Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Provides information about the identification of the last notification.
	LastNotificationIdentification *NotificationIdentification4 `xml:"LastNtfctnId,omitempty"`
}

func (e *EventInformation8) SetCorporateActionEventIdentification(value string) {
	e.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (e *EventInformation8) SetOfficialCorporateActionEventIdentification(value string) {
	e.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (e *EventInformation8) AddEventType() *CorporateActionEventType36Choice {
	e.EventType = new(CorporateActionEventType36Choice)
	return e.EventType
}

func (e *EventInformation8) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	e.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return e.MandatoryVoluntaryEventType
}

func (e *EventInformation8) AddLastNotificationIdentification() *NotificationIdentification4 {
	e.LastNotificationIdentification = new(NotificationIdentification4)
	return e.LastNotificationIdentification
}

// Provides information about event of a corporate action.
type EventInformation9 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Provides the reference of the linked official corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType52Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Provides information about the identification of the last notification.
	LastNotificationIdentification *NotificationIdentification3 `xml:"LastNtfctnId,omitempty"`
}

func (e *EventInformation9) SetCorporateActionEventIdentification(value string) {
	e.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation9) SetOfficialCorporateActionEventIdentification(value string) {
	e.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (e *EventInformation9) AddEventType() *CorporateActionEventType52Choice {
	e.EventType = new(CorporateActionEventType52Choice)
	return e.EventType
}

func (e *EventInformation9) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	e.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return e.MandatoryVoluntaryEventType
}

func (e *EventInformation9) AddLastNotificationIdentification() *NotificationIdentification3 {
	e.LastNotificationIdentification = new(NotificationIdentification3)
	return e.LastNotificationIdentification
}
