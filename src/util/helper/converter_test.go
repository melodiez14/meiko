package helper

import (
	"reflect"
	"testing"
	"time"
)

func TestStringToMD5(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				text: "allcharacter",
			},
			want: "5ed590569350a5ef57283ac2eb1754cc",
		},
		{
			name: "Test Case 2",
			args: args{
				text: "12345678",
			},
			want: "25d55ad283aa400af464c76d713c07ad",
		},
		{
			name: "Test Case 3",
			args: args{
				text: "1234abcd",
			},
			want: "ef73781effc5774100f87fe2f437a435",
		},
		{
			name: "Test Case 4",
			args: args{
				text: "!@#$abcd",
			},
			want: "2a66861a71be424de0f85eaa8f305a3b",
		},
		{
			name: "Test Case 5",
			args: args{
				text: "abcd!@#$",
			},
			want: "89a9caa3158431a0215a9598de3274f2",
		},
		{
			name: "Test Case 6",
			args: args{
				text: "%^&*1234",
			},
			want: "48d3084db3a5053f1baec2f560e72077",
		},
		{
			name: "Test Case 7",
			args: args{
				text: "%^&*!@#$",
			},
			want: "1d0ee682abc98c1a3c3c6283e7e602af",
		},
		{
			name: "Test Case 8",
			args: args{
				text: "%^&*!@#$fkdsafjksf3129382932839kfdss",
			},
			want: "ec68bff5b6a22f72ae1915238e320434",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToMD5(tt.args.text); got != tt.want {
				t.Errorf("StringToMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractExtension(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				fileName: "abcdefgh.pdf",
			},
			want:    "abcdefgh",
			want1:   "pdf",
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				fileName: "12345678.pdf",
			},
			want:    "12345678",
			want1:   "pdf",
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				fileName: "1234abcd.pdf",
			},
			want:    "1234abcd",
			want1:   "pdf",
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				fileName: "abcd1234.pdf",
			},
			want:    "abcd1234",
			want1:   "pdf",
			wantErr: false,
		},
		{
			name: "Test case 5",
			args: args{
				fileName: "!@#$%^&*.pdf",
			},
			want:    "!@#$%^&*",
			want1:   "pdf",
			wantErr: false,
		},
		{
			name: "Test case 6",
			args: args{
				fileName: "abcd!@#$.pdf",
			},
			want:    "abcd!@#$",
			want1:   "pdf",
			wantErr: false,
		},
		{
			name: "Test case 7",
			args: args{
				fileName: "..........pdf",
			},
			want:    ".........",
			want1:   "pdf",
			wantErr: false,
		},
		{
			name: "Test case 8",
			args: args{
				fileName: "!@#$abcd.pdf",
			},
			want:    "!@#$abcd",
			want1:   "pdf",
			wantErr: false,
		},
		{
			name: "Test case 9",
			args: args{
				fileName: "pdf.pdf",
			},
			want:    "pdf",
			want1:   "pdf",
			wantErr: false,
		},
		{
			name: "Test case 10",
			args: args{
				fileName: "abc.exe.pdf",
			},
			want:    "abc.exe",
			want1:   "pdf",
			wantErr: false,
		},
		{
			name: "Test case 11",
			args: args{
				fileName: "abc.exe.pdf.",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "Test case 12",
			args: args{
				fileName: "abc.exe.pdf..",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "Test case 13",
			args: args{
				fileName: "abc.exe.pdf.ex?x",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "Test case 14",
			args: args{
				fileName: "ab.pdf!@#$%^&**.@&*&*@&",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "Test case 15",
			args: args{
				fileName: "abpdf!@#$%^&**@&*&*@&",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "Test case 16",
			args: args{
				fileName: "abpdf!@#$%^&**@&*&*@&",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "Test case 17",
			args: args{
				fileName: " .pdf",
			},
			want:    " ",
			want1:   "pdf",
			wantErr: false,
		},
		{
			name: "Test case 17",
			args: args{
				fileName: " .  ",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "Test case 18",
			args: args{
				fileName: " .pdf",
			},
			want:    " ",
			want1:   "pdf",
			wantErr: false,
		},
		{
			name: "Test case 19",
			args: args{
				fileName: "abc.a pdf",
			},
			want:    "abc",
			want1:   "apdf",
			wantErr: false,
		},
		{
			name: "Test case 20",
			args: args{
				fileName: "abc.a pdf pdf pdf",
			},
			want:    "abc",
			want1:   "apdfpdfpdf",
			wantErr: false,
		},
		{
			name: "Test case 21",
			args: args{
				fileName: "abc.p!@#! df",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ExtractExtension(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractExtension() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractExtension() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExtractExtension() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDateToString(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 0, 0, 59000000000, time.UTC),
			},
			want: "Just now",
		},
		{
			name: "Test case 2",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 0, 0, 60000000000, time.UTC),
			},
			want: "1 minutes ago",
		},
		{
			name: "Test case 3",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 0, 0, 3600000000000, time.UTC),
			},
			want: "1 hours ago",
		},
		{
			name: "Test case 4",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 0, 0, 86400000000000, time.UTC),
			},
			want: "1 days ago",
		},
		{
			name: "Test case 5",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 0, 0, 259300000000000, time.UTC),
			},
			want: "3 days ago",
		},
		{
			name: "Test case 6",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 0, 0, 345600000000000, time.UTC),
			},
			want: "Monday, 10 October 2017",
		},
		{
			name: "Test case 7",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 0, 59, 0, time.UTC),
			},
			want: "Just now",
		},
		{
			name: "Test case 8",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 0, 60, 0, time.UTC),
			},
			want: "1 minutes ago",
		},
		{
			name: "Test case 9",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 0, 3600, 0, time.UTC),
			},
			want: "1 hours ago",
		},
		{
			name: "Test case 10",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 0, 86400, 0, time.UTC),
			},
			want: "1 days ago",
		},
		{
			name: "Test case 11",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 0, 345600, 0, time.UTC),
			},
			want: "Monday, 10 October 2017",
		},
		{
			name: "Test case 12",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, -1, 0, 0, time.UTC),
			},
			want: "Just now",
		},
		{
			name: "Test case 12",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 59, 0, 0, time.UTC),
			},
			want: "59 minutes ago",
		},
		{
			name: "Test case 13",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 60, 0, 0, time.UTC),
			},
			want: "1 hours ago",
		},
		{
			name: "Test case 14",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 1440, 0, 0, time.UTC),
			},
			want: "1 days ago",
		},
		{
			name: "Test case 15",
			args: args{
				t1: time.Date(2017, time.October, 5, 0, 0, 0, 0, time.UTC),
				t2: time.Date(2017, time.October, 5, 0, 5760, 0, 0, time.UTC),
			},
			want: "Monday, 10 October 2017",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateToString(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("DateToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntDayToString(t *testing.T) {
	type args struct {
		day int8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Day -1",
			args: args{
				day: -1,
			},
			want: "",
		},
		{
			name: "Sunday",
			args: args{
				day: 0,
			},
			want: "Sunday",
		},
		{
			name: "Test case 2",
			args: args{
				day: 1,
			},
			want: "Monday",
		},
		{
			name: "Test case 3",
			args: args{
				day: 2,
			},
			want: "Tuesday",
		},
		{
			name: "Test case 4",
			args: args{
				day: 3,
			},
			want: "Wednesday",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntDayToString(tt.args.day); got != tt.want {
				t.Errorf("IntDayToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinutesToTimeString(t *testing.T) {
	type args struct {
		minutes uint16
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				minutes: 1441,
			},
			want: "",
		},
		{
			name: "Test case 2",
			args: args{
				minutes: 0,
			},
			want: "00:00",
		},
		{
			name: "Test case 3",
			args: args{
				minutes: 1440,
			},
			want: "24:00",
		},
		{
			name: "Test case 4",
			args: args{
				minutes: 200,
			},
			want: "03:20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinutesToTimeString(tt.args.minutes); got != tt.want {
				t.Errorf("MinutesToTimeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDayStringToInt(t *testing.T) {
	type args struct {
		day string
	}
	tests := []struct {
		name    string
		args    args
		want    int8
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				day: "monday",
			},
			want:    int8(time.Monday),
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				day: "tuesday",
			},
			want:    int8(time.Tuesday),
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				day: "wednesday",
			},
			want:    int8(time.Wednesday),
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				day: "thursday",
			},
			want:    int8(time.Thursday),
			wantErr: false,
		},
		{
			name: "Test case 5",
			args: args{
				day: "friday",
			},
			want:    int8(time.Friday),
			wantErr: false,
		},
		{
			name: "Test case 6",
			args: args{
				day: "saturday",
			},
			want:    int8(time.Saturday),
			wantErr: false,
		},
		{
			name: "Test case 7",
			args: args{
				day: "sunday",
			},
			want:    int8(time.Sunday),
			wantErr: false,
		},
		{
			name: "Test case 8",
			args: args{
				day: "sundayyyy",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Test case 8",
			args: args{
				day: "123213213",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DayStringToInt(tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("DayStringToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DayStringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ToStringSlice(t *testing.T) {
	type args struct {
		value []int64
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty slice",
			args: args{},
		},
		{
			name: "Filled slice",
			args: args{
				value: []int64{123, 98123987123, 3232, -1230123},
			},
			want: []string{"123", "98123987123", "3232", "-1230123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64ToStringSlice(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64ToStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeToDayInt(t *testing.T) {
	type args struct {
		time []time.Time
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "27-29 October",
			args: args{
				time: []time.Time{
					time.Date(2017, time.October, 27, 0, 0, 0, 0, time.Local),
					time.Date(2017, time.October, 28, 0, 0, 0, 0, time.Local),
					time.Date(2017, time.October, 29, 0, 0, 0, 0, time.Local),
				},
			},
			want: []int8{5, 6, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeToDayInt(tt.args.time...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeToDayInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Round(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1.1234",
			args: args{
				value: 1.1234,
			},
			want: 1,
		},
		{
			name: "1.5",
			args: args{
				value: 1.51,
			},
			want: 2,
		},
		{
			name: "1.9",
			args: args{
				value: 1.9,
			},
			want: 2,
		},
		{
			name: "-102.234",
			args: args{
				value: -102.234,
			},
			want: -102,
		},
		{
			name: "-102.234",
			args: args{
				value: -102.534,
			},
			want: -103,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64Round(tt.args.value); got != tt.want {
				t.Errorf("Float64Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Round(t *testing.T) {
	type args struct {
		value float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "1.1234",
			args: args{
				value: 1.1234,
			},
			want: 1,
		},
		{
			name: "1.5",
			args: args{
				value: 1.51,
			},
			want: 2,
		},
		{
			name: "1.9",
			args: args{
				value: 1.9,
			},
			want: 2,
		},
		{
			name: "-102.234",
			args: args{
				value: -102.234,
			},
			want: -102,
		},
		{
			name: "-102.234",
			args: args{
				value: -102.534,
			},
			want: -103,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32Round(tt.args.value); got != tt.want {
				t.Errorf("Float32Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
