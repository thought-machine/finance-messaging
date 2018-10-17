package iso20022

// Information related to registration of securities.
type RegistrationParameters1 struct {

	// Identification assigned to a deposit.
	CertificationIdentification *Max35Text `xml:"CertfctnId,omitempty"`

	// Date/time at which the certificates in the deposit were validated by the agent.
	CertificationDateTime *DateAndDateTimeChoice `xml:"CertfctnDtTm,omitempty"`

	// Account at the registrar where financial instruments are registered.
	RegistrarAccount *Max35Text `xml:"RegarAcct,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate1 `xml:"CertNb,omitempty"`
}

func (r *RegistrationParameters1) SetCertificationIdentification(value string) {
	r.CertificationIdentification = (*Max35Text)(&value)
}

func (r *RegistrationParameters1) AddCertificationDateTime() *DateAndDateTimeChoice {
	r.CertificationDateTime = new(DateAndDateTimeChoice)
	return r.CertificationDateTime
}

func (r *RegistrationParameters1) SetRegistrarAccount(value string) {
	r.RegistrarAccount = (*Max35Text)(&value)
}

func (r *RegistrationParameters1) AddCertificateNumber() *SecuritiesCertificate1 {
	newValue := new(SecuritiesCertificate1)
	r.CertificateNumber = append(r.CertificateNumber, newValue)
	return newValue
}

// Information related to registration of securities.
type RegistrationParameters3 struct {

	// Identification assigned to a deposit.
	CertificationIdentification *Max35Text `xml:"CertfctnId,omitempty"`

	// Date/time at which the certificates in the deposit were validated by the agent.
	CertificationDateTime *DateAndDateTime1Choice `xml:"CertfctnDtTm,omitempty"`

	// Account at the registrar where financial instruments are registered.
	RegistrarAccount *Max35Text `xml:"RegarAcct,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate3 `xml:"CertNb,omitempty"`
}

func (r *RegistrationParameters3) SetCertificationIdentification(value string) {
	r.CertificationIdentification = (*Max35Text)(&value)
}

func (r *RegistrationParameters3) AddCertificationDateTime() *DateAndDateTime1Choice {
	r.CertificationDateTime = new(DateAndDateTime1Choice)
	return r.CertificationDateTime
}

func (r *RegistrationParameters3) SetRegistrarAccount(value string) {
	r.RegistrarAccount = (*Max35Text)(&value)
}

func (r *RegistrationParameters3) AddCertificateNumber() *SecuritiesCertificate3 {
	newValue := new(SecuritiesCertificate3)
	r.CertificateNumber = append(r.CertificateNumber, newValue)
	return newValue
}

// Information related to registration of securities.
type RegistrationParameters4 struct {

	// Identification assigned to a deposit.
	CertificationIdentification *Max35Text `xml:"CertfctnId,omitempty"`

	// Date/time at which the certificates in the deposit were validated by the agent.
	CertificationDateTime *DateAndDateTimeChoice `xml:"CertfctnDtTm,omitempty"`

	// Account at the registrar where financial instruments are registered.
	RegistrarAccount *Max35Text `xml:"RegarAcct,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate4 `xml:"CertNb,omitempty"`
}

func (r *RegistrationParameters4) SetCertificationIdentification(value string) {
	r.CertificationIdentification = (*Max35Text)(&value)
}

func (r *RegistrationParameters4) AddCertificationDateTime() *DateAndDateTimeChoice {
	r.CertificationDateTime = new(DateAndDateTimeChoice)
	return r.CertificationDateTime
}

func (r *RegistrationParameters4) SetRegistrarAccount(value string) {
	r.RegistrarAccount = (*Max35Text)(&value)
}

func (r *RegistrationParameters4) AddCertificateNumber() *SecuritiesCertificate4 {
	newValue := new(SecuritiesCertificate4)
	r.CertificateNumber = append(r.CertificateNumber, newValue)
	return newValue
}

// Information related to registration of securities.
type RegistrationParameters5 struct {

	// Identification assigned to a deposit.
	CertificationIdentification *RestrictedFINXMax16Text `xml:"CertfctnId,omitempty"`

	// Date/time at which the certificates in the deposit were validated by the agent.
	CertificationDateTime *DateAndDateTimeChoice `xml:"CertfctnDtTm,omitempty"`

	// Account at the registrar where financial instruments are registered.
	RegistrarAccount *RestrictedFINXMax35Text `xml:"RegarAcct,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate5 `xml:"CertNb,omitempty"`
}

func (r *RegistrationParameters5) SetCertificationIdentification(value string) {
	r.CertificationIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *RegistrationParameters5) AddCertificationDateTime() *DateAndDateTimeChoice {
	r.CertificationDateTime = new(DateAndDateTimeChoice)
	return r.CertificationDateTime
}

func (r *RegistrationParameters5) SetRegistrarAccount(value string) {
	r.RegistrarAccount = (*RestrictedFINXMax35Text)(&value)
}

func (r *RegistrationParameters5) AddCertificateNumber() *SecuritiesCertificate5 {
	newValue := new(SecuritiesCertificate5)
	r.CertificateNumber = append(r.CertificateNumber, newValue)
	return newValue
}
