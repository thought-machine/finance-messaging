package iso20022

// Parameters related to the net position.
type ReportParameters1 struct {

	// After netting, reference that is common to a net transaction to settle and all its underlying trades.
	NetPositionIdentification *Max35Text `xml:"NetPosId"`

	// Date and time of the net position report.
	ReportDateAndTime *DateAndDateTimeChoice `xml:"RptDtAndTm"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *StatementUpdateType1Code `xml:"UpdTp"`

	// Frequency of the report.
	Frequency *EventFrequency6Code `xml:"Frqcy"`

	// Sequential number of the report.
	ReportNumber *Exact5NumericText `xml:"RptNb,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (r *ReportParameters1) SetNetPositionIdentification(value string) {
	r.NetPositionIdentification = (*Max35Text)(&value)
}

func (r *ReportParameters1) AddReportDateAndTime() *DateAndDateTimeChoice {
	r.ReportDateAndTime = new(DateAndDateTimeChoice)
	return r.ReportDateAndTime
}

func (r *ReportParameters1) SetUpdateType(value string) {
	r.UpdateType = (*StatementUpdateType1Code)(&value)
}

func (r *ReportParameters1) SetFrequency(value string) {
	r.Frequency = (*EventFrequency6Code)(&value)
}

func (r *ReportParameters1) SetReportNumber(value string) {
	r.ReportNumber = (*Exact5NumericText)(&value)
}

func (r *ReportParameters1) SetActivityIndicator(value string) {
	r.ActivityIndicator = (*YesNoIndicator)(&value)
}

// Provides the parameters of the report.
type ReportParameters2 struct {

	// Unique identification of the report.
	ReportIdentification *Max35Text `xml:"RptId"`

	// Date (and time) at which the report was created.
	ReportDateAndTime *DateAndDateTimeChoice `xml:"RptDtAndTm"`

	// Frequency of the report.
	Frequency *EventFrequency6Code `xml:"Frqcy"`

	// Indicates the currency used for the calculation of the guarantee fund.
	ReportCurrency *CurrencyCode `xml:"RptCcy"`

	// Indicates the date of calculation of the deficit (if any).
	CalculationDate *ISODateTime `xml:"ClctnDt,omitempty"`
}

func (r *ReportParameters2) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *ReportParameters2) AddReportDateAndTime() *DateAndDateTimeChoice {
	r.ReportDateAndTime = new(DateAndDateTimeChoice)
	return r.ReportDateAndTime
}

func (r *ReportParameters2) SetFrequency(value string) {
	r.Frequency = (*EventFrequency6Code)(&value)
}

func (r *ReportParameters2) SetReportCurrency(value string) {
	r.ReportCurrency = (*CurrencyCode)(&value)
}

func (r *ReportParameters2) SetCalculationDate(value string) {
	r.CalculationDate = (*ISODateTime)(&value)
}

// Provides the parameters of the report.
type ReportParameters3 struct {

	// Unique identification of the report.
	ReportIdentification *Max35Text `xml:"RptId"`

	// Date (and time) and time of the report.
	ReportDateAndTime *DateAndDateTimeChoice `xml:"RptDtAndTm"`

	// Currency used for the calculation of the margin.
	ReportCurrency *CurrencyCode `xml:"RptCcy"`

	// Date of calculation of the margin.
	CalculationDateAndTime *ISODateTime `xml:"ClctnDtAndTm"`

	// Frequency of the report.
	Frequency *EventFrequency6Code `xml:"Frqcy"`

	// Sequential number of the report.
	ReportNumber *Exact5NumericText `xml:"RptNb,omitempty"`
}

func (r *ReportParameters3) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *ReportParameters3) AddReportDateAndTime() *DateAndDateTimeChoice {
	r.ReportDateAndTime = new(DateAndDateTimeChoice)
	return r.ReportDateAndTime
}

func (r *ReportParameters3) SetReportCurrency(value string) {
	r.ReportCurrency = (*CurrencyCode)(&value)
}

func (r *ReportParameters3) SetCalculationDateAndTime(value string) {
	r.CalculationDateAndTime = (*ISODateTime)(&value)
}

func (r *ReportParameters3) SetFrequency(value string) {
	r.Frequency = (*EventFrequency6Code)(&value)
}

func (r *ReportParameters3) SetReportNumber(value string) {
	r.ReportNumber = (*Exact5NumericText)(&value)
}

// Provides the parameters of the report.
type ReportParameters4 struct {

	// Unique identification of the report.
	ReportIdentification *Max35Text `xml:"RptId"`

	// Date and time of the report .
	ReportDateAndTime *DateAndDateTimeChoice `xml:"RptDtAndTm"`
}

func (r *ReportParameters4) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *ReportParameters4) AddReportDateAndTime() *DateAndDateTimeChoice {
	r.ReportDateAndTime = new(DateAndDateTimeChoice)
	return r.ReportDateAndTime
}

// Provides the parameters of the report.
type ReportParameters5 struct {

	// Unique identification of the report.
	ReportIdentification *Max35Text `xml:"RptId"`

	// Date (and time) at which the report was created.
	ReportDateAndTime *DateAndDateTimeChoice `xml:"RptDtAndTm"`

	// Frequency of the report.
	Frequency *EventFrequency6Code `xml:"Frqcy"`

	// Indicates the currency used for the calculation of the guarantee fund.
	ReportCurrency *ActiveCurrencyCode `xml:"RptCcy"`

	// Indicates the date of calculation of the deficit (if any).
	CalculationDate *ISODateTime `xml:"ClctnDt,omitempty"`
}

func (r *ReportParameters5) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *ReportParameters5) AddReportDateAndTime() *DateAndDateTimeChoice {
	r.ReportDateAndTime = new(DateAndDateTimeChoice)
	return r.ReportDateAndTime
}

func (r *ReportParameters5) SetFrequency(value string) {
	r.Frequency = (*EventFrequency6Code)(&value)
}

func (r *ReportParameters5) SetReportCurrency(value string) {
	r.ReportCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *ReportParameters5) SetCalculationDate(value string) {
	r.CalculationDate = (*ISODateTime)(&value)
}
