package iso20022

// Specifies the detailed parameters a service to be billed.
type BillingServiceParameters1 struct {

	// Specifies the details to fully identify the bank service.
	BankService *BillingServiceIdentification1 `xml:"BkSvc"`

	// Count or number of items (volume) involved in the charge.
	Volume *DecimalNumber `xml:"Vol,omitempty"`
}

func (b *BillingServiceParameters1) AddBankService() *BillingServiceIdentification1 {
	b.BankService = new(BillingServiceIdentification1)
	return b.BankService
}

func (b *BillingServiceParameters1) SetVolume(value string) {
	b.Volume = (*DecimalNumber)(&value)
}

// Specifies the detailed parameters a service to be billed.
type BillingServiceParameters2 struct {

	// Specifies the details to fully identify the bank service.
	BankService *BillingServiceIdentification2 `xml:"BkSvc"`

	// Count or number of items (volume) involved in the charge.
	Volume *DecimalNumber `xml:"Vol,omitempty"`

	// Price per item or unit used to calculate the charge expressed in the pricing currency.
	UnitPrice *AmountAndDirection34 `xml:"UnitPric,omitempty"`

	// Amount of the calculated charge expressed in the pricing currency, exclusive of any tax.
	ServiceChargeAmount *AmountAndDirection34 `xml:"SvcChrgAmt"`
}

func (b *BillingServiceParameters2) AddBankService() *BillingServiceIdentification2 {
	b.BankService = new(BillingServiceIdentification2)
	return b.BankService
}

func (b *BillingServiceParameters2) SetVolume(value string) {
	b.Volume = (*DecimalNumber)(&value)
}

func (b *BillingServiceParameters2) AddUnitPrice() *AmountAndDirection34 {
	b.UnitPrice = new(AmountAndDirection34)
	return b.UnitPrice
}

func (b *BillingServiceParameters2) AddServiceChargeAmount() *AmountAndDirection34 {
	b.ServiceChargeAmount = new(AmountAndDirection34)
	return b.ServiceChargeAmount
}

// Specifies the detailed parameters a service to be billed.
type BillingServiceParameters3 struct {

	// Specifies the details to fully identify the bank service.
	BankService *BillingServiceIdentification3 `xml:"BkSvc"`

	// Count or number of items (volume) involved in the charge.
	Volume *DecimalNumber `xml:"Vol,omitempty"`
}

func (b *BillingServiceParameters3) AddBankService() *BillingServiceIdentification3 {
	b.BankService = new(BillingServiceIdentification3)
	return b.BankService
}

func (b *BillingServiceParameters3) SetVolume(value string) {
	b.Volume = (*DecimalNumber)(&value)
}
