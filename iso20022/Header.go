package iso20022

// Set of characteristics related to the protocol.
type Header1 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction1Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification32 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification32 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`
}

func (h *Header1) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction1Code)(&value)
}

func (h *Header1) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header1) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header1) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header1) AddInitiatingParty() *GenericIdentification32 {
	h.InitiatingParty = new(GenericIdentification32)
	return h.InitiatingParty
}

func (h *Header1) AddRecipientParty() *GenericIdentification32 {
	h.RecipientParty = new(GenericIdentification32)
	return h.RecipientParty
}

func (h *Header1) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the protocol.
type Header10 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction4Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification53 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification53 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`
}

func (h *Header10) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction4Code)(&value)
}

func (h *Header10) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header10) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header10) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header10) AddInitiatingParty() *GenericIdentification53 {
	h.InitiatingParty = new(GenericIdentification53)
	return h.InitiatingParty
}

func (h *Header10) AddRecipientParty() *GenericIdentification53 {
	h.RecipientParty = new(GenericIdentification53)
	return h.RecipientParty
}

func (h *Header10) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the protocol.
type Header11 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction4Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Number of retransmissions of the message.
	ReTransmissionCounter *Max3NumericText `xml:"ReTrnsmssnCntr,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification53 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification53 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`
}

func (h *Header11) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction4Code)(&value)
}

func (h *Header11) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header11) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header11) SetReTransmissionCounter(value string) {
	h.ReTransmissionCounter = (*Max3NumericText)(&value)
}

func (h *Header11) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header11) AddInitiatingParty() *GenericIdentification53 {
	h.InitiatingParty = new(GenericIdentification53)
	return h.InitiatingParty
}

func (h *Header11) AddRecipientParty() *GenericIdentification53 {
	h.RecipientParty = new(GenericIdentification53)
	return h.RecipientParty
}

func (h *Header11) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the transfer of transactions.
type Header12 struct {

	// Indicates if the file transfer is a download or an upload.
	DownloadTransfer *TrueFalseIndicator `xml:"DwnldTrf"`

	// Version of file format.
	FormatVersion *Max6Text `xml:"FrmtVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification53 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification53 `xml:"RcptPty,omitempty"`
}

func (h *Header12) SetDownloadTransfer(value string) {
	h.DownloadTransfer = (*TrueFalseIndicator)(&value)
}

func (h *Header12) SetFormatVersion(value string) {
	h.FormatVersion = (*Max6Text)(&value)
}

func (h *Header12) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header12) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header12) AddInitiatingParty() *GenericIdentification53 {
	h.InitiatingParty = new(GenericIdentification53)
	return h.InitiatingParty
}

func (h *Header12) AddRecipientParty() *GenericIdentification53 {
	h.RecipientParty = new(GenericIdentification53)
	return h.RecipientParty
}

// Set of characteristics related to the protocol after a rejection.
type Header13 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction4Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification53 `xml:"InitgPty,omitempty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification53 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`
}

func (h *Header13) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction4Code)(&value)
}

func (h *Header13) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header13) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header13) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header13) AddInitiatingParty() *GenericIdentification53 {
	h.InitiatingParty = new(GenericIdentification53)
	return h.InitiatingParty
}

func (h *Header13) AddRecipientParty() *GenericIdentification53 {
	h.RecipientParty = new(GenericIdentification53)
	return h.RecipientParty
}

func (h *Header13) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the transfer of transactions.
type Header14 struct {

	// Indicates if the file transfer is a download or an upload.
	DownloadTransfer *TrueFalseIndicator `xml:"DwnldTrf"`

	// Version of file format.
	FormatVersion *Max6Text `xml:"FrmtVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification71 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification71 `xml:"RcptPty,omitempty"`
}

func (h *Header14) SetDownloadTransfer(value string) {
	h.DownloadTransfer = (*TrueFalseIndicator)(&value)
}

func (h *Header14) SetFormatVersion(value string) {
	h.FormatVersion = (*Max6Text)(&value)
}

func (h *Header14) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header14) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header14) AddInitiatingParty() *GenericIdentification71 {
	h.InitiatingParty = new(GenericIdentification71)
	return h.InitiatingParty
}

func (h *Header14) AddRecipientParty() *GenericIdentification71 {
	h.RecipientParty = new(GenericIdentification71)
	return h.RecipientParty
}

// Set of characteristics related to the reject of a transaction.
type Header15 struct {

	// Indicates if the file transfer is a download or an upload.
	DownloadTransfer *TrueFalseIndicator `xml:"DwnldTrf"`

	// Version of file format.
	FormatVersion *Max6Text `xml:"FrmtVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId,omitempty"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification71 `xml:"InitgPty,omitempty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification71 `xml:"RcptPty,omitempty"`
}

func (h *Header15) SetDownloadTransfer(value string) {
	h.DownloadTransfer = (*TrueFalseIndicator)(&value)
}

func (h *Header15) SetFormatVersion(value string) {
	h.FormatVersion = (*Max6Text)(&value)
}

func (h *Header15) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header15) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header15) AddInitiatingParty() *GenericIdentification71 {
	h.InitiatingParty = new(GenericIdentification71)
	return h.InitiatingParty
}

func (h *Header15) AddRecipientParty() *GenericIdentification71 {
	h.RecipientParty = new(GenericIdentification71)
	return h.RecipientParty
}

// Set of characteristics related to the reject of a transaction.
type Header16 struct {

	// Version of the terminal management protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId,omitempty"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification72 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification72 `xml:"RcptPty,omitempty"`
}

func (h *Header16) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header16) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header16) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header16) AddInitiatingParty() *GenericIdentification72 {
	h.InitiatingParty = new(GenericIdentification72)
	return h.InitiatingParty
}

func (h *Header16) AddRecipientParty() *GenericIdentification72 {
	h.RecipientParty = new(GenericIdentification72)
	return h.RecipientParty
}

// Set of characteristics related to the protocol.
type Header17 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction6Code `xml:"MsgFctn"`

	// Version of the acquirer to issuer protocol specifications
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Number of retransmission of the message. Incremented by one for each retransmission.
	ReTransmissionCounter *Max3NumericText `xml:"ReTrnsmssnCntr,omitempty"`

	// Date and time at which the message was sent.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification73 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification73 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability3 `xml:"Tracblt,omitempty"`
}

func (h *Header17) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction6Code)(&value)
}

func (h *Header17) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header17) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header17) SetReTransmissionCounter(value string) {
	h.ReTransmissionCounter = (*Max3NumericText)(&value)
}

func (h *Header17) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header17) AddInitiatingParty() *GenericIdentification73 {
	h.InitiatingParty = new(GenericIdentification73)
	return h.InitiatingParty
}

func (h *Header17) AddRecipientParty() *GenericIdentification73 {
	h.RecipientParty = new(GenericIdentification73)
	return h.RecipientParty
}

func (h *Header17) AddTraceability() *Traceability3 {
	newValue := new(Traceability3)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the protocol.
type Header18 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction6Code `xml:"MsgFctn"`

	// Identifies the type of process related to the message which has to be reversed.
	OriginalMessageFunction *MessageFunction6Code `xml:"OrgnlMsgFctn"`

	// Version of the acquirer to issuer protocol specifications
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Number of retransmission of the message. Incremented by 1 for each retransmission.
	ReTransmissionCounter *Max3NumericText `xml:"ReTrnsmssnCntr,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification73 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification73 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability3 `xml:"Tracblt,omitempty"`
}

func (h *Header18) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction6Code)(&value)
}

func (h *Header18) SetOriginalMessageFunction(value string) {
	h.OriginalMessageFunction = (*MessageFunction6Code)(&value)
}

func (h *Header18) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header18) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header18) SetReTransmissionCounter(value string) {
	h.ReTransmissionCounter = (*Max3NumericText)(&value)
}

func (h *Header18) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header18) AddInitiatingParty() *GenericIdentification73 {
	h.InitiatingParty = new(GenericIdentification73)
	return h.InitiatingParty
}

func (h *Header18) AddRecipientParty() *GenericIdentification73 {
	h.RecipientParty = new(GenericIdentification73)
	return h.RecipientParty
}

func (h *Header18) AddTraceability() *Traceability3 {
	newValue := new(Traceability3)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the protocol.
type Header19 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction6Code `xml:"MsgFctn"`

	// Identifies the type of process related to the message which has to be reversed.
	OriginalMessageFunction *MessageFunction6Code `xml:"OrgnlMsgFctn,omitempty"`

	// Version of the acquirer to issuer protocol specifications
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Number of retransmission of the message. Incremented by 1 for each retransmission.
	ReTransmissionCounter *Max3NumericText `xml:"ReTrnsmssnCntr,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification73 `xml:"InitgPty,omitempty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification73 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability3 `xml:"Tracblt,omitempty"`
}

func (h *Header19) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction6Code)(&value)
}

func (h *Header19) SetOriginalMessageFunction(value string) {
	h.OriginalMessageFunction = (*MessageFunction6Code)(&value)
}

func (h *Header19) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header19) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header19) SetReTransmissionCounter(value string) {
	h.ReTransmissionCounter = (*Max3NumericText)(&value)
}

func (h *Header19) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header19) AddInitiatingParty() *GenericIdentification73 {
	h.InitiatingParty = new(GenericIdentification73)
	return h.InitiatingParty
}

func (h *Header19) AddRecipientParty() *GenericIdentification73 {
	h.RecipientParty = new(GenericIdentification73)
	return h.RecipientParty
}

func (h *Header19) AddTraceability() *Traceability3 {
	newValue := new(Traceability3)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the protocol.
type Header2 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction1Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Number of retransmissions of the message.
	ReTransmissionCounter *Max3NumericText `xml:"ReTrnsmssnCntr,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification32 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification32 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`
}

func (h *Header2) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction1Code)(&value)
}

func (h *Header2) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header2) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header2) SetReTransmissionCounter(value string) {
	h.ReTransmissionCounter = (*Max3NumericText)(&value)
}

func (h *Header2) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header2) AddInitiatingParty() *GenericIdentification32 {
	h.InitiatingParty = new(GenericIdentification32)
	return h.InitiatingParty
}

func (h *Header2) AddRecipientParty() *GenericIdentification32 {
	h.RecipientParty = new(GenericIdentification32)
	return h.RecipientParty
}

func (h *Header2) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
type Header20 struct {

	// Identifies the type of process related to the message.
	MessageFunction *ATMMessageFunction1 `xml:"MsgFctn"`

	// Version of the ATM protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *Max35Text `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *Max35Text `xml:"RcptPty,omitempty"`

	// State of the sender of the message inside the process flow.
	ProcessState *Max35Text `xml:"PrcStat,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability4 `xml:"Tracblt,omitempty"`
}

func (h *Header20) AddMessageFunction() *ATMMessageFunction1 {
	h.MessageFunction = new(ATMMessageFunction1)
	return h.MessageFunction
}

func (h *Header20) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header20) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header20) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header20) SetInitiatingParty(value string) {
	h.InitiatingParty = (*Max35Text)(&value)
}

func (h *Header20) SetRecipientParty(value string) {
	h.RecipientParty = (*Max35Text)(&value)
}

func (h *Header20) SetProcessState(value string) {
	h.ProcessState = (*Max35Text)(&value)
}

func (h *Header20) AddTraceability() *Traceability4 {
	newValue := new(Traceability4)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
type Header21 struct {

	// Identifies the type of process related to the message.
	MessageFunction *ATMMessageFunction1 `xml:"MsgFctn"`

	// Version of the ATM protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	//  Retransmission counter of this advice, 0 for the first transmission.
	ReTransmissionCounter *Number `xml:"ReTrnsmssnCntr,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *Max35Text `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *Max35Text `xml:"RcptPty,omitempty"`

	// State of the sender of the message inside the process flow.
	ProcessState *Max35Text `xml:"PrcStat,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability4 `xml:"Tracblt,omitempty"`
}

func (h *Header21) AddMessageFunction() *ATMMessageFunction1 {
	h.MessageFunction = new(ATMMessageFunction1)
	return h.MessageFunction
}

func (h *Header21) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header21) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header21) SetReTransmissionCounter(value string) {
	h.ReTransmissionCounter = (*Number)(&value)
}

func (h *Header21) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header21) SetInitiatingParty(value string) {
	h.InitiatingParty = (*Max35Text)(&value)
}

func (h *Header21) SetRecipientParty(value string) {
	h.RecipientParty = (*Max35Text)(&value)
}

func (h *Header21) SetProcessState(value string) {
	h.ProcessState = (*Max35Text)(&value)
}

func (h *Header21) AddTraceability() *Traceability4 {
	newValue := new(Traceability4)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
type Header22 struct {

	// Identifies the type of process related to the message.
	MessageFunction *ATMMessageFunction1 `xml:"MsgFctn"`

	// Version of the ATM protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *Max35Text `xml:"InitgPty,omitempty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *Max35Text `xml:"RcptPty,omitempty"`

	// State of the sender of the message inside the process flow.
	ProcessState *Max35Text `xml:"PrcStat,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability4 `xml:"Tracblt,omitempty"`
}

func (h *Header22) AddMessageFunction() *ATMMessageFunction1 {
	h.MessageFunction = new(ATMMessageFunction1)
	return h.MessageFunction
}

func (h *Header22) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header22) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header22) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header22) SetInitiatingParty(value string) {
	h.InitiatingParty = (*Max35Text)(&value)
}

func (h *Header22) SetRecipientParty(value string) {
	h.RecipientParty = (*Max35Text)(&value)
}

func (h *Header22) SetProcessState(value string) {
	h.ProcessState = (*Max35Text)(&value)
}

func (h *Header22) AddTraceability() *Traceability4 {
	newValue := new(Traceability4)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the protocol.
type Header24 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction10Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Number `xml:"XchgId"`

	// Number of retransmissions of the message.
	ReTransmissionCounter *Max3NumericText `xml:"ReTrnsmssnCntr,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification53 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification94 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`
}

func (h *Header24) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction10Code)(&value)
}

func (h *Header24) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header24) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Number)(&value)
}

func (h *Header24) SetReTransmissionCounter(value string) {
	h.ReTransmissionCounter = (*Max3NumericText)(&value)
}

func (h *Header24) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header24) AddInitiatingParty() *GenericIdentification53 {
	h.InitiatingParty = new(GenericIdentification53)
	return h.InitiatingParty
}

func (h *Header24) AddRecipientParty() *GenericIdentification94 {
	h.RecipientParty = new(GenericIdentification94)
	return h.RecipientParty
}

func (h *Header24) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the transfer of transactions.
type Header25 struct {

	// Indicates if the file transfer is a download or an upload.
	DownloadTransfer *TrueFalseIndicator `xml:"DwnldTrf"`

	// Version of file format.
	FormatVersion *Max6Text `xml:"FrmtVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Number `xml:"XchgId"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification53 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification94 `xml:"RcptPty,omitempty"`
}

func (h *Header25) SetDownloadTransfer(value string) {
	h.DownloadTransfer = (*TrueFalseIndicator)(&value)
}

func (h *Header25) SetFormatVersion(value string) {
	h.FormatVersion = (*Max6Text)(&value)
}

func (h *Header25) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Number)(&value)
}

func (h *Header25) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header25) AddInitiatingParty() *GenericIdentification53 {
	h.InitiatingParty = new(GenericIdentification53)
	return h.InitiatingParty
}

func (h *Header25) AddRecipientParty() *GenericIdentification94 {
	h.RecipientParty = new(GenericIdentification94)
	return h.RecipientParty
}

// Set of characteristics related to the protocol after a rejection.
type Header26 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction9Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Number `xml:"XchgId,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification53 `xml:"InitgPty,omitempty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification94 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`
}

func (h *Header26) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction9Code)(&value)
}

func (h *Header26) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header26) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Number)(&value)
}

func (h *Header26) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header26) AddInitiatingParty() *GenericIdentification53 {
	h.InitiatingParty = new(GenericIdentification53)
	return h.InitiatingParty
}

func (h *Header26) AddRecipientParty() *GenericIdentification94 {
	h.RecipientParty = new(GenericIdentification94)
	return h.RecipientParty
}

func (h *Header26) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the transfer of transactions.
type Header27 struct {

	// Indicates if the file transfer is a download or an upload.
	DownloadTransfer *TrueFalseIndicator `xml:"DwnldTrf"`

	// Version of file format.
	FormatVersion *Max6Text `xml:"FrmtVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Number `xml:"XchgId"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification71 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification92 `xml:"RcptPty,omitempty"`
}

func (h *Header27) SetDownloadTransfer(value string) {
	h.DownloadTransfer = (*TrueFalseIndicator)(&value)
}

func (h *Header27) SetFormatVersion(value string) {
	h.FormatVersion = (*Max6Text)(&value)
}

func (h *Header27) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Number)(&value)
}

func (h *Header27) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header27) AddInitiatingParty() *GenericIdentification71 {
	h.InitiatingParty = new(GenericIdentification71)
	return h.InitiatingParty
}

func (h *Header27) AddRecipientParty() *GenericIdentification92 {
	h.RecipientParty = new(GenericIdentification92)
	return h.RecipientParty
}

// Set of characteristics related to the reject of a transaction.
type Header28 struct {

	// Indicates if the file transfer is a download or an upload.
	DownloadTransfer *TrueFalseIndicator `xml:"DwnldTrf"`

	// Version of file format.
	FormatVersion *Max6Text `xml:"FrmtVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Number `xml:"XchgId,omitempty"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification71 `xml:"InitgPty,omitempty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification92 `xml:"RcptPty,omitempty"`
}

func (h *Header28) SetDownloadTransfer(value string) {
	h.DownloadTransfer = (*TrueFalseIndicator)(&value)
}

func (h *Header28) SetFormatVersion(value string) {
	h.FormatVersion = (*Max6Text)(&value)
}

func (h *Header28) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Number)(&value)
}

func (h *Header28) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header28) AddInitiatingParty() *GenericIdentification71 {
	h.InitiatingParty = new(GenericIdentification71)
	return h.InitiatingParty
}

func (h *Header28) AddRecipientParty() *GenericIdentification92 {
	h.RecipientParty = new(GenericIdentification92)
	return h.RecipientParty
}

// Set of characteristics related to the reject of a transaction.
type Header29 struct {

	// Version of the terminal management protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Number `xml:"XchgId,omitempty"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification72 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification93 `xml:"RcptPty,omitempty"`
}

func (h *Header29) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header29) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Number)(&value)
}

func (h *Header29) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header29) AddInitiatingParty() *GenericIdentification72 {
	h.InitiatingParty = new(GenericIdentification72)
	return h.InitiatingParty
}

func (h *Header29) AddRecipientParty() *GenericIdentification93 {
	h.RecipientParty = new(GenericIdentification93)
	return h.RecipientParty
}

// Set of characteristics related to the transfer of transactions.
type Header3 struct {

	// Indicates if the file transfer is a download or an upload.
	DownloadTransfer *TrueFalseIndicator `xml:"DwnldTrf"`

	// Version of file format.
	FormatVersion *Max6Text `xml:"FrmtVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification32 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification32 `xml:"RcptPty,omitempty"`
}

func (h *Header3) SetDownloadTransfer(value string) {
	h.DownloadTransfer = (*TrueFalseIndicator)(&value)
}

func (h *Header3) SetFormatVersion(value string) {
	h.FormatVersion = (*Max6Text)(&value)
}

func (h *Header3) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header3) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header3) AddInitiatingParty() *GenericIdentification32 {
	h.InitiatingParty = new(GenericIdentification32)
	return h.InitiatingParty
}

func (h *Header3) AddRecipientParty() *GenericIdentification32 {
	h.RecipientParty = new(GenericIdentification32)
	return h.RecipientParty
}

// Set of characteristics related to the protocol.
type Header30 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction10Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Number `xml:"XchgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification53 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification94 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`
}

func (h *Header30) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction10Code)(&value)
}

func (h *Header30) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header30) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Number)(&value)
}

func (h *Header30) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header30) AddInitiatingParty() *GenericIdentification53 {
	h.InitiatingParty = new(GenericIdentification53)
	return h.InitiatingParty
}

func (h *Header30) AddRecipientParty() *GenericIdentification94 {
	h.RecipientParty = new(GenericIdentification94)
	return h.RecipientParty
}

func (h *Header30) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
type Header31 struct {

	// Identifies the type of process related to the message.
	MessageFunction *ATMMessageFunction2 `xml:"MsgFctn"`

	// Version of the ATM protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *Max35Text `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *Max35Text `xml:"RcptPty,omitempty"`

	// State of the sender of the message inside the process flow.
	ProcessState *Max35Text `xml:"PrcStat,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability4 `xml:"Tracblt,omitempty"`
}

func (h *Header31) AddMessageFunction() *ATMMessageFunction2 {
	h.MessageFunction = new(ATMMessageFunction2)
	return h.MessageFunction
}

func (h *Header31) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header31) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header31) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header31) SetInitiatingParty(value string) {
	h.InitiatingParty = (*Max35Text)(&value)
}

func (h *Header31) SetRecipientParty(value string) {
	h.RecipientParty = (*Max35Text)(&value)
}

func (h *Header31) SetProcessState(value string) {
	h.ProcessState = (*Max35Text)(&value)
}

func (h *Header31) AddTraceability() *Traceability4 {
	newValue := new(Traceability4)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
type Header32 struct {

	// Identifies the type of process related to the message.
	MessageFunction *ATMMessageFunction2 `xml:"MsgFctn"`

	// Version of the ATM protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	//  Retransmission counter of this advice, 0 for the first transmission.
	ReTransmissionCounter *Number `xml:"ReTrnsmssnCntr,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *Max35Text `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *Max35Text `xml:"RcptPty,omitempty"`

	// State of the sender of the message inside the process flow.
	ProcessState *Max35Text `xml:"PrcStat,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability4 `xml:"Tracblt,omitempty"`
}

func (h *Header32) AddMessageFunction() *ATMMessageFunction2 {
	h.MessageFunction = new(ATMMessageFunction2)
	return h.MessageFunction
}

func (h *Header32) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header32) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header32) SetReTransmissionCounter(value string) {
	h.ReTransmissionCounter = (*Number)(&value)
}

func (h *Header32) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header32) SetInitiatingParty(value string) {
	h.InitiatingParty = (*Max35Text)(&value)
}

func (h *Header32) SetRecipientParty(value string) {
	h.RecipientParty = (*Max35Text)(&value)
}

func (h *Header32) SetProcessState(value string) {
	h.ProcessState = (*Max35Text)(&value)
}

func (h *Header32) AddTraceability() *Traceability4 {
	newValue := new(Traceability4)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
type Header33 struct {

	// Identifies the type of process related to the message.
	MessageFunction *ATMMessageFunction2 `xml:"MsgFctn"`

	// Version of the ATM protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *Max35Text `xml:"InitgPty,omitempty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *Max35Text `xml:"RcptPty,omitempty"`

	// State of the sender of the message inside the process flow.
	ProcessState *Max35Text `xml:"PrcStat,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability4 `xml:"Tracblt,omitempty"`
}

func (h *Header33) AddMessageFunction() *ATMMessageFunction2 {
	h.MessageFunction = new(ATMMessageFunction2)
	return h.MessageFunction
}

func (h *Header33) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header33) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header33) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header33) SetInitiatingParty(value string) {
	h.InitiatingParty = (*Max35Text)(&value)
}

func (h *Header33) SetRecipientParty(value string) {
	h.RecipientParty = (*Max35Text)(&value)
}

func (h *Header33) SetProcessState(value string) {
	h.ProcessState = (*Max35Text)(&value)
}

func (h *Header33) AddTraceability() *Traceability4 {
	newValue := new(Traceability4)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the transfer of transactions.
type Header4 struct {

	// Indicates if the file transfer is a download or an upload.
	DownloadTransfer *TrueFalseIndicator `xml:"DwnldTrf"`

	// Version of file format.
	FormatVersion *Max6Text `xml:"FrmtVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification35 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification35 `xml:"RcptPty,omitempty"`
}

func (h *Header4) SetDownloadTransfer(value string) {
	h.DownloadTransfer = (*TrueFalseIndicator)(&value)
}

func (h *Header4) SetFormatVersion(value string) {
	h.FormatVersion = (*Max6Text)(&value)
}

func (h *Header4) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header4) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header4) AddInitiatingParty() *GenericIdentification35 {
	h.InitiatingParty = new(GenericIdentification35)
	return h.InitiatingParty
}

func (h *Header4) AddRecipientParty() *GenericIdentification35 {
	h.RecipientParty = new(GenericIdentification35)
	return h.RecipientParty
}

// Set of characteristics related to the protocol after a rejection.
type Header5 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction1Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification32 `xml:"InitgPty,omitempty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification32 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`
}

func (h *Header5) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction1Code)(&value)
}

func (h *Header5) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header5) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header5) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header5) AddInitiatingParty() *GenericIdentification32 {
	h.InitiatingParty = new(GenericIdentification32)
	return h.InitiatingParty
}

func (h *Header5) AddRecipientParty() *GenericIdentification32 {
	h.RecipientParty = new(GenericIdentification32)
	return h.RecipientParty
}

func (h *Header5) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the reject of a transaction.
type Header6 struct {

	// Indicates if the file transfer is a download or an upload.
	DownloadTransfer *TrueFalseIndicator `xml:"DwnldTrf"`

	// Version of file format.
	FormatVersion *Max6Text `xml:"FrmtVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId,omitempty"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification35 `xml:"InitgPty,omitempty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification35 `xml:"RcptPty,omitempty"`
}

func (h *Header6) SetDownloadTransfer(value string) {
	h.DownloadTransfer = (*TrueFalseIndicator)(&value)
}

func (h *Header6) SetFormatVersion(value string) {
	h.FormatVersion = (*Max6Text)(&value)
}

func (h *Header6) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header6) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header6) AddInitiatingParty() *GenericIdentification35 {
	h.InitiatingParty = new(GenericIdentification35)
	return h.InitiatingParty
}

func (h *Header6) AddRecipientParty() *GenericIdentification35 {
	h.RecipientParty = new(GenericIdentification35)
	return h.RecipientParty
}

// Set of characteristics related to the protocol.
type Header7 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction4Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification32 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification32 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`
}

func (h *Header7) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction4Code)(&value)
}

func (h *Header7) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header7) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header7) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header7) AddInitiatingParty() *GenericIdentification32 {
	h.InitiatingParty = new(GenericIdentification32)
	return h.InitiatingParty
}

func (h *Header7) AddRecipientParty() *GenericIdentification32 {
	h.RecipientParty = new(GenericIdentification32)
	return h.RecipientParty
}

func (h *Header7) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the protocol.
type Header8 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction4Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Number of retransmissions of the message.
	ReTransmissionCounter *Max3NumericText `xml:"ReTrnsmssnCntr,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification32 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification32 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`
}

func (h *Header8) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction4Code)(&value)
}

func (h *Header8) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header8) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header8) SetReTransmissionCounter(value string) {
	h.ReTransmissionCounter = (*Max3NumericText)(&value)
}

func (h *Header8) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header8) AddInitiatingParty() *GenericIdentification32 {
	h.InitiatingParty = new(GenericIdentification32)
	return h.InitiatingParty
}

func (h *Header8) AddRecipientParty() *GenericIdentification32 {
	h.RecipientParty = new(GenericIdentification32)
	return h.RecipientParty
}

func (h *Header8) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}

// Set of characteristics related to the protocol after a rejection.
type Header9 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction4Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification32 `xml:"InitgPty,omitempty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification32 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`
}

func (h *Header9) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction4Code)(&value)
}

func (h *Header9) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header9) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header9) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header9) AddInitiatingParty() *GenericIdentification32 {
	h.InitiatingParty = new(GenericIdentification32)
	return h.InitiatingParty
}

func (h *Header9) AddRecipientParty() *GenericIdentification32 {
	h.RecipientParty = new(GenericIdentification32)
	return h.RecipientParty
}

func (h *Header9) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}
