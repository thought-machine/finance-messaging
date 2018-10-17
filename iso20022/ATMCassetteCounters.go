package iso20022

// ATM cassette counter per unit value or globally.
type ATMCassetteCounters1 struct {

	// Amount of one media unit, if the media type is valued.
	UnitValue *ImpliedCurrencyAndAmount `xml:"UnitVal,omitempty"`

	// Currency of the media, if the media type is valued and different from the currency of the requested amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Type of notes.
	ItemType *ATMNoteType2Code `xml:"ItmTp,omitempty"`

	// Counters of media inside the cassette.
	Counter []*ATMCassetteCounters2 `xml:"Cntr,omitempty"`

	// Current number of media present in the cassette.
	CurrentNumber *Number `xml:"CurNb"`

	// Current amount in the cassette.
	CurrentAmount *ImpliedCurrencyAndAmount `xml:"CurAmt,omitempty"`
}

func (a *ATMCassetteCounters1) SetUnitValue(value, currency string) {
	a.UnitValue = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMCassetteCounters1) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMCassetteCounters1) SetItemType(value string) {
	a.ItemType = (*ATMNoteType2Code)(&value)
}

func (a *ATMCassetteCounters1) AddCounter() *ATMCassetteCounters2 {
	newValue := new(ATMCassetteCounters2)
	a.Counter = append(a.Counter, newValue)
	return newValue
}

func (a *ATMCassetteCounters1) SetCurrentNumber(value string) {
	a.CurrentNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters1) SetCurrentAmount(value, currency string) {
	a.CurrentAmount = NewImpliedCurrencyAndAmount(value, currency)
}

// Counters of media inside an ATM cassette.
type ATMCassetteCounters2 struct {

	// Type of counters.
	Type *ATMCounterType1Code `xml:"Tp"`

	// Number of added media during servicing operations.
	AddedNumber *Number `xml:"AddedNb,omitempty"`

	// Number of removed media during servicing operations.
	RemovedNumber *Number `xml:"RmvdNb,omitempty"`

	// Total number of media out of the cassette.
	DispensedNumber *Number `xml:"DspnsdNb,omitempty"`

	// Total number of media deposited in the cassette.
	DepositNumber *Number `xml:"DpstNb,omitempty"`

	// Total number of recycled media from the cassette.
	RecycledNumber *Number `xml:"RcycldNb,omitempty"`

	// Total number of retracted media originating from the cassette.
	RetractedNumber *Number `xml:"RtrctdNb,omitempty"`

	// Total number of media from this cassette which are on the reject bin.
	RejectedNumber *Number `xml:"RjctdNb,omitempty"`

	// Total number of media presented to the customer.
	PresentedNumber *Number `xml:"PresntdNb,omitempty"`
}

func (a *ATMCassetteCounters2) SetType(value string) {
	a.Type = (*ATMCounterType1Code)(&value)
}

func (a *ATMCassetteCounters2) SetAddedNumber(value string) {
	a.AddedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters2) SetRemovedNumber(value string) {
	a.RemovedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters2) SetDispensedNumber(value string) {
	a.DispensedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters2) SetDepositNumber(value string) {
	a.DepositNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters2) SetRecycledNumber(value string) {
	a.RecycledNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters2) SetRetractedNumber(value string) {
	a.RetractedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters2) SetRejectedNumber(value string) {
	a.RejectedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters2) SetPresentedNumber(value string) {
	a.PresentedNumber = (*Number)(&value)
}

// ATM cassette counter per unit value or globally.
type ATMCassetteCounters3 struct {

	// Amount of one media unit, if the media type is valued.
	UnitValue *ImpliedCurrencyAndAmount `xml:"UnitVal,omitempty"`

	// Currency of the media, if the media type is valued and different from the currency of the requested amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Category of media items.
	MediaCategory *ATMMediaType3Code `xml:"MdiaCtgy,omitempty"`

	// Current number of media present in the cassette.
	CurrentNumber *Number `xml:"CurNb"`

	// Current amount in the cassette.
	CurrentAmount *ImpliedCurrencyAndAmount `xml:"CurAmt,omitempty"`

	// Counters of media inside the cassette.
	FlowTotals []*ATMCassetteCounters4 `xml:"FlowTtls,omitempty"`
}

func (a *ATMCassetteCounters3) SetUnitValue(value, currency string) {
	a.UnitValue = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMCassetteCounters3) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMCassetteCounters3) SetMediaCategory(value string) {
	a.MediaCategory = (*ATMMediaType3Code)(&value)
}

func (a *ATMCassetteCounters3) SetCurrentNumber(value string) {
	a.CurrentNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters3) SetCurrentAmount(value, currency string) {
	a.CurrentAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMCassetteCounters3) AddFlowTotals() *ATMCassetteCounters4 {
	newValue := new(ATMCassetteCounters4)
	a.FlowTotals = append(a.FlowTotals, newValue)
	return newValue
}

// Counters of media inside an ATM cassette.
type ATMCassetteCounters4 struct {

	// Type of counters.
	Type *ATMCounterType1Code `xml:"Tp"`

	// Number of added media during servicing operations.
	AddedNumber *Number `xml:"AddedNb,omitempty"`

	// Number of removed media during servicing operations.
	RemovedNumber *Number `xml:"RmvdNb,omitempty"`

	// Total number of media out of the cassette.
	DispensedNumber *Number `xml:"DspnsdNb,omitempty"`

	// Total number of media deposited in the cassette.
	DepositedNumber *Number `xml:"DpstdNb,omitempty"`

	// Total number of recycled media from the cassette.
	RecycledNumber *Number `xml:"RcycldNb,omitempty"`

	// Total number of retracted media originating from the cassette.
	RetractedNumber *Number `xml:"RtrctdNb,omitempty"`

	// Total number of media from this cassette which are on the reject bin.
	RejectedNumber *Number `xml:"RjctdNb,omitempty"`

	// Total number of media presented to the customer.
	PresentedNumber *Number `xml:"PresntdNb,omitempty"`
}

func (a *ATMCassetteCounters4) SetType(value string) {
	a.Type = (*ATMCounterType1Code)(&value)
}

func (a *ATMCassetteCounters4) SetAddedNumber(value string) {
	a.AddedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetRemovedNumber(value string) {
	a.RemovedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetDispensedNumber(value string) {
	a.DispensedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetDepositedNumber(value string) {
	a.DepositedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetRecycledNumber(value string) {
	a.RecycledNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetRetractedNumber(value string) {
	a.RetractedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetRejectedNumber(value string) {
	a.RejectedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetPresentedNumber(value string) {
	a.PresentedNumber = (*Number)(&value)
}
