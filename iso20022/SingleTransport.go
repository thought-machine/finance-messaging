package iso20022

// Specifies individually each leg of a transport of goods.
type SingleTransport3 struct {

	// Information related to the transportation of goods by air.
	TransportByAir *TransportByAir2 `xml:"TrnsprtByAir,omitempty"`

	// Information related for the transportation of goods by sea.
	TransportBySea *TransportBySea4 `xml:"TrnsprtBySea,omitempty"`

	// Information related to the transportation of goods by road.
	TransportByRoad *TransportByRoad2 `xml:"TrnsprtByRoad,omitempty"`

	// Information related to the transportation of goods by rail.
	TransportByRail *TransportByRail2 `xml:"TrnsprtByRail,omitempty"`
}

func (s *SingleTransport3) AddTransportByAir() *TransportByAir2 {
	s.TransportByAir = new(TransportByAir2)
	return s.TransportByAir
}

func (s *SingleTransport3) AddTransportBySea() *TransportBySea4 {
	s.TransportBySea = new(TransportBySea4)
	return s.TransportBySea
}

func (s *SingleTransport3) AddTransportByRoad() *TransportByRoad2 {
	s.TransportByRoad = new(TransportByRoad2)
	return s.TransportByRoad
}

func (s *SingleTransport3) AddTransportByRail() *TransportByRail2 {
	s.TransportByRail = new(TransportByRail2)
	return s.TransportByRail
}

// Specifies individually each leg of a transport of goods.
type SingleTransport4 struct {

	// Moving of goods or people from one place to another by vehicle.
	TransportByAir []*TransportByAir3 `xml:"TrnsprtByAir,omitempty"`

	// Information related for the transportation of goods by sea.
	TransportBySea []*TransportBySea3 `xml:"TrnsprtBySea,omitempty"`

	// Information related to the transportation of goods by road.
	TransportByRoad []*TransportByRoad3 `xml:"TrnsprtByRoad,omitempty"`

	// Information related to the transportation of goods by rail.
	TransportByRail []*TransportByRail3 `xml:"TrnsprtByRail,omitempty"`
}

func (s *SingleTransport4) AddTransportByAir() *TransportByAir3 {
	newValue := new(TransportByAir3)
	s.TransportByAir = append(s.TransportByAir, newValue)
	return newValue
}

func (s *SingleTransport4) AddTransportBySea() *TransportBySea3 {
	newValue := new(TransportBySea3)
	s.TransportBySea = append(s.TransportBySea, newValue)
	return newValue
}

func (s *SingleTransport4) AddTransportByRoad() *TransportByRoad3 {
	newValue := new(TransportByRoad3)
	s.TransportByRoad = append(s.TransportByRoad, newValue)
	return newValue
}

func (s *SingleTransport4) AddTransportByRail() *TransportByRail3 {
	newValue := new(TransportByRail3)
	s.TransportByRail = append(s.TransportByRail, newValue)
	return newValue
}

// Specifies individually each leg of a transport of goods.
type SingleTransport5 struct {

	// Information related to the transportation of goods by air.
	TransportByAir []*TransportByAir2 `xml:"TrnsprtByAir,omitempty"`

	// Information related for the transportation of goods by sea.
	TransportBySea []*TransportBySea4 `xml:"TrnsprtBySea,omitempty"`

	// Moving of goods or people from one place to another by vehicle.
	TransportByRoad []*TransportByRoad2 `xml:"TrnsprtByRoad,omitempty"`

	// Information related to the transportation of goods by rail.
	TransportByRail []*TransportByRail2 `xml:"TrnsprtByRail,omitempty"`
}

func (s *SingleTransport5) AddTransportByAir() *TransportByAir2 {
	newValue := new(TransportByAir2)
	s.TransportByAir = append(s.TransportByAir, newValue)
	return newValue
}

func (s *SingleTransport5) AddTransportBySea() *TransportBySea4 {
	newValue := new(TransportBySea4)
	s.TransportBySea = append(s.TransportBySea, newValue)
	return newValue
}

func (s *SingleTransport5) AddTransportByRoad() *TransportByRoad2 {
	newValue := new(TransportByRoad2)
	s.TransportByRoad = append(s.TransportByRoad, newValue)
	return newValue
}

func (s *SingleTransport5) AddTransportByRail() *TransportByRail2 {
	newValue := new(TransportByRail2)
	s.TransportByRail = append(s.TransportByRail, newValue)
	return newValue
}

// Specifies individually each leg of a transport of goods.
type SingleTransport6 struct {

	// Information related to the transportation of goods by air.
	TransportByAir []*TransportByAir4 `xml:"TrnsprtByAir,omitempty"`

	// Information related for the transportation of goods by sea.
	TransportBySea []*TransportBySea5 `xml:"TrnsprtBySea,omitempty"`

	// Moving of goods or people from one place to another by vehicle.
	TransportByRoad []*TransportByRoad4 `xml:"TrnsprtByRoad,omitempty"`

	// Information related to the transportation of goods by rail.
	TransportByRail []*TransportByRail4 `xml:"TrnsprtByRail,omitempty"`
}

func (s *SingleTransport6) AddTransportByAir() *TransportByAir4 {
	newValue := new(TransportByAir4)
	s.TransportByAir = append(s.TransportByAir, newValue)
	return newValue
}

func (s *SingleTransport6) AddTransportBySea() *TransportBySea5 {
	newValue := new(TransportBySea5)
	s.TransportBySea = append(s.TransportBySea, newValue)
	return newValue
}

func (s *SingleTransport6) AddTransportByRoad() *TransportByRoad4 {
	newValue := new(TransportByRoad4)
	s.TransportByRoad = append(s.TransportByRoad, newValue)
	return newValue
}

func (s *SingleTransport6) AddTransportByRail() *TransportByRail4 {
	newValue := new(TransportByRail4)
	s.TransportByRail = append(s.TransportByRail, newValue)
	return newValue
}

// Specifies individually each leg of a transport of goods.
type SingleTransport7 struct {

	// Moving of goods or people from one place to another by vehicle.
	TransportByAir []*TransportByAir5 `xml:"TrnsprtByAir,omitempty"`

	// Information related for the transportation of goods by sea.
	TransportBySea []*TransportBySea6 `xml:"TrnsprtBySea,omitempty"`

	// Information related to the transportation of goods by road.
	TransportByRoad []*TransportByRoad5 `xml:"TrnsprtByRoad,omitempty"`

	// Information related to the transportation of goods by rail.
	TransportByRail []*TransportByRail5 `xml:"TrnsprtByRail,omitempty"`
}

func (s *SingleTransport7) AddTransportByAir() *TransportByAir5 {
	newValue := new(TransportByAir5)
	s.TransportByAir = append(s.TransportByAir, newValue)
	return newValue
}

func (s *SingleTransport7) AddTransportBySea() *TransportBySea6 {
	newValue := new(TransportBySea6)
	s.TransportBySea = append(s.TransportBySea, newValue)
	return newValue
}

func (s *SingleTransport7) AddTransportByRoad() *TransportByRoad5 {
	newValue := new(TransportByRoad5)
	s.TransportByRoad = append(s.TransportByRoad, newValue)
	return newValue
}

func (s *SingleTransport7) AddTransportByRail() *TransportByRail5 {
	newValue := new(TransportByRail5)
	s.TransportByRail = append(s.TransportByRail, newValue)
	return newValue
}

// Specifies individually each leg of a transport of goods.
type SingleTransport8 struct {

	// Information related to the transportation of goods by air.
	TransportByAir []*TransportByAir4 `xml:"TrnsprtByAir,omitempty"`

	// Information related for the transportation of goods by sea.
	TransportBySea []*TransportBySea5 `xml:"TrnsprtBySea,omitempty"`

	// Moving of goods or people from one place to another by vehicle.
	TransportByRoad []*TransportByRoad4 `xml:"TrnsprtByRoad,omitempty"`

	// Information related to the transportation of goods by rail.
	TransportByRail []*TransportByRail4 `xml:"TrnsprtByRail,omitempty"`
}

func (s *SingleTransport8) AddTransportByAir() *TransportByAir4 {
	newValue := new(TransportByAir4)
	s.TransportByAir = append(s.TransportByAir, newValue)
	return newValue
}

func (s *SingleTransport8) AddTransportBySea() *TransportBySea5 {
	newValue := new(TransportBySea5)
	s.TransportBySea = append(s.TransportBySea, newValue)
	return newValue
}

func (s *SingleTransport8) AddTransportByRoad() *TransportByRoad4 {
	newValue := new(TransportByRoad4)
	s.TransportByRoad = append(s.TransportByRoad, newValue)
	return newValue
}

func (s *SingleTransport8) AddTransportByRail() *TransportByRail4 {
	newValue := new(TransportByRail4)
	s.TransportByRail = append(s.TransportByRail, newValue)
	return newValue
}
