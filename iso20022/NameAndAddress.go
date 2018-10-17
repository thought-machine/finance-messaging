package iso20022

// Information that locates and identifies a specific address, as defined by postal services.
type NameAndAddress1 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Postal address of a party.
	Address *LongPostalAddress1Choice `xml:"Adr"`
}

func (n *NameAndAddress1) SetName(value string) {
	n.Name = (*Max35Text)(&value)
}

func (n *NameAndAddress1) AddAddress() *LongPostalAddress1Choice {
	n.Address = new(LongPostalAddress1Choice)
	return n.Address
}

// Information that locates and identifies a party.
type NameAndAddress10 struct {

	// Name by which a party is known and is usually used to identify that identity.
	Name *Max140Text `xml:"Nm"`

	// Postal address of a party.
	Address *PostalAddress6 `xml:"Adr"`
}

func (n *NameAndAddress10) SetName(value string) {
	n.Name = (*Max140Text)(&value)
}

func (n *NameAndAddress10) AddAddress() *PostalAddress6 {
	n.Address = new(PostalAddress6)
	return n.Address
}

// Information that locates and identifies a party.
type NameAndAddress12 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *RestrictedFINXMax140Text `xml:"Nm"`
}

func (n *NameAndAddress12) SetName(value string) {
	n.Name = (*RestrictedFINXMax140Text)(&value)
}

// Information that locates and identifies a party.
type NameAndAddress13 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Postal address of a party.
	Address *PostalAddress8 `xml:"Adr,omitempty"`
}

func (n *NameAndAddress13) SetName(value string) {
	n.Name = (*Max350Text)(&value)
}

func (n *NameAndAddress13) AddAddress() *PostalAddress8 {
	n.Address = new(PostalAddress8)
	return n.Address
}

// Information that locates and identifies a party.
type NameAndAddress15 struct {

	// Name by which the party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Postal address of the party.
	PostalAddress *PostalAddress21 `xml:"PstlAdr,omitempty"`
}

func (n *NameAndAddress15) SetName(value string) {
	n.Name = (*Max350Text)(&value)
}

func (n *NameAndAddress15) AddPostalAddress() *PostalAddress21 {
	n.PostalAddress = new(PostalAddress21)
	return n.PostalAddress
}

// Entity involved in an activity.
type NameAndAddress2 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm"`

	// Information that locates and identifies a specific address, as defined by postal services.
	Address *LongPostalAddress1Choice `xml:"Adr,omitempty"`
}

func (n *NameAndAddress2) SetName(value string) {
	n.Name = (*Max35Text)(&value)
}

func (n *NameAndAddress2) AddAddress() *LongPostalAddress1Choice {
	n.Address = new(LongPostalAddress1Choice)
	return n.Address
}

// Information that locates and identifies a party.
type NameAndAddress3 struct {

	// Name by which a party is known and is usually used to identify that identity.
	Name *Max70Text `xml:"Nm"`

	// Postal address of a party.
	Address *PostalAddress1 `xml:"Adr"`
}

func (n *NameAndAddress3) SetName(value string) {
	n.Name = (*Max70Text)(&value)
}

func (n *NameAndAddress3) AddAddress() *PostalAddress1 {
	n.Address = new(PostalAddress1)
	return n.Address
}

// Information that locates and identifies a party.
type NameAndAddress4 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Postal address of a party.
	Address *PostalAddress1 `xml:"Adr"`
}

func (n *NameAndAddress4) SetName(value string) {
	n.Name = (*Max350Text)(&value)
}

func (n *NameAndAddress4) AddAddress() *PostalAddress1 {
	n.Address = new(PostalAddress1)
	return n.Address
}

// Information that locates and identifies a party.
type NameAndAddress5 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Postal address of a party.
	Address *PostalAddress1 `xml:"Adr,omitempty"`
}

func (n *NameAndAddress5) SetName(value string) {
	n.Name = (*Max350Text)(&value)
}

func (n *NameAndAddress5) AddAddress() *PostalAddress1 {
	n.Address = new(PostalAddress1)
	return n.Address
}

// Name and address of an institution.
type NameAndAddress6 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm"`

	// Information that locates and identifies a specific address, as defined by postal services.
	Address *PostalAddress2 `xml:"Adr"`
}

func (n *NameAndAddress6) SetName(value string) {
	n.Name = (*Max70Text)(&value)
}

func (n *NameAndAddress6) AddAddress() *PostalAddress2 {
	n.Address = new(PostalAddress2)
	return n.Address
}

// Name and address of an institution.
type NameAndAddress7 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress1 `xml:"PstlAdr"`
}

func (n *NameAndAddress7) SetName(value string) {
	n.Name = (*Max70Text)(&value)
}

func (n *NameAndAddress7) AddPostalAddress() *PostalAddress1 {
	n.PostalAddress = new(PostalAddress1)
	return n.PostalAddress
}

// Information that locates and identifies a party.
type NameAndAddress8 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Postal address of a party.
	Address *PostalAddress1 `xml:"Adr,omitempty"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	AlternativeIdentifier []*Max35Text `xml:"AltrntvIdr,omitempty"`
}

func (n *NameAndAddress8) SetName(value string) {
	n.Name = (*Max350Text)(&value)
}

func (n *NameAndAddress8) AddAddress() *PostalAddress1 {
	n.Address = new(PostalAddress1)
	return n.Address
}

func (n *NameAndAddress8) AddAlternativeIdentifier(value string) {
	n.AlternativeIdentifier = append(n.AlternativeIdentifier, (*Max35Text)(&value))
}

// Information that locates and identifies a party.
type NameAndAddress9 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Postal address of a party.
	Address *LongPostalAddress2Choice `xml:"Adr,omitempty"`
}

func (n *NameAndAddress9) SetName(value string) {
	n.Name = (*Max350Text)(&value)
}

func (n *NameAndAddress9) AddAddress() *LongPostalAddress2Choice {
	n.Address = new(LongPostalAddress2Choice)
	return n.Address
}
