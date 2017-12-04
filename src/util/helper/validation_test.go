package helper

import (
	"testing"
)

func TestIsAlpha(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				text: "abcdefgh",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				text: "ABCDefgh",
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				text: "abcd EFGH",
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				text: "1234abcd",
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				text: "abcd!@#$",
			},
			want: false,
		},
		{
			name: "Test Case 6",
			args: args{
				text: "1234!@#$",
			},
			want: false,
		},
		{
			name: "Test Case 7",
			args: args{
				text: " ",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlpha(tt.args.text); got != tt.want {
				t.Errorf("IsAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlphaSpace(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				text: "abcdefgh",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				text: "abcdEFGH",
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				text: "abcd EFGH",
			},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				text: "1234abcd",
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				text: "abcd 1234",
			},
			want: false,
		},
		{
			name: "Test Case 6",
			args: args{
				text: "!@#$abcd",
			},
			want: false,
		},
		{
			name: "Test Case 7",
			args: args{
				text: "!@#$ abcd",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlphaSpace(tt.args.text); got != tt.want {
				t.Errorf("IsAlphaSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlphaNumericSpace(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				text: "abcdefgh",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				text: "abcdEFGH",
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				text: "abcd EFGH",
			},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				text: "abcd 1234",
			},
			want: true,
		},
		{
			name: "Test Case 5",
			args: args{
				text: "abcd!@#$",
			},
			want: false,
		},
		{
			name: "Test Case 6",
			args: args{
				text: "abcd !@@##",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlphaNumericSpace(tt.args.text); got != tt.want {
				t.Errorf("IsAlphaNumericSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				phone: "81220058838",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				phone: "081220058838",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				phone: "+6281220058838",
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				phone: "6281220058838",
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				phone: "abcdefgh",
			},
			want: false,
		},
		{
			name: "Test Case 6",
			args: args{
				phone: "812200588",
			},
			want: true,
		},
		{
			name: "Test Case 7",
			args: args{
				phone: "81220058",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPhone(tt.args.phone); got != tt.want {
				t.Errorf("IsPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				email: "khairil@gmail.com",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				email: "khairil@mail.unpad.com",
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				email: "khairil@yahoo.com",
			},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				email: "khairil@.mail.unpad.com.com",
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				email: "kh@iril@gmail.com",
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				email: " ",
			},
			want: false,
		},
		{
			name: "Test Case 6",
			args: args{
				email: "",
			},
			want: false,
		},
		{
			name: "Test Case 7",
			args: args{
				email: "Khairil_Azmi@gmail.com",
			},
			want: true,
		},
		{
			name: "Test Case 8",
			args: args{
				email: "Khairil.Azmi@gmail.com",
			},
			want: true,
		},
		{
			name: "Test Case 9",
			args: args{
				email: "Khairil*Azmi@gmail.com",
			},
			want: true,
		},
		{
			name: "Test Case 10",
			args: args{
				email: "Khairil@Azmi@gmail.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmail(tt.args.email); got != tt.want {
				t.Errorf("IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				password: "abcdefgh",
			},
			want: false,
		},
		{
			name: "Test Case 2",
			args: args{
				password: "12345678",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				password: "!@#$%^&*",
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				password: "abcd1234",
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				password: "abcd!@#$",
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				password: "1234!@#$",
			},
			want: false,
		},
		{
			name: "Test Case 6",
			args: args{
				password: "1234abcdDEFG",
			},
			want: true,
		},
		{
			name: "Test Case 7",
			args: args{
				password: "`@)()+_+?><::;`",
			},
			want: false,
		},
		{
			name: "Test Case 8",
			args: args{
				password: "<HTML>asep12</html>",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPassword(tt.args.password); got != tt.want {
				t.Errorf("IsPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				text: "abc",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				text: "abcAb def",
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				text: "abc821 def3422",
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				text: "sdkala. 21ok la la  ",
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				text: " ",
			},
			want: false,
		},
		{
			name: "Test Case 6",
			args: args{
				text: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.text); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsImageMime(t *testing.T) {
	type args struct {
		mime string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				mime: "",
			},
			want: false,
		},
		{
			name: "Test Case 2",
			args: args{
				mime: " ",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				mime: ".",
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				mime: "abc",
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				mime: "image/jpg",
			},
			want: false,
		},
		{
			name: "Test Case 6",
			args: args{
				mime: "image/png",
			},
			want: true,
		},
		{
			name: "Test Case 7",
			args: args{
				mime: "image/jpeg",
			},
			want: true,
		},
		{
			name: "Test Case 8",
			args: args{
				mime: "image/pdf",
			},
			want: false,
		},
		{
			name: "Test Case 9",
			args: args{
				mime: "image/bmp",
			},
			want: false,
		},
		{
			name: "Test Case 8",
			args: args{
				mime: "jpg/image",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsImageMime(tt.args.mime); got != tt.want {
				t.Errorf("IsImageMime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsImageExtension(t *testing.T) {
	type args struct {
		extension string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				extension: "",
			},
			want: false,
		},
		{
			name: "Test Case 2",
			args: args{
				extension: " ",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				extension: ".",
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				extension: "abc",
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				extension: "jpg",
			},
			want: true,
		},
		{
			name: "Test Case 6",
			args: args{
				extension: "image.png",
			},
			want: false,
		},
		{
			name: "Test Case 7",
			args: args{
				extension: "image.jpeg",
			},
			want: false,
		},
		{
			name: "Test Case 8",
			args: args{
				extension: "image.pdf",
			},
			want: false,
		},
		{
			name: "Test Case 9",
			args: args{
				extension: "image.jpg.pdf",
			},
			want: false,
		},
		{
			name: "Test Case 10",
			args: args{
				extension: ".jpg.image",
			},
			want: false,
		},
		{
			name: "Test Case 11",
			args: args{
				extension: ".jpgimage",
			},
			want: false,
		},
		{
			name: "Test Case 12",
			args: args{
				extension: "image.pdf.png",
			},
			want: false,
		},
		{
			name: "Test Case 12",
			args: args{
				extension: "!@$$",
			},
			want: false,
		},
		{
			name: "Test Case 13",
			args: args{
				extension: "jpeg",
			},
			want: true,
		},
		{
			name: "Test Case 14",
			args: args{
				extension: "png",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsImageExtension(tt.args.extension); got != tt.want {
				t.Errorf("IsImageExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	type args struct {
		text   string
		format func(string) bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				text:   "khairil    azmi",
				format: IsAlphaSpace,
			},
			want:    "khairil azmi",
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				text:   "khairil    azmi    ",
				format: IsAlphaSpace,
			},
			want:    "khairil azmi",
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				text:   "    khairil    azmi",
				format: IsAlphaSpace,
			},
			want:    "khairil azmi",
			wantErr: false,
		},
		{
			name: "Test Case 4",
			args: args{
				text:   "    khairil    azmi    ",
				format: IsAlphaSpace,
			},
			want:    "khairil azmi",
			wantErr: false,
		},
		{
			name: "Test Case 5",
			args: args{
				text:   "khairilazmi",
				format: IsAlphaSpace,
			},
			want:    "khairilazmi",
			wantErr: false,
		},
		{
			name: "Test Case 6",
			args: args{
				text:   "khairil14001",
				format: IsAlphaSpace,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test Case 7",
			args: args{
				text:   "khairil 14001    !  @ ",
				format: IsAlphaSpace,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Normalize(tt.args.text, tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeNPM(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				str: "",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test Case 2",
			args: args{
				str: "14081014006",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				str: "a408101400600",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test Case 4",
			args: args{
				str: "140810140060",
			},
			want:    140810140060,
			wantErr: false,
		},
		{
			name: "Test Case 5",
			args: args{
				str: "            ",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeNPM(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizeNPM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NormalizeNPM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeIdentity(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				str: "",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test Case 2",
			args: args{
				str: "140810140060",
			},
			want:    140810140060,
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				str: " ",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test Case 4",
			args: args{
				str: "140810140060",
			},
			want:    140810140060,
			wantErr: false,
		},
		{
			name: "Test Case 5",
			args: args{
				str: "1207261801970005",
			},
			want:    1207261801970005,
			wantErr: false,
		},
		{
			name: "Test Case 6",
			args: args{
				str: "1207261801970005aw",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test Case 7",
			args: args{
				str: "12072618 01970005",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test Case 8",
			args: args{
				str: "120726180197000512",
			},
			want:    120726180197000512,
			wantErr: false,
		},
		{
			name: "Test Case 9",
			args: args{
				str: "1207261801970005121",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test Case 10",
			args: args{
				str: "1207261801",
			},
			want:    1207261801,
			wantErr: false,
		},
		{
			name: "Test Case 11",
			args: args{
				str: "120726180",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeIdentity(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizeIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NormalizeIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				name: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test Case 2",
			args: args{
				name: "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				name: "1&*I(!$khairil azmi",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test Case 4",
			args: args{
				name: "abcd efgh",
			},
			want:    "abcd efgh",
			wantErr: false,
		},
		{
			name: "Test Case 5",
			args: args{
				name: "abcdefgh",
			},
			want:    "abcdefgh",
			wantErr: false,
		},
		{
			name: "Test Case 6",
			args: args{
				name: "  ",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizeName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NormalizeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeCollege(t *testing.T) {
	type args struct {
		college string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				college: "",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				college: "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				college: "1&*I(!$khairil azmi",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test Case 4",
			args: args{
				college: "abcd efgh",
			},
			want:    "abcd efgh",
			wantErr: false,
		},
		{
			name: "Test Case 5",
			args: args{
				college: "abcdefgh",
			},
			want:    "abcdefgh",
			wantErr: false,
		},
		{
			name: "Test Case 6",
			args: args{
				college: "  ",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeCollege(tt.args.college)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizeCollege() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NormalizeCollege() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				email: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test Case 2",
			args: args{
				email: " ",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				email: "a@gmail.com",
			},
			want:    "a@gmail.com",
			wantErr: false,
		},
		{
			name: "Test Case 4",
			args: args{
				email: "a@gmail.com.com",
			},
			want:    "a@gmail.com.com",
			wantErr: false,
		},
		{
			name: "Test Case 5",
			args: args{
				email: "a@mail.unpad.ac.id",
			},
			want:    "a@mail.unpad.ac.id",
			wantErr: false,
		},
		{
			name: "Test Case 6",
			args: args{
				email: "abcdefgh",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test Case 7",
			args: args{
				email: "a14@mail.unpad.ac.id",
			},
			want:    "a14@mail.unpad.ac.id",
			wantErr: false,
		},
		{
			name: "Test Case 8",
			args: args{
				email: "ssfdsdfghijklmnabcdefghijklmn@mail.unpad.ac.id",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizeEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NormalizeEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				str: "",
			},
			want: "",
		},
		{
			name: "Test Case 2",
			args: args{
				str: " Risal Falah ",
			},
			want: "Risal Falah",
		},
		{
			name: "Test Case 3",
			args: args{
				str: " Risal        Falah ",
			},
			want: "Risal Falah",
		},
		{
			name: "Test Case 4",
			args: args{
				str: " Risal    $    Falah ",
			},
			want: "Risal $ Falah",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.args.str); got != tt.want {
				t.Errorf("Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidFileID(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "12345689023.123123.123.123",
			args: args{
				str: "12345689023.123123.123.123",
			},
			want: true,
		},
		{
			name: "1",
			args: args{
				str: "1",
			},
			want: false,
		},
		{
			name: "abcdefghijklmnopqrstuvwxyz",
			args: args{
				str: "abcdefghijklmnopqrstuvwxyz",
			},
			want: false,
		},
		{
			name: "abcdefghijklmnopqrstuvwxyz12355",
			args: args{
				str: "abcdefghijklmnopqrstuvwxyz12355",
			},
			want: false,
		},
		{
			name: "123456789012345678901234567890",
			args: args{
				str: "123456789012345678901234567890",
			},
			want: true,
		},
		{
			name: "123456789012345678901234567890123",
			args: args{
				str: "123456789012345678901234567890123",
			},
			want: false,
		},
		{
			name: "1234567890123.123252123.jpg",
			args: args{
				str: "1234567890123.123252123.jpg",
			},
			want: false,
		},
		{
			name: "123456-1234567-123456754",
			args: args{
				str: "123456-1234567-123456754",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidFileID(tt.args.str); got != tt.want {
				t.Errorf("IsValidFileID() = %v, want %v", got, tt.want)
			}
		})
	}
}
