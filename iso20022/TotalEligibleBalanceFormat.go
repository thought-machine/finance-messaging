package iso20022

// Total eligible balance for the corporate action and full and part way period units.
type TotalEligibleBalanceFormat1 struct {

	// Provides information about balance related to a corporate action.
	Balance *Quantity3Choice `xml:"Bal,omitempty"`

	// Number of units of a fund that were purchased in a previous distribution period and/or held at the beginning of a distribution period, for example Group I Units in the UK.
	FullPeriodUnits *SignedQuantityFormat2 `xml:"FullPrdUnits,omitempty"`

	// Number of units of a fund that were purchased part way throughout a distribution period, for example Group II Units in the U.K.
	PartWayPeriodUnits *SignedQuantityFormat2 `xml:"PartWayPrdUnits,omitempty"`
}

func (t *TotalEligibleBalanceFormat1) AddBalance() *Quantity3Choice {
	t.Balance = new(Quantity3Choice)
	return t.Balance
}

func (t *TotalEligibleBalanceFormat1) AddFullPeriodUnits() *SignedQuantityFormat2 {
	t.FullPeriodUnits = new(SignedQuantityFormat2)
	return t.FullPeriodUnits
}

func (t *TotalEligibleBalanceFormat1) AddPartWayPeriodUnits() *SignedQuantityFormat2 {
	t.PartWayPeriodUnits = new(SignedQuantityFormat2)
	return t.PartWayPeriodUnits
}

// Total eligible balance for the corporate action and full and part way period units.
type TotalEligibleBalanceFormat8 struct {

	// Provides information about balance related to a corporate action.
	Balance *Quantity17Choice `xml:"Bal,omitempty"`

	// Number of units of a fund that were purchased in a previous distribution period and/or held at the beginning of a distribution period, for example Group I Units in the UK.
	FullPeriodUnits *SignedQuantityFormat6 `xml:"FullPrdUnits,omitempty"`

	// Number of units of a fund that were purchased part way throughout a distribution period, for example Group II Units in the U.K.
	PartWayPeriodUnits *SignedQuantityFormat6 `xml:"PartWayPrdUnits,omitempty"`
}

func (t *TotalEligibleBalanceFormat8) AddBalance() *Quantity17Choice {
	t.Balance = new(Quantity17Choice)
	return t.Balance
}

func (t *TotalEligibleBalanceFormat8) AddFullPeriodUnits() *SignedQuantityFormat6 {
	t.FullPeriodUnits = new(SignedQuantityFormat6)
	return t.FullPeriodUnits
}

func (t *TotalEligibleBalanceFormat8) AddPartWayPeriodUnits() *SignedQuantityFormat6 {
	t.PartWayPeriodUnits = new(SignedQuantityFormat6)
	return t.PartWayPeriodUnits
}

// Total eligible balance for the corporate action and full and part way period units.
type TotalEligibleBalanceFormat9 struct {

	// Provides information about balance related to a corporate action.
	Balance *Quantity22Choice `xml:"Bal,omitempty"`

	// Number of units of a fund that were purchased in a previous distribution period and/or held at the beginning of a distribution period, for example Group I Units in the UK.
	FullPeriodUnits *SignedQuantityFormat9 `xml:"FullPrdUnits,omitempty"`

	// Number of units of a fund that were purchased part way throughout a distribution period, for example Group II Units in the U.K.
	PartWayPeriodUnits *SignedQuantityFormat9 `xml:"PartWayPrdUnits,omitempty"`
}

func (t *TotalEligibleBalanceFormat9) AddBalance() *Quantity22Choice {
	t.Balance = new(Quantity22Choice)
	return t.Balance
}

func (t *TotalEligibleBalanceFormat9) AddFullPeriodUnits() *SignedQuantityFormat9 {
	t.FullPeriodUnits = new(SignedQuantityFormat9)
	return t.FullPeriodUnits
}

func (t *TotalEligibleBalanceFormat9) AddPartWayPeriodUnits() *SignedQuantityFormat9 {
	t.PartWayPeriodUnits = new(SignedQuantityFormat9)
	return t.PartWayPeriodUnits
}
