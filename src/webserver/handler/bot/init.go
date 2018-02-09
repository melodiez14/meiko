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
			"Ini dia!",
			"Sikat gan!",
			"Ini dia bro!",
			"Ieu wa!",
			"Aha!",
			"Here it is!",
			"Easy..",
			"Kuy!",
			"Cus!",
			"Ini dia yang kamu maksud",
			"Sundul Gan",
			"Sikat Mang",
			"Tarik Mang",
			"Ini nih jawabannya",
			"Pasti ini yang kamu maksud",
			"Silahkan ini dia",
		},
		"doubt": []string{
			"Kayaknya sih",
			"Mereun ya ini ge",
			"Apakah ini yang kamu maksud?",
			"Apakah ini jawabannya?",
			"Mungkin ini jawabannya",
			"Mungkin ini yang kamu maksud",
			"Sepertinya aku tahu",
			"Sepertinya ini deh gan",
			"Sepertinya ini jawabannya",
			"Sepertinya ini yang kamu maksud",
			"Barangkali ini gan?",
			"Semoga ini yang kamu maksud",
			"Inikan yang kamu maksud",
			"Jeng jeng jeng...",
		},
		"notsure": []string{
			"sorry nih bro, gua ga tau",
			"kasih tau ga ya?",
			"Kamu ngomong apa sih?",
			"Gatau, coba tanya ke temen kamu!",
			"Wah, aku belum tahu mengenai itu",
			"Pertanyaan yang bagus sekali. Tapi aku tidak tahu tuh",
			"Maaf nih, bot gak tau maksud pertanyaan kamu",
			"Ada yang mau kamu tanyakan lagi?",
			"Aduh maksud kamu apa ya? hehe :D",
			"Bisa diperjelas pertanyaannya :)",
			"Maafin bot ya, bot gak ngerti :(",
			"Maaf kaka bot tidak tahu maksud kaka",
			"Maafin bot kaka, bot gak ngerti :(",
		},
	}

	msgGreet = []greeting{
		greeting{text: "Hello my friend", isExistName: false},
		greeting{text: "Hey", isExistName: false},
		greeting{text: "Whatsup bro? What can i help?", isExistName: false},
		greeting{text: "Hola hola", isExistName: false},
		greeting{text: "Sup?", isExistName: false},
		greeting{text: "Hello %s", isExistName: true},
		greeting{text: "Hai Hai", isExistName: false},
		greeting{text: "oioi %s", isExistName: true},
		greeting{text: "Ask me anything", isExistName: false},
		greeting{text: "Heyhoo..", isExistName: false},
		greeting{text: "Hei", isExistName: false},
		greeting{text: "Wah ada apa nih?", isExistName: false},
		greeting{text: "Wah ada apa nih %s?", isExistName: true},
		greeting{text: "Hello! %s", isExistName: true},
		greeting{text: "Hello", isExistName: false},
		greeting{text: "Ada yang bisa dibantu?", isExistName: false},
		greeting{text: "Ada yang bisa dibantu %s?", isExistName: true},
		greeting{text: "Hallo, ada yang bisa bot bantu?", isExistName: false},
		greeting{text: "Hei, ada yang bisa bot bantu?", isExistName: false},
		greeting{text: "Oioi, ada yang bisa bot bantu?", isExistName: false},
		greeting{text: "Apa yg bisa dibantu %s?", isExistName: true},
	}

	msgAboutBot = []string{
		"Pengen tau aja atau pengen tau banget?",
		"Aku robot",
		"Aku owl asisten. aku bisa ngasih tau kamu mengenai jadwal, berita, tugas, jadwal, dan asisten",
		"Aku adalah Owl Assistant, aku bisa memberi informasi tentang kegiatan praktikum",
		"Aku adalah teman baik kamu yang selalu ada setiap saat",
		"Aku adalah teman mu",
		"Aku adalah sahabat kamu",
		"Aku adalah chat bot praktikum",
		"Aku Owl Assistant!",
	}

	msgAboutStudent = []student{
		student{text: "Kamu tuh orang paling pinter yang pernah aku kenal", isExistName: false},
		student{text: "My best friend", isExistName: false},
		student{text: "Anda adalah %s", isExistName: true},
		student{text: "Kamu adalah mahasiswa Teknik Informatika Unpad", isExistName: false},
		student{text: "Hai %s kamu adalah mahasiswa Teknik Informatika Unpad", isExistName: true},
		student{text: "Kamu, iya kamu!", isExistName: false},
		student{text: "Kamu siapa?", isExistName: false},
		student{text: "Kamu adalah anak ibu bapak mu..", isExistName: false},
		student{text: "Kamu bambang kan? hehe, masa gak tau sih %s", isExistName: true},
	}

	msgAboutCreator = []string{
		"Tau meiko ga? dia tuh yang bikin aku!",
		"Yang Maha Pencipta",
		"My creator",
		"Aku dibuat oleh Meiko Team",
		"Dari siapa aja yg penting aku bisa membantu kamu",
		"Aku dibuat Meiko",
		"Meiko adalah pemuatku",
		"Aku dibuat oleh tuanku",
		"Penciptaku adalah Meiko",
		"Meiko, Meiko, Meiko...",
	}

	msgKidding = []string{
		"Ciyee yang lagi bercanda wkwk",
		"Wah kamu bercanda deh...",
		"Kamu lucu ih :p",
		"Kamu bisa bercanda juga",
		"Jangan bercanda deh..",
		"Jawab gak yaa...",
		"Wkwkwkw",
		"Sa ae lu",
		"Hehehe",
		"Ckckckck",
		"Mau tau ajah? apa mau tau banget?",
		"Emang gue pikirin :p",
	}

	msgAssistant = []string{
		"Belum ada pengajar, tunggu yak",
		"Pengajar kamu sedang di atur oleh sistem",
		"Sabar yak belum ada pengajar",
		"Asisten belum ada nih, tunggu ya..",
		"Asisten sedang diatur oleh owl, tunggu sebentar ya",
		"List asisten yang kamu minta belum tersedia, tunggu ya sedang bot atur",
		"Belum ada list asisten nih bro, tunggu yaa",
		"Wait wait, tunggu sebentar ya sedang bot atur",
	}
	msgAssignment = []string{
		"Yes, tidak ada assignment, kamu boleh liburan",
		"Belum ada assignment, kamu boleh main sama pacar kamu",
		"Tidak ada tugas biat sekarang, kamu bisa melakukan hobi kamu sekarang",
		"Tidak ada tugas hari ini, selamat berlibur :D",
		"Selamat.. semua tugasmu sudah diselesaikan :D",
		"Selamat.. semua tugasmu sudah diselesaikan :D, selamat berlibur",
		"Kamu sudah menyelesaikan tugas-tugas, selamat berlibur",
		"Sudah tidak ada tugas untuk sekarang, selamat beristirahat",
		"Semua tugasmu sudah diselesaikan.. Horee",
	}

	msgInformation = []string{
		"Wah, tidak ada informasi untuk sekarang, nanti aku kabarin via notifikasi ya",
		"Informasi belum ada, nanti aku kabarin via notifikasi ya",
		"Belum ada sayang, selamat bersenang-senang",
		"Informasi yang kamu butuhkan belum ada nih",
		"Belum ada informasi untuk sekarang, jangan lupa terus mengupdate info kita ya, atau nanti aku kabarin via notifikasi :D",
		"Informasi sekarang belum tersedia, selamat bersenag-senang, tapi jangan lupa pantau terus owl ya",
		"Wait wait, tunggu sebentar ya sedang bot atur",
	}

	msgGrade = []string{
		"Nilai kamu belum keluar, tunggu ya nanti aku kabarin",
		"Nilai kamu belum keluar, santai nilai kamu pasti bagus",
		"Wah dengan sangat terpaksa nilai kamu belum dinilai asisten, tunggu ya, nilai kamu pasti bagus kok",
		"Nilai kamu masih belum ada di database nih, tunggu ya :D",
		"Wah, nilai kamu masih belum keluar nih, tunggu sebentar yaa :D",
		"Nilai yang kamu minta masih belum ada, tunggu sebentar lagi ya",
		"Tunggu sebentar ya, nilai kamu belum ada :)",
		"Sorry bro, nilai kamu belum ada nih, santai ae ya :)",
		"Maafin bro, nilai kamu belum ada nih, santai ae ya :)",
		"Maafin bro, nilai kamu masih belum keluar nih, tunggu sebentar yaa :D",
		"Wait wait, tunggu sebentar ya sedang bot atur",
		"Mohon maaf dengan sangat terpaksa nilai kamu belum dinilai asisten, hehe, tunggu sebentar ya, nilai kamu pasti bagus kok",
	}

	msgSchedule = []string{
		"yes, tidak ada praktikum. kamu boleh bersenang-senang",
		"Alhamdulilah tidak ada praktikum. yuk perbanyak berdzikir",
		"Hari ini tidak ada jadwal, yeeee.. selamat berlibur :)",
		"Hari ini tidak ada jadwal, yeeee.. jangan lupa bahagia",
		"Tidak ada praktikum nih, kamu bisa berlibur sejenak, jangan lupa kerjain tugas-tugas yaa...",
		"Yes, tidak ada praktikum, jangan lupa kerjain tugas-tugas yaa...",
		"Yes, tidak ada praktikum, selamat berlibur, jangan lupa kerjain tugas yang belum ya",
		"Kamu bisa berlibur, tidak ada praktikum",
		"Wait wait, tunggu sebentar ya sedang bot atur",
		"Selamat berlibur, tidak ada jadwal praktikum untuk kamu",
		"Selamat berlibur, tidak ada jadwal praktikum untuk kamu, jangan lupa kerjain tugas-tugas yaa...",
	}
}
