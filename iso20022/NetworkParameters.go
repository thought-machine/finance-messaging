package iso20022

// Configuration parameters to communicate with a host.
type NetworkParameters1 struct {

	// IP address or host name of the primary host.
	PrimaryAddress *Max35Text `xml:"PmryAdr"`

	// Port number of the primary host.
	PrimaryPortNumber *Number `xml:"PmryPortNb"`

	// IP address or host name of the secondary host.
	SecondaryAddress *Max35Text `xml:"ScndryAdr,omitempty"`

	// Port number of the secondary host.
	SecondaryPortNumber *Number `xml:"ScndryPortNb,omitempty"`

	// User name identifying the client.
	UserName *Max35Text `xml:"UsrNm,omitempty"`

	// Password authenticating the client.
	AccessCode *Max35Text `xml:"AccsCd,omitempty"`

	// Client certificate chain.
	ClientCertificate *Max3000Binary `xml:"ClntCert,omitempty"`
}

func (n *NetworkParameters1) SetPrimaryAddress(value string) {
	n.PrimaryAddress = (*Max35Text)(&value)
}

func (n *NetworkParameters1) SetPrimaryPortNumber(value string) {
	n.PrimaryPortNumber = (*Number)(&value)
}

func (n *NetworkParameters1) SetSecondaryAddress(value string) {
	n.SecondaryAddress = (*Max35Text)(&value)
}

func (n *NetworkParameters1) SetSecondaryPortNumber(value string) {
	n.SecondaryPortNumber = (*Number)(&value)
}

func (n *NetworkParameters1) SetUserName(value string) {
	n.UserName = (*Max35Text)(&value)
}

func (n *NetworkParameters1) SetAccessCode(value string) {
	n.AccessCode = (*Max35Text)(&value)
}

func (n *NetworkParameters1) SetClientCertificate(value string) {
	n.ClientCertificate = (*Max3000Binary)(&value)
}

// Configuration parameters to communicate with a host.
type NetworkParameters2 struct {

	// IP address or hostname.
	Address *Max35Text `xml:"Adr"`

	// Port number of the server, if the default port number is not used.
	PortNumber *Number `xml:"PortNb,omitempty"`

	// Delay between two contacts of the server..
	Delay *ISOTime `xml:"Dely,omitempty"`
}

func (n *NetworkParameters2) SetAddress(value string) {
	n.Address = (*Max35Text)(&value)
}

func (n *NetworkParameters2) SetPortNumber(value string) {
	n.PortNumber = (*Number)(&value)
}

func (n *NetworkParameters2) SetDelay(value string) {
	n.Delay = (*ISOTime)(&value)
}

// Parameters to communicate with a host.
type NetworkParameters3 struct {

	// Network addresses of the host.
	Address []*NetworkParameters4 `xml:"Adr"`

	// User name identifying the client.
	UserName *Max35Text `xml:"UsrNm,omitempty"`

	// Password authenticating the client.
	AccessCode *Max35Binary `xml:"AccsCd,omitempty"`

	// X.509 Certificate required to authenticate the server.
	ServerCertificate []*Max3000Binary `xml:"SvrCert,omitempty"`

	// Identification of the X.509 Certificate required to authenticate the server, for instance a digest of the certificate.
	ServerCertificateIdentifier []*Max140Binary `xml:"SvrCertIdr,omitempty"`
}

func (n *NetworkParameters3) AddAddress() *NetworkParameters4 {
	newValue := new(NetworkParameters4)
	n.Address = append(n.Address, newValue)
	return newValue
}

func (n *NetworkParameters3) SetUserName(value string) {
	n.UserName = (*Max35Text)(&value)
}

func (n *NetworkParameters3) SetAccessCode(value string) {
	n.AccessCode = (*Max35Binary)(&value)
}

func (n *NetworkParameters3) AddServerCertificate(value string) {
	n.ServerCertificate = append(n.ServerCertificate, (*Max3000Binary)(&value))
}

func (n *NetworkParameters3) AddServerCertificateIdentifier(value string) {
	n.ServerCertificateIdentifier = append(n.ServerCertificateIdentifier, (*Max140Binary)(&value))
}

// Parameters to communicate with a host.
type NetworkParameters4 struct {

	// Type of communication network.
	NetworkType *NetworkType1Code `xml:"NtwkTp"`

	// Value of the address. The value of an internet protocol address contains the IP address or the DNS (Domain Name Server) address, followed by the character ':' and the port number if the default port is not used. The value of a public telephone address contains the phone number with possible prefix and extensions.
	AddressValue *Max70Text `xml:"AdrVal"`
}

func (n *NetworkParameters4) SetNetworkType(value string) {
	n.NetworkType = (*NetworkType1Code)(&value)
}

func (n *NetworkParameters4) SetAddressValue(value string) {
	n.AddressValue = (*Max70Text)(&value)
}

// Parameters to communicate with a host.
type NetworkParameters5 struct {

	// Network addresses of the host.
	Address []*NetworkParameters4 `xml:"Adr"`

	// User name identifying the client.
	UserName *Max35Text `xml:"UsrNm,omitempty"`

	// Password authenticating the client.
	AccessCode *Max35Binary `xml:"AccsCd,omitempty"`

	// X.509 Certificate required to authenticate the server.
	ServerCertificate []*Max10KBinary `xml:"SvrCert,omitempty"`

	// Identification of the X.509 Certificates required to authenticate the server, for instance a digest of the certificate.
	ServerCertificateIdentifier []*Max140Binary `xml:"SvrCertIdr,omitempty"`

	// X.509 Certificate required to authenticate the client.
	ClientCertificate []*Max10KBinary `xml:"ClntCert,omitempty"`

	// Identification of the set of security elements to access the host.
	SecurityProfile *Max35Text `xml:"SctyPrfl,omitempty"`
}

func (n *NetworkParameters5) AddAddress() *NetworkParameters4 {
	newValue := new(NetworkParameters4)
	n.Address = append(n.Address, newValue)
	return newValue
}

func (n *NetworkParameters5) SetUserName(value string) {
	n.UserName = (*Max35Text)(&value)
}

func (n *NetworkParameters5) SetAccessCode(value string) {
	n.AccessCode = (*Max35Binary)(&value)
}

func (n *NetworkParameters5) AddServerCertificate(value string) {
	n.ServerCertificate = append(n.ServerCertificate, (*Max10KBinary)(&value))
}

func (n *NetworkParameters5) AddServerCertificateIdentifier(value string) {
	n.ServerCertificateIdentifier = append(n.ServerCertificateIdentifier, (*Max140Binary)(&value))
}

func (n *NetworkParameters5) AddClientCertificate(value string) {
	n.ClientCertificate = append(n.ClientCertificate, (*Max10KBinary)(&value))
}

func (n *NetworkParameters5) SetSecurityProfile(value string) {
	n.SecurityProfile = (*Max35Text)(&value)
}

// Parameters to communicate with a host.
type NetworkParameters6 struct {

	// Type of proxy.
	Type *NetworkType2Code `xml:"Tp"`

	// Access information to the proxy.
	Access *NetworkParameters5 `xml:"Accs"`
}

func (n *NetworkParameters6) SetType(value string) {
	n.Type = (*NetworkType2Code)(&value)
}

func (n *NetworkParameters6) AddAccess() *NetworkParameters5 {
	n.Access = new(NetworkParameters5)
	return n.Access
}
