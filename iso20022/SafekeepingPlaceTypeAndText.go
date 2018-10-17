package iso20022

// Identification of the place of safekeeping expressed as a code and a narrative description.
type SafekeepingPlaceTypeAndText1 struct {

	// Place of safekeeping as a code.
	SafekeepingPlaceType *SafekeepingPlace3Code `xml:"SfkpgPlcTp"`

	// Place of safekeeping.
	Identification *Max35Text `xml:"Id,omitempty"`
}

func (s *SafekeepingPlaceTypeAndText1) SetSafekeepingPlaceType(value string) {
	s.SafekeepingPlaceType = (*SafekeepingPlace3Code)(&value)
}

func (s *SafekeepingPlaceTypeAndText1) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

// Identification of the place of safekeeping expressed as a code and a narrative description.
type SafekeepingPlaceTypeAndText15 struct {

	// Place of safekeeping as a code.
	SafekeepingPlaceType *SafekeepingPlace3Code `xml:"SfkpgPlcTp"`

	// Additional information about the place of safekeeping.
	Identification *RestrictedFINXMax30Text `xml:"Id,omitempty"`
}

func (s *SafekeepingPlaceTypeAndText15) SetSafekeepingPlaceType(value string) {
	s.SafekeepingPlaceType = (*SafekeepingPlace3Code)(&value)
}

func (s *SafekeepingPlaceTypeAndText15) SetIdentification(value string) {
	s.Identification = (*RestrictedFINXMax30Text)(&value)
}

// Identification of the place of safekeeping expressed as a code and a narrative description.
type SafekeepingPlaceTypeAndText2 struct {

	// Place of safekeeping as a code.
	SafekeepingPlaceType *SafekeepingPlace2Code `xml:"SfkpgPlcTp"`

	// Additional information about the place of safekeeping.
	Identification *Max35Text `xml:"Id,omitempty"`
}

func (s *SafekeepingPlaceTypeAndText2) SetSafekeepingPlaceType(value string) {
	s.SafekeepingPlaceType = (*SafekeepingPlace2Code)(&value)
}

func (s *SafekeepingPlaceTypeAndText2) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

// Identification of the place of safekeeping expressed as a code and a narrative description.
type SafekeepingPlaceTypeAndText25 struct {

	// Place of safekeeping as a code.
	SafekeepingPlaceType *SafekeepingPlace2Code `xml:"SfkpgPlcTp"`

	// Additional information about the place of safekeeping.
	Identification *RestrictedFINXMax30Text `xml:"Id,omitempty"`
}

func (s *SafekeepingPlaceTypeAndText25) SetSafekeepingPlaceType(value string) {
	s.SafekeepingPlaceType = (*SafekeepingPlace2Code)(&value)
}

func (s *SafekeepingPlaceTypeAndText25) SetIdentification(value string) {
	s.Identification = (*RestrictedFINXMax30Text)(&value)
}

// Identification of the place of safekeeping expressed as a code and a narrative description.
type SafekeepingPlaceTypeAndText3 struct {

	// Place of safekeeping as a code.
	SafekeepingPlaceType *SafekeepingPlace3Code `xml:"SfkpgPlcTp"`

	// Additional information about the place of safekeeping.
	Identification *Max35Text `xml:"Id,omitempty"`
}

func (s *SafekeepingPlaceTypeAndText3) SetSafekeepingPlaceType(value string) {
	s.SafekeepingPlaceType = (*SafekeepingPlace3Code)(&value)
}

func (s *SafekeepingPlaceTypeAndText3) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

// Identification of the place of safekeeping expressed as a code and a narrative description.
type SafekeepingPlaceTypeAndText6 struct {

	// Place of safekeeping as a code.
	SafekeepingPlaceType *SafekeepingPlace2Code `xml:"SfkpgPlcTp"`

	// Additional information about the place of safekeeping.
	Identification *Max35Text `xml:"Id,omitempty"`
}

func (s *SafekeepingPlaceTypeAndText6) SetSafekeepingPlaceType(value string) {
	s.SafekeepingPlaceType = (*SafekeepingPlace2Code)(&value)
}

func (s *SafekeepingPlaceTypeAndText6) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

// Identification of the place of safekeeping expressed as a code and a narrative description.
type SafekeepingPlaceTypeAndText8 struct {

	// Place of safekeeping as a code.
	SafekeepingPlaceType *SafekeepingPlace3Code `xml:"SfkpgPlcTp"`

	// Additional information about the place of safekeeping.
	Identification *Max35Text `xml:"Id,omitempty"`
}

func (s *SafekeepingPlaceTypeAndText8) SetSafekeepingPlaceType(value string) {
	s.SafekeepingPlaceType = (*SafekeepingPlace3Code)(&value)
}

func (s *SafekeepingPlaceTypeAndText8) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

// Identification of the place of safekeeping expressed as a code and a narrative description.
type SafekeepingPlaceTypeAndText9 struct {

	// Place of safekeeping as a code.
	SafekeepingPlaceType *SafekeepingPlace2Code `xml:"SfkpgPlcTp"`

	// Additional information about the place of safekeeping.
	Identification *RestrictedFINXMax30Text `xml:"Id,omitempty"`
}

func (s *SafekeepingPlaceTypeAndText9) SetSafekeepingPlaceType(value string) {
	s.SafekeepingPlaceType = (*SafekeepingPlace2Code)(&value)
}

func (s *SafekeepingPlaceTypeAndText9) SetIdentification(value string) {
	s.Identification = (*RestrictedFINXMax30Text)(&value)
}
