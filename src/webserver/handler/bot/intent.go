package bot

import "regexp"
import "fmt"

func getIntent(text string) (string, error) {
	var intents []string

	// get assistant intent
	if sAssistant(text).isValidIntent() {
		intents = append(intents, intentAssistant)
	}

	// get assignment intent
	if sAssignment(text).isValidIntent() {
		intents = append(intents, intentAssignment)
	}

	// get information intent
	if sInformation(text).isValidIntent() {
		intents = append(intents, intentInformation)
	}

	// get grade intent
	if sGrade(text).isValidIntent() {
		intents = append(intents, intentGrade)
	}

	// get schedule intent
	if sSchedule(text).isValidIntent() {
		intents = append(intents, intentSchedule)
	}

	switch len(intents) {
	case 0:
		return intentUnknown, fmt.Errorf("Cannot get intent")
	case 1:
		return intents[0], nil
	default:
		return intentUnknown, fmt.Errorf("Detected more than 1 intents")
	}
}

// validate assistant intent
func (params sAssistant) isValidIntent() bool {
	pattern := []string{
		`asisten`,
		`assistant`,
		`pengajar`,
	}

	for _, val := range pattern {
		if regexp.MustCompile(val).MatchString(string(params)) {
			return true
		}
	}

	return false
}

// validate assignment intent
func (params sAssignment) isValidIntent() bool {
	pattern := []string{
		`tugas`,
		`pr`,
		`laprak`,
		`laporan`,
		`assignment`,
		`task`,
		`pekerjaan\s*rumah`,
		`pekerjaan`,
	}

	for _, val := range pattern {
		if regexp.MustCompile(val).MatchString(string(params)) {
			return true
		}
	}

	return false
}

// validate information intent
func (params sInformation) isValidIntent() bool {
	pattern := []string{
		`informasi`,
		`berita`,
		`hot`,
		`\s*trend`,
	}

	for _, val := range pattern {
		if regexp.MustCompile(val).MatchString(string(params)) {
			return true
		}
	}

	return false
}

// validate grade intent
func (params sGrade) isValidIntent() bool {
	pattern := []string{
		`nilai`,
		`score`,
		`grade`,
	}

	for _, val := range pattern {
		if regexp.MustCompile(val).MatchString(string(params)) {
			return true
		}
	}

	return false
}

// validate schedule intent
func (params sSchedule) isValidIntent() bool {
	pattern := []string{
		`jadwal`,
		`agenda`,
		`kegiatan`,
	}

	for _, val := range pattern {
		if regexp.MustCompile(val).MatchString(string(params)) {
			return true
		}
	}

	return false
}
