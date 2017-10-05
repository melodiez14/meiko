package helper

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

var ()

func TestStringToMD5(t *testing.T) {
	var (
		target1   string = "khaazas"
		MD5_want1 string = "338e19f5acace637a0c525ae510eec59"
		target2   string = "khairil azmi"
		MD5_want2 string = "ecc8f4b5e1e7613cf543f370b7ce53c6"
		target3   string = "khairil14001 !@123.,"
		MD5_want3 string = "7dd67fd22bd7bddc9e0fcb836466d047"
	)

	fmt.Println("String To MD5 Case  1 : %s", target1)
	t.Logf("MD5 : %s", StringToMD5(target1))

	if strings.Compare(MD5_want1, StringToMD5(target1)) != 0 {
		t.Errorf("Wrong! must be %s", MD5_want1)
	}

	fmt.Println("String To MD5 Case 2 : %s", target2)
	t.Logf("MD5 : %s", StringToMD5(target2))

	if strings.Compare(MD5_want2, StringToMD5(target2)) != 0 {
		t.Errorf("Wrong! must be %s", MD5_want2)
	}

	fmt.Println("String To MD5 Case 3 : %s", target3)
	t.Logf("MD5 : %s", StringToMD5(target3))

	if strings.Compare(MD5_want3, StringToMD5(target3)) != 0 {
		t.Errorf("Wrong! must be %s", MD5_want3)
	}
}

func TestExtractExtension(t *testing.T) {
	var (
		FileName1 string = "task.pdf"
		fn1_want  string = "task"
		ext1_want string = "pdf"
		FileName2 string = "task 1.pdf"
		fn2_want  string = "task 1"
		ext2_want string = "pdf"
		FileName3 string = "task.1.pdf"
		fn3_want  string = "task.1"
		ext3_want string = "pdf"
	)

	fmt.Println("Extract Extension Case 1 : %s", FileName1)

	var fn1, ext1, stat1 = ExtractExtension(FileName1)
	t.Logf("file name : %s, ext : %s  ", fn1, ext1)
	if strings.Compare(fn1, fn1_want) != 0 || strings.Compare(ext1, ext1_want) != 0 || stat1 != nil {
		t.Errorf("Wrong! must be file name : %s, ext : %s, error : %v", fn1_want, ext1_want, nil)
	}

	fmt.Println("Extract Extension Case 2 : %s", FileName2)

	var fn2, ext2, stat2 = ExtractExtension(FileName2)
	t.Logf("file name : %s, ext : %s  ", fn2, ext2)
	if strings.Compare(fn2, fn2_want) != 0 || strings.Compare(ext2, ext2_want) != 0 || stat2 != nil {
		t.Errorf("Wrong! must be file name : %s, ext : %s, error : %v", fn2_want, ext2_want, nil)
	}

	fmt.Println("Extract Extension Case 3 : %s", FileName3)

	var fn3, ext3, stat3 = ExtractExtension(FileName3)
	t.Logf("file name : %s, ext : %s  ", fn3, ext3)
	if strings.Compare(fn3, fn3_want) != 0 || strings.Compare(ext3, ext3_want) != 0 || stat3 != nil {
		t.Errorf("Wrong! must be file name : %s, ext : %s, error : %v", fn3_want, ext3_want, nil)
	}
}

func TestDateToString(t *testing.T) {
	var (
		time1_want string = "Just now"
		time2_want string = "5 minutes ago"
		time3_want string = "5 hours ago"
		time4_want string = "5 days ago"
	)

	t11_case := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	t12_case := time.Date(2009, time.November, 10, 23, 0, 5, 0, time.UTC)
	t21_case := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	t22_case := time.Date(2009, time.November, 10, 23, 5, 0, 0, time.UTC)
	t31_case := time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC)
	t32_case := time.Date(2009, time.November, 10, 5, 0, 0, 0, time.UTC)
	t41_case := time.Date(2009, time.November, 1, 23, 0, 0, 0, time.UTC)
	t42_case := time.Date(2009, time.November, 6, 23, 0, 0, 0, time.UTC)

	fmt.Println("Date To String Case 1")
	t.Logf("Time : %s", DateToString(t12_case, t11_case))
	if strings.Compare(DateToString(t12_case, t11_case), time1_want) != 0 {
		t.Errorf("Wrong! must be elapsed time : %s", time1_want)
	}

	fmt.Println("Date To String Case 2")
	t.Logf("Time : %s", DateToString(t22_case, t21_case))
	if strings.Compare(DateToString(t22_case, t21_case), time2_want) != 0 {
		t.Errorf("Wrong! must be elapsed time : %s", time2_want)
	}

	fmt.Println("Date To String Case 3")
	t.Logf("Time : %s", DateToString(t32_case, t31_case))
	if strings.Compare(DateToString(t32_case, t31_case), time3_want) != 0 {
		t.Errorf("Wrong! must be elapsed time : %s", time3_want)
	}

	fmt.Println("Date To String Case 4")
	t.Logf("Time : %s", DateToString(t42_case, t41_case))
	if strings.Compare(DateToString(t42_case, t41_case), time4_want) != 0 {
		t.Errorf("Wrong! must be elapsed time : %s", time4_want)
	}

}
