package bot

import "time"

const (
	intentAssistant   = "assistant"
	intentSchedule    = "schedule"
	intentInformation = "information"
	intentGrade       = "grade"
	intentAssignment  = "assignment"
	intentUnknown     = "unknown"

	rgxMonday    = "(senin|senen|monday)"
	rgxTuesday   = "(selasa|tuesday)"
	rgxWednesday = "(rabu|rebo|wednesday)"
	rgxThursday  = "(kamis|kemis|thursday)"
	rgxFriday    = "(jumat|jum'at|jumaah|friday)"
	rgxSaturday  = "(sabtu|septu|saturday)"
	rgxSunday    = "(minggu|sunday)"
)

var rgxAssistant string
var rgxCourse string

type sAssistant string
type sSchedule string
type sInformation string
type sGrade string
type sAssignment string

type sEntity struct {
	text   string
	userID int64
}

type messageParams struct {
	Text string
}

type messageArgs struct {
	Text           string
	NormalizedText string
}

type messageResponse struct {
	Text      string      `json:"original_text"`
	TimeStamp int64       `json:"time_stamp"`
	Response  interface{} `json:"response"`
}

type loadHistoryParams struct {
	Time     string
	Position string
}

type loadHistoryArgs struct {
	Time    time.Time
	IsAfter bool
}

type loadHistoryResponse struct {
	TimeStamp int64       `json:"time_stamp"`
	Response  interface{} `json:"response"`
}
