package format

import (
	"reflect"
	"testing"
)

func TestSameWidth(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name      string
		args      args
		wantSized []string
	}{
		{
			name:      "same width",
			args:      args{src: []string{"a", "bb", "ccc"}},
			wantSized: []string{"a  ", "bb ", "ccc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSized := SameWidth(tt.args.src); !reflect.DeepEqual(gotSized, tt.wantSized) {
				t.Errorf("SameWidth() = %v, want %v", gotSized, tt.wantSized)
			}
		})
	}
}
