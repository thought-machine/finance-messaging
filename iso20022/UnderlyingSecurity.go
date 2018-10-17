package iso20022

// Securitised right for entitlement, for example, equity or bond.
type UnderlyingSecurity1 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId"`
}

func (u *UnderlyingSecurity1) AddSecurityIdentification() *SecurityIdentification11 {
	u.SecurityIdentification = new(SecurityIdentification11)
	return u.SecurityIdentification
}

// Securitised right for entitlement, for example, equity or bond.
type UnderlyingSecurity3 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`
}

func (u *UnderlyingSecurity3) AddSecurityIdentification() *SecurityIdentification14 {
	u.SecurityIdentification = new(SecurityIdentification14)
	return u.SecurityIdentification
}
