package iso20022

// Describes the multimodal or the individual transport of goods.
type TransportMeans1 struct {

	// Moving of goods or people from one place to another by vehicle.
	IndividualTransport *SingleTransport4 `xml:"IndvTrnsprt"`

	// Specifies the different movements and places and their role in a multimodal conveyance of goods.
	MultimodalTransport *MultimodalTransport3 `xml:"MltmdlTrnsprt,omitempty"`
}

func (t *TransportMeans1) AddIndividualTransport() *SingleTransport4 {
	t.IndividualTransport = new(SingleTransport4)
	return t.IndividualTransport
}

func (t *TransportMeans1) AddMultimodalTransport() *MultimodalTransport3 {
	t.MultimodalTransport = new(MultimodalTransport3)
	return t.MultimodalTransport
}

// Describes the multimodal or the individual transport of goods.
type TransportMeans2 struct {

	// Moving of goods or people from one place to another by vehicle.
	IndividualTransport *SingleTransport5 `xml:"IndvTrnsprt"`

	// Specifies the different movements and places and their role in a multimodal conveyance of goods.
	MultimodalTransport *MultimodalTransport3 `xml:"MltmdlTrnsprt,omitempty"`
}

func (t *TransportMeans2) AddIndividualTransport() *SingleTransport5 {
	t.IndividualTransport = new(SingleTransport5)
	return t.IndividualTransport
}

func (t *TransportMeans2) AddMultimodalTransport() *MultimodalTransport3 {
	t.MultimodalTransport = new(MultimodalTransport3)
	return t.MultimodalTransport
}

// Describes the multimodal or the individual transport of goods.
type TransportMeans3 struct {

	// Code specifying the transport mode for the delivery of the consignment, such as by air, sea, rail, road or inland waterway. Reference UN/ECE Recommendation 19 - Code for Modes of Transport (www.unece.org/cefact/recommendations).
	ModeCode *Max4Text `xml:"MdCd,omitempty"`

	// Unique identification of the means of transport, such as the International Maritime Organization number of a vessel.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Name, expressed as text, of the means of transport.
	Name *Max35Text `xml:"Nm,omitempty"`
}

func (t *TransportMeans3) SetModeCode(value string) {
	t.ModeCode = (*Max4Text)(&value)
}

func (t *TransportMeans3) SetIdentification(value string) {
	t.Identification = (*Max35Text)(&value)
}

func (t *TransportMeans3) SetName(value string) {
	t.Name = (*Max35Text)(&value)
}

// Describes the multimodal or the individual transport of goods.
type TransportMeans4 struct {

	// Moving of goods or people from one place to another by vehicle.
	IndividualTransport *SingleTransport6 `xml:"IndvTrnsprt"`

	// Specifies the different movements and places and their role in a multimodal conveyance of goods.
	MultimodalTransport *MultimodalTransport3 `xml:"MltmdlTrnsprt,omitempty"`
}

func (t *TransportMeans4) AddIndividualTransport() *SingleTransport6 {
	t.IndividualTransport = new(SingleTransport6)
	return t.IndividualTransport
}

func (t *TransportMeans4) AddMultimodalTransport() *MultimodalTransport3 {
	t.MultimodalTransport = new(MultimodalTransport3)
	return t.MultimodalTransport
}

// Describes the multimodal or the individual transport of goods.
type TransportMeans5 struct {

	// Moving of goods or people from one place to another by vehicle.
	IndividualTransport *SingleTransport7 `xml:"IndvTrnsprt"`

	// Specifies the different movements and places and their role in a multimodal conveyance of goods.
	MultimodalTransport *MultimodalTransport3 `xml:"MltmdlTrnsprt,omitempty"`
}

func (t *TransportMeans5) AddIndividualTransport() *SingleTransport7 {
	t.IndividualTransport = new(SingleTransport7)
	return t.IndividualTransport
}

func (t *TransportMeans5) AddMultimodalTransport() *MultimodalTransport3 {
	t.MultimodalTransport = new(MultimodalTransport3)
	return t.MultimodalTransport
}

// Describes the multimodal or the individual transport of goods.
type TransportMeans6 struct {

	// Moving of goods or people from one place to another by vehicle.
	IndividualTransport *SingleTransport8 `xml:"IndvTrnsprt"`

	// Specifies the different movements and places and their role in a multimodal conveyance of goods.
	MultimodalTransport *MultimodalTransport3 `xml:"MltmdlTrnsprt,omitempty"`
}

func (t *TransportMeans6) AddIndividualTransport() *SingleTransport8 {
	t.IndividualTransport = new(SingleTransport8)
	return t.IndividualTransport
}

func (t *TransportMeans6) AddMultimodalTransport() *MultimodalTransport3 {
	t.MultimodalTransport = new(MultimodalTransport3)
	return t.MultimodalTransport
}
