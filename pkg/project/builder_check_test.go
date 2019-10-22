package project

import (
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestBuilder(t *testing.T) {
	builder := NewBuilder()
	okHomes, err := buildLongOkHomes()
	if err != nil {
		t.Error(err)
	}
	// ok homes
	testBuildFromHomes(t, builder, "ok", okHomes, false)
}

func testBuildFromHomes(t *testing.T, builder *Builder, name string, homes []*Button, wantErr bool) {
	t.Run(name, func(t *testing.T) {
		err := builder.BuildFromHomes(homes)
		if !wantErr && err != nil {
			t.Errorf("Builder.BuildFromHomes() error = %v", err)
		}
		if wantErr && err == nil {
			t.Error("Builder.BuildFromHomes() wanted an error but didn't get one.")
		}
	})
}

func buildShortOkHomes() ([]*Button, error) {
	okHomes := make([]*Button, 0, 5)
	home, err := getHome("testyaml/ok/home1.yaml")
	if err != nil {
		return nil, err
	}
	okHomes = append(okHomes, home)
	home, err = getHome("testyaml/ok/home2.yaml")
	if err != nil {
		return nil, err
	}
	okHomes = append(okHomes, home)
	return okHomes, nil
}

func buildDeepHomes() ([]*Button, error) {
	deepHomes := make([]*Button, 0, 5)
	home, err := getHome("testyaml/ok/home5.yaml")
	if err != nil {
		return nil, err
	}
	deepHomes = append(deepHomes, home)
	return deepHomes, nil
}

func buildLongOkHomes() ([]*Button, error) {
	okHomes := make([]*Button, 0, 5)
	home, err := getHome("testyaml/ok/home1.yaml")
	if err != nil {
		return nil, err
	}
	okHomes = append(okHomes, home)
	home, err = getHome("testyaml/ok/home2.yaml")
	if err != nil {
		return nil, err
	}
	okHomes = append(okHomes, home)
	home, err = getHome("testyaml/ok/home3.yaml")
	if err != nil {
		return nil, err
	}
	okHomes = append(okHomes, home)
	home, err = getHome("testyaml/ok/home4.yaml")
	if err != nil {
		return nil, err
	}
	okHomes = append(okHomes, home)
	return okHomes, nil
}

func getHome(path string) (*Button, error) {
	bb, err := getFileBB(path)
	if err != nil {
		return nil, err
	}
	return getHomeFromBB(bb)
}

func getHomeFromBB(bb []byte) (*Button, error) {
	home := &Button{}
	if err := yaml.Unmarshal(bb, home); err != nil {
		return nil, err
	}
	return home, nil
}

func getFileBB(fpath string) ([]byte, error) {
	ifile, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer func() {
		// close and check for the error
		if cerr := ifile.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	stat, err := ifile.Stat()
	if err != nil {
		return nil, err
	}
	l := int(stat.Size())
	ifilebb := make([]byte, l, l)
	if _, err := ifile.Read(ifilebb); err != nil {
		return nil, err
	}
	return ifilebb, nil
}

func Test_isValidButtonID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "missing button name",
			args:    args{""},
			wantErr: true,
		},
		{
			name:    "numeric button name",
			args:    args{"1 Button"},
			wantErr: true,
		},
		{
			name:    "not cc button name",
			args:    args{"One Button"},
			wantErr: true,
		},
		{
			name: "cc button name",
			args: args{"OneButton"},
		},
		{
			name:    "bad suffix button name",
			args:    args{"OneButtons"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isValidButtonID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("isValidButtonID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isValidTabID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "missing tab name",
			args:    args{""},
			wantErr: true,
		},
		{
			name:    "numeric tab name",
			args:    args{"1 Tab"},
			wantErr: true,
		},
		{
			name:    "not cc tab name",
			args:    args{"One Tab"},
			wantErr: true,
		},
		{
			name: "cc tab name",
			args: args{"OneTab"},
		},
		{
			name:    "bad suffix tab name",
			args:    args{"OneTabs"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isValidTabID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("isValidTabID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
