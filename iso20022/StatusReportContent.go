package iso20022

// Content of the status report.
type StatusReportContent1 struct {

	// Capabilities of the POI performing the status report.
	POICapabilities *PointOfInteractionCapabilities1 `xml:"POICpblties,omitempty"`

	// Data related to a component of the POI performing the status report.
	POIComponent []*PointOfInteractionComponent2 `xml:"POICmpnt,omitempty"`

	// Human attendance at the POI location during transactions.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// System date time of the point of interaction (POI) sending the status report.
	POIDateTime *ISODateTime `xml:"POIDtTm"`

	// Request the terminal management system to answer with the identified data set.
	DataSetRequired *DataSetIdentification2 `xml:"DataSetReqrd,omitempty"`

	// Result of an individual terminal management action by the point of interaction.
	Event []*TMSEvent1 `xml:"Evt,omitempty"`

	// Error log of the point of interaction since the last status report.
	Errors *Max140Text `xml:"Errs,omitempty"`
}

func (s *StatusReportContent1) AddPOICapabilities() *PointOfInteractionCapabilities1 {
	s.POICapabilities = new(PointOfInteractionCapabilities1)
	return s.POICapabilities
}

func (s *StatusReportContent1) AddPOIComponent() *PointOfInteractionComponent2 {
	newValue := new(PointOfInteractionComponent2)
	s.POIComponent = append(s.POIComponent, newValue)
	return newValue
}

func (s *StatusReportContent1) SetAttendanceContext(value string) {
	s.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (s *StatusReportContent1) SetPOIDateTime(value string) {
	s.POIDateTime = (*ISODateTime)(&value)
}

func (s *StatusReportContent1) AddDataSetRequired() *DataSetIdentification2 {
	s.DataSetRequired = new(DataSetIdentification2)
	return s.DataSetRequired
}

func (s *StatusReportContent1) AddEvent() *TMSEvent1 {
	newValue := new(TMSEvent1)
	s.Event = append(s.Event, newValue)
	return newValue
}

func (s *StatusReportContent1) SetErrors(value string) {
	s.Errors = (*Max140Text)(&value)
}

// Content of the status report.
type StatusReportContent2 struct {

	// Capabilities of the POI performing the status report.
	POICapabilities *PointOfInteractionCapabilities1 `xml:"POICpblties,omitempty"`

	// Data related to a component of the POI performing the status report.
	POIComponent []*PointOfInteractionComponent3 `xml:"POICmpnt,omitempty"`

	// Human attendance at the POI location during transactions.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// System date time of the point of interaction (POI) sending the status report.
	POIDateTime *ISODateTime `xml:"POIDtTm"`

	// Request the terminal management system to answer with the identified data set.
	DataSetRequired *TerminalManagementDataSet7 `xml:"DataSetReqrd,omitempty"`

	// Result of an individual terminal management action by the point of interaction.
	Event []*TMSEvent2 `xml:"Evt,omitempty"`

	// Error log of the point of interaction since the last status report.
	Errors *Max140Text `xml:"Errs,omitempty"`
}

func (s *StatusReportContent2) AddPOICapabilities() *PointOfInteractionCapabilities1 {
	s.POICapabilities = new(PointOfInteractionCapabilities1)
	return s.POICapabilities
}

func (s *StatusReportContent2) AddPOIComponent() *PointOfInteractionComponent3 {
	newValue := new(PointOfInteractionComponent3)
	s.POIComponent = append(s.POIComponent, newValue)
	return newValue
}

func (s *StatusReportContent2) SetAttendanceContext(value string) {
	s.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (s *StatusReportContent2) SetPOIDateTime(value string) {
	s.POIDateTime = (*ISODateTime)(&value)
}

func (s *StatusReportContent2) AddDataSetRequired() *TerminalManagementDataSet7 {
	s.DataSetRequired = new(TerminalManagementDataSet7)
	return s.DataSetRequired
}

func (s *StatusReportContent2) AddEvent() *TMSEvent2 {
	newValue := new(TMSEvent2)
	s.Event = append(s.Event, newValue)
	return newValue
}

func (s *StatusReportContent2) SetErrors(value string) {
	s.Errors = (*Max140Text)(&value)
}

// Content of the status report.
type StatusReportContent3 struct {

	// Capabilities of the POI (Point Of Interaction) performing the status report.
	POICapabilities *PointOfInteractionCapabilities2 `xml:"POICpblties,omitempty"`

	// Data related to a component of the POI (Point Of Interaction) performing the status report.
	POIComponent []*PointOfInteractionComponent4 `xml:"POICmpnt,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during transactions.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// System date time of the point of interaction (POI) sending the status report.
	POIDateTime *ISODateTime `xml:"POIDtTm"`

	// Request the terminal management system to answer with the identified data set.
	DataSetRequired *TerminalManagementDataSet8 `xml:"DataSetReqrd,omitempty"`

	// Result of an individual terminal management action by the point of interaction.
	Event []*TMSEvent2 `xml:"Evt,omitempty"`

	// Error log of the point of interaction since the last status report.
	Errors []*Max140Text `xml:"Errs,omitempty"`
}

func (s *StatusReportContent3) AddPOICapabilities() *PointOfInteractionCapabilities2 {
	s.POICapabilities = new(PointOfInteractionCapabilities2)
	return s.POICapabilities
}

func (s *StatusReportContent3) AddPOIComponent() *PointOfInteractionComponent4 {
	newValue := new(PointOfInteractionComponent4)
	s.POIComponent = append(s.POIComponent, newValue)
	return newValue
}

func (s *StatusReportContent3) SetAttendanceContext(value string) {
	s.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (s *StatusReportContent3) SetPOIDateTime(value string) {
	s.POIDateTime = (*ISODateTime)(&value)
}

func (s *StatusReportContent3) AddDataSetRequired() *TerminalManagementDataSet8 {
	s.DataSetRequired = new(TerminalManagementDataSet8)
	return s.DataSetRequired
}

func (s *StatusReportContent3) AddEvent() *TMSEvent2 {
	newValue := new(TMSEvent2)
	s.Event = append(s.Event, newValue)
	return newValue
}

func (s *StatusReportContent3) AddErrors(value string) {
	s.Errors = append(s.Errors, (*Max140Text)(&value))
}

// Content of the status report.
type StatusReportContent4 struct {

	// Capabilities of the POI (Point Of Interaction) performing the status report.
	POICapabilities *PointOfInteractionCapabilities3 `xml:"POICpblties,omitempty"`

	// Data related to a component of the POI (Point Of Interaction) performing the status report.
	POIComponent []*PointOfInteractionComponent5 `xml:"POICmpnt,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during transactions.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// System date time of the point of interaction (POI) sending the status report.
	POIDateTime *ISODateTime `xml:"POIDtTm"`

	// Request the terminal management system to answer with the identified data set.
	DataSetRequired *TerminalManagementDataSet12 `xml:"DataSetReqrd,omitempty"`

	// Result of an individual terminal management action by the point of interaction.
	Event []*TMSEvent3 `xml:"Evt,omitempty"`

	// Error log of the point of interaction since the last status report.
	Errors []*Max140Text `xml:"Errs,omitempty"`
}

func (s *StatusReportContent4) AddPOICapabilities() *PointOfInteractionCapabilities3 {
	s.POICapabilities = new(PointOfInteractionCapabilities3)
	return s.POICapabilities
}

func (s *StatusReportContent4) AddPOIComponent() *PointOfInteractionComponent5 {
	newValue := new(PointOfInteractionComponent5)
	s.POIComponent = append(s.POIComponent, newValue)
	return newValue
}

func (s *StatusReportContent4) SetAttendanceContext(value string) {
	s.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (s *StatusReportContent4) SetPOIDateTime(value string) {
	s.POIDateTime = (*ISODateTime)(&value)
}

func (s *StatusReportContent4) AddDataSetRequired() *TerminalManagementDataSet12 {
	s.DataSetRequired = new(TerminalManagementDataSet12)
	return s.DataSetRequired
}

func (s *StatusReportContent4) AddEvent() *TMSEvent3 {
	newValue := new(TMSEvent3)
	s.Event = append(s.Event, newValue)
	return newValue
}

func (s *StatusReportContent4) AddErrors(value string) {
	s.Errors = append(s.Errors, (*Max140Text)(&value))
}

// Content of the status report.
type StatusReportContent5 struct {

	// Capabilities of the POI (Point Of Interaction) performing the status report.
	POICapabilities *PointOfInteractionCapabilities6 `xml:"POICpblties,omitempty"`

	// Data related to a component of the POI (Point Of Interaction) performing the status report.
	POIComponent []*PointOfInteractionComponent6 `xml:"POICmpnt,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during transactions.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// System date time of the point of interaction (POI) sending the status report.
	POIDateTime *ISODateTime `xml:"POIDtTm"`

	// Request the terminal management system to answer with the identified data set.
	DataSetRequired *TerminalManagementDataSet17 `xml:"DataSetReqrd,omitempty"`

	// Result of an individual terminal management action by the point of interaction.
	Event []*TMSEvent4 `xml:"Evt,omitempty"`

	// Error log of the point of interaction since the last status report.
	Errors []*Max140Text `xml:"Errs,omitempty"`
}

func (s *StatusReportContent5) AddPOICapabilities() *PointOfInteractionCapabilities6 {
	s.POICapabilities = new(PointOfInteractionCapabilities6)
	return s.POICapabilities
}

func (s *StatusReportContent5) AddPOIComponent() *PointOfInteractionComponent6 {
	newValue := new(PointOfInteractionComponent6)
	s.POIComponent = append(s.POIComponent, newValue)
	return newValue
}

func (s *StatusReportContent5) SetAttendanceContext(value string) {
	s.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (s *StatusReportContent5) SetPOIDateTime(value string) {
	s.POIDateTime = (*ISODateTime)(&value)
}

func (s *StatusReportContent5) AddDataSetRequired() *TerminalManagementDataSet17 {
	s.DataSetRequired = new(TerminalManagementDataSet17)
	return s.DataSetRequired
}

func (s *StatusReportContent5) AddEvent() *TMSEvent4 {
	newValue := new(TMSEvent4)
	s.Event = append(s.Event, newValue)
	return newValue
}

func (s *StatusReportContent5) AddErrors(value string) {
	s.Errors = append(s.Errors, (*Max140Text)(&value))
}
