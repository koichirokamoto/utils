package utils

import (
	"testing"
)

func TestToUpperCamelCase(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "separator is space and first is to upper", args: args{s: "instant date", sep: " "}, want: "InstantDate"},
		{name: "separator is comma and first is to upper", args: args{s: "instant,date", sep: ","}, want: "InstantDate"},
		{name: "separator is alpabet and first is to upper", args: args{s: "instant date", sep: "a"}, want: "InstNt dTe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpperCamelCase(tt.args.s, tt.args.sep); got != tt.want {
				t.Errorf("ToUpperCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLowerCamelCaes(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "separator is space and first is to lower", args: args{s: "instant date", sep: " "}, want: "instantDate"},
		{name: "separator is comma and first is to lower", args: args{s: "instant,date", sep: ","}, want: "instantDate"},
		{name: "separator is alpabet and first is to lower", args: args{s: "instant date", sep: "a"}, want: "instNt dTe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLowerCamelCaes(tt.args.s, tt.args.sep); got != tt.want {
				t.Errorf("ToLowerCamelCaes() = %v, want %v", got, tt.want)
			}
		})
	}
}
