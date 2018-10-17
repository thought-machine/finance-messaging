package iso20022

// Reporting per financial instrument.
type FinancialInstrumentDetails1 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes4 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails3 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails1) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails1) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes4 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes4)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails1) AddSubBalance() *IntraPositionDetails3 {
	newValue := new(IntraPositionDetails3)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails10 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes36 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails20 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails10) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails10) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes36 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes36)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails10) AddSubBalance() *IntraPositionDetails20 {
	newValue := new(IntraPositionDetails20)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails13 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Information regarding the price of the instrument.
	PriceDetails *PriceInformation6 `xml:"PricDtls,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
	OpeningBalance *OpeningBalance1 `xml:"OpngBal,omitempty"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance1 `xml:"ClsgBal,omitempty"`

	// Transaction details.
	Transaction []*Transaction29 `xml:"Tx"`
}

func (f *FinancialInstrumentDetails13) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails13) AddPriceDetails() *PriceInformation6 {
	f.PriceDetails = new(PriceInformation6)
	return f.PriceDetails
}

func (f *FinancialInstrumentDetails13) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	f.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return f.SafekeepingPlace
}

func (f *FinancialInstrumentDetails13) AddOpeningBalance() *OpeningBalance1 {
	f.OpeningBalance = new(OpeningBalance1)
	return f.OpeningBalance
}

func (f *FinancialInstrumentDetails13) AddClosingBalance() *ClosingBalance1 {
	f.ClosingBalance = new(ClosingBalance1)
	return f.ClosingBalance
}

func (f *FinancialInstrumentDetails13) AddTransaction() *Transaction29 {
	newValue := new(Transaction29)
	f.Transaction = append(f.Transaction, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails14 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes36 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails28 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails14) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails14) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes36 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes36)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails14) AddSubBalance() *IntraPositionDetails28 {
	newValue := new(IntraPositionDetails28)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails17 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Information regarding the price of the instrument.
	PriceDetails *PriceInformation6 `xml:"PricDtls,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
	OpeningBalance *OpeningBalance1 `xml:"OpngBal,omitempty"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance1 `xml:"ClsgBal,omitempty"`

	// Transaction details.
	Transaction []*Transaction36 `xml:"Tx"`
}

func (f *FinancialInstrumentDetails17) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails17) AddPriceDetails() *PriceInformation6 {
	f.PriceDetails = new(PriceInformation6)
	return f.PriceDetails
}

func (f *FinancialInstrumentDetails17) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	f.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return f.SafekeepingPlace
}

func (f *FinancialInstrumentDetails17) AddOpeningBalance() *OpeningBalance1 {
	f.OpeningBalance = new(OpeningBalance1)
	return f.OpeningBalance
}

func (f *FinancialInstrumentDetails17) AddClosingBalance() *ClosingBalance1 {
	f.ClosingBalance = new(ClosingBalance1)
	return f.ClosingBalance
}

func (f *FinancialInstrumentDetails17) AddTransaction() *Transaction36 {
	newValue := new(Transaction36)
	f.Transaction = append(f.Transaction, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails2 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Information regarding the price of the instrument.
	PriceDetails *PriceInformation6 `xml:"PricDtls,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
	OpeningBalance *OpeningBalance1 `xml:"OpngBal,omitempty"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance1 `xml:"ClsgBal,omitempty"`

	// Transaction details.
	Transaction []*Transaction6 `xml:"Tx"`
}

func (f *FinancialInstrumentDetails2) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails2) AddPriceDetails() *PriceInformation6 {
	f.PriceDetails = new(PriceInformation6)
	return f.PriceDetails
}

func (f *FinancialInstrumentDetails2) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	f.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return f.SafekeepingPlace
}

func (f *FinancialInstrumentDetails2) AddOpeningBalance() *OpeningBalance1 {
	f.OpeningBalance = new(OpeningBalance1)
	return f.OpeningBalance
}

func (f *FinancialInstrumentDetails2) AddClosingBalance() *ClosingBalance1 {
	f.ClosingBalance = new(ClosingBalance1)
	return f.ClosingBalance
}

func (f *FinancialInstrumentDetails2) AddTransaction() *Transaction6 {
	newValue := new(Transaction6)
	f.Transaction = append(f.Transaction, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails20 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Information regarding the price of the instrument.
	PriceDetails *PriceInformation13 `xml:"PricDtls,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
	OpeningBalance *OpeningBalance3 `xml:"OpngBal,omitempty"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance3 `xml:"ClsgBal,omitempty"`

	// Transaction details.
	Transaction []*Transaction46 `xml:"Tx"`
}

func (f *FinancialInstrumentDetails20) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails20) AddPriceDetails() *PriceInformation13 {
	f.PriceDetails = new(PriceInformation13)
	return f.PriceDetails
}

func (f *FinancialInstrumentDetails20) AddSafekeepingPlace() *SafeKeepingPlace1 {
	f.SafekeepingPlace = new(SafeKeepingPlace1)
	return f.SafekeepingPlace
}

func (f *FinancialInstrumentDetails20) AddOpeningBalance() *OpeningBalance3 {
	f.OpeningBalance = new(OpeningBalance3)
	return f.OpeningBalance
}

func (f *FinancialInstrumentDetails20) AddClosingBalance() *ClosingBalance3 {
	f.ClosingBalance = new(ClosingBalance3)
	return f.ClosingBalance
}

func (f *FinancialInstrumentDetails20) AddTransaction() *Transaction46 {
	newValue := new(Transaction46)
	f.Transaction = append(f.Transaction, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails21 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes63 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails32 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails21) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails21) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes63 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes63)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails21) AddSubBalance() *IntraPositionDetails32 {
	newValue := new(IntraPositionDetails32)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails22 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes75 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails37 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails22) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails22) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes75 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes75)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails22) AddSubBalance() *IntraPositionDetails37 {
	newValue := new(IntraPositionDetails37)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails23 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Information regarding the price of the instrument.
	PriceDetails *PriceInformation16 `xml:"PricDtls,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
	OpeningBalance *OpeningBalance4 `xml:"OpngBal,omitempty"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance4 `xml:"ClsgBal,omitempty"`

	// Transaction details.
	Transaction []*Transaction50 `xml:"Tx"`
}

func (f *FinancialInstrumentDetails23) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails23) AddPriceDetails() *PriceInformation16 {
	f.PriceDetails = new(PriceInformation16)
	return f.PriceDetails
}

func (f *FinancialInstrumentDetails23) AddSafekeepingPlace() *SafeKeepingPlace2 {
	f.SafekeepingPlace = new(SafeKeepingPlace2)
	return f.SafekeepingPlace
}

func (f *FinancialInstrumentDetails23) AddOpeningBalance() *OpeningBalance4 {
	f.OpeningBalance = new(OpeningBalance4)
	return f.OpeningBalance
}

func (f *FinancialInstrumentDetails23) AddClosingBalance() *ClosingBalance4 {
	f.ClosingBalance = new(ClosingBalance4)
	return f.ClosingBalance
}

func (f *FinancialInstrumentDetails23) AddTransaction() *Transaction50 {
	newValue := new(Transaction50)
	f.Transaction = append(f.Transaction, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails24 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes63 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails40 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails24) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails24) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes63 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes63)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails24) AddSubBalance() *IntraPositionDetails40 {
	newValue := new(IntraPositionDetails40)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails25 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Information regarding the price of the instrument.
	PriceDetails *PriceInformation13 `xml:"PricDtls,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
	OpeningBalance *OpeningBalance3 `xml:"OpngBal,omitempty"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance3 `xml:"ClsgBal,omitempty"`

	// Transaction details.
	Transaction []*Transaction52 `xml:"Tx"`
}

func (f *FinancialInstrumentDetails25) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails25) AddPriceDetails() *PriceInformation13 {
	f.PriceDetails = new(PriceInformation13)
	return f.PriceDetails
}

func (f *FinancialInstrumentDetails25) AddSafekeepingPlace() *SafeKeepingPlace1 {
	f.SafekeepingPlace = new(SafeKeepingPlace1)
	return f.SafekeepingPlace
}

func (f *FinancialInstrumentDetails25) AddOpeningBalance() *OpeningBalance3 {
	f.OpeningBalance = new(OpeningBalance3)
	return f.OpeningBalance
}

func (f *FinancialInstrumentDetails25) AddClosingBalance() *ClosingBalance3 {
	f.ClosingBalance = new(ClosingBalance3)
	return f.ClosingBalance
}

func (f *FinancialInstrumentDetails25) AddTransaction() *Transaction52 {
	newValue := new(Transaction52)
	f.Transaction = append(f.Transaction, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails26 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes75 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails44 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails26) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails26) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes75 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes75)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails26) AddSubBalance() *IntraPositionDetails44 {
	newValue := new(IntraPositionDetails44)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails27 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Information regarding the price of the instrument.
	PriceDetails *PriceInformation16 `xml:"PricDtls,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
	OpeningBalance *OpeningBalance4 `xml:"OpngBal,omitempty"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance4 `xml:"ClsgBal,omitempty"`

	// Transaction details.
	Transaction []*Transaction55 `xml:"Tx"`
}

func (f *FinancialInstrumentDetails27) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails27) AddPriceDetails() *PriceInformation16 {
	f.PriceDetails = new(PriceInformation16)
	return f.PriceDetails
}

func (f *FinancialInstrumentDetails27) AddSafekeepingPlace() *SafeKeepingPlace2 {
	f.SafekeepingPlace = new(SafeKeepingPlace2)
	return f.SafekeepingPlace
}

func (f *FinancialInstrumentDetails27) AddOpeningBalance() *OpeningBalance4 {
	f.OpeningBalance = new(OpeningBalance4)
	return f.OpeningBalance
}

func (f *FinancialInstrumentDetails27) AddClosingBalance() *ClosingBalance4 {
	f.ClosingBalance = new(ClosingBalance4)
	return f.ClosingBalance
}

func (f *FinancialInstrumentDetails27) AddTransaction() *Transaction55 {
	newValue := new(Transaction55)
	f.Transaction = append(f.Transaction, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails5 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Information regarding the price of the instrument.
	PriceDetails *PriceInformation6 `xml:"PricDtls,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
	OpeningBalance *OpeningBalance1 `xml:"OpngBal,omitempty"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance1 `xml:"ClsgBal,omitempty"`

	// Transaction details.
	Transaction []*Transaction13 `xml:"Tx"`
}

func (f *FinancialInstrumentDetails5) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails5) AddPriceDetails() *PriceInformation6 {
	f.PriceDetails = new(PriceInformation6)
	return f.PriceDetails
}

func (f *FinancialInstrumentDetails5) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	f.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return f.SafekeepingPlace
}

func (f *FinancialInstrumentDetails5) AddOpeningBalance() *OpeningBalance1 {
	f.OpeningBalance = new(OpeningBalance1)
	return f.OpeningBalance
}

func (f *FinancialInstrumentDetails5) AddClosingBalance() *ClosingBalance1 {
	f.ClosingBalance = new(ClosingBalance1)
	return f.ClosingBalance
}

func (f *FinancialInstrumentDetails5) AddTransaction() *Transaction13 {
	newValue := new(Transaction13)
	f.Transaction = append(f.Transaction, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails6 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes21 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails17 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails6) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails6) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes21 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes21)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails6) AddSubBalance() *IntraPositionDetails17 {
	newValue := new(IntraPositionDetails17)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}

// Reporting per financial instrument.
type FinancialInstrumentDetails9 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Information regarding the price of the instrument.
	PriceDetails *PriceInformation6 `xml:"PricDtls,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
	OpeningBalance *OpeningBalance1 `xml:"OpngBal,omitempty"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance1 `xml:"ClsgBal,omitempty"`

	// Transaction details.
	Transaction []*Transaction18 `xml:"Tx"`
}

func (f *FinancialInstrumentDetails9) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails9) AddPriceDetails() *PriceInformation6 {
	f.PriceDetails = new(PriceInformation6)
	return f.PriceDetails
}

func (f *FinancialInstrumentDetails9) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	f.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return f.SafekeepingPlace
}

func (f *FinancialInstrumentDetails9) AddOpeningBalance() *OpeningBalance1 {
	f.OpeningBalance = new(OpeningBalance1)
	return f.OpeningBalance
}

func (f *FinancialInstrumentDetails9) AddClosingBalance() *ClosingBalance1 {
	f.ClosingBalance = new(ClosingBalance1)
	return f.ClosingBalance
}

func (f *FinancialInstrumentDetails9) AddTransaction() *Transaction18 {
	newValue := new(Transaction18)
	f.Transaction = append(f.Transaction, newValue)
	return newValue
}
