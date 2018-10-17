package iso20022

// Additional specific query criteria.
type AdditionalQueryParameters1 struct {

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status.
	Status *Status1Choice `xml:"Sts,omitempty"`

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status reason.
	Reason []*Reason1Choice `xml:"Rsn,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification []*SecurityIdentification11 `xml:"FinInstrmId,omitempty"`
}

func (a *AdditionalQueryParameters1) AddStatus() *Status1Choice {
	a.Status = new(Status1Choice)
	return a.Status
}

func (a *AdditionalQueryParameters1) AddReason() *Reason1Choice {
	newValue := new(Reason1Choice)
	a.Reason = append(a.Reason, newValue)
	return newValue
}

func (a *AdditionalQueryParameters1) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	newValue := new(SecurityIdentification11)
	a.FinancialInstrumentIdentification = append(a.FinancialInstrumentIdentification, newValue)
	return newValue
}

// Additional specific query criteria.
type AdditionalQueryParameters11 struct {

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status.
	Status *Status19Choice `xml:"Sts,omitempty"`

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status reason.
	Reason []*Reason16Choice `xml:"Rsn,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification []*SecurityIdentification19 `xml:"FinInstrmId,omitempty"`
}

func (a *AdditionalQueryParameters11) AddStatus() *Status19Choice {
	a.Status = new(Status19Choice)
	return a.Status
}

func (a *AdditionalQueryParameters11) AddReason() *Reason16Choice {
	newValue := new(Reason16Choice)
	a.Reason = append(a.Reason, newValue)
	return newValue
}

func (a *AdditionalQueryParameters11) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	newValue := new(SecurityIdentification19)
	a.FinancialInstrumentIdentification = append(a.FinancialInstrumentIdentification, newValue)
	return newValue
}

// Additional specific query criteria.
type AdditionalQueryParameters12 struct {

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status.
	Status *Status22Choice `xml:"Sts,omitempty"`

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status reason.
	Reason []*Reason17Choice `xml:"Rsn,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification []*SecurityIdentification20 `xml:"FinInstrmId,omitempty"`
}

func (a *AdditionalQueryParameters12) AddStatus() *Status22Choice {
	a.Status = new(Status22Choice)
	return a.Status
}

func (a *AdditionalQueryParameters12) AddReason() *Reason17Choice {
	newValue := new(Reason17Choice)
	a.Reason = append(a.Reason, newValue)
	return newValue
}

func (a *AdditionalQueryParameters12) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	newValue := new(SecurityIdentification20)
	a.FinancialInstrumentIdentification = append(a.FinancialInstrumentIdentification, newValue)
	return newValue
}

// Additional specific query criteria.
type AdditionalQueryParameters3 struct {

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status.
	Status *Status8Choice `xml:"Sts,omitempty"`

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status reason.
	Reason []*Reason6Choice `xml:"Rsn,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification []*SecurityIdentification14 `xml:"FinInstrmId,omitempty"`
}

func (a *AdditionalQueryParameters3) AddStatus() *Status8Choice {
	a.Status = new(Status8Choice)
	return a.Status
}

func (a *AdditionalQueryParameters3) AddReason() *Reason6Choice {
	newValue := new(Reason6Choice)
	a.Reason = append(a.Reason, newValue)
	return newValue
}

func (a *AdditionalQueryParameters3) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	newValue := new(SecurityIdentification14)
	a.FinancialInstrumentIdentification = append(a.FinancialInstrumentIdentification, newValue)
	return newValue
}

// Additional specific query criteria.
type AdditionalQueryParameters5 struct {

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status.
	Status *Status8Choice `xml:"Sts,omitempty"`

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status reason.
	Reason []*Reason7Choice `xml:"Rsn,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification []*SecurityIdentification14 `xml:"FinInstrmId,omitempty"`
}

func (a *AdditionalQueryParameters5) AddStatus() *Status8Choice {
	a.Status = new(Status8Choice)
	return a.Status
}

func (a *AdditionalQueryParameters5) AddReason() *Reason7Choice {
	newValue := new(Reason7Choice)
	a.Reason = append(a.Reason, newValue)
	return newValue
}

func (a *AdditionalQueryParameters5) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	newValue := new(SecurityIdentification14)
	a.FinancialInstrumentIdentification = append(a.FinancialInstrumentIdentification, newValue)
	return newValue
}

// Additional specific query criteria.
type AdditionalQueryParameters7 struct {

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status.
	Status *Status8Choice `xml:"Sts,omitempty"`

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status reason.
	Reason []*Reason12Choice `xml:"Rsn,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification []*SecurityIdentification14 `xml:"FinInstrmId,omitempty"`
}

func (a *AdditionalQueryParameters7) AddStatus() *Status8Choice {
	a.Status = new(Status8Choice)
	return a.Status
}

func (a *AdditionalQueryParameters7) AddReason() *Reason12Choice {
	newValue := new(Reason12Choice)
	a.Reason = append(a.Reason, newValue)
	return newValue
}

func (a *AdditionalQueryParameters7) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	newValue := new(SecurityIdentification14)
	a.FinancialInstrumentIdentification = append(a.FinancialInstrumentIdentification, newValue)
	return newValue
}

// Additional specific query criteria.
type AdditionalQueryParameters9 struct {

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status.
	Status *Status8Choice `xml:"Sts,omitempty"`

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status reason.
	Reason []*Reason14Choice `xml:"Rsn,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification []*SecurityIdentification14 `xml:"FinInstrmId,omitempty"`
}

func (a *AdditionalQueryParameters9) AddStatus() *Status8Choice {
	a.Status = new(Status8Choice)
	return a.Status
}

func (a *AdditionalQueryParameters9) AddReason() *Reason14Choice {
	newValue := new(Reason14Choice)
	a.Reason = append(a.Reason, newValue)
	return newValue
}

func (a *AdditionalQueryParameters9) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	newValue := new(SecurityIdentification14)
	a.FinancialInstrumentIdentification = append(a.FinancialInstrumentIdentification, newValue)
	return newValue
}
