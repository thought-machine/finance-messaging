package iso20022

// Set of elements that further details the information related to the type of payment.
type PaymentTypeInformation1 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel2Choice `xml:"SvcLvl,omitempty"`

	// Specifies the clearing channel to be used for the instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`

	// User community specific instrument.
	//
	// Usage : When available, codes provided by local authorities should be used.
	LocalInstrument *LocalInstrument1Choice `xml:"LclInstrm,omitempty"`

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	CategoryPurpose *PaymentCategoryPurpose1Code `xml:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation1) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation1) AddServiceLevel() *ServiceLevel2Choice {
	p.ServiceLevel = new(ServiceLevel2Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation1) SetClearingChannel(value string) {
	p.ClearingChannel = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation1) AddLocalInstrument() *LocalInstrument1Choice {
	p.LocalInstrument = new(LocalInstrument1Choice)
	return p.LocalInstrument
}

func (p *PaymentTypeInformation1) SetCategoryPurpose(value string) {
	p.CategoryPurpose = (*PaymentCategoryPurpose1Code)(&value)
}

// Set of elements used to provide further details of the type of payment.
type PaymentTypeInformation19 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel8Choice `xml:"SvcLvl,omitempty"`

	// User community specific instrument.
	//
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LocalInstrument *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
	CategoryPurpose *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation19) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation19) AddServiceLevel() *ServiceLevel8Choice {
	p.ServiceLevel = new(ServiceLevel8Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation19) AddLocalInstrument() *LocalInstrument2Choice {
	p.LocalInstrument = new(LocalInstrument2Choice)
	return p.LocalInstrument
}

func (p *PaymentTypeInformation19) AddCategoryPurpose() *CategoryPurpose1Choice {
	p.CategoryPurpose = new(CategoryPurpose1Choice)
	return p.CategoryPurpose
}

// Set of elements that further details the information related to the type of payment.
type PaymentTypeInformation2 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel3Choice `xml:"SvcLvl,omitempty"`

	// Specifies the clearing channel to be used for the instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`

	// User community specific instrument.
	//
	// Usage : When available, codes provided by local authorities should be used.
	LocalInstrument *LocalInstrument1Choice `xml:"LclInstrm,omitempty"`

	// Identifies the direct debit sequence, eg, first, recurrent, final or one-off.
	SequenceType *SequenceType1Code `xml:"SeqTp,omitempty"`

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	CategoryPurpose *PaymentCategoryPurpose1Code `xml:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation2) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation2) AddServiceLevel() *ServiceLevel3Choice {
	p.ServiceLevel = new(ServiceLevel3Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation2) SetClearingChannel(value string) {
	p.ClearingChannel = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation2) AddLocalInstrument() *LocalInstrument1Choice {
	p.LocalInstrument = new(LocalInstrument1Choice)
	return p.LocalInstrument
}

func (p *PaymentTypeInformation2) SetSequenceType(value string) {
	p.SequenceType = (*SequenceType1Code)(&value)
}

func (p *PaymentTypeInformation2) SetCategoryPurpose(value string) {
	p.CategoryPurpose = (*PaymentCategoryPurpose1Code)(&value)
}

// Set of elements used to provide further details of the type of payment.
type PaymentTypeInformation20 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel8Choice `xml:"SvcLvl,omitempty"`

	// User community specific instrument.
	//
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LocalInstrument *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`

	// Identifies the direct debit sequence, such as first, recurrent, final or one-off.
	SequenceType *SequenceType1Code `xml:"SeqTp,omitempty"`

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
	CategoryPurpose *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation20) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation20) AddServiceLevel() *ServiceLevel8Choice {
	p.ServiceLevel = new(ServiceLevel8Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation20) AddLocalInstrument() *LocalInstrument2Choice {
	p.LocalInstrument = new(LocalInstrument2Choice)
	return p.LocalInstrument
}

func (p *PaymentTypeInformation20) SetSequenceType(value string) {
	p.SequenceType = (*SequenceType1Code)(&value)
}

func (p *PaymentTypeInformation20) AddCategoryPurpose() *CategoryPurpose1Choice {
	p.CategoryPurpose = new(CategoryPurpose1Choice)
	return p.CategoryPurpose
}

// Set of elements used to provide further details of the type of payment.
type PaymentTypeInformation21 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Specifies the clearing channel to be used to process the payment instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel8Choice `xml:"SvcLvl,omitempty"`

	// User community specific instrument.
	//
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LocalInstrument *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
	CategoryPurpose *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation21) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation21) SetClearingChannel(value string) {
	p.ClearingChannel = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation21) AddServiceLevel() *ServiceLevel8Choice {
	p.ServiceLevel = new(ServiceLevel8Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation21) AddLocalInstrument() *LocalInstrument2Choice {
	p.LocalInstrument = new(LocalInstrument2Choice)
	return p.LocalInstrument
}

func (p *PaymentTypeInformation21) AddCategoryPurpose() *CategoryPurpose1Choice {
	p.CategoryPurpose = new(CategoryPurpose1Choice)
	return p.CategoryPurpose
}

// Set of elements used to provide further details of the type of payment.
type PaymentTypeInformation22 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Specifies the clearing channel to be used to process the payment instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel8Choice `xml:"SvcLvl,omitempty"`

	// User community specific instrument.
	//
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LocalInstrument *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`

	// Identifies the direct debit sequence, such as first, recurrent, final or one-off.
	SequenceType *SequenceType1Code `xml:"SeqTp,omitempty"`

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
	CategoryPurpose *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation22) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation22) SetClearingChannel(value string) {
	p.ClearingChannel = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation22) AddServiceLevel() *ServiceLevel8Choice {
	p.ServiceLevel = new(ServiceLevel8Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation22) AddLocalInstrument() *LocalInstrument2Choice {
	p.LocalInstrument = new(LocalInstrument2Choice)
	return p.LocalInstrument
}

func (p *PaymentTypeInformation22) SetSequenceType(value string) {
	p.SequenceType = (*SequenceType1Code)(&value)
}

func (p *PaymentTypeInformation22) AddCategoryPurpose() *CategoryPurpose1Choice {
	p.CategoryPurpose = new(CategoryPurpose1Choice)
	return p.CategoryPurpose
}

// Set of elements used to provide further details of the type of payment.
type PaymentTypeInformation23 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Specifies the clearing channel to be used to process the payment instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel8Choice `xml:"SvcLvl,omitempty"`

	// User community specific instrument.
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LocalInstrument *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`
}

func (p *PaymentTypeInformation23) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation23) SetClearingChannel(value string) {
	p.ClearingChannel = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation23) AddServiceLevel() *ServiceLevel8Choice {
	p.ServiceLevel = new(ServiceLevel8Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation23) AddLocalInstrument() *LocalInstrument2Choice {
	p.LocalInstrument = new(LocalInstrument2Choice)
	return p.LocalInstrument
}

// Set of elements used to provide further details of the type of payment.
type PaymentTypeInformation24 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel8Choice `xml:"SvcLvl,omitempty"`

	// User community specific instrument.
	//
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LocalInstrument *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`

	// Identifies the direct debit sequence, such as first, recurrent, final or one-off.
	SequenceType *SequenceType3Code `xml:"SeqTp,omitempty"`

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
	CategoryPurpose *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation24) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation24) AddServiceLevel() *ServiceLevel8Choice {
	p.ServiceLevel = new(ServiceLevel8Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation24) AddLocalInstrument() *LocalInstrument2Choice {
	p.LocalInstrument = new(LocalInstrument2Choice)
	return p.LocalInstrument
}

func (p *PaymentTypeInformation24) SetSequenceType(value string) {
	p.SequenceType = (*SequenceType3Code)(&value)
}

func (p *PaymentTypeInformation24) AddCategoryPurpose() *CategoryPurpose1Choice {
	p.CategoryPurpose = new(CategoryPurpose1Choice)
	return p.CategoryPurpose
}

// Set of elements used to provide further details of the type of payment.
type PaymentTypeInformation25 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Specifies the clearing channel to be used to process the payment instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel8Choice `xml:"SvcLvl,omitempty"`

	// User community specific instrument.
	//
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LocalInstrument *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`

	// Identifies the direct debit sequence, such as first, recurrent, final or one-off.
	SequenceType *SequenceType3Code `xml:"SeqTp,omitempty"`

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
	CategoryPurpose *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation25) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation25) SetClearingChannel(value string) {
	p.ClearingChannel = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation25) AddServiceLevel() *ServiceLevel8Choice {
	p.ServiceLevel = new(ServiceLevel8Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation25) AddLocalInstrument() *LocalInstrument2Choice {
	p.LocalInstrument = new(LocalInstrument2Choice)
	return p.LocalInstrument
}

func (p *PaymentTypeInformation25) SetSequenceType(value string) {
	p.SequenceType = (*SequenceType3Code)(&value)
}

func (p *PaymentTypeInformation25) AddCategoryPurpose() *CategoryPurpose1Choice {
	p.CategoryPurpose = new(CategoryPurpose1Choice)
	return p.CategoryPurpose
}

// Set of elements that further details the information related to the type of payment.
type PaymentTypeInformation3 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel2Choice `xml:"SvcLvl,omitempty"`

	// Specifies the clearing channel to be used to process the payment instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`

	// User community specific instrument.
	//
	// Usage : When available, codes provided by local authorities should be used.
	LocalInstrument *LocalInstrument1Choice `xml:"LclInstrm,omitempty"`

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	CategoryPurpose *PaymentCategoryPurpose1Code `xml:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation3) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation3) AddServiceLevel() *ServiceLevel2Choice {
	p.ServiceLevel = new(ServiceLevel2Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation3) SetClearingChannel(value string) {
	p.ClearingChannel = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation3) AddLocalInstrument() *LocalInstrument1Choice {
	p.LocalInstrument = new(LocalInstrument1Choice)
	return p.LocalInstrument
}

func (p *PaymentTypeInformation3) SetCategoryPurpose(value string) {
	p.CategoryPurpose = (*PaymentCategoryPurpose1Code)(&value)
}

// Set of elements that further details the information related to the type of payment.
type PaymentTypeInformation4 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel3Choice `xml:"SvcLvl,omitempty"`

	// Specifies the clearing channel to be used for the instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`

	// User community specific instrument.
	//
	// Usage : When available, codes provided by local authorities should be used.
	LocalInstrument *LocalInstrument1Choice `xml:"LclInstrm,omitempty"`

	// Identifies the direct debit sequence, eg, first, recurrent, final or one-off.
	SequenceType *SequenceType1Code `xml:"SeqTp,omitempty"`

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	CategoryPurpose *PaymentCategoryPurpose1Code `xml:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation4) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation4) AddServiceLevel() *ServiceLevel3Choice {
	p.ServiceLevel = new(ServiceLevel3Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation4) SetClearingChannel(value string) {
	p.ClearingChannel = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation4) AddLocalInstrument() *LocalInstrument1Choice {
	p.LocalInstrument = new(LocalInstrument1Choice)
	return p.LocalInstrument
}

func (p *PaymentTypeInformation4) SetSequenceType(value string) {
	p.SequenceType = (*SequenceType1Code)(&value)
}

func (p *PaymentTypeInformation4) SetCategoryPurpose(value string) {
	p.CategoryPurpose = (*PaymentCategoryPurpose1Code)(&value)
}

// Set of elements that further details the information related to the type of payment.
type PaymentTypeInformation5 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *RestrictedProprietaryChoice `xml:"SvcLvl,omitempty"`

	// Specifies the clearing channel to be used for the instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`

	// User community specific instrument required for use within that user community.
	//
	// Usage : When available, codes provided by local authorities should be used.
	LocalInstrument *RestrictedProprietaryChoice `xml:"LclInstrm,omitempty"`
}

func (p *PaymentTypeInformation5) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation5) AddServiceLevel() *RestrictedProprietaryChoice {
	p.ServiceLevel = new(RestrictedProprietaryChoice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation5) SetClearingChannel(value string) {
	p.ClearingChannel = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation5) AddLocalInstrument() *RestrictedProprietaryChoice {
	p.LocalInstrument = new(RestrictedProprietaryChoice)
	return p.LocalInstrument
}

// Set of elements that further details the information related to the type of payment.
type PaymentTypeInformation6 struct {

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Agreement under which or rules under which the transaction should be processed.
	ServiceLevel *ServiceLevel2Choice `xml:"SvcLvl,omitempty"`

	// Specifies the clearing channel to be used to process the payment instruction.
	ClearingChannel *ClearingChannel2Code `xml:"ClrChanl,omitempty"`

	// User community specific instrument.
	//
	// Usage : When available, codes provided by local authorities should be used.
	LocalInstrument *LocalInstrument1Choice `xml:"LclInstrm,omitempty"`

	// Identifies the direct debit sequence, eg, first, recurrent, final or one-off.
	SequenceType *SequenceType1Code `xml:"SeqTp,omitempty"`

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	CategoryPurpose *PaymentCategoryPurpose1Code `xml:"CtgyPurp,omitempty"`
}

func (p *PaymentTypeInformation6) SetInstructionPriority(value string) {
	p.InstructionPriority = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation6) AddServiceLevel() *ServiceLevel2Choice {
	p.ServiceLevel = new(ServiceLevel2Choice)
	return p.ServiceLevel
}

func (p *PaymentTypeInformation6) SetClearingChannel(value string) {
	p.ClearingChannel = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation6) AddLocalInstrument() *LocalInstrument1Choice {
	p.LocalInstrument = new(LocalInstrument1Choice)
	return p.LocalInstrument
}

func (p *PaymentTypeInformation6) SetSequenceType(value string) {
	p.SequenceType = (*SequenceType1Code)(&value)
}

func (p *PaymentTypeInformation6) SetCategoryPurpose(value string) {
	p.CategoryPurpose = (*PaymentCategoryPurpose1Code)(&value)
}
