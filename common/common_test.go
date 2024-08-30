package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertNumToLetter(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "a"},
		{2, "b"},
		{26, "z"},
		{27, "a"},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("Convert %d", tt.n)
		t.Run(name, func(t *testing.T) {
			if got := ConvertNumToLetter(tt.n); got != tt.want {
				t.Errorf("ConvertNumToLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateNums(t *testing.T) {

	tests := []struct {
		n    int
		want []int
	}{
		{1, []int{1}},
		{3, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("Generate %d", tt.n)
		t.Run(name, func(t *testing.T) {
			if got := GenerateNums(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
