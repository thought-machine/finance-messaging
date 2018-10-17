package iso20022

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader1 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which a (group of) payment instruction(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key that allows to check if the initiating party is allowed to initiate transactions from the account specified in the initiation.
	//
	// Usage: the content is not of a technical nature, but reflects the organisational structure at the initiating side. The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a different party than the initiating party.
	Authorisation []*Max128Text `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions in the message is requested.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether common accounting information in the transaction is included once for all transactions or repeated for each single transaction.
	Grouping *Grouping1Code `xml:"Grpg"`

	// Party initiating the payment. In the payment context, this can either be the debtor (in a credit transfer), the creditor (in a direct debit), or the party that initiates the payment on behalf of the debtor or creditor.
	InitiatingParty *PartyIdentification8 `xml:"InitgPty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain for execution.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"FwdgAgt,omitempty"`
}

func (g *GroupHeader1) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader1) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader1) AddAuthorisation(value string) {
	g.Authorisation = append(g.Authorisation, (*Max128Text)(&value))
}

func (g *GroupHeader1) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader1) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader1) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader1) SetGrouping(value string) {
	g.Grouping = (*Grouping1Code)(&value)
}

func (g *GroupHeader1) AddInitiatingParty() *PartyIdentification8 {
	g.InitiatingParty = new(PartyIdentification8)
	return g.InitiatingParty
}

func (g *GroupHeader1) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.ForwardingAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader2 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which a (group of) payment instruction(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions in the message is requested.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount of money transferred between instructing agent and instructed agent.
	TotalInterbankSettlementAmount *CurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInformation1 `xml:"SttlmInf"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation3 `xml:"PmtTpInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader2) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader2) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader2) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader2) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader2) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader2) SetTotalInterbankSettlementAmount(value, currency string) {
	g.TotalInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (g *GroupHeader2) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader2) AddSettlementInformation() *SettlementInformation1 {
	g.SettlementInformation = new(SettlementInformation1)
	return g.SettlementInformation
}

func (g *GroupHeader2) AddPaymentTypeInformation() *PaymentTypeInformation3 {
	g.PaymentTypeInformation = new(PaymentTypeInformation3)
	return g.PaymentTypeInformation
}

func (g *GroupHeader2) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructingAgent
}

func (g *GroupHeader2) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructedAgent
}

// Set of elements providing further details on the message.
type GroupHeader23 struct {

	// Point to point reference assigned by the account servicing institution and sent to the account owner to unambiguously identify the message.
	//
	// Usage: The account servicing institution has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created by the account servicer.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party that is entitled by the account owner to receive information about movements in the account.
	//
	// Guideline : MessageRecipient should only be identified when different from the account owner.
	MessageRecipient *PartyIdentification8 `xml:"MsgRcpt,omitempty"`

	// Pagination of the message.
	//
	// Usage: the pagination of the message is only allowed when agreed between the parties.
	MessagePagination *Pagination `xml:"MsgPgntn,omitempty"`

	// Further details on the message.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`
}

func (g *GroupHeader23) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader23) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader23) AddMessageRecipient() *PartyIdentification8 {
	g.MessageRecipient = new(PartyIdentification8)
	return g.MessageRecipient
}

func (g *GroupHeader23) AddMessagePagination() *Pagination {
	g.MessagePagination = new(Pagination)
	return g.MessagePagination
}

func (g *GroupHeader23) SetAdditionalInformation(value string) {
	g.AdditionalInformation = (*Max500Text)(&value)
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader3 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which a (group of) payment instruction(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key that allows to check if the initiating party is allowed to initiate transactions from the account specified in the initiation.
	//
	// Usage: the content is not of a technical nature, but reflects the organisational structure at the initiating side. The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a different party than the initiating party.
	Authorisation []*Max128Text `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions in the message is requested.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount of money transferred between instructing agent and instructed agent.
	TotalInterbankSettlementAmount *CurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInformation2 `xml:"SttlmInf"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation4 `xml:"PmtTpInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader3) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader3) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader3) AddAuthorisation(value string) {
	g.Authorisation = append(g.Authorisation, (*Max128Text)(&value))
}

func (g *GroupHeader3) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader3) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader3) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader3) SetTotalInterbankSettlementAmount(value, currency string) {
	g.TotalInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (g *GroupHeader3) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader3) AddSettlementInformation() *SettlementInformation2 {
	g.SettlementInformation = new(SettlementInformation2)
	return g.SettlementInformation
}

func (g *GroupHeader3) AddPaymentTypeInformation() *PaymentTypeInformation4 {
	g.PaymentTypeInformation = new(PaymentTypeInformation4)
	return g.PaymentTypeInformation
}

func (g *GroupHeader3) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructingAgent
}

func (g *GroupHeader3) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader31 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the instructed party, to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check the authority of the initiating party.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side. The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Party that initiates the mandate message.
	InitiatingParty *PartyIdentification32 `xml:"InitgPty,omitempty"`

	// Agent that instructs the next party in the chain to carry out an instruction.
	//
	// Usage Rule:
	// In case of amendment and cancellation request messages, the instructing agent is the party sending the amendment and cancellation request message and not the party that sent the original mandate initiation request message.
	// In case of acceptance report message, the instructing agent is the party sending the acceptance report message and not the party that sent the original mandate request message.
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out an instruction.
	//
	// Usage Rule:
	// In case of amendment and cancellation request messages, the instructed agent is the party receiving the amendment and cancellation request message and not the party that received the original mandate initiation request message.
	// In case of acceptance report message, the instructed agent is the party receiving the acceptance report message and not the party that received the original mandate request message.
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader31) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader31) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader31) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader31) AddInitiatingParty() *PartyIdentification32 {
	g.InitiatingParty = new(PartyIdentification32)
	return g.InitiatingParty
}

func (g *GroupHeader31) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructingAgent
}

func (g *GroupHeader31) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader32 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Party that initiates the payment.
	//
	// Usage: This can either be the debtor or the party that initiates the credit transfer on behalf of the debtor.
	InitiatingParty *PartyIdentification32 `xml:"InitgPty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain for execution.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"FwdgAgt,omitempty"`
}

func (g *GroupHeader32) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader32) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader32) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader32) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader32) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader32) AddInitiatingParty() *PartyIdentification32 {
	g.InitiatingParty = new(PartyIdentification32)
	return g.InitiatingParty
}

func (g *GroupHeader32) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.ForwardingAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader33 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent.
	TotalInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInformation13 `xml:"SttlmInf"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation21 `xml:"PmtTpInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader33) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader33) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader33) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader33) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader33) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader33) SetTotalInterbankSettlementAmount(value, currency string) {
	g.TotalInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader33) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader33) AddSettlementInformation() *SettlementInformation13 {
	g.SettlementInformation = new(SettlementInformation13)
	return g.SettlementInformation
}

func (g *GroupHeader33) AddPaymentTypeInformation() *PaymentTypeInformation21 {
	g.PaymentTypeInformation = new(PaymentTypeInformation21)
	return g.PaymentTypeInformation
}

func (g *GroupHeader33) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructingAgent
}

func (g *GroupHeader33) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader34 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent.
	TotalInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInformation14 `xml:"SttlmInf"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation22 `xml:"PmtTpInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader34) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader34) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader34) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader34) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader34) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader34) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader34) SetTotalInterbankSettlementAmount(value, currency string) {
	g.TotalInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader34) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader34) AddSettlementInformation() *SettlementInformation14 {
	g.SettlementInformation = new(SettlementInformation14)
	return g.SettlementInformation
}

func (g *GroupHeader34) AddPaymentTypeInformation() *PaymentTypeInformation22 {
	g.PaymentTypeInformation = new(PaymentTypeInformation22)
	return g.PaymentTypeInformation
}

func (g *GroupHeader34) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructingAgent
}

func (g *GroupHeader34) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader35 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent.
	TotalInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInformation13 `xml:"SttlmInf"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation23 `xml:"PmtTpInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader35) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader35) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader35) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader35) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader35) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader35) SetTotalInterbankSettlementAmount(value, currency string) {
	g.TotalInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader35) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader35) AddSettlementInformation() *SettlementInformation13 {
	g.SettlementInformation = new(SettlementInformation13)
	return g.SettlementInformation
}

func (g *GroupHeader35) AddPaymentTypeInformation() *PaymentTypeInformation23 {
	g.PaymentTypeInformation = new(PaymentTypeInformation23)
	return g.PaymentTypeInformation
}

func (g *GroupHeader35) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructingAgent
}

func (g *GroupHeader35) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader36 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party that initiates the status message.
	InitiatingParty *PartyIdentification32 `xml:"InitgPty,omitempty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"FwdgAgt,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`
}

func (g *GroupHeader36) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader36) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader36) AddInitiatingParty() *PartyIdentification32 {
	g.InitiatingParty = new(PartyIdentification32)
	return g.InitiatingParty
}

func (g *GroupHeader36) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.ForwardingAgent
}

func (g *GroupHeader36) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.DebtorAgent
}

func (g *GroupHeader36) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.CreditorAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader37 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader37) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader37) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader37) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructingAgent
}

func (g *GroupHeader37) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader38 struct {

	// Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the return applies to the whole group of transactions or to individual transactions within the original group(s).
	GroupReturn *TrueFalseIndicator `xml:"GrpRtr,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent in the return message.
	TotalReturnedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlRtrdIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInformation13 `xml:"SttlmInf"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	// Usage: The instructing agent is the party sending the return message and not the party that sent the original instruction that is being returned.
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	// Usage: The instructed agent is the party receiving the return message and not the party that received the original instruction that is being returned.
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader38) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader38) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader38) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader38) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader38) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader38) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader38) SetGroupReturn(value string) {
	g.GroupReturn = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader38) SetTotalReturnedInterbankSettlementAmount(value, currency string) {
	g.TotalReturnedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader38) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader38) AddSettlementInformation() *SettlementInformation13 {
	g.SettlementInformation = new(SettlementInformation13)
	return g.SettlementInformation
}

func (g *GroupHeader38) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructingAgent
}

func (g *GroupHeader38) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader39 struct {

	// Point to point reference, assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which a (group of) payment instruction(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key that allows to check if the initiating party is allowed to initiate transactions from the account specified in the initiation.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side. The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a different party than the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Party that initiates the payment.
	//
	// Usage: This can either be the creditor or a party that initiates the direct debit on behalf of the creditor.
	InitiatingParty *PartyIdentification32 `xml:"InitgPty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain for execution.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"FwdgAgt,omitempty"`
}

func (g *GroupHeader39) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader39) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader39) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader39) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader39) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader39) AddInitiatingParty() *PartyIdentification32 {
	g.InitiatingParty = new(PartyIdentification32)
	return g.InitiatingParty
}

func (g *GroupHeader39) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.ForwardingAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader4 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which a (group of) payment instruction(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions in the message is requested.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount of money transferred between instructing agent and instructed agent.
	TotalInterbankSettlementAmount *CurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInformation1 `xml:"SttlmInf"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation5 `xml:"PmtTpInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader4) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader4) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader4) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader4) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader4) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader4) SetTotalInterbankSettlementAmount(value, currency string) {
	g.TotalInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (g *GroupHeader4) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader4) AddSettlementInformation() *SettlementInformation1 {
	g.SettlementInformation = new(SettlementInformation1)
	return g.SettlementInformation
}

func (g *GroupHeader4) AddPaymentTypeInformation() *PaymentTypeInformation5 {
	g.PaymentTypeInformation = new(PaymentTypeInformation5)
	return g.PaymentTypeInformation
}

func (g *GroupHeader4) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructingAgent
}

func (g *GroupHeader4) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader40 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the reversal applies to the whole group of transactions or to individual transactions within the original group.
	GroupReversal *TrueFalseIndicator `xml:"GrpRvsl,omitempty"`

	// Party that initiates the reversal message.
	// Usage: This can be either the creditor or a party that initiates the reversal of the direct debit on behalf of the creditor.
	InitiatingParty *PartyIdentification32 `xml:"InitgPty,omitempty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"FwdgAgt,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`
}

func (g *GroupHeader40) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader40) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader40) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader40) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader40) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader40) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader40) SetGroupReversal(value string) {
	g.GroupReversal = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader40) AddInitiatingParty() *PartyIdentification32 {
	g.InitiatingParty = new(PartyIdentification32)
	return g.InitiatingParty
}

func (g *GroupHeader40) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.ForwardingAgent
}

func (g *GroupHeader40) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.DebtorAgent
}

func (g *GroupHeader40) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.CreditorAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader41 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the reversal applies to the whole group of transactions or to individual transactions within the original group.
	GroupReversal *TrueFalseIndicator `xml:"GrpRvsl,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent in the reversal message.
	TotalReversedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlRvsdIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInformation13 `xml:"SttlmInf"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the reversal message and not the party that received the original instruction that is being reversed.
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader41) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader41) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader41) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader41) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader41) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader41) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader41) SetGroupReversal(value string) {
	g.GroupReversal = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader41) SetTotalReversedInterbankSettlementAmount(value, currency string) {
	g.TotalReversedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader41) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader41) AddSettlementInformation() *SettlementInformation13 {
	g.SettlementInformation = new(SettlementInformation13)
	return g.SettlementInformation
}

func (g *GroupHeader41) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructingAgent
}

func (g *GroupHeader41) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return g.InstructedAgent
}

// Set of elements used to provide further details of the message.
type GroupHeader42 struct {

	// Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	// Usage: The account servicing institution has to make sure that MessageIdentification is unique per account owner for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party authorised by the account owner to receive information about movements on the account.
	// Usage: MessageRecipient should only be identified when different from the account owner.
	MessageRecipient *PartyIdentification32 `xml:"MsgRcpt,omitempty"`

	// Set of elements used to provide details on the page number of the message.
	//
	// Usage: The pagination of the message is only allowed when agreed between the parties.
	MessagePagination *Pagination `xml:"MsgPgntn,omitempty"`

	// Further details of the message.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`
}

func (g *GroupHeader42) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader42) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader42) AddMessageRecipient() *PartyIdentification32 {
	g.MessageRecipient = new(PartyIdentification32)
	return g.MessageRecipient
}

func (g *GroupHeader42) AddMessagePagination() *Pagination {
	g.MessagePagination = new(Pagination)
	return g.MessagePagination
}

func (g *GroupHeader42) SetAdditionalInformation(value string) {
	g.AdditionalInformation = (*Max500Text)(&value)
}

// Set of elements used to provide further details on the message.
type GroupHeader43 struct {

	// Point to point reference, as assigned by the account owner or the party authorised to send the message, and sent to the account servicing institution, to unambiguously identify the message.
	// Usage: The sender has to make sure that 'MessageIdentification' is unique per account servicing institution for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identification of the party that is sending the message, when different from the account owner.
	MessageSender *Party7Choice `xml:"MsgSndr,omitempty"`
}

func (g *GroupHeader43) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader43) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader43) AddMessageSender() *Party7Choice {
	g.MessageSender = new(Party7Choice)
	return g.MessageSender
}

// Set of elements used to provide further details on the message.
type GroupHeader44 struct {

	// Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	//
	// Usage: The account servicing institution has to make sure that 'MessageIdentification' is unique per account owner for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identification of the party that is receiving the message, when different from the account owner.
	MessageRecipient *Party7Choice `xml:"MsgRcpt,omitempty"`
}

func (g *GroupHeader44) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader44) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader44) AddMessageRecipient() *Party7Choice {
	g.MessageRecipient = new(Party7Choice)
	return g.MessageRecipient
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader45 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which a (group of) payment instruction(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Party initiating the creditor payment activation request. This can either be the creditor himself or the party that initiates the request on behalf of the creditor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty"`
}

func (g *GroupHeader45) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader45) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader45) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader45) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader45) AddInitiatingParty() *PartyIdentification43 {
	g.InitiatingParty = new(PartyIdentification43)
	return g.InitiatingParty
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader46 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the status report was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party initiating the creditor payment activation request. This can either be the creditor himself or the party that initiates the request on behalf of the creditor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`
}

func (g *GroupHeader46) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader46) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader46) AddInitiatingParty() *PartyIdentification43 {
	g.InitiatingParty = new(PartyIdentification43)
	return g.InitiatingParty
}

func (g *GroupHeader46) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.DebtorAgent
}

func (g *GroupHeader46) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.CreditorAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader47 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the instructed party, to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check the authority of the initiating party.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side. The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Party that initiates the mandate message.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty,omitempty"`

	// Agent that instructs the next party in the chain to carry out an instruction.
	//
	// Usage Rule:
	// In case of amendment and cancellation request messages, the instructing agent is the party sending the amendment and cancellation request message and not the party that sent the original mandate initiation request message.
	// In case of acceptance report message, the instructing agent is the party sending the acceptance report message and not the party that sent the original mandate request message.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out an instruction.
	//
	// Usage Rule:
	// In case of amendment and cancellation request messages, the instructed agent is the party receiving the amendment and cancellation request message and not the party that received the original mandate initiation request message.
	// In case of acceptance report message, the instructed agent is the party receiving the acceptance report message and not the party that received the original mandate request message.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader47) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader47) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader47) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader47) AddInitiatingParty() *PartyIdentification43 {
	g.InitiatingParty = new(PartyIdentification43)
	return g.InitiatingParty
}

func (g *GroupHeader47) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructingAgent
}

func (g *GroupHeader47) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader48 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Party that initiates the payment.
	//
	// Usage: This can either be the debtor or the party that initiates the credit transfer on behalf of the debtor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain for execution.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"FwdgAgt,omitempty"`
}

func (g *GroupHeader48) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader48) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader48) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader48) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader48) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader48) AddInitiatingParty() *PartyIdentification43 {
	g.InitiatingParty = new(PartyIdentification43)
	return g.InitiatingParty
}

func (g *GroupHeader48) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.ForwardingAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader49 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent.
	TotalInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInstruction1 `xml:"SttlmInf"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation21 `xml:"PmtTpInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader49) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader49) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader49) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader49) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader49) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader49) SetTotalInterbankSettlementAmount(value, currency string) {
	g.TotalInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader49) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader49) AddSettlementInformation() *SettlementInstruction1 {
	g.SettlementInformation = new(SettlementInstruction1)
	return g.SettlementInformation
}

func (g *GroupHeader49) AddPaymentTypeInformation() *PaymentTypeInformation21 {
	g.PaymentTypeInformation = new(PaymentTypeInformation21)
	return g.PaymentTypeInformation
}

func (g *GroupHeader49) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructingAgent
}

func (g *GroupHeader49) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader5 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the status report was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party initiating the payment. In the payment context, this can either be the debtor (in a credit transfer), the creditor (in a direct debit), or the party that initiates the payment on behalf of the debtor or creditor.
	InitiatingParty *PartyIdentification8 `xml:"InitgPty,omitempty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"FwdgAgt,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader5) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader5) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader5) AddInitiatingParty() *PartyIdentification8 {
	g.InitiatingParty = new(PartyIdentification8)
	return g.InitiatingParty
}

func (g *GroupHeader5) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.ForwardingAgent
}

func (g *GroupHeader5) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.DebtorAgent
}

func (g *GroupHeader5) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.CreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.CreditorAgent
}

func (g *GroupHeader5) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructingAgent
}

func (g *GroupHeader5) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader50 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent.
	TotalInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInstruction2 `xml:"SttlmInf"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader50) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader50) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader50) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader50) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader50) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader50) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader50) SetTotalInterbankSettlementAmount(value, currency string) {
	g.TotalInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader50) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader50) AddSettlementInformation() *SettlementInstruction2 {
	g.SettlementInformation = new(SettlementInstruction2)
	return g.SettlementInformation
}

func (g *GroupHeader50) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	g.PaymentTypeInformation = new(PaymentTypeInformation25)
	return g.PaymentTypeInformation
}

func (g *GroupHeader50) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructingAgent
}

func (g *GroupHeader50) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader51 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent.
	TotalInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInstruction1 `xml:"SttlmInf"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation21 `xml:"PmtTpInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader51) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader51) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader51) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader51) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader51) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader51) SetTotalInterbankSettlementAmount(value, currency string) {
	g.TotalInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader51) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader51) AddSettlementInformation() *SettlementInstruction1 {
	g.SettlementInformation = new(SettlementInstruction1)
	return g.SettlementInformation
}

func (g *GroupHeader51) AddPaymentTypeInformation() *PaymentTypeInformation21 {
	g.PaymentTypeInformation = new(PaymentTypeInformation21)
	return g.PaymentTypeInformation
}

func (g *GroupHeader51) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructingAgent
}

func (g *GroupHeader51) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader52 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party that initiates the status message.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty,omitempty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"FwdgAgt,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`
}

func (g *GroupHeader52) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader52) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader52) AddInitiatingParty() *PartyIdentification43 {
	g.InitiatingParty = new(PartyIdentification43)
	return g.InitiatingParty
}

func (g *GroupHeader52) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.ForwardingAgent
}

func (g *GroupHeader52) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.DebtorAgent
}

func (g *GroupHeader52) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.CreditorAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader53 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader53) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader53) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader53) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructingAgent
}

func (g *GroupHeader53) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader54 struct {

	// Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the return applies to the whole group of transactions or to individual transactions within the original group(s).
	GroupReturn *TrueFalseIndicator `xml:"GrpRtr,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent in the return message.
	TotalReturnedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlRtrdIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInstruction1 `xml:"SttlmInf"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	// Usage: The instructing agent is the party sending the return message and not the party that sent the original instruction that is being returned.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	// Usage: The instructed agent is the party receiving the return message and not the party that received the original instruction that is being returned.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader54) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader54) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader54) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader54) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader54) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader54) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader54) SetGroupReturn(value string) {
	g.GroupReturn = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader54) SetTotalReturnedInterbankSettlementAmount(value, currency string) {
	g.TotalReturnedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader54) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader54) AddSettlementInformation() *SettlementInstruction1 {
	g.SettlementInformation = new(SettlementInstruction1)
	return g.SettlementInformation
}

func (g *GroupHeader54) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructingAgent
}

func (g *GroupHeader54) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader55 struct {

	// Point to point reference, assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which a (group of) payment instruction(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key that allows to check if the initiating party is allowed to initiate transactions from the account specified in the initiation.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side. The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a different party than the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Party that initiates the payment.
	//
	// Usage: This can either be the creditor or a party that initiates the direct debit on behalf of the creditor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain for execution.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"FwdgAgt,omitempty"`
}

func (g *GroupHeader55) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader55) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader55) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader55) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader55) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader55) AddInitiatingParty() *PartyIdentification43 {
	g.InitiatingParty = new(PartyIdentification43)
	return g.InitiatingParty
}

func (g *GroupHeader55) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.ForwardingAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader56 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the reversal applies to the whole group of transactions or to individual transactions within the original group.
	GroupReversal *TrueFalseIndicator `xml:"GrpRvsl,omitempty"`

	// Party that initiates the reversal message.
	// Usage: This can be either the creditor or a party that initiates the reversal of the direct debit on behalf of the creditor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty,omitempty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"FwdgAgt,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`
}

func (g *GroupHeader56) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader56) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader56) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader56) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader56) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader56) SetGroupReversal(value string) {
	g.GroupReversal = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader56) AddInitiatingParty() *PartyIdentification43 {
	g.InitiatingParty = new(PartyIdentification43)
	return g.InitiatingParty
}

func (g *GroupHeader56) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.ForwardingAgent
}

func (g *GroupHeader56) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.DebtorAgent
}

func (g *GroupHeader56) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.CreditorAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader57 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the reversal applies to the whole group of transactions or to individual transactions within the original group.
	GroupReversal *TrueFalseIndicator `xml:"GrpRvsl,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent in the reversal message.
	TotalReversedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlRvsdIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInstruction1 `xml:"SttlmInf"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the reversal message and not the party that received the original instruction that is being reversed.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader57) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader57) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader57) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader57) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader57) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader57) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader57) SetGroupReversal(value string) {
	g.GroupReversal = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader57) SetTotalReversedInterbankSettlementAmount(value, currency string) {
	g.TotalReversedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader57) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader57) AddSettlementInformation() *SettlementInstruction1 {
	g.SettlementInformation = new(SettlementInstruction1)
	return g.SettlementInformation
}

func (g *GroupHeader57) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructingAgent
}

func (g *GroupHeader57) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructedAgent
}

// Provides further details on the message.
type GroupHeader58 struct {

	// Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	// Usage: The account servicing institution has to make sure that MessageIdentification is unique per account owner for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party authorised by the account owner to receive information about movements on the account.
	// Usage: MessageRecipient should only be identified when different from the account owner.
	MessageRecipient *PartyIdentification43 `xml:"MsgRcpt,omitempty"`

	// Provides details on the page number of the message.
	//
	// Usage: The pagination of the message is only allowed when agreed between the parties.
	MessagePagination *Pagination `xml:"MsgPgntn,omitempty"`

	// Unique identification, as assigned by the original requestor, to unambiguously identify the business query message.
	OriginalBusinessQuery *OriginalBusinessQuery1 `xml:"OrgnlBizQry,omitempty"`

	// Further details of the message.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`
}

func (g *GroupHeader58) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader58) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader58) AddMessageRecipient() *PartyIdentification43 {
	g.MessageRecipient = new(PartyIdentification43)
	return g.MessageRecipient
}

func (g *GroupHeader58) AddMessagePagination() *Pagination {
	g.MessagePagination = new(Pagination)
	return g.MessagePagination
}

func (g *GroupHeader58) AddOriginalBusinessQuery() *OriginalBusinessQuery1 {
	g.OriginalBusinessQuery = new(OriginalBusinessQuery1)
	return g.OriginalBusinessQuery
}

func (g *GroupHeader58) SetAdditionalInformation(value string) {
	g.AdditionalInformation = (*Max500Text)(&value)
}

// Provides further details on the message.
type GroupHeader59 struct {

	// Point to point reference, as assigned by the account owner or the party authorised to send the message, and sent to the account servicing institution, to unambiguously identify the message.
	// Usage: The sender has to make sure that 'MessageIdentification' is unique per account servicing institution for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identification of the party that is sending the message, when different from the account owner.
	MessageSender *Party12Choice `xml:"MsgSndr,omitempty"`
}

func (g *GroupHeader59) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader59) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader59) AddMessageSender() *Party12Choice {
	g.MessageSender = new(Party12Choice)
	return g.MessageSender
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader6 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which a (group of) payment instruction(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key that allows to check if the initiating party is allowed to initiate transactions from the account specified in the initiation.
	//
	// Usage: the content is not of a technical nature, but reflects the organisational structure at the initiating side. The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a different party than the initiating party.
	Authorisation []*Max128Text `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions in the message is requested.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the return applies to the whole group of transactions or to individual transactions within the original group(s).
	GroupReturn *TrueFalseIndicator `xml:"GrpRtr,omitempty"`

	// Total returned amount of money transferred between the instructing agent and the instructed agent in the return message.
	TotalReturnedInterbankSettlementAmount *CurrencyAndAmount `xml:"TtlRtrdIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInformation1 `xml:"SttlmInf"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader6) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader6) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader6) AddAuthorisation(value string) {
	g.Authorisation = append(g.Authorisation, (*Max128Text)(&value))
}

func (g *GroupHeader6) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader6) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader6) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader6) SetGroupReturn(value string) {
	g.GroupReturn = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader6) SetTotalReturnedInterbankSettlementAmount(value, currency string) {
	g.TotalReturnedInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (g *GroupHeader6) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader6) AddSettlementInformation() *SettlementInformation1 {
	g.SettlementInformation = new(SettlementInformation1)
	return g.SettlementInformation
}

func (g *GroupHeader6) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructingAgent
}

func (g *GroupHeader6) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructedAgent
}

// Provides further details on the message.
type GroupHeader60 struct {

	// Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	//
	// Usage: The account servicing institution has to make sure that 'MessageIdentification' is unique per account owner for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identification of the party that is receiving the message, when different from the account owner.
	MessageRecipient *Party12Choice `xml:"MsgRcpt,omitempty"`
}

func (g *GroupHeader60) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader60) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader60) AddMessageRecipient() *Party12Choice {
	g.MessageRecipient = new(Party12Choice)
	return g.MessageRecipient
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader62 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Specifies if this document is a copy, a duplicate, or a duplicate of a copy.
	CopyIndicator *CopyDuplicate1Code `xml:"CpyInd,omitempty"`

	// Party that initiates the payment.
	//
	// Usage: This can either be the debtor or the party that initiates the credit transfer on behalf of the debtor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty"`

	// Party authorised by the account owner to receive information about movements on the account.
	MessageRecipient *PartyIdentification43 `xml:"MsgRcpt,omitempty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain for execution.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"FwdgAgt,omitempty"`
}

func (g *GroupHeader62) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader62) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader62) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader62) SetCopyIndicator(value string) {
	g.CopyIndicator = (*CopyDuplicate1Code)(&value)
}

func (g *GroupHeader62) AddInitiatingParty() *PartyIdentification43 {
	g.InitiatingParty = new(PartyIdentification43)
	return g.InitiatingParty
}

func (g *GroupHeader62) AddMessageRecipient() *PartyIdentification43 {
	g.MessageRecipient = new(PartyIdentification43)
	return g.MessageRecipient
}

func (g *GroupHeader62) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.ForwardingAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader63 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader63) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader63) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader63) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader63) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader63) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructingAgent
}

func (g *GroupHeader63) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader69 struct {

	// Report identification, for example invoice number or report number from point of sales system.
	Identification *Max35Text `xml:"Id"`

	// Date at which the status report was created.
	IssuedDate *ISODate `xml:"IssdDt"`

	// Specifies if the report is based on debit invoice, credit invoice, card transaction or cash transaction.
	ReportCategory *ExternalDocumentType1Code `xml:"RptCtgy"`

	// Specifies if the TaxReport is new, correction or remove.
	TaxReportPurpose *ExternalDocumentType1Code `xml:"TaxRptPurp"`

	// Original tax report identification, used for example original invoice number with credit notes.
	OriginalIdentification *Max35Text `xml:"OrgnlId,omitempty"`

	// Details of tax representative. The corporate (seller) is allowed to use a tax representative for value added tax responsibilities in case the seller is not registered in a specific value added tax registry.
	SellerTaxRepresentative *PartyIdentification116 `xml:"SellrTaxRprtv,omitempty"`

	// Details of tax representative. The corporate (buyer) is allowed to use a tax representative for value added tax responsibilities in  case the buyer is not registered in a specific value added tax registry.
	BuyerTaxRepresentative *PartyIdentification116 `xml:"BuyrTaxRprtv,omitempty"`

	// Specifies the language used in the message.
	LanguageCode *LanguageCode `xml:"LangCd,omitempty"`
}

func (g *GroupHeader69) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GroupHeader69) SetIssuedDate(value string) {
	g.IssuedDate = (*ISODate)(&value)
}

func (g *GroupHeader69) SetReportCategory(value string) {
	g.ReportCategory = (*ExternalDocumentType1Code)(&value)
}

func (g *GroupHeader69) SetTaxReportPurpose(value string) {
	g.TaxReportPurpose = (*ExternalDocumentType1Code)(&value)
}

func (g *GroupHeader69) SetOriginalIdentification(value string) {
	g.OriginalIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader69) AddSellerTaxRepresentative() *PartyIdentification116 {
	g.SellerTaxRepresentative = new(PartyIdentification116)
	return g.SellerTaxRepresentative
}

func (g *GroupHeader69) AddBuyerTaxRepresentative() *PartyIdentification116 {
	g.BuyerTaxRepresentative = new(PartyIdentification116)
	return g.BuyerTaxRepresentative
}

func (g *GroupHeader69) SetLanguageCode(value string) {
	g.LanguageCode = (*LanguageCode)(&value)
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader7 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the (group of) cancellation request(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the cancellation applies to a whole group of transactions or to individual transactions within the original group.
	GroupCancellation *GroupingIndicator `xml:"GrpCxl,omitempty"`

	// Party initiating the payment. In the payment context, this can either be the debtor (in a credit transfer), the creditor (in a direct debit), or the party that initiates the payment on behalf of the debtor or creditor.
	InitiatingParty *PartyIdentification8 `xml:"InitgPty,omitempty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"FwdgAgt,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader7) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader7) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader7) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader7) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader7) SetGroupCancellation(value string) {
	g.GroupCancellation = (*GroupingIndicator)(&value)
}

func (g *GroupHeader7) AddInitiatingParty() *PartyIdentification8 {
	g.InitiatingParty = new(PartyIdentification8)
	return g.InitiatingParty
}

func (g *GroupHeader7) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.ForwardingAgent
}

func (g *GroupHeader7) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.DebtorAgent
}

func (g *GroupHeader7) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.CreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.CreditorAgent
}

func (g *GroupHeader7) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructingAgent
}

func (g *GroupHeader7) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader70 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent.
	TotalInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInstruction4 `xml:"SttlmInf"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation21 `xml:"PmtTpInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader70) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader70) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader70) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader70) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader70) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader70) SetTotalInterbankSettlementAmount(value, currency string) {
	g.TotalInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader70) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader70) AddSettlementInformation() *SettlementInstruction4 {
	g.SettlementInformation = new(SettlementInstruction4)
	return g.SettlementInformation
}

func (g *GroupHeader70) AddPaymentTypeInformation() *PaymentTypeInformation21 {
	g.PaymentTypeInformation = new(PaymentTypeInformation21)
	return g.PaymentTypeInformation
}

func (g *GroupHeader70) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructingAgent
}

func (g *GroupHeader70) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader71 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the reversal applies to the whole group of transactions or to individual transactions within the original group.
	GroupReversal *TrueFalseIndicator `xml:"GrpRvsl,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent in the reversal message.
	TotalReversedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlRvsdIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInstruction4 `xml:"SttlmInf"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the reversal message and not the party that received the original instruction that is being reversed.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader71) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader71) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader71) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader71) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader71) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader71) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader71) SetGroupReversal(value string) {
	g.GroupReversal = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader71) SetTotalReversedInterbankSettlementAmount(value, currency string) {
	g.TotalReversedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader71) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader71) AddSettlementInformation() *SettlementInstruction4 {
	g.SettlementInformation = new(SettlementInstruction4)
	return g.SettlementInformation
}

func (g *GroupHeader71) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructingAgent
}

func (g *GroupHeader71) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader72 struct {

	// Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authorisation []*Authorisation1Choice `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the return applies to the whole group of transactions or to individual transactions within the original group(s).
	GroupReturn *TrueFalseIndicator `xml:"GrpRtr,omitempty"`

	// Total amount of money moved between the instructing agent and the instructed agent in the return message.
	TotalReturnedInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlRtrdIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInstruction4 `xml:"SttlmInf"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	// Usage: The instructing agent is the party sending the return message and not the party that sent the original instruction that is being returned.
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	// Usage: The instructed agent is the party receiving the return message and not the party that received the original instruction that is being returned.
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader72) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader72) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader72) AddAuthorisation() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authorisation = append(g.Authorisation, newValue)
	return newValue
}

func (g *GroupHeader72) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader72) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader72) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader72) SetGroupReturn(value string) {
	g.GroupReturn = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader72) SetTotalReturnedInterbankSettlementAmount(value, currency string) {
	g.TotalReturnedInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader72) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader72) AddSettlementInformation() *SettlementInstruction4 {
	g.SettlementInformation = new(SettlementInstruction4)
	return g.SettlementInformation
}

func (g *GroupHeader72) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructingAgent
}

func (g *GroupHeader72) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstructedAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader8 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which a (group of) payment instruction(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key that allows to check if the initiating party is allowed to initiate transactions from the account specified in the initiation.
	//
	// Usage: the content is not of a technical nature, but reflects the organisational structure at the initiating side. The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a different party than the initiating party.
	Authorisation []*Max128Text `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions in the message is requested.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the reversal applies to the whole group of transactions or to individual transactions within the original group.
	GroupReversal *TrueFalseIndicator `xml:"GrpRvsl,omitempty"`

	// Party initiating the payment. In the payment context, this can either be the debtor (in a credit transfer), the creditor (in a direct debit), or the party that initiates the payment on behalf of the debtor or creditor.
	InitiatingParty *PartyIdentification8 `xml:"InitgPty,omitempty"`

	// Financial institution that receives the instruction from the initiating party and forwards it to the next agent in the payment chain.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"FwdgAgt,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt,omitempty"`
}

func (g *GroupHeader8) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader8) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader8) AddAuthorisation(value string) {
	g.Authorisation = append(g.Authorisation, (*Max128Text)(&value))
}

func (g *GroupHeader8) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader8) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader8) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader8) SetGroupReversal(value string) {
	g.GroupReversal = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader8) AddInitiatingParty() *PartyIdentification8 {
	g.InitiatingParty = new(PartyIdentification8)
	return g.InitiatingParty
}

func (g *GroupHeader8) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.ForwardingAgent
}

func (g *GroupHeader8) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.DebtorAgent
}

func (g *GroupHeader8) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.CreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.CreditorAgent
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader9 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which a (group of) payment instruction(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key that allows to check if the initiating party is allowed to initiate transactions from the account specified in the initiation.
	//
	// Usage: the content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a different party than the initiating party.
	Authorisation []*Max128Text `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions in the message is requested.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the reversal applies to the whole group of transactions or to individual transactions within the original group.
	GroupReversal *TrueFalseIndicator `xml:"GrpRvsl,omitempty"`

	// Total reversed amount of money transferred between the instructing agent and the instructed agent in the reversal message.
	TotalReversedInterbankSettlementAmount *CurrencyAndAmount `xml:"TtlRvsdIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInformation1 `xml:"SttlmInf"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader9) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader9) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader9) AddAuthorisation(value string) {
	g.Authorisation = append(g.Authorisation, (*Max128Text)(&value))
}

func (g *GroupHeader9) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader9) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader9) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader9) SetGroupReversal(value string) {
	g.GroupReversal = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader9) SetTotalReversedInterbankSettlementAmount(value, currency string) {
	g.TotalReversedInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (g *GroupHeader9) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader9) AddSettlementInformation() *SettlementInformation1 {
	g.SettlementInformation = new(SettlementInformation1)
	return g.SettlementInformation
}

func (g *GroupHeader9) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructingAgent
}

func (g *GroupHeader9) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructedAgent
}
