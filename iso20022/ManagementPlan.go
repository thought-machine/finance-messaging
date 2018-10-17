package iso20022

// Sequence of terminal management actions to be performed by a point of interaction (POI).
type ManagementPlan1 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the management plan.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId"`

	// Data set related to the sequence of actions to be performed by a point of interaction (POI).
	DataSet []*TerminalManagementDataSet2 `xml:"DataSet"`
}

func (m *ManagementPlan1) AddPOIIdentification() *GenericIdentification35 {
	m.POIIdentification = new(GenericIdentification35)
	return m.POIIdentification
}

func (m *ManagementPlan1) AddTerminalManagerIdentification() *GenericIdentification35 {
	m.TerminalManagerIdentification = new(GenericIdentification35)
	return m.TerminalManagerIdentification
}

func (m *ManagementPlan1) AddDataSet() *TerminalManagementDataSet2 {
	newValue := new(TerminalManagementDataSet2)
	m.DataSet = append(m.DataSet, newValue)
	return newValue
}

// Sequence of terminal management actions to be performed by a point of interaction (POI).
type ManagementPlan2 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the management plan.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId"`

	// Data set related to the sequence of actions to be performed by a point of interaction (POI).
	DataSet []*TerminalManagementDataSet5 `xml:"DataSet"`
}

func (m *ManagementPlan2) AddPOIIdentification() *GenericIdentification35 {
	m.POIIdentification = new(GenericIdentification35)
	return m.POIIdentification
}

func (m *ManagementPlan2) AddTerminalManagerIdentification() *GenericIdentification35 {
	m.TerminalManagerIdentification = new(GenericIdentification35)
	return m.TerminalManagerIdentification
}

func (m *ManagementPlan2) AddDataSet() *TerminalManagementDataSet5 {
	newValue := new(TerminalManagementDataSet5)
	m.DataSet = append(m.DataSet, newValue)
	return newValue
}

// Sequence of terminal management actions to be performed by a point of interaction (POI).
type ManagementPlan3 struct {

	// Identification of the point of interaction (POI) for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the management plan.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId"`

	// Data set related to the sequence of actions to be performed by a point of interaction (POI).
	DataSet []*TerminalManagementDataSet10 `xml:"DataSet"`
}

func (m *ManagementPlan3) AddPOIIdentification() *GenericIdentification35 {
	m.POIIdentification = new(GenericIdentification35)
	return m.POIIdentification
}

func (m *ManagementPlan3) AddTerminalManagerIdentification() *GenericIdentification35 {
	m.TerminalManagerIdentification = new(GenericIdentification35)
	return m.TerminalManagerIdentification
}

func (m *ManagementPlan3) AddDataSet() *TerminalManagementDataSet10 {
	newValue := new(TerminalManagementDataSet10)
	m.DataSet = append(m.DataSet, newValue)
	return newValue
}

// Sequence of terminal management actions to be performed by a point of interaction (POI).
type ManagementPlan4 struct {

	// Identification of the point of interaction (POI) for terminal management.
	POIIdentification *GenericIdentification71 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the management plan.
	TerminalManagerIdentification *GenericIdentification71 `xml:"TermnlMgrId"`

	// Data set related to the sequence of actions to be performed by a point of interaction (POI).
	DataSet *TerminalManagementDataSet15 `xml:"DataSet"`
}

func (m *ManagementPlan4) AddPOIIdentification() *GenericIdentification71 {
	m.POIIdentification = new(GenericIdentification71)
	return m.POIIdentification
}

func (m *ManagementPlan4) AddTerminalManagerIdentification() *GenericIdentification71 {
	m.TerminalManagerIdentification = new(GenericIdentification71)
	return m.TerminalManagerIdentification
}

func (m *ManagementPlan4) AddDataSet() *TerminalManagementDataSet15 {
	m.DataSet = new(TerminalManagementDataSet15)
	return m.DataSet
}

// Sequence of terminal management actions to be performed by a point of interaction (POI).
type ManagementPlan5 struct {

	// Identification of the point of interaction (POI) for terminal management.
	POIIdentification *GenericIdentification71 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the management plan.
	TerminalManagerIdentification *GenericIdentification71 `xml:"TermnlMgrId"`

	// Data set related to the sequence of actions to be performed by a point of interaction (POI).
	DataSet *TerminalManagementDataSet18 `xml:"DataSet"`
}

func (m *ManagementPlan5) AddPOIIdentification() *GenericIdentification71 {
	m.POIIdentification = new(GenericIdentification71)
	return m.POIIdentification
}

func (m *ManagementPlan5) AddTerminalManagerIdentification() *GenericIdentification71 {
	m.TerminalManagerIdentification = new(GenericIdentification71)
	return m.TerminalManagerIdentification
}

func (m *ManagementPlan5) AddDataSet() *TerminalManagementDataSet18 {
	m.DataSet = new(TerminalManagementDataSet18)
	return m.DataSet
}
