package iso20022

// Identification of a data set.
type DataSetIdentification1 struct {

	// Name of the data set.
	Name *Max256Text `xml:"Nm"`

	// Category of data set.
	Type *DataSetCategory1Code `xml:"Tp"`

	// Version of the data set.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Date and time of creation of the data set.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (d *DataSetIdentification1) SetName(value string) {
	d.Name = (*Max256Text)(&value)
}

func (d *DataSetIdentification1) SetType(value string) {
	d.Type = (*DataSetCategory1Code)(&value)
}

func (d *DataSetIdentification1) SetVersion(value string) {
	d.Version = (*Max256Text)(&value)
}

func (d *DataSetIdentification1) SetCreationDateTime(value string) {
	d.CreationDateTime = (*ISODateTime)(&value)
}

// Identification of a data set.
type DataSetIdentification2 struct {

	// Name of the data set.
	Name *Max256Text `xml:"Nm,omitempty"`

	// Category of data set.
	Type *DataSetCategory2Code `xml:"Tp"`

	// Version of the data set.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Date and time of creation of the data set.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (d *DataSetIdentification2) SetName(value string) {
	d.Name = (*Max256Text)(&value)
}

func (d *DataSetIdentification2) SetType(value string) {
	d.Type = (*DataSetCategory2Code)(&value)
}

func (d *DataSetIdentification2) SetVersion(value string) {
	d.Version = (*Max256Text)(&value)
}

func (d *DataSetIdentification2) SetCreationDateTime(value string) {
	d.CreationDateTime = (*ISODateTime)(&value)
}

// Identification of a data set.
type DataSetIdentification3 struct {

	// Name of the data set.
	Name *Max256Text `xml:"Nm,omitempty"`

	// Category of data set.
	Type *DataSetCategory3Code `xml:"Tp"`

	// Version of the data set.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Date and time of creation of the data set.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (d *DataSetIdentification3) SetName(value string) {
	d.Name = (*Max256Text)(&value)
}

func (d *DataSetIdentification3) SetType(value string) {
	d.Type = (*DataSetCategory3Code)(&value)
}

func (d *DataSetIdentification3) SetVersion(value string) {
	d.Version = (*Max256Text)(&value)
}

func (d *DataSetIdentification3) SetCreationDateTime(value string) {
	d.CreationDateTime = (*ISODateTime)(&value)
}

// Identification of a data set.
type DataSetIdentification4 struct {

	// Name of the data set.
	Name *Max256Text `xml:"Nm,omitempty"`

	// Category of data set.
	Type *DataSetCategory4Code `xml:"Tp"`

	// Version of the data set.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Date and time of creation of the data set.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (d *DataSetIdentification4) SetName(value string) {
	d.Name = (*Max256Text)(&value)
}

func (d *DataSetIdentification4) SetType(value string) {
	d.Type = (*DataSetCategory4Code)(&value)
}

func (d *DataSetIdentification4) SetVersion(value string) {
	d.Version = (*Max256Text)(&value)
}

func (d *DataSetIdentification4) SetCreationDateTime(value string) {
	d.CreationDateTime = (*ISODateTime)(&value)
}

// Identification of a data set.
type DataSetIdentification5 struct {

	// Name of the data set.
	Name *Max256Text `xml:"Nm"`

	// Category of data set.
	Type *DataSetCategory8Code `xml:"Tp"`

	// Version of the data set.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Date and time of creation of the data set.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (d *DataSetIdentification5) SetName(value string) {
	d.Name = (*Max256Text)(&value)
}

func (d *DataSetIdentification5) SetType(value string) {
	d.Type = (*DataSetCategory8Code)(&value)
}

func (d *DataSetIdentification5) SetVersion(value string) {
	d.Version = (*Max256Text)(&value)
}

func (d *DataSetIdentification5) SetCreationDateTime(value string) {
	d.CreationDateTime = (*ISODateTime)(&value)
}

// Identification of a data set.
type DataSetIdentification6 struct {

	// Name of the data set.
	Name *Max256Text `xml:"Nm,omitempty"`

	// Category of data set.
	Type *DataSetCategory9Code `xml:"Tp"`

	// Version of the data set.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Date and time of creation of the data set.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (d *DataSetIdentification6) SetName(value string) {
	d.Name = (*Max256Text)(&value)
}

func (d *DataSetIdentification6) SetType(value string) {
	d.Type = (*DataSetCategory9Code)(&value)
}

func (d *DataSetIdentification6) SetVersion(value string) {
	d.Version = (*Max256Text)(&value)
}

func (d *DataSetIdentification6) SetCreationDateTime(value string) {
	d.CreationDateTime = (*ISODateTime)(&value)
}
