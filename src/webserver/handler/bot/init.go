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
	initMessage()
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

func initMessage() {
	msgConf = map[string][]string{
		"confident": []string{
			"Ini bro jawabannya",
			"Ini dia",
			"Sikat gan!",
		},
		"doubt": []string{
			"Kayaknya sih",
			"Mereun ya ini ge",
		},
		"notsure": []string{
			"sorry nih bro, gua ga tau",
			"kasih tau ga ya?",
			"Kamu ngomong apa sih?",
			"Gatau, coba tanya ke temen kamu!",
		},
	}

	msgGreet = []greeting{
		greeting{text: "Hello my friend", isExistName: false},
		greeting{text: "Hey", isExistName: false},
		greeting{text: "Whatsup bro? What can i help?", isExistName: false},
		greeting{text: "Hola hola", isExistName: false},
		greeting{text: "Sup?", isExistName: false},
		greeting{text: "Hello %s", isExistName: true},
		greeting{text: "Hey %s. My buddy", isExistName: true},
	}

	msgAboutBot = []string{
		"Pengen tau aja atau pengen tau banget?",
		"Aku robot",
		"Aku owl asisten. aku bisa ngasih tau kamu mengenai jadwal, berita, tugas, jadwal, dan asisten",
	}

	msgAboutStudent = []student{
		student{text: "Kamu tuh orang paling pinter yang pernah aku kenal", isExistName: false},
		student{text: "My best friend", isExistName: false},
		student{text: "Anda adalah %s", isExistName: true},
	}

	msgAboutCreator = []string{
		"Tau meiko ga? dia tuh yang bikin aku!",
		"Yang Maha Pencipta",
		"My creator",
	}

	msgKidding = []string{
		"Ciyee yang lagi bercanda wkwk",
	}
}
