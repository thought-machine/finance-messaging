package iso20022

// Information needed due to regulatory and statutory requirements.
type StructuredRegulatoryReporting2 struct {

	// Specifies the nature, purpose, and reason for the transaction to be reported for regulatory and statutory requirements in a coded form.
	Code *Max3Text `xml:"Cd,omitempty"`

	// Amount of money to be reported for regulatory and statutory requirements.
	Amount *CurrencyAndAmount `xml:"Amt,omitempty"`

	// Additional details that cater for specific domestic regulatory requirements.
	//
	// Usage: Information is used to provide details that are not catered for in the Code or/and Amount elements.
	Information *Max35Text `xml:"Inf,omitempty"`
}

func (s *StructuredRegulatoryReporting2) SetCode(value string) {
	s.Code = (*Max3Text)(&value)
}

func (s *StructuredRegulatoryReporting2) SetAmount(value, currency string) {
	s.Amount = NewCurrencyAndAmount(value, currency)
}

func (s *StructuredRegulatoryReporting2) SetInformation(value string) {
	s.Information = (*Max35Text)(&value)
}

// Information needed due to regulatory and statutory requirements.
type StructuredRegulatoryReporting3 struct {

	// Specifies the type of the information supplied in the regulatory reporting details.
	Type *Max35Text `xml:"Tp,omitempty"`

	// Date related to the specified type of regulatory reporting details.
	Date *ISODate `xml:"Dt,omitempty"`

	// Country related to the specified type of regulatory reporting details.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Specifies the nature, purpose, and reason for the transaction to be reported for regulatory and statutory requirements in a coded form.
	Code *Max10Text `xml:"Cd,omitempty"`

	// Amount of money to be reported for regulatory and statutory requirements.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`

	// Additional details that cater for specific domestic regulatory requirements.
	Information []*Max35Text `xml:"Inf,omitempty"`
}

func (s *StructuredRegulatoryReporting3) SetType(value string) {
	s.Type = (*Max35Text)(&value)
}

func (s *StructuredRegulatoryReporting3) SetDate(value string) {
	s.Date = (*ISODate)(&value)
}

func (s *StructuredRegulatoryReporting3) SetCountry(value string) {
	s.Country = (*CountryCode)(&value)
}

func (s *StructuredRegulatoryReporting3) SetCode(value string) {
	s.Code = (*Max10Text)(&value)
}

func (s *StructuredRegulatoryReporting3) SetAmount(value, currency string) {
	s.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *StructuredRegulatoryReporting3) AddInformation(value string) {
	s.Information = append(s.Information, (*Max35Text)(&value))
}
