package iso20022

// Environment of the withdrawal transaction.
type ATMEnvironment1 struct {

	// Acquirer of the withdrawal transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine1 `xml:"ATM"`

	// Customer involved in the withdrawal transaction.
	Customer *ATMCustomer1 `xml:"Cstmr"`

	// Card performing the withdrawal transaction.
	Card *PaymentCard16 `xml:"Card,omitempty"`
}

func (a *ATMEnvironment1) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment1) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment1) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment1) AddATM() *AutomatedTellerMachine1 {
	a.ATM = new(AutomatedTellerMachine1)
	return a.ATM
}

func (a *ATMEnvironment1) AddCustomer() *ATMCustomer1 {
	a.Customer = new(ATMCustomer1)
	return a.Customer
}

func (a *ATMEnvironment1) AddCard() *PaymentCard16 {
	a.Card = new(PaymentCard16)
	return a.Card
}

// Environment of the transaction.
type ATMEnvironment10 struct {

	// Acquirer of the ATM transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine8 `xml:"ATM"`
}

func (a *ATMEnvironment10) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment10) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment10) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment10) AddATM() *AutomatedTellerMachine8 {
	a.ATM = new(AutomatedTellerMachine8)
	return a.ATM
}

// Environment of the withdrawal transaction.
type ATMEnvironment11 struct {

	// Acquirer of the transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine9 `xml:"ATM"`

	// Customer involved in the transaction.
	Customer *ATMCustomer4 `xml:"Cstmr"`

	// Card performing the transaction.
	Card *PaymentCard22 `xml:"Card,omitempty"`
}

func (a *ATMEnvironment11) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment11) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment11) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment11) AddATM() *AutomatedTellerMachine9 {
	a.ATM = new(AutomatedTellerMachine9)
	return a.ATM
}

func (a *ATMEnvironment11) AddCustomer() *ATMCustomer4 {
	a.Customer = new(ATMCustomer4)
	return a.Customer
}

func (a *ATMEnvironment11) AddCard() *PaymentCard22 {
	a.Card = new(PaymentCard22)
	return a.Card
}

// Environment of the withdrawal transaction.
type ATMEnvironment12 struct {

	// Acquirer of the transactions, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Institution in charge of managing the ATM.
	ATMManager *Acquirer8 `xml:"ATMMgr,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine2 `xml:"ATM"`

	// Customer involved in the transaction.
	Customer *ATMCustomer5 `xml:"Cstmr"`

	// Encryption of the sensitive card data.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData19 `xml:"PlainCardData,omitempty"`
}

func (a *ATMEnvironment12) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment12) AddATMManager() *Acquirer8 {
	a.ATMManager = new(Acquirer8)
	return a.ATMManager
}

func (a *ATMEnvironment12) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment12) AddATM() *AutomatedTellerMachine2 {
	a.ATM = new(AutomatedTellerMachine2)
	return a.ATM
}

func (a *ATMEnvironment12) AddCustomer() *ATMCustomer5 {
	a.Customer = new(ATMCustomer5)
	return a.Customer
}

func (a *ATMEnvironment12) AddProtectedCardData() *ContentInformationType10 {
	a.ProtectedCardData = new(ContentInformationType10)
	return a.ProtectedCardData
}

func (a *ATMEnvironment12) AddPlainCardData() *PlainCardData19 {
	a.PlainCardData = new(PlainCardData19)
	return a.PlainCardData
}

// Environment of the withdrawal transaction.
type ATMEnvironment13 struct {

	// Acquirer of the transactions, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine9 `xml:"ATM"`

	// Customer involved in the transaction.
	Customer *ATMCustomer6 `xml:"Cstmr"`

	// Card performing the transaction.
	Card *PaymentCard23 `xml:"Card,omitempty"`
}

func (a *ATMEnvironment13) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment13) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment13) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment13) AddATM() *AutomatedTellerMachine9 {
	a.ATM = new(AutomatedTellerMachine9)
	return a.ATM
}

func (a *ATMEnvironment13) AddCustomer() *ATMCustomer6 {
	a.Customer = new(ATMCustomer6)
	return a.Customer
}

func (a *ATMEnvironment13) AddCard() *PaymentCard23 {
	a.Card = new(PaymentCard23)
	return a.Card
}

// Environment of the inquiry.
type ATMEnvironment14 struct {

	// Acquirer of the transactions, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine10 `xml:"ATM"`

	// Customer involved in the withdrawal transaction.
	Customer *ATMCustomer4 `xml:"Cstmr,omitempty"`

	// Card performing the withdrawal transaction.
	Card *PaymentCard22 `xml:"Card,omitempty"`
}

func (a *ATMEnvironment14) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment14) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment14) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment14) AddATM() *AutomatedTellerMachine10 {
	a.ATM = new(AutomatedTellerMachine10)
	return a.ATM
}

func (a *ATMEnvironment14) AddCustomer() *ATMCustomer4 {
	a.Customer = new(ATMCustomer4)
	return a.Customer
}

func (a *ATMEnvironment14) AddCard() *PaymentCard22 {
	a.Card = new(PaymentCard22)
	return a.Card
}

// Environment of the key download.
type ATMEnvironment15 struct {

	// Acquirer of the ATM transactions, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine6 `xml:"ATM"`
}

func (a *ATMEnvironment15) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment15) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment15) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment15) AddATM() *AutomatedTellerMachine6 {
	a.ATM = new(AutomatedTellerMachine6)
	return a.ATM
}

// Environment of exceptions.
type ATMEnvironment16 struct {

	// Acquirer of transactions, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine9 `xml:"ATM"`

	// Customer involved in the transaction.
	Customer *ATMCustomer6 `xml:"Cstmr,omitempty"`

	// Card performing the transaction.
	Card *PaymentCard23 `xml:"Card,omitempty"`
}

func (a *ATMEnvironment16) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment16) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment16) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment16) AddATM() *AutomatedTellerMachine9 {
	a.ATM = new(AutomatedTellerMachine9)
	return a.ATM
}

func (a *ATMEnvironment16) AddCustomer() *ATMCustomer6 {
	a.Customer = new(ATMCustomer6)
	return a.Customer
}

func (a *ATMEnvironment16) AddCard() *PaymentCard23 {
	a.Card = new(PaymentCard23)
	return a.Card
}

// Environment of the withdrawal transaction.
type ATMEnvironment2 struct {

	// Acquirer of the withdrawal transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Institution in charge of managing the ATM.
	ATMManager *Acquirer8 `xml:"ATMMgr,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine2 `xml:"ATM"`

	// Customer involved in the withdrawal transaction.
	Customer *ATMCustomer2 `xml:"Cstmr"`

	// Encryption of the sensitive card data.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData14 `xml:"PlainCardData,omitempty"`
}

func (a *ATMEnvironment2) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment2) AddATMManager() *Acquirer8 {
	a.ATMManager = new(Acquirer8)
	return a.ATMManager
}

func (a *ATMEnvironment2) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment2) AddATM() *AutomatedTellerMachine2 {
	a.ATM = new(AutomatedTellerMachine2)
	return a.ATM
}

func (a *ATMEnvironment2) AddCustomer() *ATMCustomer2 {
	a.Customer = new(ATMCustomer2)
	return a.Customer
}

func (a *ATMEnvironment2) AddProtectedCardData() *ContentInformationType10 {
	a.ProtectedCardData = new(ContentInformationType10)
	return a.ProtectedCardData
}

func (a *ATMEnvironment2) AddPlainCardData() *PlainCardData14 {
	a.PlainCardData = new(PlainCardData14)
	return a.PlainCardData
}

// Environment of the withdrawal transaction.
type ATMEnvironment3 struct {

	// Acquirer of the withdrawal transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine1 `xml:"ATM"`

	// Customer involved in the transaction.
	Customer *ATMCustomer3 `xml:"Cstmr"`

	// Card performing the transaction.
	Card *PaymentCard17 `xml:"Card,omitempty"`
}

func (a *ATMEnvironment3) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment3) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment3) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment3) AddATM() *AutomatedTellerMachine1 {
	a.ATM = new(AutomatedTellerMachine1)
	return a.ATM
}

func (a *ATMEnvironment3) AddCustomer() *ATMCustomer3 {
	a.Customer = new(ATMCustomer3)
	return a.Customer
}

func (a *ATMEnvironment3) AddCard() *PaymentCard17 {
	a.Card = new(PaymentCard17)
	return a.Card
}

// Environment of the inquiry.
type ATMEnvironment4 struct {

	// Acquirer of the withdrawal transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine4 `xml:"ATM"`

	// Customer involved in the withdrawal transaction.
	Customer *ATMCustomer1 `xml:"Cstmr,omitempty"`

	// Card performing the withdrawal transaction.
	Card *PaymentCard16 `xml:"Card,omitempty"`
}

func (a *ATMEnvironment4) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment4) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment4) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment4) AddATM() *AutomatedTellerMachine4 {
	a.ATM = new(AutomatedTellerMachine4)
	return a.ATM
}

func (a *ATMEnvironment4) AddCustomer() *ATMCustomer1 {
	a.Customer = new(ATMCustomer1)
	return a.Customer
}

func (a *ATMEnvironment4) AddCard() *PaymentCard16 {
	a.Card = new(PaymentCard16)
	return a.Card
}

// Environment of the transaction.
type ATMEnvironment5 struct {

	// ATM information.
	ATM *AutomatedTellerMachine2 `xml:"ATM"`

	// Customer involved in the withdrawal transaction.
	Customer *ATMCustomer2 `xml:"Cstmr"`

	// Encryption of the sensitive card data.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData14 `xml:"PlainCardData,omitempty"`
}

func (a *ATMEnvironment5) AddATM() *AutomatedTellerMachine2 {
	a.ATM = new(AutomatedTellerMachine2)
	return a.ATM
}

func (a *ATMEnvironment5) AddCustomer() *ATMCustomer2 {
	a.Customer = new(ATMCustomer2)
	return a.Customer
}

func (a *ATMEnvironment5) AddProtectedCardData() *ContentInformationType10 {
	a.ProtectedCardData = new(ContentInformationType10)
	return a.ProtectedCardData
}

func (a *ATMEnvironment5) AddPlainCardData() *PlainCardData14 {
	a.PlainCardData = new(PlainCardData14)
	return a.PlainCardData
}

// Environment of the transaction.
type ATMEnvironment6 struct {

	// Acquirer of the ATM transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine5 `xml:"ATM"`
}

func (a *ATMEnvironment6) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment6) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment6) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment6) AddATM() *AutomatedTellerMachine5 {
	a.ATM = new(AutomatedTellerMachine5)
	return a.ATM
}

// Environment of the transaction.
type ATMEnvironment7 struct {

	// Acquirer of the ATM transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Institution in charge of managing the ATM.
	ATMManager *Acquirer8 `xml:"ATMMgr,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine3 `xml:"ATM"`
}

func (a *ATMEnvironment7) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment7) AddATMManager() *Acquirer8 {
	a.ATMManager = new(Acquirer8)
	return a.ATMManager
}

func (a *ATMEnvironment7) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment7) AddATM() *AutomatedTellerMachine3 {
	a.ATM = new(AutomatedTellerMachine3)
	return a.ATM
}

// Environment of the key download.
type ATMEnvironment8 struct {

	// Acquirer of the ATM transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// Entity hosting the ATM terminal.
	HostingEntity *TerminalHosting1 `xml:"HstgNtty,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine6 `xml:"ATM"`
}

func (a *ATMEnvironment8) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment8) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment8) AddHostingEntity() *TerminalHosting1 {
	a.HostingEntity = new(TerminalHosting1)
	return a.HostingEntity
}

func (a *ATMEnvironment8) AddATM() *AutomatedTellerMachine6 {
	a.ATM = new(AutomatedTellerMachine6)
	return a.ATM
}

// Environment of the ATM.
type ATMEnvironment9 struct {

	// Acquirer of the ATM transaction, in charge of the funds settlement with the issuer.
	Acquirer *Acquirer7 `xml:"Acqrr,omitempty"`

	// Identification of the ATM manager.
	ATMManagerIdentification *Max35Text `xml:"ATMMgrId,omitempty"`

	// ATM information.
	ATM *AutomatedTellerMachine7 `xml:"ATM"`
}

func (a *ATMEnvironment9) AddAcquirer() *Acquirer7 {
	a.Acquirer = new(Acquirer7)
	return a.Acquirer
}

func (a *ATMEnvironment9) SetATMManagerIdentification(value string) {
	a.ATMManagerIdentification = (*Max35Text)(&value)
}

func (a *ATMEnvironment9) AddATM() *AutomatedTellerMachine7 {
	a.ATM = new(AutomatedTellerMachine7)
	return a.ATM
}
