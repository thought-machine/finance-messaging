package iso20022

// Set of elements to identify a proprietary quantity.
type ProprietaryQuantity1 struct {

	// Identifies the type of proprietary quantity reported.
	Type *Max35Text `xml:"Tp"`

	// Provides the proprietary quantity in free format.
	Quantity *Max35Text `xml:"Qty"`
}

func (p *ProprietaryQuantity1) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryQuantity1) SetQuantity(value string) {
	p.Quantity = (*Max35Text)(&value)
}

// Proprietary quantity format.
type ProprietaryQuantity10 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos,omitempty"`

	// Provides the proprietary quantity with a decimal number.
	Quantity *RestrictedFINDecimalNumber `xml:"Qty"`

	// Identifies the type of proprietary quantity reported.
	QuantityType *Exact4AlphaNumericText `xml:"QtyTp"`

	// Provides information related to issuer in free format.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Name of the identification scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (p *ProprietaryQuantity10) SetShortLongPosition(value string) {
	p.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (p *ProprietaryQuantity10) SetQuantity(value string) {
	p.Quantity = (*RestrictedFINDecimalNumber)(&value)
}

func (p *ProprietaryQuantity10) SetQuantityType(value string) {
	p.QuantityType = (*Exact4AlphaNumericText)(&value)
}

func (p *ProprietaryQuantity10) SetIssuer(value string) {
	p.Issuer = (*Max4AlphaNumericText)(&value)
}

func (p *ProprietaryQuantity10) SetSchemeName(value string) {
	p.SchemeName = (*Max4AlphaNumericText)(&value)
}

// Proprietary quantity format.
type ProprietaryQuantity2 struct {

	// Provides the proprietary quantity with a decimal number.
	Quantity *DecimalNumber `xml:"Qty"`

	// Identifies the type of proprietary quantity reported.
	QuantityType *Exact4AlphaNumericText `xml:"QtyTp"`

	// Provides information related to issuer in free format.
	Issuer *Max35Text `xml:"Issr"`

	// Name of the identification scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (p *ProprietaryQuantity2) SetQuantity(value string) {
	p.Quantity = (*DecimalNumber)(&value)
}

func (p *ProprietaryQuantity2) SetQuantityType(value string) {
	p.QuantityType = (*Exact4AlphaNumericText)(&value)
}

func (p *ProprietaryQuantity2) SetIssuer(value string) {
	p.Issuer = (*Max35Text)(&value)
}

func (p *ProprietaryQuantity2) SetSchemeName(value string) {
	p.SchemeName = (*Max35Text)(&value)
}

// Proprietary quantity format.
type ProprietaryQuantity3 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos,omitempty"`

	// Provides the proprietary quantity with a decimal number.
	Quantity *DecimalNumber `xml:"Qty"`

	// Identifies the type of proprietary quantity reported.
	QuantityType *Exact4AlphaNumericText `xml:"QtyTp"`

	// Provides information related to issuer in free format.
	Issuer *Max35Text `xml:"Issr"`

	// Name of the identification scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (p *ProprietaryQuantity3) SetShortLongPosition(value string) {
	p.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (p *ProprietaryQuantity3) SetQuantity(value string) {
	p.Quantity = (*DecimalNumber)(&value)
}

func (p *ProprietaryQuantity3) SetQuantityType(value string) {
	p.QuantityType = (*Exact4AlphaNumericText)(&value)
}

func (p *ProprietaryQuantity3) SetIssuer(value string) {
	p.Issuer = (*Max35Text)(&value)
}

func (p *ProprietaryQuantity3) SetSchemeName(value string) {
	p.SchemeName = (*Max35Text)(&value)
}

// Proprietary quantity format.
type ProprietaryQuantity7 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos,omitempty"`

	// Provides the proprietary quantity with a decimal number.
	Quantity *DecimalNumber `xml:"Qty"`

	// Identifies the type of proprietary quantity reported.
	QuantityType *Exact4AlphaNumericText `xml:"QtyTp"`

	// Provides information related to issuer in free format.
	Issuer *Max35Text `xml:"Issr"`

	// Name of the identification scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (p *ProprietaryQuantity7) SetShortLongPosition(value string) {
	p.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (p *ProprietaryQuantity7) SetQuantity(value string) {
	p.Quantity = (*DecimalNumber)(&value)
}

func (p *ProprietaryQuantity7) SetQuantityType(value string) {
	p.QuantityType = (*Exact4AlphaNumericText)(&value)
}

func (p *ProprietaryQuantity7) SetIssuer(value string) {
	p.Issuer = (*Max35Text)(&value)
}

func (p *ProprietaryQuantity7) SetSchemeName(value string) {
	p.SchemeName = (*Max35Text)(&value)
}

// Proprietary quantity format.
type ProprietaryQuantity8 struct {

	// Provides the proprietary quantity with a decimal number.
	Quantity *DecimalNumber `xml:"Qty"`

	// Identifies the type of proprietary quantity reported.
	QuantityType *Exact4AlphaNumericText `xml:"QtyTp"`

	// Provides information related to issuer in free format.
	Issuer *Max35Text `xml:"Issr"`

	// Name of the identification scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (p *ProprietaryQuantity8) SetQuantity(value string) {
	p.Quantity = (*DecimalNumber)(&value)
}

func (p *ProprietaryQuantity8) SetQuantityType(value string) {
	p.QuantityType = (*Exact4AlphaNumericText)(&value)
}

func (p *ProprietaryQuantity8) SetIssuer(value string) {
	p.Issuer = (*Max35Text)(&value)
}

func (p *ProprietaryQuantity8) SetSchemeName(value string) {
	p.SchemeName = (*Max35Text)(&value)
}

// Proprietary quantity format.
type ProprietaryQuantity9 struct {

	// Provides the proprietary quantity with a decimal number.
	Quantity *RestrictedFINDecimalNumber `xml:"Qty"`

	// Identifies the type of proprietary quantity reported.
	QuantityType *Exact4AlphaNumericText `xml:"QtyTp"`

	// Provides information related to issuer in free format.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Name of the identification scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (p *ProprietaryQuantity9) SetQuantity(value string) {
	p.Quantity = (*RestrictedFINDecimalNumber)(&value)
}

func (p *ProprietaryQuantity9) SetQuantityType(value string) {
	p.QuantityType = (*Exact4AlphaNumericText)(&value)
}

func (p *ProprietaryQuantity9) SetIssuer(value string) {
	p.Issuer = (*Max4AlphaNumericText)(&value)
}

func (p *ProprietaryQuantity9) SetSchemeName(value string) {
	p.SchemeName = (*Max4AlphaNumericText)(&value)
}
