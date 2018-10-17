package iso20022

// Deposited media put in the safe.
type ATMDepositedMedia1 struct {

	// Link to the account for multi-account deposit.
	AccountSequenceNumber *Number `xml:"AcctSeqNb,omitempty"`

	// Type of deposited media.
	MediaType *ATMMediaType2Code `xml:"MdiaTp"`

	// Category of deposited media items.
	MediaCategory *ATMMediaType3Code `xml:"MdiaCtgy,omitempty"`

	// Media item that are deposited.
	MediaItems []*ATMDepositedMedia2 `xml:"MdiaItms"`
}

func (a *ATMDepositedMedia1) SetAccountSequenceNumber(value string) {
	a.AccountSequenceNumber = (*Number)(&value)
}

func (a *ATMDepositedMedia1) SetMediaType(value string) {
	a.MediaType = (*ATMMediaType2Code)(&value)
}

func (a *ATMDepositedMedia1) SetMediaCategory(value string) {
	a.MediaCategory = (*ATMMediaType3Code)(&value)
}

func (a *ATMDepositedMedia1) AddMediaItems() *ATMDepositedMedia2 {
	newValue := new(ATMDepositedMedia2)
	a.MediaItems = append(a.MediaItems, newValue)
	return newValue
}

// Media item that are deposited.
type ATMDepositedMedia2 struct {

	// Number of deposit media.
	Count *Number `xml:"Cnt,omitempty"`

	// Amount or denomination of one media item, if the media type is valued or entered by the customer.
	UnitValue *ImpliedCurrencyAndAmount `xml:"UnitVal,omitempty"`

	// Currency of media items, if valued and different from base currency.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Format of the check code line.
	CodeLineFormat *CheckCodeLine1Code `xml:"CdLineFrmt,omitempty"`

	// Check code line.
	CodeLine *Max70Text `xml:"CdLine,omitempty"`

	// Check amount scanned by the check reader.
	ScannedValue *ImpliedCurrencyAndAmount `xml:"ScnndVal,omitempty"`

	// Percentage of the confidence in the check amount.
	ConfidenceLevel *Percentage `xml:"CnfdncLvl,omitempty"`
}

func (a *ATMDepositedMedia2) SetCount(value string) {
	a.Count = (*Number)(&value)
}

func (a *ATMDepositedMedia2) SetUnitValue(value, currency string) {
	a.UnitValue = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMDepositedMedia2) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMDepositedMedia2) SetCodeLineFormat(value string) {
	a.CodeLineFormat = (*CheckCodeLine1Code)(&value)
}

func (a *ATMDepositedMedia2) SetCodeLine(value string) {
	a.CodeLine = (*Max70Text)(&value)
}

func (a *ATMDepositedMedia2) SetScannedValue(value, currency string) {
	a.ScannedValue = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMDepositedMedia2) SetConfidenceLevel(value string) {
	a.ConfidenceLevel = (*Percentage)(&value)
}

// Deposited media put in the safe.
type ATMDepositedMedia3 struct {

	// Type of deposited media.
	MediaType *ATMMediaType2Code `xml:"MdiaTp"`

	// Category of deposited media items.
	MediaCategory *ATMMediaType3Code `xml:"MdiaCtgy,omitempty"`

	// Media item that are deposited.
	MediaItems []*ATMDepositedMedia2 `xml:"MdiaItms"`
}

func (a *ATMDepositedMedia3) SetMediaType(value string) {
	a.MediaType = (*ATMMediaType2Code)(&value)
}

func (a *ATMDepositedMedia3) SetMediaCategory(value string) {
	a.MediaCategory = (*ATMMediaType3Code)(&value)
}

func (a *ATMDepositedMedia3) AddMediaItems() *ATMDepositedMedia2 {
	newValue := new(ATMDepositedMedia2)
	a.MediaItems = append(a.MediaItems, newValue)
	return newValue
}
