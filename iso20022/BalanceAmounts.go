package iso20022

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type BalanceAmounts1 struct {

	// Value of an individual financial instrument holding within a safekeeping account.
	HoldingValue *AmountAndDirection6 `xml:"HldgVal"`

	// Previous value of an individual financial instrument holding within a safekeeping account.
	PreviousHoldingValue *AmountAndDirection6 `xml:"PrvsHldgVal,omitempty"`

	// Value of a financial instrument, as booked/acquired in an account. It may be used to establish capital gain tax liability.
	BookValue *AmountAndDirection6 `xml:"BookVal,omitempty"`

	// Difference between holding value and the book value.
	UnrealisedGainLoss *AmountAndDirection6 `xml:"UrlsdGnLoss,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection6 `xml:"AcrdIntrstAmt,omitempty"`
}

func (b *BalanceAmounts1) AddHoldingValue() *AmountAndDirection6 {
	b.HoldingValue = new(AmountAndDirection6)
	return b.HoldingValue
}

func (b *BalanceAmounts1) AddPreviousHoldingValue() *AmountAndDirection6 {
	b.PreviousHoldingValue = new(AmountAndDirection6)
	return b.PreviousHoldingValue
}

func (b *BalanceAmounts1) AddBookValue() *AmountAndDirection6 {
	b.BookValue = new(AmountAndDirection6)
	return b.BookValue
}

func (b *BalanceAmounts1) AddUnrealisedGainLoss() *AmountAndDirection6 {
	b.UnrealisedGainLoss = new(AmountAndDirection6)
	return b.UnrealisedGainLoss
}

func (b *BalanceAmounts1) AddAccruedInterestAmount() *AmountAndDirection6 {
	b.AccruedInterestAmount = new(AmountAndDirection6)
	return b.AccruedInterestAmount
}

// Amounts linked to a securities balance, for example, holding value.
type BalanceAmounts2 struct {

	// Value of an individual financial instrument holding within a safekeeping account.
	HoldingValue *AmountAndDirection6 `xml:"HldgVal"`

	// Value of a financial instrument, as booked/acquired in an account. It may be used to establish capital gain tax liability.
	BookValue *AmountAndDirection6 `xml:"BookVal,omitempty"`

	// Difference between holding value and the book value.
	UnrealisedGainLoss *AmountAndDirection6 `xml:"UrlsdGnLoss,omitempty"`
}

func (b *BalanceAmounts2) AddHoldingValue() *AmountAndDirection6 {
	b.HoldingValue = new(AmountAndDirection6)
	return b.HoldingValue
}

func (b *BalanceAmounts2) AddBookValue() *AmountAndDirection6 {
	b.BookValue = new(AmountAndDirection6)
	return b.BookValue
}

func (b *BalanceAmounts2) AddUnrealisedGainLoss() *AmountAndDirection6 {
	b.UnrealisedGainLoss = new(AmountAndDirection6)
	return b.UnrealisedGainLoss
}

// Amounts linked to a securities balance, for example, holding value.
type BalanceAmounts3 struct {

	// Value of an individual financial instrument holding within a safekeeping account.
	HoldingValue *AmountAndDirection6 `xml:"HldgVal,omitempty"`

	// Previous value of an individual financial instrument holding within a safekeeping account.
	PreviousHoldingValue *AmountAndDirection6 `xml:"PrvsHldgVal,omitempty"`

	// Value of a financial instrument, as booked/acquired in an account. It may be used to establish capital gain tax liability.
	BookValue *AmountAndDirection6 `xml:"BookVal,omitempty"`

	// Value of the position eligible for collateral purposes.
	EligibleCollateralValue *AmountAndDirection6 `xml:"ElgblCollVal,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection6 `xml:"AcrdIntrstAmt,omitempty"`
}

func (b *BalanceAmounts3) AddHoldingValue() *AmountAndDirection6 {
	b.HoldingValue = new(AmountAndDirection6)
	return b.HoldingValue
}

func (b *BalanceAmounts3) AddPreviousHoldingValue() *AmountAndDirection6 {
	b.PreviousHoldingValue = new(AmountAndDirection6)
	return b.PreviousHoldingValue
}

func (b *BalanceAmounts3) AddBookValue() *AmountAndDirection6 {
	b.BookValue = new(AmountAndDirection6)
	return b.BookValue
}

func (b *BalanceAmounts3) AddEligibleCollateralValue() *AmountAndDirection6 {
	b.EligibleCollateralValue = new(AmountAndDirection6)
	return b.EligibleCollateralValue
}

func (b *BalanceAmounts3) AddAccruedInterestAmount() *AmountAndDirection6 {
	b.AccruedInterestAmount = new(AmountAndDirection6)
	return b.AccruedInterestAmount
}

// Amounts linked to a securities balance, for example, holding value.
type BalanceAmounts4 struct {

	// Value of an individual financial instrument holding within a safekeeping account.
	HoldingValue *AmountAndDirection14 `xml:"HldgVal,omitempty"`

	// Previous value of an individual financial instrument holding within a safekeeping account.
	PreviousHoldingValue *AmountAndDirection14 `xml:"PrvsHldgVal,omitempty"`

	// Value of a financial instrument, as booked/acquired in an account. It may be used to establish capital gain tax liability.
	BookValue *AmountAndDirection14 `xml:"BookVal,omitempty"`

	// Value of the position eligible for collateral purposes.
	EligibleCollateralValue *AmountAndDirection14 `xml:"ElgblCollVal,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection14 `xml:"AcrdIntrstAmt,omitempty"`
}

func (b *BalanceAmounts4) AddHoldingValue() *AmountAndDirection14 {
	b.HoldingValue = new(AmountAndDirection14)
	return b.HoldingValue
}

func (b *BalanceAmounts4) AddPreviousHoldingValue() *AmountAndDirection14 {
	b.PreviousHoldingValue = new(AmountAndDirection14)
	return b.PreviousHoldingValue
}

func (b *BalanceAmounts4) AddBookValue() *AmountAndDirection14 {
	b.BookValue = new(AmountAndDirection14)
	return b.BookValue
}

func (b *BalanceAmounts4) AddEligibleCollateralValue() *AmountAndDirection14 {
	b.EligibleCollateralValue = new(AmountAndDirection14)
	return b.EligibleCollateralValue
}

func (b *BalanceAmounts4) AddAccruedInterestAmount() *AmountAndDirection14 {
	b.AccruedInterestAmount = new(AmountAndDirection14)
	return b.AccruedInterestAmount
}

// Posting of an item to a cash account, in the context of a cash transaction, that results in an increase or decrease to the balance of the account.
type BalanceAmounts5 struct {

	// Value of an individual financial instrument holding within a safekeeping account.
	HoldingValue *AmountAndDirection14 `xml:"HldgVal"`

	// Previous value of an individual financial instrument holding within a safekeeping account.
	PreviousHoldingValue *AmountAndDirection14 `xml:"PrvsHldgVal,omitempty"`

	// Value of a financial instrument, as booked/acquired in an account. It may be used to establish capital gain tax liability.
	BookValue *AmountAndDirection14 `xml:"BookVal,omitempty"`

	// Difference between holding value and the book value.
	UnrealisedGainLoss *AmountAndDirection14 `xml:"UrlsdGnLoss,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection14 `xml:"AcrdIntrstAmt,omitempty"`
}

func (b *BalanceAmounts5) AddHoldingValue() *AmountAndDirection14 {
	b.HoldingValue = new(AmountAndDirection14)
	return b.HoldingValue
}

func (b *BalanceAmounts5) AddPreviousHoldingValue() *AmountAndDirection14 {
	b.PreviousHoldingValue = new(AmountAndDirection14)
	return b.PreviousHoldingValue
}

func (b *BalanceAmounts5) AddBookValue() *AmountAndDirection14 {
	b.BookValue = new(AmountAndDirection14)
	return b.BookValue
}

func (b *BalanceAmounts5) AddUnrealisedGainLoss() *AmountAndDirection14 {
	b.UnrealisedGainLoss = new(AmountAndDirection14)
	return b.UnrealisedGainLoss
}

func (b *BalanceAmounts5) AddAccruedInterestAmount() *AmountAndDirection14 {
	b.AccruedInterestAmount = new(AmountAndDirection14)
	return b.AccruedInterestAmount
}

// Amounts linked to a securities balance, for example, holding value.
type BalanceAmounts6 struct {

	// Value of an individual financial instrument holding within a safekeeping account.
	HoldingValue *AmountAndDirection14 `xml:"HldgVal"`

	// Value of a financial instrument, as booked/acquired in an account. It may be used to establish capital gain tax liability.
	BookValue *AmountAndDirection14 `xml:"BookVal,omitempty"`

	// Difference between holding value and the book value.
	UnrealisedGainLoss *AmountAndDirection14 `xml:"UrlsdGnLoss,omitempty"`
}

func (b *BalanceAmounts6) AddHoldingValue() *AmountAndDirection14 {
	b.HoldingValue = new(AmountAndDirection14)
	return b.HoldingValue
}

func (b *BalanceAmounts6) AddBookValue() *AmountAndDirection14 {
	b.BookValue = new(AmountAndDirection14)
	return b.BookValue
}

func (b *BalanceAmounts6) AddUnrealisedGainLoss() *AmountAndDirection14 {
	b.UnrealisedGainLoss = new(AmountAndDirection14)
	return b.UnrealisedGainLoss
}
