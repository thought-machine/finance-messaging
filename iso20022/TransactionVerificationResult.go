package iso20022

// Result of the verifications performed by the issuer to deliver or decline the authorisation.
type TransactionVerificationResult1 struct {

	// Result of an e-commerce authentication process.
	ElectronicCommerceAuthenticationResult *Max500Text `xml:"ElctrncComrcAuthntcnRslt,omitempty"`

	// Result of the printed card security code (CSC) validation.
	CSCResult *CSCResult1Code `xml:"CSCRslt,omitempty"`

	// Result of the cardholder verification address checks on the street number and the postal code.
	CardholderAddressVerificationResult *CardholderAddressVerificationResult1Code `xml:"CrdhldrAdrVrfctnRslt,omitempty"`

	// Product code for which the authorisation was declined.
	DeclinedProductCode []*Max70Text `xml:"DclndPdctCd,omitempty"`
}

func (t *TransactionVerificationResult1) SetElectronicCommerceAuthenticationResult(value string) {
	t.ElectronicCommerceAuthenticationResult = (*Max500Text)(&value)
}

func (t *TransactionVerificationResult1) SetCSCResult(value string) {
	t.CSCResult = (*CSCResult1Code)(&value)
}

func (t *TransactionVerificationResult1) SetCardholderAddressVerificationResult(value string) {
	t.CardholderAddressVerificationResult = (*CardholderAddressVerificationResult1Code)(&value)
}

func (t *TransactionVerificationResult1) AddDeclinedProductCode(value string) {
	t.DeclinedProductCode = append(t.DeclinedProductCode, (*Max70Text)(&value))
}

// Result of the verifications performed by the issuer to deliver or decline the authorisation.
type TransactionVerificationResult2 struct {

	// Result of an e-commerce authentication process.
	ElectronicCommerceAuthenticationResult *Max500Text `xml:"ElctrncComrcAuthntcnRslt,omitempty"`

	// Result of the printed card security code (CSC) validation.
	CSCResult *CSCResult1Code `xml:"CSCRslt,omitempty"`

	// Result of the cardholder verification address checks on the street number and the postal code.
	CardholderAddressVerificationResult []*CardholderAddressVerificationResult1Code `xml:"CrdhldrAdrVrfctnRslt,omitempty"`

	// Product code for which the authorisation was declined.
	DeclinedProductCode []*Max70Text `xml:"DclndPdctCd,omitempty"`
}

func (t *TransactionVerificationResult2) SetElectronicCommerceAuthenticationResult(value string) {
	t.ElectronicCommerceAuthenticationResult = (*Max500Text)(&value)
}

func (t *TransactionVerificationResult2) SetCSCResult(value string) {
	t.CSCResult = (*CSCResult1Code)(&value)
}

func (t *TransactionVerificationResult2) AddCardholderAddressVerificationResult(value string) {
	t.CardholderAddressVerificationResult = append(t.CardholderAddressVerificationResult, (*CardholderAddressVerificationResult1Code)(&value))
}

func (t *TransactionVerificationResult2) AddDeclinedProductCode(value string) {
	t.DeclinedProductCode = append(t.DeclinedProductCode, (*Max70Text)(&value))
}

// Result of performed verifications for the transaction.
type TransactionVerificationResult3 struct {

	// Method of verification that has been used.
	Method *AuthenticationMethod4Code `xml:"Mtd"`

	// Entity or device that has performed the verification.
	VerificationEntity *AuthenticationEntity2Code `xml:"VrfctnNtty,omitempty"`

	// Result of the verification.
	Result *Verification1Code `xml:"Rslt,omitempty"`

	// Additional result of the verification.
	AdditionalResult *Max500Text `xml:"AddtlRslt,omitempty"`
}

func (t *TransactionVerificationResult3) SetMethod(value string) {
	t.Method = (*AuthenticationMethod4Code)(&value)
}

func (t *TransactionVerificationResult3) SetVerificationEntity(value string) {
	t.VerificationEntity = (*AuthenticationEntity2Code)(&value)
}

func (t *TransactionVerificationResult3) SetResult(value string) {
	t.Result = (*Verification1Code)(&value)
}

func (t *TransactionVerificationResult3) SetAdditionalResult(value string) {
	t.AdditionalResult = (*Max500Text)(&value)
}

// Result of performed verifications for the transaction.
type TransactionVerificationResult4 struct {

	// Method of verification that has been performed.
	Method *AuthenticationMethod6Code `xml:"Mtd"`

	// Entity or device that has performed the verification.
	VerificationEntity *AuthenticationEntity2Code `xml:"VrfctnNtty,omitempty"`

	// Result of the verification.
	Result *Verification1Code `xml:"Rslt,omitempty"`

	// Additional result of the verification.
	AdditionalResult *Max500Text `xml:"AddtlRslt,omitempty"`
}

func (t *TransactionVerificationResult4) SetMethod(value string) {
	t.Method = (*AuthenticationMethod6Code)(&value)
}

func (t *TransactionVerificationResult4) SetVerificationEntity(value string) {
	t.VerificationEntity = (*AuthenticationEntity2Code)(&value)
}

func (t *TransactionVerificationResult4) SetResult(value string) {
	t.Result = (*Verification1Code)(&value)
}

func (t *TransactionVerificationResult4) SetAdditionalResult(value string) {
	t.AdditionalResult = (*Max500Text)(&value)
}

// Result of performed verifications for the transaction.
type TransactionVerificationResult5 struct {

	// Method of verification that has been performed.
	Method *AuthenticationMethod7Code `xml:"Mtd"`

	// Entity or device that has performed the verification.
	VerificationEntity *AuthenticationEntity2Code `xml:"VrfctnNtty,omitempty"`

	// Result of the verification.
	Result *Verification1Code `xml:"Rslt,omitempty"`

	// Additional result of the verification.
	AdditionalResult *Max500Text `xml:"AddtlRslt,omitempty"`

	// Token provided to the ATM for further proof of authentication.
	AuthenticationToken *Max140Binary `xml:"AuthntcnTkn,omitempty"`
}

func (t *TransactionVerificationResult5) SetMethod(value string) {
	t.Method = (*AuthenticationMethod7Code)(&value)
}

func (t *TransactionVerificationResult5) SetVerificationEntity(value string) {
	t.VerificationEntity = (*AuthenticationEntity2Code)(&value)
}

func (t *TransactionVerificationResult5) SetResult(value string) {
	t.Result = (*Verification1Code)(&value)
}

func (t *TransactionVerificationResult5) SetAdditionalResult(value string) {
	t.AdditionalResult = (*Max500Text)(&value)
}

func (t *TransactionVerificationResult5) SetAuthenticationToken(value string) {
	t.AuthenticationToken = (*Max140Binary)(&value)
}
