package iso20022

// Payment card performing the transaction.
type PaymentCard1 struct {

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType2 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData1 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Exact3NumericText `xml:"CardCtryCd,omitempty"`

	// Defines a category of cards related to the acceptance processing rules defined by the acquirer.
	CardProductProfile *Exact4NumericText `xml:"CardPdctPrfl,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard1) AddProtectedCardData() *ContentInformationType2 {
	p.ProtectedCardData = new(ContentInformationType2)
	return p.ProtectedCardData
}

func (p *PaymentCard1) AddPlainCardData() *PlainCardData1 {
	p.PlainCardData = new(PlainCardData1)
	return p.PlainCardData
}

func (p *PaymentCard1) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Exact3NumericText)(&value)
}

func (p *PaymentCard1) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Exact4NumericText)(&value)
}

func (p *PaymentCard1) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard1) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}

// Payment card performing the transaction.
type PaymentCard10 struct {

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData8 `xml:"PlainCardData,omitempty"`

	// Masked PAN to be printed on payment receipts or displayed to the cardholder. Masked digits may be absent or replaced by another character as '*'.
	MaskedPAN *Max30Text `xml:"MskdPAN,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Type of card product.
	CardProductType *CardProductType1Code `xml:"CardPdctTp,omitempty"`
}

func (p *PaymentCard10) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard10) AddPlainCardData() *PlainCardData8 {
	p.PlainCardData = new(PlainCardData8)
	return p.PlainCardData
}

func (p *PaymentCard10) SetMaskedPAN(value string) {
	p.MaskedPAN = (*Max30Text)(&value)
}

func (p *PaymentCard10) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard10) SetCardProductType(value string) {
	p.CardProductType = (*CardProductType1Code)(&value)
}

// Payment card performing the transaction.
type PaymentCard11 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData9 `xml:"PlainCardData,omitempty"`

	// Masked PAN to be printed the payment receipts or displayed to the cardholder. Masked digits may be absent or replaced by another character as '*'.
	MaskedPAN *Max30Text `xml:"MskdPAN,omitempty"`

	// Bank identifier number of the issuer for routing purpose.
	IssuerBIN *Max15NumericText `xml:"IssrBIN,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`

	// Defines a category of cards related to the acceptance processing rules defined by the acquirer.
	CardProductProfile *Max35Text `xml:"CardPdctPrfl,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Type of card product.
	CardProductType *CardProductType1Code `xml:"CardPdctTp,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard11) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard11) AddPlainCardData() *PlainCardData9 {
	p.PlainCardData = new(PlainCardData9)
	return p.PlainCardData
}

func (p *PaymentCard11) SetMaskedPAN(value string) {
	p.MaskedPAN = (*Max30Text)(&value)
}

func (p *PaymentCard11) SetIssuerBIN(value string) {
	p.IssuerBIN = (*Max15NumericText)(&value)
}

func (p *PaymentCard11) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard11) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

func (p *PaymentCard11) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Max35Text)(&value)
}

func (p *PaymentCard11) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard11) SetCardProductType(value string) {
	p.CardProductType = (*CardProductType1Code)(&value)
}

func (p *PaymentCard11) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}

// Payment card performing the transaction.
type PaymentCard12 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData10 `xml:"PlainCardData,omitempty"`

	// Bank identifier number of the issuer for routing purpose.
	IssuerBIN *Max15NumericText `xml:"IssrBIN,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard12) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard12) AddPlainCardData() *PlainCardData10 {
	p.PlainCardData = new(PlainCardData10)
	return p.PlainCardData
}

func (p *PaymentCard12) SetIssuerBIN(value string) {
	p.IssuerBIN = (*Max15NumericText)(&value)
}

func (p *PaymentCard12) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard12) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

func (p *PaymentCard12) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}

// Payment card performing the transaction.
type PaymentCard13 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData9 `xml:"PlainCardData,omitempty"`

	// Masked PAN to be printed the payment receipts or displayed to the cardholder. Masked digits may be absent or replaced by another character as '*'.
	MaskedPAN *Max30Text `xml:"MskdPAN,omitempty"`

	// Type of card product.
	CardProductType *CardProductType1Code `xml:"CardPdctTp,omitempty"`

	// Name of card product.
	CardProductName *Max35Text `xml:"CardPdctNm,omitempty"`
}

func (p *PaymentCard13) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard13) AddPlainCardData() *PlainCardData9 {
	p.PlainCardData = new(PlainCardData9)
	return p.PlainCardData
}

func (p *PaymentCard13) SetMaskedPAN(value string) {
	p.MaskedPAN = (*Max30Text)(&value)
}

func (p *PaymentCard13) SetCardProductType(value string) {
	p.CardProductType = (*CardProductType1Code)(&value)
}

func (p *PaymentCard13) SetCardProductName(value string) {
	p.CardProductName = (*Max35Text)(&value)
}

// Payment card performing the transaction.
type PaymentCard14 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData11 `xml:"PlainCardData,omitempty"`

	// Bank identifier number of the issuer for routing purpose.
	IssuerBIN *Max15NumericText `xml:"IssrBIN,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`
}

func (p *PaymentCard14) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard14) AddPlainCardData() *PlainCardData11 {
	p.PlainCardData = new(PlainCardData11)
	return p.PlainCardData
}

func (p *PaymentCard14) SetIssuerBIN(value string) {
	p.IssuerBIN = (*Max15NumericText)(&value)
}

func (p *PaymentCard14) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard14) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

// Payment card performing the transaction.
type PaymentCard15 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData12 `xml:"PlainCardData,omitempty"`

	// Masked PAN to be printed the payment receipts or displayed to the cardholder. Masked digits may be absent or replaced by another character as '*'.
	MaskedPAN *Max30Text `xml:"MskdPAN,omitempty"`

	// Type of card product.
	CardProductType *CardProductType1Code `xml:"CardPdctTp,omitempty"`

	// Name of card product.
	CardProductName *Max35Text `xml:"CardPdctNm,omitempty"`
}

func (p *PaymentCard15) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard15) AddPlainCardData() *PlainCardData12 {
	p.PlainCardData = new(PlainCardData12)
	return p.PlainCardData
}

func (p *PaymentCard15) SetMaskedPAN(value string) {
	p.MaskedPAN = (*Max30Text)(&value)
}

func (p *PaymentCard15) SetCardProductType(value string) {
	p.CardProductType = (*CardProductType1Code)(&value)
}

func (p *PaymentCard15) SetCardProductName(value string) {
	p.CardProductName = (*Max35Text)(&value)
}

// Card performing the withdrawal transaction.
type PaymentCard16 struct {

	// Entry mode used to obtain the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicates the occurrence of a fall-back on the card entry mode.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData13 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`
}

func (p *PaymentCard16) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentCard16) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (p *PaymentCard16) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard16) AddPlainCardData() *PlainCardData13 {
	p.PlainCardData = new(PlainCardData13)
	return p.PlainCardData
}

func (p *PaymentCard16) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard16) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

// Card performing the withdrawal transaction.
type PaymentCard17 struct {

	// Entry mode that used to obtain the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicate the occurrence of a fall-back on the card entry mode.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData14 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`

	// Balance of the captured card or epurse if available.
	RetainedCardBalance *CurrencyAndAmount `xml:"RtndCardBal,omitempty"`
}

func (p *PaymentCard17) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentCard17) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (p *PaymentCard17) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard17) AddPlainCardData() *PlainCardData14 {
	p.PlainCardData = new(PlainCardData14)
	return p.PlainCardData
}

func (p *PaymentCard17) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard17) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

func (p *PaymentCard17) SetRetainedCardBalance(value, currency string) {
	p.RetainedCardBalance = NewCurrencyAndAmount(value, currency)
}

// Electronic money product that provides the cardholder with a portable and specialised computer device, which typically contains a microprocessor.
type PaymentCard18 struct {

	// Type of card, for example, credit card.
	Type *CardType1Code `xml:"Tp"`

	// Number embossed on a card that links the card to the account owner and account servicer.
	Number *Max35Text `xml:"Nb"`

	// Party entitled by a card issuer to use a card.
	HolderName *Max35Text `xml:"HldrNm"`

	// Year and month the card is available for use.
	StartDate *ISOYearMonth `xml:"StartDt,omitempty"`

	// Year and month the card expires.
	ExpiryDate *ISOYearMonth `xml:"XpryDt"`

	// Party that issues a payment card, as expressed by a numeric identification of the card issuer according to ISO/IEC 7812-1.
	CardIssuerName *Max35Text `xml:"CardIssrNm,omitempty"`

	// Party that issues a payment card, as expressed by a numeric identification of the card issuer according to ISO/IEC 7812-1.
	CardIssuerIdentification *PartyIdentification70Choice `xml:"CardIssrId,omitempty"`

	// Security code written on, or in, the card.
	SecurityCode *Max35Text `xml:"SctyCd,omitempty"`

	// Number distinguishing two or more payment cards with the same account number.
	SequenceNumber *Max3Text `xml:"SeqNb,omitempty"`
}

func (p *PaymentCard18) SetType(value string) {
	p.Type = (*CardType1Code)(&value)
}

func (p *PaymentCard18) SetNumber(value string) {
	p.Number = (*Max35Text)(&value)
}

func (p *PaymentCard18) SetHolderName(value string) {
	p.HolderName = (*Max35Text)(&value)
}

func (p *PaymentCard18) SetStartDate(value string) {
	p.StartDate = (*ISOYearMonth)(&value)
}

func (p *PaymentCard18) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISOYearMonth)(&value)
}

func (p *PaymentCard18) SetCardIssuerName(value string) {
	p.CardIssuerName = (*Max35Text)(&value)
}

func (p *PaymentCard18) AddCardIssuerIdentification() *PartyIdentification70Choice {
	p.CardIssuerIdentification = new(PartyIdentification70Choice)
	return p.CardIssuerIdentification
}

func (p *PaymentCard18) SetSecurityCode(value string) {
	p.SecurityCode = (*Max35Text)(&value)
}

func (p *PaymentCard18) SetSequenceNumber(value string) {
	p.SequenceNumber = (*Max3Text)(&value)
}

// Payment card performing the transaction.
type PaymentCard19 struct {

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData8 `xml:"PlainCardData,omitempty"`

	// Unique reference to the card, used by both merchants and acquirers to link tokenized and non-tokenized transactions associated to the same underlying card.
	PaymentAccountReference *Max70Text `xml:"PmtAcctRef,omitempty"`

	// Masked PAN to be printed on payment receipts or displayed to the cardholder. Masked digits may be absent or replaced by another character as '*'.
	MaskedPAN *Max30Text `xml:"MskdPAN,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Type of card product.
	CardProductType *CardProductType1Code `xml:"CardPdctTp,omitempty"`
}

func (p *PaymentCard19) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard19) AddPlainCardData() *PlainCardData8 {
	p.PlainCardData = new(PlainCardData8)
	return p.PlainCardData
}

func (p *PaymentCard19) SetPaymentAccountReference(value string) {
	p.PaymentAccountReference = (*Max70Text)(&value)
}

func (p *PaymentCard19) SetMaskedPAN(value string) {
	p.MaskedPAN = (*Max30Text)(&value)
}

func (p *PaymentCard19) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard19) SetCardProductType(value string) {
	p.CardProductType = (*CardProductType1Code)(&value)
}

// Electronic money product that provides the cardholder with a portable and specialised computer device, which typically contains a microprocessor.
type PaymentCard2 struct {

	// Type of card, eg, credit card.
	Type *CardType1Code `xml:"Tp"`

	// Number embossed on a card that links the card to the account owner and account servicer.
	Number *Max35Text `xml:"Nb"`

	// Party entitled by a card issuer to use a card.
	HolderName *Max35Text `xml:"HldrNm"`

	// Year and month the card is available for use.
	StartDate *ISOYearMonth `xml:"StartDt,omitempty"`

	// Year and month the card expires.
	ExpiryDate *ISOYearMonth `xml:"XpryDt"`

	// Party that issues a payment card, as expressed by a numeric identification of the card issuer according to ISO/IEC 7812-1.
	CardIssuerName *Max35Text `xml:"CardIssrNm,omitempty"`

	// Party that issues a payment card, as expressed by a numeric identification of the card issuer according to ISO/IEC 7812-1.
	CardIssuerIdentification *PartyIdentification2Choice `xml:"CardIssrId,omitempty"`

	// Security code written on, or in, the card.
	SecurityCode *Max35Text `xml:"SctyCd,omitempty"`

	// Number distinguishing two or more payment cards with the same account number.
	SequenceNumber *Max3Text `xml:"SeqNb,omitempty"`
}

func (p *PaymentCard2) SetType(value string) {
	p.Type = (*CardType1Code)(&value)
}

func (p *PaymentCard2) SetNumber(value string) {
	p.Number = (*Max35Text)(&value)
}

func (p *PaymentCard2) SetHolderName(value string) {
	p.HolderName = (*Max35Text)(&value)
}

func (p *PaymentCard2) SetStartDate(value string) {
	p.StartDate = (*ISOYearMonth)(&value)
}

func (p *PaymentCard2) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISOYearMonth)(&value)
}

func (p *PaymentCard2) SetCardIssuerName(value string) {
	p.CardIssuerName = (*Max35Text)(&value)
}

func (p *PaymentCard2) AddCardIssuerIdentification() *PartyIdentification2Choice {
	p.CardIssuerIdentification = new(PartyIdentification2Choice)
	return p.CardIssuerIdentification
}

func (p *PaymentCard2) SetSecurityCode(value string) {
	p.SecurityCode = (*Max35Text)(&value)
}

func (p *PaymentCard2) SetSequenceNumber(value string) {
	p.SequenceNumber = (*Max3Text)(&value)
}

// Payment card performing the transaction.
type PaymentCard20 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData16 `xml:"PlainCardData,omitempty"`

	// Unique reference to the card, used by both merchants and acquirers to link tokenized and non-tokenized transactions associated to the same underlying card.
	PaymentAccountReference *Max70Text `xml:"PmtAcctRef,omitempty"`

	// Masked PAN to be printed the payment receipts or displayed to the cardholder. Masked digits may be absent or replaced by another character as '*'.
	MaskedPAN *Max30Text `xml:"MskdPAN,omitempty"`

	// Bank identifier number of the issuer for routing purpose.
	IssuerBIN *Max15NumericText `xml:"IssrBIN,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`

	// Defines a category of cards related to the acceptance processing rules defined by the acquirer.
	CardProductProfile *Max35Text `xml:"CardPdctPrfl,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Type of card product.
	CardProductType *CardProductType1Code `xml:"CardPdctTp,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard20) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard20) AddPlainCardData() *PlainCardData16 {
	p.PlainCardData = new(PlainCardData16)
	return p.PlainCardData
}

func (p *PaymentCard20) SetPaymentAccountReference(value string) {
	p.PaymentAccountReference = (*Max70Text)(&value)
}

func (p *PaymentCard20) SetMaskedPAN(value string) {
	p.MaskedPAN = (*Max30Text)(&value)
}

func (p *PaymentCard20) SetIssuerBIN(value string) {
	p.IssuerBIN = (*Max15NumericText)(&value)
}

func (p *PaymentCard20) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard20) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

func (p *PaymentCard20) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Max35Text)(&value)
}

func (p *PaymentCard20) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard20) SetCardProductType(value string) {
	p.CardProductType = (*CardProductType1Code)(&value)
}

func (p *PaymentCard20) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}

// Payment card performing the transaction.
type PaymentCard21 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData15 `xml:"PlainCardData,omitempty"`

	// Unique reference to the card, used by both merchants and acquirers to link tokenized and non-tokenized transactions associated to the same underlying card.
	PaymentAccountReference *Max70Text `xml:"PmtAcctRef,omitempty"`

	// Bank identifier number of the issuer for routing purpose.
	IssuerBIN *Max15NumericText `xml:"IssrBIN,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`

	// Defines a category of cards related to the acceptance processing rules defined by the acquirer.
	CardProductProfile *Max35Text `xml:"CardPdctPrfl,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// True if the card may be used abroad.
	InternationalCard *TrueFalseIndicator `xml:"IntrnlCard,omitempty"`

	// Product that can be purchased with the card.
	AllowedProduct []*Max70Text `xml:"AllwdPdct,omitempty"`

	// Options to the service provided by the card.
	ServiceOption *Max35Text `xml:"SvcOptn,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard21) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard21) AddPlainCardData() *PlainCardData15 {
	p.PlainCardData = new(PlainCardData15)
	return p.PlainCardData
}

func (p *PaymentCard21) SetPaymentAccountReference(value string) {
	p.PaymentAccountReference = (*Max70Text)(&value)
}

func (p *PaymentCard21) SetIssuerBIN(value string) {
	p.IssuerBIN = (*Max15NumericText)(&value)
}

func (p *PaymentCard21) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard21) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

func (p *PaymentCard21) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Max35Text)(&value)
}

func (p *PaymentCard21) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard21) SetInternationalCard(value string) {
	p.InternationalCard = (*TrueFalseIndicator)(&value)
}

func (p *PaymentCard21) AddAllowedProduct(value string) {
	p.AllowedProduct = append(p.AllowedProduct, (*Max70Text)(&value))
}

func (p *PaymentCard21) SetServiceOption(value string) {
	p.ServiceOption = (*Max35Text)(&value)
}

func (p *PaymentCard21) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}

// Card performing the withdrawal transaction.
type PaymentCard22 struct {

	// Entry mode used to obtain the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicates the occurrence of a fall-back on the card entry mode.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData18 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`
}

func (p *PaymentCard22) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentCard22) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (p *PaymentCard22) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard22) AddPlainCardData() *PlainCardData18 {
	p.PlainCardData = new(PlainCardData18)
	return p.PlainCardData
}

func (p *PaymentCard22) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard22) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

// Card performing the withdrawal transaction.
type PaymentCard23 struct {

	// Entry mode that used to obtain the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicate the occurrence of a fall-back on the card entry mode.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData19 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`

	// Balance of the captured card or epurse if available.
	ElectronicPurseBalance *CurrencyAndAmount `xml:"ElctrncPrsBal,omitempty"`
}

func (p *PaymentCard23) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentCard23) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (p *PaymentCard23) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard23) AddPlainCardData() *PlainCardData19 {
	p.PlainCardData = new(PlainCardData19)
	return p.PlainCardData
}

func (p *PaymentCard23) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard23) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

func (p *PaymentCard23) SetElectronicPurseBalance(value, currency string) {
	p.ElectronicPurseBalance = NewCurrencyAndAmount(value, currency)
}

// Electronic money product that provides the cardholder with a portable and specialised computer device, which typically contains a microprocessor.
type PaymentCard25 struct {

	// Type of card, for example, credit card.
	Type *CardType1Code `xml:"Tp"`

	// Number embossed on a card that links the card to the account owner and account servicer.
	Number *Max35Text `xml:"Nb"`

	// Party entitled by a card issuer to use a card.
	HolderName *Max35Text `xml:"HldrNm"`

	// Year and month the card is available for use.
	StartDate *ISOYearMonth `xml:"StartDt,omitempty"`

	// Year and month the card expires.
	ExpiryDate *ISOYearMonth `xml:"XpryDt"`

	// Party that issues a payment card, as expressed by a numeric identification of the card issuer according to ISO/IEC 7812-1.
	CardIssuerName *Max35Text `xml:"CardIssrNm,omitempty"`

	// Party that issues a payment card, as expressed by a numeric identification of the card issuer according to ISO/IEC 7812-1.
	CardIssuerIdentification *PartyIdentification113 `xml:"CardIssrId,omitempty"`

	// Security code written on, or in, the card.
	SecurityCode *Max35Text `xml:"SctyCd,omitempty"`

	// Number distinguishing two or more payment cards with the same account number.
	SequenceNumber *Max3Text `xml:"SeqNb,omitempty"`
}

func (p *PaymentCard25) SetType(value string) {
	p.Type = (*CardType1Code)(&value)
}

func (p *PaymentCard25) SetNumber(value string) {
	p.Number = (*Max35Text)(&value)
}

func (p *PaymentCard25) SetHolderName(value string) {
	p.HolderName = (*Max35Text)(&value)
}

func (p *PaymentCard25) SetStartDate(value string) {
	p.StartDate = (*ISOYearMonth)(&value)
}

func (p *PaymentCard25) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISOYearMonth)(&value)
}

func (p *PaymentCard25) SetCardIssuerName(value string) {
	p.CardIssuerName = (*Max35Text)(&value)
}

func (p *PaymentCard25) AddCardIssuerIdentification() *PartyIdentification113 {
	p.CardIssuerIdentification = new(PartyIdentification113)
	return p.CardIssuerIdentification
}

func (p *PaymentCard25) SetSecurityCode(value string) {
	p.SecurityCode = (*Max35Text)(&value)
}

func (p *PaymentCard25) SetSequenceNumber(value string) {
	p.SequenceNumber = (*Max3Text)(&value)
}

// Payment card performing the transaction.
type PaymentCard3 struct {

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType2 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData2 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Exact3NumericText `xml:"CardCtryCd,omitempty"`

	// Defines a category of cards related the acceptance processing rules defined by the acquirer.
	CardProductProfile *Exact4NumericText `xml:"CardPdctPrfl,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard3) AddProtectedCardData() *ContentInformationType2 {
	p.ProtectedCardData = new(ContentInformationType2)
	return p.ProtectedCardData
}

func (p *PaymentCard3) AddPlainCardData() *PlainCardData2 {
	p.PlainCardData = new(PlainCardData2)
	return p.PlainCardData
}

func (p *PaymentCard3) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Exact3NumericText)(&value)
}

func (p *PaymentCard3) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Exact4NumericText)(&value)
}

func (p *PaymentCard3) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard3) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}

// Electronic money product that provides the cardholder with a portable and specialised computer device, which typically contains a microprocessor.
type PaymentCard4 struct {

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData1 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Exact3NumericText `xml:"CardCtryCd,omitempty"`

	// Brand name of the card.
	CardBrand *GenericIdentification1 `xml:"CardBrnd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard4) AddPlainCardData() *PlainCardData1 {
	p.PlainCardData = new(PlainCardData1)
	return p.PlainCardData
}

func (p *PaymentCard4) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Exact3NumericText)(&value)
}

func (p *PaymentCard4) AddCardBrand() *GenericIdentification1 {
	p.CardBrand = new(GenericIdentification1)
	return p.CardBrand
}

func (p *PaymentCard4) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}

// Payment card performing the transaction.
type PaymentCard5 struct {

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType5 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData1 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Defines a category of cards related to the acceptance processing rules defined by the acquirer.
	CardProductProfile *Exact4NumericText `xml:"CardPdctPrfl,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard5) AddProtectedCardData() *ContentInformationType5 {
	p.ProtectedCardData = new(ContentInformationType5)
	return p.ProtectedCardData
}

func (p *PaymentCard5) AddPlainCardData() *PlainCardData1 {
	p.PlainCardData = new(PlainCardData1)
	return p.PlainCardData
}

func (p *PaymentCard5) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard5) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Exact4NumericText)(&value)
}

func (p *PaymentCard5) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard5) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}

// Payment card performing the transaction.
type PaymentCard6 struct {

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType5 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData2 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Defines a category of cards related the acceptance processing rules defined by the acquirer.
	CardProductProfile *Exact4NumericText `xml:"CardPdctPrfl,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard6) AddProtectedCardData() *ContentInformationType5 {
	p.ProtectedCardData = new(ContentInformationType5)
	return p.ProtectedCardData
}

func (p *PaymentCard6) AddPlainCardData() *PlainCardData2 {
	p.PlainCardData = new(PlainCardData2)
	return p.PlainCardData
}

func (p *PaymentCard6) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard6) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Exact4NumericText)(&value)
}

func (p *PaymentCard6) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard6) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}

// Payment card performing the transaction.
type PaymentCard7 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType7 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData4 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3NumericText `xml:"CardCcyCd,omitempty"`

	// Defines a category of cards related to the acceptance processing rules defined by the acquirer.
	CardProductProfile *Exact4NumericText `xml:"CardPdctPrfl,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard7) AddProtectedCardData() *ContentInformationType7 {
	p.ProtectedCardData = new(ContentInformationType7)
	return p.ProtectedCardData
}

func (p *PaymentCard7) AddPlainCardData() *PlainCardData4 {
	p.PlainCardData = new(PlainCardData4)
	return p.PlainCardData
}

func (p *PaymentCard7) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard7) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3NumericText)(&value)
}

func (p *PaymentCard7) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Exact4NumericText)(&value)
}

func (p *PaymentCard7) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard7) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}

// Payment card performing the transaction.
type PaymentCard8 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType7 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData6 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3NumericText `xml:"CardCcyCd,omitempty"`

	// Defines a category of cards related the acceptance processing rules defined by the acquirer.
	CardProductProfile *Exact4NumericText `xml:"CardPdctPrfl,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard8) AddProtectedCardData() *ContentInformationType7 {
	p.ProtectedCardData = new(ContentInformationType7)
	return p.ProtectedCardData
}

func (p *PaymentCard8) AddPlainCardData() *PlainCardData6 {
	p.PlainCardData = new(PlainCardData6)
	return p.PlainCardData
}

func (p *PaymentCard8) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard8) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3NumericText)(&value)
}

func (p *PaymentCard8) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Exact4NumericText)(&value)
}

func (p *PaymentCard8) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard8) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}

// Payment card performing the transaction.
type PaymentCard9 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData7 `xml:"PlainCardData,omitempty"`

	// Bank identifier number of the issuer for routing purpose.
	IssuerBIN *Max15NumericText `xml:"IssrBIN,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`

	// Defines a category of cards related to the acceptance processing rules defined by the acquirer.
	CardProductProfile *Max35Text `xml:"CardPdctPrfl,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard9) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard9) AddPlainCardData() *PlainCardData7 {
	p.PlainCardData = new(PlainCardData7)
	return p.PlainCardData
}

func (p *PaymentCard9) SetIssuerBIN(value string) {
	p.IssuerBIN = (*Max15NumericText)(&value)
}

func (p *PaymentCard9) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard9) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

func (p *PaymentCard9) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Max35Text)(&value)
}

func (p *PaymentCard9) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard9) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}
