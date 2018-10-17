package iso20022

// Specifies the quantity of a product in a trade transaction.
type Quantity10 struct {

	// Specifies a unit of measure with a code or free text.
	UnitOfMeasure *UnitOfMeasure3Choice `xml:"UnitOfMeasr"`

	// Quantity of a product on a line specified by a number. For example, 100 (kgs), 50 (pieces).
	Value *DecimalNumber `xml:"Val"`
}

func (q *Quantity10) AddUnitOfMeasure() *UnitOfMeasure3Choice {
	q.UnitOfMeasure = new(UnitOfMeasure3Choice)
	return q.UnitOfMeasure
}

func (q *Quantity10) SetValue(value string) {
	q.Value = (*DecimalNumber)(&value)
}

// Details on the quantity, account and other related information involved in a transaction.
type Quantity11 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate4 `xml:"CertNb,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown30 `xml:"QtyBrkdwn,omitempty"`
}

func (q *Quantity11) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *Quantity11) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *Quantity11) AddCertificateNumber() *SecuritiesCertificate4 {
	newValue := new(SecuritiesCertificate4)
	q.CertificateNumber = append(q.CertificateNumber, newValue)
	return newValue
}

func (q *Quantity11) AddQuantityBreakdown() *QuantityBreakdown30 {
	newValue := new(QuantityBreakdown30)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type Quantity12 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity15Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *RestrictedFINXMax210Text `xml:"DnmtnChc,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate5 `xml:"CertNb,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown38 `xml:"QtyBrkdwn,omitempty"`
}

func (q *Quantity12) AddSettlementQuantity() *FinancialInstrumentQuantity15Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity15Choice)
	return q.SettlementQuantity
}

func (q *Quantity12) SetDenominationChoice(value string) {
	q.DenominationChoice = (*RestrictedFINXMax210Text)(&value)
}

func (q *Quantity12) AddCertificateNumber() *SecuritiesCertificate5 {
	newValue := new(SecuritiesCertificate5)
	q.CertificateNumber = append(q.CertificateNumber, newValue)
	return newValue
}

func (q *Quantity12) AddQuantityBreakdown() *QuantityBreakdown38 {
	newValue := new(QuantityBreakdown38)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Specifies the quantity of a product in a trade transaction.
type Quantity3 struct {

	// Specifies the unit of measurement. For example, kilo, tons.
	UnitOfMeasureCode *UnitOfMeasure4Code `xml:"UnitOfMeasrCd"`

	// Identifies the unit of measure not present in the code list.
	OtherUnitOfMeasure *Max35Text `xml:"OthrUnitOfMeasr"`

	// Quantity of a product on a line specified by a number. For example, 100 (kgs), 50 (pieces).
	Value *DecimalNumber `xml:"Val"`
}

func (q *Quantity3) SetUnitOfMeasureCode(value string) {
	q.UnitOfMeasureCode = (*UnitOfMeasure4Code)(&value)
}

func (q *Quantity3) SetOtherUnitOfMeasure(value string) {
	q.OtherUnitOfMeasure = (*Max35Text)(&value)
}

func (q *Quantity3) SetValue(value string) {
	q.Value = (*DecimalNumber)(&value)
}

// Specifies the quantity of a product in a trade transaction.
type Quantity4 struct {

	// Specifies the unit of measurement. For example, kilo, tons.
	UnitOfMeasureCode *UnitOfMeasure4Code `xml:"UnitOfMeasrCd"`

	// Identifies the unit of measure not present in the code list.
	OtherUnitOfMeasure *Max35Text `xml:"OthrUnitOfMeasr"`

	// Quantity of a product on a line specified by a number. For example, 100 (kgs), 50 (pieces).
	Value *DecimalNumber `xml:"Val"`

	// Multiplication factor of measurement values. For example: goods that can be ordered by 36 pieces.
	Factor *Max15NumericText `xml:"Fctr,omitempty"`
}

func (q *Quantity4) SetUnitOfMeasureCode(value string) {
	q.UnitOfMeasureCode = (*UnitOfMeasure4Code)(&value)
}

func (q *Quantity4) SetOtherUnitOfMeasure(value string) {
	q.OtherUnitOfMeasure = (*Max35Text)(&value)
}

func (q *Quantity4) SetValue(value string) {
	q.Value = (*DecimalNumber)(&value)
}

func (q *Quantity4) SetFactor(value string) {
	q.Factor = (*Max15NumericText)(&value)
}

// Details on the quantity, account and other related information involved in a transaction.
type Quantity5 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate1 `xml:"CertNb,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`
}

func (q *Quantity5) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *Quantity5) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *Quantity5) AddCertificateNumber() *SecuritiesCertificate1 {
	newValue := new(SecuritiesCertificate1)
	q.CertificateNumber = append(q.CertificateNumber, newValue)
	return newValue
}

func (q *Quantity5) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Details on the quantity, account and other related information involved in a transaction.
type Quantity7 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate1 `xml:"CertNb,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown13 `xml:"QtyBrkdwn,omitempty"`
}

func (q *Quantity7) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *Quantity7) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *Quantity7) AddCertificateNumber() *SecuritiesCertificate1 {
	newValue := new(SecuritiesCertificate1)
	q.CertificateNumber = append(q.CertificateNumber, newValue)
	return newValue
}

func (q *Quantity7) AddQuantityBreakdown() *QuantityBreakdown13 {
	newValue := new(QuantityBreakdown13)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}

// Specifies the quantity of a product in a trade transaction.
type Quantity9 struct {

	// Specifies a unit of measure with a code or free text.
	UnitOfMeasure *UnitOfMeasure3Choice `xml:"UnitOfMeasr"`

	// Quantity of a product on a line specified by a number. For example, 100 (kgs), 50 (pieces).
	Value *DecimalNumber `xml:"Val"`

	// Multiplication factor of measurement values. For example: goods that can be ordered by 36 pieces.
	Factor *Max15NumericText `xml:"Fctr,omitempty"`
}

func (q *Quantity9) AddUnitOfMeasure() *UnitOfMeasure3Choice {
	q.UnitOfMeasure = new(UnitOfMeasure3Choice)
	return q.UnitOfMeasure
}

func (q *Quantity9) SetValue(value string) {
	q.Value = (*DecimalNumber)(&value)
}

func (q *Quantity9) SetFactor(value string) {
	q.Factor = (*Max15NumericText)(&value)
}
