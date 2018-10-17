package iso20022

// Content of the management plan.
type ManagementPlanContent1 struct {

	// Terminal management action to be performed by the point of interaction (POI).
	Action []*TMSAction1 `xml:"Actn"`
}

func (m *ManagementPlanContent1) AddAction() *TMSAction1 {
	newValue := new(TMSAction1)
	m.Action = append(m.Action, newValue)
	return newValue
}

// Content of the management plan.
type ManagementPlanContent2 struct {

	// Terminal management action to be performed by the point of interaction (POI).
	Action []*TMSAction2 `xml:"Actn"`
}

func (m *ManagementPlanContent2) AddAction() *TMSAction2 {
	newValue := new(TMSAction2)
	m.Action = append(m.Action, newValue)
	return newValue
}

// Content of the management plan.
type ManagementPlanContent3 struct {

	// Terminal management action to be performed by the point of interaction (POI).
	Action []*TMSAction3 `xml:"Actn"`
}

func (m *ManagementPlanContent3) AddAction() *TMSAction3 {
	newValue := new(TMSAction3)
	m.Action = append(m.Action, newValue)
	return newValue
}

// Content of the management plan.
type ManagementPlanContent4 struct {

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Certificate chain of an asymmetric encryption keys for the encryption of temporary transport key of the key to inject.
	KeyEnciphermentCertificate []*Max10KBinary `xml:"KeyNcphrmntCert,omitempty"`

	// Terminal management action to be performed by the point of interaction (POI).
	Action []*TMSAction4 `xml:"Actn"`
}

func (m *ManagementPlanContent4) SetTMChallenge(value string) {
	m.TMChallenge = (*Max140Binary)(&value)
}

func (m *ManagementPlanContent4) AddKeyEnciphermentCertificate(value string) {
	m.KeyEnciphermentCertificate = append(m.KeyEnciphermentCertificate, (*Max10KBinary)(&value))
}

func (m *ManagementPlanContent4) AddAction() *TMSAction4 {
	newValue := new(TMSAction4)
	m.Action = append(m.Action, newValue)
	return newValue
}

// Content of the management plan.
type ManagementPlanContent5 struct {

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Certificate chain of an asymmetric encryption keys for the encryption of temporary transport key of the key to inject.
	KeyEnciphermentCertificate []*Max10KBinary `xml:"KeyNcphrmntCert,omitempty"`

	// Terminal management action to be performed by the point of interaction (POI).
	Action []*TMSAction5 `xml:"Actn"`
}

func (m *ManagementPlanContent5) SetTMChallenge(value string) {
	m.TMChallenge = (*Max140Binary)(&value)
}

func (m *ManagementPlanContent5) AddKeyEnciphermentCertificate(value string) {
	m.KeyEnciphermentCertificate = append(m.KeyEnciphermentCertificate, (*Max10KBinary)(&value))
}

func (m *ManagementPlanContent5) AddAction() *TMSAction5 {
	newValue := new(TMSAction5)
	m.Action = append(m.Action, newValue)
	return newValue
}
