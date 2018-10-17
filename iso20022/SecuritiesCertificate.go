package iso20022

// Physical representation of a security.
type SecuritiesCertificate1 struct {

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	Number *Max35Text `xml:"Nb"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (s *SecuritiesCertificate1) SetNumber(value string) {
	s.Number = (*Max35Text)(&value)
}

func (s *SecuritiesCertificate1) SetIssuer(value string) {
	s.Issuer = (*Max35Text)(&value)
}

func (s *SecuritiesCertificate1) SetSchemeName(value string) {
	s.SchemeName = (*Max35Text)(&value)
}

// Physical representation of a security.
type SecuritiesCertificate3 struct {

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	Number *Max35Text `xml:"Nb"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (s *SecuritiesCertificate3) SetNumber(value string) {
	s.Number = (*Max35Text)(&value)
}

func (s *SecuritiesCertificate3) SetIssuer(value string) {
	s.Issuer = (*Max35Text)(&value)
}

func (s *SecuritiesCertificate3) SetSchemeName(value string) {
	s.SchemeName = (*Max35Text)(&value)
}

// Physical representation of a security.
type SecuritiesCertificate4 struct {

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	Number *Max35Text `xml:"Nb"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (s *SecuritiesCertificate4) SetNumber(value string) {
	s.Number = (*Max35Text)(&value)
}

func (s *SecuritiesCertificate4) SetIssuer(value string) {
	s.Issuer = (*Max35Text)(&value)
}

func (s *SecuritiesCertificate4) SetSchemeName(value string) {
	s.SchemeName = (*Max35Text)(&value)
}

// Physical representation of a security.
type SecuritiesCertificate5 struct {

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	Number *RestrictedFINXMax30Text `xml:"Nb"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	Issuer *Max4AlphaNumericText `xml:"Issr,omitempty"`

	// Short textual description of the scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (s *SecuritiesCertificate5) SetNumber(value string) {
	s.Number = (*RestrictedFINXMax30Text)(&value)
}

func (s *SecuritiesCertificate5) SetIssuer(value string) {
	s.Issuer = (*Max4AlphaNumericText)(&value)
}

func (s *SecuritiesCertificate5) SetSchemeName(value string) {
	s.SchemeName = (*Max4AlphaNumericText)(&value)
}
