package iso20022

// Place where the securities are safe-kept, physically or notionally. This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
type SafeKeepingPlace1 struct {

	// Unique identification of the party.
	SafekeepingPlaceFormat *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlcFrmt,omitempty"`

	// Legal entity identification as an alternate identification for a place of safekeeping.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (s *SafeKeepingPlace1) AddSafekeepingPlaceFormat() *SafekeepingPlaceFormat10Choice {
	s.SafekeepingPlaceFormat = new(SafekeepingPlaceFormat10Choice)
	return s.SafekeepingPlaceFormat
}

func (s *SafeKeepingPlace1) SetLEI(value string) {
	s.LEI = (*LEIIdentifier)(&value)
}

// Place where the securities are safe-kept, physically or notionally. This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
type SafeKeepingPlace2 struct {

	// Unique identification of the party.
	SafekeepingPlaceFormat *SafekeepingPlaceFormat17Choice `xml:"SfkpgPlcFrmt,omitempty"`

	// Legal entity identification as an alternate identification for a place of safekeeping.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (s *SafeKeepingPlace2) AddSafekeepingPlaceFormat() *SafekeepingPlaceFormat17Choice {
	s.SafekeepingPlaceFormat = new(SafekeepingPlaceFormat17Choice)
	return s.SafekeepingPlaceFormat
}

func (s *SafeKeepingPlace2) SetLEI(value string) {
	s.LEI = (*LEIIdentifier)(&value)
}
