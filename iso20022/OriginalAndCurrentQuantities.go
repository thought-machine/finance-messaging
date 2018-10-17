package iso20022

// Original and current value of an asset-back instrument.
type OriginalAndCurrentQuantities1 struct {

	// Quantity expressed as an amount representing the face amount, ie, the principal, of a debt instrument.
	FaceAmount *ImpliedCurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as an amount representing the current amortised face amount of a bond, for example, a periodic reduction/increase of a bond's principal amount.
	AmortisedValue *ImpliedCurrencyAndAmount `xml:"AmtsdVal"`
}

func (o *OriginalAndCurrentQuantities1) SetFaceAmount(value, currency string) {
	o.FaceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (o *OriginalAndCurrentQuantities1) SetAmortisedValue(value, currency string) {
	o.AmortisedValue = NewImpliedCurrencyAndAmount(value, currency)
}

// Signed face amount and amortised value.
type OriginalAndCurrentQuantities2 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Quantity expressed as an amount representing the face amount, that is, the principal, of a debt instrument.
	FaceAmount *ImpliedCurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as an amount representing the current amortised face amount of a bond, for example, a periodic reduction/increase of a bond's principal amount.
	AmortisedValue *ImpliedCurrencyAndAmount `xml:"AmtsdVal"`
}

func (o *OriginalAndCurrentQuantities2) SetShortLongPosition(value string) {
	o.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (o *OriginalAndCurrentQuantities2) SetFaceAmount(value, currency string) {
	o.FaceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (o *OriginalAndCurrentQuantities2) SetAmortisedValue(value, currency string) {
	o.AmortisedValue = NewImpliedCurrencyAndAmount(value, currency)
}

// Original and current value of an asset-back instrument.
type OriginalAndCurrentQuantities4 struct {

	// Quantity expressed as an amount representing the face amount, ie, the principal, of a debt instrument.
	FaceAmount *RestrictedFINImpliedCurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as an amount representing the current amortised face amount of a bond, for example, a periodic reduction/increase of a bond's principal amount.
	AmortisedValue *RestrictedFINImpliedCurrencyAndAmount `xml:"AmtsdVal"`
}

func (o *OriginalAndCurrentQuantities4) SetFaceAmount(value, currency string) {
	o.FaceAmount = NewRestrictedFINImpliedCurrencyAndAmount(value, currency)
}

func (o *OriginalAndCurrentQuantities4) SetAmortisedValue(value, currency string) {
	o.AmortisedValue = NewRestrictedFINImpliedCurrencyAndAmount(value, currency)
}

// Signed face amount and amortised value.
type OriginalAndCurrentQuantities6 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Quantity expressed as an amount representing the face amount, that is, the principal, of a debt instrument.
	FaceAmount *ImpliedCurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as an amount representing the current amortised face amount of a bond, for example, a periodic reduction/increase of a bond's principal amount.
	AmortisedValue *ImpliedCurrencyAndAmount `xml:"AmtsdVal"`
}

func (o *OriginalAndCurrentQuantities6) SetShortLongPosition(value string) {
	o.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (o *OriginalAndCurrentQuantities6) SetFaceAmount(value, currency string) {
	o.FaceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (o *OriginalAndCurrentQuantities6) SetAmortisedValue(value, currency string) {
	o.AmortisedValue = NewImpliedCurrencyAndAmount(value, currency)
}

// Signed face amount and amortised value.
type OriginalAndCurrentQuantities7 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Quantity expressed as an amount representing the face amount, that is, the principal, of a debt instrument.
	FaceAmount *RestrictedFINImpliedCurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as an amount representing the current amortised face amount of a bond, for example, a periodic reduction/increase of a bond's principal amount.
	AmortisedValue *RestrictedFINImpliedCurrencyAndAmount `xml:"AmtsdVal"`
}

func (o *OriginalAndCurrentQuantities7) SetShortLongPosition(value string) {
	o.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (o *OriginalAndCurrentQuantities7) SetFaceAmount(value, currency string) {
	o.FaceAmount = NewRestrictedFINImpliedCurrencyAndAmount(value, currency)
}

func (o *OriginalAndCurrentQuantities7) SetAmortisedValue(value, currency string) {
	o.AmortisedValue = NewRestrictedFINImpliedCurrencyAndAmount(value, currency)
}
