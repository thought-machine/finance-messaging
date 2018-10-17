package iso20022

// Specifies elements related to the confirmation sent by the central counterparty to the clearing member in the context of the buy in process.
type BuyIn2 struct {

	// Indicates the reference of the BuyInNotification message.
	BuyInNotificationIdentification *Max35Text `xml:"BuyInNtfctnId,omitempty"`

	// Indicates the reference id of the buy in.
	BuyInIdentification *Max35Text `xml:"BuyInId"`

	// Provides the date at which the buy occured.
	Date *ISODate `xml:"Dt"`

	// Provides the price of the buy-in.
	Price *Price4 `xml:"Pric,omitempty"`

	// Specifies the elements related to the securities that the central counterparty had to buy in the context of the buy-in process.
	SecuritiesBuyIn *SecuritiesCompensation1 `xml:"SctiesBuyIn,omitempty"`

	// Provides details about the cash compensation required, in case the central counterparty could not buy the securities to cover the trade(s) that failed.
	RequiredCashCompensation *CashCompensation1 `xml:"ReqrdCshCompstn,omitempty"`
}

func (b *BuyIn2) SetBuyInNotificationIdentification(value string) {
	b.BuyInNotificationIdentification = (*Max35Text)(&value)
}

func (b *BuyIn2) SetBuyInIdentification(value string) {
	b.BuyInIdentification = (*Max35Text)(&value)
}

func (b *BuyIn2) SetDate(value string) {
	b.Date = (*ISODate)(&value)
}

func (b *BuyIn2) AddPrice() *Price4 {
	b.Price = new(Price4)
	return b.Price
}

func (b *BuyIn2) AddSecuritiesBuyIn() *SecuritiesCompensation1 {
	b.SecuritiesBuyIn = new(SecuritiesCompensation1)
	return b.SecuritiesBuyIn
}

func (b *BuyIn2) AddRequiredCashCompensation() *CashCompensation1 {
	b.RequiredCashCompensation = new(CashCompensation1)
	return b.RequiredCashCompensation
}

// Specifies elements related to the response sent by the clearing member to the central counterparty in the context of the buy in process.
type BuyIn3 struct {

	// Indicates the reference of the BuyInNotification message.
	BuyInNotificationIdentification *Max35Text `xml:"BuyInNtfctnId"`

	// Specific continuous net settlement case where the central counterparty can call for buy-in at a date anterior to "theoretical" buy-in date, the clearing member may request a delay.
	RequestForDelayIndicator *YesNoIndicator `xml:"ReqForDelyInd"`

	// Number of days associated to the request for delay.
	NumberOfDays *Number `xml:"NbOfDays"`

	// Buy in quantity called initially by the central counterparty.
	InitialQuantity *FinancialInstrumentQuantity1Choice `xml:"InitlQty"`

	// Quantity amount covered by the clearing member after notification.
	CoveredQuantity *FinancialInstrumentQuantity1Choice `xml:"CvrdQty"`

	// Quantity amount non covered by the clearing member after notification (this is, new buy in amount to be executed).
	UncoveredQuantity *FinancialInstrumentQuantity1Choice `xml:"UcvrdQty"`
}

func (b *BuyIn3) SetBuyInNotificationIdentification(value string) {
	b.BuyInNotificationIdentification = (*Max35Text)(&value)
}

func (b *BuyIn3) SetRequestForDelayIndicator(value string) {
	b.RequestForDelayIndicator = (*YesNoIndicator)(&value)
}

func (b *BuyIn3) SetNumberOfDays(value string) {
	b.NumberOfDays = (*Number)(&value)
}

func (b *BuyIn3) AddInitialQuantity() *FinancialInstrumentQuantity1Choice {
	b.InitialQuantity = new(FinancialInstrumentQuantity1Choice)
	return b.InitialQuantity
}

func (b *BuyIn3) AddCoveredQuantity() *FinancialInstrumentQuantity1Choice {
	b.CoveredQuantity = new(FinancialInstrumentQuantity1Choice)
	return b.CoveredQuantity
}

func (b *BuyIn3) AddUncoveredQuantity() *FinancialInstrumentQuantity1Choice {
	b.UncoveredQuantity = new(FinancialInstrumentQuantity1Choice)
	return b.UncoveredQuantity
}

// Specifies elements related to the notification (or warn) sent by the central counterparty to the clearing member in the context of the buy in process.
type BuyIn4 struct {

	// Indicates whether the message is a warning only or a notification.
	WarningIndicator *YesNoIndicator `xml:"WrngInd,omitempty"`

	// Provides the date at which the buy-in will occur.
	ExpectedBuyInDate *DateFormat15Choice `xml:"XpctdBuyInDt"`

	// Identifies the latest date by which the buy-in operation can be cancelled.
	CancellationLimitDate *ISODate `xml:"CxlLmtDt,omitempty"`

	// Identifies the date by which the buy-in operation is reversed by the CCP.
	BuyInReversionDate *ISODate `xml:"BuyInRvrsnDt,omitempty"`
}

func (b *BuyIn4) SetWarningIndicator(value string) {
	b.WarningIndicator = (*YesNoIndicator)(&value)
}

func (b *BuyIn4) AddExpectedBuyInDate() *DateFormat15Choice {
	b.ExpectedBuyInDate = new(DateFormat15Choice)
	return b.ExpectedBuyInDate
}

func (b *BuyIn4) SetCancellationLimitDate(value string) {
	b.CancellationLimitDate = (*ISODate)(&value)
}

func (b *BuyIn4) SetBuyInReversionDate(value string) {
	b.BuyInReversionDate = (*ISODate)(&value)
}
