package cases

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CamelCase: lower separated",
			args: args{s: "abc def ghi"},
			want: "AbcDefGhi",
		},
		{
			name: "CamelCase: upper separated",
			args: args{s: "ABC DEF GHI"},
			want: "ABCDEFGHI",
		},
		{
			name: "CamelCase: lower",
			args: args{s: "abcDefGhi"},
			want: "AbcDefGhi",
		},
		{
			name: "CamelCase: upper",
			args: args{s: "AbcDefGhi"},
			want: "AbcDefGhi",
		},
		{
			name: "CamelCase: digit",
			args: args{s: "1Ab2cDefGhi"},
			want: "_1Ab2cDefGhi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.args.s); got != tt.want {
				t.Errorf("CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "LowerCamelCase: lower separated",
			args: args{s: "abc def ghi"},
			want: "abcDefGhi",
		},
		{
			name: "LowerCamelCase: upper separated",
			args: args{s: "ABC DEF GHI"},
			want: "aBCDEFGHI",
		},
		{
			name: "LowerCamelCase: lower",
			args: args{s: "abcDefGhi"},
			want: "abcDefGhi",
		},
		{
			name: "LowerCamelCase: upper",
			args: args{s: "AbcDefGhi"},
			want: "abcDefGhi",
		},
		{
			name: "LowerCamelCase: digit",
			args: args{s: "1Ab2cDefGhi"},
			want: "_1Ab2cDefGhi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerCamelCase(tt.args.s); got != tt.want {
				t.Errorf("LowerCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToGoPackageName(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name        string
		args        args
		wantNewName string
	}{
		{
			name:        "ToGoPackageName: lower separated",
			args:        args{s: "abc def ghi"},
			wantNewName: "abcdefghi",
		},
		{
			name:        "ToGoPackageName: upper separated",
			args:        args{s: "ABC DEF GHI"},
			wantNewName: "abcdefghi",
		},
		{
			name:        "ToGoPackageName: lower",
			args:        args{s: "abcDefGhi"},
			wantNewName: "abcdefghi",
		},
		{
			name:        "ToGoPackageName: upper",
			args:        args{s: "AbcDefGhi"},
			wantNewName: "abcdefghi",
		},
		{
			name:        "ToGoPackageName: digit",
			args:        args{s: "1Ab2cDefGhi"},
			wantNewName: "_1ab2cdefghi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewName := ToGoPackageName(tt.args.s); gotNewName != tt.wantNewName {
				t.Errorf("ToGoPackageName() = %v, want %v", gotNewName, tt.wantNewName)
			}
		})
	}
}
