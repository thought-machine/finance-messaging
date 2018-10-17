package iso20022

// Taxable service charge amount conversions to host currency.
type BillingServicesAmount1 struct {

	// Sum of all the individual taxes on the service expressed in the host currency.
	HostAmount *AmountAndDirection34 `xml:"HstAmt"`

	// Amount of the tax obligation expressed in the tax region's pricing currency.
	// Usage: This is the same amount as carried in the host amount but allows the sender to optionally express the value in the pricing currency.
	PricingAmount *AmountAndDirection34 `xml:"PricgAmt,omitempty"`
}

func (b *BillingServicesAmount1) AddHostAmount() *AmountAndDirection34 {
	b.HostAmount = new(AmountAndDirection34)
	return b.HostAmount
}

func (b *BillingServicesAmount1) AddPricingAmount() *AmountAndDirection34 {
	b.PricingAmount = new(AmountAndDirection34)
	return b.PricingAmount
}

// Taxable service charge amount conversions to host currency.
type BillingServicesAmount2 struct {

	// Sum of the original charge host amount and the service tax host amount values. It represents the total charge for a service (including taxes) expressed in the host currency.
	HostAmount *AmountAndDirection34 `xml:"HstAmt"`

	// Sum of the original charge host amount and the service tax host amount values but expressed in the settlement currency.
	SettlementAmount *AmountAndDirection34 `xml:"SttlmAmt,omitempty"`

	// Sum of the original charge host amount and the service tax host amount values but expressed in the pricing currency.
	PricingAmount *AmountAndDirection34 `xml:"PricgAmt,omitempty"`
}

func (b *BillingServicesAmount2) AddHostAmount() *AmountAndDirection34 {
	b.HostAmount = new(AmountAndDirection34)
	return b.HostAmount
}

func (b *BillingServicesAmount2) AddSettlementAmount() *AmountAndDirection34 {
	b.SettlementAmount = new(AmountAndDirection34)
	return b.SettlementAmount
}

func (b *BillingServicesAmount2) AddPricingAmount() *AmountAndDirection34 {
	b.PricingAmount = new(AmountAndDirection34)
	return b.PricingAmount
}

// Taxable service charge amount conversions to host currency.
type BillingServicesAmount3 struct {

	// Represents the total of all taxable services in a specific tax region for a specific currency.  For example, all taxable services for a tax region in Euro would be totalled here in the Euro currency.
	SourceAmount *AmountAndDirection34 `xml:"SrcAmt"`

	// Represents the total of all taxable services in a specific tax region for a specific currency and is a one-to-one relationship with total taxable charge of services, but represented in the host currency after conversion.
	HostAmount *AmountAndDirection34 `xml:"HstAmt"`
}

func (b *BillingServicesAmount3) AddSourceAmount() *AmountAndDirection34 {
	b.SourceAmount = new(AmountAndDirection34)
	return b.SourceAmount
}

func (b *BillingServicesAmount3) AddHostAmount() *AmountAndDirection34 {
	b.HostAmount = new(AmountAndDirection34)
	return b.HostAmount
}
