package iso20022

// Information about the units to settle.
type Unit1 struct {

	// Total quantity of securities to be settled.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Date upon which the investor purchased the financial instrument.
	AcquisitionDate *ISODate `xml:"AcqstnDt,omitempty"`

	// Certificate representing the security that is delivered.
	CertificateNumber []*Max35Text `xml:"CertNb,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Information related to the price of the transferred units.
	PriceDetails *UnitPrice3 `xml:"PricDtls,omitempty"`
}

func (u *Unit1) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	u.UnitsNumber = new(FinancialInstrumentQuantity1)
	return u.UnitsNumber
}

func (u *Unit1) SetAcquisitionDate(value string) {
	u.AcquisitionDate = (*ISODate)(&value)
}

func (u *Unit1) AddCertificateNumber(value string) {
	u.CertificateNumber = append(u.CertificateNumber, (*Max35Text)(&value))
}

func (u *Unit1) SetGroup1Or2Units(value string) {
	u.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (u *Unit1) AddPriceDetails() *UnitPrice3 {
	u.PriceDetails = new(UnitPrice3)
	return u.PriceDetails
}

// Information about the units to settle.
type Unit3 struct {

	// Total quantity of securities to be settled.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Date upon which the investor purchased the financial instrument.
	AcquisitionDate *ISODate `xml:"AcqstnDt,omitempty"`

	// Certificate representing the security that is delivered.
	CertificateNumber []*Max35Text `xml:"CertNb,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Information related to the price of the transferred units.
	PriceDetails *UnitPrice12 `xml:"PricDtls,omitempty"`
}

func (u *Unit3) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	u.UnitsNumber = new(FinancialInstrumentQuantity1)
	return u.UnitsNumber
}

func (u *Unit3) SetAcquisitionDate(value string) {
	u.AcquisitionDate = (*ISODate)(&value)
}

func (u *Unit3) AddCertificateNumber(value string) {
	u.CertificateNumber = append(u.CertificateNumber, (*Max35Text)(&value))
}

func (u *Unit3) SetGroup1Or2Units(value string) {
	u.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (u *Unit3) AddPriceDetails() *UnitPrice12 {
	u.PriceDetails = new(UnitPrice12)
	return u.PriceDetails
}

// Quantity expressed as a number and its details.
type Unit4 struct {

	// Quantity expressed as a number, for example, a number of shares.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitDetails []*Unit5 `xml:"UnitDtls,omitempty"`
}

func (u *Unit4) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	u.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return u.TotalUnitsNumber
}

func (u *Unit4) AddUnitDetails() *Unit5 {
	newValue := new(Unit5)
	u.UnitDetails = append(u.UnitDetails, newValue)
	return newValue
}

// Quantity expressed as a number and its details.
type Unit5 struct {

	// Quantity expressed as a number, for example, a number of shares.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units"`
}

func (u *Unit5) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	u.UnitsNumber = new(FinancialInstrumentQuantity1)
	return u.UnitsNumber
}

func (u *Unit5) SetGroup1Or2Units(value string) {
	u.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

// Information about the units to settle.
type Unit6 struct {

	// Total quantity of securities to be settled.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Date upon which the investor purchased the financial instrument.
	AcquisitionDate *ISODate `xml:"AcqstnDt,omitempty"`

	// Certificate representing the security that is delivered.
	CertificateNumber []*Max35Text `xml:"CertNb,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Information related to the price of the transferred units.
	PriceDetails *UnitPrice21 `xml:"PricDtls,omitempty"`
}

func (u *Unit6) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	u.UnitsNumber = new(FinancialInstrumentQuantity1)
	return u.UnitsNumber
}

func (u *Unit6) SetAcquisitionDate(value string) {
	u.AcquisitionDate = (*ISODate)(&value)
}

func (u *Unit6) AddCertificateNumber(value string) {
	u.CertificateNumber = append(u.CertificateNumber, (*Max35Text)(&value))
}

func (u *Unit6) SetGroup1Or2Units(value string) {
	u.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (u *Unit6) AddPriceDetails() *UnitPrice21 {
	u.PriceDetails = new(UnitPrice21)
	return u.PriceDetails
}
