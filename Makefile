GO=go

test:
	$(GO) test ./...

_clean:
	rm -rfv built/

coverage: _clean
	mkdir built/
	$(GO) test -v -coverprofile ./built/coverage.out ./...
	$(GO) tool cover -html=./built/coverage.out -o ./built/coverage.html
