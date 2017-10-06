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
				text: "abc",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				text: "abcAdef",
			},
			want: true,
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
				text: "abc",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				text: "abcAdef",
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				text: "abcAbD def",
			},
			want: true,
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
				text: "abc",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				text: "abcAdef",
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				text: "abcAbD def",
			},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				text: "abc821 def3422",
			},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				text: "sdkala. 21ok la la  ",
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
				phone: "0",
			},
			want: false,
		},
		{
			name: "Test Case 2",
			args: args{
				phone: "082214467300",
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				phone: "0822144673000",
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				phone: "abc821 def3422",
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				phone: "sdkala. 21ok la la  ",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				phone: "0822144673000O",
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
				email: "khairil@.mail.unpad.com",
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
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				email: "Khairil_Azmi@gmail.com",
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
				password: "",
			},
			want: false,
		},
		{
			name: "Test Case 2",
			args: args{
				password: " ",
			},
			want: false,
		},
		{
			name: "Test Case 2",
			args: args{
				password: "a",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				password: "khaazas",
			},
			want: true,
		},
		{
			name: "Test Case 4",
			args: args{
				password: "kh4zass",
			},
			want: true,
		},
		{
			name: "Test Case 5",
			args: args{
				password: "a ouco34q8hw 8O",
			},
			want: false,
		},
		{
			name: "Test Case 6",
			args: args{
				password: "aAa7@4!",
			},
			want: false,
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
			want: true,
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
				mime: "image.jpg",
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
				extension: "image.jpg",
			},
			want: true,
		},
		{
			name: "Test Case 6",
			args: args{
				extension: "image.png",
			},
			want: true,
		},
		{
			name: "Test Case 7",
			args: args{
				extension: "image.jpeg",
			},
			want: true,
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
			want: true,
		},
		{
			name: "Test Case 12",
			args: args{
				extension: "image.xyz.png",
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
				text: "khairil    azmi",
				format: IsAlphaSpace,
			},
			want: "khairil azmi",
			wantErr:false,
		},
		{
			name: "Test Case 2",
			args: args{
				text: "khairil    azmi    ",
				format: IsAlphaSpace,
			},
			want: "khairil azmi",
			wantErr:false,
		},
		{
			name: "Test Case 3",
			args: args{
				text: "    khairil    azmi",
				format: IsAlphaSpace,
			},
			want: "khairil azmi",
			wantErr:false,
		},
		{
			name: "Test Case 4",
			args: args{
				text: "    khairil    azmi    ",
				format: IsAlphaSpace,
			},
			want: "khairil azmi",
			wantErr:false,
		},
		{
			name: "Test Case 5",
			args: args{
				text: "khairilazmi",
				format: IsAlphaSpace,
			},
			want: "khairil azmi",
			wantErr:true,
		},
		{
			name: "Test Case 6",
			args: args{
				text: "khairilazmi",
				format: IsAlphaSpace,
			},
			want: "khairilazmi",
			wantErr:false,
		},
		{
			name: "Test Case 7",
			args: args{
				text: "khairil14001",
				format: IsAlphaSpace,
			},
			want: "khairil14001",
			wantErr:false,
		},
		{
			name: "Test Case 8",
			args: args{
				text: "khairil 14001    !  @ ",
				format: IsAlphaSpace,
			},
			want: "khairil 14001 ! @",
			wantErr:false,
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
				str: "140810140060",
			},
			want: 140810140060,
			wantErr:false,
		},
		{
			name: "Test Case 2",
			args: args{
				str: "14081014006",
			},
			want: 14081014006,
			wantErr:true,
		},
		{
			name: "Test Case 3",
			args: args{
				str: "1408101400600",
			},
			want: 1408101400600,
			wantErr:true,
		},
		{
			name: "Test Case 4",
			args: args{
				str: "14081014006O",
			},
			want: 14081014006,
			wantErr:true,
		},
		{
			name: "Test Case 5",
			args: args{
				str: "",
			},
			want: 0,
			wantErr:true,
		},
		{
			name: "Test Case 2",
			args: args{
				str: " ",
			},
			want: 0,
			wantErr:true,
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
			want: 0,
			wantErr:true,
		},
		{
			name: "Test Case 2",
			args: args{
				str: "140810140060",
			},
			want: 140810140060,
			wantErr:false,
		},
		{
			name: "Test Case 3",
			args: args{
				str: " ",
			},
			want: 0,
			wantErr:true,
		},
		{
			name: "Test Case 4",
			args: args{
				str: "140810140060",
			},
			want: 140810140060,
			wantErr:false,
		},
		{
			name: "Test Case 5",
			args: args{
				str: "1207261801970005",
			},
			want: 1207261801970005,
			wantErr:false,
		},
		{
			name: "Test Case 6",
			args: args{
				str: "1207261801970005aw",
			},
			want: 1207261801970005,
			wantErr:true,
		},
		{
			name: "Test Case 7",
			args: args{
				str: "12072618 01970005",
			},
			want: 1207261801970005,
			wantErr:true,
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
				name: "khairil azmi",
			},
			want: "khairil azmi",
			wantErr:false,
		},
		{
			name: "Test Case 2",
			args: args{
				name: "khairil azmi   ",
			},
			want: "khairil azmi",
			wantErr:false,
		},
		{
			name: "Test Case 3",
			args: args{
				name: " khairil azmi",
			},
			want: "khairil azmi",
			wantErr:false,
		},
		{
			name: "Test Case 4",
			args: args{
				name: " khairil14 azmi",
			},
			want: "khairil14 azmi",
			wantErr:true,
		},
		{
			name: "Test Case 4",
			args: args{
				name: "",
			},
			want: "",
			wantErr:true,
		},
		{
			name: "Test Case 5",
			args: args{
				name: " ",
			},
			want: "",
			wantErr:true,
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
				college: "abc",
			},
			want: "abc",
			wantErr:false,
		},
		{
			name: "Test Case 2",
			args: args{
				college: "abcAdef",
			},
			want: "abcAdef",
			wantErr:false,
		},
		{
			name: "Test Case 3",
			args: args{
				college: "abcAbD def",
			},
			want: "abcAbD def",
			wantErr:false,
		},
		{
			name: "Test Case 4",
			args: args{
				college: "abc821 def3422",
			},
			want: "abc821 def3422",
			wantErr:false,
		},
		{
			name: "Test Case 4",
			args: args{
				college: "sdkala. 21ok la la  ",
			},
			want: "sdkala. 21ok la la",
			wantErr:true,
		},
		{
			name: "Test Case 5",
			args: args{
				college: "",
			},
			want: "",
			wantErr:true,
		},
		{
			name: "Test Case 6",
			args: args{
				college: " ",
			},
			want: " ",
			wantErr:true,
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
			want: "",
			wantErr:true,
		},
		{
			name: "Test Case 2",
			args: args{
				email: " ",
			},
			want: " ",
			wantErr:true,
		},
		{
			name: "Test Case 3",
			args: args{
				email: "a@gmail.com",
			},
			want: "a@gmail.com",
			wantErr:false,
		},
		{
			name: "Test Case 4",
			args: args{
				email: "a@gmail.com.com",
			},
			want: "a@gmail.com.com",
			wantErr:true,
		},
		{
			name: "Test Case 5",
			args: args{
				email: "a@mail.unpad.ac.id",
			},
			want: "a@mail.unpad.ac.id",
			wantErr:false,
		},
		{
			name: "Test Case 6",
			args: args{
				email: "khaazas",
			},
			want: "khaazas",
			wantErr:true,
		},
		{
			name: "Test Case 7",
			args: args{
				email: "a14@mail.unpad.ac.id",
			},
			want: "a14@mail.unpad.ac.id",
			wantErr:false,
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
