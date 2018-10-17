package iso20022

// Other type of party.
type ExtendedParty1 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation4 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty1) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty1) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation4 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation4)
	return e.OtherPartyDetails
}

// Other type of party.
type ExtendedParty10 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation13 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty10) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty10) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation13 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation13)
	return e.OtherPartyDetails
}

// Party and account information.
type ExtendedParty11 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation14 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty11) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty11) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation14 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation14)
	return e.OtherPartyDetails
}

// Other type of party.
type ExtendedParty12 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation15 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty12) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty12) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation15 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation15)
	return e.OtherPartyDetails
}

// Other type of party.
type ExtendedParty2 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation5 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty2) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty2) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation5 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation5)
	return e.OtherPartyDetails
}

// Other type of party.
type ExtendedParty3 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation6 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty3) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty3) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation6 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation6)
	return e.OtherPartyDetails
}

// Other type of party.
type ExtendedParty4 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation7 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty4) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty4) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation7 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation7)
	return e.OtherPartyDetails
}

// Other type of party.
type ExtendedParty5 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation8 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty5) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty5) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation8 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation8)
	return e.OtherPartyDetails
}

// Other type of party.
type ExtendedParty6 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation9 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty6) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty6) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation9 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation9)
	return e.OtherPartyDetails
}

// Other type of party.
type ExtendedParty7 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation10 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty7) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty7) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation10 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation10)
	return e.OtherPartyDetails
}

// Other type of party.
type ExtendedParty8 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation11 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty8) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty8) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation11 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation11)
	return e.OtherPartyDetails
}

// Party and account information.
type ExtendedParty9 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation12 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty9) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty9) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation12 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation12)
	return e.OtherPartyDetails
}
