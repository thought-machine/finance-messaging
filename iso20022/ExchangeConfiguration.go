package iso20022

// Configuration parameters of data exchanges.
type ExchangeConfiguration1 struct {

	// Exchange policy between parties.
	ExchangePolicy []*ExchangePolicy1Code `xml:"XchgPlcy"`

	// Maximum number of transactions without exchange.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// Maximum cumulative amount of the transactions without exchange.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Timing condition for periodic exchanges.
	TimeCondition *ProcessTiming1 `xml:"TmCond,omitempty"`
}

func (e *ExchangeConfiguration1) AddExchangePolicy(value string) {
	e.ExchangePolicy = append(e.ExchangePolicy, (*ExchangePolicy1Code)(&value))
}

func (e *ExchangeConfiguration1) SetMaximumNumber(value string) {
	e.MaximumNumber = (*Number)(&value)
}

func (e *ExchangeConfiguration1) SetMaximumAmount(value, currency string) {
	e.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (e *ExchangeConfiguration1) AddTimeCondition() *ProcessTiming1 {
	e.TimeCondition = new(ProcessTiming1)
	return e.TimeCondition
}

// Configuration parameters of data exchanges.
type ExchangeConfiguration2 struct {

	// Exchange policy between parties.
	ExchangePolicy []*ExchangePolicy1Code `xml:"XchgPlcy"`

	// Maximum number of transactions without exchange.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// Maximum cumulative amount of the transactions without exchange.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Timing condition for periodic exchanges.
	TimeCondition *ProcessTiming2 `xml:"TmCond,omitempty"`
}

func (e *ExchangeConfiguration2) AddExchangePolicy(value string) {
	e.ExchangePolicy = append(e.ExchangePolicy, (*ExchangePolicy1Code)(&value))
}

func (e *ExchangeConfiguration2) SetMaximumNumber(value string) {
	e.MaximumNumber = (*Number)(&value)
}

func (e *ExchangeConfiguration2) SetMaximumAmount(value, currency string) {
	e.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (e *ExchangeConfiguration2) AddTimeCondition() *ProcessTiming2 {
	e.TimeCondition = new(ProcessTiming2)
	return e.TimeCondition
}

// Configuration parameters of data exchanges.
type ExchangeConfiguration3 struct {

	// Exchange policy between parties.
	ExchangePolicy []*ExchangePolicy1Code `xml:"XchgPlcy"`

	// Maximum number of transactions without exchange.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// Maximum cumulative amount of the transactions without exchange.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Timing condition for periodic exchanges.
	TimeCondition *ProcessTiming2 `xml:"TmCond,omitempty"`

	// Failed transaction must be exchanged.
	//
	ExchangeFailed *TrueFalseIndicator `xml:"XchgFaild,omitempty"`

	// Indicates that declined transaction must be exchanged
	ExchangeDeclined *TrueFalseIndicator `xml:"XchgDclnd,omitempty"`
}

func (e *ExchangeConfiguration3) AddExchangePolicy(value string) {
	e.ExchangePolicy = append(e.ExchangePolicy, (*ExchangePolicy1Code)(&value))
}

func (e *ExchangeConfiguration3) SetMaximumNumber(value string) {
	e.MaximumNumber = (*Number)(&value)
}

func (e *ExchangeConfiguration3) SetMaximumAmount(value, currency string) {
	e.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (e *ExchangeConfiguration3) AddTimeCondition() *ProcessTiming2 {
	e.TimeCondition = new(ProcessTiming2)
	return e.TimeCondition
}

func (e *ExchangeConfiguration3) SetExchangeFailed(value string) {
	e.ExchangeFailed = (*TrueFalseIndicator)(&value)
}

func (e *ExchangeConfiguration3) SetExchangeDeclined(value string) {
	e.ExchangeDeclined = (*TrueFalseIndicator)(&value)
}

// Configuration parameters of data exchanges.
type ExchangeConfiguration4 struct {

	// Exchange policy between parties.
	ExchangePolicy []*ExchangePolicy1Code `xml:"XchgPlcy"`

	// Maximum number of transactions without exchange.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// Maximum cumulative amount of the transactions without exchange.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Definition of retry process if activation of an action fails.
	ReTry *ProcessRetry2 `xml:"ReTry,omitempty"`

	// Timing condition for periodic exchanges.
	TimeCondition *ProcessTiming3 `xml:"TmCond,omitempty"`
}

func (e *ExchangeConfiguration4) AddExchangePolicy(value string) {
	e.ExchangePolicy = append(e.ExchangePolicy, (*ExchangePolicy1Code)(&value))
}

func (e *ExchangeConfiguration4) SetMaximumNumber(value string) {
	e.MaximumNumber = (*Number)(&value)
}

func (e *ExchangeConfiguration4) SetMaximumAmount(value, currency string) {
	e.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (e *ExchangeConfiguration4) AddReTry() *ProcessRetry2 {
	e.ReTry = new(ProcessRetry2)
	return e.ReTry
}

func (e *ExchangeConfiguration4) AddTimeCondition() *ProcessTiming3 {
	e.TimeCondition = new(ProcessTiming3)
	return e.TimeCondition
}

// Configuration parameters of data exchanges.
type ExchangeConfiguration5 struct {

	// Exchange policy between parties.
	ExchangePolicy []*ExchangePolicy1Code `xml:"XchgPlcy"`

	// Maximum number of transactions without exchange.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// Maximum cumulative amount of the transactions without exchange.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Definition of retry process if activation of an action fails.
	ReTry *ProcessRetry2 `xml:"ReTry,omitempty"`

	// Timing condition for periodic exchanges.
	TimeCondition *ProcessTiming3 `xml:"TmCond,omitempty"`

	// Failed transaction must be exchanged.
	//
	ExchangeFailed *TrueFalseIndicator `xml:"XchgFaild,omitempty"`

	// Indicates that declined transaction must be exchanged.
	ExchangeDeclined *TrueFalseIndicator `xml:"XchgDclnd,omitempty"`
}

func (e *ExchangeConfiguration5) AddExchangePolicy(value string) {
	e.ExchangePolicy = append(e.ExchangePolicy, (*ExchangePolicy1Code)(&value))
}

func (e *ExchangeConfiguration5) SetMaximumNumber(value string) {
	e.MaximumNumber = (*Number)(&value)
}

func (e *ExchangeConfiguration5) SetMaximumAmount(value, currency string) {
	e.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (e *ExchangeConfiguration5) AddReTry() *ProcessRetry2 {
	e.ReTry = new(ProcessRetry2)
	return e.ReTry
}

func (e *ExchangeConfiguration5) AddTimeCondition() *ProcessTiming3 {
	e.TimeCondition = new(ProcessTiming3)
	return e.TimeCondition
}

func (e *ExchangeConfiguration5) SetExchangeFailed(value string) {
	e.ExchangeFailed = (*TrueFalseIndicator)(&value)
}

func (e *ExchangeConfiguration5) SetExchangeDeclined(value string) {
	e.ExchangeDeclined = (*TrueFalseIndicator)(&value)
}

// Configuration parameters of data exchanges.
type ExchangeConfiguration6 struct {

	// Exchange policy between parties.
	ExchangePolicy []*ExchangePolicy1Code `xml:"XchgPlcy"`

	// Maximum number of transactions without exchange.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// Maximum cumulative amount of the transactions without exchange.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Definition of retry process if activation of an action fails.
	ReTry *ProcessRetry2 `xml:"ReTry,omitempty"`

	// Timing condition for periodic exchanges.
	TimeCondition *ProcessTiming4 `xml:"TmCond,omitempty"`
}

func (e *ExchangeConfiguration6) AddExchangePolicy(value string) {
	e.ExchangePolicy = append(e.ExchangePolicy, (*ExchangePolicy1Code)(&value))
}

func (e *ExchangeConfiguration6) SetMaximumNumber(value string) {
	e.MaximumNumber = (*Number)(&value)
}

func (e *ExchangeConfiguration6) SetMaximumAmount(value, currency string) {
	e.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (e *ExchangeConfiguration6) AddReTry() *ProcessRetry2 {
	e.ReTry = new(ProcessRetry2)
	return e.ReTry
}

func (e *ExchangeConfiguration6) AddTimeCondition() *ProcessTiming4 {
	e.TimeCondition = new(ProcessTiming4)
	return e.TimeCondition
}

// Configuration parameters of data exchanges.
type ExchangeConfiguration7 struct {

	// Exchange policy between parties.
	ExchangePolicy []*ExchangePolicy1Code `xml:"XchgPlcy"`

	// Maximum number of transactions without exchange.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// Maximum cumulative amount of the transactions without exchange.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Definition of retry process if activation of an action fails.
	ReTry *ProcessRetry2 `xml:"ReTry,omitempty"`

	// Timing condition for periodic exchanges.
	TimeCondition *ProcessTiming4 `xml:"TmCond,omitempty"`

	// Failed transaction must be exchanged.
	//
	ExchangeFailed *TrueFalseIndicator `xml:"XchgFaild,omitempty"`

	// Indicates that declined transaction must be exchanged.
	ExchangeDeclined *TrueFalseIndicator `xml:"XchgDclnd,omitempty"`
}

func (e *ExchangeConfiguration7) AddExchangePolicy(value string) {
	e.ExchangePolicy = append(e.ExchangePolicy, (*ExchangePolicy1Code)(&value))
}

func (e *ExchangeConfiguration7) SetMaximumNumber(value string) {
	e.MaximumNumber = (*Number)(&value)
}

func (e *ExchangeConfiguration7) SetMaximumAmount(value, currency string) {
	e.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (e *ExchangeConfiguration7) AddReTry() *ProcessRetry2 {
	e.ReTry = new(ProcessRetry2)
	return e.ReTry
}

func (e *ExchangeConfiguration7) AddTimeCondition() *ProcessTiming4 {
	e.TimeCondition = new(ProcessTiming4)
	return e.TimeCondition
}

func (e *ExchangeConfiguration7) SetExchangeFailed(value string) {
	e.ExchangeFailed = (*TrueFalseIndicator)(&value)
}

func (e *ExchangeConfiguration7) SetExchangeDeclined(value string) {
	e.ExchangeDeclined = (*TrueFalseIndicator)(&value)
}
