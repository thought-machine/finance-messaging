package iso20022

// Specifies detailed voting instructions.
type VoteDetails1 struct {

	// Indicates the vote instruction for the resolutions which are announced via the meeting agenda in advance of the meeting.
	VoteInstructionForAgendaResolution *Vote1Choice `xml:"VoteInstrForAgndRsltn"`

	// Indicates the vote instruction for the resolutions that will be added during the meeting.
	VoteInstructionForMeetingResolution *VoteInstructionForMeetingResolution1 `xml:"VoteInstrForMtgRsltn,omitempty"`
}

func (v *VoteDetails1) AddVoteInstructionForAgendaResolution() *Vote1Choice {
	v.VoteInstructionForAgendaResolution = new(Vote1Choice)
	return v.VoteInstructionForAgendaResolution
}

func (v *VoteDetails1) AddVoteInstructionForMeetingResolution() *VoteInstructionForMeetingResolution1 {
	v.VoteInstructionForMeetingResolution = new(VoteInstructionForMeetingResolution1)
	return v.VoteInstructionForMeetingResolution
}

// Specifies detailed voting instructions.
type VoteDetails2 struct {

	// Indicates the vote instruction for the resolutions which are announced via the meeting agenda in advance of the meeting.
	VoteInstructionForAgendaResolution *Vote2Choice `xml:"VoteInstrForAgndRsltn"`

	// Indicates the vote instruction for the resolutions that may arise at the meeting but were not previously provided in the agenda.
	VoteInstructionForMeetingResolution *VoteInstructionForMeetingResolution1Choice `xml:"VoteInstrForMtgRsltn,omitempty"`
}

func (v *VoteDetails2) AddVoteInstructionForAgendaResolution() *Vote2Choice {
	v.VoteInstructionForAgendaResolution = new(Vote2Choice)
	return v.VoteInstructionForAgendaResolution
}

func (v *VoteDetails2) AddVoteInstructionForMeetingResolution() *VoteInstructionForMeetingResolution1Choice {
	v.VoteInstructionForMeetingResolution = new(VoteInstructionForMeetingResolution1Choice)
	return v.VoteInstructionForMeetingResolution
}

// Specifies detailed voting instructions.
type VoteDetails3 struct {

	// Indicates the vote instruction for the resolutions that are announced via the meeting agenda in advance of the meeting.
	VoteInstructionForAgendaResolution *Vote3Choice `xml:"VoteInstrForAgndRsltn"`

	// Indicates the vote instruction for the resolutions that may arise at the meeting but were not previously provided in the agenda.
	VoteInstructionForMeetingResolution *VoteInstructionForMeetingResolution2Choice `xml:"VoteInstrForMtgRsltn,omitempty"`
}

func (v *VoteDetails3) AddVoteInstructionForAgendaResolution() *Vote3Choice {
	v.VoteInstructionForAgendaResolution = new(Vote3Choice)
	return v.VoteInstructionForAgendaResolution
}

func (v *VoteDetails3) AddVoteInstructionForMeetingResolution() *VoteInstructionForMeetingResolution2Choice {
	v.VoteInstructionForMeetingResolution = new(VoteInstructionForMeetingResolution2Choice)
	return v.VoteInstructionForMeetingResolution
}
