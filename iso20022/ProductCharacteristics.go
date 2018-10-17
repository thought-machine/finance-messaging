package iso20022

// Identifies the characteristic of a product.
type ProductCharacteristics1 struct {

	// Specifies the type of product characteristic by means of a code.
	Type *ProductCharacteristics1Code `xml:"Tp"`

	// Specifies the characteristic of a product.
	Characteristics *Max35Text `xml:"Chrtcs"`
}

func (p *ProductCharacteristics1) SetType(value string) {
	p.Type = (*ProductCharacteristics1Code)(&value)
}

func (p *ProductCharacteristics1) SetCharacteristics(value string) {
	p.Characteristics = (*Max35Text)(&value)
}

// Product characteristic applicable to this trade product.
type ProductCharacteristics2 struct {

	// Characteristics of the product.
	Characteristic *ProductCharacteristics1Choice `xml:"Chrtc,omitempty"`

	// Measurement value for this product characteristic.
	ValueMeasure *Quantity3 `xml:"ValMeasr,omitempty"`
}

func (p *ProductCharacteristics2) AddCharacteristic() *ProductCharacteristics1Choice {
	p.Characteristic = new(ProductCharacteristics1Choice)
	return p.Characteristic
}

func (p *ProductCharacteristics2) AddValueMeasure() *Quantity3 {
	p.ValueMeasure = new(Quantity3)
	return p.ValueMeasure
}

// Product characteristic applicable to this trade product.
type ProductCharacteristics3 struct {

	// Characteristics of the product.
	Characteristic *ProductCharacteristics1Choice `xml:"Chrtc,omitempty"`

	// Measurement value for this product characteristic.
	ValueMeasure *Quantity10 `xml:"ValMeasr,omitempty"`
}

func (p *ProductCharacteristics3) AddCharacteristic() *ProductCharacteristics1Choice {
	p.Characteristic = new(ProductCharacteristics1Choice)
	return p.Characteristic
}

func (p *ProductCharacteristics3) AddValueMeasure() *Quantity10 {
	p.ValueMeasure = new(Quantity10)
	return p.ValueMeasure
}
