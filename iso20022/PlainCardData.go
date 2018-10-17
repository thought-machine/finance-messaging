package iso20022

// Sensible data associated with the payment card performing the transaction.
type PlainCardData1 struct {

	// Primary Account Number (PAN) of the card, or card number.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card inside a set of cards with the same card number (PAN).
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date as from which the card can be used.
	EffectiveDate *ISOYearMonth `xml:"FctvDt,omitempty"`

	// Expiry date of the card.
	ExpiryDate *ISOYearMonth `xml:"XpryDt"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// Magnetic track or equivalent payment card data.
	TrackData []*TrackData1 `xml:"TrckData,omitempty"`

	// Card security code (CSC) associated with the card performing the transaction.
	CardSecurityCode *CardSecurityInformation1 `xml:"CardSctyCd,omitempty"`
}

func (p *PlainCardData1) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData1) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData1) SetEffectiveDate(value string) {
	p.EffectiveDate = (*ISOYearMonth)(&value)
}

func (p *PlainCardData1) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISOYearMonth)(&value)
}

func (p *PlainCardData1) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData1) AddTrackData() *TrackData1 {
	newValue := new(TrackData1)
	p.TrackData = append(p.TrackData, newValue)
	return newValue
}

func (p *PlainCardData1) AddCardSecurityCode() *CardSecurityInformation1 {
	p.CardSecurityCode = new(CardSecurityInformation1)
	return p.CardSecurityCode
}

// Sensible data associated with the payment card performing the transaction.
type PlainCardData10 struct {

	// Primary Account Number (PAN) of the card, or surrogate of the PAN by a payment token.
	// It correspond to the ISO 8583 field number 2.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card or a payment token inside a set of cards with the same PAN or token.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card or the payment token expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt,omitempty"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// Track issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The track value might be provided by a payment token.
	TrackData []*TrackData1 `xml:"TrckData,omitempty"`

	// Name of the cardholder stored on the card.
	CardholderName *Max45Text `xml:"CrdhldrNm,omitempty"`
}

func (p *PlainCardData10) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData10) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData10) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData10) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData10) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData10) AddTrackData() *TrackData1 {
	newValue := new(TrackData1)
	p.TrackData = append(p.TrackData, newValue)
	return newValue
}

func (p *PlainCardData10) SetCardholderName(value string) {
	p.CardholderName = (*Max45Text)(&value)
}

// Sensitive data associated with the payment card performing the transaction.
type PlainCardData11 struct {

	// Primary Account Number (PAN) of the card, or surrogate of the PAN by a payment token.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card or a payment token inside a set of cards with the same PAN or token.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card or the payment token expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt,omitempty"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`
}

func (p *PlainCardData11) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData11) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData11) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData11) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData11) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

// Sensible data associated with the payment card performing the transaction.
type PlainCardData12 struct {

	// Primary Account Number (PAN) of the card, or surrogate of the PAN by a payment token.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card or a payment token inside a set of cards with the same PAN or token.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Expiry date of the card or the payment token expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt,omitempty"`
}

func (p *PlainCardData12) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData12) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData12) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

// Sensible data associated with the payment card performing the transaction.
type PlainCardData13 struct {

	// Primary Account Number (PAN) of the card.
	PAN *Min8Max28NumericText `xml:"PAN,omitempty"`

	// Identify a card or a payment token inside a set of cards with the same PAN.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt,omitempty"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// Track number 1 from magnetic stripe card.
	Track1 *Max140Text `xml:"Trck1,omitempty"`

	// Track number 2 without control characters (start /end and LRC) issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read.
	Track2 *Max140Text `xml:"Trck2,omitempty"`

	// Track number 3 from magnetic stripe card.
	Track3 *Max140Text `xml:"Trck3,omitempty"`

	// Name of the cardholder stored on the card.
	CardholderName *Max45Text `xml:"CrdhldrNm,omitempty"`
}

func (p *PlainCardData13) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData13) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData13) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData13) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData13) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData13) SetTrack1(value string) {
	p.Track1 = (*Max140Text)(&value)
}

func (p *PlainCardData13) SetTrack2(value string) {
	p.Track2 = (*Max140Text)(&value)
}

func (p *PlainCardData13) SetTrack3(value string) {
	p.Track3 = (*Max140Text)(&value)
}

func (p *PlainCardData13) SetCardholderName(value string) {
	p.CardholderName = (*Max45Text)(&value)
}

// Sensible data associated with the payment card performing the transaction.
type PlainCardData14 struct {

	// Primary Account Number (PAN) of the card.
	PAN *Min8Max28NumericText `xml:"PAN,omitempty"`

	// Identify a card or a payment token inside a set of cards with the same PAN.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt,omitempty"`

	// Track number 1 from magnetic stripe card.
	Track1 *Max140Text `xml:"Trck1,omitempty"`

	// Track number 2 without control characters (start /end and LRC).
	Track2 *Max140Text `xml:"Trck2,omitempty"`

	// Track number 3 from magnetic stripe card.
	Track3 *Max140Text `xml:"Trck3,omitempty"`
}

func (p *PlainCardData14) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData14) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData14) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData14) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData14) SetTrack1(value string) {
	p.Track1 = (*Max140Text)(&value)
}

func (p *PlainCardData14) SetTrack2(value string) {
	p.Track2 = (*Max140Text)(&value)
}

func (p *PlainCardData14) SetTrack3(value string) {
	p.Track3 = (*Max140Text)(&value)
}

// Sensible data associated with the payment card performing the transaction.
type PlainCardData15 struct {

	// Primary Account Number (PAN) of the card, or surrogate of the PAN by a payment token.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card or a payment token inside a set of cards with the same PAN or token.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card or the payment token expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// ISO track 1 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The format is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track1 *Max76Text `xml:"Trck1,omitempty"`

	// ISO track 2 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track2 *Max37Text `xml:"Trck2,omitempty"`

	// ISO track 3 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 4909, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track3 *Max104Text `xml:"Trck3,omitempty"`

	// Name of the cardholder stored on the card.
	CardholderName *Max45Text `xml:"CrdhldrNm,omitempty"`
}

func (p *PlainCardData15) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData15) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData15) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData15) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData15) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData15) SetTrack1(value string) {
	p.Track1 = (*Max76Text)(&value)
}

func (p *PlainCardData15) SetTrack2(value string) {
	p.Track2 = (*Max37Text)(&value)
}

func (p *PlainCardData15) SetTrack3(value string) {
	p.Track3 = (*Max104Text)(&value)
}

func (p *PlainCardData15) SetCardholderName(value string) {
	p.CardholderName = (*Max45Text)(&value)
}

// Sensitive data associated with the payment card performing the transaction.
type PlainCardData16 struct {

	// Primary Account Number (PAN) of the card, or surrogate of the PAN by a payment token.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card or a payment token inside a set of cards with the same PAN or token.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card or the payment token expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// ISO track 1 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The format is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track1 *Max76Text `xml:"Trck1,omitempty"`

	// ISO track 2 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track2 *Max37Text `xml:"Trck2,omitempty"`

	// ISO track 3 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 4909, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track3 *Max104Text `xml:"Trck3,omitempty"`
}

func (p *PlainCardData16) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData16) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData16) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData16) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData16) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData16) SetTrack1(value string) {
	p.Track1 = (*Max76Text)(&value)
}

func (p *PlainCardData16) SetTrack2(value string) {
	p.Track2 = (*Max37Text)(&value)
}

func (p *PlainCardData16) SetTrack3(value string) {
	p.Track3 = (*Max104Text)(&value)
}

// Sensitive data associated with a payment card.
type PlainCardData17 struct {

	// Primary Account Number (PAN) of the card.
	PAN *Min8Max28NumericText `xml:"PAN,omitempty"`

	// ISO track 1 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The format is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track1 *Max76Text `xml:"Trck1,omitempty"`

	// ISO track 2 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track2 *Max37Text `xml:"Trck2,omitempty"`

	// ISO track 3 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 4909, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track3 *Max104Text `xml:"Trck3,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData []*Max35Text `xml:"AddtlCardData,omitempty"`

	// Entry mode of the card.
	EntryMode *CardDataReading5Code `xml:"NtryMd,omitempty"`
}

func (p *PlainCardData17) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData17) SetTrack1(value string) {
	p.Track1 = (*Max76Text)(&value)
}

func (p *PlainCardData17) SetTrack2(value string) {
	p.Track2 = (*Max37Text)(&value)
}

func (p *PlainCardData17) SetTrack3(value string) {
	p.Track3 = (*Max104Text)(&value)
}

func (p *PlainCardData17) AddAdditionalCardData(value string) {
	p.AdditionalCardData = append(p.AdditionalCardData, (*Max35Text)(&value))
}

func (p *PlainCardData17) SetEntryMode(value string) {
	p.EntryMode = (*CardDataReading5Code)(&value)
}

// Sensible data associated with the payment card performing the transaction.
type PlainCardData18 struct {

	// Primary Account Number (PAN) of the card.
	PAN *Min8Max28NumericText `xml:"PAN,omitempty"`

	// Identify a card or a payment token inside a set of cards with the same PAN.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt,omitempty"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// ISO track 1 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The format is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track1 *Max76Text `xml:"Trck1,omitempty"`

	// ISO track 2 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track2 *Max37Text `xml:"Trck2,omitempty"`

	// ISO track 3 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 4909, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track3 *Max104Text `xml:"Trck3,omitempty"`

	// Name of the cardholder stored on the card.
	CardholderName *Max45Text `xml:"CrdhldrNm,omitempty"`
}

func (p *PlainCardData18) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData18) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData18) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData18) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData18) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData18) SetTrack1(value string) {
	p.Track1 = (*Max76Text)(&value)
}

func (p *PlainCardData18) SetTrack2(value string) {
	p.Track2 = (*Max37Text)(&value)
}

func (p *PlainCardData18) SetTrack3(value string) {
	p.Track3 = (*Max104Text)(&value)
}

func (p *PlainCardData18) SetCardholderName(value string) {
	p.CardholderName = (*Max45Text)(&value)
}

// Sensible data associated with the payment card performing the transaction.
type PlainCardData19 struct {

	// Primary Account Number (PAN) of the card.
	PAN *Min8Max28NumericText `xml:"PAN,omitempty"`

	// Identify a card or a payment token inside a set of cards with the same PAN.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt,omitempty"`

	// ISO track 1 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The format is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track1 *Max76Text `xml:"Trck1,omitempty"`

	// ISO track 2 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track2 *Max37Text `xml:"Trck2,omitempty"`

	// ISO track 3 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 4909, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track3 *Max104Text `xml:"Trck3,omitempty"`
}

func (p *PlainCardData19) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData19) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData19) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData19) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData19) SetTrack1(value string) {
	p.Track1 = (*Max76Text)(&value)
}

func (p *PlainCardData19) SetTrack2(value string) {
	p.Track2 = (*Max37Text)(&value)
}

func (p *PlainCardData19) SetTrack3(value string) {
	p.Track3 = (*Max104Text)(&value)
}

// Sensible data associated with the payment card performing the transaction.
type PlainCardData2 struct {

	// Primary Account Number (PAN) of the card, or card number.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card inside a set of cards with the same card number (PAN).
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date as from which the card can be used.
	EffectiveDate *ISOYearMonth `xml:"FctvDt,omitempty"`

	// Expiry date of the card.
	ExpiryDate *ISOYearMonth `xml:"XpryDt"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// Magnetic track or equivalent payment card data.
	TrackData []*TrackData1 `xml:"TrckData,omitempty"`
}

func (p *PlainCardData2) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData2) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData2) SetEffectiveDate(value string) {
	p.EffectiveDate = (*ISOYearMonth)(&value)
}

func (p *PlainCardData2) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISOYearMonth)(&value)
}

func (p *PlainCardData2) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData2) AddTrackData() *TrackData1 {
	newValue := new(TrackData1)
	p.TrackData = append(p.TrackData, newValue)
	return newValue
}

// Sensible data associated with the payment card performing the transaction provided for verification in response.
type PlainCardData3 struct {

	// Primary Account Number (PAN) of the card, or card number.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card inside a set of cards with the same card number (PAN).
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date as from which the card can be used.
	EffectiveDate *ISOYearMonth `xml:"FctvDt,omitempty"`

	// Expiry date of the card.
	ExpiryDate *ISOYearMonth `xml:"XpryDt"`
}

func (p *PlainCardData3) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData3) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData3) SetEffectiveDate(value string) {
	p.EffectiveDate = (*ISOYearMonth)(&value)
}

func (p *PlainCardData3) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISOYearMonth)(&value)
}

// Sensible data associated with the payment card performing the transaction.
type PlainCardData4 struct {

	// Primary Account Number (PAN) of the card, or card number.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card inside a set of cards with the same card number (PAN).
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// Magnetic track or equivalent payment card data.
	TrackData []*TrackData1 `xml:"TrckData,omitempty"`

	// Card security code (CSC) associated with the card performing the transaction.
	CardSecurityCode *CardSecurityInformation1 `xml:"CardSctyCd,omitempty"`
}

func (p *PlainCardData4) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData4) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData4) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData4) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData4) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData4) AddTrackData() *TrackData1 {
	newValue := new(TrackData1)
	p.TrackData = append(p.TrackData, newValue)
	return newValue
}

func (p *PlainCardData4) AddCardSecurityCode() *CardSecurityInformation1 {
	p.CardSecurityCode = new(CardSecurityInformation1)
	return p.CardSecurityCode
}

// Sensible data associated with the payment card performing the transaction provided for verification in response.
type PlainCardData5 struct {

	// Primary Account Number (PAN) of the card, or card number.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card inside a set of cards with the same card number (PAN).
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt"`
}

func (p *PlainCardData5) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData5) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData5) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData5) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

// Sensible data associated with the payment card performing the transaction.
type PlainCardData6 struct {

	// Primary Account Number (PAN) of the card, or card number.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card inside a set of cards with the same card number (PAN).
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// Magnetic track or equivalent payment card data.
	TrackData []*TrackData1 `xml:"TrckData,omitempty"`
}

func (p *PlainCardData6) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData6) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData6) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData6) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData6) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData6) AddTrackData() *TrackData1 {
	newValue := new(TrackData1)
	p.TrackData = append(p.TrackData, newValue)
	return newValue
}

// Sensible data associated with the payment card performing the transaction.
type PlainCardData7 struct {

	// Primary Account Number (PAN) of the card, or surrogate of the PAN by a payment token.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card or a payment token inside a set of cards with the same PAN or token.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card or the payment token expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// Track issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The track value might be provided by a payment token.
	TrackData []*TrackData1 `xml:"TrckData,omitempty"`

	// Name of the cardholder stored on the card.
	CardholderName *Max45Text `xml:"CrdhldrNm,omitempty"`
}

func (p *PlainCardData7) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData7) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData7) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData7) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData7) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData7) AddTrackData() *TrackData1 {
	newValue := new(TrackData1)
	p.TrackData = append(p.TrackData, newValue)
	return newValue
}

func (p *PlainCardData7) SetCardholderName(value string) {
	p.CardholderName = (*Max45Text)(&value)
}

// Sensible data associated with the payment card performing the transaction.
type PlainCardData8 struct {

	// Primary Account Number (PAN) of the card, , or surrogate of the PAN by a payment token.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card or a payment token inside a set of cards with the same PAN or token.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date as from which the card can be used.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card or the payment token expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt"`
}

func (p *PlainCardData8) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData8) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData8) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData8) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

// Sensitive data associated with the payment card performing the transaction.
type PlainCardData9 struct {

	// Primary Account Number (PAN) of the card, or surrogate of the PAN by a payment token.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card or a payment token inside a set of cards with the same PAN or token.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card or the payment token expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// Track issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The track value might be provided by a payment token.
	TrackData []*TrackData1 `xml:"TrckData,omitempty"`
}

func (p *PlainCardData9) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData9) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData9) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData9) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData9) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData9) AddTrackData() *TrackData1 {
	newValue := new(TrackData1)
	p.TrackData = append(p.TrackData, newValue)
	return newValue
}
