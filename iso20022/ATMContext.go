package iso20022

// Context in which the transaction is performed.
type ATMContext1 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService1 `xml:"Svc,omitempty"`
}

func (a *ATMContext1) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext1) AddService() *ATMService1 {
	a.Service = new(ATMService1)
	return a.Service
}

// Context in which the transaction is performed.
type ATMContext10 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Deposit service provided by the ATM inside the session.
	Service *ATMService11 `xml:"Svc,omitempty"`
}

func (a *ATMContext10) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext10) AddService() *ATMService11 {
	a.Service = new(ATMService11)
	return a.Service
}

// Context in which the transaction is performed.
type ATMContext11 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Deposit service provided by the ATM inside the session.
	Service *ATMService12 `xml:"Svc,omitempty"`
}

func (a *ATMContext11) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext11) AddService() *ATMService12 {
	a.Service = new(ATMService12)
	return a.Service
}

// Context in which the transaction is performed.
type ATMContext12 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Deposit service provided by the ATM inside the session.
	Service *ATMService13 `xml:"Svc,omitempty"`
}

func (a *ATMContext12) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext12) AddService() *ATMService13 {
	a.Service = new(ATMService13)
	return a.Service
}

// Context in which the transaction is performed.
type ATMContext13 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Service provided by the ATM inside the session.
	Service *ATMService14 `xml:"Svc"`
}

func (a *ATMContext13) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext13) AddService() *ATMService14 {
	a.Service = new(ATMService14)
	return a.Service
}

// Context in which the inquiry is performed.
type ATMContext14 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService15 `xml:"Svc"`
}

func (a *ATMContext14) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext14) AddService() *ATMService15 {
	a.Service = new(ATMService15)
	return a.Service
}

// Context in which the inquiry is performed.
type ATMContext15 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService16 `xml:"Svc"`
}

func (a *ATMContext15) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext15) AddService() *ATMService16 {
	a.Service = new(ATMService16)
	return a.Service
}

// Context in which the transaction is performed.
type ATMContext16 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Context in which the transaction is performed.
	Service *ATMService20 `xml:"Svc"`
}

func (a *ATMContext16) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext16) AddService() *ATMService20 {
	a.Service = new(ATMService20)
	return a.Service
}

// Context in which the transaction is performed.
type ATMContext17 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService21 `xml:"Svc"`
}

func (a *ATMContext17) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext17) AddService() *ATMService21 {
	a.Service = new(ATMService21)
	return a.Service
}

// Context in which the transfer is performed.
type ATMContext18 struct {

	// Unique identification of the customer session in which the transfer is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Fund transfer service requested by the ATM inside the session.
	Service *ATMService22 `xml:"Svc"`
}

func (a *ATMContext18) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext18) AddService() *ATMService22 {
	a.Service = new(ATMService22)
	return a.Service
}

// Context in which the transfer is performed.
type ATMContext19 struct {

	// Unique identification of the customer session in which the transfer is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Fund transfer service requested by the ATM inside the session.
	Service *ATMService23 `xml:"Svc"`
}

func (a *ATMContext19) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext19) AddService() *ATMService23 {
	a.Service = new(ATMService23)
	return a.Service
}

// Context in which the transaction is performed.
type ATMContext2 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService2 `xml:"Svc,omitempty"`
}

func (a *ATMContext2) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext2) AddService() *ATMService2 {
	a.Service = new(ATMService2)
	return a.Service
}

// Context of the exception.
type ATMContext20 struct {

	// Unique identification of the customer session in which the exception occurred.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Service provided by the ATM inside the session.
	Service *ATMService24 `xml:"Svc"`
}

func (a *ATMContext20) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext20) AddService() *ATMService24 {
	a.Service = new(ATMService24)
	return a.Service
}

// Context in which the transaction is performed.
type ATMContext3 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Service provided by the ATM inside the session.
	Service *ATMService3 `xml:"Svc"`
}

func (a *ATMContext3) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext3) AddService() *ATMService3 {
	a.Service = new(ATMService3)
	return a.Service
}

// Context in which the transaction is performed.
type ATMContext4 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService4 `xml:"Svc"`
}

func (a *ATMContext4) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext4) AddService() *ATMService4 {
	a.Service = new(ATMService4)
	return a.Service
}

// Context in which the inquiry is performed.
type ATMContext5 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService5 `xml:"Svc"`
}

func (a *ATMContext5) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext5) AddService() *ATMService5 {
	a.Service = new(ATMService5)
	return a.Service
}

// Context in which the inquiry is performed.
type ATMContext6 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService6 `xml:"Svc"`
}

func (a *ATMContext6) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext6) AddService() *ATMService6 {
	a.Service = new(ATMService6)
	return a.Service
}

// Context in which the transaction is performed.
type ATMContext7 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Context in which the transaction is performed.
	Service *ATMService8 `xml:"Svc"`
}

func (a *ATMContext7) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext7) AddService() *ATMService8 {
	a.Service = new(ATMService8)
	return a.Service
}

// Context in which the transaction is performed.
type ATMContext8 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService9 `xml:"Svc,omitempty"`
}

func (a *ATMContext8) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext8) AddService() *ATMService9 {
	a.Service = new(ATMService9)
	return a.Service
}

// Context in which the transaction is performed.
type ATMContext9 struct {

	// Unique identification of the customer session in which the service is performed.
	SessionReference *Max35Text `xml:"SsnRef,omitempty"`

	// Withdrawal service provided by the ATM inside the session.
	Service *ATMService10 `xml:"Svc,omitempty"`
}

func (a *ATMContext9) SetSessionReference(value string) {
	a.SessionReference = (*Max35Text)(&value)
}

func (a *ATMContext9) AddService() *ATMService10 {
	a.Service = new(ATMService10)
	return a.Service
}
