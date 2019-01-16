package util

import (
	"reflect"
	"testing"
)

func TestToInt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"tc-1", args{"1"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.str); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayToInt(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"tc-1", args{[]string{"1"}}, []int{1}},
		{"tc-2", args{[]string{"1", "2"}}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayToInt(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaskCardNumber(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"1", args{"1234567890123456"}, "123456******3456", false},
		{"2", args{"123456789012345"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaskCardNumber(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaskCardNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaskCardNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
