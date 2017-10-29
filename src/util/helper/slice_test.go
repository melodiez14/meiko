package helper

import (
	"testing"
)

func TestInt64InSlice(t *testing.T) {
	type args struct {
		val int64
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				val: 100000,
				arr: []int64{
					123, 122, 543, 1234, 76, 41, 324, 4567, 234, 100000,
				},
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				val: 945743534,
				arr: []int64{
					123, 122, 543, 1234, 945743534, 41, 324, 4567, 234,
				},
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				val: 9457435342,
				arr: []int64{
					123, 122, 543, 1234, 945743534, 41, 324, 4567, 234,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64InSlice(tt.args.val, tt.args.arr); got != tt.want {
				t.Errorf("Error in %s, Int64InSlice() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestIsStringInSlice(t *testing.T) {
	type args struct {
		val string
		arr []string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				val: "kamu",
				arr: []string{
					"aku", "siapanya", "bukan", "kita", "berbeda", "tak", "sama", "selesai", "lelah", "kamu",
				},
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				val: "kita",
				arr: []string{
					"aku", "siapanya", "bukan", "kita", "berbeda", "tak", "sama", "selesai", "lelah", "kamu",
				},
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				val: "letih",
				arr: []string{
					"aku", "siapanya", "bukan", "kita", "berbeda", "tak", "sama", "selesai", "lelah", "kamu",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStringInSlice(tt.args.val, tt.args.arr); got != tt.want {
				t.Errorf("Error in %s, IsStringInSlice() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestInt8InSlice(t *testing.T) {
	type args struct {
		val int8
		arr []int8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				val: 127,
				arr: []int8{
					123, 122, 0, -100, 123, 41, 32, 88, 24, 127,
				},
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				val: -20,
				arr: []int8{
					-20, 12, 54, 3, 7, 8, 5, 3, 8, 8, -23,
				},
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				val: 100,
				arr: []int8{
					1, -2, 3, -4, 5, 10, 34, 67, 23,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8InSlice(tt.args.val, tt.args.arr); got != tt.want {
				t.Errorf("Int8InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
