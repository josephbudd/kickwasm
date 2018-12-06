package project

import (
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestBuilder(t *testing.T) {
	builder := NewBuilder()
	okServices, err := buildLongOkServices()
	if err != nil {
		t.Error(err)
	}
	// ok services
	testBuildFromServices(t, builder, "ok", okServices, false)
}

func testBuildFromServices(t *testing.T, builder *Builder, name string, services []*Service, wantErr bool) {
	t.Run(name, func(t *testing.T) {
		err := builder.BuildFromServices(services)
		if !wantErr && err != nil {
			t.Errorf("Builder.BuildFromServices() error = %v", err)
		}
		if wantErr && err == nil {
			t.Error("Builder.BuildFromServices() wanted an error but didn't get one.")
		}
	})
}

func buildShortOkServices() ([]*Service, error) {
	okServices := make([]*Service, 0, 5)
	service, err := getService("testyaml/ok/service1.yaml")
	if err != nil {
		return nil, err
	}
	okServices = append(okServices, service)
	service, err = getService("testyaml/ok/service2.yaml")
	if err != nil {
		return nil, err
	}
	okServices = append(okServices, service)
	return okServices, nil
}

func buildDeepServices() ([]*Service, error) {
	deepServices := make([]*Service, 0, 5)
	service, err := getService("testyaml/ok/service5.yaml")
	if err != nil {
		return nil, err
	}
	deepServices = append(deepServices, service)
	return deepServices, nil
}

func buildLongOkServices() ([]*Service, error) {
	okServices := make([]*Service, 0, 5)
	service, err := getService("testyaml/ok/service1.yaml")
	if err != nil {
		return nil, err
	}
	okServices = append(okServices, service)
	service, err = getService("testyaml/ok/service2.yaml")
	if err != nil {
		return nil, err
	}
	okServices = append(okServices, service)
	service, err = getService("testyaml/ok/service3.yaml")
	if err != nil {
		return nil, err
	}
	okServices = append(okServices, service)
	service, err = getService("testyaml/ok/service4.yaml")
	if err != nil {
		return nil, err
	}
	okServices = append(okServices, service)
	return okServices, nil
}

func getService(path string) (*Service, error) {
	bb, err := getFileBB(path)
	if err != nil {
		return nil, err
	}
	return getServiceFromBB(bb)
}

func getServiceFromBB(bb []byte) (*Service, error) {
	service := &Service{}
	if err := yaml.Unmarshal(bb, service); err != nil {
		return nil, err
	}
	return service, nil
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

func Test_isValidServiceName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "missing service name",
			args:    args{""},
			wantErr: true,
		},
		{
			name:    "numeric service name",
			args:    args{"1 Service"},
			wantErr: true,
		},
		{
			name:    "not cc service name",
			args:    args{"One Service"},
			wantErr: true,
		},
		{
			name: "cc service name",
			args: args{"OneService"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isValidServiceName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("isValidServiceName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
