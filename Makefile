test:
	go test ./pkg/slurp/
	go test ./pkg/project/
	go test ./tools/kickpack/
	go test ./tools/rekickwasm/refactor/

install:
	go install
	go install ./tools/kickbuild
	go install ./tools/kicklpc
	go install ./tools/kickpack
	go install ./tools/kickstore
	go install ./tools/rekickwasm

