package script

import (
	"reflect"
	"testing"
)

func TestRunDump(t *testing.T) {
	type args struct {
		env     []string
		dir     string
		command string
		args    []string
	}
	tests := []struct {
		name     string
		args     args
		wantDump []string
		wantErr  bool
	}{
		{
			// TODO: Add test cases.
			name: "ls Test",
			args: args{
				env:     nil,
				dir:     "./data/",
				command: "ls",
				args:    make([]string, 0),
			},
			wantDump: []string{"file1.txt", "file2.txt", ""},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDump, err := RunDump(tt.args.env, tt.args.dir, tt.args.command, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunDump() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDump, tt.wantDump) {
				t.Errorf("RunDump() = %#v, want %#v", gotDump, tt.wantDump)
			}
		})
	}
}

func TestRun(t *testing.T) {
	type args struct {
		command string
		args    []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Run test.",
			args: args{
				command: "gedit",
				args:    []string{"./script_test.go"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Run(tt.args.command, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
