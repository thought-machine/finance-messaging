package iso20022

// Information related to an identification, eg, party identification or account identification.
type SimpleIdentificationInformation struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`
}

func (s *SimpleIdentificationInformation) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

// Information related to an identification, eg, party identification or account identification.
type SimpleIdentificationInformation1 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *RestrictedFINXMax35Text `xml:"Id"`
}

func (s *SimpleIdentificationInformation1) SetIdentification(value string) {
	s.Identification = (*RestrictedFINXMax35Text)(&value)
}

// Information related to a party identification or account identification.
type SimpleIdentificationInformation2 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max34Text `xml:"Id"`
}

func (s *SimpleIdentificationInformation2) SetIdentification(value string) {
	s.Identification = (*Max34Text)(&value)
}

// Information related to an identification, eg, party identification or account identification.
type SimpleIdentificationInformation4 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`
}

func (s *SimpleIdentificationInformation4) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}
