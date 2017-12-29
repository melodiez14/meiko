package bot

import (
	"fmt"
	"log"
	"strings"

	cs "github.com/melodiez14/meiko/src/module/course"
	usr "github.com/melodiez14/meiko/src/module/user"
	"github.com/melodiez14/meiko/src/util/helper"
)

// Init used to initialize the bot
func Init() {
	log.Println("Initializing Bot")
	initRgxAsistant()
	initRgxCourse()
	log.Println("Bot successfully initialized")
}

// initRgxAssistant gets assistant lists from database and put it into rgxassistant
func initRgxAsistant() {
	var name []string
	var nameFilter []string
	var nameGroup []string
	userID, err := cs.SelectAllAssistantID()
	if err != nil {
		log.Fatalf("Bot init error: cannot get all assistant id")
	}

	user, err := usr.SelectByID(userID, false, usr.ColName)
	if err != nil {
		log.Fatalf("Bot init error: cannot get all assistant name")
	}

	// make []string{"Risal Falah", "Rifki Muhammad"} into []string{"risal", "falah", "rifki", "muhammad"}
	for _, val := range user {
		str := strings.ToLower(val.Name)
		name = append(name, strings.Split(str, " ")...)
	}

	// filter course which has less than 5 character
	for _, val := range name {
		if len(val) >= 5 {
			nameFilter = append(nameFilter, val)
		}
	}

	// make []string{"risal", "falah"} into []string{"(risal)", "(falah)"} for regex purpose
	for _, val := range nameFilter {
		str := fmt.Sprintf("(%s)", val)
		if !helper.IsStringInSlice(str, nameGroup) {
			nameGroup = append(nameGroup, str)
		}
	}

	rgxAssistant = strings.Join(nameGroup, "|")
}

// initRgxCourse gets course lists from database and put it into rgxcourse
func initRgxCourse() {
	var name []string
	var nameFilter []string
	var nameGroup []string
	courseName, err := cs.SelectAllName()
	if err != nil {
		log.Fatalf("Bot init error: cannot get all assistant id")
	}

	// make []string{"Data Warehouse", "Algoritma dan Pemrograman"} into []string{"data", "warehouse", "algoritma", "dan", "pemrograman"}
	for _, val := range courseName {
		str := strings.ToLower(val)
		name = append(name, strings.Split(str, " ")...)
	}

	// filter course which has less than 5 character
	for _, val := range name {
		if len(val) >= 5 {
			nameFilter = append(nameFilter, val)
		}
	}

	// make []string{"data", "warehouse"} into []string{"(data)", "(warehouse)"} for regex purpose
	for _, val := range nameFilter {
		str := fmt.Sprintf("(%s)", val)
		if !helper.IsStringInSlice(str, nameGroup) {
			nameGroup = append(nameGroup, str)
		}
	}

	rgxCourse = strings.Join(nameGroup, "|")
}
