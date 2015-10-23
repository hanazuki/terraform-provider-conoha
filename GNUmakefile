export GO15VENDOREXPERIMENT=1
EXE=terraform-provider-conoha

build: ${EXE}

dep:
	gom install

test: build

release:
	gom lock

${EXE}: dep
	go build -o $@

.PHONY: build dep test release
