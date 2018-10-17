package iso20022

// Specifies information about blocked accounts.
type Blocked1 struct {

	// Specifies the order type for which the account is blocked.
	OrderType []*FundOrderType1Choice `xml:"OrdrTp"`

	// Indicates whether the account is blocked.
	Blocked *YesNoIndicator `xml:"Blckd"`

	// Specifies the reason the account is blocked.
	Reason *BlockedReason1Choice `xml:"Rsn,omitempty"`
}

func (b *Blocked1) AddOrderType() *FundOrderType1Choice {
	newValue := new(FundOrderType1Choice)
	b.OrderType = append(b.OrderType, newValue)
	return newValue
}

func (b *Blocked1) SetBlocked(value string) {
	b.Blocked = (*YesNoIndicator)(&value)
}

func (b *Blocked1) AddReason() *BlockedReason1Choice {
	b.Reason = new(BlockedReason1Choice)
	return b.Reason
}

// Information about a blocked account.
type Blocked2 struct {

	// Specifies the order or transaction type for which the account is blocked.
	OrderType []*OrderType2Choice `xml:"OrdrTp"`

	// Indicates whether the account is blocked.
	Blocked *YesNoIndicator `xml:"Blckd"`

	// Specifies the reason the account is blocked.
	Reason *BlockedReason1Choice `xml:"Rsn,omitempty"`
}

func (b *Blocked2) AddOrderType() *OrderType2Choice {
	newValue := new(OrderType2Choice)
	b.OrderType = append(b.OrderType, newValue)
	return newValue
}

func (b *Blocked2) SetBlocked(value string) {
	b.Blocked = (*YesNoIndicator)(&value)
}

func (b *Blocked2) AddReason() *BlockedReason1Choice {
	b.Reason = new(BlockedReason1Choice)
	return b.Reason
}
