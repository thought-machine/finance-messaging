package iso20022

// Specifies the physical parameters of a shareholders meeting. Several dates and places can be defined for a meeting
type Meeting2 struct {

	// Date and time at which the meeting will take place.
	DateAndTime *DateFormat2Choice `xml:"DtAndTm"`

	// Indicates the status of a meeting date.
	DateStatus *MeetingDateStatus1Code `xml:"DtSts,omitempty"`

	// Specifies whether a minimum number of security representation is required to hold a meeting.
	QuorumRequired *YesNoIndicator `xml:"QrmReqrd"`

	// Specifies location where meeting will take place.
	Location []*LocationFormat1Choice `xml:"Lctn"`

	// Minimum quantity of securities required to hold a meeting.
	QuorumQuantity *Max35Text `xml:"QrmQty,omitempty"`

	// Minimum quantity of securities, expressed as a percentage, required to hold a meeting.
	QuorumQuantityPercentage *PercentageRate `xml:"QrmQtyPctg,omitempty"`
}

func (m *Meeting2) AddDateAndTime() *DateFormat2Choice {
	m.DateAndTime = new(DateFormat2Choice)
	return m.DateAndTime
}

func (m *Meeting2) SetDateStatus(value string) {
	m.DateStatus = (*MeetingDateStatus1Code)(&value)
}

func (m *Meeting2) SetQuorumRequired(value string) {
	m.QuorumRequired = (*YesNoIndicator)(&value)
}

func (m *Meeting2) AddLocation() *LocationFormat1Choice {
	newValue := new(LocationFormat1Choice)
	m.Location = append(m.Location, newValue)
	return newValue
}

func (m *Meeting2) SetQuorumQuantity(value string) {
	m.QuorumQuantity = (*Max35Text)(&value)
}

func (m *Meeting2) SetQuorumQuantityPercentage(value string) {
	m.QuorumQuantityPercentage = (*PercentageRate)(&value)
}

// Specifies the physical parameters of a shareholders meeting. Several dates and places can be defined for a meeting
type Meeting3 struct {

	// Date and time at which the meeting will take place.
	DateAndTime *DateFormat2Choice `xml:"DtAndTm"`

	// Indicates the status of a meeting date.
	DateStatus *MeetingDateStatus1Code `xml:"DtSts,omitempty"`

	// Specifies whether a minimum number of security representation is required to hold a meeting.
	QuorumRequired *YesNoIndicator `xml:"QrmReqrd"`

	// Specifies location where meeting will take place.
	Location []*LocationFormat1Choice `xml:"Lctn"`

	// Minimum quantity of securities required to hold a meeting.
	QuorumQuantity *QuorumQuantity1Choice `xml:"QrmQty,omitempty"`
}

func (m *Meeting3) AddDateAndTime() *DateFormat2Choice {
	m.DateAndTime = new(DateFormat2Choice)
	return m.DateAndTime
}

func (m *Meeting3) SetDateStatus(value string) {
	m.DateStatus = (*MeetingDateStatus1Code)(&value)
}

func (m *Meeting3) SetQuorumRequired(value string) {
	m.QuorumRequired = (*YesNoIndicator)(&value)
}

func (m *Meeting3) AddLocation() *LocationFormat1Choice {
	newValue := new(LocationFormat1Choice)
	m.Location = append(m.Location, newValue)
	return newValue
}

func (m *Meeting3) AddQuorumQuantity() *QuorumQuantity1Choice {
	m.QuorumQuantity = new(QuorumQuantity1Choice)
	return m.QuorumQuantity
}

// Specifies the physical parameters of a shareholders meeting. Several dates and places can be defined for a meeting
type Meeting4 struct {

	// Date and time at which the meeting will take place.
	DateAndTime *DateFormat29Choice `xml:"DtAndTm"`

	// Indicates the status of the meeting date.
	DateStatus *MeetingDateStatus1Code `xml:"DtSts,omitempty"`

	// Specifies whether a minimum number of security representation is required to hold a meeting.
	QuorumRequired *YesNoIndicator `xml:"QrmReqrd"`

	// Specifies the location where meeting will take place.
	Location []*LocationFormat1Choice `xml:"Lctn"`

	// Minimum quantity of securities required to hold a meeting.
	QuorumQuantity *QuorumQuantity1Choice `xml:"QrmQty,omitempty"`
}

func (m *Meeting4) AddDateAndTime() *DateFormat29Choice {
	m.DateAndTime = new(DateFormat29Choice)
	return m.DateAndTime
}

func (m *Meeting4) SetDateStatus(value string) {
	m.DateStatus = (*MeetingDateStatus1Code)(&value)
}

func (m *Meeting4) SetQuorumRequired(value string) {
	m.QuorumRequired = (*YesNoIndicator)(&value)
}

func (m *Meeting4) AddLocation() *LocationFormat1Choice {
	newValue := new(LocationFormat1Choice)
	m.Location = append(m.Location, newValue)
	return newValue
}

func (m *Meeting4) AddQuorumQuantity() *QuorumQuantity1Choice {
	m.QuorumQuantity = new(QuorumQuantity1Choice)
	return m.QuorumQuantity
}
