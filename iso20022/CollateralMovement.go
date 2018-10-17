package iso20022

// Provides the agreed amount and the collateral movement direction.
type CollateralMovement10 struct {

	// Provides the call amount that is agreed and that will result in a collateral movement.
	AgreedAmount *ActiveCurrencyAndAmount `xml:"AgrdAmt"`

	// Provides the collateral movement direction that is a delivery and optionally a return, or a return only.
	MovementDirection []*CollateralMovement5Choice `xml:"MvmntDrctn,omitempty"`
}

func (c *CollateralMovement10) SetAgreedAmount(value, currency string) {
	c.AgreedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralMovement10) AddMovementDirection() *CollateralMovement5Choice {
	newValue := new(CollateralMovement5Choice)
	c.MovementDirection = append(c.MovementDirection, newValue)
	return newValue
}

// Provides the collateral movement direction that is a delivery and optionally a return.
type CollateralMovement11 struct {

	// Provides the collateral movement direction that is a delivery only.
	Deliver *Collateral16 `xml:"Dlvr"`

	// Provides the collateral movement direction that is a return only.
	Return *Collateral17 `xml:"Rtr,omitempty"`
}

func (c *CollateralMovement11) AddDeliver() *Collateral16 {
	c.Deliver = new(Collateral16)
	return c.Deliver
}

func (c *CollateralMovement11) AddReturn() *Collateral17 {
	c.Return = new(Collateral17)
	return c.Return
}

// Provides the agreed amount and the collateral movement direction.
type CollateralMovement5 struct {

	// Provides the call amount that is agreed and that will result in a collateral movement.
	AgreedAmount *ActiveCurrencyAndAmount `xml:"AgrdAmt"`

	// Provides the collateral movement direction that is a delivery and optionaly a return, or a return only.
	MovementDirection []*CollateralMovement3Choice `xml:"MvmntDrctn,omitempty"`
}

func (c *CollateralMovement5) SetAgreedAmount(value, currency string) {
	c.AgreedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralMovement5) AddMovementDirection() *CollateralMovement3Choice {
	newValue := new(CollateralMovement3Choice)
	c.MovementDirection = append(c.MovementDirection, newValue)
	return newValue
}

// Provides the collateral movement direction that is a delivery and optionaly a return.
type CollateralMovement6 struct {

	// Provides the collateral movement direction that is a delivery only.
	Deliver *Collateral8 `xml:"Dlvr"`

	// Provides the collateral movement direction that is a return only.
	Return *Collateral7 `xml:"Rtr,omitempty"`
}

func (c *CollateralMovement6) AddDeliver() *Collateral8 {
	c.Deliver = new(Collateral8)
	return c.Deliver
}

func (c *CollateralMovement6) AddReturn() *Collateral7 {
	c.Return = new(Collateral7)
	return c.Return
}

// Provides the agreed amount and the collateral movement direction.
type CollateralMovement7 struct {

	// Provides the call amount that is agreed and that will result in a collateral movement.
	AgreedAmount *ActiveCurrencyAndAmount `xml:"AgrdAmt"`

	// Provides the collateral movement direction that is a delivery and optionaly a return, or a return only.
	MovementDirection []*CollateralMovement4Choice `xml:"MvmntDrctn,omitempty"`
}

func (c *CollateralMovement7) SetAgreedAmount(value, currency string) {
	c.AgreedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralMovement7) AddMovementDirection() *CollateralMovement4Choice {
	newValue := new(CollateralMovement4Choice)
	c.MovementDirection = append(c.MovementDirection, newValue)
	return newValue
}

// Provides the collateral movement direction that is a delivery and optionaly a return.
type CollateralMovement8 struct {

	// Provides the collateral movement direction that is a delivery only.
	Deliver *Collateral12 `xml:"Dlvr"`

	// Provides the collateral movement direction that is a return only.
	Return *Collateral11 `xml:"Rtr,omitempty"`
}

func (c *CollateralMovement8) AddDeliver() *Collateral12 {
	c.Deliver = new(Collateral12)
	return c.Deliver
}

func (c *CollateralMovement8) AddReturn() *Collateral11 {
	c.Return = new(Collateral11)
	return c.Return
}

// Specifies the type of collateral that will be delivered and the date to be reported.
type CollateralMovement9 struct {

	// Specifies the type of collateral.
	CollateralType *CollateralType1Code `xml:"CollTp"`

	// Date by which the collateral movement must be executed.
	Date *ISODate `xml:"Dt,omitempty"`
}

func (c *CollateralMovement9) SetCollateralType(value string) {
	c.CollateralType = (*CollateralType1Code)(&value)
}

func (c *CollateralMovement9) SetDate(value string) {
	c.Date = (*ISODate)(&value)
}
