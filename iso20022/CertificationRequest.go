package iso20022

// Certification request PKCS#10 (Public Key Certificate Standard 10) for creation or renewal of an X.509 certificate.
type CertificationRequest1 struct {

	// Information of the certificate to create.
	CertificateRequestInformation *CertificationRequest2 `xml:"CertReqInf"`

	// Identification of the key.
	KeyIdentification *Max140Text `xml:"KeyId,omitempty"`

	// Version of the key.
	KeyVersion *Max140Text `xml:"KeyVrsn,omitempty"`
}

func (c *CertificationRequest1) AddCertificateRequestInformation() *CertificationRequest2 {
	c.CertificateRequestInformation = new(CertificationRequest2)
	return c.CertificateRequestInformation
}

func (c *CertificationRequest1) SetKeyIdentification(value string) {
	c.KeyIdentification = (*Max140Text)(&value)
}

func (c *CertificationRequest1) SetKeyVersion(value string) {
	c.KeyVersion = (*Max140Text)(&value)
}

// Information of the certificate to create.
type CertificationRequest2 struct {

	// Version of the certificate request information data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Distinguished name of the certificate subject, the entity whose public key is to be certified.
	SubjectName *CertificateIssuer1 `xml:"SbjtNm,omitempty"`

	// Information about the public key being certified.
	SubjectPublicKeyInformation *PublicRSAKey2 `xml:"SbjtPblcKeyInf"`

	// Attribute of the certificate service to be put in the certificate extensions, or to be used for the request.
	Attribute []*RelativeDistinguishedName2 `xml:"Attr"`
}

func (c *CertificationRequest2) SetVersion(value string) {
	c.Version = (*Number)(&value)
}

func (c *CertificationRequest2) AddSubjectName() *CertificateIssuer1 {
	c.SubjectName = new(CertificateIssuer1)
	return c.SubjectName
}

func (c *CertificationRequest2) AddSubjectPublicKeyInformation() *PublicRSAKey2 {
	c.SubjectPublicKeyInformation = new(PublicRSAKey2)
	return c.SubjectPublicKeyInformation
}

func (c *CertificationRequest2) AddAttribute() *RelativeDistinguishedName2 {
	newValue := new(RelativeDistinguishedName2)
	c.Attribute = append(c.Attribute, newValue)
	return newValue
}
