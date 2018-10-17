package iso20022

// Specifies the expected collateral type and direction.
type ExpectedCollateralMovement1 struct {

	// Type of collateral that will be delivered.
	Delivery []*CollateralType1Code `xml:"Dlvry,omitempty"`

	// Type of collateral that will be returned.
	Return []*CollateralType1Code `xml:"Rtr,omitempty"`
}

func (e *ExpectedCollateralMovement1) AddDelivery(value string) {
	e.Delivery = append(e.Delivery, (*CollateralType1Code)(&value))
}

func (e *ExpectedCollateralMovement1) AddReturn(value string) {
	e.Return = append(e.Return, (*CollateralType1Code)(&value))
}

// Specifies the expected collateral type and direction.
type ExpectedCollateralMovement2 struct {

	// Type of collateral that will be delivered and date by which the collateral movement is expected.
	Delivery []*CollateralMovement9 `xml:"Dlvry,omitempty"`

	// Type of collateral that will be returned and date by which the collateral movement is expected.
	Return []*CollateralMovement9 `xml:"Rtr,omitempty"`
}

func (e *ExpectedCollateralMovement2) AddDelivery() *CollateralMovement9 {
	newValue := new(CollateralMovement9)
	e.Delivery = append(e.Delivery, newValue)
	return newValue
}

func (e *ExpectedCollateralMovement2) AddReturn() *CollateralMovement9 {
	newValue := new(CollateralMovement9)
	e.Return = append(e.Return, newValue)
	return newValue
}
