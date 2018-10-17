package iso20022

// Value of the public component of a RSA key.
type PublicRSAKey1 struct {

	// Modulus of the RSA key.
	Modulus *Max5000Binary `xml:"Mdlus"`

	// Public exponent of the RSA key.
	Exponent *Max5000Binary `xml:"Expnt"`
}

func (p *PublicRSAKey1) SetModulus(value string) {
	p.Modulus = (*Max5000Binary)(&value)
}

func (p *PublicRSAKey1) SetExponent(value string) {
	p.Exponent = (*Max5000Binary)(&value)
}

// Value of the public component of a RSA key.
type PublicRSAKey2 struct {

	// Asymmetric cryptographic algorithm.
	Algorithm *Algorithm7Code `xml:"Algo,omitempty"`

	// Public key value.
	PublicKeyValue *PublicRSAKey1 `xml:"PblcKeyVal"`
}

func (p *PublicRSAKey2) SetAlgorithm(value string) {
	p.Algorithm = (*Algorithm7Code)(&value)
}

func (p *PublicRSAKey2) AddPublicKeyValue() *PublicRSAKey1 {
	p.PublicKeyValue = new(PublicRSAKey1)
	return p.PublicKeyValue
}
