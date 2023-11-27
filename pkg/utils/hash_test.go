package utils

import (
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Test_Plain_text_one_word",
			s:    "furkan",
			want: "nakruf",
		},
		{
			name: "Test_Plain_text_multiple_word",
			s:    "furkan topaloglu",
			want: "ulgolapot nakruf",
		},
		{
			name: "Test_Plain_text_with_number",
			s:    "furkan 1994",
			want: "4991 nakruf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		s    string
		want string
	}{
		{

			"TestHash_Basic",
			"password123",
			"6e2ca4dceb77e35d3998830f23783b07b57901b98fccbf52b9c3bc88ca02f49d",
		},
		{

			"TestHash_DifferentLengthInput",
			"abc",
			"a41c91a071a166e02052c6e05d843e2a103c757cc45ace16058168079a5eb156",
		},
		{

			"TestHash_SpecialCharactersInput",
			"pass!@#",
			"90dbde68d9d73673f83b31c1a5765d8ee8b8f9eecbc14ce2bced1c4d02231b2f",
		},
		{

			"TestHash_EmptyInput",
			"",
			"cd372fb85148700fa88095e3492d3f9f5beb43e555e5ff26d95f5a6adc36f8e6",
		},
		{

			"TestHash_CaseSensitiveInput",
			"Password123",
			"93566b7784f8eb4b31aca710e4761548e028cbdd00351424ecf4b57557df2e44",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash(tt.s); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
