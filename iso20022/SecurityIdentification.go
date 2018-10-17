package iso20022

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type SecurityIdentification1 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification7 `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Name of the umbrella fund in which financial instrument is contained.
	UmbrellaName *Max35Text `xml:"UmbrllNm,omitempty"`

	// Currency of the investment fund class.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy"`

	// Country where the fund has legal domicile as reflected in the ISIN classification.
	CountryOfDomicile *CountryCode `xml:"CtryOfDmcl"`

	// Countries where the fund is registered for distribution.
	RegisteredDistributionCountry []*CountryCode `xml:"RegdDstrbtnCtry"`
}

func (s *SecurityIdentification1) AddIdentification() *SecurityIdentification7 {
	s.Identification = new(SecurityIdentification7)
	return s.Identification
}

func (s *SecurityIdentification1) SetName(value string) {
	s.Name = (*Max350Text)(&value)
}

func (s *SecurityIdentification1) SetClassType(value string) {
	s.ClassType = (*Max35Text)(&value)
}

func (s *SecurityIdentification1) SetUmbrellaName(value string) {
	s.UmbrellaName = (*Max35Text)(&value)
}

func (s *SecurityIdentification1) SetBaseCurrency(value string) {
	s.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SecurityIdentification1) SetCountryOfDomicile(value string) {
	s.CountryOfDomicile = (*CountryCode)(&value)
}

func (s *SecurityIdentification1) AddRegisteredDistributionCountry(value string) {
	s.RegisteredDistributionCountry = append(s.RegisteredDistributionCountry, (*CountryCode)(&value))
}

// Identification of a security.
type SecurityIdentification11 struct {

	// Identification of a security.
	Identification *SecurityIdentification11Choice `xml:"Id"`

	// Textual description of a security instrument.
	Description *Max140Text `xml:"Desc,omitempty"`
}

func (s *SecurityIdentification11) AddIdentification() *SecurityIdentification11Choice {
	s.Identification = new(SecurityIdentification11Choice)
	return s.Identification
}

func (s *SecurityIdentification11) SetDescription(value string) {
	s.Description = (*Max140Text)(&value)
}

// Identification of a security.
type SecurityIdentification14 struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINIdentifier `xml:"ISIN,omitempty"`

	// Identification of a security by proprietary or domestic identification scheme.
	OtherIdentification []*OtherIdentification1 `xml:"OthrId,omitempty"`

	// Textual description of a security instrument.
	Description *Max140Text `xml:"Desc,omitempty"`
}

func (s *SecurityIdentification14) SetISIN(value string) {
	s.ISIN = (*ISINIdentifier)(&value)
}

func (s *SecurityIdentification14) AddOtherIdentification() *OtherIdentification1 {
	newValue := new(OtherIdentification1)
	s.OtherIdentification = append(s.OtherIdentification, newValue)
	return newValue
}

func (s *SecurityIdentification14) SetDescription(value string) {
	s.Description = (*Max140Text)(&value)
}

// Identification of a security.
type SecurityIdentification19 struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINOct2015Identifier `xml:"ISIN,omitempty"`

	// Identification of a security by proprietary or domestic identification scheme.
	OtherIdentification []*OtherIdentification1 `xml:"OthrId,omitempty"`

	// Textual description of a security instrument.
	Description *Max140Text `xml:"Desc,omitempty"`
}

func (s *SecurityIdentification19) SetISIN(value string) {
	s.ISIN = (*ISINOct2015Identifier)(&value)
}

func (s *SecurityIdentification19) AddOtherIdentification() *OtherIdentification1 {
	newValue := new(OtherIdentification1)
	s.OtherIdentification = append(s.OtherIdentification, newValue)
	return newValue
}

func (s *SecurityIdentification19) SetDescription(value string) {
	s.Description = (*Max140Text)(&value)
}

// Identification of a security.
type SecurityIdentification20 struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINOct2015Identifier `xml:"ISIN,omitempty"`

	// Identification of a security by proprietary or domestic identification scheme.
	OtherIdentification []*OtherIdentification2 `xml:"OthrId,omitempty"`

	// Textual description of a security instrument.
	Description *RestrictedFINXMax140Text `xml:"Desc,omitempty"`
}

func (s *SecurityIdentification20) SetISIN(value string) {
	s.ISIN = (*ISINOct2015Identifier)(&value)
}

func (s *SecurityIdentification20) AddOtherIdentification() *OtherIdentification2 {
	newValue := new(OtherIdentification2)
	s.OtherIdentification = append(s.OtherIdentification, newValue)
	return newValue
}

func (s *SecurityIdentification20) SetDescription(value string) {
	s.Description = (*RestrictedFINXMax140Text)(&value)
}

// Identification of a security.
type SecurityIdentification21 struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINOct2015Identifier `xml:"ISIN,omitempty"`

	// Identification of a security by proprietary or domestic identification scheme.
	OtherIdentification []*OtherIdentification2 `xml:"OthrId,omitempty"`

	// Textual description of a security instrument.
	Description *RestrictedFINXMax140Text `xml:"Desc,omitempty"`
}

func (s *SecurityIdentification21) SetISIN(value string) {
	s.ISIN = (*ISINOct2015Identifier)(&value)
}

func (s *SecurityIdentification21) AddOtherIdentification() *OtherIdentification2 {
	newValue := new(OtherIdentification2)
	s.OtherIdentification = append(s.OtherIdentification, newValue)
	return newValue
}

func (s *SecurityIdentification21) SetDescription(value string) {
	s.Description = (*RestrictedFINXMax140Text)(&value)
}

// Identification of a security by its symbol.
type SecurityIdentification3 struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINIdentifier `xml:"ISIN"`

	// Letters that identify a stock traded on a stock exchange. The Ticker Symbol is a short and convenient way of identifying a stock, eg, RTR.L for Reuters quoted in London.
	TickerSymbol *TickerIdentifier `xml:"TckrSymb,omitempty"`

	// Committee on Uniform Securities and Identification Procedures (CUSIP). The standards body that created and maintains the securities classification system in the US. The CUSIP is composed of a 9-character number that uniquely identifies a particular security.  Non-US securities have a similar number called the CINS number.
	CUSIP *CUSIPIdentifier `xml:"CUSIP,omitempty"`

	// Stock Exchange Daily Official List (SEDOL) number.  A code used by the London Stock Exchange to identify foreign stocks, especially those that aren't actively traded in the US and don't have a CUSIP number.
	SEDOL *SEDOLIdentifier `xml:"SEDOL,omitempty"`

	// Identifier of a security assigned by the Japanese QUICK identification scheme for financial instruments.
	QUICK *QUICKIdentifier `xml:"QUICK,omitempty"`

	// Proprietary identification of a security assigned by an institution or organisation.
	OtherIdentification *AlternateFinancialInstrumentIdentification1 `xml:"OthrId,omitempty"`
}

func (s *SecurityIdentification3) SetISIN(value string) {
	s.ISIN = (*ISINIdentifier)(&value)
}

func (s *SecurityIdentification3) SetTickerSymbol(value string) {
	s.TickerSymbol = (*TickerIdentifier)(&value)
}

func (s *SecurityIdentification3) SetCUSIP(value string) {
	s.CUSIP = (*CUSIPIdentifier)(&value)
}

func (s *SecurityIdentification3) SetSEDOL(value string) {
	s.SEDOL = (*SEDOLIdentifier)(&value)
}

func (s *SecurityIdentification3) SetQUICK(value string) {
	s.QUICK = (*QUICKIdentifier)(&value)
}

func (s *SecurityIdentification3) AddOtherIdentification() *AlternateFinancialInstrumentIdentification1 {
	s.OtherIdentification = new(AlternateFinancialInstrumentIdentification1)
	return s.OtherIdentification
}

// Identification of a security.
type SecurityIdentification32 struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINOct2015Identifier `xml:"ISIN,omitempty"`

	// Identification of a security by proprietary or domestic identification scheme.
	OtherIdentification []*OtherIdentification3 `xml:"OthrId,omitempty"`

	// Textual description of a security instrument.
	Description *RestrictedFINXMax140Text `xml:"Desc,omitempty"`
}

func (s *SecurityIdentification32) SetISIN(value string) {
	s.ISIN = (*ISINOct2015Identifier)(&value)
}

func (s *SecurityIdentification32) AddOtherIdentification() *OtherIdentification3 {
	newValue := new(OtherIdentification3)
	s.OtherIdentification = append(s.OtherIdentification, newValue)
	return newValue
}

func (s *SecurityIdentification32) SetDescription(value string) {
	s.Description = (*RestrictedFINXMax140Text)(&value)
}

// Choice between ISIN and an alternative format for the identification of a security. ISIN is the preferred format.
type SecurityIdentification7 struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINIdentifier `xml:"ISIN"`

	// Proprietary identification of a security assigned by an institution or organisation.
	OtherIdentification *AlternateSecurityIdentification3 `xml:"OthrId"`

	// Textual description of a security instrument.
	Description *Max140Text `xml:"Desc,omitempty"`
}

func (s *SecurityIdentification7) SetISIN(value string) {
	s.ISIN = (*ISINIdentifier)(&value)
}

func (s *SecurityIdentification7) AddOtherIdentification() *AlternateSecurityIdentification3 {
	s.OtherIdentification = new(AlternateSecurityIdentification3)
	return s.OtherIdentification
}

func (s *SecurityIdentification7) SetDescription(value string) {
	s.Description = (*Max140Text)(&value)
}
