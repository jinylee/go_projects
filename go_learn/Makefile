# meta
NAME 		:= go_learn
VERSION 	:= $(shell git describe --tag --abbrev=0)
REVISION 	:= $(shell git rev-parse --short HEAD)
LDFLAGS 	:= -X 'main.version=$(VERSION)' \
		   	   -X 'main.revision=$(REVISION)'

# tool setup
setup:
	go get github.com/Masterminds/glide
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help


# run test
test: deps
	go test $$(glide novendor)

# dependency install using glide
deps: setup
	glide install


# update dependency using glide
update: setup
	glide update

# static source analysis using glint
lint: setup
	go vet $$(glide novendor)
	for pkg in $$(glide novendor -x); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done

# format source codes using goimport
fmt : setup
	goimports -w $$(glide nv -x)

# build binaries
bin/%: cmd/%/main.go deps
	go build -ldflags "$(LDFLAGS)" -o $@ $<

# help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: setup deps update test lint help




