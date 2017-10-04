package helper

import (
	"fmt"
	"strings"
	"testing"
)

func TestIsAlphaSpace(t *testing.T) {
	var (
		IsAlphaSpace1 string = "khaazas"
		IsAlphaSpace2 string = "kha azas"
		IsAlphaSpace3 string = "khaa zas!@"
	)

	fmt.Println("Is Alpha Space Case 1 : %s", IsAlphaSpace1)
	if IsAlphaSpace(IsAlphaSpace1) != true {
		t.Errorf("%v is Wrong! must be true", IsAlphaSpace(IsAlphaSpace1))
	}

	fmt.Println("Is Alpha Space Case 2 : %s", IsAlphaSpace2)
	if IsAlphaSpace(IsAlphaSpace2) != true {
		t.Errorf("%v is Wrong! must be true ", IsAlphaSpace(IsAlphaSpace2))
	}

	fmt.Println("Is Alpha Space Case 3 : %s", IsAlphaSpace3)
	if IsAlphaSpace(IsAlphaSpace3) != false {
		t.Errorf("%v is Wrong! must be false", IsAlphaSpace(IsAlphaSpace3))
	}
}

func TestIsAlphaNumericSpace(t *testing.T) {
	var (
		IsAlphaNumSpace1 string = "khairil14001"
		IsAlphaNumSpace2 string = "khairil 14001"
		IsAlphaNumSpace3 string = "khairil 14001!"
	)

	fmt.Println("Is Alpha Num Space Case 1 : %s", IsAlphaNumSpace1)
	if IsAlphaNumericSpace(IsAlphaNumSpace1) != true {
		t.Errorf("%v is Wrong! must be true", IsAlphaNumericSpace(IsAlphaNumSpace1))
	}

	fmt.Println("Is Alpha Num Space Case 2 : %s", IsAlphaNumSpace2)
	if IsAlphaNumericSpace(IsAlphaNumSpace2) != true {
		t.Errorf("%v is Wrong! must be true", IsAlphaNumericSpace(IsAlphaNumSpace2))
	}

	fmt.Println("Is Alpha Num Space Case 3 : %s", IsAlphaNumSpace3)
	if IsAlphaNumericSpace(IsAlphaNumSpace3) != false {
		t.Errorf("%v is Wrong! must be false", IsAlphaNumericSpace(IsAlphaNumSpace3))
	}
}

func TestIsPhone(t *testing.T) {
	var (
		IsPhone1 string = "082214467300"
		IsPhone2 string = "082214467"
		IsPhone3 string = "0822144673000"
		IsPhone4 string = "08221446730O"
	)

	fmt.Println("Is Phone Number Case 1 : %s", IsPhone1)
	if IsPhone(IsPhone1) != true {
		t.Errorf("%v is Wrong! must be true", IsPhone(IsPhone1))
	}

	fmt.Println("Is Phone Number Case 2 : %s", IsPhone2)
	if IsPhone(IsPhone2) != false {
		t.Errorf("%v is Wrong! must be false", IsPhone(IsPhone2))
	}

	fmt.Println("Is Phone Number Case 3 : %s", IsPhone3)
	if IsPhone(IsPhone3) != false {
		t.Errorf("%v is Wrong! must be false", IsPhone(IsPhone3))
	}

	fmt.Println("Is Phone Number Case 4 : %s", IsPhone4)
	if IsPhone(IsPhone4) != false {
		t.Errorf("%v is Wrong! must be false", IsPhone(IsPhone4))
	}
}

func TestIsEmail(t *testing.T) {
	var (
		IsEmail1 string = "khairil14001@mail.unpad.ac.id"
		IsEmail2 string = "khairil 14001@mail.unpad.ac.id"
		IsEmail3 string = "khairil14001@mail"
		IsEmail4 string = "khairil14001!!!@mail.unpad.ac.id"
	)

	fmt.Println("Is Email Case 1 : %s", IsEmail1)
	if IsEmail(IsEmail1) != true {
		t.Errorf("%v is Wrong! must be true", IsEmail(IsEmail1))
	}

	fmt.Println("Is Email Case 2 : %s", IsEmail2)
	if IsEmail(IsEmail2) != false {
		t.Errorf("%v is Wrong! must be false", IsEmail(IsEmail2))
	}

	fmt.Println("Is Email Case 3 : %s", IsEmail3)
	if IsEmail(IsEmail3) != false {
		t.Errorf("%v is Wrong! must be false", IsEmail(IsEmail3))
	}

	fmt.Println("Is Email Case 4 : %s", IsEmail4)
	if IsEmail(IsEmail4) != false {
		t.Errorf("%v is Wrong! must be false", IsEmail(IsEmail4))
	}
}

func TestIsPassword(t *testing.T) {
	var (
		IsPassword1 string = "sokolaRimba"
		IsPassword2 string = "sokolaRimba14"
		IsPassword3 string = "sokola Rimba"
		IsPassword4 string = "sokolaRimba@"
	)

	fmt.Println("Is Password Case 1 : %s", IsPassword1)
	if IsPassword(IsPassword1) != false {
		t.Errorf("%v is Wrong! must be true", IsPassword(IsPassword1))
	}

	fmt.Println("Is Password Case 2 : %s", IsPassword2)
	if IsPassword(IsPassword2) != true {
		t.Errorf("%v is Wrong! must be true", IsPassword(IsPassword2))
	}

	fmt.Println("Is Password Case 3 : %s", IsPassword3)
	if IsPassword(IsPassword3) != false {
		t.Errorf("%v is Wrong! must be false", IsPassword(IsPassword3))
	}

	fmt.Println("Is Password Case 4 : %s", IsPassword4)
	if IsPassword(IsPassword4) != false {
		t.Errorf("%v is Wrong! must be false", IsPassword(IsPassword4))
	}
}

func TestIsEmpty(t *testing.T) {
	var (
		IsEmpty1 string = " "
		IsEmpty2 string = ""
		IsEmpty3 string = "x"
	)

	fmt.Println("Is Empty Case 1 : %s", IsEmpty1)
	if IsEmpty(IsEmpty1) != true {
		t.Errorf("%v is Wrong! must be true", IsEmpty(IsEmpty1))
	}

	fmt.Println("Is Empty Case 2 : %s", IsEmpty2)
	if IsEmpty(IsEmpty2) != true {
		t.Errorf("%v is Wrong! must be true", IsEmpty(IsEmpty2))
	}

	fmt.Println("Is Empty Case 3 : %s", IsEmpty3)
	if IsEmpty(IsEmpty3) != false {
		t.Errorf("%v is Wrong! must be false", IsEmpty(IsEmpty3))
	}
}

func TestIsImageMime(t *testing.T) {
	var (
		IsImageMime1 string = "image/jpeg"
		IsImageMime2 string = "image/png"
		IsImageMime3 string = "image/pdf"
		IsImageMime4 string = "image.png"
	)

	fmt.Println("Is Image Mime Case 1 : %s", IsImageMime1)
	if IsImageMime(IsImageMime1) != true {
		t.Errorf("%v is Wrong! must be true", IsImageMime(IsImageMime1))
	}

	fmt.Println("Is Image Mime Case 2 : %s", IsImageMime2)
	if IsImageMime(IsImageMime2) != true {
		t.Errorf("%v is Wrong! must be true", IsImageMime(IsImageMime2))
	}

	fmt.Println("Is Image Mime Case 3 : %s", IsImageMime3)
	if IsImageMime(IsImageMime3) != false {
		t.Errorf("%v is Wrong! must be false", IsImageMime(IsImageMime3))
	}

	fmt.Println("Is Image Mime Case 4 : %s", IsImageMime4)
	if IsImageMime(IsImageMime4) != false {
		t.Errorf("%v is Wrong! must be false", IsImageMime(IsImageMime4))
	}
}

func TestIsImageExtension(t *testing.T) {
	var (
		IsImageExtension1 string = "jpeg"
		IsImageExtension2 string = "jpg"
		IsImageExtension3 string = "pdf"
	)

	fmt.Println("Is Image Extension Case 1 : %s", IsImageExtension1)
	if IsImageExtension(IsImageExtension1) != true {
		t.Errorf("%v is Wrong! must be true", IsImageExtension(IsImageExtension1))
	}

	fmt.Println("Is Image Extension Case 2 : %s", IsImageExtension2)
	if IsImageExtension(IsImageExtension2) != true {
		t.Errorf("%v is Wrong! must be true", IsImageExtension(IsImageExtension2))
	}

	fmt.Println("Is Image Extension Case 3 : %s", IsImageExtension3)
	if IsImageExtension(IsImageExtension3) != false {
		t.Errorf("%v is Wrong! must be false", IsImageExtension(IsImageExtension3))
	}
}

func TestNormalize(t *testing.T) {
	var (
		Normalize1      string = "Khairil  Azmi  Ashari"
		Normalize1_want string = "Khairil Azmi Ashari"
		Normalize2      string = "Khairil  Azmi  Ashari  "
		Normalize2_want string = "Khairil Azmi Ashari"
	)

	fmt.Println("Normalize Case 1 : %s", Normalize1)
	var Normal1, stat1 = Normalize(Normalize1, IsAlphaSpace)
	t.Logf("Text Normalize: %s", Normal1)
	if strings.Compare(Normalize1_want, Normal1) != 0 || stat1 != nil {
		t.Errorf("Wrong! must be Normalize : %s, error : %v", Normal1, nil)
	}

	fmt.Println("Normalize Case 2 : %s", Normalize2)
	var Normal2, stat2 = Normalize(Normalize2, IsAlphaSpace)
	t.Logf("Text Normalize: %s", Normal2)
	if strings.Compare(Normalize2_want, Normal2) != 0 || stat2 != nil {
		t.Errorf("Wrong! must be Normalize : %s, error : %v", Normal1, nil)
	}
}

func TestNormalizeIdentity(t *testing.T) {
	var (
		NormalizeIdentity1      string = "140810140060"
		NormalizeIdentity1_want int64  = 140810140060
		NormalizeIdentity2      string = "1408101400"
		NormalizeIdentity3      string = ""
		NormalizeIdentity4      string = "14081014006O"
	)

	fmt.Println("Normalize Identity Case 1 : %s", NormalizeIdentity1)

	var userIdentity1, stat1 = NormalizeIdentity(NormalizeIdentity1)
	if (userIdentity1 != NormalizeIdentity1_want) && (stat1 != nil) {
		t.Errorf("%d is Wrong! must be %d", userIdentity1, NormalizeIdentity1)
	}

	fmt.Println("Normalize User ID Case 2 : %s", NormalizeIdentity2)

	var _, stat2 = NormalizeIdentity(NormalizeIdentity2)
	if stat2 == nil {
		t.Errorf("%v is Wrong! Error must be not nil", stat2)
	}

	fmt.Println("Normalize Identity Case 3 : %s", NormalizeIdentity3)

	var _, stat3 = NormalizeIdentity(NormalizeIdentity3)
	if stat3 == nil {
		t.Errorf("%v is Wrong! Error must be not nil", stat3)
	}

	fmt.Println("Normalize Identity Case 4 : %s", NormalizeIdentity4)

	var _, stat4 = NormalizeIdentity(NormalizeIdentity4)
	if stat4 == nil {
		t.Errorf("%v is Wrong! Error must be not nil", stat4)
	}
}

func TestNormalizeName(t *testing.T) {
	var (
		NormalizeName1 string = "khairil azmi"
		NormalizeName2 string = "khazasX1234"
		NormalizeName3 string = ""
	)

	fmt.Println("Normalize Name Case 1 : %s", NormalizeName1)

	var _, stat1 = NormalizeName(NormalizeName1)
	if stat1 != nil {
		t.Errorf("%v is Wrong! must be nil", stat1)
	}

	fmt.Println("Normalize Name Case 2 : %s", NormalizeName2)

	var _, stat2 = NormalizeName(NormalizeName2)
	if stat2 == nil {
		t.Errorf("%v is Wrong! must be not nil", stat2)
	}

	fmt.Println("Normalize Name Case 3 : %s", NormalizeName3)

	var _, stat3 = NormalizeName(NormalizeName3)
	if stat3 == nil {
		t.Errorf("%v is Wrong! must be not nil", stat3)
	}
}

func TestNormalizeCollege(t *testing.T) {
	var (
		NormalizeCollege1 string = "Universitas Padjadjaran"
		NormalizeCollege2 string = " "
		NormalizeCollege3 string = ""
	)

	fmt.Println("Normalize Collage Case 1 : %s", NormalizeCollege1)

	var _, stat1 = NormalizeCollege(NormalizeCollege1)
	if stat1 != nil {
		t.Errorf("%v is Wrong! must be nil", stat1)
	}

	fmt.Println("Normalize Collage Case 2 : %s", NormalizeCollege2)

	var _, stat2 = NormalizeCollege(NormalizeCollege2)
	if stat2 == nil {
		t.Errorf("%v is Wrong! must be not nil", stat2)
	}

	fmt.Println("Normalize Collage Case 3 : %s", NormalizeCollege3)

	var _, stat3 = NormalizeCollege(NormalizeCollege3)
	if stat3 == nil {
		t.Errorf("%v is Wrong! must be not nil", stat3)
	}
}

func TestNormalizeEmail(t *testing.T) {
	var (
		NormalizeEmail1 string = ""
		NormalizeEmail2 string = " "
		NormalizeEmail3 string = "khairil@gmail.com"
		NormalizeEmail4 string = "khairil@yahoo.com"
	)

	fmt.Println("Normalize Email Case 1 : %s", NormalizeEmail1)

	var _, stat1 = NormalizeEmail(NormalizeEmail1)
	if stat1 == nil {
		t.Errorf("%v is Wrong! must be not nil", stat1)
	}

	fmt.Println("Normalize Email Case 2 : %s", NormalizeEmail2)

	var _, stat2 = NormalizeEmail(NormalizeEmail2)
	if stat2 == nil {
		t.Errorf("%v is Wrong! must be not nil", stat1)
	}

	fmt.Println("Normalize Email Case 3 : %s", NormalizeEmail3)

	var _, stat3 = NormalizeEmail(NormalizeEmail3)
	if stat3 != nil {
		t.Errorf("%v is Wrong! must be nil", stat3)
	}

	fmt.Println("Normalize Email Case 4 : %s", NormalizeEmail4)

	var _, stat4 = NormalizeEmail(NormalizeEmail4)
	if stat4 != nil {
		t.Errorf("%v is Wrong! must be nil", stat4)
	}
}
