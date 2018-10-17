package iso20022

// Quantity breakdown information for a specific securities balance.
type SecuritiesSubBalanceTypeAndQuantityBreakdown1 struct {

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	Type *SecuritiesBalanceType3Choice `xml:"Tp"`

	// Breakdown of a quantity into lots such as tax lots, instrument series.
	QuantityBreakdown []*QuantityBreakdown12 `xml:"QtyBrkdwn,omitempty"`
}

func (s *SecuritiesSubBalanceTypeAndQuantityBreakdown1) AddType() *SecuritiesBalanceType3Choice {
	s.Type = new(SecuritiesBalanceType3Choice)
	return s.Type
}

func (s *SecuritiesSubBalanceTypeAndQuantityBreakdown1) AddQuantityBreakdown() *QuantityBreakdown12 {
	newValue := new(QuantityBreakdown12)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}

// Quantity breakdown information for a specific securities balance.
type SecuritiesSubBalanceTypeAndQuantityBreakdown3 struct {

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	Type *SecuritiesBalanceType6Choice `xml:"Tp"`

	// Breakdown of a quantity into lots such as tax lots, instrument series.
	QuantityBreakdown []*QuantityBreakdown32 `xml:"QtyBrkdwn,omitempty"`
}

func (s *SecuritiesSubBalanceTypeAndQuantityBreakdown3) AddType() *SecuritiesBalanceType6Choice {
	s.Type = new(SecuritiesBalanceType6Choice)
	return s.Type
}

func (s *SecuritiesSubBalanceTypeAndQuantityBreakdown3) AddQuantityBreakdown() *QuantityBreakdown32 {
	newValue := new(QuantityBreakdown32)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}

// Quantity breakdown information for a specific securities balance.
type SecuritiesSubBalanceTypeAndQuantityBreakdown4 struct {

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	Type *SecuritiesBalanceType8Choice `xml:"Tp"`

	// Breakdown of a quantity into lots such as tax lots, instrument series.
	QuantityBreakdown []*QuantityBreakdown34 `xml:"QtyBrkdwn,omitempty"`
}

func (s *SecuritiesSubBalanceTypeAndQuantityBreakdown4) AddType() *SecuritiesBalanceType8Choice {
	s.Type = new(SecuritiesBalanceType8Choice)
	return s.Type
}

func (s *SecuritiesSubBalanceTypeAndQuantityBreakdown4) AddQuantityBreakdown() *QuantityBreakdown34 {
	newValue := new(QuantityBreakdown34)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}
