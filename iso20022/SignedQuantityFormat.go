package iso20022

// Signed quantity of security formats.
type SignedQuantityFormat1 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Choice between different quantity of security formats.
	QuantityChoice *Quantity2Choice `xml:"QtyChc"`
}

func (s *SignedQuantityFormat1) SetShortLongPosition(value string) {
	s.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (s *SignedQuantityFormat1) AddQuantityChoice() *Quantity2Choice {
	s.QuantityChoice = new(Quantity2Choice)
	return s.QuantityChoice
}

// Signed quantity of security formats.
type SignedQuantityFormat2 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Quantity of security.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`
}

func (s *SignedQuantityFormat2) SetShortLongPosition(value string) {
	s.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (s *SignedQuantityFormat2) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

// Signed quantity of security formats.
type SignedQuantityFormat6 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Quantity of security.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`
}

func (s *SignedQuantityFormat6) SetShortLongPosition(value string) {
	s.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (s *SignedQuantityFormat6) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

// Signed quantity of security formats.
type SignedQuantityFormat7 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Choice between different quantity of security formats.
	QuantityChoice *Quantity19Choice `xml:"QtyChc"`
}

func (s *SignedQuantityFormat7) SetShortLongPosition(value string) {
	s.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (s *SignedQuantityFormat7) AddQuantityChoice() *Quantity19Choice {
	s.QuantityChoice = new(Quantity19Choice)
	return s.QuantityChoice
}

// Signed quantity of security formats.
type SignedQuantityFormat8 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Choice between different quantity of security formats.
	QuantityChoice *Quantity21Choice `xml:"QtyChc"`
}

func (s *SignedQuantityFormat8) SetShortLongPosition(value string) {
	s.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (s *SignedQuantityFormat8) AddQuantityChoice() *Quantity21Choice {
	s.QuantityChoice = new(Quantity21Choice)
	return s.QuantityChoice
}

// Signed quantity of security formats.
type SignedQuantityFormat9 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Quantity of security.
	Quantity *FinancialInstrumentQuantity15Choice `xml:"Qty"`
}

func (s *SignedQuantityFormat9) SetShortLongPosition(value string) {
	s.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (s *SignedQuantityFormat9) AddQuantity() *FinancialInstrumentQuantity15Choice {
	s.Quantity = new(FinancialInstrumentQuantity15Choice)
	return s.Quantity
}
