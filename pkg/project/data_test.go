package project

import (
	"reflect"
	"testing"
)

const src = "{\n\"Str\":\"a\nb\"\n}"

func Test_splitBackTicked(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "only",
			args: args{src: src},
			want: []string{"{", "\"Str\":\"a", "b\"", "}"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitBackTicked(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitBackTicked() = %v, want %v", got, tt.want)
			}
		})
	}
}
