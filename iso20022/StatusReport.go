package iso20022

// Status of the acceptor system containing the identification of the POI, its components and their installed versions.
type StatusReport1 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId"`

	// Identification of the terminal management system (TMS) to contact for the maintenance.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId,omitempty"`

	// Data related to a status report of a point of interaction (POI).
	DataSet []*TerminalManagementDataSet1 `xml:"DataSet"`
}

func (s *StatusReport1) AddPOIIdentification() *GenericIdentification35 {
	s.POIIdentification = new(GenericIdentification35)
	return s.POIIdentification
}

func (s *StatusReport1) AddTerminalManagerIdentification() *GenericIdentification35 {
	s.TerminalManagerIdentification = new(GenericIdentification35)
	return s.TerminalManagerIdentification
}

func (s *StatusReport1) AddDataSet() *TerminalManagementDataSet1 {
	newValue := new(TerminalManagementDataSet1)
	s.DataSet = append(s.DataSet, newValue)
	return newValue
}

// Status of the acceptor system containing the identification of the POI, its components and their installed versions.
type StatusReport2 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId"`

	// Identification of the terminal management system (TMS) to contact for the maintenance.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId,omitempty"`

	// Data related to a status report of a point of interaction (POI).
	DataSet []*TerminalManagementDataSet4 `xml:"DataSet"`
}

func (s *StatusReport2) AddPOIIdentification() *GenericIdentification35 {
	s.POIIdentification = new(GenericIdentification35)
	return s.POIIdentification
}

func (s *StatusReport2) AddTerminalManagerIdentification() *GenericIdentification35 {
	s.TerminalManagerIdentification = new(GenericIdentification35)
	return s.TerminalManagerIdentification
}

func (s *StatusReport2) AddDataSet() *TerminalManagementDataSet4 {
	newValue := new(TerminalManagementDataSet4)
	s.DataSet = append(s.DataSet, newValue)
	return newValue
}

// Status of the acceptor system containing the identification of the POI (Point Of Interaction), its components and their installed versions.
type StatusReport3 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId"`

	// Identification of the terminal management system (TMS) to contact for the maintenance.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId,omitempty"`

	// Data related to a status report of a point of interaction (POI).
	DataSet []*TerminalManagementDataSet9 `xml:"DataSet"`
}

func (s *StatusReport3) AddPOIIdentification() *GenericIdentification35 {
	s.POIIdentification = new(GenericIdentification35)
	return s.POIIdentification
}

func (s *StatusReport3) AddTerminalManagerIdentification() *GenericIdentification35 {
	s.TerminalManagerIdentification = new(GenericIdentification35)
	return s.TerminalManagerIdentification
}

func (s *StatusReport3) AddDataSet() *TerminalManagementDataSet9 {
	newValue := new(TerminalManagementDataSet9)
	s.DataSet = append(s.DataSet, newValue)
	return newValue
}

// Status of the acceptor system containing the identification of the POI (Point Of Interaction), its components and their installed versions.
type StatusReport4 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification71 `xml:"POIId"`

	// Identification of the terminal management system (TMS) to contact for the maintenance.
	TerminalManagerIdentification *GenericIdentification71 `xml:"TermnlMgrId,omitempty"`

	// Data related to a status report of a point of interaction (POI).
	DataSet *TerminalManagementDataSet13 `xml:"DataSet"`
}

func (s *StatusReport4) AddPOIIdentification() *GenericIdentification71 {
	s.POIIdentification = new(GenericIdentification71)
	return s.POIIdentification
}

func (s *StatusReport4) AddTerminalManagerIdentification() *GenericIdentification71 {
	s.TerminalManagerIdentification = new(GenericIdentification71)
	return s.TerminalManagerIdentification
}

func (s *StatusReport4) AddDataSet() *TerminalManagementDataSet13 {
	s.DataSet = new(TerminalManagementDataSet13)
	return s.DataSet
}

// Status of the acceptor system containing the identification of the POI (Point Of Interaction), its components and their installed versions.
type StatusReport5 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification71 `xml:"POIId"`

	// Identification of the terminal management system (TMS) to contact for the maintenance.
	TerminalManagerIdentification *GenericIdentification71 `xml:"TermnlMgrId"`

	// Data related to a status report of a point of interaction (POI).
	DataSet *TerminalManagementDataSet16 `xml:"DataSet"`
}

func (s *StatusReport5) AddPOIIdentification() *GenericIdentification71 {
	s.POIIdentification = new(GenericIdentification71)
	return s.POIIdentification
}

func (s *StatusReport5) AddTerminalManagerIdentification() *GenericIdentification71 {
	s.TerminalManagerIdentification = new(GenericIdentification71)
	return s.TerminalManagerIdentification
}

func (s *StatusReport5) AddDataSet() *TerminalManagementDataSet16 {
	s.DataSet = new(TerminalManagementDataSet16)
	return s.DataSet
}
