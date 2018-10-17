package iso20022

// Identify the original notification item, to which the cancellation advice refers.
type OriginalItem2 struct {

	// Identification of the original notification item.
	OriginalItemIdentification *Max35Text `xml:"OrgnlItmId"`

	// Unique identification as assigned by the debtor to unambiguously identify the original underlying transaction to the creditor.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money expected to be credited to the account, as per the original notification to receive.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Value date on which the account was expected to be credited.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Provides further information in order to identify a previous payment notification.
	OriginalItemReference *OriginalItemReference1 `xml:"OrgnlItmRef,omitempty"`
}

func (o *OriginalItem2) SetOriginalItemIdentification(value string) {
	o.OriginalItemIdentification = (*Max35Text)(&value)
}

func (o *OriginalItem2) SetOriginalEndToEndIdentification(value string) {
	o.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (o *OriginalItem2) SetAmount(value, currency string) {
	o.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalItem2) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalItem2) AddOriginalItemReference() *OriginalItemReference1 {
	o.OriginalItemReference = new(OriginalItemReference1)
	return o.OriginalItemReference
}

// Identifies the original notification item, to which the cancellation advice refers.
type OriginalItem3 struct {

	// Identification of the original notification item.
	OriginalItemIdentification *Max35Text `xml:"OrgnlItmId"`

	// Unique identification as assigned by the debtor to unambiguously identify the original underlying transaction to the creditor.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money expected to be credited to the account, as per the original notification to receive.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Value date on which the account was expected to be credited.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Provides further information in order to identify a previous payment notification.
	OriginalItemReference *OriginalItemReference2 `xml:"OrgnlItmRef,omitempty"`
}

func (o *OriginalItem3) SetOriginalItemIdentification(value string) {
	o.OriginalItemIdentification = (*Max35Text)(&value)
}

func (o *OriginalItem3) SetOriginalEndToEndIdentification(value string) {
	o.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (o *OriginalItem3) SetAmount(value, currency string) {
	o.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalItem3) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalItem3) AddOriginalItemReference() *OriginalItemReference2 {
	o.OriginalItemReference = new(OriginalItemReference2)
	return o.OriginalItemReference
}

// Identifies the original notification item, to which the cancellation advice refers.
type OriginalItem4 struct {

	// Identification of the original notification item.
	OriginalItemIdentification *Max35Text `xml:"OrgnlItmId"`

	// Unique identification as assigned by the debtor to unambiguously identify the original underlying transaction to the creditor.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money expected to be credited to the account, as per the original notification to receive.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Value date on which the account was expected to be credited.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Provides further information in order to identify a previous payment notification.
	OriginalItemReference *OriginalItemReference3 `xml:"OrgnlItmRef,omitempty"`
}

func (o *OriginalItem4) SetOriginalItemIdentification(value string) {
	o.OriginalItemIdentification = (*Max35Text)(&value)
}

func (o *OriginalItem4) SetOriginalEndToEndIdentification(value string) {
	o.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (o *OriginalItem4) SetAmount(value, currency string) {
	o.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalItem4) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalItem4) AddOriginalItemReference() *OriginalItemReference3 {
	o.OriginalItemReference = new(OriginalItemReference3)
	return o.OriginalItemReference
}

// Identifies the original notification item, to which the cancellation advice refers.
type OriginalItem5 struct {

	// Identification of the original notification item.
	OriginalItemIdentification *Max35Text `xml:"OrgnlItmId"`

	// Unique identification as assigned by the debtor to unambiguously identify the original underlying transaction to the creditor.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money expected to be credited to the account, as per the original notification to receive.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Value date on which the account was expected to be credited.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Provides further information in order to identify a previous payment notification.
	OriginalItemReference *OriginalItemReference4 `xml:"OrgnlItmRef,omitempty"`
}

func (o *OriginalItem5) SetOriginalItemIdentification(value string) {
	o.OriginalItemIdentification = (*Max35Text)(&value)
}

func (o *OriginalItem5) SetOriginalEndToEndIdentification(value string) {
	o.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (o *OriginalItem5) SetAmount(value, currency string) {
	o.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalItem5) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalItem5) AddOriginalItemReference() *OriginalItemReference4 {
	o.OriginalItemReference = new(OriginalItemReference4)
	return o.OriginalItemReference
}
