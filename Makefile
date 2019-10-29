install:
	go install
	go install ./tools/kickbuild
	go install ./tools/kicklpc
	go install ./tools/kickpack
	go install ./tools/kickstore
	go install ./tools/rekickwasm

test:
	go test ./pkg/slurp/
	go test ./pkg/project/
	go test ./tools/kickpack/
	go test ./tools/rekickwasm/refactor/

dependencies:
	go get github.com/boltdb/bolt/...
	go get gopkg.in/yaml.v2
	go get github.com/gorilla/websocket
