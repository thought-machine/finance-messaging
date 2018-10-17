package iso20022

// Specification of the fund order type.
type FundOrderType1 struct {

	// Structured format.
	Structured *FundOrderType2Code `xml:"Strd"`

	// Additional information about the type of identification
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (f *FundOrderType1) SetStructured(value string) {
	f.Structured = (*FundOrderType2Code)(&value)
}

func (f *FundOrderType1) SetAdditionalInformation(value string) {
	f.AdditionalInformation = (*Max350Text)(&value)
}

// Specifies the category of the investment fund order.
type FundOrderType2 struct {

	// Specifies the category of the investment fund order.
	OrderType *FundOrderType3Code `xml:"OrdrTp"`

	// Specifies the category of the investment fund order.
	ExtendedOrderType *Extended350Code `xml:"XtndedOrdrTp"`
}

func (f *FundOrderType2) SetOrderType(value string) {
	f.OrderType = (*FundOrderType3Code)(&value)
}

func (f *FundOrderType2) SetExtendedOrderType(value string) {
	f.ExtendedOrderType = (*Extended350Code)(&value)
}

// Specifies the category of the investment fund order.
type FundOrderType3 struct {

	// Specifies the category of the investment fund order.
	OrderType *FundOrderType4Code `xml:"OrdrTp"`

	// Specifies the category of the investment fund order.
	ExtendedOrderType *Extended350Code `xml:"XtndedOrdrTp"`
}

func (f *FundOrderType3) SetOrderType(value string) {
	f.OrderType = (*FundOrderType4Code)(&value)
}

func (f *FundOrderType3) SetExtendedOrderType(value string) {
	f.ExtendedOrderType = (*Extended350Code)(&value)
}
