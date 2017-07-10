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
		{name: "separator is space and first is multi byte character", args: args{s: "今日は 晴れ", sep: " "}, want: "今日は晴れ"},
		{name: "separator is empty", args: args{s: "instantdate", sep: ""}, want: "Instantdate"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpperCamelCase(tt.args.s, tt.args.sep); got != tt.want {
				t.Errorf("ToUpperCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLowerCamelCase(t *testing.T) {
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
		{name: "separator is space and first is multi byte character", args: args{s: "今日は 晴れ", sep: " "}, want: "今日は晴れ"},
		{name: "separator is empty", args: args{s: "Instantdate", sep: ""}, want: "instantdate"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLowerCamelCase(tt.args.s, tt.args.sep); got != tt.want {
				t.Errorf("ToLowerCamelCaes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamelCaseToUnderscoreSeparated(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "separator is space", args: args{s: "instantDate", sep: " "}, want: "instant date"},
		{name: "separator is comma", args: args{s: "instantDate", sep: ","}, want: "instant,date"},
		{name: "separator is alpabet", args: args{s: "instantDate", sep: "a"}, want: "instantadate"},
		{name: "separator is space and string is multi byte character", args: args{s: "今日は晴れ", sep: " "}, want: "今日は晴れ"},
		{name: "separator is empty", args: args{s: "instantDate", sep: ""}, want: "instantDate"},
		{name: "last character is upper case", args: args{s: "instantDateA", sep: "_"}, want: "instant_date_a"},
		{name: "two upper case", args: args{s: "unionInformationMaster", sep: "_"}, want: "union_information_master"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCaseToUnderscoreSeparated(tt.args.s, tt.args.sep); got != tt.want {
				t.Errorf("CamelCaseToUnderscoreSeparated() = %v, want %v", got, tt.want)
			}
		})
	}
}
