package iso20022

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification1 struct {

	// Identification assigned by an institution.
	Identification *Max35Text `xml:"Id"`

	// Name of the identification scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification1) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification1) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

func (g *GenericIdentification1) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

// Information related to the identification of an individual person.
type GenericIdentification10 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identification.
	IdentificationType *PersonIdentificationType1Code `xml:"IdTp"`

	// Specifies the nature of the identification.
	ExtendedIdentificationType *Extended350Code `xml:"XtndedIdTp"`
}

func (g *GenericIdentification10) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification10) SetIdentificationType(value string) {
	g.IdentificationType = (*PersonIdentificationType1Code)(&value)
}

func (g *GenericIdentification10) SetExtendedIdentificationType(value string) {
	g.ExtendedIdentificationType = (*Extended350Code)(&value)
}

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification11 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identifier.
	IdentificationType *PersonIdentificationType2Code `xml:"IdTp"`

	// Specifies a type of individual identification.
	ExtendedIdentificationType *Extended350Code `xml:"XtndedIdTp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification11) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification11) SetIdentificationType(value string) {
	g.IdentificationType = (*PersonIdentificationType2Code)(&value)
}

func (g *GenericIdentification11) SetExtendedIdentificationType(value string) {
	g.ExtendedIdentificationType = (*Extended350Code)(&value)
}

func (g *GenericIdentification11) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification13 struct {

	// Identification assigned by an institution.
	Identification *Max4AlphaNumericText `xml:"Id"`

	// Name of the identification scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`
}

func (g *GenericIdentification13) SetIdentification(value string) {
	g.Identification = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification13) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

func (g *GenericIdentification13) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

// Proprietary information related to a balance.
type GenericIdentification144 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`

	// Value of the balance.
	Balance *RestrictedFINDecimalNumber `xml:"Bal"`
}

func (g *GenericIdentification144) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification144) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification144) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification144) SetBalance(value string) {
	g.Balance = (*RestrictedFINDecimalNumber)(&value)
}

// Specifies a generic identification.
type GenericIdentification16 struct {

	// The identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identifier.
	IdentificationType *PersonIdentificationType3Choice `xml:"IdTp"`

	// Party that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification16) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification16) AddIdentificationType() *PersonIdentificationType3Choice {
	g.IdentificationType = new(PersonIdentificationType3Choice)
	return g.IdentificationType
}

func (g *GenericIdentification16) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

// Identification using a proprietary scheme.
type GenericIdentification163 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Max35Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification163) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification163) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification163) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification164 struct {

	// Name or number assigned by an entity to enable recognition of that entity, for example, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identifier.
	IdentificationType *OtherIdentification3Choice `xml:"IdTp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification164) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification164) AddIdentificationType() *OtherIdentification3Choice {
	g.IdentificationType = new(OtherIdentification3Choice)
	return g.IdentificationType
}

func (g *GenericIdentification164) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification18 struct {

	// Identification assigned by an institution.
	Identification *RestrictedFINXMax30Text `xml:"Id"`

	// Name of the identification scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr,omitempty"`
}

func (g *GenericIdentification18) SetIdentification(value string) {
	g.Identification = (*RestrictedFINXMax30Text)(&value)
}

func (g *GenericIdentification18) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification18) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification19 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Max35Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification19) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification19) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification19) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification2 struct {

	// Name of the identification scheme.
	SchemeName *Max35Text `xml:"SchmeNm"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification2) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

func (g *GenericIdentification2) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification20 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification20) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification20) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification20) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification21 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Type *GenericIdentification20 `xml:"Tp"`

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id,omitempty"`
}

func (g *GenericIdentification21) AddType() *GenericIdentification20 {
	g.Type = new(GenericIdentification20)
	return g.Type
}

func (g *GenericIdentification21) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

// Proprietary information related to a balance.
type GenericIdentification22 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`

	// Value of the balance.
	Balance *DecimalNumber `xml:"Bal"`
}

func (g *GenericIdentification22) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification22) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification22) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

func (g *GenericIdentification22) SetBalance(value string) {
	g.Balance = (*DecimalNumber)(&value)
}

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification25 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification25) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification25) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification25) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification27 struct {

	// Identification assigned by an institution.
	Identification *Max4AlphaNumericText `xml:"Id"`

	// Name of the identification scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr"`
}

func (g *GenericIdentification27) SetIdentification(value string) {
	g.Identification = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification27) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification27) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification29 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Max35Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification29) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification29) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification29) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification3 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification3) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification3) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification30 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification30) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification30) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification30) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification31 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType3Code `xml:"Tp"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType4Code `xml:"Issr,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification31) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification31) SetType(value string) {
	g.Type = (*PartyType3Code)(&value)
}

func (g *GenericIdentification31) SetIssuer(value string) {
	g.Issuer = (*PartyType4Code)(&value)
}

func (g *GenericIdentification31) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification32 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType3Code `xml:"Tp,omitempty"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType4Code `xml:"Issr,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification32) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification32) SetType(value string) {
	g.Type = (*PartyType3Code)(&value)
}

func (g *GenericIdentification32) SetIssuer(value string) {
	g.Issuer = (*PartyType4Code)(&value)
}

func (g *GenericIdentification32) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification33 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Type of identified entity.
	Type *PartyType3Code `xml:"Tp"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType4Code `xml:"Issr,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification33) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification33) SetType(value string) {
	g.Type = (*PartyType3Code)(&value)
}

func (g *GenericIdentification33) SetIssuer(value string) {
	g.Issuer = (*PartyType4Code)(&value)
}

func (g *GenericIdentification33) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification35 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType5Code `xml:"Tp,omitempty"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType6Code `xml:"Issr,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification35) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification35) SetType(value string) {
	g.Type = (*PartyType5Code)(&value)
}

func (g *GenericIdentification35) SetIssuer(value string) {
	g.Issuer = (*PartyType6Code)(&value)
}

func (g *GenericIdentification35) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Identification using a proprietary scheme.
type GenericIdentification36 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Max35Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification36) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification36) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification36) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification37 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Max35Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification37) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification37) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification38 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification38) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification38) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification38) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification39 struct {

	// Proprietary information issued by the data source scheme issuer.
	Identification *RestrictedFINMax30Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *RestrictedFINMax8Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification39) SetIdentification(value string) {
	g.Identification = (*RestrictedFINMax30Text)(&value)
}

func (g *GenericIdentification39) SetIssuer(value string) {
	g.Issuer = (*RestrictedFINMax8Text)(&value)
}

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification4 struct {

	// Identifier issued to a person for which no specific identifier has been defined.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identifier.
	// Usage: IdentificationType is used to specify what kind of identifier is used. It should be used in case the identifier is different from the identifiers listed in the pre-defined identifier list.
	IdentificationType *Max35Text `xml:"IdTp"`
}

func (g *GenericIdentification4) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification4) SetIdentificationType(value string) {
	g.IdentificationType = (*Max35Text)(&value)
}

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification40 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification40) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification40) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification40) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification41 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Max35Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification41) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification41) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification41) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

// Information related to the identification of an individual person.
type GenericIdentification44 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identification.
	Type *OtherIdentification1Choice `xml:"Tp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Date at which the identification was issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date at which the identification expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`
}

func (g *GenericIdentification44) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification44) AddType() *OtherIdentification1Choice {
	g.Type = new(OtherIdentification1Choice)
	return g.Type
}

func (g *GenericIdentification44) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification44) SetIssueDate(value string) {
	g.IssueDate = (*ISODate)(&value)
}

func (g *GenericIdentification44) SetExpiryDate(value string) {
	g.ExpiryDate = (*ISODate)(&value)
}

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification46 struct {

	// Name or number assigned by an entity to enable recognition of that entity, for example, national registration identification number.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identifier.
	Type *OtherIdentification1Choice `xml:"Tp"`
}

func (g *GenericIdentification46) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification46) AddType() *OtherIdentification1Choice {
	g.Type = new(OtherIdentification1Choice)
	return g.Type
}

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification47 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification47) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification47) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification47) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}

// Information related to an identification.
type GenericIdentification48 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Max35Text `xml:"Id"`

	// Version of the identification.
	Version *Max35Text `xml:"Vrsn"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`
}

func (g *GenericIdentification48) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification48) SetVersion(value string) {
	g.Version = (*Max35Text)(&value)
}

func (g *GenericIdentification48) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification49 struct {

	// Identifier issued to a person or an institution for which no other specific identifier has been defined.
	Identification *Max35Text `xml:"Id"`

	// Type of identifier.
	IdentificationType *Max35Text `xml:"IdTp"`
}

func (g *GenericIdentification49) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification49) SetIdentificationType(value string) {
	g.IdentificationType = (*Max35Text)(&value)
}

// Information expressed in a proprietary manner.
type GenericIdentification5 struct {

	// Entity that assigns the identification.
	Issuer *Max8Text `xml:"Issr"`

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Information *Exact4AlphaNumericText `xml:"Inf"`

	// Additional information.
	Narrative *Max35Text `xml:"Nrrtv,omitempty"`
}

func (g *GenericIdentification5) SetIssuer(value string) {
	g.Issuer = (*Max8Text)(&value)
}

func (g *GenericIdentification5) SetInformation(value string) {
	g.Information = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification5) SetNarrative(value string) {
	g.Narrative = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification53 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType3Code `xml:"Tp,omitempty"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType4Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification53) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification53) SetType(value string) {
	g.Type = (*PartyType3Code)(&value)
}

func (g *GenericIdentification53) SetIssuer(value string) {
	g.Issuer = (*PartyType4Code)(&value)
}

func (g *GenericIdentification53) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification53) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Information related to the identification of an individual person.
type GenericIdentification55 struct {

	// Name or number assigned by an entity to enable recognition of that entity, for example, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identification.
	Type *OtherIdentification2Choice `xml:"Tp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Date at which the identification was issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date at which the identification expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Country that issued the identification document.
	IssuerCountry *CountryCode `xml:"IssrCtry,omitempty"`
}

func (g *GenericIdentification55) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification55) AddType() *OtherIdentification2Choice {
	g.Type = new(OtherIdentification2Choice)
	return g.Type
}

func (g *GenericIdentification55) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification55) SetIssueDate(value string) {
	g.IssueDate = (*ISODate)(&value)
}

func (g *GenericIdentification55) SetExpiryDate(value string) {
	g.ExpiryDate = (*ISODate)(&value)
}

func (g *GenericIdentification55) SetIssuerCountry(value string) {
	g.IssuerCountry = (*CountryCode)(&value)
}

// Proprietary information related to a balance.
type GenericIdentification56 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`

	// Value of the balance.
	Balance *DecimalNumber `xml:"Bal"`
}

func (g *GenericIdentification56) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification56) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification56) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

func (g *GenericIdentification56) SetBalance(value string) {
	g.Balance = (*DecimalNumber)(&value)
}

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification58 struct {

	// Name or number assigned by an entity to enable recognition of that entity.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Type *GenericIdentification40 `xml:"Tp"`
}

func (g *GenericIdentification58) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification58) AddType() *GenericIdentification40 {
	g.Type = new(GenericIdentification40)
	return g.Type
}

// Balance expressed with a data source scheme.
type GenericIdentification6 struct {

	// Entity that assigns the identification.
	Issuer *Max8Text `xml:"Issr"`

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Information *Exact4AlphaNumericText `xml:"Inf"`

	// Value of the balance.
	Balance *Number `xml:"Bal"`
}

func (g *GenericIdentification6) SetIssuer(value string) {
	g.Issuer = (*Max8Text)(&value)
}

func (g *GenericIdentification6) SetInformation(value string) {
	g.Information = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification6) SetBalance(value string) {
	g.Balance = (*Number)(&value)
}

// Information expressed in a proprietary manner.
type GenericIdentification7 struct {

	// Entity that assigns the identification.
	Issuer *Max8Text `xml:"Issr"`

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Information *Max35Text `xml:"Inf"`
}

func (g *GenericIdentification7) SetIssuer(value string) {
	g.Issuer = (*Max8Text)(&value)
}

func (g *GenericIdentification7) SetInformation(value string) {
	g.Information = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification70 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Type of identified entity.
	Type *PartyType3Code `xml:"Tp"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType4Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification70) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification70) SetType(value string) {
	g.Type = (*PartyType3Code)(&value)
}

func (g *GenericIdentification70) SetIssuer(value string) {
	g.Issuer = (*PartyType4Code)(&value)
}

func (g *GenericIdentification70) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification70) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification71 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType5Code `xml:"Tp,omitempty"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType6Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification71) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification71) SetType(value string) {
	g.Type = (*PartyType5Code)(&value)
}

func (g *GenericIdentification71) SetIssuer(value string) {
	g.Issuer = (*PartyType6Code)(&value)
}

func (g *GenericIdentification71) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification71) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification72 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType6Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification72) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification72) SetIssuer(value string) {
	g.Issuer = (*PartyType6Code)(&value)
}

func (g *GenericIdentification72) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification72) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification73 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType9Code `xml:"Tp,omitempty"`

	// Entity assigning the identification , for example merchant, acceptor, acquirer or card issuer.
	Issuer *PartyType9Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification73) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification73) SetType(value string) {
	g.Type = (*PartyType9Code)(&value)
}

func (g *GenericIdentification73) SetIssuer(value string) {
	g.Issuer = (*PartyType9Code)(&value)
}

func (g *GenericIdentification73) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification73) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification74 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType10Code `xml:"Tp"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType10Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3)
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification74) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification74) SetType(value string) {
	g.Type = (*PartyType10Code)(&value)
}

func (g *GenericIdentification74) SetIssuer(value string) {
	g.Issuer = (*PartyType10Code)(&value)
}

func (g *GenericIdentification74) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification74) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification75 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Type of identified entity.
	Type *PartyType11Code `xml:"Tp"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType9Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification75) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification75) SetType(value string) {
	g.Type = (*PartyType11Code)(&value)
}

func (g *GenericIdentification75) SetIssuer(value string) {
	g.Issuer = (*PartyType9Code)(&value)
}

func (g *GenericIdentification75) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification75) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification76 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType3Code `xml:"Tp"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType4Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3)
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification76) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification76) SetType(value string) {
	g.Type = (*PartyType3Code)(&value)
}

func (g *GenericIdentification76) SetIssuer(value string) {
	g.Issuer = (*PartyType4Code)(&value)
}

func (g *GenericIdentification76) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification76) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification77 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType12Code `xml:"Tp"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType12Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3)
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification77) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification77) SetType(value string) {
	g.Type = (*PartyType12Code)(&value)
}

func (g *GenericIdentification77) SetIssuer(value string) {
	g.Issuer = (*PartyType12Code)(&value)
}

func (g *GenericIdentification77) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification77) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification78 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Type *GenericIdentification30 `xml:"Tp"`

	// Name or number assigned by an entity to enable recognition of that entity, for example, account identifier.
	Identification *Max35Text `xml:"Id,omitempty"`
}

func (g *GenericIdentification78) AddType() *GenericIdentification30 {
	g.Type = new(GenericIdentification30)
	return g.Type
}

func (g *GenericIdentification78) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

// Identification using a proprietary scheme.
type GenericIdentification79 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *RestrictedFINXMax34Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification79) SetIdentification(value string) {
	g.Identification = (*RestrictedFINXMax34Text)(&value)
}

func (g *GenericIdentification79) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification79) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification8 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identifier.
	IdentificationType *IdentificationType1 `xml:"IdTp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification8) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification8) AddIdentificationType() *IdentificationType1 {
	g.IdentificationType = new(IdentificationType1)
	return g.IdentificationType
}

func (g *GenericIdentification8) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification80 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Type *GenericIdentification30 `xml:"Tp"`

	// Name or number assigned by an entity to enable recognition of that entity.
	Identification *Max35Text `xml:"Id,omitempty"`
}

func (g *GenericIdentification80) AddType() *GenericIdentification30 {
	g.Type = new(GenericIdentification30)
	return g.Type
}

func (g *GenericIdentification80) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

// Information related to the identification of a party.
type GenericIdentification81 struct {

	// Identification of a party, such as a tax or social security identifier.
	Identification *Max35Text `xml:"Id"`

	// Type of identification.
	IdentificationType *OtherIdentification3Choice `xml:"IdTp"`
}

func (g *GenericIdentification81) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification81) AddIdentificationType() *OtherIdentification3Choice {
	g.IdentificationType = new(OtherIdentification3Choice)
	return g.IdentificationType
}

// Information related to the identification of a party.
type GenericIdentification82 struct {

	// Name or number assigned by an entity to enable recognition of that entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identification.
	Type *OtherIdentification3Choice `xml:"Tp"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Date at which the identification was issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date at which the identification expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Name of the state, county or country sub-division that issued the identification document.
	State *Max70Text `xml:"Stat,omitempty"`

	// Country that issued the identification document.
	IssuerCountry *CountryCode `xml:"IssrCtry,omitempty"`
}

func (g *GenericIdentification82) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification82) AddType() *OtherIdentification3Choice {
	g.Type = new(OtherIdentification3Choice)
	return g.Type
}

func (g *GenericIdentification82) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification82) SetIssueDate(value string) {
	g.IssueDate = (*ISODate)(&value)
}

func (g *GenericIdentification82) SetExpiryDate(value string) {
	g.ExpiryDate = (*ISODate)(&value)
}

func (g *GenericIdentification82) SetState(value string) {
	g.State = (*Max70Text)(&value)
}

func (g *GenericIdentification82) SetIssuerCountry(value string) {
	g.IssuerCountry = (*CountryCode)(&value)
}

// Identification using a proprietary scheme.
type GenericIdentification84 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *RestrictedFINXMax34Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification84) SetIdentification(value string) {
	g.Identification = (*RestrictedFINXMax34Text)(&value)
}

func (g *GenericIdentification84) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification84) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification85 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Type *GenericIdentification47 `xml:"Tp"`

	// Name or number assigned by an entity to enable recognition of that entity.
	Identification *RestrictedFINXMax30Text `xml:"Id,omitempty"`
}

func (g *GenericIdentification85) AddType() *GenericIdentification47 {
	g.Type = new(GenericIdentification47)
	return g.Type
}

func (g *GenericIdentification85) SetIdentification(value string) {
	g.Identification = (*RestrictedFINXMax30Text)(&value)
}

// Identification using a proprietary scheme.
type GenericIdentification86 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *RestrictedFINXMax30Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification86) SetIdentification(value string) {
	g.Identification = (*RestrictedFINXMax30Text)(&value)
}

func (g *GenericIdentification86) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification86) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification89 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Type *GenericIdentification47 `xml:"Tp"`

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *RestrictedFINXMax30Text `xml:"Id,omitempty"`
}

func (g *GenericIdentification89) AddType() *GenericIdentification47 {
	g.Type = new(GenericIdentification47)
	return g.Type
}

func (g *GenericIdentification89) SetIdentification(value string) {
	g.Identification = (*RestrictedFINXMax30Text)(&value)
}

// Information related to the identification of an individual person.
type GenericIdentification9 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identification.
	IdentificationType *PersonIdentificationType1Code `xml:"IdTp"`

	// Specifies the nature of the identification.
	ExtendedIdentificationType *Extended350Code `xml:"XtndedIdTp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Date at which the identification was issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date at which the identification expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`
}

func (g *GenericIdentification9) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification9) SetIdentificationType(value string) {
	g.IdentificationType = (*PersonIdentificationType1Code)(&value)
}

func (g *GenericIdentification9) SetExtendedIdentificationType(value string) {
	g.ExtendedIdentificationType = (*Extended350Code)(&value)
}

func (g *GenericIdentification9) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification9) SetIssueDate(value string) {
	g.IssueDate = (*ISODate)(&value)
}

func (g *GenericIdentification9) SetExpiryDate(value string) {
	g.ExpiryDate = (*ISODate)(&value)
}

// Identification of an entity.
type GenericIdentification90 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Type of identified entity.
	Type *PartyType14Code `xml:"Tp"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType4Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification90) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification90) SetType(value string) {
	g.Type = (*PartyType14Code)(&value)
}

func (g *GenericIdentification90) SetIssuer(value string) {
	g.Issuer = (*PartyType4Code)(&value)
}

func (g *GenericIdentification90) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification90) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

// Identification of an entity.
type GenericIdentification92 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType5Code `xml:"Tp,omitempty"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType6Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Access information to reach the target host.
	RemoteAccess *NetworkParameters5 `xml:"RmotAccs,omitempty"`
}

func (g *GenericIdentification92) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification92) SetType(value string) {
	g.Type = (*PartyType5Code)(&value)
}

func (g *GenericIdentification92) SetIssuer(value string) {
	g.Issuer = (*PartyType6Code)(&value)
}

func (g *GenericIdentification92) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification92) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

func (g *GenericIdentification92) AddRemoteAccess() *NetworkParameters5 {
	g.RemoteAccess = new(NetworkParameters5)
	return g.RemoteAccess
}

// Identification of an entity.
type GenericIdentification93 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType6Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Access information to reach the target host.
	RemoteAccess *NetworkParameters5 `xml:"RmotAccs,omitempty"`
}

func (g *GenericIdentification93) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification93) SetIssuer(value string) {
	g.Issuer = (*PartyType6Code)(&value)
}

func (g *GenericIdentification93) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification93) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

func (g *GenericIdentification93) AddRemoteAccess() *NetworkParameters5 {
	g.RemoteAccess = new(NetworkParameters5)
	return g.RemoteAccess
}

// Identification of an entity.
type GenericIdentification94 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType3Code `xml:"Tp,omitempty"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType4Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Access information to reach the target host.
	RemoteAccess *NetworkParameters5 `xml:"RmotAccs,omitempty"`
}

func (g *GenericIdentification94) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification94) SetType(value string) {
	g.Type = (*PartyType3Code)(&value)
}

func (g *GenericIdentification94) SetIssuer(value string) {
	g.Issuer = (*PartyType4Code)(&value)
}

func (g *GenericIdentification94) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification94) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

func (g *GenericIdentification94) AddRemoteAccess() *NetworkParameters5 {
	g.RemoteAccess = new(NetworkParameters5)
	return g.RemoteAccess
}
