package iso20022

// Additional information with update description and date.
type UpdatedAdditionalInformation1 struct {

	// Specifies the amendments made to the narrative since the last message.
	UpdateDescription *Max140Text `xml:"UpdDesc,omitempty"`

	// Specifies the date at which the narrative has been updated.
	UpdateDate *ISODate `xml:"UpdDt,omitempty"`

	// Provides additional textual information.
	AdditionalInformation []*Max350Text `xml:"AddtlInf"`
}

func (u *UpdatedAdditionalInformation1) SetUpdateDescription(value string) {
	u.UpdateDescription = (*Max140Text)(&value)
}

func (u *UpdatedAdditionalInformation1) SetUpdateDate(value string) {
	u.UpdateDate = (*ISODate)(&value)
}

func (u *UpdatedAdditionalInformation1) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max350Text)(&value))
}

// Additional information with update description and date.
type UpdatedAdditionalInformation10 struct {

	// Specifies the amendments made to the narrative since the last message.
	UpdateDescription *RestrictedFINXMax140Text `xml:"UpdDesc,omitempty"`

	// Specifies the date at which the narrative has been updated.
	UpdateDate *ISODate `xml:"UpdDt,omitempty"`

	// Provides additional textual information.
	AdditionalInformation []*RestrictedFINZMax8000Text `xml:"AddtlInf"`
}

func (u *UpdatedAdditionalInformation10) SetUpdateDescription(value string) {
	u.UpdateDescription = (*RestrictedFINXMax140Text)(&value)
}

func (u *UpdatedAdditionalInformation10) SetUpdateDate(value string) {
	u.UpdateDate = (*ISODate)(&value)
}

func (u *UpdatedAdditionalInformation10) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*RestrictedFINZMax8000Text)(&value))
}

// Additional information with update description and date.
type UpdatedAdditionalInformation2 struct {

	// Specifies the amendments made to the narrative since the last message.
	UpdateDescription *Max140Text `xml:"UpdDesc,omitempty"`

	// Specifies the date at which the narrative has been updated.
	UpdateDate *ISODate `xml:"UpdDt,omitempty"`

	// Provides additional textual information.
	AdditionalInformation []*Max8000Text `xml:"AddtlInf"`
}

func (u *UpdatedAdditionalInformation2) SetUpdateDescription(value string) {
	u.UpdateDescription = (*Max140Text)(&value)
}

func (u *UpdatedAdditionalInformation2) SetUpdateDate(value string) {
	u.UpdateDate = (*ISODate)(&value)
}

func (u *UpdatedAdditionalInformation2) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max8000Text)(&value))
}

// Additional information with update description and date.
type UpdatedAdditionalInformation3 struct {

	// Specifies the amendments made to the narrative since the last message.
	UpdateDescription *Max140Text `xml:"UpdDesc,omitempty"`

	// Specifies the date at which the narrative has been updated.
	UpdateDate *ISODate `xml:"UpdDt,omitempty"`

	// Provides additional textual information.
	AdditionalInformation *Max350Text `xml:"AddtlInf"`
}

func (u *UpdatedAdditionalInformation3) SetUpdateDescription(value string) {
	u.UpdateDescription = (*Max140Text)(&value)
}

func (u *UpdatedAdditionalInformation3) SetUpdateDate(value string) {
	u.UpdateDate = (*ISODate)(&value)
}

func (u *UpdatedAdditionalInformation3) SetAdditionalInformation(value string) {
	u.AdditionalInformation = (*Max350Text)(&value)
}

// Additional information with update description and date.
type UpdatedAdditionalInformation5 struct {

	// Specifies the amendments made to the narrative since the last message.
	UpdateDescription *RestrictedFINXMax140Text `xml:"UpdDesc,omitempty"`

	// Specifies the date at which the narrative has been updated.
	UpdateDate *ISODate `xml:"UpdDt,omitempty"`

	// Provides additional textual information.
	AdditionalInformation []*RestrictedFINXMax350Text `xml:"AddtlInf"`
}

func (u *UpdatedAdditionalInformation5) SetUpdateDescription(value string) {
	u.UpdateDescription = (*RestrictedFINXMax140Text)(&value)
}

func (u *UpdatedAdditionalInformation5) SetUpdateDate(value string) {
	u.UpdateDate = (*ISODate)(&value)
}

func (u *UpdatedAdditionalInformation5) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*RestrictedFINXMax350Text)(&value))
}

// Additional information with update description and date.
type UpdatedAdditionalInformation6 struct {

	// Specifies the amendments made to the narrative since the last message.
	UpdateDescription *RestrictedFINXMax140Text `xml:"UpdDesc,omitempty"`

	// Specifies the date at which the narrative has been updated.
	UpdateDate *ISODate `xml:"UpdDt,omitempty"`

	// Provides additional textual information.
	AdditionalInformation *RestrictedFINXMax350Text `xml:"AddtlInf"`
}

func (u *UpdatedAdditionalInformation6) SetUpdateDescription(value string) {
	u.UpdateDescription = (*RestrictedFINXMax140Text)(&value)
}

func (u *UpdatedAdditionalInformation6) SetUpdateDate(value string) {
	u.UpdateDate = (*ISODate)(&value)
}

func (u *UpdatedAdditionalInformation6) SetAdditionalInformation(value string) {
	u.AdditionalInformation = (*RestrictedFINXMax350Text)(&value)
}

// Additional information with update description and date.
type UpdatedAdditionalInformation8 struct {

	// Specifies the amendments made to the narrative since the last message.
	UpdateDescription *Max140Text `xml:"UpdDesc,omitempty"`

	// Specifies the date at which the narrative has been updated.
	UpdateDate *ISODate `xml:"UpdDt,omitempty"`

	// Provides additional textual information.
	AdditionalInformation []*Max8000Text `xml:"AddtlInf"`
}

func (u *UpdatedAdditionalInformation8) SetUpdateDescription(value string) {
	u.UpdateDescription = (*Max140Text)(&value)
}

func (u *UpdatedAdditionalInformation8) SetUpdateDate(value string) {
	u.UpdateDate = (*ISODate)(&value)
}

func (u *UpdatedAdditionalInformation8) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max8000Text)(&value))
}
