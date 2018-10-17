package iso20022

// Specifies whether the transaction is on hold/blocked/frozen.
type HoldIndicator2 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Specifies the reason of the registration staus.
	Reason []*RegistrationReason1 `xml:"Rsn,omitempty"`
}

func (h *HoldIndicator2) SetIndicator(value string) {
	h.Indicator = (*YesNoIndicator)(&value)
}

func (h *HoldIndicator2) AddReason() *RegistrationReason1 {
	newValue := new(RegistrationReason1)
	h.Reason = append(h.Reason, newValue)
	return newValue
}

// Specifies whether the transaction is on hold/blocked/frozen.
type HoldIndicator4 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Specifies the reason of the registration staus.
	Reason []*RegistrationReason3 `xml:"Rsn,omitempty"`
}

func (h *HoldIndicator4) SetIndicator(value string) {
	h.Indicator = (*YesNoIndicator)(&value)
}

func (h *HoldIndicator4) AddReason() *RegistrationReason3 {
	newValue := new(RegistrationReason3)
	h.Reason = append(h.Reason, newValue)
	return newValue
}

// Specifies whether the transaction is on hold/blocked/frozen.
type HoldIndicator6 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Specifies the reason of the registration staus.
	Reason []*RegistrationReason5 `xml:"Rsn,omitempty"`
}

func (h *HoldIndicator6) SetIndicator(value string) {
	h.Indicator = (*YesNoIndicator)(&value)
}

func (h *HoldIndicator6) AddReason() *RegistrationReason5 {
	newValue := new(RegistrationReason5)
	h.Reason = append(h.Reason, newValue)
	return newValue
}

// Specifies whether the transaction is on hold/blocked/frozen.
type HoldIndicator7 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Specifies the reason of the registration staus.
	Reason []*RegistrationReason6 `xml:"Rsn,omitempty"`
}

func (h *HoldIndicator7) SetIndicator(value string) {
	h.Indicator = (*YesNoIndicator)(&value)
}

func (h *HoldIndicator7) AddReason() *RegistrationReason6 {
	newValue := new(RegistrationReason6)
	h.Reason = append(h.Reason, newValue)
	return newValue
}
