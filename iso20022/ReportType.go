package iso20022

// Specifies the type of report.
type ReportType1 struct {

	// Specifies whether the pushed through baseline is a new one or an amendment or a resubmission.
	Type *ReportType1Code `xml:"Tp"`
}

func (r *ReportType1) SetType(value string) {
	r.Type = (*ReportType1Code)(&value)
}

// Specifies the type of report.
type ReportType2 struct {

	// Specifies whether the report is precalculated or current.
	Type *ReportType2Code `xml:"Tp"`
}

func (r *ReportType2) SetType(value string) {
	r.Type = (*ReportType2Code)(&value)
}

// Specifies the type of report.
type ReportType3 struct {

	// Specifies whether the report is for a matched or pre-matched data set.
	Type *InstructionType3Code `xml:"Tp"`
}

func (r *ReportType3) SetType(value string) {
	r.Type = (*InstructionType3Code)(&value)
}
