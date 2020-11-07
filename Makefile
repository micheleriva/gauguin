WEBAPP=editor-webapp
BINARY_NAME=gauguin
BINARY_NAME_UNIX=$(BINARY_NAME)_unix

build-react:
	cd editor/$(WEBAPP) && yarn && yarn build
	cp -R editor/$(WEBAPP)/build/static public/$(WEBAPP)/static

build:
	go build -o $(BINARY_UNIX) -v	

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_UNIX) -v	