package iso20022

// Payment context in which the transaction is performed.
type PaymentContext1 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel1Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	AttendantLanguage *ISO2ALanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *TrueFalseIndicator `xml:"FllbckInd,omitempty"`
}

func (p *PaymentContext1) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext1) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext1) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext1) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext1) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext1) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel1Code)(&value)
}

func (p *PaymentContext1) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext1) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*ISO2ALanguageCode)(&value)
}

func (p *PaymentContext1) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext1) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*TrueFalseIndicator)(&value)
}

// Payment context in which the transaction is performed.
type PaymentContext10 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	// Reference: ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AttendantLanguage *LanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`
}

func (p *PaymentContext10) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext10) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext10) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext10) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext10) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext10) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext10) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext10) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*LanguageCode)(&value)
}

func (p *PaymentContext10) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext10) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

// Payment context in which the transaction is performed.
type PaymentContext11 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`
}

func (p *PaymentContext11) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext11) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext11) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext11) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext11) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext11) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext11) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext11) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

// Payment context in which the transaction is performed.
type PaymentContext12 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	// Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AttendantLanguage *LanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd,omitempty"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Payment options the card acceptor can support.
	SupportedOption []*SupportedPaymentOption1Code `xml:"SpprtdOptn,omitempty"`
}

func (p *PaymentContext12) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext12) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext12) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext12) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext12) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext12) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext12) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext12) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*LanguageCode)(&value)
}

func (p *PaymentContext12) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext12) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (p *PaymentContext12) AddSupportedOption(value string) {
	p.SupportedOption = append(p.SupportedOption, (*SupportedPaymentOption1Code)(&value))
}

// Payment context in which the transaction is performed.
type PaymentContext13 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	// Reference: ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AttendantLanguage *LanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd,omitempty"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`
}

func (p *PaymentContext13) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext13) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext13) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext13) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext13) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext13) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext13) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext13) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*LanguageCode)(&value)
}

func (p *PaymentContext13) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext13) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

// Payment context in which the transaction is performed.
type PaymentContext14 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	// Reference: ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AttendantLanguage *LanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading5Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Payment options the card acceptor can support.
	SupportedOption []*SupportedPaymentOption1Code `xml:"SpprtdOptn,omitempty"`
}

func (p *PaymentContext14) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext14) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext14) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext14) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext14) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext14) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext14) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*LanguageCode)(&value)
}

func (p *PaymentContext14) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading5Code)(&value)
}

func (p *PaymentContext14) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (p *PaymentContext14) AddSupportedOption(value string) {
	p.SupportedOption = append(p.SupportedOption, (*SupportedPaymentOption1Code)(&value))
}

// Payment context in which the transaction is performed.
type PaymentContext15 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading5Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Payment options the card acceptor can support.
	SupportedOption []*SupportedPaymentOption1Code `xml:"SpprtdOptn,omitempty"`
}

func (p *PaymentContext15) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext15) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext15) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext15) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext15) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext15) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext15) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading5Code)(&value)
}

func (p *PaymentContext15) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (p *PaymentContext15) AddSupportedOption(value string) {
	p.SupportedOption = append(p.SupportedOption, (*SupportedPaymentOption1Code)(&value))
}

// Payment context in which the transaction is performed.
type PaymentContext16 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// Human attendance at the POI location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	// Reference: ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AttendantLanguage *LanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading5Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`
}

func (p *PaymentContext16) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext16) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext16) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext16) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext16) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext16) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext16) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*LanguageCode)(&value)
}

func (p *PaymentContext16) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading5Code)(&value)
}

func (p *PaymentContext16) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

// Payment context in which the transaction is performed.
type PaymentContext17 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading5Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`
}

func (p *PaymentContext17) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext17) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext17) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext17) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext17) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext17) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext17) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading5Code)(&value)
}

func (p *PaymentContext17) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

// Payment context in which the transaction is performed.
type PaymentContext18 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	// Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AttendantLanguage *LanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading5Code `xml:"CardDataNtryMd,omitempty"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Payment options the card acceptor can support.
	SupportedOption []*SupportedPaymentOption1Code `xml:"SpprtdOptn,omitempty"`
}

func (p *PaymentContext18) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext18) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext18) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext18) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext18) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext18) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext18) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext18) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*LanguageCode)(&value)
}

func (p *PaymentContext18) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading5Code)(&value)
}

func (p *PaymentContext18) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (p *PaymentContext18) AddSupportedOption(value string) {
	p.SupportedOption = append(p.SupportedOption, (*SupportedPaymentOption1Code)(&value))
}

// Payment context in which the transaction is performed.
type PaymentContext19 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	// Reference: ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AttendantLanguage *LanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading5Code `xml:"CardDataNtryMd,omitempty"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`
}

func (p *PaymentContext19) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext19) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext19) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext19) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext19) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext19) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext19) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext19) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*LanguageCode)(&value)
}

func (p *PaymentContext19) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading5Code)(&value)
}

func (p *PaymentContext19) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

// Payment context in which the transaction is performed.
type PaymentContext2 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel1Code `xml:"TxChanl,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *TrueFalseIndicator `xml:"FllbckInd,omitempty"`
}

func (p *PaymentContext2) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext2) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext2) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext2) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext2) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext2) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel1Code)(&value)
}

func (p *PaymentContext2) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext2) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*TrueFalseIndicator)(&value)
}

// Payment context in which the transaction is performed.
type PaymentContext3 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel1Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	AttendantLanguage *ISO2ALanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *TrueFalseIndicator `xml:"FllbckInd,omitempty"`

	// Method used to authenticate a cardholder.
	AuthenticationMethod *CardholderAuthentication2 `xml:"AuthntcnMtd,omitempty"`
}

func (p *PaymentContext3) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext3) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext3) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext3) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext3) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext3) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel1Code)(&value)
}

func (p *PaymentContext3) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext3) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*ISO2ALanguageCode)(&value)
}

func (p *PaymentContext3) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext3) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext3) AddAuthenticationMethod() *CardholderAuthentication2 {
	p.AuthenticationMethod = new(CardholderAuthentication2)
	return p.AuthenticationMethod
}

// Payment context in which the transaction is performed.
type PaymentContext4 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel1Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	AttendantLanguage *ISO2ALanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd,omitempty"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *TrueFalseIndicator `xml:"FllbckInd,omitempty"`
}

func (p *PaymentContext4) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext4) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext4) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext4) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext4) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext4) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel1Code)(&value)
}

func (p *PaymentContext4) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext4) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*ISO2ALanguageCode)(&value)
}

func (p *PaymentContext4) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext4) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*TrueFalseIndicator)(&value)
}

// Payment context in which the transaction is performed.
type PaymentContext5 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel1Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	AttendantLanguage *ISO2ALanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *TrueFalseIndicator `xml:"FllbckInd,omitempty"`

	// Payment options the card acceptor can support.
	SupportedOption []*SupportedPaymentOption1Code `xml:"SpprtdOptn,omitempty"`
}

func (p *PaymentContext5) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext5) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext5) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext5) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext5) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel1Code)(&value)
}

func (p *PaymentContext5) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext5) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*ISO2ALanguageCode)(&value)
}

func (p *PaymentContext5) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext5) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext5) AddSupportedOption(value string) {
	p.SupportedOption = append(p.SupportedOption, (*SupportedPaymentOption1Code)(&value))
}

// Payment context in which the transaction is performed.
type PaymentContext6 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel1Code `xml:"TxChanl,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *TrueFalseIndicator `xml:"FllbckInd,omitempty"`

	// Payment options the card acceptor can support.
	SupportedOption []*SupportedPaymentOption1Code `xml:"SpprtdOptn,omitempty"`
}

func (p *PaymentContext6) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext6) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext6) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext6) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext6) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext6) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel1Code)(&value)
}

func (p *PaymentContext6) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext6) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext6) AddSupportedOption(value string) {
	p.SupportedOption = append(p.SupportedOption, (*SupportedPaymentOption1Code)(&value))
}

// Payment context in which the transaction is performed.
type PaymentContext7 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel1Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	AttendantLanguage *ISO2ALanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd,omitempty"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *TrueFalseIndicator `xml:"FllbckInd,omitempty"`

	// Payment options the card acceptor can support.
	SupportedOption []*SupportedPaymentOption1Code `xml:"SpprtdOptn,omitempty"`
}

func (p *PaymentContext7) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext7) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext7) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext7) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext7) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext7) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel1Code)(&value)
}

func (p *PaymentContext7) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext7) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*ISO2ALanguageCode)(&value)
}

func (p *PaymentContext7) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext7) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext7) AddSupportedOption(value string) {
	p.SupportedOption = append(p.SupportedOption, (*SupportedPaymentOption1Code)(&value))
}

// Payment context in which the transaction is performed.
type PaymentContext8 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Indicates whether a message can be sent or not on an attendant display (attendant display present or not).
	AttendantMessageCapable *TrueFalseIndicator `xml:"AttndntMsgCpbl,omitempty"`

	// Language used to display messages to the attendant.
	// Reference: ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AttendantLanguage *LanguageCode `xml:"AttndntLang,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Payment options the card acceptor can support.
	SupportedOption []*SupportedPaymentOption1Code `xml:"SpprtdOptn,omitempty"`
}

func (p *PaymentContext8) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext8) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext8) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext8) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext8) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext8) SetAttendantMessageCapable(value string) {
	p.AttendantMessageCapable = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext8) SetAttendantLanguage(value string) {
	p.AttendantLanguage = (*LanguageCode)(&value)
}

func (p *PaymentContext8) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext8) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (p *PaymentContext8) AddSupportedOption(value string) {
	p.SupportedOption = append(p.SupportedOption, (*SupportedPaymentOption1Code)(&value))
}

// Payment context in which the transaction is performed.
type PaymentContext9 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI (Point Of Interaction) location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Payment options the card acceptor can support.
	SupportedOption []*SupportedPaymentOption1Code `xml:"SpprtdOptn,omitempty"`
}

func (p *PaymentContext9) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext9) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext9) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext9) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext9) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext9) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext9) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentContext9) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (p *PaymentContext9) AddSupportedOption(value string) {
	p.SupportedOption = append(p.SupportedOption, (*SupportedPaymentOption1Code)(&value))
}
