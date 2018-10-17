package iso20022

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress1 struct {

	// Identifies the nature of the postal address.
	AddressType *AddressType2Code `xml:"AdrTp,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services, that is presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty"`

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of a building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm,omitempty"`

	// Identifies a subdivision of a country for example, state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PostalAddress1) SetAddressType(value string) {
	p.AddressType = (*AddressType2Code)(&value)
}

func (p *PostalAddress1) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

func (p *PostalAddress1) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress1) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress1) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress1) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress1) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress1) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress11 struct {

	// Identifies the nature of the postal address.
	AddressType *AddressType2Code `xml:"AdrTp,omitempty"`

	// Identification of a division of a large organisation or building.
	Department *Max70Text `xml:"Dept,omitempty"`

	// Identification of a sub-division of a large organisation or building.
	SubDepartment *Max70Text `xml:"SubDept,omitempty"`

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of a building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm,omitempty"`

	// Identifies a subdivision of a country such as state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services, presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty"`

	// Floor or storey within a building.
	Floor *Max70Text `xml:"Flr,omitempty"`

	// Numbered box in a post office, assigned to a person or organisation, where letters are kept until called for.
	PostBox *Max16Text `xml:"PstBx,omitempty"`

	// Name of the building or house.
	BuildingName *Max70Text `xml:"BldgNm,omitempty"`

	// Building room number.
	Room *Max70Text `xml:"Room,omitempty"`
}

func (p *PostalAddress11) SetAddressType(value string) {
	p.AddressType = (*AddressType2Code)(&value)
}

func (p *PostalAddress11) SetDepartment(value string) {
	p.Department = (*Max70Text)(&value)
}

func (p *PostalAddress11) SetSubDepartment(value string) {
	p.SubDepartment = (*Max70Text)(&value)
}

func (p *PostalAddress11) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress11) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress11) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress11) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress11) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress11) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

func (p *PostalAddress11) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

func (p *PostalAddress11) SetFloor(value string) {
	p.Floor = (*Max70Text)(&value)
}

func (p *PostalAddress11) SetPostBox(value string) {
	p.PostBox = (*Max16Text)(&value)
}

func (p *PostalAddress11) SetBuildingName(value string) {
	p.BuildingName = (*Max70Text)(&value)
}

func (p *PostalAddress11) SetRoom(value string) {
	p.Room = (*Max70Text)(&value)
}

// Address of a party expressed in a formal structure.
type PostalAddress12 struct {

	// Name of a built-up area, with defined boundaries, and a local government.
	//
	TownName *Max35Text `xml:"TwnNm,omitempty"`

	// Identifies a subdivision of a country eg, state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PostalAddress12) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress12) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress12) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress13 struct {

	// Information that locates and identifies a specific address, as defined by postal services, that is presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty"`

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of a building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm"`

	// Identifies a subdivision of a country eg, state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government, occupying a particular territory.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PostalAddress13) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

func (p *PostalAddress13) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress13) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress13) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress13) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress13) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress13) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

// Information that locates and identifies a specific address, as defined by postal services or in free format text.
type PostalAddress17 struct {

	// Information that locates and identifies a specific address, as defined by postal services, that is presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty"`

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of a building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm"`

	// Identifies a subdivision of a country, for instance state, region, county.
	CountrySubDivision []*Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government, occupying a particular territory.
	Country *CountryCode `xml:"Ctry"`

	// Geographic location of the ATM specified by geographic coordinates or UTM coordinates.
	GeoLocation *GeographicLocation1Choice `xml:"GLctn,omitempty"`
}

func (p *PostalAddress17) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

func (p *PostalAddress17) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress17) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress17) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress17) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress17) AddCountrySubDivision(value string) {
	p.CountrySubDivision = append(p.CountrySubDivision, (*Max35Text)(&value))
}

func (p *PostalAddress17) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

func (p *PostalAddress17) AddGeoLocation() *GeographicLocation1Choice {
	p.GeoLocation = new(GeographicLocation1Choice)
	return p.GeoLocation
}

// Information that locates and identifies a specific address, as defined by postal services or in free format text.
type PostalAddress18 struct {

	// Information that locates and identifies a specific address, as defined by postal services, that is presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty"`

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of a building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm"`

	// Identifies a subdivision of a country, for instance state, region, county.
	CountrySubDivision []*Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government, occupying a particular territory.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PostalAddress18) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

func (p *PostalAddress18) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress18) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress18) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress18) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress18) AddCountrySubDivision(value string) {
	p.CountrySubDivision = append(p.CountrySubDivision, (*Max35Text)(&value))
}

func (p *PostalAddress18) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress19 struct {

	// Identifies the nature of the postal address.
	AddressType *AddressType2Code `xml:"AdrTp,omitempty"`

	// Identification of a division of a large organisation or building.
	Department *Max70Text `xml:"Dept,omitempty"`

	// Identification of a sub-division of a large organisation or building.
	SubDepartment *Max70Text `xml:"SubDept,omitempty"`

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of a building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm,omitempty"`

	// Specific location name within the town.
	TownLocationName *Max35Text `xml:"TwnLctnNm,omitempty"`

	// Identifies a subdivision within a country sub-division.
	DistrictName *Max35Text `xml:"DstrctNm,omitempty"`

	// Identifies a subdivision of a country such as state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services, presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty"`
}

func (p *PostalAddress19) SetAddressType(value string) {
	p.AddressType = (*AddressType2Code)(&value)
}

func (p *PostalAddress19) SetDepartment(value string) {
	p.Department = (*Max70Text)(&value)
}

func (p *PostalAddress19) SetSubDepartment(value string) {
	p.SubDepartment = (*Max70Text)(&value)
}

func (p *PostalAddress19) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress19) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress19) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress19) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress19) SetTownLocationName(value string) {
	p.TownLocationName = (*Max35Text)(&value)
}

func (p *PostalAddress19) SetDistrictName(value string) {
	p.DistrictName = (*Max35Text)(&value)
}

func (p *PostalAddress19) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress19) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

func (p *PostalAddress19) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

// Address of a party expressed in a formal structure, usually according to the country's postal services specifications.
type PostalAddress2 struct {

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCodeIdentification *Max16Text `xml:"PstCdId"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm"`

	// Identifies a subdivision of a country for example, state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PostalAddress2) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress2) SetPostCodeIdentification(value string) {
	p.PostCodeIdentification = (*Max16Text)(&value)
}

func (p *PostalAddress2) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress2) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress2) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress21 struct {

	// Type of address.
	AddressType *AddressType2Choice `xml:"AdrTp,omitempty"`

	// Indicates whether mail should be sent to the address.
	MailingIndicator *YesNoIndicator `xml:"MlngInd,omitempty"`

	// Indicates whether the address is the official address of the party.
	RegistrationAddressIndicator *YesNoIndicator `xml:"RegnAdrInd,omitempty"`

	// When the individual resides at another person’s address, the name of the other person.
	CareOf *Max70Text `xml:"CareOf,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services, that is presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty"`

	// Name of the street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of the building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty"`

	// Name of the building or house.
	BuildingName *Max35Text `xml:"BldgNm,omitempty"`

	// Post box number of the addressee within the residential or company building.
	PostBox *Max10Text `xml:"PstBx,omitempty"`

	// Side or wing of the building, for example, ‘wing A’.
	SideInBuilding *Max35Text `xml:"SdInBldg,omitempty"`

	// Floor or storey within the building.
	Floor *Max70Text `xml:"Flr,omitempty"`

	// Identification of the suite or apartment.
	SuiteIdentification *Max10Text `xml:"SuiteId,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd,omitempty"`

	// Name of a district, that is, a part of the town or region.
	DistrictName *Max35Text `xml:"DstrctNm,omitempty"`

	// Name of the village.
	Village *Max70Text `xml:"Vllg,omitempty"`

	// Name of the town or city.
	TownName *Max35Text `xml:"TwnNm,omitempty"`

	// Name of the state, county or country sub-division.
	State *Max70Text `xml:"Stat,omitempty"`

	// Country of the address.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PostalAddress21) AddAddressType() *AddressType2Choice {
	p.AddressType = new(AddressType2Choice)
	return p.AddressType
}

func (p *PostalAddress21) SetMailingIndicator(value string) {
	p.MailingIndicator = (*YesNoIndicator)(&value)
}

func (p *PostalAddress21) SetRegistrationAddressIndicator(value string) {
	p.RegistrationAddressIndicator = (*YesNoIndicator)(&value)
}

func (p *PostalAddress21) SetCareOf(value string) {
	p.CareOf = (*Max70Text)(&value)
}

func (p *PostalAddress21) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

func (p *PostalAddress21) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress21) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress21) SetBuildingName(value string) {
	p.BuildingName = (*Max35Text)(&value)
}

func (p *PostalAddress21) SetPostBox(value string) {
	p.PostBox = (*Max10Text)(&value)
}

func (p *PostalAddress21) SetSideInBuilding(value string) {
	p.SideInBuilding = (*Max35Text)(&value)
}

func (p *PostalAddress21) SetFloor(value string) {
	p.Floor = (*Max70Text)(&value)
}

func (p *PostalAddress21) SetSuiteIdentification(value string) {
	p.SuiteIdentification = (*Max10Text)(&value)
}

func (p *PostalAddress21) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress21) SetDistrictName(value string) {
	p.DistrictName = (*Max35Text)(&value)
}

func (p *PostalAddress21) SetVillage(value string) {
	p.Village = (*Max70Text)(&value)
}

func (p *PostalAddress21) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress21) SetState(value string) {
	p.State = (*Max70Text)(&value)
}

func (p *PostalAddress21) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress3 struct {

	// Type of address.
	AddressType *AddressType1Code `xml:"AdrTp"`

	// Indicates whether mail should be sent to an address.
	MailingIndicator *YesNoIndicator `xml:"MlngInd"`

	// Indicates whether the address is the official address of the party.
	RegistrationAddressIndicator *YesNoIndicator `xml:"RegnAdrInd"`

	// Information that locates and identifies a specific address, as defined by postal services.
	NameAndAddress *NameAndAddress4 `xml:"NmAndAdr"`
}

func (p *PostalAddress3) SetAddressType(value string) {
	p.AddressType = (*AddressType1Code)(&value)
}

func (p *PostalAddress3) SetMailingIndicator(value string) {
	p.MailingIndicator = (*YesNoIndicator)(&value)
}

func (p *PostalAddress3) SetRegistrationAddressIndicator(value string) {
	p.RegistrationAddressIndicator = (*YesNoIndicator)(&value)
}

func (p *PostalAddress3) AddNameAndAddress() *NameAndAddress4 {
	p.NameAndAddress = new(NameAndAddress4)
	return p.NameAndAddress
}

// Address of a party expressed in a formal structure, usually according to the country's postal services specifications.
type PostalAddress5 struct {

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCodeIdentification *Max16Text `xml:"PstCdId,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm,omitempty"`

	// Identifies a subdivision of a country eg, state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PostalAddress5) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress5) SetPostCodeIdentification(value string) {
	p.PostCodeIdentification = (*Max16Text)(&value)
}

func (p *PostalAddress5) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress5) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress5) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress6 struct {

	// Identifies the nature of the postal address.
	AddressType *AddressType2Code `xml:"AdrTp,omitempty"`

	// Identification of a division of a large organisation or building.
	Department *Max70Text `xml:"Dept,omitempty"`

	// Identification of a sub-division of a large organisation or building.
	SubDepartment *Max70Text `xml:"SubDept,omitempty"`

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of a building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm,omitempty"`

	// Identifies a subdivision of a country such as state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services, presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty"`
}

func (p *PostalAddress6) SetAddressType(value string) {
	p.AddressType = (*AddressType2Code)(&value)
}

func (p *PostalAddress6) SetDepartment(value string) {
	p.Department = (*Max70Text)(&value)
}

func (p *PostalAddress6) SetSubDepartment(value string) {
	p.SubDepartment = (*Max70Text)(&value)
}

func (p *PostalAddress6) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress6) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress6) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress6) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress6) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress6) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

func (p *PostalAddress6) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress8 struct {

	// Identifies the nature of the postal address.
	AddressType *AddressType2Code `xml:"AdrTp,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services, that is presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty"`

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of a building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm,omitempty"`

	// Identifies a subdivision of a country for example, state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government, occupying a particular territory.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PostalAddress8) SetAddressType(value string) {
	p.AddressType = (*AddressType2Code)(&value)
}

func (p *PostalAddress8) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

func (p *PostalAddress8) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress8) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress8) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress8) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress8) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress8) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
