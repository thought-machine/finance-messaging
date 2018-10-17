package iso20022

// Provides the details for the tax calculation method A.
type BillingMethod1 struct {

	// Amount of the original charge expressed in the host currency.
	ServiceChargeHostAmount *AmountAndDirection34 `xml:"SvcChrgHstAmt"`

	// Provides for the regional taxes on the service. Up to three regional taxes may be defined for the same service.
	ServiceTax *BillingServicesAmount1 `xml:"SvcTax"`

	// Specifies the total charge for a service (including taxes).
	TotalCharge *BillingServicesAmount2 `xml:"TtlChrg"`

	// Provides for the specific tax identification within the same tax region.
	//
	// Usage: This element allows for a maximum of three regional taxes on the same service.
	TaxIdentification []*BillingServicesTax1 `xml:"TaxId"`
}

func (b *BillingMethod1) AddServiceChargeHostAmount() *AmountAndDirection34 {
	b.ServiceChargeHostAmount = new(AmountAndDirection34)
	return b.ServiceChargeHostAmount
}

func (b *BillingMethod1) AddServiceTax() *BillingServicesAmount1 {
	b.ServiceTax = new(BillingServicesAmount1)
	return b.ServiceTax
}

func (b *BillingMethod1) AddTotalCharge() *BillingServicesAmount2 {
	b.TotalCharge = new(BillingServicesAmount2)
	return b.TotalCharge
}

func (b *BillingMethod1) AddTaxIdentification() *BillingServicesTax1 {
	newValue := new(BillingServicesTax1)
	b.TaxIdentification = append(b.TaxIdentification, newValue)
	return newValue
}

// Provides the details for the tax calculation method B.
type BillingMethod2 struct {

	// Amount of the original charge expressed in the host currency.
	ServiceChargeHostAmount *AmountAndDirection34 `xml:"SvcChrgHstAmt"`

	// Provides for the regional taxes on the service. Up to three regional taxes may be defined for the same service.
	ServiceTax *BillingServicesAmount1 `xml:"SvcTax"`

	// Provides for the specific tax identification within the same tax region.
	//
	// Usage: This element allows for a maximum of three regional taxes on the same service.
	TaxIdentification []*BillingServicesTax1 `xml:"TaxId"`
}

func (b *BillingMethod2) AddServiceChargeHostAmount() *AmountAndDirection34 {
	b.ServiceChargeHostAmount = new(AmountAndDirection34)
	return b.ServiceChargeHostAmount
}

func (b *BillingMethod2) AddServiceTax() *BillingServicesAmount1 {
	b.ServiceTax = new(BillingServicesAmount1)
	return b.ServiceTax
}

func (b *BillingMethod2) AddTaxIdentification() *BillingServicesTax1 {
	newValue := new(BillingServicesTax1)
	b.TaxIdentification = append(b.TaxIdentification, newValue)
	return newValue
}

// Provides the details for the tax calculation method D.
type BillingMethod3 struct {

	// Equivalent amount to the service tax host amount but allows the sender to optionally express the value in the pricing currency.
	ServiceTaxPriceAmount *AmountAndDirection34 `xml:"SvcTaxPricAmt"`

	// Provides for the specific tax identification within the same tax region.
	//
	// Usage: This element allows for a maximum of three regional taxes on the same service.
	TaxIdentification []*BillingServicesTax2 `xml:"TaxId"`
}

func (b *BillingMethod3) AddServiceTaxPriceAmount() *AmountAndDirection34 {
	b.ServiceTaxPriceAmount = new(AmountAndDirection34)
	return b.ServiceTaxPriceAmount
}

func (b *BillingMethod3) AddTaxIdentification() *BillingServicesTax2 {
	newValue := new(BillingServicesTax2)
	b.TaxIdentification = append(b.TaxIdentification, newValue)
	return newValue
}

// Provides the details for the tax calculation method C.
type BillingMethod4 struct {

	// Specifies the details of the taxable services using tax calculation method C.
	ServiceDetail []*BillingServiceParameters2 `xml:"SvcDtl"`

	// Total amount of service charge to be taxed in the tax region’s host currency along with the supporting tax calculations.
	//
	// Usage: Used for tax calculation method C only, and only one per tax region may be specified.
	TaxCalculation *TaxCalculation1 `xml:"TaxClctn"`
}

func (b *BillingMethod4) AddServiceDetail() *BillingServiceParameters2 {
	newValue := new(BillingServiceParameters2)
	b.ServiceDetail = append(b.ServiceDetail, newValue)
	return newValue
}

func (b *BillingMethod4) AddTaxCalculation() *TaxCalculation1 {
	b.TaxCalculation = new(TaxCalculation1)
	return b.TaxCalculation
}
