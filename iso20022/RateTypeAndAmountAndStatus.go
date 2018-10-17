package iso20022

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus1 struct {

	// Value expressed as a rate type.
	RateType *RateType13Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus1) AddRateType() *RateType13Choice {
	r.RateType = new(RateType13Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus1) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus1) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus11 struct {

	// Value expressed as a rate type.
	RateType *RateType17Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus11) AddRateType() *RateType17Choice {
	r.RateType = new(RateType17Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus11) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus11) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus13 struct {

	// Value expressed as a rate type.
	RateType *RateType20Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus13) AddRateType() *RateType20Choice {
	r.RateType = new(RateType20Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus13) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus13) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus14 struct {

	// Value expressed as a rate type.
	RateType *RateType21Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus14) AddRateType() *RateType21Choice {
	r.RateType = new(RateType21Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus14) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus14) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus15 struct {

	// Value expressed as a rate type.
	RateType *RateType22Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus15) AddRateType() *RateType22Choice {
	r.RateType = new(RateType22Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus15) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus15) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus16 struct {

	// Value expressed as a rate type.
	RateType *RateType23Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus16) AddRateType() *RateType23Choice {
	r.RateType = new(RateType23Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus16) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus16) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus22 struct {

	// Value expressed as a rate type.
	RateType *RateType40Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus3Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus22) AddRateType() *RateType40Choice {
	r.RateType = new(RateType40Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus22) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus22) AddRateStatus() *RateStatus3Choice {
	r.RateStatus = new(RateStatus3Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus24 struct {

	// Value expressed as a rate type.
	RateType *RateType33Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus3Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus24) AddRateType() *RateType33Choice {
	r.RateType = new(RateType33Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus24) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus24) AddRateStatus() *RateStatus3Choice {
	r.RateStatus = new(RateStatus3Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus25 struct {

	// Value expressed as a rate type.
	RateType *RateType41Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus3Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus25) AddRateType() *RateType41Choice {
	r.RateType = new(RateType41Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus25) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus25) AddRateStatus() *RateStatus3Choice {
	r.RateStatus = new(RateStatus3Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus26 struct {

	// Value expressed as a rate type.
	RateType *RateType36Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus3Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus26) AddRateType() *RateType36Choice {
	r.RateType = new(RateType36Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus26) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus26) AddRateStatus() *RateStatus3Choice {
	r.RateStatus = new(RateStatus3Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus27 struct {

	// Value expressed as a rate type.
	RateType *RateType37Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus3Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus27) AddRateType() *RateType37Choice {
	r.RateType = new(RateType37Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus27) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus27) AddRateStatus() *RateStatus3Choice {
	r.RateStatus = new(RateStatus3Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus28 struct {

	// Value expressed as a rate type.
	RateType *RateType38Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus3Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus28) AddRateType() *RateType38Choice {
	r.RateType = new(RateType38Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus28) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus28) AddRateStatus() *RateStatus3Choice {
	r.RateStatus = new(RateStatus3Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus29 struct {

	// Value expressed as a rate type.
	RateType *RateType39Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus3Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus29) AddRateType() *RateType39Choice {
	r.RateType = new(RateType39Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus29) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus29) AddRateStatus() *RateStatus3Choice {
	r.RateStatus = new(RateStatus3Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus3 struct {

	// Value expressed as a rate type.
	RateType *RateType6Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus3) AddRateType() *RateType6Choice {
	r.RateType = new(RateType6Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus3) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus3) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus30 struct {

	// Value expressed as a rate type.
	RateType *RateType51Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus4Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus30) AddRateType() *RateType51Choice {
	r.RateType = new(RateType51Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus30) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus30) AddRateStatus() *RateStatus4Choice {
	r.RateStatus = new(RateStatus4Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus31 struct {

	// Value expressed as a rate type.
	RateType *RateType44Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus4Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus31) AddRateType() *RateType44Choice {
	r.RateType = new(RateType44Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus31) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus31) AddRateStatus() *RateStatus4Choice {
	r.RateStatus = new(RateStatus4Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus32 struct {

	// Value expressed as a rate type.
	RateType *RateType45Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus4Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus32) AddRateType() *RateType45Choice {
	r.RateType = new(RateType45Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus32) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus32) AddRateStatus() *RateStatus4Choice {
	r.RateStatus = new(RateStatus4Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus33 struct {

	// Value expressed as a rate type.
	RateType *RateType47Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus4Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus33) AddRateType() *RateType47Choice {
	r.RateType = new(RateType47Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus33) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus33) AddRateStatus() *RateStatus4Choice {
	r.RateStatus = new(RateStatus4Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus34 struct {

	// Value expressed as a rate type.
	RateType *RateType48Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus4Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus34) AddRateType() *RateType48Choice {
	r.RateType = new(RateType48Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus34) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus34) AddRateStatus() *RateStatus4Choice {
	r.RateStatus = new(RateStatus4Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus35 struct {

	// Value expressed as a rate type.
	RateType *RateType49Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus4Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus35) AddRateType() *RateType49Choice {
	r.RateType = new(RateType49Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus35) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus35) AddRateStatus() *RateStatus4Choice {
	r.RateStatus = new(RateStatus4Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus36 struct {

	// Value expressed as a rate type.
	RateType *RateType50Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus4Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus36) AddRateType() *RateType50Choice {
	r.RateType = new(RateType50Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus36) SetAmount(value, currency string) {
	r.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus36) AddRateStatus() *RateStatus4Choice {
	r.RateStatus = new(RateStatus4Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus4 struct {

	// Value expressed as a rate type.
	RateType *RateType7Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus4) AddRateType() *RateType7Choice {
	r.RateType = new(RateType7Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus4) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus4) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus5 struct {

	// Value expressed as a rate type.
	RateType *RateType10Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus5) AddRateType() *RateType10Choice {
	r.RateType = new(RateType10Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus5) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus5) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus6 struct {

	// Value expressed as a rate type.
	RateType *RateType11Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Choice `xml:"RateSts,omitempty"`
}

func (r *RateTypeAndAmountAndStatus6) AddRateType() *RateType11Choice {
	r.RateType = new(RateType11Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus6) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus6) AddRateStatus() *RateStatus1Choice {
	r.RateStatus = new(RateStatus1Choice)
	return r.RateStatus
}
